package submitter

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"github.com/metis-devops/metis-reward-submitter/internal/themis"
	"github.com/metis-devops/metis-reward-submitter/internal/utils"
	"github.com/prometheus/client_golang/prometheus"
)

func (s *Submitter) newBatch(basectx context.Context) (bool, error) {
	newctx, cancel := context.WithTimeout(basectx, time.Minute*3)
	defer cancel()

	// get current batch and epoch info
	state, err := s.QueryChainState(newctx)
	if err != nil {
		return false, fmt.Errorf("query chain state: %w", err)
	}

	if state == nil {
		return false, nil
	}

	if !state.ShouldSubmit(s.params.MinEpochLen, s.params.MinInterval) {
		slog.Debug("No need to submmit new batch",
			"finalizedEpoch", state.FinalizedEpoch,
			"lastBatchId", state.LastBatchId,
			"lastBatchEnd", state.LastBatchEndEpoch,
			"lastBatchTime", state.LastBatchTime,
		)
		return false, nil
	}

	// todo: compare nonce in local cache if it's available

	// get mpc info and check if the mpc address is initlized on chain
	// if the mpc address has not created, the LatestMpcInfo will return a not found error
	mpcInfo, err := s.ThemisClient.LatestMpcInfo(newctx, themis.RewardSubmitMpcType)
	if err != nil {
		return false, err
	}
	if !state.IsMpcInited(mpcInfo.Address) {
		return false, fmt.Errorf("mpc address is not initlized expected=%s got=%s", mpcInfo.Address, state.MpcAddress)
	}

	// check if the payer is setted
	if !state.IsPayerInited() {
		slog.Warn("payer is not initlized")
		return false, nil
	}

	slog.Info("Batch proposal",
		"finalizedEpoch", state.FinalizedEpoch,
		"lastBatchId", state.LastBatchId,
		"lastBatchEnd", state.LastBatchEndEpoch,
		"lastBatchTime", state.LastBatchTime)

	newBatchId, newStartEpoch, newEndEpoch := state.LastBatchId+1, state.LastBatchEndEpoch+1, state.FinalizedEpoch

	slog.Debug("batch input", "id", newBatchId, "start", newStartEpoch, "end", newEndEpoch)
	// build batch info
	batchInfo, batchBlocks, err := s.QueryAndBuildBatchInput(newctx, newBatchId, newStartEpoch, newEndEpoch)
	if err != nil {
		return false, fmt.Errorf("failed to get epoch info for batch %d: %w", newBatchId, err)
	}

	slog.Debug("Batch", "data", utils.JsonString(batchInfo), "blocks", batchBlocks)
	// calculate total totalReward and check if the payer has enough Metis to pay
	totalReward := new(big.Int).Mul(batchBlocks, state.RewardPerBlock)
	if !state.PayerHasSufficientBalance(totalReward) {
		s.Metric.Insufficience.With(prometheus.Labels{"alias": "payer", "address": state.PayerAddress.Hex(), "type": "balance"}).Set(1)
		slog.Error("balance of payer is insufficient to pay the reward", "payer", state.PayerAddress, "amount", utils.ToEther(totalReward), "balance", utils.ToEther(state.PayerBalance))
		return false, nil
	} else {
		s.Metric.Insufficience.With(prometheus.Labels{"alias": "payer", "address": state.PayerAddress.Hex(), "type": "balance"}).Set(0)
	}

	if !state.PayerHasSufficientAllowlance(totalReward) {
		s.Metric.Insufficience.With(prometheus.Labels{"alias": "payer", "address": state.PayerAddress.Hex(), "type": "allowlance"}).Set(1)
		slog.Error("allowlance of payer is insufficient to pay the reward", "payer", state.PayerAddress, "amount", utils.ToEther(totalReward), "allowlance", utils.ToEther(state.PayerAllowance))
		return false, nil
	} else {
		s.Metric.Insufficience.With(prometheus.Labels{"alias": "payer", "address": state.PayerAddress.Hex(), "type": "allowlance"}).Set(0)
	}

	// calculate tx input
	txInput, err := batchInfo.Encode()
	if err != nil {
		return false, fmt.Errorf("failed to encode batch tx input")
	}

	// get gas price params
	gasPriceTip, gasPriceCap, err := s.GetGasParams(newctx)
	if err != nil {
		return false, fmt.Errorf("failed to get gas params")
	}

	if gasPriceCap.Cmp(s.params.GasPriceLimit) > 0 {
		slog.Warn("gas price exceed limit", "estimate", utils.ToGWei(gasPriceCap), "limit", utils.ToGWei(s.params.GasPriceLimit))
		return false, nil
	}

	// get gas limit
	accessList, gasLimit, err := s.GetAccessList(newctx,
		state.MpcAddress, s.params.LockingPoolAddress, txInput)
	if err != nil {
		return false, err
	}
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:    s.params.EthChainId,
		Nonce:      state.MpcNonce,
		GasTipCap:  gasPriceTip,
		GasFeeCap:  gasPriceCap,
		Gas:        gasLimit,
		To:         &s.params.LockingPoolAddress,
		Value:      new(big.Int),
		AccessList: *accessList,
		Data:       txInput,
	})

	slog.Debug("Batch tx info", "nonce", state.MpcNonce, "gasTipCap", gasPriceTip, "gasFeeCap",
		gasPriceCap, "gasLimit", gasLimit)

	gasFee := new(big.Int).Mul(tx.GasPrice(), new(big.Int).SetUint64(gasLimit))
	if gasFee.Cmp(state.MpcBalance) > 0 {
		s.Metric.Insufficience.With(prometheus.Labels{"alias": "mpc", "address": state.MpcAddress.Hex(), "type": "balance"}).Set(1)
		return false, fmt.Errorf("mpc eth balance %f is not enough to pay gas fee %f",
			utils.ToEther(state.MpcBalance), utils.ToEther(gasFee))
	} else {
		s.Metric.Insufficience.With(prometheus.Labels{"alias": "mpc", "address": state.MpcAddress.Hex(), "type": "balance"}).Set(0)
	}

	// request signature from mpc
	rawTx, err := tx.MarshalBinary()
	if err != nil {
		return false, fmt.Errorf("failed to encode raw tx: %w", err)
	}

	signId, err := uuid.NewRandom()
	if err != nil {
		return false, fmt.Errorf("failed to generate uuid: %w", err)
	}

	mpcReq := &themis.SignByMpcReq{
		SignID:   signId.String(),
		MpcID:    mpcInfo.Id,
		SignType: themis.BatchRewardSignType,
		SignData: hexutil.Encode(rawTx),
		SignMsg:  s.params.TxHasher.Hash(tx).Hex(),
	}
	signReqTxid, err := s.ThemisClient.InitSign(newctx, mpcReq)
	if err != nil {
		return false, err
	}
	slog.Debug("Batch sign request", "themisTxid", signReqTxid, "batch", newBatchId,
		"signId", mpcReq.SignID, "signType", mpcReq.SignType, "mpcId", mpcInfo.Id, "raw", mpcReq.SignData, "msg", mpcReq.SignMsg)

	s.state = &State{
		Tx:         tx,
		SignId:     mpcReq.SignID,
		BatchId:    newBatchId,
		StartEpoch: newStartEpoch,
		EndEpoch:   newEndEpoch,
		Status:     StatusUnsigned,
		UpdatedAt:  time.Now(),
	}
	if err := s.saveState(); err != nil {
		return false, fmt.Errorf("failed to save state: %w", err)
	}
	slog.Info("New Batch", "signId", mpcReq.SignID, "batch", newBatchId,
		"startEpoch", newStartEpoch, "endEpoch", newEndEpoch, "reward", utils.ToEther(totalReward))
	return true, nil
}

func (s *Submitter) withSignature(basectx context.Context) (bool, error) {
	newctx, cancel := context.WithTimeout(basectx, time.Second*15)
	defer cancel()

	if time.Since(s.state.UpdatedAt) > time.Minute*3 {
		slog.Warn("Discard due to no new signatgure for long", "signId", s.state.SignId, "batch", s.state.BatchId)
		s.Metric.MpcErrs.Inc()

		s.state.Status = StatusIdle
		s.state.UpdatedAt = time.Now()
		if err := s.saveState(); err != nil {
			return false, fmt.Errorf("failed to save state: %w", err)
		}
		return false, nil
	}

	slog.Info("Try to add signature", "signId", s.state.SignId, "batch", s.state.BatchId)
	res, err := s.ThemisClient.GetMpcSign(newctx, s.state.SignId)
	if err != nil {
		var notfound themis.ClientError
		if errors.As(err, &notfound) && notfound.NotFound() {
			slog.Warn("The sign request not found, wait for next polling", "id", s.state.SignId)
			return false, nil
		}
		return false, err
	}

	// not ready
	if len(res.Signature) == 0 {
		slog.Warn("No signature, wait for next polling", "id", s.state.SignId)
		return false, nil
	}

	slog.Debug("GetMpcSign", "res", utils.JsonString(res))
	if len(res.Signature) != 65 {
		s.Metric.MpcErrs.Inc()
		return false, fmt.Errorf("invalid signature: %x", res.Signature)
	}

	signed, err := s.state.Tx.WithSignature(s.params.TxHasher, res.Signature)
	if err != nil {
		s.Metric.MpcErrs.Inc()
		return false, fmt.Errorf("failed to add signature to the tx: %w", err)
	}

	s.state.Tx = signed
	s.state.UpdatedAt = time.Now()
	s.state.Status = StatusSigned
	if err := s.saveState(); err != nil {
		return false, fmt.Errorf("failed to save state: %w", err)
	}
	txHash := signed.Hash()
	slog.Info("New signed tx generated", "signId", s.state.SignId, "batch",
		s.state.BatchId, "tx", txHash)
	return true, nil
}

func (s *Submitter) submitTx(basectx context.Context) (bool, error) {
	newctx, cancel := context.WithTimeout(basectx, time.Second*15)
	defer cancel()

	txHash := s.state.Tx.Hash()
	slog.Info("Sending tx", "signId", s.state.SignId, "tx", txHash)
	if err := s.EthClient.SendTransaction(newctx, s.state.Tx); err != nil {
		return false, fmt.Errorf("failed to send transaction: %w", err)
	}

	s.state.UpdatedAt = time.Now()
	s.state.Status = StatusSubmitted
	if err := s.saveState(); err != nil {
		return false, fmt.Errorf("failed to save state: %w", err)
	}
	slog.Info("Send tx successful", "signId", s.state.SignId, "batch",
		s.state.BatchId, "tx", txHash)
	return true, nil
}

func (s *Submitter) checkTx(basectx context.Context) (bool, error) {
	newctx, cancel := context.WithTimeout(basectx, time.Second*15)
	defer cancel()

	txHash := s.state.Tx.Hash()
	slog.Info("Check tx confirmation",
		"batchId", s.state.BatchId, "signId", s.state.SignId, "tx", txHash)

	receipt, err := s.EthClient.TransactionReceipt(newctx, txHash)
	if err == nil && receipt != nil {
		latest, err := s.EthClient.HeaderByNumber(newctx, nil)
		if err != nil {
			return false, fmt.Errorf("failed to get latest block: %w", err)
		}

		confirmationNumber := new(big.Int).Sub(latest.Number, receipt.BlockNumber).Int64() + 1
		slog.Info("confiirmation", "block", latest.Number, "number", confirmationNumber)
		// 6 confirmations number at least
		if confirmationNumber > 5 {
			s.state.UpdatedAt = time.Now()
			s.state.Status = StatusConfirmed
			if err := s.saveState(); err != nil {
				return false, fmt.Errorf("failed to save state: %w", err)
			}

			slog.Info("Batch confirmed", "batch", s.state.BatchId,
				"startEpoch", s.state.StartEpoch, "endEpoch", s.state.EndEpoch, "tx", txHash)
			return true, nil
		}

		return false, nil
	}
	// todo: rbf
	if err == ethereum.NotFound {
		// try to resend the tx
		if time.Since(s.state.UpdatedAt) > time.Minute {
			_ = s.EthClient.SendTransaction(newctx, s.state.Tx)
			s.state.UpdatedAt = time.Now()
		}
		return false, nil
	}

	return false, err
}

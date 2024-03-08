package submitter

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/metisprotocol/metis-seq-reward-submitter/internal/utils"
	"golang.org/x/sync/errgroup"
)

func (s *Submitter) QueryChainState(basectx context.Context) (*ChainState, error) {
	newctx, cancle := context.WithTimeout(basectx, time.Second*10)
	defer cancle()

	opts := &bind.CallOpts{Context: newctx}

	finallized, err := s.SequencerSet.FinalizedEpoch(opts)
	if err != nil {
		return nil, err
	}

	if !finallized.Exist {
		return nil, fmt.Errorf("No finallized epoch found")
	}

	batch, err := s.LockingPool.CurBatchState(opts)
	if err != nil {
		return nil, err
	}

	rewardPerBlock, err := s.LockingPool.BLOCKREWARD(opts)
	if err != nil {
		return nil, err
	}

	block, err := s.EthClient.HeaderByNumber(basectx, batch.Number)
	if err != nil {
		return nil, err
	}

	payer, err := s.LockingInfo.RewardPayer(opts)
	if err != nil {
		return nil, err
	}

	mpcAddress, err := s.LockingPool.MpcAddress(opts)
	if err != nil {
		return nil, err
	}

	var mpcBalance = new(big.Int)
	var mpcNonce uint64
	if !utils.IsZeroAddress(mpcAddress) {
		mpcBalance, err = s.EthClient.BalanceAt(opts.Context, mpcAddress, nil)
		if err != nil {
			return nil, err
		}
		mpcNonce, err = s.EthClient.PendingNonceAt(opts.Context, mpcAddress)
		if err != nil {
			return nil, err
		}
	}

	var payerBalance = new(big.Int)
	var payerAllowance = new(big.Int)

	if !utils.IsZeroAddress(payer) {
		payerBalance, err = s.MetisToken.BalanceOf(opts, payer)
		if err != nil {
			return nil, err
		}
		payerAllowance, err = s.MetisToken.Allowance(opts, payer, s.params.LockcingInfoAddress)
		if err != nil {
			return nil, err
		}
	}

	return &ChainState{
		FinalizedEpoch:    finallized.Epoch.Number.Uint64(),
		LastBatchId:       batch.Id.Uint64(),
		LastBatchEndEpoch: batch.EndEpoch.Uint64(),
		LastBatchTime:     time.Unix(int64(block.Time), 0),
		MpcAddress:        mpcAddress,
		MpcBalance:        mpcBalance,
		MpcNonce:          mpcNonce,
		PayerAddress:      payer,
		PayerBalance:      payerBalance,
		PayerAllowance:    payerAllowance,
		RewardPerBlock:    rewardPerBlock,
	}, nil
}

func (s *Submitter) QueryAndBuildBatchInput(basectx context.Context, batchId, startEpoch, endEpoch uint64) (*TxInput, *big.Int, error) {
	newctx, canel := context.WithTimeout(basectx, time.Minute*5)
	defer canel()

	var lock sync.Mutex
	var blocks = make(map[common.Address]uint64)

	var blockCount uint64 = 0

	eg, egctx := errgroup.WithContext(newctx)
	opt := &bind.CallOpts{Context: egctx}
	for i := startEpoch; i <= endEpoch; i++ {
		i := i
		eg.Go(func() error {
			epoch, err := s.SequencerSet.Epochs(opt, new(big.Int).SetUint64(i))
			if err != nil {
				return err
			}
			span, err := s.ThemisClient.GetEpochByID(egctx, int64(i))
			if err != nil {
				return err
			}

			if len(span.Producers) != 1 {
				return fmt.Errorf("span %d doesn't have a producer", i)
			}

			if restSigner := span.Producers[0].Signer; restSigner.Cmp(epoch.Signer) != 0 {
				return fmt.Errorf("signer of the span %d mismatch l2=%s themis=%s", i, epoch.Signer, restSigner)
			}

			if startBlock := epoch.StartBlock.Uint64(); startBlock != span.StartBlock {
				return fmt.Errorf("startBlock of the span %d mismatch l2=%d themis=%d", i, startBlock, span.StartBlock)
			}

			if endBlock := epoch.EndBlock.Uint64(); endBlock != span.EndBlock {
				return fmt.Errorf("endBlock of the span %d mismatch l2=%d themis=%d", i, endBlock, span.EndBlock)
			}

			lock.Lock()
			defer lock.Unlock()

			number := epoch.EndBlock.Uint64() - epoch.StartBlock.Uint64() + 1
			blocks[epoch.Signer] += number
			blockCount += number
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, nil, err
	}

	var input = &TxInput{
		BatchId:    new(big.Int).SetUint64(batchId),
		StartEpoch: new(big.Int).SetUint64(startEpoch),
		EndEpoch:   new(big.Int).SetUint64(endEpoch),
		Seqs:       make([]common.Address, 0, len(blocks)),
		Blocks:     make([]*big.Int, 0, len(blocks)),
	}

	for addr, count := range blocks {
		input.Seqs = append(input.Seqs, addr)
		input.Blocks = append(input.Blocks, new(big.Int).SetUint64(count))
	}

	return input, new(big.Int).SetUint64(blockCount), nil
}

func (s *Submitter) GetGasParams(basectx context.Context) (tip, cap *big.Int, err error) {
	newctx, canel := context.WithTimeout(basectx, time.Second*5)
	defer canel()

	header, err := s.EthClient.HeaderByNumber(newctx, nil)
	if err != nil {
		return nil, nil, err
	}

	tip, err = s.EthClient.SuggestGasTipCap(newctx)
	if err != nil {
		return nil, nil, err
	}

	cap = new(big.Int).Mul(header.BaseFee, common.Big2)
	cap.Add(cap, tip)
	return
}

func (s *Submitter) GetAccessList(basectx context.Context, from, to common.Address, input []byte) (*types.AccessList, uint64, error) {
	newctx, canel := context.WithTimeout(basectx, time.Second*5)
	defer canel()

	callmsg := ethereum.CallMsg{
		From: from,
		To:   &to,
		Data: input,
	}
	geth := gethclient.New(s.EthClient.Client())
	// ignore the error from CreateAccessList due to the AccessList is not required and not important
	list, _, _, err := geth.CreateAccessList(newctx, callmsg)
	if err != nil {
		list = &types.AccessList{}
	}
	gas, err := s.EthClient.EstimateGas(newctx, callmsg)
	if err != nil {
		return nil, 0, err
	}
	// Add extra 21000 to prevent the out of gas error
	return list, gas + 21000, err
}

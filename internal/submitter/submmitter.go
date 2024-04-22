package submitter

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/metisprotocol/metis-reward-submitter/internal/config"
	"github.com/metisprotocol/metis-reward-submitter/internal/contracts/erc20"
	"github.com/metisprotocol/metis-reward-submitter/internal/contracts/lockinginfo"
	"github.com/metisprotocol/metis-reward-submitter/internal/contracts/lockingpool"
	"github.com/metisprotocol/metis-reward-submitter/internal/contracts/sequencerset"
	"github.com/metisprotocol/metis-reward-submitter/internal/themis"
)

/*

process

0. if we have a pending tx in the local cache, skip
1. get current batch state on Eth
2. get current finallized epoch on Metis chain
3. compare the differences and check if the new reward batch should be submitted
4. build the transaction and request signature from themis and mpc
5. wait for the signature and put it to the transaction
6. send the transaction to Eth chain
7. wait for confirmation of the transaction

*/

type Submitter struct {
	EthClient    *ethclient.Client
	MetisClient  *ethclient.Client
	ThemisClient *themis.Client
	Metric       *Metrics

	// contract instances
	SequencerSet *sequencerset.Sequencerset
	LockingPool  *lockingpool.Lockingpool
	LockingInfo  *lockinginfo.Lockinginfo
	MetisToken   *erc20.Erc20

	params *Params
	state  *State
}

type Params struct {
	StatePath     string
	MinEpochLen   uint64
	MinInterval   time.Duration
	GasPriceLimit *big.Int
	TxHasher      types.Signer

	EthChainId          *big.Int
	MetisChainId        *big.Int
	LockingPoolAddress  common.Address
	LockcingInfoAddress common.Address
	MetisTokenAddress   common.Address
	SeqSetAddress       common.Address
}

func New(basectx context.Context, metric *Metrics, conf *config.Config) (*Submitter, error) {
	newctx, cancel := context.WithTimeout(basectx, time.Minute)
	defer cancel()

	slog.Info("Connect to Eth rpc", "rpc", conf.Eth.RPC)
	ethClient, err := ethclient.DialContext(basectx, conf.Eth.RPC)
	if err != nil {
		return nil, err
	}

	ethChainId, err := ethClient.ChainID(newctx)
	if err != nil {
		return nil, err
	}

	slog.Info("Init LockingPool instance", "address", conf.Eth.LockingPool)
	lockingPool, err := lockingpool.NewLockingpool(conf.Eth.LockingPool, ethClient)
	if err != nil {
		return nil, err
	}

	slog.Info("Init LockingInfo instance", "address", conf.Eth.LockcingInfo)
	lockingInfo, err := lockinginfo.NewLockinginfo(conf.Eth.LockcingInfo, ethClient)
	if err != nil {
		return nil, err
	}

	slog.Info("Init MetisToken instance", "address", conf.Eth.MetisToken)
	metisToken, err := erc20.NewErc20(conf.Eth.MetisToken, ethClient)
	if err != nil {
		return nil, err
	}

	slog.Info("Connect to Metis rpc", "rpc", conf.Metis.RPC)
	metisClient, err := ethclient.DialContext(newctx, conf.Metis.RPC)
	if err != nil {
		return nil, err
	}

	metisChainId, err := metisClient.ChainID(newctx)
	if err != nil {
		return nil, err
	}

	slog.Info("chainId", "eth", ethChainId, "metis", metisChainId)

	slog.Info("Init SequencerSet instance", "address", conf.Metis.SeqSet)
	sequencerSet, err := sequencerset.NewSequencerset(conf.Metis.SeqSet, metisClient)
	if err != nil {
		return nil, err
	}

	slog.Info("Init Themis rest client", "host", conf.Metis.RPC)
	themisClient, err := themis.NewClient(conf.Metis.Themis)
	if err != nil {
		return nil, err
	}

	slog.Info("Restructure local state", "path", conf.Submitter.StatePath)
	var state State
	if file, err := os.ReadFile(conf.Submitter.StatePath); err == nil {
		if err := json.Unmarshal(file, &state); err != nil {
			return nil, fmt.Errorf("failed to decode local state file: %s", err)
		}
		slog.Info("Found last state", "status", state.Status, "timestamp", state.UpdatedAt)
	} else {
		slog.Info("No local state found, use on-chain state instead to start")
	}

	minInterval, err := time.ParseDuration(conf.Submitter.Params.MinInterval)
	if err != nil {
		return nil, fmt.Errorf("Invalid MinInterval param: %s", conf.Submitter.Params.MinInterval)
	}

	if conf.Submitter.Params.MinEpochLen == 0 {
		slog.Warn("MinEpochLen should be 1 at least")
		conf.Submitter.Params.MinEpochLen = 1
	}

	slog.Info("Using params",
		"MinEpochLen", conf.Submitter.Params.MinEpochLen,
		"MinInterval", minInterval,
		"GasPricingLimitInGWei", conf.Submitter.Params.GasPriceGWeiLimit,
	)
	return &Submitter{
		EthClient:    ethClient,
		MetisClient:  metisClient,
		ThemisClient: themisClient,
		SequencerSet: sequencerSet,
		LockingPool:  lockingPool,
		LockingInfo:  lockingInfo,
		MetisToken:   metisToken,
		Metric:       metric,
		params: &Params{
			StatePath:           conf.Submitter.StatePath,
			MinEpochLen:         conf.Submitter.Params.MinEpochLen,
			GasPriceLimit:       new(big.Int).Mul(conf.Submitter.Params.GasPriceGWeiLimit, big.NewInt(params.GWei)),
			EthChainId:          ethChainId,
			MetisChainId:        metisChainId,
			MinInterval:         minInterval,
			LockingPoolAddress:  conf.Eth.LockingPool,
			LockcingInfoAddress: conf.Eth.LockcingInfo,
			MetisTokenAddress:   conf.Eth.MetisToken,
			SeqSetAddress:       conf.Metis.SeqSet,
			TxHasher:            types.NewLondonSigner(ethChainId),
		},
		state: &state,
	}, nil
}

func (s *Submitter) Start(basectx context.Context) {
	timer := time.NewTimer(0)
	defer timer.Stop()

	for {
		select {
		case <-basectx.Done():
			return
		case <-timer.C:
			switch s.state.Status {
			case StatusIdle, StatusConfirmed:
				done, err := s.newBatch(basectx)
				if err != nil {
					s.Metric.Errors.Inc()
					slog.Error("newBatch", "err", err)
				}
				if done {
					timer.Reset(time.Second * 5)
				} else {
					timer.Reset(time.Minute)
				}
			case StatusUnsigned:
				done, err := s.withSignature(basectx)
				if err != nil {
					s.Metric.Errors.Inc()
					slog.Error("withSignature", "err", err)
				}
				if done {
					timer.Reset(0)
				} else {
					timer.Reset(time.Second * 5)
				}
			case StatusSigned:
				done, err := s.submitTx(basectx)
				if err != nil {
					s.Metric.Errors.Inc()
					slog.Error("sendTransaction", "err", err)
				}
				if done {
					timer.Reset(time.Second * 5)
				} else {
					timer.Reset(time.Second)
				}
			case StatusSubmitted:
				done, err := s.checkTx(basectx)
				if err != nil {
					s.Metric.Errors.Inc()
					slog.Error("checkTx", "err", err)
				}
				if done {
					timer.Reset(time.Minute * 10)
				} else {
					timer.Reset(time.Second * 10)
				}
			}
		}
	}
}

func (s *Submitter) Stop() {
	s.EthClient.Close()
	s.MetisClient.Close()
}

func (s *Submitter) saveState() error {
	file, err := os.Create(s.params.StatePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(s.state)
}

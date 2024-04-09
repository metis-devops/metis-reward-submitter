package submitter

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/metisprotocol/metis-peripheral/internal/utils"
)

type ChainState struct {
	FinalizedEpoch    uint64 // current L2 finalized epoch
	LastBatchId       uint64 // current batch id
	LastBatchEndEpoch uint64
	LastBatchTime     time.Time
	MpcAddress        common.Address
	MpcBalance        *big.Int
	MpcNonce          uint64
	PayerAddress      common.Address
	PayerBalance      *big.Int
	PayerAllowance    *big.Int
	RewardPerBlock    *big.Int
}

func (s *ChainState) ShouldSubmit(minLen uint64, minInterval time.Duration) bool {
	return (s.FinalizedEpoch > s.LastBatchEndEpoch+minLen) && time.Since(s.LastBatchTime) > minInterval
}

func (s *ChainState) IsPayerInited() bool {
	return !utils.IsZeroAddress(s.MpcAddress)
}

func (s *ChainState) IsMpcInited(a common.Address) bool {
	return a.Cmp(s.MpcAddress) == 0
}

func (s *ChainState) CanSubmit(reward *big.Int) bool {
	return s.PayerBalance.Cmp(reward) >= 0 && s.PayerAllowance.Cmp(reward) >= 0
}

func (s *ChainState) PayerHasSufficientBalance(reward *big.Int) bool {
	return s.PayerBalance.Cmp(reward) >= 0
}

func (s *ChainState) PayerHasSufficientAllowlance(reward *big.Int) bool {
	return s.PayerAllowance.Cmp(reward) >= 0
}

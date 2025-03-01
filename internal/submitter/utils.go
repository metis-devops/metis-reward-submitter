package submitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/metis-devops/metis-reward-submitter/internal/contracts/lockingpool"
)

type TxInput struct {
	BatchId    *big.Int         `json:"batchId"`
	StartEpoch *big.Int         `json:"startEpoch"`
	EndEpoch   *big.Int         `json:"endEpoch"`
	Seqs       []common.Address `json:"seqs"`
	Blocks     []*big.Int       `json:"blocks"`
}

func (t *TxInput) Encode() ([]byte, error) {
	abi, err := lockingpool.LockingpoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return abi.Pack("batchSubmitRewards",
		t.BatchId, t.StartEpoch, t.EndEpoch, t.Seqs, t.Blocks)
}

func (t *TxInput) JsonString() string {
	data, _ := json.Marshal(t)
	return string(data)
}

func DecodeBatchRewardInput(data []byte) (*TxInput, error) {
	if len(data) < 4 {
		return nil, errors.New("invalid input data")
	}

	abi, err := lockingpool.LockingpoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	method, ok := abi.Methods["batchSubmitRewards"]
	if !ok {
		return nil, fmt.Errorf("method batchSubmitRewards not found in abi")
	}

	inputs, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return nil, err
	}

	var result = new(TxInput)
	if err := method.Inputs.Copy(result, inputs); err != nil {
		return nil, err
	}
	return result, nil
}

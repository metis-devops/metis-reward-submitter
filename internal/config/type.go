package config

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Eth struct {
	RPC          string         `json:"rpc" yaml:"rpc"`
	LockingPool  common.Address `json:"locking_pool_address" yaml:"locking_pool_address"`
	LockcingInfo common.Address `json:"locking_info_address" yaml:"locking_info_address"`
	MetisToken   common.Address `json:"metis_token_address" yaml:"metis_token_address"`
}

type Metis struct {
	RPC    string         `json:"rpc" yaml:"rpc"`
	Themis string         `json:"themis_rest" yaml:"themis_rest"`
	SeqSet common.Address `json:"seq_set_address" yaml:"seq_set_address"`
}

type Submitter struct {
	StatePath string           `json:"state_path" yaml:"state_path"` // local cache file path
	Params    *SubmitterParams `json:"params" yaml:"params"`
}

type SubmitterParams struct {
	MinEpochLen       uint64   `json:"min_epoch_len,omitempty" yaml:"min_epoch_len"`
	MinInterval       string   `json:"min_interval" yaml:"min_interval"`
	GasPriceGWeiLimit *big.Int `json:"gas_price_limit_in_gwei" yaml:"gas_price_limit_in_gwei"`
}

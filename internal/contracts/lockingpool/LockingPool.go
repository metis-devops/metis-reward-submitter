// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lockingpool

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LockingpoolMetaData contains all meta data concerning the Lockingpool contract.
var LockingpoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"NoRewardRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSuchSeq\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSeqOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSeqSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SeqNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerExisted\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReward\",\"type\":\"uint256\"}],\"name\":\"RewardUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"SequencerOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"SequencerRewardRecipientChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"}],\"name\":\"SetSignerUpdateThrottle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_yes\",\"type\":\"bool\"}],\"name\":\"SetWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newMpc\",\"type\":\"address\"}],\"name\":\"UpdateMpc\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cur\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"}],\"name\":\"WithrawDelayTimeChange\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_seqs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_blocks\",\"type\":\"uint256[]\"}],\"name\":\"batchSubmitRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curBatchState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endEpoch\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow\",\"outputs\":[{\"internalType\":\"contractILockingInfo\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"}],\"name\":\"forceUnlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_escrow\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signerPubkey\",\"type\":\"bytes\"}],\"name\":\"lockFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signerPubkey\",\"type\":\"bytes\"}],\"name\":\"lockWithRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mpcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_lockReward\",\"type\":\"bool\"}],\"name\":\"relock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"seqOwners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"seqSigners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISequencerInfo.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"seqStatuses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seqId\",\"type\":\"uint256\"}],\"name\":\"sequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activationBatch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedBatch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivationBatch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockClaimTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"},{\"internalType\":\"enumISequencerInfo.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_yes\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setSequencerOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setSequencerRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"}],\"name\":\"setSignerUpdateThrottle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_yes\",\"type\":\"bool\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerUpdateThrottle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"}],\"name\":\"unlockClaim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"updateBlockReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMpc\",\"type\":\"address\"}],\"name\":\"updateMpc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signerPubkey\",\"type\":\"bytes\"}],\"name\":\"updateSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"updateWithdrawDelayTimeValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"}]",
}

// LockingpoolABI is the input ABI used to generate the binding from.
// Deprecated: Use LockingpoolMetaData.ABI instead.
var LockingpoolABI = LockingpoolMetaData.ABI

// Lockingpool is an auto generated Go binding around an Ethereum contract.
type Lockingpool struct {
	LockingpoolCaller     // Read-only binding to the contract
	LockingpoolTransactor // Write-only binding to the contract
	LockingpoolFilterer   // Log filterer for contract events
}

// LockingpoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockingpoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockingpoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockingpoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockingpoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockingpoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockingpoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockingpoolSession struct {
	Contract     *Lockingpool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockingpoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockingpoolCallerSession struct {
	Contract *LockingpoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LockingpoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockingpoolTransactorSession struct {
	Contract     *LockingpoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LockingpoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockingpoolRaw struct {
	Contract *Lockingpool // Generic contract binding to access the raw methods on
}

// LockingpoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockingpoolCallerRaw struct {
	Contract *LockingpoolCaller // Generic read-only contract binding to access the raw methods on
}

// LockingpoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockingpoolTransactorRaw struct {
	Contract *LockingpoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockingpool creates a new instance of Lockingpool, bound to a specific deployed contract.
func NewLockingpool(address common.Address, backend bind.ContractBackend) (*Lockingpool, error) {
	contract, err := bindLockingpool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lockingpool{LockingpoolCaller: LockingpoolCaller{contract: contract}, LockingpoolTransactor: LockingpoolTransactor{contract: contract}, LockingpoolFilterer: LockingpoolFilterer{contract: contract}}, nil
}

// NewLockingpoolCaller creates a new read-only instance of Lockingpool, bound to a specific deployed contract.
func NewLockingpoolCaller(address common.Address, caller bind.ContractCaller) (*LockingpoolCaller, error) {
	contract, err := bindLockingpool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockingpoolCaller{contract: contract}, nil
}

// NewLockingpoolTransactor creates a new write-only instance of Lockingpool, bound to a specific deployed contract.
func NewLockingpoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LockingpoolTransactor, error) {
	contract, err := bindLockingpool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockingpoolTransactor{contract: contract}, nil
}

// NewLockingpoolFilterer creates a new log filterer instance of Lockingpool, bound to a specific deployed contract.
func NewLockingpoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LockingpoolFilterer, error) {
	contract, err := bindLockingpool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockingpoolFilterer{contract: contract}, nil
}

// bindLockingpool binds a generic wrapper to an already deployed contract.
func bindLockingpool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LockingpoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lockingpool *LockingpoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lockingpool.Contract.LockingpoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lockingpool *LockingpoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockingpool.Contract.LockingpoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lockingpool *LockingpoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lockingpool.Contract.LockingpoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lockingpool *LockingpoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lockingpool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lockingpool *LockingpoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockingpool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lockingpool *LockingpoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lockingpool.Contract.contract.Transact(opts, method, params...)
}

// BLOCKREWARD is a free data retrieval call binding the contract method 0x7f05b9ef.
//
// Solidity: function BLOCK_REWARD() view returns(uint256)
func (_Lockingpool *LockingpoolCaller) BLOCKREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "BLOCK_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLOCKREWARD is a free data retrieval call binding the contract method 0x7f05b9ef.
//
// Solidity: function BLOCK_REWARD() view returns(uint256)
func (_Lockingpool *LockingpoolSession) BLOCKREWARD() (*big.Int, error) {
	return _Lockingpool.Contract.BLOCKREWARD(&_Lockingpool.CallOpts)
}

// BLOCKREWARD is a free data retrieval call binding the contract method 0x7f05b9ef.
//
// Solidity: function BLOCK_REWARD() view returns(uint256)
func (_Lockingpool *LockingpoolCallerSession) BLOCKREWARD() (*big.Int, error) {
	return _Lockingpool.Contract.BLOCKREWARD(&_Lockingpool.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_Lockingpool *LockingpoolCaller) WITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "WITHDRAWAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_Lockingpool *LockingpoolSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _Lockingpool.Contract.WITHDRAWALDELAY(&_Lockingpool.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_Lockingpool *LockingpoolCallerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _Lockingpool.Contract.WITHDRAWALDELAY(&_Lockingpool.CallOpts)
}

// CurBatchState is a free data retrieval call binding the contract method 0xb4472970.
//
// Solidity: function curBatchState() view returns(uint256 id, uint256 number, uint256 startEpoch, uint256 endEpoch)
func (_Lockingpool *LockingpoolCaller) CurBatchState(opts *bind.CallOpts) (struct {
	Id         *big.Int
	Number     *big.Int
	StartEpoch *big.Int
	EndEpoch   *big.Int
}, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "curBatchState")

	outstruct := new(struct {
		Id         *big.Int
		Number     *big.Int
		StartEpoch *big.Int
		EndEpoch   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Number = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurBatchState is a free data retrieval call binding the contract method 0xb4472970.
//
// Solidity: function curBatchState() view returns(uint256 id, uint256 number, uint256 startEpoch, uint256 endEpoch)
func (_Lockingpool *LockingpoolSession) CurBatchState() (struct {
	Id         *big.Int
	Number     *big.Int
	StartEpoch *big.Int
	EndEpoch   *big.Int
}, error) {
	return _Lockingpool.Contract.CurBatchState(&_Lockingpool.CallOpts)
}

// CurBatchState is a free data retrieval call binding the contract method 0xb4472970.
//
// Solidity: function curBatchState() view returns(uint256 id, uint256 number, uint256 startEpoch, uint256 endEpoch)
func (_Lockingpool *LockingpoolCallerSession) CurBatchState() (struct {
	Id         *big.Int
	Number     *big.Int
	StartEpoch *big.Int
	EndEpoch   *big.Int
}, error) {
	return _Lockingpool.Contract.CurBatchState(&_Lockingpool.CallOpts)
}

// CurrentBatch is a free data retrieval call binding the contract method 0x76cd940e.
//
// Solidity: function currentBatch() view returns(uint256)
func (_Lockingpool *LockingpoolCaller) CurrentBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "currentBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentBatch is a free data retrieval call binding the contract method 0x76cd940e.
//
// Solidity: function currentBatch() view returns(uint256)
func (_Lockingpool *LockingpoolSession) CurrentBatch() (*big.Int, error) {
	return _Lockingpool.Contract.CurrentBatch(&_Lockingpool.CallOpts)
}

// CurrentBatch is a free data retrieval call binding the contract method 0x76cd940e.
//
// Solidity: function currentBatch() view returns(uint256)
func (_Lockingpool *LockingpoolCallerSession) CurrentBatch() (*big.Int, error) {
	return _Lockingpool.Contract.CurrentBatch(&_Lockingpool.CallOpts)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_Lockingpool *LockingpoolCaller) Escrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_Lockingpool *LockingpoolSession) Escrow() (common.Address, error) {
	return _Lockingpool.Contract.Escrow(&_Lockingpool.CallOpts)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_Lockingpool *LockingpoolCallerSession) Escrow() (common.Address, error) {
	return _Lockingpool.Contract.Escrow(&_Lockingpool.CallOpts)
}

// MpcAddress is a free data retrieval call binding the contract method 0x111f4630.
//
// Solidity: function mpcAddress() view returns(address)
func (_Lockingpool *LockingpoolCaller) MpcAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "mpcAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MpcAddress is a free data retrieval call binding the contract method 0x111f4630.
//
// Solidity: function mpcAddress() view returns(address)
func (_Lockingpool *LockingpoolSession) MpcAddress() (common.Address, error) {
	return _Lockingpool.Contract.MpcAddress(&_Lockingpool.CallOpts)
}

// MpcAddress is a free data retrieval call binding the contract method 0x111f4630.
//
// Solidity: function mpcAddress() view returns(address)
func (_Lockingpool *LockingpoolCallerSession) MpcAddress() (common.Address, error) {
	return _Lockingpool.Contract.MpcAddress(&_Lockingpool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lockingpool *LockingpoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lockingpool *LockingpoolSession) Owner() (common.Address, error) {
	return _Lockingpool.Contract.Owner(&_Lockingpool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lockingpool *LockingpoolCallerSession) Owner() (common.Address, error) {
	return _Lockingpool.Contract.Owner(&_Lockingpool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lockingpool *LockingpoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lockingpool *LockingpoolSession) Paused() (bool, error) {
	return _Lockingpool.Contract.Paused(&_Lockingpool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lockingpool *LockingpoolCallerSession) Paused() (bool, error) {
	return _Lockingpool.Contract.Paused(&_Lockingpool.CallOpts)
}

// SeqOwners is a free data retrieval call binding the contract method 0x169abefc.
//
// Solidity: function seqOwners(address owner) view returns(uint256 seqId)
func (_Lockingpool *LockingpoolCaller) SeqOwners(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "seqOwners", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SeqOwners is a free data retrieval call binding the contract method 0x169abefc.
//
// Solidity: function seqOwners(address owner) view returns(uint256 seqId)
func (_Lockingpool *LockingpoolSession) SeqOwners(owner common.Address) (*big.Int, error) {
	return _Lockingpool.Contract.SeqOwners(&_Lockingpool.CallOpts, owner)
}

// SeqOwners is a free data retrieval call binding the contract method 0x169abefc.
//
// Solidity: function seqOwners(address owner) view returns(uint256 seqId)
func (_Lockingpool *LockingpoolCallerSession) SeqOwners(owner common.Address) (*big.Int, error) {
	return _Lockingpool.Contract.SeqOwners(&_Lockingpool.CallOpts, owner)
}

// SeqSigners is a free data retrieval call binding the contract method 0xbeb26755.
//
// Solidity: function seqSigners(address signer) view returns(uint256 seqId)
func (_Lockingpool *LockingpoolCaller) SeqSigners(opts *bind.CallOpts, signer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "seqSigners", signer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SeqSigners is a free data retrieval call binding the contract method 0xbeb26755.
//
// Solidity: function seqSigners(address signer) view returns(uint256 seqId)
func (_Lockingpool *LockingpoolSession) SeqSigners(signer common.Address) (*big.Int, error) {
	return _Lockingpool.Contract.SeqSigners(&_Lockingpool.CallOpts, signer)
}

// SeqSigners is a free data retrieval call binding the contract method 0xbeb26755.
//
// Solidity: function seqSigners(address signer) view returns(uint256 seqId)
func (_Lockingpool *LockingpoolCallerSession) SeqSigners(signer common.Address) (*big.Int, error) {
	return _Lockingpool.Contract.SeqSigners(&_Lockingpool.CallOpts, signer)
}

// SeqStatuses is a free data retrieval call binding the contract method 0x86d203ab.
//
// Solidity: function seqStatuses(uint8 status) view returns(uint256 count)
func (_Lockingpool *LockingpoolCaller) SeqStatuses(opts *bind.CallOpts, status uint8) (*big.Int, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "seqStatuses", status)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SeqStatuses is a free data retrieval call binding the contract method 0x86d203ab.
//
// Solidity: function seqStatuses(uint8 status) view returns(uint256 count)
func (_Lockingpool *LockingpoolSession) SeqStatuses(status uint8) (*big.Int, error) {
	return _Lockingpool.Contract.SeqStatuses(&_Lockingpool.CallOpts, status)
}

// SeqStatuses is a free data retrieval call binding the contract method 0x86d203ab.
//
// Solidity: function seqStatuses(uint8 status) view returns(uint256 count)
func (_Lockingpool *LockingpoolCallerSession) SeqStatuses(status uint8) (*big.Int, error) {
	return _Lockingpool.Contract.SeqStatuses(&_Lockingpool.CallOpts, status)
}

// Sequencers is a free data retrieval call binding the contract method 0x6ba7ccff.
//
// Solidity: function sequencers(uint256 seqId) view returns(uint256 amount, uint256 reward, uint256 activationBatch, uint256 updatedBatch, uint256 deactivationBatch, uint256 deactivationTime, uint256 unlockClaimTime, uint256 nonce, address owner, address signer, bytes pubkey, address rewardRecipient, uint8 status)
func (_Lockingpool *LockingpoolCaller) Sequencers(opts *bind.CallOpts, seqId *big.Int) (struct {
	Amount            *big.Int
	Reward            *big.Int
	ActivationBatch   *big.Int
	UpdatedBatch      *big.Int
	DeactivationBatch *big.Int
	DeactivationTime  *big.Int
	UnlockClaimTime   *big.Int
	Nonce             *big.Int
	Owner             common.Address
	Signer            common.Address
	Pubkey            []byte
	RewardRecipient   common.Address
	Status            uint8
}, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "sequencers", seqId)

	outstruct := new(struct {
		Amount            *big.Int
		Reward            *big.Int
		ActivationBatch   *big.Int
		UpdatedBatch      *big.Int
		DeactivationBatch *big.Int
		DeactivationTime  *big.Int
		UnlockClaimTime   *big.Int
		Nonce             *big.Int
		Owner             common.Address
		Signer            common.Address
		Pubkey            []byte
		RewardRecipient   common.Address
		Status            uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ActivationBatch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedBatch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DeactivationBatch = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.DeactivationTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.UnlockClaimTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.Signer = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.Pubkey = *abi.ConvertType(out[10], new([]byte)).(*[]byte)
	outstruct.RewardRecipient = *abi.ConvertType(out[11], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[12], new(uint8)).(*uint8)

	return *outstruct, err

}

// Sequencers is a free data retrieval call binding the contract method 0x6ba7ccff.
//
// Solidity: function sequencers(uint256 seqId) view returns(uint256 amount, uint256 reward, uint256 activationBatch, uint256 updatedBatch, uint256 deactivationBatch, uint256 deactivationTime, uint256 unlockClaimTime, uint256 nonce, address owner, address signer, bytes pubkey, address rewardRecipient, uint8 status)
func (_Lockingpool *LockingpoolSession) Sequencers(seqId *big.Int) (struct {
	Amount            *big.Int
	Reward            *big.Int
	ActivationBatch   *big.Int
	UpdatedBatch      *big.Int
	DeactivationBatch *big.Int
	DeactivationTime  *big.Int
	UnlockClaimTime   *big.Int
	Nonce             *big.Int
	Owner             common.Address
	Signer            common.Address
	Pubkey            []byte
	RewardRecipient   common.Address
	Status            uint8
}, error) {
	return _Lockingpool.Contract.Sequencers(&_Lockingpool.CallOpts, seqId)
}

// Sequencers is a free data retrieval call binding the contract method 0x6ba7ccff.
//
// Solidity: function sequencers(uint256 seqId) view returns(uint256 amount, uint256 reward, uint256 activationBatch, uint256 updatedBatch, uint256 deactivationBatch, uint256 deactivationTime, uint256 unlockClaimTime, uint256 nonce, address owner, address signer, bytes pubkey, address rewardRecipient, uint8 status)
func (_Lockingpool *LockingpoolCallerSession) Sequencers(seqId *big.Int) (struct {
	Amount            *big.Int
	Reward            *big.Int
	ActivationBatch   *big.Int
	UpdatedBatch      *big.Int
	DeactivationBatch *big.Int
	DeactivationTime  *big.Int
	UnlockClaimTime   *big.Int
	Nonce             *big.Int
	Owner             common.Address
	Signer            common.Address
	Pubkey            []byte
	RewardRecipient   common.Address
	Status            uint8
}, error) {
	return _Lockingpool.Contract.Sequencers(&_Lockingpool.CallOpts, seqId)
}

// SignerUpdateThrottle is a free data retrieval call binding the contract method 0xc65066d4.
//
// Solidity: function signerUpdateThrottle() view returns(uint256)
func (_Lockingpool *LockingpoolCaller) SignerUpdateThrottle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "signerUpdateThrottle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignerUpdateThrottle is a free data retrieval call binding the contract method 0xc65066d4.
//
// Solidity: function signerUpdateThrottle() view returns(uint256)
func (_Lockingpool *LockingpoolSession) SignerUpdateThrottle() (*big.Int, error) {
	return _Lockingpool.Contract.SignerUpdateThrottle(&_Lockingpool.CallOpts)
}

// SignerUpdateThrottle is a free data retrieval call binding the contract method 0xc65066d4.
//
// Solidity: function signerUpdateThrottle() view returns(uint256)
func (_Lockingpool *LockingpoolCallerSession) SignerUpdateThrottle() (*big.Int, error) {
	return _Lockingpool.Contract.SignerUpdateThrottle(&_Lockingpool.CallOpts)
}

// TotalSequencers is a free data retrieval call binding the contract method 0xcc3ab923.
//
// Solidity: function totalSequencers() view returns(uint256)
func (_Lockingpool *LockingpoolCaller) TotalSequencers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "totalSequencers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSequencers is a free data retrieval call binding the contract method 0xcc3ab923.
//
// Solidity: function totalSequencers() view returns(uint256)
func (_Lockingpool *LockingpoolSession) TotalSequencers() (*big.Int, error) {
	return _Lockingpool.Contract.TotalSequencers(&_Lockingpool.CallOpts)
}

// TotalSequencers is a free data retrieval call binding the contract method 0xcc3ab923.
//
// Solidity: function totalSequencers() view returns(uint256)
func (_Lockingpool *LockingpoolCallerSession) TotalSequencers() (*big.Int, error) {
	return _Lockingpool.Contract.TotalSequencers(&_Lockingpool.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Lockingpool *LockingpoolCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Lockingpool.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Lockingpool *LockingpoolSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Lockingpool.Contract.Whitelist(&_Lockingpool.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Lockingpool *LockingpoolCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Lockingpool.Contract.Whitelist(&_Lockingpool.CallOpts, arg0)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Lockingpool *LockingpoolTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Lockingpool *LockingpoolSession) Admin() (*types.Transaction, error) {
	return _Lockingpool.Contract.Admin(&_Lockingpool.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Lockingpool *LockingpoolTransactorSession) Admin() (*types.Transaction, error) {
	return _Lockingpool.Contract.Admin(&_Lockingpool.TransactOpts)
}

// BatchSubmitRewards is a paid mutator transaction binding the contract method 0x11c7d144.
//
// Solidity: function batchSubmitRewards(uint256 _batchId, uint256 _startEpoch, uint256 _endEpoch, address[] _seqs, uint256[] _blocks) returns(uint256 totalReward)
func (_Lockingpool *LockingpoolTransactor) BatchSubmitRewards(opts *bind.TransactOpts, _batchId *big.Int, _startEpoch *big.Int, _endEpoch *big.Int, _seqs []common.Address, _blocks []*big.Int) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "batchSubmitRewards", _batchId, _startEpoch, _endEpoch, _seqs, _blocks)
}

// BatchSubmitRewards is a paid mutator transaction binding the contract method 0x11c7d144.
//
// Solidity: function batchSubmitRewards(uint256 _batchId, uint256 _startEpoch, uint256 _endEpoch, address[] _seqs, uint256[] _blocks) returns(uint256 totalReward)
func (_Lockingpool *LockingpoolSession) BatchSubmitRewards(_batchId *big.Int, _startEpoch *big.Int, _endEpoch *big.Int, _seqs []common.Address, _blocks []*big.Int) (*types.Transaction, error) {
	return _Lockingpool.Contract.BatchSubmitRewards(&_Lockingpool.TransactOpts, _batchId, _startEpoch, _endEpoch, _seqs, _blocks)
}

// BatchSubmitRewards is a paid mutator transaction binding the contract method 0x11c7d144.
//
// Solidity: function batchSubmitRewards(uint256 _batchId, uint256 _startEpoch, uint256 _endEpoch, address[] _seqs, uint256[] _blocks) returns(uint256 totalReward)
func (_Lockingpool *LockingpoolTransactorSession) BatchSubmitRewards(_batchId *big.Int, _startEpoch *big.Int, _endEpoch *big.Int, _seqs []common.Address, _blocks []*big.Int) (*types.Transaction, error) {
	return _Lockingpool.Contract.BatchSubmitRewards(&_Lockingpool.TransactOpts, _batchId, _startEpoch, _endEpoch, _seqs, _blocks)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Lockingpool *LockingpoolTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Lockingpool *LockingpoolSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.ChangeAdmin(&_Lockingpool.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Lockingpool *LockingpoolTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.ChangeAdmin(&_Lockingpool.TransactOpts, newAdmin)
}

// ForceUnlock is a paid mutator transaction binding the contract method 0xca99e838.
//
// Solidity: function forceUnlock(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolTransactor) ForceUnlock(opts *bind.TransactOpts, _seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "forceUnlock", _seqId, _l2Gas)
}

// ForceUnlock is a paid mutator transaction binding the contract method 0xca99e838.
//
// Solidity: function forceUnlock(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolSession) ForceUnlock(_seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.Contract.ForceUnlock(&_Lockingpool.TransactOpts, _seqId, _l2Gas)
}

// ForceUnlock is a paid mutator transaction binding the contract method 0xca99e838.
//
// Solidity: function forceUnlock(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolTransactorSession) ForceUnlock(_seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.Contract.ForceUnlock(&_Lockingpool.TransactOpts, _seqId, _l2Gas)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Lockingpool *LockingpoolTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Lockingpool *LockingpoolSession) Implementation() (*types.Transaction, error) {
	return _Lockingpool.Contract.Implementation(&_Lockingpool.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Lockingpool *LockingpoolTransactorSession) Implementation() (*types.Transaction, error) {
	return _Lockingpool.Contract.Implementation(&_Lockingpool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _escrow) returns()
func (_Lockingpool *LockingpoolTransactor) Initialize(opts *bind.TransactOpts, _escrow common.Address) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "initialize", _escrow)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _escrow) returns()
func (_Lockingpool *LockingpoolSession) Initialize(_escrow common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.Initialize(&_Lockingpool.TransactOpts, _escrow)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _escrow) returns()
func (_Lockingpool *LockingpoolTransactorSession) Initialize(_escrow common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.Initialize(&_Lockingpool.TransactOpts, _escrow)
}

// LockFor is a paid mutator transaction binding the contract method 0xaf70cba5.
//
// Solidity: function lockFor(address _signer, uint256 _amount, bytes _signerPubkey) returns()
func (_Lockingpool *LockingpoolTransactor) LockFor(opts *bind.TransactOpts, _signer common.Address, _amount *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "lockFor", _signer, _amount, _signerPubkey)
}

// LockFor is a paid mutator transaction binding the contract method 0xaf70cba5.
//
// Solidity: function lockFor(address _signer, uint256 _amount, bytes _signerPubkey) returns()
func (_Lockingpool *LockingpoolSession) LockFor(_signer common.Address, _amount *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockingpool.Contract.LockFor(&_Lockingpool.TransactOpts, _signer, _amount, _signerPubkey)
}

// LockFor is a paid mutator transaction binding the contract method 0xaf70cba5.
//
// Solidity: function lockFor(address _signer, uint256 _amount, bytes _signerPubkey) returns()
func (_Lockingpool *LockingpoolTransactorSession) LockFor(_signer common.Address, _amount *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockingpool.Contract.LockFor(&_Lockingpool.TransactOpts, _signer, _amount, _signerPubkey)
}

// LockWithRewardRecipient is a paid mutator transaction binding the contract method 0x9ad42030.
//
// Solidity: function lockWithRewardRecipient(address _signer, address _rewardRecipient, uint256 _amount, bytes _signerPubkey) returns()
func (_Lockingpool *LockingpoolTransactor) LockWithRewardRecipient(opts *bind.TransactOpts, _signer common.Address, _rewardRecipient common.Address, _amount *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "lockWithRewardRecipient", _signer, _rewardRecipient, _amount, _signerPubkey)
}

// LockWithRewardRecipient is a paid mutator transaction binding the contract method 0x9ad42030.
//
// Solidity: function lockWithRewardRecipient(address _signer, address _rewardRecipient, uint256 _amount, bytes _signerPubkey) returns()
func (_Lockingpool *LockingpoolSession) LockWithRewardRecipient(_signer common.Address, _rewardRecipient common.Address, _amount *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockingpool.Contract.LockWithRewardRecipient(&_Lockingpool.TransactOpts, _signer, _rewardRecipient, _amount, _signerPubkey)
}

// LockWithRewardRecipient is a paid mutator transaction binding the contract method 0x9ad42030.
//
// Solidity: function lockWithRewardRecipient(address _signer, address _rewardRecipient, uint256 _amount, bytes _signerPubkey) returns()
func (_Lockingpool *LockingpoolTransactorSession) LockWithRewardRecipient(_signer common.Address, _rewardRecipient common.Address, _amount *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockingpool.Contract.LockWithRewardRecipient(&_Lockingpool.TransactOpts, _signer, _rewardRecipient, _amount, _signerPubkey)
}

// Relock is a paid mutator transaction binding the contract method 0x015bb180.
//
// Solidity: function relock(uint256 _seqId, uint256 _amount, bool _lockReward) returns()
func (_Lockingpool *LockingpoolTransactor) Relock(opts *bind.TransactOpts, _seqId *big.Int, _amount *big.Int, _lockReward bool) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "relock", _seqId, _amount, _lockReward)
}

// Relock is a paid mutator transaction binding the contract method 0x015bb180.
//
// Solidity: function relock(uint256 _seqId, uint256 _amount, bool _lockReward) returns()
func (_Lockingpool *LockingpoolSession) Relock(_seqId *big.Int, _amount *big.Int, _lockReward bool) (*types.Transaction, error) {
	return _Lockingpool.Contract.Relock(&_Lockingpool.TransactOpts, _seqId, _amount, _lockReward)
}

// Relock is a paid mutator transaction binding the contract method 0x015bb180.
//
// Solidity: function relock(uint256 _seqId, uint256 _amount, bool _lockReward) returns()
func (_Lockingpool *LockingpoolTransactorSession) Relock(_seqId *big.Int, _amount *big.Int, _lockReward bool) (*types.Transaction, error) {
	return _Lockingpool.Contract.Relock(&_Lockingpool.TransactOpts, _seqId, _amount, _lockReward)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lockingpool *LockingpoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lockingpool *LockingpoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lockingpool.Contract.RenounceOwnership(&_Lockingpool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lockingpool *LockingpoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lockingpool.Contract.RenounceOwnership(&_Lockingpool.TransactOpts)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _yes) returns()
func (_Lockingpool *LockingpoolTransactor) SetPause(opts *bind.TransactOpts, _yes bool) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "setPause", _yes)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _yes) returns()
func (_Lockingpool *LockingpoolSession) SetPause(_yes bool) (*types.Transaction, error) {
	return _Lockingpool.Contract.SetPause(&_Lockingpool.TransactOpts, _yes)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _yes) returns()
func (_Lockingpool *LockingpoolTransactorSession) SetPause(_yes bool) (*types.Transaction, error) {
	return _Lockingpool.Contract.SetPause(&_Lockingpool.TransactOpts, _yes)
}

// SetSequencerOwner is a paid mutator transaction binding the contract method 0xa953791f.
//
// Solidity: function setSequencerOwner(uint256 _seqId, address _owner) returns()
func (_Lockingpool *LockingpoolTransactor) SetSequencerOwner(opts *bind.TransactOpts, _seqId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "setSequencerOwner", _seqId, _owner)
}

// SetSequencerOwner is a paid mutator transaction binding the contract method 0xa953791f.
//
// Solidity: function setSequencerOwner(uint256 _seqId, address _owner) returns()
func (_Lockingpool *LockingpoolSession) SetSequencerOwner(_seqId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.SetSequencerOwner(&_Lockingpool.TransactOpts, _seqId, _owner)
}

// SetSequencerOwner is a paid mutator transaction binding the contract method 0xa953791f.
//
// Solidity: function setSequencerOwner(uint256 _seqId, address _owner) returns()
func (_Lockingpool *LockingpoolTransactorSession) SetSequencerOwner(_seqId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.SetSequencerOwner(&_Lockingpool.TransactOpts, _seqId, _owner)
}

// SetSequencerRewardRecipient is a paid mutator transaction binding the contract method 0xd83b0e14.
//
// Solidity: function setSequencerRewardRecipient(uint256 _seqId, address _recipient) returns()
func (_Lockingpool *LockingpoolTransactor) SetSequencerRewardRecipient(opts *bind.TransactOpts, _seqId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "setSequencerRewardRecipient", _seqId, _recipient)
}

// SetSequencerRewardRecipient is a paid mutator transaction binding the contract method 0xd83b0e14.
//
// Solidity: function setSequencerRewardRecipient(uint256 _seqId, address _recipient) returns()
func (_Lockingpool *LockingpoolSession) SetSequencerRewardRecipient(_seqId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.SetSequencerRewardRecipient(&_Lockingpool.TransactOpts, _seqId, _recipient)
}

// SetSequencerRewardRecipient is a paid mutator transaction binding the contract method 0xd83b0e14.
//
// Solidity: function setSequencerRewardRecipient(uint256 _seqId, address _recipient) returns()
func (_Lockingpool *LockingpoolTransactorSession) SetSequencerRewardRecipient(_seqId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.SetSequencerRewardRecipient(&_Lockingpool.TransactOpts, _seqId, _recipient)
}

// SetSignerUpdateThrottle is a paid mutator transaction binding the contract method 0xbfd6fc3f.
//
// Solidity: function setSignerUpdateThrottle(uint256 _n) returns()
func (_Lockingpool *LockingpoolTransactor) SetSignerUpdateThrottle(opts *bind.TransactOpts, _n *big.Int) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "setSignerUpdateThrottle", _n)
}

// SetSignerUpdateThrottle is a paid mutator transaction binding the contract method 0xbfd6fc3f.
//
// Solidity: function setSignerUpdateThrottle(uint256 _n) returns()
func (_Lockingpool *LockingpoolSession) SetSignerUpdateThrottle(_n *big.Int) (*types.Transaction, error) {
	return _Lockingpool.Contract.SetSignerUpdateThrottle(&_Lockingpool.TransactOpts, _n)
}

// SetSignerUpdateThrottle is a paid mutator transaction binding the contract method 0xbfd6fc3f.
//
// Solidity: function setSignerUpdateThrottle(uint256 _n) returns()
func (_Lockingpool *LockingpoolTransactorSession) SetSignerUpdateThrottle(_n *big.Int) (*types.Transaction, error) {
	return _Lockingpool.Contract.SetSignerUpdateThrottle(&_Lockingpool.TransactOpts, _n)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x53d6fd59.
//
// Solidity: function setWhitelist(address _addr, bool _yes) returns()
func (_Lockingpool *LockingpoolTransactor) SetWhitelist(opts *bind.TransactOpts, _addr common.Address, _yes bool) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "setWhitelist", _addr, _yes)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x53d6fd59.
//
// Solidity: function setWhitelist(address _addr, bool _yes) returns()
func (_Lockingpool *LockingpoolSession) SetWhitelist(_addr common.Address, _yes bool) (*types.Transaction, error) {
	return _Lockingpool.Contract.SetWhitelist(&_Lockingpool.TransactOpts, _addr, _yes)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x53d6fd59.
//
// Solidity: function setWhitelist(address _addr, bool _yes) returns()
func (_Lockingpool *LockingpoolTransactorSession) SetWhitelist(_addr common.Address, _yes bool) (*types.Transaction, error) {
	return _Lockingpool.Contract.SetWhitelist(&_Lockingpool.TransactOpts, _addr, _yes)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lockingpool *LockingpoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lockingpool *LockingpoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.TransferOwnership(&_Lockingpool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lockingpool *LockingpoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.TransferOwnership(&_Lockingpool.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x262c0e66.
//
// Solidity: function unlock(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolTransactor) Unlock(opts *bind.TransactOpts, _seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "unlock", _seqId, _l2Gas)
}

// Unlock is a paid mutator transaction binding the contract method 0x262c0e66.
//
// Solidity: function unlock(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolSession) Unlock(_seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.Contract.Unlock(&_Lockingpool.TransactOpts, _seqId, _l2Gas)
}

// Unlock is a paid mutator transaction binding the contract method 0x262c0e66.
//
// Solidity: function unlock(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolTransactorSession) Unlock(_seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.Contract.Unlock(&_Lockingpool.TransactOpts, _seqId, _l2Gas)
}

// UnlockClaim is a paid mutator transaction binding the contract method 0x8ddc74de.
//
// Solidity: function unlockClaim(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolTransactor) UnlockClaim(opts *bind.TransactOpts, _seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "unlockClaim", _seqId, _l2Gas)
}

// UnlockClaim is a paid mutator transaction binding the contract method 0x8ddc74de.
//
// Solidity: function unlockClaim(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolSession) UnlockClaim(_seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.Contract.UnlockClaim(&_Lockingpool.TransactOpts, _seqId, _l2Gas)
}

// UnlockClaim is a paid mutator transaction binding the contract method 0x8ddc74de.
//
// Solidity: function unlockClaim(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolTransactorSession) UnlockClaim(_seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.Contract.UnlockClaim(&_Lockingpool.TransactOpts, _seqId, _l2Gas)
}

// UpdateBlockReward is a paid mutator transaction binding the contract method 0xf580ffcb.
//
// Solidity: function updateBlockReward(uint256 newReward) returns()
func (_Lockingpool *LockingpoolTransactor) UpdateBlockReward(opts *bind.TransactOpts, newReward *big.Int) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "updateBlockReward", newReward)
}

// UpdateBlockReward is a paid mutator transaction binding the contract method 0xf580ffcb.
//
// Solidity: function updateBlockReward(uint256 newReward) returns()
func (_Lockingpool *LockingpoolSession) UpdateBlockReward(newReward *big.Int) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpdateBlockReward(&_Lockingpool.TransactOpts, newReward)
}

// UpdateBlockReward is a paid mutator transaction binding the contract method 0xf580ffcb.
//
// Solidity: function updateBlockReward(uint256 newReward) returns()
func (_Lockingpool *LockingpoolTransactorSession) UpdateBlockReward(newReward *big.Int) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpdateBlockReward(&_Lockingpool.TransactOpts, newReward)
}

// UpdateMpc is a paid mutator transaction binding the contract method 0xd11d0681.
//
// Solidity: function updateMpc(address _newMpc) returns()
func (_Lockingpool *LockingpoolTransactor) UpdateMpc(opts *bind.TransactOpts, _newMpc common.Address) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "updateMpc", _newMpc)
}

// UpdateMpc is a paid mutator transaction binding the contract method 0xd11d0681.
//
// Solidity: function updateMpc(address _newMpc) returns()
func (_Lockingpool *LockingpoolSession) UpdateMpc(_newMpc common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpdateMpc(&_Lockingpool.TransactOpts, _newMpc)
}

// UpdateMpc is a paid mutator transaction binding the contract method 0xd11d0681.
//
// Solidity: function updateMpc(address _newMpc) returns()
func (_Lockingpool *LockingpoolTransactorSession) UpdateMpc(_newMpc common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpdateMpc(&_Lockingpool.TransactOpts, _newMpc)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf41a9642.
//
// Solidity: function updateSigner(uint256 _seqId, bytes _signerPubkey) returns()
func (_Lockingpool *LockingpoolTransactor) UpdateSigner(opts *bind.TransactOpts, _seqId *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "updateSigner", _seqId, _signerPubkey)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf41a9642.
//
// Solidity: function updateSigner(uint256 _seqId, bytes _signerPubkey) returns()
func (_Lockingpool *LockingpoolSession) UpdateSigner(_seqId *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpdateSigner(&_Lockingpool.TransactOpts, _seqId, _signerPubkey)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf41a9642.
//
// Solidity: function updateSigner(uint256 _seqId, bytes _signerPubkey) returns()
func (_Lockingpool *LockingpoolTransactorSession) UpdateSigner(_seqId *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpdateSigner(&_Lockingpool.TransactOpts, _seqId, _signerPubkey)
}

// UpdateWithdrawDelayTimeValue is a paid mutator transaction binding the contract method 0x71e10cfa.
//
// Solidity: function updateWithdrawDelayTimeValue(uint256 _time) returns()
func (_Lockingpool *LockingpoolTransactor) UpdateWithdrawDelayTimeValue(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "updateWithdrawDelayTimeValue", _time)
}

// UpdateWithdrawDelayTimeValue is a paid mutator transaction binding the contract method 0x71e10cfa.
//
// Solidity: function updateWithdrawDelayTimeValue(uint256 _time) returns()
func (_Lockingpool *LockingpoolSession) UpdateWithdrawDelayTimeValue(_time *big.Int) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpdateWithdrawDelayTimeValue(&_Lockingpool.TransactOpts, _time)
}

// UpdateWithdrawDelayTimeValue is a paid mutator transaction binding the contract method 0x71e10cfa.
//
// Solidity: function updateWithdrawDelayTimeValue(uint256 _time) returns()
func (_Lockingpool *LockingpoolTransactorSession) UpdateWithdrawDelayTimeValue(_time *big.Int) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpdateWithdrawDelayTimeValue(&_Lockingpool.TransactOpts, _time)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Lockingpool *LockingpoolTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Lockingpool *LockingpoolSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpgradeTo(&_Lockingpool.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Lockingpool *LockingpoolTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpgradeTo(&_Lockingpool.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lockingpool *LockingpoolTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lockingpool *LockingpoolSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpgradeToAndCall(&_Lockingpool.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lockingpool *LockingpoolTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lockingpool.Contract.UpgradeToAndCall(&_Lockingpool.TransactOpts, newImplementation, data)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x17396687.
//
// Solidity: function withdrawRewards(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolTransactor) WithdrawRewards(opts *bind.TransactOpts, _seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.contract.Transact(opts, "withdrawRewards", _seqId, _l2Gas)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x17396687.
//
// Solidity: function withdrawRewards(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolSession) WithdrawRewards(_seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.Contract.WithdrawRewards(&_Lockingpool.TransactOpts, _seqId, _l2Gas)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x17396687.
//
// Solidity: function withdrawRewards(uint256 _seqId, uint32 _l2Gas) payable returns()
func (_Lockingpool *LockingpoolTransactorSession) WithdrawRewards(_seqId *big.Int, _l2Gas uint32) (*types.Transaction, error) {
	return _Lockingpool.Contract.WithdrawRewards(&_Lockingpool.TransactOpts, _seqId, _l2Gas)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lockingpool *LockingpoolTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Lockingpool.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lockingpool *LockingpoolSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Lockingpool.Contract.Fallback(&_Lockingpool.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lockingpool *LockingpoolTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Lockingpool.Contract.Fallback(&_Lockingpool.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lockingpool *LockingpoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockingpool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lockingpool *LockingpoolSession) Receive() (*types.Transaction, error) {
	return _Lockingpool.Contract.Receive(&_Lockingpool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lockingpool *LockingpoolTransactorSession) Receive() (*types.Transaction, error) {
	return _Lockingpool.Contract.Receive(&_Lockingpool.TransactOpts)
}

// LockingpoolAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Lockingpool contract.
type LockingpoolAdminChangedIterator struct {
	Event *LockingpoolAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolAdminChanged represents a AdminChanged event raised by the Lockingpool contract.
type LockingpoolAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Lockingpool *LockingpoolFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LockingpoolAdminChangedIterator, error) {

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LockingpoolAdminChangedIterator{contract: _Lockingpool.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Lockingpool *LockingpoolFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LockingpoolAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolAdminChanged)
				if err := _Lockingpool.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Lockingpool *LockingpoolFilterer) ParseAdminChanged(log types.Log) (*LockingpoolAdminChanged, error) {
	event := new(LockingpoolAdminChanged)
	if err := _Lockingpool.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Lockingpool contract.
type LockingpoolBeaconUpgradedIterator struct {
	Event *LockingpoolBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolBeaconUpgraded represents a BeaconUpgraded event raised by the Lockingpool contract.
type LockingpoolBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Lockingpool *LockingpoolFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LockingpoolBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LockingpoolBeaconUpgradedIterator{contract: _Lockingpool.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Lockingpool *LockingpoolFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LockingpoolBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolBeaconUpgraded)
				if err := _Lockingpool.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Lockingpool *LockingpoolFilterer) ParseBeaconUpgraded(log types.Log) (*LockingpoolBeaconUpgraded, error) {
	event := new(LockingpoolBeaconUpgraded)
	if err := _Lockingpool.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Lockingpool contract.
type LockingpoolInitializedIterator struct {
	Event *LockingpoolInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolInitialized represents a Initialized event raised by the Lockingpool contract.
type LockingpoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lockingpool *LockingpoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*LockingpoolInitializedIterator, error) {

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LockingpoolInitializedIterator{contract: _Lockingpool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lockingpool *LockingpoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LockingpoolInitialized) (event.Subscription, error) {

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolInitialized)
				if err := _Lockingpool.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lockingpool *LockingpoolFilterer) ParseInitialized(log types.Log) (*LockingpoolInitialized, error) {
	event := new(LockingpoolInitialized)
	if err := _Lockingpool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lockingpool contract.
type LockingpoolOwnershipTransferredIterator struct {
	Event *LockingpoolOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolOwnershipTransferred represents a OwnershipTransferred event raised by the Lockingpool contract.
type LockingpoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lockingpool *LockingpoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LockingpoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LockingpoolOwnershipTransferredIterator{contract: _Lockingpool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lockingpool *LockingpoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockingpoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolOwnershipTransferred)
				if err := _Lockingpool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lockingpool *LockingpoolFilterer) ParseOwnershipTransferred(log types.Log) (*LockingpoolOwnershipTransferred, error) {
	event := new(LockingpoolOwnershipTransferred)
	if err := _Lockingpool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Lockingpool contract.
type LockingpoolPausedIterator struct {
	Event *LockingpoolPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolPaused represents a Paused event raised by the Lockingpool contract.
type LockingpoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lockingpool *LockingpoolFilterer) FilterPaused(opts *bind.FilterOpts) (*LockingpoolPausedIterator, error) {

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LockingpoolPausedIterator{contract: _Lockingpool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lockingpool *LockingpoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LockingpoolPaused) (event.Subscription, error) {

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolPaused)
				if err := _Lockingpool.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lockingpool *LockingpoolFilterer) ParsePaused(log types.Log) (*LockingpoolPaused, error) {
	event := new(LockingpoolPaused)
	if err := _Lockingpool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolRewardUpdateIterator is returned from FilterRewardUpdate and is used to iterate over the raw logs and unpacked data for RewardUpdate events raised by the Lockingpool contract.
type LockingpoolRewardUpdateIterator struct {
	Event *LockingpoolRewardUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolRewardUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolRewardUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolRewardUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolRewardUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolRewardUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolRewardUpdate represents a RewardUpdate event raised by the Lockingpool contract.
type LockingpoolRewardUpdate struct {
	NewReward *big.Int
	OldReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardUpdate is a free log retrieval operation binding the contract event 0xf67f33e8589d3ea0356303c0f9a8e764873692159f777ff79e4fc523d389dfcd.
//
// Solidity: event RewardUpdate(uint256 newReward, uint256 oldReward)
func (_Lockingpool *LockingpoolFilterer) FilterRewardUpdate(opts *bind.FilterOpts) (*LockingpoolRewardUpdateIterator, error) {

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "RewardUpdate")
	if err != nil {
		return nil, err
	}
	return &LockingpoolRewardUpdateIterator{contract: _Lockingpool.contract, event: "RewardUpdate", logs: logs, sub: sub}, nil
}

// WatchRewardUpdate is a free log subscription operation binding the contract event 0xf67f33e8589d3ea0356303c0f9a8e764873692159f777ff79e4fc523d389dfcd.
//
// Solidity: event RewardUpdate(uint256 newReward, uint256 oldReward)
func (_Lockingpool *LockingpoolFilterer) WatchRewardUpdate(opts *bind.WatchOpts, sink chan<- *LockingpoolRewardUpdate) (event.Subscription, error) {

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "RewardUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolRewardUpdate)
				if err := _Lockingpool.contract.UnpackLog(event, "RewardUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardUpdate is a log parse operation binding the contract event 0xf67f33e8589d3ea0356303c0f9a8e764873692159f777ff79e4fc523d389dfcd.
//
// Solidity: event RewardUpdate(uint256 newReward, uint256 oldReward)
func (_Lockingpool *LockingpoolFilterer) ParseRewardUpdate(log types.Log) (*LockingpoolRewardUpdate, error) {
	event := new(LockingpoolRewardUpdate)
	if err := _Lockingpool.contract.UnpackLog(event, "RewardUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolSequencerOwnerChangedIterator is returned from FilterSequencerOwnerChanged and is used to iterate over the raw logs and unpacked data for SequencerOwnerChanged events raised by the Lockingpool contract.
type LockingpoolSequencerOwnerChangedIterator struct {
	Event *LockingpoolSequencerOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolSequencerOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolSequencerOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolSequencerOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolSequencerOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolSequencerOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolSequencerOwnerChanged represents a SequencerOwnerChanged event raised by the Lockingpool contract.
type LockingpoolSequencerOwnerChanged struct {
	SeqId *big.Int
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSequencerOwnerChanged is a free log retrieval operation binding the contract event 0x4078101d3657bf2f1ee46f64d5c94266d244d71bb0daa460336d3d6f11c9a4ac.
//
// Solidity: event SequencerOwnerChanged(uint256 _seqId, address _owner)
func (_Lockingpool *LockingpoolFilterer) FilterSequencerOwnerChanged(opts *bind.FilterOpts) (*LockingpoolSequencerOwnerChangedIterator, error) {

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "SequencerOwnerChanged")
	if err != nil {
		return nil, err
	}
	return &LockingpoolSequencerOwnerChangedIterator{contract: _Lockingpool.contract, event: "SequencerOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchSequencerOwnerChanged is a free log subscription operation binding the contract event 0x4078101d3657bf2f1ee46f64d5c94266d244d71bb0daa460336d3d6f11c9a4ac.
//
// Solidity: event SequencerOwnerChanged(uint256 _seqId, address _owner)
func (_Lockingpool *LockingpoolFilterer) WatchSequencerOwnerChanged(opts *bind.WatchOpts, sink chan<- *LockingpoolSequencerOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "SequencerOwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolSequencerOwnerChanged)
				if err := _Lockingpool.contract.UnpackLog(event, "SequencerOwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerOwnerChanged is a log parse operation binding the contract event 0x4078101d3657bf2f1ee46f64d5c94266d244d71bb0daa460336d3d6f11c9a4ac.
//
// Solidity: event SequencerOwnerChanged(uint256 _seqId, address _owner)
func (_Lockingpool *LockingpoolFilterer) ParseSequencerOwnerChanged(log types.Log) (*LockingpoolSequencerOwnerChanged, error) {
	event := new(LockingpoolSequencerOwnerChanged)
	if err := _Lockingpool.contract.UnpackLog(event, "SequencerOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolSequencerRewardRecipientChangedIterator is returned from FilterSequencerRewardRecipientChanged and is used to iterate over the raw logs and unpacked data for SequencerRewardRecipientChanged events raised by the Lockingpool contract.
type LockingpoolSequencerRewardRecipientChangedIterator struct {
	Event *LockingpoolSequencerRewardRecipientChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolSequencerRewardRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolSequencerRewardRecipientChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolSequencerRewardRecipientChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolSequencerRewardRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolSequencerRewardRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolSequencerRewardRecipientChanged represents a SequencerRewardRecipientChanged event raised by the Lockingpool contract.
type LockingpoolSequencerRewardRecipientChanged struct {
	SeqId     *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSequencerRewardRecipientChanged is a free log retrieval operation binding the contract event 0x357bb123cabaf224688e3d8de5feb37d685dc3a27012a7bce1201c49bc369cb6.
//
// Solidity: event SequencerRewardRecipientChanged(uint256 _seqId, address _recipient)
func (_Lockingpool *LockingpoolFilterer) FilterSequencerRewardRecipientChanged(opts *bind.FilterOpts) (*LockingpoolSequencerRewardRecipientChangedIterator, error) {

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "SequencerRewardRecipientChanged")
	if err != nil {
		return nil, err
	}
	return &LockingpoolSequencerRewardRecipientChangedIterator{contract: _Lockingpool.contract, event: "SequencerRewardRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchSequencerRewardRecipientChanged is a free log subscription operation binding the contract event 0x357bb123cabaf224688e3d8de5feb37d685dc3a27012a7bce1201c49bc369cb6.
//
// Solidity: event SequencerRewardRecipientChanged(uint256 _seqId, address _recipient)
func (_Lockingpool *LockingpoolFilterer) WatchSequencerRewardRecipientChanged(opts *bind.WatchOpts, sink chan<- *LockingpoolSequencerRewardRecipientChanged) (event.Subscription, error) {

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "SequencerRewardRecipientChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolSequencerRewardRecipientChanged)
				if err := _Lockingpool.contract.UnpackLog(event, "SequencerRewardRecipientChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerRewardRecipientChanged is a log parse operation binding the contract event 0x357bb123cabaf224688e3d8de5feb37d685dc3a27012a7bce1201c49bc369cb6.
//
// Solidity: event SequencerRewardRecipientChanged(uint256 _seqId, address _recipient)
func (_Lockingpool *LockingpoolFilterer) ParseSequencerRewardRecipientChanged(log types.Log) (*LockingpoolSequencerRewardRecipientChanged, error) {
	event := new(LockingpoolSequencerRewardRecipientChanged)
	if err := _Lockingpool.contract.UnpackLog(event, "SequencerRewardRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolSetSignerUpdateThrottleIterator is returned from FilterSetSignerUpdateThrottle and is used to iterate over the raw logs and unpacked data for SetSignerUpdateThrottle events raised by the Lockingpool contract.
type LockingpoolSetSignerUpdateThrottleIterator struct {
	Event *LockingpoolSetSignerUpdateThrottle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolSetSignerUpdateThrottleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolSetSignerUpdateThrottle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolSetSignerUpdateThrottle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolSetSignerUpdateThrottleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolSetSignerUpdateThrottleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolSetSignerUpdateThrottle represents a SetSignerUpdateThrottle event raised by the Lockingpool contract.
type LockingpoolSetSignerUpdateThrottle struct {
	N   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetSignerUpdateThrottle is a free log retrieval operation binding the contract event 0xe58685f6b78e6d567d2ed9d7c5efb779c4cd91c63c76763a0e3204a5671f4705.
//
// Solidity: event SetSignerUpdateThrottle(uint256 _n)
func (_Lockingpool *LockingpoolFilterer) FilterSetSignerUpdateThrottle(opts *bind.FilterOpts) (*LockingpoolSetSignerUpdateThrottleIterator, error) {

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "SetSignerUpdateThrottle")
	if err != nil {
		return nil, err
	}
	return &LockingpoolSetSignerUpdateThrottleIterator{contract: _Lockingpool.contract, event: "SetSignerUpdateThrottle", logs: logs, sub: sub}, nil
}

// WatchSetSignerUpdateThrottle is a free log subscription operation binding the contract event 0xe58685f6b78e6d567d2ed9d7c5efb779c4cd91c63c76763a0e3204a5671f4705.
//
// Solidity: event SetSignerUpdateThrottle(uint256 _n)
func (_Lockingpool *LockingpoolFilterer) WatchSetSignerUpdateThrottle(opts *bind.WatchOpts, sink chan<- *LockingpoolSetSignerUpdateThrottle) (event.Subscription, error) {

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "SetSignerUpdateThrottle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolSetSignerUpdateThrottle)
				if err := _Lockingpool.contract.UnpackLog(event, "SetSignerUpdateThrottle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetSignerUpdateThrottle is a log parse operation binding the contract event 0xe58685f6b78e6d567d2ed9d7c5efb779c4cd91c63c76763a0e3204a5671f4705.
//
// Solidity: event SetSignerUpdateThrottle(uint256 _n)
func (_Lockingpool *LockingpoolFilterer) ParseSetSignerUpdateThrottle(log types.Log) (*LockingpoolSetSignerUpdateThrottle, error) {
	event := new(LockingpoolSetSignerUpdateThrottle)
	if err := _Lockingpool.contract.UnpackLog(event, "SetSignerUpdateThrottle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolSetWhitelistIterator is returned from FilterSetWhitelist and is used to iterate over the raw logs and unpacked data for SetWhitelist events raised by the Lockingpool contract.
type LockingpoolSetWhitelistIterator struct {
	Event *LockingpoolSetWhitelist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolSetWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolSetWhitelist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolSetWhitelist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolSetWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolSetWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolSetWhitelist represents a SetWhitelist event raised by the Lockingpool contract.
type LockingpoolSetWhitelist struct {
	User common.Address
	Yes  bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetWhitelist is a free log retrieval operation binding the contract event 0xf6019ec0a78d156d249a1ec7579e2321f6ac7521d6e1d2eacf90ba4a184dcceb.
//
// Solidity: event SetWhitelist(address _user, bool _yes)
func (_Lockingpool *LockingpoolFilterer) FilterSetWhitelist(opts *bind.FilterOpts) (*LockingpoolSetWhitelistIterator, error) {

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "SetWhitelist")
	if err != nil {
		return nil, err
	}
	return &LockingpoolSetWhitelistIterator{contract: _Lockingpool.contract, event: "SetWhitelist", logs: logs, sub: sub}, nil
}

// WatchSetWhitelist is a free log subscription operation binding the contract event 0xf6019ec0a78d156d249a1ec7579e2321f6ac7521d6e1d2eacf90ba4a184dcceb.
//
// Solidity: event SetWhitelist(address _user, bool _yes)
func (_Lockingpool *LockingpoolFilterer) WatchSetWhitelist(opts *bind.WatchOpts, sink chan<- *LockingpoolSetWhitelist) (event.Subscription, error) {

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "SetWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolSetWhitelist)
				if err := _Lockingpool.contract.UnpackLog(event, "SetWhitelist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetWhitelist is a log parse operation binding the contract event 0xf6019ec0a78d156d249a1ec7579e2321f6ac7521d6e1d2eacf90ba4a184dcceb.
//
// Solidity: event SetWhitelist(address _user, bool _yes)
func (_Lockingpool *LockingpoolFilterer) ParseSetWhitelist(log types.Log) (*LockingpoolSetWhitelist, error) {
	event := new(LockingpoolSetWhitelist)
	if err := _Lockingpool.contract.UnpackLog(event, "SetWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Lockingpool contract.
type LockingpoolUnpausedIterator struct {
	Event *LockingpoolUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolUnpaused represents a Unpaused event raised by the Lockingpool contract.
type LockingpoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lockingpool *LockingpoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LockingpoolUnpausedIterator, error) {

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LockingpoolUnpausedIterator{contract: _Lockingpool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lockingpool *LockingpoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LockingpoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolUnpaused)
				if err := _Lockingpool.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lockingpool *LockingpoolFilterer) ParseUnpaused(log types.Log) (*LockingpoolUnpaused, error) {
	event := new(LockingpoolUnpaused)
	if err := _Lockingpool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolUpdateMpcIterator is returned from FilterUpdateMpc and is used to iterate over the raw logs and unpacked data for UpdateMpc events raised by the Lockingpool contract.
type LockingpoolUpdateMpcIterator struct {
	Event *LockingpoolUpdateMpc // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolUpdateMpcIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolUpdateMpc)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolUpdateMpc)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolUpdateMpcIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolUpdateMpcIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolUpdateMpc represents a UpdateMpc event raised by the Lockingpool contract.
type LockingpoolUpdateMpc struct {
	NewMpc common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateMpc is a free log retrieval operation binding the contract event 0xc6759872346bc2093e270f2fa00d99d7bcdde60a410a3e9956b34196d42fee76.
//
// Solidity: event UpdateMpc(address _newMpc)
func (_Lockingpool *LockingpoolFilterer) FilterUpdateMpc(opts *bind.FilterOpts) (*LockingpoolUpdateMpcIterator, error) {

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "UpdateMpc")
	if err != nil {
		return nil, err
	}
	return &LockingpoolUpdateMpcIterator{contract: _Lockingpool.contract, event: "UpdateMpc", logs: logs, sub: sub}, nil
}

// WatchUpdateMpc is a free log subscription operation binding the contract event 0xc6759872346bc2093e270f2fa00d99d7bcdde60a410a3e9956b34196d42fee76.
//
// Solidity: event UpdateMpc(address _newMpc)
func (_Lockingpool *LockingpoolFilterer) WatchUpdateMpc(opts *bind.WatchOpts, sink chan<- *LockingpoolUpdateMpc) (event.Subscription, error) {

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "UpdateMpc")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolUpdateMpc)
				if err := _Lockingpool.contract.UnpackLog(event, "UpdateMpc", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateMpc is a log parse operation binding the contract event 0xc6759872346bc2093e270f2fa00d99d7bcdde60a410a3e9956b34196d42fee76.
//
// Solidity: event UpdateMpc(address _newMpc)
func (_Lockingpool *LockingpoolFilterer) ParseUpdateMpc(log types.Log) (*LockingpoolUpdateMpc, error) {
	event := new(LockingpoolUpdateMpc)
	if err := _Lockingpool.contract.UnpackLog(event, "UpdateMpc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Lockingpool contract.
type LockingpoolUpgradedIterator struct {
	Event *LockingpoolUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolUpgraded represents a Upgraded event raised by the Lockingpool contract.
type LockingpoolUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Lockingpool *LockingpoolFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LockingpoolUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LockingpoolUpgradedIterator{contract: _Lockingpool.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Lockingpool *LockingpoolFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LockingpoolUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolUpgraded)
				if err := _Lockingpool.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Lockingpool *LockingpoolFilterer) ParseUpgraded(log types.Log) (*LockingpoolUpgraded, error) {
	event := new(LockingpoolUpgraded)
	if err := _Lockingpool.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingpoolWithrawDelayTimeChangeIterator is returned from FilterWithrawDelayTimeChange and is used to iterate over the raw logs and unpacked data for WithrawDelayTimeChange events raised by the Lockingpool contract.
type LockingpoolWithrawDelayTimeChangeIterator struct {
	Event *LockingpoolWithrawDelayTimeChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingpoolWithrawDelayTimeChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingpoolWithrawDelayTimeChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingpoolWithrawDelayTimeChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingpoolWithrawDelayTimeChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingpoolWithrawDelayTimeChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingpoolWithrawDelayTimeChange represents a WithrawDelayTimeChange event raised by the Lockingpool contract.
type LockingpoolWithrawDelayTimeChange struct {
	Cur  *big.Int
	Prev *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWithrawDelayTimeChange is a free log retrieval operation binding the contract event 0x08cb0bf599c925a6240976039d9d4d3201d88d2ba83703b890049356cdbb67e6.
//
// Solidity: event WithrawDelayTimeChange(uint256 _cur, uint256 _prev)
func (_Lockingpool *LockingpoolFilterer) FilterWithrawDelayTimeChange(opts *bind.FilterOpts) (*LockingpoolWithrawDelayTimeChangeIterator, error) {

	logs, sub, err := _Lockingpool.contract.FilterLogs(opts, "WithrawDelayTimeChange")
	if err != nil {
		return nil, err
	}
	return &LockingpoolWithrawDelayTimeChangeIterator{contract: _Lockingpool.contract, event: "WithrawDelayTimeChange", logs: logs, sub: sub}, nil
}

// WatchWithrawDelayTimeChange is a free log subscription operation binding the contract event 0x08cb0bf599c925a6240976039d9d4d3201d88d2ba83703b890049356cdbb67e6.
//
// Solidity: event WithrawDelayTimeChange(uint256 _cur, uint256 _prev)
func (_Lockingpool *LockingpoolFilterer) WatchWithrawDelayTimeChange(opts *bind.WatchOpts, sink chan<- *LockingpoolWithrawDelayTimeChange) (event.Subscription, error) {

	logs, sub, err := _Lockingpool.contract.WatchLogs(opts, "WithrawDelayTimeChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingpoolWithrawDelayTimeChange)
				if err := _Lockingpool.contract.UnpackLog(event, "WithrawDelayTimeChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithrawDelayTimeChange is a log parse operation binding the contract event 0x08cb0bf599c925a6240976039d9d4d3201d88d2ba83703b890049356cdbb67e6.
//
// Solidity: event WithrawDelayTimeChange(uint256 _cur, uint256 _prev)
func (_Lockingpool *LockingpoolFilterer) ParseWithrawDelayTimeChange(log types.Log) (*LockingpoolWithrawDelayTimeChange, error) {
	event := new(LockingpoolWithrawDelayTimeChange)
	if err := _Lockingpool.contract.UnpackLog(event, "WithrawDelayTimeChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

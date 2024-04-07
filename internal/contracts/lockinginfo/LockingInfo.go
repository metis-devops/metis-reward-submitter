// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lockinginfo

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

// ISequencerInfoSequencer is an auto generated low-level Go binding around an user-defined struct.
type ISequencerInfoSequencer struct {
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
}

// LockinginfoMetaData contains all meta data concerning the Lockinginfo contract.
var LockinginfoMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newBatchId\",\"type\":\"uint256\"}],\"name\":\"BatchSubmitReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequencerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"ClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequencerId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"LockUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequencerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"activationBatch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequencerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"Relocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newMaxLock\",\"type\":\"uint256\"}],\"name\":\"SetMaxLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newMinLock\",\"type\":\"uint256\"}],\"name\":\"SetMinLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_payer\",\"type\":\"address\"}],\"name\":\"SetRewardPayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequencerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"SignerChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequencerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivationBatch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockClaimTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequencerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalReward\",\"type\":\"uint256\"}],\"name\":\"distributeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_l2gas\",\"type\":\"uint32\"}],\"name\":\"finalizeUnlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_locked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_incoming\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromReward\",\"type\":\"uint256\"}],\"name\":\"increaseLocked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"initManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_l2ChainId\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2gas\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activationBatch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedBatch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivationBatch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockClaimTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"},{\"internalType\":\"enumISequencerInfo.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structISequencerInfo.Sequencer\",\"name\":\"_seq\",\"type\":\"tuple\"}],\"name\":\"initializeUnlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seqId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_l2gas\",\"type\":\"uint32\"}],\"name\":\"liquidateReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequencerId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"logSignerChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signerPubkey\",\"type\":\"bytes\"}],\"name\":\"newSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLock\",\"type\":\"uint256\"}],\"name\":\"setMaxLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minLock\",\"type\":\"uint256\"}],\"name\":\"setMinLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_payer\",\"type\":\"address\"}],\"name\":\"setRewardPayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRewardsLiquidated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"}]",
}

// LockinginfoABI is the input ABI used to generate the binding from.
// Deprecated: Use LockinginfoMetaData.ABI instead.
var LockinginfoABI = LockinginfoMetaData.ABI

// Lockinginfo is an auto generated Go binding around an Ethereum contract.
type Lockinginfo struct {
	LockinginfoCaller     // Read-only binding to the contract
	LockinginfoTransactor // Write-only binding to the contract
	LockinginfoFilterer   // Log filterer for contract events
}

// LockinginfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockinginfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockinginfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockinginfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockinginfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockinginfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockinginfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockinginfoSession struct {
	Contract     *Lockinginfo      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockinginfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockinginfoCallerSession struct {
	Contract *LockinginfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LockinginfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockinginfoTransactorSession struct {
	Contract     *LockinginfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LockinginfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockinginfoRaw struct {
	Contract *Lockinginfo // Generic contract binding to access the raw methods on
}

// LockinginfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockinginfoCallerRaw struct {
	Contract *LockinginfoCaller // Generic read-only contract binding to access the raw methods on
}

// LockinginfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockinginfoTransactorRaw struct {
	Contract *LockinginfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockinginfo creates a new instance of Lockinginfo, bound to a specific deployed contract.
func NewLockinginfo(address common.Address, backend bind.ContractBackend) (*Lockinginfo, error) {
	contract, err := bindLockinginfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lockinginfo{LockinginfoCaller: LockinginfoCaller{contract: contract}, LockinginfoTransactor: LockinginfoTransactor{contract: contract}, LockinginfoFilterer: LockinginfoFilterer{contract: contract}}, nil
}

// NewLockinginfoCaller creates a new read-only instance of Lockinginfo, bound to a specific deployed contract.
func NewLockinginfoCaller(address common.Address, caller bind.ContractCaller) (*LockinginfoCaller, error) {
	contract, err := bindLockinginfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockinginfoCaller{contract: contract}, nil
}

// NewLockinginfoTransactor creates a new write-only instance of Lockinginfo, bound to a specific deployed contract.
func NewLockinginfoTransactor(address common.Address, transactor bind.ContractTransactor) (*LockinginfoTransactor, error) {
	contract, err := bindLockinginfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockinginfoTransactor{contract: contract}, nil
}

// NewLockinginfoFilterer creates a new log filterer instance of Lockinginfo, bound to a specific deployed contract.
func NewLockinginfoFilterer(address common.Address, filterer bind.ContractFilterer) (*LockinginfoFilterer, error) {
	contract, err := bindLockinginfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockinginfoFilterer{contract: contract}, nil
}

// bindLockinginfo binds a generic wrapper to an already deployed contract.
func bindLockinginfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LockinginfoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lockinginfo *LockinginfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lockinginfo.Contract.LockinginfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lockinginfo *LockinginfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockinginfo.Contract.LockinginfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lockinginfo *LockinginfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lockinginfo.Contract.LockinginfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lockinginfo *LockinginfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lockinginfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lockinginfo *LockinginfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockinginfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lockinginfo *LockinginfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lockinginfo.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Lockinginfo *LockinginfoCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lockinginfo.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Lockinginfo *LockinginfoSession) Bridge() (common.Address, error) {
	return _Lockinginfo.Contract.Bridge(&_Lockinginfo.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Lockinginfo *LockinginfoCallerSession) Bridge() (common.Address, error) {
	return _Lockinginfo.Contract.Bridge(&_Lockinginfo.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_Lockinginfo *LockinginfoCaller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lockinginfo.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_Lockinginfo *LockinginfoSession) L1Token() (common.Address, error) {
	return _Lockinginfo.Contract.L1Token(&_Lockinginfo.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_Lockinginfo *LockinginfoCallerSession) L1Token() (common.Address, error) {
	return _Lockinginfo.Contract.L1Token(&_Lockinginfo.CallOpts)
}

// L2ChainId is a free data retrieval call binding the contract method 0xd6ae3cd5.
//
// Solidity: function l2ChainId() view returns(uint256)
func (_Lockinginfo *LockinginfoCaller) L2ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lockinginfo.contract.Call(opts, &out, "l2ChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ChainId is a free data retrieval call binding the contract method 0xd6ae3cd5.
//
// Solidity: function l2ChainId() view returns(uint256)
func (_Lockinginfo *LockinginfoSession) L2ChainId() (*big.Int, error) {
	return _Lockinginfo.Contract.L2ChainId(&_Lockinginfo.CallOpts)
}

// L2ChainId is a free data retrieval call binding the contract method 0xd6ae3cd5.
//
// Solidity: function l2ChainId() view returns(uint256)
func (_Lockinginfo *LockinginfoCallerSession) L2ChainId() (*big.Int, error) {
	return _Lockinginfo.Contract.L2ChainId(&_Lockinginfo.CallOpts)
}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_Lockinginfo *LockinginfoCaller) L2Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lockinginfo.contract.Call(opts, &out, "l2Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_Lockinginfo *LockinginfoSession) L2Token() (common.Address, error) {
	return _Lockinginfo.Contract.L2Token(&_Lockinginfo.CallOpts)
}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_Lockinginfo *LockinginfoCallerSession) L2Token() (common.Address, error) {
	return _Lockinginfo.Contract.L2Token(&_Lockinginfo.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Lockinginfo *LockinginfoCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lockinginfo.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Lockinginfo *LockinginfoSession) Manager() (common.Address, error) {
	return _Lockinginfo.Contract.Manager(&_Lockinginfo.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Lockinginfo *LockinginfoCallerSession) Manager() (common.Address, error) {
	return _Lockinginfo.Contract.Manager(&_Lockinginfo.CallOpts)
}

// MaxLock is a free data retrieval call binding the contract method 0x6c0b3e46.
//
// Solidity: function maxLock() view returns(uint256)
func (_Lockinginfo *LockinginfoCaller) MaxLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lockinginfo.contract.Call(opts, &out, "maxLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLock is a free data retrieval call binding the contract method 0x6c0b3e46.
//
// Solidity: function maxLock() view returns(uint256)
func (_Lockinginfo *LockinginfoSession) MaxLock() (*big.Int, error) {
	return _Lockinginfo.Contract.MaxLock(&_Lockinginfo.CallOpts)
}

// MaxLock is a free data retrieval call binding the contract method 0x6c0b3e46.
//
// Solidity: function maxLock() view returns(uint256)
func (_Lockinginfo *LockinginfoCallerSession) MaxLock() (*big.Int, error) {
	return _Lockinginfo.Contract.MaxLock(&_Lockinginfo.CallOpts)
}

// MinLock is a free data retrieval call binding the contract method 0xf037c630.
//
// Solidity: function minLock() view returns(uint256)
func (_Lockinginfo *LockinginfoCaller) MinLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lockinginfo.contract.Call(opts, &out, "minLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinLock is a free data retrieval call binding the contract method 0xf037c630.
//
// Solidity: function minLock() view returns(uint256)
func (_Lockinginfo *LockinginfoSession) MinLock() (*big.Int, error) {
	return _Lockinginfo.Contract.MinLock(&_Lockinginfo.CallOpts)
}

// MinLock is a free data retrieval call binding the contract method 0xf037c630.
//
// Solidity: function minLock() view returns(uint256)
func (_Lockinginfo *LockinginfoCallerSession) MinLock() (*big.Int, error) {
	return _Lockinginfo.Contract.MinLock(&_Lockinginfo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lockinginfo *LockinginfoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lockinginfo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lockinginfo *LockinginfoSession) Owner() (common.Address, error) {
	return _Lockinginfo.Contract.Owner(&_Lockinginfo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lockinginfo *LockinginfoCallerSession) Owner() (common.Address, error) {
	return _Lockinginfo.Contract.Owner(&_Lockinginfo.CallOpts)
}

// RewardPayer is a free data retrieval call binding the contract method 0x6eb27154.
//
// Solidity: function rewardPayer() view returns(address)
func (_Lockinginfo *LockinginfoCaller) RewardPayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lockinginfo.contract.Call(opts, &out, "rewardPayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardPayer is a free data retrieval call binding the contract method 0x6eb27154.
//
// Solidity: function rewardPayer() view returns(address)
func (_Lockinginfo *LockinginfoSession) RewardPayer() (common.Address, error) {
	return _Lockinginfo.Contract.RewardPayer(&_Lockinginfo.CallOpts)
}

// RewardPayer is a free data retrieval call binding the contract method 0x6eb27154.
//
// Solidity: function rewardPayer() view returns(address)
func (_Lockinginfo *LockinginfoCallerSession) RewardPayer() (common.Address, error) {
	return _Lockinginfo.Contract.RewardPayer(&_Lockinginfo.CallOpts)
}

// TotalLocked is a free data retrieval call binding the contract method 0x56891412.
//
// Solidity: function totalLocked() view returns(uint256)
func (_Lockinginfo *LockinginfoCaller) TotalLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lockinginfo.contract.Call(opts, &out, "totalLocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLocked is a free data retrieval call binding the contract method 0x56891412.
//
// Solidity: function totalLocked() view returns(uint256)
func (_Lockinginfo *LockinginfoSession) TotalLocked() (*big.Int, error) {
	return _Lockinginfo.Contract.TotalLocked(&_Lockinginfo.CallOpts)
}

// TotalLocked is a free data retrieval call binding the contract method 0x56891412.
//
// Solidity: function totalLocked() view returns(uint256)
func (_Lockinginfo *LockinginfoCallerSession) TotalLocked() (*big.Int, error) {
	return _Lockinginfo.Contract.TotalLocked(&_Lockinginfo.CallOpts)
}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() view returns(uint256)
func (_Lockinginfo *LockinginfoCaller) TotalRewardsLiquidated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lockinginfo.contract.Call(opts, &out, "totalRewardsLiquidated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() view returns(uint256)
func (_Lockinginfo *LockinginfoSession) TotalRewardsLiquidated() (*big.Int, error) {
	return _Lockinginfo.Contract.TotalRewardsLiquidated(&_Lockinginfo.CallOpts)
}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() view returns(uint256)
func (_Lockinginfo *LockinginfoCallerSession) TotalRewardsLiquidated() (*big.Int, error) {
	return _Lockinginfo.Contract.TotalRewardsLiquidated(&_Lockinginfo.CallOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Lockinginfo *LockinginfoTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Lockinginfo *LockinginfoSession) Admin() (*types.Transaction, error) {
	return _Lockinginfo.Contract.Admin(&_Lockinginfo.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Lockinginfo *LockinginfoTransactorSession) Admin() (*types.Transaction, error) {
	return _Lockinginfo.Contract.Admin(&_Lockinginfo.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Lockinginfo *LockinginfoTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Lockinginfo *LockinginfoSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Lockinginfo.Contract.ChangeAdmin(&_Lockinginfo.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Lockinginfo *LockinginfoTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Lockinginfo.Contract.ChangeAdmin(&_Lockinginfo.TransactOpts, newAdmin)
}

// DistributeReward is a paid mutator transaction binding the contract method 0xe3bcd27c.
//
// Solidity: function distributeReward(uint256 _batchId, uint256 _totalReward) returns()
func (_Lockinginfo *LockinginfoTransactor) DistributeReward(opts *bind.TransactOpts, _batchId *big.Int, _totalReward *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "distributeReward", _batchId, _totalReward)
}

// DistributeReward is a paid mutator transaction binding the contract method 0xe3bcd27c.
//
// Solidity: function distributeReward(uint256 _batchId, uint256 _totalReward) returns()
func (_Lockinginfo *LockinginfoSession) DistributeReward(_batchId *big.Int, _totalReward *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.Contract.DistributeReward(&_Lockinginfo.TransactOpts, _batchId, _totalReward)
}

// DistributeReward is a paid mutator transaction binding the contract method 0xe3bcd27c.
//
// Solidity: function distributeReward(uint256 _batchId, uint256 _totalReward) returns()
func (_Lockinginfo *LockinginfoTransactorSession) DistributeReward(_batchId *big.Int, _totalReward *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.Contract.DistributeReward(&_Lockinginfo.TransactOpts, _batchId, _totalReward)
}

// FinalizeUnlock is a paid mutator transaction binding the contract method 0x528ed12a.
//
// Solidity: function finalizeUnlock(address _operator, uint256 _seqId, uint256 _amount, uint256 _reward, address _recipient, uint32 _l2gas) payable returns()
func (_Lockinginfo *LockinginfoTransactor) FinalizeUnlock(opts *bind.TransactOpts, _operator common.Address, _seqId *big.Int, _amount *big.Int, _reward *big.Int, _recipient common.Address, _l2gas uint32) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "finalizeUnlock", _operator, _seqId, _amount, _reward, _recipient, _l2gas)
}

// FinalizeUnlock is a paid mutator transaction binding the contract method 0x528ed12a.
//
// Solidity: function finalizeUnlock(address _operator, uint256 _seqId, uint256 _amount, uint256 _reward, address _recipient, uint32 _l2gas) payable returns()
func (_Lockinginfo *LockinginfoSession) FinalizeUnlock(_operator common.Address, _seqId *big.Int, _amount *big.Int, _reward *big.Int, _recipient common.Address, _l2gas uint32) (*types.Transaction, error) {
	return _Lockinginfo.Contract.FinalizeUnlock(&_Lockinginfo.TransactOpts, _operator, _seqId, _amount, _reward, _recipient, _l2gas)
}

// FinalizeUnlock is a paid mutator transaction binding the contract method 0x528ed12a.
//
// Solidity: function finalizeUnlock(address _operator, uint256 _seqId, uint256 _amount, uint256 _reward, address _recipient, uint32 _l2gas) payable returns()
func (_Lockinginfo *LockinginfoTransactorSession) FinalizeUnlock(_operator common.Address, _seqId *big.Int, _amount *big.Int, _reward *big.Int, _recipient common.Address, _l2gas uint32) (*types.Transaction, error) {
	return _Lockinginfo.Contract.FinalizeUnlock(&_Lockinginfo.TransactOpts, _operator, _seqId, _amount, _reward, _recipient, _l2gas)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Lockinginfo *LockinginfoTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Lockinginfo *LockinginfoSession) Implementation() (*types.Transaction, error) {
	return _Lockinginfo.Contract.Implementation(&_Lockinginfo.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Lockinginfo *LockinginfoTransactorSession) Implementation() (*types.Transaction, error) {
	return _Lockinginfo.Contract.Implementation(&_Lockinginfo.TransactOpts)
}

// IncreaseLocked is a paid mutator transaction binding the contract method 0x2684b8ec.
//
// Solidity: function increaseLocked(uint256 _seqId, uint256 _nonce, address _owner, uint256 _locked, uint256 _incoming, uint256 _fromReward) returns()
func (_Lockinginfo *LockinginfoTransactor) IncreaseLocked(opts *bind.TransactOpts, _seqId *big.Int, _nonce *big.Int, _owner common.Address, _locked *big.Int, _incoming *big.Int, _fromReward *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "increaseLocked", _seqId, _nonce, _owner, _locked, _incoming, _fromReward)
}

// IncreaseLocked is a paid mutator transaction binding the contract method 0x2684b8ec.
//
// Solidity: function increaseLocked(uint256 _seqId, uint256 _nonce, address _owner, uint256 _locked, uint256 _incoming, uint256 _fromReward) returns()
func (_Lockinginfo *LockinginfoSession) IncreaseLocked(_seqId *big.Int, _nonce *big.Int, _owner common.Address, _locked *big.Int, _incoming *big.Int, _fromReward *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.Contract.IncreaseLocked(&_Lockinginfo.TransactOpts, _seqId, _nonce, _owner, _locked, _incoming, _fromReward)
}

// IncreaseLocked is a paid mutator transaction binding the contract method 0x2684b8ec.
//
// Solidity: function increaseLocked(uint256 _seqId, uint256 _nonce, address _owner, uint256 _locked, uint256 _incoming, uint256 _fromReward) returns()
func (_Lockinginfo *LockinginfoTransactorSession) IncreaseLocked(_seqId *big.Int, _nonce *big.Int, _owner common.Address, _locked *big.Int, _incoming *big.Int, _fromReward *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.Contract.IncreaseLocked(&_Lockinginfo.TransactOpts, _seqId, _nonce, _owner, _locked, _incoming, _fromReward)
}

// InitManager is a paid mutator transaction binding the contract method 0xb1fc19d3.
//
// Solidity: function initManager(address _manager) returns()
func (_Lockinginfo *LockinginfoTransactor) InitManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "initManager", _manager)
}

// InitManager is a paid mutator transaction binding the contract method 0xb1fc19d3.
//
// Solidity: function initManager(address _manager) returns()
func (_Lockinginfo *LockinginfoSession) InitManager(_manager common.Address) (*types.Transaction, error) {
	return _Lockinginfo.Contract.InitManager(&_Lockinginfo.TransactOpts, _manager)
}

// InitManager is a paid mutator transaction binding the contract method 0xb1fc19d3.
//
// Solidity: function initManager(address _manager) returns()
func (_Lockinginfo *LockinginfoTransactorSession) InitManager(_manager common.Address) (*types.Transaction, error) {
	return _Lockinginfo.Contract.InitManager(&_Lockinginfo.TransactOpts, _manager)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _bridge, address _l1Token, address _l2Token, uint256 _l2ChainId) returns()
func (_Lockinginfo *LockinginfoTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address, _l1Token common.Address, _l2Token common.Address, _l2ChainId *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "initialize", _bridge, _l1Token, _l2Token, _l2ChainId)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _bridge, address _l1Token, address _l2Token, uint256 _l2ChainId) returns()
func (_Lockinginfo *LockinginfoSession) Initialize(_bridge common.Address, _l1Token common.Address, _l2Token common.Address, _l2ChainId *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.Contract.Initialize(&_Lockinginfo.TransactOpts, _bridge, _l1Token, _l2Token, _l2ChainId)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _bridge, address _l1Token, address _l2Token, uint256 _l2ChainId) returns()
func (_Lockinginfo *LockinginfoTransactorSession) Initialize(_bridge common.Address, _l1Token common.Address, _l2Token common.Address, _l2ChainId *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.Contract.Initialize(&_Lockinginfo.TransactOpts, _bridge, _l1Token, _l2Token, _l2ChainId)
}

// InitializeUnlock is a paid mutator transaction binding the contract method 0x2243069c.
//
// Solidity: function initializeUnlock(uint256 _seqId, uint256 _reward, uint32 _l2gas, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,bytes,address,uint8) _seq) payable returns()
func (_Lockinginfo *LockinginfoTransactor) InitializeUnlock(opts *bind.TransactOpts, _seqId *big.Int, _reward *big.Int, _l2gas uint32, _seq ISequencerInfoSequencer) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "initializeUnlock", _seqId, _reward, _l2gas, _seq)
}

// InitializeUnlock is a paid mutator transaction binding the contract method 0x2243069c.
//
// Solidity: function initializeUnlock(uint256 _seqId, uint256 _reward, uint32 _l2gas, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,bytes,address,uint8) _seq) payable returns()
func (_Lockinginfo *LockinginfoSession) InitializeUnlock(_seqId *big.Int, _reward *big.Int, _l2gas uint32, _seq ISequencerInfoSequencer) (*types.Transaction, error) {
	return _Lockinginfo.Contract.InitializeUnlock(&_Lockinginfo.TransactOpts, _seqId, _reward, _l2gas, _seq)
}

// InitializeUnlock is a paid mutator transaction binding the contract method 0x2243069c.
//
// Solidity: function initializeUnlock(uint256 _seqId, uint256 _reward, uint32 _l2gas, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,bytes,address,uint8) _seq) payable returns()
func (_Lockinginfo *LockinginfoTransactorSession) InitializeUnlock(_seqId *big.Int, _reward *big.Int, _l2gas uint32, _seq ISequencerInfoSequencer) (*types.Transaction, error) {
	return _Lockinginfo.Contract.InitializeUnlock(&_Lockinginfo.TransactOpts, _seqId, _reward, _l2gas, _seq)
}

// LiquidateReward is a paid mutator transaction binding the contract method 0x5d7878a8.
//
// Solidity: function liquidateReward(uint256 _seqId, uint256 _amount, address _recipient, uint32 _l2gas) payable returns()
func (_Lockinginfo *LockinginfoTransactor) LiquidateReward(opts *bind.TransactOpts, _seqId *big.Int, _amount *big.Int, _recipient common.Address, _l2gas uint32) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "liquidateReward", _seqId, _amount, _recipient, _l2gas)
}

// LiquidateReward is a paid mutator transaction binding the contract method 0x5d7878a8.
//
// Solidity: function liquidateReward(uint256 _seqId, uint256 _amount, address _recipient, uint32 _l2gas) payable returns()
func (_Lockinginfo *LockinginfoSession) LiquidateReward(_seqId *big.Int, _amount *big.Int, _recipient common.Address, _l2gas uint32) (*types.Transaction, error) {
	return _Lockinginfo.Contract.LiquidateReward(&_Lockinginfo.TransactOpts, _seqId, _amount, _recipient, _l2gas)
}

// LiquidateReward is a paid mutator transaction binding the contract method 0x5d7878a8.
//
// Solidity: function liquidateReward(uint256 _seqId, uint256 _amount, address _recipient, uint32 _l2gas) payable returns()
func (_Lockinginfo *LockinginfoTransactorSession) LiquidateReward(_seqId *big.Int, _amount *big.Int, _recipient common.Address, _l2gas uint32) (*types.Transaction, error) {
	return _Lockinginfo.Contract.LiquidateReward(&_Lockinginfo.TransactOpts, _seqId, _amount, _recipient, _l2gas)
}

// LogSignerChange is a paid mutator transaction binding the contract method 0xb3285702.
//
// Solidity: function logSignerChange(uint256 sequencerId, address oldSigner, address newSigner, uint256 nonce, bytes signerPubkey) returns()
func (_Lockinginfo *LockinginfoTransactor) LogSignerChange(opts *bind.TransactOpts, sequencerId *big.Int, oldSigner common.Address, newSigner common.Address, nonce *big.Int, signerPubkey []byte) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "logSignerChange", sequencerId, oldSigner, newSigner, nonce, signerPubkey)
}

// LogSignerChange is a paid mutator transaction binding the contract method 0xb3285702.
//
// Solidity: function logSignerChange(uint256 sequencerId, address oldSigner, address newSigner, uint256 nonce, bytes signerPubkey) returns()
func (_Lockinginfo *LockinginfoSession) LogSignerChange(sequencerId *big.Int, oldSigner common.Address, newSigner common.Address, nonce *big.Int, signerPubkey []byte) (*types.Transaction, error) {
	return _Lockinginfo.Contract.LogSignerChange(&_Lockinginfo.TransactOpts, sequencerId, oldSigner, newSigner, nonce, signerPubkey)
}

// LogSignerChange is a paid mutator transaction binding the contract method 0xb3285702.
//
// Solidity: function logSignerChange(uint256 sequencerId, address oldSigner, address newSigner, uint256 nonce, bytes signerPubkey) returns()
func (_Lockinginfo *LockinginfoTransactorSession) LogSignerChange(sequencerId *big.Int, oldSigner common.Address, newSigner common.Address, nonce *big.Int, signerPubkey []byte) (*types.Transaction, error) {
	return _Lockinginfo.Contract.LogSignerChange(&_Lockinginfo.TransactOpts, sequencerId, oldSigner, newSigner, nonce, signerPubkey)
}

// NewSequencer is a paid mutator transaction binding the contract method 0x1badded5.
//
// Solidity: function newSequencer(uint256 _id, address _owner, address _signer, uint256 _amount, uint256 _batchId, bytes _signerPubkey) returns()
func (_Lockinginfo *LockinginfoTransactor) NewSequencer(opts *bind.TransactOpts, _id *big.Int, _owner common.Address, _signer common.Address, _amount *big.Int, _batchId *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "newSequencer", _id, _owner, _signer, _amount, _batchId, _signerPubkey)
}

// NewSequencer is a paid mutator transaction binding the contract method 0x1badded5.
//
// Solidity: function newSequencer(uint256 _id, address _owner, address _signer, uint256 _amount, uint256 _batchId, bytes _signerPubkey) returns()
func (_Lockinginfo *LockinginfoSession) NewSequencer(_id *big.Int, _owner common.Address, _signer common.Address, _amount *big.Int, _batchId *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockinginfo.Contract.NewSequencer(&_Lockinginfo.TransactOpts, _id, _owner, _signer, _amount, _batchId, _signerPubkey)
}

// NewSequencer is a paid mutator transaction binding the contract method 0x1badded5.
//
// Solidity: function newSequencer(uint256 _id, address _owner, address _signer, uint256 _amount, uint256 _batchId, bytes _signerPubkey) returns()
func (_Lockinginfo *LockinginfoTransactorSession) NewSequencer(_id *big.Int, _owner common.Address, _signer common.Address, _amount *big.Int, _batchId *big.Int, _signerPubkey []byte) (*types.Transaction, error) {
	return _Lockinginfo.Contract.NewSequencer(&_Lockinginfo.TransactOpts, _id, _owner, _signer, _amount, _batchId, _signerPubkey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lockinginfo *LockinginfoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lockinginfo *LockinginfoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lockinginfo.Contract.RenounceOwnership(&_Lockinginfo.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lockinginfo *LockinginfoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lockinginfo.Contract.RenounceOwnership(&_Lockinginfo.TransactOpts)
}

// SetMaxLock is a paid mutator transaction binding the contract method 0xcd15b2a5.
//
// Solidity: function setMaxLock(uint256 _maxLock) returns()
func (_Lockinginfo *LockinginfoTransactor) SetMaxLock(opts *bind.TransactOpts, _maxLock *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "setMaxLock", _maxLock)
}

// SetMaxLock is a paid mutator transaction binding the contract method 0xcd15b2a5.
//
// Solidity: function setMaxLock(uint256 _maxLock) returns()
func (_Lockinginfo *LockinginfoSession) SetMaxLock(_maxLock *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.Contract.SetMaxLock(&_Lockinginfo.TransactOpts, _maxLock)
}

// SetMaxLock is a paid mutator transaction binding the contract method 0xcd15b2a5.
//
// Solidity: function setMaxLock(uint256 _maxLock) returns()
func (_Lockinginfo *LockinginfoTransactorSession) SetMaxLock(_maxLock *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.Contract.SetMaxLock(&_Lockinginfo.TransactOpts, _maxLock)
}

// SetMinLock is a paid mutator transaction binding the contract method 0xaa15af6a.
//
// Solidity: function setMinLock(uint256 _minLock) returns()
func (_Lockinginfo *LockinginfoTransactor) SetMinLock(opts *bind.TransactOpts, _minLock *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "setMinLock", _minLock)
}

// SetMinLock is a paid mutator transaction binding the contract method 0xaa15af6a.
//
// Solidity: function setMinLock(uint256 _minLock) returns()
func (_Lockinginfo *LockinginfoSession) SetMinLock(_minLock *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.Contract.SetMinLock(&_Lockinginfo.TransactOpts, _minLock)
}

// SetMinLock is a paid mutator transaction binding the contract method 0xaa15af6a.
//
// Solidity: function setMinLock(uint256 _minLock) returns()
func (_Lockinginfo *LockinginfoTransactorSession) SetMinLock(_minLock *big.Int) (*types.Transaction, error) {
	return _Lockinginfo.Contract.SetMinLock(&_Lockinginfo.TransactOpts, _minLock)
}

// SetRewardPayer is a paid mutator transaction binding the contract method 0xe8b8b413.
//
// Solidity: function setRewardPayer(address _payer) returns()
func (_Lockinginfo *LockinginfoTransactor) SetRewardPayer(opts *bind.TransactOpts, _payer common.Address) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "setRewardPayer", _payer)
}

// SetRewardPayer is a paid mutator transaction binding the contract method 0xe8b8b413.
//
// Solidity: function setRewardPayer(address _payer) returns()
func (_Lockinginfo *LockinginfoSession) SetRewardPayer(_payer common.Address) (*types.Transaction, error) {
	return _Lockinginfo.Contract.SetRewardPayer(&_Lockinginfo.TransactOpts, _payer)
}

// SetRewardPayer is a paid mutator transaction binding the contract method 0xe8b8b413.
//
// Solidity: function setRewardPayer(address _payer) returns()
func (_Lockinginfo *LockinginfoTransactorSession) SetRewardPayer(_payer common.Address) (*types.Transaction, error) {
	return _Lockinginfo.Contract.SetRewardPayer(&_Lockinginfo.TransactOpts, _payer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lockinginfo *LockinginfoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lockinginfo *LockinginfoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lockinginfo.Contract.TransferOwnership(&_Lockinginfo.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lockinginfo *LockinginfoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lockinginfo.Contract.TransferOwnership(&_Lockinginfo.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Lockinginfo *LockinginfoTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Lockinginfo *LockinginfoSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Lockinginfo.Contract.UpgradeTo(&_Lockinginfo.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Lockinginfo *LockinginfoTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Lockinginfo.Contract.UpgradeTo(&_Lockinginfo.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lockinginfo *LockinginfoTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lockinginfo.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lockinginfo *LockinginfoSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lockinginfo.Contract.UpgradeToAndCall(&_Lockinginfo.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lockinginfo *LockinginfoTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lockinginfo.Contract.UpgradeToAndCall(&_Lockinginfo.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lockinginfo *LockinginfoTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Lockinginfo.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lockinginfo *LockinginfoSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Lockinginfo.Contract.Fallback(&_Lockinginfo.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lockinginfo *LockinginfoTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Lockinginfo.Contract.Fallback(&_Lockinginfo.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lockinginfo *LockinginfoTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockinginfo.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lockinginfo *LockinginfoSession) Receive() (*types.Transaction, error) {
	return _Lockinginfo.Contract.Receive(&_Lockinginfo.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lockinginfo *LockinginfoTransactorSession) Receive() (*types.Transaction, error) {
	return _Lockinginfo.Contract.Receive(&_Lockinginfo.TransactOpts)
}

// LockinginfoAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Lockinginfo contract.
type LockinginfoAdminChangedIterator struct {
	Event *LockinginfoAdminChanged // Event containing the contract specifics and raw log

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
func (it *LockinginfoAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoAdminChanged)
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
		it.Event = new(LockinginfoAdminChanged)
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
func (it *LockinginfoAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoAdminChanged represents a AdminChanged event raised by the Lockinginfo contract.
type LockinginfoAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Lockinginfo *LockinginfoFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LockinginfoAdminChangedIterator, error) {

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LockinginfoAdminChangedIterator{contract: _Lockinginfo.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Lockinginfo *LockinginfoFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LockinginfoAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoAdminChanged)
				if err := _Lockinginfo.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Lockinginfo *LockinginfoFilterer) ParseAdminChanged(log types.Log) (*LockinginfoAdminChanged, error) {
	event := new(LockinginfoAdminChanged)
	if err := _Lockinginfo.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoBatchSubmitRewardIterator is returned from FilterBatchSubmitReward and is used to iterate over the raw logs and unpacked data for BatchSubmitReward events raised by the Lockinginfo contract.
type LockinginfoBatchSubmitRewardIterator struct {
	Event *LockinginfoBatchSubmitReward // Event containing the contract specifics and raw log

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
func (it *LockinginfoBatchSubmitRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoBatchSubmitReward)
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
		it.Event = new(LockinginfoBatchSubmitReward)
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
func (it *LockinginfoBatchSubmitRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoBatchSubmitRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoBatchSubmitReward represents a BatchSubmitReward event raised by the Lockinginfo contract.
type LockinginfoBatchSubmitReward struct {
	NewBatchId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBatchSubmitReward is a free log retrieval operation binding the contract event 0x9e5aedd489785d05ba086064386f2e75b3e497d3dc00a54ed1c78bfc50a3953f.
//
// Solidity: event BatchSubmitReward(uint256 _newBatchId)
func (_Lockinginfo *LockinginfoFilterer) FilterBatchSubmitReward(opts *bind.FilterOpts) (*LockinginfoBatchSubmitRewardIterator, error) {

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "BatchSubmitReward")
	if err != nil {
		return nil, err
	}
	return &LockinginfoBatchSubmitRewardIterator{contract: _Lockinginfo.contract, event: "BatchSubmitReward", logs: logs, sub: sub}, nil
}

// WatchBatchSubmitReward is a free log subscription operation binding the contract event 0x9e5aedd489785d05ba086064386f2e75b3e497d3dc00a54ed1c78bfc50a3953f.
//
// Solidity: event BatchSubmitReward(uint256 _newBatchId)
func (_Lockinginfo *LockinginfoFilterer) WatchBatchSubmitReward(opts *bind.WatchOpts, sink chan<- *LockinginfoBatchSubmitReward) (event.Subscription, error) {

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "BatchSubmitReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoBatchSubmitReward)
				if err := _Lockinginfo.contract.UnpackLog(event, "BatchSubmitReward", log); err != nil {
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

// ParseBatchSubmitReward is a log parse operation binding the contract event 0x9e5aedd489785d05ba086064386f2e75b3e497d3dc00a54ed1c78bfc50a3953f.
//
// Solidity: event BatchSubmitReward(uint256 _newBatchId)
func (_Lockinginfo *LockinginfoFilterer) ParseBatchSubmitReward(log types.Log) (*LockinginfoBatchSubmitReward, error) {
	event := new(LockinginfoBatchSubmitReward)
	if err := _Lockinginfo.contract.UnpackLog(event, "BatchSubmitReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Lockinginfo contract.
type LockinginfoBeaconUpgradedIterator struct {
	Event *LockinginfoBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *LockinginfoBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoBeaconUpgraded)
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
		it.Event = new(LockinginfoBeaconUpgraded)
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
func (it *LockinginfoBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoBeaconUpgraded represents a BeaconUpgraded event raised by the Lockinginfo contract.
type LockinginfoBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Lockinginfo *LockinginfoFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LockinginfoBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LockinginfoBeaconUpgradedIterator{contract: _Lockinginfo.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Lockinginfo *LockinginfoFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LockinginfoBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoBeaconUpgraded)
				if err := _Lockinginfo.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Lockinginfo *LockinginfoFilterer) ParseBeaconUpgraded(log types.Log) (*LockinginfoBeaconUpgraded, error) {
	event := new(LockinginfoBeaconUpgraded)
	if err := _Lockinginfo.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoClaimRewardsIterator is returned from FilterClaimRewards and is used to iterate over the raw logs and unpacked data for ClaimRewards events raised by the Lockinginfo contract.
type LockinginfoClaimRewardsIterator struct {
	Event *LockinginfoClaimRewards // Event containing the contract specifics and raw log

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
func (it *LockinginfoClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoClaimRewards)
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
		it.Event = new(LockinginfoClaimRewards)
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
func (it *LockinginfoClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoClaimRewards represents a ClaimRewards event raised by the Lockinginfo contract.
type LockinginfoClaimRewards struct {
	SequencerId *big.Int
	Recipient   common.Address
	Amount      *big.Int
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimRewards is a free log retrieval operation binding the contract event 0x18c7dc2a1800c409227dc12c0c05ada9c995ebfe0e42ae6d65f1b3ae3e6111de.
//
// Solidity: event ClaimRewards(uint256 indexed sequencerId, address recipient, uint256 indexed amount, uint256 indexed totalAmount)
func (_Lockinginfo *LockinginfoFilterer) FilterClaimRewards(opts *bind.FilterOpts, sequencerId []*big.Int, amount []*big.Int, totalAmount []*big.Int) (*LockinginfoClaimRewardsIterator, error) {

	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var totalAmountRule []interface{}
	for _, totalAmountItem := range totalAmount {
		totalAmountRule = append(totalAmountRule, totalAmountItem)
	}

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "ClaimRewards", sequencerIdRule, amountRule, totalAmountRule)
	if err != nil {
		return nil, err
	}
	return &LockinginfoClaimRewardsIterator{contract: _Lockinginfo.contract, event: "ClaimRewards", logs: logs, sub: sub}, nil
}

// WatchClaimRewards is a free log subscription operation binding the contract event 0x18c7dc2a1800c409227dc12c0c05ada9c995ebfe0e42ae6d65f1b3ae3e6111de.
//
// Solidity: event ClaimRewards(uint256 indexed sequencerId, address recipient, uint256 indexed amount, uint256 indexed totalAmount)
func (_Lockinginfo *LockinginfoFilterer) WatchClaimRewards(opts *bind.WatchOpts, sink chan<- *LockinginfoClaimRewards, sequencerId []*big.Int, amount []*big.Int, totalAmount []*big.Int) (event.Subscription, error) {

	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var totalAmountRule []interface{}
	for _, totalAmountItem := range totalAmount {
		totalAmountRule = append(totalAmountRule, totalAmountItem)
	}

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "ClaimRewards", sequencerIdRule, amountRule, totalAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoClaimRewards)
				if err := _Lockinginfo.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
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

// ParseClaimRewards is a log parse operation binding the contract event 0x18c7dc2a1800c409227dc12c0c05ada9c995ebfe0e42ae6d65f1b3ae3e6111de.
//
// Solidity: event ClaimRewards(uint256 indexed sequencerId, address recipient, uint256 indexed amount, uint256 indexed totalAmount)
func (_Lockinginfo *LockinginfoFilterer) ParseClaimRewards(log types.Log) (*LockinginfoClaimRewards, error) {
	event := new(LockinginfoClaimRewards)
	if err := _Lockinginfo.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Lockinginfo contract.
type LockinginfoInitializedIterator struct {
	Event *LockinginfoInitialized // Event containing the contract specifics and raw log

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
func (it *LockinginfoInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoInitialized)
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
		it.Event = new(LockinginfoInitialized)
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
func (it *LockinginfoInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoInitialized represents a Initialized event raised by the Lockinginfo contract.
type LockinginfoInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lockinginfo *LockinginfoFilterer) FilterInitialized(opts *bind.FilterOpts) (*LockinginfoInitializedIterator, error) {

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LockinginfoInitializedIterator{contract: _Lockinginfo.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lockinginfo *LockinginfoFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LockinginfoInitialized) (event.Subscription, error) {

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoInitialized)
				if err := _Lockinginfo.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Lockinginfo *LockinginfoFilterer) ParseInitialized(log types.Log) (*LockinginfoInitialized, error) {
	event := new(LockinginfoInitialized)
	if err := _Lockinginfo.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoLockUpdateIterator is returned from FilterLockUpdate and is used to iterate over the raw logs and unpacked data for LockUpdate events raised by the Lockinginfo contract.
type LockinginfoLockUpdateIterator struct {
	Event *LockinginfoLockUpdate // Event containing the contract specifics and raw log

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
func (it *LockinginfoLockUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoLockUpdate)
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
		it.Event = new(LockinginfoLockUpdate)
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
func (it *LockinginfoLockUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoLockUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoLockUpdate represents a LockUpdate event raised by the Lockinginfo contract.
type LockinginfoLockUpdate struct {
	SequencerId *big.Int
	Nonce       *big.Int
	NewAmount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLockUpdate is a free log retrieval operation binding the contract event 0xd716c027b3dd610e4534df756848128bbb159a757724c17d89fcc4d0151b1f30.
//
// Solidity: event LockUpdate(uint256 indexed sequencerId, uint256 indexed nonce, uint256 indexed newAmount)
func (_Lockinginfo *LockinginfoFilterer) FilterLockUpdate(opts *bind.FilterOpts, sequencerId []*big.Int, nonce []*big.Int, newAmount []*big.Int) (*LockinginfoLockUpdateIterator, error) {

	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "LockUpdate", sequencerIdRule, nonceRule, newAmountRule)
	if err != nil {
		return nil, err
	}
	return &LockinginfoLockUpdateIterator{contract: _Lockinginfo.contract, event: "LockUpdate", logs: logs, sub: sub}, nil
}

// WatchLockUpdate is a free log subscription operation binding the contract event 0xd716c027b3dd610e4534df756848128bbb159a757724c17d89fcc4d0151b1f30.
//
// Solidity: event LockUpdate(uint256 indexed sequencerId, uint256 indexed nonce, uint256 indexed newAmount)
func (_Lockinginfo *LockinginfoFilterer) WatchLockUpdate(opts *bind.WatchOpts, sink chan<- *LockinginfoLockUpdate, sequencerId []*big.Int, nonce []*big.Int, newAmount []*big.Int) (event.Subscription, error) {

	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "LockUpdate", sequencerIdRule, nonceRule, newAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoLockUpdate)
				if err := _Lockinginfo.contract.UnpackLog(event, "LockUpdate", log); err != nil {
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

// ParseLockUpdate is a log parse operation binding the contract event 0xd716c027b3dd610e4534df756848128bbb159a757724c17d89fcc4d0151b1f30.
//
// Solidity: event LockUpdate(uint256 indexed sequencerId, uint256 indexed nonce, uint256 indexed newAmount)
func (_Lockinginfo *LockinginfoFilterer) ParseLockUpdate(log types.Log) (*LockinginfoLockUpdate, error) {
	event := new(LockinginfoLockUpdate)
	if err := _Lockinginfo.contract.UnpackLog(event, "LockUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the Lockinginfo contract.
type LockinginfoLockedIterator struct {
	Event *LockinginfoLocked // Event containing the contract specifics and raw log

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
func (it *LockinginfoLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoLocked)
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
		it.Event = new(LockinginfoLocked)
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
func (it *LockinginfoLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoLocked represents a Locked event raised by the Lockinginfo contract.
type LockinginfoLocked struct {
	Signer          common.Address
	SequencerId     *big.Int
	Nonce           *big.Int
	ActivationBatch *big.Int
	Amount          *big.Int
	Total           *big.Int
	SignerPubkey    []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0xe6f1eb1f1d0ca344d03cf47b9e6ece8a7d3b196e38dd7dd2307cca75e26682a8.
//
// Solidity: event Locked(address indexed signer, uint256 indexed sequencerId, uint256 nonce, uint256 indexed activationBatch, uint256 amount, uint256 total, bytes signerPubkey)
func (_Lockinginfo *LockinginfoFilterer) FilterLocked(opts *bind.FilterOpts, signer []common.Address, sequencerId []*big.Int, activationBatch []*big.Int) (*LockinginfoLockedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	var activationBatchRule []interface{}
	for _, activationBatchItem := range activationBatch {
		activationBatchRule = append(activationBatchRule, activationBatchItem)
	}

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "Locked", signerRule, sequencerIdRule, activationBatchRule)
	if err != nil {
		return nil, err
	}
	return &LockinginfoLockedIterator{contract: _Lockinginfo.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0xe6f1eb1f1d0ca344d03cf47b9e6ece8a7d3b196e38dd7dd2307cca75e26682a8.
//
// Solidity: event Locked(address indexed signer, uint256 indexed sequencerId, uint256 nonce, uint256 indexed activationBatch, uint256 amount, uint256 total, bytes signerPubkey)
func (_Lockinginfo *LockinginfoFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *LockinginfoLocked, signer []common.Address, sequencerId []*big.Int, activationBatch []*big.Int) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	var activationBatchRule []interface{}
	for _, activationBatchItem := range activationBatch {
		activationBatchRule = append(activationBatchRule, activationBatchItem)
	}

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "Locked", signerRule, sequencerIdRule, activationBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoLocked)
				if err := _Lockinginfo.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0xe6f1eb1f1d0ca344d03cf47b9e6ece8a7d3b196e38dd7dd2307cca75e26682a8.
//
// Solidity: event Locked(address indexed signer, uint256 indexed sequencerId, uint256 nonce, uint256 indexed activationBatch, uint256 amount, uint256 total, bytes signerPubkey)
func (_Lockinginfo *LockinginfoFilterer) ParseLocked(log types.Log) (*LockinginfoLocked, error) {
	event := new(LockinginfoLocked)
	if err := _Lockinginfo.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lockinginfo contract.
type LockinginfoOwnershipTransferredIterator struct {
	Event *LockinginfoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LockinginfoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoOwnershipTransferred)
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
		it.Event = new(LockinginfoOwnershipTransferred)
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
func (it *LockinginfoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoOwnershipTransferred represents a OwnershipTransferred event raised by the Lockinginfo contract.
type LockinginfoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lockinginfo *LockinginfoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LockinginfoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LockinginfoOwnershipTransferredIterator{contract: _Lockinginfo.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lockinginfo *LockinginfoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockinginfoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoOwnershipTransferred)
				if err := _Lockinginfo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Lockinginfo *LockinginfoFilterer) ParseOwnershipTransferred(log types.Log) (*LockinginfoOwnershipTransferred, error) {
	event := new(LockinginfoOwnershipTransferred)
	if err := _Lockinginfo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoRelockedIterator is returned from FilterRelocked and is used to iterate over the raw logs and unpacked data for Relocked events raised by the Lockinginfo contract.
type LockinginfoRelockedIterator struct {
	Event *LockinginfoRelocked // Event containing the contract specifics and raw log

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
func (it *LockinginfoRelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoRelocked)
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
		it.Event = new(LockinginfoRelocked)
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
func (it *LockinginfoRelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoRelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoRelocked represents a Relocked event raised by the Lockinginfo contract.
type LockinginfoRelocked struct {
	SequencerId *big.Int
	Amount      *big.Int
	Total       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelocked is a free log retrieval operation binding the contract event 0x33a87ba488658b3d1319098cd49c6d65b72a79c0f3530fec611e7afffed04395.
//
// Solidity: event Relocked(uint256 indexed sequencerId, uint256 amount, uint256 total)
func (_Lockinginfo *LockinginfoFilterer) FilterRelocked(opts *bind.FilterOpts, sequencerId []*big.Int) (*LockinginfoRelockedIterator, error) {

	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "Relocked", sequencerIdRule)
	if err != nil {
		return nil, err
	}
	return &LockinginfoRelockedIterator{contract: _Lockinginfo.contract, event: "Relocked", logs: logs, sub: sub}, nil
}

// WatchRelocked is a free log subscription operation binding the contract event 0x33a87ba488658b3d1319098cd49c6d65b72a79c0f3530fec611e7afffed04395.
//
// Solidity: event Relocked(uint256 indexed sequencerId, uint256 amount, uint256 total)
func (_Lockinginfo *LockinginfoFilterer) WatchRelocked(opts *bind.WatchOpts, sink chan<- *LockinginfoRelocked, sequencerId []*big.Int) (event.Subscription, error) {

	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "Relocked", sequencerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoRelocked)
				if err := _Lockinginfo.contract.UnpackLog(event, "Relocked", log); err != nil {
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

// ParseRelocked is a log parse operation binding the contract event 0x33a87ba488658b3d1319098cd49c6d65b72a79c0f3530fec611e7afffed04395.
//
// Solidity: event Relocked(uint256 indexed sequencerId, uint256 amount, uint256 total)
func (_Lockinginfo *LockinginfoFilterer) ParseRelocked(log types.Log) (*LockinginfoRelocked, error) {
	event := new(LockinginfoRelocked)
	if err := _Lockinginfo.contract.UnpackLog(event, "Relocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoSetMaxLockIterator is returned from FilterSetMaxLock and is used to iterate over the raw logs and unpacked data for SetMaxLock events raised by the Lockinginfo contract.
type LockinginfoSetMaxLockIterator struct {
	Event *LockinginfoSetMaxLock // Event containing the contract specifics and raw log

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
func (it *LockinginfoSetMaxLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoSetMaxLock)
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
		it.Event = new(LockinginfoSetMaxLock)
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
func (it *LockinginfoSetMaxLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoSetMaxLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoSetMaxLock represents a SetMaxLock event raised by the Lockinginfo contract.
type LockinginfoSetMaxLock struct {
	NewMaxLock *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMaxLock is a free log retrieval operation binding the contract event 0xbe23e9641c545443c3c625039b327c0eee88e9024040be7b03c5d73862d425e0.
//
// Solidity: event SetMaxLock(uint256 _newMaxLock)
func (_Lockinginfo *LockinginfoFilterer) FilterSetMaxLock(opts *bind.FilterOpts) (*LockinginfoSetMaxLockIterator, error) {

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "SetMaxLock")
	if err != nil {
		return nil, err
	}
	return &LockinginfoSetMaxLockIterator{contract: _Lockinginfo.contract, event: "SetMaxLock", logs: logs, sub: sub}, nil
}

// WatchSetMaxLock is a free log subscription operation binding the contract event 0xbe23e9641c545443c3c625039b327c0eee88e9024040be7b03c5d73862d425e0.
//
// Solidity: event SetMaxLock(uint256 _newMaxLock)
func (_Lockinginfo *LockinginfoFilterer) WatchSetMaxLock(opts *bind.WatchOpts, sink chan<- *LockinginfoSetMaxLock) (event.Subscription, error) {

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "SetMaxLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoSetMaxLock)
				if err := _Lockinginfo.contract.UnpackLog(event, "SetMaxLock", log); err != nil {
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

// ParseSetMaxLock is a log parse operation binding the contract event 0xbe23e9641c545443c3c625039b327c0eee88e9024040be7b03c5d73862d425e0.
//
// Solidity: event SetMaxLock(uint256 _newMaxLock)
func (_Lockinginfo *LockinginfoFilterer) ParseSetMaxLock(log types.Log) (*LockinginfoSetMaxLock, error) {
	event := new(LockinginfoSetMaxLock)
	if err := _Lockinginfo.contract.UnpackLog(event, "SetMaxLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoSetMinLockIterator is returned from FilterSetMinLock and is used to iterate over the raw logs and unpacked data for SetMinLock events raised by the Lockinginfo contract.
type LockinginfoSetMinLockIterator struct {
	Event *LockinginfoSetMinLock // Event containing the contract specifics and raw log

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
func (it *LockinginfoSetMinLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoSetMinLock)
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
		it.Event = new(LockinginfoSetMinLock)
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
func (it *LockinginfoSetMinLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoSetMinLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoSetMinLock represents a SetMinLock event raised by the Lockinginfo contract.
type LockinginfoSetMinLock struct {
	NewMinLock *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMinLock is a free log retrieval operation binding the contract event 0xabb05374bb45ebfef33afb21ec5aa52333708d8217fd8e5c0616efd2530b2145.
//
// Solidity: event SetMinLock(uint256 _newMinLock)
func (_Lockinginfo *LockinginfoFilterer) FilterSetMinLock(opts *bind.FilterOpts) (*LockinginfoSetMinLockIterator, error) {

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "SetMinLock")
	if err != nil {
		return nil, err
	}
	return &LockinginfoSetMinLockIterator{contract: _Lockinginfo.contract, event: "SetMinLock", logs: logs, sub: sub}, nil
}

// WatchSetMinLock is a free log subscription operation binding the contract event 0xabb05374bb45ebfef33afb21ec5aa52333708d8217fd8e5c0616efd2530b2145.
//
// Solidity: event SetMinLock(uint256 _newMinLock)
func (_Lockinginfo *LockinginfoFilterer) WatchSetMinLock(opts *bind.WatchOpts, sink chan<- *LockinginfoSetMinLock) (event.Subscription, error) {

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "SetMinLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoSetMinLock)
				if err := _Lockinginfo.contract.UnpackLog(event, "SetMinLock", log); err != nil {
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

// ParseSetMinLock is a log parse operation binding the contract event 0xabb05374bb45ebfef33afb21ec5aa52333708d8217fd8e5c0616efd2530b2145.
//
// Solidity: event SetMinLock(uint256 _newMinLock)
func (_Lockinginfo *LockinginfoFilterer) ParseSetMinLock(log types.Log) (*LockinginfoSetMinLock, error) {
	event := new(LockinginfoSetMinLock)
	if err := _Lockinginfo.contract.UnpackLog(event, "SetMinLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoSetRewardPayerIterator is returned from FilterSetRewardPayer and is used to iterate over the raw logs and unpacked data for SetRewardPayer events raised by the Lockinginfo contract.
type LockinginfoSetRewardPayerIterator struct {
	Event *LockinginfoSetRewardPayer // Event containing the contract specifics and raw log

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
func (it *LockinginfoSetRewardPayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoSetRewardPayer)
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
		it.Event = new(LockinginfoSetRewardPayer)
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
func (it *LockinginfoSetRewardPayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoSetRewardPayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoSetRewardPayer represents a SetRewardPayer event raised by the Lockinginfo contract.
type LockinginfoSetRewardPayer struct {
	Payer common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetRewardPayer is a free log retrieval operation binding the contract event 0x30b92f5a89d7473895c4e9ce260fa7d0eefef2d59d5e68192e2e8cca4b9473a0.
//
// Solidity: event SetRewardPayer(address _payer)
func (_Lockinginfo *LockinginfoFilterer) FilterSetRewardPayer(opts *bind.FilterOpts) (*LockinginfoSetRewardPayerIterator, error) {

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "SetRewardPayer")
	if err != nil {
		return nil, err
	}
	return &LockinginfoSetRewardPayerIterator{contract: _Lockinginfo.contract, event: "SetRewardPayer", logs: logs, sub: sub}, nil
}

// WatchSetRewardPayer is a free log subscription operation binding the contract event 0x30b92f5a89d7473895c4e9ce260fa7d0eefef2d59d5e68192e2e8cca4b9473a0.
//
// Solidity: event SetRewardPayer(address _payer)
func (_Lockinginfo *LockinginfoFilterer) WatchSetRewardPayer(opts *bind.WatchOpts, sink chan<- *LockinginfoSetRewardPayer) (event.Subscription, error) {

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "SetRewardPayer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoSetRewardPayer)
				if err := _Lockinginfo.contract.UnpackLog(event, "SetRewardPayer", log); err != nil {
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

// ParseSetRewardPayer is a log parse operation binding the contract event 0x30b92f5a89d7473895c4e9ce260fa7d0eefef2d59d5e68192e2e8cca4b9473a0.
//
// Solidity: event SetRewardPayer(address _payer)
func (_Lockinginfo *LockinginfoFilterer) ParseSetRewardPayer(log types.Log) (*LockinginfoSetRewardPayer, error) {
	event := new(LockinginfoSetRewardPayer)
	if err := _Lockinginfo.contract.UnpackLog(event, "SetRewardPayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoSignerChangeIterator is returned from FilterSignerChange and is used to iterate over the raw logs and unpacked data for SignerChange events raised by the Lockinginfo contract.
type LockinginfoSignerChangeIterator struct {
	Event *LockinginfoSignerChange // Event containing the contract specifics and raw log

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
func (it *LockinginfoSignerChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoSignerChange)
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
		it.Event = new(LockinginfoSignerChange)
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
func (it *LockinginfoSignerChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoSignerChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoSignerChange represents a SignerChange event raised by the Lockinginfo contract.
type LockinginfoSignerChange struct {
	SequencerId  *big.Int
	Nonce        *big.Int
	OldSigner    common.Address
	NewSigner    common.Address
	SignerPubkey []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSignerChange is a free log retrieval operation binding the contract event 0x086044c0612a8c965d4cccd907f0d588e40ad68438bd4c1274cac60f4c3a9d1f.
//
// Solidity: event SignerChange(uint256 indexed sequencerId, uint256 nonce, address indexed oldSigner, address indexed newSigner, bytes signerPubkey)
func (_Lockinginfo *LockinginfoFilterer) FilterSignerChange(opts *bind.FilterOpts, sequencerId []*big.Int, oldSigner []common.Address, newSigner []common.Address) (*LockinginfoSignerChangeIterator, error) {

	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "SignerChange", sequencerIdRule, oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return &LockinginfoSignerChangeIterator{contract: _Lockinginfo.contract, event: "SignerChange", logs: logs, sub: sub}, nil
}

// WatchSignerChange is a free log subscription operation binding the contract event 0x086044c0612a8c965d4cccd907f0d588e40ad68438bd4c1274cac60f4c3a9d1f.
//
// Solidity: event SignerChange(uint256 indexed sequencerId, uint256 nonce, address indexed oldSigner, address indexed newSigner, bytes signerPubkey)
func (_Lockinginfo *LockinginfoFilterer) WatchSignerChange(opts *bind.WatchOpts, sink chan<- *LockinginfoSignerChange, sequencerId []*big.Int, oldSigner []common.Address, newSigner []common.Address) (event.Subscription, error) {

	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "SignerChange", sequencerIdRule, oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoSignerChange)
				if err := _Lockinginfo.contract.UnpackLog(event, "SignerChange", log); err != nil {
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

// ParseSignerChange is a log parse operation binding the contract event 0x086044c0612a8c965d4cccd907f0d588e40ad68438bd4c1274cac60f4c3a9d1f.
//
// Solidity: event SignerChange(uint256 indexed sequencerId, uint256 nonce, address indexed oldSigner, address indexed newSigner, bytes signerPubkey)
func (_Lockinginfo *LockinginfoFilterer) ParseSignerChange(log types.Log) (*LockinginfoSignerChange, error) {
	event := new(LockinginfoSignerChange)
	if err := _Lockinginfo.contract.UnpackLog(event, "SignerChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoUnlockInitIterator is returned from FilterUnlockInit and is used to iterate over the raw logs and unpacked data for UnlockInit events raised by the Lockinginfo contract.
type LockinginfoUnlockInitIterator struct {
	Event *LockinginfoUnlockInit // Event containing the contract specifics and raw log

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
func (it *LockinginfoUnlockInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoUnlockInit)
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
		it.Event = new(LockinginfoUnlockInit)
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
func (it *LockinginfoUnlockInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoUnlockInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoUnlockInit represents a UnlockInit event raised by the Lockinginfo contract.
type LockinginfoUnlockInit struct {
	User              common.Address
	SequencerId       *big.Int
	Nonce             *big.Int
	DeactivationBatch *big.Int
	DeactivationTime  *big.Int
	UnlockClaimTime   *big.Int
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUnlockInit is a free log retrieval operation binding the contract event 0x06d9e13438f0daf13a71d63f3f8579db8bdeb299e4b651942313c73224d7af69.
//
// Solidity: event UnlockInit(address indexed user, uint256 indexed sequencerId, uint256 nonce, uint256 deactivationBatch, uint256 deactivationTime, uint256 unlockClaimTime, uint256 indexed amount)
func (_Lockinginfo *LockinginfoFilterer) FilterUnlockInit(opts *bind.FilterOpts, user []common.Address, sequencerId []*big.Int, amount []*big.Int) (*LockinginfoUnlockInitIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "UnlockInit", userRule, sequencerIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LockinginfoUnlockInitIterator{contract: _Lockinginfo.contract, event: "UnlockInit", logs: logs, sub: sub}, nil
}

// WatchUnlockInit is a free log subscription operation binding the contract event 0x06d9e13438f0daf13a71d63f3f8579db8bdeb299e4b651942313c73224d7af69.
//
// Solidity: event UnlockInit(address indexed user, uint256 indexed sequencerId, uint256 nonce, uint256 deactivationBatch, uint256 deactivationTime, uint256 unlockClaimTime, uint256 indexed amount)
func (_Lockinginfo *LockinginfoFilterer) WatchUnlockInit(opts *bind.WatchOpts, sink chan<- *LockinginfoUnlockInit, user []common.Address, sequencerId []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "UnlockInit", userRule, sequencerIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoUnlockInit)
				if err := _Lockinginfo.contract.UnpackLog(event, "UnlockInit", log); err != nil {
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

// ParseUnlockInit is a log parse operation binding the contract event 0x06d9e13438f0daf13a71d63f3f8579db8bdeb299e4b651942313c73224d7af69.
//
// Solidity: event UnlockInit(address indexed user, uint256 indexed sequencerId, uint256 nonce, uint256 deactivationBatch, uint256 deactivationTime, uint256 unlockClaimTime, uint256 indexed amount)
func (_Lockinginfo *LockinginfoFilterer) ParseUnlockInit(log types.Log) (*LockinginfoUnlockInit, error) {
	event := new(LockinginfoUnlockInit)
	if err := _Lockinginfo.contract.UnpackLog(event, "UnlockInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the Lockinginfo contract.
type LockinginfoUnlockedIterator struct {
	Event *LockinginfoUnlocked // Event containing the contract specifics and raw log

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
func (it *LockinginfoUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoUnlocked)
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
		it.Event = new(LockinginfoUnlocked)
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
func (it *LockinginfoUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoUnlocked represents a Unlocked event raised by the Lockinginfo contract.
type LockinginfoUnlocked struct {
	User        common.Address
	SequencerId *big.Int
	Amount      *big.Int
	Total       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x5245d528087a96a64f4589a764f00061e4671eab90cb1e019b1a5b24b2e4c2a8.
//
// Solidity: event Unlocked(address indexed user, uint256 indexed sequencerId, uint256 amount, uint256 total)
func (_Lockinginfo *LockinginfoFilterer) FilterUnlocked(opts *bind.FilterOpts, user []common.Address, sequencerId []*big.Int) (*LockinginfoUnlockedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "Unlocked", userRule, sequencerIdRule)
	if err != nil {
		return nil, err
	}
	return &LockinginfoUnlockedIterator{contract: _Lockinginfo.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x5245d528087a96a64f4589a764f00061e4671eab90cb1e019b1a5b24b2e4c2a8.
//
// Solidity: event Unlocked(address indexed user, uint256 indexed sequencerId, uint256 amount, uint256 total)
func (_Lockinginfo *LockinginfoFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *LockinginfoUnlocked, user []common.Address, sequencerId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var sequencerIdRule []interface{}
	for _, sequencerIdItem := range sequencerId {
		sequencerIdRule = append(sequencerIdRule, sequencerIdItem)
	}

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "Unlocked", userRule, sequencerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoUnlocked)
				if err := _Lockinginfo.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0x5245d528087a96a64f4589a764f00061e4671eab90cb1e019b1a5b24b2e4c2a8.
//
// Solidity: event Unlocked(address indexed user, uint256 indexed sequencerId, uint256 amount, uint256 total)
func (_Lockinginfo *LockinginfoFilterer) ParseUnlocked(log types.Log) (*LockinginfoUnlocked, error) {
	event := new(LockinginfoUnlocked)
	if err := _Lockinginfo.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockinginfoUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Lockinginfo contract.
type LockinginfoUpgradedIterator struct {
	Event *LockinginfoUpgraded // Event containing the contract specifics and raw log

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
func (it *LockinginfoUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockinginfoUpgraded)
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
		it.Event = new(LockinginfoUpgraded)
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
func (it *LockinginfoUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockinginfoUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockinginfoUpgraded represents a Upgraded event raised by the Lockinginfo contract.
type LockinginfoUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Lockinginfo *LockinginfoFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LockinginfoUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Lockinginfo.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LockinginfoUpgradedIterator{contract: _Lockinginfo.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Lockinginfo *LockinginfoFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LockinginfoUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Lockinginfo.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockinginfoUpgraded)
				if err := _Lockinginfo.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Lockinginfo *LockinginfoFilterer) ParseUpgraded(log types.Log) (*LockinginfoUpgraded, error) {
	event := new(LockinginfoUpgraded)
	if err := _Lockinginfo.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

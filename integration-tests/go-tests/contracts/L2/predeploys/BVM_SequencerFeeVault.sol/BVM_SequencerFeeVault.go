// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BVM_SequencerFeeVault

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

// BVMSequencerFeeVaultMetaData contains all meta data concerning the BVMSequencerFeeVault contract.
var BVMSequencerFeeVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1FeeWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bvmGasPriceOracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_burner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minWithdrawalAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L1Gas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bvmGasPriceOracleAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1FeeWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minWithdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_burner\",\"type\":\"address\"}],\"name\":\"setBurner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1FeeWallet\",\"type\":\"address\"}],\"name\":\"setL1FeeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minWithdrawalAmount\",\"type\":\"uint256\"}],\"name\":\"setMinWithdrawalAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BVMSequencerFeeVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use BVMSequencerFeeVaultMetaData.ABI instead.
var BVMSequencerFeeVaultABI = BVMSequencerFeeVaultMetaData.ABI

// BVMSequencerFeeVault is an auto generated Go binding around an Ethereum contract.
type BVMSequencerFeeVault struct {
	BVMSequencerFeeVaultCaller     // Read-only binding to the contract
	BVMSequencerFeeVaultTransactor // Write-only binding to the contract
	BVMSequencerFeeVaultFilterer   // Log filterer for contract events
}

// BVMSequencerFeeVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type BVMSequencerFeeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMSequencerFeeVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BVMSequencerFeeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMSequencerFeeVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BVMSequencerFeeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMSequencerFeeVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BVMSequencerFeeVaultSession struct {
	Contract     *BVMSequencerFeeVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BVMSequencerFeeVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BVMSequencerFeeVaultCallerSession struct {
	Contract *BVMSequencerFeeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BVMSequencerFeeVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BVMSequencerFeeVaultTransactorSession struct {
	Contract     *BVMSequencerFeeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BVMSequencerFeeVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type BVMSequencerFeeVaultRaw struct {
	Contract *BVMSequencerFeeVault // Generic contract binding to access the raw methods on
}

// BVMSequencerFeeVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BVMSequencerFeeVaultCallerRaw struct {
	Contract *BVMSequencerFeeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// BVMSequencerFeeVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BVMSequencerFeeVaultTransactorRaw struct {
	Contract *BVMSequencerFeeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBVMSequencerFeeVault creates a new instance of BVMSequencerFeeVault, bound to a specific deployed contract.
func NewBVMSequencerFeeVault(address common.Address, backend bind.ContractBackend) (*BVMSequencerFeeVault, error) {
	contract, err := bindBVMSequencerFeeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BVMSequencerFeeVault{BVMSequencerFeeVaultCaller: BVMSequencerFeeVaultCaller{contract: contract}, BVMSequencerFeeVaultTransactor: BVMSequencerFeeVaultTransactor{contract: contract}, BVMSequencerFeeVaultFilterer: BVMSequencerFeeVaultFilterer{contract: contract}}, nil
}

// NewBVMSequencerFeeVaultCaller creates a new read-only instance of BVMSequencerFeeVault, bound to a specific deployed contract.
func NewBVMSequencerFeeVaultCaller(address common.Address, caller bind.ContractCaller) (*BVMSequencerFeeVaultCaller, error) {
	contract, err := bindBVMSequencerFeeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BVMSequencerFeeVaultCaller{contract: contract}, nil
}

// NewBVMSequencerFeeVaultTransactor creates a new write-only instance of BVMSequencerFeeVault, bound to a specific deployed contract.
func NewBVMSequencerFeeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*BVMSequencerFeeVaultTransactor, error) {
	contract, err := bindBVMSequencerFeeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BVMSequencerFeeVaultTransactor{contract: contract}, nil
}

// NewBVMSequencerFeeVaultFilterer creates a new log filterer instance of BVMSequencerFeeVault, bound to a specific deployed contract.
func NewBVMSequencerFeeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*BVMSequencerFeeVaultFilterer, error) {
	contract, err := bindBVMSequencerFeeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BVMSequencerFeeVaultFilterer{contract: contract}, nil
}

// bindBVMSequencerFeeVault binds a generic wrapper to an already deployed contract.
func bindBVMSequencerFeeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BVMSequencerFeeVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMSequencerFeeVault.Contract.BVMSequencerFeeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.BVMSequencerFeeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.BVMSequencerFeeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMSequencerFeeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.contract.Transact(opts, method, params...)
}

// L1Gas is a free data retrieval call binding the contract method 0x5558979e.
//
// Solidity: function L1Gas() view returns(uint256)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCaller) L1Gas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMSequencerFeeVault.contract.Call(opts, &out, "L1Gas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1Gas is a free data retrieval call binding the contract method 0x5558979e.
//
// Solidity: function L1Gas() view returns(uint256)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) L1Gas() (*big.Int, error) {
	return _BVMSequencerFeeVault.Contract.L1Gas(&_BVMSequencerFeeVault.CallOpts)
}

// L1Gas is a free data retrieval call binding the contract method 0x5558979e.
//
// Solidity: function L1Gas() view returns(uint256)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCallerSession) L1Gas() (*big.Int, error) {
	return _BVMSequencerFeeVault.Contract.L1Gas(&_BVMSequencerFeeVault.CallOpts)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCaller) Burner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMSequencerFeeVault.contract.Call(opts, &out, "burner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) Burner() (common.Address, error) {
	return _BVMSequencerFeeVault.Contract.Burner(&_BVMSequencerFeeVault.CallOpts)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCallerSession) Burner() (common.Address, error) {
	return _BVMSequencerFeeVault.Contract.Burner(&_BVMSequencerFeeVault.CallOpts)
}

// BvmGasPriceOracleAddress is a free data retrieval call binding the contract method 0x110b7eb0.
//
// Solidity: function bvmGasPriceOracleAddress() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCaller) BvmGasPriceOracleAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMSequencerFeeVault.contract.Call(opts, &out, "bvmGasPriceOracleAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BvmGasPriceOracleAddress is a free data retrieval call binding the contract method 0x110b7eb0.
//
// Solidity: function bvmGasPriceOracleAddress() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) BvmGasPriceOracleAddress() (common.Address, error) {
	return _BVMSequencerFeeVault.Contract.BvmGasPriceOracleAddress(&_BVMSequencerFeeVault.CallOpts)
}

// BvmGasPriceOracleAddress is a free data retrieval call binding the contract method 0x110b7eb0.
//
// Solidity: function bvmGasPriceOracleAddress() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCallerSession) BvmGasPriceOracleAddress() (common.Address, error) {
	return _BVMSequencerFeeVault.Contract.BvmGasPriceOracleAddress(&_BVMSequencerFeeVault.CallOpts)
}

// L1FeeWallet is a free data retrieval call binding the contract method 0xd4ff9218.
//
// Solidity: function l1FeeWallet() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCaller) L1FeeWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMSequencerFeeVault.contract.Call(opts, &out, "l1FeeWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1FeeWallet is a free data retrieval call binding the contract method 0xd4ff9218.
//
// Solidity: function l1FeeWallet() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) L1FeeWallet() (common.Address, error) {
	return _BVMSequencerFeeVault.Contract.L1FeeWallet(&_BVMSequencerFeeVault.CallOpts)
}

// L1FeeWallet is a free data retrieval call binding the contract method 0xd4ff9218.
//
// Solidity: function l1FeeWallet() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCallerSession) L1FeeWallet() (common.Address, error) {
	return _BVMSequencerFeeVault.Contract.L1FeeWallet(&_BVMSequencerFeeVault.CallOpts)
}

// MinWithdrawalAmount is a free data retrieval call binding the contract method 0x8312f149.
//
// Solidity: function minWithdrawalAmount() view returns(uint256)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCaller) MinWithdrawalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMSequencerFeeVault.contract.Call(opts, &out, "minWithdrawalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawalAmount is a free data retrieval call binding the contract method 0x8312f149.
//
// Solidity: function minWithdrawalAmount() view returns(uint256)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) MinWithdrawalAmount() (*big.Int, error) {
	return _BVMSequencerFeeVault.Contract.MinWithdrawalAmount(&_BVMSequencerFeeVault.CallOpts)
}

// MinWithdrawalAmount is a free data retrieval call binding the contract method 0x8312f149.
//
// Solidity: function minWithdrawalAmount() view returns(uint256)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCallerSession) MinWithdrawalAmount() (*big.Int, error) {
	return _BVMSequencerFeeVault.Contract.MinWithdrawalAmount(&_BVMSequencerFeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMSequencerFeeVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) Owner() (common.Address, error) {
	return _BVMSequencerFeeVault.Contract.Owner(&_BVMSequencerFeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCallerSession) Owner() (common.Address, error) {
	return _BVMSequencerFeeVault.Contract.Owner(&_BVMSequencerFeeVault.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.RenounceOwnership(&_BVMSequencerFeeVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.RenounceOwnership(&_BVMSequencerFeeVault.TransactOpts)
}

// SetBurner is a paid mutator transaction binding the contract method 0xa996d6ce.
//
// Solidity: function setBurner(address _burner) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactor) SetBurner(opts *bind.TransactOpts, _burner common.Address) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.contract.Transact(opts, "setBurner", _burner)
}

// SetBurner is a paid mutator transaction binding the contract method 0xa996d6ce.
//
// Solidity: function setBurner(address _burner) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) SetBurner(_burner common.Address) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.SetBurner(&_BVMSequencerFeeVault.TransactOpts, _burner)
}

// SetBurner is a paid mutator transaction binding the contract method 0xa996d6ce.
//
// Solidity: function setBurner(address _burner) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactorSession) SetBurner(_burner common.Address) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.SetBurner(&_BVMSequencerFeeVault.TransactOpts, _burner)
}

// SetL1FeeWallet is a paid mutator transaction binding the contract method 0xbfb08462.
//
// Solidity: function setL1FeeWallet(address _l1FeeWallet) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactor) SetL1FeeWallet(opts *bind.TransactOpts, _l1FeeWallet common.Address) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.contract.Transact(opts, "setL1FeeWallet", _l1FeeWallet)
}

// SetL1FeeWallet is a paid mutator transaction binding the contract method 0xbfb08462.
//
// Solidity: function setL1FeeWallet(address _l1FeeWallet) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) SetL1FeeWallet(_l1FeeWallet common.Address) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.SetL1FeeWallet(&_BVMSequencerFeeVault.TransactOpts, _l1FeeWallet)
}

// SetL1FeeWallet is a paid mutator transaction binding the contract method 0xbfb08462.
//
// Solidity: function setL1FeeWallet(address _l1FeeWallet) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactorSession) SetL1FeeWallet(_l1FeeWallet common.Address) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.SetL1FeeWallet(&_BVMSequencerFeeVault.TransactOpts, _l1FeeWallet)
}

// SetMinWithdrawalAmount is a paid mutator transaction binding the contract method 0x85b5b14d.
//
// Solidity: function setMinWithdrawalAmount(uint256 _minWithdrawalAmount) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactor) SetMinWithdrawalAmount(opts *bind.TransactOpts, _minWithdrawalAmount *big.Int) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.contract.Transact(opts, "setMinWithdrawalAmount", _minWithdrawalAmount)
}

// SetMinWithdrawalAmount is a paid mutator transaction binding the contract method 0x85b5b14d.
//
// Solidity: function setMinWithdrawalAmount(uint256 _minWithdrawalAmount) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) SetMinWithdrawalAmount(_minWithdrawalAmount *big.Int) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.SetMinWithdrawalAmount(&_BVMSequencerFeeVault.TransactOpts, _minWithdrawalAmount)
}

// SetMinWithdrawalAmount is a paid mutator transaction binding the contract method 0x85b5b14d.
//
// Solidity: function setMinWithdrawalAmount(uint256 _minWithdrawalAmount) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactorSession) SetMinWithdrawalAmount(_minWithdrawalAmount *big.Int) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.SetMinWithdrawalAmount(&_BVMSequencerFeeVault.TransactOpts, _minWithdrawalAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.TransferOwnership(&_BVMSequencerFeeVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.TransferOwnership(&_BVMSequencerFeeVault.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) Withdraw() (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.Withdraw(&_BVMSequencerFeeVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactorSession) Withdraw() (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.Withdraw(&_BVMSequencerFeeVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMSequencerFeeVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) Receive() (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.Receive(&_BVMSequencerFeeVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _BVMSequencerFeeVault.Contract.Receive(&_BVMSequencerFeeVault.TransactOpts)
}

// BVMSequencerFeeVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BVMSequencerFeeVault contract.
type BVMSequencerFeeVaultOwnershipTransferredIterator struct {
	Event *BVMSequencerFeeVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BVMSequencerFeeVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMSequencerFeeVaultOwnershipTransferred)
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
		it.Event = new(BVMSequencerFeeVaultOwnershipTransferred)
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
func (it *BVMSequencerFeeVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMSequencerFeeVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMSequencerFeeVaultOwnershipTransferred represents a OwnershipTransferred event raised by the BVMSequencerFeeVault contract.
type BVMSequencerFeeVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BVMSequencerFeeVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BVMSequencerFeeVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BVMSequencerFeeVaultOwnershipTransferredIterator{contract: _BVMSequencerFeeVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BVMSequencerFeeVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BVMSequencerFeeVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMSequencerFeeVaultOwnershipTransferred)
				if err := _BVMSequencerFeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultFilterer) ParseOwnershipTransferred(log types.Log) (*BVMSequencerFeeVaultOwnershipTransferred, error) {
	event := new(BVMSequencerFeeVaultOwnershipTransferred)
	if err := _BVMSequencerFeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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
)

// BVMSequencerFeeVaultMetaData contains all meta data concerning the BVMSequencerFeeVault contract.
var BVMSequencerFeeVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1FeeWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bvmGasPriceOracleAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"L1Gas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bvmGasPriceOracleAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1FeeWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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
	parsed, err := abi.JSON(strings.NewReader(BVMSequencerFeeVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCaller) MINWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMSequencerFeeVault.contract.Call(opts, &out, "MIN_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _BVMSequencerFeeVault.Contract.MINWITHDRAWALAMOUNT(&_BVMSequencerFeeVault.CallOpts)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_BVMSequencerFeeVault *BVMSequencerFeeVaultCallerSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _BVMSequencerFeeVault.Contract.MINWITHDRAWALAMOUNT(&_BVMSequencerFeeVault.CallOpts)
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

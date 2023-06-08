// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Arrays

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

// ArraysMetaData contains all meta data concerning the Arrays contract.
var ArraysMetaData = &bind.MetaData{
	ABI: "[]",
}

// ArraysABI is the input ABI used to generate the binding from.
// Deprecated: Use ArraysMetaData.ABI instead.
var ArraysABI = ArraysMetaData.ABI

// Arrays is an auto generated Go binding around an Ethereum contract.
type Arrays struct {
	ArraysCaller     // Read-only binding to the contract
	ArraysTransactor // Write-only binding to the contract
	ArraysFilterer   // Log filterer for contract events
}

// ArraysCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArraysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArraysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArraysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArraysSession struct {
	Contract     *Arrays           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArraysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArraysCallerSession struct {
	Contract *ArraysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArraysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArraysTransactorSession struct {
	Contract     *ArraysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArraysRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArraysRaw struct {
	Contract *Arrays // Generic contract binding to access the raw methods on
}

// ArraysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArraysCallerRaw struct {
	Contract *ArraysCaller // Generic read-only contract binding to access the raw methods on
}

// ArraysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArraysTransactorRaw struct {
	Contract *ArraysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArrays creates a new instance of Arrays, bound to a specific deployed contract.
func NewArrays(address common.Address, backend bind.ContractBackend) (*Arrays, error) {
	contract, err := bindArrays(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arrays{ArraysCaller: ArraysCaller{contract: contract}, ArraysTransactor: ArraysTransactor{contract: contract}, ArraysFilterer: ArraysFilterer{contract: contract}}, nil
}

// NewArraysCaller creates a new read-only instance of Arrays, bound to a specific deployed contract.
func NewArraysCaller(address common.Address, caller bind.ContractCaller) (*ArraysCaller, error) {
	contract, err := bindArrays(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArraysCaller{contract: contract}, nil
}

// NewArraysTransactor creates a new write-only instance of Arrays, bound to a specific deployed contract.
func NewArraysTransactor(address common.Address, transactor bind.ContractTransactor) (*ArraysTransactor, error) {
	contract, err := bindArrays(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArraysTransactor{contract: contract}, nil
}

// NewArraysFilterer creates a new log filterer instance of Arrays, bound to a specific deployed contract.
func NewArraysFilterer(address common.Address, filterer bind.ContractFilterer) (*ArraysFilterer, error) {
	contract, err := bindArrays(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArraysFilterer{contract: contract}, nil
}

// bindArrays binds a generic wrapper to an already deployed contract.
func bindArrays(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArraysMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arrays *ArraysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arrays.Contract.ArraysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arrays *ArraysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arrays.Contract.ArraysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arrays *ArraysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arrays.Contract.ArraysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arrays *ArraysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arrays.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arrays *ArraysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arrays.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arrays *ArraysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arrays.Contract.contract.Transact(opts, method, params...)
}

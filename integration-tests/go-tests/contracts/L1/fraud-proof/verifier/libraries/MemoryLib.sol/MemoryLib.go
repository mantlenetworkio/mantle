// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MemoryLib

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

// MemoryLibMetaData contains all meta data concerning the MemoryLib contract.
var MemoryLibMetaData = &bind.MetaData{
	ABI: "[]",
}

// MemoryLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MemoryLibMetaData.ABI instead.
var MemoryLibABI = MemoryLibMetaData.ABI

// MemoryLib is an auto generated Go binding around an Ethereum contract.
type MemoryLib struct {
	MemoryLibCaller     // Read-only binding to the contract
	MemoryLibTransactor // Write-only binding to the contract
	MemoryLibFilterer   // Log filterer for contract events
}

// MemoryLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemoryLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemoryLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemoryLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemoryLibSession struct {
	Contract     *MemoryLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemoryLibCallerSession struct {
	Contract *MemoryLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MemoryLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemoryLibTransactorSession struct {
	Contract     *MemoryLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MemoryLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemoryLibRaw struct {
	Contract *MemoryLib // Generic contract binding to access the raw methods on
}

// MemoryLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemoryLibCallerRaw struct {
	Contract *MemoryLibCaller // Generic read-only contract binding to access the raw methods on
}

// MemoryLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemoryLibTransactorRaw struct {
	Contract *MemoryLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemoryLib creates a new instance of MemoryLib, bound to a specific deployed contract.
func NewMemoryLib(address common.Address, backend bind.ContractBackend) (*MemoryLib, error) {
	contract, err := bindMemoryLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MemoryLib{MemoryLibCaller: MemoryLibCaller{contract: contract}, MemoryLibTransactor: MemoryLibTransactor{contract: contract}, MemoryLibFilterer: MemoryLibFilterer{contract: contract}}, nil
}

// NewMemoryLibCaller creates a new read-only instance of MemoryLib, bound to a specific deployed contract.
func NewMemoryLibCaller(address common.Address, caller bind.ContractCaller) (*MemoryLibCaller, error) {
	contract, err := bindMemoryLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryLibCaller{contract: contract}, nil
}

// NewMemoryLibTransactor creates a new write-only instance of MemoryLib, bound to a specific deployed contract.
func NewMemoryLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MemoryLibTransactor, error) {
	contract, err := bindMemoryLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryLibTransactor{contract: contract}, nil
}

// NewMemoryLibFilterer creates a new log filterer instance of MemoryLib, bound to a specific deployed contract.
func NewMemoryLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MemoryLibFilterer, error) {
	contract, err := bindMemoryLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemoryLibFilterer{contract: contract}, nil
}

// bindMemoryLib binds a generic wrapper to an already deployed contract.
func bindMemoryLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MemoryLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemoryLib *MemoryLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemoryLib.Contract.MemoryLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemoryLib *MemoryLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemoryLib.Contract.MemoryLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemoryLib *MemoryLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemoryLib.Contract.MemoryLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemoryLib *MemoryLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemoryLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemoryLib *MemoryLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemoryLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemoryLib *MemoryLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemoryLib.Contract.contract.Transact(opts, method, params...)
}

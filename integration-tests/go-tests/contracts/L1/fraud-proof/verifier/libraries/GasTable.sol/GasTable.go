// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GasTable

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

// GasTableMetaData contains all meta data concerning the GasTable contract.
var GasTableMetaData = &bind.MetaData{
	ABI: "[]",
}

// GasTableABI is the input ABI used to generate the binding from.
// Deprecated: Use GasTableMetaData.ABI instead.
var GasTableABI = GasTableMetaData.ABI

// GasTable is an auto generated Go binding around an Ethereum contract.
type GasTable struct {
	GasTableCaller     // Read-only binding to the contract
	GasTableTransactor // Write-only binding to the contract
	GasTableFilterer   // Log filterer for contract events
}

// GasTableCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasTableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasTableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasTableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasTableSession struct {
	Contract     *GasTable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasTableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasTableCallerSession struct {
	Contract *GasTableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GasTableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasTableTransactorSession struct {
	Contract     *GasTableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GasTableRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasTableRaw struct {
	Contract *GasTable // Generic contract binding to access the raw methods on
}

// GasTableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasTableCallerRaw struct {
	Contract *GasTableCaller // Generic read-only contract binding to access the raw methods on
}

// GasTableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasTableTransactorRaw struct {
	Contract *GasTableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasTable creates a new instance of GasTable, bound to a specific deployed contract.
func NewGasTable(address common.Address, backend bind.ContractBackend) (*GasTable, error) {
	contract, err := bindGasTable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasTable{GasTableCaller: GasTableCaller{contract: contract}, GasTableTransactor: GasTableTransactor{contract: contract}, GasTableFilterer: GasTableFilterer{contract: contract}}, nil
}

// NewGasTableCaller creates a new read-only instance of GasTable, bound to a specific deployed contract.
func NewGasTableCaller(address common.Address, caller bind.ContractCaller) (*GasTableCaller, error) {
	contract, err := bindGasTable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasTableCaller{contract: contract}, nil
}

// NewGasTableTransactor creates a new write-only instance of GasTable, bound to a specific deployed contract.
func NewGasTableTransactor(address common.Address, transactor bind.ContractTransactor) (*GasTableTransactor, error) {
	contract, err := bindGasTable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasTableTransactor{contract: contract}, nil
}

// NewGasTableFilterer creates a new log filterer instance of GasTable, bound to a specific deployed contract.
func NewGasTableFilterer(address common.Address, filterer bind.ContractFilterer) (*GasTableFilterer, error) {
	contract, err := bindGasTable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasTableFilterer{contract: contract}, nil
}

// bindGasTable binds a generic wrapper to an already deployed contract.
func bindGasTable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GasTableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasTable *GasTableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasTable.Contract.GasTableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasTable *GasTableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasTable.Contract.GasTableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasTable *GasTableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasTable.Contract.GasTableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasTable *GasTableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasTable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasTable *GasTableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasTable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasTable *GasTableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasTable.Contract.contract.Transact(opts, method, params...)
}

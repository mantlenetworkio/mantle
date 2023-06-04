// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RLPWriter

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

// RLPWriterMetaData contains all meta data concerning the RLPWriter contract.
var RLPWriterMetaData = &bind.MetaData{
	ABI: "[]",
}

// RLPWriterABI is the input ABI used to generate the binding from.
// Deprecated: Use RLPWriterMetaData.ABI instead.
var RLPWriterABI = RLPWriterMetaData.ABI

// RLPWriter is an auto generated Go binding around an Ethereum contract.
type RLPWriter struct {
	RLPWriterCaller     // Read-only binding to the contract
	RLPWriterTransactor // Write-only binding to the contract
	RLPWriterFilterer   // Log filterer for contract events
}

// RLPWriterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPWriterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPWriterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPWriterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPWriterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPWriterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPWriterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPWriterSession struct {
	Contract     *RLPWriter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPWriterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPWriterCallerSession struct {
	Contract *RLPWriterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPWriterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPWriterTransactorSession struct {
	Contract     *RLPWriterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPWriterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPWriterRaw struct {
	Contract *RLPWriter // Generic contract binding to access the raw methods on
}

// RLPWriterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPWriterCallerRaw struct {
	Contract *RLPWriterCaller // Generic read-only contract binding to access the raw methods on
}

// RLPWriterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPWriterTransactorRaw struct {
	Contract *RLPWriterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPWriter creates a new instance of RLPWriter, bound to a specific deployed contract.
func NewRLPWriter(address common.Address, backend bind.ContractBackend) (*RLPWriter, error) {
	contract, err := bindRLPWriter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPWriter{RLPWriterCaller: RLPWriterCaller{contract: contract}, RLPWriterTransactor: RLPWriterTransactor{contract: contract}, RLPWriterFilterer: RLPWriterFilterer{contract: contract}}, nil
}

// NewRLPWriterCaller creates a new read-only instance of RLPWriter, bound to a specific deployed contract.
func NewRLPWriterCaller(address common.Address, caller bind.ContractCaller) (*RLPWriterCaller, error) {
	contract, err := bindRLPWriter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPWriterCaller{contract: contract}, nil
}

// NewRLPWriterTransactor creates a new write-only instance of RLPWriter, bound to a specific deployed contract.
func NewRLPWriterTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPWriterTransactor, error) {
	contract, err := bindRLPWriter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPWriterTransactor{contract: contract}, nil
}

// NewRLPWriterFilterer creates a new log filterer instance of RLPWriter, bound to a specific deployed contract.
func NewRLPWriterFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPWriterFilterer, error) {
	contract, err := bindRLPWriter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPWriterFilterer{contract: contract}, nil
}

// bindRLPWriter binds a generic wrapper to an already deployed contract.
func bindRLPWriter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RLPWriterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPWriter *RLPWriterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPWriter.Contract.RLPWriterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPWriter *RLPWriterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPWriter.Contract.RLPWriterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPWriter *RLPWriterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPWriter.Contract.RLPWriterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPWriter *RLPWriterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPWriter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPWriter *RLPWriterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPWriter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPWriter *RLPWriterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPWriter.Contract.contract.Transact(opts, method, params...)
}

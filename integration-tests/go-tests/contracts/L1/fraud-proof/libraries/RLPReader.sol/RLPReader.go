// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RLPReader

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

// RLPReaderMetaData contains all meta data concerning the RLPReader contract.
var RLPReaderMetaData = &bind.MetaData{
	ABI: "[]",
}

// RLPReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use RLPReaderMetaData.ABI instead.
var RLPReaderABI = RLPReaderMetaData.ABI

// RLPReader is an auto generated Go binding around an Ethereum contract.
type RLPReader struct {
	RLPReaderCaller     // Read-only binding to the contract
	RLPReaderTransactor // Write-only binding to the contract
	RLPReaderFilterer   // Log filterer for contract events
}

// RLPReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPReaderSession struct {
	Contract     *RLPReader        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPReaderCallerSession struct {
	Contract *RLPReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPReaderTransactorSession struct {
	Contract     *RLPReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPReaderRaw struct {
	Contract *RLPReader // Generic contract binding to access the raw methods on
}

// RLPReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPReaderCallerRaw struct {
	Contract *RLPReaderCaller // Generic read-only contract binding to access the raw methods on
}

// RLPReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPReaderTransactorRaw struct {
	Contract *RLPReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPReader creates a new instance of RLPReader, bound to a specific deployed contract.
func NewRLPReader(address common.Address, backend bind.ContractBackend) (*RLPReader, error) {
	contract, err := bindRLPReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPReader{RLPReaderCaller: RLPReaderCaller{contract: contract}, RLPReaderTransactor: RLPReaderTransactor{contract: contract}, RLPReaderFilterer: RLPReaderFilterer{contract: contract}}, nil
}

// NewRLPReaderCaller creates a new read-only instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderCaller(address common.Address, caller bind.ContractCaller) (*RLPReaderCaller, error) {
	contract, err := bindRLPReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPReaderCaller{contract: contract}, nil
}

// NewRLPReaderTransactor creates a new write-only instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPReaderTransactor, error) {
	contract, err := bindRLPReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPReaderTransactor{contract: contract}, nil
}

// NewRLPReaderFilterer creates a new log filterer instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPReaderFilterer, error) {
	contract, err := bindRLPReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPReaderFilterer{contract: contract}, nil
}

// bindRLPReader binds a generic wrapper to an already deployed contract.
func bindRLPReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RLPReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPReader *RLPReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPReader.Contract.RLPReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPReader *RLPReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPReader.Contract.RLPReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPReader *RLPReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPReader.Contract.RLPReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPReader *RLPReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPReader *RLPReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPReader *RLPReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPReader.Contract.contract.Transact(opts, method, params...)
}

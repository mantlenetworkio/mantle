// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DeserializationLib

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

// DeserializationLibMetaData contains all meta data concerning the DeserializationLib contract.
var DeserializationLibMetaData = &bind.MetaData{
	ABI: "[]",
}

// DeserializationLibABI is the input ABI used to generate the binding from.
// Deprecated: Use DeserializationLibMetaData.ABI instead.
var DeserializationLibABI = DeserializationLibMetaData.ABI

// DeserializationLib is an auto generated Go binding around an Ethereum contract.
type DeserializationLib struct {
	DeserializationLibCaller     // Read-only binding to the contract
	DeserializationLibTransactor // Write-only binding to the contract
	DeserializationLibFilterer   // Log filterer for contract events
}

// DeserializationLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeserializationLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeserializationLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeserializationLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeserializationLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeserializationLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeserializationLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeserializationLibSession struct {
	Contract     *DeserializationLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DeserializationLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeserializationLibCallerSession struct {
	Contract *DeserializationLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DeserializationLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeserializationLibTransactorSession struct {
	Contract     *DeserializationLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DeserializationLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeserializationLibRaw struct {
	Contract *DeserializationLib // Generic contract binding to access the raw methods on
}

// DeserializationLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeserializationLibCallerRaw struct {
	Contract *DeserializationLibCaller // Generic read-only contract binding to access the raw methods on
}

// DeserializationLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeserializationLibTransactorRaw struct {
	Contract *DeserializationLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeserializationLib creates a new instance of DeserializationLib, bound to a specific deployed contract.
func NewDeserializationLib(address common.Address, backend bind.ContractBackend) (*DeserializationLib, error) {
	contract, err := bindDeserializationLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeserializationLib{DeserializationLibCaller: DeserializationLibCaller{contract: contract}, DeserializationLibTransactor: DeserializationLibTransactor{contract: contract}, DeserializationLibFilterer: DeserializationLibFilterer{contract: contract}}, nil
}

// NewDeserializationLibCaller creates a new read-only instance of DeserializationLib, bound to a specific deployed contract.
func NewDeserializationLibCaller(address common.Address, caller bind.ContractCaller) (*DeserializationLibCaller, error) {
	contract, err := bindDeserializationLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeserializationLibCaller{contract: contract}, nil
}

// NewDeserializationLibTransactor creates a new write-only instance of DeserializationLib, bound to a specific deployed contract.
func NewDeserializationLibTransactor(address common.Address, transactor bind.ContractTransactor) (*DeserializationLibTransactor, error) {
	contract, err := bindDeserializationLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeserializationLibTransactor{contract: contract}, nil
}

// NewDeserializationLibFilterer creates a new log filterer instance of DeserializationLib, bound to a specific deployed contract.
func NewDeserializationLibFilterer(address common.Address, filterer bind.ContractFilterer) (*DeserializationLibFilterer, error) {
	contract, err := bindDeserializationLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeserializationLibFilterer{contract: contract}, nil
}

// bindDeserializationLib binds a generic wrapper to an already deployed contract.
func bindDeserializationLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeserializationLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeserializationLib *DeserializationLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeserializationLib.Contract.DeserializationLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeserializationLib *DeserializationLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeserializationLib.Contract.DeserializationLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeserializationLib *DeserializationLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeserializationLib.Contract.DeserializationLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeserializationLib *DeserializationLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeserializationLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeserializationLib *DeserializationLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeserializationLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeserializationLib *DeserializationLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeserializationLib.Contract.contract.Transact(opts, method, params...)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VerificationContext

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

// VerificationContextMetaData contains all meta data concerning the VerificationContext contract.
var VerificationContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// VerificationContextABI is the input ABI used to generate the binding from.
// Deprecated: Use VerificationContextMetaData.ABI instead.
var VerificationContextABI = VerificationContextMetaData.ABI

// VerificationContext is an auto generated Go binding around an Ethereum contract.
type VerificationContext struct {
	VerificationContextCaller     // Read-only binding to the contract
	VerificationContextTransactor // Write-only binding to the contract
	VerificationContextFilterer   // Log filterer for contract events
}

// VerificationContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerificationContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerificationContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerificationContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerificationContextSession struct {
	Contract     *VerificationContext // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VerificationContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerificationContextCallerSession struct {
	Contract *VerificationContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// VerificationContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerificationContextTransactorSession struct {
	Contract     *VerificationContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// VerificationContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerificationContextRaw struct {
	Contract *VerificationContext // Generic contract binding to access the raw methods on
}

// VerificationContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerificationContextCallerRaw struct {
	Contract *VerificationContextCaller // Generic read-only contract binding to access the raw methods on
}

// VerificationContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerificationContextTransactorRaw struct {
	Contract *VerificationContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerificationContext creates a new instance of VerificationContext, bound to a specific deployed contract.
func NewVerificationContext(address common.Address, backend bind.ContractBackend) (*VerificationContext, error) {
	contract, err := bindVerificationContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerificationContext{VerificationContextCaller: VerificationContextCaller{contract: contract}, VerificationContextTransactor: VerificationContextTransactor{contract: contract}, VerificationContextFilterer: VerificationContextFilterer{contract: contract}}, nil
}

// NewVerificationContextCaller creates a new read-only instance of VerificationContext, bound to a specific deployed contract.
func NewVerificationContextCaller(address common.Address, caller bind.ContractCaller) (*VerificationContextCaller, error) {
	contract, err := bindVerificationContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationContextCaller{contract: contract}, nil
}

// NewVerificationContextTransactor creates a new write-only instance of VerificationContext, bound to a specific deployed contract.
func NewVerificationContextTransactor(address common.Address, transactor bind.ContractTransactor) (*VerificationContextTransactor, error) {
	contract, err := bindVerificationContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationContextTransactor{contract: contract}, nil
}

// NewVerificationContextFilterer creates a new log filterer instance of VerificationContext, bound to a specific deployed contract.
func NewVerificationContextFilterer(address common.Address, filterer bind.ContractFilterer) (*VerificationContextFilterer, error) {
	contract, err := bindVerificationContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerificationContextFilterer{contract: contract}, nil
}

// bindVerificationContext binds a generic wrapper to an already deployed contract.
func bindVerificationContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerificationContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerificationContext *VerificationContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerificationContext.Contract.VerificationContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerificationContext *VerificationContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerificationContext.Contract.VerificationContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerificationContext *VerificationContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerificationContext.Contract.VerificationContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerificationContext *VerificationContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerificationContext.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerificationContext *VerificationContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerificationContext.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerificationContext *VerificationContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerificationContext.Contract.contract.Transact(opts, method, params...)
}

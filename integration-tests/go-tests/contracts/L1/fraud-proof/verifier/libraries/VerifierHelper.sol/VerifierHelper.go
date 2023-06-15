// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VerifierHelper

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

// VerifierHelperMetaData contains all meta data concerning the VerifierHelper contract.
var VerifierHelperMetaData = &bind.MetaData{
	ABI: "[]",
}

// VerifierHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierHelperMetaData.ABI instead.
var VerifierHelperABI = VerifierHelperMetaData.ABI

// VerifierHelper is an auto generated Go binding around an Ethereum contract.
type VerifierHelper struct {
	VerifierHelperCaller     // Read-only binding to the contract
	VerifierHelperTransactor // Write-only binding to the contract
	VerifierHelperFilterer   // Log filterer for contract events
}

// VerifierHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierHelperSession struct {
	Contract     *VerifierHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierHelperCallerSession struct {
	Contract *VerifierHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VerifierHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierHelperTransactorSession struct {
	Contract     *VerifierHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VerifierHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierHelperRaw struct {
	Contract *VerifierHelper // Generic contract binding to access the raw methods on
}

// VerifierHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierHelperCallerRaw struct {
	Contract *VerifierHelperCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierHelperTransactorRaw struct {
	Contract *VerifierHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifierHelper creates a new instance of VerifierHelper, bound to a specific deployed contract.
func NewVerifierHelper(address common.Address, backend bind.ContractBackend) (*VerifierHelper, error) {
	contract, err := bindVerifierHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierHelper{VerifierHelperCaller: VerifierHelperCaller{contract: contract}, VerifierHelperTransactor: VerifierHelperTransactor{contract: contract}, VerifierHelperFilterer: VerifierHelperFilterer{contract: contract}}, nil
}

// NewVerifierHelperCaller creates a new read-only instance of VerifierHelper, bound to a specific deployed contract.
func NewVerifierHelperCaller(address common.Address, caller bind.ContractCaller) (*VerifierHelperCaller, error) {
	contract, err := bindVerifierHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierHelperCaller{contract: contract}, nil
}

// NewVerifierHelperTransactor creates a new write-only instance of VerifierHelper, bound to a specific deployed contract.
func NewVerifierHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierHelperTransactor, error) {
	contract, err := bindVerifierHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierHelperTransactor{contract: contract}, nil
}

// NewVerifierHelperFilterer creates a new log filterer instance of VerifierHelper, bound to a specific deployed contract.
func NewVerifierHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierHelperFilterer, error) {
	contract, err := bindVerifierHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierHelperFilterer{contract: contract}, nil
}

// bindVerifierHelper binds a generic wrapper to an already deployed contract.
func bindVerifierHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierHelper *VerifierHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierHelper.Contract.VerifierHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierHelper *VerifierHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierHelper.Contract.VerifierHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierHelper *VerifierHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierHelper.Contract.VerifierHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierHelper *VerifierHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierHelper *VerifierHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierHelper *VerifierHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierHelper.Contract.contract.Transact(opts, method, params...)
}

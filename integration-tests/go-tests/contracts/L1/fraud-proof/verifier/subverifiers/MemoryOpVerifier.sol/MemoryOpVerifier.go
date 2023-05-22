// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MemoryOpVerifier

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

// MemoryOpVerifierMetaData contains all meta data concerning the MemoryOpVerifier contract.
var MemoryOpVerifierMetaData = &bind.MetaData{
	ABI: "[]",
}

// MemoryOpVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use MemoryOpVerifierMetaData.ABI instead.
var MemoryOpVerifierABI = MemoryOpVerifierMetaData.ABI

// MemoryOpVerifier is an auto generated Go binding around an Ethereum contract.
type MemoryOpVerifier struct {
	MemoryOpVerifierCaller     // Read-only binding to the contract
	MemoryOpVerifierTransactor // Write-only binding to the contract
	MemoryOpVerifierFilterer   // Log filterer for contract events
}

// MemoryOpVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemoryOpVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryOpVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemoryOpVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryOpVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemoryOpVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryOpVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemoryOpVerifierSession struct {
	Contract     *MemoryOpVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryOpVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemoryOpVerifierCallerSession struct {
	Contract *MemoryOpVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MemoryOpVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemoryOpVerifierTransactorSession struct {
	Contract     *MemoryOpVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MemoryOpVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemoryOpVerifierRaw struct {
	Contract *MemoryOpVerifier // Generic contract binding to access the raw methods on
}

// MemoryOpVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemoryOpVerifierCallerRaw struct {
	Contract *MemoryOpVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// MemoryOpVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemoryOpVerifierTransactorRaw struct {
	Contract *MemoryOpVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemoryOpVerifier creates a new instance of MemoryOpVerifier, bound to a specific deployed contract.
func NewMemoryOpVerifier(address common.Address, backend bind.ContractBackend) (*MemoryOpVerifier, error) {
	contract, err := bindMemoryOpVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MemoryOpVerifier{MemoryOpVerifierCaller: MemoryOpVerifierCaller{contract: contract}, MemoryOpVerifierTransactor: MemoryOpVerifierTransactor{contract: contract}, MemoryOpVerifierFilterer: MemoryOpVerifierFilterer{contract: contract}}, nil
}

// NewMemoryOpVerifierCaller creates a new read-only instance of MemoryOpVerifier, bound to a specific deployed contract.
func NewMemoryOpVerifierCaller(address common.Address, caller bind.ContractCaller) (*MemoryOpVerifierCaller, error) {
	contract, err := bindMemoryOpVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryOpVerifierCaller{contract: contract}, nil
}

// NewMemoryOpVerifierTransactor creates a new write-only instance of MemoryOpVerifier, bound to a specific deployed contract.
func NewMemoryOpVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*MemoryOpVerifierTransactor, error) {
	contract, err := bindMemoryOpVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryOpVerifierTransactor{contract: contract}, nil
}

// NewMemoryOpVerifierFilterer creates a new log filterer instance of MemoryOpVerifier, bound to a specific deployed contract.
func NewMemoryOpVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*MemoryOpVerifierFilterer, error) {
	contract, err := bindMemoryOpVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemoryOpVerifierFilterer{contract: contract}, nil
}

// bindMemoryOpVerifier binds a generic wrapper to an already deployed contract.
func bindMemoryOpVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MemoryOpVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemoryOpVerifier *MemoryOpVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemoryOpVerifier.Contract.MemoryOpVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemoryOpVerifier *MemoryOpVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemoryOpVerifier.Contract.MemoryOpVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemoryOpVerifier *MemoryOpVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemoryOpVerifier.Contract.MemoryOpVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemoryOpVerifier *MemoryOpVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemoryOpVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemoryOpVerifier *MemoryOpVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemoryOpVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemoryOpVerifier *MemoryOpVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemoryOpVerifier.Contract.contract.Transact(opts, method, params...)
}

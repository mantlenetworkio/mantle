// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StorageOpVerifier

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

// StorageOpVerifierMetaData contains all meta data concerning the StorageOpVerifier contract.
var StorageOpVerifierMetaData = &bind.MetaData{
	ABI: "[]",
}

// StorageOpVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageOpVerifierMetaData.ABI instead.
var StorageOpVerifierABI = StorageOpVerifierMetaData.ABI

// StorageOpVerifier is an auto generated Go binding around an Ethereum contract.
type StorageOpVerifier struct {
	StorageOpVerifierCaller     // Read-only binding to the contract
	StorageOpVerifierTransactor // Write-only binding to the contract
	StorageOpVerifierFilterer   // Log filterer for contract events
}

// StorageOpVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageOpVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageOpVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageOpVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageOpVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageOpVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageOpVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageOpVerifierSession struct {
	Contract     *StorageOpVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageOpVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageOpVerifierCallerSession struct {
	Contract *StorageOpVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StorageOpVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageOpVerifierTransactorSession struct {
	Contract     *StorageOpVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StorageOpVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageOpVerifierRaw struct {
	Contract *StorageOpVerifier // Generic contract binding to access the raw methods on
}

// StorageOpVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageOpVerifierCallerRaw struct {
	Contract *StorageOpVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// StorageOpVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageOpVerifierTransactorRaw struct {
	Contract *StorageOpVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageOpVerifier creates a new instance of StorageOpVerifier, bound to a specific deployed contract.
func NewStorageOpVerifier(address common.Address, backend bind.ContractBackend) (*StorageOpVerifier, error) {
	contract, err := bindStorageOpVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageOpVerifier{StorageOpVerifierCaller: StorageOpVerifierCaller{contract: contract}, StorageOpVerifierTransactor: StorageOpVerifierTransactor{contract: contract}, StorageOpVerifierFilterer: StorageOpVerifierFilterer{contract: contract}}, nil
}

// NewStorageOpVerifierCaller creates a new read-only instance of StorageOpVerifier, bound to a specific deployed contract.
func NewStorageOpVerifierCaller(address common.Address, caller bind.ContractCaller) (*StorageOpVerifierCaller, error) {
	contract, err := bindStorageOpVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageOpVerifierCaller{contract: contract}, nil
}

// NewStorageOpVerifierTransactor creates a new write-only instance of StorageOpVerifier, bound to a specific deployed contract.
func NewStorageOpVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageOpVerifierTransactor, error) {
	contract, err := bindStorageOpVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageOpVerifierTransactor{contract: contract}, nil
}

// NewStorageOpVerifierFilterer creates a new log filterer instance of StorageOpVerifier, bound to a specific deployed contract.
func NewStorageOpVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageOpVerifierFilterer, error) {
	contract, err := bindStorageOpVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageOpVerifierFilterer{contract: contract}, nil
}

// bindStorageOpVerifier binds a generic wrapper to an already deployed contract.
func bindStorageOpVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageOpVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageOpVerifier *StorageOpVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageOpVerifier.Contract.StorageOpVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageOpVerifier *StorageOpVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageOpVerifier.Contract.StorageOpVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageOpVerifier *StorageOpVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageOpVerifier.Contract.StorageOpVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageOpVerifier *StorageOpVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageOpVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageOpVerifier *StorageOpVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageOpVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageOpVerifier *StorageOpVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageOpVerifier.Contract.contract.Transact(opts, method, params...)
}

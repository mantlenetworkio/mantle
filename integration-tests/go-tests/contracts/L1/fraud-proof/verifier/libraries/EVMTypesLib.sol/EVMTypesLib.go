// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EVMTypesLib

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

// EVMTypesLibMetaData contains all meta data concerning the EVMTypesLib contract.
var EVMTypesLibMetaData = &bind.MetaData{
	ABI: "[]",
}

// EVMTypesLibABI is the input ABI used to generate the binding from.
// Deprecated: Use EVMTypesLibMetaData.ABI instead.
var EVMTypesLibABI = EVMTypesLibMetaData.ABI

// EVMTypesLib is an auto generated Go binding around an Ethereum contract.
type EVMTypesLib struct {
	EVMTypesLibCaller     // Read-only binding to the contract
	EVMTypesLibTransactor // Write-only binding to the contract
	EVMTypesLibFilterer   // Log filterer for contract events
}

// EVMTypesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type EVMTypesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVMTypesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EVMTypesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVMTypesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EVMTypesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVMTypesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EVMTypesLibSession struct {
	Contract     *EVMTypesLib      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EVMTypesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EVMTypesLibCallerSession struct {
	Contract *EVMTypesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EVMTypesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EVMTypesLibTransactorSession struct {
	Contract     *EVMTypesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EVMTypesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type EVMTypesLibRaw struct {
	Contract *EVMTypesLib // Generic contract binding to access the raw methods on
}

// EVMTypesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EVMTypesLibCallerRaw struct {
	Contract *EVMTypesLibCaller // Generic read-only contract binding to access the raw methods on
}

// EVMTypesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EVMTypesLibTransactorRaw struct {
	Contract *EVMTypesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEVMTypesLib creates a new instance of EVMTypesLib, bound to a specific deployed contract.
func NewEVMTypesLib(address common.Address, backend bind.ContractBackend) (*EVMTypesLib, error) {
	contract, err := bindEVMTypesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVMTypesLib{EVMTypesLibCaller: EVMTypesLibCaller{contract: contract}, EVMTypesLibTransactor: EVMTypesLibTransactor{contract: contract}, EVMTypesLibFilterer: EVMTypesLibFilterer{contract: contract}}, nil
}

// NewEVMTypesLibCaller creates a new read-only instance of EVMTypesLib, bound to a specific deployed contract.
func NewEVMTypesLibCaller(address common.Address, caller bind.ContractCaller) (*EVMTypesLibCaller, error) {
	contract, err := bindEVMTypesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVMTypesLibCaller{contract: contract}, nil
}

// NewEVMTypesLibTransactor creates a new write-only instance of EVMTypesLib, bound to a specific deployed contract.
func NewEVMTypesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*EVMTypesLibTransactor, error) {
	contract, err := bindEVMTypesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVMTypesLibTransactor{contract: contract}, nil
}

// NewEVMTypesLibFilterer creates a new log filterer instance of EVMTypesLib, bound to a specific deployed contract.
func NewEVMTypesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*EVMTypesLibFilterer, error) {
	contract, err := bindEVMTypesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVMTypesLibFilterer{contract: contract}, nil
}

// bindEVMTypesLib binds a generic wrapper to an already deployed contract.
func bindEVMTypesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EVMTypesLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EVMTypesLib *EVMTypesLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVMTypesLib.Contract.EVMTypesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EVMTypesLib *EVMTypesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVMTypesLib.Contract.EVMTypesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EVMTypesLib *EVMTypesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVMTypesLib.Contract.EVMTypesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EVMTypesLib *EVMTypesLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVMTypesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EVMTypesLib *EVMTypesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVMTypesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EVMTypesLib *EVMTypesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVMTypesLib.Contract.contract.Transact(opts, method, params...)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ChallengeLib

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

// ChallengeLibMetaData contains all meta data concerning the ChallengeLib contract.
var ChallengeLibMetaData = &bind.MetaData{
	ABI: "[]",
}

// ChallengeLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ChallengeLibMetaData.ABI instead.
var ChallengeLibABI = ChallengeLibMetaData.ABI

// ChallengeLib is an auto generated Go binding around an Ethereum contract.
type ChallengeLib struct {
	ChallengeLibCaller     // Read-only binding to the contract
	ChallengeLibTransactor // Write-only binding to the contract
	ChallengeLibFilterer   // Log filterer for contract events
}

// ChallengeLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeLibSession struct {
	Contract     *ChallengeLib     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeLibCallerSession struct {
	Contract *ChallengeLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ChallengeLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeLibTransactorSession struct {
	Contract     *ChallengeLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ChallengeLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeLibRaw struct {
	Contract *ChallengeLib // Generic contract binding to access the raw methods on
}

// ChallengeLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeLibCallerRaw struct {
	Contract *ChallengeLibCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeLibTransactorRaw struct {
	Contract *ChallengeLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeLib creates a new instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLib(address common.Address, backend bind.ContractBackend) (*ChallengeLib, error) {
	contract, err := bindChallengeLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeLib{ChallengeLibCaller: ChallengeLibCaller{contract: contract}, ChallengeLibTransactor: ChallengeLibTransactor{contract: contract}, ChallengeLibFilterer: ChallengeLibFilterer{contract: contract}}, nil
}

// NewChallengeLibCaller creates a new read-only instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLibCaller(address common.Address, caller bind.ContractCaller) (*ChallengeLibCaller, error) {
	contract, err := bindChallengeLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeLibCaller{contract: contract}, nil
}

// NewChallengeLibTransactor creates a new write-only instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeLibTransactor, error) {
	contract, err := bindChallengeLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeLibTransactor{contract: contract}, nil
}

// NewChallengeLibFilterer creates a new log filterer instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeLibFilterer, error) {
	contract, err := bindChallengeLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeLibFilterer{contract: contract}, nil
}

// bindChallengeLib binds a generic wrapper to an already deployed contract.
func bindChallengeLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChallengeLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeLib *ChallengeLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeLib.Contract.ChallengeLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeLib *ChallengeLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeLib.Contract.ChallengeLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeLib *ChallengeLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeLib.Contract.ChallengeLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeLib *ChallengeLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeLib *ChallengeLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeLib *ChallengeLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeLib.Contract.contract.Transact(opts, method, params...)
}

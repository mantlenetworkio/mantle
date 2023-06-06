// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Params

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

// ParamsMetaData contains all meta data concerning the Params contract.
var ParamsMetaData = &bind.MetaData{
	ABI: "[]",
}

// ParamsABI is the input ABI used to generate the binding from.
// Deprecated: Use ParamsMetaData.ABI instead.
var ParamsABI = ParamsMetaData.ABI

// Params is an auto generated Go binding around an Ethereum contract.
type Params struct {
	ParamsCaller     // Read-only binding to the contract
	ParamsTransactor // Write-only binding to the contract
	ParamsFilterer   // Log filterer for contract events
}

// ParamsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParamsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParamsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParamsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParamsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParamsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParamsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParamsSession struct {
	Contract     *Params           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParamsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParamsCallerSession struct {
	Contract *ParamsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ParamsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParamsTransactorSession struct {
	Contract     *ParamsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParamsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParamsRaw struct {
	Contract *Params // Generic contract binding to access the raw methods on
}

// ParamsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParamsCallerRaw struct {
	Contract *ParamsCaller // Generic read-only contract binding to access the raw methods on
}

// ParamsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParamsTransactorRaw struct {
	Contract *ParamsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParams creates a new instance of Params, bound to a specific deployed contract.
func NewParams(address common.Address, backend bind.ContractBackend) (*Params, error) {
	contract, err := bindParams(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Params{ParamsCaller: ParamsCaller{contract: contract}, ParamsTransactor: ParamsTransactor{contract: contract}, ParamsFilterer: ParamsFilterer{contract: contract}}, nil
}

// NewParamsCaller creates a new read-only instance of Params, bound to a specific deployed contract.
func NewParamsCaller(address common.Address, caller bind.ContractCaller) (*ParamsCaller, error) {
	contract, err := bindParams(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParamsCaller{contract: contract}, nil
}

// NewParamsTransactor creates a new write-only instance of Params, bound to a specific deployed contract.
func NewParamsTransactor(address common.Address, transactor bind.ContractTransactor) (*ParamsTransactor, error) {
	contract, err := bindParams(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParamsTransactor{contract: contract}, nil
}

// NewParamsFilterer creates a new log filterer instance of Params, bound to a specific deployed contract.
func NewParamsFilterer(address common.Address, filterer bind.ContractFilterer) (*ParamsFilterer, error) {
	contract, err := bindParams(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParamsFilterer{contract: contract}, nil
}

// bindParams binds a generic wrapper to an already deployed contract.
func bindParams(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParamsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Params *ParamsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Params.Contract.ParamsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Params *ParamsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Params.Contract.ParamsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Params *ParamsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Params.Contract.ParamsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Params *ParamsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Params.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Params *ParamsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Params.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Params *ParamsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Params.Contract.contract.Transact(opts, method, params...)
}

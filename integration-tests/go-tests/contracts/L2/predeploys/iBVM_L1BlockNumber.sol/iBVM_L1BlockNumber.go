// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iBVM_L1BlockNumber

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
)

// IBVML1BlockNumberMetaData contains all meta data concerning the IBVML1BlockNumber contract.
var IBVML1BlockNumberMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getL1BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBVML1BlockNumberABI is the input ABI used to generate the binding from.
// Deprecated: Use IBVML1BlockNumberMetaData.ABI instead.
var IBVML1BlockNumberABI = IBVML1BlockNumberMetaData.ABI

// IBVML1BlockNumber is an auto generated Go binding around an Ethereum contract.
type IBVML1BlockNumber struct {
	IBVML1BlockNumberCaller     // Read-only binding to the contract
	IBVML1BlockNumberTransactor // Write-only binding to the contract
	IBVML1BlockNumberFilterer   // Log filterer for contract events
}

// IBVML1BlockNumberCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBVML1BlockNumberCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBVML1BlockNumberTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBVML1BlockNumberTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBVML1BlockNumberFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBVML1BlockNumberFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBVML1BlockNumberSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBVML1BlockNumberSession struct {
	Contract     *IBVML1BlockNumber // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IBVML1BlockNumberCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBVML1BlockNumberCallerSession struct {
	Contract *IBVML1BlockNumberCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IBVML1BlockNumberTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBVML1BlockNumberTransactorSession struct {
	Contract     *IBVML1BlockNumberTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IBVML1BlockNumberRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBVML1BlockNumberRaw struct {
	Contract *IBVML1BlockNumber // Generic contract binding to access the raw methods on
}

// IBVML1BlockNumberCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBVML1BlockNumberCallerRaw struct {
	Contract *IBVML1BlockNumberCaller // Generic read-only contract binding to access the raw methods on
}

// IBVML1BlockNumberTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBVML1BlockNumberTransactorRaw struct {
	Contract *IBVML1BlockNumberTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBVML1BlockNumber creates a new instance of IBVML1BlockNumber, bound to a specific deployed contract.
func NewIBVML1BlockNumber(address common.Address, backend bind.ContractBackend) (*IBVML1BlockNumber, error) {
	contract, err := bindIBVML1BlockNumber(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBVML1BlockNumber{IBVML1BlockNumberCaller: IBVML1BlockNumberCaller{contract: contract}, IBVML1BlockNumberTransactor: IBVML1BlockNumberTransactor{contract: contract}, IBVML1BlockNumberFilterer: IBVML1BlockNumberFilterer{contract: contract}}, nil
}

// NewIBVML1BlockNumberCaller creates a new read-only instance of IBVML1BlockNumber, bound to a specific deployed contract.
func NewIBVML1BlockNumberCaller(address common.Address, caller bind.ContractCaller) (*IBVML1BlockNumberCaller, error) {
	contract, err := bindIBVML1BlockNumber(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBVML1BlockNumberCaller{contract: contract}, nil
}

// NewIBVML1BlockNumberTransactor creates a new write-only instance of IBVML1BlockNumber, bound to a specific deployed contract.
func NewIBVML1BlockNumberTransactor(address common.Address, transactor bind.ContractTransactor) (*IBVML1BlockNumberTransactor, error) {
	contract, err := bindIBVML1BlockNumber(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBVML1BlockNumberTransactor{contract: contract}, nil
}

// NewIBVML1BlockNumberFilterer creates a new log filterer instance of IBVML1BlockNumber, bound to a specific deployed contract.
func NewIBVML1BlockNumberFilterer(address common.Address, filterer bind.ContractFilterer) (*IBVML1BlockNumberFilterer, error) {
	contract, err := bindIBVML1BlockNumber(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBVML1BlockNumberFilterer{contract: contract}, nil
}

// bindIBVML1BlockNumber binds a generic wrapper to an already deployed contract.
func bindIBVML1BlockNumber(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBVML1BlockNumberABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBVML1BlockNumber *IBVML1BlockNumberRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBVML1BlockNumber.Contract.IBVML1BlockNumberCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBVML1BlockNumber *IBVML1BlockNumberRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBVML1BlockNumber.Contract.IBVML1BlockNumberTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBVML1BlockNumber *IBVML1BlockNumberRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBVML1BlockNumber.Contract.IBVML1BlockNumberTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBVML1BlockNumber *IBVML1BlockNumberCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBVML1BlockNumber.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBVML1BlockNumber *IBVML1BlockNumberTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBVML1BlockNumber.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBVML1BlockNumber *IBVML1BlockNumberTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBVML1BlockNumber.Contract.contract.Transact(opts, method, params...)
}

// GetL1BlockNumber is a free data retrieval call binding the contract method 0xb9b3efe9.
//
// Solidity: function getL1BlockNumber() view returns(uint256)
func (_IBVML1BlockNumber *IBVML1BlockNumberCaller) GetL1BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBVML1BlockNumber.contract.Call(opts, &out, "getL1BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL1BlockNumber is a free data retrieval call binding the contract method 0xb9b3efe9.
//
// Solidity: function getL1BlockNumber() view returns(uint256)
func (_IBVML1BlockNumber *IBVML1BlockNumberSession) GetL1BlockNumber() (*big.Int, error) {
	return _IBVML1BlockNumber.Contract.GetL1BlockNumber(&_IBVML1BlockNumber.CallOpts)
}

// GetL1BlockNumber is a free data retrieval call binding the contract method 0xb9b3efe9.
//
// Solidity: function getL1BlockNumber() view returns(uint256)
func (_IBVML1BlockNumber *IBVML1BlockNumberCallerSession) GetL1BlockNumber() (*big.Int, error) {
	return _IBVML1BlockNumber.Contract.GetL1BlockNumber(&_IBVML1BlockNumber.CallOpts)
}

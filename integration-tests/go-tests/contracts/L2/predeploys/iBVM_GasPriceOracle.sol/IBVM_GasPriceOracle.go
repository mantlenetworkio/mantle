// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IBVM_GasPriceOracle

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

// IBVMGasPriceOracleMetaData contains all meta data concerning the IBVMGasPriceOracle contract.
var IBVMGasPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IsBurning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBVMGasPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IBVMGasPriceOracleMetaData.ABI instead.
var IBVMGasPriceOracleABI = IBVMGasPriceOracleMetaData.ABI

// IBVMGasPriceOracle is an auto generated Go binding around an Ethereum contract.
type IBVMGasPriceOracle struct {
	IBVMGasPriceOracleCaller     // Read-only binding to the contract
	IBVMGasPriceOracleTransactor // Write-only binding to the contract
	IBVMGasPriceOracleFilterer   // Log filterer for contract events
}

// IBVMGasPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBVMGasPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBVMGasPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBVMGasPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBVMGasPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBVMGasPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBVMGasPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBVMGasPriceOracleSession struct {
	Contract     *IBVMGasPriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IBVMGasPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBVMGasPriceOracleCallerSession struct {
	Contract *IBVMGasPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IBVMGasPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBVMGasPriceOracleTransactorSession struct {
	Contract     *IBVMGasPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IBVMGasPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBVMGasPriceOracleRaw struct {
	Contract *IBVMGasPriceOracle // Generic contract binding to access the raw methods on
}

// IBVMGasPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBVMGasPriceOracleCallerRaw struct {
	Contract *IBVMGasPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IBVMGasPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBVMGasPriceOracleTransactorRaw struct {
	Contract *IBVMGasPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBVMGasPriceOracle creates a new instance of IBVMGasPriceOracle, bound to a specific deployed contract.
func NewIBVMGasPriceOracle(address common.Address, backend bind.ContractBackend) (*IBVMGasPriceOracle, error) {
	contract, err := bindIBVMGasPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBVMGasPriceOracle{IBVMGasPriceOracleCaller: IBVMGasPriceOracleCaller{contract: contract}, IBVMGasPriceOracleTransactor: IBVMGasPriceOracleTransactor{contract: contract}, IBVMGasPriceOracleFilterer: IBVMGasPriceOracleFilterer{contract: contract}}, nil
}

// NewIBVMGasPriceOracleCaller creates a new read-only instance of IBVMGasPriceOracle, bound to a specific deployed contract.
func NewIBVMGasPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*IBVMGasPriceOracleCaller, error) {
	contract, err := bindIBVMGasPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBVMGasPriceOracleCaller{contract: contract}, nil
}

// NewIBVMGasPriceOracleTransactor creates a new write-only instance of IBVMGasPriceOracle, bound to a specific deployed contract.
func NewIBVMGasPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IBVMGasPriceOracleTransactor, error) {
	contract, err := bindIBVMGasPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBVMGasPriceOracleTransactor{contract: contract}, nil
}

// NewIBVMGasPriceOracleFilterer creates a new log filterer instance of IBVMGasPriceOracle, bound to a specific deployed contract.
func NewIBVMGasPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IBVMGasPriceOracleFilterer, error) {
	contract, err := bindIBVMGasPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBVMGasPriceOracleFilterer{contract: contract}, nil
}

// bindIBVMGasPriceOracle binds a generic wrapper to an already deployed contract.
func bindIBVMGasPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBVMGasPriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBVMGasPriceOracle *IBVMGasPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBVMGasPriceOracle.Contract.IBVMGasPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBVMGasPriceOracle *IBVMGasPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBVMGasPriceOracle.Contract.IBVMGasPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBVMGasPriceOracle *IBVMGasPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBVMGasPriceOracle.Contract.IBVMGasPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBVMGasPriceOracle *IBVMGasPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBVMGasPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBVMGasPriceOracle *IBVMGasPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBVMGasPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBVMGasPriceOracle *IBVMGasPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBVMGasPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// IsBurning is a free data retrieval call binding the contract method 0x0d1e43a0.
//
// Solidity: function IsBurning() view returns(bool)
func (_IBVMGasPriceOracle *IBVMGasPriceOracleCaller) IsBurning(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IBVMGasPriceOracle.contract.Call(opts, &out, "IsBurning")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBurning is a free data retrieval call binding the contract method 0x0d1e43a0.
//
// Solidity: function IsBurning() view returns(bool)
func (_IBVMGasPriceOracle *IBVMGasPriceOracleSession) IsBurning() (bool, error) {
	return _IBVMGasPriceOracle.Contract.IsBurning(&_IBVMGasPriceOracle.CallOpts)
}

// IsBurning is a free data retrieval call binding the contract method 0x0d1e43a0.
//
// Solidity: function IsBurning() view returns(bool)
func (_IBVMGasPriceOracle *IBVMGasPriceOracleCallerSession) IsBurning() (bool, error) {
	return _IBVMGasPriceOracle.Contract.IsBurning(&_IBVMGasPriceOracle.CallOpts)
}

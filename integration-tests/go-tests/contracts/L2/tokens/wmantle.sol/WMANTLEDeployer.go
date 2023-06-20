// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WMANTLEDeployer

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

// WMANTLEDeployerMetaData contains all meta data concerning the WMANTLEDeployer contract.
var WMANTLEDeployerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"calculateAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"predictedAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WMANTLEDeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use WMANTLEDeployerMetaData.ABI instead.
var WMANTLEDeployerABI = WMANTLEDeployerMetaData.ABI

// WMANTLEDeployer is an auto generated Go binding around an Ethereum contract.
type WMANTLEDeployer struct {
	WMANTLEDeployerCaller     // Read-only binding to the contract
	WMANTLEDeployerTransactor // Write-only binding to the contract
	WMANTLEDeployerFilterer   // Log filterer for contract events
}

// WMANTLEDeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WMANTLEDeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WMANTLEDeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WMANTLEDeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WMANTLEDeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WMANTLEDeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WMANTLEDeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WMANTLEDeployerSession struct {
	Contract     *WMANTLEDeployer  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WMANTLEDeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WMANTLEDeployerCallerSession struct {
	Contract *WMANTLEDeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WMANTLEDeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WMANTLEDeployerTransactorSession struct {
	Contract     *WMANTLEDeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WMANTLEDeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WMANTLEDeployerRaw struct {
	Contract *WMANTLEDeployer // Generic contract binding to access the raw methods on
}

// WMANTLEDeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WMANTLEDeployerCallerRaw struct {
	Contract *WMANTLEDeployerCaller // Generic read-only contract binding to access the raw methods on
}

// WMANTLEDeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WMANTLEDeployerTransactorRaw struct {
	Contract *WMANTLEDeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWMANTLEDeployer creates a new instance of WMANTLEDeployer, bound to a specific deployed contract.
func NewWMANTLEDeployer(address common.Address, backend bind.ContractBackend) (*WMANTLEDeployer, error) {
	contract, err := bindWMANTLEDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WMANTLEDeployer{WMANTLEDeployerCaller: WMANTLEDeployerCaller{contract: contract}, WMANTLEDeployerTransactor: WMANTLEDeployerTransactor{contract: contract}, WMANTLEDeployerFilterer: WMANTLEDeployerFilterer{contract: contract}}, nil
}

// NewWMANTLEDeployerCaller creates a new read-only instance of WMANTLEDeployer, bound to a specific deployed contract.
func NewWMANTLEDeployerCaller(address common.Address, caller bind.ContractCaller) (*WMANTLEDeployerCaller, error) {
	contract, err := bindWMANTLEDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WMANTLEDeployerCaller{contract: contract}, nil
}

// NewWMANTLEDeployerTransactor creates a new write-only instance of WMANTLEDeployer, bound to a specific deployed contract.
func NewWMANTLEDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*WMANTLEDeployerTransactor, error) {
	contract, err := bindWMANTLEDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WMANTLEDeployerTransactor{contract: contract}, nil
}

// NewWMANTLEDeployerFilterer creates a new log filterer instance of WMANTLEDeployer, bound to a specific deployed contract.
func NewWMANTLEDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*WMANTLEDeployerFilterer, error) {
	contract, err := bindWMANTLEDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WMANTLEDeployerFilterer{contract: contract}, nil
}

// bindWMANTLEDeployer binds a generic wrapper to an already deployed contract.
func bindWMANTLEDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WMANTLEDeployerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WMANTLEDeployer *WMANTLEDeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WMANTLEDeployer.Contract.WMANTLEDeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WMANTLEDeployer *WMANTLEDeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WMANTLEDeployer.Contract.WMANTLEDeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WMANTLEDeployer *WMANTLEDeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WMANTLEDeployer.Contract.WMANTLEDeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WMANTLEDeployer *WMANTLEDeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WMANTLEDeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WMANTLEDeployer *WMANTLEDeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WMANTLEDeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WMANTLEDeployer *WMANTLEDeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WMANTLEDeployer.Contract.contract.Transact(opts, method, params...)
}

// CalculateAddr is a free data retrieval call binding the contract method 0x6c9ba972.
//
// Solidity: function calculateAddr() view returns(address predictedAddress)
func (_WMANTLEDeployer *WMANTLEDeployerCaller) CalculateAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WMANTLEDeployer.contract.Call(opts, &out, "calculateAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateAddr is a free data retrieval call binding the contract method 0x6c9ba972.
//
// Solidity: function calculateAddr() view returns(address predictedAddress)
func (_WMANTLEDeployer *WMANTLEDeployerSession) CalculateAddr() (common.Address, error) {
	return _WMANTLEDeployer.Contract.CalculateAddr(&_WMANTLEDeployer.CallOpts)
}

// CalculateAddr is a free data retrieval call binding the contract method 0x6c9ba972.
//
// Solidity: function calculateAddr() view returns(address predictedAddress)
func (_WMANTLEDeployer *WMANTLEDeployerCallerSession) CalculateAddr() (common.Address, error) {
	return _WMANTLEDeployer.Contract.CalculateAddr(&_WMANTLEDeployer.CallOpts)
}

// Deploy is a paid mutator transaction binding the contract method 0x775c300c.
//
// Solidity: function deploy() returns(address addr)
func (_WMANTLEDeployer *WMANTLEDeployerTransactor) Deploy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WMANTLEDeployer.contract.Transact(opts, "deploy")
}

// Deploy is a paid mutator transaction binding the contract method 0x775c300c.
//
// Solidity: function deploy() returns(address addr)
func (_WMANTLEDeployer *WMANTLEDeployerSession) Deploy() (*types.Transaction, error) {
	return _WMANTLEDeployer.Contract.Deploy(&_WMANTLEDeployer.TransactOpts)
}

// Deploy is a paid mutator transaction binding the contract method 0x775c300c.
//
// Solidity: function deploy() returns(address addr)
func (_WMANTLEDeployer *WMANTLEDeployerTransactorSession) Deploy() (*types.Transaction, error) {
	return _WMANTLEDeployer.Contract.Deploy(&_WMANTLEDeployer.TransactOpts)
}

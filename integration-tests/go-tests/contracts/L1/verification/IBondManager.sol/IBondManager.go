// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IBondManager

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

// IBondManagerMetaData contains all meta data concerning the IBondManager contract.
var IBondManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"isCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBondManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IBondManagerMetaData.ABI instead.
var IBondManagerABI = IBondManagerMetaData.ABI

// IBondManager is an auto generated Go binding around an Ethereum contract.
type IBondManager struct {
	IBondManagerCaller     // Read-only binding to the contract
	IBondManagerTransactor // Write-only binding to the contract
	IBondManagerFilterer   // Log filterer for contract events
}

// IBondManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBondManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBondManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBondManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBondManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBondManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBondManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBondManagerSession struct {
	Contract     *IBondManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBondManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBondManagerCallerSession struct {
	Contract *IBondManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IBondManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBondManagerTransactorSession struct {
	Contract     *IBondManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IBondManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBondManagerRaw struct {
	Contract *IBondManager // Generic contract binding to access the raw methods on
}

// IBondManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBondManagerCallerRaw struct {
	Contract *IBondManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IBondManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBondManagerTransactorRaw struct {
	Contract *IBondManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBondManager creates a new instance of IBondManager, bound to a specific deployed contract.
func NewIBondManager(address common.Address, backend bind.ContractBackend) (*IBondManager, error) {
	contract, err := bindIBondManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBondManager{IBondManagerCaller: IBondManagerCaller{contract: contract}, IBondManagerTransactor: IBondManagerTransactor{contract: contract}, IBondManagerFilterer: IBondManagerFilterer{contract: contract}}, nil
}

// NewIBondManagerCaller creates a new read-only instance of IBondManager, bound to a specific deployed contract.
func NewIBondManagerCaller(address common.Address, caller bind.ContractCaller) (*IBondManagerCaller, error) {
	contract, err := bindIBondManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBondManagerCaller{contract: contract}, nil
}

// NewIBondManagerTransactor creates a new write-only instance of IBondManager, bound to a specific deployed contract.
func NewIBondManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IBondManagerTransactor, error) {
	contract, err := bindIBondManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBondManagerTransactor{contract: contract}, nil
}

// NewIBondManagerFilterer creates a new log filterer instance of IBondManager, bound to a specific deployed contract.
func NewIBondManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IBondManagerFilterer, error) {
	contract, err := bindIBondManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBondManagerFilterer{contract: contract}, nil
}

// bindIBondManager binds a generic wrapper to an already deployed contract.
func bindIBondManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBondManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBondManager *IBondManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBondManager.Contract.IBondManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBondManager *IBondManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBondManager.Contract.IBondManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBondManager *IBondManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBondManager.Contract.IBondManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBondManager *IBondManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBondManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBondManager *IBondManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBondManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBondManager *IBondManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBondManager.Contract.contract.Transact(opts, method, params...)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x02ad4d2a.
//
// Solidity: function isCollateralized(address _who) view returns(bool)
func (_IBondManager *IBondManagerCaller) IsCollateralized(opts *bind.CallOpts, _who common.Address) (bool, error) {
	var out []interface{}
	err := _IBondManager.contract.Call(opts, &out, "isCollateralized", _who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralized is a free data retrieval call binding the contract method 0x02ad4d2a.
//
// Solidity: function isCollateralized(address _who) view returns(bool)
func (_IBondManager *IBondManagerSession) IsCollateralized(_who common.Address) (bool, error) {
	return _IBondManager.Contract.IsCollateralized(&_IBondManager.CallOpts, _who)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x02ad4d2a.
//
// Solidity: function isCollateralized(address _who) view returns(bool)
func (_IBondManager *IBondManagerCallerSession) IsCollateralized(_who common.Address) (bool, error) {
	return _IBondManager.Contract.IsCollateralized(&_IBondManager.CallOpts, _who)
}

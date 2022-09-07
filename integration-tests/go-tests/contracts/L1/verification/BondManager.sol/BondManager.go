// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BondManager

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

// BondManagerMetaData contains all meta data concerning the BondManager contract.
var BondManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_libAddressManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"isCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libAddressManager\",\"outputs\":[{\"internalType\":\"contractLib_AddressManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BondManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BondManagerMetaData.ABI instead.
var BondManagerABI = BondManagerMetaData.ABI

// BondManager is an auto generated Go binding around an Ethereum contract.
type BondManager struct {
	BondManagerCaller     // Read-only binding to the contract
	BondManagerTransactor // Write-only binding to the contract
	BondManagerFilterer   // Log filterer for contract events
}

// BondManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BondManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BondManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BondManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BondManagerSession struct {
	Contract     *BondManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BondManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BondManagerCallerSession struct {
	Contract *BondManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BondManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BondManagerTransactorSession struct {
	Contract     *BondManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BondManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BondManagerRaw struct {
	Contract *BondManager // Generic contract binding to access the raw methods on
}

// BondManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BondManagerCallerRaw struct {
	Contract *BondManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BondManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BondManagerTransactorRaw struct {
	Contract *BondManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBondManager creates a new instance of BondManager, bound to a specific deployed contract.
func NewBondManager(address common.Address, backend bind.ContractBackend) (*BondManager, error) {
	contract, err := bindBondManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BondManager{BondManagerCaller: BondManagerCaller{contract: contract}, BondManagerTransactor: BondManagerTransactor{contract: contract}, BondManagerFilterer: BondManagerFilterer{contract: contract}}, nil
}

// NewBondManagerCaller creates a new read-only instance of BondManager, bound to a specific deployed contract.
func NewBondManagerCaller(address common.Address, caller bind.ContractCaller) (*BondManagerCaller, error) {
	contract, err := bindBondManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BondManagerCaller{contract: contract}, nil
}

// NewBondManagerTransactor creates a new write-only instance of BondManager, bound to a specific deployed contract.
func NewBondManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BondManagerTransactor, error) {
	contract, err := bindBondManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BondManagerTransactor{contract: contract}, nil
}

// NewBondManagerFilterer creates a new log filterer instance of BondManager, bound to a specific deployed contract.
func NewBondManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BondManagerFilterer, error) {
	contract, err := bindBondManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BondManagerFilterer{contract: contract}, nil
}

// bindBondManager binds a generic wrapper to an already deployed contract.
func bindBondManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BondManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondManager *BondManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondManager.Contract.BondManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondManager *BondManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondManager.Contract.BondManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondManager *BondManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondManager.Contract.BondManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondManager *BondManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondManager *BondManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondManager *BondManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondManager.Contract.contract.Transact(opts, method, params...)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x02ad4d2a.
//
// Solidity: function isCollateralized(address _who) view returns(bool)
func (_BondManager *BondManagerCaller) IsCollateralized(opts *bind.CallOpts, _who common.Address) (bool, error) {
	var out []interface{}
	err := _BondManager.contract.Call(opts, &out, "isCollateralized", _who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralized is a free data retrieval call binding the contract method 0x02ad4d2a.
//
// Solidity: function isCollateralized(address _who) view returns(bool)
func (_BondManager *BondManagerSession) IsCollateralized(_who common.Address) (bool, error) {
	return _BondManager.Contract.IsCollateralized(&_BondManager.CallOpts, _who)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x02ad4d2a.
//
// Solidity: function isCollateralized(address _who) view returns(bool)
func (_BondManager *BondManagerCallerSession) IsCollateralized(_who common.Address) (bool, error) {
	return _BondManager.Contract.IsCollateralized(&_BondManager.CallOpts, _who)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_BondManager *BondManagerCaller) LibAddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BondManager.contract.Call(opts, &out, "libAddressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_BondManager *BondManagerSession) LibAddressManager() (common.Address, error) {
	return _BondManager.Contract.LibAddressManager(&_BondManager.CallOpts)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_BondManager *BondManagerCallerSession) LibAddressManager() (common.Address, error) {
	return _BondManager.Contract.LibAddressManager(&_BondManager.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_BondManager *BondManagerCaller) Resolve(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _BondManager.contract.Call(opts, &out, "resolve", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_BondManager *BondManagerSession) Resolve(_name string) (common.Address, error) {
	return _BondManager.Contract.Resolve(&_BondManager.CallOpts, _name)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_BondManager *BondManagerCallerSession) Resolve(_name string) (common.Address, error) {
	return _BondManager.Contract.Resolve(&_BondManager.CallOpts, _name)
}

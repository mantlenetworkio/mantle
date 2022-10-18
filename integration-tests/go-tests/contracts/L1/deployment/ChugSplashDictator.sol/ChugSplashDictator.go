// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ChugSplashDictator

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

// ChugSplashDictatorMetaData contains all meta data concerning the ChugSplashDictator contract.
var ChugSplashDictatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractL1ChugSplashProxy\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_finalOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_messengerSlotKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_messengerSlotVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bridgeSlotKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bridgeSlotVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bitAddressSlotKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bitAddressSlotVal\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bitAddressSlotKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bitAddressSlotVal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeSlotKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeSlotVal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"codeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_code\",\"type\":\"bytes\"}],\"name\":\"doActions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUpgrading\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messengerSlotKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messengerSlotVal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"returnOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"internalType\":\"contractL1ChugSplashProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ChugSplashDictatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ChugSplashDictatorMetaData.ABI instead.
var ChugSplashDictatorABI = ChugSplashDictatorMetaData.ABI

// ChugSplashDictator is an auto generated Go binding around an Ethereum contract.
type ChugSplashDictator struct {
	ChugSplashDictatorCaller     // Read-only binding to the contract
	ChugSplashDictatorTransactor // Write-only binding to the contract
	ChugSplashDictatorFilterer   // Log filterer for contract events
}

// ChugSplashDictatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChugSplashDictatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChugSplashDictatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChugSplashDictatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChugSplashDictatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChugSplashDictatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChugSplashDictatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChugSplashDictatorSession struct {
	Contract     *ChugSplashDictator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ChugSplashDictatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChugSplashDictatorCallerSession struct {
	Contract *ChugSplashDictatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ChugSplashDictatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChugSplashDictatorTransactorSession struct {
	Contract     *ChugSplashDictatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ChugSplashDictatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChugSplashDictatorRaw struct {
	Contract *ChugSplashDictator // Generic contract binding to access the raw methods on
}

// ChugSplashDictatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChugSplashDictatorCallerRaw struct {
	Contract *ChugSplashDictatorCaller // Generic read-only contract binding to access the raw methods on
}

// ChugSplashDictatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChugSplashDictatorTransactorRaw struct {
	Contract *ChugSplashDictatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChugSplashDictator creates a new instance of ChugSplashDictator, bound to a specific deployed contract.
func NewChugSplashDictator(address common.Address, backend bind.ContractBackend) (*ChugSplashDictator, error) {
	contract, err := bindChugSplashDictator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChugSplashDictator{ChugSplashDictatorCaller: ChugSplashDictatorCaller{contract: contract}, ChugSplashDictatorTransactor: ChugSplashDictatorTransactor{contract: contract}, ChugSplashDictatorFilterer: ChugSplashDictatorFilterer{contract: contract}}, nil
}

// NewChugSplashDictatorCaller creates a new read-only instance of ChugSplashDictator, bound to a specific deployed contract.
func NewChugSplashDictatorCaller(address common.Address, caller bind.ContractCaller) (*ChugSplashDictatorCaller, error) {
	contract, err := bindChugSplashDictator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChugSplashDictatorCaller{contract: contract}, nil
}

// NewChugSplashDictatorTransactor creates a new write-only instance of ChugSplashDictator, bound to a specific deployed contract.
func NewChugSplashDictatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ChugSplashDictatorTransactor, error) {
	contract, err := bindChugSplashDictator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChugSplashDictatorTransactor{contract: contract}, nil
}

// NewChugSplashDictatorFilterer creates a new log filterer instance of ChugSplashDictator, bound to a specific deployed contract.
func NewChugSplashDictatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ChugSplashDictatorFilterer, error) {
	contract, err := bindChugSplashDictator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChugSplashDictatorFilterer{contract: contract}, nil
}

// bindChugSplashDictator binds a generic wrapper to an already deployed contract.
func bindChugSplashDictator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChugSplashDictatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChugSplashDictator *ChugSplashDictatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChugSplashDictator.Contract.ChugSplashDictatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChugSplashDictator *ChugSplashDictatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChugSplashDictator.Contract.ChugSplashDictatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChugSplashDictator *ChugSplashDictatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChugSplashDictator.Contract.ChugSplashDictatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChugSplashDictator *ChugSplashDictatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChugSplashDictator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChugSplashDictator *ChugSplashDictatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChugSplashDictator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChugSplashDictator *ChugSplashDictatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChugSplashDictator.Contract.contract.Transact(opts, method, params...)
}

// BitAddressSlotKey is a free data retrieval call binding the contract method 0x830c9df8.
//
// Solidity: function bitAddressSlotKey() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCaller) BitAddressSlotKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChugSplashDictator.contract.Call(opts, &out, "bitAddressSlotKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BitAddressSlotKey is a free data retrieval call binding the contract method 0x830c9df8.
//
// Solidity: function bitAddressSlotKey() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorSession) BitAddressSlotKey() ([32]byte, error) {
	return _ChugSplashDictator.Contract.BitAddressSlotKey(&_ChugSplashDictator.CallOpts)
}

// BitAddressSlotKey is a free data retrieval call binding the contract method 0x830c9df8.
//
// Solidity: function bitAddressSlotKey() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCallerSession) BitAddressSlotKey() ([32]byte, error) {
	return _ChugSplashDictator.Contract.BitAddressSlotKey(&_ChugSplashDictator.CallOpts)
}

// BitAddressSlotVal is a free data retrieval call binding the contract method 0x61aaa41b.
//
// Solidity: function bitAddressSlotVal() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCaller) BitAddressSlotVal(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChugSplashDictator.contract.Call(opts, &out, "bitAddressSlotVal")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BitAddressSlotVal is a free data retrieval call binding the contract method 0x61aaa41b.
//
// Solidity: function bitAddressSlotVal() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorSession) BitAddressSlotVal() ([32]byte, error) {
	return _ChugSplashDictator.Contract.BitAddressSlotVal(&_ChugSplashDictator.CallOpts)
}

// BitAddressSlotVal is a free data retrieval call binding the contract method 0x61aaa41b.
//
// Solidity: function bitAddressSlotVal() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCallerSession) BitAddressSlotVal() ([32]byte, error) {
	return _ChugSplashDictator.Contract.BitAddressSlotVal(&_ChugSplashDictator.CallOpts)
}

// BridgeSlotKey is a free data retrieval call binding the contract method 0xa3b2d8a5.
//
// Solidity: function bridgeSlotKey() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCaller) BridgeSlotKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChugSplashDictator.contract.Call(opts, &out, "bridgeSlotKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BridgeSlotKey is a free data retrieval call binding the contract method 0xa3b2d8a5.
//
// Solidity: function bridgeSlotKey() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorSession) BridgeSlotKey() ([32]byte, error) {
	return _ChugSplashDictator.Contract.BridgeSlotKey(&_ChugSplashDictator.CallOpts)
}

// BridgeSlotKey is a free data retrieval call binding the contract method 0xa3b2d8a5.
//
// Solidity: function bridgeSlotKey() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCallerSession) BridgeSlotKey() ([32]byte, error) {
	return _ChugSplashDictator.Contract.BridgeSlotKey(&_ChugSplashDictator.CallOpts)
}

// BridgeSlotVal is a free data retrieval call binding the contract method 0x907023dd.
//
// Solidity: function bridgeSlotVal() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCaller) BridgeSlotVal(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChugSplashDictator.contract.Call(opts, &out, "bridgeSlotVal")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BridgeSlotVal is a free data retrieval call binding the contract method 0x907023dd.
//
// Solidity: function bridgeSlotVal() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorSession) BridgeSlotVal() ([32]byte, error) {
	return _ChugSplashDictator.Contract.BridgeSlotVal(&_ChugSplashDictator.CallOpts)
}

// BridgeSlotVal is a free data retrieval call binding the contract method 0x907023dd.
//
// Solidity: function bridgeSlotVal() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCallerSession) BridgeSlotVal() ([32]byte, error) {
	return _ChugSplashDictator.Contract.BridgeSlotVal(&_ChugSplashDictator.CallOpts)
}

// CodeHash is a free data retrieval call binding the contract method 0x18edaaf2.
//
// Solidity: function codeHash() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCaller) CodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChugSplashDictator.contract.Call(opts, &out, "codeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CodeHash is a free data retrieval call binding the contract method 0x18edaaf2.
//
// Solidity: function codeHash() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorSession) CodeHash() ([32]byte, error) {
	return _ChugSplashDictator.Contract.CodeHash(&_ChugSplashDictator.CallOpts)
}

// CodeHash is a free data retrieval call binding the contract method 0x18edaaf2.
//
// Solidity: function codeHash() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCallerSession) CodeHash() ([32]byte, error) {
	return _ChugSplashDictator.Contract.CodeHash(&_ChugSplashDictator.CallOpts)
}

// FinalOwner is a free data retrieval call binding the contract method 0x17ad94ec.
//
// Solidity: function finalOwner() view returns(address)
func (_ChugSplashDictator *ChugSplashDictatorCaller) FinalOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChugSplashDictator.contract.Call(opts, &out, "finalOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FinalOwner is a free data retrieval call binding the contract method 0x17ad94ec.
//
// Solidity: function finalOwner() view returns(address)
func (_ChugSplashDictator *ChugSplashDictatorSession) FinalOwner() (common.Address, error) {
	return _ChugSplashDictator.Contract.FinalOwner(&_ChugSplashDictator.CallOpts)
}

// FinalOwner is a free data retrieval call binding the contract method 0x17ad94ec.
//
// Solidity: function finalOwner() view returns(address)
func (_ChugSplashDictator *ChugSplashDictatorCallerSession) FinalOwner() (common.Address, error) {
	return _ChugSplashDictator.Contract.FinalOwner(&_ChugSplashDictator.CallOpts)
}

// IsUpgrading is a free data retrieval call binding the contract method 0xb7947262.
//
// Solidity: function isUpgrading() view returns(bool)
func (_ChugSplashDictator *ChugSplashDictatorCaller) IsUpgrading(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ChugSplashDictator.contract.Call(opts, &out, "isUpgrading")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpgrading is a free data retrieval call binding the contract method 0xb7947262.
//
// Solidity: function isUpgrading() view returns(bool)
func (_ChugSplashDictator *ChugSplashDictatorSession) IsUpgrading() (bool, error) {
	return _ChugSplashDictator.Contract.IsUpgrading(&_ChugSplashDictator.CallOpts)
}

// IsUpgrading is a free data retrieval call binding the contract method 0xb7947262.
//
// Solidity: function isUpgrading() view returns(bool)
func (_ChugSplashDictator *ChugSplashDictatorCallerSession) IsUpgrading() (bool, error) {
	return _ChugSplashDictator.Contract.IsUpgrading(&_ChugSplashDictator.CallOpts)
}

// MessengerSlotKey is a free data retrieval call binding the contract method 0x708518de.
//
// Solidity: function messengerSlotKey() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCaller) MessengerSlotKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChugSplashDictator.contract.Call(opts, &out, "messengerSlotKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessengerSlotKey is a free data retrieval call binding the contract method 0x708518de.
//
// Solidity: function messengerSlotKey() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorSession) MessengerSlotKey() ([32]byte, error) {
	return _ChugSplashDictator.Contract.MessengerSlotKey(&_ChugSplashDictator.CallOpts)
}

// MessengerSlotKey is a free data retrieval call binding the contract method 0x708518de.
//
// Solidity: function messengerSlotKey() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCallerSession) MessengerSlotKey() ([32]byte, error) {
	return _ChugSplashDictator.Contract.MessengerSlotKey(&_ChugSplashDictator.CallOpts)
}

// MessengerSlotVal is a free data retrieval call binding the contract method 0x5307023b.
//
// Solidity: function messengerSlotVal() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCaller) MessengerSlotVal(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChugSplashDictator.contract.Call(opts, &out, "messengerSlotVal")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessengerSlotVal is a free data retrieval call binding the contract method 0x5307023b.
//
// Solidity: function messengerSlotVal() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorSession) MessengerSlotVal() ([32]byte, error) {
	return _ChugSplashDictator.Contract.MessengerSlotVal(&_ChugSplashDictator.CallOpts)
}

// MessengerSlotVal is a free data retrieval call binding the contract method 0x5307023b.
//
// Solidity: function messengerSlotVal() view returns(bytes32)
func (_ChugSplashDictator *ChugSplashDictatorCallerSession) MessengerSlotVal() ([32]byte, error) {
	return _ChugSplashDictator.Contract.MessengerSlotVal(&_ChugSplashDictator.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_ChugSplashDictator *ChugSplashDictatorCaller) Target(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChugSplashDictator.contract.Call(opts, &out, "target")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_ChugSplashDictator *ChugSplashDictatorSession) Target() (common.Address, error) {
	return _ChugSplashDictator.Contract.Target(&_ChugSplashDictator.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_ChugSplashDictator *ChugSplashDictatorCallerSession) Target() (common.Address, error) {
	return _ChugSplashDictator.Contract.Target(&_ChugSplashDictator.CallOpts)
}

// DoActions is a paid mutator transaction binding the contract method 0x0bf56f21.
//
// Solidity: function doActions(bytes _code) returns()
func (_ChugSplashDictator *ChugSplashDictatorTransactor) DoActions(opts *bind.TransactOpts, _code []byte) (*types.Transaction, error) {
	return _ChugSplashDictator.contract.Transact(opts, "doActions", _code)
}

// DoActions is a paid mutator transaction binding the contract method 0x0bf56f21.
//
// Solidity: function doActions(bytes _code) returns()
func (_ChugSplashDictator *ChugSplashDictatorSession) DoActions(_code []byte) (*types.Transaction, error) {
	return _ChugSplashDictator.Contract.DoActions(&_ChugSplashDictator.TransactOpts, _code)
}

// DoActions is a paid mutator transaction binding the contract method 0x0bf56f21.
//
// Solidity: function doActions(bytes _code) returns()
func (_ChugSplashDictator *ChugSplashDictatorTransactorSession) DoActions(_code []byte) (*types.Transaction, error) {
	return _ChugSplashDictator.Contract.DoActions(&_ChugSplashDictator.TransactOpts, _code)
}

// ReturnOwnership is a paid mutator transaction binding the contract method 0x297d1a34.
//
// Solidity: function returnOwnership() returns()
func (_ChugSplashDictator *ChugSplashDictatorTransactor) ReturnOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChugSplashDictator.contract.Transact(opts, "returnOwnership")
}

// ReturnOwnership is a paid mutator transaction binding the contract method 0x297d1a34.
//
// Solidity: function returnOwnership() returns()
func (_ChugSplashDictator *ChugSplashDictatorSession) ReturnOwnership() (*types.Transaction, error) {
	return _ChugSplashDictator.Contract.ReturnOwnership(&_ChugSplashDictator.TransactOpts)
}

// ReturnOwnership is a paid mutator transaction binding the contract method 0x297d1a34.
//
// Solidity: function returnOwnership() returns()
func (_ChugSplashDictator *ChugSplashDictatorTransactorSession) ReturnOwnership() (*types.Transaction, error) {
	return _ChugSplashDictator.Contract.ReturnOwnership(&_ChugSplashDictator.TransactOpts)
}

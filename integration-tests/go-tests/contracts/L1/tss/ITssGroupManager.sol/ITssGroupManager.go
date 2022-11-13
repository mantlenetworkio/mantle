// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ITssGroupManager

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

// ITssGroupManagerTssMember is an auto generated low-level Go binding around an user-defined struct.
type ITssGroupManagerTssMember struct {
	PublicKey   []byte
	NodeAddress common.Address
	Status      uint8
}

// ITssGroupManagerMetaData contains all meta data concerning the ITssGroupManager contract.
var ITssGroupManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getTssGroupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTssGroupMembers\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTssGroupUnJailMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTssInactiveGroupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"getTssMember\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"enumITssGroupManager.MemberStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structITssGroupManager.TssMember\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inActiveIsEmpty\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"memberExistActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"memberExistInActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"memberJail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"memberUnJail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"publicKeyToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"removeMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_groupPublicKey\",\"type\":\"bytes\"}],\"name\":\"setGroupPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_batchPublicKey\",\"type\":\"bytes[]\"}],\"name\":\"setTssGroupMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"verifySign\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITssGroupManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ITssGroupManagerMetaData.ABI instead.
var ITssGroupManagerABI = ITssGroupManagerMetaData.ABI

// ITssGroupManager is an auto generated Go binding around an Ethereum contract.
type ITssGroupManager struct {
	ITssGroupManagerCaller     // Read-only binding to the contract
	ITssGroupManagerTransactor // Write-only binding to the contract
	ITssGroupManagerFilterer   // Log filterer for contract events
}

// ITssGroupManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITssGroupManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITssGroupManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITssGroupManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITssGroupManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITssGroupManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITssGroupManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITssGroupManagerSession struct {
	Contract     *ITssGroupManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITssGroupManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITssGroupManagerCallerSession struct {
	Contract *ITssGroupManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ITssGroupManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITssGroupManagerTransactorSession struct {
	Contract     *ITssGroupManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ITssGroupManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITssGroupManagerRaw struct {
	Contract *ITssGroupManager // Generic contract binding to access the raw methods on
}

// ITssGroupManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITssGroupManagerCallerRaw struct {
	Contract *ITssGroupManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ITssGroupManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITssGroupManagerTransactorRaw struct {
	Contract *ITssGroupManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITssGroupManager creates a new instance of ITssGroupManager, bound to a specific deployed contract.
func NewITssGroupManager(address common.Address, backend bind.ContractBackend) (*ITssGroupManager, error) {
	contract, err := bindITssGroupManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITssGroupManager{ITssGroupManagerCaller: ITssGroupManagerCaller{contract: contract}, ITssGroupManagerTransactor: ITssGroupManagerTransactor{contract: contract}, ITssGroupManagerFilterer: ITssGroupManagerFilterer{contract: contract}}, nil
}

// NewITssGroupManagerCaller creates a new read-only instance of ITssGroupManager, bound to a specific deployed contract.
func NewITssGroupManagerCaller(address common.Address, caller bind.ContractCaller) (*ITssGroupManagerCaller, error) {
	contract, err := bindITssGroupManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITssGroupManagerCaller{contract: contract}, nil
}

// NewITssGroupManagerTransactor creates a new write-only instance of ITssGroupManager, bound to a specific deployed contract.
func NewITssGroupManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ITssGroupManagerTransactor, error) {
	contract, err := bindITssGroupManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITssGroupManagerTransactor{contract: contract}, nil
}

// NewITssGroupManagerFilterer creates a new log filterer instance of ITssGroupManager, bound to a specific deployed contract.
func NewITssGroupManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ITssGroupManagerFilterer, error) {
	contract, err := bindITssGroupManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITssGroupManagerFilterer{contract: contract}, nil
}

// bindITssGroupManager binds a generic wrapper to an already deployed contract.
func bindITssGroupManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITssGroupManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITssGroupManager *ITssGroupManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITssGroupManager.Contract.ITssGroupManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITssGroupManager *ITssGroupManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.ITssGroupManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITssGroupManager *ITssGroupManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.ITssGroupManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITssGroupManager *ITssGroupManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITssGroupManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITssGroupManager *ITssGroupManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITssGroupManager *ITssGroupManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.contract.Transact(opts, method, params...)
}

// GetTssGroupInfo is a paid mutator transaction binding the contract method 0xfcca5592.
//
// Solidity: function getTssGroupInfo() returns(uint256, uint256, bytes, bytes[])
func (_ITssGroupManager *ITssGroupManagerTransactor) GetTssGroupInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "getTssGroupInfo")
}

// GetTssGroupInfo is a paid mutator transaction binding the contract method 0xfcca5592.
//
// Solidity: function getTssGroupInfo() returns(uint256, uint256, bytes, bytes[])
func (_ITssGroupManager *ITssGroupManagerSession) GetTssGroupInfo() (*types.Transaction, error) {
	return _ITssGroupManager.Contract.GetTssGroupInfo(&_ITssGroupManager.TransactOpts)
}

// GetTssGroupInfo is a paid mutator transaction binding the contract method 0xfcca5592.
//
// Solidity: function getTssGroupInfo() returns(uint256, uint256, bytes, bytes[])
func (_ITssGroupManager *ITssGroupManagerTransactorSession) GetTssGroupInfo() (*types.Transaction, error) {
	return _ITssGroupManager.Contract.GetTssGroupInfo(&_ITssGroupManager.TransactOpts)
}

// GetTssGroupMembers is a paid mutator transaction binding the contract method 0x95f645bc.
//
// Solidity: function getTssGroupMembers() returns(bytes[])
func (_ITssGroupManager *ITssGroupManagerTransactor) GetTssGroupMembers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "getTssGroupMembers")
}

// GetTssGroupMembers is a paid mutator transaction binding the contract method 0x95f645bc.
//
// Solidity: function getTssGroupMembers() returns(bytes[])
func (_ITssGroupManager *ITssGroupManagerSession) GetTssGroupMembers() (*types.Transaction, error) {
	return _ITssGroupManager.Contract.GetTssGroupMembers(&_ITssGroupManager.TransactOpts)
}

// GetTssGroupMembers is a paid mutator transaction binding the contract method 0x95f645bc.
//
// Solidity: function getTssGroupMembers() returns(bytes[])
func (_ITssGroupManager *ITssGroupManagerTransactorSession) GetTssGroupMembers() (*types.Transaction, error) {
	return _ITssGroupManager.Contract.GetTssGroupMembers(&_ITssGroupManager.TransactOpts)
}

// GetTssGroupUnJailMembers is a paid mutator transaction binding the contract method 0x2cd00d53.
//
// Solidity: function getTssGroupUnJailMembers() returns(address[])
func (_ITssGroupManager *ITssGroupManagerTransactor) GetTssGroupUnJailMembers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "getTssGroupUnJailMembers")
}

// GetTssGroupUnJailMembers is a paid mutator transaction binding the contract method 0x2cd00d53.
//
// Solidity: function getTssGroupUnJailMembers() returns(address[])
func (_ITssGroupManager *ITssGroupManagerSession) GetTssGroupUnJailMembers() (*types.Transaction, error) {
	return _ITssGroupManager.Contract.GetTssGroupUnJailMembers(&_ITssGroupManager.TransactOpts)
}

// GetTssGroupUnJailMembers is a paid mutator transaction binding the contract method 0x2cd00d53.
//
// Solidity: function getTssGroupUnJailMembers() returns(address[])
func (_ITssGroupManager *ITssGroupManagerTransactorSession) GetTssGroupUnJailMembers() (*types.Transaction, error) {
	return _ITssGroupManager.Contract.GetTssGroupUnJailMembers(&_ITssGroupManager.TransactOpts)
}

// GetTssInactiveGroupInfo is a paid mutator transaction binding the contract method 0x404a6986.
//
// Solidity: function getTssInactiveGroupInfo() returns(uint256, uint256, bytes[])
func (_ITssGroupManager *ITssGroupManagerTransactor) GetTssInactiveGroupInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "getTssInactiveGroupInfo")
}

// GetTssInactiveGroupInfo is a paid mutator transaction binding the contract method 0x404a6986.
//
// Solidity: function getTssInactiveGroupInfo() returns(uint256, uint256, bytes[])
func (_ITssGroupManager *ITssGroupManagerSession) GetTssInactiveGroupInfo() (*types.Transaction, error) {
	return _ITssGroupManager.Contract.GetTssInactiveGroupInfo(&_ITssGroupManager.TransactOpts)
}

// GetTssInactiveGroupInfo is a paid mutator transaction binding the contract method 0x404a6986.
//
// Solidity: function getTssInactiveGroupInfo() returns(uint256, uint256, bytes[])
func (_ITssGroupManager *ITssGroupManagerTransactorSession) GetTssInactiveGroupInfo() (*types.Transaction, error) {
	return _ITssGroupManager.Contract.GetTssInactiveGroupInfo(&_ITssGroupManager.TransactOpts)
}

// GetTssMember is a paid mutator transaction binding the contract method 0x54bbb914.
//
// Solidity: function getTssMember(bytes _publicKey) returns((bytes,address,uint8))
func (_ITssGroupManager *ITssGroupManagerTransactor) GetTssMember(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "getTssMember", _publicKey)
}

// GetTssMember is a paid mutator transaction binding the contract method 0x54bbb914.
//
// Solidity: function getTssMember(bytes _publicKey) returns((bytes,address,uint8))
func (_ITssGroupManager *ITssGroupManagerSession) GetTssMember(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.GetTssMember(&_ITssGroupManager.TransactOpts, _publicKey)
}

// GetTssMember is a paid mutator transaction binding the contract method 0x54bbb914.
//
// Solidity: function getTssMember(bytes _publicKey) returns((bytes,address,uint8))
func (_ITssGroupManager *ITssGroupManagerTransactorSession) GetTssMember(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.GetTssMember(&_ITssGroupManager.TransactOpts, _publicKey)
}

// InActiveIsEmpty is a paid mutator transaction binding the contract method 0x8a51d8b4.
//
// Solidity: function inActiveIsEmpty() returns(bool)
func (_ITssGroupManager *ITssGroupManagerTransactor) InActiveIsEmpty(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "inActiveIsEmpty")
}

// InActiveIsEmpty is a paid mutator transaction binding the contract method 0x8a51d8b4.
//
// Solidity: function inActiveIsEmpty() returns(bool)
func (_ITssGroupManager *ITssGroupManagerSession) InActiveIsEmpty() (*types.Transaction, error) {
	return _ITssGroupManager.Contract.InActiveIsEmpty(&_ITssGroupManager.TransactOpts)
}

// InActiveIsEmpty is a paid mutator transaction binding the contract method 0x8a51d8b4.
//
// Solidity: function inActiveIsEmpty() returns(bool)
func (_ITssGroupManager *ITssGroupManagerTransactorSession) InActiveIsEmpty() (*types.Transaction, error) {
	return _ITssGroupManager.Contract.InActiveIsEmpty(&_ITssGroupManager.TransactOpts)
}

// MemberExistActive is a paid mutator transaction binding the contract method 0xbcbc5fc6.
//
// Solidity: function memberExistActive(bytes _publicKey) returns(bool)
func (_ITssGroupManager *ITssGroupManagerTransactor) MemberExistActive(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "memberExistActive", _publicKey)
}

// MemberExistActive is a paid mutator transaction binding the contract method 0xbcbc5fc6.
//
// Solidity: function memberExistActive(bytes _publicKey) returns(bool)
func (_ITssGroupManager *ITssGroupManagerSession) MemberExistActive(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.MemberExistActive(&_ITssGroupManager.TransactOpts, _publicKey)
}

// MemberExistActive is a paid mutator transaction binding the contract method 0xbcbc5fc6.
//
// Solidity: function memberExistActive(bytes _publicKey) returns(bool)
func (_ITssGroupManager *ITssGroupManagerTransactorSession) MemberExistActive(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.MemberExistActive(&_ITssGroupManager.TransactOpts, _publicKey)
}

// MemberExistInActive is a paid mutator transaction binding the contract method 0x4866e2e0.
//
// Solidity: function memberExistInActive(bytes _publicKey) returns(bool)
func (_ITssGroupManager *ITssGroupManagerTransactor) MemberExistInActive(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "memberExistInActive", _publicKey)
}

// MemberExistInActive is a paid mutator transaction binding the contract method 0x4866e2e0.
//
// Solidity: function memberExistInActive(bytes _publicKey) returns(bool)
func (_ITssGroupManager *ITssGroupManagerSession) MemberExistInActive(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.MemberExistInActive(&_ITssGroupManager.TransactOpts, _publicKey)
}

// MemberExistInActive is a paid mutator transaction binding the contract method 0x4866e2e0.
//
// Solidity: function memberExistInActive(bytes _publicKey) returns(bool)
func (_ITssGroupManager *ITssGroupManagerTransactorSession) MemberExistInActive(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.MemberExistInActive(&_ITssGroupManager.TransactOpts, _publicKey)
}

// MemberJail is a paid mutator transaction binding the contract method 0x254ff981.
//
// Solidity: function memberJail(bytes _publicKey) returns()
func (_ITssGroupManager *ITssGroupManagerTransactor) MemberJail(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "memberJail", _publicKey)
}

// MemberJail is a paid mutator transaction binding the contract method 0x254ff981.
//
// Solidity: function memberJail(bytes _publicKey) returns()
func (_ITssGroupManager *ITssGroupManagerSession) MemberJail(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.MemberJail(&_ITssGroupManager.TransactOpts, _publicKey)
}

// MemberJail is a paid mutator transaction binding the contract method 0x254ff981.
//
// Solidity: function memberJail(bytes _publicKey) returns()
func (_ITssGroupManager *ITssGroupManagerTransactorSession) MemberJail(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.MemberJail(&_ITssGroupManager.TransactOpts, _publicKey)
}

// MemberUnJail is a paid mutator transaction binding the contract method 0x429ec49c.
//
// Solidity: function memberUnJail(bytes _publicKey) returns()
func (_ITssGroupManager *ITssGroupManagerTransactor) MemberUnJail(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "memberUnJail", _publicKey)
}

// MemberUnJail is a paid mutator transaction binding the contract method 0x429ec49c.
//
// Solidity: function memberUnJail(bytes _publicKey) returns()
func (_ITssGroupManager *ITssGroupManagerSession) MemberUnJail(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.MemberUnJail(&_ITssGroupManager.TransactOpts, _publicKey)
}

// MemberUnJail is a paid mutator transaction binding the contract method 0x429ec49c.
//
// Solidity: function memberUnJail(bytes _publicKey) returns()
func (_ITssGroupManager *ITssGroupManagerTransactorSession) MemberUnJail(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.MemberUnJail(&_ITssGroupManager.TransactOpts, _publicKey)
}

// PublicKeyToAddress is a paid mutator transaction binding the contract method 0x43ae656c.
//
// Solidity: function publicKeyToAddress(bytes publicKey) returns(address)
func (_ITssGroupManager *ITssGroupManagerTransactor) PublicKeyToAddress(opts *bind.TransactOpts, publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "publicKeyToAddress", publicKey)
}

// PublicKeyToAddress is a paid mutator transaction binding the contract method 0x43ae656c.
//
// Solidity: function publicKeyToAddress(bytes publicKey) returns(address)
func (_ITssGroupManager *ITssGroupManagerSession) PublicKeyToAddress(publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.PublicKeyToAddress(&_ITssGroupManager.TransactOpts, publicKey)
}

// PublicKeyToAddress is a paid mutator transaction binding the contract method 0x43ae656c.
//
// Solidity: function publicKeyToAddress(bytes publicKey) returns(address)
func (_ITssGroupManager *ITssGroupManagerTransactorSession) PublicKeyToAddress(publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.PublicKeyToAddress(&_ITssGroupManager.TransactOpts, publicKey)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x7a952c11.
//
// Solidity: function removeMember(bytes _publicKey) returns()
func (_ITssGroupManager *ITssGroupManagerTransactor) RemoveMember(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "removeMember", _publicKey)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x7a952c11.
//
// Solidity: function removeMember(bytes _publicKey) returns()
func (_ITssGroupManager *ITssGroupManagerSession) RemoveMember(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.RemoveMember(&_ITssGroupManager.TransactOpts, _publicKey)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x7a952c11.
//
// Solidity: function removeMember(bytes _publicKey) returns()
func (_ITssGroupManager *ITssGroupManagerTransactorSession) RemoveMember(_publicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.RemoveMember(&_ITssGroupManager.TransactOpts, _publicKey)
}

// SetGroupPublicKey is a paid mutator transaction binding the contract method 0x4679b366.
//
// Solidity: function setGroupPublicKey(bytes _publicKey, bytes _groupPublicKey) returns()
func (_ITssGroupManager *ITssGroupManagerTransactor) SetGroupPublicKey(opts *bind.TransactOpts, _publicKey []byte, _groupPublicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "setGroupPublicKey", _publicKey, _groupPublicKey)
}

// SetGroupPublicKey is a paid mutator transaction binding the contract method 0x4679b366.
//
// Solidity: function setGroupPublicKey(bytes _publicKey, bytes _groupPublicKey) returns()
func (_ITssGroupManager *ITssGroupManagerSession) SetGroupPublicKey(_publicKey []byte, _groupPublicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.SetGroupPublicKey(&_ITssGroupManager.TransactOpts, _publicKey, _groupPublicKey)
}

// SetGroupPublicKey is a paid mutator transaction binding the contract method 0x4679b366.
//
// Solidity: function setGroupPublicKey(bytes _publicKey, bytes _groupPublicKey) returns()
func (_ITssGroupManager *ITssGroupManagerTransactorSession) SetGroupPublicKey(_publicKey []byte, _groupPublicKey []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.SetGroupPublicKey(&_ITssGroupManager.TransactOpts, _publicKey, _groupPublicKey)
}

// SetTssGroupMember is a paid mutator transaction binding the contract method 0x1a47931b.
//
// Solidity: function setTssGroupMember(uint256 _threshold, bytes[] _batchPublicKey) returns()
func (_ITssGroupManager *ITssGroupManagerTransactor) SetTssGroupMember(opts *bind.TransactOpts, _threshold *big.Int, _batchPublicKey [][]byte) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "setTssGroupMember", _threshold, _batchPublicKey)
}

// SetTssGroupMember is a paid mutator transaction binding the contract method 0x1a47931b.
//
// Solidity: function setTssGroupMember(uint256 _threshold, bytes[] _batchPublicKey) returns()
func (_ITssGroupManager *ITssGroupManagerSession) SetTssGroupMember(_threshold *big.Int, _batchPublicKey [][]byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.SetTssGroupMember(&_ITssGroupManager.TransactOpts, _threshold, _batchPublicKey)
}

// SetTssGroupMember is a paid mutator transaction binding the contract method 0x1a47931b.
//
// Solidity: function setTssGroupMember(uint256 _threshold, bytes[] _batchPublicKey) returns()
func (_ITssGroupManager *ITssGroupManagerTransactorSession) SetTssGroupMember(_threshold *big.Int, _batchPublicKey [][]byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.SetTssGroupMember(&_ITssGroupManager.TransactOpts, _threshold, _batchPublicKey)
}

// VerifySign is a paid mutator transaction binding the contract method 0x3231a7f0.
//
// Solidity: function verifySign(bytes32 _message, bytes _sig) returns(bool)
func (_ITssGroupManager *ITssGroupManagerTransactor) VerifySign(opts *bind.TransactOpts, _message [32]byte, _sig []byte) (*types.Transaction, error) {
	return _ITssGroupManager.contract.Transact(opts, "verifySign", _message, _sig)
}

// VerifySign is a paid mutator transaction binding the contract method 0x3231a7f0.
//
// Solidity: function verifySign(bytes32 _message, bytes _sig) returns(bool)
func (_ITssGroupManager *ITssGroupManagerSession) VerifySign(_message [32]byte, _sig []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.VerifySign(&_ITssGroupManager.TransactOpts, _message, _sig)
}

// VerifySign is a paid mutator transaction binding the contract method 0x3231a7f0.
//
// Solidity: function verifySign(bytes32 _message, bytes _sig) returns(bool)
func (_ITssGroupManager *ITssGroupManagerTransactorSession) VerifySign(_message [32]byte, _sig []byte) (*types.Transaction, error) {
	return _ITssGroupManager.Contract.VerifySign(&_ITssGroupManager.TransactOpts, _message, _sig)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IChainStorageContainer

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

// IChainStorageContainerMetaData contains all meta data concerning the IChainStorageContainer contract.
var IChainStorageContainerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes27\",\"name\":\"_globalMetadata\",\"type\":\"bytes27\"}],\"name\":\"deleteElementsAfterInclusive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"deleteElementsAfterInclusive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalMetadata\",\"outputs\":[{\"internalType\":\"bytes27\",\"name\":\"\",\"type\":\"bytes27\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_object\",\"type\":\"bytes32\"},{\"internalType\":\"bytes27\",\"name\":\"_globalMetadata\",\"type\":\"bytes27\"}],\"name\":\"push\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_object\",\"type\":\"bytes32\"}],\"name\":\"push\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes27\",\"name\":\"_globalMetadata\",\"type\":\"bytes27\"}],\"name\":\"setGlobalMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IChainStorageContainerABI is the input ABI used to generate the binding from.
// Deprecated: Use IChainStorageContainerMetaData.ABI instead.
var IChainStorageContainerABI = IChainStorageContainerMetaData.ABI

// IChainStorageContainer is an auto generated Go binding around an Ethereum contract.
type IChainStorageContainer struct {
	IChainStorageContainerCaller     // Read-only binding to the contract
	IChainStorageContainerTransactor // Write-only binding to the contract
	IChainStorageContainerFilterer   // Log filterer for contract events
}

// IChainStorageContainerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChainStorageContainerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChainStorageContainerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChainStorageContainerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChainStorageContainerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChainStorageContainerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChainStorageContainerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChainStorageContainerSession struct {
	Contract     *IChainStorageContainer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IChainStorageContainerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChainStorageContainerCallerSession struct {
	Contract *IChainStorageContainerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IChainStorageContainerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChainStorageContainerTransactorSession struct {
	Contract     *IChainStorageContainerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IChainStorageContainerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChainStorageContainerRaw struct {
	Contract *IChainStorageContainer // Generic contract binding to access the raw methods on
}

// IChainStorageContainerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChainStorageContainerCallerRaw struct {
	Contract *IChainStorageContainerCaller // Generic read-only contract binding to access the raw methods on
}

// IChainStorageContainerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChainStorageContainerTransactorRaw struct {
	Contract *IChainStorageContainerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChainStorageContainer creates a new instance of IChainStorageContainer, bound to a specific deployed contract.
func NewIChainStorageContainer(address common.Address, backend bind.ContractBackend) (*IChainStorageContainer, error) {
	contract, err := bindIChainStorageContainer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChainStorageContainer{IChainStorageContainerCaller: IChainStorageContainerCaller{contract: contract}, IChainStorageContainerTransactor: IChainStorageContainerTransactor{contract: contract}, IChainStorageContainerFilterer: IChainStorageContainerFilterer{contract: contract}}, nil
}

// NewIChainStorageContainerCaller creates a new read-only instance of IChainStorageContainer, bound to a specific deployed contract.
func NewIChainStorageContainerCaller(address common.Address, caller bind.ContractCaller) (*IChainStorageContainerCaller, error) {
	contract, err := bindIChainStorageContainer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChainStorageContainerCaller{contract: contract}, nil
}

// NewIChainStorageContainerTransactor creates a new write-only instance of IChainStorageContainer, bound to a specific deployed contract.
func NewIChainStorageContainerTransactor(address common.Address, transactor bind.ContractTransactor) (*IChainStorageContainerTransactor, error) {
	contract, err := bindIChainStorageContainer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChainStorageContainerTransactor{contract: contract}, nil
}

// NewIChainStorageContainerFilterer creates a new log filterer instance of IChainStorageContainer, bound to a specific deployed contract.
func NewIChainStorageContainerFilterer(address common.Address, filterer bind.ContractFilterer) (*IChainStorageContainerFilterer, error) {
	contract, err := bindIChainStorageContainer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChainStorageContainerFilterer{contract: contract}, nil
}

// bindIChainStorageContainer binds a generic wrapper to an already deployed contract.
func bindIChainStorageContainer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChainStorageContainerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChainStorageContainer *IChainStorageContainerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChainStorageContainer.Contract.IChainStorageContainerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChainStorageContainer *IChainStorageContainerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.IChainStorageContainerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChainStorageContainer *IChainStorageContainerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.IChainStorageContainerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChainStorageContainer *IChainStorageContainerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChainStorageContainer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChainStorageContainer *IChainStorageContainerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChainStorageContainer *IChainStorageContainerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _index) view returns(bytes32)
func (_IChainStorageContainer *IChainStorageContainerCaller) Get(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IChainStorageContainer.contract.Call(opts, &out, "get", _index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _index) view returns(bytes32)
func (_IChainStorageContainer *IChainStorageContainerSession) Get(_index *big.Int) ([32]byte, error) {
	return _IChainStorageContainer.Contract.Get(&_IChainStorageContainer.CallOpts, _index)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _index) view returns(bytes32)
func (_IChainStorageContainer *IChainStorageContainerCallerSession) Get(_index *big.Int) ([32]byte, error) {
	return _IChainStorageContainer.Contract.Get(&_IChainStorageContainer.CallOpts, _index)
}

// GetGlobalMetadata is a free data retrieval call binding the contract method 0xccf8f969.
//
// Solidity: function getGlobalMetadata() view returns(bytes27)
func (_IChainStorageContainer *IChainStorageContainerCaller) GetGlobalMetadata(opts *bind.CallOpts) ([27]byte, error) {
	var out []interface{}
	err := _IChainStorageContainer.contract.Call(opts, &out, "getGlobalMetadata")

	if err != nil {
		return *new([27]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([27]byte)).(*[27]byte)

	return out0, err

}

// GetGlobalMetadata is a free data retrieval call binding the contract method 0xccf8f969.
//
// Solidity: function getGlobalMetadata() view returns(bytes27)
func (_IChainStorageContainer *IChainStorageContainerSession) GetGlobalMetadata() ([27]byte, error) {
	return _IChainStorageContainer.Contract.GetGlobalMetadata(&_IChainStorageContainer.CallOpts)
}

// GetGlobalMetadata is a free data retrieval call binding the contract method 0xccf8f969.
//
// Solidity: function getGlobalMetadata() view returns(bytes27)
func (_IChainStorageContainer *IChainStorageContainerCallerSession) GetGlobalMetadata() ([27]byte, error) {
	return _IChainStorageContainer.Contract.GetGlobalMetadata(&_IChainStorageContainer.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_IChainStorageContainer *IChainStorageContainerCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IChainStorageContainer.contract.Call(opts, &out, "length")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_IChainStorageContainer *IChainStorageContainerSession) Length() (*big.Int, error) {
	return _IChainStorageContainer.Contract.Length(&_IChainStorageContainer.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_IChainStorageContainer *IChainStorageContainerCallerSession) Length() (*big.Int, error) {
	return _IChainStorageContainer.Contract.Length(&_IChainStorageContainer.CallOpts)
}

// DeleteElementsAfterInclusive is a paid mutator transaction binding the contract method 0x167fd681.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index, bytes27 _globalMetadata) returns()
func (_IChainStorageContainer *IChainStorageContainerTransactor) DeleteElementsAfterInclusive(opts *bind.TransactOpts, _index *big.Int, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.contract.Transact(opts, "deleteElementsAfterInclusive", _index, _globalMetadata)
}

// DeleteElementsAfterInclusive is a paid mutator transaction binding the contract method 0x167fd681.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index, bytes27 _globalMetadata) returns()
func (_IChainStorageContainer *IChainStorageContainerSession) DeleteElementsAfterInclusive(_index *big.Int, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.DeleteElementsAfterInclusive(&_IChainStorageContainer.TransactOpts, _index, _globalMetadata)
}

// DeleteElementsAfterInclusive is a paid mutator transaction binding the contract method 0x167fd681.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index, bytes27 _globalMetadata) returns()
func (_IChainStorageContainer *IChainStorageContainerTransactorSession) DeleteElementsAfterInclusive(_index *big.Int, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.DeleteElementsAfterInclusive(&_IChainStorageContainer.TransactOpts, _index, _globalMetadata)
}

// DeleteElementsAfterInclusive0 is a paid mutator transaction binding the contract method 0x4651d91e.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index) returns()
func (_IChainStorageContainer *IChainStorageContainerTransactor) DeleteElementsAfterInclusive0(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _IChainStorageContainer.contract.Transact(opts, "deleteElementsAfterInclusive0", _index)
}

// DeleteElementsAfterInclusive0 is a paid mutator transaction binding the contract method 0x4651d91e.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index) returns()
func (_IChainStorageContainer *IChainStorageContainerSession) DeleteElementsAfterInclusive0(_index *big.Int) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.DeleteElementsAfterInclusive0(&_IChainStorageContainer.TransactOpts, _index)
}

// DeleteElementsAfterInclusive0 is a paid mutator transaction binding the contract method 0x4651d91e.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index) returns()
func (_IChainStorageContainer *IChainStorageContainerTransactorSession) DeleteElementsAfterInclusive0(_index *big.Int) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.DeleteElementsAfterInclusive0(&_IChainStorageContainer.TransactOpts, _index)
}

// Push is a paid mutator transaction binding the contract method 0x2015276c.
//
// Solidity: function push(bytes32 _object, bytes27 _globalMetadata) returns()
func (_IChainStorageContainer *IChainStorageContainerTransactor) Push(opts *bind.TransactOpts, _object [32]byte, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.contract.Transact(opts, "push", _object, _globalMetadata)
}

// Push is a paid mutator transaction binding the contract method 0x2015276c.
//
// Solidity: function push(bytes32 _object, bytes27 _globalMetadata) returns()
func (_IChainStorageContainer *IChainStorageContainerSession) Push(_object [32]byte, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.Push(&_IChainStorageContainer.TransactOpts, _object, _globalMetadata)
}

// Push is a paid mutator transaction binding the contract method 0x2015276c.
//
// Solidity: function push(bytes32 _object, bytes27 _globalMetadata) returns()
func (_IChainStorageContainer *IChainStorageContainerTransactorSession) Push(_object [32]byte, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.Push(&_IChainStorageContainer.TransactOpts, _object, _globalMetadata)
}

// Push0 is a paid mutator transaction binding the contract method 0xb298e36b.
//
// Solidity: function push(bytes32 _object) returns()
func (_IChainStorageContainer *IChainStorageContainerTransactor) Push0(opts *bind.TransactOpts, _object [32]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.contract.Transact(opts, "push0", _object)
}

// Push0 is a paid mutator transaction binding the contract method 0xb298e36b.
//
// Solidity: function push(bytes32 _object) returns()
func (_IChainStorageContainer *IChainStorageContainerSession) Push0(_object [32]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.Push0(&_IChainStorageContainer.TransactOpts, _object)
}

// Push0 is a paid mutator transaction binding the contract method 0xb298e36b.
//
// Solidity: function push(bytes32 _object) returns()
func (_IChainStorageContainer *IChainStorageContainerTransactorSession) Push0(_object [32]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.Push0(&_IChainStorageContainer.TransactOpts, _object)
}

// SetGlobalMetadata is a paid mutator transaction binding the contract method 0x29061de2.
//
// Solidity: function setGlobalMetadata(bytes27 _globalMetadata) returns()
func (_IChainStorageContainer *IChainStorageContainerTransactor) SetGlobalMetadata(opts *bind.TransactOpts, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.contract.Transact(opts, "setGlobalMetadata", _globalMetadata)
}

// SetGlobalMetadata is a paid mutator transaction binding the contract method 0x29061de2.
//
// Solidity: function setGlobalMetadata(bytes27 _globalMetadata) returns()
func (_IChainStorageContainer *IChainStorageContainerSession) SetGlobalMetadata(_globalMetadata [27]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.SetGlobalMetadata(&_IChainStorageContainer.TransactOpts, _globalMetadata)
}

// SetGlobalMetadata is a paid mutator transaction binding the contract method 0x29061de2.
//
// Solidity: function setGlobalMetadata(bytes27 _globalMetadata) returns()
func (_IChainStorageContainer *IChainStorageContainerTransactorSession) SetGlobalMetadata(_globalMetadata [27]byte) (*types.Transaction, error) {
	return _IChainStorageContainer.Contract.SetGlobalMetadata(&_IChainStorageContainer.TransactOpts, _globalMetadata)
}

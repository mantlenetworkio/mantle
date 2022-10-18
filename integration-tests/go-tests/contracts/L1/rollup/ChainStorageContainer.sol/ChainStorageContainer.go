// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ChainStorageContainer

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

// ChainStorageContainerMetaData contains all meta data concerning the ChainStorageContainer contract.
var ChainStorageContainerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_libAddressManager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_owner\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes27\",\"name\":\"_globalMetadata\",\"type\":\"bytes27\"}],\"name\":\"deleteElementsAfterInclusive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"deleteElementsAfterInclusive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalMetadata\",\"outputs\":[{\"internalType\":\"bytes27\",\"name\":\"\",\"type\":\"bytes27\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libAddressManager\",\"outputs\":[{\"internalType\":\"contractLib_AddressManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_object\",\"type\":\"bytes32\"},{\"internalType\":\"bytes27\",\"name\":\"_globalMetadata\",\"type\":\"bytes27\"}],\"name\":\"push\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_object\",\"type\":\"bytes32\"}],\"name\":\"push\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes27\",\"name\":\"_globalMetadata\",\"type\":\"bytes27\"}],\"name\":\"setGlobalMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ChainStorageContainerABI is the input ABI used to generate the binding from.
// Deprecated: Use ChainStorageContainerMetaData.ABI instead.
var ChainStorageContainerABI = ChainStorageContainerMetaData.ABI

// ChainStorageContainer is an auto generated Go binding around an Ethereum contract.
type ChainStorageContainer struct {
	ChainStorageContainerCaller     // Read-only binding to the contract
	ChainStorageContainerTransactor // Write-only binding to the contract
	ChainStorageContainerFilterer   // Log filterer for contract events
}

// ChainStorageContainerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainStorageContainerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainStorageContainerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainStorageContainerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainStorageContainerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainStorageContainerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainStorageContainerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainStorageContainerSession struct {
	Contract     *ChainStorageContainer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ChainStorageContainerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainStorageContainerCallerSession struct {
	Contract *ChainStorageContainerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ChainStorageContainerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainStorageContainerTransactorSession struct {
	Contract     *ChainStorageContainerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ChainStorageContainerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainStorageContainerRaw struct {
	Contract *ChainStorageContainer // Generic contract binding to access the raw methods on
}

// ChainStorageContainerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainStorageContainerCallerRaw struct {
	Contract *ChainStorageContainerCaller // Generic read-only contract binding to access the raw methods on
}

// ChainStorageContainerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainStorageContainerTransactorRaw struct {
	Contract *ChainStorageContainerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainStorageContainer creates a new instance of ChainStorageContainer, bound to a specific deployed contract.
func NewChainStorageContainer(address common.Address, backend bind.ContractBackend) (*ChainStorageContainer, error) {
	contract, err := bindChainStorageContainer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainStorageContainer{ChainStorageContainerCaller: ChainStorageContainerCaller{contract: contract}, ChainStorageContainerTransactor: ChainStorageContainerTransactor{contract: contract}, ChainStorageContainerFilterer: ChainStorageContainerFilterer{contract: contract}}, nil
}

// NewChainStorageContainerCaller creates a new read-only instance of ChainStorageContainer, bound to a specific deployed contract.
func NewChainStorageContainerCaller(address common.Address, caller bind.ContractCaller) (*ChainStorageContainerCaller, error) {
	contract, err := bindChainStorageContainer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainStorageContainerCaller{contract: contract}, nil
}

// NewChainStorageContainerTransactor creates a new write-only instance of ChainStorageContainer, bound to a specific deployed contract.
func NewChainStorageContainerTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainStorageContainerTransactor, error) {
	contract, err := bindChainStorageContainer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainStorageContainerTransactor{contract: contract}, nil
}

// NewChainStorageContainerFilterer creates a new log filterer instance of ChainStorageContainer, bound to a specific deployed contract.
func NewChainStorageContainerFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainStorageContainerFilterer, error) {
	contract, err := bindChainStorageContainer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainStorageContainerFilterer{contract: contract}, nil
}

// bindChainStorageContainer binds a generic wrapper to an already deployed contract.
func bindChainStorageContainer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainStorageContainerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainStorageContainer *ChainStorageContainerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainStorageContainer.Contract.ChainStorageContainerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainStorageContainer *ChainStorageContainerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.ChainStorageContainerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainStorageContainer *ChainStorageContainerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.ChainStorageContainerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainStorageContainer *ChainStorageContainerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainStorageContainer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainStorageContainer *ChainStorageContainerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainStorageContainer *ChainStorageContainerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _index) view returns(bytes32)
func (_ChainStorageContainer *ChainStorageContainerCaller) Get(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ChainStorageContainer.contract.Call(opts, &out, "get", _index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _index) view returns(bytes32)
func (_ChainStorageContainer *ChainStorageContainerSession) Get(_index *big.Int) ([32]byte, error) {
	return _ChainStorageContainer.Contract.Get(&_ChainStorageContainer.CallOpts, _index)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _index) view returns(bytes32)
func (_ChainStorageContainer *ChainStorageContainerCallerSession) Get(_index *big.Int) ([32]byte, error) {
	return _ChainStorageContainer.Contract.Get(&_ChainStorageContainer.CallOpts, _index)
}

// GetGlobalMetadata is a free data retrieval call binding the contract method 0xccf8f969.
//
// Solidity: function getGlobalMetadata() view returns(bytes27)
func (_ChainStorageContainer *ChainStorageContainerCaller) GetGlobalMetadata(opts *bind.CallOpts) ([27]byte, error) {
	var out []interface{}
	err := _ChainStorageContainer.contract.Call(opts, &out, "getGlobalMetadata")

	if err != nil {
		return *new([27]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([27]byte)).(*[27]byte)

	return out0, err

}

// GetGlobalMetadata is a free data retrieval call binding the contract method 0xccf8f969.
//
// Solidity: function getGlobalMetadata() view returns(bytes27)
func (_ChainStorageContainer *ChainStorageContainerSession) GetGlobalMetadata() ([27]byte, error) {
	return _ChainStorageContainer.Contract.GetGlobalMetadata(&_ChainStorageContainer.CallOpts)
}

// GetGlobalMetadata is a free data retrieval call binding the contract method 0xccf8f969.
//
// Solidity: function getGlobalMetadata() view returns(bytes27)
func (_ChainStorageContainer *ChainStorageContainerCallerSession) GetGlobalMetadata() ([27]byte, error) {
	return _ChainStorageContainer.Contract.GetGlobalMetadata(&_ChainStorageContainer.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_ChainStorageContainer *ChainStorageContainerCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainStorageContainer.contract.Call(opts, &out, "length")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_ChainStorageContainer *ChainStorageContainerSession) Length() (*big.Int, error) {
	return _ChainStorageContainer.Contract.Length(&_ChainStorageContainer.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_ChainStorageContainer *ChainStorageContainerCallerSession) Length() (*big.Int, error) {
	return _ChainStorageContainer.Contract.Length(&_ChainStorageContainer.CallOpts)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_ChainStorageContainer *ChainStorageContainerCaller) LibAddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainStorageContainer.contract.Call(opts, &out, "libAddressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_ChainStorageContainer *ChainStorageContainerSession) LibAddressManager() (common.Address, error) {
	return _ChainStorageContainer.Contract.LibAddressManager(&_ChainStorageContainer.CallOpts)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_ChainStorageContainer *ChainStorageContainerCallerSession) LibAddressManager() (common.Address, error) {
	return _ChainStorageContainer.Contract.LibAddressManager(&_ChainStorageContainer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(string)
func (_ChainStorageContainer *ChainStorageContainerCaller) Owner(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChainStorageContainer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(string)
func (_ChainStorageContainer *ChainStorageContainerSession) Owner() (string, error) {
	return _ChainStorageContainer.Contract.Owner(&_ChainStorageContainer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(string)
func (_ChainStorageContainer *ChainStorageContainerCallerSession) Owner() (string, error) {
	return _ChainStorageContainer.Contract.Owner(&_ChainStorageContainer.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_ChainStorageContainer *ChainStorageContainerCaller) Resolve(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _ChainStorageContainer.contract.Call(opts, &out, "resolve", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_ChainStorageContainer *ChainStorageContainerSession) Resolve(_name string) (common.Address, error) {
	return _ChainStorageContainer.Contract.Resolve(&_ChainStorageContainer.CallOpts, _name)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_ChainStorageContainer *ChainStorageContainerCallerSession) Resolve(_name string) (common.Address, error) {
	return _ChainStorageContainer.Contract.Resolve(&_ChainStorageContainer.CallOpts, _name)
}

// DeleteElementsAfterInclusive is a paid mutator transaction binding the contract method 0x167fd681.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index, bytes27 _globalMetadata) returns()
func (_ChainStorageContainer *ChainStorageContainerTransactor) DeleteElementsAfterInclusive(opts *bind.TransactOpts, _index *big.Int, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.contract.Transact(opts, "deleteElementsAfterInclusive", _index, _globalMetadata)
}

// DeleteElementsAfterInclusive is a paid mutator transaction binding the contract method 0x167fd681.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index, bytes27 _globalMetadata) returns()
func (_ChainStorageContainer *ChainStorageContainerSession) DeleteElementsAfterInclusive(_index *big.Int, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.DeleteElementsAfterInclusive(&_ChainStorageContainer.TransactOpts, _index, _globalMetadata)
}

// DeleteElementsAfterInclusive is a paid mutator transaction binding the contract method 0x167fd681.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index, bytes27 _globalMetadata) returns()
func (_ChainStorageContainer *ChainStorageContainerTransactorSession) DeleteElementsAfterInclusive(_index *big.Int, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.DeleteElementsAfterInclusive(&_ChainStorageContainer.TransactOpts, _index, _globalMetadata)
}

// DeleteElementsAfterInclusive0 is a paid mutator transaction binding the contract method 0x4651d91e.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index) returns()
func (_ChainStorageContainer *ChainStorageContainerTransactor) DeleteElementsAfterInclusive0(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _ChainStorageContainer.contract.Transact(opts, "deleteElementsAfterInclusive0", _index)
}

// DeleteElementsAfterInclusive0 is a paid mutator transaction binding the contract method 0x4651d91e.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index) returns()
func (_ChainStorageContainer *ChainStorageContainerSession) DeleteElementsAfterInclusive0(_index *big.Int) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.DeleteElementsAfterInclusive0(&_ChainStorageContainer.TransactOpts, _index)
}

// DeleteElementsAfterInclusive0 is a paid mutator transaction binding the contract method 0x4651d91e.
//
// Solidity: function deleteElementsAfterInclusive(uint256 _index) returns()
func (_ChainStorageContainer *ChainStorageContainerTransactorSession) DeleteElementsAfterInclusive0(_index *big.Int) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.DeleteElementsAfterInclusive0(&_ChainStorageContainer.TransactOpts, _index)
}

// Push is a paid mutator transaction binding the contract method 0x2015276c.
//
// Solidity: function push(bytes32 _object, bytes27 _globalMetadata) returns()
func (_ChainStorageContainer *ChainStorageContainerTransactor) Push(opts *bind.TransactOpts, _object [32]byte, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.contract.Transact(opts, "push", _object, _globalMetadata)
}

// Push is a paid mutator transaction binding the contract method 0x2015276c.
//
// Solidity: function push(bytes32 _object, bytes27 _globalMetadata) returns()
func (_ChainStorageContainer *ChainStorageContainerSession) Push(_object [32]byte, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.Push(&_ChainStorageContainer.TransactOpts, _object, _globalMetadata)
}

// Push is a paid mutator transaction binding the contract method 0x2015276c.
//
// Solidity: function push(bytes32 _object, bytes27 _globalMetadata) returns()
func (_ChainStorageContainer *ChainStorageContainerTransactorSession) Push(_object [32]byte, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.Push(&_ChainStorageContainer.TransactOpts, _object, _globalMetadata)
}

// Push0 is a paid mutator transaction binding the contract method 0xb298e36b.
//
// Solidity: function push(bytes32 _object) returns()
func (_ChainStorageContainer *ChainStorageContainerTransactor) Push0(opts *bind.TransactOpts, _object [32]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.contract.Transact(opts, "push0", _object)
}

// Push0 is a paid mutator transaction binding the contract method 0xb298e36b.
//
// Solidity: function push(bytes32 _object) returns()
func (_ChainStorageContainer *ChainStorageContainerSession) Push0(_object [32]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.Push0(&_ChainStorageContainer.TransactOpts, _object)
}

// Push0 is a paid mutator transaction binding the contract method 0xb298e36b.
//
// Solidity: function push(bytes32 _object) returns()
func (_ChainStorageContainer *ChainStorageContainerTransactorSession) Push0(_object [32]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.Push0(&_ChainStorageContainer.TransactOpts, _object)
}

// SetGlobalMetadata is a paid mutator transaction binding the contract method 0x29061de2.
//
// Solidity: function setGlobalMetadata(bytes27 _globalMetadata) returns()
func (_ChainStorageContainer *ChainStorageContainerTransactor) SetGlobalMetadata(opts *bind.TransactOpts, _globalMetadata [27]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.contract.Transact(opts, "setGlobalMetadata", _globalMetadata)
}

// SetGlobalMetadata is a paid mutator transaction binding the contract method 0x29061de2.
//
// Solidity: function setGlobalMetadata(bytes27 _globalMetadata) returns()
func (_ChainStorageContainer *ChainStorageContainerSession) SetGlobalMetadata(_globalMetadata [27]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.SetGlobalMetadata(&_ChainStorageContainer.TransactOpts, _globalMetadata)
}

// SetGlobalMetadata is a paid mutator transaction binding the contract method 0x29061de2.
//
// Solidity: function setGlobalMetadata(bytes27 _globalMetadata) returns()
func (_ChainStorageContainer *ChainStorageContainerTransactorSession) SetGlobalMetadata(_globalMetadata [27]byte) (*types.Transaction, error) {
	return _ChainStorageContainer.Contract.SetGlobalMetadata(&_ChainStorageContainer.TransactOpts, _globalMetadata)
}

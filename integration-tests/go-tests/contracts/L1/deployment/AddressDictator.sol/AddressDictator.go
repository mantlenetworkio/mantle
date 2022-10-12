// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AddressDictator

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

// AddressDictatorNamedAddress is an auto generated low-level Go binding around an user-defined struct.
type AddressDictatorNamedAddress struct {
	Name string
	Addr common.Address
}

// AddressDictatorMetaData contains all meta data concerning the AddressDictator contract.
var AddressDictatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLib_AddressManager\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_finalOwner\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_names\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"finalOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNamedAddresses\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structAddressDictator.NamedAddress[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"contractLib_AddressManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"returnOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AddressDictatorABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressDictatorMetaData.ABI instead.
var AddressDictatorABI = AddressDictatorMetaData.ABI

// AddressDictator is an auto generated Go binding around an Ethereum contract.
type AddressDictator struct {
	AddressDictatorCaller     // Read-only binding to the contract
	AddressDictatorTransactor // Write-only binding to the contract
	AddressDictatorFilterer   // Log filterer for contract events
}

// AddressDictatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressDictatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressDictatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressDictatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressDictatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressDictatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressDictatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressDictatorSession struct {
	Contract     *AddressDictator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressDictatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressDictatorCallerSession struct {
	Contract *AddressDictatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AddressDictatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressDictatorTransactorSession struct {
	Contract     *AddressDictatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AddressDictatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressDictatorRaw struct {
	Contract *AddressDictator // Generic contract binding to access the raw methods on
}

// AddressDictatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressDictatorCallerRaw struct {
	Contract *AddressDictatorCaller // Generic read-only contract binding to access the raw methods on
}

// AddressDictatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressDictatorTransactorRaw struct {
	Contract *AddressDictatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressDictator creates a new instance of AddressDictator, bound to a specific deployed contract.
func NewAddressDictator(address common.Address, backend bind.ContractBackend) (*AddressDictator, error) {
	contract, err := bindAddressDictator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressDictator{AddressDictatorCaller: AddressDictatorCaller{contract: contract}, AddressDictatorTransactor: AddressDictatorTransactor{contract: contract}, AddressDictatorFilterer: AddressDictatorFilterer{contract: contract}}, nil
}

// NewAddressDictatorCaller creates a new read-only instance of AddressDictator, bound to a specific deployed contract.
func NewAddressDictatorCaller(address common.Address, caller bind.ContractCaller) (*AddressDictatorCaller, error) {
	contract, err := bindAddressDictator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressDictatorCaller{contract: contract}, nil
}

// NewAddressDictatorTransactor creates a new write-only instance of AddressDictator, bound to a specific deployed contract.
func NewAddressDictatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressDictatorTransactor, error) {
	contract, err := bindAddressDictator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressDictatorTransactor{contract: contract}, nil
}

// NewAddressDictatorFilterer creates a new log filterer instance of AddressDictator, bound to a specific deployed contract.
func NewAddressDictatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressDictatorFilterer, error) {
	contract, err := bindAddressDictator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressDictatorFilterer{contract: contract}, nil
}

// bindAddressDictator binds a generic wrapper to an already deployed contract.
func bindAddressDictator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressDictatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressDictator *AddressDictatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressDictator.Contract.AddressDictatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressDictator *AddressDictatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressDictator.Contract.AddressDictatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressDictator *AddressDictatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressDictator.Contract.AddressDictatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressDictator *AddressDictatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressDictator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressDictator *AddressDictatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressDictator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressDictator *AddressDictatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressDictator.Contract.contract.Transact(opts, method, params...)
}

// FinalOwner is a free data retrieval call binding the contract method 0x17ad94ec.
//
// Solidity: function finalOwner() view returns(address)
func (_AddressDictator *AddressDictatorCaller) FinalOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressDictator.contract.Call(opts, &out, "finalOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FinalOwner is a free data retrieval call binding the contract method 0x17ad94ec.
//
// Solidity: function finalOwner() view returns(address)
func (_AddressDictator *AddressDictatorSession) FinalOwner() (common.Address, error) {
	return _AddressDictator.Contract.FinalOwner(&_AddressDictator.CallOpts)
}

// FinalOwner is a free data retrieval call binding the contract method 0x17ad94ec.
//
// Solidity: function finalOwner() view returns(address)
func (_AddressDictator *AddressDictatorCallerSession) FinalOwner() (common.Address, error) {
	return _AddressDictator.Contract.FinalOwner(&_AddressDictator.CallOpts)
}

// GetNamedAddresses is a free data retrieval call binding the contract method 0x3ccad6fc.
//
// Solidity: function getNamedAddresses() view returns((string,address)[])
func (_AddressDictator *AddressDictatorCaller) GetNamedAddresses(opts *bind.CallOpts) ([]AddressDictatorNamedAddress, error) {
	var out []interface{}
	err := _AddressDictator.contract.Call(opts, &out, "getNamedAddresses")

	if err != nil {
		return *new([]AddressDictatorNamedAddress), err
	}

	out0 := *abi.ConvertType(out[0], new([]AddressDictatorNamedAddress)).(*[]AddressDictatorNamedAddress)

	return out0, err

}

// GetNamedAddresses is a free data retrieval call binding the contract method 0x3ccad6fc.
//
// Solidity: function getNamedAddresses() view returns((string,address)[])
func (_AddressDictator *AddressDictatorSession) GetNamedAddresses() ([]AddressDictatorNamedAddress, error) {
	return _AddressDictator.Contract.GetNamedAddresses(&_AddressDictator.CallOpts)
}

// GetNamedAddresses is a free data retrieval call binding the contract method 0x3ccad6fc.
//
// Solidity: function getNamedAddresses() view returns((string,address)[])
func (_AddressDictator *AddressDictatorCallerSession) GetNamedAddresses() ([]AddressDictatorNamedAddress, error) {
	return _AddressDictator.Contract.GetNamedAddresses(&_AddressDictator.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_AddressDictator *AddressDictatorCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressDictator.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_AddressDictator *AddressDictatorSession) Manager() (common.Address, error) {
	return _AddressDictator.Contract.Manager(&_AddressDictator.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_AddressDictator *AddressDictatorCallerSession) Manager() (common.Address, error) {
	return _AddressDictator.Contract.Manager(&_AddressDictator.CallOpts)
}

// ReturnOwnership is a paid mutator transaction binding the contract method 0x297d1a34.
//
// Solidity: function returnOwnership() returns()
func (_AddressDictator *AddressDictatorTransactor) ReturnOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressDictator.contract.Transact(opts, "returnOwnership")
}

// ReturnOwnership is a paid mutator transaction binding the contract method 0x297d1a34.
//
// Solidity: function returnOwnership() returns()
func (_AddressDictator *AddressDictatorSession) ReturnOwnership() (*types.Transaction, error) {
	return _AddressDictator.Contract.ReturnOwnership(&_AddressDictator.TransactOpts)
}

// ReturnOwnership is a paid mutator transaction binding the contract method 0x297d1a34.
//
// Solidity: function returnOwnership() returns()
func (_AddressDictator *AddressDictatorTransactorSession) ReturnOwnership() (*types.Transaction, error) {
	return _AddressDictator.Contract.ReturnOwnership(&_AddressDictator.TransactOpts)
}

// SetAddresses is a paid mutator transaction binding the contract method 0xbc3a429b.
//
// Solidity: function setAddresses() returns()
func (_AddressDictator *AddressDictatorTransactor) SetAddresses(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressDictator.contract.Transact(opts, "setAddresses")
}

// SetAddresses is a paid mutator transaction binding the contract method 0xbc3a429b.
//
// Solidity: function setAddresses() returns()
func (_AddressDictator *AddressDictatorSession) SetAddresses() (*types.Transaction, error) {
	return _AddressDictator.Contract.SetAddresses(&_AddressDictator.TransactOpts)
}

// SetAddresses is a paid mutator transaction binding the contract method 0xbc3a429b.
//
// Solidity: function setAddresses() returns()
func (_AddressDictator *AddressDictatorTransactorSession) SetAddresses() (*types.Transaction, error) {
	return _AddressDictator.Contract.SetAddresses(&_AddressDictator.TransactOpts)
}

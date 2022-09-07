// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BVM_DeployerWhitelist

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

// BVMDeployerWhitelistMetaData contains all meta data concerning the BVMDeployerWhitelist contract.
var BVMDeployerWhitelistMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"}],\"name\":\"WhitelistDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"WhitelistStatusChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"enableArbitraryContractDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"}],\"name\":\"isDeployerAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isWhitelisted\",\"type\":\"bool\"}],\"name\":\"setWhitelistedDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BVMDeployerWhitelistABI is the input ABI used to generate the binding from.
// Deprecated: Use BVMDeployerWhitelistMetaData.ABI instead.
var BVMDeployerWhitelistABI = BVMDeployerWhitelistMetaData.ABI

// BVMDeployerWhitelist is an auto generated Go binding around an Ethereum contract.
type BVMDeployerWhitelist struct {
	BVMDeployerWhitelistCaller     // Read-only binding to the contract
	BVMDeployerWhitelistTransactor // Write-only binding to the contract
	BVMDeployerWhitelistFilterer   // Log filterer for contract events
}

// BVMDeployerWhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type BVMDeployerWhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMDeployerWhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BVMDeployerWhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMDeployerWhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BVMDeployerWhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMDeployerWhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BVMDeployerWhitelistSession struct {
	Contract     *BVMDeployerWhitelist // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BVMDeployerWhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BVMDeployerWhitelistCallerSession struct {
	Contract *BVMDeployerWhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BVMDeployerWhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BVMDeployerWhitelistTransactorSession struct {
	Contract     *BVMDeployerWhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BVMDeployerWhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type BVMDeployerWhitelistRaw struct {
	Contract *BVMDeployerWhitelist // Generic contract binding to access the raw methods on
}

// BVMDeployerWhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BVMDeployerWhitelistCallerRaw struct {
	Contract *BVMDeployerWhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// BVMDeployerWhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BVMDeployerWhitelistTransactorRaw struct {
	Contract *BVMDeployerWhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBVMDeployerWhitelist creates a new instance of BVMDeployerWhitelist, bound to a specific deployed contract.
func NewBVMDeployerWhitelist(address common.Address, backend bind.ContractBackend) (*BVMDeployerWhitelist, error) {
	contract, err := bindBVMDeployerWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BVMDeployerWhitelist{BVMDeployerWhitelistCaller: BVMDeployerWhitelistCaller{contract: contract}, BVMDeployerWhitelistTransactor: BVMDeployerWhitelistTransactor{contract: contract}, BVMDeployerWhitelistFilterer: BVMDeployerWhitelistFilterer{contract: contract}}, nil
}

// NewBVMDeployerWhitelistCaller creates a new read-only instance of BVMDeployerWhitelist, bound to a specific deployed contract.
func NewBVMDeployerWhitelistCaller(address common.Address, caller bind.ContractCaller) (*BVMDeployerWhitelistCaller, error) {
	contract, err := bindBVMDeployerWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BVMDeployerWhitelistCaller{contract: contract}, nil
}

// NewBVMDeployerWhitelistTransactor creates a new write-only instance of BVMDeployerWhitelist, bound to a specific deployed contract.
func NewBVMDeployerWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*BVMDeployerWhitelistTransactor, error) {
	contract, err := bindBVMDeployerWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BVMDeployerWhitelistTransactor{contract: contract}, nil
}

// NewBVMDeployerWhitelistFilterer creates a new log filterer instance of BVMDeployerWhitelist, bound to a specific deployed contract.
func NewBVMDeployerWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*BVMDeployerWhitelistFilterer, error) {
	contract, err := bindBVMDeployerWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BVMDeployerWhitelistFilterer{contract: contract}, nil
}

// bindBVMDeployerWhitelist binds a generic wrapper to an already deployed contract.
func bindBVMDeployerWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BVMDeployerWhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMDeployerWhitelist *BVMDeployerWhitelistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMDeployerWhitelist.Contract.BVMDeployerWhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMDeployerWhitelist *BVMDeployerWhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMDeployerWhitelist.Contract.BVMDeployerWhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMDeployerWhitelist *BVMDeployerWhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMDeployerWhitelist.Contract.BVMDeployerWhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMDeployerWhitelist *BVMDeployerWhitelistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMDeployerWhitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMDeployerWhitelist *BVMDeployerWhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMDeployerWhitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMDeployerWhitelist *BVMDeployerWhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMDeployerWhitelist.Contract.contract.Transact(opts, method, params...)
}

// IsDeployerAllowed is a free data retrieval call binding the contract method 0xb1540a01.
//
// Solidity: function isDeployerAllowed(address _deployer) view returns(bool)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistCaller) IsDeployerAllowed(opts *bind.CallOpts, _deployer common.Address) (bool, error) {
	var out []interface{}
	err := _BVMDeployerWhitelist.contract.Call(opts, &out, "isDeployerAllowed", _deployer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeployerAllowed is a free data retrieval call binding the contract method 0xb1540a01.
//
// Solidity: function isDeployerAllowed(address _deployer) view returns(bool)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistSession) IsDeployerAllowed(_deployer common.Address) (bool, error) {
	return _BVMDeployerWhitelist.Contract.IsDeployerAllowed(&_BVMDeployerWhitelist.CallOpts, _deployer)
}

// IsDeployerAllowed is a free data retrieval call binding the contract method 0xb1540a01.
//
// Solidity: function isDeployerAllowed(address _deployer) view returns(bool)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistCallerSession) IsDeployerAllowed(_deployer common.Address) (bool, error) {
	return _BVMDeployerWhitelist.Contract.IsDeployerAllowed(&_BVMDeployerWhitelist.CallOpts, _deployer)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMDeployerWhitelist.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistSession) Owner() (common.Address, error) {
	return _BVMDeployerWhitelist.Contract.Owner(&_BVMDeployerWhitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistCallerSession) Owner() (common.Address, error) {
	return _BVMDeployerWhitelist.Contract.Owner(&_BVMDeployerWhitelist.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BVMDeployerWhitelist.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistSession) Whitelist(arg0 common.Address) (bool, error) {
	return _BVMDeployerWhitelist.Contract.Whitelist(&_BVMDeployerWhitelist.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _BVMDeployerWhitelist.Contract.Whitelist(&_BVMDeployerWhitelist.CallOpts, arg0)
}

// EnableArbitraryContractDeployment is a paid mutator transaction binding the contract method 0xbdc7b54f.
//
// Solidity: function enableArbitraryContractDeployment() returns()
func (_BVMDeployerWhitelist *BVMDeployerWhitelistTransactor) EnableArbitraryContractDeployment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMDeployerWhitelist.contract.Transact(opts, "enableArbitraryContractDeployment")
}

// EnableArbitraryContractDeployment is a paid mutator transaction binding the contract method 0xbdc7b54f.
//
// Solidity: function enableArbitraryContractDeployment() returns()
func (_BVMDeployerWhitelist *BVMDeployerWhitelistSession) EnableArbitraryContractDeployment() (*types.Transaction, error) {
	return _BVMDeployerWhitelist.Contract.EnableArbitraryContractDeployment(&_BVMDeployerWhitelist.TransactOpts)
}

// EnableArbitraryContractDeployment is a paid mutator transaction binding the contract method 0xbdc7b54f.
//
// Solidity: function enableArbitraryContractDeployment() returns()
func (_BVMDeployerWhitelist *BVMDeployerWhitelistTransactorSession) EnableArbitraryContractDeployment() (*types.Transaction, error) {
	return _BVMDeployerWhitelist.Contract.EnableArbitraryContractDeployment(&_BVMDeployerWhitelist.TransactOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_BVMDeployerWhitelist *BVMDeployerWhitelistTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _BVMDeployerWhitelist.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_BVMDeployerWhitelist *BVMDeployerWhitelistSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _BVMDeployerWhitelist.Contract.SetOwner(&_BVMDeployerWhitelist.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_BVMDeployerWhitelist *BVMDeployerWhitelistTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _BVMDeployerWhitelist.Contract.SetOwner(&_BVMDeployerWhitelist.TransactOpts, _owner)
}

// SetWhitelistedDeployer is a paid mutator transaction binding the contract method 0x08fd6322.
//
// Solidity: function setWhitelistedDeployer(address _deployer, bool _isWhitelisted) returns()
func (_BVMDeployerWhitelist *BVMDeployerWhitelistTransactor) SetWhitelistedDeployer(opts *bind.TransactOpts, _deployer common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _BVMDeployerWhitelist.contract.Transact(opts, "setWhitelistedDeployer", _deployer, _isWhitelisted)
}

// SetWhitelistedDeployer is a paid mutator transaction binding the contract method 0x08fd6322.
//
// Solidity: function setWhitelistedDeployer(address _deployer, bool _isWhitelisted) returns()
func (_BVMDeployerWhitelist *BVMDeployerWhitelistSession) SetWhitelistedDeployer(_deployer common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _BVMDeployerWhitelist.Contract.SetWhitelistedDeployer(&_BVMDeployerWhitelist.TransactOpts, _deployer, _isWhitelisted)
}

// SetWhitelistedDeployer is a paid mutator transaction binding the contract method 0x08fd6322.
//
// Solidity: function setWhitelistedDeployer(address _deployer, bool _isWhitelisted) returns()
func (_BVMDeployerWhitelist *BVMDeployerWhitelistTransactorSession) SetWhitelistedDeployer(_deployer common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _BVMDeployerWhitelist.Contract.SetWhitelistedDeployer(&_BVMDeployerWhitelist.TransactOpts, _deployer, _isWhitelisted)
}

// BVMDeployerWhitelistOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the BVMDeployerWhitelist contract.
type BVMDeployerWhitelistOwnerChangedIterator struct {
	Event *BVMDeployerWhitelistOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BVMDeployerWhitelistOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMDeployerWhitelistOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BVMDeployerWhitelistOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BVMDeployerWhitelistOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMDeployerWhitelistOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMDeployerWhitelistOwnerChanged represents a OwnerChanged event raised by the BVMDeployerWhitelist contract.
type BVMDeployerWhitelistOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*BVMDeployerWhitelistOwnerChangedIterator, error) {

	logs, sub, err := _BVMDeployerWhitelist.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &BVMDeployerWhitelistOwnerChangedIterator{contract: _BVMDeployerWhitelist.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *BVMDeployerWhitelistOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _BVMDeployerWhitelist.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMDeployerWhitelistOwnerChanged)
				if err := _BVMDeployerWhitelist.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistFilterer) ParseOwnerChanged(log types.Log) (*BVMDeployerWhitelistOwnerChanged, error) {
	event := new(BVMDeployerWhitelistOwnerChanged)
	if err := _BVMDeployerWhitelist.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMDeployerWhitelistWhitelistDisabledIterator is returned from FilterWhitelistDisabled and is used to iterate over the raw logs and unpacked data for WhitelistDisabled events raised by the BVMDeployerWhitelist contract.
type BVMDeployerWhitelistWhitelistDisabledIterator struct {
	Event *BVMDeployerWhitelistWhitelistDisabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BVMDeployerWhitelistWhitelistDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMDeployerWhitelistWhitelistDisabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BVMDeployerWhitelistWhitelistDisabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BVMDeployerWhitelistWhitelistDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMDeployerWhitelistWhitelistDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMDeployerWhitelistWhitelistDisabled represents a WhitelistDisabled event raised by the BVMDeployerWhitelist contract.
type BVMDeployerWhitelistWhitelistDisabled struct {
	OldOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistDisabled is a free log retrieval operation binding the contract event 0xc0e106cf568e50698fdbde1eff56f5a5c966cc7958e37e276918e9e4ccdf8cd4.
//
// Solidity: event WhitelistDisabled(address oldOwner)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistFilterer) FilterWhitelistDisabled(opts *bind.FilterOpts) (*BVMDeployerWhitelistWhitelistDisabledIterator, error) {

	logs, sub, err := _BVMDeployerWhitelist.contract.FilterLogs(opts, "WhitelistDisabled")
	if err != nil {
		return nil, err
	}
	return &BVMDeployerWhitelistWhitelistDisabledIterator{contract: _BVMDeployerWhitelist.contract, event: "WhitelistDisabled", logs: logs, sub: sub}, nil
}

// WatchWhitelistDisabled is a free log subscription operation binding the contract event 0xc0e106cf568e50698fdbde1eff56f5a5c966cc7958e37e276918e9e4ccdf8cd4.
//
// Solidity: event WhitelistDisabled(address oldOwner)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistFilterer) WatchWhitelistDisabled(opts *bind.WatchOpts, sink chan<- *BVMDeployerWhitelistWhitelistDisabled) (event.Subscription, error) {

	logs, sub, err := _BVMDeployerWhitelist.contract.WatchLogs(opts, "WhitelistDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMDeployerWhitelistWhitelistDisabled)
				if err := _BVMDeployerWhitelist.contract.UnpackLog(event, "WhitelistDisabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelistDisabled is a log parse operation binding the contract event 0xc0e106cf568e50698fdbde1eff56f5a5c966cc7958e37e276918e9e4ccdf8cd4.
//
// Solidity: event WhitelistDisabled(address oldOwner)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistFilterer) ParseWhitelistDisabled(log types.Log) (*BVMDeployerWhitelistWhitelistDisabled, error) {
	event := new(BVMDeployerWhitelistWhitelistDisabled)
	if err := _BVMDeployerWhitelist.contract.UnpackLog(event, "WhitelistDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMDeployerWhitelistWhitelistStatusChangedIterator is returned from FilterWhitelistStatusChanged and is used to iterate over the raw logs and unpacked data for WhitelistStatusChanged events raised by the BVMDeployerWhitelist contract.
type BVMDeployerWhitelistWhitelistStatusChangedIterator struct {
	Event *BVMDeployerWhitelistWhitelistStatusChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BVMDeployerWhitelistWhitelistStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMDeployerWhitelistWhitelistStatusChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BVMDeployerWhitelistWhitelistStatusChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BVMDeployerWhitelistWhitelistStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMDeployerWhitelistWhitelistStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMDeployerWhitelistWhitelistStatusChanged represents a WhitelistStatusChanged event raised by the BVMDeployerWhitelist contract.
type BVMDeployerWhitelistWhitelistStatusChanged struct {
	Deployer    common.Address
	Whitelisted bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWhitelistStatusChanged is a free log retrieval operation binding the contract event 0x8daaf060c3306c38e068a75c054bf96ecd85a3db1252712c4d93632744c42e0d.
//
// Solidity: event WhitelistStatusChanged(address deployer, bool whitelisted)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistFilterer) FilterWhitelistStatusChanged(opts *bind.FilterOpts) (*BVMDeployerWhitelistWhitelistStatusChangedIterator, error) {

	logs, sub, err := _BVMDeployerWhitelist.contract.FilterLogs(opts, "WhitelistStatusChanged")
	if err != nil {
		return nil, err
	}
	return &BVMDeployerWhitelistWhitelistStatusChangedIterator{contract: _BVMDeployerWhitelist.contract, event: "WhitelistStatusChanged", logs: logs, sub: sub}, nil
}

// WatchWhitelistStatusChanged is a free log subscription operation binding the contract event 0x8daaf060c3306c38e068a75c054bf96ecd85a3db1252712c4d93632744c42e0d.
//
// Solidity: event WhitelistStatusChanged(address deployer, bool whitelisted)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistFilterer) WatchWhitelistStatusChanged(opts *bind.WatchOpts, sink chan<- *BVMDeployerWhitelistWhitelistStatusChanged) (event.Subscription, error) {

	logs, sub, err := _BVMDeployerWhitelist.contract.WatchLogs(opts, "WhitelistStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMDeployerWhitelistWhitelistStatusChanged)
				if err := _BVMDeployerWhitelist.contract.UnpackLog(event, "WhitelistStatusChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelistStatusChanged is a log parse operation binding the contract event 0x8daaf060c3306c38e068a75c054bf96ecd85a3db1252712c4d93632744c42e0d.
//
// Solidity: event WhitelistStatusChanged(address deployer, bool whitelisted)
func (_BVMDeployerWhitelist *BVMDeployerWhitelistFilterer) ParseWhitelistStatusChanged(log types.Log) (*BVMDeployerWhitelistWhitelistStatusChanged, error) {
	event := new(BVMDeployerWhitelistWhitelistStatusChanged)
	if err := _BVMDeployerWhitelist.contract.UnpackLog(event, "WhitelistStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

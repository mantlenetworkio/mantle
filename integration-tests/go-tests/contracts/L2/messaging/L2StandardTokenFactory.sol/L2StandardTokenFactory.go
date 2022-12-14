// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package L2StandardTokenFactory

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

// L2StandardTokenFactoryMetaData contains all meta data concerning the L2StandardTokenFactory contract.
var L2StandardTokenFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"}],\"name\":\"StandardL2TokenCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimal\",\"type\":\"uint8\"}],\"name\":\"createStandardL2Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L2StandardTokenFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use L2StandardTokenFactoryMetaData.ABI instead.
var L2StandardTokenFactoryABI = L2StandardTokenFactoryMetaData.ABI

// L2StandardTokenFactory is an auto generated Go binding around an Ethereum contract.
type L2StandardTokenFactory struct {
	L2StandardTokenFactoryCaller     // Read-only binding to the contract
	L2StandardTokenFactoryTransactor // Write-only binding to the contract
	L2StandardTokenFactoryFilterer   // Log filterer for contract events
}

// L2StandardTokenFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2StandardTokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StandardTokenFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2StandardTokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StandardTokenFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2StandardTokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StandardTokenFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2StandardTokenFactorySession struct {
	Contract     *L2StandardTokenFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2StandardTokenFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2StandardTokenFactoryCallerSession struct {
	Contract *L2StandardTokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// L2StandardTokenFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2StandardTokenFactoryTransactorSession struct {
	Contract     *L2StandardTokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L2StandardTokenFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2StandardTokenFactoryRaw struct {
	Contract *L2StandardTokenFactory // Generic contract binding to access the raw methods on
}

// L2StandardTokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2StandardTokenFactoryCallerRaw struct {
	Contract *L2StandardTokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// L2StandardTokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2StandardTokenFactoryTransactorRaw struct {
	Contract *L2StandardTokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2StandardTokenFactory creates a new instance of L2StandardTokenFactory, bound to a specific deployed contract.
func NewL2StandardTokenFactory(address common.Address, backend bind.ContractBackend) (*L2StandardTokenFactory, error) {
	contract, err := bindL2StandardTokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2StandardTokenFactory{L2StandardTokenFactoryCaller: L2StandardTokenFactoryCaller{contract: contract}, L2StandardTokenFactoryTransactor: L2StandardTokenFactoryTransactor{contract: contract}, L2StandardTokenFactoryFilterer: L2StandardTokenFactoryFilterer{contract: contract}}, nil
}

// NewL2StandardTokenFactoryCaller creates a new read-only instance of L2StandardTokenFactory, bound to a specific deployed contract.
func NewL2StandardTokenFactoryCaller(address common.Address, caller bind.ContractCaller) (*L2StandardTokenFactoryCaller, error) {
	contract, err := bindL2StandardTokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2StandardTokenFactoryCaller{contract: contract}, nil
}

// NewL2StandardTokenFactoryTransactor creates a new write-only instance of L2StandardTokenFactory, bound to a specific deployed contract.
func NewL2StandardTokenFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*L2StandardTokenFactoryTransactor, error) {
	contract, err := bindL2StandardTokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2StandardTokenFactoryTransactor{contract: contract}, nil
}

// NewL2StandardTokenFactoryFilterer creates a new log filterer instance of L2StandardTokenFactory, bound to a specific deployed contract.
func NewL2StandardTokenFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*L2StandardTokenFactoryFilterer, error) {
	contract, err := bindL2StandardTokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2StandardTokenFactoryFilterer{contract: contract}, nil
}

// bindL2StandardTokenFactory binds a generic wrapper to an already deployed contract.
func bindL2StandardTokenFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2StandardTokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2StandardTokenFactory *L2StandardTokenFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2StandardTokenFactory.Contract.L2StandardTokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2StandardTokenFactory *L2StandardTokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2StandardTokenFactory.Contract.L2StandardTokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2StandardTokenFactory *L2StandardTokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2StandardTokenFactory.Contract.L2StandardTokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2StandardTokenFactory *L2StandardTokenFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2StandardTokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2StandardTokenFactory *L2StandardTokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2StandardTokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2StandardTokenFactory *L2StandardTokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2StandardTokenFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x0769a969.
//
// Solidity: function createStandardL2Token(address _l1Token, string _name, string _symbol, uint8 _decimal) returns()
func (_L2StandardTokenFactory *L2StandardTokenFactoryTransactor) CreateStandardL2Token(opts *bind.TransactOpts, _l1Token common.Address, _name string, _symbol string, _decimal uint8) (*types.Transaction, error) {
	return _L2StandardTokenFactory.contract.Transact(opts, "createStandardL2Token", _l1Token, _name, _symbol, _decimal)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x0769a969.
//
// Solidity: function createStandardL2Token(address _l1Token, string _name, string _symbol, uint8 _decimal) returns()
func (_L2StandardTokenFactory *L2StandardTokenFactorySession) CreateStandardL2Token(_l1Token common.Address, _name string, _symbol string, _decimal uint8) (*types.Transaction, error) {
	return _L2StandardTokenFactory.Contract.CreateStandardL2Token(&_L2StandardTokenFactory.TransactOpts, _l1Token, _name, _symbol, _decimal)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x0769a969.
//
// Solidity: function createStandardL2Token(address _l1Token, string _name, string _symbol, uint8 _decimal) returns()
func (_L2StandardTokenFactory *L2StandardTokenFactoryTransactorSession) CreateStandardL2Token(_l1Token common.Address, _name string, _symbol string, _decimal uint8) (*types.Transaction, error) {
	return _L2StandardTokenFactory.Contract.CreateStandardL2Token(&_L2StandardTokenFactory.TransactOpts, _l1Token, _name, _symbol, _decimal)
}

// L2StandardTokenFactoryStandardL2TokenCreatedIterator is returned from FilterStandardL2TokenCreated and is used to iterate over the raw logs and unpacked data for StandardL2TokenCreated events raised by the L2StandardTokenFactory contract.
type L2StandardTokenFactoryStandardL2TokenCreatedIterator struct {
	Event *L2StandardTokenFactoryStandardL2TokenCreated // Event containing the contract specifics and raw log

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
func (it *L2StandardTokenFactoryStandardL2TokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StandardTokenFactoryStandardL2TokenCreated)
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
		it.Event = new(L2StandardTokenFactoryStandardL2TokenCreated)
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
func (it *L2StandardTokenFactoryStandardL2TokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StandardTokenFactoryStandardL2TokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StandardTokenFactoryStandardL2TokenCreated represents a StandardL2TokenCreated event raised by the L2StandardTokenFactory contract.
type L2StandardTokenFactoryStandardL2TokenCreated struct {
	L1Token common.Address
	L2Token common.Address
	Decimal uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStandardL2TokenCreated is a free log retrieval operation binding the contract event 0x41e27481c6f764357db26ae29b68b9f3aafc40b87444459cbf50d338c7531732.
//
// Solidity: event StandardL2TokenCreated(address indexed _l1Token, address indexed _l2Token, uint8 decimal)
func (_L2StandardTokenFactory *L2StandardTokenFactoryFilterer) FilterStandardL2TokenCreated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address) (*L2StandardTokenFactoryStandardL2TokenCreatedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}

	logs, sub, err := _L2StandardTokenFactory.contract.FilterLogs(opts, "StandardL2TokenCreated", _l1TokenRule, _l2TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2StandardTokenFactoryStandardL2TokenCreatedIterator{contract: _L2StandardTokenFactory.contract, event: "StandardL2TokenCreated", logs: logs, sub: sub}, nil
}

// WatchStandardL2TokenCreated is a free log subscription operation binding the contract event 0x41e27481c6f764357db26ae29b68b9f3aafc40b87444459cbf50d338c7531732.
//
// Solidity: event StandardL2TokenCreated(address indexed _l1Token, address indexed _l2Token, uint8 decimal)
func (_L2StandardTokenFactory *L2StandardTokenFactoryFilterer) WatchStandardL2TokenCreated(opts *bind.WatchOpts, sink chan<- *L2StandardTokenFactoryStandardL2TokenCreated, _l1Token []common.Address, _l2Token []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}

	logs, sub, err := _L2StandardTokenFactory.contract.WatchLogs(opts, "StandardL2TokenCreated", _l1TokenRule, _l2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StandardTokenFactoryStandardL2TokenCreated)
				if err := _L2StandardTokenFactory.contract.UnpackLog(event, "StandardL2TokenCreated", log); err != nil {
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

// ParseStandardL2TokenCreated is a log parse operation binding the contract event 0x41e27481c6f764357db26ae29b68b9f3aafc40b87444459cbf50d338c7531732.
//
// Solidity: event StandardL2TokenCreated(address indexed _l1Token, address indexed _l2Token, uint8 decimal)
func (_L2StandardTokenFactory *L2StandardTokenFactoryFilterer) ParseStandardL2TokenCreated(log types.Log) (*L2StandardTokenFactoryStandardL2TokenCreated, error) {
	event := new(L2StandardTokenFactoryStandardL2TokenCreated)
	if err := _L2StandardTokenFactory.contract.UnpackLog(event, "StandardL2TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

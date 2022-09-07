// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IL2ERC20Bridge

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

// IL2ERC20BridgeMetaData contains all meta data concerning the IL2ERC20Bridge contract.
var IL2ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DepositFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DepositFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l1Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l1Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IL2ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IL2ERC20BridgeMetaData.ABI instead.
var IL2ERC20BridgeABI = IL2ERC20BridgeMetaData.ABI

// IL2ERC20Bridge is an auto generated Go binding around an Ethereum contract.
type IL2ERC20Bridge struct {
	IL2ERC20BridgeCaller     // Read-only binding to the contract
	IL2ERC20BridgeTransactor // Write-only binding to the contract
	IL2ERC20BridgeFilterer   // Log filterer for contract events
}

// IL2ERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL2ERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL2ERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL2ERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL2ERC20BridgeSession struct {
	Contract     *IL2ERC20Bridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL2ERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL2ERC20BridgeCallerSession struct {
	Contract *IL2ERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IL2ERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL2ERC20BridgeTransactorSession struct {
	Contract     *IL2ERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IL2ERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL2ERC20BridgeRaw struct {
	Contract *IL2ERC20Bridge // Generic contract binding to access the raw methods on
}

// IL2ERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL2ERC20BridgeCallerRaw struct {
	Contract *IL2ERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IL2ERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL2ERC20BridgeTransactorRaw struct {
	Contract *IL2ERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL2ERC20Bridge creates a new instance of IL2ERC20Bridge, bound to a specific deployed contract.
func NewIL2ERC20Bridge(address common.Address, backend bind.ContractBackend) (*IL2ERC20Bridge, error) {
	contract, err := bindIL2ERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20Bridge{IL2ERC20BridgeCaller: IL2ERC20BridgeCaller{contract: contract}, IL2ERC20BridgeTransactor: IL2ERC20BridgeTransactor{contract: contract}, IL2ERC20BridgeFilterer: IL2ERC20BridgeFilterer{contract: contract}}, nil
}

// NewIL2ERC20BridgeCaller creates a new read-only instance of IL2ERC20Bridge, bound to a specific deployed contract.
func NewIL2ERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*IL2ERC20BridgeCaller, error) {
	contract, err := bindIL2ERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20BridgeCaller{contract: contract}, nil
}

// NewIL2ERC20BridgeTransactor creates a new write-only instance of IL2ERC20Bridge, bound to a specific deployed contract.
func NewIL2ERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IL2ERC20BridgeTransactor, error) {
	contract, err := bindIL2ERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20BridgeTransactor{contract: contract}, nil
}

// NewIL2ERC20BridgeFilterer creates a new log filterer instance of IL2ERC20Bridge, bound to a specific deployed contract.
func NewIL2ERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IL2ERC20BridgeFilterer, error) {
	contract, err := bindIL2ERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20BridgeFilterer{contract: contract}, nil
}

// bindIL2ERC20Bridge binds a generic wrapper to an already deployed contract.
func bindIL2ERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IL2ERC20BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2ERC20Bridge *IL2ERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2ERC20Bridge.Contract.IL2ERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2ERC20Bridge *IL2ERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.IL2ERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2ERC20Bridge *IL2ERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.IL2ERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2ERC20Bridge *IL2ERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2ERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2ERC20Bridge *IL2ERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2ERC20Bridge *IL2ERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x662a633a.
//
// Solidity: function finalizeDeposit(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_IL2ERC20Bridge *IL2ERC20BridgeTransactor) FinalizeDeposit(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL2ERC20Bridge.contract.Transact(opts, "finalizeDeposit", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x662a633a.
//
// Solidity: function finalizeDeposit(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_IL2ERC20Bridge *IL2ERC20BridgeSession) FinalizeDeposit(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.FinalizeDeposit(&_IL2ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x662a633a.
//
// Solidity: function finalizeDeposit(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_IL2ERC20Bridge *IL2ERC20BridgeTransactorSession) FinalizeDeposit(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.FinalizeDeposit(&_IL2ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// L1TokenBridge is a paid mutator transaction binding the contract method 0x36c717c1.
//
// Solidity: function l1TokenBridge() returns(address)
func (_IL2ERC20Bridge *IL2ERC20BridgeTransactor) L1TokenBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2ERC20Bridge.contract.Transact(opts, "l1TokenBridge")
}

// L1TokenBridge is a paid mutator transaction binding the contract method 0x36c717c1.
//
// Solidity: function l1TokenBridge() returns(address)
func (_IL2ERC20Bridge *IL2ERC20BridgeSession) L1TokenBridge() (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.L1TokenBridge(&_IL2ERC20Bridge.TransactOpts)
}

// L1TokenBridge is a paid mutator transaction binding the contract method 0x36c717c1.
//
// Solidity: function l1TokenBridge() returns(address)
func (_IL2ERC20Bridge *IL2ERC20BridgeTransactorSession) L1TokenBridge() (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.L1TokenBridge(&_IL2ERC20Bridge.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32b7006d.
//
// Solidity: function withdraw(address _l2Token, uint256 _amount, uint32 _l1Gas, bytes _data) returns()
func (_IL2ERC20Bridge *IL2ERC20BridgeTransactor) Withdraw(opts *bind.TransactOpts, _l2Token common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL2ERC20Bridge.contract.Transact(opts, "withdraw", _l2Token, _amount, _l1Gas, _data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32b7006d.
//
// Solidity: function withdraw(address _l2Token, uint256 _amount, uint32 _l1Gas, bytes _data) returns()
func (_IL2ERC20Bridge *IL2ERC20BridgeSession) Withdraw(_l2Token common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.Withdraw(&_IL2ERC20Bridge.TransactOpts, _l2Token, _amount, _l1Gas, _data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32b7006d.
//
// Solidity: function withdraw(address _l2Token, uint256 _amount, uint32 _l1Gas, bytes _data) returns()
func (_IL2ERC20Bridge *IL2ERC20BridgeTransactorSession) Withdraw(_l2Token common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.Withdraw(&_IL2ERC20Bridge.TransactOpts, _l2Token, _amount, _l1Gas, _data)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xa3a79548.
//
// Solidity: function withdrawTo(address _l2Token, address _to, uint256 _amount, uint32 _l1Gas, bytes _data) returns()
func (_IL2ERC20Bridge *IL2ERC20BridgeTransactor) WithdrawTo(opts *bind.TransactOpts, _l2Token common.Address, _to common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL2ERC20Bridge.contract.Transact(opts, "withdrawTo", _l2Token, _to, _amount, _l1Gas, _data)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xa3a79548.
//
// Solidity: function withdrawTo(address _l2Token, address _to, uint256 _amount, uint32 _l1Gas, bytes _data) returns()
func (_IL2ERC20Bridge *IL2ERC20BridgeSession) WithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.WithdrawTo(&_IL2ERC20Bridge.TransactOpts, _l2Token, _to, _amount, _l1Gas, _data)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xa3a79548.
//
// Solidity: function withdrawTo(address _l2Token, address _to, uint256 _amount, uint32 _l1Gas, bytes _data) returns()
func (_IL2ERC20Bridge *IL2ERC20BridgeTransactorSession) WithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL2ERC20Bridge.Contract.WithdrawTo(&_IL2ERC20Bridge.TransactOpts, _l2Token, _to, _amount, _l1Gas, _data)
}

// IL2ERC20BridgeDepositFailedIterator is returned from FilterDepositFailed and is used to iterate over the raw logs and unpacked data for DepositFailed events raised by the IL2ERC20Bridge contract.
type IL2ERC20BridgeDepositFailedIterator struct {
	Event *IL2ERC20BridgeDepositFailed // Event containing the contract specifics and raw log

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
func (it *IL2ERC20BridgeDepositFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2ERC20BridgeDepositFailed)
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
		it.Event = new(IL2ERC20BridgeDepositFailed)
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
func (it *IL2ERC20BridgeDepositFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2ERC20BridgeDepositFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2ERC20BridgeDepositFailed represents a DepositFailed event raised by the IL2ERC20Bridge contract.
type IL2ERC20BridgeDepositFailed struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositFailed is a free log retrieval operation binding the contract event 0x7ea89a4591614515571c2b51f5ea06494056f261c10ab1ed8c03c7590d87bce0.
//
// Solidity: event DepositFailed(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL2ERC20Bridge *IL2ERC20BridgeFilterer) FilterDepositFailed(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*IL2ERC20BridgeDepositFailedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IL2ERC20Bridge.contract.FilterLogs(opts, "DepositFailed", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20BridgeDepositFailedIterator{contract: _IL2ERC20Bridge.contract, event: "DepositFailed", logs: logs, sub: sub}, nil
}

// WatchDepositFailed is a free log subscription operation binding the contract event 0x7ea89a4591614515571c2b51f5ea06494056f261c10ab1ed8c03c7590d87bce0.
//
// Solidity: event DepositFailed(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL2ERC20Bridge *IL2ERC20BridgeFilterer) WatchDepositFailed(opts *bind.WatchOpts, sink chan<- *IL2ERC20BridgeDepositFailed, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IL2ERC20Bridge.contract.WatchLogs(opts, "DepositFailed", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2ERC20BridgeDepositFailed)
				if err := _IL2ERC20Bridge.contract.UnpackLog(event, "DepositFailed", log); err != nil {
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

// ParseDepositFailed is a log parse operation binding the contract event 0x7ea89a4591614515571c2b51f5ea06494056f261c10ab1ed8c03c7590d87bce0.
//
// Solidity: event DepositFailed(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL2ERC20Bridge *IL2ERC20BridgeFilterer) ParseDepositFailed(log types.Log) (*IL2ERC20BridgeDepositFailed, error) {
	event := new(IL2ERC20BridgeDepositFailed)
	if err := _IL2ERC20Bridge.contract.UnpackLog(event, "DepositFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2ERC20BridgeDepositFinalizedIterator is returned from FilterDepositFinalized and is used to iterate over the raw logs and unpacked data for DepositFinalized events raised by the IL2ERC20Bridge contract.
type IL2ERC20BridgeDepositFinalizedIterator struct {
	Event *IL2ERC20BridgeDepositFinalized // Event containing the contract specifics and raw log

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
func (it *IL2ERC20BridgeDepositFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2ERC20BridgeDepositFinalized)
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
		it.Event = new(IL2ERC20BridgeDepositFinalized)
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
func (it *IL2ERC20BridgeDepositFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2ERC20BridgeDepositFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2ERC20BridgeDepositFinalized represents a DepositFinalized event raised by the IL2ERC20Bridge contract.
type IL2ERC20BridgeDepositFinalized struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositFinalized is a free log retrieval operation binding the contract event 0xb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89.
//
// Solidity: event DepositFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL2ERC20Bridge *IL2ERC20BridgeFilterer) FilterDepositFinalized(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*IL2ERC20BridgeDepositFinalizedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IL2ERC20Bridge.contract.FilterLogs(opts, "DepositFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20BridgeDepositFinalizedIterator{contract: _IL2ERC20Bridge.contract, event: "DepositFinalized", logs: logs, sub: sub}, nil
}

// WatchDepositFinalized is a free log subscription operation binding the contract event 0xb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89.
//
// Solidity: event DepositFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL2ERC20Bridge *IL2ERC20BridgeFilterer) WatchDepositFinalized(opts *bind.WatchOpts, sink chan<- *IL2ERC20BridgeDepositFinalized, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IL2ERC20Bridge.contract.WatchLogs(opts, "DepositFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2ERC20BridgeDepositFinalized)
				if err := _IL2ERC20Bridge.contract.UnpackLog(event, "DepositFinalized", log); err != nil {
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

// ParseDepositFinalized is a log parse operation binding the contract event 0xb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89.
//
// Solidity: event DepositFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL2ERC20Bridge *IL2ERC20BridgeFilterer) ParseDepositFinalized(log types.Log) (*IL2ERC20BridgeDepositFinalized, error) {
	event := new(IL2ERC20BridgeDepositFinalized)
	if err := _IL2ERC20Bridge.contract.UnpackLog(event, "DepositFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2ERC20BridgeWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the IL2ERC20Bridge contract.
type IL2ERC20BridgeWithdrawalInitiatedIterator struct {
	Event *IL2ERC20BridgeWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *IL2ERC20BridgeWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2ERC20BridgeWithdrawalInitiated)
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
		it.Event = new(IL2ERC20BridgeWithdrawalInitiated)
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
func (it *IL2ERC20BridgeWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2ERC20BridgeWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2ERC20BridgeWithdrawalInitiated represents a WithdrawalInitiated event raised by the IL2ERC20Bridge contract.
type IL2ERC20BridgeWithdrawalInitiated struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e.
//
// Solidity: event WithdrawalInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL2ERC20Bridge *IL2ERC20BridgeFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*IL2ERC20BridgeWithdrawalInitiatedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IL2ERC20Bridge.contract.FilterLogs(opts, "WithdrawalInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20BridgeWithdrawalInitiatedIterator{contract: _IL2ERC20Bridge.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e.
//
// Solidity: event WithdrawalInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL2ERC20Bridge *IL2ERC20BridgeFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *IL2ERC20BridgeWithdrawalInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IL2ERC20Bridge.contract.WatchLogs(opts, "WithdrawalInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2ERC20BridgeWithdrawalInitiated)
				if err := _IL2ERC20Bridge.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e.
//
// Solidity: event WithdrawalInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL2ERC20Bridge *IL2ERC20BridgeFilterer) ParseWithdrawalInitiated(log types.Log) (*IL2ERC20BridgeWithdrawalInitiated, error) {
	event := new(IL2ERC20BridgeWithdrawalInitiated)
	if err := _IL2ERC20Bridge.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

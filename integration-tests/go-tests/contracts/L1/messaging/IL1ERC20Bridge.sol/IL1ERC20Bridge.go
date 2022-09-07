// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IL1ERC20Bridge

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

// IL1ERC20BridgeMetaData contains all meta data concerning the IL1ERC20Bridge contract.
var IL1ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20WithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeBitWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeERC20Withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IL1ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1ERC20BridgeMetaData.ABI instead.
var IL1ERC20BridgeABI = IL1ERC20BridgeMetaData.ABI

// IL1ERC20Bridge is an auto generated Go binding around an Ethereum contract.
type IL1ERC20Bridge struct {
	IL1ERC20BridgeCaller     // Read-only binding to the contract
	IL1ERC20BridgeTransactor // Write-only binding to the contract
	IL1ERC20BridgeFilterer   // Log filterer for contract events
}

// IL1ERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1ERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1ERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1ERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1ERC20BridgeSession struct {
	Contract     *IL1ERC20Bridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL1ERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1ERC20BridgeCallerSession struct {
	Contract *IL1ERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IL1ERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1ERC20BridgeTransactorSession struct {
	Contract     *IL1ERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IL1ERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1ERC20BridgeRaw struct {
	Contract *IL1ERC20Bridge // Generic contract binding to access the raw methods on
}

// IL1ERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1ERC20BridgeCallerRaw struct {
	Contract *IL1ERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IL1ERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1ERC20BridgeTransactorRaw struct {
	Contract *IL1ERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1ERC20Bridge creates a new instance of IL1ERC20Bridge, bound to a specific deployed contract.
func NewIL1ERC20Bridge(address common.Address, backend bind.ContractBackend) (*IL1ERC20Bridge, error) {
	contract, err := bindIL1ERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20Bridge{IL1ERC20BridgeCaller: IL1ERC20BridgeCaller{contract: contract}, IL1ERC20BridgeTransactor: IL1ERC20BridgeTransactor{contract: contract}, IL1ERC20BridgeFilterer: IL1ERC20BridgeFilterer{contract: contract}}, nil
}

// NewIL1ERC20BridgeCaller creates a new read-only instance of IL1ERC20Bridge, bound to a specific deployed contract.
func NewIL1ERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*IL1ERC20BridgeCaller, error) {
	contract, err := bindIL1ERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20BridgeCaller{contract: contract}, nil
}

// NewIL1ERC20BridgeTransactor creates a new write-only instance of IL1ERC20Bridge, bound to a specific deployed contract.
func NewIL1ERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1ERC20BridgeTransactor, error) {
	contract, err := bindIL1ERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20BridgeTransactor{contract: contract}, nil
}

// NewIL1ERC20BridgeFilterer creates a new log filterer instance of IL1ERC20Bridge, bound to a specific deployed contract.
func NewIL1ERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1ERC20BridgeFilterer, error) {
	contract, err := bindIL1ERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20BridgeFilterer{contract: contract}, nil
}

// bindIL1ERC20Bridge binds a generic wrapper to an already deployed contract.
func bindIL1ERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IL1ERC20BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1ERC20Bridge *IL1ERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1ERC20Bridge.Contract.IL1ERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1ERC20Bridge *IL1ERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.IL1ERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1ERC20Bridge *IL1ERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.IL1ERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1ERC20Bridge *IL1ERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1ERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactor) DepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.contract.Transact(opts, "depositERC20", _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.DepositERC20(&_IL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.DepositERC20(&_IL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactor) DepositERC20To(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.contract.Transact(opts, "depositERC20To", _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.DepositERC20To(&_IL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.DepositERC20To(&_IL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// FinalizeBitWithdrawal is a paid mutator transaction binding the contract method 0x839f0ec6.
//
// Solidity: function finalizeBitWithdrawal(address _from, address _to, uint256 _amount, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactor) FinalizeBitWithdrawal(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.contract.Transact(opts, "finalizeBitWithdrawal", _from, _to, _amount, _data)
}

// FinalizeBitWithdrawal is a paid mutator transaction binding the contract method 0x839f0ec6.
//
// Solidity: function finalizeBitWithdrawal(address _from, address _to, uint256 _amount, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) FinalizeBitWithdrawal(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.FinalizeBitWithdrawal(&_IL1ERC20Bridge.TransactOpts, _from, _to, _amount, _data)
}

// FinalizeBitWithdrawal is a paid mutator transaction binding the contract method 0x839f0ec6.
//
// Solidity: function finalizeBitWithdrawal(address _from, address _to, uint256 _amount, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorSession) FinalizeBitWithdrawal(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.FinalizeBitWithdrawal(&_IL1ERC20Bridge.TransactOpts, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactor) FinalizeERC20Withdrawal(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.contract.Transact(opts, "finalizeERC20Withdrawal", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.FinalizeERC20Withdrawal(&_IL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.FinalizeERC20Withdrawal(&_IL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// L2TokenBridge is a paid mutator transaction binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactor) L2TokenBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1ERC20Bridge.contract.Transact(opts, "l2TokenBridge")
}

// L2TokenBridge is a paid mutator transaction binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) L2TokenBridge() (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.L2TokenBridge(&_IL1ERC20Bridge.TransactOpts)
}

// L2TokenBridge is a paid mutator transaction binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorSession) L2TokenBridge() (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.L2TokenBridge(&_IL1ERC20Bridge.TransactOpts)
}

// IL1ERC20BridgeERC20DepositInitiatedIterator is returned from FilterERC20DepositInitiated and is used to iterate over the raw logs and unpacked data for ERC20DepositInitiated events raised by the IL1ERC20Bridge contract.
type IL1ERC20BridgeERC20DepositInitiatedIterator struct {
	Event *IL1ERC20BridgeERC20DepositInitiated // Event containing the contract specifics and raw log

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
func (it *IL1ERC20BridgeERC20DepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ERC20BridgeERC20DepositInitiated)
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
		it.Event = new(IL1ERC20BridgeERC20DepositInitiated)
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
func (it *IL1ERC20BridgeERC20DepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ERC20BridgeERC20DepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ERC20BridgeERC20DepositInitiated represents a ERC20DepositInitiated event raised by the IL1ERC20Bridge contract.
type IL1ERC20BridgeERC20DepositInitiated struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterERC20DepositInitiated is a free log retrieval operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) FilterERC20DepositInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*IL1ERC20BridgeERC20DepositInitiatedIterator, error) {

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

	logs, sub, err := _IL1ERC20Bridge.contract.FilterLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20BridgeERC20DepositInitiatedIterator{contract: _IL1ERC20Bridge.contract, event: "ERC20DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20DepositInitiated is a free log subscription operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) WatchERC20DepositInitiated(opts *bind.WatchOpts, sink chan<- *IL1ERC20BridgeERC20DepositInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IL1ERC20Bridge.contract.WatchLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ERC20BridgeERC20DepositInitiated)
				if err := _IL1ERC20Bridge.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
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

// ParseERC20DepositInitiated is a log parse operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) ParseERC20DepositInitiated(log types.Log) (*IL1ERC20BridgeERC20DepositInitiated, error) {
	event := new(IL1ERC20BridgeERC20DepositInitiated)
	if err := _IL1ERC20Bridge.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ERC20BridgeERC20WithdrawalFinalizedIterator is returned from FilterERC20WithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ERC20WithdrawalFinalized events raised by the IL1ERC20Bridge contract.
type IL1ERC20BridgeERC20WithdrawalFinalizedIterator struct {
	Event *IL1ERC20BridgeERC20WithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *IL1ERC20BridgeERC20WithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ERC20BridgeERC20WithdrawalFinalized)
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
		it.Event = new(IL1ERC20BridgeERC20WithdrawalFinalized)
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
func (it *IL1ERC20BridgeERC20WithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ERC20BridgeERC20WithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ERC20BridgeERC20WithdrawalFinalized represents a ERC20WithdrawalFinalized event raised by the IL1ERC20Bridge contract.
type IL1ERC20BridgeERC20WithdrawalFinalized struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterERC20WithdrawalFinalized is a free log retrieval operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) FilterERC20WithdrawalFinalized(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*IL1ERC20BridgeERC20WithdrawalFinalizedIterator, error) {

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

	logs, sub, err := _IL1ERC20Bridge.contract.FilterLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20BridgeERC20WithdrawalFinalizedIterator{contract: _IL1ERC20Bridge.contract, event: "ERC20WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20WithdrawalFinalized is a free log subscription operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) WatchERC20WithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *IL1ERC20BridgeERC20WithdrawalFinalized, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IL1ERC20Bridge.contract.WatchLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ERC20BridgeERC20WithdrawalFinalized)
				if err := _IL1ERC20Bridge.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
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

// ParseERC20WithdrawalFinalized is a log parse operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) ParseERC20WithdrawalFinalized(log types.Log) (*IL1ERC20BridgeERC20WithdrawalFinalized, error) {
	event := new(IL1ERC20BridgeERC20WithdrawalFinalized)
	if err := _IL1ERC20Bridge.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

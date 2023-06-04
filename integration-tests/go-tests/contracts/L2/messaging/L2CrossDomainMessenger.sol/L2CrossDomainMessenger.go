// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package L2CrossDomainMessenger

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
	_ = abi.ConvertType
)

// L2CrossDomainMessengerMetaData contains all meta data concerning the L2CrossDomainMessenger contract.
var L2CrossDomainMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1CrossDomainMessenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"l1CrossDomainMessenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_messageNonce\",\"type\":\"uint256\"}],\"name\":\"relayMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"relayedMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_gasLimit\",\"type\":\"uint32\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sentMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"successfulMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// L2CrossDomainMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use L2CrossDomainMessengerMetaData.ABI instead.
var L2CrossDomainMessengerABI = L2CrossDomainMessengerMetaData.ABI

// L2CrossDomainMessenger is an auto generated Go binding around an Ethereum contract.
type L2CrossDomainMessenger struct {
	L2CrossDomainMessengerCaller     // Read-only binding to the contract
	L2CrossDomainMessengerTransactor // Write-only binding to the contract
	L2CrossDomainMessengerFilterer   // Log filterer for contract events
}

// L2CrossDomainMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CrossDomainMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CrossDomainMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2CrossDomainMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CrossDomainMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2CrossDomainMessengerSession struct {
	Contract     *L2CrossDomainMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2CrossDomainMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2CrossDomainMessengerCallerSession struct {
	Contract *L2CrossDomainMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// L2CrossDomainMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2CrossDomainMessengerTransactorSession struct {
	Contract     *L2CrossDomainMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L2CrossDomainMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2CrossDomainMessengerRaw struct {
	Contract *L2CrossDomainMessenger // Generic contract binding to access the raw methods on
}

// L2CrossDomainMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerCallerRaw struct {
	Contract *L2CrossDomainMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// L2CrossDomainMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerTransactorRaw struct {
	Contract *L2CrossDomainMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2CrossDomainMessenger creates a new instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessenger(address common.Address, backend bind.ContractBackend) (*L2CrossDomainMessenger, error) {
	contract, err := bindL2CrossDomainMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessenger{L2CrossDomainMessengerCaller: L2CrossDomainMessengerCaller{contract: contract}, L2CrossDomainMessengerTransactor: L2CrossDomainMessengerTransactor{contract: contract}, L2CrossDomainMessengerFilterer: L2CrossDomainMessengerFilterer{contract: contract}}, nil
}

// NewL2CrossDomainMessengerCaller creates a new read-only instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessengerCaller(address common.Address, caller bind.ContractCaller) (*L2CrossDomainMessengerCaller, error) {
	contract, err := bindL2CrossDomainMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerCaller{contract: contract}, nil
}

// NewL2CrossDomainMessengerTransactor creates a new write-only instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*L2CrossDomainMessengerTransactor, error) {
	contract, err := bindL2CrossDomainMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerTransactor{contract: contract}, nil
}

// NewL2CrossDomainMessengerFilterer creates a new log filterer instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*L2CrossDomainMessengerFilterer, error) {
	contract, err := bindL2CrossDomainMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerFilterer{contract: contract}, nil
}

// bindL2CrossDomainMessenger binds a generic wrapper to an already deployed contract.
func bindL2CrossDomainMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2CrossDomainMessenger.Contract.L2CrossDomainMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.L2CrossDomainMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.L2CrossDomainMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2CrossDomainMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.contract.Transact(opts, method, params...)
}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) L1CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "l1CrossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) L1CrossDomainMessenger() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.L1CrossDomainMessenger(&_L2CrossDomainMessenger.CallOpts)
}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) L1CrossDomainMessenger() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.L1CrossDomainMessenger(&_L2CrossDomainMessenger.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) MessageNonce() (*big.Int, error) {
	return _L2CrossDomainMessenger.Contract.MessageNonce(&_L2CrossDomainMessenger.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) MessageNonce() (*big.Int, error) {
	return _L2CrossDomainMessenger.Contract.MessageNonce(&_L2CrossDomainMessenger.CallOpts)
}

// RelayedMessages is a free data retrieval call binding the contract method 0x21d800ec.
//
// Solidity: function relayedMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) RelayedMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "relayedMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RelayedMessages is a free data retrieval call binding the contract method 0x21d800ec.
//
// Solidity: function relayedMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) RelayedMessages(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.RelayedMessages(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// RelayedMessages is a free data retrieval call binding the contract method 0x21d800ec.
//
// Solidity: function relayedMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) RelayedMessages(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.RelayedMessages(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) SentMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "sentMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.SentMessages(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.SentMessages(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// SuccessfulMessages is a free data retrieval call binding the contract method 0xb1b1b209.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) SuccessfulMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "successfulMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SuccessfulMessages is a free data retrieval call binding the contract method 0xb1b1b209.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) SuccessfulMessages(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.SuccessfulMessages(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// SuccessfulMessages is a free data retrieval call binding the contract method 0xb1b1b209.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) SuccessfulMessages(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.SuccessfulMessages(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) XDomainMessageSender() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.XDomainMessageSender(&_L2CrossDomainMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.XDomainMessageSender(&_L2CrossDomainMessenger.CallOpts)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xcbd4ece9.
//
// Solidity: function relayMessage(address _target, address _sender, bytes _message, uint256 _messageNonce) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) RelayMessage(opts *bind.TransactOpts, _target common.Address, _sender common.Address, _message []byte, _messageNonce *big.Int) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "relayMessage", _target, _sender, _message, _messageNonce)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xcbd4ece9.
//
// Solidity: function relayMessage(address _target, address _sender, bytes _message, uint256 _messageNonce) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) RelayMessage(_target common.Address, _sender common.Address, _message []byte, _messageNonce *big.Int) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.RelayMessage(&_L2CrossDomainMessenger.TransactOpts, _target, _sender, _message, _messageNonce)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xcbd4ece9.
//
// Solidity: function relayMessage(address _target, address _sender, bytes _message, uint256 _messageNonce) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) RelayMessage(_target common.Address, _sender common.Address, _message []byte, _messageNonce *big.Int) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.RelayMessage(&_L2CrossDomainMessenger.TransactOpts, _target, _sender, _message, _messageNonce)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _gasLimit) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) SendMessage(opts *bind.TransactOpts, _target common.Address, _message []byte, _gasLimit uint32) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "sendMessage", _target, _message, _gasLimit)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _gasLimit) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) SendMessage(_target common.Address, _message []byte, _gasLimit uint32) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.SendMessage(&_L2CrossDomainMessenger.TransactOpts, _target, _message, _gasLimit)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _gasLimit) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) SendMessage(_target common.Address, _message []byte, _gasLimit uint32) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.SendMessage(&_L2CrossDomainMessenger.TransactOpts, _target, _message, _gasLimit)
}

// L2CrossDomainMessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerFailedRelayedMessageIterator struct {
	Event *L2CrossDomainMessengerFailedRelayedMessage // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerFailedRelayedMessage)
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
		it.Event = new(L2CrossDomainMessengerFailedRelayedMessage)
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
func (it *L2CrossDomainMessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerFailedRelayedMessage struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, msgHash [][32]byte) (*L2CrossDomainMessengerFailedRelayedMessageIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "FailedRelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerFailedRelayedMessageIterator{contract: _L2CrossDomainMessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerFailedRelayedMessage, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "FailedRelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerFailedRelayedMessage)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
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

// ParseFailedRelayedMessage is a log parse operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*L2CrossDomainMessengerFailedRelayedMessage, error) {
	event := new(L2CrossDomainMessengerFailedRelayedMessage)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerRelayedMessageIterator struct {
	Event *L2CrossDomainMessengerRelayedMessage // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerRelayedMessage)
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
		it.Event = new(L2CrossDomainMessengerRelayedMessage)
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
func (it *L2CrossDomainMessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerRelayedMessage represents a RelayedMessage event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerRelayedMessage struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, msgHash [][32]byte) (*L2CrossDomainMessengerRelayedMessageIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "RelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerRelayedMessageIterator{contract: _L2CrossDomainMessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerRelayedMessage, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "RelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerRelayedMessage)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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

// ParseRelayedMessage is a log parse operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseRelayedMessage(log types.Log) (*L2CrossDomainMessengerRelayedMessage, error) {
	event := new(L2CrossDomainMessengerRelayedMessage)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerSentMessageIterator struct {
	Event *L2CrossDomainMessengerSentMessage // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerSentMessage)
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
		it.Event = new(L2CrossDomainMessengerSentMessage)
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
func (it *L2CrossDomainMessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerSentMessage represents a SentMessage event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerSentMessage struct {
	Target       common.Address
	Sender       common.Address
	Message      []byte
	MessageNonce *big.Int
	GasLimit     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentMessage is a free log retrieval operation binding the contract event 0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a.
//
// Solidity: event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, target []common.Address) (*L2CrossDomainMessengerSentMessageIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "SentMessage", targetRule)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerSentMessageIterator{contract: _L2CrossDomainMessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a.
//
// Solidity: event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerSentMessage, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "SentMessage", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerSentMessage)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
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

// ParseSentMessage is a log parse operation binding the contract event 0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a.
//
// Solidity: event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseSentMessage(log types.Log) (*L2CrossDomainMessengerSentMessage, error) {
	event := new(L2CrossDomainMessengerSentMessage)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

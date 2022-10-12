// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BVM_L2ToL1MessagePasser

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

// BVML2ToL1MessagePasserMetaData contains all meta data concerning the BVML2ToL1MessagePasser contract.
var BVML2ToL1MessagePasserMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"L2ToL1Message\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"passMessageToL1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sentMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BVML2ToL1MessagePasserABI is the input ABI used to generate the binding from.
// Deprecated: Use BVML2ToL1MessagePasserMetaData.ABI instead.
var BVML2ToL1MessagePasserABI = BVML2ToL1MessagePasserMetaData.ABI

// BVML2ToL1MessagePasser is an auto generated Go binding around an Ethereum contract.
type BVML2ToL1MessagePasser struct {
	BVML2ToL1MessagePasserCaller     // Read-only binding to the contract
	BVML2ToL1MessagePasserTransactor // Write-only binding to the contract
	BVML2ToL1MessagePasserFilterer   // Log filterer for contract events
}

// BVML2ToL1MessagePasserCaller is an auto generated read-only Go binding around an Ethereum contract.
type BVML2ToL1MessagePasserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVML2ToL1MessagePasserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BVML2ToL1MessagePasserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVML2ToL1MessagePasserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BVML2ToL1MessagePasserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVML2ToL1MessagePasserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BVML2ToL1MessagePasserSession struct {
	Contract     *BVML2ToL1MessagePasser // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BVML2ToL1MessagePasserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BVML2ToL1MessagePasserCallerSession struct {
	Contract *BVML2ToL1MessagePasserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// BVML2ToL1MessagePasserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BVML2ToL1MessagePasserTransactorSession struct {
	Contract     *BVML2ToL1MessagePasserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BVML2ToL1MessagePasserRaw is an auto generated low-level Go binding around an Ethereum contract.
type BVML2ToL1MessagePasserRaw struct {
	Contract *BVML2ToL1MessagePasser // Generic contract binding to access the raw methods on
}

// BVML2ToL1MessagePasserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BVML2ToL1MessagePasserCallerRaw struct {
	Contract *BVML2ToL1MessagePasserCaller // Generic read-only contract binding to access the raw methods on
}

// BVML2ToL1MessagePasserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BVML2ToL1MessagePasserTransactorRaw struct {
	Contract *BVML2ToL1MessagePasserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBVML2ToL1MessagePasser creates a new instance of BVML2ToL1MessagePasser, bound to a specific deployed contract.
func NewBVML2ToL1MessagePasser(address common.Address, backend bind.ContractBackend) (*BVML2ToL1MessagePasser, error) {
	contract, err := bindBVML2ToL1MessagePasser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BVML2ToL1MessagePasser{BVML2ToL1MessagePasserCaller: BVML2ToL1MessagePasserCaller{contract: contract}, BVML2ToL1MessagePasserTransactor: BVML2ToL1MessagePasserTransactor{contract: contract}, BVML2ToL1MessagePasserFilterer: BVML2ToL1MessagePasserFilterer{contract: contract}}, nil
}

// NewBVML2ToL1MessagePasserCaller creates a new read-only instance of BVML2ToL1MessagePasser, bound to a specific deployed contract.
func NewBVML2ToL1MessagePasserCaller(address common.Address, caller bind.ContractCaller) (*BVML2ToL1MessagePasserCaller, error) {
	contract, err := bindBVML2ToL1MessagePasser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BVML2ToL1MessagePasserCaller{contract: contract}, nil
}

// NewBVML2ToL1MessagePasserTransactor creates a new write-only instance of BVML2ToL1MessagePasser, bound to a specific deployed contract.
func NewBVML2ToL1MessagePasserTransactor(address common.Address, transactor bind.ContractTransactor) (*BVML2ToL1MessagePasserTransactor, error) {
	contract, err := bindBVML2ToL1MessagePasser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BVML2ToL1MessagePasserTransactor{contract: contract}, nil
}

// NewBVML2ToL1MessagePasserFilterer creates a new log filterer instance of BVML2ToL1MessagePasser, bound to a specific deployed contract.
func NewBVML2ToL1MessagePasserFilterer(address common.Address, filterer bind.ContractFilterer) (*BVML2ToL1MessagePasserFilterer, error) {
	contract, err := bindBVML2ToL1MessagePasser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BVML2ToL1MessagePasserFilterer{contract: contract}, nil
}

// bindBVML2ToL1MessagePasser binds a generic wrapper to an already deployed contract.
func bindBVML2ToL1MessagePasser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BVML2ToL1MessagePasserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVML2ToL1MessagePasser.Contract.BVML2ToL1MessagePasserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVML2ToL1MessagePasser.Contract.BVML2ToL1MessagePasserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVML2ToL1MessagePasser.Contract.BVML2ToL1MessagePasserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVML2ToL1MessagePasser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVML2ToL1MessagePasser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVML2ToL1MessagePasser.Contract.contract.Transact(opts, method, params...)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserCaller) SentMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _BVML2ToL1MessagePasser.contract.Call(opts, &out, "sentMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _BVML2ToL1MessagePasser.Contract.SentMessages(&_BVML2ToL1MessagePasser.CallOpts, arg0)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserCallerSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _BVML2ToL1MessagePasser.Contract.SentMessages(&_BVML2ToL1MessagePasser.CallOpts, arg0)
}

// PassMessageToL1 is a paid mutator transaction binding the contract method 0xcafa81dc.
//
// Solidity: function passMessageToL1(bytes _message) returns()
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserTransactor) PassMessageToL1(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _BVML2ToL1MessagePasser.contract.Transact(opts, "passMessageToL1", _message)
}

// PassMessageToL1 is a paid mutator transaction binding the contract method 0xcafa81dc.
//
// Solidity: function passMessageToL1(bytes _message) returns()
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserSession) PassMessageToL1(_message []byte) (*types.Transaction, error) {
	return _BVML2ToL1MessagePasser.Contract.PassMessageToL1(&_BVML2ToL1MessagePasser.TransactOpts, _message)
}

// PassMessageToL1 is a paid mutator transaction binding the contract method 0xcafa81dc.
//
// Solidity: function passMessageToL1(bytes _message) returns()
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserTransactorSession) PassMessageToL1(_message []byte) (*types.Transaction, error) {
	return _BVML2ToL1MessagePasser.Contract.PassMessageToL1(&_BVML2ToL1MessagePasser.TransactOpts, _message)
}

// BVML2ToL1MessagePasserL2ToL1MessageIterator is returned from FilterL2ToL1Message and is used to iterate over the raw logs and unpacked data for L2ToL1Message events raised by the BVML2ToL1MessagePasser contract.
type BVML2ToL1MessagePasserL2ToL1MessageIterator struct {
	Event *BVML2ToL1MessagePasserL2ToL1Message // Event containing the contract specifics and raw log

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
func (it *BVML2ToL1MessagePasserL2ToL1MessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVML2ToL1MessagePasserL2ToL1Message)
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
		it.Event = new(BVML2ToL1MessagePasserL2ToL1Message)
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
func (it *BVML2ToL1MessagePasserL2ToL1MessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVML2ToL1MessagePasserL2ToL1MessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVML2ToL1MessagePasserL2ToL1Message represents a L2ToL1Message event raised by the BVML2ToL1MessagePasser contract.
type BVML2ToL1MessagePasserL2ToL1Message struct {
	Nonce  *big.Int
	Sender common.Address
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterL2ToL1Message is a free log retrieval operation binding the contract event 0x47b65c6c9adf9c9a1f4d661cea00e3a0be49b77b90d9b5a02347d55cbfb7c3f5.
//
// Solidity: event L2ToL1Message(uint256 _nonce, address _sender, bytes _data)
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserFilterer) FilterL2ToL1Message(opts *bind.FilterOpts) (*BVML2ToL1MessagePasserL2ToL1MessageIterator, error) {

	logs, sub, err := _BVML2ToL1MessagePasser.contract.FilterLogs(opts, "L2ToL1Message")
	if err != nil {
		return nil, err
	}
	return &BVML2ToL1MessagePasserL2ToL1MessageIterator{contract: _BVML2ToL1MessagePasser.contract, event: "L2ToL1Message", logs: logs, sub: sub}, nil
}

// WatchL2ToL1Message is a free log subscription operation binding the contract event 0x47b65c6c9adf9c9a1f4d661cea00e3a0be49b77b90d9b5a02347d55cbfb7c3f5.
//
// Solidity: event L2ToL1Message(uint256 _nonce, address _sender, bytes _data)
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserFilterer) WatchL2ToL1Message(opts *bind.WatchOpts, sink chan<- *BVML2ToL1MessagePasserL2ToL1Message) (event.Subscription, error) {

	logs, sub, err := _BVML2ToL1MessagePasser.contract.WatchLogs(opts, "L2ToL1Message")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVML2ToL1MessagePasserL2ToL1Message)
				if err := _BVML2ToL1MessagePasser.contract.UnpackLog(event, "L2ToL1Message", log); err != nil {
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

// ParseL2ToL1Message is a log parse operation binding the contract event 0x47b65c6c9adf9c9a1f4d661cea00e3a0be49b77b90d9b5a02347d55cbfb7c3f5.
//
// Solidity: event L2ToL1Message(uint256 _nonce, address _sender, bytes _data)
func (_BVML2ToL1MessagePasser *BVML2ToL1MessagePasserFilterer) ParseL2ToL1Message(log types.Log) (*BVML2ToL1MessagePasserL2ToL1Message, error) {
	event := new(BVML2ToL1MessagePasserL2ToL1Message)
	if err := _BVML2ToL1MessagePasser.contract.UnpackLog(event, "L2ToL1Message", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

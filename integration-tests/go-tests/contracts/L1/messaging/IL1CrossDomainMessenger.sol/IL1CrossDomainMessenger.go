// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IL1CrossDomainMessenger

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

// IL1CrossDomainMessengerL2MessageInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type IL1CrossDomainMessengerL2MessageInclusionProof struct {
	StateRoot            [32]byte
	StateRootBatchHeader LibBVMCodecChainBatchHeader
	StateRootProof       LibBVMCodecChainInclusionProof
	StateTrieWitness     []byte
	StorageTrieWitness   []byte
}

// LibBVMCodecChainBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type LibBVMCodecChainBatchHeader struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	Signature         []byte
	ExtraData         []byte
}

// LibBVMCodecChainInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type LibBVMCodecChainInclusionProof struct {
	Index    *big.Int
	Siblings [][32]byte
}

// IL1CrossDomainMessengerMetaData contains all meta data concerning the IL1CrossDomainMessenger contract.
var IL1CrossDomainMessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_messageNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_BVMCodec.ChainBatchHeader\",\"name\":\"stateRootBatchHeader\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLib_BVMCodec.ChainInclusionProof\",\"name\":\"stateRootProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"stateTrieWitness\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"storageTrieWitness\",\"type\":\"bytes\"}],\"internalType\":\"structIL1CrossDomainMessenger.L2MessageInclusionProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"relayMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_oldGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_newGasLimit\",\"type\":\"uint32\"}],\"name\":\"replayMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_gasLimit\",\"type\":\"uint32\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IL1CrossDomainMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1CrossDomainMessengerMetaData.ABI instead.
var IL1CrossDomainMessengerABI = IL1CrossDomainMessengerMetaData.ABI

// IL1CrossDomainMessenger is an auto generated Go binding around an Ethereum contract.
type IL1CrossDomainMessenger struct {
	IL1CrossDomainMessengerCaller     // Read-only binding to the contract
	IL1CrossDomainMessengerTransactor // Write-only binding to the contract
	IL1CrossDomainMessengerFilterer   // Log filterer for contract events
}

// IL1CrossDomainMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1CrossDomainMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1CrossDomainMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1CrossDomainMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1CrossDomainMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1CrossDomainMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1CrossDomainMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1CrossDomainMessengerSession struct {
	Contract     *IL1CrossDomainMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IL1CrossDomainMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1CrossDomainMessengerCallerSession struct {
	Contract *IL1CrossDomainMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// IL1CrossDomainMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1CrossDomainMessengerTransactorSession struct {
	Contract     *IL1CrossDomainMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// IL1CrossDomainMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1CrossDomainMessengerRaw struct {
	Contract *IL1CrossDomainMessenger // Generic contract binding to access the raw methods on
}

// IL1CrossDomainMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1CrossDomainMessengerCallerRaw struct {
	Contract *IL1CrossDomainMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// IL1CrossDomainMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1CrossDomainMessengerTransactorRaw struct {
	Contract *IL1CrossDomainMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1CrossDomainMessenger creates a new instance of IL1CrossDomainMessenger, bound to a specific deployed contract.
func NewIL1CrossDomainMessenger(address common.Address, backend bind.ContractBackend) (*IL1CrossDomainMessenger, error) {
	contract, err := bindIL1CrossDomainMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1CrossDomainMessenger{IL1CrossDomainMessengerCaller: IL1CrossDomainMessengerCaller{contract: contract}, IL1CrossDomainMessengerTransactor: IL1CrossDomainMessengerTransactor{contract: contract}, IL1CrossDomainMessengerFilterer: IL1CrossDomainMessengerFilterer{contract: contract}}, nil
}

// NewIL1CrossDomainMessengerCaller creates a new read-only instance of IL1CrossDomainMessenger, bound to a specific deployed contract.
func NewIL1CrossDomainMessengerCaller(address common.Address, caller bind.ContractCaller) (*IL1CrossDomainMessengerCaller, error) {
	contract, err := bindIL1CrossDomainMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1CrossDomainMessengerCaller{contract: contract}, nil
}

// NewIL1CrossDomainMessengerTransactor creates a new write-only instance of IL1CrossDomainMessenger, bound to a specific deployed contract.
func NewIL1CrossDomainMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1CrossDomainMessengerTransactor, error) {
	contract, err := bindIL1CrossDomainMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1CrossDomainMessengerTransactor{contract: contract}, nil
}

// NewIL1CrossDomainMessengerFilterer creates a new log filterer instance of IL1CrossDomainMessenger, bound to a specific deployed contract.
func NewIL1CrossDomainMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1CrossDomainMessengerFilterer, error) {
	contract, err := bindIL1CrossDomainMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1CrossDomainMessengerFilterer{contract: contract}, nil
}

// bindIL1CrossDomainMessenger binds a generic wrapper to an already deployed contract.
func bindIL1CrossDomainMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IL1CrossDomainMessengerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1CrossDomainMessenger.Contract.IL1CrossDomainMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.Contract.IL1CrossDomainMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.Contract.IL1CrossDomainMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1CrossDomainMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.Contract.contract.Transact(opts, method, params...)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1CrossDomainMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerSession) XDomainMessageSender() (common.Address, error) {
	return _IL1CrossDomainMessenger.Contract.XDomainMessageSender(&_IL1CrossDomainMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _IL1CrossDomainMessenger.Contract.XDomainMessageSender(&_IL1CrossDomainMessenger.CallOpts)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xd9ce1901.
//
// Solidity: function relayMessage(address _target, address _sender, bytes _message, uint256 _messageNonce, (bytes32,(uint256,bytes32,uint256,uint256,bytes,bytes),(uint256,bytes32[]),bytes,bytes) _proof) returns()
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerTransactor) RelayMessage(opts *bind.TransactOpts, _target common.Address, _sender common.Address, _message []byte, _messageNonce *big.Int, _proof IL1CrossDomainMessengerL2MessageInclusionProof) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.contract.Transact(opts, "relayMessage", _target, _sender, _message, _messageNonce, _proof)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xd9ce1901.
//
// Solidity: function relayMessage(address _target, address _sender, bytes _message, uint256 _messageNonce, (bytes32,(uint256,bytes32,uint256,uint256,bytes,bytes),(uint256,bytes32[]),bytes,bytes) _proof) returns()
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerSession) RelayMessage(_target common.Address, _sender common.Address, _message []byte, _messageNonce *big.Int, _proof IL1CrossDomainMessengerL2MessageInclusionProof) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.Contract.RelayMessage(&_IL1CrossDomainMessenger.TransactOpts, _target, _sender, _message, _messageNonce, _proof)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xd9ce1901.
//
// Solidity: function relayMessage(address _target, address _sender, bytes _message, uint256 _messageNonce, (bytes32,(uint256,bytes32,uint256,uint256,bytes,bytes),(uint256,bytes32[]),bytes,bytes) _proof) returns()
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerTransactorSession) RelayMessage(_target common.Address, _sender common.Address, _message []byte, _messageNonce *big.Int, _proof IL1CrossDomainMessengerL2MessageInclusionProof) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.Contract.RelayMessage(&_IL1CrossDomainMessenger.TransactOpts, _target, _sender, _message, _messageNonce, _proof)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x6f1c8d47.
//
// Solidity: function replayMessage(address _target, address _sender, bytes _message, uint256 _queueIndex, uint32 _oldGasLimit, uint32 _newGasLimit) returns()
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerTransactor) ReplayMessage(opts *bind.TransactOpts, _target common.Address, _sender common.Address, _message []byte, _queueIndex *big.Int, _oldGasLimit uint32, _newGasLimit uint32) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.contract.Transact(opts, "replayMessage", _target, _sender, _message, _queueIndex, _oldGasLimit, _newGasLimit)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x6f1c8d47.
//
// Solidity: function replayMessage(address _target, address _sender, bytes _message, uint256 _queueIndex, uint32 _oldGasLimit, uint32 _newGasLimit) returns()
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerSession) ReplayMessage(_target common.Address, _sender common.Address, _message []byte, _queueIndex *big.Int, _oldGasLimit uint32, _newGasLimit uint32) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.Contract.ReplayMessage(&_IL1CrossDomainMessenger.TransactOpts, _target, _sender, _message, _queueIndex, _oldGasLimit, _newGasLimit)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x6f1c8d47.
//
// Solidity: function replayMessage(address _target, address _sender, bytes _message, uint256 _queueIndex, uint32 _oldGasLimit, uint32 _newGasLimit) returns()
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerTransactorSession) ReplayMessage(_target common.Address, _sender common.Address, _message []byte, _queueIndex *big.Int, _oldGasLimit uint32, _newGasLimit uint32) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.Contract.ReplayMessage(&_IL1CrossDomainMessenger.TransactOpts, _target, _sender, _message, _queueIndex, _oldGasLimit, _newGasLimit)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _gasLimit) returns()
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerTransactor) SendMessage(opts *bind.TransactOpts, _target common.Address, _message []byte, _gasLimit uint32) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.contract.Transact(opts, "sendMessage", _target, _message, _gasLimit)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _gasLimit) returns()
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerSession) SendMessage(_target common.Address, _message []byte, _gasLimit uint32) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.Contract.SendMessage(&_IL1CrossDomainMessenger.TransactOpts, _target, _message, _gasLimit)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _gasLimit) returns()
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerTransactorSession) SendMessage(_target common.Address, _message []byte, _gasLimit uint32) (*types.Transaction, error) {
	return _IL1CrossDomainMessenger.Contract.SendMessage(&_IL1CrossDomainMessenger.TransactOpts, _target, _message, _gasLimit)
}

// IL1CrossDomainMessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the IL1CrossDomainMessenger contract.
type IL1CrossDomainMessengerFailedRelayedMessageIterator struct {
	Event *IL1CrossDomainMessengerFailedRelayedMessage // Event containing the contract specifics and raw log

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
func (it *IL1CrossDomainMessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1CrossDomainMessengerFailedRelayedMessage)
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
		it.Event = new(IL1CrossDomainMessengerFailedRelayedMessage)
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
func (it *IL1CrossDomainMessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1CrossDomainMessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1CrossDomainMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the IL1CrossDomainMessenger contract.
type IL1CrossDomainMessengerFailedRelayedMessage struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, msgHash [][32]byte) (*IL1CrossDomainMessengerFailedRelayedMessageIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _IL1CrossDomainMessenger.contract.FilterLogs(opts, "FailedRelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &IL1CrossDomainMessengerFailedRelayedMessageIterator{contract: _IL1CrossDomainMessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *IL1CrossDomainMessengerFailedRelayedMessage, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _IL1CrossDomainMessenger.contract.WatchLogs(opts, "FailedRelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1CrossDomainMessengerFailedRelayedMessage)
				if err := _IL1CrossDomainMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
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
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*IL1CrossDomainMessengerFailedRelayedMessage, error) {
	event := new(IL1CrossDomainMessengerFailedRelayedMessage)
	if err := _IL1CrossDomainMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1CrossDomainMessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the IL1CrossDomainMessenger contract.
type IL1CrossDomainMessengerRelayedMessageIterator struct {
	Event *IL1CrossDomainMessengerRelayedMessage // Event containing the contract specifics and raw log

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
func (it *IL1CrossDomainMessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1CrossDomainMessengerRelayedMessage)
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
		it.Event = new(IL1CrossDomainMessengerRelayedMessage)
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
func (it *IL1CrossDomainMessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1CrossDomainMessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1CrossDomainMessengerRelayedMessage represents a RelayedMessage event raised by the IL1CrossDomainMessenger contract.
type IL1CrossDomainMessengerRelayedMessage struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, msgHash [][32]byte) (*IL1CrossDomainMessengerRelayedMessageIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _IL1CrossDomainMessenger.contract.FilterLogs(opts, "RelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &IL1CrossDomainMessengerRelayedMessageIterator{contract: _IL1CrossDomainMessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *IL1CrossDomainMessengerRelayedMessage, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _IL1CrossDomainMessenger.contract.WatchLogs(opts, "RelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1CrossDomainMessengerRelayedMessage)
				if err := _IL1CrossDomainMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerFilterer) ParseRelayedMessage(log types.Log) (*IL1CrossDomainMessengerRelayedMessage, error) {
	event := new(IL1CrossDomainMessengerRelayedMessage)
	if err := _IL1CrossDomainMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1CrossDomainMessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the IL1CrossDomainMessenger contract.
type IL1CrossDomainMessengerSentMessageIterator struct {
	Event *IL1CrossDomainMessengerSentMessage // Event containing the contract specifics and raw log

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
func (it *IL1CrossDomainMessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1CrossDomainMessengerSentMessage)
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
		it.Event = new(IL1CrossDomainMessengerSentMessage)
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
func (it *IL1CrossDomainMessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1CrossDomainMessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1CrossDomainMessengerSentMessage represents a SentMessage event raised by the IL1CrossDomainMessenger contract.
type IL1CrossDomainMessengerSentMessage struct {
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
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, target []common.Address) (*IL1CrossDomainMessengerSentMessageIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IL1CrossDomainMessenger.contract.FilterLogs(opts, "SentMessage", targetRule)
	if err != nil {
		return nil, err
	}
	return &IL1CrossDomainMessengerSentMessageIterator{contract: _IL1CrossDomainMessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a.
//
// Solidity: event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *IL1CrossDomainMessengerSentMessage, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IL1CrossDomainMessenger.contract.WatchLogs(opts, "SentMessage", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1CrossDomainMessengerSentMessage)
				if err := _IL1CrossDomainMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
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
func (_IL1CrossDomainMessenger *IL1CrossDomainMessengerFilterer) ParseSentMessage(log types.Log) (*IL1CrossDomainMessengerSentMessage, error) {
	event := new(IL1CrossDomainMessengerSentMessage)
	if err := _IL1CrossDomainMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

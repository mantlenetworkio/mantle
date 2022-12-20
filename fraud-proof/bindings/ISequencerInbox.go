// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ISequencerInboxMetaData contains all meta data concerning the ISequencerInbox contract.
var ISequencerInboxMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTxNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTxNumber\",\"type\":\"uint256\"}],\"name\":\"TxBatchAppended\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"contexts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"txLengths\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"txBatch\",\"type\":\"bytes\"}],\"name\":\"appendTxBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInboxSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"verifyTxInclusion\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ISequencerInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use ISequencerInboxMetaData.ABI instead.
var ISequencerInboxABI = ISequencerInboxMetaData.ABI

// ISequencerInbox is an auto generated Go binding around an Ethereum contract.
type ISequencerInbox struct {
	ISequencerInboxCaller     // Read-only binding to the contract
	ISequencerInboxTransactor // Write-only binding to the contract
	ISequencerInboxFilterer   // Log filterer for contract events
}

// ISequencerInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISequencerInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISequencerInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISequencerInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISequencerInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISequencerInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISequencerInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISequencerInboxSession struct {
	Contract     *ISequencerInbox  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISequencerInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISequencerInboxCallerSession struct {
	Contract *ISequencerInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISequencerInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISequencerInboxTransactorSession struct {
	Contract     *ISequencerInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISequencerInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISequencerInboxRaw struct {
	Contract *ISequencerInbox // Generic contract binding to access the raw methods on
}

// ISequencerInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISequencerInboxCallerRaw struct {
	Contract *ISequencerInboxCaller // Generic read-only contract binding to access the raw methods on
}

// ISequencerInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISequencerInboxTransactorRaw struct {
	Contract *ISequencerInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISequencerInbox creates a new instance of ISequencerInbox, bound to a specific deployed contract.
func NewISequencerInbox(address common.Address, backend bind.ContractBackend) (*ISequencerInbox, error) {
	contract, err := bindISequencerInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISequencerInbox{ISequencerInboxCaller: ISequencerInboxCaller{contract: contract}, ISequencerInboxTransactor: ISequencerInboxTransactor{contract: contract}, ISequencerInboxFilterer: ISequencerInboxFilterer{contract: contract}}, nil
}

// NewISequencerInboxCaller creates a new read-only instance of ISequencerInbox, bound to a specific deployed contract.
func NewISequencerInboxCaller(address common.Address, caller bind.ContractCaller) (*ISequencerInboxCaller, error) {
	contract, err := bindISequencerInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxCaller{contract: contract}, nil
}

// NewISequencerInboxTransactor creates a new write-only instance of ISequencerInbox, bound to a specific deployed contract.
func NewISequencerInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*ISequencerInboxTransactor, error) {
	contract, err := bindISequencerInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxTransactor{contract: contract}, nil
}

// NewISequencerInboxFilterer creates a new log filterer instance of ISequencerInbox, bound to a specific deployed contract.
func NewISequencerInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*ISequencerInboxFilterer, error) {
	contract, err := bindISequencerInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxFilterer{contract: contract}, nil
}

// bindISequencerInbox binds a generic wrapper to an already deployed contract.
func bindISequencerInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISequencerInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISequencerInbox *ISequencerInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISequencerInbox.Contract.ISequencerInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISequencerInbox *ISequencerInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.ISequencerInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISequencerInbox *ISequencerInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.ISequencerInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISequencerInbox *ISequencerInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISequencerInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISequencerInbox *ISequencerInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISequencerInbox *ISequencerInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.contract.Transact(opts, method, params...)
}

// GetInboxSize is a free data retrieval call binding the contract method 0x29869a7f.
//
// Solidity: function getInboxSize() view returns(uint256)
func (_ISequencerInbox *ISequencerInboxCaller) GetInboxSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISequencerInbox.contract.Call(opts, &out, "getInboxSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInboxSize is a free data retrieval call binding the contract method 0x29869a7f.
//
// Solidity: function getInboxSize() view returns(uint256)
func (_ISequencerInbox *ISequencerInboxSession) GetInboxSize() (*big.Int, error) {
	return _ISequencerInbox.Contract.GetInboxSize(&_ISequencerInbox.CallOpts)
}

// GetInboxSize is a free data retrieval call binding the contract method 0x29869a7f.
//
// Solidity: function getInboxSize() view returns(uint256)
func (_ISequencerInbox *ISequencerInboxCallerSession) GetInboxSize() (*big.Int, error) {
	return _ISequencerInbox.Contract.GetInboxSize(&_ISequencerInbox.CallOpts)
}

// VerifyTxInclusion is a free data retrieval call binding the contract method 0xe0faa15c.
//
// Solidity: function verifyTxInclusion(bytes proof) view returns()
func (_ISequencerInbox *ISequencerInboxCaller) VerifyTxInclusion(opts *bind.CallOpts, proof []byte) error {
	var out []interface{}
	err := _ISequencerInbox.contract.Call(opts, &out, "verifyTxInclusion", proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyTxInclusion is a free data retrieval call binding the contract method 0xe0faa15c.
//
// Solidity: function verifyTxInclusion(bytes proof) view returns()
func (_ISequencerInbox *ISequencerInboxSession) VerifyTxInclusion(proof []byte) error {
	return _ISequencerInbox.Contract.VerifyTxInclusion(&_ISequencerInbox.CallOpts, proof)
}

// VerifyTxInclusion is a free data retrieval call binding the contract method 0xe0faa15c.
//
// Solidity: function verifyTxInclusion(bytes proof) view returns()
func (_ISequencerInbox *ISequencerInboxCallerSession) VerifyTxInclusion(proof []byte) error {
	return _ISequencerInbox.Contract.VerifyTxInclusion(&_ISequencerInbox.CallOpts, proof)
}

// AppendTxBatch is a paid mutator transaction binding the contract method 0x66ac980a.
//
// Solidity: function appendTxBatch(uint256[] contexts, uint256[] txLengths, bytes txBatch) returns()
func (_ISequencerInbox *ISequencerInboxTransactor) AppendTxBatch(opts *bind.TransactOpts, contexts []*big.Int, txLengths []*big.Int, txBatch []byte) (*types.Transaction, error) {
	return _ISequencerInbox.contract.Transact(opts, "appendTxBatch", contexts, txLengths, txBatch)
}

// AppendTxBatch is a paid mutator transaction binding the contract method 0x66ac980a.
//
// Solidity: function appendTxBatch(uint256[] contexts, uint256[] txLengths, bytes txBatch) returns()
func (_ISequencerInbox *ISequencerInboxSession) AppendTxBatch(contexts []*big.Int, txLengths []*big.Int, txBatch []byte) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.AppendTxBatch(&_ISequencerInbox.TransactOpts, contexts, txLengths, txBatch)
}

// AppendTxBatch is a paid mutator transaction binding the contract method 0x66ac980a.
//
// Solidity: function appendTxBatch(uint256[] contexts, uint256[] txLengths, bytes txBatch) returns()
func (_ISequencerInbox *ISequencerInboxTransactorSession) AppendTxBatch(contexts []*big.Int, txLengths []*big.Int, txBatch []byte) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.AppendTxBatch(&_ISequencerInbox.TransactOpts, contexts, txLengths, txBatch)
}

// ISequencerInboxTxBatchAppendedIterator is returned from FilterTxBatchAppended and is used to iterate over the raw logs and unpacked data for TxBatchAppended events raised by the ISequencerInbox contract.
type ISequencerInboxTxBatchAppendedIterator struct {
	Event *ISequencerInboxTxBatchAppended // Event containing the contract specifics and raw log

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
func (it *ISequencerInboxTxBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISequencerInboxTxBatchAppended)
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
		it.Event = new(ISequencerInboxTxBatchAppended)
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
func (it *ISequencerInboxTxBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISequencerInboxTxBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISequencerInboxTxBatchAppended represents a TxBatchAppended event raised by the ISequencerInbox contract.
type ISequencerInboxTxBatchAppended struct {
	BatchNumber   *big.Int
	StartTxNumber *big.Int
	EndTxNumber   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTxBatchAppended is a free log retrieval operation binding the contract event 0x0611f4b7c4b81a9158f0c0e9aae0a88c50eaa3efdaee2f87bd796c8c6ef7bffc.
//
// Solidity: event TxBatchAppended(uint256 batchNumber, uint256 startTxNumber, uint256 endTxNumber)
func (_ISequencerInbox *ISequencerInboxFilterer) FilterTxBatchAppended(opts *bind.FilterOpts) (*ISequencerInboxTxBatchAppendedIterator, error) {

	logs, sub, err := _ISequencerInbox.contract.FilterLogs(opts, "TxBatchAppended")
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxTxBatchAppendedIterator{contract: _ISequencerInbox.contract, event: "TxBatchAppended", logs: logs, sub: sub}, nil
}

// WatchTxBatchAppended is a free log subscription operation binding the contract event 0x0611f4b7c4b81a9158f0c0e9aae0a88c50eaa3efdaee2f87bd796c8c6ef7bffc.
//
// Solidity: event TxBatchAppended(uint256 batchNumber, uint256 startTxNumber, uint256 endTxNumber)
func (_ISequencerInbox *ISequencerInboxFilterer) WatchTxBatchAppended(opts *bind.WatchOpts, sink chan<- *ISequencerInboxTxBatchAppended) (event.Subscription, error) {

	logs, sub, err := _ISequencerInbox.contract.WatchLogs(opts, "TxBatchAppended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISequencerInboxTxBatchAppended)
				if err := _ISequencerInbox.contract.UnpackLog(event, "TxBatchAppended", log); err != nil {
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

// ParseTxBatchAppended is a log parse operation binding the contract event 0x0611f4b7c4b81a9158f0c0e9aae0a88c50eaa3efdaee2f87bd796c8c6ef7bffc.
//
// Solidity: event TxBatchAppended(uint256 batchNumber, uint256 startTxNumber, uint256 endTxNumber)
func (_ISequencerInbox *ISequencerInboxFilterer) ParseTxBatchAppended(log types.Log) (*ISequencerInboxTxBatchAppended, error) {
	event := new(ISequencerInboxTxBatchAppended)
	if err := _ISequencerInbox.contract.UnpackLog(event, "TxBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

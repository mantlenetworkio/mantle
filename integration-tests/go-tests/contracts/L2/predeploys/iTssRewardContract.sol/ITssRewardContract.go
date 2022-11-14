// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ITssRewardContract

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

// ITssRewardContractMetaData contains all meta data concerning the ITssRewardContract contract.
var ITssRewardContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tssMembers\",\"type\":\"address[]\"}],\"name\":\"DistributeTssReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockStartHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"length\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tssMembers\",\"type\":\"address[]\"}],\"name\":\"DistributeTssRewardByBlock\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockStartHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_length\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_batchTime\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_tssMembers\",\"type\":\"address[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"updateReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDust\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITssRewardContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ITssRewardContractMetaData.ABI instead.
var ITssRewardContractABI = ITssRewardContractMetaData.ABI

// ITssRewardContract is an auto generated Go binding around an Ethereum contract.
type ITssRewardContract struct {
	ITssRewardContractCaller     // Read-only binding to the contract
	ITssRewardContractTransactor // Write-only binding to the contract
	ITssRewardContractFilterer   // Log filterer for contract events
}

// ITssRewardContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITssRewardContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITssRewardContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITssRewardContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITssRewardContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITssRewardContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITssRewardContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITssRewardContractSession struct {
	Contract     *ITssRewardContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ITssRewardContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITssRewardContractCallerSession struct {
	Contract *ITssRewardContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ITssRewardContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITssRewardContractTransactorSession struct {
	Contract     *ITssRewardContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ITssRewardContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITssRewardContractRaw struct {
	Contract *ITssRewardContract // Generic contract binding to access the raw methods on
}

// ITssRewardContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITssRewardContractCallerRaw struct {
	Contract *ITssRewardContractCaller // Generic read-only contract binding to access the raw methods on
}

// ITssRewardContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITssRewardContractTransactorRaw struct {
	Contract *ITssRewardContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITssRewardContract creates a new instance of ITssRewardContract, bound to a specific deployed contract.
func NewITssRewardContract(address common.Address, backend bind.ContractBackend) (*ITssRewardContract, error) {
	contract, err := bindITssRewardContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITssRewardContract{ITssRewardContractCaller: ITssRewardContractCaller{contract: contract}, ITssRewardContractTransactor: ITssRewardContractTransactor{contract: contract}, ITssRewardContractFilterer: ITssRewardContractFilterer{contract: contract}}, nil
}

// NewITssRewardContractCaller creates a new read-only instance of ITssRewardContract, bound to a specific deployed contract.
func NewITssRewardContractCaller(address common.Address, caller bind.ContractCaller) (*ITssRewardContractCaller, error) {
	contract, err := bindITssRewardContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITssRewardContractCaller{contract: contract}, nil
}

// NewITssRewardContractTransactor creates a new write-only instance of ITssRewardContract, bound to a specific deployed contract.
func NewITssRewardContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ITssRewardContractTransactor, error) {
	contract, err := bindITssRewardContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITssRewardContractTransactor{contract: contract}, nil
}

// NewITssRewardContractFilterer creates a new log filterer instance of ITssRewardContract, bound to a specific deployed contract.
func NewITssRewardContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ITssRewardContractFilterer, error) {
	contract, err := bindITssRewardContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITssRewardContractFilterer{contract: contract}, nil
}

// bindITssRewardContract binds a generic wrapper to an already deployed contract.
func bindITssRewardContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITssRewardContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITssRewardContract *ITssRewardContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITssRewardContract.Contract.ITssRewardContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITssRewardContract *ITssRewardContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITssRewardContract.Contract.ITssRewardContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITssRewardContract *ITssRewardContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITssRewardContract.Contract.ITssRewardContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITssRewardContract *ITssRewardContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITssRewardContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITssRewardContract *ITssRewardContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITssRewardContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITssRewardContract *ITssRewardContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITssRewardContract.Contract.contract.Transact(opts, method, params...)
}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_ITssRewardContract *ITssRewardContractCaller) QueryReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ITssRewardContract.contract.Call(opts, &out, "queryReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_ITssRewardContract *ITssRewardContractSession) QueryReward() (*big.Int, error) {
	return _ITssRewardContract.Contract.QueryReward(&_ITssRewardContract.CallOpts)
}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_ITssRewardContract *ITssRewardContractCallerSession) QueryReward() (*big.Int, error) {
	return _ITssRewardContract.Contract.QueryReward(&_ITssRewardContract.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0fae75d9.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] _tssMembers) returns()
func (_ITssRewardContract *ITssRewardContractTransactor) ClaimReward(opts *bind.TransactOpts, _blockStartHeight *big.Int, _length uint32, _batchTime *big.Int, _tssMembers []common.Address) (*types.Transaction, error) {
	return _ITssRewardContract.contract.Transact(opts, "claimReward", _blockStartHeight, _length, _batchTime, _tssMembers)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0fae75d9.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] _tssMembers) returns()
func (_ITssRewardContract *ITssRewardContractSession) ClaimReward(_blockStartHeight *big.Int, _length uint32, _batchTime *big.Int, _tssMembers []common.Address) (*types.Transaction, error) {
	return _ITssRewardContract.Contract.ClaimReward(&_ITssRewardContract.TransactOpts, _blockStartHeight, _length, _batchTime, _tssMembers)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0fae75d9.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] _tssMembers) returns()
func (_ITssRewardContract *ITssRewardContractTransactorSession) ClaimReward(_blockStartHeight *big.Int, _length uint32, _batchTime *big.Int, _tssMembers []common.Address) (*types.Transaction, error) {
	return _ITssRewardContract.Contract.ClaimReward(&_ITssRewardContract.TransactOpts, _blockStartHeight, _length, _batchTime, _tssMembers)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_ITssRewardContract *ITssRewardContractTransactor) UpdateReward(opts *bind.TransactOpts, _blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ITssRewardContract.contract.Transact(opts, "updateReward", _blockID, _amount)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_ITssRewardContract *ITssRewardContractSession) UpdateReward(_blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ITssRewardContract.Contract.UpdateReward(&_ITssRewardContract.TransactOpts, _blockID, _amount)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_ITssRewardContract *ITssRewardContractTransactorSession) UpdateReward(_blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ITssRewardContract.Contract.UpdateReward(&_ITssRewardContract.TransactOpts, _blockID, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ITssRewardContract *ITssRewardContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITssRewardContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ITssRewardContract *ITssRewardContractSession) Withdraw() (*types.Transaction, error) {
	return _ITssRewardContract.Contract.Withdraw(&_ITssRewardContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ITssRewardContract *ITssRewardContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ITssRewardContract.Contract.Withdraw(&_ITssRewardContract.TransactOpts)
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_ITssRewardContract *ITssRewardContractTransactor) WithdrawDust(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITssRewardContract.contract.Transact(opts, "withdrawDust")
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_ITssRewardContract *ITssRewardContractSession) WithdrawDust() (*types.Transaction, error) {
	return _ITssRewardContract.Contract.WithdrawDust(&_ITssRewardContract.TransactOpts)
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_ITssRewardContract *ITssRewardContractTransactorSession) WithdrawDust() (*types.Transaction, error) {
	return _ITssRewardContract.Contract.WithdrawDust(&_ITssRewardContract.TransactOpts)
}

// ITssRewardContractDistributeTssRewardIterator is returned from FilterDistributeTssReward and is used to iterate over the raw logs and unpacked data for DistributeTssReward events raised by the ITssRewardContract contract.
type ITssRewardContractDistributeTssRewardIterator struct {
	Event *ITssRewardContractDistributeTssReward // Event containing the contract specifics and raw log

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
func (it *ITssRewardContractDistributeTssRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITssRewardContractDistributeTssReward)
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
		it.Event = new(ITssRewardContractDistributeTssReward)
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
func (it *ITssRewardContractDistributeTssRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITssRewardContractDistributeTssRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITssRewardContractDistributeTssReward represents a DistributeTssReward event raised by the ITssRewardContract contract.
type ITssRewardContractDistributeTssReward struct {
	BatchTime  *big.Int
	TssMembers []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDistributeTssReward is a free log retrieval operation binding the contract event 0xf8c88edb629fd6d7636c9252b157475c11ff64f1be0cd034423e9e1046499b00.
//
// Solidity: event DistributeTssReward(uint256 batchTime, address[] tssMembers)
func (_ITssRewardContract *ITssRewardContractFilterer) FilterDistributeTssReward(opts *bind.FilterOpts) (*ITssRewardContractDistributeTssRewardIterator, error) {

	logs, sub, err := _ITssRewardContract.contract.FilterLogs(opts, "DistributeTssReward")
	if err != nil {
		return nil, err
	}
	return &ITssRewardContractDistributeTssRewardIterator{contract: _ITssRewardContract.contract, event: "DistributeTssReward", logs: logs, sub: sub}, nil
}

// WatchDistributeTssReward is a free log subscription operation binding the contract event 0xf8c88edb629fd6d7636c9252b157475c11ff64f1be0cd034423e9e1046499b00.
//
// Solidity: event DistributeTssReward(uint256 batchTime, address[] tssMembers)
func (_ITssRewardContract *ITssRewardContractFilterer) WatchDistributeTssReward(opts *bind.WatchOpts, sink chan<- *ITssRewardContractDistributeTssReward) (event.Subscription, error) {

	logs, sub, err := _ITssRewardContract.contract.WatchLogs(opts, "DistributeTssReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITssRewardContractDistributeTssReward)
				if err := _ITssRewardContract.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
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

// ParseDistributeTssReward is a log parse operation binding the contract event 0xf8c88edb629fd6d7636c9252b157475c11ff64f1be0cd034423e9e1046499b00.
//
// Solidity: event DistributeTssReward(uint256 batchTime, address[] tssMembers)
func (_ITssRewardContract *ITssRewardContractFilterer) ParseDistributeTssReward(log types.Log) (*ITssRewardContractDistributeTssReward, error) {
	event := new(ITssRewardContractDistributeTssReward)
	if err := _ITssRewardContract.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITssRewardContractDistributeTssRewardByBlockIterator is returned from FilterDistributeTssRewardByBlock and is used to iterate over the raw logs and unpacked data for DistributeTssRewardByBlock events raised by the ITssRewardContract contract.
type ITssRewardContractDistributeTssRewardByBlockIterator struct {
	Event *ITssRewardContractDistributeTssRewardByBlock // Event containing the contract specifics and raw log

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
func (it *ITssRewardContractDistributeTssRewardByBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITssRewardContractDistributeTssRewardByBlock)
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
		it.Event = new(ITssRewardContractDistributeTssRewardByBlock)
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
func (it *ITssRewardContractDistributeTssRewardByBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITssRewardContractDistributeTssRewardByBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITssRewardContractDistributeTssRewardByBlock represents a DistributeTssRewardByBlock event raised by the ITssRewardContract contract.
type ITssRewardContractDistributeTssRewardByBlock struct {
	BlockStartHeight *big.Int
	Length           uint32
	TssMembers       []common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributeTssRewardByBlock is a free log retrieval operation binding the contract event 0x417ed5c981c4836fcb057421eaeb9defc15ab95bfadab190ec10e11aecaeeeb9.
//
// Solidity: event DistributeTssRewardByBlock(uint256 blockStartHeight, uint32 length, address[] tssMembers)
func (_ITssRewardContract *ITssRewardContractFilterer) FilterDistributeTssRewardByBlock(opts *bind.FilterOpts) (*ITssRewardContractDistributeTssRewardByBlockIterator, error) {

	logs, sub, err := _ITssRewardContract.contract.FilterLogs(opts, "DistributeTssRewardByBlock")
	if err != nil {
		return nil, err
	}
	return &ITssRewardContractDistributeTssRewardByBlockIterator{contract: _ITssRewardContract.contract, event: "DistributeTssRewardByBlock", logs: logs, sub: sub}, nil
}

// WatchDistributeTssRewardByBlock is a free log subscription operation binding the contract event 0x417ed5c981c4836fcb057421eaeb9defc15ab95bfadab190ec10e11aecaeeeb9.
//
// Solidity: event DistributeTssRewardByBlock(uint256 blockStartHeight, uint32 length, address[] tssMembers)
func (_ITssRewardContract *ITssRewardContractFilterer) WatchDistributeTssRewardByBlock(opts *bind.WatchOpts, sink chan<- *ITssRewardContractDistributeTssRewardByBlock) (event.Subscription, error) {

	logs, sub, err := _ITssRewardContract.contract.WatchLogs(opts, "DistributeTssRewardByBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITssRewardContractDistributeTssRewardByBlock)
				if err := _ITssRewardContract.contract.UnpackLog(event, "DistributeTssRewardByBlock", log); err != nil {
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

// ParseDistributeTssRewardByBlock is a log parse operation binding the contract event 0x417ed5c981c4836fcb057421eaeb9defc15ab95bfadab190ec10e11aecaeeeb9.
//
// Solidity: event DistributeTssRewardByBlock(uint256 blockStartHeight, uint32 length, address[] tssMembers)
func (_ITssRewardContract *ITssRewardContractFilterer) ParseDistributeTssRewardByBlock(log types.Log) (*ITssRewardContractDistributeTssRewardByBlock, error) {
	event := new(ITssRewardContractDistributeTssRewardByBlock)
	if err := _ITssRewardContract.contract.UnpackLog(event, "DistributeTssRewardByBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

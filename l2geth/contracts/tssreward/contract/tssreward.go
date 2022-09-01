// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deadAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockStartHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tssMembers\",\"type\":\"address[]\"}],\"name\":\"DistributeTssReward\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bestBlockID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockStartHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_length\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"_tssMembers\",\"type\":\"address[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dust\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"updateReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDust\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// BestBlockID is a free data retrieval call binding the contract method 0x19d509a1.
//
// Solidity: function bestBlockID() view returns(uint256)
func (_Contract *ContractCaller) BestBlockID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "bestBlockID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BestBlockID is a free data retrieval call binding the contract method 0x19d509a1.
//
// Solidity: function bestBlockID() view returns(uint256)
func (_Contract *ContractSession) BestBlockID() (*big.Int, error) {
	return _Contract.Contract.BestBlockID(&_Contract.CallOpts)
}

// BestBlockID is a free data retrieval call binding the contract method 0x19d509a1.
//
// Solidity: function bestBlockID() view returns(uint256)
func (_Contract *ContractCallerSession) BestBlockID() (*big.Int, error) {
	return _Contract.Contract.BestBlockID(&_Contract.CallOpts)
}

// DeadAddress is a free data retrieval call binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() view returns(address)
func (_Contract *ContractCaller) DeadAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "deadAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeadAddress is a free data retrieval call binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() view returns(address)
func (_Contract *ContractSession) DeadAddress() (common.Address, error) {
	return _Contract.Contract.DeadAddress(&_Contract.CallOpts)
}

// DeadAddress is a free data retrieval call binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() view returns(address)
func (_Contract *ContractCallerSession) DeadAddress() (common.Address, error) {
	return _Contract.Contract.DeadAddress(&_Contract.CallOpts)
}

// Dust is a free data retrieval call binding the contract method 0xfad9aba3.
//
// Solidity: function dust() view returns(uint256)
func (_Contract *ContractCaller) Dust(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dust")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dust is a free data retrieval call binding the contract method 0xfad9aba3.
//
// Solidity: function dust() view returns(uint256)
func (_Contract *ContractSession) Dust() (*big.Int, error) {
	return _Contract.Contract.Dust(&_Contract.CallOpts)
}

// Dust is a free data retrieval call binding the contract method 0xfad9aba3.
//
// Solidity: function dust() view returns(uint256)
func (_Contract *ContractCallerSession) Dust() (*big.Int, error) {
	return _Contract.Contract.Dust(&_Contract.CallOpts)
}

// Ledger is a free data retrieval call binding the contract method 0x10a7fd7b.
//
// Solidity: function ledger(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) Ledger(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ledger", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ledger is a free data retrieval call binding the contract method 0x10a7fd7b.
//
// Solidity: function ledger(uint256 ) view returns(uint256)
func (_Contract *ContractSession) Ledger(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.Ledger(&_Contract.CallOpts, arg0)
}

// Ledger is a free data retrieval call binding the contract method 0x10a7fd7b.
//
// Solidity: function ledger(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) Ledger(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.Ledger(&_Contract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_Contract *ContractCaller) QueryReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "queryReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_Contract *ContractSession) QueryReward() (*big.Int, error) {
	return _Contract.Contract.QueryReward(&_Contract.CallOpts)
}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_Contract *ContractCallerSession) QueryReward() (*big.Int, error) {
	return _Contract.Contract.QueryReward(&_Contract.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_Contract *ContractCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_Contract *ContractSession) TotalAmount() (*big.Int, error) {
	return _Contract.Contract.TotalAmount(&_Contract.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_Contract *ContractCallerSession) TotalAmount() (*big.Int, error) {
	return _Contract.Contract.TotalAmount(&_Contract.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xe04f6e35.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, address[] _tssMembers) returns()
func (_Contract *ContractTransactor) ClaimReward(opts *bind.TransactOpts, _blockStartHeight *big.Int, _length uint32, _tssMembers []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimReward", _blockStartHeight, _length, _tssMembers)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xe04f6e35.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, address[] _tssMembers) returns()
func (_Contract *ContractSession) ClaimReward(_blockStartHeight *big.Int, _length uint32, _tssMembers []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ClaimReward(&_Contract.TransactOpts, _blockStartHeight, _length, _tssMembers)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xe04f6e35.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, address[] _tssMembers) returns()
func (_Contract *ContractTransactorSession) ClaimReward(_blockStartHeight *big.Int, _length uint32, _tssMembers []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ClaimReward(&_Contract.TransactOpts, _blockStartHeight, _length, _tssMembers)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_Contract *ContractTransactor) UpdateReward(opts *bind.TransactOpts, _blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateReward", _blockID, _amount)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_Contract *ContractSession) UpdateReward(_blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateReward(&_Contract.TransactOpts, _blockID, _amount)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_Contract *ContractTransactorSession) UpdateReward(_blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateReward(&_Contract.TransactOpts, _blockID, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_Contract *ContractTransactor) WithdrawDust(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawDust")
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_Contract *ContractSession) WithdrawDust() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawDust(&_Contract.TransactOpts)
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_Contract *ContractTransactorSession) WithdrawDust() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawDust(&_Contract.TransactOpts)
}

// ContractDistributeTssRewardIterator is returned from FilterDistributeTssReward and is used to iterate over the raw logs and unpacked data for DistributeTssReward events raised by the Contract contract.
type ContractDistributeTssRewardIterator struct {
	Event *ContractDistributeTssReward // Event containing the contract specifics and raw log

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
func (it *ContractDistributeTssRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDistributeTssReward)
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
		it.Event = new(ContractDistributeTssReward)
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
func (it *ContractDistributeTssRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDistributeTssRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDistributeTssReward represents a DistributeTssReward event raised by the Contract contract.
type ContractDistributeTssReward struct {
	BlockStartHeight *big.Int
	Length           *big.Int
	TssMembers       []common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributeTssReward is a free log retrieval operation binding the contract event 0xf630cba6d450d736e85735388d4fe67a177b8a3685cdd7dee2bea7727b47860a.
//
// Solidity: event DistributeTssReward(uint256 blockStartHeight, uint256 length, address[] tssMembers)
func (_Contract *ContractFilterer) FilterDistributeTssReward(opts *bind.FilterOpts) (*ContractDistributeTssRewardIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DistributeTssReward")
	if err != nil {
		return nil, err
	}
	return &ContractDistributeTssRewardIterator{contract: _Contract.contract, event: "DistributeTssReward", logs: logs, sub: sub}, nil
}

// WatchDistributeTssReward is a free log subscription operation binding the contract event 0xf630cba6d450d736e85735388d4fe67a177b8a3685cdd7dee2bea7727b47860a.
//
// Solidity: event DistributeTssReward(uint256 blockStartHeight, uint256 length, address[] tssMembers)
func (_Contract *ContractFilterer) WatchDistributeTssReward(opts *bind.WatchOpts, sink chan<- *ContractDistributeTssReward) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DistributeTssReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDistributeTssReward)
				if err := _Contract.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
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

// ParseDistributeTssReward is a log parse operation binding the contract event 0xf630cba6d450d736e85735388d4fe67a177b8a3685cdd7dee2bea7727b47860a.
//
// Solidity: event DistributeTssReward(uint256 blockStartHeight, uint256 length, address[] tssMembers)
func (_Contract *ContractFilterer) ParseDistributeTssReward(log types.Log) (*ContractDistributeTssReward, error) {
	event := new(ContractDistributeTssReward)
	if err := _Contract.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

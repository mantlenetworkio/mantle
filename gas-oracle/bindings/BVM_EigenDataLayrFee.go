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

// BVMEigenDataLayrFeeMetaData contains all meta data concerning the BVMEigenDataLayrFee contract.
var BVMEigenDataLayrFeeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userRollupFee\",\"type\":\"uint256\"}],\"name\":\"RollupFeeHistory\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gasFeeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userRollupFee\",\"type\":\"uint256\"}],\"name\":\"setRollupFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BVMEigenDataLayrFeeABI is the input ABI used to generate the binding from.
// Deprecated: Use BVMEigenDataLayrFeeMetaData.ABI instead.
var BVMEigenDataLayrFeeABI = BVMEigenDataLayrFeeMetaData.ABI

// BVMEigenDataLayrFee is an auto generated Go binding around an Ethereum contract.
type BVMEigenDataLayrFee struct {
	BVMEigenDataLayrFeeCaller     // Read-only binding to the contract
	BVMEigenDataLayrFeeTransactor // Write-only binding to the contract
	BVMEigenDataLayrFeeFilterer   // Log filterer for contract events
}

// BVMEigenDataLayrFeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BVMEigenDataLayrFeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMEigenDataLayrFeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BVMEigenDataLayrFeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMEigenDataLayrFeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BVMEigenDataLayrFeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMEigenDataLayrFeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BVMEigenDataLayrFeeSession struct {
	Contract     *BVMEigenDataLayrFee // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BVMEigenDataLayrFeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BVMEigenDataLayrFeeCallerSession struct {
	Contract *BVMEigenDataLayrFeeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BVMEigenDataLayrFeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BVMEigenDataLayrFeeTransactorSession struct {
	Contract     *BVMEigenDataLayrFeeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BVMEigenDataLayrFeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BVMEigenDataLayrFeeRaw struct {
	Contract *BVMEigenDataLayrFee // Generic contract binding to access the raw methods on
}

// BVMEigenDataLayrFeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BVMEigenDataLayrFeeCallerRaw struct {
	Contract *BVMEigenDataLayrFeeCaller // Generic read-only contract binding to access the raw methods on
}

// BVMEigenDataLayrFeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BVMEigenDataLayrFeeTransactorRaw struct {
	Contract *BVMEigenDataLayrFeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBVMEigenDataLayrFee creates a new instance of BVMEigenDataLayrFee, bound to a specific deployed contract.
func NewBVMEigenDataLayrFee(address common.Address, backend bind.ContractBackend) (*BVMEigenDataLayrFee, error) {
	contract, err := bindBVMEigenDataLayrFee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrFee{BVMEigenDataLayrFeeCaller: BVMEigenDataLayrFeeCaller{contract: contract}, BVMEigenDataLayrFeeTransactor: BVMEigenDataLayrFeeTransactor{contract: contract}, BVMEigenDataLayrFeeFilterer: BVMEigenDataLayrFeeFilterer{contract: contract}}, nil
}

// NewBVMEigenDataLayrFeeCaller creates a new read-only instance of BVMEigenDataLayrFee, bound to a specific deployed contract.
func NewBVMEigenDataLayrFeeCaller(address common.Address, caller bind.ContractCaller) (*BVMEigenDataLayrFeeCaller, error) {
	contract, err := bindBVMEigenDataLayrFee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrFeeCaller{contract: contract}, nil
}

// NewBVMEigenDataLayrFeeTransactor creates a new write-only instance of BVMEigenDataLayrFee, bound to a specific deployed contract.
func NewBVMEigenDataLayrFeeTransactor(address common.Address, transactor bind.ContractTransactor) (*BVMEigenDataLayrFeeTransactor, error) {
	contract, err := bindBVMEigenDataLayrFee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrFeeTransactor{contract: contract}, nil
}

// NewBVMEigenDataLayrFeeFilterer creates a new log filterer instance of BVMEigenDataLayrFee, bound to a specific deployed contract.
func NewBVMEigenDataLayrFeeFilterer(address common.Address, filterer bind.ContractFilterer) (*BVMEigenDataLayrFeeFilterer, error) {
	contract, err := bindBVMEigenDataLayrFee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrFeeFilterer{contract: contract}, nil
}

// bindBVMEigenDataLayrFee binds a generic wrapper to an already deployed contract.
func bindBVMEigenDataLayrFee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BVMEigenDataLayrFeeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMEigenDataLayrFee.Contract.BVMEigenDataLayrFeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.BVMEigenDataLayrFeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.BVMEigenDataLayrFeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMEigenDataLayrFee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.contract.Transact(opts, method, params...)
}

// GasFeeAddress is a free data retrieval call binding the contract method 0xf7d8f616.
//
// Solidity: function gasFeeAddress() view returns(address)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeCaller) GasFeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMEigenDataLayrFee.contract.Call(opts, &out, "gasFeeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasFeeAddress is a free data retrieval call binding the contract method 0xf7d8f616.
//
// Solidity: function gasFeeAddress() view returns(address)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeSession) GasFeeAddress() (common.Address, error) {
	return _BVMEigenDataLayrFee.Contract.GasFeeAddress(&_BVMEigenDataLayrFee.CallOpts)
}

// GasFeeAddress is a free data retrieval call binding the contract method 0xf7d8f616.
//
// Solidity: function gasFeeAddress() view returns(address)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeCallerSession) GasFeeAddress() (common.Address, error) {
	return _BVMEigenDataLayrFee.Contract.GasFeeAddress(&_BVMEigenDataLayrFee.CallOpts)
}

// GetRollupFee is a free data retrieval call binding the contract method 0xced808e7.
//
// Solidity: function getRollupFee() view returns(uint256)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeCaller) GetRollupFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrFee.contract.Call(opts, &out, "getRollupFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRollupFee is a free data retrieval call binding the contract method 0xced808e7.
//
// Solidity: function getRollupFee() view returns(uint256)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeSession) GetRollupFee() (*big.Int, error) {
	return _BVMEigenDataLayrFee.Contract.GetRollupFee(&_BVMEigenDataLayrFee.CallOpts)
}

// GetRollupFee is a free data retrieval call binding the contract method 0xced808e7.
//
// Solidity: function getRollupFee() view returns(uint256)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeCallerSession) GetRollupFee() (*big.Int, error) {
	return _BVMEigenDataLayrFee.Contract.GetRollupFee(&_BVMEigenDataLayrFee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMEigenDataLayrFee.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeSession) Owner() (common.Address, error) {
	return _BVMEigenDataLayrFee.Contract.Owner(&_BVMEigenDataLayrFee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeCallerSession) Owner() (common.Address, error) {
	return _BVMEigenDataLayrFee.Contract.Owner(&_BVMEigenDataLayrFee.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _address) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactor) Initialize(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.contract.Transact(opts, "initialize", _address)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _address) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeSession) Initialize(_address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.Initialize(&_BVMEigenDataLayrFee.TransactOpts, _address)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _address) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactorSession) Initialize(_address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.Initialize(&_BVMEigenDataLayrFee.TransactOpts, _address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeSession) RenounceOwnership() (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.RenounceOwnership(&_BVMEigenDataLayrFee.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.RenounceOwnership(&_BVMEigenDataLayrFee.TransactOpts)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _address) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactor) SetFeeAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.contract.Transact(opts, "setFeeAddress", _address)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _address) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeSession) SetFeeAddress(_address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.SetFeeAddress(&_BVMEigenDataLayrFee.TransactOpts, _address)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _address) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactorSession) SetFeeAddress(_address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.SetFeeAddress(&_BVMEigenDataLayrFee.TransactOpts, _address)
}

// SetRollupFee is a paid mutator transaction binding the contract method 0x66ab598b.
//
// Solidity: function setRollupFee(uint256 _l2Block, uint256 _userRollupFee) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactor) SetRollupFee(opts *bind.TransactOpts, _l2Block *big.Int, _userRollupFee *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.contract.Transact(opts, "setRollupFee", _l2Block, _userRollupFee)
}

// SetRollupFee is a paid mutator transaction binding the contract method 0x66ab598b.
//
// Solidity: function setRollupFee(uint256 _l2Block, uint256 _userRollupFee) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeSession) SetRollupFee(_l2Block *big.Int, _userRollupFee *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.SetRollupFee(&_BVMEigenDataLayrFee.TransactOpts, _l2Block, _userRollupFee)
}

// SetRollupFee is a paid mutator transaction binding the contract method 0x66ab598b.
//
// Solidity: function setRollupFee(uint256 _l2Block, uint256 _userRollupFee) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactorSession) SetRollupFee(_l2Block *big.Int, _userRollupFee *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.SetRollupFee(&_BVMEigenDataLayrFee.TransactOpts, _l2Block, _userRollupFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.TransferOwnership(&_BVMEigenDataLayrFee.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrFee.Contract.TransferOwnership(&_BVMEigenDataLayrFee.TransactOpts, newOwner)
}

// BVMEigenDataLayrFeeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BVMEigenDataLayrFee contract.
type BVMEigenDataLayrFeeInitializedIterator struct {
	Event *BVMEigenDataLayrFeeInitialized // Event containing the contract specifics and raw log

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
func (it *BVMEigenDataLayrFeeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMEigenDataLayrFeeInitialized)
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
		it.Event = new(BVMEigenDataLayrFeeInitialized)
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
func (it *BVMEigenDataLayrFeeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMEigenDataLayrFeeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMEigenDataLayrFeeInitialized represents a Initialized event raised by the BVMEigenDataLayrFee contract.
type BVMEigenDataLayrFeeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BVMEigenDataLayrFeeInitializedIterator, error) {

	logs, sub, err := _BVMEigenDataLayrFee.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrFeeInitializedIterator{contract: _BVMEigenDataLayrFee.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BVMEigenDataLayrFeeInitialized) (event.Subscription, error) {

	logs, sub, err := _BVMEigenDataLayrFee.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMEigenDataLayrFeeInitialized)
				if err := _BVMEigenDataLayrFee.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeFilterer) ParseInitialized(log types.Log) (*BVMEigenDataLayrFeeInitialized, error) {
	event := new(BVMEigenDataLayrFeeInitialized)
	if err := _BVMEigenDataLayrFee.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMEigenDataLayrFeeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BVMEigenDataLayrFee contract.
type BVMEigenDataLayrFeeOwnershipTransferredIterator struct {
	Event *BVMEigenDataLayrFeeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BVMEigenDataLayrFeeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMEigenDataLayrFeeOwnershipTransferred)
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
		it.Event = new(BVMEigenDataLayrFeeOwnershipTransferred)
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
func (it *BVMEigenDataLayrFeeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMEigenDataLayrFeeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMEigenDataLayrFeeOwnershipTransferred represents a OwnershipTransferred event raised by the BVMEigenDataLayrFee contract.
type BVMEigenDataLayrFeeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BVMEigenDataLayrFeeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BVMEigenDataLayrFee.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrFeeOwnershipTransferredIterator{contract: _BVMEigenDataLayrFee.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BVMEigenDataLayrFeeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BVMEigenDataLayrFee.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMEigenDataLayrFeeOwnershipTransferred)
				if err := _BVMEigenDataLayrFee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeFilterer) ParseOwnershipTransferred(log types.Log) (*BVMEigenDataLayrFeeOwnershipTransferred, error) {
	event := new(BVMEigenDataLayrFeeOwnershipTransferred)
	if err := _BVMEigenDataLayrFee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMEigenDataLayrFeeRollupFeeHistoryIterator is returned from FilterRollupFeeHistory and is used to iterate over the raw logs and unpacked data for RollupFeeHistory events raised by the BVMEigenDataLayrFee contract.
type BVMEigenDataLayrFeeRollupFeeHistoryIterator struct {
	Event *BVMEigenDataLayrFeeRollupFeeHistory // Event containing the contract specifics and raw log

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
func (it *BVMEigenDataLayrFeeRollupFeeHistoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMEigenDataLayrFeeRollupFeeHistory)
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
		it.Event = new(BVMEigenDataLayrFeeRollupFeeHistory)
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
func (it *BVMEigenDataLayrFeeRollupFeeHistoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMEigenDataLayrFeeRollupFeeHistoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMEigenDataLayrFeeRollupFeeHistory represents a RollupFeeHistory event raised by the BVMEigenDataLayrFee contract.
type BVMEigenDataLayrFeeRollupFeeHistory struct {
	L2Block       *big.Int
	UserRollupFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRollupFeeHistory is a free log retrieval operation binding the contract event 0x56b7c04de50cf9d2c6d1c31bff463ab5fc4f80431d2a1292062f73a6de2bbf71.
//
// Solidity: event RollupFeeHistory(uint256 l2Block, uint256 userRollupFee)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeFilterer) FilterRollupFeeHistory(opts *bind.FilterOpts) (*BVMEigenDataLayrFeeRollupFeeHistoryIterator, error) {

	logs, sub, err := _BVMEigenDataLayrFee.contract.FilterLogs(opts, "RollupFeeHistory")
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrFeeRollupFeeHistoryIterator{contract: _BVMEigenDataLayrFee.contract, event: "RollupFeeHistory", logs: logs, sub: sub}, nil
}

// WatchRollupFeeHistory is a free log subscription operation binding the contract event 0x56b7c04de50cf9d2c6d1c31bff463ab5fc4f80431d2a1292062f73a6de2bbf71.
//
// Solidity: event RollupFeeHistory(uint256 l2Block, uint256 userRollupFee)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeFilterer) WatchRollupFeeHistory(opts *bind.WatchOpts, sink chan<- *BVMEigenDataLayrFeeRollupFeeHistory) (event.Subscription, error) {

	logs, sub, err := _BVMEigenDataLayrFee.contract.WatchLogs(opts, "RollupFeeHistory")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMEigenDataLayrFeeRollupFeeHistory)
				if err := _BVMEigenDataLayrFee.contract.UnpackLog(event, "RollupFeeHistory", log); err != nil {
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

// ParseRollupFeeHistory is a log parse operation binding the contract event 0x56b7c04de50cf9d2c6d1c31bff463ab5fc4f80431d2a1292062f73a6de2bbf71.
//
// Solidity: event RollupFeeHistory(uint256 l2Block, uint256 userRollupFee)
func (_BVMEigenDataLayrFee *BVMEigenDataLayrFeeFilterer) ParseRollupFeeHistory(log types.Log) (*BVMEigenDataLayrFeeRollupFeeHistory, error) {
	event := new(BVMEigenDataLayrFeeRollupFeeHistory)
	if err := _BVMEigenDataLayrFee.contract.UnpackLog(event, "RollupFeeHistory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// IChallengeMetaData contains all meta data concerning the IChallenge contract.
var IChallengeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challengeState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIChallenge.CompletionReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"ChallengeCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"bisection\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengedSegmentIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"prevBisection\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"prevChallengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevChallengedSegmentLength\",\"type\":\"uint256\"}],\"name\":\"bisectExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentResponderTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierEntry\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resultReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_startStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_endStateHash\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numSteps\",\"type\":\"uint256\"}],\"name\":\"initializeChallengeLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"challengedStepIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"prevBisection\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"prevChallengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevChallengedSegmentLength\",\"type\":\"uint256\"}],\"name\":\"verifyOneStepProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IChallengeABI is the input ABI used to generate the binding from.
// Deprecated: Use IChallengeMetaData.ABI instead.
var IChallengeABI = IChallengeMetaData.ABI

// IChallenge is an auto generated Go binding around an Ethereum contract.
type IChallenge struct {
	IChallengeCaller     // Read-only binding to the contract
	IChallengeTransactor // Write-only binding to the contract
	IChallengeFilterer   // Log filterer for contract events
}

// IChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeSession struct {
	Contract     *IChallenge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeCallerSession struct {
	Contract *IChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeTransactorSession struct {
	Contract     *IChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeRaw struct {
	Contract *IChallenge // Generic contract binding to access the raw methods on
}

// IChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeCallerRaw struct {
	Contract *IChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeTransactorRaw struct {
	Contract *IChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallenge creates a new instance of IChallenge, bound to a specific deployed contract.
func NewIChallenge(address common.Address, backend bind.ContractBackend) (*IChallenge, error) {
	contract, err := bindIChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallenge{IChallengeCaller: IChallengeCaller{contract: contract}, IChallengeTransactor: IChallengeTransactor{contract: contract}, IChallengeFilterer: IChallengeFilterer{contract: contract}}, nil
}

// NewIChallengeCaller creates a new read-only instance of IChallenge, bound to a specific deployed contract.
func NewIChallengeCaller(address common.Address, caller bind.ContractCaller) (*IChallengeCaller, error) {
	contract, err := bindIChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeCaller{contract: contract}, nil
}

// NewIChallengeTransactor creates a new write-only instance of IChallenge, bound to a specific deployed contract.
func NewIChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeTransactor, error) {
	contract, err := bindIChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeTransactor{contract: contract}, nil
}

// NewIChallengeFilterer creates a new log filterer instance of IChallenge, bound to a specific deployed contract.
func NewIChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeFilterer, error) {
	contract, err := bindIChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeFilterer{contract: contract}, nil
}

// bindIChallenge binds a generic wrapper to an already deployed contract.
func bindIChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallenge *IChallengeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallenge.Contract.IChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallenge *IChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallenge.Contract.IChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallenge *IChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallenge.Contract.IChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallenge *IChallengeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallenge *IChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallenge *IChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallenge.Contract.contract.Transact(opts, method, params...)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x8a8cd218.
//
// Solidity: function currentResponder() view returns(address)
func (_IChallenge *IChallengeCaller) CurrentResponder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IChallenge.contract.Call(opts, &out, "currentResponder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x8a8cd218.
//
// Solidity: function currentResponder() view returns(address)
func (_IChallenge *IChallengeSession) CurrentResponder() (common.Address, error) {
	return _IChallenge.Contract.CurrentResponder(&_IChallenge.CallOpts)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x8a8cd218.
//
// Solidity: function currentResponder() view returns(address)
func (_IChallenge *IChallengeCallerSession) CurrentResponder() (common.Address, error) {
	return _IChallenge.Contract.CurrentResponder(&_IChallenge.CallOpts)
}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xe87e3589.
//
// Solidity: function currentResponderTimeLeft() view returns(uint256)
func (_IChallenge *IChallengeCaller) CurrentResponderTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IChallenge.contract.Call(opts, &out, "currentResponderTimeLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xe87e3589.
//
// Solidity: function currentResponderTimeLeft() view returns(uint256)
func (_IChallenge *IChallengeSession) CurrentResponderTimeLeft() (*big.Int, error) {
	return _IChallenge.Contract.CurrentResponderTimeLeft(&_IChallenge.CallOpts)
}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xe87e3589.
//
// Solidity: function currentResponderTimeLeft() view returns(uint256)
func (_IChallenge *IChallengeCallerSession) CurrentResponderTimeLeft() (*big.Int, error) {
	return _IChallenge.Contract.CurrentResponderTimeLeft(&_IChallenge.CallOpts)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xcc8f6677.
//
// Solidity: function bisectExecution(bytes32[] bisection, uint256 challengedSegmentIndex, bytes32[] prevBisection, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_IChallenge *IChallengeTransactor) BisectExecution(opts *bind.TransactOpts, bisection [][32]byte, challengedSegmentIndex *big.Int, prevBisection [][32]byte, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _IChallenge.contract.Transact(opts, "bisectExecution", bisection, challengedSegmentIndex, prevBisection, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xcc8f6677.
//
// Solidity: function bisectExecution(bytes32[] bisection, uint256 challengedSegmentIndex, bytes32[] prevBisection, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_IChallenge *IChallengeSession) BisectExecution(bisection [][32]byte, challengedSegmentIndex *big.Int, prevBisection [][32]byte, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _IChallenge.Contract.BisectExecution(&_IChallenge.TransactOpts, bisection, challengedSegmentIndex, prevBisection, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xcc8f6677.
//
// Solidity: function bisectExecution(bytes32[] bisection, uint256 challengedSegmentIndex, bytes32[] prevBisection, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_IChallenge *IChallengeTransactorSession) BisectExecution(bisection [][32]byte, challengedSegmentIndex *big.Int, prevBisection [][32]byte, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _IChallenge.Contract.BisectExecution(&_IChallenge.TransactOpts, bisection, challengedSegmentIndex, prevBisection, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// Initialize is a paid mutator transaction binding the contract method 0xf28c3e37.
//
// Solidity: function initialize(address _defender, address _challenger, address _verifier, address _resultReceiver, bytes32 _startStateHash, bytes32 _endStateHash) returns()
func (_IChallenge *IChallengeTransactor) Initialize(opts *bind.TransactOpts, _defender common.Address, _challenger common.Address, _verifier common.Address, _resultReceiver common.Address, _startStateHash [32]byte, _endStateHash [32]byte) (*types.Transaction, error) {
	return _IChallenge.contract.Transact(opts, "initialize", _defender, _challenger, _verifier, _resultReceiver, _startStateHash, _endStateHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xf28c3e37.
//
// Solidity: function initialize(address _defender, address _challenger, address _verifier, address _resultReceiver, bytes32 _startStateHash, bytes32 _endStateHash) returns()
func (_IChallenge *IChallengeSession) Initialize(_defender common.Address, _challenger common.Address, _verifier common.Address, _resultReceiver common.Address, _startStateHash [32]byte, _endStateHash [32]byte) (*types.Transaction, error) {
	return _IChallenge.Contract.Initialize(&_IChallenge.TransactOpts, _defender, _challenger, _verifier, _resultReceiver, _startStateHash, _endStateHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xf28c3e37.
//
// Solidity: function initialize(address _defender, address _challenger, address _verifier, address _resultReceiver, bytes32 _startStateHash, bytes32 _endStateHash) returns()
func (_IChallenge *IChallengeTransactorSession) Initialize(_defender common.Address, _challenger common.Address, _verifier common.Address, _resultReceiver common.Address, _startStateHash [32]byte, _endStateHash [32]byte) (*types.Transaction, error) {
	return _IChallenge.Contract.Initialize(&_IChallenge.TransactOpts, _defender, _challenger, _verifier, _resultReceiver, _startStateHash, _endStateHash)
}

// InitializeChallengeLength is a paid mutator transaction binding the contract method 0x9909e0d9.
//
// Solidity: function initializeChallengeLength(uint256 _numSteps) returns()
func (_IChallenge *IChallengeTransactor) InitializeChallengeLength(opts *bind.TransactOpts, _numSteps *big.Int) (*types.Transaction, error) {
	return _IChallenge.contract.Transact(opts, "initializeChallengeLength", _numSteps)
}

// InitializeChallengeLength is a paid mutator transaction binding the contract method 0x9909e0d9.
//
// Solidity: function initializeChallengeLength(uint256 _numSteps) returns()
func (_IChallenge *IChallengeSession) InitializeChallengeLength(_numSteps *big.Int) (*types.Transaction, error) {
	return _IChallenge.Contract.InitializeChallengeLength(&_IChallenge.TransactOpts, _numSteps)
}

// InitializeChallengeLength is a paid mutator transaction binding the contract method 0x9909e0d9.
//
// Solidity: function initializeChallengeLength(uint256 _numSteps) returns()
func (_IChallenge *IChallengeTransactorSession) InitializeChallengeLength(_numSteps *big.Int) (*types.Transaction, error) {
	return _IChallenge.Contract.InitializeChallengeLength(&_IChallenge.TransactOpts, _numSteps)
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_IChallenge *IChallengeTransactor) Timeout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallenge.contract.Transact(opts, "timeout")
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_IChallenge *IChallengeSession) Timeout() (*types.Transaction, error) {
	return _IChallenge.Contract.Timeout(&_IChallenge.TransactOpts)
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_IChallenge *IChallengeTransactorSession) Timeout() (*types.Transaction, error) {
	return _IChallenge.Contract.Timeout(&_IChallenge.TransactOpts)
}

// VerifyOneStepProof is a paid mutator transaction binding the contract method 0x45b258a7.
//
// Solidity: function verifyOneStepProof(bytes proof, uint256 challengedStepIndex, bytes32[] prevBisection, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_IChallenge *IChallengeTransactor) VerifyOneStepProof(opts *bind.TransactOpts, proof []byte, challengedStepIndex *big.Int, prevBisection [][32]byte, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _IChallenge.contract.Transact(opts, "verifyOneStepProof", proof, challengedStepIndex, prevBisection, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// VerifyOneStepProof is a paid mutator transaction binding the contract method 0x45b258a7.
//
// Solidity: function verifyOneStepProof(bytes proof, uint256 challengedStepIndex, bytes32[] prevBisection, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_IChallenge *IChallengeSession) VerifyOneStepProof(proof []byte, challengedStepIndex *big.Int, prevBisection [][32]byte, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _IChallenge.Contract.VerifyOneStepProof(&_IChallenge.TransactOpts, proof, challengedStepIndex, prevBisection, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// VerifyOneStepProof is a paid mutator transaction binding the contract method 0x45b258a7.
//
// Solidity: function verifyOneStepProof(bytes proof, uint256 challengedStepIndex, bytes32[] prevBisection, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_IChallenge *IChallengeTransactorSession) VerifyOneStepProof(proof []byte, challengedStepIndex *big.Int, prevBisection [][32]byte, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _IChallenge.Contract.VerifyOneStepProof(&_IChallenge.TransactOpts, proof, challengedStepIndex, prevBisection, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// IChallengeBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the IChallenge contract.
type IChallengeBisectedIterator struct {
	Event *IChallengeBisected // Event containing the contract specifics and raw log

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
func (it *IChallengeBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeBisected)
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
		it.Event = new(IChallengeBisected)
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
func (it *IChallengeBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeBisected represents a Bisected event raised by the IChallenge contract.
type IChallengeBisected struct {
	ChallengeState          [32]byte
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x8c3cfc522d91af51bb14f6db452f8c212ba664a426c79e5ef78872e7a1072074.
//
// Solidity: event Bisected(bytes32 challengeState, uint256 challengedSegmentStart, uint256 challengedSegmentLength)
func (_IChallenge *IChallengeFilterer) FilterBisected(opts *bind.FilterOpts) (*IChallengeBisectedIterator, error) {

	logs, sub, err := _IChallenge.contract.FilterLogs(opts, "Bisected")
	if err != nil {
		return nil, err
	}
	return &IChallengeBisectedIterator{contract: _IChallenge.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x8c3cfc522d91af51bb14f6db452f8c212ba664a426c79e5ef78872e7a1072074.
//
// Solidity: event Bisected(bytes32 challengeState, uint256 challengedSegmentStart, uint256 challengedSegmentLength)
func (_IChallenge *IChallengeFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *IChallengeBisected) (event.Subscription, error) {

	logs, sub, err := _IChallenge.contract.WatchLogs(opts, "Bisected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeBisected)
				if err := _IChallenge.contract.UnpackLog(event, "Bisected", log); err != nil {
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

// ParseBisected is a log parse operation binding the contract event 0x8c3cfc522d91af51bb14f6db452f8c212ba664a426c79e5ef78872e7a1072074.
//
// Solidity: event Bisected(bytes32 challengeState, uint256 challengedSegmentStart, uint256 challengedSegmentLength)
func (_IChallenge *IChallengeFilterer) ParseBisected(log types.Log) (*IChallengeBisected, error) {
	event := new(IChallengeBisected)
	if err := _IChallenge.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeChallengeCompletedIterator is returned from FilterChallengeCompleted and is used to iterate over the raw logs and unpacked data for ChallengeCompleted events raised by the IChallenge contract.
type IChallengeChallengeCompletedIterator struct {
	Event *IChallengeChallengeCompleted // Event containing the contract specifics and raw log

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
func (it *IChallengeChallengeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeChallengeCompleted)
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
		it.Event = new(IChallengeChallengeCompleted)
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
func (it *IChallengeChallengeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeChallengeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeChallengeCompleted represents a ChallengeCompleted event raised by the IChallenge contract.
type IChallengeChallengeCompleted struct {
	Winner common.Address
	Loser  common.Address
	Reason uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChallengeCompleted is a free log retrieval operation binding the contract event 0x03f929a9a6b1f0aef5e43cb12b56f862da97ec3de3fda02a52e85f9f3974fb6a.
//
// Solidity: event ChallengeCompleted(address winner, address loser, uint8 reason)
func (_IChallenge *IChallengeFilterer) FilterChallengeCompleted(opts *bind.FilterOpts) (*IChallengeChallengeCompletedIterator, error) {

	logs, sub, err := _IChallenge.contract.FilterLogs(opts, "ChallengeCompleted")
	if err != nil {
		return nil, err
	}
	return &IChallengeChallengeCompletedIterator{contract: _IChallenge.contract, event: "ChallengeCompleted", logs: logs, sub: sub}, nil
}

// WatchChallengeCompleted is a free log subscription operation binding the contract event 0x03f929a9a6b1f0aef5e43cb12b56f862da97ec3de3fda02a52e85f9f3974fb6a.
//
// Solidity: event ChallengeCompleted(address winner, address loser, uint8 reason)
func (_IChallenge *IChallengeFilterer) WatchChallengeCompleted(opts *bind.WatchOpts, sink chan<- *IChallengeChallengeCompleted) (event.Subscription, error) {

	logs, sub, err := _IChallenge.contract.WatchLogs(opts, "ChallengeCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeChallengeCompleted)
				if err := _IChallenge.contract.UnpackLog(event, "ChallengeCompleted", log); err != nil {
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

// ParseChallengeCompleted is a log parse operation binding the contract event 0x03f929a9a6b1f0aef5e43cb12b56f862da97ec3de3fda02a52e85f9f3974fb6a.
//
// Solidity: event ChallengeCompleted(address winner, address loser, uint8 reason)
func (_IChallenge *IChallengeFilterer) ParseChallengeCompleted(log types.Log) (*IChallengeChallengeCompleted, error) {
	event := new(IChallengeChallengeCompleted)
	if err := _IChallenge.contract.UnpackLog(event, "ChallengeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

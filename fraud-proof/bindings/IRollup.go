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

// IRollupMetaData contains all meta data concerning the IRollup contract.
var IRollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AssertionAlreadyResolved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssertionOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengePeriodPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengedStaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DifferentParent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAssertion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker1Challenge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker2Challenge\",\"type\":\"address\"}],\"name\":\"InDifferentChallenge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InboxReadLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinimumAssertionPeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoStaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoUnresolvedAssertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInChallenge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParentAssertionUnstaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreviousStateHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedundantInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakedOnUnconfirmedAssertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakerStakedOnTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakersPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnproposedAssertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeAddr\",\"type\":\"address\"}],\"name\":\"AssertionChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"AssertionConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserterAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vmHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"}],\"name\":\"AssertionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"AssertionRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"StakerStaked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"advanceStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assertions\",\"outputs\":[{\"internalType\":\"contractAssertionMap\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"players\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"assertionIDs\",\"type\":\"uint256[2]\"}],\"name\":\"challengeAssertion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmFirstUnresolvedAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmedInboxSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vmHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"}],\"name\":\"createAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vmHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_batch\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_shouldStartAtElement\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"createAssertionWithStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectFirstUnresolvedAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"removeStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRollupABI is the input ABI used to generate the binding from.
// Deprecated: Use IRollupMetaData.ABI instead.
var IRollupABI = IRollupMetaData.ABI

// IRollup is an auto generated Go binding around an Ethereum contract.
type IRollup struct {
	IRollupCaller     // Read-only binding to the contract
	IRollupTransactor // Write-only binding to the contract
	IRollupFilterer   // Log filterer for contract events
}

// IRollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRollupSession struct {
	Contract     *IRollup          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRollupCallerSession struct {
	Contract *IRollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IRollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRollupTransactorSession struct {
	Contract     *IRollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IRollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRollupRaw struct {
	Contract *IRollup // Generic contract binding to access the raw methods on
}

// IRollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRollupCallerRaw struct {
	Contract *IRollupCaller // Generic read-only contract binding to access the raw methods on
}

// IRollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRollupTransactorRaw struct {
	Contract *IRollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRollup creates a new instance of IRollup, bound to a specific deployed contract.
func NewIRollup(address common.Address, backend bind.ContractBackend) (*IRollup, error) {
	contract, err := bindIRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRollup{IRollupCaller: IRollupCaller{contract: contract}, IRollupTransactor: IRollupTransactor{contract: contract}, IRollupFilterer: IRollupFilterer{contract: contract}}, nil
}

// NewIRollupCaller creates a new read-only instance of IRollup, bound to a specific deployed contract.
func NewIRollupCaller(address common.Address, caller bind.ContractCaller) (*IRollupCaller, error) {
	contract, err := bindIRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRollupCaller{contract: contract}, nil
}

// NewIRollupTransactor creates a new write-only instance of IRollup, bound to a specific deployed contract.
func NewIRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*IRollupTransactor, error) {
	contract, err := bindIRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRollupTransactor{contract: contract}, nil
}

// NewIRollupFilterer creates a new log filterer instance of IRollup, bound to a specific deployed contract.
func NewIRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*IRollupFilterer, error) {
	contract, err := bindIRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRollupFilterer{contract: contract}, nil
}

// bindIRollup binds a generic wrapper to an already deployed contract.
func bindIRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRollup *IRollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRollup.Contract.IRollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRollup *IRollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRollup.Contract.IRollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRollup *IRollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRollup.Contract.IRollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRollup *IRollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRollup *IRollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRollup *IRollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRollup.Contract.contract.Transact(opts, method, params...)
}

// Assertions is a free data retrieval call binding the contract method 0x40d9224b.
//
// Solidity: function assertions() view returns(address)
func (_IRollup *IRollupCaller) Assertions(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRollup.contract.Call(opts, &out, "assertions")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Assertions is a free data retrieval call binding the contract method 0x40d9224b.
//
// Solidity: function assertions() view returns(address)
func (_IRollup *IRollupSession) Assertions() (common.Address, error) {
	return _IRollup.Contract.Assertions(&_IRollup.CallOpts)
}

// Assertions is a free data retrieval call binding the contract method 0x40d9224b.
//
// Solidity: function assertions() view returns(address)
func (_IRollup *IRollupCallerSession) Assertions() (common.Address, error) {
	return _IRollup.Contract.Assertions(&_IRollup.CallOpts)
}

// ConfirmedInboxSize is a free data retrieval call binding the contract method 0xc94b5847.
//
// Solidity: function confirmedInboxSize() view returns(uint256)
func (_IRollup *IRollupCaller) ConfirmedInboxSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRollup.contract.Call(opts, &out, "confirmedInboxSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmedInboxSize is a free data retrieval call binding the contract method 0xc94b5847.
//
// Solidity: function confirmedInboxSize() view returns(uint256)
func (_IRollup *IRollupSession) ConfirmedInboxSize() (*big.Int, error) {
	return _IRollup.Contract.ConfirmedInboxSize(&_IRollup.CallOpts)
}

// ConfirmedInboxSize is a free data retrieval call binding the contract method 0xc94b5847.
//
// Solidity: function confirmedInboxSize() view returns(uint256)
func (_IRollup *IRollupCallerSession) ConfirmedInboxSize() (*big.Int, error) {
	return _IRollup.Contract.ConfirmedInboxSize(&_IRollup.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_IRollup *IRollupCaller) CurrentRequiredStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRollup.contract.Call(opts, &out, "currentRequiredStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_IRollup *IRollupSession) CurrentRequiredStake() (*big.Int, error) {
	return _IRollup.Contract.CurrentRequiredStake(&_IRollup.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_IRollup *IRollupCallerSession) CurrentRequiredStake() (*big.Int, error) {
	return _IRollup.Contract.CurrentRequiredStake(&_IRollup.CallOpts)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_IRollup *IRollupCaller) IsStaked(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _IRollup.contract.Call(opts, &out, "isStaked", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_IRollup *IRollupSession) IsStaked(addr common.Address) (bool, error) {
	return _IRollup.Contract.IsStaked(&_IRollup.CallOpts, addr)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_IRollup *IRollupCallerSession) IsStaked(addr common.Address) (bool, error) {
	return _IRollup.Contract.IsStaked(&_IRollup.CallOpts, addr)
}

// AdvanceStake is a paid mutator transaction binding the contract method 0x8821b2ae.
//
// Solidity: function advanceStake(uint256 assertionID) returns()
func (_IRollup *IRollupTransactor) AdvanceStake(opts *bind.TransactOpts, assertionID *big.Int) (*types.Transaction, error) {
	return _IRollup.contract.Transact(opts, "advanceStake", assertionID)
}

// AdvanceStake is a paid mutator transaction binding the contract method 0x8821b2ae.
//
// Solidity: function advanceStake(uint256 assertionID) returns()
func (_IRollup *IRollupSession) AdvanceStake(assertionID *big.Int) (*types.Transaction, error) {
	return _IRollup.Contract.AdvanceStake(&_IRollup.TransactOpts, assertionID)
}

// AdvanceStake is a paid mutator transaction binding the contract method 0x8821b2ae.
//
// Solidity: function advanceStake(uint256 assertionID) returns()
func (_IRollup *IRollupTransactorSession) AdvanceStake(assertionID *big.Int) (*types.Transaction, error) {
	return _IRollup.Contract.AdvanceStake(&_IRollup.TransactOpts, assertionID)
}

// ChallengeAssertion is a paid mutator transaction binding the contract method 0x2f06d1b0.
//
// Solidity: function challengeAssertion(address[2] players, uint256[2] assertionIDs) returns(address)
func (_IRollup *IRollupTransactor) ChallengeAssertion(opts *bind.TransactOpts, players [2]common.Address, assertionIDs [2]*big.Int) (*types.Transaction, error) {
	return _IRollup.contract.Transact(opts, "challengeAssertion", players, assertionIDs)
}

// ChallengeAssertion is a paid mutator transaction binding the contract method 0x2f06d1b0.
//
// Solidity: function challengeAssertion(address[2] players, uint256[2] assertionIDs) returns(address)
func (_IRollup *IRollupSession) ChallengeAssertion(players [2]common.Address, assertionIDs [2]*big.Int) (*types.Transaction, error) {
	return _IRollup.Contract.ChallengeAssertion(&_IRollup.TransactOpts, players, assertionIDs)
}

// ChallengeAssertion is a paid mutator transaction binding the contract method 0x2f06d1b0.
//
// Solidity: function challengeAssertion(address[2] players, uint256[2] assertionIDs) returns(address)
func (_IRollup *IRollupTransactorSession) ChallengeAssertion(players [2]common.Address, assertionIDs [2]*big.Int) (*types.Transaction, error) {
	return _IRollup.Contract.ChallengeAssertion(&_IRollup.TransactOpts, players, assertionIDs)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winner, address loser) returns()
func (_IRollup *IRollupTransactor) CompleteChallenge(opts *bind.TransactOpts, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IRollup.contract.Transact(opts, "completeChallenge", winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winner, address loser) returns()
func (_IRollup *IRollupSession) CompleteChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IRollup.Contract.CompleteChallenge(&_IRollup.TransactOpts, winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winner, address loser) returns()
func (_IRollup *IRollupTransactorSession) CompleteChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IRollup.Contract.CompleteChallenge(&_IRollup.TransactOpts, winner, loser)
}

// ConfirmFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x2906040e.
//
// Solidity: function confirmFirstUnresolvedAssertion() returns()
func (_IRollup *IRollupTransactor) ConfirmFirstUnresolvedAssertion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRollup.contract.Transact(opts, "confirmFirstUnresolvedAssertion")
}

// ConfirmFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x2906040e.
//
// Solidity: function confirmFirstUnresolvedAssertion() returns()
func (_IRollup *IRollupSession) ConfirmFirstUnresolvedAssertion() (*types.Transaction, error) {
	return _IRollup.Contract.ConfirmFirstUnresolvedAssertion(&_IRollup.TransactOpts)
}

// ConfirmFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x2906040e.
//
// Solidity: function confirmFirstUnresolvedAssertion() returns()
func (_IRollup *IRollupTransactorSession) ConfirmFirstUnresolvedAssertion() (*types.Transaction, error) {
	return _IRollup.Contract.ConfirmFirstUnresolvedAssertion(&_IRollup.TransactOpts)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0xb6da898f.
//
// Solidity: function createAssertion(bytes32 vmHash, uint256 inboxSize) returns()
func (_IRollup *IRollupTransactor) CreateAssertion(opts *bind.TransactOpts, vmHash [32]byte, inboxSize *big.Int) (*types.Transaction, error) {
	return _IRollup.contract.Transact(opts, "createAssertion", vmHash, inboxSize)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0xb6da898f.
//
// Solidity: function createAssertion(bytes32 vmHash, uint256 inboxSize) returns()
func (_IRollup *IRollupSession) CreateAssertion(vmHash [32]byte, inboxSize *big.Int) (*types.Transaction, error) {
	return _IRollup.Contract.CreateAssertion(&_IRollup.TransactOpts, vmHash, inboxSize)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0xb6da898f.
//
// Solidity: function createAssertion(bytes32 vmHash, uint256 inboxSize) returns()
func (_IRollup *IRollupTransactorSession) CreateAssertion(vmHash [32]byte, inboxSize *big.Int) (*types.Transaction, error) {
	return _IRollup.Contract.CreateAssertion(&_IRollup.TransactOpts, vmHash, inboxSize)
}

// CreateAssertionWithStateBatch is a paid mutator transaction binding the contract method 0x49cd3004.
//
// Solidity: function createAssertionWithStateBatch(bytes32 vmHash, uint256 inboxSize, bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_IRollup *IRollupTransactor) CreateAssertionWithStateBatch(opts *bind.TransactOpts, vmHash [32]byte, inboxSize *big.Int, _batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _IRollup.contract.Transact(opts, "createAssertionWithStateBatch", vmHash, inboxSize, _batch, _shouldStartAtElement, _signature)
}

// CreateAssertionWithStateBatch is a paid mutator transaction binding the contract method 0x49cd3004.
//
// Solidity: function createAssertionWithStateBatch(bytes32 vmHash, uint256 inboxSize, bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_IRollup *IRollupSession) CreateAssertionWithStateBatch(vmHash [32]byte, inboxSize *big.Int, _batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _IRollup.Contract.CreateAssertionWithStateBatch(&_IRollup.TransactOpts, vmHash, inboxSize, _batch, _shouldStartAtElement, _signature)
}

// CreateAssertionWithStateBatch is a paid mutator transaction binding the contract method 0x49cd3004.
//
// Solidity: function createAssertionWithStateBatch(bytes32 vmHash, uint256 inboxSize, bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_IRollup *IRollupTransactorSession) CreateAssertionWithStateBatch(vmHash [32]byte, inboxSize *big.Int, _batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _IRollup.Contract.CreateAssertionWithStateBatch(&_IRollup.TransactOpts, vmHash, inboxSize, _batch, _shouldStartAtElement, _signature)
}

// RejectFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x042dca93.
//
// Solidity: function rejectFirstUnresolvedAssertion(address stakerAddress) returns()
func (_IRollup *IRollupTransactor) RejectFirstUnresolvedAssertion(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _IRollup.contract.Transact(opts, "rejectFirstUnresolvedAssertion", stakerAddress)
}

// RejectFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x042dca93.
//
// Solidity: function rejectFirstUnresolvedAssertion(address stakerAddress) returns()
func (_IRollup *IRollupSession) RejectFirstUnresolvedAssertion(stakerAddress common.Address) (*types.Transaction, error) {
	return _IRollup.Contract.RejectFirstUnresolvedAssertion(&_IRollup.TransactOpts, stakerAddress)
}

// RejectFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x042dca93.
//
// Solidity: function rejectFirstUnresolvedAssertion(address stakerAddress) returns()
func (_IRollup *IRollupTransactorSession) RejectFirstUnresolvedAssertion(stakerAddress common.Address) (*types.Transaction, error) {
	return _IRollup.Contract.RejectFirstUnresolvedAssertion(&_IRollup.TransactOpts, stakerAddress)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address stakerAddress) returns()
func (_IRollup *IRollupTransactor) RemoveStake(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _IRollup.contract.Transact(opts, "removeStake", stakerAddress)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address stakerAddress) returns()
func (_IRollup *IRollupSession) RemoveStake(stakerAddress common.Address) (*types.Transaction, error) {
	return _IRollup.Contract.RemoveStake(&_IRollup.TransactOpts, stakerAddress)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address stakerAddress) returns()
func (_IRollup *IRollupTransactorSession) RemoveStake(stakerAddress common.Address) (*types.Transaction, error) {
	return _IRollup.Contract.RemoveStake(&_IRollup.TransactOpts, stakerAddress)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_IRollup *IRollupTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRollup.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_IRollup *IRollupSession) Stake() (*types.Transaction, error) {
	return _IRollup.Contract.Stake(&_IRollup.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_IRollup *IRollupTransactorSession) Stake() (*types.Transaction, error) {
	return _IRollup.Contract.Stake(&_IRollup.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeAmount) returns()
func (_IRollup *IRollupTransactor) Unstake(opts *bind.TransactOpts, stakeAmount *big.Int) (*types.Transaction, error) {
	return _IRollup.contract.Transact(opts, "unstake", stakeAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeAmount) returns()
func (_IRollup *IRollupSession) Unstake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _IRollup.Contract.Unstake(&_IRollup.TransactOpts, stakeAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeAmount) returns()
func (_IRollup *IRollupTransactorSession) Unstake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _IRollup.Contract.Unstake(&_IRollup.TransactOpts, stakeAmount)
}

// IRollupAssertionChallengedIterator is returned from FilterAssertionChallenged and is used to iterate over the raw logs and unpacked data for AssertionChallenged events raised by the IRollup contract.
type IRollupAssertionChallengedIterator struct {
	Event *IRollupAssertionChallenged // Event containing the contract specifics and raw log

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
func (it *IRollupAssertionChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRollupAssertionChallenged)
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
		it.Event = new(IRollupAssertionChallenged)
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
func (it *IRollupAssertionChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRollupAssertionChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRollupAssertionChallenged represents a AssertionChallenged event raised by the IRollup contract.
type IRollupAssertionChallenged struct {
	AssertionID   *big.Int
	ChallengeAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAssertionChallenged is a free log retrieval operation binding the contract event 0xd0ebe74b4f7d89a9b0fdc9d95f887a7b925c6c7300b5c4b2c3304d97925840fa.
//
// Solidity: event AssertionChallenged(uint256 assertionID, address challengeAddr)
func (_IRollup *IRollupFilterer) FilterAssertionChallenged(opts *bind.FilterOpts) (*IRollupAssertionChallengedIterator, error) {

	logs, sub, err := _IRollup.contract.FilterLogs(opts, "AssertionChallenged")
	if err != nil {
		return nil, err
	}
	return &IRollupAssertionChallengedIterator{contract: _IRollup.contract, event: "AssertionChallenged", logs: logs, sub: sub}, nil
}

// WatchAssertionChallenged is a free log subscription operation binding the contract event 0xd0ebe74b4f7d89a9b0fdc9d95f887a7b925c6c7300b5c4b2c3304d97925840fa.
//
// Solidity: event AssertionChallenged(uint256 assertionID, address challengeAddr)
func (_IRollup *IRollupFilterer) WatchAssertionChallenged(opts *bind.WatchOpts, sink chan<- *IRollupAssertionChallenged) (event.Subscription, error) {

	logs, sub, err := _IRollup.contract.WatchLogs(opts, "AssertionChallenged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRollupAssertionChallenged)
				if err := _IRollup.contract.UnpackLog(event, "AssertionChallenged", log); err != nil {
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

// ParseAssertionChallenged is a log parse operation binding the contract event 0xd0ebe74b4f7d89a9b0fdc9d95f887a7b925c6c7300b5c4b2c3304d97925840fa.
//
// Solidity: event AssertionChallenged(uint256 assertionID, address challengeAddr)
func (_IRollup *IRollupFilterer) ParseAssertionChallenged(log types.Log) (*IRollupAssertionChallenged, error) {
	event := new(IRollupAssertionChallenged)
	if err := _IRollup.contract.UnpackLog(event, "AssertionChallenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRollupAssertionConfirmedIterator is returned from FilterAssertionConfirmed and is used to iterate over the raw logs and unpacked data for AssertionConfirmed events raised by the IRollup contract.
type IRollupAssertionConfirmedIterator struct {
	Event *IRollupAssertionConfirmed // Event containing the contract specifics and raw log

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
func (it *IRollupAssertionConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRollupAssertionConfirmed)
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
		it.Event = new(IRollupAssertionConfirmed)
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
func (it *IRollupAssertionConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRollupAssertionConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRollupAssertionConfirmed represents a AssertionConfirmed event raised by the IRollup contract.
type IRollupAssertionConfirmed struct {
	AssertionID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssertionConfirmed is a free log retrieval operation binding the contract event 0x453430d123684340024ae0a229704bdab39c93dc48bb5a0b4bc83142d95d48ef.
//
// Solidity: event AssertionConfirmed(uint256 assertionID)
func (_IRollup *IRollupFilterer) FilterAssertionConfirmed(opts *bind.FilterOpts) (*IRollupAssertionConfirmedIterator, error) {

	logs, sub, err := _IRollup.contract.FilterLogs(opts, "AssertionConfirmed")
	if err != nil {
		return nil, err
	}
	return &IRollupAssertionConfirmedIterator{contract: _IRollup.contract, event: "AssertionConfirmed", logs: logs, sub: sub}, nil
}

// WatchAssertionConfirmed is a free log subscription operation binding the contract event 0x453430d123684340024ae0a229704bdab39c93dc48bb5a0b4bc83142d95d48ef.
//
// Solidity: event AssertionConfirmed(uint256 assertionID)
func (_IRollup *IRollupFilterer) WatchAssertionConfirmed(opts *bind.WatchOpts, sink chan<- *IRollupAssertionConfirmed) (event.Subscription, error) {

	logs, sub, err := _IRollup.contract.WatchLogs(opts, "AssertionConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRollupAssertionConfirmed)
				if err := _IRollup.contract.UnpackLog(event, "AssertionConfirmed", log); err != nil {
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

// ParseAssertionConfirmed is a log parse operation binding the contract event 0x453430d123684340024ae0a229704bdab39c93dc48bb5a0b4bc83142d95d48ef.
//
// Solidity: event AssertionConfirmed(uint256 assertionID)
func (_IRollup *IRollupFilterer) ParseAssertionConfirmed(log types.Log) (*IRollupAssertionConfirmed, error) {
	event := new(IRollupAssertionConfirmed)
	if err := _IRollup.contract.UnpackLog(event, "AssertionConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRollupAssertionCreatedIterator is returned from FilterAssertionCreated and is used to iterate over the raw logs and unpacked data for AssertionCreated events raised by the IRollup contract.
type IRollupAssertionCreatedIterator struct {
	Event *IRollupAssertionCreated // Event containing the contract specifics and raw log

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
func (it *IRollupAssertionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRollupAssertionCreated)
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
		it.Event = new(IRollupAssertionCreated)
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
func (it *IRollupAssertionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRollupAssertionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRollupAssertionCreated represents a AssertionCreated event raised by the IRollup contract.
type IRollupAssertionCreated struct {
	AssertionID  *big.Int
	AsserterAddr common.Address
	VmHash       [32]byte
	InboxSize    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssertionCreated is a free log retrieval operation binding the contract event 0x5c610f28399ecc14b66149012a0197a5e3257a8c397125afee95d1cf4b950734.
//
// Solidity: event AssertionCreated(uint256 assertionID, address asserterAddr, bytes32 vmHash, uint256 inboxSize)
func (_IRollup *IRollupFilterer) FilterAssertionCreated(opts *bind.FilterOpts) (*IRollupAssertionCreatedIterator, error) {

	logs, sub, err := _IRollup.contract.FilterLogs(opts, "AssertionCreated")
	if err != nil {
		return nil, err
	}
	return &IRollupAssertionCreatedIterator{contract: _IRollup.contract, event: "AssertionCreated", logs: logs, sub: sub}, nil
}

// WatchAssertionCreated is a free log subscription operation binding the contract event 0x5c610f28399ecc14b66149012a0197a5e3257a8c397125afee95d1cf4b950734.
//
// Solidity: event AssertionCreated(uint256 assertionID, address asserterAddr, bytes32 vmHash, uint256 inboxSize)
func (_IRollup *IRollupFilterer) WatchAssertionCreated(opts *bind.WatchOpts, sink chan<- *IRollupAssertionCreated) (event.Subscription, error) {

	logs, sub, err := _IRollup.contract.WatchLogs(opts, "AssertionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRollupAssertionCreated)
				if err := _IRollup.contract.UnpackLog(event, "AssertionCreated", log); err != nil {
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

// ParseAssertionCreated is a log parse operation binding the contract event 0x5c610f28399ecc14b66149012a0197a5e3257a8c397125afee95d1cf4b950734.
//
// Solidity: event AssertionCreated(uint256 assertionID, address asserterAddr, bytes32 vmHash, uint256 inboxSize)
func (_IRollup *IRollupFilterer) ParseAssertionCreated(log types.Log) (*IRollupAssertionCreated, error) {
	event := new(IRollupAssertionCreated)
	if err := _IRollup.contract.UnpackLog(event, "AssertionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRollupAssertionRejectedIterator is returned from FilterAssertionRejected and is used to iterate over the raw logs and unpacked data for AssertionRejected events raised by the IRollup contract.
type IRollupAssertionRejectedIterator struct {
	Event *IRollupAssertionRejected // Event containing the contract specifics and raw log

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
func (it *IRollupAssertionRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRollupAssertionRejected)
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
		it.Event = new(IRollupAssertionRejected)
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
func (it *IRollupAssertionRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRollupAssertionRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRollupAssertionRejected represents a AssertionRejected event raised by the IRollup contract.
type IRollupAssertionRejected struct {
	AssertionID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssertionRejected is a free log retrieval operation binding the contract event 0x5b24ab8ceb442373727ac5c559a027521cb52db451c74710ebed9faa5fe15a7c.
//
// Solidity: event AssertionRejected(uint256 assertionID)
func (_IRollup *IRollupFilterer) FilterAssertionRejected(opts *bind.FilterOpts) (*IRollupAssertionRejectedIterator, error) {

	logs, sub, err := _IRollup.contract.FilterLogs(opts, "AssertionRejected")
	if err != nil {
		return nil, err
	}
	return &IRollupAssertionRejectedIterator{contract: _IRollup.contract, event: "AssertionRejected", logs: logs, sub: sub}, nil
}

// WatchAssertionRejected is a free log subscription operation binding the contract event 0x5b24ab8ceb442373727ac5c559a027521cb52db451c74710ebed9faa5fe15a7c.
//
// Solidity: event AssertionRejected(uint256 assertionID)
func (_IRollup *IRollupFilterer) WatchAssertionRejected(opts *bind.WatchOpts, sink chan<- *IRollupAssertionRejected) (event.Subscription, error) {

	logs, sub, err := _IRollup.contract.WatchLogs(opts, "AssertionRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRollupAssertionRejected)
				if err := _IRollup.contract.UnpackLog(event, "AssertionRejected", log); err != nil {
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

// ParseAssertionRejected is a log parse operation binding the contract event 0x5b24ab8ceb442373727ac5c559a027521cb52db451c74710ebed9faa5fe15a7c.
//
// Solidity: event AssertionRejected(uint256 assertionID)
func (_IRollup *IRollupFilterer) ParseAssertionRejected(log types.Log) (*IRollupAssertionRejected, error) {
	event := new(IRollupAssertionRejected)
	if err := _IRollup.contract.UnpackLog(event, "AssertionRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRollupStakerStakedIterator is returned from FilterStakerStaked and is used to iterate over the raw logs and unpacked data for StakerStaked events raised by the IRollup contract.
type IRollupStakerStakedIterator struct {
	Event *IRollupStakerStaked // Event containing the contract specifics and raw log

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
func (it *IRollupStakerStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRollupStakerStaked)
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
		it.Event = new(IRollupStakerStaked)
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
func (it *IRollupStakerStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRollupStakerStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRollupStakerStaked represents a StakerStaked event raised by the IRollup contract.
type IRollupStakerStaked struct {
	StakerAddr  common.Address
	AssertionID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakerStaked is a free log retrieval operation binding the contract event 0x617d31491414a4ab2bd831e566a31837fa7fb6582921c91dffbbe83fbca789f3.
//
// Solidity: event StakerStaked(address stakerAddr, uint256 assertionID)
func (_IRollup *IRollupFilterer) FilterStakerStaked(opts *bind.FilterOpts) (*IRollupStakerStakedIterator, error) {

	logs, sub, err := _IRollup.contract.FilterLogs(opts, "StakerStaked")
	if err != nil {
		return nil, err
	}
	return &IRollupStakerStakedIterator{contract: _IRollup.contract, event: "StakerStaked", logs: logs, sub: sub}, nil
}

// WatchStakerStaked is a free log subscription operation binding the contract event 0x617d31491414a4ab2bd831e566a31837fa7fb6582921c91dffbbe83fbca789f3.
//
// Solidity: event StakerStaked(address stakerAddr, uint256 assertionID)
func (_IRollup *IRollupFilterer) WatchStakerStaked(opts *bind.WatchOpts, sink chan<- *IRollupStakerStaked) (event.Subscription, error) {

	logs, sub, err := _IRollup.contract.WatchLogs(opts, "StakerStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRollupStakerStaked)
				if err := _IRollup.contract.UnpackLog(event, "StakerStaked", log); err != nil {
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

// ParseStakerStaked is a log parse operation binding the contract event 0x617d31491414a4ab2bd831e566a31837fa7fb6582921c91dffbbe83fbca789f3.
//
// Solidity: event StakerStaked(address stakerAddr, uint256 assertionID)
func (_IRollup *IRollupFilterer) ParseStakerStaked(log types.Log) (*IRollupStakerStaked, error) {
	event := new(IRollupStakerStaked)
	if err := _IRollup.contract.UnpackLog(event, "StakerStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

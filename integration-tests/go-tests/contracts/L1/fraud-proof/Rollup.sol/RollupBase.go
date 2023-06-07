// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RollupBase

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

// RollupBaseMetaData contains all meta data concerning the RollupBase contract.
var RollupBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AssertionAlreadyResolved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssertionOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengePeriodPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengedStaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DifferentParent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAssertion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker1Challenge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker2Challenge\",\"type\":\"address\"}],\"name\":\"InDifferentChallenge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InboxReadLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinimumAssertionPeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoStaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoUnresolvedAssertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInChallenge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParentAssertionUnstaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreviousStateHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedundantInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakedOnUnconfirmedAssertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakerStakedOnTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakersPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnproposedAssertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeAddr\",\"type\":\"address\"}],\"name\":\"AssertionChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"AssertionConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserterAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vmHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"}],\"name\":\"AssertionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"AssertionRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"StakerStaked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"advanceStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assertions\",\"outputs\":[{\"internalType\":\"contractAssertionMap\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"players\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"assertionIDs\",\"type\":\"uint256[2]\"}],\"name\":\"challengeAssertion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmFirstUnresolvedAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmedInboxSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vmHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"}],\"name\":\"createAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vmHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_batch\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_shouldStartAtElement\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"createAssertionWithStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rejectFirstUnresolvedAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"removeStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifierEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RollupBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupBaseMetaData.ABI instead.
var RollupBaseABI = RollupBaseMetaData.ABI

// RollupBase is an auto generated Go binding around an Ethereum contract.
type RollupBase struct {
	RollupBaseCaller     // Read-only binding to the contract
	RollupBaseTransactor // Write-only binding to the contract
	RollupBaseFilterer   // Log filterer for contract events
}

// RollupBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupBaseSession struct {
	Contract     *RollupBase       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupBaseCallerSession struct {
	Contract *RollupBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupBaseTransactorSession struct {
	Contract     *RollupBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupBaseRaw struct {
	Contract *RollupBase // Generic contract binding to access the raw methods on
}

// RollupBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupBaseCallerRaw struct {
	Contract *RollupBaseCaller // Generic read-only contract binding to access the raw methods on
}

// RollupBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupBaseTransactorRaw struct {
	Contract *RollupBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupBase creates a new instance of RollupBase, bound to a specific deployed contract.
func NewRollupBase(address common.Address, backend bind.ContractBackend) (*RollupBase, error) {
	contract, err := bindRollupBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupBase{RollupBaseCaller: RollupBaseCaller{contract: contract}, RollupBaseTransactor: RollupBaseTransactor{contract: contract}, RollupBaseFilterer: RollupBaseFilterer{contract: contract}}, nil
}

// NewRollupBaseCaller creates a new read-only instance of RollupBase, bound to a specific deployed contract.
func NewRollupBaseCaller(address common.Address, caller bind.ContractCaller) (*RollupBaseCaller, error) {
	contract, err := bindRollupBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupBaseCaller{contract: contract}, nil
}

// NewRollupBaseTransactor creates a new write-only instance of RollupBase, bound to a specific deployed contract.
func NewRollupBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupBaseTransactor, error) {
	contract, err := bindRollupBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupBaseTransactor{contract: contract}, nil
}

// NewRollupBaseFilterer creates a new log filterer instance of RollupBase, bound to a specific deployed contract.
func NewRollupBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupBaseFilterer, error) {
	contract, err := bindRollupBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupBaseFilterer{contract: contract}, nil
}

// bindRollupBase binds a generic wrapper to an already deployed contract.
func bindRollupBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupBase *RollupBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupBase.Contract.RollupBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupBase *RollupBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupBase.Contract.RollupBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupBase *RollupBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupBase.Contract.RollupBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupBase *RollupBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupBase *RollupBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupBase *RollupBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupBase.Contract.contract.Transact(opts, method, params...)
}

// Assertions is a free data retrieval call binding the contract method 0x40d9224b.
//
// Solidity: function assertions() view returns(address)
func (_RollupBase *RollupBaseCaller) Assertions(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupBase.contract.Call(opts, &out, "assertions")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Assertions is a free data retrieval call binding the contract method 0x40d9224b.
//
// Solidity: function assertions() view returns(address)
func (_RollupBase *RollupBaseSession) Assertions() (common.Address, error) {
	return _RollupBase.Contract.Assertions(&_RollupBase.CallOpts)
}

// Assertions is a free data retrieval call binding the contract method 0x40d9224b.
//
// Solidity: function assertions() view returns(address)
func (_RollupBase *RollupBaseCallerSession) Assertions() (common.Address, error) {
	return _RollupBase.Contract.Assertions(&_RollupBase.CallOpts)
}

// BaseStakeAmount is a free data retrieval call binding the contract method 0x71129559.
//
// Solidity: function baseStakeAmount() view returns(uint256)
func (_RollupBase *RollupBaseCaller) BaseStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupBase.contract.Call(opts, &out, "baseStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStakeAmount is a free data retrieval call binding the contract method 0x71129559.
//
// Solidity: function baseStakeAmount() view returns(uint256)
func (_RollupBase *RollupBaseSession) BaseStakeAmount() (*big.Int, error) {
	return _RollupBase.Contract.BaseStakeAmount(&_RollupBase.CallOpts)
}

// BaseStakeAmount is a free data retrieval call binding the contract method 0x71129559.
//
// Solidity: function baseStakeAmount() view returns(uint256)
func (_RollupBase *RollupBaseCallerSession) BaseStakeAmount() (*big.Int, error) {
	return _RollupBase.Contract.BaseStakeAmount(&_RollupBase.CallOpts)
}

// ConfirmedInboxSize is a free data retrieval call binding the contract method 0xc94b5847.
//
// Solidity: function confirmedInboxSize() view returns(uint256)
func (_RollupBase *RollupBaseCaller) ConfirmedInboxSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupBase.contract.Call(opts, &out, "confirmedInboxSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmedInboxSize is a free data retrieval call binding the contract method 0xc94b5847.
//
// Solidity: function confirmedInboxSize() view returns(uint256)
func (_RollupBase *RollupBaseSession) ConfirmedInboxSize() (*big.Int, error) {
	return _RollupBase.Contract.ConfirmedInboxSize(&_RollupBase.CallOpts)
}

// ConfirmedInboxSize is a free data retrieval call binding the contract method 0xc94b5847.
//
// Solidity: function confirmedInboxSize() view returns(uint256)
func (_RollupBase *RollupBaseCallerSession) ConfirmedInboxSize() (*big.Int, error) {
	return _RollupBase.Contract.ConfirmedInboxSize(&_RollupBase.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_RollupBase *RollupBaseCaller) CurrentRequiredStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupBase.contract.Call(opts, &out, "currentRequiredStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_RollupBase *RollupBaseSession) CurrentRequiredStake() (*big.Int, error) {
	return _RollupBase.Contract.CurrentRequiredStake(&_RollupBase.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_RollupBase *RollupBaseCallerSession) CurrentRequiredStake() (*big.Int, error) {
	return _RollupBase.Contract.CurrentRequiredStake(&_RollupBase.CallOpts)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_RollupBase *RollupBaseCaller) IsStaked(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _RollupBase.contract.Call(opts, &out, "isStaked", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_RollupBase *RollupBaseSession) IsStaked(addr common.Address) (bool, error) {
	return _RollupBase.Contract.IsStaked(&_RollupBase.CallOpts, addr)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_RollupBase *RollupBaseCallerSession) IsStaked(addr common.Address) (bool, error) {
	return _RollupBase.Contract.IsStaked(&_RollupBase.CallOpts, addr)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_RollupBase *RollupBaseCaller) MinimumAssertionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupBase.contract.Call(opts, &out, "minimumAssertionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_RollupBase *RollupBaseSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _RollupBase.Contract.MinimumAssertionPeriod(&_RollupBase.CallOpts)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_RollupBase *RollupBaseCallerSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _RollupBase.Contract.MinimumAssertionPeriod(&_RollupBase.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_RollupBase *RollupBaseCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupBase.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_RollupBase *RollupBaseSession) StakeToken() (common.Address, error) {
	return _RollupBase.Contract.StakeToken(&_RollupBase.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_RollupBase *RollupBaseCallerSession) StakeToken() (common.Address, error) {
	return _RollupBase.Contract.StakeToken(&_RollupBase.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_RollupBase *RollupBaseCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupBase.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_RollupBase *RollupBaseSession) Verifier() (common.Address, error) {
	return _RollupBase.Contract.Verifier(&_RollupBase.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_RollupBase *RollupBaseCallerSession) Verifier() (common.Address, error) {
	return _RollupBase.Contract.Verifier(&_RollupBase.CallOpts)
}

// AdvanceStake is a paid mutator transaction binding the contract method 0x8821b2ae.
//
// Solidity: function advanceStake(uint256 assertionID) returns()
func (_RollupBase *RollupBaseTransactor) AdvanceStake(opts *bind.TransactOpts, assertionID *big.Int) (*types.Transaction, error) {
	return _RollupBase.contract.Transact(opts, "advanceStake", assertionID)
}

// AdvanceStake is a paid mutator transaction binding the contract method 0x8821b2ae.
//
// Solidity: function advanceStake(uint256 assertionID) returns()
func (_RollupBase *RollupBaseSession) AdvanceStake(assertionID *big.Int) (*types.Transaction, error) {
	return _RollupBase.Contract.AdvanceStake(&_RollupBase.TransactOpts, assertionID)
}

// AdvanceStake is a paid mutator transaction binding the contract method 0x8821b2ae.
//
// Solidity: function advanceStake(uint256 assertionID) returns()
func (_RollupBase *RollupBaseTransactorSession) AdvanceStake(assertionID *big.Int) (*types.Transaction, error) {
	return _RollupBase.Contract.AdvanceStake(&_RollupBase.TransactOpts, assertionID)
}

// ChallengeAssertion is a paid mutator transaction binding the contract method 0x2f06d1b0.
//
// Solidity: function challengeAssertion(address[2] players, uint256[2] assertionIDs) returns(address)
func (_RollupBase *RollupBaseTransactor) ChallengeAssertion(opts *bind.TransactOpts, players [2]common.Address, assertionIDs [2]*big.Int) (*types.Transaction, error) {
	return _RollupBase.contract.Transact(opts, "challengeAssertion", players, assertionIDs)
}

// ChallengeAssertion is a paid mutator transaction binding the contract method 0x2f06d1b0.
//
// Solidity: function challengeAssertion(address[2] players, uint256[2] assertionIDs) returns(address)
func (_RollupBase *RollupBaseSession) ChallengeAssertion(players [2]common.Address, assertionIDs [2]*big.Int) (*types.Transaction, error) {
	return _RollupBase.Contract.ChallengeAssertion(&_RollupBase.TransactOpts, players, assertionIDs)
}

// ChallengeAssertion is a paid mutator transaction binding the contract method 0x2f06d1b0.
//
// Solidity: function challengeAssertion(address[2] players, uint256[2] assertionIDs) returns(address)
func (_RollupBase *RollupBaseTransactorSession) ChallengeAssertion(players [2]common.Address, assertionIDs [2]*big.Int) (*types.Transaction, error) {
	return _RollupBase.Contract.ChallengeAssertion(&_RollupBase.TransactOpts, players, assertionIDs)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winner, address loser) returns()
func (_RollupBase *RollupBaseTransactor) CompleteChallenge(opts *bind.TransactOpts, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _RollupBase.contract.Transact(opts, "completeChallenge", winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winner, address loser) returns()
func (_RollupBase *RollupBaseSession) CompleteChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _RollupBase.Contract.CompleteChallenge(&_RollupBase.TransactOpts, winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winner, address loser) returns()
func (_RollupBase *RollupBaseTransactorSession) CompleteChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _RollupBase.Contract.CompleteChallenge(&_RollupBase.TransactOpts, winner, loser)
}

// ConfirmFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x2906040e.
//
// Solidity: function confirmFirstUnresolvedAssertion() returns()
func (_RollupBase *RollupBaseTransactor) ConfirmFirstUnresolvedAssertion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupBase.contract.Transact(opts, "confirmFirstUnresolvedAssertion")
}

// ConfirmFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x2906040e.
//
// Solidity: function confirmFirstUnresolvedAssertion() returns()
func (_RollupBase *RollupBaseSession) ConfirmFirstUnresolvedAssertion() (*types.Transaction, error) {
	return _RollupBase.Contract.ConfirmFirstUnresolvedAssertion(&_RollupBase.TransactOpts)
}

// ConfirmFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x2906040e.
//
// Solidity: function confirmFirstUnresolvedAssertion() returns()
func (_RollupBase *RollupBaseTransactorSession) ConfirmFirstUnresolvedAssertion() (*types.Transaction, error) {
	return _RollupBase.Contract.ConfirmFirstUnresolvedAssertion(&_RollupBase.TransactOpts)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0xb6da898f.
//
// Solidity: function createAssertion(bytes32 vmHash, uint256 inboxSize) returns()
func (_RollupBase *RollupBaseTransactor) CreateAssertion(opts *bind.TransactOpts, vmHash [32]byte, inboxSize *big.Int) (*types.Transaction, error) {
	return _RollupBase.contract.Transact(opts, "createAssertion", vmHash, inboxSize)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0xb6da898f.
//
// Solidity: function createAssertion(bytes32 vmHash, uint256 inboxSize) returns()
func (_RollupBase *RollupBaseSession) CreateAssertion(vmHash [32]byte, inboxSize *big.Int) (*types.Transaction, error) {
	return _RollupBase.Contract.CreateAssertion(&_RollupBase.TransactOpts, vmHash, inboxSize)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0xb6da898f.
//
// Solidity: function createAssertion(bytes32 vmHash, uint256 inboxSize) returns()
func (_RollupBase *RollupBaseTransactorSession) CreateAssertion(vmHash [32]byte, inboxSize *big.Int) (*types.Transaction, error) {
	return _RollupBase.Contract.CreateAssertion(&_RollupBase.TransactOpts, vmHash, inboxSize)
}

// CreateAssertionWithStateBatch is a paid mutator transaction binding the contract method 0x49cd3004.
//
// Solidity: function createAssertionWithStateBatch(bytes32 vmHash, uint256 inboxSize, bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_RollupBase *RollupBaseTransactor) CreateAssertionWithStateBatch(opts *bind.TransactOpts, vmHash [32]byte, inboxSize *big.Int, _batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _RollupBase.contract.Transact(opts, "createAssertionWithStateBatch", vmHash, inboxSize, _batch, _shouldStartAtElement, _signature)
}

// CreateAssertionWithStateBatch is a paid mutator transaction binding the contract method 0x49cd3004.
//
// Solidity: function createAssertionWithStateBatch(bytes32 vmHash, uint256 inboxSize, bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_RollupBase *RollupBaseSession) CreateAssertionWithStateBatch(vmHash [32]byte, inboxSize *big.Int, _batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _RollupBase.Contract.CreateAssertionWithStateBatch(&_RollupBase.TransactOpts, vmHash, inboxSize, _batch, _shouldStartAtElement, _signature)
}

// CreateAssertionWithStateBatch is a paid mutator transaction binding the contract method 0x49cd3004.
//
// Solidity: function createAssertionWithStateBatch(bytes32 vmHash, uint256 inboxSize, bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_RollupBase *RollupBaseTransactorSession) CreateAssertionWithStateBatch(vmHash [32]byte, inboxSize *big.Int, _batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _RollupBase.Contract.CreateAssertionWithStateBatch(&_RollupBase.TransactOpts, vmHash, inboxSize, _batch, _shouldStartAtElement, _signature)
}

// RejectFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x30b26075.
//
// Solidity: function rejectFirstUnresolvedAssertion() returns()
func (_RollupBase *RollupBaseTransactor) RejectFirstUnresolvedAssertion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupBase.contract.Transact(opts, "rejectFirstUnresolvedAssertion")
}

// RejectFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x30b26075.
//
// Solidity: function rejectFirstUnresolvedAssertion() returns()
func (_RollupBase *RollupBaseSession) RejectFirstUnresolvedAssertion() (*types.Transaction, error) {
	return _RollupBase.Contract.RejectFirstUnresolvedAssertion(&_RollupBase.TransactOpts)
}

// RejectFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x30b26075.
//
// Solidity: function rejectFirstUnresolvedAssertion() returns()
func (_RollupBase *RollupBaseTransactorSession) RejectFirstUnresolvedAssertion() (*types.Transaction, error) {
	return _RollupBase.Contract.RejectFirstUnresolvedAssertion(&_RollupBase.TransactOpts)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address stakerAddress) returns()
func (_RollupBase *RollupBaseTransactor) RemoveStake(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _RollupBase.contract.Transact(opts, "removeStake", stakerAddress)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address stakerAddress) returns()
func (_RollupBase *RollupBaseSession) RemoveStake(stakerAddress common.Address) (*types.Transaction, error) {
	return _RollupBase.Contract.RemoveStake(&_RollupBase.TransactOpts, stakerAddress)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address stakerAddress) returns()
func (_RollupBase *RollupBaseTransactorSession) RemoveStake(stakerAddress common.Address) (*types.Transaction, error) {
	return _RollupBase.Contract.RemoveStake(&_RollupBase.TransactOpts, stakerAddress)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 stakeAmount, address operator) returns()
func (_RollupBase *RollupBaseTransactor) Stake(opts *bind.TransactOpts, stakeAmount *big.Int, operator common.Address) (*types.Transaction, error) {
	return _RollupBase.contract.Transact(opts, "stake", stakeAmount, operator)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 stakeAmount, address operator) returns()
func (_RollupBase *RollupBaseSession) Stake(stakeAmount *big.Int, operator common.Address) (*types.Transaction, error) {
	return _RollupBase.Contract.Stake(&_RollupBase.TransactOpts, stakeAmount, operator)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 stakeAmount, address operator) returns()
func (_RollupBase *RollupBaseTransactorSession) Stake(stakeAmount *big.Int, operator common.Address) (*types.Transaction, error) {
	return _RollupBase.Contract.Stake(&_RollupBase.TransactOpts, stakeAmount, operator)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeAmount) returns()
func (_RollupBase *RollupBaseTransactor) Unstake(opts *bind.TransactOpts, stakeAmount *big.Int) (*types.Transaction, error) {
	return _RollupBase.contract.Transact(opts, "unstake", stakeAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeAmount) returns()
func (_RollupBase *RollupBaseSession) Unstake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _RollupBase.Contract.Unstake(&_RollupBase.TransactOpts, stakeAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeAmount) returns()
func (_RollupBase *RollupBaseTransactorSession) Unstake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _RollupBase.Contract.Unstake(&_RollupBase.TransactOpts, stakeAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RollupBase *RollupBaseTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupBase.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RollupBase *RollupBaseSession) Withdraw() (*types.Transaction, error) {
	return _RollupBase.Contract.Withdraw(&_RollupBase.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RollupBase *RollupBaseTransactorSession) Withdraw() (*types.Transaction, error) {
	return _RollupBase.Contract.Withdraw(&_RollupBase.TransactOpts)
}

// RollupBaseAssertionChallengedIterator is returned from FilterAssertionChallenged and is used to iterate over the raw logs and unpacked data for AssertionChallenged events raised by the RollupBase contract.
type RollupBaseAssertionChallengedIterator struct {
	Event *RollupBaseAssertionChallenged // Event containing the contract specifics and raw log

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
func (it *RollupBaseAssertionChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupBaseAssertionChallenged)
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
		it.Event = new(RollupBaseAssertionChallenged)
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
func (it *RollupBaseAssertionChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupBaseAssertionChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupBaseAssertionChallenged represents a AssertionChallenged event raised by the RollupBase contract.
type RollupBaseAssertionChallenged struct {
	AssertionID   *big.Int
	ChallengeAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAssertionChallenged is a free log retrieval operation binding the contract event 0xd0ebe74b4f7d89a9b0fdc9d95f887a7b925c6c7300b5c4b2c3304d97925840fa.
//
// Solidity: event AssertionChallenged(uint256 assertionID, address challengeAddr)
func (_RollupBase *RollupBaseFilterer) FilterAssertionChallenged(opts *bind.FilterOpts) (*RollupBaseAssertionChallengedIterator, error) {

	logs, sub, err := _RollupBase.contract.FilterLogs(opts, "AssertionChallenged")
	if err != nil {
		return nil, err
	}
	return &RollupBaseAssertionChallengedIterator{contract: _RollupBase.contract, event: "AssertionChallenged", logs: logs, sub: sub}, nil
}

// WatchAssertionChallenged is a free log subscription operation binding the contract event 0xd0ebe74b4f7d89a9b0fdc9d95f887a7b925c6c7300b5c4b2c3304d97925840fa.
//
// Solidity: event AssertionChallenged(uint256 assertionID, address challengeAddr)
func (_RollupBase *RollupBaseFilterer) WatchAssertionChallenged(opts *bind.WatchOpts, sink chan<- *RollupBaseAssertionChallenged) (event.Subscription, error) {

	logs, sub, err := _RollupBase.contract.WatchLogs(opts, "AssertionChallenged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupBaseAssertionChallenged)
				if err := _RollupBase.contract.UnpackLog(event, "AssertionChallenged", log); err != nil {
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
func (_RollupBase *RollupBaseFilterer) ParseAssertionChallenged(log types.Log) (*RollupBaseAssertionChallenged, error) {
	event := new(RollupBaseAssertionChallenged)
	if err := _RollupBase.contract.UnpackLog(event, "AssertionChallenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupBaseAssertionConfirmedIterator is returned from FilterAssertionConfirmed and is used to iterate over the raw logs and unpacked data for AssertionConfirmed events raised by the RollupBase contract.
type RollupBaseAssertionConfirmedIterator struct {
	Event *RollupBaseAssertionConfirmed // Event containing the contract specifics and raw log

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
func (it *RollupBaseAssertionConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupBaseAssertionConfirmed)
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
		it.Event = new(RollupBaseAssertionConfirmed)
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
func (it *RollupBaseAssertionConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupBaseAssertionConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupBaseAssertionConfirmed represents a AssertionConfirmed event raised by the RollupBase contract.
type RollupBaseAssertionConfirmed struct {
	AssertionID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssertionConfirmed is a free log retrieval operation binding the contract event 0x453430d123684340024ae0a229704bdab39c93dc48bb5a0b4bc83142d95d48ef.
//
// Solidity: event AssertionConfirmed(uint256 assertionID)
func (_RollupBase *RollupBaseFilterer) FilterAssertionConfirmed(opts *bind.FilterOpts) (*RollupBaseAssertionConfirmedIterator, error) {

	logs, sub, err := _RollupBase.contract.FilterLogs(opts, "AssertionConfirmed")
	if err != nil {
		return nil, err
	}
	return &RollupBaseAssertionConfirmedIterator{contract: _RollupBase.contract, event: "AssertionConfirmed", logs: logs, sub: sub}, nil
}

// WatchAssertionConfirmed is a free log subscription operation binding the contract event 0x453430d123684340024ae0a229704bdab39c93dc48bb5a0b4bc83142d95d48ef.
//
// Solidity: event AssertionConfirmed(uint256 assertionID)
func (_RollupBase *RollupBaseFilterer) WatchAssertionConfirmed(opts *bind.WatchOpts, sink chan<- *RollupBaseAssertionConfirmed) (event.Subscription, error) {

	logs, sub, err := _RollupBase.contract.WatchLogs(opts, "AssertionConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupBaseAssertionConfirmed)
				if err := _RollupBase.contract.UnpackLog(event, "AssertionConfirmed", log); err != nil {
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
func (_RollupBase *RollupBaseFilterer) ParseAssertionConfirmed(log types.Log) (*RollupBaseAssertionConfirmed, error) {
	event := new(RollupBaseAssertionConfirmed)
	if err := _RollupBase.contract.UnpackLog(event, "AssertionConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupBaseAssertionCreatedIterator is returned from FilterAssertionCreated and is used to iterate over the raw logs and unpacked data for AssertionCreated events raised by the RollupBase contract.
type RollupBaseAssertionCreatedIterator struct {
	Event *RollupBaseAssertionCreated // Event containing the contract specifics and raw log

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
func (it *RollupBaseAssertionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupBaseAssertionCreated)
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
		it.Event = new(RollupBaseAssertionCreated)
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
func (it *RollupBaseAssertionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupBaseAssertionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupBaseAssertionCreated represents a AssertionCreated event raised by the RollupBase contract.
type RollupBaseAssertionCreated struct {
	AssertionID  *big.Int
	AsserterAddr common.Address
	VmHash       [32]byte
	InboxSize    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssertionCreated is a free log retrieval operation binding the contract event 0x5c610f28399ecc14b66149012a0197a5e3257a8c397125afee95d1cf4b950734.
//
// Solidity: event AssertionCreated(uint256 assertionID, address asserterAddr, bytes32 vmHash, uint256 inboxSize)
func (_RollupBase *RollupBaseFilterer) FilterAssertionCreated(opts *bind.FilterOpts) (*RollupBaseAssertionCreatedIterator, error) {

	logs, sub, err := _RollupBase.contract.FilterLogs(opts, "AssertionCreated")
	if err != nil {
		return nil, err
	}
	return &RollupBaseAssertionCreatedIterator{contract: _RollupBase.contract, event: "AssertionCreated", logs: logs, sub: sub}, nil
}

// WatchAssertionCreated is a free log subscription operation binding the contract event 0x5c610f28399ecc14b66149012a0197a5e3257a8c397125afee95d1cf4b950734.
//
// Solidity: event AssertionCreated(uint256 assertionID, address asserterAddr, bytes32 vmHash, uint256 inboxSize)
func (_RollupBase *RollupBaseFilterer) WatchAssertionCreated(opts *bind.WatchOpts, sink chan<- *RollupBaseAssertionCreated) (event.Subscription, error) {

	logs, sub, err := _RollupBase.contract.WatchLogs(opts, "AssertionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupBaseAssertionCreated)
				if err := _RollupBase.contract.UnpackLog(event, "AssertionCreated", log); err != nil {
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
func (_RollupBase *RollupBaseFilterer) ParseAssertionCreated(log types.Log) (*RollupBaseAssertionCreated, error) {
	event := new(RollupBaseAssertionCreated)
	if err := _RollupBase.contract.UnpackLog(event, "AssertionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupBaseAssertionRejectedIterator is returned from FilterAssertionRejected and is used to iterate over the raw logs and unpacked data for AssertionRejected events raised by the RollupBase contract.
type RollupBaseAssertionRejectedIterator struct {
	Event *RollupBaseAssertionRejected // Event containing the contract specifics and raw log

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
func (it *RollupBaseAssertionRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupBaseAssertionRejected)
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
		it.Event = new(RollupBaseAssertionRejected)
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
func (it *RollupBaseAssertionRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupBaseAssertionRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupBaseAssertionRejected represents a AssertionRejected event raised by the RollupBase contract.
type RollupBaseAssertionRejected struct {
	AssertionID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssertionRejected is a free log retrieval operation binding the contract event 0x5b24ab8ceb442373727ac5c559a027521cb52db451c74710ebed9faa5fe15a7c.
//
// Solidity: event AssertionRejected(uint256 assertionID)
func (_RollupBase *RollupBaseFilterer) FilterAssertionRejected(opts *bind.FilterOpts) (*RollupBaseAssertionRejectedIterator, error) {

	logs, sub, err := _RollupBase.contract.FilterLogs(opts, "AssertionRejected")
	if err != nil {
		return nil, err
	}
	return &RollupBaseAssertionRejectedIterator{contract: _RollupBase.contract, event: "AssertionRejected", logs: logs, sub: sub}, nil
}

// WatchAssertionRejected is a free log subscription operation binding the contract event 0x5b24ab8ceb442373727ac5c559a027521cb52db451c74710ebed9faa5fe15a7c.
//
// Solidity: event AssertionRejected(uint256 assertionID)
func (_RollupBase *RollupBaseFilterer) WatchAssertionRejected(opts *bind.WatchOpts, sink chan<- *RollupBaseAssertionRejected) (event.Subscription, error) {

	logs, sub, err := _RollupBase.contract.WatchLogs(opts, "AssertionRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupBaseAssertionRejected)
				if err := _RollupBase.contract.UnpackLog(event, "AssertionRejected", log); err != nil {
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
func (_RollupBase *RollupBaseFilterer) ParseAssertionRejected(log types.Log) (*RollupBaseAssertionRejected, error) {
	event := new(RollupBaseAssertionRejected)
	if err := _RollupBase.contract.UnpackLog(event, "AssertionRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RollupBase contract.
type RollupBaseInitializedIterator struct {
	Event *RollupBaseInitialized // Event containing the contract specifics and raw log

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
func (it *RollupBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupBaseInitialized)
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
		it.Event = new(RollupBaseInitialized)
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
func (it *RollupBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupBaseInitialized represents a Initialized event raised by the RollupBase contract.
type RollupBaseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RollupBase *RollupBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*RollupBaseInitializedIterator, error) {

	logs, sub, err := _RollupBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RollupBaseInitializedIterator{contract: _RollupBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RollupBase *RollupBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RollupBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _RollupBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupBaseInitialized)
				if err := _RollupBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RollupBase *RollupBaseFilterer) ParseInitialized(log types.Log) (*RollupBaseInitialized, error) {
	event := new(RollupBaseInitialized)
	if err := _RollupBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupBaseStakerStakedIterator is returned from FilterStakerStaked and is used to iterate over the raw logs and unpacked data for StakerStaked events raised by the RollupBase contract.
type RollupBaseStakerStakedIterator struct {
	Event *RollupBaseStakerStaked // Event containing the contract specifics and raw log

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
func (it *RollupBaseStakerStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupBaseStakerStaked)
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
		it.Event = new(RollupBaseStakerStaked)
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
func (it *RollupBaseStakerStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupBaseStakerStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupBaseStakerStaked represents a StakerStaked event raised by the RollupBase contract.
type RollupBaseStakerStaked struct {
	StakerAddr  common.Address
	AssertionID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakerStaked is a free log retrieval operation binding the contract event 0x617d31491414a4ab2bd831e566a31837fa7fb6582921c91dffbbe83fbca789f3.
//
// Solidity: event StakerStaked(address stakerAddr, uint256 assertionID)
func (_RollupBase *RollupBaseFilterer) FilterStakerStaked(opts *bind.FilterOpts) (*RollupBaseStakerStakedIterator, error) {

	logs, sub, err := _RollupBase.contract.FilterLogs(opts, "StakerStaked")
	if err != nil {
		return nil, err
	}
	return &RollupBaseStakerStakedIterator{contract: _RollupBase.contract, event: "StakerStaked", logs: logs, sub: sub}, nil
}

// WatchStakerStaked is a free log subscription operation binding the contract event 0x617d31491414a4ab2bd831e566a31837fa7fb6582921c91dffbbe83fbca789f3.
//
// Solidity: event StakerStaked(address stakerAddr, uint256 assertionID)
func (_RollupBase *RollupBaseFilterer) WatchStakerStaked(opts *bind.WatchOpts, sink chan<- *RollupBaseStakerStaked) (event.Subscription, error) {

	logs, sub, err := _RollupBase.contract.WatchLogs(opts, "StakerStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupBaseStakerStaked)
				if err := _RollupBase.contract.UnpackLog(event, "StakerStaked", log); err != nil {
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
func (_RollupBase *RollupBaseFilterer) ParseStakerStaked(log types.Log) (*RollupBaseStakerStaked, error) {
	event := new(RollupBaseStakerStaked)
	if err := _RollupBase.contract.UnpackLog(event, "StakerStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// RollupMetaData contains all meta data concerning the Rollup contract.
var RollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AssertionAlreadyResolved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssertionOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengePeriodPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengedStaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DifferentParent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAssertion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker1Challenge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker2Challenge\",\"type\":\"address\"}],\"name\":\"InDifferentChallenge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InboxReadLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinimumAssertionPeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoStaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoUnresolvedAssertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInChallenge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParentAssertionUnstaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreviousStateHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedundantInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakedOnUnconfirmedAssertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakerStakedOnTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakersPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnproposedAssertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeAddr\",\"type\":\"address\"}],\"name\":\"AssertionChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"AssertionConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserterAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vmHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"}],\"name\":\"AssertionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"AssertionRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"StakerStaked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"advanceStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assertions\",\"outputs\":[{\"internalType\":\"contractAssertionMap\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"players\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"assertionIDs\",\"type\":\"uint256[2]\"}],\"name\":\"challengeAssertion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeCtx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"challengeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"defenderAssertionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerAssertionID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmFirstUnresolvedAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmationPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmedInboxSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vmHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"}],\"name\":\"createAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vmHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_batch\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_shouldStartAtElement\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"createAssertionWithStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_libAddressManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_assertionMap\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_confirmationPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumAssertionPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_initialVMhash\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastConfirmedAssertionID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCreatedAssertionID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastResolvedAssertionID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libAddressManager\",\"outputs\":[{\"internalType\":\"contractLib_AddressManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectFirstUnresolvedAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"removeStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifierEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"zombies\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastAssertionID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RollupABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupMetaData.ABI instead.
var RollupABI = RollupMetaData.ABI

// Rollup is an auto generated Go binding around an Ethereum contract.
type Rollup struct {
	RollupCaller     // Read-only binding to the contract
	RollupTransactor // Write-only binding to the contract
	RollupFilterer   // Log filterer for contract events
}

// RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupSession struct {
	Contract     *Rollup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCallerSession struct {
	Contract *RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTransactorSession struct {
	Contract     *RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupRaw struct {
	Contract *Rollup // Generic contract binding to access the raw methods on
}

// RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCallerRaw struct {
	Contract *RollupCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTransactorRaw struct {
	Contract *RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollup creates a new instance of Rollup, bound to a specific deployed contract.
func NewRollup(address common.Address, backend bind.ContractBackend) (*Rollup, error) {
	contract, err := bindRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// NewRollupCaller creates a new read-only instance of Rollup, bound to a specific deployed contract.
func NewRollupCaller(address common.Address, caller bind.ContractCaller) (*RollupCaller, error) {
	contract, err := bindRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCaller{contract: contract}, nil
}

// NewRollupTransactor creates a new write-only instance of Rollup, bound to a specific deployed contract.
func NewRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTransactor, error) {
	contract, err := bindRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTransactor{contract: contract}, nil
}

// NewRollupFilterer creates a new log filterer instance of Rollup, bound to a specific deployed contract.
func NewRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupFilterer, error) {
	contract, err := bindRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupFilterer{contract: contract}, nil
}

// bindRollup binds a generic wrapper to an already deployed contract.
func bindRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transact(opts, method, params...)
}

// Assertions is a free data retrieval call binding the contract method 0x40d9224b.
//
// Solidity: function assertions() view returns(address)
func (_Rollup *RollupCaller) Assertions(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "assertions")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Assertions is a free data retrieval call binding the contract method 0x40d9224b.
//
// Solidity: function assertions() view returns(address)
func (_Rollup *RollupSession) Assertions() (common.Address, error) {
	return _Rollup.Contract.Assertions(&_Rollup.CallOpts)
}

// Assertions is a free data retrieval call binding the contract method 0x40d9224b.
//
// Solidity: function assertions() view returns(address)
func (_Rollup *RollupCallerSession) Assertions() (common.Address, error) {
	return _Rollup.Contract.Assertions(&_Rollup.CallOpts)
}

// BaseStakeAmount is a free data retrieval call binding the contract method 0x71129559.
//
// Solidity: function baseStakeAmount() view returns(uint256)
func (_Rollup *RollupCaller) BaseStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "baseStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStakeAmount is a free data retrieval call binding the contract method 0x71129559.
//
// Solidity: function baseStakeAmount() view returns(uint256)
func (_Rollup *RollupSession) BaseStakeAmount() (*big.Int, error) {
	return _Rollup.Contract.BaseStakeAmount(&_Rollup.CallOpts)
}

// BaseStakeAmount is a free data retrieval call binding the contract method 0x71129559.
//
// Solidity: function baseStakeAmount() view returns(uint256)
func (_Rollup *RollupCallerSession) BaseStakeAmount() (*big.Int, error) {
	return _Rollup.Contract.BaseStakeAmount(&_Rollup.CallOpts)
}

// ChallengeCtx is a free data retrieval call binding the contract method 0x0e456acf.
//
// Solidity: function challengeCtx() view returns(bool completed, address challengeAddress, address defenderAddress, address challengerAddress, uint256 defenderAssertionID, uint256 challengerAssertionID)
func (_Rollup *RollupCaller) ChallengeCtx(opts *bind.CallOpts) (struct {
	Completed             bool
	ChallengeAddress      common.Address
	DefenderAddress       common.Address
	ChallengerAddress     common.Address
	DefenderAssertionID   *big.Int
	ChallengerAssertionID *big.Int
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "challengeCtx")

	outstruct := new(struct {
		Completed             bool
		ChallengeAddress      common.Address
		DefenderAddress       common.Address
		ChallengerAddress     common.Address
		DefenderAssertionID   *big.Int
		ChallengerAssertionID *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Completed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ChallengeAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.DefenderAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ChallengerAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.DefenderAssertionID = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ChallengerAssertionID = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ChallengeCtx is a free data retrieval call binding the contract method 0x0e456acf.
//
// Solidity: function challengeCtx() view returns(bool completed, address challengeAddress, address defenderAddress, address challengerAddress, uint256 defenderAssertionID, uint256 challengerAssertionID)
func (_Rollup *RollupSession) ChallengeCtx() (struct {
	Completed             bool
	ChallengeAddress      common.Address
	DefenderAddress       common.Address
	ChallengerAddress     common.Address
	DefenderAssertionID   *big.Int
	ChallengerAssertionID *big.Int
}, error) {
	return _Rollup.Contract.ChallengeCtx(&_Rollup.CallOpts)
}

// ChallengeCtx is a free data retrieval call binding the contract method 0x0e456acf.
//
// Solidity: function challengeCtx() view returns(bool completed, address challengeAddress, address defenderAddress, address challengerAddress, uint256 defenderAssertionID, uint256 challengerAssertionID)
func (_Rollup *RollupCallerSession) ChallengeCtx() (struct {
	Completed             bool
	ChallengeAddress      common.Address
	DefenderAddress       common.Address
	ChallengerAddress     common.Address
	DefenderAssertionID   *big.Int
	ChallengerAssertionID *big.Int
}, error) {
	return _Rollup.Contract.ChallengeCtx(&_Rollup.CallOpts)
}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() view returns(uint256)
func (_Rollup *RollupCaller) ChallengePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "challengePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() view returns(uint256)
func (_Rollup *RollupSession) ChallengePeriod() (*big.Int, error) {
	return _Rollup.Contract.ChallengePeriod(&_Rollup.CallOpts)
}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() view returns(uint256)
func (_Rollup *RollupCallerSession) ChallengePeriod() (*big.Int, error) {
	return _Rollup.Contract.ChallengePeriod(&_Rollup.CallOpts)
}

// ConfirmationPeriod is a free data retrieval call binding the contract method 0x0429b880.
//
// Solidity: function confirmationPeriod() view returns(uint256)
func (_Rollup *RollupCaller) ConfirmationPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "confirmationPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmationPeriod is a free data retrieval call binding the contract method 0x0429b880.
//
// Solidity: function confirmationPeriod() view returns(uint256)
func (_Rollup *RollupSession) ConfirmationPeriod() (*big.Int, error) {
	return _Rollup.Contract.ConfirmationPeriod(&_Rollup.CallOpts)
}

// ConfirmationPeriod is a free data retrieval call binding the contract method 0x0429b880.
//
// Solidity: function confirmationPeriod() view returns(uint256)
func (_Rollup *RollupCallerSession) ConfirmationPeriod() (*big.Int, error) {
	return _Rollup.Contract.ConfirmationPeriod(&_Rollup.CallOpts)
}

// ConfirmedInboxSize is a free data retrieval call binding the contract method 0xc94b5847.
//
// Solidity: function confirmedInboxSize() view returns(uint256)
func (_Rollup *RollupCaller) ConfirmedInboxSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "confirmedInboxSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmedInboxSize is a free data retrieval call binding the contract method 0xc94b5847.
//
// Solidity: function confirmedInboxSize() view returns(uint256)
func (_Rollup *RollupSession) ConfirmedInboxSize() (*big.Int, error) {
	return _Rollup.Contract.ConfirmedInboxSize(&_Rollup.CallOpts)
}

// ConfirmedInboxSize is a free data retrieval call binding the contract method 0xc94b5847.
//
// Solidity: function confirmedInboxSize() view returns(uint256)
func (_Rollup *RollupCallerSession) ConfirmedInboxSize() (*big.Int, error) {
	return _Rollup.Contract.ConfirmedInboxSize(&_Rollup.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupCaller) CurrentRequiredStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "currentRequiredStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupSession) CurrentRequiredStake() (*big.Int, error) {
	return _Rollup.Contract.CurrentRequiredStake(&_Rollup.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupCallerSession) CurrentRequiredStake() (*big.Int, error) {
	return _Rollup.Contract.CurrentRequiredStake(&_Rollup.CallOpts)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_Rollup *RollupCaller) IsStaked(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isStaked", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_Rollup *RollupSession) IsStaked(addr common.Address) (bool, error) {
	return _Rollup.Contract.IsStaked(&_Rollup.CallOpts, addr)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_Rollup *RollupCallerSession) IsStaked(addr common.Address) (bool, error) {
	return _Rollup.Contract.IsStaked(&_Rollup.CallOpts, addr)
}

// LastConfirmedAssertionID is a free data retrieval call binding the contract method 0xa56ba93b.
//
// Solidity: function lastConfirmedAssertionID() view returns(uint256)
func (_Rollup *RollupCaller) LastConfirmedAssertionID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "lastConfirmedAssertionID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastConfirmedAssertionID is a free data retrieval call binding the contract method 0xa56ba93b.
//
// Solidity: function lastConfirmedAssertionID() view returns(uint256)
func (_Rollup *RollupSession) LastConfirmedAssertionID() (*big.Int, error) {
	return _Rollup.Contract.LastConfirmedAssertionID(&_Rollup.CallOpts)
}

// LastConfirmedAssertionID is a free data retrieval call binding the contract method 0xa56ba93b.
//
// Solidity: function lastConfirmedAssertionID() view returns(uint256)
func (_Rollup *RollupCallerSession) LastConfirmedAssertionID() (*big.Int, error) {
	return _Rollup.Contract.LastConfirmedAssertionID(&_Rollup.CallOpts)
}

// LastCreatedAssertionID is a free data retrieval call binding the contract method 0x107035a4.
//
// Solidity: function lastCreatedAssertionID() view returns(uint256)
func (_Rollup *RollupCaller) LastCreatedAssertionID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "lastCreatedAssertionID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCreatedAssertionID is a free data retrieval call binding the contract method 0x107035a4.
//
// Solidity: function lastCreatedAssertionID() view returns(uint256)
func (_Rollup *RollupSession) LastCreatedAssertionID() (*big.Int, error) {
	return _Rollup.Contract.LastCreatedAssertionID(&_Rollup.CallOpts)
}

// LastCreatedAssertionID is a free data retrieval call binding the contract method 0x107035a4.
//
// Solidity: function lastCreatedAssertionID() view returns(uint256)
func (_Rollup *RollupCallerSession) LastCreatedAssertionID() (*big.Int, error) {
	return _Rollup.Contract.LastCreatedAssertionID(&_Rollup.CallOpts)
}

// LastResolvedAssertionID is a free data retrieval call binding the contract method 0xb553ee84.
//
// Solidity: function lastResolvedAssertionID() view returns(uint256)
func (_Rollup *RollupCaller) LastResolvedAssertionID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "lastResolvedAssertionID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastResolvedAssertionID is a free data retrieval call binding the contract method 0xb553ee84.
//
// Solidity: function lastResolvedAssertionID() view returns(uint256)
func (_Rollup *RollupSession) LastResolvedAssertionID() (*big.Int, error) {
	return _Rollup.Contract.LastResolvedAssertionID(&_Rollup.CallOpts)
}

// LastResolvedAssertionID is a free data retrieval call binding the contract method 0xb553ee84.
//
// Solidity: function lastResolvedAssertionID() view returns(uint256)
func (_Rollup *RollupCallerSession) LastResolvedAssertionID() (*big.Int, error) {
	return _Rollup.Contract.LastResolvedAssertionID(&_Rollup.CallOpts)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_Rollup *RollupCaller) LibAddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "libAddressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_Rollup *RollupSession) LibAddressManager() (common.Address, error) {
	return _Rollup.Contract.LibAddressManager(&_Rollup.CallOpts)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_Rollup *RollupCallerSession) LibAddressManager() (common.Address, error) {
	return _Rollup.Contract.LibAddressManager(&_Rollup.CallOpts)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_Rollup *RollupCaller) MinimumAssertionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "minimumAssertionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_Rollup *RollupSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _Rollup.Contract.MinimumAssertionPeriod(&_Rollup.CallOpts)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_Rollup *RollupCallerSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _Rollup.Contract.MinimumAssertionPeriod(&_Rollup.CallOpts)
}

// NumStakers is a free data retrieval call binding the contract method 0x6c8b052a.
//
// Solidity: function numStakers() view returns(uint256)
func (_Rollup *RollupCaller) NumStakers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "numStakers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakers is a free data retrieval call binding the contract method 0x6c8b052a.
//
// Solidity: function numStakers() view returns(uint256)
func (_Rollup *RollupSession) NumStakers() (*big.Int, error) {
	return _Rollup.Contract.NumStakers(&_Rollup.CallOpts)
}

// NumStakers is a free data retrieval call binding the contract method 0x6c8b052a.
//
// Solidity: function numStakers() view returns(uint256)
func (_Rollup *RollupCallerSession) NumStakers() (*big.Int, error) {
	return _Rollup.Contract.NumStakers(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCallerSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_Rollup *RollupCaller) Resolve(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "resolve", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_Rollup *RollupSession) Resolve(_name string) (common.Address, error) {
	return _Rollup.Contract.Resolve(&_Rollup.CallOpts, _name)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_Rollup *RollupCallerSession) Resolve(_name string) (common.Address, error) {
	return _Rollup.Contract.Resolve(&_Rollup.CallOpts, _name)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Rollup *RollupCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Rollup *RollupSession) StakeToken() (common.Address, error) {
	return _Rollup.Contract.StakeToken(&_Rollup.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Rollup *RollupCallerSession) StakeToken() (common.Address, error) {
	return _Rollup.Contract.StakeToken(&_Rollup.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bool isStaked, uint256 amountStaked, uint256 assertionID, address currentChallenge)
func (_Rollup *RollupCaller) Stakers(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsStaked         bool
	AmountStaked     *big.Int
	AssertionID      *big.Int
	CurrentChallenge common.Address
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakers", arg0)

	outstruct := new(struct {
		IsStaked         bool
		AmountStaked     *big.Int
		AssertionID      *big.Int
		CurrentChallenge common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsStaked = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.AmountStaked = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AssertionID = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentChallenge = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bool isStaked, uint256 amountStaked, uint256 assertionID, address currentChallenge)
func (_Rollup *RollupSession) Stakers(arg0 common.Address) (struct {
	IsStaked         bool
	AmountStaked     *big.Int
	AssertionID      *big.Int
	CurrentChallenge common.Address
}, error) {
	return _Rollup.Contract.Stakers(&_Rollup.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bool isStaked, uint256 amountStaked, uint256 assertionID, address currentChallenge)
func (_Rollup *RollupCallerSession) Stakers(arg0 common.Address) (struct {
	IsStaked         bool
	AmountStaked     *big.Int
	AssertionID      *big.Int
	CurrentChallenge common.Address
}, error) {
	return _Rollup.Contract.Stakers(&_Rollup.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Rollup *RollupCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Rollup *RollupSession) Verifier() (common.Address, error) {
	return _Rollup.Contract.Verifier(&_Rollup.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Rollup *RollupCallerSession) Verifier() (common.Address, error) {
	return _Rollup.Contract.Verifier(&_Rollup.CallOpts)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address ) view returns(uint256)
func (_Rollup *RollupCaller) WithdrawableFunds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "withdrawableFunds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address ) view returns(uint256)
func (_Rollup *RollupSession) WithdrawableFunds(arg0 common.Address) (*big.Int, error) {
	return _Rollup.Contract.WithdrawableFunds(&_Rollup.CallOpts, arg0)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address ) view returns(uint256)
func (_Rollup *RollupCallerSession) WithdrawableFunds(arg0 common.Address) (*big.Int, error) {
	return _Rollup.Contract.WithdrawableFunds(&_Rollup.CallOpts, arg0)
}

// Zombies is a free data retrieval call binding the contract method 0x2052465e.
//
// Solidity: function zombies(uint256 ) view returns(address stakerAddress, uint256 lastAssertionID)
func (_Rollup *RollupCaller) Zombies(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakerAddress   common.Address
	LastAssertionID *big.Int
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "zombies", arg0)

	outstruct := new(struct {
		StakerAddress   common.Address
		LastAssertionID *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakerAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LastAssertionID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Zombies is a free data retrieval call binding the contract method 0x2052465e.
//
// Solidity: function zombies(uint256 ) view returns(address stakerAddress, uint256 lastAssertionID)
func (_Rollup *RollupSession) Zombies(arg0 *big.Int) (struct {
	StakerAddress   common.Address
	LastAssertionID *big.Int
}, error) {
	return _Rollup.Contract.Zombies(&_Rollup.CallOpts, arg0)
}

// Zombies is a free data retrieval call binding the contract method 0x2052465e.
//
// Solidity: function zombies(uint256 ) view returns(address stakerAddress, uint256 lastAssertionID)
func (_Rollup *RollupCallerSession) Zombies(arg0 *big.Int) (struct {
	StakerAddress   common.Address
	LastAssertionID *big.Int
}, error) {
	return _Rollup.Contract.Zombies(&_Rollup.CallOpts, arg0)
}

// AdvanceStake is a paid mutator transaction binding the contract method 0x8821b2ae.
//
// Solidity: function advanceStake(uint256 assertionID) returns()
func (_Rollup *RollupTransactor) AdvanceStake(opts *bind.TransactOpts, assertionID *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "advanceStake", assertionID)
}

// AdvanceStake is a paid mutator transaction binding the contract method 0x8821b2ae.
//
// Solidity: function advanceStake(uint256 assertionID) returns()
func (_Rollup *RollupSession) AdvanceStake(assertionID *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.AdvanceStake(&_Rollup.TransactOpts, assertionID)
}

// AdvanceStake is a paid mutator transaction binding the contract method 0x8821b2ae.
//
// Solidity: function advanceStake(uint256 assertionID) returns()
func (_Rollup *RollupTransactorSession) AdvanceStake(assertionID *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.AdvanceStake(&_Rollup.TransactOpts, assertionID)
}

// ChallengeAssertion is a paid mutator transaction binding the contract method 0x2f06d1b0.
//
// Solidity: function challengeAssertion(address[2] players, uint256[2] assertionIDs) returns(address)
func (_Rollup *RollupTransactor) ChallengeAssertion(opts *bind.TransactOpts, players [2]common.Address, assertionIDs [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "challengeAssertion", players, assertionIDs)
}

// ChallengeAssertion is a paid mutator transaction binding the contract method 0x2f06d1b0.
//
// Solidity: function challengeAssertion(address[2] players, uint256[2] assertionIDs) returns(address)
func (_Rollup *RollupSession) ChallengeAssertion(players [2]common.Address, assertionIDs [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ChallengeAssertion(&_Rollup.TransactOpts, players, assertionIDs)
}

// ChallengeAssertion is a paid mutator transaction binding the contract method 0x2f06d1b0.
//
// Solidity: function challengeAssertion(address[2] players, uint256[2] assertionIDs) returns(address)
func (_Rollup *RollupTransactorSession) ChallengeAssertion(players [2]common.Address, assertionIDs [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ChallengeAssertion(&_Rollup.TransactOpts, players, assertionIDs)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winner, address loser) returns()
func (_Rollup *RollupTransactor) CompleteChallenge(opts *bind.TransactOpts, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "completeChallenge", winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winner, address loser) returns()
func (_Rollup *RollupSession) CompleteChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.CompleteChallenge(&_Rollup.TransactOpts, winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winner, address loser) returns()
func (_Rollup *RollupTransactorSession) CompleteChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.CompleteChallenge(&_Rollup.TransactOpts, winner, loser)
}

// ConfirmFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x2906040e.
//
// Solidity: function confirmFirstUnresolvedAssertion() returns()
func (_Rollup *RollupTransactor) ConfirmFirstUnresolvedAssertion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "confirmFirstUnresolvedAssertion")
}

// ConfirmFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x2906040e.
//
// Solidity: function confirmFirstUnresolvedAssertion() returns()
func (_Rollup *RollupSession) ConfirmFirstUnresolvedAssertion() (*types.Transaction, error) {
	return _Rollup.Contract.ConfirmFirstUnresolvedAssertion(&_Rollup.TransactOpts)
}

// ConfirmFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x2906040e.
//
// Solidity: function confirmFirstUnresolvedAssertion() returns()
func (_Rollup *RollupTransactorSession) ConfirmFirstUnresolvedAssertion() (*types.Transaction, error) {
	return _Rollup.Contract.ConfirmFirstUnresolvedAssertion(&_Rollup.TransactOpts)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0xb6da898f.
//
// Solidity: function createAssertion(bytes32 vmHash, uint256 inboxSize) returns()
func (_Rollup *RollupTransactor) CreateAssertion(opts *bind.TransactOpts, vmHash [32]byte, inboxSize *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "createAssertion", vmHash, inboxSize)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0xb6da898f.
//
// Solidity: function createAssertion(bytes32 vmHash, uint256 inboxSize) returns()
func (_Rollup *RollupSession) CreateAssertion(vmHash [32]byte, inboxSize *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateAssertion(&_Rollup.TransactOpts, vmHash, inboxSize)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0xb6da898f.
//
// Solidity: function createAssertion(bytes32 vmHash, uint256 inboxSize) returns()
func (_Rollup *RollupTransactorSession) CreateAssertion(vmHash [32]byte, inboxSize *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateAssertion(&_Rollup.TransactOpts, vmHash, inboxSize)
}

// CreateAssertionWithStateBatch is a paid mutator transaction binding the contract method 0x49cd3004.
//
// Solidity: function createAssertionWithStateBatch(bytes32 vmHash, uint256 inboxSize, bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_Rollup *RollupTransactor) CreateAssertionWithStateBatch(opts *bind.TransactOpts, vmHash [32]byte, inboxSize *big.Int, _batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "createAssertionWithStateBatch", vmHash, inboxSize, _batch, _shouldStartAtElement, _signature)
}

// CreateAssertionWithStateBatch is a paid mutator transaction binding the contract method 0x49cd3004.
//
// Solidity: function createAssertionWithStateBatch(bytes32 vmHash, uint256 inboxSize, bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_Rollup *RollupSession) CreateAssertionWithStateBatch(vmHash [32]byte, inboxSize *big.Int, _batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Rollup.Contract.CreateAssertionWithStateBatch(&_Rollup.TransactOpts, vmHash, inboxSize, _batch, _shouldStartAtElement, _signature)
}

// CreateAssertionWithStateBatch is a paid mutator transaction binding the contract method 0x49cd3004.
//
// Solidity: function createAssertionWithStateBatch(bytes32 vmHash, uint256 inboxSize, bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_Rollup *RollupTransactorSession) CreateAssertionWithStateBatch(vmHash [32]byte, inboxSize *big.Int, _batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Rollup.Contract.CreateAssertionWithStateBatch(&_Rollup.TransactOpts, vmHash, inboxSize, _batch, _shouldStartAtElement, _signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xce4e5ffd.
//
// Solidity: function initialize(address _owner, address _verifier, address _stakeToken, address _libAddressManager, address _assertionMap, uint256 _confirmationPeriod, uint256 _challengePeriod, uint256 _minimumAssertionPeriod, uint256 _baseStakeAmount, bytes32 _initialVMhash) returns()
func (_Rollup *RollupTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _verifier common.Address, _stakeToken common.Address, _libAddressManager common.Address, _assertionMap common.Address, _confirmationPeriod *big.Int, _challengePeriod *big.Int, _minimumAssertionPeriod *big.Int, _baseStakeAmount *big.Int, _initialVMhash [32]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize", _owner, _verifier, _stakeToken, _libAddressManager, _assertionMap, _confirmationPeriod, _challengePeriod, _minimumAssertionPeriod, _baseStakeAmount, _initialVMhash)
}

// Initialize is a paid mutator transaction binding the contract method 0xce4e5ffd.
//
// Solidity: function initialize(address _owner, address _verifier, address _stakeToken, address _libAddressManager, address _assertionMap, uint256 _confirmationPeriod, uint256 _challengePeriod, uint256 _minimumAssertionPeriod, uint256 _baseStakeAmount, bytes32 _initialVMhash) returns()
func (_Rollup *RollupSession) Initialize(_owner common.Address, _verifier common.Address, _stakeToken common.Address, _libAddressManager common.Address, _assertionMap common.Address, _confirmationPeriod *big.Int, _challengePeriod *big.Int, _minimumAssertionPeriod *big.Int, _baseStakeAmount *big.Int, _initialVMhash [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _owner, _verifier, _stakeToken, _libAddressManager, _assertionMap, _confirmationPeriod, _challengePeriod, _minimumAssertionPeriod, _baseStakeAmount, _initialVMhash)
}

// Initialize is a paid mutator transaction binding the contract method 0xce4e5ffd.
//
// Solidity: function initialize(address _owner, address _verifier, address _stakeToken, address _libAddressManager, address _assertionMap, uint256 _confirmationPeriod, uint256 _challengePeriod, uint256 _minimumAssertionPeriod, uint256 _baseStakeAmount, bytes32 _initialVMhash) returns()
func (_Rollup *RollupTransactorSession) Initialize(_owner common.Address, _verifier common.Address, _stakeToken common.Address, _libAddressManager common.Address, _assertionMap common.Address, _confirmationPeriod *big.Int, _challengePeriod *big.Int, _minimumAssertionPeriod *big.Int, _baseStakeAmount *big.Int, _initialVMhash [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _owner, _verifier, _stakeToken, _libAddressManager, _assertionMap, _confirmationPeriod, _challengePeriod, _minimumAssertionPeriod, _baseStakeAmount, _initialVMhash)
}

// RejectFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x042dca93.
//
// Solidity: function rejectFirstUnresolvedAssertion(address stakerAddress) returns()
func (_Rollup *RollupTransactor) RejectFirstUnresolvedAssertion(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "rejectFirstUnresolvedAssertion", stakerAddress)
}

// RejectFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x042dca93.
//
// Solidity: function rejectFirstUnresolvedAssertion(address stakerAddress) returns()
func (_Rollup *RollupSession) RejectFirstUnresolvedAssertion(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RejectFirstUnresolvedAssertion(&_Rollup.TransactOpts, stakerAddress)
}

// RejectFirstUnresolvedAssertion is a paid mutator transaction binding the contract method 0x042dca93.
//
// Solidity: function rejectFirstUnresolvedAssertion(address stakerAddress) returns()
func (_Rollup *RollupTransactorSession) RejectFirstUnresolvedAssertion(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RejectFirstUnresolvedAssertion(&_Rollup.TransactOpts, stakerAddress)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address stakerAddress) returns()
func (_Rollup *RollupTransactor) RemoveStake(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeStake", stakerAddress)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address stakerAddress) returns()
func (_Rollup *RollupSession) RemoveStake(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveStake(&_Rollup.TransactOpts, stakerAddress)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address stakerAddress) returns()
func (_Rollup *RollupTransactorSession) RemoveStake(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveStake(&_Rollup.TransactOpts, stakerAddress)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Rollup *RollupTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Rollup *RollupSession) Stake() (*types.Transaction, error) {
	return _Rollup.Contract.Stake(&_Rollup.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Rollup *RollupTransactorSession) Stake() (*types.Transaction, error) {
	return _Rollup.Contract.Stake(&_Rollup.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeAmount) returns()
func (_Rollup *RollupTransactor) Unstake(opts *bind.TransactOpts, stakeAmount *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "unstake", stakeAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeAmount) returns()
func (_Rollup *RollupSession) Unstake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Unstake(&_Rollup.TransactOpts, stakeAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeAmount) returns()
func (_Rollup *RollupTransactorSession) Unstake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Unstake(&_Rollup.TransactOpts, stakeAmount)
}

// RollupAssertionChallengedIterator is returned from FilterAssertionChallenged and is used to iterate over the raw logs and unpacked data for AssertionChallenged events raised by the Rollup contract.
type RollupAssertionChallengedIterator struct {
	Event *RollupAssertionChallenged // Event containing the contract specifics and raw log

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
func (it *RollupAssertionChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAssertionChallenged)
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
		it.Event = new(RollupAssertionChallenged)
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
func (it *RollupAssertionChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAssertionChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAssertionChallenged represents a AssertionChallenged event raised by the Rollup contract.
type RollupAssertionChallenged struct {
	AssertionID   *big.Int
	ChallengeAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAssertionChallenged is a free log retrieval operation binding the contract event 0xd0ebe74b4f7d89a9b0fdc9d95f887a7b925c6c7300b5c4b2c3304d97925840fa.
//
// Solidity: event AssertionChallenged(uint256 assertionID, address challengeAddr)
func (_Rollup *RollupFilterer) FilterAssertionChallenged(opts *bind.FilterOpts) (*RollupAssertionChallengedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "AssertionChallenged")
	if err != nil {
		return nil, err
	}
	return &RollupAssertionChallengedIterator{contract: _Rollup.contract, event: "AssertionChallenged", logs: logs, sub: sub}, nil
}

// WatchAssertionChallenged is a free log subscription operation binding the contract event 0xd0ebe74b4f7d89a9b0fdc9d95f887a7b925c6c7300b5c4b2c3304d97925840fa.
//
// Solidity: event AssertionChallenged(uint256 assertionID, address challengeAddr)
func (_Rollup *RollupFilterer) WatchAssertionChallenged(opts *bind.WatchOpts, sink chan<- *RollupAssertionChallenged) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "AssertionChallenged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAssertionChallenged)
				if err := _Rollup.contract.UnpackLog(event, "AssertionChallenged", log); err != nil {
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
func (_Rollup *RollupFilterer) ParseAssertionChallenged(log types.Log) (*RollupAssertionChallenged, error) {
	event := new(RollupAssertionChallenged)
	if err := _Rollup.contract.UnpackLog(event, "AssertionChallenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAssertionConfirmedIterator is returned from FilterAssertionConfirmed and is used to iterate over the raw logs and unpacked data for AssertionConfirmed events raised by the Rollup contract.
type RollupAssertionConfirmedIterator struct {
	Event *RollupAssertionConfirmed // Event containing the contract specifics and raw log

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
func (it *RollupAssertionConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAssertionConfirmed)
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
		it.Event = new(RollupAssertionConfirmed)
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
func (it *RollupAssertionConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAssertionConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAssertionConfirmed represents a AssertionConfirmed event raised by the Rollup contract.
type RollupAssertionConfirmed struct {
	AssertionID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssertionConfirmed is a free log retrieval operation binding the contract event 0x453430d123684340024ae0a229704bdab39c93dc48bb5a0b4bc83142d95d48ef.
//
// Solidity: event AssertionConfirmed(uint256 assertionID)
func (_Rollup *RollupFilterer) FilterAssertionConfirmed(opts *bind.FilterOpts) (*RollupAssertionConfirmedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "AssertionConfirmed")
	if err != nil {
		return nil, err
	}
	return &RollupAssertionConfirmedIterator{contract: _Rollup.contract, event: "AssertionConfirmed", logs: logs, sub: sub}, nil
}

// WatchAssertionConfirmed is a free log subscription operation binding the contract event 0x453430d123684340024ae0a229704bdab39c93dc48bb5a0b4bc83142d95d48ef.
//
// Solidity: event AssertionConfirmed(uint256 assertionID)
func (_Rollup *RollupFilterer) WatchAssertionConfirmed(opts *bind.WatchOpts, sink chan<- *RollupAssertionConfirmed) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "AssertionConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAssertionConfirmed)
				if err := _Rollup.contract.UnpackLog(event, "AssertionConfirmed", log); err != nil {
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
func (_Rollup *RollupFilterer) ParseAssertionConfirmed(log types.Log) (*RollupAssertionConfirmed, error) {
	event := new(RollupAssertionConfirmed)
	if err := _Rollup.contract.UnpackLog(event, "AssertionConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAssertionCreatedIterator is returned from FilterAssertionCreated and is used to iterate over the raw logs and unpacked data for AssertionCreated events raised by the Rollup contract.
type RollupAssertionCreatedIterator struct {
	Event *RollupAssertionCreated // Event containing the contract specifics and raw log

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
func (it *RollupAssertionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAssertionCreated)
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
		it.Event = new(RollupAssertionCreated)
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
func (it *RollupAssertionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAssertionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAssertionCreated represents a AssertionCreated event raised by the Rollup contract.
type RollupAssertionCreated struct {
	AssertionID  *big.Int
	AsserterAddr common.Address
	VmHash       [32]byte
	InboxSize    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssertionCreated is a free log retrieval operation binding the contract event 0x5c610f28399ecc14b66149012a0197a5e3257a8c397125afee95d1cf4b950734.
//
// Solidity: event AssertionCreated(uint256 assertionID, address asserterAddr, bytes32 vmHash, uint256 inboxSize)
func (_Rollup *RollupFilterer) FilterAssertionCreated(opts *bind.FilterOpts) (*RollupAssertionCreatedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "AssertionCreated")
	if err != nil {
		return nil, err
	}
	return &RollupAssertionCreatedIterator{contract: _Rollup.contract, event: "AssertionCreated", logs: logs, sub: sub}, nil
}

// WatchAssertionCreated is a free log subscription operation binding the contract event 0x5c610f28399ecc14b66149012a0197a5e3257a8c397125afee95d1cf4b950734.
//
// Solidity: event AssertionCreated(uint256 assertionID, address asserterAddr, bytes32 vmHash, uint256 inboxSize)
func (_Rollup *RollupFilterer) WatchAssertionCreated(opts *bind.WatchOpts, sink chan<- *RollupAssertionCreated) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "AssertionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAssertionCreated)
				if err := _Rollup.contract.UnpackLog(event, "AssertionCreated", log); err != nil {
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
func (_Rollup *RollupFilterer) ParseAssertionCreated(log types.Log) (*RollupAssertionCreated, error) {
	event := new(RollupAssertionCreated)
	if err := _Rollup.contract.UnpackLog(event, "AssertionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAssertionRejectedIterator is returned from FilterAssertionRejected and is used to iterate over the raw logs and unpacked data for AssertionRejected events raised by the Rollup contract.
type RollupAssertionRejectedIterator struct {
	Event *RollupAssertionRejected // Event containing the contract specifics and raw log

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
func (it *RollupAssertionRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAssertionRejected)
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
		it.Event = new(RollupAssertionRejected)
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
func (it *RollupAssertionRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAssertionRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAssertionRejected represents a AssertionRejected event raised by the Rollup contract.
type RollupAssertionRejected struct {
	AssertionID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssertionRejected is a free log retrieval operation binding the contract event 0x5b24ab8ceb442373727ac5c559a027521cb52db451c74710ebed9faa5fe15a7c.
//
// Solidity: event AssertionRejected(uint256 assertionID)
func (_Rollup *RollupFilterer) FilterAssertionRejected(opts *bind.FilterOpts) (*RollupAssertionRejectedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "AssertionRejected")
	if err != nil {
		return nil, err
	}
	return &RollupAssertionRejectedIterator{contract: _Rollup.contract, event: "AssertionRejected", logs: logs, sub: sub}, nil
}

// WatchAssertionRejected is a free log subscription operation binding the contract event 0x5b24ab8ceb442373727ac5c559a027521cb52db451c74710ebed9faa5fe15a7c.
//
// Solidity: event AssertionRejected(uint256 assertionID)
func (_Rollup *RollupFilterer) WatchAssertionRejected(opts *bind.WatchOpts, sink chan<- *RollupAssertionRejected) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "AssertionRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAssertionRejected)
				if err := _Rollup.contract.UnpackLog(event, "AssertionRejected", log); err != nil {
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
func (_Rollup *RollupFilterer) ParseAssertionRejected(log types.Log) (*RollupAssertionRejected, error) {
	event := new(RollupAssertionRejected)
	if err := _Rollup.contract.UnpackLog(event, "AssertionRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Rollup contract.
type RollupInitializedIterator struct {
	Event *RollupInitialized // Event containing the contract specifics and raw log

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
func (it *RollupInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupInitialized)
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
		it.Event = new(RollupInitialized)
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
func (it *RollupInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupInitialized represents a Initialized event raised by the Rollup contract.
type RollupInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Rollup *RollupFilterer) FilterInitialized(opts *bind.FilterOpts) (*RollupInitializedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RollupInitializedIterator{contract: _Rollup.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Rollup *RollupFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RollupInitialized) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupInitialized)
				if err := _Rollup.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Rollup *RollupFilterer) ParseInitialized(log types.Log) (*RollupInitialized, error) {
	event := new(RollupInitialized)
	if err := _Rollup.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupStakerStakedIterator is returned from FilterStakerStaked and is used to iterate over the raw logs and unpacked data for StakerStaked events raised by the Rollup contract.
type RollupStakerStakedIterator struct {
	Event *RollupStakerStaked // Event containing the contract specifics and raw log

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
func (it *RollupStakerStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupStakerStaked)
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
		it.Event = new(RollupStakerStaked)
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
func (it *RollupStakerStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupStakerStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupStakerStaked represents a StakerStaked event raised by the Rollup contract.
type RollupStakerStaked struct {
	StakerAddr  common.Address
	AssertionID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakerStaked is a free log retrieval operation binding the contract event 0x617d31491414a4ab2bd831e566a31837fa7fb6582921c91dffbbe83fbca789f3.
//
// Solidity: event StakerStaked(address stakerAddr, uint256 assertionID)
func (_Rollup *RollupFilterer) FilterStakerStaked(opts *bind.FilterOpts) (*RollupStakerStakedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "StakerStaked")
	if err != nil {
		return nil, err
	}
	return &RollupStakerStakedIterator{contract: _Rollup.contract, event: "StakerStaked", logs: logs, sub: sub}, nil
}

// WatchStakerStaked is a free log subscription operation binding the contract event 0x617d31491414a4ab2bd831e566a31837fa7fb6582921c91dffbbe83fbca789f3.
//
// Solidity: event StakerStaked(address stakerAddr, uint256 assertionID)
func (_Rollup *RollupFilterer) WatchStakerStaked(opts *bind.WatchOpts, sink chan<- *RollupStakerStaked) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "StakerStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupStakerStaked)
				if err := _Rollup.contract.UnpackLog(event, "StakerStaked", log); err != nil {
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
func (_Rollup *RollupFilterer) ParseStakerStaked(log types.Log) (*RollupStakerStaked, error) {
	event := new(RollupStakerStaked)
	if err := _Rollup.contract.UnpackLog(event, "StakerStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

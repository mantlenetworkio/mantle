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

// EVMTypesLibTransaction is an auto generated low-level Go binding around an user-defined struct.
type EVMTypesLibTransaction struct {
	Nonce    uint64
	GasPrice *big.Int
	Gas      uint64
	To       common.Address
	Value    *big.Int
	Data     []byte
	V        *big.Int
	R        *big.Int
	S        *big.Int
}

// VerificationContextContext is an auto generated low-level Go binding around an user-defined struct.
type VerificationContextContext struct {
	Coinbase    common.Address
	Timestamp   *big.Int
	Number      *big.Int
	Origin      common.Address
	Transaction EVMTypesLibTransaction
	InputRoot   [32]byte
	TxHash      [32]byte
}

// ChallengeMetaData contains all meta data concerning the Challenge contract.
var ChallengeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"startState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"midState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"endState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIChallenge.CompletionReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"ChallengeCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"bisection\",\"type\":\"bytes32[3]\"},{\"internalType\":\"uint256\",\"name\":\"challengedSegmentIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevChallengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevChallengedSegmentLength\",\"type\":\"uint256\"}],\"name\":\"bisectExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bisectionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengerTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentBisected\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"startState\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"midState\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endState\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentResponderTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defenderTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierEntry\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resultReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startInboxSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_startStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_endStateHash\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"checkStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_numSteps\",\"type\":\"uint256\"}],\"name\":\"initializeChallengeLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMoveBlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prevBisection\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setRollback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startInboxSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"turn\",\"outputs\":[{\"internalType\":\"enumChallenge.Turn\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structEVMTypesLib.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"inputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerificationContext.Context\",\"name\":\"ctx\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"verifyType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"prevChallengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevChallengedSegmentLength\",\"type\":\"uint256\"}],\"name\":\"verifyOneStepProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ChallengeABI is the input ABI used to generate the binding from.
// Deprecated: Use ChallengeMetaData.ABI instead.
var ChallengeABI = ChallengeMetaData.ABI

// Challenge is an auto generated Go binding around an Ethereum contract.
type Challenge struct {
	ChallengeCaller     // Read-only binding to the contract
	ChallengeTransactor // Write-only binding to the contract
	ChallengeFilterer   // Log filterer for contract events
}

// ChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeSession struct {
	Contract     *Challenge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeCallerSession struct {
	Contract *ChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTransactorSession struct {
	Contract     *ChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeRaw struct {
	Contract *Challenge // Generic contract binding to access the raw methods on
}

// ChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeCallerRaw struct {
	Contract *ChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTransactorRaw struct {
	Contract *ChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallenge creates a new instance of Challenge, bound to a specific deployed contract.
func NewChallenge(address common.Address, backend bind.ContractBackend) (*Challenge, error) {
	contract, err := bindChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// NewChallengeCaller creates a new read-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeCaller(address common.Address, caller bind.ContractCaller) (*ChallengeCaller, error) {
	contract, err := bindChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeCaller{contract: contract}, nil
}

// NewChallengeTransactor creates a new write-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTransactor, error) {
	contract, err := bindChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTransactor{contract: contract}, nil
}

// NewChallengeFilterer creates a new log filterer instance of Challenge, bound to a specific deployed contract.
func NewChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeFilterer, error) {
	contract, err := bindChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeFilterer{contract: contract}, nil
}

// bindChallenge binds a generic wrapper to an already deployed contract.
func bindChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.ChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transact(opts, method, params...)
}

// BisectionHash is a free data retrieval call binding the contract method 0x2a51f6f7.
//
// Solidity: function bisectionHash() view returns(bytes32)
func (_Challenge *ChallengeCaller) BisectionHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "bisectionHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BisectionHash is a free data retrieval call binding the contract method 0x2a51f6f7.
//
// Solidity: function bisectionHash() view returns(bytes32)
func (_Challenge *ChallengeSession) BisectionHash() ([32]byte, error) {
	return _Challenge.Contract.BisectionHash(&_Challenge.CallOpts)
}

// BisectionHash is a free data retrieval call binding the contract method 0x2a51f6f7.
//
// Solidity: function bisectionHash() view returns(bytes32)
func (_Challenge *ChallengeCallerSession) BisectionHash() ([32]byte, error) {
	return _Challenge.Contract.BisectionHash(&_Challenge.CallOpts)
}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_Challenge *ChallengeCaller) Challenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "challenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_Challenge *ChallengeSession) Challenger() (common.Address, error) {
	return _Challenge.Contract.Challenger(&_Challenge.CallOpts)
}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_Challenge *ChallengeCallerSession) Challenger() (common.Address, error) {
	return _Challenge.Contract.Challenger(&_Challenge.CallOpts)
}

// ChallengerTimeLeft is a free data retrieval call binding the contract method 0x41e8510c.
//
// Solidity: function challengerTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCaller) ChallengerTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "challengerTimeLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengerTimeLeft is a free data retrieval call binding the contract method 0x41e8510c.
//
// Solidity: function challengerTimeLeft() view returns(uint256)
func (_Challenge *ChallengeSession) ChallengerTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.ChallengerTimeLeft(&_Challenge.CallOpts)
}

// ChallengerTimeLeft is a free data retrieval call binding the contract method 0x41e8510c.
//
// Solidity: function challengerTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCallerSession) ChallengerTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.ChallengerTimeLeft(&_Challenge.CallOpts)
}

// CurrentBisected is a free data retrieval call binding the contract method 0x732e6961.
//
// Solidity: function currentBisected() view returns(bytes32 startState, bytes32 midState, bytes32 endState, uint256 blockNum, uint256 blockTime, uint256 challengedSegmentStart, uint256 challengedSegmentLength)
func (_Challenge *ChallengeCaller) CurrentBisected(opts *bind.CallOpts) (struct {
	StartState              [32]byte
	MidState                [32]byte
	EndState                [32]byte
	BlockNum                *big.Int
	BlockTime               *big.Int
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
}, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "currentBisected")

	outstruct := new(struct {
		StartState              [32]byte
		MidState                [32]byte
		EndState                [32]byte
		BlockNum                *big.Int
		BlockTime               *big.Int
		ChallengedSegmentStart  *big.Int
		ChallengedSegmentLength *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartState = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.MidState = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.EndState = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.BlockNum = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BlockTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ChallengedSegmentStart = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ChallengedSegmentLength = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentBisected is a free data retrieval call binding the contract method 0x732e6961.
//
// Solidity: function currentBisected() view returns(bytes32 startState, bytes32 midState, bytes32 endState, uint256 blockNum, uint256 blockTime, uint256 challengedSegmentStart, uint256 challengedSegmentLength)
func (_Challenge *ChallengeSession) CurrentBisected() (struct {
	StartState              [32]byte
	MidState                [32]byte
	EndState                [32]byte
	BlockNum                *big.Int
	BlockTime               *big.Int
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
}, error) {
	return _Challenge.Contract.CurrentBisected(&_Challenge.CallOpts)
}

// CurrentBisected is a free data retrieval call binding the contract method 0x732e6961.
//
// Solidity: function currentBisected() view returns(bytes32 startState, bytes32 midState, bytes32 endState, uint256 blockNum, uint256 blockTime, uint256 challengedSegmentStart, uint256 challengedSegmentLength)
func (_Challenge *ChallengeCallerSession) CurrentBisected() (struct {
	StartState              [32]byte
	MidState                [32]byte
	EndState                [32]byte
	BlockNum                *big.Int
	BlockTime               *big.Int
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
}, error) {
	return _Challenge.Contract.CurrentBisected(&_Challenge.CallOpts)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x8a8cd218.
//
// Solidity: function currentResponder() view returns(address)
func (_Challenge *ChallengeCaller) CurrentResponder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "currentResponder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x8a8cd218.
//
// Solidity: function currentResponder() view returns(address)
func (_Challenge *ChallengeSession) CurrentResponder() (common.Address, error) {
	return _Challenge.Contract.CurrentResponder(&_Challenge.CallOpts)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x8a8cd218.
//
// Solidity: function currentResponder() view returns(address)
func (_Challenge *ChallengeCallerSession) CurrentResponder() (common.Address, error) {
	return _Challenge.Contract.CurrentResponder(&_Challenge.CallOpts)
}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xe87e3589.
//
// Solidity: function currentResponderTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCaller) CurrentResponderTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "currentResponderTimeLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xe87e3589.
//
// Solidity: function currentResponderTimeLeft() view returns(uint256)
func (_Challenge *ChallengeSession) CurrentResponderTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.CurrentResponderTimeLeft(&_Challenge.CallOpts)
}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xe87e3589.
//
// Solidity: function currentResponderTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCallerSession) CurrentResponderTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.CurrentResponderTimeLeft(&_Challenge.CallOpts)
}

// Defender is a free data retrieval call binding the contract method 0x7f4c91c5.
//
// Solidity: function defender() view returns(address)
func (_Challenge *ChallengeCaller) Defender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "defender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Defender is a free data retrieval call binding the contract method 0x7f4c91c5.
//
// Solidity: function defender() view returns(address)
func (_Challenge *ChallengeSession) Defender() (common.Address, error) {
	return _Challenge.Contract.Defender(&_Challenge.CallOpts)
}

// Defender is a free data retrieval call binding the contract method 0x7f4c91c5.
//
// Solidity: function defender() view returns(address)
func (_Challenge *ChallengeCallerSession) Defender() (common.Address, error) {
	return _Challenge.Contract.Defender(&_Challenge.CallOpts)
}

// DefenderTimeLeft is a free data retrieval call binding the contract method 0x5f41e3d6.
//
// Solidity: function defenderTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCaller) DefenderTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "defenderTimeLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefenderTimeLeft is a free data retrieval call binding the contract method 0x5f41e3d6.
//
// Solidity: function defenderTimeLeft() view returns(uint256)
func (_Challenge *ChallengeSession) DefenderTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.DefenderTimeLeft(&_Challenge.CallOpts)
}

// DefenderTimeLeft is a free data retrieval call binding the contract method 0x5f41e3d6.
//
// Solidity: function defenderTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCallerSession) DefenderTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.DefenderTimeLeft(&_Challenge.CallOpts)
}

// LastMoveBlockTime is a free data retrieval call binding the contract method 0xed5b1303.
//
// Solidity: function lastMoveBlockTime() view returns(uint256)
func (_Challenge *ChallengeCaller) LastMoveBlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "lastMoveBlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMoveBlockTime is a free data retrieval call binding the contract method 0xed5b1303.
//
// Solidity: function lastMoveBlockTime() view returns(uint256)
func (_Challenge *ChallengeSession) LastMoveBlockTime() (*big.Int, error) {
	return _Challenge.Contract.LastMoveBlockTime(&_Challenge.CallOpts)
}

// LastMoveBlockTime is a free data retrieval call binding the contract method 0xed5b1303.
//
// Solidity: function lastMoveBlockTime() view returns(uint256)
func (_Challenge *ChallengeCallerSession) LastMoveBlockTime() (*big.Int, error) {
	return _Challenge.Contract.LastMoveBlockTime(&_Challenge.CallOpts)
}

// PrevBisection is a free data retrieval call binding the contract method 0xafeae965.
//
// Solidity: function prevBisection(uint256 ) view returns(bytes32)
func (_Challenge *ChallengeCaller) PrevBisection(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "prevBisection", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PrevBisection is a free data retrieval call binding the contract method 0xafeae965.
//
// Solidity: function prevBisection(uint256 ) view returns(bytes32)
func (_Challenge *ChallengeSession) PrevBisection(arg0 *big.Int) ([32]byte, error) {
	return _Challenge.Contract.PrevBisection(&_Challenge.CallOpts, arg0)
}

// PrevBisection is a free data retrieval call binding the contract method 0xafeae965.
//
// Solidity: function prevBisection(uint256 ) view returns(bytes32)
func (_Challenge *ChallengeCallerSession) PrevBisection(arg0 *big.Int) ([32]byte, error) {
	return _Challenge.Contract.PrevBisection(&_Challenge.CallOpts, arg0)
}

// Rollback is a free data retrieval call binding the contract method 0x9afd9d78.
//
// Solidity: function rollback() view returns(bool)
func (_Challenge *ChallengeCaller) Rollback(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "rollback")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Rollback is a free data retrieval call binding the contract method 0x9afd9d78.
//
// Solidity: function rollback() view returns(bool)
func (_Challenge *ChallengeSession) Rollback() (bool, error) {
	return _Challenge.Contract.Rollback(&_Challenge.CallOpts)
}

// Rollback is a free data retrieval call binding the contract method 0x9afd9d78.
//
// Solidity: function rollback() view returns(bool)
func (_Challenge *ChallengeCallerSession) Rollback() (bool, error) {
	return _Challenge.Contract.Rollback(&_Challenge.CallOpts)
}

// StartInboxSize is a free data retrieval call binding the contract method 0xfaeff41b.
//
// Solidity: function startInboxSize() view returns(uint256)
func (_Challenge *ChallengeCaller) StartInboxSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "startInboxSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartInboxSize is a free data retrieval call binding the contract method 0xfaeff41b.
//
// Solidity: function startInboxSize() view returns(uint256)
func (_Challenge *ChallengeSession) StartInboxSize() (*big.Int, error) {
	return _Challenge.Contract.StartInboxSize(&_Challenge.CallOpts)
}

// StartInboxSize is a free data retrieval call binding the contract method 0xfaeff41b.
//
// Solidity: function startInboxSize() view returns(uint256)
func (_Challenge *ChallengeCallerSession) StartInboxSize() (*big.Int, error) {
	return _Challenge.Contract.StartInboxSize(&_Challenge.CallOpts)
}

// Turn is a free data retrieval call binding the contract method 0x8b299903.
//
// Solidity: function turn() view returns(uint8)
func (_Challenge *ChallengeCaller) Turn(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "turn")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Turn is a free data retrieval call binding the contract method 0x8b299903.
//
// Solidity: function turn() view returns(uint8)
func (_Challenge *ChallengeSession) Turn() (uint8, error) {
	return _Challenge.Contract.Turn(&_Challenge.CallOpts)
}

// Turn is a free data retrieval call binding the contract method 0x8b299903.
//
// Solidity: function turn() view returns(uint8)
func (_Challenge *ChallengeCallerSession) Turn() (uint8, error) {
	return _Challenge.Contract.Turn(&_Challenge.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_Challenge *ChallengeCaller) Winner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "winner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_Challenge *ChallengeSession) Winner() (common.Address, error) {
	return _Challenge.Contract.Winner(&_Challenge.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_Challenge *ChallengeCallerSession) Winner() (common.Address, error) {
	return _Challenge.Contract.Winner(&_Challenge.CallOpts)
}

// BisectExecution is a paid mutator transaction binding the contract method 0x8f2400a8.
//
// Solidity: function bisectExecution(bytes32[3] bisection, uint256 challengedSegmentIndex, uint256 challengedSegmentStart, uint256 challengedSegmentLength, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_Challenge *ChallengeTransactor) BisectExecution(opts *bind.TransactOpts, bisection [3][32]byte, challengedSegmentIndex *big.Int, challengedSegmentStart *big.Int, challengedSegmentLength *big.Int, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "bisectExecution", bisection, challengedSegmentIndex, challengedSegmentStart, challengedSegmentLength, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// BisectExecution is a paid mutator transaction binding the contract method 0x8f2400a8.
//
// Solidity: function bisectExecution(bytes32[3] bisection, uint256 challengedSegmentIndex, uint256 challengedSegmentStart, uint256 challengedSegmentLength, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_Challenge *ChallengeSession) BisectExecution(bisection [3][32]byte, challengedSegmentIndex *big.Int, challengedSegmentStart *big.Int, challengedSegmentLength *big.Int, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.BisectExecution(&_Challenge.TransactOpts, bisection, challengedSegmentIndex, challengedSegmentStart, challengedSegmentLength, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// BisectExecution is a paid mutator transaction binding the contract method 0x8f2400a8.
//
// Solidity: function bisectExecution(bytes32[3] bisection, uint256 challengedSegmentIndex, uint256 challengedSegmentStart, uint256 challengedSegmentLength, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_Challenge *ChallengeTransactorSession) BisectExecution(bisection [3][32]byte, challengedSegmentIndex *big.Int, challengedSegmentStart *big.Int, challengedSegmentLength *big.Int, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.BisectExecution(&_Challenge.TransactOpts, bisection, challengedSegmentIndex, challengedSegmentStart, challengedSegmentLength, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// Initialize is a paid mutator transaction binding the contract method 0xabf48013.
//
// Solidity: function initialize(address _defender, address _challenger, address _verifier, address _resultReceiver, uint256 _startInboxSize, bytes32 _startStateHash, bytes32 _endStateHash) returns()
func (_Challenge *ChallengeTransactor) Initialize(opts *bind.TransactOpts, _defender common.Address, _challenger common.Address, _verifier common.Address, _resultReceiver common.Address, _startInboxSize *big.Int, _startStateHash [32]byte, _endStateHash [32]byte) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "initialize", _defender, _challenger, _verifier, _resultReceiver, _startInboxSize, _startStateHash, _endStateHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xabf48013.
//
// Solidity: function initialize(address _defender, address _challenger, address _verifier, address _resultReceiver, uint256 _startInboxSize, bytes32 _startStateHash, bytes32 _endStateHash) returns()
func (_Challenge *ChallengeSession) Initialize(_defender common.Address, _challenger common.Address, _verifier common.Address, _resultReceiver common.Address, _startInboxSize *big.Int, _startStateHash [32]byte, _endStateHash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.Initialize(&_Challenge.TransactOpts, _defender, _challenger, _verifier, _resultReceiver, _startInboxSize, _startStateHash, _endStateHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xabf48013.
//
// Solidity: function initialize(address _defender, address _challenger, address _verifier, address _resultReceiver, uint256 _startInboxSize, bytes32 _startStateHash, bytes32 _endStateHash) returns()
func (_Challenge *ChallengeTransactorSession) Initialize(_defender common.Address, _challenger common.Address, _verifier common.Address, _resultReceiver common.Address, _startInboxSize *big.Int, _startStateHash [32]byte, _endStateHash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.Initialize(&_Challenge.TransactOpts, _defender, _challenger, _verifier, _resultReceiver, _startInboxSize, _startStateHash, _endStateHash)
}

// InitializeChallengeLength is a paid mutator transaction binding the contract method 0x18ef160d.
//
// Solidity: function initializeChallengeLength(bytes32 checkStateHash, uint256 _numSteps) returns()
func (_Challenge *ChallengeTransactor) InitializeChallengeLength(opts *bind.TransactOpts, checkStateHash [32]byte, _numSteps *big.Int) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "initializeChallengeLength", checkStateHash, _numSteps)
}

// InitializeChallengeLength is a paid mutator transaction binding the contract method 0x18ef160d.
//
// Solidity: function initializeChallengeLength(bytes32 checkStateHash, uint256 _numSteps) returns()
func (_Challenge *ChallengeSession) InitializeChallengeLength(checkStateHash [32]byte, _numSteps *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.InitializeChallengeLength(&_Challenge.TransactOpts, checkStateHash, _numSteps)
}

// InitializeChallengeLength is a paid mutator transaction binding the contract method 0x18ef160d.
//
// Solidity: function initializeChallengeLength(bytes32 checkStateHash, uint256 _numSteps) returns()
func (_Challenge *ChallengeTransactorSession) InitializeChallengeLength(checkStateHash [32]byte, _numSteps *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.InitializeChallengeLength(&_Challenge.TransactOpts, checkStateHash, _numSteps)
}

// SetRollback is a paid mutator transaction binding the contract method 0x631acced.
//
// Solidity: function setRollback() returns()
func (_Challenge *ChallengeTransactor) SetRollback(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "setRollback")
}

// SetRollback is a paid mutator transaction binding the contract method 0x631acced.
//
// Solidity: function setRollback() returns()
func (_Challenge *ChallengeSession) SetRollback() (*types.Transaction, error) {
	return _Challenge.Contract.SetRollback(&_Challenge.TransactOpts)
}

// SetRollback is a paid mutator transaction binding the contract method 0x631acced.
//
// Solidity: function setRollback() returns()
func (_Challenge *ChallengeTransactorSession) SetRollback() (*types.Transaction, error) {
	return _Challenge.Contract.SetRollback(&_Challenge.TransactOpts)
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_Challenge *ChallengeTransactor) Timeout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "timeout")
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_Challenge *ChallengeSession) Timeout() (*types.Transaction, error) {
	return _Challenge.Contract.Timeout(&_Challenge.TransactOpts)
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_Challenge *ChallengeTransactorSession) Timeout() (*types.Transaction, error) {
	return _Challenge.Contract.Timeout(&_Challenge.TransactOpts)
}

// VerifyOneStepProof is a paid mutator transaction binding the contract method 0x0095c958.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, uint8 verifyType, bytes proof, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_Challenge *ChallengeTransactor) VerifyOneStepProof(opts *bind.TransactOpts, ctx VerificationContextContext, verifyType uint8, proof []byte, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "verifyOneStepProof", ctx, verifyType, proof, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// VerifyOneStepProof is a paid mutator transaction binding the contract method 0x0095c958.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, uint8 verifyType, bytes proof, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_Challenge *ChallengeSession) VerifyOneStepProof(ctx VerificationContextContext, verifyType uint8, proof []byte, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.VerifyOneStepProof(&_Challenge.TransactOpts, ctx, verifyType, proof, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// VerifyOneStepProof is a paid mutator transaction binding the contract method 0x0095c958.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, uint8 verifyType, bytes proof, uint256 prevChallengedSegmentStart, uint256 prevChallengedSegmentLength) returns()
func (_Challenge *ChallengeTransactorSession) VerifyOneStepProof(ctx VerificationContextContext, verifyType uint8, proof []byte, prevChallengedSegmentStart *big.Int, prevChallengedSegmentLength *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.VerifyOneStepProof(&_Challenge.TransactOpts, ctx, verifyType, proof, prevChallengedSegmentStart, prevChallengedSegmentLength)
}

// ChallengeBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the Challenge contract.
type ChallengeBisectedIterator struct {
	Event *ChallengeBisected // Event containing the contract specifics and raw log

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
func (it *ChallengeBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeBisected)
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
		it.Event = new(ChallengeBisected)
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
func (it *ChallengeBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeBisected represents a Bisected event raised by the Challenge contract.
type ChallengeBisected struct {
	StartState              [32]byte
	MidState                [32]byte
	EndState                [32]byte
	BlockNum                *big.Int
	BlockTime               *big.Int
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x71809f4d4f7bf3c208a85ccd3c922c984024f8e3cef51e3d03ae677e4217097d.
//
// Solidity: event Bisected(bytes32 startState, bytes32 midState, bytes32 endState, uint256 blockNum, uint256 blockTime, uint256 challengedSegmentStart, uint256 challengedSegmentLength)
func (_Challenge *ChallengeFilterer) FilterBisected(opts *bind.FilterOpts) (*ChallengeBisectedIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "Bisected")
	if err != nil {
		return nil, err
	}
	return &ChallengeBisectedIterator{contract: _Challenge.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x71809f4d4f7bf3c208a85ccd3c922c984024f8e3cef51e3d03ae677e4217097d.
//
// Solidity: event Bisected(bytes32 startState, bytes32 midState, bytes32 endState, uint256 blockNum, uint256 blockTime, uint256 challengedSegmentStart, uint256 challengedSegmentLength)
func (_Challenge *ChallengeFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *ChallengeBisected) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "Bisected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeBisected)
				if err := _Challenge.contract.UnpackLog(event, "Bisected", log); err != nil {
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

// ParseBisected is a log parse operation binding the contract event 0x71809f4d4f7bf3c208a85ccd3c922c984024f8e3cef51e3d03ae677e4217097d.
//
// Solidity: event Bisected(bytes32 startState, bytes32 midState, bytes32 endState, uint256 blockNum, uint256 blockTime, uint256 challengedSegmentStart, uint256 challengedSegmentLength)
func (_Challenge *ChallengeFilterer) ParseBisected(log types.Log) (*ChallengeBisected, error) {
	event := new(ChallengeBisected)
	if err := _Challenge.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeChallengeCompletedIterator is returned from FilterChallengeCompleted and is used to iterate over the raw logs and unpacked data for ChallengeCompleted events raised by the Challenge contract.
type ChallengeChallengeCompletedIterator struct {
	Event *ChallengeChallengeCompleted // Event containing the contract specifics and raw log

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
func (it *ChallengeChallengeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeChallengeCompleted)
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
		it.Event = new(ChallengeChallengeCompleted)
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
func (it *ChallengeChallengeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeChallengeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeChallengeCompleted represents a ChallengeCompleted event raised by the Challenge contract.
type ChallengeChallengeCompleted struct {
	Winner common.Address
	Loser  common.Address
	Reason uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChallengeCompleted is a free log retrieval operation binding the contract event 0x03f929a9a6b1f0aef5e43cb12b56f862da97ec3de3fda02a52e85f9f3974fb6a.
//
// Solidity: event ChallengeCompleted(address winner, address loser, uint8 reason)
func (_Challenge *ChallengeFilterer) FilterChallengeCompleted(opts *bind.FilterOpts) (*ChallengeChallengeCompletedIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "ChallengeCompleted")
	if err != nil {
		return nil, err
	}
	return &ChallengeChallengeCompletedIterator{contract: _Challenge.contract, event: "ChallengeCompleted", logs: logs, sub: sub}, nil
}

// WatchChallengeCompleted is a free log subscription operation binding the contract event 0x03f929a9a6b1f0aef5e43cb12b56f862da97ec3de3fda02a52e85f9f3974fb6a.
//
// Solidity: event ChallengeCompleted(address winner, address loser, uint8 reason)
func (_Challenge *ChallengeFilterer) WatchChallengeCompleted(opts *bind.WatchOpts, sink chan<- *ChallengeChallengeCompleted) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "ChallengeCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeChallengeCompleted)
				if err := _Challenge.contract.UnpackLog(event, "ChallengeCompleted", log); err != nil {
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
func (_Challenge *ChallengeFilterer) ParseChallengeCompleted(log types.Log) (*ChallengeChallengeCompleted, error) {
	event := new(ChallengeChallengeCompleted)
	if err := _Challenge.contract.UnpackLog(event, "ChallengeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

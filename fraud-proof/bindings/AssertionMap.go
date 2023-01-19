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

// AssertionMapMetaData contains all meta data concerning the AssertionMap contract.
var AssertionMapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assertions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numStakers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"childInboxSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"createAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"deleteAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"getDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"getGasUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"getInboxSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestAssertionID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"getNumStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"getParentID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"getProposalTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"}],\"name\":\"getStateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"isStaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupAddress\",\"type\":\"address\"}],\"name\":\"setRollupAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assertionID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"stakeOnAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AssertionMapABI is the input ABI used to generate the binding from.
// Deprecated: Use AssertionMapMetaData.ABI instead.
var AssertionMapABI = AssertionMapMetaData.ABI

// AssertionMap is an auto generated Go binding around an Ethereum contract.
type AssertionMap struct {
	AssertionMapCaller     // Read-only binding to the contract
	AssertionMapTransactor // Write-only binding to the contract
	AssertionMapFilterer   // Log filterer for contract events
}

// AssertionMapCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssertionMapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionMapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssertionMapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionMapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssertionMapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionMapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssertionMapSession struct {
	Contract     *AssertionMap     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssertionMapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssertionMapCallerSession struct {
	Contract *AssertionMapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AssertionMapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssertionMapTransactorSession struct {
	Contract     *AssertionMapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AssertionMapRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssertionMapRaw struct {
	Contract *AssertionMap // Generic contract binding to access the raw methods on
}

// AssertionMapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssertionMapCallerRaw struct {
	Contract *AssertionMapCaller // Generic read-only contract binding to access the raw methods on
}

// AssertionMapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssertionMapTransactorRaw struct {
	Contract *AssertionMapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssertionMap creates a new instance of AssertionMap, bound to a specific deployed contract.
func NewAssertionMap(address common.Address, backend bind.ContractBackend) (*AssertionMap, error) {
	contract, err := bindAssertionMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssertionMap{AssertionMapCaller: AssertionMapCaller{contract: contract}, AssertionMapTransactor: AssertionMapTransactor{contract: contract}, AssertionMapFilterer: AssertionMapFilterer{contract: contract}}, nil
}

// NewAssertionMapCaller creates a new read-only instance of AssertionMap, bound to a specific deployed contract.
func NewAssertionMapCaller(address common.Address, caller bind.ContractCaller) (*AssertionMapCaller, error) {
	contract, err := bindAssertionMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionMapCaller{contract: contract}, nil
}

// NewAssertionMapTransactor creates a new write-only instance of AssertionMap, bound to a specific deployed contract.
func NewAssertionMapTransactor(address common.Address, transactor bind.ContractTransactor) (*AssertionMapTransactor, error) {
	contract, err := bindAssertionMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionMapTransactor{contract: contract}, nil
}

// NewAssertionMapFilterer creates a new log filterer instance of AssertionMap, bound to a specific deployed contract.
func NewAssertionMapFilterer(address common.Address, filterer bind.ContractFilterer) (*AssertionMapFilterer, error) {
	contract, err := bindAssertionMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssertionMapFilterer{contract: contract}, nil
}

// bindAssertionMap binds a generic wrapper to an already deployed contract.
func bindAssertionMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssertionMapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionMap *AssertionMapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionMap.Contract.AssertionMapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionMap *AssertionMapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionMap.Contract.AssertionMapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionMap *AssertionMapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionMap.Contract.AssertionMapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionMap *AssertionMapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionMap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionMap *AssertionMapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionMap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionMap *AssertionMapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionMap.Contract.contract.Transact(opts, method, params...)
}

// Assertions is a free data retrieval call binding the contract method 0x524232f6.
//
// Solidity: function assertions(uint256 ) view returns(bytes32 stateHash, uint256 inboxSize, uint256 gasUsed, uint256 parent, uint256 deadline, uint256 proposalTime, uint256 numStakers, uint256 childInboxSize)
func (_AssertionMap *AssertionMapCaller) Assertions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StateHash      [32]byte
	InboxSize      *big.Int
	GasUsed        *big.Int
	Parent         *big.Int
	Deadline       *big.Int
	ProposalTime   *big.Int
	NumStakers     *big.Int
	ChildInboxSize *big.Int
}, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "assertions", arg0)

	outstruct := new(struct {
		StateHash      [32]byte
		InboxSize      *big.Int
		GasUsed        *big.Int
		Parent         *big.Int
		Deadline       *big.Int
		ProposalTime   *big.Int
		NumStakers     *big.Int
		ChildInboxSize *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StateHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.InboxSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.GasUsed = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Parent = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ProposalTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.NumStakers = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.ChildInboxSize = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assertions is a free data retrieval call binding the contract method 0x524232f6.
//
// Solidity: function assertions(uint256 ) view returns(bytes32 stateHash, uint256 inboxSize, uint256 gasUsed, uint256 parent, uint256 deadline, uint256 proposalTime, uint256 numStakers, uint256 childInboxSize)
func (_AssertionMap *AssertionMapSession) Assertions(arg0 *big.Int) (struct {
	StateHash      [32]byte
	InboxSize      *big.Int
	GasUsed        *big.Int
	Parent         *big.Int
	Deadline       *big.Int
	ProposalTime   *big.Int
	NumStakers     *big.Int
	ChildInboxSize *big.Int
}, error) {
	return _AssertionMap.Contract.Assertions(&_AssertionMap.CallOpts, arg0)
}

// Assertions is a free data retrieval call binding the contract method 0x524232f6.
//
// Solidity: function assertions(uint256 ) view returns(bytes32 stateHash, uint256 inboxSize, uint256 gasUsed, uint256 parent, uint256 deadline, uint256 proposalTime, uint256 numStakers, uint256 childInboxSize)
func (_AssertionMap *AssertionMapCallerSession) Assertions(arg0 *big.Int) (struct {
	StateHash      [32]byte
	InboxSize      *big.Int
	GasUsed        *big.Int
	Parent         *big.Int
	Deadline       *big.Int
	ProposalTime   *big.Int
	NumStakers     *big.Int
	ChildInboxSize *big.Int
}, error) {
	return _AssertionMap.Contract.Assertions(&_AssertionMap.CallOpts, arg0)
}

// GetDeadline is a free data retrieval call binding the contract method 0x82862275.
//
// Solidity: function getDeadline(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCaller) GetDeadline(opts *bind.CallOpts, assertionID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "getDeadline", assertionID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeadline is a free data retrieval call binding the contract method 0x82862275.
//
// Solidity: function getDeadline(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapSession) GetDeadline(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetDeadline(&_AssertionMap.CallOpts, assertionID)
}

// GetDeadline is a free data retrieval call binding the contract method 0x82862275.
//
// Solidity: function getDeadline(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCallerSession) GetDeadline(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetDeadline(&_AssertionMap.CallOpts, assertionID)
}

// GetGasUsed is a free data retrieval call binding the contract method 0xb919bf93.
//
// Solidity: function getGasUsed(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCaller) GetGasUsed(opts *bind.CallOpts, assertionID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "getGasUsed", assertionID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasUsed is a free data retrieval call binding the contract method 0xb919bf93.
//
// Solidity: function getGasUsed(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapSession) GetGasUsed(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetGasUsed(&_AssertionMap.CallOpts, assertionID)
}

// GetGasUsed is a free data retrieval call binding the contract method 0xb919bf93.
//
// Solidity: function getGasUsed(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCallerSession) GetGasUsed(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetGasUsed(&_AssertionMap.CallOpts, assertionID)
}

// GetInboxSize is a free data retrieval call binding the contract method 0x2b27e93b.
//
// Solidity: function getInboxSize(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCaller) GetInboxSize(opts *bind.CallOpts, assertionID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "getInboxSize", assertionID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInboxSize is a free data retrieval call binding the contract method 0x2b27e93b.
//
// Solidity: function getInboxSize(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapSession) GetInboxSize(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetInboxSize(&_AssertionMap.CallOpts, assertionID)
}

// GetInboxSize is a free data retrieval call binding the contract method 0x2b27e93b.
//
// Solidity: function getInboxSize(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCallerSession) GetInboxSize(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetInboxSize(&_AssertionMap.CallOpts, assertionID)
}

// GetLatestAssertionID is a free data retrieval call binding the contract method 0x4935e6e2.
//
// Solidity: function getLatestAssertionID() view returns(uint256)
func (_AssertionMap *AssertionMapCaller) GetLatestAssertionID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "getLatestAssertionID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestAssertionID is a free data retrieval call binding the contract method 0x4935e6e2.
//
// Solidity: function getLatestAssertionID() view returns(uint256)
func (_AssertionMap *AssertionMapSession) GetLatestAssertionID() (*big.Int, error) {
	return _AssertionMap.Contract.GetLatestAssertionID(&_AssertionMap.CallOpts)
}

// GetLatestAssertionID is a free data retrieval call binding the contract method 0x4935e6e2.
//
// Solidity: function getLatestAssertionID() view returns(uint256)
func (_AssertionMap *AssertionMapCallerSession) GetLatestAssertionID() (*big.Int, error) {
	return _AssertionMap.Contract.GetLatestAssertionID(&_AssertionMap.CallOpts)
}

// GetNumStakers is a free data retrieval call binding the contract method 0x366b2b69.
//
// Solidity: function getNumStakers(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCaller) GetNumStakers(opts *bind.CallOpts, assertionID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "getNumStakers", assertionID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumStakers is a free data retrieval call binding the contract method 0x366b2b69.
//
// Solidity: function getNumStakers(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapSession) GetNumStakers(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetNumStakers(&_AssertionMap.CallOpts, assertionID)
}

// GetNumStakers is a free data retrieval call binding the contract method 0x366b2b69.
//
// Solidity: function getNumStakers(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCallerSession) GetNumStakers(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetNumStakers(&_AssertionMap.CallOpts, assertionID)
}

// GetParentID is a free data retrieval call binding the contract method 0x30b94770.
//
// Solidity: function getParentID(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCaller) GetParentID(opts *bind.CallOpts, assertionID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "getParentID", assertionID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetParentID is a free data retrieval call binding the contract method 0x30b94770.
//
// Solidity: function getParentID(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapSession) GetParentID(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetParentID(&_AssertionMap.CallOpts, assertionID)
}

// GetParentID is a free data retrieval call binding the contract method 0x30b94770.
//
// Solidity: function getParentID(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCallerSession) GetParentID(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetParentID(&_AssertionMap.CallOpts, assertionID)
}

// GetProposalTime is a free data retrieval call binding the contract method 0x4e04886d.
//
// Solidity: function getProposalTime(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCaller) GetProposalTime(opts *bind.CallOpts, assertionID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "getProposalTime", assertionID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalTime is a free data retrieval call binding the contract method 0x4e04886d.
//
// Solidity: function getProposalTime(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapSession) GetProposalTime(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetProposalTime(&_AssertionMap.CallOpts, assertionID)
}

// GetProposalTime is a free data retrieval call binding the contract method 0x4e04886d.
//
// Solidity: function getProposalTime(uint256 assertionID) view returns(uint256)
func (_AssertionMap *AssertionMapCallerSession) GetProposalTime(assertionID *big.Int) (*big.Int, error) {
	return _AssertionMap.Contract.GetProposalTime(&_AssertionMap.CallOpts, assertionID)
}

// GetStateHash is a free data retrieval call binding the contract method 0x54823e66.
//
// Solidity: function getStateHash(uint256 assertionID) view returns(bytes32)
func (_AssertionMap *AssertionMapCaller) GetStateHash(opts *bind.CallOpts, assertionID *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "getStateHash", assertionID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStateHash is a free data retrieval call binding the contract method 0x54823e66.
//
// Solidity: function getStateHash(uint256 assertionID) view returns(bytes32)
func (_AssertionMap *AssertionMapSession) GetStateHash(assertionID *big.Int) ([32]byte, error) {
	return _AssertionMap.Contract.GetStateHash(&_AssertionMap.CallOpts, assertionID)
}

// GetStateHash is a free data retrieval call binding the contract method 0x54823e66.
//
// Solidity: function getStateHash(uint256 assertionID) view returns(bytes32)
func (_AssertionMap *AssertionMapCallerSession) GetStateHash(assertionID *big.Int) ([32]byte, error) {
	return _AssertionMap.Contract.GetStateHash(&_AssertionMap.CallOpts, assertionID)
}

// IsStaker is a free data retrieval call binding the contract method 0x873fd089.
//
// Solidity: function isStaker(uint256 assertionID, address stakerAddress) view returns(bool)
func (_AssertionMap *AssertionMapCaller) IsStaker(opts *bind.CallOpts, assertionID *big.Int, stakerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "isStaker", assertionID, stakerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaker is a free data retrieval call binding the contract method 0x873fd089.
//
// Solidity: function isStaker(uint256 assertionID, address stakerAddress) view returns(bool)
func (_AssertionMap *AssertionMapSession) IsStaker(assertionID *big.Int, stakerAddress common.Address) (bool, error) {
	return _AssertionMap.Contract.IsStaker(&_AssertionMap.CallOpts, assertionID, stakerAddress)
}

// IsStaker is a free data retrieval call binding the contract method 0x873fd089.
//
// Solidity: function isStaker(uint256 assertionID, address stakerAddress) view returns(bool)
func (_AssertionMap *AssertionMapCallerSession) IsStaker(assertionID *big.Int, stakerAddress common.Address) (bool, error) {
	return _AssertionMap.Contract.IsStaker(&_AssertionMap.CallOpts, assertionID, stakerAddress)
}

// RollupAddress is a free data retrieval call binding the contract method 0x5ec6a8df.
//
// Solidity: function rollupAddress() view returns(address)
func (_AssertionMap *AssertionMapCaller) RollupAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "rollupAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupAddress is a free data retrieval call binding the contract method 0x5ec6a8df.
//
// Solidity: function rollupAddress() view returns(address)
func (_AssertionMap *AssertionMapSession) RollupAddress() (common.Address, error) {
	return _AssertionMap.Contract.RollupAddress(&_AssertionMap.CallOpts)
}

// RollupAddress is a free data retrieval call binding the contract method 0x5ec6a8df.
//
// Solidity: function rollupAddress() view returns(address)
func (_AssertionMap *AssertionMapCallerSession) RollupAddress() (common.Address, error) {
	return _AssertionMap.Contract.RollupAddress(&_AssertionMap.CallOpts)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0x05fa5c53.
//
// Solidity: function createAssertion(uint256 assertionID, bytes32 stateHash, uint256 inboxSize, uint256 gasUsed, uint256 parentID, uint256 deadline) returns()
func (_AssertionMap *AssertionMapTransactor) CreateAssertion(opts *bind.TransactOpts, assertionID *big.Int, stateHash [32]byte, inboxSize *big.Int, gasUsed *big.Int, parentID *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AssertionMap.contract.Transact(opts, "createAssertion", assertionID, stateHash, inboxSize, gasUsed, parentID, deadline)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0x05fa5c53.
//
// Solidity: function createAssertion(uint256 assertionID, bytes32 stateHash, uint256 inboxSize, uint256 gasUsed, uint256 parentID, uint256 deadline) returns()
func (_AssertionMap *AssertionMapSession) CreateAssertion(assertionID *big.Int, stateHash [32]byte, inboxSize *big.Int, gasUsed *big.Int, parentID *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AssertionMap.Contract.CreateAssertion(&_AssertionMap.TransactOpts, assertionID, stateHash, inboxSize, gasUsed, parentID, deadline)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0x05fa5c53.
//
// Solidity: function createAssertion(uint256 assertionID, bytes32 stateHash, uint256 inboxSize, uint256 gasUsed, uint256 parentID, uint256 deadline) returns()
func (_AssertionMap *AssertionMapTransactorSession) CreateAssertion(assertionID *big.Int, stateHash [32]byte, inboxSize *big.Int, gasUsed *big.Int, parentID *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AssertionMap.Contract.CreateAssertion(&_AssertionMap.TransactOpts, assertionID, stateHash, inboxSize, gasUsed, parentID, deadline)
}

// DeleteAssertion is a paid mutator transaction binding the contract method 0xd8a4e5af.
//
// Solidity: function deleteAssertion(uint256 assertionID) returns()
func (_AssertionMap *AssertionMapTransactor) DeleteAssertion(opts *bind.TransactOpts, assertionID *big.Int) (*types.Transaction, error) {
	return _AssertionMap.contract.Transact(opts, "deleteAssertion", assertionID)
}

// DeleteAssertion is a paid mutator transaction binding the contract method 0xd8a4e5af.
//
// Solidity: function deleteAssertion(uint256 assertionID) returns()
func (_AssertionMap *AssertionMapSession) DeleteAssertion(assertionID *big.Int) (*types.Transaction, error) {
	return _AssertionMap.Contract.DeleteAssertion(&_AssertionMap.TransactOpts, assertionID)
}

// DeleteAssertion is a paid mutator transaction binding the contract method 0xd8a4e5af.
//
// Solidity: function deleteAssertion(uint256 assertionID) returns()
func (_AssertionMap *AssertionMapTransactorSession) DeleteAssertion(assertionID *big.Int) (*types.Transaction, error) {
	return _AssertionMap.Contract.DeleteAssertion(&_AssertionMap.TransactOpts, assertionID)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AssertionMap *AssertionMapTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionMap.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AssertionMap *AssertionMapSession) Initialize() (*types.Transaction, error) {
	return _AssertionMap.Contract.Initialize(&_AssertionMap.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AssertionMap *AssertionMapTransactorSession) Initialize() (*types.Transaction, error) {
	return _AssertionMap.Contract.Initialize(&_AssertionMap.TransactOpts)
}

// SetRollupAddress is a paid mutator transaction binding the contract method 0x07663706.
//
// Solidity: function setRollupAddress(address _rollupAddress) returns()
func (_AssertionMap *AssertionMapTransactor) SetRollupAddress(opts *bind.TransactOpts, _rollupAddress common.Address) (*types.Transaction, error) {
	return _AssertionMap.contract.Transact(opts, "setRollupAddress", _rollupAddress)
}

// SetRollupAddress is a paid mutator transaction binding the contract method 0x07663706.
//
// Solidity: function setRollupAddress(address _rollupAddress) returns()
func (_AssertionMap *AssertionMapSession) SetRollupAddress(_rollupAddress common.Address) (*types.Transaction, error) {
	return _AssertionMap.Contract.SetRollupAddress(&_AssertionMap.TransactOpts, _rollupAddress)
}

// SetRollupAddress is a paid mutator transaction binding the contract method 0x07663706.
//
// Solidity: function setRollupAddress(address _rollupAddress) returns()
func (_AssertionMap *AssertionMapTransactorSession) SetRollupAddress(_rollupAddress common.Address) (*types.Transaction, error) {
	return _AssertionMap.Contract.SetRollupAddress(&_AssertionMap.TransactOpts, _rollupAddress)
}

// StakeOnAssertion is a paid mutator transaction binding the contract method 0xa832c3ae.
//
// Solidity: function stakeOnAssertion(uint256 assertionID, address stakerAddress) returns()
func (_AssertionMap *AssertionMapTransactor) StakeOnAssertion(opts *bind.TransactOpts, assertionID *big.Int, stakerAddress common.Address) (*types.Transaction, error) {
	return _AssertionMap.contract.Transact(opts, "stakeOnAssertion", assertionID, stakerAddress)
}

// StakeOnAssertion is a paid mutator transaction binding the contract method 0xa832c3ae.
//
// Solidity: function stakeOnAssertion(uint256 assertionID, address stakerAddress) returns()
func (_AssertionMap *AssertionMapSession) StakeOnAssertion(assertionID *big.Int, stakerAddress common.Address) (*types.Transaction, error) {
	return _AssertionMap.Contract.StakeOnAssertion(&_AssertionMap.TransactOpts, assertionID, stakerAddress)
}

// StakeOnAssertion is a paid mutator transaction binding the contract method 0xa832c3ae.
//
// Solidity: function stakeOnAssertion(uint256 assertionID, address stakerAddress) returns()
func (_AssertionMap *AssertionMapTransactorSession) StakeOnAssertion(assertionID *big.Int, stakerAddress common.Address) (*types.Transaction, error) {
	return _AssertionMap.Contract.StakeOnAssertion(&_AssertionMap.TransactOpts, assertionID, stakerAddress)
}

// AssertionMapInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AssertionMap contract.
type AssertionMapInitializedIterator struct {
	Event *AssertionMapInitialized // Event containing the contract specifics and raw log

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
func (it *AssertionMapInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssertionMapInitialized)
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
		it.Event = new(AssertionMapInitialized)
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
func (it *AssertionMapInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssertionMapInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssertionMapInitialized represents a Initialized event raised by the AssertionMap contract.
type AssertionMapInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AssertionMap *AssertionMapFilterer) FilterInitialized(opts *bind.FilterOpts) (*AssertionMapInitializedIterator, error) {

	logs, sub, err := _AssertionMap.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AssertionMapInitializedIterator{contract: _AssertionMap.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AssertionMap *AssertionMapFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AssertionMapInitialized) (event.Subscription, error) {

	logs, sub, err := _AssertionMap.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssertionMapInitialized)
				if err := _AssertionMap.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AssertionMap *AssertionMapFilterer) ParseInitialized(log types.Log) (*AssertionMapInitialized, error) {
	event := new(AssertionMapInitialized)
	if err := _AssertionMap.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

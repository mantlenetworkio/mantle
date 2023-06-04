// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VerifierEntry

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

// VerifierEntryMetaData contains all meta data concerning the VerifierEntry contract.
var VerifierEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"blockFinalizationVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockInitiationVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callOpVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"environmentalOpVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interTxVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invalidOpVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"memoryOpVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"verifier\",\"type\":\"uint8\"},{\"internalType\":\"contractIVerifier\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stackOpVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageOpVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structEVMTypesLib.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"inputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerificationContext.Context\",\"name\":\"ctx\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"verifier\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"currStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"verifyOneStepProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VerifierEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierEntryMetaData.ABI instead.
var VerifierEntryABI = VerifierEntryMetaData.ABI

// VerifierEntry is an auto generated Go binding around an Ethereum contract.
type VerifierEntry struct {
	VerifierEntryCaller     // Read-only binding to the contract
	VerifierEntryTransactor // Write-only binding to the contract
	VerifierEntryFilterer   // Log filterer for contract events
}

// VerifierEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierEntrySession struct {
	Contract     *VerifierEntry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierEntryCallerSession struct {
	Contract *VerifierEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VerifierEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierEntryTransactorSession struct {
	Contract     *VerifierEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VerifierEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierEntryRaw struct {
	Contract *VerifierEntry // Generic contract binding to access the raw methods on
}

// VerifierEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierEntryCallerRaw struct {
	Contract *VerifierEntryCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierEntryTransactorRaw struct {
	Contract *VerifierEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifierEntry creates a new instance of VerifierEntry, bound to a specific deployed contract.
func NewVerifierEntry(address common.Address, backend bind.ContractBackend) (*VerifierEntry, error) {
	contract, err := bindVerifierEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierEntry{VerifierEntryCaller: VerifierEntryCaller{contract: contract}, VerifierEntryTransactor: VerifierEntryTransactor{contract: contract}, VerifierEntryFilterer: VerifierEntryFilterer{contract: contract}}, nil
}

// NewVerifierEntryCaller creates a new read-only instance of VerifierEntry, bound to a specific deployed contract.
func NewVerifierEntryCaller(address common.Address, caller bind.ContractCaller) (*VerifierEntryCaller, error) {
	contract, err := bindVerifierEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierEntryCaller{contract: contract}, nil
}

// NewVerifierEntryTransactor creates a new write-only instance of VerifierEntry, bound to a specific deployed contract.
func NewVerifierEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierEntryTransactor, error) {
	contract, err := bindVerifierEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierEntryTransactor{contract: contract}, nil
}

// NewVerifierEntryFilterer creates a new log filterer instance of VerifierEntry, bound to a specific deployed contract.
func NewVerifierEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierEntryFilterer, error) {
	contract, err := bindVerifierEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierEntryFilterer{contract: contract}, nil
}

// bindVerifierEntry binds a generic wrapper to an already deployed contract.
func bindVerifierEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierEntry *VerifierEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierEntry.Contract.VerifierEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierEntry *VerifierEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierEntry.Contract.VerifierEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierEntry *VerifierEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierEntry.Contract.VerifierEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierEntry *VerifierEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierEntry *VerifierEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierEntry *VerifierEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierEntry.Contract.contract.Transact(opts, method, params...)
}

// BlockFinalizationVerifier is a free data retrieval call binding the contract method 0x9d765195.
//
// Solidity: function blockFinalizationVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCaller) BlockFinalizationVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierEntry.contract.Call(opts, &out, "blockFinalizationVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockFinalizationVerifier is a free data retrieval call binding the contract method 0x9d765195.
//
// Solidity: function blockFinalizationVerifier() view returns(address)
func (_VerifierEntry *VerifierEntrySession) BlockFinalizationVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.BlockFinalizationVerifier(&_VerifierEntry.CallOpts)
}

// BlockFinalizationVerifier is a free data retrieval call binding the contract method 0x9d765195.
//
// Solidity: function blockFinalizationVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCallerSession) BlockFinalizationVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.BlockFinalizationVerifier(&_VerifierEntry.CallOpts)
}

// BlockInitiationVerifier is a free data retrieval call binding the contract method 0xaa8841a2.
//
// Solidity: function blockInitiationVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCaller) BlockInitiationVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierEntry.contract.Call(opts, &out, "blockInitiationVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockInitiationVerifier is a free data retrieval call binding the contract method 0xaa8841a2.
//
// Solidity: function blockInitiationVerifier() view returns(address)
func (_VerifierEntry *VerifierEntrySession) BlockInitiationVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.BlockInitiationVerifier(&_VerifierEntry.CallOpts)
}

// BlockInitiationVerifier is a free data retrieval call binding the contract method 0xaa8841a2.
//
// Solidity: function blockInitiationVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCallerSession) BlockInitiationVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.BlockInitiationVerifier(&_VerifierEntry.CallOpts)
}

// CallOpVerifier is a free data retrieval call binding the contract method 0x199efa46.
//
// Solidity: function callOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCaller) CallOpVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierEntry.contract.Call(opts, &out, "callOpVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CallOpVerifier is a free data retrieval call binding the contract method 0x199efa46.
//
// Solidity: function callOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntrySession) CallOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.CallOpVerifier(&_VerifierEntry.CallOpts)
}

// CallOpVerifier is a free data retrieval call binding the contract method 0x199efa46.
//
// Solidity: function callOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCallerSession) CallOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.CallOpVerifier(&_VerifierEntry.CallOpts)
}

// EnvironmentalOpVerifier is a free data retrieval call binding the contract method 0x4539aaef.
//
// Solidity: function environmentalOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCaller) EnvironmentalOpVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierEntry.contract.Call(opts, &out, "environmentalOpVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnvironmentalOpVerifier is a free data retrieval call binding the contract method 0x4539aaef.
//
// Solidity: function environmentalOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntrySession) EnvironmentalOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.EnvironmentalOpVerifier(&_VerifierEntry.CallOpts)
}

// EnvironmentalOpVerifier is a free data retrieval call binding the contract method 0x4539aaef.
//
// Solidity: function environmentalOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCallerSession) EnvironmentalOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.EnvironmentalOpVerifier(&_VerifierEntry.CallOpts)
}

// InterTxVerifier is a free data retrieval call binding the contract method 0x6ab263d0.
//
// Solidity: function interTxVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCaller) InterTxVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierEntry.contract.Call(opts, &out, "interTxVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterTxVerifier is a free data retrieval call binding the contract method 0x6ab263d0.
//
// Solidity: function interTxVerifier() view returns(address)
func (_VerifierEntry *VerifierEntrySession) InterTxVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.InterTxVerifier(&_VerifierEntry.CallOpts)
}

// InterTxVerifier is a free data retrieval call binding the contract method 0x6ab263d0.
//
// Solidity: function interTxVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCallerSession) InterTxVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.InterTxVerifier(&_VerifierEntry.CallOpts)
}

// InvalidOpVerifier is a free data retrieval call binding the contract method 0x3f1bf192.
//
// Solidity: function invalidOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCaller) InvalidOpVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierEntry.contract.Call(opts, &out, "invalidOpVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InvalidOpVerifier is a free data retrieval call binding the contract method 0x3f1bf192.
//
// Solidity: function invalidOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntrySession) InvalidOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.InvalidOpVerifier(&_VerifierEntry.CallOpts)
}

// InvalidOpVerifier is a free data retrieval call binding the contract method 0x3f1bf192.
//
// Solidity: function invalidOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCallerSession) InvalidOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.InvalidOpVerifier(&_VerifierEntry.CallOpts)
}

// MemoryOpVerifier is a free data retrieval call binding the contract method 0x17facb7d.
//
// Solidity: function memoryOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCaller) MemoryOpVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierEntry.contract.Call(opts, &out, "memoryOpVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MemoryOpVerifier is a free data retrieval call binding the contract method 0x17facb7d.
//
// Solidity: function memoryOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntrySession) MemoryOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.MemoryOpVerifier(&_VerifierEntry.CallOpts)
}

// MemoryOpVerifier is a free data retrieval call binding the contract method 0x17facb7d.
//
// Solidity: function memoryOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCallerSession) MemoryOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.MemoryOpVerifier(&_VerifierEntry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifierEntry *VerifierEntryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierEntry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifierEntry *VerifierEntrySession) Owner() (common.Address, error) {
	return _VerifierEntry.Contract.Owner(&_VerifierEntry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifierEntry *VerifierEntryCallerSession) Owner() (common.Address, error) {
	return _VerifierEntry.Contract.Owner(&_VerifierEntry.CallOpts)
}

// StackOpVerifier is a free data retrieval call binding the contract method 0x53660e70.
//
// Solidity: function stackOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCaller) StackOpVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierEntry.contract.Call(opts, &out, "stackOpVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StackOpVerifier is a free data retrieval call binding the contract method 0x53660e70.
//
// Solidity: function stackOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntrySession) StackOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.StackOpVerifier(&_VerifierEntry.CallOpts)
}

// StackOpVerifier is a free data retrieval call binding the contract method 0x53660e70.
//
// Solidity: function stackOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCallerSession) StackOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.StackOpVerifier(&_VerifierEntry.CallOpts)
}

// StorageOpVerifier is a free data retrieval call binding the contract method 0x54f5b858.
//
// Solidity: function storageOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCaller) StorageOpVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierEntry.contract.Call(opts, &out, "storageOpVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageOpVerifier is a free data retrieval call binding the contract method 0x54f5b858.
//
// Solidity: function storageOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntrySession) StorageOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.StorageOpVerifier(&_VerifierEntry.CallOpts)
}

// StorageOpVerifier is a free data retrieval call binding the contract method 0x54f5b858.
//
// Solidity: function storageOpVerifier() view returns(address)
func (_VerifierEntry *VerifierEntryCallerSession) StorageOpVerifier() (common.Address, error) {
	return _VerifierEntry.Contract.StorageOpVerifier(&_VerifierEntry.CallOpts)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x625eb72e.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, uint8 verifier, bytes32 currStateHash, bytes encoded) view returns(bytes32)
func (_VerifierEntry *VerifierEntryCaller) VerifyOneStepProof(opts *bind.CallOpts, ctx VerificationContextContext, verifier uint8, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	var out []interface{}
	err := _VerifierEntry.contract.Call(opts, &out, "verifyOneStepProof", ctx, verifier, currStateHash, encoded)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x625eb72e.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, uint8 verifier, bytes32 currStateHash, bytes encoded) view returns(bytes32)
func (_VerifierEntry *VerifierEntrySession) VerifyOneStepProof(ctx VerificationContextContext, verifier uint8, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _VerifierEntry.Contract.VerifyOneStepProof(&_VerifierEntry.CallOpts, ctx, verifier, currStateHash, encoded)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x625eb72e.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, uint8 verifier, bytes32 currStateHash, bytes encoded) view returns(bytes32)
func (_VerifierEntry *VerifierEntryCallerSession) VerifyOneStepProof(ctx VerificationContextContext, verifier uint8, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _VerifierEntry.Contract.VerifyOneStepProof(&_VerifierEntry.CallOpts, ctx, verifier, currStateHash, encoded)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VerifierEntry *VerifierEntryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierEntry.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VerifierEntry *VerifierEntrySession) Initialize() (*types.Transaction, error) {
	return _VerifierEntry.Contract.Initialize(&_VerifierEntry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VerifierEntry *VerifierEntryTransactorSession) Initialize() (*types.Transaction, error) {
	return _VerifierEntry.Contract.Initialize(&_VerifierEntry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierEntry *VerifierEntryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierEntry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierEntry *VerifierEntrySession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifierEntry.Contract.RenounceOwnership(&_VerifierEntry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierEntry *VerifierEntryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifierEntry.Contract.RenounceOwnership(&_VerifierEntry.TransactOpts)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x2fb8ff97.
//
// Solidity: function setVerifier(uint8 verifier, address impl) returns()
func (_VerifierEntry *VerifierEntryTransactor) SetVerifier(opts *bind.TransactOpts, verifier uint8, impl common.Address) (*types.Transaction, error) {
	return _VerifierEntry.contract.Transact(opts, "setVerifier", verifier, impl)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x2fb8ff97.
//
// Solidity: function setVerifier(uint8 verifier, address impl) returns()
func (_VerifierEntry *VerifierEntrySession) SetVerifier(verifier uint8, impl common.Address) (*types.Transaction, error) {
	return _VerifierEntry.Contract.SetVerifier(&_VerifierEntry.TransactOpts, verifier, impl)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x2fb8ff97.
//
// Solidity: function setVerifier(uint8 verifier, address impl) returns()
func (_VerifierEntry *VerifierEntryTransactorSession) SetVerifier(verifier uint8, impl common.Address) (*types.Transaction, error) {
	return _VerifierEntry.Contract.SetVerifier(&_VerifierEntry.TransactOpts, verifier, impl)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierEntry *VerifierEntryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VerifierEntry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierEntry *VerifierEntrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifierEntry.Contract.TransferOwnership(&_VerifierEntry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierEntry *VerifierEntryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifierEntry.Contract.TransferOwnership(&_VerifierEntry.TransactOpts, newOwner)
}

// VerifierEntryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VerifierEntry contract.
type VerifierEntryInitializedIterator struct {
	Event *VerifierEntryInitialized // Event containing the contract specifics and raw log

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
func (it *VerifierEntryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierEntryInitialized)
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
		it.Event = new(VerifierEntryInitialized)
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
func (it *VerifierEntryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifierEntryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifierEntryInitialized represents a Initialized event raised by the VerifierEntry contract.
type VerifierEntryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VerifierEntry *VerifierEntryFilterer) FilterInitialized(opts *bind.FilterOpts) (*VerifierEntryInitializedIterator, error) {

	logs, sub, err := _VerifierEntry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VerifierEntryInitializedIterator{contract: _VerifierEntry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VerifierEntry *VerifierEntryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VerifierEntryInitialized) (event.Subscription, error) {

	logs, sub, err := _VerifierEntry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifierEntryInitialized)
				if err := _VerifierEntry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VerifierEntry *VerifierEntryFilterer) ParseInitialized(log types.Log) (*VerifierEntryInitialized, error) {
	event := new(VerifierEntryInitialized)
	if err := _VerifierEntry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifierEntryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VerifierEntry contract.
type VerifierEntryOwnershipTransferredIterator struct {
	Event *VerifierEntryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VerifierEntryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierEntryOwnershipTransferred)
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
		it.Event = new(VerifierEntryOwnershipTransferred)
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
func (it *VerifierEntryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifierEntryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifierEntryOwnershipTransferred represents a OwnershipTransferred event raised by the VerifierEntry contract.
type VerifierEntryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifierEntry *VerifierEntryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifierEntryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifierEntry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifierEntryOwnershipTransferredIterator{contract: _VerifierEntry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifierEntry *VerifierEntryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierEntryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifierEntry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifierEntryOwnershipTransferred)
				if err := _VerifierEntry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VerifierEntry *VerifierEntryFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierEntryOwnershipTransferred, error) {
	event := new(VerifierEntryOwnershipTransferred)
	if err := _VerifierEntry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

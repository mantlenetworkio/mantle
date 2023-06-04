// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IVerifierEntry

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

// IVerifierEntryMetaData contains all meta data concerning the IVerifierEntry contract.
var IVerifierEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structEVMTypesLib.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"inputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerificationContext.Context\",\"name\":\"ctx\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"verifier\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"currStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"verifyOneStepProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IVerifierEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use IVerifierEntryMetaData.ABI instead.
var IVerifierEntryABI = IVerifierEntryMetaData.ABI

// IVerifierEntry is an auto generated Go binding around an Ethereum contract.
type IVerifierEntry struct {
	IVerifierEntryCaller     // Read-only binding to the contract
	IVerifierEntryTransactor // Write-only binding to the contract
	IVerifierEntryFilterer   // Log filterer for contract events
}

// IVerifierEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVerifierEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVerifierEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVerifierEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVerifierEntrySession struct {
	Contract     *IVerifierEntry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVerifierEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVerifierEntryCallerSession struct {
	Contract *IVerifierEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IVerifierEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVerifierEntryTransactorSession struct {
	Contract     *IVerifierEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IVerifierEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVerifierEntryRaw struct {
	Contract *IVerifierEntry // Generic contract binding to access the raw methods on
}

// IVerifierEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVerifierEntryCallerRaw struct {
	Contract *IVerifierEntryCaller // Generic read-only contract binding to access the raw methods on
}

// IVerifierEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVerifierEntryTransactorRaw struct {
	Contract *IVerifierEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVerifierEntry creates a new instance of IVerifierEntry, bound to a specific deployed contract.
func NewIVerifierEntry(address common.Address, backend bind.ContractBackend) (*IVerifierEntry, error) {
	contract, err := bindIVerifierEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVerifierEntry{IVerifierEntryCaller: IVerifierEntryCaller{contract: contract}, IVerifierEntryTransactor: IVerifierEntryTransactor{contract: contract}, IVerifierEntryFilterer: IVerifierEntryFilterer{contract: contract}}, nil
}

// NewIVerifierEntryCaller creates a new read-only instance of IVerifierEntry, bound to a specific deployed contract.
func NewIVerifierEntryCaller(address common.Address, caller bind.ContractCaller) (*IVerifierEntryCaller, error) {
	contract, err := bindIVerifierEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVerifierEntryCaller{contract: contract}, nil
}

// NewIVerifierEntryTransactor creates a new write-only instance of IVerifierEntry, bound to a specific deployed contract.
func NewIVerifierEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*IVerifierEntryTransactor, error) {
	contract, err := bindIVerifierEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVerifierEntryTransactor{contract: contract}, nil
}

// NewIVerifierEntryFilterer creates a new log filterer instance of IVerifierEntry, bound to a specific deployed contract.
func NewIVerifierEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*IVerifierEntryFilterer, error) {
	contract, err := bindIVerifierEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVerifierEntryFilterer{contract: contract}, nil
}

// bindIVerifierEntry binds a generic wrapper to an already deployed contract.
func bindIVerifierEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IVerifierEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVerifierEntry *IVerifierEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVerifierEntry.Contract.IVerifierEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVerifierEntry *IVerifierEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVerifierEntry.Contract.IVerifierEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVerifierEntry *IVerifierEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVerifierEntry.Contract.IVerifierEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVerifierEntry *IVerifierEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVerifierEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVerifierEntry *IVerifierEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVerifierEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVerifierEntry *IVerifierEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVerifierEntry.Contract.contract.Transact(opts, method, params...)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x625eb72e.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, uint8 verifier, bytes32 currStateHash, bytes encoded) view returns(bytes32)
func (_IVerifierEntry *IVerifierEntryCaller) VerifyOneStepProof(opts *bind.CallOpts, ctx VerificationContextContext, verifier uint8, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	var out []interface{}
	err := _IVerifierEntry.contract.Call(opts, &out, "verifyOneStepProof", ctx, verifier, currStateHash, encoded)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x625eb72e.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, uint8 verifier, bytes32 currStateHash, bytes encoded) view returns(bytes32)
func (_IVerifierEntry *IVerifierEntrySession) VerifyOneStepProof(ctx VerificationContextContext, verifier uint8, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _IVerifierEntry.Contract.VerifyOneStepProof(&_IVerifierEntry.CallOpts, ctx, verifier, currStateHash, encoded)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x625eb72e.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, uint8 verifier, bytes32 currStateHash, bytes encoded) view returns(bytes32)
func (_IVerifierEntry *IVerifierEntryCallerSession) VerifyOneStepProof(ctx VerificationContextContext, verifier uint8, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _IVerifierEntry.Contract.VerifyOneStepProof(&_IVerifierEntry.CallOpts, ctx, verifier, currStateHash, encoded)
}

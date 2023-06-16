// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VerifierTestDriver

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

// VerifierTestDriverMetaData contains all meta data concerning the VerifierTestDriver contract.
var VerifierTestDriverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blockInitiationVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_blockFinalizationVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_interTxVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stackOpVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_environmentalOpVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_memoryOpVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_storageOpVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callOpVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_invalidOpVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structEVMTypesLib.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"verifier\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"currStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VerifierTestDriverABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierTestDriverMetaData.ABI instead.
var VerifierTestDriverABI = VerifierTestDriverMetaData.ABI

// VerifierTestDriver is an auto generated Go binding around an Ethereum contract.
type VerifierTestDriver struct {
	VerifierTestDriverCaller     // Read-only binding to the contract
	VerifierTestDriverTransactor // Write-only binding to the contract
	VerifierTestDriverFilterer   // Log filterer for contract events
}

// VerifierTestDriverCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierTestDriverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTestDriverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTestDriverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTestDriverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierTestDriverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTestDriverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierTestDriverSession struct {
	Contract     *VerifierTestDriver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierTestDriverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierTestDriverCallerSession struct {
	Contract *VerifierTestDriverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VerifierTestDriverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTestDriverTransactorSession struct {
	Contract     *VerifierTestDriverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VerifierTestDriverRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierTestDriverRaw struct {
	Contract *VerifierTestDriver // Generic contract binding to access the raw methods on
}

// VerifierTestDriverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierTestDriverCallerRaw struct {
	Contract *VerifierTestDriverCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTestDriverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTestDriverTransactorRaw struct {
	Contract *VerifierTestDriverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifierTestDriver creates a new instance of VerifierTestDriver, bound to a specific deployed contract.
func NewVerifierTestDriver(address common.Address, backend bind.ContractBackend) (*VerifierTestDriver, error) {
	contract, err := bindVerifierTestDriver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierTestDriver{VerifierTestDriverCaller: VerifierTestDriverCaller{contract: contract}, VerifierTestDriverTransactor: VerifierTestDriverTransactor{contract: contract}, VerifierTestDriverFilterer: VerifierTestDriverFilterer{contract: contract}}, nil
}

// NewVerifierTestDriverCaller creates a new read-only instance of VerifierTestDriver, bound to a specific deployed contract.
func NewVerifierTestDriverCaller(address common.Address, caller bind.ContractCaller) (*VerifierTestDriverCaller, error) {
	contract, err := bindVerifierTestDriver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTestDriverCaller{contract: contract}, nil
}

// NewVerifierTestDriverTransactor creates a new write-only instance of VerifierTestDriver, bound to a specific deployed contract.
func NewVerifierTestDriverTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTestDriverTransactor, error) {
	contract, err := bindVerifierTestDriver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTestDriverTransactor{contract: contract}, nil
}

// NewVerifierTestDriverFilterer creates a new log filterer instance of VerifierTestDriver, bound to a specific deployed contract.
func NewVerifierTestDriverFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierTestDriverFilterer, error) {
	contract, err := bindVerifierTestDriver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierTestDriverFilterer{contract: contract}, nil
}

// bindVerifierTestDriver binds a generic wrapper to an already deployed contract.
func bindVerifierTestDriver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierTestDriverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierTestDriver *VerifierTestDriverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierTestDriver.Contract.VerifierTestDriverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierTestDriver *VerifierTestDriverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierTestDriver.Contract.VerifierTestDriverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierTestDriver *VerifierTestDriverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierTestDriver.Contract.VerifierTestDriverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierTestDriver *VerifierTestDriverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierTestDriver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierTestDriver *VerifierTestDriverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierTestDriver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierTestDriver *VerifierTestDriverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierTestDriver.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x86f73e4e.
//
// Solidity: function verifyProof(address sequencerAddress, uint256 timestamp, uint256 number, address origin, bytes32 txHash, (uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256) transaction, uint8 verifier, bytes32 currStateHash, bytes proof) view returns(bytes32)
func (_VerifierTestDriver *VerifierTestDriverCaller) VerifyProof(opts *bind.CallOpts, sequencerAddress common.Address, timestamp *big.Int, number *big.Int, origin common.Address, txHash [32]byte, transaction EVMTypesLibTransaction, verifier uint8, currStateHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _VerifierTestDriver.contract.Call(opts, &out, "verifyProof", sequencerAddress, timestamp, number, origin, txHash, transaction, verifier, currStateHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x86f73e4e.
//
// Solidity: function verifyProof(address sequencerAddress, uint256 timestamp, uint256 number, address origin, bytes32 txHash, (uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256) transaction, uint8 verifier, bytes32 currStateHash, bytes proof) view returns(bytes32)
func (_VerifierTestDriver *VerifierTestDriverSession) VerifyProof(sequencerAddress common.Address, timestamp *big.Int, number *big.Int, origin common.Address, txHash [32]byte, transaction EVMTypesLibTransaction, verifier uint8, currStateHash [32]byte, proof []byte) ([32]byte, error) {
	return _VerifierTestDriver.Contract.VerifyProof(&_VerifierTestDriver.CallOpts, sequencerAddress, timestamp, number, origin, txHash, transaction, verifier, currStateHash, proof)
}

// VerifyProof is a free data retrieval call binding the contract method 0x86f73e4e.
//
// Solidity: function verifyProof(address sequencerAddress, uint256 timestamp, uint256 number, address origin, bytes32 txHash, (uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256) transaction, uint8 verifier, bytes32 currStateHash, bytes proof) view returns(bytes32)
func (_VerifierTestDriver *VerifierTestDriverCallerSession) VerifyProof(sequencerAddress common.Address, timestamp *big.Int, number *big.Int, origin common.Address, txHash [32]byte, transaction EVMTypesLibTransaction, verifier uint8, currStateHash [32]byte, proof []byte) ([32]byte, error) {
	return _VerifierTestDriver.Contract.VerifyProof(&_VerifierTestDriver.CallOpts, sequencerAddress, timestamp, number, origin, txHash, transaction, verifier, currStateHash, proof)
}

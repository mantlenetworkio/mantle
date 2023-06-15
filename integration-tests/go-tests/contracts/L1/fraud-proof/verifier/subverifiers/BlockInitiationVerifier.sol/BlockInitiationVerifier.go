// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BlockInitiationVerifier

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

// BloomLibBloom is an auto generated low-level Go binding around an user-defined struct.
type BloomLibBloom struct {
	Data [8][32]byte
}

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

// OneStepProofInterStateProof is an auto generated low-level Go binding around an user-defined struct.
type OneStepProofInterStateProof struct {
	BlockNumber         uint64
	TransactionIdx      uint64
	GlobalStateRoot     [32]byte
	CumulativeGasUsed   *big.Int
	BlockGasUsed        *big.Int
	BlockHashRoot       [32]byte
	TransactionTrieRoot [32]byte
	ReceiptTrieRoot     [32]byte
	LogsBloom           BloomLibBloom
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

// BlockInitiationVerifierMetaData contains all meta data concerning the BlockInitiationVerifier contract.
var BlockInitiationVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"executeOneStepProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionIdx\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeGasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockGasUsed\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHashRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionTrieRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptTrieRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32[8]\",\"name\":\"data\",\"type\":\"bytes32[8]\"}],\"internalType\":\"structBloomLib.Bloom\",\"name\":\"logsBloom\",\"type\":\"tuple\"}],\"internalType\":\"structOneStepProof.InterStateProof\",\"name\":\"endState\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structEVMTypesLib.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"inputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerificationContext.Context\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"currStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"verifyOneStepProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// BlockInitiationVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockInitiationVerifierMetaData.ABI instead.
var BlockInitiationVerifierABI = BlockInitiationVerifierMetaData.ABI

// BlockInitiationVerifier is an auto generated Go binding around an Ethereum contract.
type BlockInitiationVerifier struct {
	BlockInitiationVerifierCaller     // Read-only binding to the contract
	BlockInitiationVerifierTransactor // Write-only binding to the contract
	BlockInitiationVerifierFilterer   // Log filterer for contract events
}

// BlockInitiationVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockInitiationVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockInitiationVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockInitiationVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockInitiationVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockInitiationVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockInitiationVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockInitiationVerifierSession struct {
	Contract     *BlockInitiationVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BlockInitiationVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockInitiationVerifierCallerSession struct {
	Contract *BlockInitiationVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// BlockInitiationVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockInitiationVerifierTransactorSession struct {
	Contract     *BlockInitiationVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// BlockInitiationVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockInitiationVerifierRaw struct {
	Contract *BlockInitiationVerifier // Generic contract binding to access the raw methods on
}

// BlockInitiationVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockInitiationVerifierCallerRaw struct {
	Contract *BlockInitiationVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// BlockInitiationVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockInitiationVerifierTransactorRaw struct {
	Contract *BlockInitiationVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockInitiationVerifier creates a new instance of BlockInitiationVerifier, bound to a specific deployed contract.
func NewBlockInitiationVerifier(address common.Address, backend bind.ContractBackend) (*BlockInitiationVerifier, error) {
	contract, err := bindBlockInitiationVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockInitiationVerifier{BlockInitiationVerifierCaller: BlockInitiationVerifierCaller{contract: contract}, BlockInitiationVerifierTransactor: BlockInitiationVerifierTransactor{contract: contract}, BlockInitiationVerifierFilterer: BlockInitiationVerifierFilterer{contract: contract}}, nil
}

// NewBlockInitiationVerifierCaller creates a new read-only instance of BlockInitiationVerifier, bound to a specific deployed contract.
func NewBlockInitiationVerifierCaller(address common.Address, caller bind.ContractCaller) (*BlockInitiationVerifierCaller, error) {
	contract, err := bindBlockInitiationVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockInitiationVerifierCaller{contract: contract}, nil
}

// NewBlockInitiationVerifierTransactor creates a new write-only instance of BlockInitiationVerifier, bound to a specific deployed contract.
func NewBlockInitiationVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockInitiationVerifierTransactor, error) {
	contract, err := bindBlockInitiationVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockInitiationVerifierTransactor{contract: contract}, nil
}

// NewBlockInitiationVerifierFilterer creates a new log filterer instance of BlockInitiationVerifier, bound to a specific deployed contract.
func NewBlockInitiationVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockInitiationVerifierFilterer, error) {
	contract, err := bindBlockInitiationVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockInitiationVerifierFilterer{contract: contract}, nil
}

// bindBlockInitiationVerifier binds a generic wrapper to an already deployed contract.
func bindBlockInitiationVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockInitiationVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockInitiationVerifier *BlockInitiationVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockInitiationVerifier.Contract.BlockInitiationVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockInitiationVerifier *BlockInitiationVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockInitiationVerifier.Contract.BlockInitiationVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockInitiationVerifier *BlockInitiationVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockInitiationVerifier.Contract.BlockInitiationVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockInitiationVerifier *BlockInitiationVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockInitiationVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockInitiationVerifier *BlockInitiationVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockInitiationVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockInitiationVerifier *BlockInitiationVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockInitiationVerifier.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0xa1c1f6ab.
//
// Solidity: function executeOneStepProof(bytes32 currStateHash, bytes encoded) pure returns((uint64,uint64,bytes32,uint256,uint256,bytes32,bytes32,bytes32,(bytes32[8])) endState)
func (_BlockInitiationVerifier *BlockInitiationVerifierCaller) ExecuteOneStepProof(opts *bind.CallOpts, currStateHash [32]byte, encoded []byte) (OneStepProofInterStateProof, error) {
	var out []interface{}
	err := _BlockInitiationVerifier.contract.Call(opts, &out, "executeOneStepProof", currStateHash, encoded)

	if err != nil {
		return *new(OneStepProofInterStateProof), err
	}

	out0 := *abi.ConvertType(out[0], new(OneStepProofInterStateProof)).(*OneStepProofInterStateProof)

	return out0, err

}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0xa1c1f6ab.
//
// Solidity: function executeOneStepProof(bytes32 currStateHash, bytes encoded) pure returns((uint64,uint64,bytes32,uint256,uint256,bytes32,bytes32,bytes32,(bytes32[8])) endState)
func (_BlockInitiationVerifier *BlockInitiationVerifierSession) ExecuteOneStepProof(currStateHash [32]byte, encoded []byte) (OneStepProofInterStateProof, error) {
	return _BlockInitiationVerifier.Contract.ExecuteOneStepProof(&_BlockInitiationVerifier.CallOpts, currStateHash, encoded)
}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0xa1c1f6ab.
//
// Solidity: function executeOneStepProof(bytes32 currStateHash, bytes encoded) pure returns((uint64,uint64,bytes32,uint256,uint256,bytes32,bytes32,bytes32,(bytes32[8])) endState)
func (_BlockInitiationVerifier *BlockInitiationVerifierCallerSession) ExecuteOneStepProof(currStateHash [32]byte, encoded []byte) (OneStepProofInterStateProof, error) {
	return _BlockInitiationVerifier.Contract.ExecuteOneStepProof(&_BlockInitiationVerifier.CallOpts, currStateHash, encoded)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) , bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_BlockInitiationVerifier *BlockInitiationVerifierCaller) VerifyOneStepProof(opts *bind.CallOpts, arg0 VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	var out []interface{}
	err := _BlockInitiationVerifier.contract.Call(opts, &out, "verifyOneStepProof", arg0, currStateHash, encoded)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) , bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_BlockInitiationVerifier *BlockInitiationVerifierSession) VerifyOneStepProof(arg0 VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _BlockInitiationVerifier.Contract.VerifyOneStepProof(&_BlockInitiationVerifier.CallOpts, arg0, currStateHash, encoded)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) , bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_BlockInitiationVerifier *BlockInitiationVerifierCallerSession) VerifyOneStepProof(arg0 VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _BlockInitiationVerifier.Contract.VerifyOneStepProof(&_BlockInitiationVerifier.CallOpts, arg0, currStateHash, encoded)
}

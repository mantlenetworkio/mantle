// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BlockFinalizationVerifier

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

// OneStepProofBlockStateProof is an auto generated low-level Go binding around an user-defined struct.
type OneStepProofBlockStateProof struct {
	BlockNumber       uint64
	GlobalStateRoot   [32]byte
	CumulativeGasUsed *big.Int
	BlockHashRoot     [32]byte
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

// BlockFinalizationVerifierMetaData contains all meta data concerning the BlockFinalizationVerifier contract.
var BlockFinalizationVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structEVMTypesLib.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"inputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerificationContext.Context\",\"name\":\"ctx\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"currStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"executeOneStepProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeGasUsed\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHashRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structOneStepProof.BlockStateProof\",\"name\":\"endState\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structEVMTypesLib.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"inputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerificationContext.Context\",\"name\":\"ctx\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"currStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"verifyOneStepProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// BlockFinalizationVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockFinalizationVerifierMetaData.ABI instead.
var BlockFinalizationVerifierABI = BlockFinalizationVerifierMetaData.ABI

// BlockFinalizationVerifier is an auto generated Go binding around an Ethereum contract.
type BlockFinalizationVerifier struct {
	BlockFinalizationVerifierCaller     // Read-only binding to the contract
	BlockFinalizationVerifierTransactor // Write-only binding to the contract
	BlockFinalizationVerifierFilterer   // Log filterer for contract events
}

// BlockFinalizationVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockFinalizationVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockFinalizationVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockFinalizationVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockFinalizationVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockFinalizationVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockFinalizationVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockFinalizationVerifierSession struct {
	Contract     *BlockFinalizationVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BlockFinalizationVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockFinalizationVerifierCallerSession struct {
	Contract *BlockFinalizationVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// BlockFinalizationVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockFinalizationVerifierTransactorSession struct {
	Contract     *BlockFinalizationVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// BlockFinalizationVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockFinalizationVerifierRaw struct {
	Contract *BlockFinalizationVerifier // Generic contract binding to access the raw methods on
}

// BlockFinalizationVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockFinalizationVerifierCallerRaw struct {
	Contract *BlockFinalizationVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// BlockFinalizationVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockFinalizationVerifierTransactorRaw struct {
	Contract *BlockFinalizationVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockFinalizationVerifier creates a new instance of BlockFinalizationVerifier, bound to a specific deployed contract.
func NewBlockFinalizationVerifier(address common.Address, backend bind.ContractBackend) (*BlockFinalizationVerifier, error) {
	contract, err := bindBlockFinalizationVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockFinalizationVerifier{BlockFinalizationVerifierCaller: BlockFinalizationVerifierCaller{contract: contract}, BlockFinalizationVerifierTransactor: BlockFinalizationVerifierTransactor{contract: contract}, BlockFinalizationVerifierFilterer: BlockFinalizationVerifierFilterer{contract: contract}}, nil
}

// NewBlockFinalizationVerifierCaller creates a new read-only instance of BlockFinalizationVerifier, bound to a specific deployed contract.
func NewBlockFinalizationVerifierCaller(address common.Address, caller bind.ContractCaller) (*BlockFinalizationVerifierCaller, error) {
	contract, err := bindBlockFinalizationVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockFinalizationVerifierCaller{contract: contract}, nil
}

// NewBlockFinalizationVerifierTransactor creates a new write-only instance of BlockFinalizationVerifier, bound to a specific deployed contract.
func NewBlockFinalizationVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockFinalizationVerifierTransactor, error) {
	contract, err := bindBlockFinalizationVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockFinalizationVerifierTransactor{contract: contract}, nil
}

// NewBlockFinalizationVerifierFilterer creates a new log filterer instance of BlockFinalizationVerifier, bound to a specific deployed contract.
func NewBlockFinalizationVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockFinalizationVerifierFilterer, error) {
	contract, err := bindBlockFinalizationVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockFinalizationVerifierFilterer{contract: contract}, nil
}

// bindBlockFinalizationVerifier binds a generic wrapper to an already deployed contract.
func bindBlockFinalizationVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockFinalizationVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockFinalizationVerifier *BlockFinalizationVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockFinalizationVerifier.Contract.BlockFinalizationVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockFinalizationVerifier *BlockFinalizationVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockFinalizationVerifier.Contract.BlockFinalizationVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockFinalizationVerifier *BlockFinalizationVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockFinalizationVerifier.Contract.BlockFinalizationVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockFinalizationVerifier *BlockFinalizationVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockFinalizationVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockFinalizationVerifier *BlockFinalizationVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockFinalizationVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockFinalizationVerifier *BlockFinalizationVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockFinalizationVerifier.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0x89c3ad0f.
//
// Solidity: function executeOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns((uint64,bytes32,uint256,bytes32) endState)
func (_BlockFinalizationVerifier *BlockFinalizationVerifierCaller) ExecuteOneStepProof(opts *bind.CallOpts, ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) (OneStepProofBlockStateProof, error) {
	var out []interface{}
	err := _BlockFinalizationVerifier.contract.Call(opts, &out, "executeOneStepProof", ctx, currStateHash, encoded)

	if err != nil {
		return *new(OneStepProofBlockStateProof), err
	}

	out0 := *abi.ConvertType(out[0], new(OneStepProofBlockStateProof)).(*OneStepProofBlockStateProof)

	return out0, err

}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0x89c3ad0f.
//
// Solidity: function executeOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns((uint64,bytes32,uint256,bytes32) endState)
func (_BlockFinalizationVerifier *BlockFinalizationVerifierSession) ExecuteOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) (OneStepProofBlockStateProof, error) {
	return _BlockFinalizationVerifier.Contract.ExecuteOneStepProof(&_BlockFinalizationVerifier.CallOpts, ctx, currStateHash, encoded)
}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0x89c3ad0f.
//
// Solidity: function executeOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns((uint64,bytes32,uint256,bytes32) endState)
func (_BlockFinalizationVerifier *BlockFinalizationVerifierCallerSession) ExecuteOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) (OneStepProofBlockStateProof, error) {
	return _BlockFinalizationVerifier.Contract.ExecuteOneStepProof(&_BlockFinalizationVerifier.CallOpts, ctx, currStateHash, encoded)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_BlockFinalizationVerifier *BlockFinalizationVerifierCaller) VerifyOneStepProof(opts *bind.CallOpts, ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	var out []interface{}
	err := _BlockFinalizationVerifier.contract.Call(opts, &out, "verifyOneStepProof", ctx, currStateHash, encoded)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_BlockFinalizationVerifier *BlockFinalizationVerifierSession) VerifyOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _BlockFinalizationVerifier.Contract.VerifyOneStepProof(&_BlockFinalizationVerifier.CallOpts, ctx, currStateHash, encoded)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_BlockFinalizationVerifier *BlockFinalizationVerifierCallerSession) VerifyOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _BlockFinalizationVerifier.Contract.VerifyOneStepProof(&_BlockFinalizationVerifier.CallOpts, ctx, currStateHash, encoded)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package InvalidOpVerifier

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

// OneStepProofStateProof is an auto generated low-level Go binding around an user-defined struct.
type OneStepProofStateProof struct {
	BlockNumber              uint64
	TransactionIdx           uint64
	Depth                    uint16
	Gas                      uint64
	Refund                   uint64
	LastDepthHash            [32]byte
	ContractAddress          common.Address
	Caller                   common.Address
	Value                    *big.Int
	CallFlag                 uint8
	Out                      uint64
	OutSize                  uint64
	Pc                       uint64
	OpCode                   uint8
	CodeHash                 [32]byte
	StackSize                uint64
	StackHash                [32]byte
	MemSize                  uint64
	MemRoot                  [32]byte
	InputDataSize            uint64
	InputDataRoot            [32]byte
	ReturnDataSize           uint64
	ReturnDataRoot           [32]byte
	CommittedGlobalStateRoot [32]byte
	GlobalStateRoot          [32]byte
	SelfDestructAcc          [32]byte
	LogAcc                   [32]byte
	BlockHashRoot            [32]byte
	AccessListRoot           [32]byte
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

// InvalidOpVerifierMetaData contains all meta data concerning the InvalidOpVerifier contract.
var InvalidOpVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structEVMTypesLib.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"inputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerificationContext.Context\",\"name\":\"ctx\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"currStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"executeOneStepProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionIdx\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"depth\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"refund\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastDepthHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"callFlag\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"out\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pc\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"opCode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"stackSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"stackHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"memSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"memRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"inputDataSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"inputDataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"returnDataSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"returnDataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"committedGlobalStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"selfDestructAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"logAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHashRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accessListRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structOneStepProof.StateProof\",\"name\":\"endState\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structEVMTypesLib.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"inputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerificationContext.Context\",\"name\":\"ctx\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"currStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"verifyOneStepProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// InvalidOpVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use InvalidOpVerifierMetaData.ABI instead.
var InvalidOpVerifierABI = InvalidOpVerifierMetaData.ABI

// InvalidOpVerifier is an auto generated Go binding around an Ethereum contract.
type InvalidOpVerifier struct {
	InvalidOpVerifierCaller     // Read-only binding to the contract
	InvalidOpVerifierTransactor // Write-only binding to the contract
	InvalidOpVerifierFilterer   // Log filterer for contract events
}

// InvalidOpVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type InvalidOpVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvalidOpVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InvalidOpVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvalidOpVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InvalidOpVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvalidOpVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InvalidOpVerifierSession struct {
	Contract     *InvalidOpVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InvalidOpVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InvalidOpVerifierCallerSession struct {
	Contract *InvalidOpVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// InvalidOpVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InvalidOpVerifierTransactorSession struct {
	Contract     *InvalidOpVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// InvalidOpVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type InvalidOpVerifierRaw struct {
	Contract *InvalidOpVerifier // Generic contract binding to access the raw methods on
}

// InvalidOpVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InvalidOpVerifierCallerRaw struct {
	Contract *InvalidOpVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// InvalidOpVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InvalidOpVerifierTransactorRaw struct {
	Contract *InvalidOpVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInvalidOpVerifier creates a new instance of InvalidOpVerifier, bound to a specific deployed contract.
func NewInvalidOpVerifier(address common.Address, backend bind.ContractBackend) (*InvalidOpVerifier, error) {
	contract, err := bindInvalidOpVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InvalidOpVerifier{InvalidOpVerifierCaller: InvalidOpVerifierCaller{contract: contract}, InvalidOpVerifierTransactor: InvalidOpVerifierTransactor{contract: contract}, InvalidOpVerifierFilterer: InvalidOpVerifierFilterer{contract: contract}}, nil
}

// NewInvalidOpVerifierCaller creates a new read-only instance of InvalidOpVerifier, bound to a specific deployed contract.
func NewInvalidOpVerifierCaller(address common.Address, caller bind.ContractCaller) (*InvalidOpVerifierCaller, error) {
	contract, err := bindInvalidOpVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InvalidOpVerifierCaller{contract: contract}, nil
}

// NewInvalidOpVerifierTransactor creates a new write-only instance of InvalidOpVerifier, bound to a specific deployed contract.
func NewInvalidOpVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*InvalidOpVerifierTransactor, error) {
	contract, err := bindInvalidOpVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InvalidOpVerifierTransactor{contract: contract}, nil
}

// NewInvalidOpVerifierFilterer creates a new log filterer instance of InvalidOpVerifier, bound to a specific deployed contract.
func NewInvalidOpVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*InvalidOpVerifierFilterer, error) {
	contract, err := bindInvalidOpVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InvalidOpVerifierFilterer{contract: contract}, nil
}

// bindInvalidOpVerifier binds a generic wrapper to an already deployed contract.
func bindInvalidOpVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InvalidOpVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InvalidOpVerifier *InvalidOpVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InvalidOpVerifier.Contract.InvalidOpVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InvalidOpVerifier *InvalidOpVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InvalidOpVerifier.Contract.InvalidOpVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InvalidOpVerifier *InvalidOpVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InvalidOpVerifier.Contract.InvalidOpVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InvalidOpVerifier *InvalidOpVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InvalidOpVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InvalidOpVerifier *InvalidOpVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InvalidOpVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InvalidOpVerifier *InvalidOpVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InvalidOpVerifier.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0x89c3ad0f.
//
// Solidity: function executeOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns((uint64,uint64,uint16,uint64,uint64,bytes32,address,address,uint256,uint8,uint64,uint64,uint64,uint8,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32) endState)
func (_InvalidOpVerifier *InvalidOpVerifierCaller) ExecuteOneStepProof(opts *bind.CallOpts, ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) (OneStepProofStateProof, error) {
	var out []interface{}
	err := _InvalidOpVerifier.contract.Call(opts, &out, "executeOneStepProof", ctx, currStateHash, encoded)

	if err != nil {
		return *new(OneStepProofStateProof), err
	}

	out0 := *abi.ConvertType(out[0], new(OneStepProofStateProof)).(*OneStepProofStateProof)

	return out0, err

}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0x89c3ad0f.
//
// Solidity: function executeOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns((uint64,uint64,uint16,uint64,uint64,bytes32,address,address,uint256,uint8,uint64,uint64,uint64,uint8,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32) endState)
func (_InvalidOpVerifier *InvalidOpVerifierSession) ExecuteOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) (OneStepProofStateProof, error) {
	return _InvalidOpVerifier.Contract.ExecuteOneStepProof(&_InvalidOpVerifier.CallOpts, ctx, currStateHash, encoded)
}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0x89c3ad0f.
//
// Solidity: function executeOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns((uint64,uint64,uint16,uint64,uint64,bytes32,address,address,uint256,uint8,uint64,uint64,uint64,uint8,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32) endState)
func (_InvalidOpVerifier *InvalidOpVerifierCallerSession) ExecuteOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) (OneStepProofStateProof, error) {
	return _InvalidOpVerifier.Contract.ExecuteOneStepProof(&_InvalidOpVerifier.CallOpts, ctx, currStateHash, encoded)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_InvalidOpVerifier *InvalidOpVerifierCaller) VerifyOneStepProof(opts *bind.CallOpts, ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	var out []interface{}
	err := _InvalidOpVerifier.contract.Call(opts, &out, "verifyOneStepProof", ctx, currStateHash, encoded)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_InvalidOpVerifier *InvalidOpVerifierSession) VerifyOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _InvalidOpVerifier.Contract.VerifyOneStepProof(&_InvalidOpVerifier.CallOpts, ctx, currStateHash, encoded)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_InvalidOpVerifier *InvalidOpVerifierCallerSession) VerifyOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _InvalidOpVerifier.Contract.VerifyOneStepProof(&_InvalidOpVerifier.CallOpts, ctx, currStateHash, encoded)
}

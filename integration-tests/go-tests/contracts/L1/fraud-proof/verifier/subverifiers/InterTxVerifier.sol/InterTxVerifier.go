// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package InterTxVerifier

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

// InterTxVerifierMetaData contains all meta data concerning the InterTxVerifier contract.
var InterTxVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structEVMTypesLib.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"inputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerificationContext.Context\",\"name\":\"ctx\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"currStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"executeOneStepProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionIdx\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"depth\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"refund\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastDepthHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"callFlag\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"out\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pc\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"opCode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"stackSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"stackHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"memSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"memRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"inputDataSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"inputDataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"returnDataSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"returnDataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"committedGlobalStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"selfDestructAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"logAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHashRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accessListRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structOneStepProof.StateProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structEVMTypesLib.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"inputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerificationContext.Context\",\"name\":\"ctx\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"currStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"verifyOneStepProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// InterTxVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use InterTxVerifierMetaData.ABI instead.
var InterTxVerifierABI = InterTxVerifierMetaData.ABI

// InterTxVerifier is an auto generated Go binding around an Ethereum contract.
type InterTxVerifier struct {
	InterTxVerifierCaller     // Read-only binding to the contract
	InterTxVerifierTransactor // Write-only binding to the contract
	InterTxVerifierFilterer   // Log filterer for contract events
}

// InterTxVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterTxVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterTxVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterTxVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterTxVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterTxVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterTxVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterTxVerifierSession struct {
	Contract     *InterTxVerifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InterTxVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterTxVerifierCallerSession struct {
	Contract *InterTxVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// InterTxVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterTxVerifierTransactorSession struct {
	Contract     *InterTxVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// InterTxVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterTxVerifierRaw struct {
	Contract *InterTxVerifier // Generic contract binding to access the raw methods on
}

// InterTxVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterTxVerifierCallerRaw struct {
	Contract *InterTxVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// InterTxVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterTxVerifierTransactorRaw struct {
	Contract *InterTxVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterTxVerifier creates a new instance of InterTxVerifier, bound to a specific deployed contract.
func NewInterTxVerifier(address common.Address, backend bind.ContractBackend) (*InterTxVerifier, error) {
	contract, err := bindInterTxVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterTxVerifier{InterTxVerifierCaller: InterTxVerifierCaller{contract: contract}, InterTxVerifierTransactor: InterTxVerifierTransactor{contract: contract}, InterTxVerifierFilterer: InterTxVerifierFilterer{contract: contract}}, nil
}

// NewInterTxVerifierCaller creates a new read-only instance of InterTxVerifier, bound to a specific deployed contract.
func NewInterTxVerifierCaller(address common.Address, caller bind.ContractCaller) (*InterTxVerifierCaller, error) {
	contract, err := bindInterTxVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterTxVerifierCaller{contract: contract}, nil
}

// NewInterTxVerifierTransactor creates a new write-only instance of InterTxVerifier, bound to a specific deployed contract.
func NewInterTxVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*InterTxVerifierTransactor, error) {
	contract, err := bindInterTxVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterTxVerifierTransactor{contract: contract}, nil
}

// NewInterTxVerifierFilterer creates a new log filterer instance of InterTxVerifier, bound to a specific deployed contract.
func NewInterTxVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*InterTxVerifierFilterer, error) {
	contract, err := bindInterTxVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterTxVerifierFilterer{contract: contract}, nil
}

// bindInterTxVerifier binds a generic wrapper to an already deployed contract.
func bindInterTxVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterTxVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterTxVerifier *InterTxVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterTxVerifier.Contract.InterTxVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterTxVerifier *InterTxVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterTxVerifier.Contract.InterTxVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterTxVerifier *InterTxVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterTxVerifier.Contract.InterTxVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterTxVerifier *InterTxVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterTxVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterTxVerifier *InterTxVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterTxVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterTxVerifier *InterTxVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterTxVerifier.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0x89c3ad0f.
//
// Solidity: function executeOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns((uint64,uint64,uint16,uint64,uint64,bytes32,address,address,uint256,uint8,uint64,uint64,uint64,uint8,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32))
func (_InterTxVerifier *InterTxVerifierCaller) ExecuteOneStepProof(opts *bind.CallOpts, ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) (OneStepProofStateProof, error) {
	var out []interface{}
	err := _InterTxVerifier.contract.Call(opts, &out, "executeOneStepProof", ctx, currStateHash, encoded)

	if err != nil {
		return *new(OneStepProofStateProof), err
	}

	out0 := *abi.ConvertType(out[0], new(OneStepProofStateProof)).(*OneStepProofStateProof)

	return out0, err

}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0x89c3ad0f.
//
// Solidity: function executeOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns((uint64,uint64,uint16,uint64,uint64,bytes32,address,address,uint256,uint8,uint64,uint64,uint64,uint8,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32))
func (_InterTxVerifier *InterTxVerifierSession) ExecuteOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) (OneStepProofStateProof, error) {
	return _InterTxVerifier.Contract.ExecuteOneStepProof(&_InterTxVerifier.CallOpts, ctx, currStateHash, encoded)
}

// ExecuteOneStepProof is a free data retrieval call binding the contract method 0x89c3ad0f.
//
// Solidity: function executeOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns((uint64,uint64,uint16,uint64,uint64,bytes32,address,address,uint256,uint8,uint64,uint64,uint64,uint8,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,uint64,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32))
func (_InterTxVerifier *InterTxVerifierCallerSession) ExecuteOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) (OneStepProofStateProof, error) {
	return _InterTxVerifier.Contract.ExecuteOneStepProof(&_InterTxVerifier.CallOpts, ctx, currStateHash, encoded)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_InterTxVerifier *InterTxVerifierCaller) VerifyOneStepProof(opts *bind.CallOpts, ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	var out []interface{}
	err := _InterTxVerifier.contract.Call(opts, &out, "verifyOneStepProof", ctx, currStateHash, encoded)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_InterTxVerifier *InterTxVerifierSession) VerifyOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _InterTxVerifier.Contract.VerifyOneStepProof(&_InterTxVerifier.CallOpts, ctx, currStateHash, encoded)
}

// VerifyOneStepProof is a free data retrieval call binding the contract method 0x2138b3e4.
//
// Solidity: function verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32) ctx, bytes32 currStateHash, bytes encoded) pure returns(bytes32)
func (_InterTxVerifier *InterTxVerifierCallerSession) VerifyOneStepProof(ctx VerificationContextContext, currStateHash [32]byte, encoded []byte) ([32]byte, error) {
	return _InterTxVerifier.Contract.VerifyOneStepProof(&_InterTxVerifier.CallOpts, ctx, currStateHash, encoded)
}

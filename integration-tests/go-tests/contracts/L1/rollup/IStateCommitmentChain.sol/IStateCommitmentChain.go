// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IStateCommitmentChain

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

// LibBVMCodecChainBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type LibBVMCodecChainBatchHeader struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	Signature         []byte
	ExtraData         []byte
}

// LibBVMCodecChainInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type LibBVMCodecChainInclusionProof struct {
	Index    *big.Int
	Siblings [][32]byte
}

// IStateCommitmentChainMetaData contains all meta data concerning the IStateCommitmentChain contract.
var IStateCommitmentChainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_tssMembers\",\"type\":\"address[]\"}],\"name\":\"DistributeTssReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_batchSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prevTotalElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"StateBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_batch\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_shouldStartAtElement\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"appendStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_BVMCodec.ChainBatchHeader\",\"name\":\"_batchHeader\",\"type\":\"tuple\"}],\"name\":\"deleteStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastSequencerTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastSequencerTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBatches\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalElements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_BVMCodec.ChainBatchHeader\",\"name\":\"_batchHeader\",\"type\":\"tuple\"}],\"name\":\"insideFraudProofWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_inside\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_element\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_BVMCodec.ChainBatchHeader\",\"name\":\"_batchHeader\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLib_BVMCodec.ChainInclusionProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyStateCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_verified\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IStateCommitmentChainABI is the input ABI used to generate the binding from.
// Deprecated: Use IStateCommitmentChainMetaData.ABI instead.
var IStateCommitmentChainABI = IStateCommitmentChainMetaData.ABI

// IStateCommitmentChain is an auto generated Go binding around an Ethereum contract.
type IStateCommitmentChain struct {
	IStateCommitmentChainCaller     // Read-only binding to the contract
	IStateCommitmentChainTransactor // Write-only binding to the contract
	IStateCommitmentChainFilterer   // Log filterer for contract events
}

// IStateCommitmentChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStateCommitmentChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateCommitmentChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStateCommitmentChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateCommitmentChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStateCommitmentChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateCommitmentChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStateCommitmentChainSession struct {
	Contract     *IStateCommitmentChain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IStateCommitmentChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStateCommitmentChainCallerSession struct {
	Contract *IStateCommitmentChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IStateCommitmentChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStateCommitmentChainTransactorSession struct {
	Contract     *IStateCommitmentChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IStateCommitmentChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStateCommitmentChainRaw struct {
	Contract *IStateCommitmentChain // Generic contract binding to access the raw methods on
}

// IStateCommitmentChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStateCommitmentChainCallerRaw struct {
	Contract *IStateCommitmentChainCaller // Generic read-only contract binding to access the raw methods on
}

// IStateCommitmentChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStateCommitmentChainTransactorRaw struct {
	Contract *IStateCommitmentChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStateCommitmentChain creates a new instance of IStateCommitmentChain, bound to a specific deployed contract.
func NewIStateCommitmentChain(address common.Address, backend bind.ContractBackend) (*IStateCommitmentChain, error) {
	contract, err := bindIStateCommitmentChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStateCommitmentChain{IStateCommitmentChainCaller: IStateCommitmentChainCaller{contract: contract}, IStateCommitmentChainTransactor: IStateCommitmentChainTransactor{contract: contract}, IStateCommitmentChainFilterer: IStateCommitmentChainFilterer{contract: contract}}, nil
}

// NewIStateCommitmentChainCaller creates a new read-only instance of IStateCommitmentChain, bound to a specific deployed contract.
func NewIStateCommitmentChainCaller(address common.Address, caller bind.ContractCaller) (*IStateCommitmentChainCaller, error) {
	contract, err := bindIStateCommitmentChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStateCommitmentChainCaller{contract: contract}, nil
}

// NewIStateCommitmentChainTransactor creates a new write-only instance of IStateCommitmentChain, bound to a specific deployed contract.
func NewIStateCommitmentChainTransactor(address common.Address, transactor bind.ContractTransactor) (*IStateCommitmentChainTransactor, error) {
	contract, err := bindIStateCommitmentChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStateCommitmentChainTransactor{contract: contract}, nil
}

// NewIStateCommitmentChainFilterer creates a new log filterer instance of IStateCommitmentChain, bound to a specific deployed contract.
func NewIStateCommitmentChainFilterer(address common.Address, filterer bind.ContractFilterer) (*IStateCommitmentChainFilterer, error) {
	contract, err := bindIStateCommitmentChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStateCommitmentChainFilterer{contract: contract}, nil
}

// bindIStateCommitmentChain binds a generic wrapper to an already deployed contract.
func bindIStateCommitmentChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStateCommitmentChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStateCommitmentChain *IStateCommitmentChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStateCommitmentChain.Contract.IStateCommitmentChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStateCommitmentChain *IStateCommitmentChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStateCommitmentChain.Contract.IStateCommitmentChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStateCommitmentChain *IStateCommitmentChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStateCommitmentChain.Contract.IStateCommitmentChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStateCommitmentChain *IStateCommitmentChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStateCommitmentChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStateCommitmentChain *IStateCommitmentChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStateCommitmentChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStateCommitmentChain *IStateCommitmentChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStateCommitmentChain.Contract.contract.Transact(opts, method, params...)
}

// GetLastSequencerTimestamp is a free data retrieval call binding the contract method 0x7ad168a0.
//
// Solidity: function getLastSequencerTimestamp() view returns(uint256 _lastSequencerTimestamp)
func (_IStateCommitmentChain *IStateCommitmentChainCaller) GetLastSequencerTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStateCommitmentChain.contract.Call(opts, &out, "getLastSequencerTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastSequencerTimestamp is a free data retrieval call binding the contract method 0x7ad168a0.
//
// Solidity: function getLastSequencerTimestamp() view returns(uint256 _lastSequencerTimestamp)
func (_IStateCommitmentChain *IStateCommitmentChainSession) GetLastSequencerTimestamp() (*big.Int, error) {
	return _IStateCommitmentChain.Contract.GetLastSequencerTimestamp(&_IStateCommitmentChain.CallOpts)
}

// GetLastSequencerTimestamp is a free data retrieval call binding the contract method 0x7ad168a0.
//
// Solidity: function getLastSequencerTimestamp() view returns(uint256 _lastSequencerTimestamp)
func (_IStateCommitmentChain *IStateCommitmentChainCallerSession) GetLastSequencerTimestamp() (*big.Int, error) {
	return _IStateCommitmentChain.Contract.GetLastSequencerTimestamp(&_IStateCommitmentChain.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_IStateCommitmentChain *IStateCommitmentChainCaller) GetTotalBatches(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStateCommitmentChain.contract.Call(opts, &out, "getTotalBatches")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_IStateCommitmentChain *IStateCommitmentChainSession) GetTotalBatches() (*big.Int, error) {
	return _IStateCommitmentChain.Contract.GetTotalBatches(&_IStateCommitmentChain.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_IStateCommitmentChain *IStateCommitmentChainCallerSession) GetTotalBatches() (*big.Int, error) {
	return _IStateCommitmentChain.Contract.GetTotalBatches(&_IStateCommitmentChain.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_IStateCommitmentChain *IStateCommitmentChainCaller) GetTotalElements(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStateCommitmentChain.contract.Call(opts, &out, "getTotalElements")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_IStateCommitmentChain *IStateCommitmentChainSession) GetTotalElements() (*big.Int, error) {
	return _IStateCommitmentChain.Contract.GetTotalElements(&_IStateCommitmentChain.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_IStateCommitmentChain *IStateCommitmentChainCallerSession) GetTotalElements() (*big.Int, error) {
	return _IStateCommitmentChain.Contract.GetTotalElements(&_IStateCommitmentChain.CallOpts)
}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0x89a1d980.
//
// Solidity: function insideFraudProofWindow((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) view returns(bool _inside)
func (_IStateCommitmentChain *IStateCommitmentChainCaller) InsideFraudProofWindow(opts *bind.CallOpts, _batchHeader LibBVMCodecChainBatchHeader) (bool, error) {
	var out []interface{}
	err := _IStateCommitmentChain.contract.Call(opts, &out, "insideFraudProofWindow", _batchHeader)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0x89a1d980.
//
// Solidity: function insideFraudProofWindow((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) view returns(bool _inside)
func (_IStateCommitmentChain *IStateCommitmentChainSession) InsideFraudProofWindow(_batchHeader LibBVMCodecChainBatchHeader) (bool, error) {
	return _IStateCommitmentChain.Contract.InsideFraudProofWindow(&_IStateCommitmentChain.CallOpts, _batchHeader)
}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0x89a1d980.
//
// Solidity: function insideFraudProofWindow((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) view returns(bool _inside)
func (_IStateCommitmentChain *IStateCommitmentChainCallerSession) InsideFraudProofWindow(_batchHeader LibBVMCodecChainBatchHeader) (bool, error) {
	return _IStateCommitmentChain.Contract.InsideFraudProofWindow(&_IStateCommitmentChain.CallOpts, _batchHeader)
}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0xb768bb17.
//
// Solidity: function verifyStateCommitment(bytes32 _element, (uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader, (uint256,bytes32[]) _proof) view returns(bool _verified)
func (_IStateCommitmentChain *IStateCommitmentChainCaller) VerifyStateCommitment(opts *bind.CallOpts, _element [32]byte, _batchHeader LibBVMCodecChainBatchHeader, _proof LibBVMCodecChainInclusionProof) (bool, error) {
	var out []interface{}
	err := _IStateCommitmentChain.contract.Call(opts, &out, "verifyStateCommitment", _element, _batchHeader, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0xb768bb17.
//
// Solidity: function verifyStateCommitment(bytes32 _element, (uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader, (uint256,bytes32[]) _proof) view returns(bool _verified)
func (_IStateCommitmentChain *IStateCommitmentChainSession) VerifyStateCommitment(_element [32]byte, _batchHeader LibBVMCodecChainBatchHeader, _proof LibBVMCodecChainInclusionProof) (bool, error) {
	return _IStateCommitmentChain.Contract.VerifyStateCommitment(&_IStateCommitmentChain.CallOpts, _element, _batchHeader, _proof)
}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0xb768bb17.
//
// Solidity: function verifyStateCommitment(bytes32 _element, (uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader, (uint256,bytes32[]) _proof) view returns(bool _verified)
func (_IStateCommitmentChain *IStateCommitmentChainCallerSession) VerifyStateCommitment(_element [32]byte, _batchHeader LibBVMCodecChainBatchHeader, _proof LibBVMCodecChainInclusionProof) (bool, error) {
	return _IStateCommitmentChain.Contract.VerifyStateCommitment(&_IStateCommitmentChain.CallOpts, _element, _batchHeader, _proof)
}

// AppendStateBatch is a paid mutator transaction binding the contract method 0x2169f79f.
//
// Solidity: function appendStateBatch(bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_IStateCommitmentChain *IStateCommitmentChainTransactor) AppendStateBatch(opts *bind.TransactOpts, _batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _IStateCommitmentChain.contract.Transact(opts, "appendStateBatch", _batch, _shouldStartAtElement, _signature)
}

// AppendStateBatch is a paid mutator transaction binding the contract method 0x2169f79f.
//
// Solidity: function appendStateBatch(bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_IStateCommitmentChain *IStateCommitmentChainSession) AppendStateBatch(_batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _IStateCommitmentChain.Contract.AppendStateBatch(&_IStateCommitmentChain.TransactOpts, _batch, _shouldStartAtElement, _signature)
}

// AppendStateBatch is a paid mutator transaction binding the contract method 0x2169f79f.
//
// Solidity: function appendStateBatch(bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_IStateCommitmentChain *IStateCommitmentChainTransactorSession) AppendStateBatch(_batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _IStateCommitmentChain.Contract.AppendStateBatch(&_IStateCommitmentChain.TransactOpts, _batch, _shouldStartAtElement, _signature)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0x5b4d90e2.
//
// Solidity: function deleteStateBatch((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) returns()
func (_IStateCommitmentChain *IStateCommitmentChainTransactor) DeleteStateBatch(opts *bind.TransactOpts, _batchHeader LibBVMCodecChainBatchHeader) (*types.Transaction, error) {
	return _IStateCommitmentChain.contract.Transact(opts, "deleteStateBatch", _batchHeader)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0x5b4d90e2.
//
// Solidity: function deleteStateBatch((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) returns()
func (_IStateCommitmentChain *IStateCommitmentChainSession) DeleteStateBatch(_batchHeader LibBVMCodecChainBatchHeader) (*types.Transaction, error) {
	return _IStateCommitmentChain.Contract.DeleteStateBatch(&_IStateCommitmentChain.TransactOpts, _batchHeader)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0x5b4d90e2.
//
// Solidity: function deleteStateBatch((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) returns()
func (_IStateCommitmentChain *IStateCommitmentChainTransactorSession) DeleteStateBatch(_batchHeader LibBVMCodecChainBatchHeader) (*types.Transaction, error) {
	return _IStateCommitmentChain.Contract.DeleteStateBatch(&_IStateCommitmentChain.TransactOpts, _batchHeader)
}

// IStateCommitmentChainDistributeTssRewardIterator is returned from FilterDistributeTssReward and is used to iterate over the raw logs and unpacked data for DistributeTssReward events raised by the IStateCommitmentChain contract.
type IStateCommitmentChainDistributeTssRewardIterator struct {
	Event *IStateCommitmentChainDistributeTssReward // Event containing the contract specifics and raw log

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
func (it *IStateCommitmentChainDistributeTssRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStateCommitmentChainDistributeTssReward)
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
		it.Event = new(IStateCommitmentChainDistributeTssReward)
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
func (it *IStateCommitmentChainDistributeTssRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStateCommitmentChainDistributeTssRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStateCommitmentChainDistributeTssReward represents a DistributeTssReward event raised by the IStateCommitmentChain contract.
type IStateCommitmentChainDistributeTssReward struct {
	StartBlockNumber *big.Int
	Length           *big.Int
	BatchTime        *big.Int
	TssMembers       []common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributeTssReward is a free log retrieval operation binding the contract event 0xf533ef50019763ee9d95ad46e28350b533c11edd472ae7be93e8fae83c1b6d99.
//
// Solidity: event DistributeTssReward(uint256 indexed _startBlockNumber, uint256 _length, uint256 indexed _batchTime, address[] _tssMembers)
func (_IStateCommitmentChain *IStateCommitmentChainFilterer) FilterDistributeTssReward(opts *bind.FilterOpts, _startBlockNumber []*big.Int, _batchTime []*big.Int) (*IStateCommitmentChainDistributeTssRewardIterator, error) {

	var _startBlockNumberRule []interface{}
	for _, _startBlockNumberItem := range _startBlockNumber {
		_startBlockNumberRule = append(_startBlockNumberRule, _startBlockNumberItem)
	}

	var _batchTimeRule []interface{}
	for _, _batchTimeItem := range _batchTime {
		_batchTimeRule = append(_batchTimeRule, _batchTimeItem)
	}

	logs, sub, err := _IStateCommitmentChain.contract.FilterLogs(opts, "DistributeTssReward", _startBlockNumberRule, _batchTimeRule)
	if err != nil {
		return nil, err
	}
	return &IStateCommitmentChainDistributeTssRewardIterator{contract: _IStateCommitmentChain.contract, event: "DistributeTssReward", logs: logs, sub: sub}, nil
}

// WatchDistributeTssReward is a free log subscription operation binding the contract event 0xf533ef50019763ee9d95ad46e28350b533c11edd472ae7be93e8fae83c1b6d99.
//
// Solidity: event DistributeTssReward(uint256 indexed _startBlockNumber, uint256 _length, uint256 indexed _batchTime, address[] _tssMembers)
func (_IStateCommitmentChain *IStateCommitmentChainFilterer) WatchDistributeTssReward(opts *bind.WatchOpts, sink chan<- *IStateCommitmentChainDistributeTssReward, _startBlockNumber []*big.Int, _batchTime []*big.Int) (event.Subscription, error) {

	var _startBlockNumberRule []interface{}
	for _, _startBlockNumberItem := range _startBlockNumber {
		_startBlockNumberRule = append(_startBlockNumberRule, _startBlockNumberItem)
	}

	var _batchTimeRule []interface{}
	for _, _batchTimeItem := range _batchTime {
		_batchTimeRule = append(_batchTimeRule, _batchTimeItem)
	}

	logs, sub, err := _IStateCommitmentChain.contract.WatchLogs(opts, "DistributeTssReward", _startBlockNumberRule, _batchTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStateCommitmentChainDistributeTssReward)
				if err := _IStateCommitmentChain.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
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

// ParseDistributeTssReward is a log parse operation binding the contract event 0xf533ef50019763ee9d95ad46e28350b533c11edd472ae7be93e8fae83c1b6d99.
//
// Solidity: event DistributeTssReward(uint256 indexed _startBlockNumber, uint256 _length, uint256 indexed _batchTime, address[] _tssMembers)
func (_IStateCommitmentChain *IStateCommitmentChainFilterer) ParseDistributeTssReward(log types.Log) (*IStateCommitmentChainDistributeTssReward, error) {
	event := new(IStateCommitmentChainDistributeTssReward)
	if err := _IStateCommitmentChain.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStateCommitmentChainStateBatchAppendedIterator is returned from FilterStateBatchAppended and is used to iterate over the raw logs and unpacked data for StateBatchAppended events raised by the IStateCommitmentChain contract.
type IStateCommitmentChainStateBatchAppendedIterator struct {
	Event *IStateCommitmentChainStateBatchAppended // Event containing the contract specifics and raw log

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
func (it *IStateCommitmentChainStateBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStateCommitmentChainStateBatchAppended)
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
		it.Event = new(IStateCommitmentChainStateBatchAppended)
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
func (it *IStateCommitmentChainStateBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStateCommitmentChainStateBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStateCommitmentChainStateBatchAppended represents a StateBatchAppended event raised by the IStateCommitmentChain contract.
type IStateCommitmentChainStateBatchAppended struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	Signature         []byte
	ExtraData         []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStateBatchAppended is a free log retrieval operation binding the contract event 0x9cf3ad24eae3fd6d461e2f566b35b95b6d671871d9fcb45f8ac8030e4a8d21b3.
//
// Solidity: event StateBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _signature, bytes _extraData)
func (_IStateCommitmentChain *IStateCommitmentChainFilterer) FilterStateBatchAppended(opts *bind.FilterOpts, _batchIndex []*big.Int) (*IStateCommitmentChainStateBatchAppendedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _IStateCommitmentChain.contract.FilterLogs(opts, "StateBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &IStateCommitmentChainStateBatchAppendedIterator{contract: _IStateCommitmentChain.contract, event: "StateBatchAppended", logs: logs, sub: sub}, nil
}

// WatchStateBatchAppended is a free log subscription operation binding the contract event 0x9cf3ad24eae3fd6d461e2f566b35b95b6d671871d9fcb45f8ac8030e4a8d21b3.
//
// Solidity: event StateBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _signature, bytes _extraData)
func (_IStateCommitmentChain *IStateCommitmentChainFilterer) WatchStateBatchAppended(opts *bind.WatchOpts, sink chan<- *IStateCommitmentChainStateBatchAppended, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _IStateCommitmentChain.contract.WatchLogs(opts, "StateBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStateCommitmentChainStateBatchAppended)
				if err := _IStateCommitmentChain.contract.UnpackLog(event, "StateBatchAppended", log); err != nil {
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

// ParseStateBatchAppended is a log parse operation binding the contract event 0x9cf3ad24eae3fd6d461e2f566b35b95b6d671871d9fcb45f8ac8030e4a8d21b3.
//
// Solidity: event StateBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _signature, bytes _extraData)
func (_IStateCommitmentChain *IStateCommitmentChainFilterer) ParseStateBatchAppended(log types.Log) (*IStateCommitmentChainStateBatchAppended, error) {
	event := new(IStateCommitmentChainStateBatchAppended)
	if err := _IStateCommitmentChain.contract.UnpackLog(event, "StateBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStateCommitmentChainStateBatchDeletedIterator is returned from FilterStateBatchDeleted and is used to iterate over the raw logs and unpacked data for StateBatchDeleted events raised by the IStateCommitmentChain contract.
type IStateCommitmentChainStateBatchDeletedIterator struct {
	Event *IStateCommitmentChainStateBatchDeleted // Event containing the contract specifics and raw log

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
func (it *IStateCommitmentChainStateBatchDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStateCommitmentChainStateBatchDeleted)
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
		it.Event = new(IStateCommitmentChainStateBatchDeleted)
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
func (it *IStateCommitmentChainStateBatchDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStateCommitmentChainStateBatchDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStateCommitmentChainStateBatchDeleted represents a StateBatchDeleted event raised by the IStateCommitmentChain contract.
type IStateCommitmentChainStateBatchDeleted struct {
	BatchIndex *big.Int
	BatchRoot  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStateBatchDeleted is a free log retrieval operation binding the contract event 0x8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64.
//
// Solidity: event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_IStateCommitmentChain *IStateCommitmentChainFilterer) FilterStateBatchDeleted(opts *bind.FilterOpts, _batchIndex []*big.Int) (*IStateCommitmentChainStateBatchDeletedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _IStateCommitmentChain.contract.FilterLogs(opts, "StateBatchDeleted", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &IStateCommitmentChainStateBatchDeletedIterator{contract: _IStateCommitmentChain.contract, event: "StateBatchDeleted", logs: logs, sub: sub}, nil
}

// WatchStateBatchDeleted is a free log subscription operation binding the contract event 0x8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64.
//
// Solidity: event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_IStateCommitmentChain *IStateCommitmentChainFilterer) WatchStateBatchDeleted(opts *bind.WatchOpts, sink chan<- *IStateCommitmentChainStateBatchDeleted, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _IStateCommitmentChain.contract.WatchLogs(opts, "StateBatchDeleted", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStateCommitmentChainStateBatchDeleted)
				if err := _IStateCommitmentChain.contract.UnpackLog(event, "StateBatchDeleted", log); err != nil {
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

// ParseStateBatchDeleted is a log parse operation binding the contract event 0x8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64.
//
// Solidity: event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_IStateCommitmentChain *IStateCommitmentChainFilterer) ParseStateBatchDeleted(log types.Log) (*IStateCommitmentChainStateBatchDeleted, error) {
	event := new(IStateCommitmentChainStateBatchDeleted)
	if err := _IStateCommitmentChain.contract.UnpackLog(event, "StateBatchDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StateCommitmentChain

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

// StateCommitmentChainMetaData contains all meta data concerning the StateCommitmentChain contract.
var StateCommitmentChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_libAddressManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1messenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fraudProofWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sequencerPublishWindow\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_tssMembers\",\"type\":\"address[]\"}],\"name\":\"DistributeTssReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_batchSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prevTotalElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"StateBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchDeleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FRAUD_PROOF_WINDOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_PUBLISH_WINDOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_batch\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_shouldStartAtElement\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"appendStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"contractIChainStorageContainer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_BVMCodec.ChainBatchHeader\",\"name\":\"_batchHeader\",\"type\":\"tuple\"}],\"name\":\"deleteStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastSequencerTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastSequencerTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBatches\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalElements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_BVMCodec.ChainBatchHeader\",\"name\":\"_batchHeader\",\"type\":\"tuple\"}],\"name\":\"insideFraudProofWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_inside\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libAddressManager\",\"outputs\":[{\"internalType\":\"contractLib_AddressManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_element\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_BVMCodec.ChainBatchHeader\",\"name\":\"_batchHeader\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLib_BVMCodec.ChainInclusionProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyStateCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StateCommitmentChainABI is the input ABI used to generate the binding from.
// Deprecated: Use StateCommitmentChainMetaData.ABI instead.
var StateCommitmentChainABI = StateCommitmentChainMetaData.ABI

// StateCommitmentChain is an auto generated Go binding around an Ethereum contract.
type StateCommitmentChain struct {
	StateCommitmentChainCaller     // Read-only binding to the contract
	StateCommitmentChainTransactor // Write-only binding to the contract
	StateCommitmentChainFilterer   // Log filterer for contract events
}

// StateCommitmentChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateCommitmentChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateCommitmentChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateCommitmentChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateCommitmentChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateCommitmentChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateCommitmentChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateCommitmentChainSession struct {
	Contract     *StateCommitmentChain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StateCommitmentChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateCommitmentChainCallerSession struct {
	Contract *StateCommitmentChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// StateCommitmentChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateCommitmentChainTransactorSession struct {
	Contract     *StateCommitmentChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// StateCommitmentChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateCommitmentChainRaw struct {
	Contract *StateCommitmentChain // Generic contract binding to access the raw methods on
}

// StateCommitmentChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateCommitmentChainCallerRaw struct {
	Contract *StateCommitmentChainCaller // Generic read-only contract binding to access the raw methods on
}

// StateCommitmentChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateCommitmentChainTransactorRaw struct {
	Contract *StateCommitmentChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateCommitmentChain creates a new instance of StateCommitmentChain, bound to a specific deployed contract.
func NewStateCommitmentChain(address common.Address, backend bind.ContractBackend) (*StateCommitmentChain, error) {
	contract, err := bindStateCommitmentChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChain{StateCommitmentChainCaller: StateCommitmentChainCaller{contract: contract}, StateCommitmentChainTransactor: StateCommitmentChainTransactor{contract: contract}, StateCommitmentChainFilterer: StateCommitmentChainFilterer{contract: contract}}, nil
}

// NewStateCommitmentChainCaller creates a new read-only instance of StateCommitmentChain, bound to a specific deployed contract.
func NewStateCommitmentChainCaller(address common.Address, caller bind.ContractCaller) (*StateCommitmentChainCaller, error) {
	contract, err := bindStateCommitmentChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChainCaller{contract: contract}, nil
}

// NewStateCommitmentChainTransactor creates a new write-only instance of StateCommitmentChain, bound to a specific deployed contract.
func NewStateCommitmentChainTransactor(address common.Address, transactor bind.ContractTransactor) (*StateCommitmentChainTransactor, error) {
	contract, err := bindStateCommitmentChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChainTransactor{contract: contract}, nil
}

// NewStateCommitmentChainFilterer creates a new log filterer instance of StateCommitmentChain, bound to a specific deployed contract.
func NewStateCommitmentChainFilterer(address common.Address, filterer bind.ContractFilterer) (*StateCommitmentChainFilterer, error) {
	contract, err := bindStateCommitmentChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChainFilterer{contract: contract}, nil
}

// bindStateCommitmentChain binds a generic wrapper to an already deployed contract.
func bindStateCommitmentChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateCommitmentChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateCommitmentChain *StateCommitmentChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateCommitmentChain.Contract.StateCommitmentChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateCommitmentChain *StateCommitmentChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.StateCommitmentChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateCommitmentChain *StateCommitmentChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.StateCommitmentChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateCommitmentChain *StateCommitmentChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateCommitmentChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateCommitmentChain *StateCommitmentChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateCommitmentChain *StateCommitmentChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.contract.Transact(opts, method, params...)
}

// FRAUDPROOFWINDOW is a free data retrieval call binding the contract method 0xc17b291b.
//
// Solidity: function FRAUD_PROOF_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainCaller) FRAUDPROOFWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "FRAUD_PROOF_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FRAUDPROOFWINDOW is a free data retrieval call binding the contract method 0xc17b291b.
//
// Solidity: function FRAUD_PROOF_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainSession) FRAUDPROOFWINDOW() (*big.Int, error) {
	return _StateCommitmentChain.Contract.FRAUDPROOFWINDOW(&_StateCommitmentChain.CallOpts)
}

// FRAUDPROOFWINDOW is a free data retrieval call binding the contract method 0xc17b291b.
//
// Solidity: function FRAUD_PROOF_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) FRAUDPROOFWINDOW() (*big.Int, error) {
	return _StateCommitmentChain.Contract.FRAUDPROOFWINDOW(&_StateCommitmentChain.CallOpts)
}

// SEQUENCERPUBLISHWINDOW is a free data retrieval call binding the contract method 0x81eb62ef.
//
// Solidity: function SEQUENCER_PUBLISH_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainCaller) SEQUENCERPUBLISHWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "SEQUENCER_PUBLISH_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SEQUENCERPUBLISHWINDOW is a free data retrieval call binding the contract method 0x81eb62ef.
//
// Solidity: function SEQUENCER_PUBLISH_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainSession) SEQUENCERPUBLISHWINDOW() (*big.Int, error) {
	return _StateCommitmentChain.Contract.SEQUENCERPUBLISHWINDOW(&_StateCommitmentChain.CallOpts)
}

// SEQUENCERPUBLISHWINDOW is a free data retrieval call binding the contract method 0x81eb62ef.
//
// Solidity: function SEQUENCER_PUBLISH_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) SEQUENCERPUBLISHWINDOW() (*big.Int, error) {
	return _StateCommitmentChain.Contract.SEQUENCERPUBLISHWINDOW(&_StateCommitmentChain.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCaller) Batches(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "batches")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainSession) Batches() (common.Address, error) {
	return _StateCommitmentChain.Contract.Batches(&_StateCommitmentChain.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) Batches() (common.Address, error) {
	return _StateCommitmentChain.Contract.Batches(&_StateCommitmentChain.CallOpts)
}

// GetLastSequencerTimestamp is a free data retrieval call binding the contract method 0x7ad168a0.
//
// Solidity: function getLastSequencerTimestamp() view returns(uint256 _lastSequencerTimestamp)
func (_StateCommitmentChain *StateCommitmentChainCaller) GetLastSequencerTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "getLastSequencerTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastSequencerTimestamp is a free data retrieval call binding the contract method 0x7ad168a0.
//
// Solidity: function getLastSequencerTimestamp() view returns(uint256 _lastSequencerTimestamp)
func (_StateCommitmentChain *StateCommitmentChainSession) GetLastSequencerTimestamp() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetLastSequencerTimestamp(&_StateCommitmentChain.CallOpts)
}

// GetLastSequencerTimestamp is a free data retrieval call binding the contract method 0x7ad168a0.
//
// Solidity: function getLastSequencerTimestamp() view returns(uint256 _lastSequencerTimestamp)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) GetLastSequencerTimestamp() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetLastSequencerTimestamp(&_StateCommitmentChain.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_StateCommitmentChain *StateCommitmentChainCaller) GetTotalBatches(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "getTotalBatches")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_StateCommitmentChain *StateCommitmentChainSession) GetTotalBatches() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetTotalBatches(&_StateCommitmentChain.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) GetTotalBatches() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetTotalBatches(&_StateCommitmentChain.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_StateCommitmentChain *StateCommitmentChainCaller) GetTotalElements(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "getTotalElements")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_StateCommitmentChain *StateCommitmentChainSession) GetTotalElements() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetTotalElements(&_StateCommitmentChain.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) GetTotalElements() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetTotalElements(&_StateCommitmentChain.CallOpts)
}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0x89a1d980.
//
// Solidity: function insideFraudProofWindow((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) view returns(bool _inside)
func (_StateCommitmentChain *StateCommitmentChainCaller) InsideFraudProofWindow(opts *bind.CallOpts, _batchHeader LibBVMCodecChainBatchHeader) (bool, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "insideFraudProofWindow", _batchHeader)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0x89a1d980.
//
// Solidity: function insideFraudProofWindow((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) view returns(bool _inside)
func (_StateCommitmentChain *StateCommitmentChainSession) InsideFraudProofWindow(_batchHeader LibBVMCodecChainBatchHeader) (bool, error) {
	return _StateCommitmentChain.Contract.InsideFraudProofWindow(&_StateCommitmentChain.CallOpts, _batchHeader)
}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0x89a1d980.
//
// Solidity: function insideFraudProofWindow((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) view returns(bool _inside)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) InsideFraudProofWindow(_batchHeader LibBVMCodecChainBatchHeader) (bool, error) {
	return _StateCommitmentChain.Contract.InsideFraudProofWindow(&_StateCommitmentChain.CallOpts, _batchHeader)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCaller) LibAddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "libAddressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainSession) LibAddressManager() (common.Address, error) {
	return _StateCommitmentChain.Contract.LibAddressManager(&_StateCommitmentChain.CallOpts)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) LibAddressManager() (common.Address, error) {
	return _StateCommitmentChain.Contract.LibAddressManager(&_StateCommitmentChain.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainSession) Messenger() (common.Address, error) {
	return _StateCommitmentChain.Contract.Messenger(&_StateCommitmentChain.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) Messenger() (common.Address, error) {
	return _StateCommitmentChain.Contract.Messenger(&_StateCommitmentChain.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCaller) Resolve(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "resolve", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_StateCommitmentChain *StateCommitmentChainSession) Resolve(_name string) (common.Address, error) {
	return _StateCommitmentChain.Contract.Resolve(&_StateCommitmentChain.CallOpts, _name)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) Resolve(_name string) (common.Address, error) {
	return _StateCommitmentChain.Contract.Resolve(&_StateCommitmentChain.CallOpts, _name)
}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0xb768bb17.
//
// Solidity: function verifyStateCommitment(bytes32 _element, (uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader, (uint256,bytes32[]) _proof) view returns(bool)
func (_StateCommitmentChain *StateCommitmentChainCaller) VerifyStateCommitment(opts *bind.CallOpts, _element [32]byte, _batchHeader LibBVMCodecChainBatchHeader, _proof LibBVMCodecChainInclusionProof) (bool, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "verifyStateCommitment", _element, _batchHeader, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0xb768bb17.
//
// Solidity: function verifyStateCommitment(bytes32 _element, (uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader, (uint256,bytes32[]) _proof) view returns(bool)
func (_StateCommitmentChain *StateCommitmentChainSession) VerifyStateCommitment(_element [32]byte, _batchHeader LibBVMCodecChainBatchHeader, _proof LibBVMCodecChainInclusionProof) (bool, error) {
	return _StateCommitmentChain.Contract.VerifyStateCommitment(&_StateCommitmentChain.CallOpts, _element, _batchHeader, _proof)
}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0xb768bb17.
//
// Solidity: function verifyStateCommitment(bytes32 _element, (uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader, (uint256,bytes32[]) _proof) view returns(bool)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) VerifyStateCommitment(_element [32]byte, _batchHeader LibBVMCodecChainBatchHeader, _proof LibBVMCodecChainInclusionProof) (bool, error) {
	return _StateCommitmentChain.Contract.VerifyStateCommitment(&_StateCommitmentChain.CallOpts, _element, _batchHeader, _proof)
}

// AppendStateBatch is a paid mutator transaction binding the contract method 0x2169f79f.
//
// Solidity: function appendStateBatch(bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_StateCommitmentChain *StateCommitmentChainTransactor) AppendStateBatch(opts *bind.TransactOpts, _batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _StateCommitmentChain.contract.Transact(opts, "appendStateBatch", _batch, _shouldStartAtElement, _signature)
}

// AppendStateBatch is a paid mutator transaction binding the contract method 0x2169f79f.
//
// Solidity: function appendStateBatch(bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_StateCommitmentChain *StateCommitmentChainSession) AppendStateBatch(_batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.AppendStateBatch(&_StateCommitmentChain.TransactOpts, _batch, _shouldStartAtElement, _signature)
}

// AppendStateBatch is a paid mutator transaction binding the contract method 0x2169f79f.
//
// Solidity: function appendStateBatch(bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) returns()
func (_StateCommitmentChain *StateCommitmentChainTransactorSession) AppendStateBatch(_batch [][32]byte, _shouldStartAtElement *big.Int, _signature []byte) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.AppendStateBatch(&_StateCommitmentChain.TransactOpts, _batch, _shouldStartAtElement, _signature)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0x5b4d90e2.
//
// Solidity: function deleteStateBatch((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) returns()
func (_StateCommitmentChain *StateCommitmentChainTransactor) DeleteStateBatch(opts *bind.TransactOpts, _batchHeader LibBVMCodecChainBatchHeader) (*types.Transaction, error) {
	return _StateCommitmentChain.contract.Transact(opts, "deleteStateBatch", _batchHeader)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0x5b4d90e2.
//
// Solidity: function deleteStateBatch((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) returns()
func (_StateCommitmentChain *StateCommitmentChainSession) DeleteStateBatch(_batchHeader LibBVMCodecChainBatchHeader) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.DeleteStateBatch(&_StateCommitmentChain.TransactOpts, _batchHeader)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0x5b4d90e2.
//
// Solidity: function deleteStateBatch((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) returns()
func (_StateCommitmentChain *StateCommitmentChainTransactorSession) DeleteStateBatch(_batchHeader LibBVMCodecChainBatchHeader) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.DeleteStateBatch(&_StateCommitmentChain.TransactOpts, _batchHeader)
}

// StateCommitmentChainDistributeTssRewardIterator is returned from FilterDistributeTssReward and is used to iterate over the raw logs and unpacked data for DistributeTssReward events raised by the StateCommitmentChain contract.
type StateCommitmentChainDistributeTssRewardIterator struct {
	Event *StateCommitmentChainDistributeTssReward // Event containing the contract specifics and raw log

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
func (it *StateCommitmentChainDistributeTssRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateCommitmentChainDistributeTssReward)
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
		it.Event = new(StateCommitmentChainDistributeTssReward)
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
func (it *StateCommitmentChainDistributeTssRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateCommitmentChainDistributeTssRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateCommitmentChainDistributeTssReward represents a DistributeTssReward event raised by the StateCommitmentChain contract.
type StateCommitmentChainDistributeTssReward struct {
	StartBlockNumber *big.Int
	Length           *big.Int
	BatchTime        *big.Int
	TssMembers       []common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributeTssReward is a free log retrieval operation binding the contract event 0xf533ef50019763ee9d95ad46e28350b533c11edd472ae7be93e8fae83c1b6d99.
//
// Solidity: event DistributeTssReward(uint256 indexed _startBlockNumber, uint256 _length, uint256 indexed _batchTime, address[] _tssMembers)
func (_StateCommitmentChain *StateCommitmentChainFilterer) FilterDistributeTssReward(opts *bind.FilterOpts, _startBlockNumber []*big.Int, _batchTime []*big.Int) (*StateCommitmentChainDistributeTssRewardIterator, error) {

	var _startBlockNumberRule []interface{}
	for _, _startBlockNumberItem := range _startBlockNumber {
		_startBlockNumberRule = append(_startBlockNumberRule, _startBlockNumberItem)
	}

	var _batchTimeRule []interface{}
	for _, _batchTimeItem := range _batchTime {
		_batchTimeRule = append(_batchTimeRule, _batchTimeItem)
	}

	logs, sub, err := _StateCommitmentChain.contract.FilterLogs(opts, "DistributeTssReward", _startBlockNumberRule, _batchTimeRule)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChainDistributeTssRewardIterator{contract: _StateCommitmentChain.contract, event: "DistributeTssReward", logs: logs, sub: sub}, nil
}

// WatchDistributeTssReward is a free log subscription operation binding the contract event 0xf533ef50019763ee9d95ad46e28350b533c11edd472ae7be93e8fae83c1b6d99.
//
// Solidity: event DistributeTssReward(uint256 indexed _startBlockNumber, uint256 _length, uint256 indexed _batchTime, address[] _tssMembers)
func (_StateCommitmentChain *StateCommitmentChainFilterer) WatchDistributeTssReward(opts *bind.WatchOpts, sink chan<- *StateCommitmentChainDistributeTssReward, _startBlockNumber []*big.Int, _batchTime []*big.Int) (event.Subscription, error) {

	var _startBlockNumberRule []interface{}
	for _, _startBlockNumberItem := range _startBlockNumber {
		_startBlockNumberRule = append(_startBlockNumberRule, _startBlockNumberItem)
	}

	var _batchTimeRule []interface{}
	for _, _batchTimeItem := range _batchTime {
		_batchTimeRule = append(_batchTimeRule, _batchTimeItem)
	}

	logs, sub, err := _StateCommitmentChain.contract.WatchLogs(opts, "DistributeTssReward", _startBlockNumberRule, _batchTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateCommitmentChainDistributeTssReward)
				if err := _StateCommitmentChain.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
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
func (_StateCommitmentChain *StateCommitmentChainFilterer) ParseDistributeTssReward(log types.Log) (*StateCommitmentChainDistributeTssReward, error) {
	event := new(StateCommitmentChainDistributeTssReward)
	if err := _StateCommitmentChain.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateCommitmentChainStateBatchAppendedIterator is returned from FilterStateBatchAppended and is used to iterate over the raw logs and unpacked data for StateBatchAppended events raised by the StateCommitmentChain contract.
type StateCommitmentChainStateBatchAppendedIterator struct {
	Event *StateCommitmentChainStateBatchAppended // Event containing the contract specifics and raw log

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
func (it *StateCommitmentChainStateBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateCommitmentChainStateBatchAppended)
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
		it.Event = new(StateCommitmentChainStateBatchAppended)
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
func (it *StateCommitmentChainStateBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateCommitmentChainStateBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateCommitmentChainStateBatchAppended represents a StateBatchAppended event raised by the StateCommitmentChain contract.
type StateCommitmentChainStateBatchAppended struct {
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
func (_StateCommitmentChain *StateCommitmentChainFilterer) FilterStateBatchAppended(opts *bind.FilterOpts, _batchIndex []*big.Int) (*StateCommitmentChainStateBatchAppendedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _StateCommitmentChain.contract.FilterLogs(opts, "StateBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChainStateBatchAppendedIterator{contract: _StateCommitmentChain.contract, event: "StateBatchAppended", logs: logs, sub: sub}, nil
}

// WatchStateBatchAppended is a free log subscription operation binding the contract event 0x9cf3ad24eae3fd6d461e2f566b35b95b6d671871d9fcb45f8ac8030e4a8d21b3.
//
// Solidity: event StateBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _signature, bytes _extraData)
func (_StateCommitmentChain *StateCommitmentChainFilterer) WatchStateBatchAppended(opts *bind.WatchOpts, sink chan<- *StateCommitmentChainStateBatchAppended, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _StateCommitmentChain.contract.WatchLogs(opts, "StateBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateCommitmentChainStateBatchAppended)
				if err := _StateCommitmentChain.contract.UnpackLog(event, "StateBatchAppended", log); err != nil {
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
func (_StateCommitmentChain *StateCommitmentChainFilterer) ParseStateBatchAppended(log types.Log) (*StateCommitmentChainStateBatchAppended, error) {
	event := new(StateCommitmentChainStateBatchAppended)
	if err := _StateCommitmentChain.contract.UnpackLog(event, "StateBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateCommitmentChainStateBatchDeletedIterator is returned from FilterStateBatchDeleted and is used to iterate over the raw logs and unpacked data for StateBatchDeleted events raised by the StateCommitmentChain contract.
type StateCommitmentChainStateBatchDeletedIterator struct {
	Event *StateCommitmentChainStateBatchDeleted // Event containing the contract specifics and raw log

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
func (it *StateCommitmentChainStateBatchDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateCommitmentChainStateBatchDeleted)
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
		it.Event = new(StateCommitmentChainStateBatchDeleted)
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
func (it *StateCommitmentChainStateBatchDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateCommitmentChainStateBatchDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateCommitmentChainStateBatchDeleted represents a StateBatchDeleted event raised by the StateCommitmentChain contract.
type StateCommitmentChainStateBatchDeleted struct {
	BatchIndex *big.Int
	BatchRoot  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStateBatchDeleted is a free log retrieval operation binding the contract event 0x8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64.
//
// Solidity: event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_StateCommitmentChain *StateCommitmentChainFilterer) FilterStateBatchDeleted(opts *bind.FilterOpts, _batchIndex []*big.Int) (*StateCommitmentChainStateBatchDeletedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _StateCommitmentChain.contract.FilterLogs(opts, "StateBatchDeleted", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChainStateBatchDeletedIterator{contract: _StateCommitmentChain.contract, event: "StateBatchDeleted", logs: logs, sub: sub}, nil
}

// WatchStateBatchDeleted is a free log subscription operation binding the contract event 0x8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64.
//
// Solidity: event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_StateCommitmentChain *StateCommitmentChainFilterer) WatchStateBatchDeleted(opts *bind.WatchOpts, sink chan<- *StateCommitmentChainStateBatchDeleted, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _StateCommitmentChain.contract.WatchLogs(opts, "StateBatchDeleted", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateCommitmentChainStateBatchDeleted)
				if err := _StateCommitmentChain.contract.UnpackLog(event, "StateBatchDeleted", log); err != nil {
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
func (_StateCommitmentChain *StateCommitmentChainFilterer) ParseStateBatchDeleted(log types.Log) (*StateCommitmentChainStateBatchDeleted, error) {
	event := new(StateCommitmentChainStateBatchDeleted)
	if err := _StateCommitmentChain.contract.UnpackLog(event, "StateBatchDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

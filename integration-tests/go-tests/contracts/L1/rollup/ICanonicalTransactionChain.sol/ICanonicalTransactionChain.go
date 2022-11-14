// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ICanonicalTransactionChain

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

// LibBVMCodecQueueElement is an auto generated low-level Go binding around an user-defined struct.
type LibBVMCodecQueueElement struct {
	TransactionHash [32]byte
	Timestamp       *big.Int
	BlockNumber     *big.Int
}

// ICanonicalTransactionChainMetaData contains all meta data concerning the ICanonicalTransactionChain contract.
var ICanonicalTransactionChainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l2GasDiscountDivisor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"enqueueGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"enqueueL2GasPrepaid\",\"type\":\"uint256\"}],\"name\":\"L2GasParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startingQueueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_numQueueElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"name\":\"QueueBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startingQueueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_numQueueElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"name\":\"SequencerBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_batchSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prevTotalElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"TransactionBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1TxOrigin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"TransactionEnqueued\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"appendSequencerBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"contractIChainStorageContainer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"enqueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockNumber\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastTimestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextQueueIndex\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumPendingQueueElements\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getQueueElement\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"}],\"internalType\":\"structLib_BVMCodec.QueueElement\",\"name\":\"_element\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQueueLength\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBatches\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalElements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2GasDiscountDivisor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_enqueueGasCost\",\"type\":\"uint256\"}],\"name\":\"setGasParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ICanonicalTransactionChainABI is the input ABI used to generate the binding from.
// Deprecated: Use ICanonicalTransactionChainMetaData.ABI instead.
var ICanonicalTransactionChainABI = ICanonicalTransactionChainMetaData.ABI

// ICanonicalTransactionChain is an auto generated Go binding around an Ethereum contract.
type ICanonicalTransactionChain struct {
	ICanonicalTransactionChainCaller     // Read-only binding to the contract
	ICanonicalTransactionChainTransactor // Write-only binding to the contract
	ICanonicalTransactionChainFilterer   // Log filterer for contract events
}

// ICanonicalTransactionChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICanonicalTransactionChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICanonicalTransactionChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICanonicalTransactionChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICanonicalTransactionChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICanonicalTransactionChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICanonicalTransactionChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICanonicalTransactionChainSession struct {
	Contract     *ICanonicalTransactionChain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ICanonicalTransactionChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICanonicalTransactionChainCallerSession struct {
	Contract *ICanonicalTransactionChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ICanonicalTransactionChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICanonicalTransactionChainTransactorSession struct {
	Contract     *ICanonicalTransactionChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ICanonicalTransactionChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICanonicalTransactionChainRaw struct {
	Contract *ICanonicalTransactionChain // Generic contract binding to access the raw methods on
}

// ICanonicalTransactionChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICanonicalTransactionChainCallerRaw struct {
	Contract *ICanonicalTransactionChainCaller // Generic read-only contract binding to access the raw methods on
}

// ICanonicalTransactionChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICanonicalTransactionChainTransactorRaw struct {
	Contract *ICanonicalTransactionChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICanonicalTransactionChain creates a new instance of ICanonicalTransactionChain, bound to a specific deployed contract.
func NewICanonicalTransactionChain(address common.Address, backend bind.ContractBackend) (*ICanonicalTransactionChain, error) {
	contract, err := bindICanonicalTransactionChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICanonicalTransactionChain{ICanonicalTransactionChainCaller: ICanonicalTransactionChainCaller{contract: contract}, ICanonicalTransactionChainTransactor: ICanonicalTransactionChainTransactor{contract: contract}, ICanonicalTransactionChainFilterer: ICanonicalTransactionChainFilterer{contract: contract}}, nil
}

// NewICanonicalTransactionChainCaller creates a new read-only instance of ICanonicalTransactionChain, bound to a specific deployed contract.
func NewICanonicalTransactionChainCaller(address common.Address, caller bind.ContractCaller) (*ICanonicalTransactionChainCaller, error) {
	contract, err := bindICanonicalTransactionChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICanonicalTransactionChainCaller{contract: contract}, nil
}

// NewICanonicalTransactionChainTransactor creates a new write-only instance of ICanonicalTransactionChain, bound to a specific deployed contract.
func NewICanonicalTransactionChainTransactor(address common.Address, transactor bind.ContractTransactor) (*ICanonicalTransactionChainTransactor, error) {
	contract, err := bindICanonicalTransactionChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICanonicalTransactionChainTransactor{contract: contract}, nil
}

// NewICanonicalTransactionChainFilterer creates a new log filterer instance of ICanonicalTransactionChain, bound to a specific deployed contract.
func NewICanonicalTransactionChainFilterer(address common.Address, filterer bind.ContractFilterer) (*ICanonicalTransactionChainFilterer, error) {
	contract, err := bindICanonicalTransactionChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICanonicalTransactionChainFilterer{contract: contract}, nil
}

// bindICanonicalTransactionChain binds a generic wrapper to an already deployed contract.
func bindICanonicalTransactionChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICanonicalTransactionChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICanonicalTransactionChain *ICanonicalTransactionChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICanonicalTransactionChain.Contract.ICanonicalTransactionChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICanonicalTransactionChain *ICanonicalTransactionChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICanonicalTransactionChain.Contract.ICanonicalTransactionChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICanonicalTransactionChain *ICanonicalTransactionChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICanonicalTransactionChain.Contract.ICanonicalTransactionChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICanonicalTransactionChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICanonicalTransactionChain *ICanonicalTransactionChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICanonicalTransactionChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICanonicalTransactionChain *ICanonicalTransactionChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICanonicalTransactionChain.Contract.contract.Transact(opts, method, params...)
}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCaller) Batches(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICanonicalTransactionChain.contract.Call(opts, &out, "batches")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) Batches() (common.Address, error) {
	return _ICanonicalTransactionChain.Contract.Batches(&_ICanonicalTransactionChain.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCallerSession) Batches() (common.Address, error) {
	return _ICanonicalTransactionChain.Contract.Batches(&_ICanonicalTransactionChain.CallOpts)
}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCaller) GetLastBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICanonicalTransactionChain.contract.Call(opts, &out, "getLastBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) GetLastBlockNumber() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetLastBlockNumber(&_ICanonicalTransactionChain.CallOpts)
}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCallerSession) GetLastBlockNumber() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetLastBlockNumber(&_ICanonicalTransactionChain.CallOpts)
}

// GetLastTimestamp is a free data retrieval call binding the contract method 0x37899770.
//
// Solidity: function getLastTimestamp() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCaller) GetLastTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICanonicalTransactionChain.contract.Call(opts, &out, "getLastTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastTimestamp is a free data retrieval call binding the contract method 0x37899770.
//
// Solidity: function getLastTimestamp() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) GetLastTimestamp() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetLastTimestamp(&_ICanonicalTransactionChain.CallOpts)
}

// GetLastTimestamp is a free data retrieval call binding the contract method 0x37899770.
//
// Solidity: function getLastTimestamp() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCallerSession) GetLastTimestamp() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetLastTimestamp(&_ICanonicalTransactionChain.CallOpts)
}

// GetNextQueueIndex is a free data retrieval call binding the contract method 0x7a167a8a.
//
// Solidity: function getNextQueueIndex() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCaller) GetNextQueueIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICanonicalTransactionChain.contract.Call(opts, &out, "getNextQueueIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextQueueIndex is a free data retrieval call binding the contract method 0x7a167a8a.
//
// Solidity: function getNextQueueIndex() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) GetNextQueueIndex() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetNextQueueIndex(&_ICanonicalTransactionChain.CallOpts)
}

// GetNextQueueIndex is a free data retrieval call binding the contract method 0x7a167a8a.
//
// Solidity: function getNextQueueIndex() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCallerSession) GetNextQueueIndex() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetNextQueueIndex(&_ICanonicalTransactionChain.CallOpts)
}

// GetNumPendingQueueElements is a free data retrieval call binding the contract method 0xf722b41a.
//
// Solidity: function getNumPendingQueueElements() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCaller) GetNumPendingQueueElements(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICanonicalTransactionChain.contract.Call(opts, &out, "getNumPendingQueueElements")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumPendingQueueElements is a free data retrieval call binding the contract method 0xf722b41a.
//
// Solidity: function getNumPendingQueueElements() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) GetNumPendingQueueElements() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetNumPendingQueueElements(&_ICanonicalTransactionChain.CallOpts)
}

// GetNumPendingQueueElements is a free data retrieval call binding the contract method 0xf722b41a.
//
// Solidity: function getNumPendingQueueElements() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCallerSession) GetNumPendingQueueElements() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetNumPendingQueueElements(&_ICanonicalTransactionChain.CallOpts)
}

// GetQueueElement is a free data retrieval call binding the contract method 0x2a7f18be.
//
// Solidity: function getQueueElement(uint256 _index) view returns((bytes32,uint40,uint40) _element)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCaller) GetQueueElement(opts *bind.CallOpts, _index *big.Int) (LibBVMCodecQueueElement, error) {
	var out []interface{}
	err := _ICanonicalTransactionChain.contract.Call(opts, &out, "getQueueElement", _index)

	if err != nil {
		return *new(LibBVMCodecQueueElement), err
	}

	out0 := *abi.ConvertType(out[0], new(LibBVMCodecQueueElement)).(*LibBVMCodecQueueElement)

	return out0, err

}

// GetQueueElement is a free data retrieval call binding the contract method 0x2a7f18be.
//
// Solidity: function getQueueElement(uint256 _index) view returns((bytes32,uint40,uint40) _element)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) GetQueueElement(_index *big.Int) (LibBVMCodecQueueElement, error) {
	return _ICanonicalTransactionChain.Contract.GetQueueElement(&_ICanonicalTransactionChain.CallOpts, _index)
}

// GetQueueElement is a free data retrieval call binding the contract method 0x2a7f18be.
//
// Solidity: function getQueueElement(uint256 _index) view returns((bytes32,uint40,uint40) _element)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCallerSession) GetQueueElement(_index *big.Int) (LibBVMCodecQueueElement, error) {
	return _ICanonicalTransactionChain.Contract.GetQueueElement(&_ICanonicalTransactionChain.CallOpts, _index)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCaller) GetQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICanonicalTransactionChain.contract.Call(opts, &out, "getQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) GetQueueLength() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetQueueLength(&_ICanonicalTransactionChain.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint40)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCallerSession) GetQueueLength() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetQueueLength(&_ICanonicalTransactionChain.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCaller) GetTotalBatches(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICanonicalTransactionChain.contract.Call(opts, &out, "getTotalBatches")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) GetTotalBatches() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetTotalBatches(&_ICanonicalTransactionChain.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCallerSession) GetTotalBatches() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetTotalBatches(&_ICanonicalTransactionChain.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCaller) GetTotalElements(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICanonicalTransactionChain.contract.Call(opts, &out, "getTotalElements")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) GetTotalElements() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetTotalElements(&_ICanonicalTransactionChain.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainCallerSession) GetTotalElements() (*big.Int, error) {
	return _ICanonicalTransactionChain.Contract.GetTotalElements(&_ICanonicalTransactionChain.CallOpts)
}

// AppendSequencerBatch is a paid mutator transaction binding the contract method 0xd0f89344.
//
// Solidity: function appendSequencerBatch() returns()
func (_ICanonicalTransactionChain *ICanonicalTransactionChainTransactor) AppendSequencerBatch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICanonicalTransactionChain.contract.Transact(opts, "appendSequencerBatch")
}

// AppendSequencerBatch is a paid mutator transaction binding the contract method 0xd0f89344.
//
// Solidity: function appendSequencerBatch() returns()
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) AppendSequencerBatch() (*types.Transaction, error) {
	return _ICanonicalTransactionChain.Contract.AppendSequencerBatch(&_ICanonicalTransactionChain.TransactOpts)
}

// AppendSequencerBatch is a paid mutator transaction binding the contract method 0xd0f89344.
//
// Solidity: function appendSequencerBatch() returns()
func (_ICanonicalTransactionChain *ICanonicalTransactionChainTransactorSession) AppendSequencerBatch() (*types.Transaction, error) {
	return _ICanonicalTransactionChain.Contract.AppendSequencerBatch(&_ICanonicalTransactionChain.TransactOpts)
}

// Enqueue is a paid mutator transaction binding the contract method 0x6fee07e0.
//
// Solidity: function enqueue(address _target, uint256 _gasLimit, bytes _data) returns()
func (_ICanonicalTransactionChain *ICanonicalTransactionChainTransactor) Enqueue(opts *bind.TransactOpts, _target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _ICanonicalTransactionChain.contract.Transact(opts, "enqueue", _target, _gasLimit, _data)
}

// Enqueue is a paid mutator transaction binding the contract method 0x6fee07e0.
//
// Solidity: function enqueue(address _target, uint256 _gasLimit, bytes _data) returns()
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) Enqueue(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _ICanonicalTransactionChain.Contract.Enqueue(&_ICanonicalTransactionChain.TransactOpts, _target, _gasLimit, _data)
}

// Enqueue is a paid mutator transaction binding the contract method 0x6fee07e0.
//
// Solidity: function enqueue(address _target, uint256 _gasLimit, bytes _data) returns()
func (_ICanonicalTransactionChain *ICanonicalTransactionChainTransactorSession) Enqueue(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _ICanonicalTransactionChain.Contract.Enqueue(&_ICanonicalTransactionChain.TransactOpts, _target, _gasLimit, _data)
}

// SetGasParams is a paid mutator transaction binding the contract method 0xedcc4a45.
//
// Solidity: function setGasParams(uint256 _l2GasDiscountDivisor, uint256 _enqueueGasCost) returns()
func (_ICanonicalTransactionChain *ICanonicalTransactionChainTransactor) SetGasParams(opts *bind.TransactOpts, _l2GasDiscountDivisor *big.Int, _enqueueGasCost *big.Int) (*types.Transaction, error) {
	return _ICanonicalTransactionChain.contract.Transact(opts, "setGasParams", _l2GasDiscountDivisor, _enqueueGasCost)
}

// SetGasParams is a paid mutator transaction binding the contract method 0xedcc4a45.
//
// Solidity: function setGasParams(uint256 _l2GasDiscountDivisor, uint256 _enqueueGasCost) returns()
func (_ICanonicalTransactionChain *ICanonicalTransactionChainSession) SetGasParams(_l2GasDiscountDivisor *big.Int, _enqueueGasCost *big.Int) (*types.Transaction, error) {
	return _ICanonicalTransactionChain.Contract.SetGasParams(&_ICanonicalTransactionChain.TransactOpts, _l2GasDiscountDivisor, _enqueueGasCost)
}

// SetGasParams is a paid mutator transaction binding the contract method 0xedcc4a45.
//
// Solidity: function setGasParams(uint256 _l2GasDiscountDivisor, uint256 _enqueueGasCost) returns()
func (_ICanonicalTransactionChain *ICanonicalTransactionChainTransactorSession) SetGasParams(_l2GasDiscountDivisor *big.Int, _enqueueGasCost *big.Int) (*types.Transaction, error) {
	return _ICanonicalTransactionChain.Contract.SetGasParams(&_ICanonicalTransactionChain.TransactOpts, _l2GasDiscountDivisor, _enqueueGasCost)
}

// ICanonicalTransactionChainL2GasParamsUpdatedIterator is returned from FilterL2GasParamsUpdated and is used to iterate over the raw logs and unpacked data for L2GasParamsUpdated events raised by the ICanonicalTransactionChain contract.
type ICanonicalTransactionChainL2GasParamsUpdatedIterator struct {
	Event *ICanonicalTransactionChainL2GasParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ICanonicalTransactionChainL2GasParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICanonicalTransactionChainL2GasParamsUpdated)
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
		it.Event = new(ICanonicalTransactionChainL2GasParamsUpdated)
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
func (it *ICanonicalTransactionChainL2GasParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICanonicalTransactionChainL2GasParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICanonicalTransactionChainL2GasParamsUpdated represents a L2GasParamsUpdated event raised by the ICanonicalTransactionChain contract.
type ICanonicalTransactionChainL2GasParamsUpdated struct {
	L2GasDiscountDivisor *big.Int
	EnqueueGasCost       *big.Int
	EnqueueL2GasPrepaid  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterL2GasParamsUpdated is a free log retrieval operation binding the contract event 0xc6ed75e96b8b18b71edc1a6e82a9d677f8268c774a262c624eeb2cf0a8b3e07e.
//
// Solidity: event L2GasParamsUpdated(uint256 l2GasDiscountDivisor, uint256 enqueueGasCost, uint256 enqueueL2GasPrepaid)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) FilterL2GasParamsUpdated(opts *bind.FilterOpts) (*ICanonicalTransactionChainL2GasParamsUpdatedIterator, error) {

	logs, sub, err := _ICanonicalTransactionChain.contract.FilterLogs(opts, "L2GasParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &ICanonicalTransactionChainL2GasParamsUpdatedIterator{contract: _ICanonicalTransactionChain.contract, event: "L2GasParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchL2GasParamsUpdated is a free log subscription operation binding the contract event 0xc6ed75e96b8b18b71edc1a6e82a9d677f8268c774a262c624eeb2cf0a8b3e07e.
//
// Solidity: event L2GasParamsUpdated(uint256 l2GasDiscountDivisor, uint256 enqueueGasCost, uint256 enqueueL2GasPrepaid)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) WatchL2GasParamsUpdated(opts *bind.WatchOpts, sink chan<- *ICanonicalTransactionChainL2GasParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _ICanonicalTransactionChain.contract.WatchLogs(opts, "L2GasParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICanonicalTransactionChainL2GasParamsUpdated)
				if err := _ICanonicalTransactionChain.contract.UnpackLog(event, "L2GasParamsUpdated", log); err != nil {
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

// ParseL2GasParamsUpdated is a log parse operation binding the contract event 0xc6ed75e96b8b18b71edc1a6e82a9d677f8268c774a262c624eeb2cf0a8b3e07e.
//
// Solidity: event L2GasParamsUpdated(uint256 l2GasDiscountDivisor, uint256 enqueueGasCost, uint256 enqueueL2GasPrepaid)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) ParseL2GasParamsUpdated(log types.Log) (*ICanonicalTransactionChainL2GasParamsUpdated, error) {
	event := new(ICanonicalTransactionChainL2GasParamsUpdated)
	if err := _ICanonicalTransactionChain.contract.UnpackLog(event, "L2GasParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICanonicalTransactionChainQueueBatchAppendedIterator is returned from FilterQueueBatchAppended and is used to iterate over the raw logs and unpacked data for QueueBatchAppended events raised by the ICanonicalTransactionChain contract.
type ICanonicalTransactionChainQueueBatchAppendedIterator struct {
	Event *ICanonicalTransactionChainQueueBatchAppended // Event containing the contract specifics and raw log

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
func (it *ICanonicalTransactionChainQueueBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICanonicalTransactionChainQueueBatchAppended)
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
		it.Event = new(ICanonicalTransactionChainQueueBatchAppended)
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
func (it *ICanonicalTransactionChainQueueBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICanonicalTransactionChainQueueBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICanonicalTransactionChainQueueBatchAppended represents a QueueBatchAppended event raised by the ICanonicalTransactionChain contract.
type ICanonicalTransactionChainQueueBatchAppended struct {
	StartingQueueIndex *big.Int
	NumQueueElements   *big.Int
	TotalElements      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterQueueBatchAppended is a free log retrieval operation binding the contract event 0x64d7f508348c70dea42d5302a393987e4abc20e45954ab3f9d320207751956f0.
//
// Solidity: event QueueBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) FilterQueueBatchAppended(opts *bind.FilterOpts) (*ICanonicalTransactionChainQueueBatchAppendedIterator, error) {

	logs, sub, err := _ICanonicalTransactionChain.contract.FilterLogs(opts, "QueueBatchAppended")
	if err != nil {
		return nil, err
	}
	return &ICanonicalTransactionChainQueueBatchAppendedIterator{contract: _ICanonicalTransactionChain.contract, event: "QueueBatchAppended", logs: logs, sub: sub}, nil
}

// WatchQueueBatchAppended is a free log subscription operation binding the contract event 0x64d7f508348c70dea42d5302a393987e4abc20e45954ab3f9d320207751956f0.
//
// Solidity: event QueueBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) WatchQueueBatchAppended(opts *bind.WatchOpts, sink chan<- *ICanonicalTransactionChainQueueBatchAppended) (event.Subscription, error) {

	logs, sub, err := _ICanonicalTransactionChain.contract.WatchLogs(opts, "QueueBatchAppended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICanonicalTransactionChainQueueBatchAppended)
				if err := _ICanonicalTransactionChain.contract.UnpackLog(event, "QueueBatchAppended", log); err != nil {
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

// ParseQueueBatchAppended is a log parse operation binding the contract event 0x64d7f508348c70dea42d5302a393987e4abc20e45954ab3f9d320207751956f0.
//
// Solidity: event QueueBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) ParseQueueBatchAppended(log types.Log) (*ICanonicalTransactionChainQueueBatchAppended, error) {
	event := new(ICanonicalTransactionChainQueueBatchAppended)
	if err := _ICanonicalTransactionChain.contract.UnpackLog(event, "QueueBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICanonicalTransactionChainSequencerBatchAppendedIterator is returned from FilterSequencerBatchAppended and is used to iterate over the raw logs and unpacked data for SequencerBatchAppended events raised by the ICanonicalTransactionChain contract.
type ICanonicalTransactionChainSequencerBatchAppendedIterator struct {
	Event *ICanonicalTransactionChainSequencerBatchAppended // Event containing the contract specifics and raw log

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
func (it *ICanonicalTransactionChainSequencerBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICanonicalTransactionChainSequencerBatchAppended)
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
		it.Event = new(ICanonicalTransactionChainSequencerBatchAppended)
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
func (it *ICanonicalTransactionChainSequencerBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICanonicalTransactionChainSequencerBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICanonicalTransactionChainSequencerBatchAppended represents a SequencerBatchAppended event raised by the ICanonicalTransactionChain contract.
type ICanonicalTransactionChainSequencerBatchAppended struct {
	StartingQueueIndex *big.Int
	NumQueueElements   *big.Int
	TotalElements      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchAppended is a free log retrieval operation binding the contract event 0x602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f899.
//
// Solidity: event SequencerBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) FilterSequencerBatchAppended(opts *bind.FilterOpts) (*ICanonicalTransactionChainSequencerBatchAppendedIterator, error) {

	logs, sub, err := _ICanonicalTransactionChain.contract.FilterLogs(opts, "SequencerBatchAppended")
	if err != nil {
		return nil, err
	}
	return &ICanonicalTransactionChainSequencerBatchAppendedIterator{contract: _ICanonicalTransactionChain.contract, event: "SequencerBatchAppended", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchAppended is a free log subscription operation binding the contract event 0x602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f899.
//
// Solidity: event SequencerBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) WatchSequencerBatchAppended(opts *bind.WatchOpts, sink chan<- *ICanonicalTransactionChainSequencerBatchAppended) (event.Subscription, error) {

	logs, sub, err := _ICanonicalTransactionChain.contract.WatchLogs(opts, "SequencerBatchAppended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICanonicalTransactionChainSequencerBatchAppended)
				if err := _ICanonicalTransactionChain.contract.UnpackLog(event, "SequencerBatchAppended", log); err != nil {
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

// ParseSequencerBatchAppended is a log parse operation binding the contract event 0x602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f899.
//
// Solidity: event SequencerBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) ParseSequencerBatchAppended(log types.Log) (*ICanonicalTransactionChainSequencerBatchAppended, error) {
	event := new(ICanonicalTransactionChainSequencerBatchAppended)
	if err := _ICanonicalTransactionChain.contract.UnpackLog(event, "SequencerBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICanonicalTransactionChainTransactionBatchAppendedIterator is returned from FilterTransactionBatchAppended and is used to iterate over the raw logs and unpacked data for TransactionBatchAppended events raised by the ICanonicalTransactionChain contract.
type ICanonicalTransactionChainTransactionBatchAppendedIterator struct {
	Event *ICanonicalTransactionChainTransactionBatchAppended // Event containing the contract specifics and raw log

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
func (it *ICanonicalTransactionChainTransactionBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICanonicalTransactionChainTransactionBatchAppended)
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
		it.Event = new(ICanonicalTransactionChainTransactionBatchAppended)
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
func (it *ICanonicalTransactionChainTransactionBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICanonicalTransactionChainTransactionBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICanonicalTransactionChainTransactionBatchAppended represents a TransactionBatchAppended event raised by the ICanonicalTransactionChain contract.
type ICanonicalTransactionChainTransactionBatchAppended struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	Signature         []byte
	ExtraData         []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTransactionBatchAppended is a free log retrieval operation binding the contract event 0xa47512905cf577d4cfae2efc3df461008ddb7234e91ce7f4eefcdb51e1077ccf.
//
// Solidity: event TransactionBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _signature, bytes _extraData)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) FilterTransactionBatchAppended(opts *bind.FilterOpts, _batchIndex []*big.Int) (*ICanonicalTransactionChainTransactionBatchAppendedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _ICanonicalTransactionChain.contract.FilterLogs(opts, "TransactionBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &ICanonicalTransactionChainTransactionBatchAppendedIterator{contract: _ICanonicalTransactionChain.contract, event: "TransactionBatchAppended", logs: logs, sub: sub}, nil
}

// WatchTransactionBatchAppended is a free log subscription operation binding the contract event 0xa47512905cf577d4cfae2efc3df461008ddb7234e91ce7f4eefcdb51e1077ccf.
//
// Solidity: event TransactionBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _signature, bytes _extraData)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) WatchTransactionBatchAppended(opts *bind.WatchOpts, sink chan<- *ICanonicalTransactionChainTransactionBatchAppended, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _ICanonicalTransactionChain.contract.WatchLogs(opts, "TransactionBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICanonicalTransactionChainTransactionBatchAppended)
				if err := _ICanonicalTransactionChain.contract.UnpackLog(event, "TransactionBatchAppended", log); err != nil {
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

// ParseTransactionBatchAppended is a log parse operation binding the contract event 0xa47512905cf577d4cfae2efc3df461008ddb7234e91ce7f4eefcdb51e1077ccf.
//
// Solidity: event TransactionBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _signature, bytes _extraData)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) ParseTransactionBatchAppended(log types.Log) (*ICanonicalTransactionChainTransactionBatchAppended, error) {
	event := new(ICanonicalTransactionChainTransactionBatchAppended)
	if err := _ICanonicalTransactionChain.contract.UnpackLog(event, "TransactionBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICanonicalTransactionChainTransactionEnqueuedIterator is returned from FilterTransactionEnqueued and is used to iterate over the raw logs and unpacked data for TransactionEnqueued events raised by the ICanonicalTransactionChain contract.
type ICanonicalTransactionChainTransactionEnqueuedIterator struct {
	Event *ICanonicalTransactionChainTransactionEnqueued // Event containing the contract specifics and raw log

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
func (it *ICanonicalTransactionChainTransactionEnqueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICanonicalTransactionChainTransactionEnqueued)
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
		it.Event = new(ICanonicalTransactionChainTransactionEnqueued)
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
func (it *ICanonicalTransactionChainTransactionEnqueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICanonicalTransactionChainTransactionEnqueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICanonicalTransactionChainTransactionEnqueued represents a TransactionEnqueued event raised by the ICanonicalTransactionChain contract.
type ICanonicalTransactionChainTransactionEnqueued struct {
	L1TxOrigin common.Address
	Target     common.Address
	GasLimit   *big.Int
	Data       []byte
	QueueIndex *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionEnqueued is a free log retrieval operation binding the contract event 0x4b388aecf9fa6cc92253704e5975a6129a4f735bdbd99567df4ed0094ee4ceb5.
//
// Solidity: event TransactionEnqueued(address indexed _l1TxOrigin, address indexed _target, uint256 _gasLimit, bytes _data, uint256 indexed _queueIndex, uint256 _timestamp)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) FilterTransactionEnqueued(opts *bind.FilterOpts, _l1TxOrigin []common.Address, _target []common.Address, _queueIndex []*big.Int) (*ICanonicalTransactionChainTransactionEnqueuedIterator, error) {

	var _l1TxOriginRule []interface{}
	for _, _l1TxOriginItem := range _l1TxOrigin {
		_l1TxOriginRule = append(_l1TxOriginRule, _l1TxOriginItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	var _queueIndexRule []interface{}
	for _, _queueIndexItem := range _queueIndex {
		_queueIndexRule = append(_queueIndexRule, _queueIndexItem)
	}

	logs, sub, err := _ICanonicalTransactionChain.contract.FilterLogs(opts, "TransactionEnqueued", _l1TxOriginRule, _targetRule, _queueIndexRule)
	if err != nil {
		return nil, err
	}
	return &ICanonicalTransactionChainTransactionEnqueuedIterator{contract: _ICanonicalTransactionChain.contract, event: "TransactionEnqueued", logs: logs, sub: sub}, nil
}

// WatchTransactionEnqueued is a free log subscription operation binding the contract event 0x4b388aecf9fa6cc92253704e5975a6129a4f735bdbd99567df4ed0094ee4ceb5.
//
// Solidity: event TransactionEnqueued(address indexed _l1TxOrigin, address indexed _target, uint256 _gasLimit, bytes _data, uint256 indexed _queueIndex, uint256 _timestamp)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) WatchTransactionEnqueued(opts *bind.WatchOpts, sink chan<- *ICanonicalTransactionChainTransactionEnqueued, _l1TxOrigin []common.Address, _target []common.Address, _queueIndex []*big.Int) (event.Subscription, error) {

	var _l1TxOriginRule []interface{}
	for _, _l1TxOriginItem := range _l1TxOrigin {
		_l1TxOriginRule = append(_l1TxOriginRule, _l1TxOriginItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	var _queueIndexRule []interface{}
	for _, _queueIndexItem := range _queueIndex {
		_queueIndexRule = append(_queueIndexRule, _queueIndexItem)
	}

	logs, sub, err := _ICanonicalTransactionChain.contract.WatchLogs(opts, "TransactionEnqueued", _l1TxOriginRule, _targetRule, _queueIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICanonicalTransactionChainTransactionEnqueued)
				if err := _ICanonicalTransactionChain.contract.UnpackLog(event, "TransactionEnqueued", log); err != nil {
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

// ParseTransactionEnqueued is a log parse operation binding the contract event 0x4b388aecf9fa6cc92253704e5975a6129a4f735bdbd99567df4ed0094ee4ceb5.
//
// Solidity: event TransactionEnqueued(address indexed _l1TxOrigin, address indexed _target, uint256 _gasLimit, bytes _data, uint256 indexed _queueIndex, uint256 _timestamp)
func (_ICanonicalTransactionChain *ICanonicalTransactionChainFilterer) ParseTransactionEnqueued(log types.Log) (*ICanonicalTransactionChainTransactionEnqueued, error) {
	event := new(ICanonicalTransactionChainTransactionEnqueued)
	if err := _ICanonicalTransactionChain.contract.UnpackLog(event, "TransactionEnqueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

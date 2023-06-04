// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TssRewardContract

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

// TssRewardContractMetaData contains all meta data concerning the TssRewardContract contract.
var TssRewardContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deadAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sendAmountPerYear\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_bvmGasPriceOracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2CrossDomainMessenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sccAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastBatchTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tssMembers\",\"type\":\"address[]\"}],\"name\":\"DistributeTssReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockStartHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"length\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tssMembers\",\"type\":\"address[]\"}],\"name\":\"DistributeTssRewardByBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bestBlockID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bvmGasPriceOracleAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockStartHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_length\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_batchTime\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_tssMembers\",\"type\":\"address[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dust\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBatchTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"querySendAmountPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sccAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendAmountPerYear\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sendAmountPerYear\",\"type\":\"uint256\"}],\"name\":\"setSendAmountPerYear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"updateReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDust\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TssRewardContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TssRewardContractMetaData.ABI instead.
var TssRewardContractABI = TssRewardContractMetaData.ABI

// TssRewardContract is an auto generated Go binding around an Ethereum contract.
type TssRewardContract struct {
	TssRewardContractCaller     // Read-only binding to the contract
	TssRewardContractTransactor // Write-only binding to the contract
	TssRewardContractFilterer   // Log filterer for contract events
}

// TssRewardContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TssRewardContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssRewardContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TssRewardContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssRewardContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TssRewardContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssRewardContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TssRewardContractSession struct {
	Contract     *TssRewardContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TssRewardContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TssRewardContractCallerSession struct {
	Contract *TssRewardContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TssRewardContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TssRewardContractTransactorSession struct {
	Contract     *TssRewardContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TssRewardContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TssRewardContractRaw struct {
	Contract *TssRewardContract // Generic contract binding to access the raw methods on
}

// TssRewardContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TssRewardContractCallerRaw struct {
	Contract *TssRewardContractCaller // Generic read-only contract binding to access the raw methods on
}

// TssRewardContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TssRewardContractTransactorRaw struct {
	Contract *TssRewardContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTssRewardContract creates a new instance of TssRewardContract, bound to a specific deployed contract.
func NewTssRewardContract(address common.Address, backend bind.ContractBackend) (*TssRewardContract, error) {
	contract, err := bindTssRewardContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TssRewardContract{TssRewardContractCaller: TssRewardContractCaller{contract: contract}, TssRewardContractTransactor: TssRewardContractTransactor{contract: contract}, TssRewardContractFilterer: TssRewardContractFilterer{contract: contract}}, nil
}

// NewTssRewardContractCaller creates a new read-only instance of TssRewardContract, bound to a specific deployed contract.
func NewTssRewardContractCaller(address common.Address, caller bind.ContractCaller) (*TssRewardContractCaller, error) {
	contract, err := bindTssRewardContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TssRewardContractCaller{contract: contract}, nil
}

// NewTssRewardContractTransactor creates a new write-only instance of TssRewardContract, bound to a specific deployed contract.
func NewTssRewardContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TssRewardContractTransactor, error) {
	contract, err := bindTssRewardContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TssRewardContractTransactor{contract: contract}, nil
}

// NewTssRewardContractFilterer creates a new log filterer instance of TssRewardContract, bound to a specific deployed contract.
func NewTssRewardContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TssRewardContractFilterer, error) {
	contract, err := bindTssRewardContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TssRewardContractFilterer{contract: contract}, nil
}

// bindTssRewardContract binds a generic wrapper to an already deployed contract.
func bindTssRewardContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TssRewardContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TssRewardContract *TssRewardContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TssRewardContract.Contract.TssRewardContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TssRewardContract *TssRewardContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssRewardContract.Contract.TssRewardContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TssRewardContract *TssRewardContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TssRewardContract.Contract.TssRewardContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TssRewardContract *TssRewardContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TssRewardContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TssRewardContract *TssRewardContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssRewardContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TssRewardContract *TssRewardContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TssRewardContract.Contract.contract.Transact(opts, method, params...)
}

// BestBlockID is a free data retrieval call binding the contract method 0x19d509a1.
//
// Solidity: function bestBlockID() view returns(uint256)
func (_TssRewardContract *TssRewardContractCaller) BestBlockID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "bestBlockID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BestBlockID is a free data retrieval call binding the contract method 0x19d509a1.
//
// Solidity: function bestBlockID() view returns(uint256)
func (_TssRewardContract *TssRewardContractSession) BestBlockID() (*big.Int, error) {
	return _TssRewardContract.Contract.BestBlockID(&_TssRewardContract.CallOpts)
}

// BestBlockID is a free data retrieval call binding the contract method 0x19d509a1.
//
// Solidity: function bestBlockID() view returns(uint256)
func (_TssRewardContract *TssRewardContractCallerSession) BestBlockID() (*big.Int, error) {
	return _TssRewardContract.Contract.BestBlockID(&_TssRewardContract.CallOpts)
}

// BvmGasPriceOracleAddress is a free data retrieval call binding the contract method 0x110b7eb0.
//
// Solidity: function bvmGasPriceOracleAddress() view returns(address)
func (_TssRewardContract *TssRewardContractCaller) BvmGasPriceOracleAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "bvmGasPriceOracleAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BvmGasPriceOracleAddress is a free data retrieval call binding the contract method 0x110b7eb0.
//
// Solidity: function bvmGasPriceOracleAddress() view returns(address)
func (_TssRewardContract *TssRewardContractSession) BvmGasPriceOracleAddress() (common.Address, error) {
	return _TssRewardContract.Contract.BvmGasPriceOracleAddress(&_TssRewardContract.CallOpts)
}

// BvmGasPriceOracleAddress is a free data retrieval call binding the contract method 0x110b7eb0.
//
// Solidity: function bvmGasPriceOracleAddress() view returns(address)
func (_TssRewardContract *TssRewardContractCallerSession) BvmGasPriceOracleAddress() (common.Address, error) {
	return _TssRewardContract.Contract.BvmGasPriceOracleAddress(&_TssRewardContract.CallOpts)
}

// DeadAddress is a free data retrieval call binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() view returns(address)
func (_TssRewardContract *TssRewardContractCaller) DeadAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "deadAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeadAddress is a free data retrieval call binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() view returns(address)
func (_TssRewardContract *TssRewardContractSession) DeadAddress() (common.Address, error) {
	return _TssRewardContract.Contract.DeadAddress(&_TssRewardContract.CallOpts)
}

// DeadAddress is a free data retrieval call binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() view returns(address)
func (_TssRewardContract *TssRewardContractCallerSession) DeadAddress() (common.Address, error) {
	return _TssRewardContract.Contract.DeadAddress(&_TssRewardContract.CallOpts)
}

// Dust is a free data retrieval call binding the contract method 0xfad9aba3.
//
// Solidity: function dust() view returns(uint256)
func (_TssRewardContract *TssRewardContractCaller) Dust(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "dust")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dust is a free data retrieval call binding the contract method 0xfad9aba3.
//
// Solidity: function dust() view returns(uint256)
func (_TssRewardContract *TssRewardContractSession) Dust() (*big.Int, error) {
	return _TssRewardContract.Contract.Dust(&_TssRewardContract.CallOpts)
}

// Dust is a free data retrieval call binding the contract method 0xfad9aba3.
//
// Solidity: function dust() view returns(uint256)
func (_TssRewardContract *TssRewardContractCallerSession) Dust() (*big.Int, error) {
	return _TssRewardContract.Contract.Dust(&_TssRewardContract.CallOpts)
}

// LastBatchTime is a free data retrieval call binding the contract method 0xe5efd585.
//
// Solidity: function lastBatchTime() view returns(uint256)
func (_TssRewardContract *TssRewardContractCaller) LastBatchTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "lastBatchTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBatchTime is a free data retrieval call binding the contract method 0xe5efd585.
//
// Solidity: function lastBatchTime() view returns(uint256)
func (_TssRewardContract *TssRewardContractSession) LastBatchTime() (*big.Int, error) {
	return _TssRewardContract.Contract.LastBatchTime(&_TssRewardContract.CallOpts)
}

// LastBatchTime is a free data retrieval call binding the contract method 0xe5efd585.
//
// Solidity: function lastBatchTime() view returns(uint256)
func (_TssRewardContract *TssRewardContractCallerSession) LastBatchTime() (*big.Int, error) {
	return _TssRewardContract.Contract.LastBatchTime(&_TssRewardContract.CallOpts)
}

// Ledger is a free data retrieval call binding the contract method 0x10a7fd7b.
//
// Solidity: function ledger(uint256 ) view returns(uint256)
func (_TssRewardContract *TssRewardContractCaller) Ledger(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "ledger", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ledger is a free data retrieval call binding the contract method 0x10a7fd7b.
//
// Solidity: function ledger(uint256 ) view returns(uint256)
func (_TssRewardContract *TssRewardContractSession) Ledger(arg0 *big.Int) (*big.Int, error) {
	return _TssRewardContract.Contract.Ledger(&_TssRewardContract.CallOpts, arg0)
}

// Ledger is a free data retrieval call binding the contract method 0x10a7fd7b.
//
// Solidity: function ledger(uint256 ) view returns(uint256)
func (_TssRewardContract *TssRewardContractCallerSession) Ledger(arg0 *big.Int) (*big.Int, error) {
	return _TssRewardContract.Contract.Ledger(&_TssRewardContract.CallOpts, arg0)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_TssRewardContract *TssRewardContractCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_TssRewardContract *TssRewardContractSession) Messenger() (common.Address, error) {
	return _TssRewardContract.Contract.Messenger(&_TssRewardContract.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_TssRewardContract *TssRewardContractCallerSession) Messenger() (common.Address, error) {
	return _TssRewardContract.Contract.Messenger(&_TssRewardContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssRewardContract *TssRewardContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssRewardContract *TssRewardContractSession) Owner() (common.Address, error) {
	return _TssRewardContract.Contract.Owner(&_TssRewardContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssRewardContract *TssRewardContractCallerSession) Owner() (common.Address, error) {
	return _TssRewardContract.Contract.Owner(&_TssRewardContract.CallOpts)
}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_TssRewardContract *TssRewardContractCaller) QueryReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "queryReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_TssRewardContract *TssRewardContractSession) QueryReward() (*big.Int, error) {
	return _TssRewardContract.Contract.QueryReward(&_TssRewardContract.CallOpts)
}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_TssRewardContract *TssRewardContractCallerSession) QueryReward() (*big.Int, error) {
	return _TssRewardContract.Contract.QueryReward(&_TssRewardContract.CallOpts)
}

// QuerySendAmountPerSecond is a free data retrieval call binding the contract method 0x15c6f166.
//
// Solidity: function querySendAmountPerSecond() view returns(uint256)
func (_TssRewardContract *TssRewardContractCaller) QuerySendAmountPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "querySendAmountPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySendAmountPerSecond is a free data retrieval call binding the contract method 0x15c6f166.
//
// Solidity: function querySendAmountPerSecond() view returns(uint256)
func (_TssRewardContract *TssRewardContractSession) QuerySendAmountPerSecond() (*big.Int, error) {
	return _TssRewardContract.Contract.QuerySendAmountPerSecond(&_TssRewardContract.CallOpts)
}

// QuerySendAmountPerSecond is a free data retrieval call binding the contract method 0x15c6f166.
//
// Solidity: function querySendAmountPerSecond() view returns(uint256)
func (_TssRewardContract *TssRewardContractCallerSession) QuerySendAmountPerSecond() (*big.Int, error) {
	return _TssRewardContract.Contract.QuerySendAmountPerSecond(&_TssRewardContract.CallOpts)
}

// SccAddress is a free data retrieval call binding the contract method 0xea01cd36.
//
// Solidity: function sccAddress() view returns(address)
func (_TssRewardContract *TssRewardContractCaller) SccAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "sccAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SccAddress is a free data retrieval call binding the contract method 0xea01cd36.
//
// Solidity: function sccAddress() view returns(address)
func (_TssRewardContract *TssRewardContractSession) SccAddress() (common.Address, error) {
	return _TssRewardContract.Contract.SccAddress(&_TssRewardContract.CallOpts)
}

// SccAddress is a free data retrieval call binding the contract method 0xea01cd36.
//
// Solidity: function sccAddress() view returns(address)
func (_TssRewardContract *TssRewardContractCallerSession) SccAddress() (common.Address, error) {
	return _TssRewardContract.Contract.SccAddress(&_TssRewardContract.CallOpts)
}

// SendAmountPerYear is a free data retrieval call binding the contract method 0xd8111a57.
//
// Solidity: function sendAmountPerYear() view returns(uint256)
func (_TssRewardContract *TssRewardContractCaller) SendAmountPerYear(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "sendAmountPerYear")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SendAmountPerYear is a free data retrieval call binding the contract method 0xd8111a57.
//
// Solidity: function sendAmountPerYear() view returns(uint256)
func (_TssRewardContract *TssRewardContractSession) SendAmountPerYear() (*big.Int, error) {
	return _TssRewardContract.Contract.SendAmountPerYear(&_TssRewardContract.CallOpts)
}

// SendAmountPerYear is a free data retrieval call binding the contract method 0xd8111a57.
//
// Solidity: function sendAmountPerYear() view returns(uint256)
func (_TssRewardContract *TssRewardContractCallerSession) SendAmountPerYear() (*big.Int, error) {
	return _TssRewardContract.Contract.SendAmountPerYear(&_TssRewardContract.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TssRewardContract *TssRewardContractCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TssRewardContract.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TssRewardContract *TssRewardContractSession) TotalAmount() (*big.Int, error) {
	return _TssRewardContract.Contract.TotalAmount(&_TssRewardContract.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TssRewardContract *TssRewardContractCallerSession) TotalAmount() (*big.Int, error) {
	return _TssRewardContract.Contract.TotalAmount(&_TssRewardContract.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0fae75d9.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] _tssMembers) returns()
func (_TssRewardContract *TssRewardContractTransactor) ClaimReward(opts *bind.TransactOpts, _blockStartHeight *big.Int, _length uint32, _batchTime *big.Int, _tssMembers []common.Address) (*types.Transaction, error) {
	return _TssRewardContract.contract.Transact(opts, "claimReward", _blockStartHeight, _length, _batchTime, _tssMembers)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0fae75d9.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] _tssMembers) returns()
func (_TssRewardContract *TssRewardContractSession) ClaimReward(_blockStartHeight *big.Int, _length uint32, _batchTime *big.Int, _tssMembers []common.Address) (*types.Transaction, error) {
	return _TssRewardContract.Contract.ClaimReward(&_TssRewardContract.TransactOpts, _blockStartHeight, _length, _batchTime, _tssMembers)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0fae75d9.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] _tssMembers) returns()
func (_TssRewardContract *TssRewardContractTransactorSession) ClaimReward(_blockStartHeight *big.Int, _length uint32, _batchTime *big.Int, _tssMembers []common.Address) (*types.Transaction, error) {
	return _TssRewardContract.Contract.ClaimReward(&_TssRewardContract.TransactOpts, _blockStartHeight, _length, _batchTime, _tssMembers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TssRewardContract *TssRewardContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssRewardContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TssRewardContract *TssRewardContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _TssRewardContract.Contract.RenounceOwnership(&_TssRewardContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TssRewardContract *TssRewardContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TssRewardContract.Contract.RenounceOwnership(&_TssRewardContract.TransactOpts)
}

// SetSendAmountPerYear is a paid mutator transaction binding the contract method 0x3b52c31e.
//
// Solidity: function setSendAmountPerYear(uint256 _sendAmountPerYear) returns()
func (_TssRewardContract *TssRewardContractTransactor) SetSendAmountPerYear(opts *bind.TransactOpts, _sendAmountPerYear *big.Int) (*types.Transaction, error) {
	return _TssRewardContract.contract.Transact(opts, "setSendAmountPerYear", _sendAmountPerYear)
}

// SetSendAmountPerYear is a paid mutator transaction binding the contract method 0x3b52c31e.
//
// Solidity: function setSendAmountPerYear(uint256 _sendAmountPerYear) returns()
func (_TssRewardContract *TssRewardContractSession) SetSendAmountPerYear(_sendAmountPerYear *big.Int) (*types.Transaction, error) {
	return _TssRewardContract.Contract.SetSendAmountPerYear(&_TssRewardContract.TransactOpts, _sendAmountPerYear)
}

// SetSendAmountPerYear is a paid mutator transaction binding the contract method 0x3b52c31e.
//
// Solidity: function setSendAmountPerYear(uint256 _sendAmountPerYear) returns()
func (_TssRewardContract *TssRewardContractTransactorSession) SetSendAmountPerYear(_sendAmountPerYear *big.Int) (*types.Transaction, error) {
	return _TssRewardContract.Contract.SetSendAmountPerYear(&_TssRewardContract.TransactOpts, _sendAmountPerYear)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TssRewardContract *TssRewardContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TssRewardContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TssRewardContract *TssRewardContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TssRewardContract.Contract.TransferOwnership(&_TssRewardContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TssRewardContract *TssRewardContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TssRewardContract.Contract.TransferOwnership(&_TssRewardContract.TransactOpts, newOwner)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_TssRewardContract *TssRewardContractTransactor) UpdateReward(opts *bind.TransactOpts, _blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TssRewardContract.contract.Transact(opts, "updateReward", _blockID, _amount)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_TssRewardContract *TssRewardContractSession) UpdateReward(_blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TssRewardContract.Contract.UpdateReward(&_TssRewardContract.TransactOpts, _blockID, _amount)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_TssRewardContract *TssRewardContractTransactorSession) UpdateReward(_blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TssRewardContract.Contract.UpdateReward(&_TssRewardContract.TransactOpts, _blockID, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TssRewardContract *TssRewardContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssRewardContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TssRewardContract *TssRewardContractSession) Withdraw() (*types.Transaction, error) {
	return _TssRewardContract.Contract.Withdraw(&_TssRewardContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TssRewardContract *TssRewardContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _TssRewardContract.Contract.Withdraw(&_TssRewardContract.TransactOpts)
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_TssRewardContract *TssRewardContractTransactor) WithdrawDust(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssRewardContract.contract.Transact(opts, "withdrawDust")
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_TssRewardContract *TssRewardContractSession) WithdrawDust() (*types.Transaction, error) {
	return _TssRewardContract.Contract.WithdrawDust(&_TssRewardContract.TransactOpts)
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_TssRewardContract *TssRewardContractTransactorSession) WithdrawDust() (*types.Transaction, error) {
	return _TssRewardContract.Contract.WithdrawDust(&_TssRewardContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TssRewardContract *TssRewardContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssRewardContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TssRewardContract *TssRewardContractSession) Receive() (*types.Transaction, error) {
	return _TssRewardContract.Contract.Receive(&_TssRewardContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TssRewardContract *TssRewardContractTransactorSession) Receive() (*types.Transaction, error) {
	return _TssRewardContract.Contract.Receive(&_TssRewardContract.TransactOpts)
}

// TssRewardContractDistributeTssRewardIterator is returned from FilterDistributeTssReward and is used to iterate over the raw logs and unpacked data for DistributeTssReward events raised by the TssRewardContract contract.
type TssRewardContractDistributeTssRewardIterator struct {
	Event *TssRewardContractDistributeTssReward // Event containing the contract specifics and raw log

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
func (it *TssRewardContractDistributeTssRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssRewardContractDistributeTssReward)
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
		it.Event = new(TssRewardContractDistributeTssReward)
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
func (it *TssRewardContractDistributeTssRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssRewardContractDistributeTssRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssRewardContractDistributeTssReward represents a DistributeTssReward event raised by the TssRewardContract contract.
type TssRewardContractDistributeTssReward struct {
	LastBatchTime *big.Int
	BatchTime     *big.Int
	Amount        *big.Int
	TssMembers    []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDistributeTssReward is a free log retrieval operation binding the contract event 0xf533ef50019763ee9d95ad46e28350b533c11edd472ae7be93e8fae83c1b6d99.
//
// Solidity: event DistributeTssReward(uint256 lastBatchTime, uint256 batchTime, uint256 amount, address[] tssMembers)
func (_TssRewardContract *TssRewardContractFilterer) FilterDistributeTssReward(opts *bind.FilterOpts) (*TssRewardContractDistributeTssRewardIterator, error) {

	logs, sub, err := _TssRewardContract.contract.FilterLogs(opts, "DistributeTssReward")
	if err != nil {
		return nil, err
	}
	return &TssRewardContractDistributeTssRewardIterator{contract: _TssRewardContract.contract, event: "DistributeTssReward", logs: logs, sub: sub}, nil
}

// WatchDistributeTssReward is a free log subscription operation binding the contract event 0xf533ef50019763ee9d95ad46e28350b533c11edd472ae7be93e8fae83c1b6d99.
//
// Solidity: event DistributeTssReward(uint256 lastBatchTime, uint256 batchTime, uint256 amount, address[] tssMembers)
func (_TssRewardContract *TssRewardContractFilterer) WatchDistributeTssReward(opts *bind.WatchOpts, sink chan<- *TssRewardContractDistributeTssReward) (event.Subscription, error) {

	logs, sub, err := _TssRewardContract.contract.WatchLogs(opts, "DistributeTssReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssRewardContractDistributeTssReward)
				if err := _TssRewardContract.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
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
// Solidity: event DistributeTssReward(uint256 lastBatchTime, uint256 batchTime, uint256 amount, address[] tssMembers)
func (_TssRewardContract *TssRewardContractFilterer) ParseDistributeTssReward(log types.Log) (*TssRewardContractDistributeTssReward, error) {
	event := new(TssRewardContractDistributeTssReward)
	if err := _TssRewardContract.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TssRewardContractDistributeTssRewardByBlockIterator is returned from FilterDistributeTssRewardByBlock and is used to iterate over the raw logs and unpacked data for DistributeTssRewardByBlock events raised by the TssRewardContract contract.
type TssRewardContractDistributeTssRewardByBlockIterator struct {
	Event *TssRewardContractDistributeTssRewardByBlock // Event containing the contract specifics and raw log

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
func (it *TssRewardContractDistributeTssRewardByBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssRewardContractDistributeTssRewardByBlock)
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
		it.Event = new(TssRewardContractDistributeTssRewardByBlock)
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
func (it *TssRewardContractDistributeTssRewardByBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssRewardContractDistributeTssRewardByBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssRewardContractDistributeTssRewardByBlock represents a DistributeTssRewardByBlock event raised by the TssRewardContract contract.
type TssRewardContractDistributeTssRewardByBlock struct {
	BlockStartHeight *big.Int
	Length           uint32
	Amount           *big.Int
	TssMembers       []common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributeTssRewardByBlock is a free log retrieval operation binding the contract event 0x2dae6f3d42a2c50d6baa3ea3f2423a9e1ff0ba26875f8ba6ba25c40df98009fe.
//
// Solidity: event DistributeTssRewardByBlock(uint256 blockStartHeight, uint32 length, uint256 amount, address[] tssMembers)
func (_TssRewardContract *TssRewardContractFilterer) FilterDistributeTssRewardByBlock(opts *bind.FilterOpts) (*TssRewardContractDistributeTssRewardByBlockIterator, error) {

	logs, sub, err := _TssRewardContract.contract.FilterLogs(opts, "DistributeTssRewardByBlock")
	if err != nil {
		return nil, err
	}
	return &TssRewardContractDistributeTssRewardByBlockIterator{contract: _TssRewardContract.contract, event: "DistributeTssRewardByBlock", logs: logs, sub: sub}, nil
}

// WatchDistributeTssRewardByBlock is a free log subscription operation binding the contract event 0x2dae6f3d42a2c50d6baa3ea3f2423a9e1ff0ba26875f8ba6ba25c40df98009fe.
//
// Solidity: event DistributeTssRewardByBlock(uint256 blockStartHeight, uint32 length, uint256 amount, address[] tssMembers)
func (_TssRewardContract *TssRewardContractFilterer) WatchDistributeTssRewardByBlock(opts *bind.WatchOpts, sink chan<- *TssRewardContractDistributeTssRewardByBlock) (event.Subscription, error) {

	logs, sub, err := _TssRewardContract.contract.WatchLogs(opts, "DistributeTssRewardByBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssRewardContractDistributeTssRewardByBlock)
				if err := _TssRewardContract.contract.UnpackLog(event, "DistributeTssRewardByBlock", log); err != nil {
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

// ParseDistributeTssRewardByBlock is a log parse operation binding the contract event 0x2dae6f3d42a2c50d6baa3ea3f2423a9e1ff0ba26875f8ba6ba25c40df98009fe.
//
// Solidity: event DistributeTssRewardByBlock(uint256 blockStartHeight, uint32 length, uint256 amount, address[] tssMembers)
func (_TssRewardContract *TssRewardContractFilterer) ParseDistributeTssRewardByBlock(log types.Log) (*TssRewardContractDistributeTssRewardByBlock, error) {
	event := new(TssRewardContractDistributeTssRewardByBlock)
	if err := _TssRewardContract.contract.UnpackLog(event, "DistributeTssRewardByBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TssRewardContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TssRewardContract contract.
type TssRewardContractOwnershipTransferredIterator struct {
	Event *TssRewardContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TssRewardContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssRewardContractOwnershipTransferred)
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
		it.Event = new(TssRewardContractOwnershipTransferred)
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
func (it *TssRewardContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssRewardContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssRewardContractOwnershipTransferred represents a OwnershipTransferred event raised by the TssRewardContract contract.
type TssRewardContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TssRewardContract *TssRewardContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TssRewardContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TssRewardContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TssRewardContractOwnershipTransferredIterator{contract: _TssRewardContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TssRewardContract *TssRewardContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TssRewardContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TssRewardContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssRewardContractOwnershipTransferred)
				if err := _TssRewardContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TssRewardContract *TssRewardContractFilterer) ParseOwnershipTransferred(log types.Log) (*TssRewardContractOwnershipTransferred, error) {
	event := new(TssRewardContractOwnershipTransferred)
	if err := _TssRewardContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

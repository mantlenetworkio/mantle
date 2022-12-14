// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IStakingSlashing

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

// IStakingSlashingDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakingSlashingDepositInfo struct {
	Pledgor common.Address
	PubKey  []byte
	Amount  *big.Int
}

// IStakingSlashingMetaData contains all meta data concerning the IStakingSlashing contract.
var IStakingSlashingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"batchGetDeposits\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pledgor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingSlashing.DepositInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearQuitRequestList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getDeposits\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pledgor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingSlashing.DepositInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuitRequestList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getSlashRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlashingParams\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isJailed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quitRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"name\":\"setSlashingParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"slashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"staking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unJail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IStakingSlashingABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingSlashingMetaData.ABI instead.
var IStakingSlashingABI = IStakingSlashingMetaData.ABI

// IStakingSlashing is an auto generated Go binding around an Ethereum contract.
type IStakingSlashing struct {
	IStakingSlashingCaller     // Read-only binding to the contract
	IStakingSlashingTransactor // Write-only binding to the contract
	IStakingSlashingFilterer   // Log filterer for contract events
}

// IStakingSlashingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingSlashingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingSlashingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingSlashingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingSlashingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingSlashingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingSlashingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingSlashingSession struct {
	Contract     *IStakingSlashing // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakingSlashingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingSlashingCallerSession struct {
	Contract *IStakingSlashingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IStakingSlashingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingSlashingTransactorSession struct {
	Contract     *IStakingSlashingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IStakingSlashingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingSlashingRaw struct {
	Contract *IStakingSlashing // Generic contract binding to access the raw methods on
}

// IStakingSlashingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingSlashingCallerRaw struct {
	Contract *IStakingSlashingCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingSlashingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingSlashingTransactorRaw struct {
	Contract *IStakingSlashingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakingSlashing creates a new instance of IStakingSlashing, bound to a specific deployed contract.
func NewIStakingSlashing(address common.Address, backend bind.ContractBackend) (*IStakingSlashing, error) {
	contract, err := bindIStakingSlashing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakingSlashing{IStakingSlashingCaller: IStakingSlashingCaller{contract: contract}, IStakingSlashingTransactor: IStakingSlashingTransactor{contract: contract}, IStakingSlashingFilterer: IStakingSlashingFilterer{contract: contract}}, nil
}

// NewIStakingSlashingCaller creates a new read-only instance of IStakingSlashing, bound to a specific deployed contract.
func NewIStakingSlashingCaller(address common.Address, caller bind.ContractCaller) (*IStakingSlashingCaller, error) {
	contract, err := bindIStakingSlashing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingSlashingCaller{contract: contract}, nil
}

// NewIStakingSlashingTransactor creates a new write-only instance of IStakingSlashing, bound to a specific deployed contract.
func NewIStakingSlashingTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingSlashingTransactor, error) {
	contract, err := bindIStakingSlashing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingSlashingTransactor{contract: contract}, nil
}

// NewIStakingSlashingFilterer creates a new log filterer instance of IStakingSlashing, bound to a specific deployed contract.
func NewIStakingSlashingFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingSlashingFilterer, error) {
	contract, err := bindIStakingSlashing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingSlashingFilterer{contract: contract}, nil
}

// bindIStakingSlashing binds a generic wrapper to an already deployed contract.
func bindIStakingSlashing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStakingSlashingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingSlashing *IStakingSlashingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingSlashing.Contract.IStakingSlashingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingSlashing *IStakingSlashingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.IStakingSlashingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingSlashing *IStakingSlashingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.IStakingSlashingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingSlashing *IStakingSlashingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingSlashing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingSlashing *IStakingSlashingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingSlashing *IStakingSlashingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.contract.Transact(opts, method, params...)
}

// BatchGetDeposits is a free data retrieval call binding the contract method 0x793505c8.
//
// Solidity: function batchGetDeposits(address[] ) view returns((address,bytes,uint256)[])
func (_IStakingSlashing *IStakingSlashingCaller) BatchGetDeposits(opts *bind.CallOpts, arg0 []common.Address) ([]IStakingSlashingDepositInfo, error) {
	var out []interface{}
	err := _IStakingSlashing.contract.Call(opts, &out, "batchGetDeposits", arg0)

	if err != nil {
		return *new([]IStakingSlashingDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingSlashingDepositInfo)).(*[]IStakingSlashingDepositInfo)

	return out0, err

}

// BatchGetDeposits is a free data retrieval call binding the contract method 0x793505c8.
//
// Solidity: function batchGetDeposits(address[] ) view returns((address,bytes,uint256)[])
func (_IStakingSlashing *IStakingSlashingSession) BatchGetDeposits(arg0 []common.Address) ([]IStakingSlashingDepositInfo, error) {
	return _IStakingSlashing.Contract.BatchGetDeposits(&_IStakingSlashing.CallOpts, arg0)
}

// BatchGetDeposits is a free data retrieval call binding the contract method 0x793505c8.
//
// Solidity: function batchGetDeposits(address[] ) view returns((address,bytes,uint256)[])
func (_IStakingSlashing *IStakingSlashingCallerSession) BatchGetDeposits(arg0 []common.Address) ([]IStakingSlashingDepositInfo, error) {
	return _IStakingSlashing.Contract.BatchGetDeposits(&_IStakingSlashing.CallOpts, arg0)
}

// GetQuitRequestList is a free data retrieval call binding the contract method 0x490ab9d6.
//
// Solidity: function getQuitRequestList() view returns(address[])
func (_IStakingSlashing *IStakingSlashingCaller) GetQuitRequestList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IStakingSlashing.contract.Call(opts, &out, "getQuitRequestList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetQuitRequestList is a free data retrieval call binding the contract method 0x490ab9d6.
//
// Solidity: function getQuitRequestList() view returns(address[])
func (_IStakingSlashing *IStakingSlashingSession) GetQuitRequestList() ([]common.Address, error) {
	return _IStakingSlashing.Contract.GetQuitRequestList(&_IStakingSlashing.CallOpts)
}

// GetQuitRequestList is a free data retrieval call binding the contract method 0x490ab9d6.
//
// Solidity: function getQuitRequestList() view returns(address[])
func (_IStakingSlashing *IStakingSlashingCallerSession) GetQuitRequestList() ([]common.Address, error) {
	return _IStakingSlashing.Contract.GetQuitRequestList(&_IStakingSlashing.CallOpts)
}

// GetSlashRecord is a free data retrieval call binding the contract method 0x829673ef.
//
// Solidity: function getSlashRecord(uint256 , address ) view returns(bool)
func (_IStakingSlashing *IStakingSlashingCaller) GetSlashRecord(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _IStakingSlashing.contract.Call(opts, &out, "getSlashRecord", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSlashRecord is a free data retrieval call binding the contract method 0x829673ef.
//
// Solidity: function getSlashRecord(uint256 , address ) view returns(bool)
func (_IStakingSlashing *IStakingSlashingSession) GetSlashRecord(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _IStakingSlashing.Contract.GetSlashRecord(&_IStakingSlashing.CallOpts, arg0, arg1)
}

// GetSlashRecord is a free data retrieval call binding the contract method 0x829673ef.
//
// Solidity: function getSlashRecord(uint256 , address ) view returns(bool)
func (_IStakingSlashing *IStakingSlashingCallerSession) GetSlashRecord(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _IStakingSlashing.Contract.GetSlashRecord(&_IStakingSlashing.CallOpts, arg0, arg1)
}

// GetSlashingParams is a free data retrieval call binding the contract method 0xb3fc1cb2.
//
// Solidity: function getSlashingParams() view returns(uint256[2], uint256[2])
func (_IStakingSlashing *IStakingSlashingCaller) GetSlashingParams(opts *bind.CallOpts) ([2]*big.Int, [2]*big.Int, error) {
	var out []interface{}
	err := _IStakingSlashing.contract.Call(opts, &out, "getSlashingParams")

	if err != nil {
		return *new([2]*big.Int), *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)
	out1 := *abi.ConvertType(out[1], new([2]*big.Int)).(*[2]*big.Int)

	return out0, out1, err

}

// GetSlashingParams is a free data retrieval call binding the contract method 0xb3fc1cb2.
//
// Solidity: function getSlashingParams() view returns(uint256[2], uint256[2])
func (_IStakingSlashing *IStakingSlashingSession) GetSlashingParams() ([2]*big.Int, [2]*big.Int, error) {
	return _IStakingSlashing.Contract.GetSlashingParams(&_IStakingSlashing.CallOpts)
}

// GetSlashingParams is a free data retrieval call binding the contract method 0xb3fc1cb2.
//
// Solidity: function getSlashingParams() view returns(uint256[2], uint256[2])
func (_IStakingSlashing *IStakingSlashingCallerSession) GetSlashingParams() ([2]*big.Int, [2]*big.Int, error) {
	return _IStakingSlashing.Contract.GetSlashingParams(&_IStakingSlashing.CallOpts)
}

// ClearQuitRequestList is a paid mutator transaction binding the contract method 0x111b8c02.
//
// Solidity: function clearQuitRequestList() returns()
func (_IStakingSlashing *IStakingSlashingTransactor) ClearQuitRequestList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingSlashing.contract.Transact(opts, "clearQuitRequestList")
}

// ClearQuitRequestList is a paid mutator transaction binding the contract method 0x111b8c02.
//
// Solidity: function clearQuitRequestList() returns()
func (_IStakingSlashing *IStakingSlashingSession) ClearQuitRequestList() (*types.Transaction, error) {
	return _IStakingSlashing.Contract.ClearQuitRequestList(&_IStakingSlashing.TransactOpts)
}

// ClearQuitRequestList is a paid mutator transaction binding the contract method 0x111b8c02.
//
// Solidity: function clearQuitRequestList() returns()
func (_IStakingSlashing *IStakingSlashingTransactorSession) ClearQuitRequestList() (*types.Transaction, error) {
	return _IStakingSlashing.Contract.ClearQuitRequestList(&_IStakingSlashing.TransactOpts)
}

// GetDeposits is a paid mutator transaction binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address ) returns((address,bytes,uint256))
func (_IStakingSlashing *IStakingSlashingTransactor) GetDeposits(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IStakingSlashing.contract.Transact(opts, "getDeposits", arg0)
}

// GetDeposits is a paid mutator transaction binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address ) returns((address,bytes,uint256))
func (_IStakingSlashing *IStakingSlashingSession) GetDeposits(arg0 common.Address) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.GetDeposits(&_IStakingSlashing.TransactOpts, arg0)
}

// GetDeposits is a paid mutator transaction binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address ) returns((address,bytes,uint256))
func (_IStakingSlashing *IStakingSlashingTransactorSession) GetDeposits(arg0 common.Address) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.GetDeposits(&_IStakingSlashing.TransactOpts, arg0)
}

// IsJailed is a paid mutator transaction binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address ) returns(bool)
func (_IStakingSlashing *IStakingSlashingTransactor) IsJailed(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IStakingSlashing.contract.Transact(opts, "isJailed", arg0)
}

// IsJailed is a paid mutator transaction binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address ) returns(bool)
func (_IStakingSlashing *IStakingSlashingSession) IsJailed(arg0 common.Address) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.IsJailed(&_IStakingSlashing.TransactOpts, arg0)
}

// IsJailed is a paid mutator transaction binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address ) returns(bool)
func (_IStakingSlashing *IStakingSlashingTransactorSession) IsJailed(arg0 common.Address) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.IsJailed(&_IStakingSlashing.TransactOpts, arg0)
}

// QuitRequest is a paid mutator transaction binding the contract method 0x740efec3.
//
// Solidity: function quitRequest() returns()
func (_IStakingSlashing *IStakingSlashingTransactor) QuitRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingSlashing.contract.Transact(opts, "quitRequest")
}

// QuitRequest is a paid mutator transaction binding the contract method 0x740efec3.
//
// Solidity: function quitRequest() returns()
func (_IStakingSlashing *IStakingSlashingSession) QuitRequest() (*types.Transaction, error) {
	return _IStakingSlashing.Contract.QuitRequest(&_IStakingSlashing.TransactOpts)
}

// QuitRequest is a paid mutator transaction binding the contract method 0x740efec3.
//
// Solidity: function quitRequest() returns()
func (_IStakingSlashing *IStakingSlashingTransactorSession) QuitRequest() (*types.Transaction, error) {
	return _IStakingSlashing.Contract.QuitRequest(&_IStakingSlashing.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address , address ) returns()
func (_IStakingSlashing *IStakingSlashingTransactor) SetAddress(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStakingSlashing.contract.Transact(opts, "setAddress", arg0, arg1)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address , address ) returns()
func (_IStakingSlashing *IStakingSlashingSession) SetAddress(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.SetAddress(&_IStakingSlashing.TransactOpts, arg0, arg1)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address , address ) returns()
func (_IStakingSlashing *IStakingSlashingTransactorSession) SetAddress(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.SetAddress(&_IStakingSlashing.TransactOpts, arg0, arg1)
}

// SetSlashingParams is a paid mutator transaction binding the contract method 0xde6fb88a.
//
// Solidity: function setSlashingParams(uint256[2] , uint256[2] ) returns()
func (_IStakingSlashing *IStakingSlashingTransactor) SetSlashingParams(opts *bind.TransactOpts, arg0 [2]*big.Int, arg1 [2]*big.Int) (*types.Transaction, error) {
	return _IStakingSlashing.contract.Transact(opts, "setSlashingParams", arg0, arg1)
}

// SetSlashingParams is a paid mutator transaction binding the contract method 0xde6fb88a.
//
// Solidity: function setSlashingParams(uint256[2] , uint256[2] ) returns()
func (_IStakingSlashing *IStakingSlashingSession) SetSlashingParams(arg0 [2]*big.Int, arg1 [2]*big.Int) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.SetSlashingParams(&_IStakingSlashing.TransactOpts, arg0, arg1)
}

// SetSlashingParams is a paid mutator transaction binding the contract method 0xde6fb88a.
//
// Solidity: function setSlashingParams(uint256[2] , uint256[2] ) returns()
func (_IStakingSlashing *IStakingSlashingTransactorSession) SetSlashingParams(arg0 [2]*big.Int, arg1 [2]*big.Int) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.SetSlashingParams(&_IStakingSlashing.TransactOpts, arg0, arg1)
}

// Slashing is a paid mutator transaction binding the contract method 0x5887c33c.
//
// Solidity: function slashing(bytes , bytes ) returns()
func (_IStakingSlashing *IStakingSlashingTransactor) Slashing(opts *bind.TransactOpts, arg0 []byte, arg1 []byte) (*types.Transaction, error) {
	return _IStakingSlashing.contract.Transact(opts, "slashing", arg0, arg1)
}

// Slashing is a paid mutator transaction binding the contract method 0x5887c33c.
//
// Solidity: function slashing(bytes , bytes ) returns()
func (_IStakingSlashing *IStakingSlashingSession) Slashing(arg0 []byte, arg1 []byte) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.Slashing(&_IStakingSlashing.TransactOpts, arg0, arg1)
}

// Slashing is a paid mutator transaction binding the contract method 0x5887c33c.
//
// Solidity: function slashing(bytes , bytes ) returns()
func (_IStakingSlashing *IStakingSlashingTransactorSession) Slashing(arg0 []byte, arg1 []byte) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.Slashing(&_IStakingSlashing.TransactOpts, arg0, arg1)
}

// Staking is a paid mutator transaction binding the contract method 0x5df6db49.
//
// Solidity: function staking(uint256 , bytes ) returns()
func (_IStakingSlashing *IStakingSlashingTransactor) Staking(opts *bind.TransactOpts, arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _IStakingSlashing.contract.Transact(opts, "staking", arg0, arg1)
}

// Staking is a paid mutator transaction binding the contract method 0x5df6db49.
//
// Solidity: function staking(uint256 , bytes ) returns()
func (_IStakingSlashing *IStakingSlashingSession) Staking(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.Staking(&_IStakingSlashing.TransactOpts, arg0, arg1)
}

// Staking is a paid mutator transaction binding the contract method 0x5df6db49.
//
// Solidity: function staking(uint256 , bytes ) returns()
func (_IStakingSlashing *IStakingSlashingTransactorSession) Staking(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _IStakingSlashing.Contract.Staking(&_IStakingSlashing.TransactOpts, arg0, arg1)
}

// UnJail is a paid mutator transaction binding the contract method 0x6eae5b11.
//
// Solidity: function unJail() returns()
func (_IStakingSlashing *IStakingSlashingTransactor) UnJail(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingSlashing.contract.Transact(opts, "unJail")
}

// UnJail is a paid mutator transaction binding the contract method 0x6eae5b11.
//
// Solidity: function unJail() returns()
func (_IStakingSlashing *IStakingSlashingSession) UnJail() (*types.Transaction, error) {
	return _IStakingSlashing.Contract.UnJail(&_IStakingSlashing.TransactOpts)
}

// UnJail is a paid mutator transaction binding the contract method 0x6eae5b11.
//
// Solidity: function unJail() returns()
func (_IStakingSlashing *IStakingSlashingTransactorSession) UnJail() (*types.Transaction, error) {
	return _IStakingSlashing.Contract.UnJail(&_IStakingSlashing.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_IStakingSlashing *IStakingSlashingTransactor) WithdrawToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingSlashing.contract.Transact(opts, "withdrawToken")
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_IStakingSlashing *IStakingSlashingSession) WithdrawToken() (*types.Transaction, error) {
	return _IStakingSlashing.Contract.WithdrawToken(&_IStakingSlashing.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_IStakingSlashing *IStakingSlashingTransactorSession) WithdrawToken() (*types.Transaction, error) {
	return _IStakingSlashing.Contract.WithdrawToken(&_IStakingSlashing.TransactOpts)
}

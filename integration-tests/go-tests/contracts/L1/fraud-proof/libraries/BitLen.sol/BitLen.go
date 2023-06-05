// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BitLen

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

// BitLenMetaData contains all meta data concerning the BitLen contract.
var BitLenMetaData = &bind.MetaData{
	ABI: "[]",
}

// BitLenABI is the input ABI used to generate the binding from.
// Deprecated: Use BitLenMetaData.ABI instead.
var BitLenABI = BitLenMetaData.ABI

// BitLen is an auto generated Go binding around an Ethereum contract.
type BitLen struct {
	BitLenCaller     // Read-only binding to the contract
	BitLenTransactor // Write-only binding to the contract
	BitLenFilterer   // Log filterer for contract events
}

// BitLenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BitLenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitLenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BitLenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitLenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitLenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitLenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitLenSession struct {
	Contract     *BitLen           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BitLenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitLenCallerSession struct {
	Contract *BitLenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BitLenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitLenTransactorSession struct {
	Contract     *BitLenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BitLenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BitLenRaw struct {
	Contract *BitLen // Generic contract binding to access the raw methods on
}

// BitLenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitLenCallerRaw struct {
	Contract *BitLenCaller // Generic read-only contract binding to access the raw methods on
}

// BitLenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitLenTransactorRaw struct {
	Contract *BitLenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBitLen creates a new instance of BitLen, bound to a specific deployed contract.
func NewBitLen(address common.Address, backend bind.ContractBackend) (*BitLen, error) {
	contract, err := bindBitLen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BitLen{BitLenCaller: BitLenCaller{contract: contract}, BitLenTransactor: BitLenTransactor{contract: contract}, BitLenFilterer: BitLenFilterer{contract: contract}}, nil
}

// NewBitLenCaller creates a new read-only instance of BitLen, bound to a specific deployed contract.
func NewBitLenCaller(address common.Address, caller bind.ContractCaller) (*BitLenCaller, error) {
	contract, err := bindBitLen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitLenCaller{contract: contract}, nil
}

// NewBitLenTransactor creates a new write-only instance of BitLen, bound to a specific deployed contract.
func NewBitLenTransactor(address common.Address, transactor bind.ContractTransactor) (*BitLenTransactor, error) {
	contract, err := bindBitLen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitLenTransactor{contract: contract}, nil
}

// NewBitLenFilterer creates a new log filterer instance of BitLen, bound to a specific deployed contract.
func NewBitLenFilterer(address common.Address, filterer bind.ContractFilterer) (*BitLenFilterer, error) {
	contract, err := bindBitLen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitLenFilterer{contract: contract}, nil
}

// bindBitLen binds a generic wrapper to an already deployed contract.
func bindBitLen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BitLenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitLen *BitLenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitLen.Contract.BitLenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitLen *BitLenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitLen.Contract.BitLenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitLen *BitLenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitLen.Contract.BitLenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitLen *BitLenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitLen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitLen *BitLenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitLen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitLen *BitLenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitLen.Contract.contract.Transact(opts, method, params...)
}

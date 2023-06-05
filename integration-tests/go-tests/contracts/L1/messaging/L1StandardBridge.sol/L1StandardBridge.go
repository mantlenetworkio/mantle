// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package L1StandardBridge

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

// L1StandardBridgeMetaData contains all meta data concerning the L1StandardBridge contract.
var L1StandardBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ETHDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ETHWithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donateETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeERC20Withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeETHWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeMantleWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2TokenBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1MantleAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1MantleAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// L1StandardBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L1StandardBridgeMetaData.ABI instead.
var L1StandardBridgeABI = L1StandardBridgeMetaData.ABI

// L1StandardBridge is an auto generated Go binding around an Ethereum contract.
type L1StandardBridge struct {
	L1StandardBridgeCaller     // Read-only binding to the contract
	L1StandardBridgeTransactor // Write-only binding to the contract
	L1StandardBridgeFilterer   // Log filterer for contract events
}

// L1StandardBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1StandardBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1StandardBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1StandardBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1StandardBridgeSession struct {
	Contract     *L1StandardBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1StandardBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1StandardBridgeCallerSession struct {
	Contract *L1StandardBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// L1StandardBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1StandardBridgeTransactorSession struct {
	Contract     *L1StandardBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L1StandardBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1StandardBridgeRaw struct {
	Contract *L1StandardBridge // Generic contract binding to access the raw methods on
}

// L1StandardBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1StandardBridgeCallerRaw struct {
	Contract *L1StandardBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L1StandardBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1StandardBridgeTransactorRaw struct {
	Contract *L1StandardBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1StandardBridge creates a new instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridge(address common.Address, backend bind.ContractBackend) (*L1StandardBridge, error) {
	contract, err := bindL1StandardBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridge{L1StandardBridgeCaller: L1StandardBridgeCaller{contract: contract}, L1StandardBridgeTransactor: L1StandardBridgeTransactor{contract: contract}, L1StandardBridgeFilterer: L1StandardBridgeFilterer{contract: contract}}, nil
}

// NewL1StandardBridgeCaller creates a new read-only instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridgeCaller(address common.Address, caller bind.ContractCaller) (*L1StandardBridgeCaller, error) {
	contract, err := bindL1StandardBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeCaller{contract: contract}, nil
}

// NewL1StandardBridgeTransactor creates a new write-only instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L1StandardBridgeTransactor, error) {
	contract, err := bindL1StandardBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeTransactor{contract: contract}, nil
}

// NewL1StandardBridgeFilterer creates a new log filterer instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L1StandardBridgeFilterer, error) {
	contract, err := bindL1StandardBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeFilterer{contract: contract}, nil
}

// bindL1StandardBridge binds a generic wrapper to an already deployed contract.
func bindL1StandardBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1StandardBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1StandardBridge *L1StandardBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1StandardBridge.Contract.L1StandardBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1StandardBridge *L1StandardBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.L1StandardBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1StandardBridge *L1StandardBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.L1StandardBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1StandardBridge *L1StandardBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1StandardBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1StandardBridge *L1StandardBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1StandardBridge *L1StandardBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.contract.Transact(opts, method, params...)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L1StandardBridge *L1StandardBridgeCaller) Deposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "deposits", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L1StandardBridge *L1StandardBridgeSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L1StandardBridge.Contract.Deposits(&_L1StandardBridge.CallOpts, arg0, arg1)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L1StandardBridge *L1StandardBridgeCallerSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L1StandardBridge.Contract.Deposits(&_L1StandardBridge.CallOpts, arg0, arg1)
}

// L1MantleAddress is a free data retrieval call binding the contract method 0x33151848.
//
// Solidity: function l1MantleAddress() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) L1MantleAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "l1MantleAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1MantleAddress is a free data retrieval call binding the contract method 0x33151848.
//
// Solidity: function l1MantleAddress() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) L1MantleAddress() (common.Address, error) {
	return _L1StandardBridge.Contract.L1MantleAddress(&_L1StandardBridge.CallOpts)
}

// L1MantleAddress is a free data retrieval call binding the contract method 0x33151848.
//
// Solidity: function l1MantleAddress() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) L1MantleAddress() (common.Address, error) {
	return _L1StandardBridge.Contract.L1MantleAddress(&_L1StandardBridge.CallOpts)
}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) L2TokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "l2TokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) L2TokenBridge() (common.Address, error) {
	return _L1StandardBridge.Contract.L2TokenBridge(&_L1StandardBridge.CallOpts)
}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) L2TokenBridge() (common.Address, error) {
	return _L1StandardBridge.Contract.L2TokenBridge(&_L1StandardBridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) Messenger() (common.Address, error) {
	return _L1StandardBridge.Contract.Messenger(&_L1StandardBridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) Messenger() (common.Address, error) {
	return _L1StandardBridge.Contract.Messenger(&_L1StandardBridge.CallOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) DepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "depositERC20", _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) DepositERC20To(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "depositERC20To", _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20To(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20To(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// DepositETH is a paid mutator transaction binding the contract method 0xb1a1a882.
//
// Solidity: function depositETH(uint32 _l2Gas, bytes _data) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) DepositETH(opts *bind.TransactOpts, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "depositETH", _l2Gas, _data)
}

// DepositETH is a paid mutator transaction binding the contract method 0xb1a1a882.
//
// Solidity: function depositETH(uint32 _l2Gas, bytes _data) payable returns()
func (_L1StandardBridge *L1StandardBridgeSession) DepositETH(_l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositETH(&_L1StandardBridge.TransactOpts, _l2Gas, _data)
}

// DepositETH is a paid mutator transaction binding the contract method 0xb1a1a882.
//
// Solidity: function depositETH(uint32 _l2Gas, bytes _data) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) DepositETH(_l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositETH(&_L1StandardBridge.TransactOpts, _l2Gas, _data)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0x9a2ac6d5.
//
// Solidity: function depositETHTo(address _to, uint32 _l2Gas, bytes _data) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) DepositETHTo(opts *bind.TransactOpts, _to common.Address, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "depositETHTo", _to, _l2Gas, _data)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0x9a2ac6d5.
//
// Solidity: function depositETHTo(address _to, uint32 _l2Gas, bytes _data) payable returns()
func (_L1StandardBridge *L1StandardBridgeSession) DepositETHTo(_to common.Address, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositETHTo(&_L1StandardBridge.TransactOpts, _to, _l2Gas, _data)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0x9a2ac6d5.
//
// Solidity: function depositETHTo(address _to, uint32 _l2Gas, bytes _data) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) DepositETHTo(_to common.Address, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositETHTo(&_L1StandardBridge.TransactOpts, _to, _l2Gas, _data)
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) DonateETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "donateETH")
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_L1StandardBridge *L1StandardBridgeSession) DonateETH() (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DonateETH(&_L1StandardBridge.TransactOpts)
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) DonateETH() (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DonateETH(&_L1StandardBridge.TransactOpts)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) FinalizeERC20Withdrawal(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "finalizeERC20Withdrawal", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeERC20Withdrawal(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeERC20Withdrawal(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeETHWithdrawal is a paid mutator transaction binding the contract method 0x1532ec34.
//
// Solidity: function finalizeETHWithdrawal(address _from, address _to, uint256 _amount, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) FinalizeETHWithdrawal(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "finalizeETHWithdrawal", _from, _to, _amount, _data)
}

// FinalizeETHWithdrawal is a paid mutator transaction binding the contract method 0x1532ec34.
//
// Solidity: function finalizeETHWithdrawal(address _from, address _to, uint256 _amount, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeSession) FinalizeETHWithdrawal(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeETHWithdrawal(&_L1StandardBridge.TransactOpts, _from, _to, _amount, _data)
}

// FinalizeETHWithdrawal is a paid mutator transaction binding the contract method 0x1532ec34.
//
// Solidity: function finalizeETHWithdrawal(address _from, address _to, uint256 _amount, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) FinalizeETHWithdrawal(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeETHWithdrawal(&_L1StandardBridge.TransactOpts, _from, _to, _amount, _data)
}

// FinalizeMantleWithdrawal is a paid mutator transaction binding the contract method 0xf82b418e.
//
// Solidity: function finalizeMantleWithdrawal(address _from, address _to, uint256 _amount, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) FinalizeMantleWithdrawal(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "finalizeMantleWithdrawal", _from, _to, _amount, _data)
}

// FinalizeMantleWithdrawal is a paid mutator transaction binding the contract method 0xf82b418e.
//
// Solidity: function finalizeMantleWithdrawal(address _from, address _to, uint256 _amount, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeSession) FinalizeMantleWithdrawal(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeMantleWithdrawal(&_L1StandardBridge.TransactOpts, _from, _to, _amount, _data)
}

// FinalizeMantleWithdrawal is a paid mutator transaction binding the contract method 0xf82b418e.
//
// Solidity: function finalizeMantleWithdrawal(address _from, address _to, uint256 _amount, bytes _data) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) FinalizeMantleWithdrawal(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeMantleWithdrawal(&_L1StandardBridge.TransactOpts, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _l1messenger, address _l2TokenBridge, address _l1MantleAddress) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) Initialize(opts *bind.TransactOpts, _l1messenger common.Address, _l2TokenBridge common.Address, _l1MantleAddress common.Address) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "initialize", _l1messenger, _l2TokenBridge, _l1MantleAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _l1messenger, address _l2TokenBridge, address _l1MantleAddress) returns()
func (_L1StandardBridge *L1StandardBridgeSession) Initialize(_l1messenger common.Address, _l2TokenBridge common.Address, _l1MantleAddress common.Address) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.Initialize(&_L1StandardBridge.TransactOpts, _l1messenger, _l2TokenBridge, _l1MantleAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _l1messenger, address _l2TokenBridge, address _l1MantleAddress) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) Initialize(_l1messenger common.Address, _l2TokenBridge common.Address, _l1MantleAddress common.Address) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.Initialize(&_L1StandardBridge.TransactOpts, _l1messenger, _l2TokenBridge, _l1MantleAddress)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1StandardBridge *L1StandardBridgeSession) Receive() (*types.Transaction, error) {
	return _L1StandardBridge.Contract.Receive(&_L1StandardBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _L1StandardBridge.Contract.Receive(&_L1StandardBridge.TransactOpts)
}

// L1StandardBridgeERC20DepositInitiatedIterator is returned from FilterERC20DepositInitiated and is used to iterate over the raw logs and unpacked data for ERC20DepositInitiated events raised by the L1StandardBridge contract.
type L1StandardBridgeERC20DepositInitiatedIterator struct {
	Event *L1StandardBridgeERC20DepositInitiated // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeERC20DepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeERC20DepositInitiated)
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
		it.Event = new(L1StandardBridgeERC20DepositInitiated)
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
func (it *L1StandardBridgeERC20DepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeERC20DepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeERC20DepositInitiated represents a ERC20DepositInitiated event raised by the L1StandardBridge contract.
type L1StandardBridgeERC20DepositInitiated struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterERC20DepositInitiated is a free log retrieval operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterERC20DepositInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*L1StandardBridgeERC20DepositInitiatedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeERC20DepositInitiatedIterator{contract: _L1StandardBridge.contract, event: "ERC20DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20DepositInitiated is a free log subscription operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchERC20DepositInitiated(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeERC20DepositInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeERC20DepositInitiated)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
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

// ParseERC20DepositInitiated is a log parse operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseERC20DepositInitiated(log types.Log) (*L1StandardBridgeERC20DepositInitiated, error) {
	event := new(L1StandardBridgeERC20DepositInitiated)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeERC20WithdrawalFinalizedIterator is returned from FilterERC20WithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ERC20WithdrawalFinalized events raised by the L1StandardBridge contract.
type L1StandardBridgeERC20WithdrawalFinalizedIterator struct {
	Event *L1StandardBridgeERC20WithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeERC20WithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeERC20WithdrawalFinalized)
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
		it.Event = new(L1StandardBridgeERC20WithdrawalFinalized)
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
func (it *L1StandardBridgeERC20WithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeERC20WithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeERC20WithdrawalFinalized represents a ERC20WithdrawalFinalized event raised by the L1StandardBridge contract.
type L1StandardBridgeERC20WithdrawalFinalized struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterERC20WithdrawalFinalized is a free log retrieval operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterERC20WithdrawalFinalized(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*L1StandardBridgeERC20WithdrawalFinalizedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeERC20WithdrawalFinalizedIterator{contract: _L1StandardBridge.contract, event: "ERC20WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20WithdrawalFinalized is a free log subscription operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchERC20WithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeERC20WithdrawalFinalized, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeERC20WithdrawalFinalized)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
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

// ParseERC20WithdrawalFinalized is a log parse operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseERC20WithdrawalFinalized(log types.Log) (*L1StandardBridgeERC20WithdrawalFinalized, error) {
	event := new(L1StandardBridgeERC20WithdrawalFinalized)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeETHDepositInitiatedIterator is returned from FilterETHDepositInitiated and is used to iterate over the raw logs and unpacked data for ETHDepositInitiated events raised by the L1StandardBridge contract.
type L1StandardBridgeETHDepositInitiatedIterator struct {
	Event *L1StandardBridgeETHDepositInitiated // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeETHDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeETHDepositInitiated)
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
		it.Event = new(L1StandardBridgeETHDepositInitiated)
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
func (it *L1StandardBridgeETHDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeETHDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeETHDepositInitiated represents a ETHDepositInitiated event raised by the L1StandardBridge contract.
type L1StandardBridgeETHDepositInitiated struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterETHDepositInitiated is a free log retrieval operation binding the contract event 0x35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23.
//
// Solidity: event ETHDepositInitiated(address indexed _from, address indexed _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterETHDepositInitiated(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*L1StandardBridgeETHDepositInitiatedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ETHDepositInitiated", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeETHDepositInitiatedIterator{contract: _L1StandardBridge.contract, event: "ETHDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchETHDepositInitiated is a free log subscription operation binding the contract event 0x35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23.
//
// Solidity: event ETHDepositInitiated(address indexed _from, address indexed _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchETHDepositInitiated(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeETHDepositInitiated, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ETHDepositInitiated", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeETHDepositInitiated)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ETHDepositInitiated", log); err != nil {
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

// ParseETHDepositInitiated is a log parse operation binding the contract event 0x35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23.
//
// Solidity: event ETHDepositInitiated(address indexed _from, address indexed _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseETHDepositInitiated(log types.Log) (*L1StandardBridgeETHDepositInitiated, error) {
	event := new(L1StandardBridgeETHDepositInitiated)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ETHDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeETHWithdrawalFinalizedIterator is returned from FilterETHWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ETHWithdrawalFinalized events raised by the L1StandardBridge contract.
type L1StandardBridgeETHWithdrawalFinalizedIterator struct {
	Event *L1StandardBridgeETHWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeETHWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeETHWithdrawalFinalized)
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
		it.Event = new(L1StandardBridgeETHWithdrawalFinalized)
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
func (it *L1StandardBridgeETHWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeETHWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeETHWithdrawalFinalized represents a ETHWithdrawalFinalized event raised by the L1StandardBridge contract.
type L1StandardBridgeETHWithdrawalFinalized struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterETHWithdrawalFinalized is a free log retrieval operation binding the contract event 0x2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631.
//
// Solidity: event ETHWithdrawalFinalized(address indexed _from, address indexed _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterETHWithdrawalFinalized(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*L1StandardBridgeETHWithdrawalFinalizedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ETHWithdrawalFinalized", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeETHWithdrawalFinalizedIterator{contract: _L1StandardBridge.contract, event: "ETHWithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchETHWithdrawalFinalized is a free log subscription operation binding the contract event 0x2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631.
//
// Solidity: event ETHWithdrawalFinalized(address indexed _from, address indexed _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchETHWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeETHWithdrawalFinalized, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ETHWithdrawalFinalized", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeETHWithdrawalFinalized)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ETHWithdrawalFinalized", log); err != nil {
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

// ParseETHWithdrawalFinalized is a log parse operation binding the contract event 0x2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631.
//
// Solidity: event ETHWithdrawalFinalized(address indexed _from, address indexed _to, uint256 _amount, bytes _data)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseETHWithdrawalFinalized(log types.Log) (*L1StandardBridgeETHWithdrawalFinalized, error) {
	event := new(L1StandardBridgeETHWithdrawalFinalized)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ETHWithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1StandardBridge contract.
type L1StandardBridgeInitializedIterator struct {
	Event *L1StandardBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeInitialized)
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
		it.Event = new(L1StandardBridgeInitialized)
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
func (it *L1StandardBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeInitialized represents a Initialized event raised by the L1StandardBridge contract.
type L1StandardBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1StandardBridgeInitializedIterator, error) {

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeInitializedIterator{contract: _L1StandardBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeInitialized)
				if err := _L1StandardBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseInitialized(log types.Log) (*L1StandardBridgeInitialized, error) {
	event := new(L1StandardBridgeInitialized)
	if err := _L1StandardBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

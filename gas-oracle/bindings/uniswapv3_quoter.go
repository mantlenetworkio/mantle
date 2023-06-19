// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// Uniswapv3QuoterMetaData contains all meta data concerning the Uniswapv3Quoter contract.
var Uniswapv3QuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Uniswapv3QuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapv3QuoterMetaData.ABI instead.
var Uniswapv3QuoterABI = Uniswapv3QuoterMetaData.ABI

// Uniswapv3Quoter is an auto generated Go binding around an Ethereum contract.
type Uniswapv3Quoter struct {
	Uniswapv3QuoterCaller     // Read-only binding to the contract
	Uniswapv3QuoterTransactor // Write-only binding to the contract
	Uniswapv3QuoterFilterer   // Log filterer for contract events
}

// Uniswapv3QuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv3QuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3QuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv3QuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3QuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv3QuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3QuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv3QuoterSession struct {
	Contract     *Uniswapv3Quoter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapv3QuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv3QuoterCallerSession struct {
	Contract *Uniswapv3QuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Uniswapv3QuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv3QuoterTransactorSession struct {
	Contract     *Uniswapv3QuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Uniswapv3QuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv3QuoterRaw struct {
	Contract *Uniswapv3Quoter // Generic contract binding to access the raw methods on
}

// Uniswapv3QuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv3QuoterCallerRaw struct {
	Contract *Uniswapv3QuoterCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv3QuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv3QuoterTransactorRaw struct {
	Contract *Uniswapv3QuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv3Quoter creates a new instance of Uniswapv3Quoter, bound to a specific deployed contract.
func NewUniswapv3Quoter(address common.Address, backend bind.ContractBackend) (*Uniswapv3Quoter, error) {
	contract, err := bindUniswapv3Quoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3Quoter{Uniswapv3QuoterCaller: Uniswapv3QuoterCaller{contract: contract}, Uniswapv3QuoterTransactor: Uniswapv3QuoterTransactor{contract: contract}, Uniswapv3QuoterFilterer: Uniswapv3QuoterFilterer{contract: contract}}, nil
}

// NewUniswapv3QuoterCaller creates a new read-only instance of Uniswapv3Quoter, bound to a specific deployed contract.
func NewUniswapv3QuoterCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv3QuoterCaller, error) {
	contract, err := bindUniswapv3Quoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3QuoterCaller{contract: contract}, nil
}

// NewUniswapv3QuoterTransactor creates a new write-only instance of Uniswapv3Quoter, bound to a specific deployed contract.
func NewUniswapv3QuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv3QuoterTransactor, error) {
	contract, err := bindUniswapv3Quoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3QuoterTransactor{contract: contract}, nil
}

// NewUniswapv3QuoterFilterer creates a new log filterer instance of Uniswapv3Quoter, bound to a specific deployed contract.
func NewUniswapv3QuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv3QuoterFilterer, error) {
	contract, err := bindUniswapv3Quoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3QuoterFilterer{contract: contract}, nil
}

// bindUniswapv3Quoter binds a generic wrapper to an already deployed contract.
func bindUniswapv3Quoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Uniswapv3QuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3Quoter *Uniswapv3QuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3Quoter.Contract.Uniswapv3QuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3Quoter *Uniswapv3QuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.Uniswapv3QuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3Quoter *Uniswapv3QuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.Uniswapv3QuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3Quoter *Uniswapv3QuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3Quoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3Quoter *Uniswapv3QuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3Quoter *Uniswapv3QuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Uniswapv3Quoter *Uniswapv3QuoterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3Quoter.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Uniswapv3Quoter *Uniswapv3QuoterSession) WETH9() (common.Address, error) {
	return _Uniswapv3Quoter.Contract.WETH9(&_Uniswapv3Quoter.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Uniswapv3Quoter *Uniswapv3QuoterCallerSession) WETH9() (common.Address, error) {
	return _Uniswapv3Quoter.Contract.WETH9(&_Uniswapv3Quoter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv3Quoter *Uniswapv3QuoterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3Quoter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv3Quoter *Uniswapv3QuoterSession) Factory() (common.Address, error) {
	return _Uniswapv3Quoter.Contract.Factory(&_Uniswapv3Quoter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv3Quoter *Uniswapv3QuoterCallerSession) Factory() (common.Address, error) {
	return _Uniswapv3Quoter.Contract.Factory(&_Uniswapv3Quoter.CallOpts)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_Uniswapv3Quoter *Uniswapv3QuoterCaller) UniswapV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _Uniswapv3Quoter.contract.Call(opts, &out, "uniswapV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_Uniswapv3Quoter *Uniswapv3QuoterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _Uniswapv3Quoter.Contract.UniswapV3SwapCallback(&_Uniswapv3Quoter.CallOpts, amount0Delta, amount1Delta, path)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_Uniswapv3Quoter *Uniswapv3QuoterCallerSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _Uniswapv3Quoter.Contract.UniswapV3SwapCallback(&_Uniswapv3Quoter.CallOpts, amount0Delta, amount1Delta, path)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut)
func (_Uniswapv3Quoter *Uniswapv3QuoterTransactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut)
func (_Uniswapv3Quoter *Uniswapv3QuoterSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.QuoteExactInput(&_Uniswapv3Quoter.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut)
func (_Uniswapv3Quoter *Uniswapv3QuoterTransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.QuoteExactInput(&_Uniswapv3Quoter.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) returns(uint256 amountOut)
func (_Uniswapv3Quoter *Uniswapv3QuoterTransactor) QuoteExactInputSingle(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.contract.Transact(opts, "quoteExactInputSingle", tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) returns(uint256 amountOut)
func (_Uniswapv3Quoter *Uniswapv3QuoterSession) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.QuoteExactInputSingle(&_Uniswapv3Quoter.TransactOpts, tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) returns(uint256 amountOut)
func (_Uniswapv3Quoter *Uniswapv3QuoterTransactorSession) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.QuoteExactInputSingle(&_Uniswapv3Quoter.TransactOpts, tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn)
func (_Uniswapv3Quoter *Uniswapv3QuoterTransactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn)
func (_Uniswapv3Quoter *Uniswapv3QuoterSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.QuoteExactOutput(&_Uniswapv3Quoter.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn)
func (_Uniswapv3Quoter *Uniswapv3QuoterTransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.QuoteExactOutput(&_Uniswapv3Quoter.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x30d07f21.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountOut, uint160 sqrtPriceLimitX96) returns(uint256 amountIn)
func (_Uniswapv3Quoter *Uniswapv3QuoterTransactor) QuoteExactOutputSingle(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountOut *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.contract.Transact(opts, "quoteExactOutputSingle", tokenIn, tokenOut, fee, amountOut, sqrtPriceLimitX96)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x30d07f21.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountOut, uint160 sqrtPriceLimitX96) returns(uint256 amountIn)
func (_Uniswapv3Quoter *Uniswapv3QuoterSession) QuoteExactOutputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountOut *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.QuoteExactOutputSingle(&_Uniswapv3Quoter.TransactOpts, tokenIn, tokenOut, fee, amountOut, sqrtPriceLimitX96)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x30d07f21.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountOut, uint160 sqrtPriceLimitX96) returns(uint256 amountIn)
func (_Uniswapv3Quoter *Uniswapv3QuoterTransactorSession) QuoteExactOutputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountOut *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Uniswapv3Quoter.Contract.QuoteExactOutputSingle(&_Uniswapv3Quoter.TransactOpts, tokenIn, tokenOut, fee, amountOut, sqrtPriceLimitX96)
}

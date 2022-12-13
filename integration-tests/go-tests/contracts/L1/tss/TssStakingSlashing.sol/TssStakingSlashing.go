// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TssStakingSlashing

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

// TssStakingSlashingMetaData contains all meta data concerning the TssStakingSlashing contract.
var TssStakingSlashingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"pledgor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIStakingSlashing.DepositInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"AddDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTssStakingSlashing.SlashType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"Slashing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BitToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"batchGetDeposits\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pledgor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingSlashing.DepositInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearQuitRequestList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pledgor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getDeposits\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pledgor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingSlashing.DepositInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuitRequestList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getSlashRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlashingParams\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bitToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tssGroupContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"byteListA\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"byteListB\",\"type\":\"bytes\"}],\"name\":\"isEqual\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isJailed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quitRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"quitRequestList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tssGroup\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_slashAmount\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_exIncome\",\"type\":\"uint256[2]\"}],\"name\":\"setSlashingParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messageBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"slashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_pubKey\",\"type\":\"bytes\"}],\"name\":\"staking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssGroupContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unJail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TssStakingSlashingABI is the input ABI used to generate the binding from.
// Deprecated: Use TssStakingSlashingMetaData.ABI instead.
var TssStakingSlashingABI = TssStakingSlashingMetaData.ABI

// TssStakingSlashing is an auto generated Go binding around an Ethereum contract.
type TssStakingSlashing struct {
	TssStakingSlashingCaller     // Read-only binding to the contract
	TssStakingSlashingTransactor // Write-only binding to the contract
	TssStakingSlashingFilterer   // Log filterer for contract events
}

// TssStakingSlashingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TssStakingSlashingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssStakingSlashingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TssStakingSlashingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssStakingSlashingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TssStakingSlashingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssStakingSlashingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TssStakingSlashingSession struct {
	Contract     *TssStakingSlashing // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TssStakingSlashingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TssStakingSlashingCallerSession struct {
	Contract *TssStakingSlashingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TssStakingSlashingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TssStakingSlashingTransactorSession struct {
	Contract     *TssStakingSlashingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TssStakingSlashingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TssStakingSlashingRaw struct {
	Contract *TssStakingSlashing // Generic contract binding to access the raw methods on
}

// TssStakingSlashingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TssStakingSlashingCallerRaw struct {
	Contract *TssStakingSlashingCaller // Generic read-only contract binding to access the raw methods on
}

// TssStakingSlashingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TssStakingSlashingTransactorRaw struct {
	Contract *TssStakingSlashingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTssStakingSlashing creates a new instance of TssStakingSlashing, bound to a specific deployed contract.
func NewTssStakingSlashing(address common.Address, backend bind.ContractBackend) (*TssStakingSlashing, error) {
	contract, err := bindTssStakingSlashing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TssStakingSlashing{TssStakingSlashingCaller: TssStakingSlashingCaller{contract: contract}, TssStakingSlashingTransactor: TssStakingSlashingTransactor{contract: contract}, TssStakingSlashingFilterer: TssStakingSlashingFilterer{contract: contract}}, nil
}

// NewTssStakingSlashingCaller creates a new read-only instance of TssStakingSlashing, bound to a specific deployed contract.
func NewTssStakingSlashingCaller(address common.Address, caller bind.ContractCaller) (*TssStakingSlashingCaller, error) {
	contract, err := bindTssStakingSlashing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TssStakingSlashingCaller{contract: contract}, nil
}

// NewTssStakingSlashingTransactor creates a new write-only instance of TssStakingSlashing, bound to a specific deployed contract.
func NewTssStakingSlashingTransactor(address common.Address, transactor bind.ContractTransactor) (*TssStakingSlashingTransactor, error) {
	contract, err := bindTssStakingSlashing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TssStakingSlashingTransactor{contract: contract}, nil
}

// NewTssStakingSlashingFilterer creates a new log filterer instance of TssStakingSlashing, bound to a specific deployed contract.
func NewTssStakingSlashingFilterer(address common.Address, filterer bind.ContractFilterer) (*TssStakingSlashingFilterer, error) {
	contract, err := bindTssStakingSlashing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TssStakingSlashingFilterer{contract: contract}, nil
}

// bindTssStakingSlashing binds a generic wrapper to an already deployed contract.
func bindTssStakingSlashing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TssStakingSlashingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TssStakingSlashing *TssStakingSlashingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TssStakingSlashing.Contract.TssStakingSlashingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TssStakingSlashing *TssStakingSlashingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.TssStakingSlashingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TssStakingSlashing *TssStakingSlashingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.TssStakingSlashingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TssStakingSlashing *TssStakingSlashingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TssStakingSlashing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TssStakingSlashing *TssStakingSlashingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TssStakingSlashing *TssStakingSlashingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.contract.Transact(opts, method, params...)
}

// BitToken is a free data retrieval call binding the contract method 0xeb20b589.
//
// Solidity: function BitToken() view returns(address)
func (_TssStakingSlashing *TssStakingSlashingCaller) BitToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "BitToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BitToken is a free data retrieval call binding the contract method 0xeb20b589.
//
// Solidity: function BitToken() view returns(address)
func (_TssStakingSlashing *TssStakingSlashingSession) BitToken() (common.Address, error) {
	return _TssStakingSlashing.Contract.BitToken(&_TssStakingSlashing.CallOpts)
}

// BitToken is a free data retrieval call binding the contract method 0xeb20b589.
//
// Solidity: function BitToken() view returns(address)
func (_TssStakingSlashing *TssStakingSlashingCallerSession) BitToken() (common.Address, error) {
	return _TssStakingSlashing.Contract.BitToken(&_TssStakingSlashing.CallOpts)
}

// BatchGetDeposits is a free data retrieval call binding the contract method 0x793505c8.
//
// Solidity: function batchGetDeposits(address[] users) view returns((address,bytes,uint256)[])
func (_TssStakingSlashing *TssStakingSlashingCaller) BatchGetDeposits(opts *bind.CallOpts, users []common.Address) ([]IStakingSlashingDepositInfo, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "batchGetDeposits", users)

	if err != nil {
		return *new([]IStakingSlashingDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingSlashingDepositInfo)).(*[]IStakingSlashingDepositInfo)

	return out0, err

}

// BatchGetDeposits is a free data retrieval call binding the contract method 0x793505c8.
//
// Solidity: function batchGetDeposits(address[] users) view returns((address,bytes,uint256)[])
func (_TssStakingSlashing *TssStakingSlashingSession) BatchGetDeposits(users []common.Address) ([]IStakingSlashingDepositInfo, error) {
	return _TssStakingSlashing.Contract.BatchGetDeposits(&_TssStakingSlashing.CallOpts, users)
}

// BatchGetDeposits is a free data retrieval call binding the contract method 0x793505c8.
//
// Solidity: function batchGetDeposits(address[] users) view returns((address,bytes,uint256)[])
func (_TssStakingSlashing *TssStakingSlashingCallerSession) BatchGetDeposits(users []common.Address) ([]IStakingSlashingDepositInfo, error) {
	return _TssStakingSlashing.Contract.BatchGetDeposits(&_TssStakingSlashing.CallOpts, users)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(address pledgor, bytes pubKey, uint256 amount)
func (_TssStakingSlashing *TssStakingSlashingCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Pledgor common.Address
	PubKey  []byte
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Pledgor common.Address
		PubKey  []byte
		Amount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pledgor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PubKey = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(address pledgor, bytes pubKey, uint256 amount)
func (_TssStakingSlashing *TssStakingSlashingSession) Deposits(arg0 common.Address) (struct {
	Pledgor common.Address
	PubKey  []byte
	Amount  *big.Int
}, error) {
	return _TssStakingSlashing.Contract.Deposits(&_TssStakingSlashing.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(address pledgor, bytes pubKey, uint256 amount)
func (_TssStakingSlashing *TssStakingSlashingCallerSession) Deposits(arg0 common.Address) (struct {
	Pledgor common.Address
	PubKey  []byte
	Amount  *big.Int
}, error) {
	return _TssStakingSlashing.Contract.Deposits(&_TssStakingSlashing.CallOpts, arg0)
}

// ExIncome is a free data retrieval call binding the contract method 0x3f950438.
//
// Solidity: function exIncome(uint256 ) view returns(uint256)
func (_TssStakingSlashing *TssStakingSlashingCaller) ExIncome(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "exIncome", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExIncome is a free data retrieval call binding the contract method 0x3f950438.
//
// Solidity: function exIncome(uint256 ) view returns(uint256)
func (_TssStakingSlashing *TssStakingSlashingSession) ExIncome(arg0 *big.Int) (*big.Int, error) {
	return _TssStakingSlashing.Contract.ExIncome(&_TssStakingSlashing.CallOpts, arg0)
}

// ExIncome is a free data retrieval call binding the contract method 0x3f950438.
//
// Solidity: function exIncome(uint256 ) view returns(uint256)
func (_TssStakingSlashing *TssStakingSlashingCallerSession) ExIncome(arg0 *big.Int) (*big.Int, error) {
	return _TssStakingSlashing.Contract.ExIncome(&_TssStakingSlashing.CallOpts, arg0)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address user) view returns((address,bytes,uint256))
func (_TssStakingSlashing *TssStakingSlashingCaller) GetDeposits(opts *bind.CallOpts, user common.Address) (IStakingSlashingDepositInfo, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "getDeposits", user)

	if err != nil {
		return *new(IStakingSlashingDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakingSlashingDepositInfo)).(*IStakingSlashingDepositInfo)

	return out0, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address user) view returns((address,bytes,uint256))
func (_TssStakingSlashing *TssStakingSlashingSession) GetDeposits(user common.Address) (IStakingSlashingDepositInfo, error) {
	return _TssStakingSlashing.Contract.GetDeposits(&_TssStakingSlashing.CallOpts, user)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address user) view returns((address,bytes,uint256))
func (_TssStakingSlashing *TssStakingSlashingCallerSession) GetDeposits(user common.Address) (IStakingSlashingDepositInfo, error) {
	return _TssStakingSlashing.Contract.GetDeposits(&_TssStakingSlashing.CallOpts, user)
}

// GetQuitRequestList is a free data retrieval call binding the contract method 0x490ab9d6.
//
// Solidity: function getQuitRequestList() view returns(address[])
func (_TssStakingSlashing *TssStakingSlashingCaller) GetQuitRequestList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "getQuitRequestList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetQuitRequestList is a free data retrieval call binding the contract method 0x490ab9d6.
//
// Solidity: function getQuitRequestList() view returns(address[])
func (_TssStakingSlashing *TssStakingSlashingSession) GetQuitRequestList() ([]common.Address, error) {
	return _TssStakingSlashing.Contract.GetQuitRequestList(&_TssStakingSlashing.CallOpts)
}

// GetQuitRequestList is a free data retrieval call binding the contract method 0x490ab9d6.
//
// Solidity: function getQuitRequestList() view returns(address[])
func (_TssStakingSlashing *TssStakingSlashingCallerSession) GetQuitRequestList() ([]common.Address, error) {
	return _TssStakingSlashing.Contract.GetQuitRequestList(&_TssStakingSlashing.CallOpts)
}

// GetSlashRecord is a free data retrieval call binding the contract method 0x829673ef.
//
// Solidity: function getSlashRecord(uint256 batchIndex, address user) view returns(bool)
func (_TssStakingSlashing *TssStakingSlashingCaller) GetSlashRecord(opts *bind.CallOpts, batchIndex *big.Int, user common.Address) (bool, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "getSlashRecord", batchIndex, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSlashRecord is a free data retrieval call binding the contract method 0x829673ef.
//
// Solidity: function getSlashRecord(uint256 batchIndex, address user) view returns(bool)
func (_TssStakingSlashing *TssStakingSlashingSession) GetSlashRecord(batchIndex *big.Int, user common.Address) (bool, error) {
	return _TssStakingSlashing.Contract.GetSlashRecord(&_TssStakingSlashing.CallOpts, batchIndex, user)
}

// GetSlashRecord is a free data retrieval call binding the contract method 0x829673ef.
//
// Solidity: function getSlashRecord(uint256 batchIndex, address user) view returns(bool)
func (_TssStakingSlashing *TssStakingSlashingCallerSession) GetSlashRecord(batchIndex *big.Int, user common.Address) (bool, error) {
	return _TssStakingSlashing.Contract.GetSlashRecord(&_TssStakingSlashing.CallOpts, batchIndex, user)
}

// GetSlashingParams is a free data retrieval call binding the contract method 0xb3fc1cb2.
//
// Solidity: function getSlashingParams() view returns(uint256[2], uint256[2])
func (_TssStakingSlashing *TssStakingSlashingCaller) GetSlashingParams(opts *bind.CallOpts) ([2]*big.Int, [2]*big.Int, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "getSlashingParams")

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
func (_TssStakingSlashing *TssStakingSlashingSession) GetSlashingParams() ([2]*big.Int, [2]*big.Int, error) {
	return _TssStakingSlashing.Contract.GetSlashingParams(&_TssStakingSlashing.CallOpts)
}

// GetSlashingParams is a free data retrieval call binding the contract method 0xb3fc1cb2.
//
// Solidity: function getSlashingParams() view returns(uint256[2], uint256[2])
func (_TssStakingSlashing *TssStakingSlashingCallerSession) GetSlashingParams() ([2]*big.Int, [2]*big.Int, error) {
	return _TssStakingSlashing.Contract.GetSlashingParams(&_TssStakingSlashing.CallOpts)
}

// IsEqual is a free data retrieval call binding the contract method 0x34359808.
//
// Solidity: function isEqual(bytes byteListA, bytes byteListB) pure returns(bool)
func (_TssStakingSlashing *TssStakingSlashingCaller) IsEqual(opts *bind.CallOpts, byteListA []byte, byteListB []byte) (bool, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "isEqual", byteListA, byteListB)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEqual is a free data retrieval call binding the contract method 0x34359808.
//
// Solidity: function isEqual(bytes byteListA, bytes byteListB) pure returns(bool)
func (_TssStakingSlashing *TssStakingSlashingSession) IsEqual(byteListA []byte, byteListB []byte) (bool, error) {
	return _TssStakingSlashing.Contract.IsEqual(&_TssStakingSlashing.CallOpts, byteListA, byteListB)
}

// IsEqual is a free data retrieval call binding the contract method 0x34359808.
//
// Solidity: function isEqual(bytes byteListA, bytes byteListB) pure returns(bool)
func (_TssStakingSlashing *TssStakingSlashingCallerSession) IsEqual(byteListA []byte, byteListB []byte) (bool, error) {
	return _TssStakingSlashing.Contract.IsEqual(&_TssStakingSlashing.CallOpts, byteListA, byteListB)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssStakingSlashing *TssStakingSlashingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssStakingSlashing *TssStakingSlashingSession) Owner() (common.Address, error) {
	return _TssStakingSlashing.Contract.Owner(&_TssStakingSlashing.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssStakingSlashing *TssStakingSlashingCallerSession) Owner() (common.Address, error) {
	return _TssStakingSlashing.Contract.Owner(&_TssStakingSlashing.CallOpts)
}

// QuitRequestList is a free data retrieval call binding the contract method 0x57b0f05f.
//
// Solidity: function quitRequestList(uint256 ) view returns(address)
func (_TssStakingSlashing *TssStakingSlashingCaller) QuitRequestList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "quitRequestList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuitRequestList is a free data retrieval call binding the contract method 0x57b0f05f.
//
// Solidity: function quitRequestList(uint256 ) view returns(address)
func (_TssStakingSlashing *TssStakingSlashingSession) QuitRequestList(arg0 *big.Int) (common.Address, error) {
	return _TssStakingSlashing.Contract.QuitRequestList(&_TssStakingSlashing.CallOpts, arg0)
}

// QuitRequestList is a free data retrieval call binding the contract method 0x57b0f05f.
//
// Solidity: function quitRequestList(uint256 ) view returns(address)
func (_TssStakingSlashing *TssStakingSlashingCallerSession) QuitRequestList(arg0 *big.Int) (common.Address, error) {
	return _TssStakingSlashing.Contract.QuitRequestList(&_TssStakingSlashing.CallOpts, arg0)
}

// SlashAmount is a free data retrieval call binding the contract method 0xf2bd7400.
//
// Solidity: function slashAmount(uint256 ) view returns(uint256)
func (_TssStakingSlashing *TssStakingSlashingCaller) SlashAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "slashAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashAmount is a free data retrieval call binding the contract method 0xf2bd7400.
//
// Solidity: function slashAmount(uint256 ) view returns(uint256)
func (_TssStakingSlashing *TssStakingSlashingSession) SlashAmount(arg0 *big.Int) (*big.Int, error) {
	return _TssStakingSlashing.Contract.SlashAmount(&_TssStakingSlashing.CallOpts, arg0)
}

// SlashAmount is a free data retrieval call binding the contract method 0xf2bd7400.
//
// Solidity: function slashAmount(uint256 ) view returns(uint256)
func (_TssStakingSlashing *TssStakingSlashingCallerSession) SlashAmount(arg0 *big.Int) (*big.Int, error) {
	return _TssStakingSlashing.Contract.SlashAmount(&_TssStakingSlashing.CallOpts, arg0)
}

// TssGroupContract is a free data retrieval call binding the contract method 0xd323041d.
//
// Solidity: function tssGroupContract() view returns(address)
func (_TssStakingSlashing *TssStakingSlashingCaller) TssGroupContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssStakingSlashing.contract.Call(opts, &out, "tssGroupContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssGroupContract is a free data retrieval call binding the contract method 0xd323041d.
//
// Solidity: function tssGroupContract() view returns(address)
func (_TssStakingSlashing *TssStakingSlashingSession) TssGroupContract() (common.Address, error) {
	return _TssStakingSlashing.Contract.TssGroupContract(&_TssStakingSlashing.CallOpts)
}

// TssGroupContract is a free data retrieval call binding the contract method 0xd323041d.
//
// Solidity: function tssGroupContract() view returns(address)
func (_TssStakingSlashing *TssStakingSlashingCallerSession) TssGroupContract() (common.Address, error) {
	return _TssStakingSlashing.Contract.TssGroupContract(&_TssStakingSlashing.CallOpts)
}

// ClearQuitRequestList is a paid mutator transaction binding the contract method 0x111b8c02.
//
// Solidity: function clearQuitRequestList() returns()
func (_TssStakingSlashing *TssStakingSlashingTransactor) ClearQuitRequestList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "clearQuitRequestList")
}

// ClearQuitRequestList is a paid mutator transaction binding the contract method 0x111b8c02.
//
// Solidity: function clearQuitRequestList() returns()
func (_TssStakingSlashing *TssStakingSlashingSession) ClearQuitRequestList() (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.ClearQuitRequestList(&_TssStakingSlashing.TransactOpts)
}

// ClearQuitRequestList is a paid mutator transaction binding the contract method 0x111b8c02.
//
// Solidity: function clearQuitRequestList() returns()
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) ClearQuitRequestList() (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.ClearQuitRequestList(&_TssStakingSlashing.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bitToken, address _tssGroupContract) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactor) Initialize(opts *bind.TransactOpts, _bitToken common.Address, _tssGroupContract common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "initialize", _bitToken, _tssGroupContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bitToken, address _tssGroupContract) returns()
func (_TssStakingSlashing *TssStakingSlashingSession) Initialize(_bitToken common.Address, _tssGroupContract common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.Initialize(&_TssStakingSlashing.TransactOpts, _bitToken, _tssGroupContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bitToken, address _tssGroupContract) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) Initialize(_bitToken common.Address, _tssGroupContract common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.Initialize(&_TssStakingSlashing.TransactOpts, _bitToken, _tssGroupContract)
}

// IsJailed is a paid mutator transaction binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address user) returns(bool)
func (_TssStakingSlashing *TssStakingSlashingTransactor) IsJailed(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "isJailed", user)
}

// IsJailed is a paid mutator transaction binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address user) returns(bool)
func (_TssStakingSlashing *TssStakingSlashingSession) IsJailed(user common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.IsJailed(&_TssStakingSlashing.TransactOpts, user)
}

// IsJailed is a paid mutator transaction binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address user) returns(bool)
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) IsJailed(user common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.IsJailed(&_TssStakingSlashing.TransactOpts, user)
}

// QuitRequest is a paid mutator transaction binding the contract method 0x740efec3.
//
// Solidity: function quitRequest() returns()
func (_TssStakingSlashing *TssStakingSlashingTransactor) QuitRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "quitRequest")
}

// QuitRequest is a paid mutator transaction binding the contract method 0x740efec3.
//
// Solidity: function quitRequest() returns()
func (_TssStakingSlashing *TssStakingSlashingSession) QuitRequest() (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.QuitRequest(&_TssStakingSlashing.TransactOpts)
}

// QuitRequest is a paid mutator transaction binding the contract method 0x740efec3.
//
// Solidity: function quitRequest() returns()
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) QuitRequest() (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.QuitRequest(&_TssStakingSlashing.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TssStakingSlashing *TssStakingSlashingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TssStakingSlashing *TssStakingSlashingSession) RenounceOwnership() (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.RenounceOwnership(&_TssStakingSlashing.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.RenounceOwnership(&_TssStakingSlashing.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address _token, address _tssGroup) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactor) SetAddress(opts *bind.TransactOpts, _token common.Address, _tssGroup common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "setAddress", _token, _tssGroup)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address _token, address _tssGroup) returns()
func (_TssStakingSlashing *TssStakingSlashingSession) SetAddress(_token common.Address, _tssGroup common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.SetAddress(&_TssStakingSlashing.TransactOpts, _token, _tssGroup)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address _token, address _tssGroup) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) SetAddress(_token common.Address, _tssGroup common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.SetAddress(&_TssStakingSlashing.TransactOpts, _token, _tssGroup)
}

// SetSlashingParams is a paid mutator transaction binding the contract method 0xde6fb88a.
//
// Solidity: function setSlashingParams(uint256[2] _slashAmount, uint256[2] _exIncome) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactor) SetSlashingParams(opts *bind.TransactOpts, _slashAmount [2]*big.Int, _exIncome [2]*big.Int) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "setSlashingParams", _slashAmount, _exIncome)
}

// SetSlashingParams is a paid mutator transaction binding the contract method 0xde6fb88a.
//
// Solidity: function setSlashingParams(uint256[2] _slashAmount, uint256[2] _exIncome) returns()
func (_TssStakingSlashing *TssStakingSlashingSession) SetSlashingParams(_slashAmount [2]*big.Int, _exIncome [2]*big.Int) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.SetSlashingParams(&_TssStakingSlashing.TransactOpts, _slashAmount, _exIncome)
}

// SetSlashingParams is a paid mutator transaction binding the contract method 0xde6fb88a.
//
// Solidity: function setSlashingParams(uint256[2] _slashAmount, uint256[2] _exIncome) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) SetSlashingParams(_slashAmount [2]*big.Int, _exIncome [2]*big.Int) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.SetSlashingParams(&_TssStakingSlashing.TransactOpts, _slashAmount, _exIncome)
}

// Slashing is a paid mutator transaction binding the contract method 0x5887c33c.
//
// Solidity: function slashing(bytes _messageBytes, bytes _sig) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactor) Slashing(opts *bind.TransactOpts, _messageBytes []byte, _sig []byte) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "slashing", _messageBytes, _sig)
}

// Slashing is a paid mutator transaction binding the contract method 0x5887c33c.
//
// Solidity: function slashing(bytes _messageBytes, bytes _sig) returns()
func (_TssStakingSlashing *TssStakingSlashingSession) Slashing(_messageBytes []byte, _sig []byte) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.Slashing(&_TssStakingSlashing.TransactOpts, _messageBytes, _sig)
}

// Slashing is a paid mutator transaction binding the contract method 0x5887c33c.
//
// Solidity: function slashing(bytes _messageBytes, bytes _sig) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) Slashing(_messageBytes []byte, _sig []byte) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.Slashing(&_TssStakingSlashing.TransactOpts, _messageBytes, _sig)
}

// Staking is a paid mutator transaction binding the contract method 0x5df6db49.
//
// Solidity: function staking(uint256 _amount, bytes _pubKey) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactor) Staking(opts *bind.TransactOpts, _amount *big.Int, _pubKey []byte) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "staking", _amount, _pubKey)
}

// Staking is a paid mutator transaction binding the contract method 0x5df6db49.
//
// Solidity: function staking(uint256 _amount, bytes _pubKey) returns()
func (_TssStakingSlashing *TssStakingSlashingSession) Staking(_amount *big.Int, _pubKey []byte) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.Staking(&_TssStakingSlashing.TransactOpts, _amount, _pubKey)
}

// Staking is a paid mutator transaction binding the contract method 0x5df6db49.
//
// Solidity: function staking(uint256 _amount, bytes _pubKey) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) Staking(_amount *big.Int, _pubKey []byte) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.Staking(&_TssStakingSlashing.TransactOpts, _amount, _pubKey)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TssStakingSlashing *TssStakingSlashingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.TransferOwnership(&_TssStakingSlashing.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.TransferOwnership(&_TssStakingSlashing.TransactOpts, newOwner)
}

// UnJail is a paid mutator transaction binding the contract method 0x6eae5b11.
//
// Solidity: function unJail() returns()
func (_TssStakingSlashing *TssStakingSlashingTransactor) UnJail(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "unJail")
}

// UnJail is a paid mutator transaction binding the contract method 0x6eae5b11.
//
// Solidity: function unJail() returns()
func (_TssStakingSlashing *TssStakingSlashingSession) UnJail() (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.UnJail(&_TssStakingSlashing.TransactOpts)
}

// UnJail is a paid mutator transaction binding the contract method 0x6eae5b11.
//
// Solidity: function unJail() returns()
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) UnJail() (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.UnJail(&_TssStakingSlashing.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_TssStakingSlashing *TssStakingSlashingTransactor) WithdrawToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssStakingSlashing.contract.Transact(opts, "withdrawToken")
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_TssStakingSlashing *TssStakingSlashingSession) WithdrawToken() (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.WithdrawToken(&_TssStakingSlashing.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_TssStakingSlashing *TssStakingSlashingTransactorSession) WithdrawToken() (*types.Transaction, error) {
	return _TssStakingSlashing.Contract.WithdrawToken(&_TssStakingSlashing.TransactOpts)
}

// TssStakingSlashingAddDepositIterator is returned from FilterAddDeposit and is used to iterate over the raw logs and unpacked data for AddDeposit events raised by the TssStakingSlashing contract.
type TssStakingSlashingAddDepositIterator struct {
	Event *TssStakingSlashingAddDeposit // Event containing the contract specifics and raw log

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
func (it *TssStakingSlashingAddDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssStakingSlashingAddDeposit)
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
		it.Event = new(TssStakingSlashingAddDeposit)
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
func (it *TssStakingSlashingAddDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssStakingSlashingAddDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssStakingSlashingAddDeposit represents a AddDeposit event raised by the TssStakingSlashing contract.
type TssStakingSlashingAddDeposit struct {
	Arg0 common.Address
	Arg1 IStakingSlashingDepositInfo
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddDeposit is a free log retrieval operation binding the contract event 0x3183ac23cb0f1e36109abc28e6807d6282c166cb0a538cb5617024770fc73835.
//
// Solidity: event AddDeposit(address arg0, (address,bytes,uint256) arg1)
func (_TssStakingSlashing *TssStakingSlashingFilterer) FilterAddDeposit(opts *bind.FilterOpts) (*TssStakingSlashingAddDepositIterator, error) {

	logs, sub, err := _TssStakingSlashing.contract.FilterLogs(opts, "AddDeposit")
	if err != nil {
		return nil, err
	}
	return &TssStakingSlashingAddDepositIterator{contract: _TssStakingSlashing.contract, event: "AddDeposit", logs: logs, sub: sub}, nil
}

// WatchAddDeposit is a free log subscription operation binding the contract event 0x3183ac23cb0f1e36109abc28e6807d6282c166cb0a538cb5617024770fc73835.
//
// Solidity: event AddDeposit(address arg0, (address,bytes,uint256) arg1)
func (_TssStakingSlashing *TssStakingSlashingFilterer) WatchAddDeposit(opts *bind.WatchOpts, sink chan<- *TssStakingSlashingAddDeposit) (event.Subscription, error) {

	logs, sub, err := _TssStakingSlashing.contract.WatchLogs(opts, "AddDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssStakingSlashingAddDeposit)
				if err := _TssStakingSlashing.contract.UnpackLog(event, "AddDeposit", log); err != nil {
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

// ParseAddDeposit is a log parse operation binding the contract event 0x3183ac23cb0f1e36109abc28e6807d6282c166cb0a538cb5617024770fc73835.
//
// Solidity: event AddDeposit(address arg0, (address,bytes,uint256) arg1)
func (_TssStakingSlashing *TssStakingSlashingFilterer) ParseAddDeposit(log types.Log) (*TssStakingSlashingAddDeposit, error) {
	event := new(TssStakingSlashingAddDeposit)
	if err := _TssStakingSlashing.contract.UnpackLog(event, "AddDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TssStakingSlashingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TssStakingSlashing contract.
type TssStakingSlashingInitializedIterator struct {
	Event *TssStakingSlashingInitialized // Event containing the contract specifics and raw log

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
func (it *TssStakingSlashingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssStakingSlashingInitialized)
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
		it.Event = new(TssStakingSlashingInitialized)
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
func (it *TssStakingSlashingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssStakingSlashingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssStakingSlashingInitialized represents a Initialized event raised by the TssStakingSlashing contract.
type TssStakingSlashingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TssStakingSlashing *TssStakingSlashingFilterer) FilterInitialized(opts *bind.FilterOpts) (*TssStakingSlashingInitializedIterator, error) {

	logs, sub, err := _TssStakingSlashing.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TssStakingSlashingInitializedIterator{contract: _TssStakingSlashing.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TssStakingSlashing *TssStakingSlashingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TssStakingSlashingInitialized) (event.Subscription, error) {

	logs, sub, err := _TssStakingSlashing.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssStakingSlashingInitialized)
				if err := _TssStakingSlashing.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TssStakingSlashing *TssStakingSlashingFilterer) ParseInitialized(log types.Log) (*TssStakingSlashingInitialized, error) {
	event := new(TssStakingSlashingInitialized)
	if err := _TssStakingSlashing.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TssStakingSlashingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TssStakingSlashing contract.
type TssStakingSlashingOwnershipTransferredIterator struct {
	Event *TssStakingSlashingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TssStakingSlashingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssStakingSlashingOwnershipTransferred)
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
		it.Event = new(TssStakingSlashingOwnershipTransferred)
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
func (it *TssStakingSlashingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssStakingSlashingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssStakingSlashingOwnershipTransferred represents a OwnershipTransferred event raised by the TssStakingSlashing contract.
type TssStakingSlashingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TssStakingSlashing *TssStakingSlashingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TssStakingSlashingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TssStakingSlashing.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TssStakingSlashingOwnershipTransferredIterator{contract: _TssStakingSlashing.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TssStakingSlashing *TssStakingSlashingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TssStakingSlashingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TssStakingSlashing.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssStakingSlashingOwnershipTransferred)
				if err := _TssStakingSlashing.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TssStakingSlashing *TssStakingSlashingFilterer) ParseOwnershipTransferred(log types.Log) (*TssStakingSlashingOwnershipTransferred, error) {
	event := new(TssStakingSlashingOwnershipTransferred)
	if err := _TssStakingSlashing.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TssStakingSlashingSlashingIterator is returned from FilterSlashing and is used to iterate over the raw logs and unpacked data for Slashing events raised by the TssStakingSlashing contract.
type TssStakingSlashingSlashingIterator struct {
	Event *TssStakingSlashingSlashing // Event containing the contract specifics and raw log

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
func (it *TssStakingSlashingSlashingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssStakingSlashingSlashing)
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
		it.Event = new(TssStakingSlashingSlashing)
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
func (it *TssStakingSlashingSlashingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssStakingSlashingSlashingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssStakingSlashingSlashing represents a Slashing event raised by the TssStakingSlashing contract.
type TssStakingSlashingSlashing struct {
	Arg0 common.Address
	Arg1 uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSlashing is a free log retrieval operation binding the contract event 0x9453459a6e9fa2069f1490c99cec15646afa157300f218a7f7840b23d09dbd3e.
//
// Solidity: event Slashing(address arg0, uint8 arg1)
func (_TssStakingSlashing *TssStakingSlashingFilterer) FilterSlashing(opts *bind.FilterOpts) (*TssStakingSlashingSlashingIterator, error) {

	logs, sub, err := _TssStakingSlashing.contract.FilterLogs(opts, "Slashing")
	if err != nil {
		return nil, err
	}
	return &TssStakingSlashingSlashingIterator{contract: _TssStakingSlashing.contract, event: "Slashing", logs: logs, sub: sub}, nil
}

// WatchSlashing is a free log subscription operation binding the contract event 0x9453459a6e9fa2069f1490c99cec15646afa157300f218a7f7840b23d09dbd3e.
//
// Solidity: event Slashing(address arg0, uint8 arg1)
func (_TssStakingSlashing *TssStakingSlashingFilterer) WatchSlashing(opts *bind.WatchOpts, sink chan<- *TssStakingSlashingSlashing) (event.Subscription, error) {

	logs, sub, err := _TssStakingSlashing.contract.WatchLogs(opts, "Slashing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssStakingSlashingSlashing)
				if err := _TssStakingSlashing.contract.UnpackLog(event, "Slashing", log); err != nil {
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

// ParseSlashing is a log parse operation binding the contract event 0x9453459a6e9fa2069f1490c99cec15646afa157300f218a7f7840b23d09dbd3e.
//
// Solidity: event Slashing(address arg0, uint8 arg1)
func (_TssStakingSlashing *TssStakingSlashingFilterer) ParseSlashing(log types.Log) (*TssStakingSlashingSlashing, error) {
	event := new(TssStakingSlashingSlashing)
	if err := _TssStakingSlashing.contract.UnpackLog(event, "Slashing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TssStakingSlashingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the TssStakingSlashing contract.
type TssStakingSlashingWithdrawIterator struct {
	Event *TssStakingSlashingWithdraw // Event containing the contract specifics and raw log

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
func (it *TssStakingSlashingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssStakingSlashingWithdraw)
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
		it.Event = new(TssStakingSlashingWithdraw)
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
func (it *TssStakingSlashingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssStakingSlashingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssStakingSlashingWithdraw represents a Withdraw event raised by the TssStakingSlashing contract.
type TssStakingSlashingWithdraw struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address arg0, uint256 arg1)
func (_TssStakingSlashing *TssStakingSlashingFilterer) FilterWithdraw(opts *bind.FilterOpts) (*TssStakingSlashingWithdrawIterator, error) {

	logs, sub, err := _TssStakingSlashing.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &TssStakingSlashingWithdrawIterator{contract: _TssStakingSlashing.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address arg0, uint256 arg1)
func (_TssStakingSlashing *TssStakingSlashingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TssStakingSlashingWithdraw) (event.Subscription, error) {

	logs, sub, err := _TssStakingSlashing.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssStakingSlashingWithdraw)
				if err := _TssStakingSlashing.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address arg0, uint256 arg1)
func (_TssStakingSlashing *TssStakingSlashingFilterer) ParseWithdraw(log types.Log) (*TssStakingSlashingWithdraw, error) {
	event := new(TssStakingSlashingWithdraw)
	if err := _TssStakingSlashing.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

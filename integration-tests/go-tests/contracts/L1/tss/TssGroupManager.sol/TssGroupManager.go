// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TssGroupManager

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

// ITssGroupManagerTssMember is an auto generated low-level Go binding around an user-defined struct.
type ITssGroupManagerTssMember struct {
	PublicKey   []byte
	NodeAddress common.Address
	Status      uint8
}

// TssGroupManagerMetaData contains all meta data concerning the TssGroupManager contract.
var TssGroupManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_groupKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"activeTssMembers\",\"type\":\"bytes[]\"}],\"name\":\"tssActiveMemberAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"_inActiveTssMembers\",\"type\":\"bytes[]\"}],\"name\":\"tssGroupMemberAppend\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getTssGroupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTssGroupMembers\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTssGroupUnJailMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTssInactiveGroupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"getTssMember\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"enumITssGroupManager.MemberStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structITssGroupManager.TssMember\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inActiveIsEmpty\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"byteListA\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"byteListB\",\"type\":\"bytes\"}],\"name\":\"isEqual\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isInActiveMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"memberExistActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"memberExistInActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"memberJail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"memberUnJail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"publicKeyToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"removeMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_groupPublicKey\",\"type\":\"bytes\"}],\"name\":\"setGroupPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setStakingSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_batchPublicKey\",\"type\":\"bytes[]\"}],\"name\":\"setTssGroupMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingSlash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tssActiveMemberInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"enumITssGroupManager.MemberStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"verifySign\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TssGroupManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TssGroupManagerMetaData.ABI instead.
var TssGroupManagerABI = TssGroupManagerMetaData.ABI

// TssGroupManager is an auto generated Go binding around an Ethereum contract.
type TssGroupManager struct {
	TssGroupManagerCaller     // Read-only binding to the contract
	TssGroupManagerTransactor // Write-only binding to the contract
	TssGroupManagerFilterer   // Log filterer for contract events
}

// TssGroupManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TssGroupManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssGroupManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TssGroupManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssGroupManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TssGroupManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssGroupManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TssGroupManagerSession struct {
	Contract     *TssGroupManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TssGroupManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TssGroupManagerCallerSession struct {
	Contract *TssGroupManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TssGroupManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TssGroupManagerTransactorSession struct {
	Contract     *TssGroupManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TssGroupManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TssGroupManagerRaw struct {
	Contract *TssGroupManager // Generic contract binding to access the raw methods on
}

// TssGroupManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TssGroupManagerCallerRaw struct {
	Contract *TssGroupManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TssGroupManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TssGroupManagerTransactorRaw struct {
	Contract *TssGroupManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTssGroupManager creates a new instance of TssGroupManager, bound to a specific deployed contract.
func NewTssGroupManager(address common.Address, backend bind.ContractBackend) (*TssGroupManager, error) {
	contract, err := bindTssGroupManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TssGroupManager{TssGroupManagerCaller: TssGroupManagerCaller{contract: contract}, TssGroupManagerTransactor: TssGroupManagerTransactor{contract: contract}, TssGroupManagerFilterer: TssGroupManagerFilterer{contract: contract}}, nil
}

// NewTssGroupManagerCaller creates a new read-only instance of TssGroupManager, bound to a specific deployed contract.
func NewTssGroupManagerCaller(address common.Address, caller bind.ContractCaller) (*TssGroupManagerCaller, error) {
	contract, err := bindTssGroupManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TssGroupManagerCaller{contract: contract}, nil
}

// NewTssGroupManagerTransactor creates a new write-only instance of TssGroupManager, bound to a specific deployed contract.
func NewTssGroupManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TssGroupManagerTransactor, error) {
	contract, err := bindTssGroupManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TssGroupManagerTransactor{contract: contract}, nil
}

// NewTssGroupManagerFilterer creates a new log filterer instance of TssGroupManager, bound to a specific deployed contract.
func NewTssGroupManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TssGroupManagerFilterer, error) {
	contract, err := bindTssGroupManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TssGroupManagerFilterer{contract: contract}, nil
}

// bindTssGroupManager binds a generic wrapper to an already deployed contract.
func bindTssGroupManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TssGroupManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TssGroupManager *TssGroupManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TssGroupManager.Contract.TssGroupManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TssGroupManager *TssGroupManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssGroupManager.Contract.TssGroupManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TssGroupManager *TssGroupManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TssGroupManager.Contract.TssGroupManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TssGroupManager *TssGroupManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TssGroupManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TssGroupManager *TssGroupManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssGroupManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TssGroupManager *TssGroupManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TssGroupManager.Contract.contract.Transact(opts, method, params...)
}

// GetTssGroupInfo is a free data retrieval call binding the contract method 0xfcca5592.
//
// Solidity: function getTssGroupInfo() view returns(uint256, uint256, bytes, bytes[])
func (_TssGroupManager *TssGroupManagerCaller) GetTssGroupInfo(opts *bind.CallOpts) (*big.Int, *big.Int, []byte, [][]byte, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "getTssGroupInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new([]byte), *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	out3 := *abi.ConvertType(out[3], new([][]byte)).(*[][]byte)

	return out0, out1, out2, out3, err

}

// GetTssGroupInfo is a free data retrieval call binding the contract method 0xfcca5592.
//
// Solidity: function getTssGroupInfo() view returns(uint256, uint256, bytes, bytes[])
func (_TssGroupManager *TssGroupManagerSession) GetTssGroupInfo() (*big.Int, *big.Int, []byte, [][]byte, error) {
	return _TssGroupManager.Contract.GetTssGroupInfo(&_TssGroupManager.CallOpts)
}

// GetTssGroupInfo is a free data retrieval call binding the contract method 0xfcca5592.
//
// Solidity: function getTssGroupInfo() view returns(uint256, uint256, bytes, bytes[])
func (_TssGroupManager *TssGroupManagerCallerSession) GetTssGroupInfo() (*big.Int, *big.Int, []byte, [][]byte, error) {
	return _TssGroupManager.Contract.GetTssGroupInfo(&_TssGroupManager.CallOpts)
}

// GetTssGroupMembers is a free data retrieval call binding the contract method 0x95f645bc.
//
// Solidity: function getTssGroupMembers() view returns(bytes[])
func (_TssGroupManager *TssGroupManagerCaller) GetTssGroupMembers(opts *bind.CallOpts) ([][]byte, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "getTssGroupMembers")

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetTssGroupMembers is a free data retrieval call binding the contract method 0x95f645bc.
//
// Solidity: function getTssGroupMembers() view returns(bytes[])
func (_TssGroupManager *TssGroupManagerSession) GetTssGroupMembers() ([][]byte, error) {
	return _TssGroupManager.Contract.GetTssGroupMembers(&_TssGroupManager.CallOpts)
}

// GetTssGroupMembers is a free data retrieval call binding the contract method 0x95f645bc.
//
// Solidity: function getTssGroupMembers() view returns(bytes[])
func (_TssGroupManager *TssGroupManagerCallerSession) GetTssGroupMembers() ([][]byte, error) {
	return _TssGroupManager.Contract.GetTssGroupMembers(&_TssGroupManager.CallOpts)
}

// GetTssGroupUnJailMembers is a free data retrieval call binding the contract method 0x2cd00d53.
//
// Solidity: function getTssGroupUnJailMembers() view returns(address[])
func (_TssGroupManager *TssGroupManagerCaller) GetTssGroupUnJailMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "getTssGroupUnJailMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTssGroupUnJailMembers is a free data retrieval call binding the contract method 0x2cd00d53.
//
// Solidity: function getTssGroupUnJailMembers() view returns(address[])
func (_TssGroupManager *TssGroupManagerSession) GetTssGroupUnJailMembers() ([]common.Address, error) {
	return _TssGroupManager.Contract.GetTssGroupUnJailMembers(&_TssGroupManager.CallOpts)
}

// GetTssGroupUnJailMembers is a free data retrieval call binding the contract method 0x2cd00d53.
//
// Solidity: function getTssGroupUnJailMembers() view returns(address[])
func (_TssGroupManager *TssGroupManagerCallerSession) GetTssGroupUnJailMembers() ([]common.Address, error) {
	return _TssGroupManager.Contract.GetTssGroupUnJailMembers(&_TssGroupManager.CallOpts)
}

// GetTssInactiveGroupInfo is a free data retrieval call binding the contract method 0x404a6986.
//
// Solidity: function getTssInactiveGroupInfo() view returns(uint256, uint256, bytes[])
func (_TssGroupManager *TssGroupManagerCaller) GetTssInactiveGroupInfo(opts *bind.CallOpts) (*big.Int, *big.Int, [][]byte, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "getTssInactiveGroupInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new([][]byte)).(*[][]byte)

	return out0, out1, out2, err

}

// GetTssInactiveGroupInfo is a free data retrieval call binding the contract method 0x404a6986.
//
// Solidity: function getTssInactiveGroupInfo() view returns(uint256, uint256, bytes[])
func (_TssGroupManager *TssGroupManagerSession) GetTssInactiveGroupInfo() (*big.Int, *big.Int, [][]byte, error) {
	return _TssGroupManager.Contract.GetTssInactiveGroupInfo(&_TssGroupManager.CallOpts)
}

// GetTssInactiveGroupInfo is a free data retrieval call binding the contract method 0x404a6986.
//
// Solidity: function getTssInactiveGroupInfo() view returns(uint256, uint256, bytes[])
func (_TssGroupManager *TssGroupManagerCallerSession) GetTssInactiveGroupInfo() (*big.Int, *big.Int, [][]byte, error) {
	return _TssGroupManager.Contract.GetTssInactiveGroupInfo(&_TssGroupManager.CallOpts)
}

// GetTssMember is a free data retrieval call binding the contract method 0x54bbb914.
//
// Solidity: function getTssMember(bytes _publicKey) view returns((bytes,address,uint8))
func (_TssGroupManager *TssGroupManagerCaller) GetTssMember(opts *bind.CallOpts, _publicKey []byte) (ITssGroupManagerTssMember, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "getTssMember", _publicKey)

	if err != nil {
		return *new(ITssGroupManagerTssMember), err
	}

	out0 := *abi.ConvertType(out[0], new(ITssGroupManagerTssMember)).(*ITssGroupManagerTssMember)

	return out0, err

}

// GetTssMember is a free data retrieval call binding the contract method 0x54bbb914.
//
// Solidity: function getTssMember(bytes _publicKey) view returns((bytes,address,uint8))
func (_TssGroupManager *TssGroupManagerSession) GetTssMember(_publicKey []byte) (ITssGroupManagerTssMember, error) {
	return _TssGroupManager.Contract.GetTssMember(&_TssGroupManager.CallOpts, _publicKey)
}

// GetTssMember is a free data retrieval call binding the contract method 0x54bbb914.
//
// Solidity: function getTssMember(bytes _publicKey) view returns((bytes,address,uint8))
func (_TssGroupManager *TssGroupManagerCallerSession) GetTssMember(_publicKey []byte) (ITssGroupManagerTssMember, error) {
	return _TssGroupManager.Contract.GetTssMember(&_TssGroupManager.CallOpts, _publicKey)
}

// InActiveIsEmpty is a free data retrieval call binding the contract method 0x8a51d8b4.
//
// Solidity: function inActiveIsEmpty() view returns(bool)
func (_TssGroupManager *TssGroupManagerCaller) InActiveIsEmpty(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "inActiveIsEmpty")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InActiveIsEmpty is a free data retrieval call binding the contract method 0x8a51d8b4.
//
// Solidity: function inActiveIsEmpty() view returns(bool)
func (_TssGroupManager *TssGroupManagerSession) InActiveIsEmpty() (bool, error) {
	return _TssGroupManager.Contract.InActiveIsEmpty(&_TssGroupManager.CallOpts)
}

// InActiveIsEmpty is a free data retrieval call binding the contract method 0x8a51d8b4.
//
// Solidity: function inActiveIsEmpty() view returns(bool)
func (_TssGroupManager *TssGroupManagerCallerSession) InActiveIsEmpty() (bool, error) {
	return _TssGroupManager.Contract.InActiveIsEmpty(&_TssGroupManager.CallOpts)
}

// IsEqual is a free data retrieval call binding the contract method 0x34359808.
//
// Solidity: function isEqual(bytes byteListA, bytes byteListB) pure returns(bool)
func (_TssGroupManager *TssGroupManagerCaller) IsEqual(opts *bind.CallOpts, byteListA []byte, byteListB []byte) (bool, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "isEqual", byteListA, byteListB)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEqual is a free data retrieval call binding the contract method 0x34359808.
//
// Solidity: function isEqual(bytes byteListA, bytes byteListB) pure returns(bool)
func (_TssGroupManager *TssGroupManagerSession) IsEqual(byteListA []byte, byteListB []byte) (bool, error) {
	return _TssGroupManager.Contract.IsEqual(&_TssGroupManager.CallOpts, byteListA, byteListB)
}

// IsEqual is a free data retrieval call binding the contract method 0x34359808.
//
// Solidity: function isEqual(bytes byteListA, bytes byteListB) pure returns(bool)
func (_TssGroupManager *TssGroupManagerCallerSession) IsEqual(byteListA []byte, byteListB []byte) (bool, error) {
	return _TssGroupManager.Contract.IsEqual(&_TssGroupManager.CallOpts, byteListA, byteListB)
}

// IsInActiveMember is a free data retrieval call binding the contract method 0xe46a6368.
//
// Solidity: function isInActiveMember(bytes ) view returns(bool)
func (_TssGroupManager *TssGroupManagerCaller) IsInActiveMember(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "isInActiveMember", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInActiveMember is a free data retrieval call binding the contract method 0xe46a6368.
//
// Solidity: function isInActiveMember(bytes ) view returns(bool)
func (_TssGroupManager *TssGroupManagerSession) IsInActiveMember(arg0 []byte) (bool, error) {
	return _TssGroupManager.Contract.IsInActiveMember(&_TssGroupManager.CallOpts, arg0)
}

// IsInActiveMember is a free data retrieval call binding the contract method 0xe46a6368.
//
// Solidity: function isInActiveMember(bytes ) view returns(bool)
func (_TssGroupManager *TssGroupManagerCallerSession) IsInActiveMember(arg0 []byte) (bool, error) {
	return _TssGroupManager.Contract.IsInActiveMember(&_TssGroupManager.CallOpts, arg0)
}

// MemberExistActive is a free data retrieval call binding the contract method 0xbcbc5fc6.
//
// Solidity: function memberExistActive(bytes _publicKey) view returns(bool)
func (_TssGroupManager *TssGroupManagerCaller) MemberExistActive(opts *bind.CallOpts, _publicKey []byte) (bool, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "memberExistActive", _publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MemberExistActive is a free data retrieval call binding the contract method 0xbcbc5fc6.
//
// Solidity: function memberExistActive(bytes _publicKey) view returns(bool)
func (_TssGroupManager *TssGroupManagerSession) MemberExistActive(_publicKey []byte) (bool, error) {
	return _TssGroupManager.Contract.MemberExistActive(&_TssGroupManager.CallOpts, _publicKey)
}

// MemberExistActive is a free data retrieval call binding the contract method 0xbcbc5fc6.
//
// Solidity: function memberExistActive(bytes _publicKey) view returns(bool)
func (_TssGroupManager *TssGroupManagerCallerSession) MemberExistActive(_publicKey []byte) (bool, error) {
	return _TssGroupManager.Contract.MemberExistActive(&_TssGroupManager.CallOpts, _publicKey)
}

// MemberExistInActive is a free data retrieval call binding the contract method 0x4866e2e0.
//
// Solidity: function memberExistInActive(bytes _publicKey) view returns(bool)
func (_TssGroupManager *TssGroupManagerCaller) MemberExistInActive(opts *bind.CallOpts, _publicKey []byte) (bool, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "memberExistInActive", _publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MemberExistInActive is a free data retrieval call binding the contract method 0x4866e2e0.
//
// Solidity: function memberExistInActive(bytes _publicKey) view returns(bool)
func (_TssGroupManager *TssGroupManagerSession) MemberExistInActive(_publicKey []byte) (bool, error) {
	return _TssGroupManager.Contract.MemberExistInActive(&_TssGroupManager.CallOpts, _publicKey)
}

// MemberExistInActive is a free data retrieval call binding the contract method 0x4866e2e0.
//
// Solidity: function memberExistInActive(bytes _publicKey) view returns(bool)
func (_TssGroupManager *TssGroupManagerCallerSession) MemberExistInActive(_publicKey []byte) (bool, error) {
	return _TssGroupManager.Contract.MemberExistInActive(&_TssGroupManager.CallOpts, _publicKey)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssGroupManager *TssGroupManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssGroupManager *TssGroupManagerSession) Owner() (common.Address, error) {
	return _TssGroupManager.Contract.Owner(&_TssGroupManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssGroupManager *TssGroupManagerCallerSession) Owner() (common.Address, error) {
	return _TssGroupManager.Contract.Owner(&_TssGroupManager.CallOpts)
}

// PublicKeyToAddress is a free data retrieval call binding the contract method 0x43ae656c.
//
// Solidity: function publicKeyToAddress(bytes publicKey) pure returns(address)
func (_TssGroupManager *TssGroupManagerCaller) PublicKeyToAddress(opts *bind.CallOpts, publicKey []byte) (common.Address, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "publicKeyToAddress", publicKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PublicKeyToAddress is a free data retrieval call binding the contract method 0x43ae656c.
//
// Solidity: function publicKeyToAddress(bytes publicKey) pure returns(address)
func (_TssGroupManager *TssGroupManagerSession) PublicKeyToAddress(publicKey []byte) (common.Address, error) {
	return _TssGroupManager.Contract.PublicKeyToAddress(&_TssGroupManager.CallOpts, publicKey)
}

// PublicKeyToAddress is a free data retrieval call binding the contract method 0x43ae656c.
//
// Solidity: function publicKeyToAddress(bytes publicKey) pure returns(address)
func (_TssGroupManager *TssGroupManagerCallerSession) PublicKeyToAddress(publicKey []byte) (common.Address, error) {
	return _TssGroupManager.Contract.PublicKeyToAddress(&_TssGroupManager.CallOpts, publicKey)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _ethSignedMessageHash, bytes _sig) pure returns(address)
func (_TssGroupManager *TssGroupManagerCaller) Recover(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _sig []byte) (common.Address, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "recover", _ethSignedMessageHash, _sig)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _ethSignedMessageHash, bytes _sig) pure returns(address)
func (_TssGroupManager *TssGroupManagerSession) Recover(_ethSignedMessageHash [32]byte, _sig []byte) (common.Address, error) {
	return _TssGroupManager.Contract.Recover(&_TssGroupManager.CallOpts, _ethSignedMessageHash, _sig)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _ethSignedMessageHash, bytes _sig) pure returns(address)
func (_TssGroupManager *TssGroupManagerCallerSession) Recover(_ethSignedMessageHash [32]byte, _sig []byte) (common.Address, error) {
	return _TssGroupManager.Contract.Recover(&_TssGroupManager.CallOpts, _ethSignedMessageHash, _sig)
}

// StakingSlash is a free data retrieval call binding the contract method 0xcf4c8a7b.
//
// Solidity: function stakingSlash() view returns(address)
func (_TssGroupManager *TssGroupManagerCaller) StakingSlash(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "stakingSlash")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingSlash is a free data retrieval call binding the contract method 0xcf4c8a7b.
//
// Solidity: function stakingSlash() view returns(address)
func (_TssGroupManager *TssGroupManagerSession) StakingSlash() (common.Address, error) {
	return _TssGroupManager.Contract.StakingSlash(&_TssGroupManager.CallOpts)
}

// StakingSlash is a free data retrieval call binding the contract method 0xcf4c8a7b.
//
// Solidity: function stakingSlash() view returns(address)
func (_TssGroupManager *TssGroupManagerCallerSession) StakingSlash() (common.Address, error) {
	return _TssGroupManager.Contract.StakingSlash(&_TssGroupManager.CallOpts)
}

// TssActiveMemberInfo is a free data retrieval call binding the contract method 0xb3bd9ac4.
//
// Solidity: function tssActiveMemberInfo(bytes ) view returns(bytes publicKey, address nodeAddress, uint8 status)
func (_TssGroupManager *TssGroupManagerCaller) TssActiveMemberInfo(opts *bind.CallOpts, arg0 []byte) (struct {
	PublicKey   []byte
	NodeAddress common.Address
	Status      uint8
}, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "tssActiveMemberInfo", arg0)

	outstruct := new(struct {
		PublicKey   []byte
		NodeAddress common.Address
		Status      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PublicKey = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.NodeAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// TssActiveMemberInfo is a free data retrieval call binding the contract method 0xb3bd9ac4.
//
// Solidity: function tssActiveMemberInfo(bytes ) view returns(bytes publicKey, address nodeAddress, uint8 status)
func (_TssGroupManager *TssGroupManagerSession) TssActiveMemberInfo(arg0 []byte) (struct {
	PublicKey   []byte
	NodeAddress common.Address
	Status      uint8
}, error) {
	return _TssGroupManager.Contract.TssActiveMemberInfo(&_TssGroupManager.CallOpts, arg0)
}

// TssActiveMemberInfo is a free data retrieval call binding the contract method 0xb3bd9ac4.
//
// Solidity: function tssActiveMemberInfo(bytes ) view returns(bytes publicKey, address nodeAddress, uint8 status)
func (_TssGroupManager *TssGroupManagerCallerSession) TssActiveMemberInfo(arg0 []byte) (struct {
	PublicKey   []byte
	NodeAddress common.Address
	Status      uint8
}, error) {
	return _TssGroupManager.Contract.TssActiveMemberInfo(&_TssGroupManager.CallOpts, arg0)
}

// VerifySign is a free data retrieval call binding the contract method 0x3231a7f0.
//
// Solidity: function verifySign(bytes32 _message, bytes _sig) view returns(bool)
func (_TssGroupManager *TssGroupManagerCaller) VerifySign(opts *bind.CallOpts, _message [32]byte, _sig []byte) (bool, error) {
	var out []interface{}
	err := _TssGroupManager.contract.Call(opts, &out, "verifySign", _message, _sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySign is a free data retrieval call binding the contract method 0x3231a7f0.
//
// Solidity: function verifySign(bytes32 _message, bytes _sig) view returns(bool)
func (_TssGroupManager *TssGroupManagerSession) VerifySign(_message [32]byte, _sig []byte) (bool, error) {
	return _TssGroupManager.Contract.VerifySign(&_TssGroupManager.CallOpts, _message, _sig)
}

// VerifySign is a free data retrieval call binding the contract method 0x3231a7f0.
//
// Solidity: function verifySign(bytes32 _message, bytes _sig) view returns(bool)
func (_TssGroupManager *TssGroupManagerCallerSession) VerifySign(_message [32]byte, _sig []byte) (bool, error) {
	return _TssGroupManager.Contract.VerifySign(&_TssGroupManager.CallOpts, _message, _sig)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TssGroupManager *TssGroupManagerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssGroupManager.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TssGroupManager *TssGroupManagerSession) Initialize() (*types.Transaction, error) {
	return _TssGroupManager.Contract.Initialize(&_TssGroupManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TssGroupManager *TssGroupManagerTransactorSession) Initialize() (*types.Transaction, error) {
	return _TssGroupManager.Contract.Initialize(&_TssGroupManager.TransactOpts)
}

// MemberJail is a paid mutator transaction binding the contract method 0x254ff981.
//
// Solidity: function memberJail(bytes _publicKey) returns()
func (_TssGroupManager *TssGroupManagerTransactor) MemberJail(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.contract.Transact(opts, "memberJail", _publicKey)
}

// MemberJail is a paid mutator transaction binding the contract method 0x254ff981.
//
// Solidity: function memberJail(bytes _publicKey) returns()
func (_TssGroupManager *TssGroupManagerSession) MemberJail(_publicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.Contract.MemberJail(&_TssGroupManager.TransactOpts, _publicKey)
}

// MemberJail is a paid mutator transaction binding the contract method 0x254ff981.
//
// Solidity: function memberJail(bytes _publicKey) returns()
func (_TssGroupManager *TssGroupManagerTransactorSession) MemberJail(_publicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.Contract.MemberJail(&_TssGroupManager.TransactOpts, _publicKey)
}

// MemberUnJail is a paid mutator transaction binding the contract method 0x429ec49c.
//
// Solidity: function memberUnJail(bytes _publicKey) returns()
func (_TssGroupManager *TssGroupManagerTransactor) MemberUnJail(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.contract.Transact(opts, "memberUnJail", _publicKey)
}

// MemberUnJail is a paid mutator transaction binding the contract method 0x429ec49c.
//
// Solidity: function memberUnJail(bytes _publicKey) returns()
func (_TssGroupManager *TssGroupManagerSession) MemberUnJail(_publicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.Contract.MemberUnJail(&_TssGroupManager.TransactOpts, _publicKey)
}

// MemberUnJail is a paid mutator transaction binding the contract method 0x429ec49c.
//
// Solidity: function memberUnJail(bytes _publicKey) returns()
func (_TssGroupManager *TssGroupManagerTransactorSession) MemberUnJail(_publicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.Contract.MemberUnJail(&_TssGroupManager.TransactOpts, _publicKey)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x7a952c11.
//
// Solidity: function removeMember(bytes _publicKey) returns()
func (_TssGroupManager *TssGroupManagerTransactor) RemoveMember(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.contract.Transact(opts, "removeMember", _publicKey)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x7a952c11.
//
// Solidity: function removeMember(bytes _publicKey) returns()
func (_TssGroupManager *TssGroupManagerSession) RemoveMember(_publicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.Contract.RemoveMember(&_TssGroupManager.TransactOpts, _publicKey)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x7a952c11.
//
// Solidity: function removeMember(bytes _publicKey) returns()
func (_TssGroupManager *TssGroupManagerTransactorSession) RemoveMember(_publicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.Contract.RemoveMember(&_TssGroupManager.TransactOpts, _publicKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TssGroupManager *TssGroupManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssGroupManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TssGroupManager *TssGroupManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _TssGroupManager.Contract.RenounceOwnership(&_TssGroupManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TssGroupManager *TssGroupManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TssGroupManager.Contract.RenounceOwnership(&_TssGroupManager.TransactOpts)
}

// SetGroupPublicKey is a paid mutator transaction binding the contract method 0x4679b366.
//
// Solidity: function setGroupPublicKey(bytes _publicKey, bytes _groupPublicKey) returns()
func (_TssGroupManager *TssGroupManagerTransactor) SetGroupPublicKey(opts *bind.TransactOpts, _publicKey []byte, _groupPublicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.contract.Transact(opts, "setGroupPublicKey", _publicKey, _groupPublicKey)
}

// SetGroupPublicKey is a paid mutator transaction binding the contract method 0x4679b366.
//
// Solidity: function setGroupPublicKey(bytes _publicKey, bytes _groupPublicKey) returns()
func (_TssGroupManager *TssGroupManagerSession) SetGroupPublicKey(_publicKey []byte, _groupPublicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.Contract.SetGroupPublicKey(&_TssGroupManager.TransactOpts, _publicKey, _groupPublicKey)
}

// SetGroupPublicKey is a paid mutator transaction binding the contract method 0x4679b366.
//
// Solidity: function setGroupPublicKey(bytes _publicKey, bytes _groupPublicKey) returns()
func (_TssGroupManager *TssGroupManagerTransactorSession) SetGroupPublicKey(_publicKey []byte, _groupPublicKey []byte) (*types.Transaction, error) {
	return _TssGroupManager.Contract.SetGroupPublicKey(&_TssGroupManager.TransactOpts, _publicKey, _groupPublicKey)
}

// SetStakingSlash is a paid mutator transaction binding the contract method 0xfd3eb81f.
//
// Solidity: function setStakingSlash(address _address) returns()
func (_TssGroupManager *TssGroupManagerTransactor) SetStakingSlash(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _TssGroupManager.contract.Transact(opts, "setStakingSlash", _address)
}

// SetStakingSlash is a paid mutator transaction binding the contract method 0xfd3eb81f.
//
// Solidity: function setStakingSlash(address _address) returns()
func (_TssGroupManager *TssGroupManagerSession) SetStakingSlash(_address common.Address) (*types.Transaction, error) {
	return _TssGroupManager.Contract.SetStakingSlash(&_TssGroupManager.TransactOpts, _address)
}

// SetStakingSlash is a paid mutator transaction binding the contract method 0xfd3eb81f.
//
// Solidity: function setStakingSlash(address _address) returns()
func (_TssGroupManager *TssGroupManagerTransactorSession) SetStakingSlash(_address common.Address) (*types.Transaction, error) {
	return _TssGroupManager.Contract.SetStakingSlash(&_TssGroupManager.TransactOpts, _address)
}

// SetTssGroupMember is a paid mutator transaction binding the contract method 0x1a47931b.
//
// Solidity: function setTssGroupMember(uint256 _threshold, bytes[] _batchPublicKey) returns()
func (_TssGroupManager *TssGroupManagerTransactor) SetTssGroupMember(opts *bind.TransactOpts, _threshold *big.Int, _batchPublicKey [][]byte) (*types.Transaction, error) {
	return _TssGroupManager.contract.Transact(opts, "setTssGroupMember", _threshold, _batchPublicKey)
}

// SetTssGroupMember is a paid mutator transaction binding the contract method 0x1a47931b.
//
// Solidity: function setTssGroupMember(uint256 _threshold, bytes[] _batchPublicKey) returns()
func (_TssGroupManager *TssGroupManagerSession) SetTssGroupMember(_threshold *big.Int, _batchPublicKey [][]byte) (*types.Transaction, error) {
	return _TssGroupManager.Contract.SetTssGroupMember(&_TssGroupManager.TransactOpts, _threshold, _batchPublicKey)
}

// SetTssGroupMember is a paid mutator transaction binding the contract method 0x1a47931b.
//
// Solidity: function setTssGroupMember(uint256 _threshold, bytes[] _batchPublicKey) returns()
func (_TssGroupManager *TssGroupManagerTransactorSession) SetTssGroupMember(_threshold *big.Int, _batchPublicKey [][]byte) (*types.Transaction, error) {
	return _TssGroupManager.Contract.SetTssGroupMember(&_TssGroupManager.TransactOpts, _threshold, _batchPublicKey)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TssGroupManager *TssGroupManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TssGroupManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TssGroupManager *TssGroupManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TssGroupManager.Contract.TransferOwnership(&_TssGroupManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TssGroupManager *TssGroupManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TssGroupManager.Contract.TransferOwnership(&_TssGroupManager.TransactOpts, newOwner)
}

// TssGroupManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TssGroupManager contract.
type TssGroupManagerInitializedIterator struct {
	Event *TssGroupManagerInitialized // Event containing the contract specifics and raw log

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
func (it *TssGroupManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssGroupManagerInitialized)
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
		it.Event = new(TssGroupManagerInitialized)
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
func (it *TssGroupManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssGroupManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssGroupManagerInitialized represents a Initialized event raised by the TssGroupManager contract.
type TssGroupManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TssGroupManager *TssGroupManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*TssGroupManagerInitializedIterator, error) {

	logs, sub, err := _TssGroupManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TssGroupManagerInitializedIterator{contract: _TssGroupManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TssGroupManager *TssGroupManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TssGroupManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _TssGroupManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssGroupManagerInitialized)
				if err := _TssGroupManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TssGroupManager *TssGroupManagerFilterer) ParseInitialized(log types.Log) (*TssGroupManagerInitialized, error) {
	event := new(TssGroupManagerInitialized)
	if err := _TssGroupManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TssGroupManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TssGroupManager contract.
type TssGroupManagerOwnershipTransferredIterator struct {
	Event *TssGroupManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TssGroupManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssGroupManagerOwnershipTransferred)
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
		it.Event = new(TssGroupManagerOwnershipTransferred)
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
func (it *TssGroupManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssGroupManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssGroupManagerOwnershipTransferred represents a OwnershipTransferred event raised by the TssGroupManager contract.
type TssGroupManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TssGroupManager *TssGroupManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TssGroupManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TssGroupManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TssGroupManagerOwnershipTransferredIterator{contract: _TssGroupManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TssGroupManager *TssGroupManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TssGroupManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TssGroupManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssGroupManagerOwnershipTransferred)
				if err := _TssGroupManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TssGroupManager *TssGroupManagerFilterer) ParseOwnershipTransferred(log types.Log) (*TssGroupManagerOwnershipTransferred, error) {
	event := new(TssGroupManagerOwnershipTransferred)
	if err := _TssGroupManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TssGroupManagerTssActiveMemberAppendedIterator is returned from FilterTssActiveMemberAppended and is used to iterate over the raw logs and unpacked data for TssActiveMemberAppended events raised by the TssGroupManager contract.
type TssGroupManagerTssActiveMemberAppendedIterator struct {
	Event *TssGroupManagerTssActiveMemberAppended // Event containing the contract specifics and raw log

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
func (it *TssGroupManagerTssActiveMemberAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssGroupManagerTssActiveMemberAppended)
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
		it.Event = new(TssGroupManagerTssActiveMemberAppended)
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
func (it *TssGroupManagerTssActiveMemberAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssGroupManagerTssActiveMemberAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssGroupManagerTssActiveMemberAppended represents a TssActiveMemberAppended event raised by the TssGroupManager contract.
type TssGroupManagerTssActiveMemberAppended struct {
	RoundId          *big.Int
	GroupKey         []byte
	ActiveTssMembers [][]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTssActiveMemberAppended is a free log retrieval operation binding the contract event 0x61c7922e50ad7ea3d35879a4a819ae97eb09b665f2113451d7ada7cb2b689b66.
//
// Solidity: event tssActiveMemberAppended(uint256 _roundId, bytes _groupKey, bytes[] activeTssMembers)
func (_TssGroupManager *TssGroupManagerFilterer) FilterTssActiveMemberAppended(opts *bind.FilterOpts) (*TssGroupManagerTssActiveMemberAppendedIterator, error) {

	logs, sub, err := _TssGroupManager.contract.FilterLogs(opts, "tssActiveMemberAppended")
	if err != nil {
		return nil, err
	}
	return &TssGroupManagerTssActiveMemberAppendedIterator{contract: _TssGroupManager.contract, event: "tssActiveMemberAppended", logs: logs, sub: sub}, nil
}

// WatchTssActiveMemberAppended is a free log subscription operation binding the contract event 0x61c7922e50ad7ea3d35879a4a819ae97eb09b665f2113451d7ada7cb2b689b66.
//
// Solidity: event tssActiveMemberAppended(uint256 _roundId, bytes _groupKey, bytes[] activeTssMembers)
func (_TssGroupManager *TssGroupManagerFilterer) WatchTssActiveMemberAppended(opts *bind.WatchOpts, sink chan<- *TssGroupManagerTssActiveMemberAppended) (event.Subscription, error) {

	logs, sub, err := _TssGroupManager.contract.WatchLogs(opts, "tssActiveMemberAppended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssGroupManagerTssActiveMemberAppended)
				if err := _TssGroupManager.contract.UnpackLog(event, "tssActiveMemberAppended", log); err != nil {
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

// ParseTssActiveMemberAppended is a log parse operation binding the contract event 0x61c7922e50ad7ea3d35879a4a819ae97eb09b665f2113451d7ada7cb2b689b66.
//
// Solidity: event tssActiveMemberAppended(uint256 _roundId, bytes _groupKey, bytes[] activeTssMembers)
func (_TssGroupManager *TssGroupManagerFilterer) ParseTssActiveMemberAppended(log types.Log) (*TssGroupManagerTssActiveMemberAppended, error) {
	event := new(TssGroupManagerTssActiveMemberAppended)
	if err := _TssGroupManager.contract.UnpackLog(event, "tssActiveMemberAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TssGroupManagerTssGroupMemberAppendIterator is returned from FilterTssGroupMemberAppend and is used to iterate over the raw logs and unpacked data for TssGroupMemberAppend events raised by the TssGroupManager contract.
type TssGroupManagerTssGroupMemberAppendIterator struct {
	Event *TssGroupManagerTssGroupMemberAppend // Event containing the contract specifics and raw log

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
func (it *TssGroupManagerTssGroupMemberAppendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssGroupManagerTssGroupMemberAppend)
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
		it.Event = new(TssGroupManagerTssGroupMemberAppend)
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
func (it *TssGroupManagerTssGroupMemberAppendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssGroupManagerTssGroupMemberAppendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssGroupManagerTssGroupMemberAppend represents a TssGroupMemberAppend event raised by the TssGroupManager contract.
type TssGroupManagerTssGroupMemberAppend struct {
	RoundId            *big.Int
	Threshold          *big.Int
	InActiveTssMembers [][]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTssGroupMemberAppend is a free log retrieval operation binding the contract event 0xfa26843b7ced045dd12994ee0e73c6f0904c3a4608bce58ae304d606035b079e.
//
// Solidity: event tssGroupMemberAppend(uint256 _roundId, uint256 _threshold, bytes[] _inActiveTssMembers)
func (_TssGroupManager *TssGroupManagerFilterer) FilterTssGroupMemberAppend(opts *bind.FilterOpts) (*TssGroupManagerTssGroupMemberAppendIterator, error) {

	logs, sub, err := _TssGroupManager.contract.FilterLogs(opts, "tssGroupMemberAppend")
	if err != nil {
		return nil, err
	}
	return &TssGroupManagerTssGroupMemberAppendIterator{contract: _TssGroupManager.contract, event: "tssGroupMemberAppend", logs: logs, sub: sub}, nil
}

// WatchTssGroupMemberAppend is a free log subscription operation binding the contract event 0xfa26843b7ced045dd12994ee0e73c6f0904c3a4608bce58ae304d606035b079e.
//
// Solidity: event tssGroupMemberAppend(uint256 _roundId, uint256 _threshold, bytes[] _inActiveTssMembers)
func (_TssGroupManager *TssGroupManagerFilterer) WatchTssGroupMemberAppend(opts *bind.WatchOpts, sink chan<- *TssGroupManagerTssGroupMemberAppend) (event.Subscription, error) {

	logs, sub, err := _TssGroupManager.contract.WatchLogs(opts, "tssGroupMemberAppend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssGroupManagerTssGroupMemberAppend)
				if err := _TssGroupManager.contract.UnpackLog(event, "tssGroupMemberAppend", log); err != nil {
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

// ParseTssGroupMemberAppend is a log parse operation binding the contract event 0xfa26843b7ced045dd12994ee0e73c6f0904c3a4608bce58ae304d606035b079e.
//
// Solidity: event tssGroupMemberAppend(uint256 _roundId, uint256 _threshold, bytes[] _inActiveTssMembers)
func (_TssGroupManager *TssGroupManagerFilterer) ParseTssGroupMemberAppend(log types.Log) (*TssGroupManagerTssGroupMemberAppend, error) {
	event := new(TssGroupManagerTssGroupMemberAppend)
	if err := _TssGroupManager.contract.UnpackLog(event, "tssGroupMemberAppend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

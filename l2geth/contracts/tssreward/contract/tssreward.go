// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TssReward

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

// TssRewardMetaData contains all meta data concerning the TssReward contract.
var TssRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deadAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockStartHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tssMembers\",\"type\":\"address[]\"}],\"name\":\"DistributeTssReward\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bestBlockID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockStartHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_length\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"_tssMembers\",\"type\":\"address[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dust\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"updateReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDust\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620013793803806200137983398181016040528101906200003791906200012b565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000172565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000f382620000c6565b9050919050565b6200010581620000e6565b81146200011157600080fd5b50565b6000815190506200012581620000fa565b92915050565b60008060408385031215620001455762000144620000c1565b5b6000620001558582860162000114565b9250506020620001688582860162000114565b9150509250929050565b6111f780620001826000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80632c79db11116100715780632c79db11146101685780633ccfd60b146101865780638da5cb5b14610190578063cfb550f1146101ae578063e04f6e35146101b8578063fad9aba3146101d4576100a9565b80630b50cd3e146100ae57806310a7fd7b146100de57806319d509a11461010e5780631a39d8ef1461012c57806327c8f8351461014a575b600080fd5b6100c860048036038101906100c39190610a28565b6101f2565b6040516100d59190610a83565b60405180910390f35b6100f860048036038101906100f39190610a9e565b61035c565b6040516101059190610ada565b60405180910390f35b610116610374565b6040516101239190610ada565b60405180910390f35b61013461037a565b6040516101419190610ada565b60405180910390f35b610152610380565b60405161015f9190610b36565b60405180910390f35b6101706103a6565b60405161017d9190610ada565b60405180910390f35b61018e6103f3565b005b610198610545565b6040516101a59190610b36565b60405180910390f35b6101b661056b565b005b6101d260048036038101906101cd9190610bf2565b6106e4565b005b6101dc6109a0565b6040516101e99190610ada565b60405180910390f35b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610284576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027b90610ce9565b60405180910390fd5b6005544710156102c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c090610d7b565b60405180910390fd5b60016003546102d89190610dca565b8314610319576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031090610e6c565b60405180910390fd5b82600381905550610335826005546109a690919063ffffffff16565b60058190555081600080858152602001908152602001600020819055506001905092915050565b60006020528060005260406000206000915090505481565b60035481565b60055481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006005544710156103ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e490610d7b565b60405180910390fd5b47905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610483576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047a90610efe565b60405180910390fd5b6005544710156104c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bf90610d7b565b60405180910390fd5b6000600581905550600047111561054357600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015610541573d6000803e3d6000fd5b505b565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f290610efe565b60405180910390fd5b600554471015610640576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063790610d7b565b60405180910390fd5b6000600454905061065e6004546005546109bc90919063ffffffff16565b600581905550600060048190555060008111156106e157600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6004549081150290604051600060405180830381858888f193505050501580156106df573d6000803e3d6000fd5b505b50565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610774576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076b90610ce9565b60405180910390fd5b6005544710156107b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b090610d7b565b60405180910390fd5b6000806000805b8663ffffffff16811015610834576107fe600080838b6107e09190610dca565b815260200190815260200160002054846109a690919063ffffffff16565b9250600080828a61080f9190610dca565b815260200190815260200160002060009055808061082c90610f1e565b9150506107c0565b50600082111561095a5761085485859050836109d290919063ffffffff16565b925060005b8585905081101561091b57600086868381811061087957610878610f67565b5b905060200201602081019061088e9190610fc2565b90506108a385846109a690919063ffffffff16565b92506108ba856005546109bc90919063ffffffff16565b6005819055508073ffffffffffffffffffffffffffffffffffffffff166108fc869081150290604051600060405180830381858888f19350505050158015610906573d6000803e3d6000fd5b5050808061091390610f1e565b915050610859565b50600061093182846109bc90919063ffffffff16565b9050600081111561095857610951816004546109a690919063ffffffff16565b6004819055505b505b7ff630cba6d450d736e85735388d4fe67a177b8a3685cdd7dee2bea7727b47860a8787878760405161098f94939291906110ed565b60405180910390a150505050505050565b60045481565b600081836109b49190610dca565b905092915050565b600081836109ca919061112d565b905092915050565b600081836109e09190611190565b905092915050565b600080fd5b600080fd5b6000819050919050565b610a05816109f2565b8114610a1057600080fd5b50565b600081359050610a22816109fc565b92915050565b60008060408385031215610a3f57610a3e6109e8565b5b6000610a4d85828601610a13565b9250506020610a5e85828601610a13565b9150509250929050565b60008115159050919050565b610a7d81610a68565b82525050565b6000602082019050610a986000830184610a74565b92915050565b600060208284031215610ab457610ab36109e8565b5b6000610ac284828501610a13565b91505092915050565b610ad4816109f2565b82525050565b6000602082019050610aef6000830184610acb565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b2082610af5565b9050919050565b610b3081610b15565b82525050565b6000602082019050610b4b6000830184610b27565b92915050565b600063ffffffff82169050919050565b610b6a81610b51565b8114610b7557600080fd5b50565b600081359050610b8781610b61565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610bb257610bb1610b8d565b5b8235905067ffffffffffffffff811115610bcf57610bce610b92565b5b602083019150836020820283011115610beb57610bea610b97565b5b9250929050565b60008060008060608587031215610c0c57610c0b6109e8565b5b6000610c1a87828801610a13565b9450506020610c2b87828801610b78565b935050604085013567ffffffffffffffff811115610c4c57610c4b6109ed565b5b610c5887828801610b9c565b925092505092959194509250565b600082825260208201905092915050565b7f747373207265776172642063616c6c206d65737361676520756e61757468656e60008201527f7469636174656400000000000000000000000000000000000000000000000000602082015250565b6000610cd3602783610c66565b9150610cde82610c77565b604082019050919050565b60006020820190508181036000830152610d0281610cc6565b9050919050565b7f62616c616e6365207265636f726420616e6420636f6e74726163742062616c6160008201527f6e636520617265206e6f7420657175616c000000000000000000000000000000602082015250565b6000610d65603183610c66565b9150610d7082610d09565b604082019050919050565b60006020820190508181036000830152610d9481610d58565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610dd5826109f2565b9150610de0836109f2565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610e1557610e14610d9b565b5b828201905092915050565b7f626c6f636b2069642075706461746520696c6c6567616c000000000000000000600082015250565b6000610e56601783610c66565b9150610e6182610e20565b602082019050919050565b60006020820190508181036000830152610e8581610e49565b9050919050565b7f6f6e6c792062652063616c6c656420627920746865206f776e6572206f66207460008201527f68697320636f6e74726163740000000000000000000000000000000000000000602082015250565b6000610ee8602c83610c66565b9150610ef382610e8c565b604082019050919050565b60006020820190508181036000830152610f1781610edb565b9050919050565b6000610f29826109f2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610f5c57610f5b610d9b565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b610f9f81610b15565b8114610faa57600080fd5b50565b600081359050610fbc81610f96565b92915050565b600060208284031215610fd857610fd76109e8565b5b6000610fe684828501610fad565b91505092915050565b6000819050919050565b600061101461100f61100a84610b51565b610fef565b6109f2565b9050919050565b61102481610ff9565b82525050565b600082825260208201905092915050565b6000819050919050565b61104e81610b15565b82525050565b60006110608383611045565b60208301905092915050565b600061107b6020840184610fad565b905092915050565b6000602082019050919050565b600061109c838561102a565b93506110a78261103b565b8060005b858110156110e0576110bd828461106c565b6110c78882611054565b97506110d283611083565b9250506001810190506110ab565b5085925050509392505050565b60006060820190506111026000830187610acb565b61110f602083018661101b565b8181036040830152611122818486611090565b905095945050505050565b6000611138826109f2565b9150611143836109f2565b92508282101561115657611155610d9b565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061119b826109f2565b91506111a6836109f2565b9250826111b6576111b5611161565b5b82820490509291505056fea2646970667358221220c357406ca5ad9f3d6e68796737f17d9167fb463b128033a344818d24879906f864736f6c63430008090033",
}

// TssRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use TssRewardMetaData.ABI instead.
var TssRewardABI = TssRewardMetaData.ABI

// TssRewardBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TssRewardMetaData.Bin instead.
var TssRewardBin = TssRewardMetaData.Bin

// DeployTssReward deploys a new Ethereum contract, binding an instance of TssReward to it.
func DeployTssReward(auth *bind.TransactOpts, backend bind.ContractBackend, _deadAddress common.Address, _owner common.Address) (common.Address, *types.Transaction, *TssReward, error) {
	parsed, err := TssRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TssRewardBin), backend, _deadAddress, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TssReward{TssRewardCaller: TssRewardCaller{contract: contract}, TssRewardTransactor: TssRewardTransactor{contract: contract}, TssRewardFilterer: TssRewardFilterer{contract: contract}}, nil
}

// TssReward is an auto generated Go binding around an Ethereum contract.
type TssReward struct {
	TssRewardCaller     // Read-only binding to the contract
	TssRewardTransactor // Write-only binding to the contract
	TssRewardFilterer   // Log filterer for contract events
}

// TssRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type TssRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TssRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TssRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TssRewardSession struct {
	Contract     *TssReward        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TssRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TssRewardCallerSession struct {
	Contract *TssRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TssRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TssRewardTransactorSession struct {
	Contract     *TssRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TssRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type TssRewardRaw struct {
	Contract *TssReward // Generic contract binding to access the raw methods on
}

// TssRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TssRewardCallerRaw struct {
	Contract *TssRewardCaller // Generic read-only contract binding to access the raw methods on
}

// TssRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TssRewardTransactorRaw struct {
	Contract *TssRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTssReward creates a new instance of TssReward, bound to a specific deployed contract.
func NewTssReward(address common.Address, backend bind.ContractBackend) (*TssReward, error) {
	contract, err := bindTssReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TssReward{TssRewardCaller: TssRewardCaller{contract: contract}, TssRewardTransactor: TssRewardTransactor{contract: contract}, TssRewardFilterer: TssRewardFilterer{contract: contract}}, nil
}

// NewTssRewardCaller creates a new read-only instance of TssReward, bound to a specific deployed contract.
func NewTssRewardCaller(address common.Address, caller bind.ContractCaller) (*TssRewardCaller, error) {
	contract, err := bindTssReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TssRewardCaller{contract: contract}, nil
}

// NewTssRewardTransactor creates a new write-only instance of TssReward, bound to a specific deployed contract.
func NewTssRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*TssRewardTransactor, error) {
	contract, err := bindTssReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TssRewardTransactor{contract: contract}, nil
}

// NewTssRewardFilterer creates a new log filterer instance of TssReward, bound to a specific deployed contract.
func NewTssRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*TssRewardFilterer, error) {
	contract, err := bindTssReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TssRewardFilterer{contract: contract}, nil
}

// bindTssReward binds a generic wrapper to an already deployed contract.
func bindTssReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TssRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TssReward *TssRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TssReward.Contract.TssRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TssReward *TssRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssReward.Contract.TssRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TssReward *TssRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TssReward.Contract.TssRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TssReward *TssRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TssReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TssReward *TssRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TssReward *TssRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TssReward.Contract.contract.Transact(opts, method, params...)
}

// BestBlockID is a free data retrieval call binding the contract method 0x19d509a1.
//
// Solidity: function bestBlockID() view returns(uint256)
func (_TssReward *TssRewardCaller) BestBlockID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TssReward.contract.Call(opts, &out, "bestBlockID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BestBlockID is a free data retrieval call binding the contract method 0x19d509a1.
//
// Solidity: function bestBlockID() view returns(uint256)
func (_TssReward *TssRewardSession) BestBlockID() (*big.Int, error) {
	return _TssReward.Contract.BestBlockID(&_TssReward.CallOpts)
}

// BestBlockID is a free data retrieval call binding the contract method 0x19d509a1.
//
// Solidity: function bestBlockID() view returns(uint256)
func (_TssReward *TssRewardCallerSession) BestBlockID() (*big.Int, error) {
	return _TssReward.Contract.BestBlockID(&_TssReward.CallOpts)
}

// DeadAddress is a free data retrieval call binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() view returns(address)
func (_TssReward *TssRewardCaller) DeadAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssReward.contract.Call(opts, &out, "deadAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeadAddress is a free data retrieval call binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() view returns(address)
func (_TssReward *TssRewardSession) DeadAddress() (common.Address, error) {
	return _TssReward.Contract.DeadAddress(&_TssReward.CallOpts)
}

// DeadAddress is a free data retrieval call binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() view returns(address)
func (_TssReward *TssRewardCallerSession) DeadAddress() (common.Address, error) {
	return _TssReward.Contract.DeadAddress(&_TssReward.CallOpts)
}

// Dust is a free data retrieval call binding the contract method 0xfad9aba3.
//
// Solidity: function dust() view returns(uint256)
func (_TssReward *TssRewardCaller) Dust(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TssReward.contract.Call(opts, &out, "dust")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dust is a free data retrieval call binding the contract method 0xfad9aba3.
//
// Solidity: function dust() view returns(uint256)
func (_TssReward *TssRewardSession) Dust() (*big.Int, error) {
	return _TssReward.Contract.Dust(&_TssReward.CallOpts)
}

// Dust is a free data retrieval call binding the contract method 0xfad9aba3.
//
// Solidity: function dust() view returns(uint256)
func (_TssReward *TssRewardCallerSession) Dust() (*big.Int, error) {
	return _TssReward.Contract.Dust(&_TssReward.CallOpts)
}

// Ledger is a free data retrieval call binding the contract method 0x10a7fd7b.
//
// Solidity: function ledger(uint256 ) view returns(uint256)
func (_TssReward *TssRewardCaller) Ledger(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TssReward.contract.Call(opts, &out, "ledger", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ledger is a free data retrieval call binding the contract method 0x10a7fd7b.
//
// Solidity: function ledger(uint256 ) view returns(uint256)
func (_TssReward *TssRewardSession) Ledger(arg0 *big.Int) (*big.Int, error) {
	return _TssReward.Contract.Ledger(&_TssReward.CallOpts, arg0)
}

// Ledger is a free data retrieval call binding the contract method 0x10a7fd7b.
//
// Solidity: function ledger(uint256 ) view returns(uint256)
func (_TssReward *TssRewardCallerSession) Ledger(arg0 *big.Int) (*big.Int, error) {
	return _TssReward.Contract.Ledger(&_TssReward.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssReward *TssRewardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TssReward.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssReward *TssRewardSession) Owner() (common.Address, error) {
	return _TssReward.Contract.Owner(&_TssReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TssReward *TssRewardCallerSession) Owner() (common.Address, error) {
	return _TssReward.Contract.Owner(&_TssReward.CallOpts)
}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_TssReward *TssRewardCaller) QueryReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TssReward.contract.Call(opts, &out, "queryReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_TssReward *TssRewardSession) QueryReward() (*big.Int, error) {
	return _TssReward.Contract.QueryReward(&_TssReward.CallOpts)
}

// QueryReward is a free data retrieval call binding the contract method 0x2c79db11.
//
// Solidity: function queryReward() view returns(uint256)
func (_TssReward *TssRewardCallerSession) QueryReward() (*big.Int, error) {
	return _TssReward.Contract.QueryReward(&_TssReward.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TssReward *TssRewardCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TssReward.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TssReward *TssRewardSession) TotalAmount() (*big.Int, error) {
	return _TssReward.Contract.TotalAmount(&_TssReward.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TssReward *TssRewardCallerSession) TotalAmount() (*big.Int, error) {
	return _TssReward.Contract.TotalAmount(&_TssReward.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xe04f6e35.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, address[] _tssMembers) returns()
func (_TssReward *TssRewardTransactor) ClaimReward(opts *bind.TransactOpts, _blockStartHeight *big.Int, _length uint32, _tssMembers []common.Address) (*types.Transaction, error) {
	return _TssReward.contract.Transact(opts, "claimReward", _blockStartHeight, _length, _tssMembers)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xe04f6e35.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, address[] _tssMembers) returns()
func (_TssReward *TssRewardSession) ClaimReward(_blockStartHeight *big.Int, _length uint32, _tssMembers []common.Address) (*types.Transaction, error) {
	return _TssReward.Contract.ClaimReward(&_TssReward.TransactOpts, _blockStartHeight, _length, _tssMembers)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xe04f6e35.
//
// Solidity: function claimReward(uint256 _blockStartHeight, uint32 _length, address[] _tssMembers) returns()
func (_TssReward *TssRewardTransactorSession) ClaimReward(_blockStartHeight *big.Int, _length uint32, _tssMembers []common.Address) (*types.Transaction, error) {
	return _TssReward.Contract.ClaimReward(&_TssReward.TransactOpts, _blockStartHeight, _length, _tssMembers)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_TssReward *TssRewardTransactor) UpdateReward(opts *bind.TransactOpts, _blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TssReward.contract.Transact(opts, "updateReward", _blockID, _amount)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_TssReward *TssRewardSession) UpdateReward(_blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TssReward.Contract.UpdateReward(&_TssReward.TransactOpts, _blockID, _amount)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x0b50cd3e.
//
// Solidity: function updateReward(uint256 _blockID, uint256 _amount) returns(bool)
func (_TssReward *TssRewardTransactorSession) UpdateReward(_blockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TssReward.Contract.UpdateReward(&_TssReward.TransactOpts, _blockID, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TssReward *TssRewardTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssReward.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TssReward *TssRewardSession) Withdraw() (*types.Transaction, error) {
	return _TssReward.Contract.Withdraw(&_TssReward.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TssReward *TssRewardTransactorSession) Withdraw() (*types.Transaction, error) {
	return _TssReward.Contract.Withdraw(&_TssReward.TransactOpts)
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_TssReward *TssRewardTransactor) WithdrawDust(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TssReward.contract.Transact(opts, "withdrawDust")
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_TssReward *TssRewardSession) WithdrawDust() (*types.Transaction, error) {
	return _TssReward.Contract.WithdrawDust(&_TssReward.TransactOpts)
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_TssReward *TssRewardTransactorSession) WithdrawDust() (*types.Transaction, error) {
	return _TssReward.Contract.WithdrawDust(&_TssReward.TransactOpts)
}

// TssRewardDistributeTssRewardIterator is returned from FilterDistributeTssReward and is used to iterate over the raw logs and unpacked data for DistributeTssReward events raised by the TssReward contract.
type TssRewardDistributeTssRewardIterator struct {
	Event *TssRewardDistributeTssReward // Event containing the contract specifics and raw log

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
func (it *TssRewardDistributeTssRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssRewardDistributeTssReward)
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
		it.Event = new(TssRewardDistributeTssReward)
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
func (it *TssRewardDistributeTssRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssRewardDistributeTssRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssRewardDistributeTssReward represents a DistributeTssReward event raised by the TssReward contract.
type TssRewardDistributeTssReward struct {
	BlockStartHeight *big.Int
	Length           *big.Int
	TssMembers       []common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributeTssReward is a free log retrieval operation binding the contract event 0xf630cba6d450d736e85735388d4fe67a177b8a3685cdd7dee2bea7727b47860a.
//
// Solidity: event DistributeTssReward(uint256 blockStartHeight, uint256 length, address[] tssMembers)
func (_TssReward *TssRewardFilterer) FilterDistributeTssReward(opts *bind.FilterOpts) (*TssRewardDistributeTssRewardIterator, error) {

	logs, sub, err := _TssReward.contract.FilterLogs(opts, "DistributeTssReward")
	if err != nil {
		return nil, err
	}
	return &TssRewardDistributeTssRewardIterator{contract: _TssReward.contract, event: "DistributeTssReward", logs: logs, sub: sub}, nil
}

// WatchDistributeTssReward is a free log subscription operation binding the contract event 0xf630cba6d450d736e85735388d4fe67a177b8a3685cdd7dee2bea7727b47860a.
//
// Solidity: event DistributeTssReward(uint256 blockStartHeight, uint256 length, address[] tssMembers)
func (_TssReward *TssRewardFilterer) WatchDistributeTssReward(opts *bind.WatchOpts, sink chan<- *TssRewardDistributeTssReward) (event.Subscription, error) {

	logs, sub, err := _TssReward.contract.WatchLogs(opts, "DistributeTssReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssRewardDistributeTssReward)
				if err := _TssReward.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
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

// ParseDistributeTssReward is a log parse operation binding the contract event 0xf630cba6d450d736e85735388d4fe67a177b8a3685cdd7dee2bea7727b47860a.
//
// Solidity: event DistributeTssReward(uint256 blockStartHeight, uint256 length, address[] tssMembers)
func (_TssReward *TssRewardFilterer) ParseDistributeTssReward(log types.Log) (*TssRewardDistributeTssReward, error) {
	event := new(TssRewardDistributeTssReward)
	if err := _TssReward.contract.UnpackLog(event, "DistributeTssReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

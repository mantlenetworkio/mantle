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
)

// IDataLayrServiceManagerDataStoreMetadata is an auto generated low-level Go binding around an user-defined struct.
type IDataLayrServiceManagerDataStoreMetadata struct {
	HeaderHash          [32]byte
	DurationDataStoreId uint32
	GlobalDataStoreId   uint32
	BlockNumber         uint32
	Fee                 *big.Int
	Confirmer           common.Address
	SignatoryRecordHash [32]byte
}

// IDataLayrServiceManagerDataStoreSearchData is an auto generated low-level Go binding around an user-defined struct.
type IDataLayrServiceManagerDataStoreSearchData struct {
	Metadata  IDataLayrServiceManagerDataStoreMetadata
	Duration  uint8
	Timestamp *big.Int
	Index     uint32
}

// BVMEigenDataLayrChainMetaData contains all meta data concerning the BVMEigenDataLayrChain contract.
var BVMEigenDataLayrChainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupStoreNumber\",\"type\":\"uint32\"}],\"name\":\"RollupStoreConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"}],\"name\":\"RollupStoreInitialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_STALE_MEASURE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreSearchData\",\"name\":\"searchData\",\"type\":\"tuple\"}],\"name\":\"confirmData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataManageAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"dataStoreIdToRollupStoreNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fraudProofPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSunmitL2Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dataManageAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupStoreNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rollupStores\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"confirmAt\",\"type\":\"uint32\"},{\"internalType\":\"enumBVM_EigenDataLayrChain.RollupStoreStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalOperatorsIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405262015180609b5534801561001757600080fd5b506114ea806100276000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80637ef01d5e11610097578063b537c4c711610066578063b537c4c714610225578063f249502914610245578063f2a8f12414610265578063f2fde38b1461026e57600080fd5b80637ef01d5e146101985780638b85902b146101ab5780638da5cb5b146101b4578063b0393a37146101d257600080fd5b80635c1bba38116100d35780635c1bba381461012f5780635e8b3f2d14610174578063715018a61461017d5780637a6847851461018557600080fd5b806312aeb89c146100fa578063428bba0914610111578063485cc9551461011a575b600080fd5b609a545b6040519081526020015b60405180910390f35b6100fe609e5481565b61012d610128366004610e92565b610281565b005b60975461014f9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610108565b6100fe60995481565b61012d610475565b61012d610193366004610f3c565b610489565b61012d6101a636600461105a565b61080e565b6100fe609a5481565b60335473ffffffffffffffffffffffffffffffffffffffff1661014f565b6102166101e03660046111a8565b609c6020526000908152604090205463ffffffff8082169164010000000081049091169068010000000000000000900460ff1683565b604051610108939291906111f0565b6100fe61023336600461124a565b609d6020526000908152604090205481565b60985461014f9073ffffffffffffffffffffffffffffffffffffffff1681565b6100fe609b5481565b61012d61027c36600461126e565b610b7b565b600054610100900460ff16158080156102a15750600054600160ff909116105b806102bb5750303b1580156102bb575060005460ff166001145b61034c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156103aa57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6103b2610c32565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560988054928516929091169190911790556064609955801561047057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b505050565b61047d610cd1565b6104876000610d52565b565b60975473ffffffffffffffffffffffffffffffffffffffff163314610530576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4f6e6c79207468652073657175656e6365722063616e2073746f72652064617460448201527f61000000000000000000000000000000000000000000000000000000000000006064820152608401610343565b60995461054363ffffffff8516436112b8565b106105aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7374616b65732074616b656e2066726f6d20746f6f206c6f6e672061676f00006044820152606401610343565b609a54811161063b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f5f6c32426c6f636b4e756d626572206d757374206c61726765207468616e206c60448201527f32426c6f636b4e756d62657200000000000000000000000000000000000000006064820152608401610343565b609a819055609854604080517f72d18e8d000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff16916372d18e8d916004808301926020929190829003018186803b1580156106ab57600080fd5b505afa1580156106bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e391906112cf565b6098546040517fdcf49ea700000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff169063dcf49ea79061074690339030908a908a908a908f908f90600401611335565b602060405180830381600087803b15801561076057600080fd5b505af1158015610774573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079891906112cf565b5063ffffffff81166000818152609d60209081526040918290207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff905590519182527f957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b5910160405180910390a150505050505050565b60975473ffffffffffffffffffffffffffffffffffffffff1633146108b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4f6e6c79207468652073657175656e6365722063616e2073746f72652064617460448201527f61000000000000000000000000000000000000000000000000000000000000006064820152608401610343565b805160409081015163ffffffff166000908152609d60205220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff146109a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605560248201527f446174612073746f72652065697468657220776173206e6f7420696e6974696160448201527f6c697a65642062792074686520726f6c6c757020636f6e74726163742c206f7260648201527f20697320616c726561647920636f6e6669726d65640000000000000000000000608482015260a401610343565b6098546040517f5189951500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116906351899515906109fd90869086908690600401611399565b600060405180830381600087803b158015610a1757600080fd5b505af1158015610a2b573d6000803e3d6000fd5b50505050604051806060016040528082600001516040015163ffffffff168152602001609b5442610a5c9190611463565b63ffffffff16815260200160019052609e546000908152609c6020908152604091829020835181549285015163ffffffff908116640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009094169116179190911780825591830151909182907fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1668010000000000000000836002811115610b0b57610b0b6111c1565b021790555050609e8054835160409081015163ffffffff166000908152609d6020529081208290557fe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a99350909190610b628361147b565b9091555060405163ffffffff9091168152602001610467565b610b83610cd1565b73ffffffffffffffffffffffffffffffffffffffff8116610c26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610343565b610c2f81610d52565b50565b600054610100900460ff16610cc9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610343565b610487610dc9565b60335473ffffffffffffffffffffffffffffffffffffffff163314610487576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610343565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610e60576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610343565b61048733610d52565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e8d57600080fd5b919050565b60008060408385031215610ea557600080fd5b610eae83610e69565b9150610ebc60208401610e69565b90509250929050565b60008083601f840112610ed757600080fd5b50813567ffffffffffffffff811115610eef57600080fd5b602083019150836020828501011115610f0757600080fd5b9250929050565b803560ff81168114610e8d57600080fd5b63ffffffff81168114610c2f57600080fd5b8035610e8d81610f1f565b60008060008060008060a08789031215610f5557600080fd5b863567ffffffffffffffff811115610f6c57600080fd5b610f7889828a01610ec5565b9097509550610f8b905060208801610f0e565b93506040870135610f9b81610f1f565b92506060870135610fab81610f1f565b80925050608087013590509295509295509295565b6040516080810167ffffffffffffffff8111828210171561100a577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b60405160e0810167ffffffffffffffff8111828210171561100a577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080600083850361016081121561107157600080fd5b843567ffffffffffffffff81111561108857600080fd5b61109487828801610ec5565b9095509350507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001610140808212156110cc57600080fd5b6110d4610fc0565b60e08312156110e257600080fd5b6110ea611010565b925060208701358352604087013561110181610f1f565b6020840152606087013561111481610f1f565b6040840152608087013561112781610f1f565b606084015260a08701356bffffffffffffffffffffffff8116811461114b57600080fd5b608084015261115c60c08801610e69565b60a084015260e087013560c084015282815261117b6101008801610f0e565b60208201526101208701356040820152611196828801610f31565b60608201528093505050509250925092565b6000602082840312156111ba57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b63ffffffff848116825283166020820152606081016003831061123c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826040830152949350505050565b60006020828403121561125c57600080fd5b813561126781610f1f565b9392505050565b60006020828403121561128057600080fd5b61126782610e69565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156112ca576112ca611289565b500390565b6000602082840312156112e157600080fd5b815161126781610f1f565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff88811682528716602082015260ff8616604082015263ffffffff85811660608301528416608082015260c060a0820181905260009061138c90830184866112ec565b9998505050505050505050565b60006101608083526113ae81840186886112ec565b915050825180516020840152602081015163ffffffff808216604086015280604084015116606086015280606084015116608086015250506bffffffffffffffffffffffff60808201511660a084015273ffffffffffffffffffffffffffffffffffffffff60a08201511660c084015260c081015160e084015250602083015161143e61010084018260ff169052565b50604083015161012083015260609092015163ffffffff166101409091015292915050565b6000821982111561147657611476611289565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156114ad576114ad611289565b506001019056fea26469706673582212209a4b323247309f400723fbc52423c565649f9037041d32a1892f8cbd74f7295164736f6c63430008090033",
}

// BVMEigenDataLayrChainABI is the input ABI used to generate the binding from.
// Deprecated: Use BVMEigenDataLayrChainMetaData.ABI instead.
var BVMEigenDataLayrChainABI = BVMEigenDataLayrChainMetaData.ABI

// BVMEigenDataLayrChainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BVMEigenDataLayrChainMetaData.Bin instead.
var BVMEigenDataLayrChainBin = BVMEigenDataLayrChainMetaData.Bin

// DeployBVMEigenDataLayrChain deploys a new Ethereum contract, binding an instance of BVMEigenDataLayrChain to it.
func DeployBVMEigenDataLayrChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BVMEigenDataLayrChain, error) {
	parsed, err := BVMEigenDataLayrChainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BVMEigenDataLayrChainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BVMEigenDataLayrChain{BVMEigenDataLayrChainCaller: BVMEigenDataLayrChainCaller{contract: contract}, BVMEigenDataLayrChainTransactor: BVMEigenDataLayrChainTransactor{contract: contract}, BVMEigenDataLayrChainFilterer: BVMEigenDataLayrChainFilterer{contract: contract}}, nil
}

// BVMEigenDataLayrChain is an auto generated Go binding around an Ethereum contract.
type BVMEigenDataLayrChain struct {
	BVMEigenDataLayrChainCaller     // Read-only binding to the contract
	BVMEigenDataLayrChainTransactor // Write-only binding to the contract
	BVMEigenDataLayrChainFilterer   // Log filterer for contract events
}

// BVMEigenDataLayrChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BVMEigenDataLayrChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMEigenDataLayrChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BVMEigenDataLayrChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMEigenDataLayrChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BVMEigenDataLayrChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMEigenDataLayrChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BVMEigenDataLayrChainSession struct {
	Contract     *BVMEigenDataLayrChain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BVMEigenDataLayrChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BVMEigenDataLayrChainCallerSession struct {
	Contract *BVMEigenDataLayrChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// BVMEigenDataLayrChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BVMEigenDataLayrChainTransactorSession struct {
	Contract     *BVMEigenDataLayrChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// BVMEigenDataLayrChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BVMEigenDataLayrChainRaw struct {
	Contract *BVMEigenDataLayrChain // Generic contract binding to access the raw methods on
}

// BVMEigenDataLayrChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BVMEigenDataLayrChainCallerRaw struct {
	Contract *BVMEigenDataLayrChainCaller // Generic read-only contract binding to access the raw methods on
}

// BVMEigenDataLayrChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BVMEigenDataLayrChainTransactorRaw struct {
	Contract *BVMEigenDataLayrChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBVMEigenDataLayrChain creates a new instance of BVMEigenDataLayrChain, bound to a specific deployed contract.
func NewBVMEigenDataLayrChain(address common.Address, backend bind.ContractBackend) (*BVMEigenDataLayrChain, error) {
	contract, err := bindBVMEigenDataLayrChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChain{BVMEigenDataLayrChainCaller: BVMEigenDataLayrChainCaller{contract: contract}, BVMEigenDataLayrChainTransactor: BVMEigenDataLayrChainTransactor{contract: contract}, BVMEigenDataLayrChainFilterer: BVMEigenDataLayrChainFilterer{contract: contract}}, nil
}

// NewBVMEigenDataLayrChainCaller creates a new read-only instance of BVMEigenDataLayrChain, bound to a specific deployed contract.
func NewBVMEigenDataLayrChainCaller(address common.Address, caller bind.ContractCaller) (*BVMEigenDataLayrChainCaller, error) {
	contract, err := bindBVMEigenDataLayrChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainCaller{contract: contract}, nil
}

// NewBVMEigenDataLayrChainTransactor creates a new write-only instance of BVMEigenDataLayrChain, bound to a specific deployed contract.
func NewBVMEigenDataLayrChainTransactor(address common.Address, transactor bind.ContractTransactor) (*BVMEigenDataLayrChainTransactor, error) {
	contract, err := bindBVMEigenDataLayrChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainTransactor{contract: contract}, nil
}

// NewBVMEigenDataLayrChainFilterer creates a new log filterer instance of BVMEigenDataLayrChain, bound to a specific deployed contract.
func NewBVMEigenDataLayrChainFilterer(address common.Address, filterer bind.ContractFilterer) (*BVMEigenDataLayrChainFilterer, error) {
	contract, err := bindBVMEigenDataLayrChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainFilterer{contract: contract}, nil
}

// bindBVMEigenDataLayrChain binds a generic wrapper to an already deployed contract.
func bindBVMEigenDataLayrChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BVMEigenDataLayrChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMEigenDataLayrChain.Contract.BVMEigenDataLayrChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.BVMEigenDataLayrChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.BVMEigenDataLayrChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMEigenDataLayrChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.contract.Transact(opts, method, params...)
}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) BLOCKSTALEMEASURE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "BLOCK_STALE_MEASURE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) BLOCKSTALEMEASURE() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.BLOCKSTALEMEASURE(&_BVMEigenDataLayrChain.CallOpts)
}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) BLOCKSTALEMEASURE() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.BLOCKSTALEMEASURE(&_BVMEigenDataLayrChain.CallOpts)
}

// DataManageAddress is a free data retrieval call binding the contract method 0xf2495029.
//
// Solidity: function dataManageAddress() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) DataManageAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "dataManageAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataManageAddress is a free data retrieval call binding the contract method 0xf2495029.
//
// Solidity: function dataManageAddress() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) DataManageAddress() (common.Address, error) {
	return _BVMEigenDataLayrChain.Contract.DataManageAddress(&_BVMEigenDataLayrChain.CallOpts)
}

// DataManageAddress is a free data retrieval call binding the contract method 0xf2495029.
//
// Solidity: function dataManageAddress() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) DataManageAddress() (common.Address, error) {
	return _BVMEigenDataLayrChain.Contract.DataManageAddress(&_BVMEigenDataLayrChain.CallOpts)
}

// DataStoreIdToRollupStoreNumber is a free data retrieval call binding the contract method 0xb537c4c7.
//
// Solidity: function dataStoreIdToRollupStoreNumber(uint32 ) view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) DataStoreIdToRollupStoreNumber(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "dataStoreIdToRollupStoreNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DataStoreIdToRollupStoreNumber is a free data retrieval call binding the contract method 0xb537c4c7.
//
// Solidity: function dataStoreIdToRollupStoreNumber(uint32 ) view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) DataStoreIdToRollupStoreNumber(arg0 uint32) (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.DataStoreIdToRollupStoreNumber(&_BVMEigenDataLayrChain.CallOpts, arg0)
}

// DataStoreIdToRollupStoreNumber is a free data retrieval call binding the contract method 0xb537c4c7.
//
// Solidity: function dataStoreIdToRollupStoreNumber(uint32 ) view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) DataStoreIdToRollupStoreNumber(arg0 uint32) (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.DataStoreIdToRollupStoreNumber(&_BVMEigenDataLayrChain.CallOpts, arg0)
}

// FraudProofPeriod is a free data retrieval call binding the contract method 0xf2a8f124.
//
// Solidity: function fraudProofPeriod() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) FraudProofPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "fraudProofPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FraudProofPeriod is a free data retrieval call binding the contract method 0xf2a8f124.
//
// Solidity: function fraudProofPeriod() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) FraudProofPeriod() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.FraudProofPeriod(&_BVMEigenDataLayrChain.CallOpts)
}

// FraudProofPeriod is a free data retrieval call binding the contract method 0xf2a8f124.
//
// Solidity: function fraudProofPeriod() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) FraudProofPeriod() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.FraudProofPeriod(&_BVMEigenDataLayrChain.CallOpts)
}

// GetSunmitL2Block is a free data retrieval call binding the contract method 0x12aeb89c.
//
// Solidity: function getSunmitL2Block() view returns(uint256 _l2BlockNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) GetSunmitL2Block(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "getSunmitL2Block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSunmitL2Block is a free data retrieval call binding the contract method 0x12aeb89c.
//
// Solidity: function getSunmitL2Block() view returns(uint256 _l2BlockNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) GetSunmitL2Block() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.GetSunmitL2Block(&_BVMEigenDataLayrChain.CallOpts)
}

// GetSunmitL2Block is a free data retrieval call binding the contract method 0x12aeb89c.
//
// Solidity: function getSunmitL2Block() view returns(uint256 _l2BlockNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) GetSunmitL2Block() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.GetSunmitL2Block(&_BVMEigenDataLayrChain.CallOpts)
}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) L2BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "l2BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) L2BlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.L2BlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) L2BlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.L2BlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) Owner() (common.Address, error) {
	return _BVMEigenDataLayrChain.Contract.Owner(&_BVMEigenDataLayrChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) Owner() (common.Address, error) {
	return _BVMEigenDataLayrChain.Contract.Owner(&_BVMEigenDataLayrChain.CallOpts)
}

// RollupStoreNumber is a free data retrieval call binding the contract method 0x428bba09.
//
// Solidity: function rollupStoreNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) RollupStoreNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "rollupStoreNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RollupStoreNumber is a free data retrieval call binding the contract method 0x428bba09.
//
// Solidity: function rollupStoreNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) RollupStoreNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.RollupStoreNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// RollupStoreNumber is a free data retrieval call binding the contract method 0x428bba09.
//
// Solidity: function rollupStoreNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) RollupStoreNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.RollupStoreNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// RollupStores is a free data retrieval call binding the contract method 0xb0393a37.
//
// Solidity: function rollupStores(uint256 ) view returns(uint32 dataStoreId, uint32 confirmAt, uint8 status)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) RollupStores(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DataStoreId uint32
	ConfirmAt   uint32
	Status      uint8
}, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "rollupStores", arg0)

	outstruct := new(struct {
		DataStoreId uint32
		ConfirmAt   uint32
		Status      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DataStoreId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ConfirmAt = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// RollupStores is a free data retrieval call binding the contract method 0xb0393a37.
//
// Solidity: function rollupStores(uint256 ) view returns(uint32 dataStoreId, uint32 confirmAt, uint8 status)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) RollupStores(arg0 *big.Int) (struct {
	DataStoreId uint32
	ConfirmAt   uint32
	Status      uint8
}, error) {
	return _BVMEigenDataLayrChain.Contract.RollupStores(&_BVMEigenDataLayrChain.CallOpts, arg0)
}

// RollupStores is a free data retrieval call binding the contract method 0xb0393a37.
//
// Solidity: function rollupStores(uint256 ) view returns(uint32 dataStoreId, uint32 confirmAt, uint8 status)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) RollupStores(arg0 *big.Int) (struct {
	DataStoreId uint32
	ConfirmAt   uint32
	Status      uint8
}, error) {
	return _BVMEigenDataLayrChain.Contract.RollupStores(&_BVMEigenDataLayrChain.CallOpts, arg0)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) Sequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "sequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) Sequencer() (common.Address, error) {
	return _BVMEigenDataLayrChain.Contract.Sequencer(&_BVMEigenDataLayrChain.CallOpts)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) Sequencer() (common.Address, error) {
	return _BVMEigenDataLayrChain.Contract.Sequencer(&_BVMEigenDataLayrChain.CallOpts)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x7ef01d5e.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) ConfirmData(opts *bind.TransactOpts, data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "confirmData", data, searchData)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x7ef01d5e.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) ConfirmData(data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.ConfirmData(&_BVMEigenDataLayrChain.TransactOpts, data, searchData)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x7ef01d5e.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) ConfirmData(data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.ConfirmData(&_BVMEigenDataLayrChain.TransactOpts, data, searchData)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) Initialize(opts *bind.TransactOpts, _sequencer common.Address, _dataManageAddress common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "initialize", _sequencer, _dataManageAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) Initialize(_sequencer common.Address, _dataManageAddress common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.Initialize(&_BVMEigenDataLayrChain.TransactOpts, _sequencer, _dataManageAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) Initialize(_sequencer common.Address, _dataManageAddress common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.Initialize(&_BVMEigenDataLayrChain.TransactOpts, _sequencer, _dataManageAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) RenounceOwnership() (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.RenounceOwnership(&_BVMEigenDataLayrChain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.RenounceOwnership(&_BVMEigenDataLayrChain.TransactOpts)
}

// StoreData is a paid mutator transaction binding the contract method 0x7a684785.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex, uint256 _l2BlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) StoreData(opts *bind.TransactOpts, header []byte, duration uint8, blockNumber uint32, totalOperatorsIndex uint32, _l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "storeData", header, duration, blockNumber, totalOperatorsIndex, _l2BlockNumber)
}

// StoreData is a paid mutator transaction binding the contract method 0x7a684785.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex, uint256 _l2BlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) StoreData(header []byte, duration uint8, blockNumber uint32, totalOperatorsIndex uint32, _l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.StoreData(&_BVMEigenDataLayrChain.TransactOpts, header, duration, blockNumber, totalOperatorsIndex, _l2BlockNumber)
}

// StoreData is a paid mutator transaction binding the contract method 0x7a684785.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex, uint256 _l2BlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) StoreData(header []byte, duration uint8, blockNumber uint32, totalOperatorsIndex uint32, _l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.StoreData(&_BVMEigenDataLayrChain.TransactOpts, header, duration, blockNumber, totalOperatorsIndex, _l2BlockNumber)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.TransferOwnership(&_BVMEigenDataLayrChain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.TransferOwnership(&_BVMEigenDataLayrChain.TransactOpts, newOwner)
}

// BVMEigenDataLayrChainInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainInitializedIterator struct {
	Event *BVMEigenDataLayrChainInitialized // Event containing the contract specifics and raw log

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
func (it *BVMEigenDataLayrChainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMEigenDataLayrChainInitialized)
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
		it.Event = new(BVMEigenDataLayrChainInitialized)
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
func (it *BVMEigenDataLayrChainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMEigenDataLayrChainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMEigenDataLayrChainInitialized represents a Initialized event raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) FilterInitialized(opts *bind.FilterOpts) (*BVMEigenDataLayrChainInitializedIterator, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainInitializedIterator{contract: _BVMEigenDataLayrChain.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BVMEigenDataLayrChainInitialized) (event.Subscription, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMEigenDataLayrChainInitialized)
				if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) ParseInitialized(log types.Log) (*BVMEigenDataLayrChainInitialized, error) {
	event := new(BVMEigenDataLayrChainInitialized)
	if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMEigenDataLayrChainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainOwnershipTransferredIterator struct {
	Event *BVMEigenDataLayrChainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BVMEigenDataLayrChainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMEigenDataLayrChainOwnershipTransferred)
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
		it.Event = new(BVMEigenDataLayrChainOwnershipTransferred)
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
func (it *BVMEigenDataLayrChainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMEigenDataLayrChainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMEigenDataLayrChainOwnershipTransferred represents a OwnershipTransferred event raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BVMEigenDataLayrChainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BVMEigenDataLayrChain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainOwnershipTransferredIterator{contract: _BVMEigenDataLayrChain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BVMEigenDataLayrChainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BVMEigenDataLayrChain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMEigenDataLayrChainOwnershipTransferred)
				if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) ParseOwnershipTransferred(log types.Log) (*BVMEigenDataLayrChainOwnershipTransferred, error) {
	event := new(BVMEigenDataLayrChainOwnershipTransferred)
	if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMEigenDataLayrChainRollupStoreConfirmedIterator is returned from FilterRollupStoreConfirmed and is used to iterate over the raw logs and unpacked data for RollupStoreConfirmed events raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainRollupStoreConfirmedIterator struct {
	Event *BVMEigenDataLayrChainRollupStoreConfirmed // Event containing the contract specifics and raw log

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
func (it *BVMEigenDataLayrChainRollupStoreConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMEigenDataLayrChainRollupStoreConfirmed)
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
		it.Event = new(BVMEigenDataLayrChainRollupStoreConfirmed)
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
func (it *BVMEigenDataLayrChainRollupStoreConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMEigenDataLayrChainRollupStoreConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMEigenDataLayrChainRollupStoreConfirmed represents a RollupStoreConfirmed event raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainRollupStoreConfirmed struct {
	RollupStoreNumber uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreConfirmed is a free log retrieval operation binding the contract event 0xe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a9.
//
// Solidity: event RollupStoreConfirmed(uint32 rollupStoreNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) FilterRollupStoreConfirmed(opts *bind.FilterOpts) (*BVMEigenDataLayrChainRollupStoreConfirmedIterator, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.FilterLogs(opts, "RollupStoreConfirmed")
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainRollupStoreConfirmedIterator{contract: _BVMEigenDataLayrChain.contract, event: "RollupStoreConfirmed", logs: logs, sub: sub}, nil
}

// WatchRollupStoreConfirmed is a free log subscription operation binding the contract event 0xe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a9.
//
// Solidity: event RollupStoreConfirmed(uint32 rollupStoreNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) WatchRollupStoreConfirmed(opts *bind.WatchOpts, sink chan<- *BVMEigenDataLayrChainRollupStoreConfirmed) (event.Subscription, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.WatchLogs(opts, "RollupStoreConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMEigenDataLayrChainRollupStoreConfirmed)
				if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "RollupStoreConfirmed", log); err != nil {
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

// ParseRollupStoreConfirmed is a log parse operation binding the contract event 0xe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a9.
//
// Solidity: event RollupStoreConfirmed(uint32 rollupStoreNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) ParseRollupStoreConfirmed(log types.Log) (*BVMEigenDataLayrChainRollupStoreConfirmed, error) {
	event := new(BVMEigenDataLayrChainRollupStoreConfirmed)
	if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "RollupStoreConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMEigenDataLayrChainRollupStoreInitializedIterator is returned from FilterRollupStoreInitialized and is used to iterate over the raw logs and unpacked data for RollupStoreInitialized events raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainRollupStoreInitializedIterator struct {
	Event *BVMEigenDataLayrChainRollupStoreInitialized // Event containing the contract specifics and raw log

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
func (it *BVMEigenDataLayrChainRollupStoreInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMEigenDataLayrChainRollupStoreInitialized)
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
		it.Event = new(BVMEigenDataLayrChainRollupStoreInitialized)
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
func (it *BVMEigenDataLayrChainRollupStoreInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMEigenDataLayrChainRollupStoreInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMEigenDataLayrChainRollupStoreInitialized represents a RollupStoreInitialized event raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainRollupStoreInitialized struct {
	DataStoreId uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreInitialized is a free log retrieval operation binding the contract event 0x957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b5.
//
// Solidity: event RollupStoreInitialized(uint32 dataStoreId)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) FilterRollupStoreInitialized(opts *bind.FilterOpts) (*BVMEigenDataLayrChainRollupStoreInitializedIterator, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.FilterLogs(opts, "RollupStoreInitialized")
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainRollupStoreInitializedIterator{contract: _BVMEigenDataLayrChain.contract, event: "RollupStoreInitialized", logs: logs, sub: sub}, nil
}

// WatchRollupStoreInitialized is a free log subscription operation binding the contract event 0x957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b5.
//
// Solidity: event RollupStoreInitialized(uint32 dataStoreId)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) WatchRollupStoreInitialized(opts *bind.WatchOpts, sink chan<- *BVMEigenDataLayrChainRollupStoreInitialized) (event.Subscription, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.WatchLogs(opts, "RollupStoreInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMEigenDataLayrChainRollupStoreInitialized)
				if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "RollupStoreInitialized", log); err != nil {
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

// ParseRollupStoreInitialized is a log parse operation binding the contract event 0x957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b5.
//
// Solidity: event RollupStoreInitialized(uint32 dataStoreId)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) ParseRollupStoreInitialized(log types.Log) (*BVMEigenDataLayrChainRollupStoreInitialized, error) {
	event := new(BVMEigenDataLayrChainRollupStoreInitialized)
	if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "RollupStoreInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

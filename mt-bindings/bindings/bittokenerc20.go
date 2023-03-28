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

// BitTokenERC20MetaData contains all meta data concerning the BitTokenERC20 contract.
var BitTokenERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200129e3803806200129e833981016040819052620000349162000256565b818160036200004483826200034e565b5060046200005382826200034e565b5050506200008f336200006b6200009760201b60201c565b6200007b9060ff16600a6200052f565b620000899061271062000544565b6200009c565b505062000581565b601290565b6001600160a01b038216620000f75760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546200010b919062000566565b90915550506001600160a01b038216600090815260208190526040812080548392906200013a90849062000566565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b505050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001b157600080fd5b81516001600160401b0380821115620001ce57620001ce62000189565b604051601f8301601f19908116603f01168101908282118183101715620001f957620001f962000189565b816040528381526020925086838588010111156200021657600080fd5b600091505b838210156200023a57858201830151818301840152908201906200021b565b838211156200024c5760008385830101525b9695505050505050565b600080604083850312156200026a57600080fd5b82516001600160401b03808211156200028257600080fd5b62000290868387016200019f565b93506020850151915080821115620002a757600080fd5b50620002b6858286016200019f565b9150509250929050565b600181811c90821680620002d557607f821691505b602082108103620002f657634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200018457600081815260208120601f850160051c81016020861015620003255750805b601f850160051c820191505b81811015620003465782815560010162000331565b505050505050565b81516001600160401b038111156200036a576200036a62000189565b62000382816200037b8454620002c0565b84620002fc565b602080601f831160018114620003ba5760008415620003a15750858301515b600019600386901b1c1916600185901b17855562000346565b600085815260208120601f198616915b82811015620003eb57888601518255948401946001909101908401620003ca565b50858210156200040a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b600181815b80851115620004715781600019048211156200045557620004556200041a565b808516156200046357918102915b93841c939080029062000435565b509250929050565b6000826200048a5750600162000529565b81620004995750600062000529565b8160018114620004b25760028114620004bd57620004dd565b600191505062000529565b60ff841115620004d157620004d16200041a565b50506001821b62000529565b5060208310610133831016604e8410600b841016171562000502575081810a62000529565b6200050e838362000430565b80600019048211156200052557620005256200041a565b0290505b92915050565b60006200053d838362000479565b9392505050565b60008160001904831182151516156200056157620005616200041a565b500290565b600082198211156200057c576200057c6200041a565b500190565b610d0d80620005916000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c806370a0823111610081578063a457c2d71161005b578063a457c2d7146101b4578063a9059cbb146101c7578063dd62ed3e146101da57600080fd5b806370a082311461016157806395d89b4114610197578063a0712d681461019f57600080fd5b806323b872dd116100b257806323b872dd1461012c578063313ce5671461013f578063395093511461014e57600080fd5b806306fdde03146100d9578063095ea7b3146100f757806318160ddd1461011a575b600080fd5b6100e1610220565b6040516100ee9190610afe565b60405180910390f35b61010a610105366004610b9a565b6102b2565b60405190151581526020016100ee565b6002545b6040519081526020016100ee565b61010a61013a366004610bc4565b6102ca565b604051601281526020016100ee565b61010a61015c366004610b9a565b6102ee565b61011e61016f366004610c00565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100e161033a565b6101b26101ad366004610c22565b610349565b005b61010a6101c2366004610b9a565b6103c2565b61010a6101d5366004610b9a565b610493565b61011e6101e8366004610c3b565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461022f90610c6e565b80601f016020809104026020016040519081016040528092919081815260200182805461025b90610c6e565b80156102a85780601f1061027d576101008083540402835291602001916102a8565b820191906000526020600020905b81548152906001019060200180831161028b57829003601f168201915b5050505050905090565b6000336102c08185856104a1565b5060019392505050565b6000336102d8858285610654565b6102e385858561072b565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906102c09082908690610335908790610cc1565b6104a1565b60606004805461022f90610c6e565b336103b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064015b60405180910390fd5b6103bf33826109de565b50565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610486576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016103ac565b6102e382868684036104a1565b6000336102c081858561072b565b73ffffffffffffffffffffffffffffffffffffffff8316610543576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016103ac565b73ffffffffffffffffffffffffffffffffffffffff82166105e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016103ac565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107255781811015610718576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103ac565b61072584848484036104a1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166107ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016103ac565b73ffffffffffffffffffffffffffffffffffffffff8216610871576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016103ac565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610927576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016103ac565b73ffffffffffffffffffffffffffffffffffffffff80851660009081526020819052604080822085850390559185168152908120805484929061096b908490610cc1565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516109d191815260200190565b60405180910390a3610725565b73ffffffffffffffffffffffffffffffffffffffff8216610a5b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103ac565b8060026000828254610a6d9190610cc1565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610aa7908490610cc1565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b600060208083528351808285015260005b81811015610b2b57858101830151858201604001528201610b0f565b81811115610b3d576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b9557600080fd5b919050565b60008060408385031215610bad57600080fd5b610bb683610b71565b946020939093013593505050565b600080600060608486031215610bd957600080fd5b610be284610b71565b9250610bf060208501610b71565b9150604084013590509250925092565b600060208284031215610c1257600080fd5b610c1b82610b71565b9392505050565b600060208284031215610c3457600080fd5b5035919050565b60008060408385031215610c4e57600080fd5b610c5783610b71565b9150610c6560208401610b71565b90509250929050565b600181811c90821680610c8257607f821691505b602082108103610cbb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60008219821115610cfb577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50019056fea164736f6c634300080f000a",
}

// BitTokenERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use BitTokenERC20MetaData.ABI instead.
var BitTokenERC20ABI = BitTokenERC20MetaData.ABI

// BitTokenERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BitTokenERC20MetaData.Bin instead.
var BitTokenERC20Bin = BitTokenERC20MetaData.Bin

// DeployBitTokenERC20 deploys a new Ethereum contract, binding an instance of BitTokenERC20 to it.
func DeployBitTokenERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *BitTokenERC20, error) {
	parsed, err := BitTokenERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BitTokenERC20Bin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BitTokenERC20{BitTokenERC20Caller: BitTokenERC20Caller{contract: contract}, BitTokenERC20Transactor: BitTokenERC20Transactor{contract: contract}, BitTokenERC20Filterer: BitTokenERC20Filterer{contract: contract}}, nil
}

// BitTokenERC20 is an auto generated Go binding around an Ethereum contract.
type BitTokenERC20 struct {
	BitTokenERC20Caller     // Read-only binding to the contract
	BitTokenERC20Transactor // Write-only binding to the contract
	BitTokenERC20Filterer   // Log filterer for contract events
}

// BitTokenERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type BitTokenERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitTokenERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BitTokenERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitTokenERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitTokenERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitTokenERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitTokenERC20Session struct {
	Contract     *BitTokenERC20    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BitTokenERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitTokenERC20CallerSession struct {
	Contract *BitTokenERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BitTokenERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitTokenERC20TransactorSession struct {
	Contract     *BitTokenERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BitTokenERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type BitTokenERC20Raw struct {
	Contract *BitTokenERC20 // Generic contract binding to access the raw methods on
}

// BitTokenERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitTokenERC20CallerRaw struct {
	Contract *BitTokenERC20Caller // Generic read-only contract binding to access the raw methods on
}

// BitTokenERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitTokenERC20TransactorRaw struct {
	Contract *BitTokenERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBitTokenERC20 creates a new instance of BitTokenERC20, bound to a specific deployed contract.
func NewBitTokenERC20(address common.Address, backend bind.ContractBackend) (*BitTokenERC20, error) {
	contract, err := bindBitTokenERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20{BitTokenERC20Caller: BitTokenERC20Caller{contract: contract}, BitTokenERC20Transactor: BitTokenERC20Transactor{contract: contract}, BitTokenERC20Filterer: BitTokenERC20Filterer{contract: contract}}, nil
}

// NewBitTokenERC20Caller creates a new read-only instance of BitTokenERC20, bound to a specific deployed contract.
func NewBitTokenERC20Caller(address common.Address, caller bind.ContractCaller) (*BitTokenERC20Caller, error) {
	contract, err := bindBitTokenERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20Caller{contract: contract}, nil
}

// NewBitTokenERC20Transactor creates a new write-only instance of BitTokenERC20, bound to a specific deployed contract.
func NewBitTokenERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*BitTokenERC20Transactor, error) {
	contract, err := bindBitTokenERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20Transactor{contract: contract}, nil
}

// NewBitTokenERC20Filterer creates a new log filterer instance of BitTokenERC20, bound to a specific deployed contract.
func NewBitTokenERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*BitTokenERC20Filterer, error) {
	contract, err := bindBitTokenERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20Filterer{contract: contract}, nil
}

// bindBitTokenERC20 binds a generic wrapper to an already deployed contract.
func bindBitTokenERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BitTokenERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitTokenERC20 *BitTokenERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitTokenERC20.Contract.BitTokenERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitTokenERC20 *BitTokenERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.BitTokenERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitTokenERC20 *BitTokenERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.BitTokenERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitTokenERC20 *BitTokenERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitTokenERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitTokenERC20 *BitTokenERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitTokenERC20 *BitTokenERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BitTokenERC20.Contract.Allowance(&_BitTokenERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BitTokenERC20.Contract.Allowance(&_BitTokenERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _BitTokenERC20.Contract.BalanceOf(&_BitTokenERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BitTokenERC20.Contract.BalanceOf(&_BitTokenERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BitTokenERC20 *BitTokenERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BitTokenERC20 *BitTokenERC20Session) Decimals() (uint8, error) {
	return _BitTokenERC20.Contract.Decimals(&_BitTokenERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BitTokenERC20 *BitTokenERC20CallerSession) Decimals() (uint8, error) {
	return _BitTokenERC20.Contract.Decimals(&_BitTokenERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BitTokenERC20 *BitTokenERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BitTokenERC20 *BitTokenERC20Session) Name() (string, error) {
	return _BitTokenERC20.Contract.Name(&_BitTokenERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BitTokenERC20 *BitTokenERC20CallerSession) Name() (string, error) {
	return _BitTokenERC20.Contract.Name(&_BitTokenERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BitTokenERC20 *BitTokenERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BitTokenERC20 *BitTokenERC20Session) Symbol() (string, error) {
	return _BitTokenERC20.Contract.Symbol(&_BitTokenERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BitTokenERC20 *BitTokenERC20CallerSession) Symbol() (string, error) {
	return _BitTokenERC20.Contract.Symbol(&_BitTokenERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Session) TotalSupply() (*big.Int, error) {
	return _BitTokenERC20.Contract.TotalSupply(&_BitTokenERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _BitTokenERC20.Contract.TotalSupply(&_BitTokenERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Approve(&_BitTokenERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Approve(&_BitTokenERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.DecreaseAllowance(&_BitTokenERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.DecreaseAllowance(&_BitTokenERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.IncreaseAllowance(&_BitTokenERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.IncreaseAllowance(&_BitTokenERC20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_BitTokenERC20 *BitTokenERC20Transactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_BitTokenERC20 *BitTokenERC20Session) Mint(amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Mint(&_BitTokenERC20.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_BitTokenERC20 *BitTokenERC20TransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Mint(&_BitTokenERC20.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Transfer(&_BitTokenERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Transfer(&_BitTokenERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.TransferFrom(&_BitTokenERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.TransferFrom(&_BitTokenERC20.TransactOpts, from, to, amount)
}

// BitTokenERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BitTokenERC20 contract.
type BitTokenERC20ApprovalIterator struct {
	Event *BitTokenERC20Approval // Event containing the contract specifics and raw log

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
func (it *BitTokenERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitTokenERC20Approval)
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
		it.Event = new(BitTokenERC20Approval)
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
func (it *BitTokenERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitTokenERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitTokenERC20Approval represents a Approval event raised by the BitTokenERC20 contract.
type BitTokenERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BitTokenERC20 *BitTokenERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BitTokenERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BitTokenERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20ApprovalIterator{contract: _BitTokenERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BitTokenERC20 *BitTokenERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BitTokenERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BitTokenERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitTokenERC20Approval)
				if err := _BitTokenERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BitTokenERC20 *BitTokenERC20Filterer) ParseApproval(log types.Log) (*BitTokenERC20Approval, error) {
	event := new(BitTokenERC20Approval)
	if err := _BitTokenERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BitTokenERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BitTokenERC20 contract.
type BitTokenERC20TransferIterator struct {
	Event *BitTokenERC20Transfer // Event containing the contract specifics and raw log

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
func (it *BitTokenERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitTokenERC20Transfer)
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
		it.Event = new(BitTokenERC20Transfer)
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
func (it *BitTokenERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitTokenERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitTokenERC20Transfer represents a Transfer event raised by the BitTokenERC20 contract.
type BitTokenERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BitTokenERC20 *BitTokenERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BitTokenERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BitTokenERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20TransferIterator{contract: _BitTokenERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BitTokenERC20 *BitTokenERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BitTokenERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BitTokenERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitTokenERC20Transfer)
				if err := _BitTokenERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BitTokenERC20 *BitTokenERC20Filterer) ParseTransfer(log types.Log) (*BitTokenERC20Transfer, error) {
	event := new(BitTokenERC20Transfer)
	if err := _BitTokenERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

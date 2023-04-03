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

// BVMGasPriceOracleMetaData contains all meta data concerning the BVMGasPriceOracle contract.
var BVMGasPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ChargeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"DAGasPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"DASwitchUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"DecimalsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"GasPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"IsBurningUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"L1BaseFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"OverheadUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ScalarUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"IsBurning\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daSwitch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getL1Fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getL1GasUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1BaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"overhead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sccAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_charge\",\"type\":\"uint256\"}],\"name\":\"setCharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_daGasPrice\",\"type\":\"uint256\"}],\"name\":\"setDAGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_daSwitch\",\"type\":\"uint256\"}],\"name\":\"setDaSwitch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"setDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"setGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_isBurning\",\"type\":\"uint256\"}],\"name\":\"setIsBurning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseFee\",\"type\":\"uint256\"}],\"name\":\"setL1BaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_overhead\",\"type\":\"uint256\"}],\"name\":\"setOverhead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_scalar\",\"type\":\"uint256\"}],\"name\":\"setScalar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610ee4380380610ee483398101604081905261002f91610171565b61003833610047565b61004181610097565b506101a1565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61009f610115565b6001600160a01b0381166101095760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b61011281610047565b50565b6000546001600160a01b0316331461016f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610100565b565b60006020828403121561018357600080fd5b81516001600160a01b038116811461019a57600080fd5b9392505050565b610d34806101b06000396000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c8063715018a6116100e3578063de26c4a11161008c578063f45e65d811610066578063f45e65d814610328578063fc55b13814610331578063fe173b971461034457600080fd5b8063de26c4a1146102e2578063ea01cd36146102f5578063f2fde38b1461031557600080fd5b8063bede39b5116100bd578063bede39b5146102a9578063bf1fe420146102bc578063c7647832146102cf57600080fd5b8063715018a61461024f5780638c8885c8146102575780638da5cb5b1461026a57600080fd5b806345c51a3811610145578063551619131161011f578063551619131461022a5780635cbe497a14610233578063704655971461023c57600080fd5b806345c51a38146101fb57806349948e0e1461020e578063519b4bd31461022157600080fd5b8063288005781161017657806328800578146101ca578063313ce567146101df5780633577afc5146101e857600080fd5b80630c18c1621461019d5780630d1e43a0146101b95780630e6faf1e146101c1575b600080fd5b6101a660035481565b6040519081526020015b60405180910390f35b6006546101a6565b6101a6600a5481565b6101dd6101d83660046109a6565b61034d565b005b6101a660055481565b6101dd6101f63660046109a6565b61040b565b6101dd6102093660046109a6565b61044f565b6101a661021c3660046109ee565b61048c565b6101a660025481565b6101a660075481565b6101a660095481565b6101dd61024a3660046109a6565b6104e8565b6101dd610525565b6101dd6102653660046109a6565b610539565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b0565b6101dd6102b73660046109a6565b610576565b6101dd6102ca3660046109a6565b6105b3565b6101dd6102dd3660046109a6565b6105f0565b6101a66102f03660046109ee565b6106a1565b6008546102849073ffffffffffffffffffffffffffffffffffffffff1681565b6101dd610323366004610abd565b610748565b6101a660045481565b6101dd61033f3660046109a6565b6107ff565b6101a660015481565b6103556108b0565b808015806103635750806001145b6103ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f696e76616c69642076616c75652c6d7573742062652030206f7220310000000060448201526064015b60405180910390fd5b600a8290556040518281527f65cacb453bbeab72658947058c43b2a6c7dfcca1c9d96ba1bc470d346929b288906020015b60405180910390a15050565b6104136108b0565b60038190556040518181527f32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4906020015b60405180910390a150565b6104576108b0565b60098190556040518181527f5af81f5214eaf8c64101a8fde536abc770ef62af9e14d15e2b0b68760b2028f590602001610444565b600080610498836106a1565b90506000600254826104aa9190610b29565b90506000600554600a6104bd9190610c66565b90506000600454836104cf9190610b29565b905060006104dd8383610c72565b979650505050505050565b6104f06108b0565b60048190556040518181527f3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a90602001610444565b61052d6108b0565b6105376000610931565b565b6105416108b0565b60058190556040518181527fd68112a8707e326d08be3656b528c1bcc5bbbfc47f4177e2179b14d8640838c190602001610444565b61057e6108b0565b60028190556040518181527f351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c4490602001610444565b6105bb6108b0565b60018190556040518181527ffcdccc6074c6c42e4bd578aa9870c697dc976a270968452d2b8c8dc369fae39690602001610444565b6105f86108b0565b808015806106065750806001145b61066c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f696e76616c69642076616c75652c6d7573742062652030206f7220310000000060448201526064016103c5565b60068290556040518281527fd1eaae13a99b475ddca546a1b4a45052c66c14049997f44a1731a8e7167981a7906020016103ff565b600080805b8351811015610721578381815181106106c1576106c1610cad565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016600003610701576106fa600483610cdc565b915061070f565b61070c601083610cdc565b91505b8061071981610cef565b9150506106a6565b506000600354826107329190610cdc565b905061074081610440610cdc565b949350505050565b6107506108b0565b73ffffffffffffffffffffffffffffffffffffffff81166107f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103c5565b6107fc81610931565b50565b6108076108b0565b808015806108155750806001145b61087b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f696e76616c69642076616c75652c6d7573742062652030206f7220310000000060448201526064016103c5565b60078290556040518281527f49244d4195584d0644398167ca8caa7b98ee36b674e4b4d2a2640749b27eafb7906020016103ff565b60005473ffffffffffffffffffffffffffffffffffffffff163314610537576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103c5565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156109b857600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060208284031215610a0057600080fd5b813567ffffffffffffffff80821115610a1857600080fd5b818401915084601f830112610a2c57600080fd5b813581811115610a3e57610a3e6109bf565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610a8457610a846109bf565b81604052828152876020848701011115610a9d57600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208284031215610acf57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610af357600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082028115828204841417610b4057610b40610afa565b92915050565b600181815b80851115610b9f57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610b8557610b85610afa565b80851615610b9257918102915b93841c9390800290610b4b565b509250929050565b600082610bb657506001610b40565b81610bc357506000610b40565b8160018114610bd95760028114610be357610bff565b6001915050610b40565b60ff841115610bf457610bf4610afa565b50506001821b610b40565b5060208310610133831016604e8410600b8410161715610c22575081810a610b40565b610c2c8383610b46565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610c5e57610c5e610afa565b029392505050565b6000610af38383610ba7565b600082610ca8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80820180821115610b4057610b40610afa565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d2057610d20610afa565b506001019056fea164736f6c6343000811000a",
}

// BVMGasPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use BVMGasPriceOracleMetaData.ABI instead.
var BVMGasPriceOracleABI = BVMGasPriceOracleMetaData.ABI

// BVMGasPriceOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BVMGasPriceOracleMetaData.Bin instead.
var BVMGasPriceOracleBin = BVMGasPriceOracleMetaData.Bin

// DeployBVMGasPriceOracle deploys a new Ethereum contract, binding an instance of BVMGasPriceOracle to it.
func DeployBVMGasPriceOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *BVMGasPriceOracle, error) {
	parsed, err := BVMGasPriceOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BVMGasPriceOracleBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BVMGasPriceOracle{BVMGasPriceOracleCaller: BVMGasPriceOracleCaller{contract: contract}, BVMGasPriceOracleTransactor: BVMGasPriceOracleTransactor{contract: contract}, BVMGasPriceOracleFilterer: BVMGasPriceOracleFilterer{contract: contract}}, nil
}

// BVMGasPriceOracle is an auto generated Go binding around an Ethereum contract.
type BVMGasPriceOracle struct {
	BVMGasPriceOracleCaller     // Read-only binding to the contract
	BVMGasPriceOracleTransactor // Write-only binding to the contract
	BVMGasPriceOracleFilterer   // Log filterer for contract events
}

// BVMGasPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BVMGasPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMGasPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BVMGasPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMGasPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BVMGasPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMGasPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BVMGasPriceOracleSession struct {
	Contract     *BVMGasPriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BVMGasPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BVMGasPriceOracleCallerSession struct {
	Contract *BVMGasPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BVMGasPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BVMGasPriceOracleTransactorSession struct {
	Contract     *BVMGasPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BVMGasPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BVMGasPriceOracleRaw struct {
	Contract *BVMGasPriceOracle // Generic contract binding to access the raw methods on
}

// BVMGasPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BVMGasPriceOracleCallerRaw struct {
	Contract *BVMGasPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// BVMGasPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BVMGasPriceOracleTransactorRaw struct {
	Contract *BVMGasPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBVMGasPriceOracle creates a new instance of BVMGasPriceOracle, bound to a specific deployed contract.
func NewBVMGasPriceOracle(address common.Address, backend bind.ContractBackend) (*BVMGasPriceOracle, error) {
	contract, err := bindBVMGasPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracle{BVMGasPriceOracleCaller: BVMGasPriceOracleCaller{contract: contract}, BVMGasPriceOracleTransactor: BVMGasPriceOracleTransactor{contract: contract}, BVMGasPriceOracleFilterer: BVMGasPriceOracleFilterer{contract: contract}}, nil
}

// NewBVMGasPriceOracleCaller creates a new read-only instance of BVMGasPriceOracle, bound to a specific deployed contract.
func NewBVMGasPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*BVMGasPriceOracleCaller, error) {
	contract, err := bindBVMGasPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleCaller{contract: contract}, nil
}

// NewBVMGasPriceOracleTransactor creates a new write-only instance of BVMGasPriceOracle, bound to a specific deployed contract.
func NewBVMGasPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*BVMGasPriceOracleTransactor, error) {
	contract, err := bindBVMGasPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleTransactor{contract: contract}, nil
}

// NewBVMGasPriceOracleFilterer creates a new log filterer instance of BVMGasPriceOracle, bound to a specific deployed contract.
func NewBVMGasPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*BVMGasPriceOracleFilterer, error) {
	contract, err := bindBVMGasPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleFilterer{contract: contract}, nil
}

// bindBVMGasPriceOracle binds a generic wrapper to an already deployed contract.
func bindBVMGasPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BVMGasPriceOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMGasPriceOracle *BVMGasPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMGasPriceOracle.Contract.BVMGasPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMGasPriceOracle *BVMGasPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.BVMGasPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMGasPriceOracle *BVMGasPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.BVMGasPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMGasPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// IsBurning is a free data retrieval call binding the contract method 0x0d1e43a0.
//
// Solidity: function IsBurning() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) IsBurning(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "IsBurning")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsBurning is a free data retrieval call binding the contract method 0x0d1e43a0.
//
// Solidity: function IsBurning() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) IsBurning() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.IsBurning(&_BVMGasPriceOracle.CallOpts)
}

// IsBurning is a free data retrieval call binding the contract method 0x0d1e43a0.
//
// Solidity: function IsBurning() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) IsBurning() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.IsBurning(&_BVMGasPriceOracle.CallOpts)
}

// Charge is a free data retrieval call binding the contract method 0x55161913.
//
// Solidity: function charge() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) Charge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "charge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Charge is a free data retrieval call binding the contract method 0x55161913.
//
// Solidity: function charge() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) Charge() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.Charge(&_BVMGasPriceOracle.CallOpts)
}

// Charge is a free data retrieval call binding the contract method 0x55161913.
//
// Solidity: function charge() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) Charge() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.Charge(&_BVMGasPriceOracle.CallOpts)
}

// DaGasPrice is a free data retrieval call binding the contract method 0x5cbe497a.
//
// Solidity: function daGasPrice() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) DaGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "daGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DaGasPrice is a free data retrieval call binding the contract method 0x5cbe497a.
//
// Solidity: function daGasPrice() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) DaGasPrice() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.DaGasPrice(&_BVMGasPriceOracle.CallOpts)
}

// DaGasPrice is a free data retrieval call binding the contract method 0x5cbe497a.
//
// Solidity: function daGasPrice() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) DaGasPrice() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.DaGasPrice(&_BVMGasPriceOracle.CallOpts)
}

// DaSwitch is a free data retrieval call binding the contract method 0x0e6faf1e.
//
// Solidity: function daSwitch() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) DaSwitch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "daSwitch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DaSwitch is a free data retrieval call binding the contract method 0x0e6faf1e.
//
// Solidity: function daSwitch() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) DaSwitch() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.DaSwitch(&_BVMGasPriceOracle.CallOpts)
}

// DaSwitch is a free data retrieval call binding the contract method 0x0e6faf1e.
//
// Solidity: function daSwitch() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) DaSwitch() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.DaSwitch(&_BVMGasPriceOracle.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) Decimals() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.Decimals(&_BVMGasPriceOracle.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) Decimals() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.Decimals(&_BVMGasPriceOracle.CallOpts)
}

// GasPrice is a free data retrieval call binding the contract method 0xfe173b97.
//
// Solidity: function gasPrice() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) GasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "gasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasPrice is a free data retrieval call binding the contract method 0xfe173b97.
//
// Solidity: function gasPrice() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) GasPrice() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.GasPrice(&_BVMGasPriceOracle.CallOpts)
}

// GasPrice is a free data retrieval call binding the contract method 0xfe173b97.
//
// Solidity: function gasPrice() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) GasPrice() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.GasPrice(&_BVMGasPriceOracle.CallOpts)
}

// GetL1Fee is a free data retrieval call binding the contract method 0x49948e0e.
//
// Solidity: function getL1Fee(bytes _data) view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) GetL1Fee(opts *bind.CallOpts, _data []byte) (*big.Int, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "getL1Fee", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL1Fee is a free data retrieval call binding the contract method 0x49948e0e.
//
// Solidity: function getL1Fee(bytes _data) view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) GetL1Fee(_data []byte) (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.GetL1Fee(&_BVMGasPriceOracle.CallOpts, _data)
}

// GetL1Fee is a free data retrieval call binding the contract method 0x49948e0e.
//
// Solidity: function getL1Fee(bytes _data) view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) GetL1Fee(_data []byte) (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.GetL1Fee(&_BVMGasPriceOracle.CallOpts, _data)
}

// GetL1GasUsed is a free data retrieval call binding the contract method 0xde26c4a1.
//
// Solidity: function getL1GasUsed(bytes _data) view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) GetL1GasUsed(opts *bind.CallOpts, _data []byte) (*big.Int, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "getL1GasUsed", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL1GasUsed is a free data retrieval call binding the contract method 0xde26c4a1.
//
// Solidity: function getL1GasUsed(bytes _data) view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) GetL1GasUsed(_data []byte) (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.GetL1GasUsed(&_BVMGasPriceOracle.CallOpts, _data)
}

// GetL1GasUsed is a free data retrieval call binding the contract method 0xde26c4a1.
//
// Solidity: function getL1GasUsed(bytes _data) view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) GetL1GasUsed(_data []byte) (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.GetL1GasUsed(&_BVMGasPriceOracle.CallOpts, _data)
}

// L1BaseFee is a free data retrieval call binding the contract method 0x519b4bd3.
//
// Solidity: function l1BaseFee() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) L1BaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "l1BaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1BaseFee is a free data retrieval call binding the contract method 0x519b4bd3.
//
// Solidity: function l1BaseFee() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) L1BaseFee() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.L1BaseFee(&_BVMGasPriceOracle.CallOpts)
}

// L1BaseFee is a free data retrieval call binding the contract method 0x519b4bd3.
//
// Solidity: function l1BaseFee() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) L1BaseFee() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.L1BaseFee(&_BVMGasPriceOracle.CallOpts)
}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) Overhead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "overhead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) Overhead() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.Overhead(&_BVMGasPriceOracle.CallOpts)
}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) Overhead() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.Overhead(&_BVMGasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) Owner() (common.Address, error) {
	return _BVMGasPriceOracle.Contract.Owner(&_BVMGasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) Owner() (common.Address, error) {
	return _BVMGasPriceOracle.Contract.Owner(&_BVMGasPriceOracle.CallOpts)
}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) Scalar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "scalar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) Scalar() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.Scalar(&_BVMGasPriceOracle.CallOpts)
}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) Scalar() (*big.Int, error) {
	return _BVMGasPriceOracle.Contract.Scalar(&_BVMGasPriceOracle.CallOpts)
}

// SccAddress is a free data retrieval call binding the contract method 0xea01cd36.
//
// Solidity: function sccAddress() view returns(address)
func (_BVMGasPriceOracle *BVMGasPriceOracleCaller) SccAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMGasPriceOracle.contract.Call(opts, &out, "sccAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SccAddress is a free data retrieval call binding the contract method 0xea01cd36.
//
// Solidity: function sccAddress() view returns(address)
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) SccAddress() (common.Address, error) {
	return _BVMGasPriceOracle.Contract.SccAddress(&_BVMGasPriceOracle.CallOpts)
}

// SccAddress is a free data retrieval call binding the contract method 0xea01cd36.
//
// Solidity: function sccAddress() view returns(address)
func (_BVMGasPriceOracle *BVMGasPriceOracleCallerSession) SccAddress() (common.Address, error) {
	return _BVMGasPriceOracle.Contract.SccAddress(&_BVMGasPriceOracle.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMGasPriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.RenounceOwnership(&_BVMGasPriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.RenounceOwnership(&_BVMGasPriceOracle.TransactOpts)
}

// SetCharge is a paid mutator transaction binding the contract method 0xfc55b138.
//
// Solidity: function setCharge(uint256 _charge) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactor) SetCharge(opts *bind.TransactOpts, _charge *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.contract.Transact(opts, "setCharge", _charge)
}

// SetCharge is a paid mutator transaction binding the contract method 0xfc55b138.
//
// Solidity: function setCharge(uint256 _charge) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) SetCharge(_charge *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetCharge(&_BVMGasPriceOracle.TransactOpts, _charge)
}

// SetCharge is a paid mutator transaction binding the contract method 0xfc55b138.
//
// Solidity: function setCharge(uint256 _charge) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorSession) SetCharge(_charge *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetCharge(&_BVMGasPriceOracle.TransactOpts, _charge)
}

// SetDAGasPrice is a paid mutator transaction binding the contract method 0x45c51a38.
//
// Solidity: function setDAGasPrice(uint256 _daGasPrice) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactor) SetDAGasPrice(opts *bind.TransactOpts, _daGasPrice *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.contract.Transact(opts, "setDAGasPrice", _daGasPrice)
}

// SetDAGasPrice is a paid mutator transaction binding the contract method 0x45c51a38.
//
// Solidity: function setDAGasPrice(uint256 _daGasPrice) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) SetDAGasPrice(_daGasPrice *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetDAGasPrice(&_BVMGasPriceOracle.TransactOpts, _daGasPrice)
}

// SetDAGasPrice is a paid mutator transaction binding the contract method 0x45c51a38.
//
// Solidity: function setDAGasPrice(uint256 _daGasPrice) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorSession) SetDAGasPrice(_daGasPrice *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetDAGasPrice(&_BVMGasPriceOracle.TransactOpts, _daGasPrice)
}

// SetDaSwitch is a paid mutator transaction binding the contract method 0x28800578.
//
// Solidity: function setDaSwitch(uint256 _daSwitch) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactor) SetDaSwitch(opts *bind.TransactOpts, _daSwitch *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.contract.Transact(opts, "setDaSwitch", _daSwitch)
}

// SetDaSwitch is a paid mutator transaction binding the contract method 0x28800578.
//
// Solidity: function setDaSwitch(uint256 _daSwitch) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) SetDaSwitch(_daSwitch *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetDaSwitch(&_BVMGasPriceOracle.TransactOpts, _daSwitch)
}

// SetDaSwitch is a paid mutator transaction binding the contract method 0x28800578.
//
// Solidity: function setDaSwitch(uint256 _daSwitch) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorSession) SetDaSwitch(_daSwitch *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetDaSwitch(&_BVMGasPriceOracle.TransactOpts, _daSwitch)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 _decimals) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactor) SetDecimals(opts *bind.TransactOpts, _decimals *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.contract.Transact(opts, "setDecimals", _decimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 _decimals) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) SetDecimals(_decimals *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetDecimals(&_BVMGasPriceOracle.TransactOpts, _decimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 _decimals) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorSession) SetDecimals(_decimals *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetDecimals(&_BVMGasPriceOracle.TransactOpts, _decimals)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xbf1fe420.
//
// Solidity: function setGasPrice(uint256 _gasPrice) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactor) SetGasPrice(opts *bind.TransactOpts, _gasPrice *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.contract.Transact(opts, "setGasPrice", _gasPrice)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xbf1fe420.
//
// Solidity: function setGasPrice(uint256 _gasPrice) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) SetGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetGasPrice(&_BVMGasPriceOracle.TransactOpts, _gasPrice)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xbf1fe420.
//
// Solidity: function setGasPrice(uint256 _gasPrice) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorSession) SetGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetGasPrice(&_BVMGasPriceOracle.TransactOpts, _gasPrice)
}

// SetIsBurning is a paid mutator transaction binding the contract method 0xc7647832.
//
// Solidity: function setIsBurning(uint256 _isBurning) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactor) SetIsBurning(opts *bind.TransactOpts, _isBurning *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.contract.Transact(opts, "setIsBurning", _isBurning)
}

// SetIsBurning is a paid mutator transaction binding the contract method 0xc7647832.
//
// Solidity: function setIsBurning(uint256 _isBurning) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) SetIsBurning(_isBurning *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetIsBurning(&_BVMGasPriceOracle.TransactOpts, _isBurning)
}

// SetIsBurning is a paid mutator transaction binding the contract method 0xc7647832.
//
// Solidity: function setIsBurning(uint256 _isBurning) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorSession) SetIsBurning(_isBurning *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetIsBurning(&_BVMGasPriceOracle.TransactOpts, _isBurning)
}

// SetL1BaseFee is a paid mutator transaction binding the contract method 0xbede39b5.
//
// Solidity: function setL1BaseFee(uint256 _baseFee) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactor) SetL1BaseFee(opts *bind.TransactOpts, _baseFee *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.contract.Transact(opts, "setL1BaseFee", _baseFee)
}

// SetL1BaseFee is a paid mutator transaction binding the contract method 0xbede39b5.
//
// Solidity: function setL1BaseFee(uint256 _baseFee) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) SetL1BaseFee(_baseFee *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetL1BaseFee(&_BVMGasPriceOracle.TransactOpts, _baseFee)
}

// SetL1BaseFee is a paid mutator transaction binding the contract method 0xbede39b5.
//
// Solidity: function setL1BaseFee(uint256 _baseFee) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorSession) SetL1BaseFee(_baseFee *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetL1BaseFee(&_BVMGasPriceOracle.TransactOpts, _baseFee)
}

// SetOverhead is a paid mutator transaction binding the contract method 0x3577afc5.
//
// Solidity: function setOverhead(uint256 _overhead) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactor) SetOverhead(opts *bind.TransactOpts, _overhead *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.contract.Transact(opts, "setOverhead", _overhead)
}

// SetOverhead is a paid mutator transaction binding the contract method 0x3577afc5.
//
// Solidity: function setOverhead(uint256 _overhead) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) SetOverhead(_overhead *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetOverhead(&_BVMGasPriceOracle.TransactOpts, _overhead)
}

// SetOverhead is a paid mutator transaction binding the contract method 0x3577afc5.
//
// Solidity: function setOverhead(uint256 _overhead) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorSession) SetOverhead(_overhead *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetOverhead(&_BVMGasPriceOracle.TransactOpts, _overhead)
}

// SetScalar is a paid mutator transaction binding the contract method 0x70465597.
//
// Solidity: function setScalar(uint256 _scalar) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactor) SetScalar(opts *bind.TransactOpts, _scalar *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.contract.Transact(opts, "setScalar", _scalar)
}

// SetScalar is a paid mutator transaction binding the contract method 0x70465597.
//
// Solidity: function setScalar(uint256 _scalar) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) SetScalar(_scalar *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetScalar(&_BVMGasPriceOracle.TransactOpts, _scalar)
}

// SetScalar is a paid mutator transaction binding the contract method 0x70465597.
//
// Solidity: function setScalar(uint256 _scalar) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorSession) SetScalar(_scalar *big.Int) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.SetScalar(&_BVMGasPriceOracle.TransactOpts, _scalar)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BVMGasPriceOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.TransferOwnership(&_BVMGasPriceOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BVMGasPriceOracle *BVMGasPriceOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BVMGasPriceOracle.Contract.TransferOwnership(&_BVMGasPriceOracle.TransactOpts, newOwner)
}

// BVMGasPriceOracleChargeUpdatedIterator is returned from FilterChargeUpdated and is used to iterate over the raw logs and unpacked data for ChargeUpdated events raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleChargeUpdatedIterator struct {
	Event *BVMGasPriceOracleChargeUpdated // Event containing the contract specifics and raw log

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
func (it *BVMGasPriceOracleChargeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMGasPriceOracleChargeUpdated)
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
		it.Event = new(BVMGasPriceOracleChargeUpdated)
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
func (it *BVMGasPriceOracleChargeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMGasPriceOracleChargeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMGasPriceOracleChargeUpdated represents a ChargeUpdated event raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleChargeUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChargeUpdated is a free log retrieval operation binding the contract event 0x49244d4195584d0644398167ca8caa7b98ee36b674e4b4d2a2640749b27eafb7.
//
// Solidity: event ChargeUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) FilterChargeUpdated(opts *bind.FilterOpts) (*BVMGasPriceOracleChargeUpdatedIterator, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.FilterLogs(opts, "ChargeUpdated")
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleChargeUpdatedIterator{contract: _BVMGasPriceOracle.contract, event: "ChargeUpdated", logs: logs, sub: sub}, nil
}

// WatchChargeUpdated is a free log subscription operation binding the contract event 0x49244d4195584d0644398167ca8caa7b98ee36b674e4b4d2a2640749b27eafb7.
//
// Solidity: event ChargeUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) WatchChargeUpdated(opts *bind.WatchOpts, sink chan<- *BVMGasPriceOracleChargeUpdated) (event.Subscription, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.WatchLogs(opts, "ChargeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMGasPriceOracleChargeUpdated)
				if err := _BVMGasPriceOracle.contract.UnpackLog(event, "ChargeUpdated", log); err != nil {
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

// ParseChargeUpdated is a log parse operation binding the contract event 0x49244d4195584d0644398167ca8caa7b98ee36b674e4b4d2a2640749b27eafb7.
//
// Solidity: event ChargeUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) ParseChargeUpdated(log types.Log) (*BVMGasPriceOracleChargeUpdated, error) {
	event := new(BVMGasPriceOracleChargeUpdated)
	if err := _BVMGasPriceOracle.contract.UnpackLog(event, "ChargeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMGasPriceOracleDAGasPriceUpdatedIterator is returned from FilterDAGasPriceUpdated and is used to iterate over the raw logs and unpacked data for DAGasPriceUpdated events raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleDAGasPriceUpdatedIterator struct {
	Event *BVMGasPriceOracleDAGasPriceUpdated // Event containing the contract specifics and raw log

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
func (it *BVMGasPriceOracleDAGasPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMGasPriceOracleDAGasPriceUpdated)
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
		it.Event = new(BVMGasPriceOracleDAGasPriceUpdated)
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
func (it *BVMGasPriceOracleDAGasPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMGasPriceOracleDAGasPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMGasPriceOracleDAGasPriceUpdated represents a DAGasPriceUpdated event raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleDAGasPriceUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDAGasPriceUpdated is a free log retrieval operation binding the contract event 0x5af81f5214eaf8c64101a8fde536abc770ef62af9e14d15e2b0b68760b2028f5.
//
// Solidity: event DAGasPriceUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) FilterDAGasPriceUpdated(opts *bind.FilterOpts) (*BVMGasPriceOracleDAGasPriceUpdatedIterator, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.FilterLogs(opts, "DAGasPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleDAGasPriceUpdatedIterator{contract: _BVMGasPriceOracle.contract, event: "DAGasPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchDAGasPriceUpdated is a free log subscription operation binding the contract event 0x5af81f5214eaf8c64101a8fde536abc770ef62af9e14d15e2b0b68760b2028f5.
//
// Solidity: event DAGasPriceUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) WatchDAGasPriceUpdated(opts *bind.WatchOpts, sink chan<- *BVMGasPriceOracleDAGasPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.WatchLogs(opts, "DAGasPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMGasPriceOracleDAGasPriceUpdated)
				if err := _BVMGasPriceOracle.contract.UnpackLog(event, "DAGasPriceUpdated", log); err != nil {
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

// ParseDAGasPriceUpdated is a log parse operation binding the contract event 0x5af81f5214eaf8c64101a8fde536abc770ef62af9e14d15e2b0b68760b2028f5.
//
// Solidity: event DAGasPriceUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) ParseDAGasPriceUpdated(log types.Log) (*BVMGasPriceOracleDAGasPriceUpdated, error) {
	event := new(BVMGasPriceOracleDAGasPriceUpdated)
	if err := _BVMGasPriceOracle.contract.UnpackLog(event, "DAGasPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMGasPriceOracleDASwitchUpdatedIterator is returned from FilterDASwitchUpdated and is used to iterate over the raw logs and unpacked data for DASwitchUpdated events raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleDASwitchUpdatedIterator struct {
	Event *BVMGasPriceOracleDASwitchUpdated // Event containing the contract specifics and raw log

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
func (it *BVMGasPriceOracleDASwitchUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMGasPriceOracleDASwitchUpdated)
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
		it.Event = new(BVMGasPriceOracleDASwitchUpdated)
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
func (it *BVMGasPriceOracleDASwitchUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMGasPriceOracleDASwitchUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMGasPriceOracleDASwitchUpdated represents a DASwitchUpdated event raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleDASwitchUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDASwitchUpdated is a free log retrieval operation binding the contract event 0x65cacb453bbeab72658947058c43b2a6c7dfcca1c9d96ba1bc470d346929b288.
//
// Solidity: event DASwitchUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) FilterDASwitchUpdated(opts *bind.FilterOpts) (*BVMGasPriceOracleDASwitchUpdatedIterator, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.FilterLogs(opts, "DASwitchUpdated")
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleDASwitchUpdatedIterator{contract: _BVMGasPriceOracle.contract, event: "DASwitchUpdated", logs: logs, sub: sub}, nil
}

// WatchDASwitchUpdated is a free log subscription operation binding the contract event 0x65cacb453bbeab72658947058c43b2a6c7dfcca1c9d96ba1bc470d346929b288.
//
// Solidity: event DASwitchUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) WatchDASwitchUpdated(opts *bind.WatchOpts, sink chan<- *BVMGasPriceOracleDASwitchUpdated) (event.Subscription, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.WatchLogs(opts, "DASwitchUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMGasPriceOracleDASwitchUpdated)
				if err := _BVMGasPriceOracle.contract.UnpackLog(event, "DASwitchUpdated", log); err != nil {
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

// ParseDASwitchUpdated is a log parse operation binding the contract event 0x65cacb453bbeab72658947058c43b2a6c7dfcca1c9d96ba1bc470d346929b288.
//
// Solidity: event DASwitchUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) ParseDASwitchUpdated(log types.Log) (*BVMGasPriceOracleDASwitchUpdated, error) {
	event := new(BVMGasPriceOracleDASwitchUpdated)
	if err := _BVMGasPriceOracle.contract.UnpackLog(event, "DASwitchUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMGasPriceOracleDecimalsUpdatedIterator is returned from FilterDecimalsUpdated and is used to iterate over the raw logs and unpacked data for DecimalsUpdated events raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleDecimalsUpdatedIterator struct {
	Event *BVMGasPriceOracleDecimalsUpdated // Event containing the contract specifics and raw log

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
func (it *BVMGasPriceOracleDecimalsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMGasPriceOracleDecimalsUpdated)
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
		it.Event = new(BVMGasPriceOracleDecimalsUpdated)
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
func (it *BVMGasPriceOracleDecimalsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMGasPriceOracleDecimalsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMGasPriceOracleDecimalsUpdated represents a DecimalsUpdated event raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleDecimalsUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDecimalsUpdated is a free log retrieval operation binding the contract event 0xd68112a8707e326d08be3656b528c1bcc5bbbfc47f4177e2179b14d8640838c1.
//
// Solidity: event DecimalsUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) FilterDecimalsUpdated(opts *bind.FilterOpts) (*BVMGasPriceOracleDecimalsUpdatedIterator, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.FilterLogs(opts, "DecimalsUpdated")
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleDecimalsUpdatedIterator{contract: _BVMGasPriceOracle.contract, event: "DecimalsUpdated", logs: logs, sub: sub}, nil
}

// WatchDecimalsUpdated is a free log subscription operation binding the contract event 0xd68112a8707e326d08be3656b528c1bcc5bbbfc47f4177e2179b14d8640838c1.
//
// Solidity: event DecimalsUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) WatchDecimalsUpdated(opts *bind.WatchOpts, sink chan<- *BVMGasPriceOracleDecimalsUpdated) (event.Subscription, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.WatchLogs(opts, "DecimalsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMGasPriceOracleDecimalsUpdated)
				if err := _BVMGasPriceOracle.contract.UnpackLog(event, "DecimalsUpdated", log); err != nil {
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

// ParseDecimalsUpdated is a log parse operation binding the contract event 0xd68112a8707e326d08be3656b528c1bcc5bbbfc47f4177e2179b14d8640838c1.
//
// Solidity: event DecimalsUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) ParseDecimalsUpdated(log types.Log) (*BVMGasPriceOracleDecimalsUpdated, error) {
	event := new(BVMGasPriceOracleDecimalsUpdated)
	if err := _BVMGasPriceOracle.contract.UnpackLog(event, "DecimalsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMGasPriceOracleGasPriceUpdatedIterator is returned from FilterGasPriceUpdated and is used to iterate over the raw logs and unpacked data for GasPriceUpdated events raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleGasPriceUpdatedIterator struct {
	Event *BVMGasPriceOracleGasPriceUpdated // Event containing the contract specifics and raw log

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
func (it *BVMGasPriceOracleGasPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMGasPriceOracleGasPriceUpdated)
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
		it.Event = new(BVMGasPriceOracleGasPriceUpdated)
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
func (it *BVMGasPriceOracleGasPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMGasPriceOracleGasPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMGasPriceOracleGasPriceUpdated represents a GasPriceUpdated event raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleGasPriceUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGasPriceUpdated is a free log retrieval operation binding the contract event 0xfcdccc6074c6c42e4bd578aa9870c697dc976a270968452d2b8c8dc369fae396.
//
// Solidity: event GasPriceUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) FilterGasPriceUpdated(opts *bind.FilterOpts) (*BVMGasPriceOracleGasPriceUpdatedIterator, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.FilterLogs(opts, "GasPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleGasPriceUpdatedIterator{contract: _BVMGasPriceOracle.contract, event: "GasPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchGasPriceUpdated is a free log subscription operation binding the contract event 0xfcdccc6074c6c42e4bd578aa9870c697dc976a270968452d2b8c8dc369fae396.
//
// Solidity: event GasPriceUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) WatchGasPriceUpdated(opts *bind.WatchOpts, sink chan<- *BVMGasPriceOracleGasPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.WatchLogs(opts, "GasPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMGasPriceOracleGasPriceUpdated)
				if err := _BVMGasPriceOracle.contract.UnpackLog(event, "GasPriceUpdated", log); err != nil {
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

// ParseGasPriceUpdated is a log parse operation binding the contract event 0xfcdccc6074c6c42e4bd578aa9870c697dc976a270968452d2b8c8dc369fae396.
//
// Solidity: event GasPriceUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) ParseGasPriceUpdated(log types.Log) (*BVMGasPriceOracleGasPriceUpdated, error) {
	event := new(BVMGasPriceOracleGasPriceUpdated)
	if err := _BVMGasPriceOracle.contract.UnpackLog(event, "GasPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMGasPriceOracleIsBurningUpdatedIterator is returned from FilterIsBurningUpdated and is used to iterate over the raw logs and unpacked data for IsBurningUpdated events raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleIsBurningUpdatedIterator struct {
	Event *BVMGasPriceOracleIsBurningUpdated // Event containing the contract specifics and raw log

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
func (it *BVMGasPriceOracleIsBurningUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMGasPriceOracleIsBurningUpdated)
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
		it.Event = new(BVMGasPriceOracleIsBurningUpdated)
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
func (it *BVMGasPriceOracleIsBurningUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMGasPriceOracleIsBurningUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMGasPriceOracleIsBurningUpdated represents a IsBurningUpdated event raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleIsBurningUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterIsBurningUpdated is a free log retrieval operation binding the contract event 0xd1eaae13a99b475ddca546a1b4a45052c66c14049997f44a1731a8e7167981a7.
//
// Solidity: event IsBurningUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) FilterIsBurningUpdated(opts *bind.FilterOpts) (*BVMGasPriceOracleIsBurningUpdatedIterator, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.FilterLogs(opts, "IsBurningUpdated")
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleIsBurningUpdatedIterator{contract: _BVMGasPriceOracle.contract, event: "IsBurningUpdated", logs: logs, sub: sub}, nil
}

// WatchIsBurningUpdated is a free log subscription operation binding the contract event 0xd1eaae13a99b475ddca546a1b4a45052c66c14049997f44a1731a8e7167981a7.
//
// Solidity: event IsBurningUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) WatchIsBurningUpdated(opts *bind.WatchOpts, sink chan<- *BVMGasPriceOracleIsBurningUpdated) (event.Subscription, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.WatchLogs(opts, "IsBurningUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMGasPriceOracleIsBurningUpdated)
				if err := _BVMGasPriceOracle.contract.UnpackLog(event, "IsBurningUpdated", log); err != nil {
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

// ParseIsBurningUpdated is a log parse operation binding the contract event 0xd1eaae13a99b475ddca546a1b4a45052c66c14049997f44a1731a8e7167981a7.
//
// Solidity: event IsBurningUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) ParseIsBurningUpdated(log types.Log) (*BVMGasPriceOracleIsBurningUpdated, error) {
	event := new(BVMGasPriceOracleIsBurningUpdated)
	if err := _BVMGasPriceOracle.contract.UnpackLog(event, "IsBurningUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMGasPriceOracleL1BaseFeeUpdatedIterator is returned from FilterL1BaseFeeUpdated and is used to iterate over the raw logs and unpacked data for L1BaseFeeUpdated events raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleL1BaseFeeUpdatedIterator struct {
	Event *BVMGasPriceOracleL1BaseFeeUpdated // Event containing the contract specifics and raw log

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
func (it *BVMGasPriceOracleL1BaseFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMGasPriceOracleL1BaseFeeUpdated)
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
		it.Event = new(BVMGasPriceOracleL1BaseFeeUpdated)
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
func (it *BVMGasPriceOracleL1BaseFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMGasPriceOracleL1BaseFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMGasPriceOracleL1BaseFeeUpdated represents a L1BaseFeeUpdated event raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleL1BaseFeeUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterL1BaseFeeUpdated is a free log retrieval operation binding the contract event 0x351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c44.
//
// Solidity: event L1BaseFeeUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) FilterL1BaseFeeUpdated(opts *bind.FilterOpts) (*BVMGasPriceOracleL1BaseFeeUpdatedIterator, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.FilterLogs(opts, "L1BaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleL1BaseFeeUpdatedIterator{contract: _BVMGasPriceOracle.contract, event: "L1BaseFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchL1BaseFeeUpdated is a free log subscription operation binding the contract event 0x351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c44.
//
// Solidity: event L1BaseFeeUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) WatchL1BaseFeeUpdated(opts *bind.WatchOpts, sink chan<- *BVMGasPriceOracleL1BaseFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.WatchLogs(opts, "L1BaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMGasPriceOracleL1BaseFeeUpdated)
				if err := _BVMGasPriceOracle.contract.UnpackLog(event, "L1BaseFeeUpdated", log); err != nil {
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

// ParseL1BaseFeeUpdated is a log parse operation binding the contract event 0x351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c44.
//
// Solidity: event L1BaseFeeUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) ParseL1BaseFeeUpdated(log types.Log) (*BVMGasPriceOracleL1BaseFeeUpdated, error) {
	event := new(BVMGasPriceOracleL1BaseFeeUpdated)
	if err := _BVMGasPriceOracle.contract.UnpackLog(event, "L1BaseFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMGasPriceOracleOverheadUpdatedIterator is returned from FilterOverheadUpdated and is used to iterate over the raw logs and unpacked data for OverheadUpdated events raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleOverheadUpdatedIterator struct {
	Event *BVMGasPriceOracleOverheadUpdated // Event containing the contract specifics and raw log

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
func (it *BVMGasPriceOracleOverheadUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMGasPriceOracleOverheadUpdated)
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
		it.Event = new(BVMGasPriceOracleOverheadUpdated)
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
func (it *BVMGasPriceOracleOverheadUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMGasPriceOracleOverheadUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMGasPriceOracleOverheadUpdated represents a OverheadUpdated event raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleOverheadUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOverheadUpdated is a free log retrieval operation binding the contract event 0x32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4.
//
// Solidity: event OverheadUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) FilterOverheadUpdated(opts *bind.FilterOpts) (*BVMGasPriceOracleOverheadUpdatedIterator, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.FilterLogs(opts, "OverheadUpdated")
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleOverheadUpdatedIterator{contract: _BVMGasPriceOracle.contract, event: "OverheadUpdated", logs: logs, sub: sub}, nil
}

// WatchOverheadUpdated is a free log subscription operation binding the contract event 0x32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4.
//
// Solidity: event OverheadUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) WatchOverheadUpdated(opts *bind.WatchOpts, sink chan<- *BVMGasPriceOracleOverheadUpdated) (event.Subscription, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.WatchLogs(opts, "OverheadUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMGasPriceOracleOverheadUpdated)
				if err := _BVMGasPriceOracle.contract.UnpackLog(event, "OverheadUpdated", log); err != nil {
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

// ParseOverheadUpdated is a log parse operation binding the contract event 0x32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4.
//
// Solidity: event OverheadUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) ParseOverheadUpdated(log types.Log) (*BVMGasPriceOracleOverheadUpdated, error) {
	event := new(BVMGasPriceOracleOverheadUpdated)
	if err := _BVMGasPriceOracle.contract.UnpackLog(event, "OverheadUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMGasPriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleOwnershipTransferredIterator struct {
	Event *BVMGasPriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BVMGasPriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMGasPriceOracleOwnershipTransferred)
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
		it.Event = new(BVMGasPriceOracleOwnershipTransferred)
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
func (it *BVMGasPriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMGasPriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMGasPriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BVMGasPriceOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BVMGasPriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleOwnershipTransferredIterator{contract: _BVMGasPriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BVMGasPriceOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BVMGasPriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMGasPriceOracleOwnershipTransferred)
				if err := _BVMGasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*BVMGasPriceOracleOwnershipTransferred, error) {
	event := new(BVMGasPriceOracleOwnershipTransferred)
	if err := _BVMGasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMGasPriceOracleScalarUpdatedIterator is returned from FilterScalarUpdated and is used to iterate over the raw logs and unpacked data for ScalarUpdated events raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleScalarUpdatedIterator struct {
	Event *BVMGasPriceOracleScalarUpdated // Event containing the contract specifics and raw log

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
func (it *BVMGasPriceOracleScalarUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMGasPriceOracleScalarUpdated)
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
		it.Event = new(BVMGasPriceOracleScalarUpdated)
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
func (it *BVMGasPriceOracleScalarUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMGasPriceOracleScalarUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMGasPriceOracleScalarUpdated represents a ScalarUpdated event raised by the BVMGasPriceOracle contract.
type BVMGasPriceOracleScalarUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterScalarUpdated is a free log retrieval operation binding the contract event 0x3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a.
//
// Solidity: event ScalarUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) FilterScalarUpdated(opts *bind.FilterOpts) (*BVMGasPriceOracleScalarUpdatedIterator, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.FilterLogs(opts, "ScalarUpdated")
	if err != nil {
		return nil, err
	}
	return &BVMGasPriceOracleScalarUpdatedIterator{contract: _BVMGasPriceOracle.contract, event: "ScalarUpdated", logs: logs, sub: sub}, nil
}

// WatchScalarUpdated is a free log subscription operation binding the contract event 0x3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a.
//
// Solidity: event ScalarUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) WatchScalarUpdated(opts *bind.WatchOpts, sink chan<- *BVMGasPriceOracleScalarUpdated) (event.Subscription, error) {

	logs, sub, err := _BVMGasPriceOracle.contract.WatchLogs(opts, "ScalarUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMGasPriceOracleScalarUpdated)
				if err := _BVMGasPriceOracle.contract.UnpackLog(event, "ScalarUpdated", log); err != nil {
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

// ParseScalarUpdated is a log parse operation binding the contract event 0x3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a.
//
// Solidity: event ScalarUpdated(uint256 arg0)
func (_BVMGasPriceOracle *BVMGasPriceOracleFilterer) ParseScalarUpdated(log types.Log) (*BVMGasPriceOracleScalarUpdated, error) {
	event := new(BVMGasPriceOracleScalarUpdated)
	if err := _BVMGasPriceOracle.contract.UnpackLog(event, "ScalarUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

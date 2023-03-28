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

// BVMBITMetaData contains all meta data concerning the BVMBIT contract.
var BVMBITMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REMOTE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b50734200000000000000000000000000000000000010731a4b46696b2bb4794eb3d4c26f1c55f9170fa4c5604051806040016040528060098152602001682134ba102a37b5b2b760b91b8152506040518060400160405280600381526020016210925560ea1b81525060016000808484816003908162000093919062000171565b506004620000a2828262000171565b50505060809290925260a05260c05250506001600160a01b0390811660e05216610100526200023d565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000f757607f821691505b6020821081036200011857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200016c57600081815260208120601f850160051c81016020861015620001475750805b601f850160051c820191505b81811015620001685782815560010162000153565b5050505b505050565b81516001600160401b038111156200018d576200018d620000cc565b620001a5816200019e8454620000e2565b846200011e565b602080601f831160018114620001dd5760008415620001c45750858301515b600019600386901b1c1916600185901b17855562000168565b600085815260208120601f198616915b828110156200020e57888601518255948401946001909101908401620001ed565b50858210156200022d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e051610100516113a76200029e600039600081816102f50152818161038a0152818161075f01526109340152600081816101a9015261031b015260006108c30152600061089a0152600061087101526113a76000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e1461033f578063e78cea92146102f3578063ee9a31a21461038557600080fd5b8063ae1f6aaf146102f3578063c01e1bd614610319578063d6c0b2c41461031957600080fd5b80639dc29fac116100bd5780639dc29fac146102ba578063a457c2d7146102cd578063a9059cbb146102e057600080fd5b806370a082311461027c57806395d89b41146102b257600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461024c57806340c10f191461025f57806354fd4d501461027457600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a366004610fdf565b6103ac565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f861049d565b60405161019b9190611058565b61018f6102133660046110d2565b61052f565b6002545b60405190815260200161019b565b61018f6102383660046110fc565b6105e5565b6040516012815260200161019b565b61018f61025a3660046110d2565b610696565b61027261026d3660046110d2565b610747565b005b6101f861086a565b61021c61028a366004611138565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101f861090d565b6102726102c83660046110d2565b61091c565b61018f6102db3660046110d2565b610a33565b61018f6102ee3660046110d2565b610ae4565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c61034d366004611153565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000851683148061046557507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b8061049457507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104ac90611186565b80601f01602080910402602001604051908101604052809291908181526020018280546104d890611186565b80156105255780601f106104fa57610100808354040283529160200191610525565b820191906000526020600020905b81548152906001019060200180831161050857829003601f168201915b5050505050905090565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f42564d5f4249543a20617070726f76652069732064697361626c65642070656e60448201527f64696e67206675727468657220636f6d6d756e6974792064697363757373696f60648201527f6e2e000000000000000000000000000000000000000000000000000000000000608482015260009060a4015b60405180910390fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604760248201527f42564d5f4249543a207472616e7366657246726f6d2069732064697361626c6560448201527f642070656e64696e67206675727468657220636f6d6d756e697479206469736360648201527f757373696f6e2e00000000000000000000000000000000000000000000000000608482015260009060a4016105dc565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604c60248201527f42564d5f4249543a20696e637265617365416c6c6f77616e636520697320646960448201527f7361626c65642070656e64696e67206675727468657220636f6d6d756e69747960648201527f2064697363757373696f6e2e0000000000000000000000000000000000000000608482015260009060a4016105dc565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461080c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4d616e746c654d696e7461626c6545524332303a206f6e6c792062726964676560448201527f2063616e206d696e7420616e64206275726e000000000000000000000000000060648201526084016105dc565b6108168282610b95565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161085e91815260200190565b60405180910390a25050565b60606108957f0000000000000000000000000000000000000000000000000000000000000000610cb5565b6108be7f0000000000000000000000000000000000000000000000000000000000000000610cb5565b6108e77f0000000000000000000000000000000000000000000000000000000000000000610cb5565b6040516020016108f9939291906111d9565b604051602081830303815290604052905090565b6060600480546104ac90611186565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146109e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4d616e746c654d696e7461626c6545524332303a206f6e6c792062726964676560448201527f2063616e206d696e7420616e64206275726e000000000000000000000000000060648201526084016105dc565b6109eb8282610df2565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405161085e91815260200190565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604c60248201527f42564d5f4249543a206465637265617365416c6c6f77616e636520697320646960448201527f7361626c65642070656e64696e67206675727468657220636f6d6d756e69747960648201527f2064697363757373696f6e2e0000000000000000000000000000000000000000608482015260009060a4016105dc565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f42564d5f4249543a207472616e736665722069732064697361626c656420706560448201527f6e64696e67206675727468657220636f6d6d756e69747920646973637573736960648201527f6f6e2e0000000000000000000000000000000000000000000000000000000000608482015260009060a4016105dc565b73ffffffffffffffffffffffffffffffffffffffff8216610c12576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016105dc565b8060026000828254610c24919061127e565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610c5e90849061127e565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b606081600003610cf857505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610d225780610d0c81611296565b9150610d1b9050600a836112fd565b9150610cfc565b60008167ffffffffffffffff811115610d3d57610d3d611311565b6040519080825280601f01601f191660200182016040528015610d67576020820181803683370190505b5090505b8415610dea57610d7c600183611340565b9150610d89600a86611357565b610d9490603061127e565b60f81b818381518110610da957610da961136b565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610de3600a866112fd565b9450610d6b565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff8216610e95576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016105dc565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015610f4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016105dc565b73ffffffffffffffffffffffffffffffffffffffff83166000908152602081905260408120838303905560028054849290610f87908490611340565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3505050565b600060208284031215610ff157600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461102157600080fd5b9392505050565b60005b8381101561104357818101518382015260200161102b565b83811115611052576000848401525b50505050565b6020815260008251806020840152611077816040850160208701611028565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146110cd57600080fd5b919050565b600080604083850312156110e557600080fd5b6110ee836110a9565b946020939093013593505050565b60008060006060848603121561111157600080fd5b61111a846110a9565b9250611128602085016110a9565b9150604084013590509250925092565b60006020828403121561114a57600080fd5b611021826110a9565b6000806040838503121561116657600080fd5b61116f836110a9565b915061117d602084016110a9565b90509250929050565b600181811c9082168061119a57607f821691505b6020821081036111d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b600084516111eb818460208901611028565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611227816001850160208a01611028565b60019201918201528351611242816002840160208801611028565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156112915761129161124f565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036112c7576112c761124f565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261130c5761130c6112ce565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156113525761135261124f565b500390565b600082611366576113666112ce565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a",
}

// BVMBITABI is the input ABI used to generate the binding from.
// Deprecated: Use BVMBITMetaData.ABI instead.
var BVMBITABI = BVMBITMetaData.ABI

// BVMBITBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BVMBITMetaData.Bin instead.
var BVMBITBin = BVMBITMetaData.Bin

// DeployBVMBIT deploys a new Ethereum contract, binding an instance of BVMBIT to it.
func DeployBVMBIT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BVMBIT, error) {
	parsed, err := BVMBITMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BVMBITBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BVMBIT{BVMBITCaller: BVMBITCaller{contract: contract}, BVMBITTransactor: BVMBITTransactor{contract: contract}, BVMBITFilterer: BVMBITFilterer{contract: contract}}, nil
}

// BVMBIT is an auto generated Go binding around an Ethereum contract.
type BVMBIT struct {
	BVMBITCaller     // Read-only binding to the contract
	BVMBITTransactor // Write-only binding to the contract
	BVMBITFilterer   // Log filterer for contract events
}

// BVMBITCaller is an auto generated read-only Go binding around an Ethereum contract.
type BVMBITCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMBITTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BVMBITTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMBITFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BVMBITFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BVMBITSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BVMBITSession struct {
	Contract     *BVMBIT           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BVMBITCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BVMBITCallerSession struct {
	Contract *BVMBITCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BVMBITTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BVMBITTransactorSession struct {
	Contract     *BVMBITTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BVMBITRaw is an auto generated low-level Go binding around an Ethereum contract.
type BVMBITRaw struct {
	Contract *BVMBIT // Generic contract binding to access the raw methods on
}

// BVMBITCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BVMBITCallerRaw struct {
	Contract *BVMBITCaller // Generic read-only contract binding to access the raw methods on
}

// BVMBITTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BVMBITTransactorRaw struct {
	Contract *BVMBITTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBVMBIT creates a new instance of BVMBIT, bound to a specific deployed contract.
func NewBVMBIT(address common.Address, backend bind.ContractBackend) (*BVMBIT, error) {
	contract, err := bindBVMBIT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BVMBIT{BVMBITCaller: BVMBITCaller{contract: contract}, BVMBITTransactor: BVMBITTransactor{contract: contract}, BVMBITFilterer: BVMBITFilterer{contract: contract}}, nil
}

// NewBVMBITCaller creates a new read-only instance of BVMBIT, bound to a specific deployed contract.
func NewBVMBITCaller(address common.Address, caller bind.ContractCaller) (*BVMBITCaller, error) {
	contract, err := bindBVMBIT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BVMBITCaller{contract: contract}, nil
}

// NewBVMBITTransactor creates a new write-only instance of BVMBIT, bound to a specific deployed contract.
func NewBVMBITTransactor(address common.Address, transactor bind.ContractTransactor) (*BVMBITTransactor, error) {
	contract, err := bindBVMBIT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BVMBITTransactor{contract: contract}, nil
}

// NewBVMBITFilterer creates a new log filterer instance of BVMBIT, bound to a specific deployed contract.
func NewBVMBITFilterer(address common.Address, filterer bind.ContractFilterer) (*BVMBITFilterer, error) {
	contract, err := bindBVMBIT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BVMBITFilterer{contract: contract}, nil
}

// bindBVMBIT binds a generic wrapper to an already deployed contract.
func bindBVMBIT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BVMBITABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMBIT *BVMBITRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMBIT.Contract.BVMBITCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMBIT *BVMBITRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMBIT.Contract.BVMBITTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMBIT *BVMBITRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMBIT.Contract.BVMBITTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BVMBIT *BVMBITCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BVMBIT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BVMBIT *BVMBITTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BVMBIT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BVMBIT *BVMBITTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BVMBIT.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_BVMBIT *BVMBITCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_BVMBIT *BVMBITSession) BRIDGE() (common.Address, error) {
	return _BVMBIT.Contract.BRIDGE(&_BVMBIT.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_BVMBIT *BVMBITCallerSession) BRIDGE() (common.Address, error) {
	return _BVMBIT.Contract.BRIDGE(&_BVMBIT.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_BVMBIT *BVMBITCaller) REMOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "REMOTE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_BVMBIT *BVMBITSession) REMOTETOKEN() (common.Address, error) {
	return _BVMBIT.Contract.REMOTETOKEN(&_BVMBIT.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_BVMBIT *BVMBITCallerSession) REMOTETOKEN() (common.Address, error) {
	return _BVMBIT.Contract.REMOTETOKEN(&_BVMBIT.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BVMBIT *BVMBITCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BVMBIT *BVMBITSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BVMBIT.Contract.Allowance(&_BVMBIT.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BVMBIT *BVMBITCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BVMBIT.Contract.Allowance(&_BVMBIT.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BVMBIT *BVMBITCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BVMBIT *BVMBITSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BVMBIT.Contract.BalanceOf(&_BVMBIT.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BVMBIT *BVMBITCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BVMBIT.Contract.BalanceOf(&_BVMBIT.CallOpts, account)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BVMBIT *BVMBITCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BVMBIT *BVMBITSession) Bridge() (common.Address, error) {
	return _BVMBIT.Contract.Bridge(&_BVMBIT.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BVMBIT *BVMBITCallerSession) Bridge() (common.Address, error) {
	return _BVMBIT.Contract.Bridge(&_BVMBIT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BVMBIT *BVMBITCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BVMBIT *BVMBITSession) Decimals() (uint8, error) {
	return _BVMBIT.Contract.Decimals(&_BVMBIT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BVMBIT *BVMBITCallerSession) Decimals() (uint8, error) {
	return _BVMBIT.Contract.Decimals(&_BVMBIT.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_BVMBIT *BVMBITCaller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_BVMBIT *BVMBITSession) L1Token() (common.Address, error) {
	return _BVMBIT.Contract.L1Token(&_BVMBIT.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_BVMBIT *BVMBITCallerSession) L1Token() (common.Address, error) {
	return _BVMBIT.Contract.L1Token(&_BVMBIT.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_BVMBIT *BVMBITCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_BVMBIT *BVMBITSession) L2Bridge() (common.Address, error) {
	return _BVMBIT.Contract.L2Bridge(&_BVMBIT.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_BVMBIT *BVMBITCallerSession) L2Bridge() (common.Address, error) {
	return _BVMBIT.Contract.L2Bridge(&_BVMBIT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BVMBIT *BVMBITCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BVMBIT *BVMBITSession) Name() (string, error) {
	return _BVMBIT.Contract.Name(&_BVMBIT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BVMBIT *BVMBITCallerSession) Name() (string, error) {
	return _BVMBIT.Contract.Name(&_BVMBIT.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_BVMBIT *BVMBITCaller) RemoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "remoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_BVMBIT *BVMBITSession) RemoteToken() (common.Address, error) {
	return _BVMBIT.Contract.RemoteToken(&_BVMBIT.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_BVMBIT *BVMBITCallerSession) RemoteToken() (common.Address, error) {
	return _BVMBIT.Contract.RemoteToken(&_BVMBIT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_BVMBIT *BVMBITCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_BVMBIT *BVMBITSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _BVMBIT.Contract.SupportsInterface(&_BVMBIT.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_BVMBIT *BVMBITCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _BVMBIT.Contract.SupportsInterface(&_BVMBIT.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BVMBIT *BVMBITCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BVMBIT *BVMBITSession) Symbol() (string, error) {
	return _BVMBIT.Contract.Symbol(&_BVMBIT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BVMBIT *BVMBITCallerSession) Symbol() (string, error) {
	return _BVMBIT.Contract.Symbol(&_BVMBIT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BVMBIT *BVMBITCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BVMBIT *BVMBITSession) TotalSupply() (*big.Int, error) {
	return _BVMBIT.Contract.TotalSupply(&_BVMBIT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BVMBIT *BVMBITCallerSession) TotalSupply() (*big.Int, error) {
	return _BVMBIT.Contract.TotalSupply(&_BVMBIT.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BVMBIT *BVMBITCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BVMBIT.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BVMBIT *BVMBITSession) Version() (string, error) {
	return _BVMBIT.Contract.Version(&_BVMBIT.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BVMBIT *BVMBITCallerSession) Version() (string, error) {
	return _BVMBIT.Contract.Version(&_BVMBIT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BVMBIT *BVMBITTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BVMBIT *BVMBITSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.Approve(&_BVMBIT.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BVMBIT *BVMBITTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.Approve(&_BVMBIT.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_BVMBIT *BVMBITTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_BVMBIT *BVMBITSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.Burn(&_BVMBIT.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_BVMBIT *BVMBITTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.Burn(&_BVMBIT.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BVMBIT *BVMBITTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BVMBIT.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BVMBIT *BVMBITSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.DecreaseAllowance(&_BVMBIT.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BVMBIT *BVMBITTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.DecreaseAllowance(&_BVMBIT.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BVMBIT *BVMBITTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BVMBIT.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BVMBIT *BVMBITSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.IncreaseAllowance(&_BVMBIT.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BVMBIT *BVMBITTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.IncreaseAllowance(&_BVMBIT.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_BVMBIT *BVMBITTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_BVMBIT *BVMBITSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.Mint(&_BVMBIT.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_BVMBIT *BVMBITTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.Mint(&_BVMBIT.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BVMBIT *BVMBITTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BVMBIT *BVMBITSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.Transfer(&_BVMBIT.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BVMBIT *BVMBITTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.Transfer(&_BVMBIT.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BVMBIT *BVMBITTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BVMBIT *BVMBITSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.TransferFrom(&_BVMBIT.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BVMBIT *BVMBITTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BVMBIT.Contract.TransferFrom(&_BVMBIT.TransactOpts, sender, recipient, amount)
}

// BVMBITApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BVMBIT contract.
type BVMBITApprovalIterator struct {
	Event *BVMBITApproval // Event containing the contract specifics and raw log

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
func (it *BVMBITApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMBITApproval)
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
		it.Event = new(BVMBITApproval)
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
func (it *BVMBITApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMBITApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMBITApproval represents a Approval event raised by the BVMBIT contract.
type BVMBITApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BVMBIT *BVMBITFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BVMBITApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BVMBIT.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BVMBITApprovalIterator{contract: _BVMBIT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BVMBIT *BVMBITFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BVMBITApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BVMBIT.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMBITApproval)
				if err := _BVMBIT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BVMBIT *BVMBITFilterer) ParseApproval(log types.Log) (*BVMBITApproval, error) {
	event := new(BVMBITApproval)
	if err := _BVMBIT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMBITBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BVMBIT contract.
type BVMBITBurnIterator struct {
	Event *BVMBITBurn // Event containing the contract specifics and raw log

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
func (it *BVMBITBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMBITBurn)
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
		it.Event = new(BVMBITBurn)
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
func (it *BVMBITBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMBITBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMBITBurn represents a Burn event raised by the BVMBIT contract.
type BVMBITBurn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_BVMBIT *BVMBITFilterer) FilterBurn(opts *bind.FilterOpts, account []common.Address) (*BVMBITBurnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BVMBIT.contract.FilterLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return &BVMBITBurnIterator{contract: _BVMBIT.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_BVMBIT *BVMBITFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BVMBITBurn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BVMBIT.contract.WatchLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMBITBurn)
				if err := _BVMBIT.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_BVMBIT *BVMBITFilterer) ParseBurn(log types.Log) (*BVMBITBurn, error) {
	event := new(BVMBITBurn)
	if err := _BVMBIT.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMBITMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the BVMBIT contract.
type BVMBITMintIterator struct {
	Event *BVMBITMint // Event containing the contract specifics and raw log

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
func (it *BVMBITMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMBITMint)
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
		it.Event = new(BVMBITMint)
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
func (it *BVMBITMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMBITMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMBITMint represents a Mint event raised by the BVMBIT contract.
type BVMBITMint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_BVMBIT *BVMBITFilterer) FilterMint(opts *bind.FilterOpts, account []common.Address) (*BVMBITMintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BVMBIT.contract.FilterLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return &BVMBITMintIterator{contract: _BVMBIT.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_BVMBIT *BVMBITFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BVMBITMint, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BVMBIT.contract.WatchLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMBITMint)
				if err := _BVMBIT.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_BVMBIT *BVMBITFilterer) ParseMint(log types.Log) (*BVMBITMint, error) {
	event := new(BVMBITMint)
	if err := _BVMBIT.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BVMBITTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BVMBIT contract.
type BVMBITTransferIterator struct {
	Event *BVMBITTransfer // Event containing the contract specifics and raw log

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
func (it *BVMBITTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMBITTransfer)
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
		it.Event = new(BVMBITTransfer)
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
func (it *BVMBITTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMBITTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMBITTransfer represents a Transfer event raised by the BVMBIT contract.
type BVMBITTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BVMBIT *BVMBITFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BVMBITTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BVMBIT.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BVMBITTransferIterator{contract: _BVMBIT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BVMBIT *BVMBITFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BVMBITTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BVMBIT.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMBITTransfer)
				if err := _BVMBIT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BVMBIT *BVMBITFilterer) ParseTransfer(log types.Log) (*BVMBITTransfer, error) {
	event := new(BVMBITTransfer)
	if err := _BVMBIT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

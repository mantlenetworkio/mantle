// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

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

// SequencerSequencerInfo is an auto generated low-level Go binding around an user-defined struct.
type SequencerSequencerInfo struct {
	Owner       common.Address
	MintAddress common.Address
	NodeID      []byte
	Amount      *big.Int
	KeyIndex    *big.Int
}

// SequencerMetaData contains all meta data concerning the Sequencer contract.
var SequencerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"SequencerCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"SequencerDelete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SequencerUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bitToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_nodeID\",\"type\":\"bytes\"}],\"name\":\"createSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getSequencer\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"internalType\":\"structSequencer.SequencerInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencers\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"internalType\":\"structSequencer.SequencerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bitToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduler\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerLimit\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sequencers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bitToken\",\"type\":\"address\"}],\"name\":\"updateBitAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"updateEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"updateScheduler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_limit\",\"type\":\"uint8\"}],\"name\":\"updateSequencerLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612572806100206000396000f3fe608060405234801561001057600080fd5b506004361061018d5760003560e01c8063900cf0cf116100e3578063d1ad17bf1161008c578063e90f218f11610066578063e90f218f14610380578063ee43b5d9146103a0578063f2fde38b146103b357600080fd5b8063d1ad17bf14610345578063d84e9f921461035a578063e58c7bb51461036d57600080fd5b8063b6b55f25116100bd578063b6b55f25146102f6578063c4d66de814610309578063cab2ea2a1461031c57600080fd5b8063900cf0cf146102ab5780639c13b6b5146102c2578063a0e67e2b146102e157600080fd5b80632e1a7d4d11610145578063715018a61161011f578063715018a61461028a578063853828b6146102925780638da5cb5b1461029a57600080fd5b80632e1a7d4d1461022357806343dfc471146102365780636d46e9871461024957600080fd5b806315ca0dc01161017657806315ca0dc0146101d75780631c7a07ee146101ec57806326b0c1141461021057600080fd5b8063025e7c2714610192578063125c5f16146101c2575b600080fd5b6101a56101a0366004611e2e565b6103c6565b6040516001600160a01b0390911681526020015b60405180910390f35b6101ca6103f0565b6040516101b99190611f13565b6101ea6101e5366004611e2e565b6105d3565b005b6101ff6101fa366004611faf565b6105e0565b6040516101b9959493929190611fca565b6101ea61021e36600461200a565b6106a4565b6101ea610231366004611e2e565b6106e0565b6101ea61024436600461202d565b6108e9565b61027a610257366004611faf565b6001600160a01b0390811660009081526097602052604090206001015416151590565b60405190151581526020016101b9565b6101ea610cc4565b6101ea610cd8565b6065546001600160a01b03166101a5565b6102b4609b5481565b6040519081526020016101b9565b609c546102cf9060ff1681565b60405160ff90911681526020016101b9565b6102e9610e2c565b6040516101b991906120b4565b6101ea610304366004611e2e565b610e8e565b6101ea610317366004611faf565b611041565b6101a561032a366004611faf565b6098602052600090815260409020546001600160a01b031681565b61034d611222565b6040516101b99190612101565b609a546101a5906001600160a01b031681565b6101ea61037b366004612143565b6112b0565b61039361038e366004611faf565b6112cb565b6040516101b99190612212565b6101ea6103ae366004611faf565b6113fa565b6101ea6103c1366004611faf565b61143c565b609981815481106103d657600080fd5b6000918252602090912001546001600160a01b0316905081565b60995460609060009067ffffffffffffffff81111561041157610411612114565b60405190808252806020026020018201604052801561048657816020015b6104736040518060a0016040528060006001600160a01b0316815260200160006001600160a01b031681526020016060815260200160008152602001600081525090565b81526020019060019003908161042f5790505b50905060005b6099548110156105cd576000609982815481106104ab576104ab612225565b60009182526020808320909101546001600160a01b0390811680845260978352604093849020845160a08101865281548416815260018201549093169383019390935260028301805491955091938401919061050690612254565b80601f016020809104026020016040519081016040528092919081815260200182805461053290612254565b801561057f5780601f106105545761010080835404028352916020019161057f565b820191906000526020600020905b81548152906001019060200180831161056257829003601f168201915b50505050508152602001600382015481526020016004820154815250508383815181106105ae576105ae612225565b60200260200101819052505080806105c5906122d1565b91505061048c565b50919050565b6105db6114cc565b609b55565b6097602052600090815260409020805460018201546002830180546001600160a01b0393841694929093169261061590612254565b80601f016020809104026020016040519081016040528092919081815260200182805461064190612254565b801561068e5780601f106106635761010080835404028352916020019161068e565b820191906000526020600020905b81548152906001019060200180831161067157829003601f168201915b5050505050908060030154908060040154905085565b6106ac6114cc565b609c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055565b600260015414156107385760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b60026001558061078a5760405162461bcd60e51b815260206004820152600d60248201527f496e76696c6420616d6f756e7400000000000000000000000000000000000000604482015260640161072f565b336000908152609760205260409020600101546001600160a01b03166107f25760405162461bcd60e51b815260206004820152601360248201527f53657175656e636572206e6f7420657869737400000000000000000000000000604482015260640161072f565b3360009081526097602052604090206003015481908111156108235750336000908152609760205260409020600301545b609a5461083a906001600160a01b03163383611526565b336000908152609760205260408120600301805483929061085c90849061230a565b909155505033600090815260976020526040908190206001810154600382015492517fb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b16824936108b9936001600160a01b0390931692600201916123f8565b60405180910390a1336000908152609760205260409020600301546108e1576108e1336115f2565b505060018055565b6002600154141561093c5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161072f565b60026001558361098e5760405162461bcd60e51b815260206004820152600d60248201527f496e76696c6420616d6f756e7400000000000000000000000000000000000000604482015260640161072f565b6001600160a01b038316610a095760405162461bcd60e51b8152602060048201526024808201527f496e76696c6420616464726573732c20616464726573732063616e206e6f742060448201527f6265203000000000000000000000000000000000000000000000000000000000606482015260840161072f565b336000908152609760205260409020600101546001600160a01b031615610a725760405162461bcd60e51b815260206004820152601860248201527f416c726561647920686173206265656e20637265617465640000000000000000604482015260640161072f565b6001600160a01b038381166000908152609860205260409020541615610b005760405162461bcd60e51b815260206004820152602360248201527f54686973206d696e74206164647265737320616c726561647920686173206f7760448201527f6e65720000000000000000000000000000000000000000000000000000000000606482015260840161072f565b609a54610b18906001600160a01b0316333087611815565b6099546040805160a0810182523381526001600160a01b0386166020808301919091528251601f8601829004820281018201845285815291928301919086908690819084018382808284376000920182905250938552505050602080830189905260409283018590523382526097815290829020835181546001600160a01b039182167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617835585840151600184018054919093169116179055918301518051610beb9260028501920190611d5f565b506060820151600382015560809091015160049091015560998054600181019091557f72a152ddfb8e864297c917af52ea6c1c68aead0fee1a62673fcc7e0c94979d00018054337fffffffffffffffffffffffff000000000000000000000000000000000000000091821681179092556001600160a01b0386166000908152609860205260409081902080549092168317909155517f4bc782d7487c41941075eb09650b2eb45a57e23e2241db6a958b8af8485324f691610cb19187908790879061242a565b60405180910390a1505060018055505050565b610ccc6114cc565b610cd6600061186c565b565b60026001541415610d2b5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161072f565b6002600190815533600090815260976020526040902001546001600160a01b0316610d985760405162461bcd60e51b815260206004820152601260248201527f446f206e6f742068617665206372656174650000000000000000000000000000604482015260640161072f565b33600081815260976020526040902060030154609a549091610dc4916001600160a01b03169083611526565b33600090815260976020526040808220600181015491517fb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b1682493610e14936001600160a01b031692600201916123f8565b60405180910390a1610e25336115f2565b5060018055565b60606099805480602002602001604051908101604052809291908181526020018280548015610e8457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e66575b5050505050905090565b60026001541415610ee15760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161072f565b600260015580610f335760405162461bcd60e51b815260206004820152600d60248201527f496e76696c6420616d6f756e7400000000000000000000000000000000000000604482015260640161072f565b336000908152609760205260409020600101546001600160a01b0316610f9b5760405162461bcd60e51b815260206004820152601360248201527f53657175656e636572206e6f7420657869737400000000000000000000000000604482015260640161072f565b609a54610fb3906001600160a01b0316333084611815565b3360009081526097602052604081206003018054839290610fd5908490612495565b909155505033600090815260976020526040908190206001810154600382015492517fb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b1682493611032936001600160a01b0390931692600201916123f8565b60405180910390a15060018055565b600054610100900460ff16158080156110615750600054600160ff909116105b8061107b5750303b15801561107b575060005460ff166001145b6110ed5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161072f565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561114b57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6111536118d6565b61115b61195b565b609a80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384161790556000609b55609c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055801561121e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b609d805461122f90612254565b80601f016020809104026020016040519081016040528092919081815260200182805461125b90612254565b80156112a85780601f1061127d576101008083540402835291602001916112a8565b820191906000526020600020905b81548152906001019060200180831161128b57829003601f168201915b505050505081565b6112b86114cc565b805161121e90609d906020840190611d5f565b61130f6040518060a0016040528060006001600160a01b0316815260200160006001600160a01b031681526020016060815260200160008152602001600081525090565b6001600160a01b03808316600090815260976020908152604091829020825160a081018452815485168152600182015490941691840191909152600281018054919284019161135d90612254565b80601f016020809104026020016040519081016040528092919081815260200182805461138990612254565b80156113d65780601f106113ab576101008083540402835291602001916113d6565b820191906000526020600020905b8154815290600101906020018083116113b957829003601f168201915b50505050508152602001600382015481526020016004820154815250509050919050565b6114026114cc565b609a80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6114446114cc565b6001600160a01b0381166114c05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161072f565b6114c98161186c565b50565b6065546001600160a01b03163314610cd65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161072f565b6040516001600160a01b0383166024820152604481018290526115ed9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526119e0565b505050565b6001600160a01b0380821660009081526097602052604090819020600481015460995460018301549351919490937f37c49a6eae060065fbdeff05623918ea0969b4b1146b4c2fd33658caa72ed0db93611654939190921691600201906124ad565b60405180910390a1609961166960018361230a565b8154811061167957611679612225565b600091825260209091200154609980546001600160a01b0390921691849081106116a5576116a5612225565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160976000609985815481106116eb576116eb612225565b60009182526020808320909101546001600160a01b03908116845283820194909452604092830182206004019490945586831681526097845281812060010154909216825260989092522080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055609980548061176d5761176d6124cf565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000090811690915593019093556001600160a01b0386168152609790925260408220805482168155600181018054909216909155906118006002830182611de3565b50600060038201819055600490910155505050565b6040516001600160a01b03808516602483015283166044820152606481018290526118669085907f23b872dd000000000000000000000000000000000000000000000000000000009060840161156b565b50505050565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166119535760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161072f565b610cd6611ac5565b600054610100900460ff166119d85760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161072f565b610cd6611b4b565b6000611a35826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611bce9092919063ffffffff16565b8051909150156115ed5780806020019051810190611a5391906124fe565b6115ed5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161072f565b600054610100900460ff16611b425760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161072f565b610cd63361186c565b600054610100900460ff16611bc85760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161072f565b60018055565b6060611bdd8484600085611be7565b90505b9392505050565b606082471015611c5f5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161072f565b843b611cad5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161072f565b600080866001600160a01b03168587604051611cc99190612520565b60006040518083038185875af1925050503d8060008114611d06576040519150601f19603f3d011682016040523d82523d6000602084013e611d0b565b606091505b5091509150611d1b828286611d26565b979650505050505050565b60608315611d35575081611be0565b825115611d455782518084602001fd5b8160405162461bcd60e51b815260040161072f9190612101565b828054611d6b90612254565b90600052602060002090601f016020900481019282611d8d5760008555611dd3565b82601f10611da657805160ff1916838001178555611dd3565b82800160010185558215611dd3579182015b82811115611dd3578251825591602001919060010190611db8565b50611ddf929150611e19565b5090565b508054611def90612254565b6000825580601f10611dff575050565b601f0160209004906000526020600020908101906114c991905b5b80821115611ddf5760008155600101611e1a565b600060208284031215611e4057600080fd5b5035919050565b60005b83811015611e62578181015183820152602001611e4a565b838111156118665750506000910152565b60008151808452611e8b816020860160208601611e47565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60006001600160a01b0380835116845280602084015116602085015250604082015160a06040850152611ef360a0850182611e73565b905060608301516060850152608083015160808501528091505092915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015611f86577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0888603018452611f74858351611ebd565b94509285019290850190600101611f3a565b5092979650505050505050565b80356001600160a01b0381168114611faa57600080fd5b919050565b600060208284031215611fc157600080fd5b611be082611f93565b60006001600160a01b03808816835280871660208401525060a06040830152611ff660a0830186611e73565b606083019490945250608001529392505050565b60006020828403121561201c57600080fd5b813560ff81168114611be057600080fd5b6000806000806060858703121561204357600080fd5b8435935061205360208601611f93565b9250604085013567ffffffffffffffff8082111561207057600080fd5b818701915087601f83011261208457600080fd5b81358181111561209357600080fd5b8860208285010111156120a557600080fd5b95989497505060200194505050565b6020808252825182820181905260009190848201906040850190845b818110156120f55783516001600160a01b0316835292840192918401916001016120d0565b50909695505050505050565b602081526000611be06020830184611e73565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561215557600080fd5b813567ffffffffffffffff8082111561216d57600080fd5b818401915084601f83011261218157600080fd5b81358181111561219357612193612114565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156121d9576121d9612114565b816040528281528760208487010111156121f257600080fd5b826020860160208301376000928101602001929092525095945050505050565b602081526000611be06020830184611ebd565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c9082168061226857607f821691505b602082108114156105cd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415612303576123036122a2565b5060010190565b60008282101561231c5761231c6122a2565b500390565b8054600090600181811c908083168061233b57607f831692505b6020808410821415612376577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b8388526020880182801561239157600181146123c0576123eb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008716825282820197506123eb565b60008981526020902060005b878110156123e5578154848201529086019084016123cc565b83019850505b5050505050505092915050565b6001600160a01b038416815260606020820152600061241a6060830185612321565b9050826040830152949350505050565b60006001600160a01b038087168352808616602084015250606060408301528260608301528284608084013760006080848401015260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f850116830101905095945050505050565b600082198211156124a8576124a86122a2565b500190565b6001600160a01b0383168152604060208201526000611bdd6040830184612321565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60006020828403121561251057600080fd5b81518015158114611be057600080fd5b60008251612532818460208701611e47565b919091019291505056fea26469706673582212202e06db721591222522407b4708c52f86e09d437c519c9bf3deac6053896d86b864736f6c63430008090033",
}

// SequencerABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerMetaData.ABI instead.
var SequencerABI = SequencerMetaData.ABI

// SequencerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SequencerMetaData.Bin instead.
var SequencerBin = SequencerMetaData.Bin

// DeploySequencer deploys a new Ethereum contract, binding an instance of Sequencer to it.
func DeploySequencer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sequencer, error) {
	parsed, err := SequencerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SequencerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sequencer{SequencerCaller: SequencerCaller{contract: contract}, SequencerTransactor: SequencerTransactor{contract: contract}, SequencerFilterer: SequencerFilterer{contract: contract}}, nil
}

// Sequencer is an auto generated Go binding around an Ethereum contract.
type Sequencer struct {
	SequencerCaller     // Read-only binding to the contract
	SequencerTransactor // Write-only binding to the contract
	SequencerFilterer   // Log filterer for contract events
}

// SequencerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerSession struct {
	Contract     *Sequencer        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SequencerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerCallerSession struct {
	Contract *SequencerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SequencerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerTransactorSession struct {
	Contract     *SequencerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SequencerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerRaw struct {
	Contract *Sequencer // Generic contract binding to access the raw methods on
}

// SequencerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerCallerRaw struct {
	Contract *SequencerCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerTransactorRaw struct {
	Contract *SequencerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencer creates a new instance of Sequencer, bound to a specific deployed contract.
func NewSequencer(address common.Address, backend bind.ContractBackend) (*Sequencer, error) {
	contract, err := bindSequencer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sequencer{SequencerCaller: SequencerCaller{contract: contract}, SequencerTransactor: SequencerTransactor{contract: contract}, SequencerFilterer: SequencerFilterer{contract: contract}}, nil
}

// NewSequencerCaller creates a new read-only instance of Sequencer, bound to a specific deployed contract.
func NewSequencerCaller(address common.Address, caller bind.ContractCaller) (*SequencerCaller, error) {
	contract, err := bindSequencer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerCaller{contract: contract}, nil
}

// NewSequencerTransactor creates a new write-only instance of Sequencer, bound to a specific deployed contract.
func NewSequencerTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerTransactor, error) {
	contract, err := bindSequencer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerTransactor{contract: contract}, nil
}

// NewSequencerFilterer creates a new log filterer instance of Sequencer, bound to a specific deployed contract.
func NewSequencerFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerFilterer, error) {
	contract, err := bindSequencer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerFilterer{contract: contract}, nil
}

// bindSequencer binds a generic wrapper to an already deployed contract.
func bindSequencer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SequencerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sequencer *SequencerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sequencer.Contract.SequencerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sequencer *SequencerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.Contract.SequencerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sequencer *SequencerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sequencer.Contract.SequencerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sequencer *SequencerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sequencer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sequencer *SequencerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sequencer *SequencerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sequencer.Contract.contract.Transact(opts, method, params...)
}

// BitToken is a free data retrieval call binding the contract method 0xd84e9f92.
//
// Solidity: function bitToken() view returns(address)
func (_Sequencer *SequencerCaller) BitToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "bitToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BitToken is a free data retrieval call binding the contract method 0xd84e9f92.
//
// Solidity: function bitToken() view returns(address)
func (_Sequencer *SequencerSession) BitToken() (common.Address, error) {
	return _Sequencer.Contract.BitToken(&_Sequencer.CallOpts)
}

// BitToken is a free data retrieval call binding the contract method 0xd84e9f92.
//
// Solidity: function bitToken() view returns(address)
func (_Sequencer *SequencerCallerSession) BitToken() (common.Address, error) {
	return _Sequencer.Contract.BitToken(&_Sequencer.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Sequencer *SequencerCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Sequencer *SequencerSession) Epoch() (*big.Int, error) {
	return _Sequencer.Contract.Epoch(&_Sequencer.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Sequencer *SequencerCallerSession) Epoch() (*big.Int, error) {
	return _Sequencer.Contract.Epoch(&_Sequencer.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Sequencer *SequencerCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Sequencer *SequencerSession) GetOwners() ([]common.Address, error) {
	return _Sequencer.Contract.GetOwners(&_Sequencer.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Sequencer *SequencerCallerSession) GetOwners() ([]common.Address, error) {
	return _Sequencer.Contract.GetOwners(&_Sequencer.CallOpts)
}

// GetSequencer is a free data retrieval call binding the contract method 0xe90f218f.
//
// Solidity: function getSequencer(address signer) view returns((address,address,bytes,uint256,uint256))
func (_Sequencer *SequencerCaller) GetSequencer(opts *bind.CallOpts, signer common.Address) (SequencerSequencerInfo, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getSequencer", signer)

	if err != nil {
		return *new(SequencerSequencerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SequencerSequencerInfo)).(*SequencerSequencerInfo)

	return out0, err

}

// GetSequencer is a free data retrieval call binding the contract method 0xe90f218f.
//
// Solidity: function getSequencer(address signer) view returns((address,address,bytes,uint256,uint256))
func (_Sequencer *SequencerSession) GetSequencer(signer common.Address) (SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencer(&_Sequencer.CallOpts, signer)
}

// GetSequencer is a free data retrieval call binding the contract method 0xe90f218f.
//
// Solidity: function getSequencer(address signer) view returns((address,address,bytes,uint256,uint256))
func (_Sequencer *SequencerCallerSession) GetSequencer(signer common.Address) (SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencer(&_Sequencer.CallOpts, signer)
}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns((address,address,bytes,uint256,uint256)[])
func (_Sequencer *SequencerCaller) GetSequencers(opts *bind.CallOpts) ([]SequencerSequencerInfo, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getSequencers")

	if err != nil {
		return *new([]SequencerSequencerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]SequencerSequencerInfo)).(*[]SequencerSequencerInfo)

	return out0, err

}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns((address,address,bytes,uint256,uint256)[])
func (_Sequencer *SequencerSession) GetSequencers() ([]SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencers(&_Sequencer.CallOpts)
}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns((address,address,bytes,uint256,uint256)[])
func (_Sequencer *SequencerCallerSession) GetSequencers() ([]SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencers(&_Sequencer.CallOpts)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address signer) view returns(bool)
func (_Sequencer *SequencerCaller) IsSequencer(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "isSequencer", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address signer) view returns(bool)
func (_Sequencer *SequencerSession) IsSequencer(signer common.Address) (bool, error) {
	return _Sequencer.Contract.IsSequencer(&_Sequencer.CallOpts, signer)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address signer) view returns(bool)
func (_Sequencer *SequencerCallerSession) IsSequencer(signer common.Address) (bool, error) {
	return _Sequencer.Contract.IsSequencer(&_Sequencer.CallOpts, signer)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sequencer *SequencerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sequencer *SequencerSession) Owner() (common.Address, error) {
	return _Sequencer.Contract.Owner(&_Sequencer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sequencer *SequencerCallerSession) Owner() (common.Address, error) {
	return _Sequencer.Contract.Owner(&_Sequencer.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Sequencer *SequencerCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Sequencer *SequencerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.Owners(&_Sequencer.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Sequencer *SequencerCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.Owners(&_Sequencer.CallOpts, arg0)
}

// Rel is a free data retrieval call binding the contract method 0xcab2ea2a.
//
// Solidity: function rel(address ) view returns(address)
func (_Sequencer *SequencerCaller) Rel(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "rel", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xcab2ea2a.
//
// Solidity: function rel(address ) view returns(address)
func (_Sequencer *SequencerSession) Rel(arg0 common.Address) (common.Address, error) {
	return _Sequencer.Contract.Rel(&_Sequencer.CallOpts, arg0)
}

// Rel is a free data retrieval call binding the contract method 0xcab2ea2a.
//
// Solidity: function rel(address ) view returns(address)
func (_Sequencer *SequencerCallerSession) Rel(arg0 common.Address) (common.Address, error) {
	return _Sequencer.Contract.Rel(&_Sequencer.CallOpts, arg0)
}

// Scheduler is a free data retrieval call binding the contract method 0xd1ad17bf.
//
// Solidity: function scheduler() view returns(bytes)
func (_Sequencer *SequencerCaller) Scheduler(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "scheduler")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Scheduler is a free data retrieval call binding the contract method 0xd1ad17bf.
//
// Solidity: function scheduler() view returns(bytes)
func (_Sequencer *SequencerSession) Scheduler() ([]byte, error) {
	return _Sequencer.Contract.Scheduler(&_Sequencer.CallOpts)
}

// Scheduler is a free data retrieval call binding the contract method 0xd1ad17bf.
//
// Solidity: function scheduler() view returns(bytes)
func (_Sequencer *SequencerCallerSession) Scheduler() ([]byte, error) {
	return _Sequencer.Contract.Scheduler(&_Sequencer.CallOpts)
}

// SequencerLimit is a free data retrieval call binding the contract method 0x9c13b6b5.
//
// Solidity: function sequencerLimit() view returns(uint8)
func (_Sequencer *SequencerCaller) SequencerLimit(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "sequencerLimit")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SequencerLimit is a free data retrieval call binding the contract method 0x9c13b6b5.
//
// Solidity: function sequencerLimit() view returns(uint8)
func (_Sequencer *SequencerSession) SequencerLimit() (uint8, error) {
	return _Sequencer.Contract.SequencerLimit(&_Sequencer.CallOpts)
}

// SequencerLimit is a free data retrieval call binding the contract method 0x9c13b6b5.
//
// Solidity: function sequencerLimit() view returns(uint8)
func (_Sequencer *SequencerCallerSession) SequencerLimit() (uint8, error) {
	return _Sequencer.Contract.SequencerLimit(&_Sequencer.CallOpts)
}

// Sequencers is a free data retrieval call binding the contract method 0x1c7a07ee.
//
// Solidity: function sequencers(address ) view returns(address owner, address mintAddress, bytes nodeID, uint256 amount, uint256 keyIndex)
func (_Sequencer *SequencerCaller) Sequencers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner       common.Address
	MintAddress common.Address
	NodeID      []byte
	Amount      *big.Int
	KeyIndex    *big.Int
}, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "sequencers", arg0)

	outstruct := new(struct {
		Owner       common.Address
		MintAddress common.Address
		NodeID      []byte
		Amount      *big.Int
		KeyIndex    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MintAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NodeID = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.KeyIndex = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Sequencers is a free data retrieval call binding the contract method 0x1c7a07ee.
//
// Solidity: function sequencers(address ) view returns(address owner, address mintAddress, bytes nodeID, uint256 amount, uint256 keyIndex)
func (_Sequencer *SequencerSession) Sequencers(arg0 common.Address) (struct {
	Owner       common.Address
	MintAddress common.Address
	NodeID      []byte
	Amount      *big.Int
	KeyIndex    *big.Int
}, error) {
	return _Sequencer.Contract.Sequencers(&_Sequencer.CallOpts, arg0)
}

// Sequencers is a free data retrieval call binding the contract method 0x1c7a07ee.
//
// Solidity: function sequencers(address ) view returns(address owner, address mintAddress, bytes nodeID, uint256 amount, uint256 keyIndex)
func (_Sequencer *SequencerCallerSession) Sequencers(arg0 common.Address) (struct {
	Owner       common.Address
	MintAddress common.Address
	NodeID      []byte
	Amount      *big.Int
	KeyIndex    *big.Int
}, error) {
	return _Sequencer.Contract.Sequencers(&_Sequencer.CallOpts, arg0)
}

// CreateSequencer is a paid mutator transaction binding the contract method 0x43dfc471.
//
// Solidity: function createSequencer(uint256 _amount, address _mintAddress, bytes _nodeID) returns()
func (_Sequencer *SequencerTransactor) CreateSequencer(opts *bind.TransactOpts, _amount *big.Int, _mintAddress common.Address, _nodeID []byte) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "createSequencer", _amount, _mintAddress, _nodeID)
}

// CreateSequencer is a paid mutator transaction binding the contract method 0x43dfc471.
//
// Solidity: function createSequencer(uint256 _amount, address _mintAddress, bytes _nodeID) returns()
func (_Sequencer *SequencerSession) CreateSequencer(_amount *big.Int, _mintAddress common.Address, _nodeID []byte) (*types.Transaction, error) {
	return _Sequencer.Contract.CreateSequencer(&_Sequencer.TransactOpts, _amount, _mintAddress, _nodeID)
}

// CreateSequencer is a paid mutator transaction binding the contract method 0x43dfc471.
//
// Solidity: function createSequencer(uint256 _amount, address _mintAddress, bytes _nodeID) returns()
func (_Sequencer *SequencerTransactorSession) CreateSequencer(_amount *big.Int, _mintAddress common.Address, _nodeID []byte) (*types.Transaction, error) {
	return _Sequencer.Contract.CreateSequencer(&_Sequencer.TransactOpts, _amount, _mintAddress, _nodeID)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Sequencer *SequencerTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Sequencer *SequencerSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.Deposit(&_Sequencer.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Sequencer *SequencerTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.Deposit(&_Sequencer.TransactOpts, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bitToken) returns()
func (_Sequencer *SequencerTransactor) Initialize(opts *bind.TransactOpts, _bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "initialize", _bitToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bitToken) returns()
func (_Sequencer *SequencerSession) Initialize(_bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.Initialize(&_Sequencer.TransactOpts, _bitToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bitToken) returns()
func (_Sequencer *SequencerTransactorSession) Initialize(_bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.Initialize(&_Sequencer.TransactOpts, _bitToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sequencer *SequencerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sequencer *SequencerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sequencer.Contract.RenounceOwnership(&_Sequencer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sequencer *SequencerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sequencer.Contract.RenounceOwnership(&_Sequencer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sequencer *SequencerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sequencer *SequencerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.TransferOwnership(&_Sequencer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sequencer *SequencerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.TransferOwnership(&_Sequencer.TransactOpts, newOwner)
}

// UpdateBitAddress is a paid mutator transaction binding the contract method 0xee43b5d9.
//
// Solidity: function updateBitAddress(address _bitToken) returns()
func (_Sequencer *SequencerTransactor) UpdateBitAddress(opts *bind.TransactOpts, _bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "updateBitAddress", _bitToken)
}

// UpdateBitAddress is a paid mutator transaction binding the contract method 0xee43b5d9.
//
// Solidity: function updateBitAddress(address _bitToken) returns()
func (_Sequencer *SequencerSession) UpdateBitAddress(_bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateBitAddress(&_Sequencer.TransactOpts, _bitToken)
}

// UpdateBitAddress is a paid mutator transaction binding the contract method 0xee43b5d9.
//
// Solidity: function updateBitAddress(address _bitToken) returns()
func (_Sequencer *SequencerTransactorSession) UpdateBitAddress(_bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateBitAddress(&_Sequencer.TransactOpts, _bitToken)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x15ca0dc0.
//
// Solidity: function updateEpoch(uint256 _epoch) returns()
func (_Sequencer *SequencerTransactor) UpdateEpoch(opts *bind.TransactOpts, _epoch *big.Int) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "updateEpoch", _epoch)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x15ca0dc0.
//
// Solidity: function updateEpoch(uint256 _epoch) returns()
func (_Sequencer *SequencerSession) UpdateEpoch(_epoch *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateEpoch(&_Sequencer.TransactOpts, _epoch)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x15ca0dc0.
//
// Solidity: function updateEpoch(uint256 _epoch) returns()
func (_Sequencer *SequencerTransactorSession) UpdateEpoch(_epoch *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateEpoch(&_Sequencer.TransactOpts, _epoch)
}

// UpdateScheduler is a paid mutator transaction binding the contract method 0xe58c7bb5.
//
// Solidity: function updateScheduler(bytes nodeID) returns()
func (_Sequencer *SequencerTransactor) UpdateScheduler(opts *bind.TransactOpts, nodeID []byte) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "updateScheduler", nodeID)
}

// UpdateScheduler is a paid mutator transaction binding the contract method 0xe58c7bb5.
//
// Solidity: function updateScheduler(bytes nodeID) returns()
func (_Sequencer *SequencerSession) UpdateScheduler(nodeID []byte) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateScheduler(&_Sequencer.TransactOpts, nodeID)
}

// UpdateScheduler is a paid mutator transaction binding the contract method 0xe58c7bb5.
//
// Solidity: function updateScheduler(bytes nodeID) returns()
func (_Sequencer *SequencerTransactorSession) UpdateScheduler(nodeID []byte) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateScheduler(&_Sequencer.TransactOpts, nodeID)
}

// UpdateSequencerLimit is a paid mutator transaction binding the contract method 0x26b0c114.
//
// Solidity: function updateSequencerLimit(uint8 _limit) returns()
func (_Sequencer *SequencerTransactor) UpdateSequencerLimit(opts *bind.TransactOpts, _limit uint8) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "updateSequencerLimit", _limit)
}

// UpdateSequencerLimit is a paid mutator transaction binding the contract method 0x26b0c114.
//
// Solidity: function updateSequencerLimit(uint8 _limit) returns()
func (_Sequencer *SequencerSession) UpdateSequencerLimit(_limit uint8) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateSequencerLimit(&_Sequencer.TransactOpts, _limit)
}

// UpdateSequencerLimit is a paid mutator transaction binding the contract method 0x26b0c114.
//
// Solidity: function updateSequencerLimit(uint8 _limit) returns()
func (_Sequencer *SequencerTransactorSession) UpdateSequencerLimit(_limit uint8) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateSequencerLimit(&_Sequencer.TransactOpts, _limit)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Sequencer *SequencerTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Sequencer *SequencerSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.Withdraw(&_Sequencer.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Sequencer *SequencerTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.Withdraw(&_Sequencer.TransactOpts, _amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Sequencer *SequencerTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Sequencer *SequencerSession) WithdrawAll() (*types.Transaction, error) {
	return _Sequencer.Contract.WithdrawAll(&_Sequencer.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Sequencer *SequencerTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Sequencer.Contract.WithdrawAll(&_Sequencer.TransactOpts)
}

// SequencerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Sequencer contract.
type SequencerInitializedIterator struct {
	Event *SequencerInitialized // Event containing the contract specifics and raw log

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
func (it *SequencerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInitialized)
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
		it.Event = new(SequencerInitialized)
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
func (it *SequencerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInitialized represents a Initialized event raised by the Sequencer contract.
type SequencerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Sequencer *SequencerFilterer) FilterInitialized(opts *bind.FilterOpts) (*SequencerInitializedIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SequencerInitializedIterator{contract: _Sequencer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Sequencer *SequencerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SequencerInitialized) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInitialized)
				if err := _Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Sequencer *SequencerFilterer) ParseInitialized(log types.Log) (*SequencerInitialized, error) {
	event := new(SequencerInitialized)
	if err := _Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Sequencer contract.
type SequencerOwnershipTransferredIterator struct {
	Event *SequencerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SequencerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerOwnershipTransferred)
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
		it.Event = new(SequencerOwnershipTransferred)
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
func (it *SequencerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerOwnershipTransferred represents a OwnershipTransferred event raised by the Sequencer contract.
type SequencerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sequencer *SequencerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SequencerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SequencerOwnershipTransferredIterator{contract: _Sequencer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sequencer *SequencerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SequencerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerOwnershipTransferred)
				if err := _Sequencer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Sequencer *SequencerFilterer) ParseOwnershipTransferred(log types.Log) (*SequencerOwnershipTransferred, error) {
	event := new(SequencerOwnershipTransferred)
	if err := _Sequencer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerSequencerCreateIterator is returned from FilterSequencerCreate and is used to iterate over the raw logs and unpacked data for SequencerCreate events raised by the Sequencer contract.
type SequencerSequencerCreateIterator struct {
	Event *SequencerSequencerCreate // Event containing the contract specifics and raw log

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
func (it *SequencerSequencerCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerSequencerCreate)
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
		it.Event = new(SequencerSequencerCreate)
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
func (it *SequencerSequencerCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerSequencerCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerSequencerCreate represents a SequencerCreate event raised by the Sequencer contract.
type SequencerSequencerCreate struct {
	Arg0 common.Address
	Arg1 common.Address
	Arg2 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSequencerCreate is a free log retrieval operation binding the contract event 0x4bc782d7487c41941075eb09650b2eb45a57e23e2241db6a958b8af8485324f6.
//
// Solidity: event SequencerCreate(address arg0, address arg1, bytes arg2)
func (_Sequencer *SequencerFilterer) FilterSequencerCreate(opts *bind.FilterOpts) (*SequencerSequencerCreateIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "SequencerCreate")
	if err != nil {
		return nil, err
	}
	return &SequencerSequencerCreateIterator{contract: _Sequencer.contract, event: "SequencerCreate", logs: logs, sub: sub}, nil
}

// WatchSequencerCreate is a free log subscription operation binding the contract event 0x4bc782d7487c41941075eb09650b2eb45a57e23e2241db6a958b8af8485324f6.
//
// Solidity: event SequencerCreate(address arg0, address arg1, bytes arg2)
func (_Sequencer *SequencerFilterer) WatchSequencerCreate(opts *bind.WatchOpts, sink chan<- *SequencerSequencerCreate) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "SequencerCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerSequencerCreate)
				if err := _Sequencer.contract.UnpackLog(event, "SequencerCreate", log); err != nil {
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

// ParseSequencerCreate is a log parse operation binding the contract event 0x4bc782d7487c41941075eb09650b2eb45a57e23e2241db6a958b8af8485324f6.
//
// Solidity: event SequencerCreate(address arg0, address arg1, bytes arg2)
func (_Sequencer *SequencerFilterer) ParseSequencerCreate(log types.Log) (*SequencerSequencerCreate, error) {
	event := new(SequencerSequencerCreate)
	if err := _Sequencer.contract.UnpackLog(event, "SequencerCreate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerSequencerDeleteIterator is returned from FilterSequencerDelete and is used to iterate over the raw logs and unpacked data for SequencerDelete events raised by the Sequencer contract.
type SequencerSequencerDeleteIterator struct {
	Event *SequencerSequencerDelete // Event containing the contract specifics and raw log

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
func (it *SequencerSequencerDeleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerSequencerDelete)
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
		it.Event = new(SequencerSequencerDelete)
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
func (it *SequencerSequencerDeleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerSequencerDeleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerSequencerDelete represents a SequencerDelete event raised by the Sequencer contract.
type SequencerSequencerDelete struct {
	Arg0 common.Address
	Arg1 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSequencerDelete is a free log retrieval operation binding the contract event 0x37c49a6eae060065fbdeff05623918ea0969b4b1146b4c2fd33658caa72ed0db.
//
// Solidity: event SequencerDelete(address arg0, bytes arg1)
func (_Sequencer *SequencerFilterer) FilterSequencerDelete(opts *bind.FilterOpts) (*SequencerSequencerDeleteIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "SequencerDelete")
	if err != nil {
		return nil, err
	}
	return &SequencerSequencerDeleteIterator{contract: _Sequencer.contract, event: "SequencerDelete", logs: logs, sub: sub}, nil
}

// WatchSequencerDelete is a free log subscription operation binding the contract event 0x37c49a6eae060065fbdeff05623918ea0969b4b1146b4c2fd33658caa72ed0db.
//
// Solidity: event SequencerDelete(address arg0, bytes arg1)
func (_Sequencer *SequencerFilterer) WatchSequencerDelete(opts *bind.WatchOpts, sink chan<- *SequencerSequencerDelete) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "SequencerDelete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerSequencerDelete)
				if err := _Sequencer.contract.UnpackLog(event, "SequencerDelete", log); err != nil {
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

// ParseSequencerDelete is a log parse operation binding the contract event 0x37c49a6eae060065fbdeff05623918ea0969b4b1146b4c2fd33658caa72ed0db.
//
// Solidity: event SequencerDelete(address arg0, bytes arg1)
func (_Sequencer *SequencerFilterer) ParseSequencerDelete(log types.Log) (*SequencerSequencerDelete, error) {
	event := new(SequencerSequencerDelete)
	if err := _Sequencer.contract.UnpackLog(event, "SequencerDelete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerSequencerUpdateIterator is returned from FilterSequencerUpdate and is used to iterate over the raw logs and unpacked data for SequencerUpdate events raised by the Sequencer contract.
type SequencerSequencerUpdateIterator struct {
	Event *SequencerSequencerUpdate // Event containing the contract specifics and raw log

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
func (it *SequencerSequencerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerSequencerUpdate)
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
		it.Event = new(SequencerSequencerUpdate)
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
func (it *SequencerSequencerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerSequencerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerSequencerUpdate represents a SequencerUpdate event raised by the Sequencer contract.
type SequencerSequencerUpdate struct {
	Arg0 common.Address
	Arg1 []byte
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSequencerUpdate is a free log retrieval operation binding the contract event 0xb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b16824.
//
// Solidity: event SequencerUpdate(address arg0, bytes arg1, uint256 arg2)
func (_Sequencer *SequencerFilterer) FilterSequencerUpdate(opts *bind.FilterOpts) (*SequencerSequencerUpdateIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "SequencerUpdate")
	if err != nil {
		return nil, err
	}
	return &SequencerSequencerUpdateIterator{contract: _Sequencer.contract, event: "SequencerUpdate", logs: logs, sub: sub}, nil
}

// WatchSequencerUpdate is a free log subscription operation binding the contract event 0xb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b16824.
//
// Solidity: event SequencerUpdate(address arg0, bytes arg1, uint256 arg2)
func (_Sequencer *SequencerFilterer) WatchSequencerUpdate(opts *bind.WatchOpts, sink chan<- *SequencerSequencerUpdate) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "SequencerUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerSequencerUpdate)
				if err := _Sequencer.contract.UnpackLog(event, "SequencerUpdate", log); err != nil {
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

// ParseSequencerUpdate is a log parse operation binding the contract event 0xb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b16824.
//
// Solidity: event SequencerUpdate(address arg0, bytes arg1, uint256 arg2)
func (_Sequencer *SequencerFilterer) ParseSequencerUpdate(log types.Log) (*SequencerSequencerUpdate, error) {
	event := new(SequencerSequencerUpdate)
	if err := _Sequencer.contract.UnpackLog(event, "SequencerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

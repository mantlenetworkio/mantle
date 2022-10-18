// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package seq

import (
	"math/big"
	"strings"

	ethereum "github.com/bitdao-io/bitnetwork/l2geth"
	"github.com/bitdao-io/bitnetwork/l2geth/accounts/abi"
	"github.com/bitdao-io/bitnetwork/l2geth/accounts/abi/bind"
	"github.com/bitdao-io/bitnetwork/l2geth/common"
	"github.com/bitdao-io/bitnetwork/l2geth/core/types"
	"github.com/bitdao-io/bitnetwork/l2geth/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SequencerSequencerInfo is an auto generated low-level Go binding around an user-defined struct.
type SequencerSequencerInfo struct {
	Owner       common.Address
	MintAddress common.Address
	NodeID      []byte
	Amount      *big.Int
	KeyIndex    *big.Int
}

// SequencerABI is the input ABI used to generate the binding from.
const SequencerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"SequencerCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"SequencerDelete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SequencerUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bitToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bitToken\",\"type\":\"address\"}],\"name\":\"changeBitAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_nodeID\",\"type\":\"bytes\"}],\"name\":\"createSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getSequencer\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"internalType\":\"structSequencer.SequencerInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencers\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"internalType\":\"structSequencer.SequencerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bitToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sequencers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SequencerBin is the compiled bytecode used for deploying new contracts.
var SequencerBin = "0x608060405234801561001057600080fd5b50612213806100206000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c8063853828b6116100b2578063c4d66de811610081578063d84e9f9211610066578063d84e9f921461028a578063e90f218f1461029d578063f2fde38b146102bd57600080fd5b8063c4d66de81461024e578063cab2ea2a1461026157600080fd5b8063853828b61461020d5780638da5cb5b14610215578063a0e67e2b14610226578063b6b55f251461023b57600080fd5b806343dfc471116100ee57806343dfc4711461019e5780636d46e987146101b1578063715018a6146101f25780637a0674a0146101fa57600080fd5b8063025e7c2714610120578063125c5f16146101505780631c7a07ee146101655780632e1a7d4d14610189575b600080fd5b61013361012e366004611bc1565b6102d0565b6040516001600160a01b0390911681526020015b60405180910390f35b6101586102fa565b6040516101479190611ca6565b610178610173366004611d42565b6104dd565b604051610147959493929190611d5d565b61019c610197366004611bc1565b6105a1565b005b61019c6101ac366004611d9d565b6107aa565b6101e26101bf366004611d42565b6001600160a01b0390811660009081526097602052604090206001015416151590565b6040519015158152602001610147565b61019c610b85565b61019c610208366004611d42565b610b99565b61019c610bdb565b6065546001600160a01b0316610133565b61022e610cd7565b6040516101479190611e24565b61019c610249366004611bc1565b610d39565b61019c61025c366004611d42565b610eec565b61013361026f366004611d42565b6098602052600090815260409020546001600160a01b031681565b609a54610133906001600160a01b031681565b6102b06102ab366004611d42565b6110a0565b6040516101479190611e71565b61019c6102cb366004611d42565b6111cf565b609981815481106102e057600080fd5b6000918252602090912001546001600160a01b0316905081565b60995460609060009067ffffffffffffffff81111561031b5761031b611e84565b60405190808252806020026020018201604052801561039057816020015b61037d6040518060a0016040528060006001600160a01b0316815260200160006001600160a01b031681526020016060815260200160008152602001600081525090565b8152602001906001900390816103395790505b50905060005b6099548110156104d7576000609982815481106103b5576103b5611eb3565b60009182526020808320909101546001600160a01b0390811680845260978352604093849020845160a08101865281548416815260018201549093169383019390935260028301805491955091938401919061041090611ee2565b80601f016020809104026020016040519081016040528092919081815260200182805461043c90611ee2565b80156104895780601f1061045e57610100808354040283529160200191610489565b820191906000526020600020905b81548152906001019060200180831161046c57829003601f168201915b50505050508152602001600382015481526020016004820154815250508383815181106104b8576104b8611eb3565b60200260200101819052505080806104cf90611f5f565b915050610396565b50919050565b6097602052600090815260409020805460018201546002830180546001600160a01b0393841694929093169261051290611ee2565b80601f016020809104026020016040519081016040528092919081815260200182805461053e90611ee2565b801561058b5780601f106105605761010080835404028352916020019161058b565b820191906000526020600020905b81548152906001019060200180831161056e57829003601f168201915b5050505050908060030154908060040154905085565b600260015414156105f95760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b60026001558061064b5760405162461bcd60e51b815260206004820152600d60248201527f496e76696c6420616d6f756e740000000000000000000000000000000000000060448201526064016105f0565b336000908152609760205260409020600101546001600160a01b03166106b35760405162461bcd60e51b815260206004820152601360248201527f53657175656e636572206e6f742065786973740000000000000000000000000060448201526064016105f0565b3360009081526097602052604090206003015481908111156106e45750336000908152609760205260409020600301545b609a546106fb906001600160a01b0316338361125f565b336000908152609760205260408120600301805483929061071d908490611f98565b909155505033600090815260976020526040908190206001810154600382015492517fb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b168249361077a936001600160a01b039093169260020191612086565b60405180910390a1336000908152609760205260409020600301546107a2576107a23361132b565b505060018055565b600260015414156107fd5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016105f0565b60026001558361084f5760405162461bcd60e51b815260206004820152600d60248201527f496e76696c6420616d6f756e740000000000000000000000000000000000000060448201526064016105f0565b6001600160a01b0383166108ca5760405162461bcd60e51b8152602060048201526024808201527f496e76696c6420616464726573732c20616464726573732063616e206e6f742060448201527f626520300000000000000000000000000000000000000000000000000000000060648201526084016105f0565b336000908152609760205260409020600101546001600160a01b0316156109335760405162461bcd60e51b815260206004820152601460248201527f416c72656164792068617665206465706f73697400000000000000000000000060448201526064016105f0565b6001600160a01b0383811660009081526098602052604090205416156109c15760405162461bcd60e51b815260206004820152602560248201527f54686973206d696e74206164647265737320616c72656164792068617665207360448201527f656e64657200000000000000000000000000000000000000000000000000000060648201526084016105f0565b609a546109d9906001600160a01b031633308761154e565b6099546040805160a0810182523381526001600160a01b0386166020808301919091528251601f8601829004820281018201845285815291928301919086908690819084018382808284376000920182905250938552505050602080830189905260409283018590523382526097815290829020835181546001600160a01b039182167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617835585840151600184018054919093169116179055918301518051610aac9260028501920190611af2565b506060820151600382015560809091015160049091015560998054600181019091557f72a152ddfb8e864297c917af52ea6c1c68aead0fee1a62673fcc7e0c94979d00018054337fffffffffffffffffffffffff000000000000000000000000000000000000000091821681179092556001600160a01b0386166000908152609860205260409081902080549092168317909155517f4bc782d7487c41941075eb09650b2eb45a57e23e2241db6a958b8af8485324f691610b72918790879087906120b8565b60405180910390a1505060018055505050565b610b8d6115a5565b610b9760006115ff565b565b610ba16115a5565b609a80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60026001541415610c2e5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016105f0565b6002600190815533600090815260976020526040902001546001600160a01b0316610c9b5760405162461bcd60e51b815260206004820152601260248201527f446f206e6f74206861766520637265617465000000000000000000000000000060448201526064016105f0565b33600081815260976020526040902060030154609a549091610cc7916001600160a01b0316908361125f565b610cd03361132b565b5060018055565b60606099805480602002602001604051908101604052809291908181526020018280548015610d2f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d11575b5050505050905090565b60026001541415610d8c5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016105f0565b600260015580610dde5760405162461bcd60e51b815260206004820152600d60248201527f496e76696c6420616d6f756e740000000000000000000000000000000000000060448201526064016105f0565b336000908152609760205260409020600101546001600160a01b0316610e465760405162461bcd60e51b815260206004820152601360248201527f53657175656e636572206e6f742065786973740000000000000000000000000060448201526064016105f0565b609a54610e5e906001600160a01b031633308461154e565b3360009081526097602052604081206003018054839290610e80908490612123565b909155505033600090815260976020526040908190206001810154600382015492517fb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b1682493610edd936001600160a01b039093169260020191612086565b60405180910390a15060018055565b600054610100900460ff1615808015610f0c5750600054600160ff909116105b80610f265750303b158015610f26575060005460ff166001145b610f985760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105f0565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610ff657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610ffe611669565b6110066116ee565b609a80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038416179055801561109c57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6110e46040518060a0016040528060006001600160a01b0316815260200160006001600160a01b031681526020016060815260200160008152602001600081525090565b6001600160a01b03808316600090815260976020908152604091829020825160a081018452815485168152600182015490941691840191909152600281018054919284019161113290611ee2565b80601f016020809104026020016040519081016040528092919081815260200182805461115e90611ee2565b80156111ab5780601f10611180576101008083540402835291602001916111ab565b820191906000526020600020905b81548152906001019060200180831161118e57829003601f168201915b50505050508152602001600382015481526020016004820154815250509050919050565b6111d76115a5565b6001600160a01b0381166112535760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016105f0565b61125c816115ff565b50565b6040516001600160a01b0383166024820152604481018290526113269084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611773565b505050565b6001600160a01b0380821660009081526097602052604090819020600481015460995460018301549351919490937f37c49a6eae060065fbdeff05623918ea0969b4b1146b4c2fd33658caa72ed0db9361138d9391909216916002019061213b565b60405180910390a160996113a2600183611f98565b815481106113b2576113b2611eb3565b600091825260209091200154609980546001600160a01b0390921691849081106113de576113de611eb3565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081609760006099858154811061142457611424611eb3565b60009182526020808320909101546001600160a01b03908116845283820194909452604092830182206004019490945586831681526097845281812060010154909216825260989092522080547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560998054806114a6576114a661215d565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000090811690915593019093556001600160a01b0386168152609790925260408220805482168155600181018054909216909155906115396002830182611b76565b50600060038201819055600490910155505050565b6040516001600160a01b038085166024830152831660448201526064810182905261159f9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016112a4565b50505050565b6065546001600160a01b03163314610b975760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105f0565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166116e65760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105f0565b610b97611858565b600054610100900460ff1661176b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105f0565b610b976118de565b60006117c8826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166119619092919063ffffffff16565b80519091501561132657808060200190518101906117e6919061218c565b6113265760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016105f0565b600054610100900460ff166118d55760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105f0565b610b97336115ff565b600054610100900460ff1661195b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105f0565b60018055565b6060611970848460008561197a565b90505b9392505050565b6060824710156119f25760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016105f0565b843b611a405760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016105f0565b600080866001600160a01b03168587604051611a5c91906121ae565b60006040518083038185875af1925050503d8060008114611a99576040519150601f19603f3d011682016040523d82523d6000602084013e611a9e565b606091505b5091509150611aae828286611ab9565b979650505050505050565b60608315611ac8575081611973565b825115611ad85782518084602001fd5b8160405162461bcd60e51b81526004016105f091906121ca565b828054611afe90611ee2565b90600052602060002090601f016020900481019282611b205760008555611b66565b82601f10611b3957805160ff1916838001178555611b66565b82800160010185558215611b66579182015b82811115611b66578251825591602001919060010190611b4b565b50611b72929150611bac565b5090565b508054611b8290611ee2565b6000825580601f10611b92575050565b601f01602090049060005260206000209081019061125c91905b5b80821115611b725760008155600101611bad565b600060208284031215611bd357600080fd5b5035919050565b60005b83811015611bf5578181015183820152602001611bdd565b8381111561159f5750506000910152565b60008151808452611c1e816020860160208601611bda565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60006001600160a01b0380835116845280602084015116602085015250604082015160a06040850152611c8660a0850182611c06565b905060608301516060850152608083015160808501528091505092915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015611d19577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0888603018452611d07858351611c50565b94509285019290850190600101611ccd565b5092979650505050505050565b80356001600160a01b0381168114611d3d57600080fd5b919050565b600060208284031215611d5457600080fd5b61197382611d26565b60006001600160a01b03808816835280871660208401525060a06040830152611d8960a0830186611c06565b606083019490945250608001529392505050565b60008060008060608587031215611db357600080fd5b84359350611dc360208601611d26565b9250604085013567ffffffffffffffff80821115611de057600080fd5b818701915087601f830112611df457600080fd5b813581811115611e0357600080fd5b886020828501011115611e1557600080fd5b95989497505060200194505050565b6020808252825182820181905260009190848201906040850190845b81811015611e655783516001600160a01b031683529284019291840191600101611e40565b50909695505050505050565b6020815260006119736020830184611c50565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c90821680611ef657607f821691505b602082108114156104d7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611f9157611f91611f30565b5060010190565b600082821015611faa57611faa611f30565b500390565b8054600090600181811c9080831680611fc957607f831692505b6020808410821415612004577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b8388526020880182801561201f576001811461204e57612079565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00871682528282019750612079565b60008981526020902060005b878110156120735781548482015290860190840161205a565b83019850505b5050505050505092915050565b6001600160a01b03841681526060602082015260006120a86060830185611faf565b9050826040830152949350505050565b60006001600160a01b038087168352808616602084015250606060408301528260608301528284608084013760006080848401015260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f850116830101905095945050505050565b6000821982111561213657612136611f30565b500190565b6001600160a01b03831681526040602082015260006119706040830184611faf565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60006020828403121561219e57600080fd5b8151801515811461197357600080fd5b600082516121c0818460208701611bda565b9190910192915050565b6020815260006119736020830184611c0656fea2646970667358221220257051a3d41c4c557970aeacd0ddf1ceaa72f5bf2c642737583b6f41a89af7c064736f6c63430008090033"

// DeploySequencer deploys a new Ethereum contract, binding an instance of Sequencer to it.
func DeploySequencer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sequencer, error) {
	parsed, err := abi.JSON(strings.NewReader(SequencerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SequencerBin), backend)
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
	parsed, err := abi.JSON(strings.NewReader(SequencerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sequencer *SequencerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Sequencer *SequencerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function bitToken() constant returns(address)
func (_Sequencer *SequencerCaller) BitToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sequencer.contract.Call(opts, out, "bitToken")
	return *ret0, err
}

// BitToken is a free data retrieval call binding the contract method 0xd84e9f92.
//
// Solidity: function bitToken() constant returns(address)
func (_Sequencer *SequencerSession) BitToken() (common.Address, error) {
	return _Sequencer.Contract.BitToken(&_Sequencer.CallOpts)
}

// BitToken is a free data retrieval call binding the contract method 0xd84e9f92.
//
// Solidity: function bitToken() constant returns(address)
func (_Sequencer *SequencerCallerSession) BitToken() (common.Address, error) {
	return _Sequencer.Contract.BitToken(&_Sequencer.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Sequencer *SequencerCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Sequencer.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Sequencer *SequencerSession) GetOwners() ([]common.Address, error) {
	return _Sequencer.Contract.GetOwners(&_Sequencer.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Sequencer *SequencerCallerSession) GetOwners() ([]common.Address, error) {
	return _Sequencer.Contract.GetOwners(&_Sequencer.CallOpts)
}

// GetSequencer is a free data retrieval call binding the contract method 0xe90f218f.
//
// Solidity: function getSequencer(address signer) constant returns(SequencerSequencerInfo)
func (_Sequencer *SequencerCaller) GetSequencer(opts *bind.CallOpts, signer common.Address) (SequencerSequencerInfo, error) {
	var (
		ret0 = new(SequencerSequencerInfo)
	)
	out := ret0
	err := _Sequencer.contract.Call(opts, out, "getSequencer", signer)
	return *ret0, err
}

// GetSequencer is a free data retrieval call binding the contract method 0xe90f218f.
//
// Solidity: function getSequencer(address signer) constant returns(SequencerSequencerInfo)
func (_Sequencer *SequencerSession) GetSequencer(signer common.Address) (SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencer(&_Sequencer.CallOpts, signer)
}

// GetSequencer is a free data retrieval call binding the contract method 0xe90f218f.
//
// Solidity: function getSequencer(address signer) constant returns(SequencerSequencerInfo)
func (_Sequencer *SequencerCallerSession) GetSequencer(signer common.Address) (SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencer(&_Sequencer.CallOpts, signer)
}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() constant returns([]SequencerSequencerInfo)
func (_Sequencer *SequencerCaller) GetSequencers(opts *bind.CallOpts) ([]SequencerSequencerInfo, error) {
	var (
		ret0 = new([]SequencerSequencerInfo)
	)
	out := ret0
	err := _Sequencer.contract.Call(opts, out, "getSequencers")
	return *ret0, err
}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() constant returns([]SequencerSequencerInfo)
func (_Sequencer *SequencerSession) GetSequencers() ([]SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencers(&_Sequencer.CallOpts)
}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() constant returns([]SequencerSequencerInfo)
func (_Sequencer *SequencerCallerSession) GetSequencers() ([]SequencerSequencerInfo, error) {
	return _Sequencer.Contract.GetSequencers(&_Sequencer.CallOpts)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address signer) constant returns(bool)
func (_Sequencer *SequencerCaller) IsSequencer(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Sequencer.contract.Call(opts, out, "isSequencer", signer)
	return *ret0, err
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address signer) constant returns(bool)
func (_Sequencer *SequencerSession) IsSequencer(signer common.Address) (bool, error) {
	return _Sequencer.Contract.IsSequencer(&_Sequencer.CallOpts, signer)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address signer) constant returns(bool)
func (_Sequencer *SequencerCallerSession) IsSequencer(signer common.Address) (bool, error) {
	return _Sequencer.Contract.IsSequencer(&_Sequencer.CallOpts, signer)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sequencer *SequencerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sequencer.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sequencer *SequencerSession) Owner() (common.Address, error) {
	return _Sequencer.Contract.Owner(&_Sequencer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sequencer *SequencerCallerSession) Owner() (common.Address, error) {
	return _Sequencer.Contract.Owner(&_Sequencer.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_Sequencer *SequencerCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sequencer.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_Sequencer *SequencerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.Owners(&_Sequencer.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_Sequencer *SequencerCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.Owners(&_Sequencer.CallOpts, arg0)
}

// Rel is a free data retrieval call binding the contract method 0xcab2ea2a.
//
// Solidity: function rel(address ) constant returns(address)
func (_Sequencer *SequencerCaller) Rel(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sequencer.contract.Call(opts, out, "rel", arg0)
	return *ret0, err
}

// Rel is a free data retrieval call binding the contract method 0xcab2ea2a.
//
// Solidity: function rel(address ) constant returns(address)
func (_Sequencer *SequencerSession) Rel(arg0 common.Address) (common.Address, error) {
	return _Sequencer.Contract.Rel(&_Sequencer.CallOpts, arg0)
}

// Rel is a free data retrieval call binding the contract method 0xcab2ea2a.
//
// Solidity: function rel(address ) constant returns(address)
func (_Sequencer *SequencerCallerSession) Rel(arg0 common.Address) (common.Address, error) {
	return _Sequencer.Contract.Rel(&_Sequencer.CallOpts, arg0)
}

// Sequencers is a free data retrieval call binding the contract method 0x1c7a07ee.
//
// Solidity: function sequencers(address ) constant returns(address owner, address mintAddress, bytes nodeID, uint256 amount, uint256 keyIndex)
func (_Sequencer *SequencerCaller) Sequencers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner       common.Address
	MintAddress common.Address
	NodeID      []byte
	Amount      *big.Int
	KeyIndex    *big.Int
}, error) {
	ret := new(struct {
		Owner       common.Address
		MintAddress common.Address
		NodeID      []byte
		Amount      *big.Int
		KeyIndex    *big.Int
	})
	out := ret
	err := _Sequencer.contract.Call(opts, out, "sequencers", arg0)
	return *ret, err
}

// Sequencers is a free data retrieval call binding the contract method 0x1c7a07ee.
//
// Solidity: function sequencers(address ) constant returns(address owner, address mintAddress, bytes nodeID, uint256 amount, uint256 keyIndex)
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
// Solidity: function sequencers(address ) constant returns(address owner, address mintAddress, bytes nodeID, uint256 amount, uint256 keyIndex)
func (_Sequencer *SequencerCallerSession) Sequencers(arg0 common.Address) (struct {
	Owner       common.Address
	MintAddress common.Address
	NodeID      []byte
	Amount      *big.Int
	KeyIndex    *big.Int
}, error) {
	return _Sequencer.Contract.Sequencers(&_Sequencer.CallOpts, arg0)
}

// ChangeBitAddress is a paid mutator transaction binding the contract method 0x7a0674a0.
//
// Solidity: function changeBitAddress(address _bitToken) returns()
func (_Sequencer *SequencerTransactor) ChangeBitAddress(opts *bind.TransactOpts, _bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "changeBitAddress", _bitToken)
}

// ChangeBitAddress is a paid mutator transaction binding the contract method 0x7a0674a0.
//
// Solidity: function changeBitAddress(address _bitToken) returns()
func (_Sequencer *SequencerSession) ChangeBitAddress(_bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.ChangeBitAddress(&_Sequencer.TransactOpts, _bitToken)
}

// ChangeBitAddress is a paid mutator transaction binding the contract method 0x7a0674a0.
//
// Solidity: function changeBitAddress(address _bitToken) returns()
func (_Sequencer *SequencerTransactorSession) ChangeBitAddress(_bitToken common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.ChangeBitAddress(&_Sequencer.TransactOpts, _bitToken)
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
// Solidity: event SequencerCreate(address , address , bytes )
func (_Sequencer *SequencerFilterer) FilterSequencerCreate(opts *bind.FilterOpts) (*SequencerSequencerCreateIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "SequencerCreate")
	if err != nil {
		return nil, err
	}
	return &SequencerSequencerCreateIterator{contract: _Sequencer.contract, event: "SequencerCreate", logs: logs, sub: sub}, nil
}

// WatchSequencerCreate is a free log subscription operation binding the contract event 0x4bc782d7487c41941075eb09650b2eb45a57e23e2241db6a958b8af8485324f6.
//
// Solidity: event SequencerCreate(address , address , bytes )
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
// Solidity: event SequencerCreate(address , address , bytes )
func (_Sequencer *SequencerFilterer) ParseSequencerCreate(log types.Log) (*SequencerSequencerCreate, error) {
	event := new(SequencerSequencerCreate)
	if err := _Sequencer.contract.UnpackLog(event, "SequencerCreate", log); err != nil {
		return nil, err
	}
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
// Solidity: event SequencerDelete(address , bytes )
func (_Sequencer *SequencerFilterer) FilterSequencerDelete(opts *bind.FilterOpts) (*SequencerSequencerDeleteIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "SequencerDelete")
	if err != nil {
		return nil, err
	}
	return &SequencerSequencerDeleteIterator{contract: _Sequencer.contract, event: "SequencerDelete", logs: logs, sub: sub}, nil
}

// WatchSequencerDelete is a free log subscription operation binding the contract event 0x37c49a6eae060065fbdeff05623918ea0969b4b1146b4c2fd33658caa72ed0db.
//
// Solidity: event SequencerDelete(address , bytes )
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
// Solidity: event SequencerDelete(address , bytes )
func (_Sequencer *SequencerFilterer) ParseSequencerDelete(log types.Log) (*SequencerSequencerDelete, error) {
	event := new(SequencerSequencerDelete)
	if err := _Sequencer.contract.UnpackLog(event, "SequencerDelete", log); err != nil {
		return nil, err
	}
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
// Solidity: event SequencerUpdate(address , bytes , uint256 )
func (_Sequencer *SequencerFilterer) FilterSequencerUpdate(opts *bind.FilterOpts) (*SequencerSequencerUpdateIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "SequencerUpdate")
	if err != nil {
		return nil, err
	}
	return &SequencerSequencerUpdateIterator{contract: _Sequencer.contract, event: "SequencerUpdate", logs: logs, sub: sub}, nil
}

// WatchSequencerUpdate is a free log subscription operation binding the contract event 0xb60fdb7a00b5d98213e969843dc87e0f330162b9a4dc216e9b21412854b16824.
//
// Solidity: event SequencerUpdate(address , bytes , uint256 )
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
// Solidity: event SequencerUpdate(address , bytes , uint256 )
func (_Sequencer *SequencerFilterer) ParseSequencerUpdate(log types.Log) (*SequencerSequencerUpdate, error) {
	event := new(SequencerSequencerUpdate)
	if err := _Sequencer.contract.UnpackLog(event, "SequencerUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scc

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

// LibBVMCodecChainBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type LibBVMCodecChainBatchHeader struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	Signature         []byte
	ExtraData         []byte
}

// LibBVMCodecChainInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type LibBVMCodecChainInclusionProof struct {
	Index    *big.Int
	Siblings [][32]byte
}

// StateCommitmentChainMetaData contains all meta data concerning the StateCommitmentChain contract.
var StateCommitmentChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_libAddressManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fraudProofWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sequencerPublishWindow\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_batchSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prevTotalElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"StateBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchDeleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FRAUD_PROOF_WINDOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_PUBLISH_WINDOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_batch\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_shouldStartAtElement\",\"type\":\"uint256\"}],\"name\":\"appendStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"contractIChainStorageContainer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_BVMCodec.ChainBatchHeader\",\"name\":\"_batchHeader\",\"type\":\"tuple\"}],\"name\":\"deleteStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastSequencerTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastSequencerTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBatches\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalElements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_BVMCodec.ChainBatchHeader\",\"name\":\"_batchHeader\",\"type\":\"tuple\"}],\"name\":\"insideFraudProofWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_inside\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libAddressManager\",\"outputs\":[{\"internalType\":\"contractLib_AddressManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_element\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_BVMCodec.ChainBatchHeader\",\"name\":\"_batchHeader\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLib_BVMCodec.ChainInclusionProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyStateCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516200216a3803806200216a8339810160408190526100319161005d565b600080546001600160a01b0319166001600160a01b0394909416939093179092556001556002556100a0565b60008060006060848603121561007257600080fd5b83516001600160a01b038116811461008957600080fd5b602085015160409095015190969495509392505050565b6120ba80620000b06000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c806389a1d98011610081578063c84c85bf1161005b578063c84c85bf146101b1578063cfdf677e146101c4578063e561dddc146101cc57600080fd5b806389a1d98014610172578063b768bb1714610195578063c17b291b146101a857600080fd5b80637aa63a86116100b25780637aa63a861461014b5780637ad168a01461016157806381eb62ef1461016957600080fd5b8063299ca478146100d9578063461a4478146101235780635b4d90e214610136575b600080fd5b6000546100f99073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100f9610131366004611a75565b6101d4565b610149610144366004611b7d565b610281565b005b610153610442565b60405190815260200161011a565b61015361045b565b61015360025481565b610185610180366004611b7d565b610474565b604051901515815260200161011a565b6101856101a3366004611c32565b61051c565b61015360015481565b6101496101bf366004611cef565b6105e6565b6100f96109e1565b610153610a09565b600080546040517fbf40fac100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063bf40fac19061022b908590600401611dc7565b60206040518083038186803b15801561024357600080fd5b505afa158015610257573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061027b9190611dfc565b92915050565b6102bf6040518060400160405280601181526020017f42564d5f467261756456657269666965720000000000000000000000000000008152506101d4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103645760405162461bcd60e51b815260206004820152603b60248201527f537461746520626174636865732063616e206f6e6c792062652064656c65746560448201527f64206279207468652042564d5f467261756456657269666965722e000000000060648201526084015b60405180910390fd5b61036d81610a90565b6103b95760405162461bcd60e51b815260206004820152601560248201527f496e76616c6964206261746368206865616465722e0000000000000000000000604482015260640161035b565b6103c281610474565b610436576040805162461bcd60e51b81526020600482015260248101919091527f537461746520626174636865732063616e206f6e6c792062652064656c65746560448201527f642077697468696e207468652066726175642070726f6f662077696e646f772e606482015260840161035b565b61043f81610b52565b50565b60008061044d610d83565b5064ffffffffff1692915050565b600080610466610d83565b64ffffffffff169392505050565b6000808260a0015180602001905181019061048f9190611e19565b509050806105055760405162461bcd60e51b815260206004820152602560248201527f4261746368206865616465722074696d657374616d702063616e6e6f7420626560448201527f207a65726f000000000000000000000000000000000000000000000000000000606482015260840161035b565b42600154826105149190611e78565b119392505050565b600061052783610a90565b6105735760405162461bcd60e51b815260206004820152601560248201527f496e76616c6964206261746368206865616465722e0000000000000000000000604482015260640161035b565b610590836020015185846000015185602001518760400151610e27565b6105dc5760405162461bcd60e51b815260206004820152601860248201527f496e76616c696420696e636c7573696f6e2070726f6f662e0000000000000000604482015260640161035b565b5060019392505050565b6105ee610442565b81146106625760405162461bcd60e51b815260206004820152603d60248201527f41637475616c20626174636820737461727420696e64657820646f6573206e6f60448201527f74206d6174636820657870656374656420737461727420696e6465782e000000606482015260840161035b565b6106a06040518060400160405280600b81526020017f426f6e644d616e616765720000000000000000000000000000000000000000008152506101d4565b6040517f02ad4d2a00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff91909116906302ad4d2a9060240160206040518083038186803b15801561070757600080fd5b505afa15801561071b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073f9190611e90565b6107b15760405162461bcd60e51b815260206004820152602f60248201527f50726f706f73657220646f6573206e6f74206861766520656e6f75676820636f60448201527f6c6c61746572616c20706f737465640000000000000000000000000000000000606482015260840161035b565b60008351116108285760405162461bcd60e51b815260206004820152602360248201527f43616e6e6f74207375626d697420616e20656d7074792073746174652062617460448201527f63682e0000000000000000000000000000000000000000000000000000000000606482015260840161035b565b6108666040518060400160405280601981526020017f43616e6f6e6963616c5472616e73616374696f6e436861696e000000000000008152506101d4565b73ffffffffffffffffffffffffffffffffffffffff16637aa63a866040518163ffffffff1660e01b815260040160206040518083038186803b1580156108ab57600080fd5b505afa1580156108bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e39190611eb2565b83516108ed610442565b6108f79190611e78565b11156109915760405162461bcd60e51b815260206004820152604960248201527f4e756d626572206f6620737461746520726f6f74732063616e6e6f742065786360448201527f65656420746865206e756d626572206f662063616e6f6e6963616c207472616e60648201527f73616374696f6e732e0000000000000000000000000000000000000000000000608482015260a40161035b565b6109dc838342336040516020016109c892919091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b604051602081830303815290604052611095565b505050565b6000610a04604051806060016040528060218152602001612064602191396101d4565b905090565b6000610a136109e1565b73ffffffffffffffffffffffffffffffffffffffff16631f7b6d326040518163ffffffff1660e01b815260040160206040518083038186803b158015610a5857600080fd5b505afa158015610a6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a049190611eb2565b6000610a9a6109e1565b82516040517f9507d39a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9290921691639507d39a91610af29160040190815260200190565b60206040518083038186803b158015610b0a57600080fd5b505afa158015610b1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b429190611eb2565b610b4b83611345565b1492915050565b610b5a6109e1565b73ffffffffffffffffffffffffffffffffffffffff16631f7b6d326040518163ffffffff1660e01b815260040160206040518083038186803b158015610b9f57600080fd5b505afa158015610bb3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bd79190611eb2565b815110610c265760405162461bcd60e51b815260206004820152601460248201527f496e76616c696420626174636820696e6465782e000000000000000000000000604482015260640161035b565b610c2f81610a90565b610c7b5760405162461bcd60e51b815260206004820152601560248201527f496e76616c6964206261746368206865616465722e0000000000000000000000604482015260640161035b565b610c836109e1565b8151606083015173ffffffffffffffffffffffffffffffffffffffff929092169163167fd681919060281b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815260048101929092527fffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000166024820152604401600060405180830381600087803b158015610d2657600080fd5b505af1158015610d3a573d6000803e3d6000fd5b5050505080600001517f8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd648260200151604051610d7891815260200190565b60405180910390a250565b6000806000610d906109e1565b73ffffffffffffffffffffffffffffffffffffffff1663ccf8f9696040518163ffffffff1660e01b815260040160206040518083038186803b158015610dd557600080fd5b505afa158015610de9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e0d9190611ecb565b64ffffffffff602882901c169460509190911c9350915050565b6000808211610e9e5760405162461bcd60e51b815260206004820152603760248201527f4c69625f4d65726b6c65547265653a20546f74616c206c6561766573206d757360448201527f742062652067726561746572207468616e207a65726f2e000000000000000000606482015260840161035b565b818410610f125760405162461bcd60e51b8152602060048201526024808201527f4c69625f4d65726b6c65547265653a20496e646578206f7574206f6620626f7560448201527f6e64732e00000000000000000000000000000000000000000000000000000000606482015260840161035b565b610f1b8261138e565b835114610fb65760405162461bcd60e51b815260206004820152604d60248201527f4c69625f4d65726b6c65547265653a20546f74616c207369626c696e6773206460448201527f6f6573206e6f7420636f72726563746c7920636f72726573706f6e6420746f2060648201527f746f74616c206c65617665732e00000000000000000000000000000000000000608482015260a40161035b565b8460005b845181101561108857856001166001141561102157848181518110610fe157610fe1611f0d565b602002602001015182604051602001611004929190918252602082015260400190565b60405160208183030381529060405280519060200120915061106f565b8185828151811061103457611034611f0d565b6020026020010151604051602001611056929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b60019590951c948061108081611f3c565b915050610fba565b5090951495945050505050565b60006110d56040518060400160405280600c81526020017f42564d5f50726f706f73657200000000000000000000000000000000000000008152506101d4565b90506000806110e2610d83565b90925090503373ffffffffffffffffffffffffffffffffffffffff8416141561110c5750426111bb565b426002548264ffffffffff166111229190611e78565b106111bb5760405162461bcd60e51b815260206004820152604360248201527f43616e6e6f74207075626c69736820737461746520726f6f747320776974686960448201527f6e207468652073657175656e636572207075626c69636174696f6e2077696e6460648201527f6f772e0000000000000000000000000000000000000000000000000000000000608482015260a40161035b565b60006040518060c001604052806111d0610a09565b81526020016111de89611474565b8152602001885181526020018464ffffffffff16815260200187815260200186815250905080600001517f9cf3ad24eae3fd6d461e2f566b35b95b6d671871d9fcb45f8ac8030e4a8d21b382602001518360400151846060015185608001518660a00151604051611253959493929190611f75565b60405180910390a26112636109e1565b73ffffffffffffffffffffffffffffffffffffffff16632015276c61128783611345565b6112ac8460400151856060015161129e9190611e78565b602887811b91909117901b90565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815260048101929092527fffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000166024820152604401600060405180830381600087803b15801561132457600080fd5b505af1158015611338573d6000803e3d6000fd5b5050505050505050505050565b6020808201516040808401516060850151608086015160a0870151935160009661137196959101611f75565b604051602081830303815290604052805190602001209050919050565b60008082116114055760405162461bcd60e51b815260206004820152603060248201527f4c69625f4d65726b6c65547265653a2043616e6e6f7420636f6d70757465206360448201527f65696c286c6f675f3229206f6620302e00000000000000000000000000000000606482015260840161035b565b816001141561141657506000919050565b81600060805b600181106114545780611432600180831b611fb8565b901b83161561144c576114458183611e78565b92811c9291505b60011c61141c565b506001811b841461146d5761146a600182611e78565b90505b9392505050565b6000808251116114ec5760405162461bcd60e51b815260206004820152603460248201527f4c69625f4d65726b6c65547265653a204d7573742070726f766964652061742060448201527f6c65617374206f6e65206c65616620686173682e000000000000000000000000606482015260840161035b565b815160011415611518578160008151811061150957611509611f0d565b60200260200101519050919050565b60408051610200810182527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56381527f633dc4d7da7256660a892f8f1604a44b5432649cc8ec5cb3ced4c4e6ac94dd1d60208201527f890740a8eb06ce9be422cb8da5cdafc2b58c0a5e24036c578de2a433c828ff7d818301527f3b8ec09e026fdc305365dfc94e189a81b38c7597b3d941c279f042e8206e0bd86060808301919091527fecd50eee38e386bd62be9bedb990706951b65fe053bd9d8a521af753d139e2da60808301527fdefff6d330bb5403f63b14f33b578274160de3a50df4efecf0e0db73bcdd3da560a08301527f617bdd11f7c0a11f49db22f629387a12da7596f9d1704d7465177c63d88ec7d760c08301527f292c23a9aa1d8bea7e2435e555a4a60e379a5a35f3f452bae60121073fb6eead60e08301527fe1cea92ed99acdcb045a6726b2f87107e8a61620a232cf4d7d5b5766b3952e106101008301527f7ad66c0a68c72cb89e4fb4303841966e4062a76ab97451e3b9fb526a5ceb7f826101208301527fe026cc5a4aed3c22a58cbd3d2ac754c9352c5436f638042dca99034e836365166101408301527f3d04cffd8b46a874edf5cfae63077de85f849a660426697b06a829c70dd1409c6101608301527fad676aa337a485e4728a0b240d92b3ef7b3c372d06d189322bfd5f61f1e7203e6101808301527fa2fca4a49658f9fab7aa63289c91b7c7b6c832a6d0e69334ff5b0a3483d09dab6101a08301527f4ebfd9cd7bca2505f7bef59cc1c12ecc708fff26ae4af19abe852afe9e20c8626101c08301527f2def10d13dd169f550f578bda343d9717a138562e0093b380a1120789d53cf106101e083015282518381529081018352909160009190602082018180368337505085519192506000918291508180805b600184111561192e576117c9600285611ffe565b91506117d6600285612012565b600114905060005b82811015611882578a6117f2826002612026565b8151811061180257611802611f0d565b602002602001015196508a81600261181a9190612026565b611825906001611e78565b8151811061183557611835611f0d565b6020026020010151955086602089015285604089015287805190602001208b828151811061186557611865611f0d565b60209081029190910101528061187a81611f3c565b9150506117de565b5080156118fe5789611895600186611fb8565b815181106118a5576118a5611f0d565b602002602001015195508783601081106118c1576118c1611f0d565b602002015160001b945085602088015284604088015286805190602001208a83815181106118f1576118f1611f0d565b6020026020010181815250505b8061190a57600061190d565b60015b61191a9060ff1683611e78565b93508261192681611f3c565b9350506117b5565b8960008151811061194157611941611f0d565b602002602001015198505050505050505050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160c0810167ffffffffffffffff811182821017156119aa576119aa611958565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156119f7576119f7611958565b604052919050565b600067ffffffffffffffff831115611a1957611a19611958565b611a4a60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f860116016119b0565b9050828152838383011115611a5e57600080fd5b828260208301376000602084830101529392505050565b600060208284031215611a8757600080fd5b813567ffffffffffffffff811115611a9e57600080fd5b8201601f81018413611aaf57600080fd5b611abe848235602084016119ff565b949350505050565b600082601f830112611ad757600080fd5b61146d838335602085016119ff565b600060c08284031215611af857600080fd5b611b00611987565b905081358152602082013560208201526040820135604082015260608201356060820152608082013567ffffffffffffffff80821115611b3f57600080fd5b611b4b85838601611ac6565b608084015260a0840135915080821115611b6457600080fd5b50611b7184828501611ac6565b60a08301525092915050565b600060208284031215611b8f57600080fd5b813567ffffffffffffffff811115611ba657600080fd5b611abe84828501611ae6565b600082601f830112611bc357600080fd5b8135602067ffffffffffffffff821115611bdf57611bdf611958565b8160051b611bee8282016119b0565b9283528481018201928281019087851115611c0857600080fd5b83870192505b84831015611c2757823582529183019190830190611c0e565b979650505050505050565b600080600060608486031215611c4757600080fd5b83359250602084013567ffffffffffffffff80821115611c6657600080fd5b611c7287838801611ae6565b93506040860135915080821115611c8857600080fd5b9085019060408288031215611c9c57600080fd5b604051604081018181108382111715611cb757611cb7611958565b60405282358152602083013582811115611cd057600080fd5b611cdc89828601611bb2565b6020830152508093505050509250925092565b600080600060608486031215611d0457600080fd5b833567ffffffffffffffff80821115611d1c57600080fd5b611d2887838801611bb2565b94506020860135915080821115611d3e57600080fd5b50611d4b86828701611ac6565b925050604084013590509250925092565b6000815180845260005b81811015611d8257602081850181015186830182015201611d66565b81811115611d94576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061146d6020830184611d5c565b73ffffffffffffffffffffffffffffffffffffffff8116811461043f57600080fd5b600060208284031215611e0e57600080fd5b815161146d81611dda565b60008060408385031215611e2c57600080fd5b825191506020830151611e3e81611dda565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611e8b57611e8b611e49565b500190565b600060208284031215611ea257600080fd5b8151801515811461146d57600080fd5b600060208284031215611ec457600080fd5b5051919050565b600060208284031215611edd57600080fd5b81517fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000008116811461146d57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611f6e57611f6e611e49565b5060010190565b85815284602082015283604082015260a060608201526000611f9a60a0830185611d5c565b8281036080840152611fac8185611d5c565b98975050505050505050565b600082821015611fca57611fca611e49565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261200d5761200d611fcf565b500490565b60008261202157612021611fcf565b500690565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561205e5761205e611e49565b50029056fe436861696e53746f72616765436f6e7461696e65722d5343432d62617463686573a2646970667358221220fe8e95c7c72157e1e0b3bdc7a8375f496e0f2e5a004c82e3bfb384602378e1d764736f6c63430008090033",
}

// StateCommitmentChainABI is the input ABI used to generate the binding from.
// Deprecated: Use StateCommitmentChainMetaData.ABI instead.
var StateCommitmentChainABI = StateCommitmentChainMetaData.ABI

// StateCommitmentChainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateCommitmentChainMetaData.Bin instead.
var StateCommitmentChainBin = StateCommitmentChainMetaData.Bin

// DeployStateCommitmentChain deploys a new Ethereum contract, binding an instance of StateCommitmentChain to it.
func DeployStateCommitmentChain(auth *bind.TransactOpts, backend bind.ContractBackend, _libAddressManager common.Address, _fraudProofWindow *big.Int, _sequencerPublishWindow *big.Int) (common.Address, *types.Transaction, *StateCommitmentChain, error) {
	parsed, err := StateCommitmentChainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateCommitmentChainBin), backend, _libAddressManager, _fraudProofWindow, _sequencerPublishWindow)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateCommitmentChain{StateCommitmentChainCaller: StateCommitmentChainCaller{contract: contract}, StateCommitmentChainTransactor: StateCommitmentChainTransactor{contract: contract}, StateCommitmentChainFilterer: StateCommitmentChainFilterer{contract: contract}}, nil
}

// StateCommitmentChain is an auto generated Go binding around an Ethereum contract.
type StateCommitmentChain struct {
	StateCommitmentChainCaller     // Read-only binding to the contract
	StateCommitmentChainTransactor // Write-only binding to the contract
	StateCommitmentChainFilterer   // Log filterer for contract events
}

// StateCommitmentChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateCommitmentChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateCommitmentChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateCommitmentChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateCommitmentChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateCommitmentChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateCommitmentChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateCommitmentChainSession struct {
	Contract     *StateCommitmentChain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StateCommitmentChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateCommitmentChainCallerSession struct {
	Contract *StateCommitmentChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// StateCommitmentChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateCommitmentChainTransactorSession struct {
	Contract     *StateCommitmentChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// StateCommitmentChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateCommitmentChainRaw struct {
	Contract *StateCommitmentChain // Generic contract binding to access the raw methods on
}

// StateCommitmentChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateCommitmentChainCallerRaw struct {
	Contract *StateCommitmentChainCaller // Generic read-only contract binding to access the raw methods on
}

// StateCommitmentChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateCommitmentChainTransactorRaw struct {
	Contract *StateCommitmentChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateCommitmentChain creates a new instance of StateCommitmentChain, bound to a specific deployed contract.
func NewStateCommitmentChain(address common.Address, backend bind.ContractBackend) (*StateCommitmentChain, error) {
	contract, err := bindStateCommitmentChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChain{StateCommitmentChainCaller: StateCommitmentChainCaller{contract: contract}, StateCommitmentChainTransactor: StateCommitmentChainTransactor{contract: contract}, StateCommitmentChainFilterer: StateCommitmentChainFilterer{contract: contract}}, nil
}

// NewStateCommitmentChainCaller creates a new read-only instance of StateCommitmentChain, bound to a specific deployed contract.
func NewStateCommitmentChainCaller(address common.Address, caller bind.ContractCaller) (*StateCommitmentChainCaller, error) {
	contract, err := bindStateCommitmentChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChainCaller{contract: contract}, nil
}

// NewStateCommitmentChainTransactor creates a new write-only instance of StateCommitmentChain, bound to a specific deployed contract.
func NewStateCommitmentChainTransactor(address common.Address, transactor bind.ContractTransactor) (*StateCommitmentChainTransactor, error) {
	contract, err := bindStateCommitmentChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChainTransactor{contract: contract}, nil
}

// NewStateCommitmentChainFilterer creates a new log filterer instance of StateCommitmentChain, bound to a specific deployed contract.
func NewStateCommitmentChainFilterer(address common.Address, filterer bind.ContractFilterer) (*StateCommitmentChainFilterer, error) {
	contract, err := bindStateCommitmentChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChainFilterer{contract: contract}, nil
}

// bindStateCommitmentChain binds a generic wrapper to an already deployed contract.
func bindStateCommitmentChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateCommitmentChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateCommitmentChain *StateCommitmentChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateCommitmentChain.Contract.StateCommitmentChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateCommitmentChain *StateCommitmentChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.StateCommitmentChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateCommitmentChain *StateCommitmentChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.StateCommitmentChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateCommitmentChain *StateCommitmentChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateCommitmentChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateCommitmentChain *StateCommitmentChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateCommitmentChain *StateCommitmentChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.contract.Transact(opts, method, params...)
}

// FRAUDPROOFWINDOW is a free data retrieval call binding the contract method 0xc17b291b.
//
// Solidity: function FRAUD_PROOF_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainCaller) FRAUDPROOFWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "FRAUD_PROOF_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FRAUDPROOFWINDOW is a free data retrieval call binding the contract method 0xc17b291b.
//
// Solidity: function FRAUD_PROOF_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainSession) FRAUDPROOFWINDOW() (*big.Int, error) {
	return _StateCommitmentChain.Contract.FRAUDPROOFWINDOW(&_StateCommitmentChain.CallOpts)
}

// FRAUDPROOFWINDOW is a free data retrieval call binding the contract method 0xc17b291b.
//
// Solidity: function FRAUD_PROOF_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) FRAUDPROOFWINDOW() (*big.Int, error) {
	return _StateCommitmentChain.Contract.FRAUDPROOFWINDOW(&_StateCommitmentChain.CallOpts)
}

// SEQUENCERPUBLISHWINDOW is a free data retrieval call binding the contract method 0x81eb62ef.
//
// Solidity: function SEQUENCER_PUBLISH_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainCaller) SEQUENCERPUBLISHWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "SEQUENCER_PUBLISH_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SEQUENCERPUBLISHWINDOW is a free data retrieval call binding the contract method 0x81eb62ef.
//
// Solidity: function SEQUENCER_PUBLISH_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainSession) SEQUENCERPUBLISHWINDOW() (*big.Int, error) {
	return _StateCommitmentChain.Contract.SEQUENCERPUBLISHWINDOW(&_StateCommitmentChain.CallOpts)
}

// SEQUENCERPUBLISHWINDOW is a free data retrieval call binding the contract method 0x81eb62ef.
//
// Solidity: function SEQUENCER_PUBLISH_WINDOW() view returns(uint256)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) SEQUENCERPUBLISHWINDOW() (*big.Int, error) {
	return _StateCommitmentChain.Contract.SEQUENCERPUBLISHWINDOW(&_StateCommitmentChain.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCaller) Batches(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "batches")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainSession) Batches() (common.Address, error) {
	return _StateCommitmentChain.Contract.Batches(&_StateCommitmentChain.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) Batches() (common.Address, error) {
	return _StateCommitmentChain.Contract.Batches(&_StateCommitmentChain.CallOpts)
}

// GetLastSequencerTimestamp is a free data retrieval call binding the contract method 0x7ad168a0.
//
// Solidity: function getLastSequencerTimestamp() view returns(uint256 _lastSequencerTimestamp)
func (_StateCommitmentChain *StateCommitmentChainCaller) GetLastSequencerTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "getLastSequencerTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastSequencerTimestamp is a free data retrieval call binding the contract method 0x7ad168a0.
//
// Solidity: function getLastSequencerTimestamp() view returns(uint256 _lastSequencerTimestamp)
func (_StateCommitmentChain *StateCommitmentChainSession) GetLastSequencerTimestamp() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetLastSequencerTimestamp(&_StateCommitmentChain.CallOpts)
}

// GetLastSequencerTimestamp is a free data retrieval call binding the contract method 0x7ad168a0.
//
// Solidity: function getLastSequencerTimestamp() view returns(uint256 _lastSequencerTimestamp)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) GetLastSequencerTimestamp() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetLastSequencerTimestamp(&_StateCommitmentChain.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_StateCommitmentChain *StateCommitmentChainCaller) GetTotalBatches(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "getTotalBatches")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_StateCommitmentChain *StateCommitmentChainSession) GetTotalBatches() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetTotalBatches(&_StateCommitmentChain.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) GetTotalBatches() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetTotalBatches(&_StateCommitmentChain.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_StateCommitmentChain *StateCommitmentChainCaller) GetTotalElements(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "getTotalElements")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_StateCommitmentChain *StateCommitmentChainSession) GetTotalElements() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetTotalElements(&_StateCommitmentChain.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) GetTotalElements() (*big.Int, error) {
	return _StateCommitmentChain.Contract.GetTotalElements(&_StateCommitmentChain.CallOpts)
}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0x89a1d980.
//
// Solidity: function insideFraudProofWindow((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) view returns(bool _inside)
func (_StateCommitmentChain *StateCommitmentChainCaller) InsideFraudProofWindow(opts *bind.CallOpts, _batchHeader LibBVMCodecChainBatchHeader) (bool, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "insideFraudProofWindow", _batchHeader)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0x89a1d980.
//
// Solidity: function insideFraudProofWindow((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) view returns(bool _inside)
func (_StateCommitmentChain *StateCommitmentChainSession) InsideFraudProofWindow(_batchHeader LibBVMCodecChainBatchHeader) (bool, error) {
	return _StateCommitmentChain.Contract.InsideFraudProofWindow(&_StateCommitmentChain.CallOpts, _batchHeader)
}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0x89a1d980.
//
// Solidity: function insideFraudProofWindow((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) view returns(bool _inside)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) InsideFraudProofWindow(_batchHeader LibBVMCodecChainBatchHeader) (bool, error) {
	return _StateCommitmentChain.Contract.InsideFraudProofWindow(&_StateCommitmentChain.CallOpts, _batchHeader)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCaller) LibAddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "libAddressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainSession) LibAddressManager() (common.Address, error) {
	return _StateCommitmentChain.Contract.LibAddressManager(&_StateCommitmentChain.CallOpts)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) LibAddressManager() (common.Address, error) {
	return _StateCommitmentChain.Contract.LibAddressManager(&_StateCommitmentChain.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCaller) Resolve(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "resolve", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_StateCommitmentChain *StateCommitmentChainSession) Resolve(_name string) (common.Address, error) {
	return _StateCommitmentChain.Contract.Resolve(&_StateCommitmentChain.CallOpts, _name)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) Resolve(_name string) (common.Address, error) {
	return _StateCommitmentChain.Contract.Resolve(&_StateCommitmentChain.CallOpts, _name)
}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0xb768bb17.
//
// Solidity: function verifyStateCommitment(bytes32 _element, (uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader, (uint256,bytes32[]) _proof) view returns(bool)
func (_StateCommitmentChain *StateCommitmentChainCaller) VerifyStateCommitment(opts *bind.CallOpts, _element [32]byte, _batchHeader LibBVMCodecChainBatchHeader, _proof LibBVMCodecChainInclusionProof) (bool, error) {
	var out []interface{}
	err := _StateCommitmentChain.contract.Call(opts, &out, "verifyStateCommitment", _element, _batchHeader, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0xb768bb17.
//
// Solidity: function verifyStateCommitment(bytes32 _element, (uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader, (uint256,bytes32[]) _proof) view returns(bool)
func (_StateCommitmentChain *StateCommitmentChainSession) VerifyStateCommitment(_element [32]byte, _batchHeader LibBVMCodecChainBatchHeader, _proof LibBVMCodecChainInclusionProof) (bool, error) {
	return _StateCommitmentChain.Contract.VerifyStateCommitment(&_StateCommitmentChain.CallOpts, _element, _batchHeader, _proof)
}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0xb768bb17.
//
// Solidity: function verifyStateCommitment(bytes32 _element, (uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader, (uint256,bytes32[]) _proof) view returns(bool)
func (_StateCommitmentChain *StateCommitmentChainCallerSession) VerifyStateCommitment(_element [32]byte, _batchHeader LibBVMCodecChainBatchHeader, _proof LibBVMCodecChainInclusionProof) (bool, error) {
	return _StateCommitmentChain.Contract.VerifyStateCommitment(&_StateCommitmentChain.CallOpts, _element, _batchHeader, _proof)
}

// AppendStateBatch is a paid mutator transaction binding the contract method 0xc84c85bf.
//
// Solidity: function appendStateBatch(bytes32[] _batch, bytes _signature, uint256 _shouldStartAtElement) returns()
func (_StateCommitmentChain *StateCommitmentChainTransactor) AppendStateBatch(opts *bind.TransactOpts, _batch [][32]byte, _signature []byte, _shouldStartAtElement *big.Int) (*types.Transaction, error) {
	return _StateCommitmentChain.contract.Transact(opts, "appendStateBatch", _batch, _signature, _shouldStartAtElement)
}

// AppendStateBatch is a paid mutator transaction binding the contract method 0xc84c85bf.
//
// Solidity: function appendStateBatch(bytes32[] _batch, bytes _signature, uint256 _shouldStartAtElement) returns()
func (_StateCommitmentChain *StateCommitmentChainSession) AppendStateBatch(_batch [][32]byte, _signature []byte, _shouldStartAtElement *big.Int) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.AppendStateBatch(&_StateCommitmentChain.TransactOpts, _batch, _signature, _shouldStartAtElement)
}

// AppendStateBatch is a paid mutator transaction binding the contract method 0xc84c85bf.
//
// Solidity: function appendStateBatch(bytes32[] _batch, bytes _signature, uint256 _shouldStartAtElement) returns()
func (_StateCommitmentChain *StateCommitmentChainTransactorSession) AppendStateBatch(_batch [][32]byte, _signature []byte, _shouldStartAtElement *big.Int) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.AppendStateBatch(&_StateCommitmentChain.TransactOpts, _batch, _signature, _shouldStartAtElement)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0x5b4d90e2.
//
// Solidity: function deleteStateBatch((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) returns()
func (_StateCommitmentChain *StateCommitmentChainTransactor) DeleteStateBatch(opts *bind.TransactOpts, _batchHeader LibBVMCodecChainBatchHeader) (*types.Transaction, error) {
	return _StateCommitmentChain.contract.Transact(opts, "deleteStateBatch", _batchHeader)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0x5b4d90e2.
//
// Solidity: function deleteStateBatch((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) returns()
func (_StateCommitmentChain *StateCommitmentChainSession) DeleteStateBatch(_batchHeader LibBVMCodecChainBatchHeader) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.DeleteStateBatch(&_StateCommitmentChain.TransactOpts, _batchHeader)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0x5b4d90e2.
//
// Solidity: function deleteStateBatch((uint256,bytes32,uint256,uint256,bytes,bytes) _batchHeader) returns()
func (_StateCommitmentChain *StateCommitmentChainTransactorSession) DeleteStateBatch(_batchHeader LibBVMCodecChainBatchHeader) (*types.Transaction, error) {
	return _StateCommitmentChain.Contract.DeleteStateBatch(&_StateCommitmentChain.TransactOpts, _batchHeader)
}

// StateCommitmentChainStateBatchAppendedIterator is returned from FilterStateBatchAppended and is used to iterate over the raw logs and unpacked data for StateBatchAppended events raised by the StateCommitmentChain contract.
type StateCommitmentChainStateBatchAppendedIterator struct {
	Event *StateCommitmentChainStateBatchAppended // Event containing the contract specifics and raw log

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
func (it *StateCommitmentChainStateBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateCommitmentChainStateBatchAppended)
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
		it.Event = new(StateCommitmentChainStateBatchAppended)
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
func (it *StateCommitmentChainStateBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateCommitmentChainStateBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateCommitmentChainStateBatchAppended represents a StateBatchAppended event raised by the StateCommitmentChain contract.
type StateCommitmentChainStateBatchAppended struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	Signature         []byte
	ExtraData         []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStateBatchAppended is a free log retrieval operation binding the contract event 0x9cf3ad24eae3fd6d461e2f566b35b95b6d671871d9fcb45f8ac8030e4a8d21b3.
//
// Solidity: event StateBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes signature, bytes _extraData)
func (_StateCommitmentChain *StateCommitmentChainFilterer) FilterStateBatchAppended(opts *bind.FilterOpts, _batchIndex []*big.Int) (*StateCommitmentChainStateBatchAppendedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _StateCommitmentChain.contract.FilterLogs(opts, "StateBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChainStateBatchAppendedIterator{contract: _StateCommitmentChain.contract, event: "StateBatchAppended", logs: logs, sub: sub}, nil
}

// WatchStateBatchAppended is a free log subscription operation binding the contract event 0x9cf3ad24eae3fd6d461e2f566b35b95b6d671871d9fcb45f8ac8030e4a8d21b3.
//
// Solidity: event StateBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes signature, bytes _extraData)
func (_StateCommitmentChain *StateCommitmentChainFilterer) WatchStateBatchAppended(opts *bind.WatchOpts, sink chan<- *StateCommitmentChainStateBatchAppended, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _StateCommitmentChain.contract.WatchLogs(opts, "StateBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateCommitmentChainStateBatchAppended)
				if err := _StateCommitmentChain.contract.UnpackLog(event, "StateBatchAppended", log); err != nil {
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

// ParseStateBatchAppended is a log parse operation binding the contract event 0x9cf3ad24eae3fd6d461e2f566b35b95b6d671871d9fcb45f8ac8030e4a8d21b3.
//
// Solidity: event StateBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes signature, bytes _extraData)
func (_StateCommitmentChain *StateCommitmentChainFilterer) ParseStateBatchAppended(log types.Log) (*StateCommitmentChainStateBatchAppended, error) {
	event := new(StateCommitmentChainStateBatchAppended)
	if err := _StateCommitmentChain.contract.UnpackLog(event, "StateBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateCommitmentChainStateBatchDeletedIterator is returned from FilterStateBatchDeleted and is used to iterate over the raw logs and unpacked data for StateBatchDeleted events raised by the StateCommitmentChain contract.
type StateCommitmentChainStateBatchDeletedIterator struct {
	Event *StateCommitmentChainStateBatchDeleted // Event containing the contract specifics and raw log

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
func (it *StateCommitmentChainStateBatchDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateCommitmentChainStateBatchDeleted)
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
		it.Event = new(StateCommitmentChainStateBatchDeleted)
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
func (it *StateCommitmentChainStateBatchDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateCommitmentChainStateBatchDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateCommitmentChainStateBatchDeleted represents a StateBatchDeleted event raised by the StateCommitmentChain contract.
type StateCommitmentChainStateBatchDeleted struct {
	BatchIndex *big.Int
	BatchRoot  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStateBatchDeleted is a free log retrieval operation binding the contract event 0x8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64.
//
// Solidity: event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_StateCommitmentChain *StateCommitmentChainFilterer) FilterStateBatchDeleted(opts *bind.FilterOpts, _batchIndex []*big.Int) (*StateCommitmentChainStateBatchDeletedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _StateCommitmentChain.contract.FilterLogs(opts, "StateBatchDeleted", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &StateCommitmentChainStateBatchDeletedIterator{contract: _StateCommitmentChain.contract, event: "StateBatchDeleted", logs: logs, sub: sub}, nil
}

// WatchStateBatchDeleted is a free log subscription operation binding the contract event 0x8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64.
//
// Solidity: event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_StateCommitmentChain *StateCommitmentChainFilterer) WatchStateBatchDeleted(opts *bind.WatchOpts, sink chan<- *StateCommitmentChainStateBatchDeleted, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _StateCommitmentChain.contract.WatchLogs(opts, "StateBatchDeleted", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateCommitmentChainStateBatchDeleted)
				if err := _StateCommitmentChain.contract.UnpackLog(event, "StateBatchDeleted", log); err != nil {
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

// ParseStateBatchDeleted is a log parse operation binding the contract event 0x8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64.
//
// Solidity: event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_StateCommitmentChain *StateCommitmentChainFilterer) ParseStateBatchDeleted(log types.Log) (*StateCommitmentChainStateBatchDeleted, error) {
	event := new(StateCommitmentChainStateBatchDeleted)
	if err := _StateCommitmentChain.contract.UnpackLog(event, "StateBatchDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

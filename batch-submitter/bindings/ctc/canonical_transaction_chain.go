// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ctc

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

// LibBVMCodecQueueElement is an auto generated low-level Go binding around an user-defined struct.
type LibBVMCodecQueueElement struct {
	TransactionHash [32]byte
	Timestamp       *big.Int
	BlockNumber     *big.Int
}

// CanonicalTransactionChainMetaData contains all meta data concerning the CanonicalTransactionChain contract.
var CanonicalTransactionChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_libAddressManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxTransactionGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasDiscountDivisor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_enqueueGasCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mt_batcher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l2GasDiscountDivisor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"enqueueGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"enqueueL2GasPrepaid\",\"type\":\"uint256\"}],\"name\":\"L2GasParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startingQueueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_numQueueElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"name\":\"QueueBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startingQueueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_numQueueElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"name\":\"SequencerBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_batchSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prevTotalElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"TransactionBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1TxOrigin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"TransactionEnqueued\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ROLLUP_TX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_ROLLUP_TX_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_shouldStartAtElement\",\"type\":\"uint40\"},{\"internalType\":\"uint24\",\"name\":\"_totalElementsToAppend\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"_numSequencedTransactions\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"_numSubsequentQueueTransactions\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"_timestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"_blockNumber\",\"type\":\"uint40\"}],\"name\":\"appendDaSequencerBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appendSequencerBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"contractIChainStorageContainer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"enqueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enqueueGasCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enqueueL2GasPrepaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockNumber\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastTimestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextQueueIndex\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumPendingQueueElements\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getQueueElement\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"}],\"internalType\":\"structLib_BVMCodec.QueueElement\",\"name\":\"_element\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQueueLength\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBatches\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalElements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2GasDiscountDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libAddressManager\",\"outputs\":[{\"internalType\":\"contractLib_AddressManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTransactionGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mt_batcher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2GasDiscountDivisor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_enqueueGasCost\",\"type\":\"uint256\"}],\"name\":\"setGasParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001f7938038062001f798339810160408190526200003491620000b5565b600080546001600160a01b0319166001600160a01b0387161790556004849055600283905560018290556200006a82846200010a565b600355600580546001600160a01b0319166001600160a01b0392909216919091179055506200013892505050565b80516001600160a01b0381168114620000b057600080fd5b919050565b600080600080600060a08688031215620000ce57600080fd5b620000d98662000098565b9450602086015193506040860151925060608601519150620000fe6080870162000098565b90509295509295909350565b60008160001904831182151516156200013357634e487b7160e01b600052601160045260246000fd5b500290565b611e3180620001486000396000f3fe608060405234801561001057600080fd5b50600436106101825760003560e01c8063876ed5cb116100d8578063d0f893441161008c578063edcc4a4511610066578063edcc4a4514610317578063f4f7911d1461032a578063f722b41a1461033d57600080fd5b8063d0f89344146102fe578063e561dddc14610306578063e654b1fb1461030e57600080fd5b8063b8f77005116100bd578063b8f77005146102e5578063ccf987c8146102ed578063cfdf677e146102f657600080fd5b8063876ed5cb146102d35780638d38c6c1146102dc57600080fd5b8063461a44781161013a57806378f4b2f21161011457806378f4b2f21461029a5780637a167a8a146102a45780637aa63a86146102cb57600080fd5b8063461a44781461026a5780635ae6256d1461027d5780636fee07e01461028557600080fd5b8063299ca4781161016b578063299ca478146101e85780632a7f18be14610208578063378997701461024c57600080fd5b8063016f906f146101875780630b3dfa97146101d1575b600080fd5b6005546101a79073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101da60035481565b6040519081526020016101c8565b6000546101a79073ffffffffffffffffffffffffffffffffffffffff1681565b61021b6102163660046117b7565b610345565b604080518251815260208084015164ffffffffff9081169183019190915292820151909216908201526060016101c8565b6102546103c3565b60405164ffffffffff90911681526020016101c8565b6101a7610278366004611893565b6103d7565b610254610484565b610298610293366004611909565b610498565b005b6101da620186a081565b60055474010000000000000000000000000000000000000000900464ffffffffff16610254565b6101da610892565b6101da61c35081565b6101da60045481565b600654610254565b6101da60025481565b6101a76108ad565b6102986108d5565b6101da610db9565b6101da60015481565b610298610325366004611976565b610e40565b6102986103383660046119b2565b610fbd565b6102546113b3565b60408051606081018252600080825260208201819052918101919091526006828154811061037557610375611a3b565b6000918252602091829020604080516060810182526002909302909101805483526001015464ffffffffff808216948401949094526501000000000090049092169181019190915292915050565b6000806103ce6113e6565b50949350505050565b600080546040517fbf40fac100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063bf40fac19061042e908590600401611ad5565b60206040518083038186803b15801561044657600080fd5b505afa15801561045a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047e9190611aef565b92915050565b60008061048f6113e6565b95945050505050565b61c350815111156105165760405162461bcd60e51b815260206004820152603d60248201527f5472616e73616374696f6e20646174612073697a652065786365656473206d6160448201527f78696d756d20666f7220726f6c6c7570207472616e73616374696f6e2e00000060648201526084015b60405180910390fd5b60045482111561058e5760405162461bcd60e51b815260206004820152603d60248201527f5472616e73616374696f6e20676173206c696d69742065786365656473206d6160448201527f78696d756d20666f7220726f6c6c7570207472616e73616374696f6e2e000000606482015260840161050d565b620186a08210156106075760405162461bcd60e51b815260206004820152602960248201527f5472616e73616374696f6e20676173206c696d697420746f6f206c6f7720746f60448201527f20656e71756575652e0000000000000000000000000000000000000000000000606482015260840161050d565b6003548211156106d5576000600254600354846106249190611b3b565b61062e9190611b52565b905060005a90508181116106aa5760405162461bcd60e51b815260206004820152602b60248201527f496e73756666696369656e742067617320666f72204c322072617465206c696d60448201527f6974696e67206275726e2e000000000000000000000000000000000000000000606482015260840161050d565b60005b825a6106b99084611b3b565b10156106d157806106c981611b8d565b9150506106ad565b5050505b6000333214156106e65750336106ff565b5033731111000000000000000000000000000000001111015b6000818585856040516020016107189493929190611bc6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152828252805160209182012060608401835280845264ffffffffff42811692850192835243811693850193845260068054600181810183556000838152975160029092027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f81019290925594517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d4090910180549651841665010000000000027fffffffffffffffffffffffffffffffffffffffffffff000000000000000000009097169190931617949094179055915491935061081e91611b3b565b9050808673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f4b388aecf9fa6cc92253704e5975a6129a4f735bdbd99567df4ed0094ee4ceb588884260405161088293929190611c0f565b60405180910390a4505050505050565b60008061089d6113e6565b50505064ffffffffff1692915050565b60006108d0604051806060016040528060218152602001611ddb602191396103d7565b905090565b60043560d81c60093560e890811c90600c35901c6108f1610892565b8364ffffffffff161461096c5760405162461bcd60e51b815260206004820152603d60248201527f41637475616c20626174636820737461727420696e64657820646f6573206e6f60448201527f74206d6174636820657870656374656420737461727420696e6465782e000000606482015260840161050d565b6109aa6040518060400160405280600d81526020017f42564d5f53657175656e636572000000000000000000000000000000000000008152506103d7565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a4a5760405162461bcd60e51b815260206004820152602d60248201527f46756e6374696f6e2063616e206f6e6c792062652063616c6c6564206279207460448201527f68652053657175656e6365722e00000000000000000000000000000000000000606482015260840161050d565b6000610a5c62ffffff83166010611c38565b610a6790600f611c75565b905064ffffffffff8116361015610ae65760405162461bcd60e51b815260206004820152602260248201527f4e6f7420656e6f756768204261746368436f6e74657874732070726f7669646560448201527f642e000000000000000000000000000000000000000000000000000000000000606482015260840161050d565b6005546040805160808101825260008082526020820181905291810182905260608101829052909174010000000000000000000000000000000000000000900464ffffffffff169060005b8562ffffff168163ffffffff161015610b8f576000610b558263ffffffff166114a1565b8051909350839150610b679086611c8d565b9450826020015184610b799190611cb5565b9350508080610b8790611cd5565b915050610b31565b5060065464ffffffffff83161115610c355760405162461bcd60e51b815260206004820152604260248201527f417474656d7074656420746f20617070656e64206d6f726520656c656d656e7460448201527f73207468616e2061726520617661696c61626c6520696e20746865207175657560648201527f652e000000000000000000000000000000000000000000000000000000000000608482015260a40161050d565b6000610c468462ffffff8916611cf9565b63ffffffff169050600080836020015160001415610c6f57505060408201516060830151610ce0565b60006006610c7e600188611d1e565b64ffffffffff1681548110610c9557610c95611a3b565b6000918252602091829020604080516060810182526002909302909101805483526001015464ffffffffff808216948401859052650100000000009091041691018190529093509150505b610d04610cee600143611b3b565b408a62ffffff168564ffffffffff168585611528565b7f602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f899610d2f8487611d1e565b84610d38610892565b6040805164ffffffffff94851681529390921660208401529082015260600160405180910390a150506005805464ffffffffff90941674010000000000000000000000000000000000000000027fffffffffffffff0000000000ffffffffffffffffffffffffffffffffffffffff9094169390931790925550505050505050565b6000610dc36108ad565b73ffffffffffffffffffffffffffffffffffffffff16631f7b6d326040518163ffffffff1660e01b815260040160206040518083038186803b158015610e0857600080fd5b505afa158015610e1c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d09190611d3c565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610ea657600080fd5b505afa158015610eba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ede9190611aef565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f585760405162461bcd60e51b815260206004820181905260248201527f4f6e6c792063616c6c61626c6520627920746865204275726e2041646d696e2e604482015260640161050d565b60018190556002829055610f6c8183611c38565b60038190556002546001546040805192835260208301919091528101919091527fc6ed75e96b8b18b71edc1a6e82a9d677f8268c774a262c624eeb2cf0a8b3e07e9060600160405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561102357600080fd5b505afa158015611037573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061105b9190611aef565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146110d55760405162461bcd60e51b815260206004820181905260248201527f4f6e6c792063616c6c61626c6520627920746865206d7420626174636865722e604482015260640161050d565b85856110df610892565b8264ffffffffff161461115a5760405162461bcd60e51b815260206004820152603d60248201527f41637475616c20626174636820737461727420696e64657820646f6573206e6f60448201527f74206d6174636820657870656374656420737461727420696e6465782e000000606482015260840161050d565b600554869060009061118c90889074010000000000000000000000000000000000000000900464ffffffffff16611cb5565b60065490915064ffffffffff821611156112345760405162461bcd60e51b815260206004820152604260248201527f417474656d7074656420746f20617070656e64206d6f726520656c656d656e7460448201527f73207468616e2061726520617661696c61626c6520696e20746865207175657560648201527f652e000000000000000000000000000000000000000000000000000000000000608482015260a40161050d565b60006112458362ffffff8616611cf9565b63ffffffff16905060008064ffffffffff8a166112665750879050866112d7565b60006006611275600187611d1e565b64ffffffffff168154811061128c5761128c611a3b565b6000918252602091829020604080516060810182526002909302909101805483526001015464ffffffffff808216948401859052650100000000009091041691018190529093509150505b6112fb6112e5600143611b3b565b408762ffffff168564ffffffffff168585611528565b7f602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f8996113268486611d1e565b8461132f610892565b6040805164ffffffffff94851681529390921660208401529082015260600160405180910390a150506005805464ffffffffff90931674010000000000000000000000000000000000000000027fffffffffffffff0000000000ffffffffffffffffffffffffffffffffffffffff9093169290921790915550505050505050505050565b6005546006546000916108d0917401000000000000000000000000000000000000000090910464ffffffffff1690611d1e565b60008060008060006113f66108ad565b73ffffffffffffffffffffffffffffffffffffffff1663ccf8f9696040518163ffffffff1660e01b815260040160206040518083038186803b15801561143b57600080fd5b505afa15801561144f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114739190611d55565b64ffffffffff602882901c811697605083901c82169750607883901c8216965060a09290921c169350915050565b6114cc6040518060800160405280600081526020016000815260200160008152602001600081525090565b60006114d9601084611c38565b6114e490600f611c75565b60408051608081018252823560e890811c82526003840135901c6020820152600683013560d890811c92820192909252600b90920135901c60608201529392505050565b60006115326108ad565b905060008061153f6113e6565b50509150915060006040518060c001604052808573ffffffffffffffffffffffffffffffffffffffff16631f7b6d326040518163ffffffff1660e01b815260040160206040518083038186803b15801561159857600080fd5b505afa1580156115ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d09190611d3c565b81526020018a81526020018981526020018464ffffffffff16815260200160405180602001604052806000815250815260200160405180602001604052806000815250815250905080600001517fa47512905cf577d4cfae2efc3df461008ddb7234e91ce7f4eefcdb51e1077ccf82602001518360400151846060015185608001518660a00151604051611668959493929190611d97565b60405180910390a2600061167b8261176e565b905060006116b68360400151866116929190611cb5565b61169c8b87611cb5565b602890811b9190911760508b901b1760788a901b17901b90565b6040517f2015276c000000000000000000000000000000000000000000000000000000008152600481018490527fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000008216602482015290915073ffffffffffffffffffffffffffffffffffffffff871690632015276c90604401600060405180830381600087803b15801561174957600080fd5b505af115801561175d573d6000803e3d6000fd5b505050505050505050505050505050565b6020808201516040808401516060850151608086015160a0870151935160009661179a96959101611d97565b604051602081830303815290604052805190602001209050919050565b6000602082840312156117c957600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600067ffffffffffffffff8084111561181a5761181a6117d0565b604051601f85017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611860576118606117d0565b8160405280935085815286868601111561187957600080fd5b858560208301376000602087830101525050509392505050565b6000602082840312156118a557600080fd5b813567ffffffffffffffff8111156118bc57600080fd5b8201601f810184136118cd57600080fd5b6118dc848235602084016117ff565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461190657600080fd5b50565b60008060006060848603121561191e57600080fd5b8335611929816118e4565b925060208401359150604084013567ffffffffffffffff81111561194c57600080fd5b8401601f8101861361195d57600080fd5b61196c868235602084016117ff565b9150509250925092565b6000806040838503121561198957600080fd5b50508035926020909101359150565b803564ffffffffff811681146119ad57600080fd5b919050565b60008060008060008060c087890312156119cb57600080fd5b6119d487611998565b9550602087013562ffffff811681146119ec57600080fd5b9450604087013563ffffffff81168114611a0557600080fd5b9350611a1360608801611998565b9250611a2160808801611998565b9150611a2f60a08801611998565b90509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000815180845260005b81811015611a9057602081850181015186830182015201611a74565b81811115611aa2576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611ae86020830184611a6a565b9392505050565b600060208284031215611b0157600080fd5b8151611ae8816118e4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611b4d57611b4d611b0c565b500390565b600082611b88577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611bbf57611bbf611b0c565b5060010190565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525083604083015260806060830152611c056080830184611a6a565b9695505050505050565b838152606060208201526000611c286060830185611a6a565b9050826040830152949350505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611c7057611c70611b0c565b500290565b60008219821115611c8857611c88611b0c565b500190565b600063ffffffff808316818516808303821115611cac57611cac611b0c565b01949350505050565b600064ffffffffff808316818516808303821115611cac57611cac611b0c565b600063ffffffff80831681811415611cef57611cef611b0c565b6001019392505050565b600063ffffffff83811690831681811015611d1657611d16611b0c565b039392505050565b600064ffffffffff83811690831681811015611d1657611d16611b0c565b600060208284031215611d4e57600080fd5b5051919050565b600060208284031215611d6757600080fd5b81517fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000081168114611ae857600080fd5b85815284602082015283604082015260a060608201526000611dbc60a0830185611a6a565b8281036080840152611dce8185611a6a565b9897505050505050505056fe436861696e53746f72616765436f6e7461696e65722d4354432d62617463686573a2646970667358221220e9d0405c248be67b1a2c789eeb753a0a163e82efe3c819a462f405db5891e94d64736f6c63430008090033",
}

// CanonicalTransactionChainABI is the input ABI used to generate the binding from.
// Deprecated: Use CanonicalTransactionChainMetaData.ABI instead.
var CanonicalTransactionChainABI = CanonicalTransactionChainMetaData.ABI

// CanonicalTransactionChainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CanonicalTransactionChainMetaData.Bin instead.
var CanonicalTransactionChainBin = CanonicalTransactionChainMetaData.Bin

// DeployCanonicalTransactionChain deploys a new Ethereum contract, binding an instance of CanonicalTransactionChain to it.
func DeployCanonicalTransactionChain(auth *bind.TransactOpts, backend bind.ContractBackend, _libAddressManager common.Address, _maxTransactionGasLimit *big.Int, _l2GasDiscountDivisor *big.Int, _enqueueGasCost *big.Int, _mt_batcher common.Address) (common.Address, *types.Transaction, *CanonicalTransactionChain, error) {
	parsed, err := CanonicalTransactionChainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CanonicalTransactionChainBin), backend, _libAddressManager, _maxTransactionGasLimit, _l2GasDiscountDivisor, _enqueueGasCost, _mt_batcher)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CanonicalTransactionChain{CanonicalTransactionChainCaller: CanonicalTransactionChainCaller{contract: contract}, CanonicalTransactionChainTransactor: CanonicalTransactionChainTransactor{contract: contract}, CanonicalTransactionChainFilterer: CanonicalTransactionChainFilterer{contract: contract}}, nil
}

// CanonicalTransactionChain is an auto generated Go binding around an Ethereum contract.
type CanonicalTransactionChain struct {
	CanonicalTransactionChainCaller     // Read-only binding to the contract
	CanonicalTransactionChainTransactor // Write-only binding to the contract
	CanonicalTransactionChainFilterer   // Log filterer for contract events
}

// CanonicalTransactionChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type CanonicalTransactionChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalTransactionChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CanonicalTransactionChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalTransactionChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CanonicalTransactionChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalTransactionChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CanonicalTransactionChainSession struct {
	Contract     *CanonicalTransactionChain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CanonicalTransactionChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CanonicalTransactionChainCallerSession struct {
	Contract *CanonicalTransactionChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// CanonicalTransactionChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CanonicalTransactionChainTransactorSession struct {
	Contract     *CanonicalTransactionChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// CanonicalTransactionChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type CanonicalTransactionChainRaw struct {
	Contract *CanonicalTransactionChain // Generic contract binding to access the raw methods on
}

// CanonicalTransactionChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CanonicalTransactionChainCallerRaw struct {
	Contract *CanonicalTransactionChainCaller // Generic read-only contract binding to access the raw methods on
}

// CanonicalTransactionChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CanonicalTransactionChainTransactorRaw struct {
	Contract *CanonicalTransactionChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCanonicalTransactionChain creates a new instance of CanonicalTransactionChain, bound to a specific deployed contract.
func NewCanonicalTransactionChain(address common.Address, backend bind.ContractBackend) (*CanonicalTransactionChain, error) {
	contract, err := bindCanonicalTransactionChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CanonicalTransactionChain{CanonicalTransactionChainCaller: CanonicalTransactionChainCaller{contract: contract}, CanonicalTransactionChainTransactor: CanonicalTransactionChainTransactor{contract: contract}, CanonicalTransactionChainFilterer: CanonicalTransactionChainFilterer{contract: contract}}, nil
}

// NewCanonicalTransactionChainCaller creates a new read-only instance of CanonicalTransactionChain, bound to a specific deployed contract.
func NewCanonicalTransactionChainCaller(address common.Address, caller bind.ContractCaller) (*CanonicalTransactionChainCaller, error) {
	contract, err := bindCanonicalTransactionChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CanonicalTransactionChainCaller{contract: contract}, nil
}

// NewCanonicalTransactionChainTransactor creates a new write-only instance of CanonicalTransactionChain, bound to a specific deployed contract.
func NewCanonicalTransactionChainTransactor(address common.Address, transactor bind.ContractTransactor) (*CanonicalTransactionChainTransactor, error) {
	contract, err := bindCanonicalTransactionChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CanonicalTransactionChainTransactor{contract: contract}, nil
}

// NewCanonicalTransactionChainFilterer creates a new log filterer instance of CanonicalTransactionChain, bound to a specific deployed contract.
func NewCanonicalTransactionChainFilterer(address common.Address, filterer bind.ContractFilterer) (*CanonicalTransactionChainFilterer, error) {
	contract, err := bindCanonicalTransactionChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CanonicalTransactionChainFilterer{contract: contract}, nil
}

// bindCanonicalTransactionChain binds a generic wrapper to an already deployed contract.
func bindCanonicalTransactionChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CanonicalTransactionChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanonicalTransactionChain *CanonicalTransactionChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CanonicalTransactionChain.Contract.CanonicalTransactionChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanonicalTransactionChain *CanonicalTransactionChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.CanonicalTransactionChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanonicalTransactionChain *CanonicalTransactionChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.CanonicalTransactionChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CanonicalTransactionChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanonicalTransactionChain *CanonicalTransactionChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanonicalTransactionChain *CanonicalTransactionChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.contract.Transact(opts, method, params...)
}

// MAXROLLUPTXSIZE is a free data retrieval call binding the contract method 0x876ed5cb.
//
// Solidity: function MAX_ROLLUP_TX_SIZE() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) MAXROLLUPTXSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "MAX_ROLLUP_TX_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXROLLUPTXSIZE is a free data retrieval call binding the contract method 0x876ed5cb.
//
// Solidity: function MAX_ROLLUP_TX_SIZE() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) MAXROLLUPTXSIZE() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.MAXROLLUPTXSIZE(&_CanonicalTransactionChain.CallOpts)
}

// MAXROLLUPTXSIZE is a free data retrieval call binding the contract method 0x876ed5cb.
//
// Solidity: function MAX_ROLLUP_TX_SIZE() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) MAXROLLUPTXSIZE() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.MAXROLLUPTXSIZE(&_CanonicalTransactionChain.CallOpts)
}

// MINROLLUPTXGAS is a free data retrieval call binding the contract method 0x78f4b2f2.
//
// Solidity: function MIN_ROLLUP_TX_GAS() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) MINROLLUPTXGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "MIN_ROLLUP_TX_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINROLLUPTXGAS is a free data retrieval call binding the contract method 0x78f4b2f2.
//
// Solidity: function MIN_ROLLUP_TX_GAS() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) MINROLLUPTXGAS() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.MINROLLUPTXGAS(&_CanonicalTransactionChain.CallOpts)
}

// MINROLLUPTXGAS is a free data retrieval call binding the contract method 0x78f4b2f2.
//
// Solidity: function MIN_ROLLUP_TX_GAS() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) MINROLLUPTXGAS() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.MINROLLUPTXGAS(&_CanonicalTransactionChain.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) Batches(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "batches")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) Batches() (common.Address, error) {
	return _CanonicalTransactionChain.Contract.Batches(&_CanonicalTransactionChain.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) Batches() (common.Address, error) {
	return _CanonicalTransactionChain.Contract.Batches(&_CanonicalTransactionChain.CallOpts)
}

// EnqueueGasCost is a free data retrieval call binding the contract method 0xe654b1fb.
//
// Solidity: function enqueueGasCost() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) EnqueueGasCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "enqueueGasCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnqueueGasCost is a free data retrieval call binding the contract method 0xe654b1fb.
//
// Solidity: function enqueueGasCost() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) EnqueueGasCost() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.EnqueueGasCost(&_CanonicalTransactionChain.CallOpts)
}

// EnqueueGasCost is a free data retrieval call binding the contract method 0xe654b1fb.
//
// Solidity: function enqueueGasCost() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) EnqueueGasCost() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.EnqueueGasCost(&_CanonicalTransactionChain.CallOpts)
}

// EnqueueL2GasPrepaid is a free data retrieval call binding the contract method 0x0b3dfa97.
//
// Solidity: function enqueueL2GasPrepaid() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) EnqueueL2GasPrepaid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "enqueueL2GasPrepaid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnqueueL2GasPrepaid is a free data retrieval call binding the contract method 0x0b3dfa97.
//
// Solidity: function enqueueL2GasPrepaid() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) EnqueueL2GasPrepaid() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.EnqueueL2GasPrepaid(&_CanonicalTransactionChain.CallOpts)
}

// EnqueueL2GasPrepaid is a free data retrieval call binding the contract method 0x0b3dfa97.
//
// Solidity: function enqueueL2GasPrepaid() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) EnqueueL2GasPrepaid() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.EnqueueL2GasPrepaid(&_CanonicalTransactionChain.CallOpts)
}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) GetLastBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "getLastBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) GetLastBlockNumber() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetLastBlockNumber(&_CanonicalTransactionChain.CallOpts)
}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) GetLastBlockNumber() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetLastBlockNumber(&_CanonicalTransactionChain.CallOpts)
}

// GetLastTimestamp is a free data retrieval call binding the contract method 0x37899770.
//
// Solidity: function getLastTimestamp() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) GetLastTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "getLastTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastTimestamp is a free data retrieval call binding the contract method 0x37899770.
//
// Solidity: function getLastTimestamp() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) GetLastTimestamp() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetLastTimestamp(&_CanonicalTransactionChain.CallOpts)
}

// GetLastTimestamp is a free data retrieval call binding the contract method 0x37899770.
//
// Solidity: function getLastTimestamp() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) GetLastTimestamp() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetLastTimestamp(&_CanonicalTransactionChain.CallOpts)
}

// GetNextQueueIndex is a free data retrieval call binding the contract method 0x7a167a8a.
//
// Solidity: function getNextQueueIndex() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) GetNextQueueIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "getNextQueueIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextQueueIndex is a free data retrieval call binding the contract method 0x7a167a8a.
//
// Solidity: function getNextQueueIndex() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) GetNextQueueIndex() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetNextQueueIndex(&_CanonicalTransactionChain.CallOpts)
}

// GetNextQueueIndex is a free data retrieval call binding the contract method 0x7a167a8a.
//
// Solidity: function getNextQueueIndex() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) GetNextQueueIndex() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetNextQueueIndex(&_CanonicalTransactionChain.CallOpts)
}

// GetNumPendingQueueElements is a free data retrieval call binding the contract method 0xf722b41a.
//
// Solidity: function getNumPendingQueueElements() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) GetNumPendingQueueElements(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "getNumPendingQueueElements")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumPendingQueueElements is a free data retrieval call binding the contract method 0xf722b41a.
//
// Solidity: function getNumPendingQueueElements() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) GetNumPendingQueueElements() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetNumPendingQueueElements(&_CanonicalTransactionChain.CallOpts)
}

// GetNumPendingQueueElements is a free data retrieval call binding the contract method 0xf722b41a.
//
// Solidity: function getNumPendingQueueElements() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) GetNumPendingQueueElements() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetNumPendingQueueElements(&_CanonicalTransactionChain.CallOpts)
}

// GetQueueElement is a free data retrieval call binding the contract method 0x2a7f18be.
//
// Solidity: function getQueueElement(uint256 _index) view returns((bytes32,uint40,uint40) _element)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) GetQueueElement(opts *bind.CallOpts, _index *big.Int) (LibBVMCodecQueueElement, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "getQueueElement", _index)

	if err != nil {
		return *new(LibBVMCodecQueueElement), err
	}

	out0 := *abi.ConvertType(out[0], new(LibBVMCodecQueueElement)).(*LibBVMCodecQueueElement)

	return out0, err

}

// GetQueueElement is a free data retrieval call binding the contract method 0x2a7f18be.
//
// Solidity: function getQueueElement(uint256 _index) view returns((bytes32,uint40,uint40) _element)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) GetQueueElement(_index *big.Int) (LibBVMCodecQueueElement, error) {
	return _CanonicalTransactionChain.Contract.GetQueueElement(&_CanonicalTransactionChain.CallOpts, _index)
}

// GetQueueElement is a free data retrieval call binding the contract method 0x2a7f18be.
//
// Solidity: function getQueueElement(uint256 _index) view returns((bytes32,uint40,uint40) _element)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) GetQueueElement(_index *big.Int) (LibBVMCodecQueueElement, error) {
	return _CanonicalTransactionChain.Contract.GetQueueElement(&_CanonicalTransactionChain.CallOpts, _index)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) GetQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "getQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) GetQueueLength() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetQueueLength(&_CanonicalTransactionChain.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint40)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) GetQueueLength() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetQueueLength(&_CanonicalTransactionChain.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) GetTotalBatches(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "getTotalBatches")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) GetTotalBatches() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetTotalBatches(&_CanonicalTransactionChain.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) GetTotalBatches() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetTotalBatches(&_CanonicalTransactionChain.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) GetTotalElements(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "getTotalElements")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) GetTotalElements() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetTotalElements(&_CanonicalTransactionChain.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) GetTotalElements() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.GetTotalElements(&_CanonicalTransactionChain.CallOpts)
}

// L2GasDiscountDivisor is a free data retrieval call binding the contract method 0xccf987c8.
//
// Solidity: function l2GasDiscountDivisor() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) L2GasDiscountDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "l2GasDiscountDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2GasDiscountDivisor is a free data retrieval call binding the contract method 0xccf987c8.
//
// Solidity: function l2GasDiscountDivisor() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) L2GasDiscountDivisor() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.L2GasDiscountDivisor(&_CanonicalTransactionChain.CallOpts)
}

// L2GasDiscountDivisor is a free data retrieval call binding the contract method 0xccf987c8.
//
// Solidity: function l2GasDiscountDivisor() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) L2GasDiscountDivisor() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.L2GasDiscountDivisor(&_CanonicalTransactionChain.CallOpts)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) LibAddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "libAddressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) LibAddressManager() (common.Address, error) {
	return _CanonicalTransactionChain.Contract.LibAddressManager(&_CanonicalTransactionChain.CallOpts)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) LibAddressManager() (common.Address, error) {
	return _CanonicalTransactionChain.Contract.LibAddressManager(&_CanonicalTransactionChain.CallOpts)
}

// MaxTransactionGasLimit is a free data retrieval call binding the contract method 0x8d38c6c1.
//
// Solidity: function maxTransactionGasLimit() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) MaxTransactionGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "maxTransactionGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTransactionGasLimit is a free data retrieval call binding the contract method 0x8d38c6c1.
//
// Solidity: function maxTransactionGasLimit() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) MaxTransactionGasLimit() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.MaxTransactionGasLimit(&_CanonicalTransactionChain.CallOpts)
}

// MaxTransactionGasLimit is a free data retrieval call binding the contract method 0x8d38c6c1.
//
// Solidity: function maxTransactionGasLimit() view returns(uint256)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) MaxTransactionGasLimit() (*big.Int, error) {
	return _CanonicalTransactionChain.Contract.MaxTransactionGasLimit(&_CanonicalTransactionChain.CallOpts)
}

// MtBatcher is a free data retrieval call binding the contract method 0x016f906f.
//
// Solidity: function mt_batcher() view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) MtBatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "mt_batcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MtBatcher is a free data retrieval call binding the contract method 0x016f906f.
//
// Solidity: function mt_batcher() view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) MtBatcher() (common.Address, error) {
	return _CanonicalTransactionChain.Contract.MtBatcher(&_CanonicalTransactionChain.CallOpts)
}

// MtBatcher is a free data retrieval call binding the contract method 0x016f906f.
//
// Solidity: function mt_batcher() view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) MtBatcher() (common.Address, error) {
	return _CanonicalTransactionChain.Contract.MtBatcher(&_CanonicalTransactionChain.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainCaller) Resolve(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _CanonicalTransactionChain.contract.Call(opts, &out, "resolve", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) Resolve(_name string) (common.Address, error) {
	return _CanonicalTransactionChain.Contract.Resolve(&_CanonicalTransactionChain.CallOpts, _name)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_CanonicalTransactionChain *CanonicalTransactionChainCallerSession) Resolve(_name string) (common.Address, error) {
	return _CanonicalTransactionChain.Contract.Resolve(&_CanonicalTransactionChain.CallOpts, _name)
}

// AppendDaSequencerBatch is a paid mutator transaction binding the contract method 0xf4f7911d.
//
// Solidity: function appendDaSequencerBatch(uint40 _shouldStartAtElement, uint24 _totalElementsToAppend, uint32 _numSequencedTransactions, uint40 _numSubsequentQueueTransactions, uint40 _timestamp, uint40 _blockNumber) returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainTransactor) AppendDaSequencerBatch(opts *bind.TransactOpts, _shouldStartAtElement *big.Int, _totalElementsToAppend *big.Int, _numSequencedTransactions uint32, _numSubsequentQueueTransactions *big.Int, _timestamp *big.Int, _blockNumber *big.Int) (*types.Transaction, error) {
	return _CanonicalTransactionChain.contract.Transact(opts, "appendDaSequencerBatch", _shouldStartAtElement, _totalElementsToAppend, _numSequencedTransactions, _numSubsequentQueueTransactions, _timestamp, _blockNumber)
}

// AppendDaSequencerBatch is a paid mutator transaction binding the contract method 0xf4f7911d.
//
// Solidity: function appendDaSequencerBatch(uint40 _shouldStartAtElement, uint24 _totalElementsToAppend, uint32 _numSequencedTransactions, uint40 _numSubsequentQueueTransactions, uint40 _timestamp, uint40 _blockNumber) returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) AppendDaSequencerBatch(_shouldStartAtElement *big.Int, _totalElementsToAppend *big.Int, _numSequencedTransactions uint32, _numSubsequentQueueTransactions *big.Int, _timestamp *big.Int, _blockNumber *big.Int) (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.AppendDaSequencerBatch(&_CanonicalTransactionChain.TransactOpts, _shouldStartAtElement, _totalElementsToAppend, _numSequencedTransactions, _numSubsequentQueueTransactions, _timestamp, _blockNumber)
}

// AppendDaSequencerBatch is a paid mutator transaction binding the contract method 0xf4f7911d.
//
// Solidity: function appendDaSequencerBatch(uint40 _shouldStartAtElement, uint24 _totalElementsToAppend, uint32 _numSequencedTransactions, uint40 _numSubsequentQueueTransactions, uint40 _timestamp, uint40 _blockNumber) returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainTransactorSession) AppendDaSequencerBatch(_shouldStartAtElement *big.Int, _totalElementsToAppend *big.Int, _numSequencedTransactions uint32, _numSubsequentQueueTransactions *big.Int, _timestamp *big.Int, _blockNumber *big.Int) (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.AppendDaSequencerBatch(&_CanonicalTransactionChain.TransactOpts, _shouldStartAtElement, _totalElementsToAppend, _numSequencedTransactions, _numSubsequentQueueTransactions, _timestamp, _blockNumber)
}

// AppendSequencerBatch is a paid mutator transaction binding the contract method 0xd0f89344.
//
// Solidity: function appendSequencerBatch() returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainTransactor) AppendSequencerBatch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalTransactionChain.contract.Transact(opts, "appendSequencerBatch")
}

// AppendSequencerBatch is a paid mutator transaction binding the contract method 0xd0f89344.
//
// Solidity: function appendSequencerBatch() returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) AppendSequencerBatch() (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.AppendSequencerBatch(&_CanonicalTransactionChain.TransactOpts)
}

// AppendSequencerBatch is a paid mutator transaction binding the contract method 0xd0f89344.
//
// Solidity: function appendSequencerBatch() returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainTransactorSession) AppendSequencerBatch() (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.AppendSequencerBatch(&_CanonicalTransactionChain.TransactOpts)
}

// Enqueue is a paid mutator transaction binding the contract method 0x6fee07e0.
//
// Solidity: function enqueue(address _target, uint256 _gasLimit, bytes _data) returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainTransactor) Enqueue(opts *bind.TransactOpts, _target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _CanonicalTransactionChain.contract.Transact(opts, "enqueue", _target, _gasLimit, _data)
}

// Enqueue is a paid mutator transaction binding the contract method 0x6fee07e0.
//
// Solidity: function enqueue(address _target, uint256 _gasLimit, bytes _data) returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) Enqueue(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.Enqueue(&_CanonicalTransactionChain.TransactOpts, _target, _gasLimit, _data)
}

// Enqueue is a paid mutator transaction binding the contract method 0x6fee07e0.
//
// Solidity: function enqueue(address _target, uint256 _gasLimit, bytes _data) returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainTransactorSession) Enqueue(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.Enqueue(&_CanonicalTransactionChain.TransactOpts, _target, _gasLimit, _data)
}

// SetGasParams is a paid mutator transaction binding the contract method 0xedcc4a45.
//
// Solidity: function setGasParams(uint256 _l2GasDiscountDivisor, uint256 _enqueueGasCost) returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainTransactor) SetGasParams(opts *bind.TransactOpts, _l2GasDiscountDivisor *big.Int, _enqueueGasCost *big.Int) (*types.Transaction, error) {
	return _CanonicalTransactionChain.contract.Transact(opts, "setGasParams", _l2GasDiscountDivisor, _enqueueGasCost)
}

// SetGasParams is a paid mutator transaction binding the contract method 0xedcc4a45.
//
// Solidity: function setGasParams(uint256 _l2GasDiscountDivisor, uint256 _enqueueGasCost) returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainSession) SetGasParams(_l2GasDiscountDivisor *big.Int, _enqueueGasCost *big.Int) (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.SetGasParams(&_CanonicalTransactionChain.TransactOpts, _l2GasDiscountDivisor, _enqueueGasCost)
}

// SetGasParams is a paid mutator transaction binding the contract method 0xedcc4a45.
//
// Solidity: function setGasParams(uint256 _l2GasDiscountDivisor, uint256 _enqueueGasCost) returns()
func (_CanonicalTransactionChain *CanonicalTransactionChainTransactorSession) SetGasParams(_l2GasDiscountDivisor *big.Int, _enqueueGasCost *big.Int) (*types.Transaction, error) {
	return _CanonicalTransactionChain.Contract.SetGasParams(&_CanonicalTransactionChain.TransactOpts, _l2GasDiscountDivisor, _enqueueGasCost)
}

// CanonicalTransactionChainL2GasParamsUpdatedIterator is returned from FilterL2GasParamsUpdated and is used to iterate over the raw logs and unpacked data for L2GasParamsUpdated events raised by the CanonicalTransactionChain contract.
type CanonicalTransactionChainL2GasParamsUpdatedIterator struct {
	Event *CanonicalTransactionChainL2GasParamsUpdated // Event containing the contract specifics and raw log

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
func (it *CanonicalTransactionChainL2GasParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalTransactionChainL2GasParamsUpdated)
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
		it.Event = new(CanonicalTransactionChainL2GasParamsUpdated)
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
func (it *CanonicalTransactionChainL2GasParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalTransactionChainL2GasParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalTransactionChainL2GasParamsUpdated represents a L2GasParamsUpdated event raised by the CanonicalTransactionChain contract.
type CanonicalTransactionChainL2GasParamsUpdated struct {
	L2GasDiscountDivisor *big.Int
	EnqueueGasCost       *big.Int
	EnqueueL2GasPrepaid  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterL2GasParamsUpdated is a free log retrieval operation binding the contract event 0xc6ed75e96b8b18b71edc1a6e82a9d677f8268c774a262c624eeb2cf0a8b3e07e.
//
// Solidity: event L2GasParamsUpdated(uint256 l2GasDiscountDivisor, uint256 enqueueGasCost, uint256 enqueueL2GasPrepaid)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) FilterL2GasParamsUpdated(opts *bind.FilterOpts) (*CanonicalTransactionChainL2GasParamsUpdatedIterator, error) {

	logs, sub, err := _CanonicalTransactionChain.contract.FilterLogs(opts, "L2GasParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &CanonicalTransactionChainL2GasParamsUpdatedIterator{contract: _CanonicalTransactionChain.contract, event: "L2GasParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchL2GasParamsUpdated is a free log subscription operation binding the contract event 0xc6ed75e96b8b18b71edc1a6e82a9d677f8268c774a262c624eeb2cf0a8b3e07e.
//
// Solidity: event L2GasParamsUpdated(uint256 l2GasDiscountDivisor, uint256 enqueueGasCost, uint256 enqueueL2GasPrepaid)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) WatchL2GasParamsUpdated(opts *bind.WatchOpts, sink chan<- *CanonicalTransactionChainL2GasParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _CanonicalTransactionChain.contract.WatchLogs(opts, "L2GasParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalTransactionChainL2GasParamsUpdated)
				if err := _CanonicalTransactionChain.contract.UnpackLog(event, "L2GasParamsUpdated", log); err != nil {
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

// ParseL2GasParamsUpdated is a log parse operation binding the contract event 0xc6ed75e96b8b18b71edc1a6e82a9d677f8268c774a262c624eeb2cf0a8b3e07e.
//
// Solidity: event L2GasParamsUpdated(uint256 l2GasDiscountDivisor, uint256 enqueueGasCost, uint256 enqueueL2GasPrepaid)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) ParseL2GasParamsUpdated(log types.Log) (*CanonicalTransactionChainL2GasParamsUpdated, error) {
	event := new(CanonicalTransactionChainL2GasParamsUpdated)
	if err := _CanonicalTransactionChain.contract.UnpackLog(event, "L2GasParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CanonicalTransactionChainQueueBatchAppendedIterator is returned from FilterQueueBatchAppended and is used to iterate over the raw logs and unpacked data for QueueBatchAppended events raised by the CanonicalTransactionChain contract.
type CanonicalTransactionChainQueueBatchAppendedIterator struct {
	Event *CanonicalTransactionChainQueueBatchAppended // Event containing the contract specifics and raw log

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
func (it *CanonicalTransactionChainQueueBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalTransactionChainQueueBatchAppended)
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
		it.Event = new(CanonicalTransactionChainQueueBatchAppended)
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
func (it *CanonicalTransactionChainQueueBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalTransactionChainQueueBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalTransactionChainQueueBatchAppended represents a QueueBatchAppended event raised by the CanonicalTransactionChain contract.
type CanonicalTransactionChainQueueBatchAppended struct {
	StartingQueueIndex *big.Int
	NumQueueElements   *big.Int
	TotalElements      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterQueueBatchAppended is a free log retrieval operation binding the contract event 0x64d7f508348c70dea42d5302a393987e4abc20e45954ab3f9d320207751956f0.
//
// Solidity: event QueueBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) FilterQueueBatchAppended(opts *bind.FilterOpts) (*CanonicalTransactionChainQueueBatchAppendedIterator, error) {

	logs, sub, err := _CanonicalTransactionChain.contract.FilterLogs(opts, "QueueBatchAppended")
	if err != nil {
		return nil, err
	}
	return &CanonicalTransactionChainQueueBatchAppendedIterator{contract: _CanonicalTransactionChain.contract, event: "QueueBatchAppended", logs: logs, sub: sub}, nil
}

// WatchQueueBatchAppended is a free log subscription operation binding the contract event 0x64d7f508348c70dea42d5302a393987e4abc20e45954ab3f9d320207751956f0.
//
// Solidity: event QueueBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) WatchQueueBatchAppended(opts *bind.WatchOpts, sink chan<- *CanonicalTransactionChainQueueBatchAppended) (event.Subscription, error) {

	logs, sub, err := _CanonicalTransactionChain.contract.WatchLogs(opts, "QueueBatchAppended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalTransactionChainQueueBatchAppended)
				if err := _CanonicalTransactionChain.contract.UnpackLog(event, "QueueBatchAppended", log); err != nil {
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

// ParseQueueBatchAppended is a log parse operation binding the contract event 0x64d7f508348c70dea42d5302a393987e4abc20e45954ab3f9d320207751956f0.
//
// Solidity: event QueueBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) ParseQueueBatchAppended(log types.Log) (*CanonicalTransactionChainQueueBatchAppended, error) {
	event := new(CanonicalTransactionChainQueueBatchAppended)
	if err := _CanonicalTransactionChain.contract.UnpackLog(event, "QueueBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CanonicalTransactionChainSequencerBatchAppendedIterator is returned from FilterSequencerBatchAppended and is used to iterate over the raw logs and unpacked data for SequencerBatchAppended events raised by the CanonicalTransactionChain contract.
type CanonicalTransactionChainSequencerBatchAppendedIterator struct {
	Event *CanonicalTransactionChainSequencerBatchAppended // Event containing the contract specifics and raw log

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
func (it *CanonicalTransactionChainSequencerBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalTransactionChainSequencerBatchAppended)
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
		it.Event = new(CanonicalTransactionChainSequencerBatchAppended)
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
func (it *CanonicalTransactionChainSequencerBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalTransactionChainSequencerBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalTransactionChainSequencerBatchAppended represents a SequencerBatchAppended event raised by the CanonicalTransactionChain contract.
type CanonicalTransactionChainSequencerBatchAppended struct {
	StartingQueueIndex *big.Int
	NumQueueElements   *big.Int
	TotalElements      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchAppended is a free log retrieval operation binding the contract event 0x602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f899.
//
// Solidity: event SequencerBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) FilterSequencerBatchAppended(opts *bind.FilterOpts) (*CanonicalTransactionChainSequencerBatchAppendedIterator, error) {

	logs, sub, err := _CanonicalTransactionChain.contract.FilterLogs(opts, "SequencerBatchAppended")
	if err != nil {
		return nil, err
	}
	return &CanonicalTransactionChainSequencerBatchAppendedIterator{contract: _CanonicalTransactionChain.contract, event: "SequencerBatchAppended", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchAppended is a free log subscription operation binding the contract event 0x602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f899.
//
// Solidity: event SequencerBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) WatchSequencerBatchAppended(opts *bind.WatchOpts, sink chan<- *CanonicalTransactionChainSequencerBatchAppended) (event.Subscription, error) {

	logs, sub, err := _CanonicalTransactionChain.contract.WatchLogs(opts, "SequencerBatchAppended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalTransactionChainSequencerBatchAppended)
				if err := _CanonicalTransactionChain.contract.UnpackLog(event, "SequencerBatchAppended", log); err != nil {
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

// ParseSequencerBatchAppended is a log parse operation binding the contract event 0x602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f899.
//
// Solidity: event SequencerBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) ParseSequencerBatchAppended(log types.Log) (*CanonicalTransactionChainSequencerBatchAppended, error) {
	event := new(CanonicalTransactionChainSequencerBatchAppended)
	if err := _CanonicalTransactionChain.contract.UnpackLog(event, "SequencerBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CanonicalTransactionChainTransactionBatchAppendedIterator is returned from FilterTransactionBatchAppended and is used to iterate over the raw logs and unpacked data for TransactionBatchAppended events raised by the CanonicalTransactionChain contract.
type CanonicalTransactionChainTransactionBatchAppendedIterator struct {
	Event *CanonicalTransactionChainTransactionBatchAppended // Event containing the contract specifics and raw log

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
func (it *CanonicalTransactionChainTransactionBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalTransactionChainTransactionBatchAppended)
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
		it.Event = new(CanonicalTransactionChainTransactionBatchAppended)
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
func (it *CanonicalTransactionChainTransactionBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalTransactionChainTransactionBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalTransactionChainTransactionBatchAppended represents a TransactionBatchAppended event raised by the CanonicalTransactionChain contract.
type CanonicalTransactionChainTransactionBatchAppended struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	Signature         []byte
	ExtraData         []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTransactionBatchAppended is a free log retrieval operation binding the contract event 0xa47512905cf577d4cfae2efc3df461008ddb7234e91ce7f4eefcdb51e1077ccf.
//
// Solidity: event TransactionBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _signature, bytes _extraData)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) FilterTransactionBatchAppended(opts *bind.FilterOpts, _batchIndex []*big.Int) (*CanonicalTransactionChainTransactionBatchAppendedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _CanonicalTransactionChain.contract.FilterLogs(opts, "TransactionBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &CanonicalTransactionChainTransactionBatchAppendedIterator{contract: _CanonicalTransactionChain.contract, event: "TransactionBatchAppended", logs: logs, sub: sub}, nil
}

// WatchTransactionBatchAppended is a free log subscription operation binding the contract event 0xa47512905cf577d4cfae2efc3df461008ddb7234e91ce7f4eefcdb51e1077ccf.
//
// Solidity: event TransactionBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _signature, bytes _extraData)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) WatchTransactionBatchAppended(opts *bind.WatchOpts, sink chan<- *CanonicalTransactionChainTransactionBatchAppended, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _CanonicalTransactionChain.contract.WatchLogs(opts, "TransactionBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalTransactionChainTransactionBatchAppended)
				if err := _CanonicalTransactionChain.contract.UnpackLog(event, "TransactionBatchAppended", log); err != nil {
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

// ParseTransactionBatchAppended is a log parse operation binding the contract event 0xa47512905cf577d4cfae2efc3df461008ddb7234e91ce7f4eefcdb51e1077ccf.
//
// Solidity: event TransactionBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _signature, bytes _extraData)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) ParseTransactionBatchAppended(log types.Log) (*CanonicalTransactionChainTransactionBatchAppended, error) {
	event := new(CanonicalTransactionChainTransactionBatchAppended)
	if err := _CanonicalTransactionChain.contract.UnpackLog(event, "TransactionBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CanonicalTransactionChainTransactionEnqueuedIterator is returned from FilterTransactionEnqueued and is used to iterate over the raw logs and unpacked data for TransactionEnqueued events raised by the CanonicalTransactionChain contract.
type CanonicalTransactionChainTransactionEnqueuedIterator struct {
	Event *CanonicalTransactionChainTransactionEnqueued // Event containing the contract specifics and raw log

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
func (it *CanonicalTransactionChainTransactionEnqueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalTransactionChainTransactionEnqueued)
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
		it.Event = new(CanonicalTransactionChainTransactionEnqueued)
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
func (it *CanonicalTransactionChainTransactionEnqueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalTransactionChainTransactionEnqueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalTransactionChainTransactionEnqueued represents a TransactionEnqueued event raised by the CanonicalTransactionChain contract.
type CanonicalTransactionChainTransactionEnqueued struct {
	L1TxOrigin common.Address
	Target     common.Address
	GasLimit   *big.Int
	Data       []byte
	QueueIndex *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionEnqueued is a free log retrieval operation binding the contract event 0x4b388aecf9fa6cc92253704e5975a6129a4f735bdbd99567df4ed0094ee4ceb5.
//
// Solidity: event TransactionEnqueued(address indexed _l1TxOrigin, address indexed _target, uint256 _gasLimit, bytes _data, uint256 indexed _queueIndex, uint256 _timestamp)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) FilterTransactionEnqueued(opts *bind.FilterOpts, _l1TxOrigin []common.Address, _target []common.Address, _queueIndex []*big.Int) (*CanonicalTransactionChainTransactionEnqueuedIterator, error) {

	var _l1TxOriginRule []interface{}
	for _, _l1TxOriginItem := range _l1TxOrigin {
		_l1TxOriginRule = append(_l1TxOriginRule, _l1TxOriginItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	var _queueIndexRule []interface{}
	for _, _queueIndexItem := range _queueIndex {
		_queueIndexRule = append(_queueIndexRule, _queueIndexItem)
	}

	logs, sub, err := _CanonicalTransactionChain.contract.FilterLogs(opts, "TransactionEnqueued", _l1TxOriginRule, _targetRule, _queueIndexRule)
	if err != nil {
		return nil, err
	}
	return &CanonicalTransactionChainTransactionEnqueuedIterator{contract: _CanonicalTransactionChain.contract, event: "TransactionEnqueued", logs: logs, sub: sub}, nil
}

// WatchTransactionEnqueued is a free log subscription operation binding the contract event 0x4b388aecf9fa6cc92253704e5975a6129a4f735bdbd99567df4ed0094ee4ceb5.
//
// Solidity: event TransactionEnqueued(address indexed _l1TxOrigin, address indexed _target, uint256 _gasLimit, bytes _data, uint256 indexed _queueIndex, uint256 _timestamp)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) WatchTransactionEnqueued(opts *bind.WatchOpts, sink chan<- *CanonicalTransactionChainTransactionEnqueued, _l1TxOrigin []common.Address, _target []common.Address, _queueIndex []*big.Int) (event.Subscription, error) {

	var _l1TxOriginRule []interface{}
	for _, _l1TxOriginItem := range _l1TxOrigin {
		_l1TxOriginRule = append(_l1TxOriginRule, _l1TxOriginItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	var _queueIndexRule []interface{}
	for _, _queueIndexItem := range _queueIndex {
		_queueIndexRule = append(_queueIndexRule, _queueIndexItem)
	}

	logs, sub, err := _CanonicalTransactionChain.contract.WatchLogs(opts, "TransactionEnqueued", _l1TxOriginRule, _targetRule, _queueIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalTransactionChainTransactionEnqueued)
				if err := _CanonicalTransactionChain.contract.UnpackLog(event, "TransactionEnqueued", log); err != nil {
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

// ParseTransactionEnqueued is a log parse operation binding the contract event 0x4b388aecf9fa6cc92253704e5975a6129a4f735bdbd99567df4ed0094ee4ceb5.
//
// Solidity: event TransactionEnqueued(address indexed _l1TxOrigin, address indexed _target, uint256 _gasLimit, bytes _data, uint256 indexed _queueIndex, uint256 _timestamp)
func (_CanonicalTransactionChain *CanonicalTransactionChainFilterer) ParseTransactionEnqueued(log types.Log) (*CanonicalTransactionChainTransactionEnqueued, error) {
	event := new(CanonicalTransactionChainTransactionEnqueued)
	if err := _CanonicalTransactionChain.contract.UnpackLog(event, "TransactionEnqueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

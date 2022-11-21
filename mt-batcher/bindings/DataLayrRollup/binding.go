// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractDataLayrRollup

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// DataLayrDisclosureLogicMultiRevealProof is an auto generated low-level Go binding around an user-defined struct.
type DataLayrDisclosureLogicMultiRevealProof struct {
	InterpolationPoly BN254G1Point
	RevealProof       BN254G1Point
	ZeroPoly          BN254G2Point
	ZeroPolyProof     []byte
}

// DataLayrRollupDisclosureProofs is an auto generated low-level Go binding around an user-defined struct.
type DataLayrRollupDisclosureProofs struct {
	Header               []byte
	FirstChunkNumber     uint32
	Polys                [][]byte
	MultiRevealProofs    []DataLayrDisclosureLogicMultiRevealProof
	PolyEquivalenceProof BN254G2Point
}

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

// ContractDataLayrRollupMetaData contains all meta data concerning the ContractDataLayrRollup contract.
var ContractDataLayrRollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_neededStake\",\"type\":\"uint256\"},{\"internalType\":\"contractIDataLayrServiceManager\",\"name\":\"_dlsm\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupStoreNumber\",\"type\":\"uint32\"}],\"name\":\"RollupStoreConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"}],\"name\":\"RollupStoreInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupStoreNumber\",\"type\":\"uint32\"}],\"name\":\"RollupStoreReverted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_STALE_MEASURE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FRAUD_STRING\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreSearchData\",\"name\":\"searchData\",\"type\":\"tuple\"}],\"name\":\"confirmData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"dataStoreIdToRollupStoreNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dlsm\",\"outputs\":[{\"internalType\":\"contractIDataLayrServiceManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fraudProofPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"neededStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"polys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"parse\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"provenString\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fraudulentStoreNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreSearchData\",\"name\":\"searchData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"firstChunkNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"polys\",\"type\":\"bytes[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"interpolationPoly\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"revealProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"zeroPoly\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"zeroPolyProof\",\"type\":\"bytes\"}],\"internalType\":\"structDataLayrDisclosureLogic.MultiRevealProof[]\",\"name\":\"multiRevealProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"polyEquivalenceProof\",\"type\":\"tuple\"}],\"internalType\":\"structDataLayrRollup.DisclosureProofs\",\"name\":\"disclosureProofs\",\"type\":\"tuple\"}],\"name\":\"proveFraud\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupStoreNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rollupStores\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"confirmAt\",\"type\":\"uint32\"},{\"internalType\":\"enumDataLayrRollup.RollupStoreStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerStatus\",\"outputs\":[{\"internalType\":\"enumDataLayrRollup.SequencerStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalOperatorsIndex\",\"type\":\"uint32\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040526201518060e0523480156200001957600080fd5b50604051620020aa380380620020aa8339810160408190526200003c91620000a3565b60008054610100600160a81b0319166101006001600160a01b0396871602179055918316608052600480546001600160a01b031916929093169190911790915560a052606460c052620000ff565b6001600160a01b0381168114620000a057600080fd5b50565b60008060008060808587031215620000ba57600080fd5b8451620000c7816200008a565b6020860151909450620000da816200008a565b604086015160608701519194509250620000f4816200008a565b939692955090935050565b60805160a05160c05160e051611f4762000163600039600081816102b6015261102d0152600081816101e801526111ed0152600081816101090152818161057d0152610d6f0152600081816101a9015281816105a30152610d9d0152611f476000f3fe608060405234801561001057600080fd5b50600436106100ff5760003560e01c80635e8b3f2d11610097578063b537c4c711610066578063b537c4c71461027e578063bcb8ec901461029e578063f2a8f124146102b1578063f83099ca146102d857600080fd5b80635e8b3f2d146101e35780637bd858791461020a5780637ef01d5e1461021d578063b0393a371461023057600080fd5b8063428bba09116100d3578063428bba091461019357806346b2eb9b1461019c57806351ed6a30146101a45780635c1bba38146101cb57600080fd5b80621526cc146101045780631f944c8f1461013e57806329c8beb11461015e5780633a4b66f114610189575b600080fd5b61012b7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b61015161014c36600461148f565b6102f2565b6040516101359190611541565b600454610171906001600160a01b031681565b6040516001600160a01b039091168152602001610135565b610191610488565b005b61012b60035481565b610151610628565b6101717f000000000000000000000000000000000000000000000000000000000000000081565b6000546101719061010090046001600160a01b031681565b61012b7f000000000000000000000000000000000000000000000000000000000000000081565b610191610218366004611710565b610644565b61019161022b3660046117c3565b610e4d565b61026f61023e366004611819565b60016020526000908152604090205463ffffffff80821691640100000000810490911690600160401b900460ff1683565b60405161013593929190611866565b61012b61028c366004611892565b60026020526000908152604090205481565b6101916102ac3660046118af565b611142565b61012b7f000000000000000000000000000000000000000000000000000000000000000081565b6000546102e59060ff1681565b604051610135919061192b565b60606102ff602084611954565b6000036103635760405162461bcd60e51b815260206004820152602760248201527f43616e6e6f742073746172742072656164696e672066726f6d206120706164646044820152666564206279746560c81b60648201526084015b60405180910390fd5b6000835b838351101561047e5760006103b28261038160208261197e565b61038c906001611992565b6103979060206119aa565b6103a191906119c9565b85516103ad90886119c9565b6113b8565b9050838888858181106103c7576103c76119e0565b90506020028101906103d991906119f6565b84906103e58583611992565b926103f293929190611a3d565b60405160200161040493929190611a67565b6040516020818303038152906040529350878784818110610427576104276119e0565b905060200281019061043991906119f6565b90506104458284611992565b03610460578261045481611a8f565b93505060019150610478565b61046b816001611992565b6104759083611992565b91505b50610367565b5050949350505050565b60005461010090046001600160a01b031633146104e75760405162461bcd60e51b815260206004820152601c60248201527f4f6e6c79207468652073657175656e6365722063616e207374616b6500000000604482015260640161035a565b6000805460ff1660028111156104ff576104ff611832565b146105625760405162461bcd60e51b815260206004820152602d60248201527f53657175656e6365722063616e206f6e6c79207374616b65206966207468657960448201526c08185c99481d5b9cdd185ad959609a1b606482015260840161035a565b6040516323b872dd60e01b81523360048201523060248201527f000000000000000000000000000000000000000000000000000000000000000060448201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af11580156105f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106189190611aa8565b506000805460ff19166001179055565b604051806080016040528060608152602001611eb26060913981565b600160005460ff16600281111561065d5761065d611832565b146106be5760405162461bcd60e51b815260206004820152602b60248201527f43616e206f6e6c792070726f76652066726175642069662073657175656e636560448201526a1c881a5cc81cdd185ad95960aa1b606482015260840161035a565b60008481526001602090815260408083208151606081018352815463ffffffff808216835264010000000082041694820194909452929091830190600160401b900460ff16600281111561071457610714611832565b600281111561072557610725611832565b905250905060018160400151600281111561074257610742611832565b148015610758575042816020015163ffffffff16115b6107ba5760405162461bcd60e51b815260206004820152602d60248201527f526f6c6c757053746f7265206d75737420626520636f6d6d697474656420616e60448201526c19081d5b98dbdb999a5c9b5959609a1b606482015260840161035a565b82516107c5906113d0565b600480546020860151604080880151606089015191516376c1607760e11b815260ff90931694830194909452602482019390935263ffffffff90921660448301526001600160a01b03169063ed82c0ee90606401602060405180830381865afa158015610836573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085a9190611aca565b146108a75760405162461bcd60e51b815260206004820152601e60248201527f6d6574616461746120707265696d61676520697320696e636f72726563740000604482015260640161035a565b805183516040015163ffffffff9081169116146109375760405162461bcd60e51b815260206004820152604260248201527f7365616368446174612773206461746173746f7265206964206973206e6f742060448201527f636f6e73697374656e74207769746820676976656e20726f6c6c75702073746f606482015261726560f01b608482015260a40161035a565b61094182806119f6565b60405161094f929190611ae3565b604051908190039020835151146109c35760405162461bcd60e51b815260206004820152603260248201527f646973636c6f737572652070726f6f66732068656164657268617368207072656044820152711a5b5859d9481a5cc81a5b98dbdc9c9958dd60721b606482015260840161035a565b73__$82829e2b3546386e49154dd085205fe43f$__63899c1afc6109e784806119f6565b6109f76040870160208801611892565b610a046040880188611af3565b610a1160608a018a611af3565b8a6080016040518963ffffffff1660e01b8152600401610a38989796959493929190611c81565b602060405180830381865af4158015610a55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a799190611aa8565b610ac55760405162461bcd60e51b815260206004820152601d60248201527f646973636c6f737572652070726f6f66732061726520696e76616c6964000000604482015260640161035a565b600073__$82829e2b3546386e49154dd085205fe43f$__63e34c39d6610aeb85806119f6565b6040518363ffffffff1660e01b8152600401610b08929190611d31565b602060405180830381865af4158015610b25573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b499190611d4d565b905063ffffffff8116610b5f6040850185611af3565b9050610b716040860160208701611892565b63ffffffff16610b819190611992565b1115610be65760405162461bcd60e51b815260206004820152602e60248201527f43616e206f6e6c792070726f766520646174612066726f6d207468652073797360448201526d74656d61746963206368756e6b7360901b606482015260840161035a565b6000610c18610bf86040860186611af3565b88604051806080016040528060608152602001611eb260609139516102f2565b9050604051806080016040528060608152602001611eb26060913951815114610cb45760405162461bcd60e51b815260206004820152604260248201527f50617273696e67206572726f722c2070726f76656e20737472696e672069732060448201527f646966666572656e74206c656e677468207468616e20667261756420737472696064820152616e6760f01b608482015260a40161035a565b604051806080016040528060608152602001611eb26060913980519060200120818051906020012014610d295760405162461bcd60e51b815260206004820152601d60248201527f70726f76656e20737472696e6720213d20667261756420737472696e67000000604482015260640161035a565b60008781526001602052604080822080546802000000000000000060ff60401b19909116179055815460ff19166002179091555163a9059cbb60e01b81523360048201527f000000000000000000000000000000000000000000000000000000000000000060248201526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af1158015610de6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e0a9190611aa8565b5060405163ffffffff881681527f18407f85390b98bfd30ac355138feffbf9d2519036bddba3429ec57ed328e7a89060200160405180910390a150505050505050565b60005461010090046001600160a01b03163314610e7c5760405162461bcd60e51b815260040161035a90611d6a565b600160005460ff166002811115610e9557610e95611832565b14610ef85760405162461bcd60e51b815260206004820152602d60248201527f53657175656e6365722063616e206f6e6c7920636f6e6669726d20696620746860448201526c195e48185c99481cdd185ad959609a1b606482015260840161035a565b805160409081015163ffffffff1660009081526002602052205460001914610fa65760405162461bcd60e51b815260206004820152605560248201527f446174612073746f72652065697468657220776173206e6f7420696e6974696160448201527f6c697a65642062792074686520726f6c6c757020636f6e74726163742c206f72606482015274081a5cc8185b1c9958591e4818dbdb999a5c9b5959605a1b608482015260a40161035a565b60048054604051635189951560e01b81526001600160a01b0390911691635189951591610fd99187918791879101611dab565b600060405180830381600087803b158015610ff357600080fd5b505af1158015611007573d6000803e3d6000fd5b50506040805160608101825284519091015163ffffffff168152915050602081016110527f000000000000000000000000000000000000000000000000000000000000000042611992565b63ffffffff16815260200160019052600354600090815260016020908152604091829020835181549285015163ffffffff9081166401000000000267ffffffffffffffff1990941691161791909117808255918301519091829060ff60401b1916600160401b8360028111156110ca576110ca611832565b02179055505060038054835160409081015163ffffffff16600090815260026020529081208290557fe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a9935090919061112183611a8f565b9091555060405163ffffffff909116815260200160405180910390a1505050565b60005461010090046001600160a01b031633146111715760405162461bcd60e51b815260040161035a90611d6a565b600160005460ff16600281111561118a5761118a611832565b146111eb5760405162461bcd60e51b815260206004820152602b60248201527f53657175656e6365722063616e206f6e6c792073746f7265206966207468657960448201526a08185c99481cdd185ad95960aa1b606482015260840161035a565b7f000000000000000000000000000000000000000000000000000000000000000061121c63ffffffff8416436119c9565b106112695760405162461bcd60e51b815260206004820152601e60248201527f7374616b65732074616b656e2066726f6d20746f6f206c6f6e672061676f0000604482015260640161035a565b6000600460009054906101000a90046001600160a01b03166001600160a01b03166372d18e8d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112e29190611d4d565b6004805460405163dcf49ea760e01b81529293506001600160a01b03169163dcf49ea79161131e91339130918a918a918a918f918f9101611e67565b6020604051808303816000875af115801561133d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113619190611d4d565b5063ffffffff8116600081815260026020908152604091829020600019905590519182527f957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b5910160405180910390a1505050505050565b60008183106113c757816113c9565b825b9392505050565b600080826000015183602001518460400151856060015186608001518760a001518860c00151604051602001611470979695949392919096875260e095861b6001600160e01b0319908116602089015294861b851660248801529290941b909216602885015260a09190911b6001600160a01b031916602c84015260609190911b6bffffffffffffffffffffffff19166038830152604c820152606c0190565b60408051601f1981840301815291905280516020909101209392505050565b600080600080606085870312156114a557600080fd5b843567ffffffffffffffff808211156114bd57600080fd5b818701915087601f8301126114d157600080fd5b8135818111156114e057600080fd5b8860208260051b85010111156114f557600080fd5b6020928301999098509187013596604001359550909350505050565b60005b8381101561152c578181015183820152602001611514565b8381111561153b576000848401525b50505050565b6020815260008251806020840152611560816040850160208701611511565b601f01601f19169190910160400192915050565b6040516080810167ffffffffffffffff811182821017156115a557634e487b7160e01b600052604160045260246000fd5b60405290565b60405160e0810167ffffffffffffffff811182821017156115a557634e487b7160e01b600052604160045260246000fd5b63ffffffff811681146115ee57600080fd5b50565b80356115fc816115dc565b919050565b80356001600160a01b03811681146115fc57600080fd5b803560ff811681146115fc57600080fd5b600081830361014081121561163d57600080fd5b611645611574565b915060e081121561165557600080fd5b5061165e6115ab565b823581526020830135611670816115dc565b60208201526040830135611683816115dc565b60408201526060830135611696816115dc565b606082015260808301356bffffffffffffffffffffffff811681146116ba57600080fd5b60808201526116cb60a08401611601565b60a082015260c0838101359082015281526116e860e08301611618565b6020820152610100820135604082015261170561012083016115f1565b606082015292915050565b6000806000806101a0858703121561172757600080fd5b843593506020850135925061173f8660408701611629565b915061018085013567ffffffffffffffff81111561175c57600080fd5b8501610100818803121561176f57600080fd5b939692955090935050565b60008083601f84011261178c57600080fd5b50813567ffffffffffffffff8111156117a457600080fd5b6020830191508360208285010111156117bc57600080fd5b9250929050565b600080600061016084860312156117d957600080fd5b833567ffffffffffffffff8111156117f057600080fd5b6117fc8682870161177a565b909450925061181090508560208601611629565b90509250925092565b60006020828403121561182b57600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b600381106115ee57634e487b7160e01b600052602160045260246000fd5b63ffffffff8481168252831660208201526060810161188483611848565b826040830152949350505050565b6000602082840312156118a457600080fd5b81356113c9816115dc565b6000806000806000608086880312156118c757600080fd5b853567ffffffffffffffff8111156118de57600080fd5b6118ea8882890161177a565b90965094506118fd905060208701611618565b9250604086013561190d816115dc565b9150606086013561191d816115dc565b809150509295509295909350565b6020810161193883611848565b91905290565b634e487b7160e01b600052601260045260246000fd5b6000826119635761196361193e565b500690565b634e487b7160e01b600052601160045260246000fd5b60008261198d5761198d61193e565b500490565b600082198211156119a5576119a5611968565b500190565b60008160001904831182151516156119c4576119c4611968565b500290565b6000828210156119db576119db611968565b500390565b634e487b7160e01b600052603260045260246000fd5b6000808335601e19843603018112611a0d57600080fd5b83018035915067ffffffffffffffff821115611a2857600080fd5b6020019150368190038213156117bc57600080fd5b60008085851115611a4d57600080fd5b83861115611a5a57600080fd5b5050820193919092039150565b60008451611a79818460208901611511565b8201838582376000930192835250909392505050565b600060018201611aa157611aa1611968565b5060010190565b600060208284031215611aba57600080fd5b815180151581146113c957600080fd5b600060208284031215611adc57600080fd5b5051919050565b8183823760009101908152919050565b6000808335601e19843603018112611b0a57600080fd5b83018035915067ffffffffffffffff821115611b2557600080fd5b6020019150600581901b36038213156117bc57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e19843603018112611b7d57600080fd5b830160208101925035905067ffffffffffffffff811115611b9d57600080fd5b8036038213156117bc57600080fd5b604081833760408082016040840137506000608082015250565b81835260006020808501808196508560051b81019150846000805b88811015611c73578385038a52823561011e19893603018112611c02578283fd5b8801803586526020808201359087015260408082013590870152606080820135908701526101206080611c39818901828501611bac565b50610100611c4981840184611b66565b935082828a0152611c5d838a018583611b3d565b9d8a019d98505050938701935050600101611be1565b509298975050505050505050565b6000610100808352611c968184018b8d611b3d565b9050602063ffffffff8a16818501528382036040850152818883528183019050818960051b8401018a60005b8b811015611cfc57858303601f19018452611cdd828e611b66565b611ce8858284611b3d565b958701959450505090840190600101611cc2565b50508581036060870152611d1181898b611bc6565b945050505050611d246080830184611bac565b9998505050505050505050565b602081526000611d45602083018486611b3d565b949350505050565b600060208284031215611d5f57600080fd5b81516113c9816115dc565b60208082526021908201527f4f6e6c79207468652073657175656e6365722063616e2073746f7265206461746040820152606160f81b606082015260800190565b6000610160808352611dc08184018688611b3d565b915050825180516020840152602081015163ffffffff808216604086015280604084015116606086015280606084015116608086015250506bffffffffffffffffffffffff60808201511660a084015260018060a01b0360a08201511660c084015260c081015160e0840152506020830151611e4261010084018260ff169052565b50604083015161012083015260609092015163ffffffff166101409091015292915050565b6001600160a01b0388811682528716602082015260ff8616604082015263ffffffff85811660608301528416608082015260c060a08201819052600090611d249083018486611b3d56fe2d5f2860204f2060295f2d202d5f2860206f2060295f2d202d5f286020512060295f2d2042495444414f204a5553542052454b5420594f55207c5f2860204f2060295f7c202d207c5f2860206f2060295f7c202d207c5f286020512060295f7ca2646970667358221220ea2503676fba62cc94d78ed3d38c761284f173163dd507215419602a47afcb4d64736f6c634300080f0033",
}

// ContractDataLayrRollupABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractDataLayrRollupMetaData.ABI instead.
var ContractDataLayrRollupABI = ContractDataLayrRollupMetaData.ABI

// ContractDataLayrRollupBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractDataLayrRollupMetaData.Bin instead.
var ContractDataLayrRollupBin = ContractDataLayrRollupMetaData.Bin

// DeployContractDataLayrRollup deploys a new Ethereum contract, binding an instance of ContractDataLayrRollup to it.
func DeployContractDataLayrRollup(auth *bind.TransactOpts, backend bind.ContractBackend, _sequencer common.Address, _stakeToken common.Address, _neededStake *big.Int, _dlsm common.Address) (common.Address, *types.Transaction, *ContractDataLayrRollup, error) {
	parsed, err := ContractDataLayrRollupMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractDataLayrRollupBin), backend, _sequencer, _stakeToken, _neededStake, _dlsm)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractDataLayrRollup{ContractDataLayrRollupCaller: ContractDataLayrRollupCaller{contract: contract}, ContractDataLayrRollupTransactor: ContractDataLayrRollupTransactor{contract: contract}, ContractDataLayrRollupFilterer: ContractDataLayrRollupFilterer{contract: contract}}, nil
}

// ContractDataLayrRollup is an auto generated Go binding around an Ethereum contract.
type ContractDataLayrRollup struct {
	ContractDataLayrRollupCaller     // Read-only binding to the contract
	ContractDataLayrRollupTransactor // Write-only binding to the contract
	ContractDataLayrRollupFilterer   // Log filterer for contract events
}

// ContractDataLayrRollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractDataLayrRollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDataLayrRollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractDataLayrRollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDataLayrRollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractDataLayrRollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDataLayrRollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractDataLayrRollupSession struct {
	Contract     *ContractDataLayrRollup // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractDataLayrRollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractDataLayrRollupCallerSession struct {
	Contract *ContractDataLayrRollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ContractDataLayrRollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractDataLayrRollupTransactorSession struct {
	Contract     *ContractDataLayrRollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ContractDataLayrRollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractDataLayrRollupRaw struct {
	Contract *ContractDataLayrRollup // Generic contract binding to access the raw methods on
}

// ContractDataLayrRollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractDataLayrRollupCallerRaw struct {
	Contract *ContractDataLayrRollupCaller // Generic read-only contract binding to access the raw methods on
}

// ContractDataLayrRollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractDataLayrRollupTransactorRaw struct {
	Contract *ContractDataLayrRollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractDataLayrRollup creates a new instance of ContractDataLayrRollup, bound to a specific deployed contract.
func NewContractDataLayrRollup(address common.Address, backend bind.ContractBackend) (*ContractDataLayrRollup, error) {
	contract, err := bindContractDataLayrRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractDataLayrRollup{ContractDataLayrRollupCaller: ContractDataLayrRollupCaller{contract: contract}, ContractDataLayrRollupTransactor: ContractDataLayrRollupTransactor{contract: contract}, ContractDataLayrRollupFilterer: ContractDataLayrRollupFilterer{contract: contract}}, nil
}

// NewContractDataLayrRollupCaller creates a new read-only instance of ContractDataLayrRollup, bound to a specific deployed contract.
func NewContractDataLayrRollupCaller(address common.Address, caller bind.ContractCaller) (*ContractDataLayrRollupCaller, error) {
	contract, err := bindContractDataLayrRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractDataLayrRollupCaller{contract: contract}, nil
}

// NewContractDataLayrRollupTransactor creates a new write-only instance of ContractDataLayrRollup, bound to a specific deployed contract.
func NewContractDataLayrRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractDataLayrRollupTransactor, error) {
	contract, err := bindContractDataLayrRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractDataLayrRollupTransactor{contract: contract}, nil
}

// NewContractDataLayrRollupFilterer creates a new log filterer instance of ContractDataLayrRollup, bound to a specific deployed contract.
func NewContractDataLayrRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractDataLayrRollupFilterer, error) {
	contract, err := bindContractDataLayrRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractDataLayrRollupFilterer{contract: contract}, nil
}

// bindContractDataLayrRollup binds a generic wrapper to an already deployed contract.
func bindContractDataLayrRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractDataLayrRollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDataLayrRollup *ContractDataLayrRollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractDataLayrRollup.Contract.ContractDataLayrRollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDataLayrRollup *ContractDataLayrRollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.ContractDataLayrRollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDataLayrRollup *ContractDataLayrRollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.ContractDataLayrRollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractDataLayrRollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDataLayrRollup *ContractDataLayrRollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDataLayrRollup *ContractDataLayrRollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.contract.Transact(opts, method, params...)
}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) BLOCKSTALEMEASURE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "BLOCK_STALE_MEASURE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) BLOCKSTALEMEASURE() (*big.Int, error) {
	return _ContractDataLayrRollup.Contract.BLOCKSTALEMEASURE(&_ContractDataLayrRollup.CallOpts)
}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) BLOCKSTALEMEASURE() (*big.Int, error) {
	return _ContractDataLayrRollup.Contract.BLOCKSTALEMEASURE(&_ContractDataLayrRollup.CallOpts)
}

// FRAUDSTRING is a free data retrieval call binding the contract method 0x46b2eb9b.
//
// Solidity: function FRAUD_STRING() view returns(bytes)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) FRAUDSTRING(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "FRAUD_STRING")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FRAUDSTRING is a free data retrieval call binding the contract method 0x46b2eb9b.
//
// Solidity: function FRAUD_STRING() view returns(bytes)
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) FRAUDSTRING() ([]byte, error) {
	return _ContractDataLayrRollup.Contract.FRAUDSTRING(&_ContractDataLayrRollup.CallOpts)
}

// FRAUDSTRING is a free data retrieval call binding the contract method 0x46b2eb9b.
//
// Solidity: function FRAUD_STRING() view returns(bytes)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) FRAUDSTRING() ([]byte, error) {
	return _ContractDataLayrRollup.Contract.FRAUDSTRING(&_ContractDataLayrRollup.CallOpts)
}

// DataStoreIdToRollupStoreNumber is a free data retrieval call binding the contract method 0xb537c4c7.
//
// Solidity: function dataStoreIdToRollupStoreNumber(uint32 ) view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) DataStoreIdToRollupStoreNumber(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "dataStoreIdToRollupStoreNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DataStoreIdToRollupStoreNumber is a free data retrieval call binding the contract method 0xb537c4c7.
//
// Solidity: function dataStoreIdToRollupStoreNumber(uint32 ) view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) DataStoreIdToRollupStoreNumber(arg0 uint32) (*big.Int, error) {
	return _ContractDataLayrRollup.Contract.DataStoreIdToRollupStoreNumber(&_ContractDataLayrRollup.CallOpts, arg0)
}

// DataStoreIdToRollupStoreNumber is a free data retrieval call binding the contract method 0xb537c4c7.
//
// Solidity: function dataStoreIdToRollupStoreNumber(uint32 ) view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) DataStoreIdToRollupStoreNumber(arg0 uint32) (*big.Int, error) {
	return _ContractDataLayrRollup.Contract.DataStoreIdToRollupStoreNumber(&_ContractDataLayrRollup.CallOpts, arg0)
}

// Dlsm is a free data retrieval call binding the contract method 0x29c8beb1.
//
// Solidity: function dlsm() view returns(address)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) Dlsm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "dlsm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dlsm is a free data retrieval call binding the contract method 0x29c8beb1.
//
// Solidity: function dlsm() view returns(address)
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) Dlsm() (common.Address, error) {
	return _ContractDataLayrRollup.Contract.Dlsm(&_ContractDataLayrRollup.CallOpts)
}

// Dlsm is a free data retrieval call binding the contract method 0x29c8beb1.
//
// Solidity: function dlsm() view returns(address)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) Dlsm() (common.Address, error) {
	return _ContractDataLayrRollup.Contract.Dlsm(&_ContractDataLayrRollup.CallOpts)
}

// FraudProofPeriod is a free data retrieval call binding the contract method 0xf2a8f124.
//
// Solidity: function fraudProofPeriod() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) FraudProofPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "fraudProofPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FraudProofPeriod is a free data retrieval call binding the contract method 0xf2a8f124.
//
// Solidity: function fraudProofPeriod() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) FraudProofPeriod() (*big.Int, error) {
	return _ContractDataLayrRollup.Contract.FraudProofPeriod(&_ContractDataLayrRollup.CallOpts)
}

// FraudProofPeriod is a free data retrieval call binding the contract method 0xf2a8f124.
//
// Solidity: function fraudProofPeriod() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) FraudProofPeriod() (*big.Int, error) {
	return _ContractDataLayrRollup.Contract.FraudProofPeriod(&_ContractDataLayrRollup.CallOpts)
}

// NeededStake is a free data retrieval call binding the contract method 0x001526cc.
//
// Solidity: function neededStake() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) NeededStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "neededStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NeededStake is a free data retrieval call binding the contract method 0x001526cc.
//
// Solidity: function neededStake() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) NeededStake() (*big.Int, error) {
	return _ContractDataLayrRollup.Contract.NeededStake(&_ContractDataLayrRollup.CallOpts)
}

// NeededStake is a free data retrieval call binding the contract method 0x001526cc.
//
// Solidity: function neededStake() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) NeededStake() (*big.Int, error) {
	return _ContractDataLayrRollup.Contract.NeededStake(&_ContractDataLayrRollup.CallOpts)
}

// Parse is a free data retrieval call binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) pure returns(bytes provenString)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) Parse(opts *bind.CallOpts, polys [][]byte, startIndex *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "parse", polys, startIndex, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Parse is a free data retrieval call binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) pure returns(bytes provenString)
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) Parse(polys [][]byte, startIndex *big.Int, length *big.Int) ([]byte, error) {
	return _ContractDataLayrRollup.Contract.Parse(&_ContractDataLayrRollup.CallOpts, polys, startIndex, length)
}

// Parse is a free data retrieval call binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) pure returns(bytes provenString)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) Parse(polys [][]byte, startIndex *big.Int, length *big.Int) ([]byte, error) {
	return _ContractDataLayrRollup.Contract.Parse(&_ContractDataLayrRollup.CallOpts, polys, startIndex, length)
}

// RollupStoreNumber is a free data retrieval call binding the contract method 0x428bba09.
//
// Solidity: function rollupStoreNumber() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) RollupStoreNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "rollupStoreNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RollupStoreNumber is a free data retrieval call binding the contract method 0x428bba09.
//
// Solidity: function rollupStoreNumber() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) RollupStoreNumber() (*big.Int, error) {
	return _ContractDataLayrRollup.Contract.RollupStoreNumber(&_ContractDataLayrRollup.CallOpts)
}

// RollupStoreNumber is a free data retrieval call binding the contract method 0x428bba09.
//
// Solidity: function rollupStoreNumber() view returns(uint256)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) RollupStoreNumber() (*big.Int, error) {
	return _ContractDataLayrRollup.Contract.RollupStoreNumber(&_ContractDataLayrRollup.CallOpts)
}

// RollupStores is a free data retrieval call binding the contract method 0xb0393a37.
//
// Solidity: function rollupStores(uint256 ) view returns(uint32 dataStoreId, uint32 confirmAt, uint8 status)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) RollupStores(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DataStoreId uint32
	ConfirmAt   uint32
	Status      uint8
}, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "rollupStores", arg0)

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
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) RollupStores(arg0 *big.Int) (struct {
	DataStoreId uint32
	ConfirmAt   uint32
	Status      uint8
}, error) {
	return _ContractDataLayrRollup.Contract.RollupStores(&_ContractDataLayrRollup.CallOpts, arg0)
}

// RollupStores is a free data retrieval call binding the contract method 0xb0393a37.
//
// Solidity: function rollupStores(uint256 ) view returns(uint32 dataStoreId, uint32 confirmAt, uint8 status)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) RollupStores(arg0 *big.Int) (struct {
	DataStoreId uint32
	ConfirmAt   uint32
	Status      uint8
}, error) {
	return _ContractDataLayrRollup.Contract.RollupStores(&_ContractDataLayrRollup.CallOpts, arg0)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) Sequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "sequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) Sequencer() (common.Address, error) {
	return _ContractDataLayrRollup.Contract.Sequencer(&_ContractDataLayrRollup.CallOpts)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) Sequencer() (common.Address, error) {
	return _ContractDataLayrRollup.Contract.Sequencer(&_ContractDataLayrRollup.CallOpts)
}

// SequencerStatus is a free data retrieval call binding the contract method 0xf83099ca.
//
// Solidity: function sequencerStatus() view returns(uint8)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) SequencerStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "sequencerStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SequencerStatus is a free data retrieval call binding the contract method 0xf83099ca.
//
// Solidity: function sequencerStatus() view returns(uint8)
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) SequencerStatus() (uint8, error) {
	return _ContractDataLayrRollup.Contract.SequencerStatus(&_ContractDataLayrRollup.CallOpts)
}

// SequencerStatus is a free data retrieval call binding the contract method 0xf83099ca.
//
// Solidity: function sequencerStatus() view returns(uint8)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) SequencerStatus() (uint8, error) {
	return _ContractDataLayrRollup.Contract.SequencerStatus(&_ContractDataLayrRollup.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ContractDataLayrRollup *ContractDataLayrRollupCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDataLayrRollup.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) StakeToken() (common.Address, error) {
	return _ContractDataLayrRollup.Contract.StakeToken(&_ContractDataLayrRollup.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ContractDataLayrRollup *ContractDataLayrRollupCallerSession) StakeToken() (common.Address, error) {
	return _ContractDataLayrRollup.Contract.StakeToken(&_ContractDataLayrRollup.CallOpts)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x7ef01d5e.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupTransactor) ConfirmData(opts *bind.TransactOpts, data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _ContractDataLayrRollup.contract.Transact(opts, "confirmData", data, searchData)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x7ef01d5e.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) ConfirmData(data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.ConfirmData(&_ContractDataLayrRollup.TransactOpts, data, searchData)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x7ef01d5e.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupTransactorSession) ConfirmData(data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.ConfirmData(&_ContractDataLayrRollup.TransactOpts, data, searchData)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x7bd85879.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupTransactor) ProveFraud(opts *bind.TransactOpts, fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs DataLayrRollupDisclosureProofs) (*types.Transaction, error) {
	return _ContractDataLayrRollup.contract.Transact(opts, "proveFraud", fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x7bd85879.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) ProveFraud(fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs DataLayrRollupDisclosureProofs) (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.ProveFraud(&_ContractDataLayrRollup.TransactOpts, fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x7bd85879.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupTransactorSession) ProveFraud(fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs DataLayrRollupDisclosureProofs) (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.ProveFraud(&_ContractDataLayrRollup.TransactOpts, fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDataLayrRollup.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) Stake() (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.Stake(&_ContractDataLayrRollup.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupTransactorSession) Stake() (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.Stake(&_ContractDataLayrRollup.TransactOpts)
}

// StoreData is a paid mutator transaction binding the contract method 0xbcb8ec90.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex) returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupTransactor) StoreData(opts *bind.TransactOpts, header []byte, duration uint8, blockNumber uint32, totalOperatorsIndex uint32) (*types.Transaction, error) {
	return _ContractDataLayrRollup.contract.Transact(opts, "storeData", header, duration, blockNumber, totalOperatorsIndex)
}

// StoreData is a paid mutator transaction binding the contract method 0xbcb8ec90.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex) returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupSession) StoreData(header []byte, duration uint8, blockNumber uint32, totalOperatorsIndex uint32) (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.StoreData(&_ContractDataLayrRollup.TransactOpts, header, duration, blockNumber, totalOperatorsIndex)
}

// StoreData is a paid mutator transaction binding the contract method 0xbcb8ec90.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex) returns()
func (_ContractDataLayrRollup *ContractDataLayrRollupTransactorSession) StoreData(header []byte, duration uint8, blockNumber uint32, totalOperatorsIndex uint32) (*types.Transaction, error) {
	return _ContractDataLayrRollup.Contract.StoreData(&_ContractDataLayrRollup.TransactOpts, header, duration, blockNumber, totalOperatorsIndex)
}

// ContractDataLayrRollupRollupStoreConfirmedIterator is returned from FilterRollupStoreConfirmed and is used to iterate over the raw logs and unpacked data for RollupStoreConfirmed events raised by the ContractDataLayrRollup contract.
type ContractDataLayrRollupRollupStoreConfirmedIterator struct {
	Event *ContractDataLayrRollupRollupStoreConfirmed // Event containing the contract specifics and raw log

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
func (it *ContractDataLayrRollupRollupStoreConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDataLayrRollupRollupStoreConfirmed)
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
		it.Event = new(ContractDataLayrRollupRollupStoreConfirmed)
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
func (it *ContractDataLayrRollupRollupStoreConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDataLayrRollupRollupStoreConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDataLayrRollupRollupStoreConfirmed represents a RollupStoreConfirmed event raised by the ContractDataLayrRollup contract.
type ContractDataLayrRollupRollupStoreConfirmed struct {
	RollupStoreNumber uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreConfirmed is a free log retrieval operation binding the contract event 0xe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a9.
//
// Solidity: event RollupStoreConfirmed(uint32 rollupStoreNumber)
func (_ContractDataLayrRollup *ContractDataLayrRollupFilterer) FilterRollupStoreConfirmed(opts *bind.FilterOpts) (*ContractDataLayrRollupRollupStoreConfirmedIterator, error) {

	logs, sub, err := _ContractDataLayrRollup.contract.FilterLogs(opts, "RollupStoreConfirmed")
	if err != nil {
		return nil, err
	}
	return &ContractDataLayrRollupRollupStoreConfirmedIterator{contract: _ContractDataLayrRollup.contract, event: "RollupStoreConfirmed", logs: logs, sub: sub}, nil
}

// WatchRollupStoreConfirmed is a free log subscription operation binding the contract event 0xe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a9.
//
// Solidity: event RollupStoreConfirmed(uint32 rollupStoreNumber)
func (_ContractDataLayrRollup *ContractDataLayrRollupFilterer) WatchRollupStoreConfirmed(opts *bind.WatchOpts, sink chan<- *ContractDataLayrRollupRollupStoreConfirmed) (event.Subscription, error) {

	logs, sub, err := _ContractDataLayrRollup.contract.WatchLogs(opts, "RollupStoreConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDataLayrRollupRollupStoreConfirmed)
				if err := _ContractDataLayrRollup.contract.UnpackLog(event, "RollupStoreConfirmed", log); err != nil {
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
func (_ContractDataLayrRollup *ContractDataLayrRollupFilterer) ParseRollupStoreConfirmed(log types.Log) (*ContractDataLayrRollupRollupStoreConfirmed, error) {
	event := new(ContractDataLayrRollupRollupStoreConfirmed)
	if err := _ContractDataLayrRollup.contract.UnpackLog(event, "RollupStoreConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDataLayrRollupRollupStoreInitializedIterator is returned from FilterRollupStoreInitialized and is used to iterate over the raw logs and unpacked data for RollupStoreInitialized events raised by the ContractDataLayrRollup contract.
type ContractDataLayrRollupRollupStoreInitializedIterator struct {
	Event *ContractDataLayrRollupRollupStoreInitialized // Event containing the contract specifics and raw log

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
func (it *ContractDataLayrRollupRollupStoreInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDataLayrRollupRollupStoreInitialized)
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
		it.Event = new(ContractDataLayrRollupRollupStoreInitialized)
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
func (it *ContractDataLayrRollupRollupStoreInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDataLayrRollupRollupStoreInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDataLayrRollupRollupStoreInitialized represents a RollupStoreInitialized event raised by the ContractDataLayrRollup contract.
type ContractDataLayrRollupRollupStoreInitialized struct {
	DataStoreId uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreInitialized is a free log retrieval operation binding the contract event 0x957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b5.
//
// Solidity: event RollupStoreInitialized(uint32 dataStoreId)
func (_ContractDataLayrRollup *ContractDataLayrRollupFilterer) FilterRollupStoreInitialized(opts *bind.FilterOpts) (*ContractDataLayrRollupRollupStoreInitializedIterator, error) {

	logs, sub, err := _ContractDataLayrRollup.contract.FilterLogs(opts, "RollupStoreInitialized")
	if err != nil {
		return nil, err
	}
	return &ContractDataLayrRollupRollupStoreInitializedIterator{contract: _ContractDataLayrRollup.contract, event: "RollupStoreInitialized", logs: logs, sub: sub}, nil
}

// WatchRollupStoreInitialized is a free log subscription operation binding the contract event 0x957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b5.
//
// Solidity: event RollupStoreInitialized(uint32 dataStoreId)
func (_ContractDataLayrRollup *ContractDataLayrRollupFilterer) WatchRollupStoreInitialized(opts *bind.WatchOpts, sink chan<- *ContractDataLayrRollupRollupStoreInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractDataLayrRollup.contract.WatchLogs(opts, "RollupStoreInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDataLayrRollupRollupStoreInitialized)
				if err := _ContractDataLayrRollup.contract.UnpackLog(event, "RollupStoreInitialized", log); err != nil {
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
func (_ContractDataLayrRollup *ContractDataLayrRollupFilterer) ParseRollupStoreInitialized(log types.Log) (*ContractDataLayrRollupRollupStoreInitialized, error) {
	event := new(ContractDataLayrRollupRollupStoreInitialized)
	if err := _ContractDataLayrRollup.contract.UnpackLog(event, "RollupStoreInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDataLayrRollupRollupStoreRevertedIterator is returned from FilterRollupStoreReverted and is used to iterate over the raw logs and unpacked data for RollupStoreReverted events raised by the ContractDataLayrRollup contract.
type ContractDataLayrRollupRollupStoreRevertedIterator struct {
	Event *ContractDataLayrRollupRollupStoreReverted // Event containing the contract specifics and raw log

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
func (it *ContractDataLayrRollupRollupStoreRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDataLayrRollupRollupStoreReverted)
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
		it.Event = new(ContractDataLayrRollupRollupStoreReverted)
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
func (it *ContractDataLayrRollupRollupStoreRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDataLayrRollupRollupStoreRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDataLayrRollupRollupStoreReverted represents a RollupStoreReverted event raised by the ContractDataLayrRollup contract.
type ContractDataLayrRollupRollupStoreReverted struct {
	RollupStoreNumber uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreReverted is a free log retrieval operation binding the contract event 0x18407f85390b98bfd30ac355138feffbf9d2519036bddba3429ec57ed328e7a8.
//
// Solidity: event RollupStoreReverted(uint32 rollupStoreNumber)
func (_ContractDataLayrRollup *ContractDataLayrRollupFilterer) FilterRollupStoreReverted(opts *bind.FilterOpts) (*ContractDataLayrRollupRollupStoreRevertedIterator, error) {

	logs, sub, err := _ContractDataLayrRollup.contract.FilterLogs(opts, "RollupStoreReverted")
	if err != nil {
		return nil, err
	}
	return &ContractDataLayrRollupRollupStoreRevertedIterator{contract: _ContractDataLayrRollup.contract, event: "RollupStoreReverted", logs: logs, sub: sub}, nil
}

// WatchRollupStoreReverted is a free log subscription operation binding the contract event 0x18407f85390b98bfd30ac355138feffbf9d2519036bddba3429ec57ed328e7a8.
//
// Solidity: event RollupStoreReverted(uint32 rollupStoreNumber)
func (_ContractDataLayrRollup *ContractDataLayrRollupFilterer) WatchRollupStoreReverted(opts *bind.WatchOpts, sink chan<- *ContractDataLayrRollupRollupStoreReverted) (event.Subscription, error) {

	logs, sub, err := _ContractDataLayrRollup.contract.WatchLogs(opts, "RollupStoreReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDataLayrRollupRollupStoreReverted)
				if err := _ContractDataLayrRollup.contract.UnpackLog(event, "RollupStoreReverted", log); err != nil {
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

// ParseRollupStoreReverted is a log parse operation binding the contract event 0x18407f85390b98bfd30ac355138feffbf9d2519036bddba3429ec57ed328e7a8.
//
// Solidity: event RollupStoreReverted(uint32 rollupStoreNumber)
func (_ContractDataLayrRollup *ContractDataLayrRollupFilterer) ParseRollupStoreReverted(log types.Log) (*ContractDataLayrRollupRollupStoreReverted, error) {
	event := new(ContractDataLayrRollupRollupStoreReverted)
	if err := _ContractDataLayrRollup.contract.UnpackLog(event, "RollupStoreReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

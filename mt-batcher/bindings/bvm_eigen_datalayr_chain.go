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

// BVMEigenDataLayrChainDisclosureProofs is an auto generated low-level Go binding around an user-defined struct.
type BVMEigenDataLayrChainDisclosureProofs struct {
	Header               []byte
	FirstChunkNumber     uint32
	Polys                [][]byte
	MultiRevealProofs    []DataLayrDisclosureLogicMultiRevealProof
	PolyEquivalenceProof BN254G2Point
}

// DataLayrDisclosureLogicMultiRevealProof is an auto generated low-level Go binding around an user-defined struct.
type DataLayrDisclosureLogicMultiRevealProof struct {
	InterpolationPoly BN254G1Point
	RevealProof       BN254G1Point
	ZeroPoly          BN254G2Point
	ZeroPolyProof     []byte
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

// BVMEigenDataLayrChainMetaData contains all meta data concerning the BVMEigenDataLayrChain contract.
var BVMEigenDataLayrChainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupStoreNumber\",\"type\":\"uint32\"}],\"name\":\"RollupStoreConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"}],\"name\":\"RollupStoreInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupStoreNumber\",\"type\":\"uint32\"}],\"name\":\"RollupStoreReverted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_STALE_MEASURE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FRAUD_STRING\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMISSION_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreSearchData\",\"name\":\"searchData\",\"type\":\"tuple\"}],\"name\":\"confirmData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataManageAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"dataStoreIdToRollupStoreNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fraudProofPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dataManageAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_submissionInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_block_stale_measure\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"polys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"parse\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fraudulentStoreNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreSearchData\",\"name\":\"searchData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"firstChunkNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"polys\",\"type\":\"bytes[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"interpolationPoly\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"revealProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"zeroPoly\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"zeroPolyProof\",\"type\":\"bytes\"}],\"internalType\":\"structDataLayrDisclosureLogic.MultiRevealProof[]\",\"name\":\"multiRevealProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"polyEquivalenceProof\",\"type\":\"tuple\"}],\"internalType\":\"structBVM_EigenDataLayrChain.DisclosureProofs\",\"name\":\"disclosureProofs\",\"type\":\"tuple\"}],\"name\":\"proveFraud\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupStoreNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rollupStores\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"confirmAt\",\"type\":\"uint32\"},{\"internalType\":\"enumBVM_EigenDataLayrChain.RollupStoreStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"totalOperatorsIndex\",\"type\":\"uint32\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405262015180609c5534801561001757600080fd5b506124de806100276000396000f3fe608060405234801561001057600080fd5b506004361061016c5760003560e01c80638b85902b116100cd578063e8d2aeed11610081578063f249502911610066578063f249502914610310578063f2a8f12414610330578063f2fde38b1461033957600080fd5b8063e8d2aeed146102ea578063eb990c59146102fd57600080fd5b8063b0393a37116100b2578063b0393a371461026f578063b537c4c7146102c2578063dcec3348146102e257600080fd5b80638b85902b146102485780638da5cb5b1461025157600080fd5b80635c1bba3811610124578063715018a611610109578063715018a6146102185780637bd85879146102225780637ef01d5e1461023557600080fd5b80635c1bba38146101ca5780635e8b3f2d1461020f57600080fd5b80634599c788116101555780634599c788146101b157806346b2eb9b146101b9578063529933df146101c157600080fd5b80631f944c8f14610171578063428bba091461019a575b600080fd5b61018461017f366004611850565b61034c565b6040516101919190611902565b60405180910390f35b6101a3609f5481565b604051908152602001610191565b609b546101a3565b6101846104b5565b6101a360975481565b6098546101ea9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610191565b6101a3609a5481565b6102206104d1565b005b610220610230366004611b2b565b6104e5565b610220610243366004611bde565b610c8a565b6101a3609b5481565b60335473ffffffffffffffffffffffffffffffffffffffff166101ea565b6102b361027d366004611c34565b609d6020526000908152604090205463ffffffff8082169164010000000081049091169068010000000000000000900460ff1683565b60405161019193929190611c7c565b6101a36102d0366004611cd6565b609e6020526000908152604090205481565b6101a3610fcb565b6102206102f8366004611cfa565b610fe8565b61022061030b366004611d7e565b61129c565b6099546101ea9073ffffffffffffffffffffffffffffffffffffffff1681565b6101a3609c5481565b610220610347366004611dc0565b61147c565b6060806000845b84835110156104a95760006103ce8261036d602082611e0a565b610378906001611e45565b610383906020611e5d565b61038d9190611e9a565b85516103999089611e9a565b848c8c888181106103ac576103ac611eb1565b90506020028101906103be9190611ee0565b6103c9929150611e9a565b611519565b9050838989858181106103e3576103e3611eb1565b90506020028101906103f59190611ee0565b84906104018583611e45565b9261040e93929190611f45565b60405160200161042093929190611f6f565b60405160208183030381529060405293508351861161043f57506104a9565b88888481811061045157610451611eb1565b90506020028101906104639190611ee0565b905061046f8284611e45565b141561048b578261047f81611f97565b935050600191506104a3565b610496816001611e45565b6104a09083611e45565b91505b50610353565b50909695505050505050565b6040518060800160405280606081526020016124496060913981565b6104d961154c565b6104e360006115b3565b565b6000848152609d602090815260408083208151606081018352815463ffffffff80821683526401000000008204169482019490945292909183019068010000000000000000900460ff16600281111561054057610540611c4d565b600281111561055157610551611c4d565b905250905060018160400151600281111561056e5761056e611c4d565b148015610584575042816020015163ffffffff16115b6105fb5760405162461bcd60e51b815260206004820152602d60248201527f526f6c6c757053746f7265206d75737420626520636f6d6d697474656420616e60448201527f6420756e636f6e6669726d65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b82516106069061162a565b6099546020850151604080870151606088015191517fed82c0ee00000000000000000000000000000000000000000000000000000000815260ff9093166004840152602483015263ffffffff16604482015273ffffffffffffffffffffffffffffffffffffffff9091169063ed82c0ee9060640160206040518083038186803b15801561069257600080fd5b505afa1580156106a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ca9190611fd0565b146107175760405162461bcd60e51b815260206004820152601e60248201527f6d6574616461746120707265696d61676520697320696e636f7272656374000060448201526064016105f2565b805183516040015163ffffffff9081169116146107c25760405162461bcd60e51b815260206004820152604260248201527f7365616368446174612773206461746173746f7265206964206973206e6f742060448201527f636f6e73697374656e74207769746820676976656e20726f6c6c75702073746f60648201527f7265000000000000000000000000000000000000000000000000000000000000608482015260a4016105f2565b6107cc8280611ee0565b6040516107da929190611fe9565b604051908190039020835151146108595760405162461bcd60e51b815260206004820152603260248201527f646973636c6f737572652070726f6f667320686561646572686173682070726560448201527f696d61676520697320696e636f7272656374000000000000000000000000000060648201526084016105f2565b73__$487cc7a1b58f8f6e007e8cd55d94b79fe8$__63899c1afc61087d8480611ee0565b61088d6040870160208801611cd6565b61089a6040880188611ff9565b6108a760608a018a611ff9565b8a6080016040518963ffffffff1660e01b81526004016108ce989796959493929190612206565b60206040518083038186803b1580156108e657600080fd5b505af41580156108fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061091e91906122d4565b61096a5760405162461bcd60e51b815260206004820152601d60248201527f646973636c6f737572652070726f6f66732061726520696e76616c696400000060448201526064016105f2565b600073__$487cc7a1b58f8f6e007e8cd55d94b79fe8$__63e34c39d66109908580611ee0565b6040518363ffffffff1660e01b81526004016109ad9291906122f6565b60206040518083038186803b1580156109c557600080fd5b505af41580156109d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109fd919061230a565b905063ffffffff8116610a136040850185611ff9565b9050610a256040860160208701611cd6565b63ffffffff16610a359190611e45565b1115610aa95760405162461bcd60e51b815260206004820152602e60248201527f43616e206f6e6c792070726f766520646174612066726f6d207468652073797360448201527f74656d61746963206368756e6b7300000000000000000000000000000000000060648201526084016105f2565b6000610adb610abb6040860186611ff9565b88604051806080016040528060608152602001612449606091395161034c565b90506040518060800160405280606081526020016124496060913951815114610b925760405162461bcd60e51b815260206004820152604260248201527f50617273696e67206572726f722c2070726f76656e20737472696e672069732060448201527f646966666572656e74206c656e677468207468616e206672617564207374726960648201527f6e67000000000000000000000000000000000000000000000000000000000000608482015260a4016105f2565b6040518060800160405280606081526020016124496060913980519060200120818051906020012014610c075760405162461bcd60e51b815260206004820152601d60248201527f70726f76656e20737472696e6720213d20667261756420737472696e6700000060448201526064016105f2565b6000878152609d602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1668020000000000000000179055905163ffffffff891681527f18407f85390b98bfd30ac355138feffbf9d2519036bddba3429ec57ed328e7a891015b60405180910390a150505050505050565b60985473ffffffffffffffffffffffffffffffffffffffff163314610d175760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79207468652073657175656e6365722063616e2073746f72652064617460448201527f610000000000000000000000000000000000000000000000000000000000000060648201526084016105f2565b805160409081015163ffffffff166000908152609e60205220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14610deb5760405162461bcd60e51b815260206004820152605560248201527f446174612073746f72652065697468657220776173206e6f7420696e6974696160448201527f6c697a65642062792074686520726f6c6c757020636f6e74726163742c206f7260648201527f20697320616c726561647920636f6e6669726d65640000000000000000000000608482015260a4016105f2565b6099546040517f5189951500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690635189951590610e4590869086908690600401612327565b600060405180830381600087803b158015610e5f57600080fd5b505af1158015610e73573d6000803e3d6000fd5b50505050604051806060016040528082600001516040015163ffffffff168152602001609c5442610ea49190611e45565b63ffffffff16815260200160019052609f546000908152609d6020908152604091829020835181549285015163ffffffff908116640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009094169116179190911780825591830151909182907fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1668010000000000000000836002811115610f5357610f53611c4d565b021790555050609f8054835160409081015163ffffffff166000908152609e6020529081208290557fe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a99350909190610faa83611f97565b9091555060405163ffffffff909116815260200160405180910390a1505050565b6000609754610fd9609b5490565b610fe39190611e45565b905090565b60985473ffffffffffffffffffffffffffffffffffffffff1633146110755760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79207468652073657175656e6365722063616e2073746f72652064617460448201527f610000000000000000000000000000000000000000000000000000000000000060648201526084016105f2565b609a5461108863ffffffff851643611e9a565b106110d55760405162461bcd60e51b815260206004820152601e60248201527f7374616b65732074616b656e2066726f6d20746f6f206c6f6e672061676f000060448201526064016105f2565b609954604080517f72d18e8d000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff16916372d18e8d916004808301926020929190829003018186803b15801561114057600080fd5b505afa158015611154573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611178919061230a565b609b8490556099546040517fdcf49ea700000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff169063dcf49ea7906111e090339030908a908a9089908f908f906004016123f1565b602060405180830381600087803b1580156111fa57600080fd5b505af115801561120e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611232919061230a565b5063ffffffff81166000818152609e60209081526040918290207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff905590519182527f957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b59101610c79565b600054610100900460ff16158080156112bc5750600054600160ff909116105b806112d65750303b1580156112d6575060005460ff166001145b6113485760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105f2565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156113a657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6113ae61174a565b6098805473ffffffffffffffffffffffffffffffffffffffff8088167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560998054928716929091169190911790556097839055609a8290556000609b55801561147557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b61148461154c565b73ffffffffffffffffffffffffffffffffffffffff811661150d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016105f2565b611516816115b3565b50565b60008284106115355781831061152f5781611544565b82611544565b8184106115425781611544565b835b949350505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146104e35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105f2565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600080826000015183602001518460400151856060015186608001518760a001518860c0015160405160200161170d979695949392919096875260e095861b7fffffffff00000000000000000000000000000000000000000000000000000000908116602089015294861b851660248801529290941b909216602885015260a09190911b7fffffffffffffffffffffffff000000000000000000000000000000000000000016602c84015260609190911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166038830152604c820152606c0190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209392505050565b600054610100900460ff166117c75760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105f2565b6104e3600054610100900460ff166118475760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105f2565b6104e3336115b3565b6000806000806060858703121561186657600080fd5b843567ffffffffffffffff8082111561187e57600080fd5b818701915087601f83011261189257600080fd5b8135818111156118a157600080fd5b8860208260051b85010111156118b657600080fd5b6020928301999098509187013596604001359550909350505050565b60005b838110156118ed5781810151838201526020016118d5565b838111156118fc576000848401525b50505050565b60208152600082518060208401526119218160408501602087016118d2565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b6040516080810167ffffffffffffffff8111828210171561199d577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b60405160e0810167ffffffffffffffff8111828210171561199d577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b63ffffffff8116811461151657600080fd5b8035611a0a816119ed565b919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611a0a57600080fd5b803560ff81168114611a0a57600080fd5b6000818303610140811215611a5857600080fd5b611a60611953565b915060e0811215611a7057600080fd5b50611a796119a3565b823581526020830135611a8b816119ed565b60208201526040830135611a9e816119ed565b60408201526060830135611ab1816119ed565b606082015260808301356bffffffffffffffffffffffff81168114611ad557600080fd5b6080820152611ae660a08401611a0f565b60a082015260c083810135908201528152611b0360e08301611a33565b60208201526101008201356040820152611b2061012083016119ff565b606082015292915050565b6000806000806101a08587031215611b4257600080fd5b8435935060208501359250611b5a8660408701611a44565b915061018085013567ffffffffffffffff811115611b7757600080fd5b85016101008188031215611b8a57600080fd5b939692955090935050565b60008083601f840112611ba757600080fd5b50813567ffffffffffffffff811115611bbf57600080fd5b602083019150836020828501011115611bd757600080fd5b9250929050565b60008060006101608486031215611bf457600080fd5b833567ffffffffffffffff811115611c0b57600080fd5b611c1786828701611b95565b9094509250611c2b90508560208601611a44565b90509250925092565b600060208284031215611c4657600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b63ffffffff8481168252831660208201526060810160038310611cc8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826040830152949350505050565b600060208284031215611ce857600080fd5b8135611cf3816119ed565b9392505050565b60008060008060008060a08789031215611d1357600080fd5b863567ffffffffffffffff811115611d2a57600080fd5b611d3689828a01611b95565b9097509550611d49905060208801611a33565b93506040870135611d59816119ed565b9250606087013591506080870135611d70816119ed565b809150509295509295509295565b60008060008060808587031215611d9457600080fd5b611d9d85611a0f565b9350611dab60208601611a0f565b93969395505050506040820135916060013590565b600060208284031215611dd257600080fd5b611cf382611a0f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082611e40577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60008219821115611e5857611e58611ddb565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611e9557611e95611ddb565b500290565b600082821015611eac57611eac611ddb565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611f1557600080fd5b83018035915067ffffffffffffffff821115611f3057600080fd5b602001915036819003821315611bd757600080fd5b60008085851115611f5557600080fd5b83861115611f6257600080fd5b5050820193919092039150565b60008451611f818184602089016118d2565b8201838582376000930192835250909392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611fc957611fc9611ddb565b5060010190565b600060208284031215611fe257600080fd5b5051919050565b8183823760009101908152919050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261202e57600080fd5b83018035915067ffffffffffffffff82111561204957600080fd5b6020019150600581901b3603821315611bd757600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126120df57600080fd5b830160208101925035905067ffffffffffffffff8111156120ff57600080fd5b803603831315611bd757600080fd5b604081833760408201600081526040808301823750600060808301525050565b81835260006020808501808196508560051b81019150846000805b888110156121f8578385038a5282357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1893603018112612187578283fd5b88018035865260208082013590870152604080820135908701526060808201359087015261012060806121be81890182850161210e565b506101006121ce818401846120aa565b935082828a01526121e2838a018583612061565b9d8a019d98505050938701935050600101612149565b509298975050505050505050565b600061010080835261221b8184018b8d612061565b9050602063ffffffff8a16818501528382036040850152818883528183019050818960051b8401018a60005b8b81101561229f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868403018452612280828e6120aa565b61228b858284612061565b958701959450505090840190600101612247565b505085810360608701526122b481898b61212e565b9450505050506122c7608083018461210e565b9998505050505050505050565b6000602082840312156122e657600080fd5b81518015158114611cf357600080fd5b602081526000611544602083018486612061565b60006020828403121561231c57600080fd5b8151611cf3816119ed565b600061016080835261233c8184018688612061565b915050825180516020840152602081015163ffffffff808216604086015280604084015116606086015280606084015116608086015250506bffffffffffffffffffffffff60808201511660a084015273ffffffffffffffffffffffffffffffffffffffff60a08201511660c084015260c081015160e08401525060208301516123cc61010084018260ff169052565b50604083015161012083015260609092015163ffffffff166101409091015292915050565b73ffffffffffffffffffffffffffffffffffffffff88811682528716602082015260ff8616604082015263ffffffff85811660608301528416608082015260c060a082018190526000906122c7908301848661206156fe2d5f2860204f2060295f2d202d5f2860206f2060295f2d202d5f286020512060295f2d2042495444414f204a5553542052454b5420594f55207c5f2860204f2060295f7c202d207c5f2860206f2060295f7c202d207c5f286020512060295f7ca264697066735822122071846f7b2c98007b064021221741e08e0b71a6b7d1c1327527375e8b3436563364736f6c63430008090033",
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

// FRAUDSTRING is a free data retrieval call binding the contract method 0x46b2eb9b.
//
// Solidity: function FRAUD_STRING() view returns(bytes)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) FRAUDSTRING(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "FRAUD_STRING")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FRAUDSTRING is a free data retrieval call binding the contract method 0x46b2eb9b.
//
// Solidity: function FRAUD_STRING() view returns(bytes)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) FRAUDSTRING() ([]byte, error) {
	return _BVMEigenDataLayrChain.Contract.FRAUDSTRING(&_BVMEigenDataLayrChain.CallOpts)
}

// FRAUDSTRING is a free data retrieval call binding the contract method 0x46b2eb9b.
//
// Solidity: function FRAUD_STRING() view returns(bytes)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) FRAUDSTRING() ([]byte, error) {
	return _BVMEigenDataLayrChain.Contract.FRAUDSTRING(&_BVMEigenDataLayrChain.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) SUBMISSIONINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "SUBMISSION_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.SUBMISSIONINTERVAL(&_BVMEigenDataLayrChain.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.SUBMISSIONINTERVAL(&_BVMEigenDataLayrChain.CallOpts)
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

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "latestBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) LatestBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.LatestBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) LatestBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.LatestBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) NextBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "nextBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) NextBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.NextBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) NextBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.NextBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress, uint256 _submissionInterval, uint256 _block_stale_measure) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) Initialize(opts *bind.TransactOpts, _sequencer common.Address, _dataManageAddress common.Address, _submissionInterval *big.Int, _block_stale_measure *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "initialize", _sequencer, _dataManageAddress, _submissionInterval, _block_stale_measure)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress, uint256 _submissionInterval, uint256 _block_stale_measure) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) Initialize(_sequencer common.Address, _dataManageAddress common.Address, _submissionInterval *big.Int, _block_stale_measure *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.Initialize(&_BVMEigenDataLayrChain.TransactOpts, _sequencer, _dataManageAddress, _submissionInterval, _block_stale_measure)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress, uint256 _submissionInterval, uint256 _block_stale_measure) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) Initialize(_sequencer common.Address, _dataManageAddress common.Address, _submissionInterval *big.Int, _block_stale_measure *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.Initialize(&_BVMEigenDataLayrChain.TransactOpts, _sequencer, _dataManageAddress, _submissionInterval, _block_stale_measure)
}

// Parse is a paid mutator transaction binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) returns(bytes)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) Parse(opts *bind.TransactOpts, polys [][]byte, startIndex *big.Int, length *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "parse", polys, startIndex, length)
}

// Parse is a paid mutator transaction binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) returns(bytes)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) Parse(polys [][]byte, startIndex *big.Int, length *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.Parse(&_BVMEigenDataLayrChain.TransactOpts, polys, startIndex, length)
}

// Parse is a paid mutator transaction binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) returns(bytes)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) Parse(polys [][]byte, startIndex *big.Int, length *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.Parse(&_BVMEigenDataLayrChain.TransactOpts, polys, startIndex, length)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x7bd85879.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) ProveFraud(opts *bind.TransactOpts, fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs BVMEigenDataLayrChainDisclosureProofs) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "proveFraud", fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x7bd85879.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) ProveFraud(fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs BVMEigenDataLayrChainDisclosureProofs) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.ProveFraud(&_BVMEigenDataLayrChain.TransactOpts, fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x7bd85879.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) ProveFraud(fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs BVMEigenDataLayrChainDisclosureProofs) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.ProveFraud(&_BVMEigenDataLayrChain.TransactOpts, fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
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

// StoreData is a paid mutator transaction binding the contract method 0xe8d2aeed.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint256 _l2BlockNumber, uint32 totalOperatorsIndex) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) StoreData(opts *bind.TransactOpts, header []byte, duration uint8, blockNumber uint32, _l2BlockNumber *big.Int, totalOperatorsIndex uint32) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "storeData", header, duration, blockNumber, _l2BlockNumber, totalOperatorsIndex)
}

// StoreData is a paid mutator transaction binding the contract method 0xe8d2aeed.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint256 _l2BlockNumber, uint32 totalOperatorsIndex) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) StoreData(header []byte, duration uint8, blockNumber uint32, _l2BlockNumber *big.Int, totalOperatorsIndex uint32) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.StoreData(&_BVMEigenDataLayrChain.TransactOpts, header, duration, blockNumber, _l2BlockNumber, totalOperatorsIndex)
}

// StoreData is a paid mutator transaction binding the contract method 0xe8d2aeed.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint256 _l2BlockNumber, uint32 totalOperatorsIndex) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) StoreData(header []byte, duration uint8, blockNumber uint32, _l2BlockNumber *big.Int, totalOperatorsIndex uint32) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.StoreData(&_BVMEigenDataLayrChain.TransactOpts, header, duration, blockNumber, _l2BlockNumber, totalOperatorsIndex)
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

// BVMEigenDataLayrChainRollupStoreRevertedIterator is returned from FilterRollupStoreReverted and is used to iterate over the raw logs and unpacked data for RollupStoreReverted events raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainRollupStoreRevertedIterator struct {
	Event *BVMEigenDataLayrChainRollupStoreReverted // Event containing the contract specifics and raw log

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
func (it *BVMEigenDataLayrChainRollupStoreRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMEigenDataLayrChainRollupStoreReverted)
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
		it.Event = new(BVMEigenDataLayrChainRollupStoreReverted)
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
func (it *BVMEigenDataLayrChainRollupStoreRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMEigenDataLayrChainRollupStoreRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMEigenDataLayrChainRollupStoreReverted represents a RollupStoreReverted event raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainRollupStoreReverted struct {
	RollupStoreNumber uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreReverted is a free log retrieval operation binding the contract event 0x18407f85390b98bfd30ac355138feffbf9d2519036bddba3429ec57ed328e7a8.
//
// Solidity: event RollupStoreReverted(uint32 rollupStoreNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) FilterRollupStoreReverted(opts *bind.FilterOpts) (*BVMEigenDataLayrChainRollupStoreRevertedIterator, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.FilterLogs(opts, "RollupStoreReverted")
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainRollupStoreRevertedIterator{contract: _BVMEigenDataLayrChain.contract, event: "RollupStoreReverted", logs: logs, sub: sub}, nil
}

// WatchRollupStoreReverted is a free log subscription operation binding the contract event 0x18407f85390b98bfd30ac355138feffbf9d2519036bddba3429ec57ed328e7a8.
//
// Solidity: event RollupStoreReverted(uint32 rollupStoreNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) WatchRollupStoreReverted(opts *bind.WatchOpts, sink chan<- *BVMEigenDataLayrChainRollupStoreReverted) (event.Subscription, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.WatchLogs(opts, "RollupStoreReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMEigenDataLayrChainRollupStoreReverted)
				if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "RollupStoreReverted", log); err != nil {
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
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) ParseRollupStoreReverted(log types.Log) (*BVMEigenDataLayrChainRollupStoreReverted, error) {
	event := new(BVMEigenDataLayrChainRollupStoreReverted)
	if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "RollupStoreReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// BVMEigenDataLayrChainBatchRollupBlock is an auto generated low-level Go binding around an user-defined struct.
type BVMEigenDataLayrChainBatchRollupBlock struct {
	StartL2BlockNumber *big.Int
	EndBL2BlockNumber  *big.Int
	IsReRollup         bool
}

// BVMEigenDataLayrChainDisclosureProofs is an auto generated low-level Go binding around an user-defined struct.
type BVMEigenDataLayrChainDisclosureProofs struct {
	Header               []byte
	FirstChunkNumber     uint32
	Polys                [][]byte
	MultiRevealProofs    []DataLayrDisclosureLogicMultiRevealProof
	PolyEquivalenceProof BN254G2Point
}

// BVMEigenDataLayrChainRollupStore is an auto generated low-level Go binding around an user-defined struct.
type BVMEigenDataLayrChainRollupStore struct {
	OriginDataStoreId uint32
	DataStoreId       uint32
	ConfirmAt         uint32
	Status            uint8
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
	HeaderHash           [32]byte
	DurationDataStoreId  uint32
	GlobalDataStoreId    uint32
	ReferenceBlockNumber uint32
	BlockNumber          uint32
	Fee                  *big.Int
	Confirmer            common.Address
	SignatoryRecordHash  [32]byte
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reRollupIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rollupBatchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stratL2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endL2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"ReRollupBatchData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rollupBatchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stratL2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endL2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"RollupStoreConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stratL2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endL2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"RollupStoreInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rollupBatchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stratL2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endL2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"RollupStoreReverted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_STALE_MEASURE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FRAUD_STRING\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreSearchData\",\"name\":\"searchData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"startL2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endL2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"originDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"reConfirmedBatchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isReRollup\",\"type\":\"bool\"}],\"name\":\"confirmData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataManageAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"dataStoreIdToL2RollUpBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startL2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBL2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isReRollup\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"dataStoreIdToRollupStoreNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fraudProofPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2ConfirmedBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_dataStoreId\",\"type\":\"uint32\"}],\"name\":\"getL2RollUpBlockByDataStoreId\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startL2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBL2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isReRollup\",\"type\":\"bool\"}],\"internalType\":\"structBVM_EigenDataLayrChain.BatchRollupBlock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2StoredBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rollupBatchIndex\",\"type\":\"uint256\"}],\"name\":\"getRollupStoreByRollupBatchIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"confirmAt\",\"type\":\"uint32\"},{\"internalType\":\"enumBVM_EigenDataLayrChain.RollupStoreStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structBVM_EigenDataLayrChain.RollupStore\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dataManageAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reSubmitterAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_block_stale_measure\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fraudProofPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2SubmittedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ConfirmedBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2StoredBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"polys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"parse\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"provenString\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fraudulentStoreNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreSearchData\",\"name\":\"searchData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"firstChunkNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"polys\",\"type\":\"bytes[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"interpolationPoly\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"revealProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"zeroPoly\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"zeroPolyProof\",\"type\":\"bytes\"}],\"internalType\":\"structDataLayrDisclosureLogic.MultiRevealProof[]\",\"name\":\"multiRevealProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"polyEquivalenceProof\",\"type\":\"tuple\"}],\"internalType\":\"structBVM_EigenDataLayrChain.DisclosureProofs\",\"name\":\"disclosureProofs\",\"type\":\"tuple\"}],\"name\":\"proveFraud\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reRollupBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reRollupIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reSubmitterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeFraudProofAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rollupBatchIndex\",\"type\":\"uint256\"}],\"name\":\"resetRollupBatchData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rollupBatchIndexRollupStores\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"confirmAt\",\"type\":\"uint32\"},{\"internalType\":\"enumBVM_EigenDataLayrChain.RollupStoreStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setFraudProofAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"startL2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endL2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"totalOperatorsIndex\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isReRollup\",\"type\":\"bool\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"submitReRollUpInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"unavailableFraudProofAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dataManageAddress\",\"type\":\"address\"}],\"name\":\"updateDataLayrManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fraudProofPeriod\",\"type\":\"uint256\"}],\"name\":\"updateFraudProofPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2ConfirmedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"updateL2ConfirmedBlockNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2StoredBlockNumber\",\"type\":\"uint256\"}],\"name\":\"updateL2StoredBlockNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reSubmitterAddress\",\"type\":\"address\"}],\"name\":\"updateReSubmitterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"}],\"name\":\"updateSequencerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50614d54806100206000396000f3fe608060405234801561001057600080fd5b50600436106102775760003560e01c8063728cdbca11610160578063afab4ac5116100d8578063f24950291161008c578063f2fde38b11610071578063f2fde38b146105cc578063f7db9795146105df578063ff2e0749146105f257600080fd5b8063f2495029146105a3578063f2a8f124146105c357600080fd5b8063c8fff01b116100bd578063c8fff01b14610546578063c96c0d3814610559578063d7fbc2e21461059057600080fd5b8063afab4ac514610513578063b537c4c71461052657600080fd5b8063927f20321161012f5780639495de40116101145780639495de40146104e4578063990fca66146104f75780639a71e29c1461050057600080fd5b8063927f20321461048c57806392f30a451461049557600080fd5b8063728cdbca14610433578063758b8147146104465780638bea6cae146104665780638da5cb5b1461046e57600080fd5b80633c762984116101f35780635c1bba38116101c25780635e4a3056116101a75780635e4a30561461040f5780635e8b3f2d14610422578063715018a61461042b57600080fd5b80635c1bba38146103c15780635d42ffb71461040657600080fd5b80633c762984146103385780634618ed871461034157806346b2eb9b1461035457806359cb63911461035c57600080fd5b80631f944c8f1161024a5780632e72866b1161022f5780632e72866b146102f3578063301b39ab1461031357806332c58f7a1461032557600080fd5b80631f944c8f146102b75780632e64b4c0146102e057600080fd5b806302d777de1461027c578063060ee9a4146102915780630a33202e1461029157806315fda737146102a4575b600080fd5b61028f61028a366004613f01565b610612565b005b61028f61029f366004613f01565b6106eb565b61028f6102b23660046140eb565b6107c4565b6102ca6102c5366004614155565b610f4a565b6040516102d79190614251565b60405180910390f35b61028f6102ee366004614264565b6110ef565b610306610301366004614264565b611181565b6040516102d791906142e7565b609a545b6040519081526020016102d7565b61028f610333366004613f01565b61121a565b610317609d5481565b61028f61034f36600461437e565b6112f6565b6102ca611a00565b6103b161036a366004614264565b609e6020526000908152604090205463ffffffff80821691640100000000810482169168010000000000000000820416906c01000000000000000000000000900460ff1684565b6040516102d79493929190614419565b6097546103e19073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102d7565b610317609b5481565b61028f61041d36600461444c565b611a1c565b61031760995481565b61028f611d5d565b61028f6104413660046144da565b611d71565b60a2546103e19073ffffffffffffffffffffffffffffffffffffffff1681565b609b54610317565b60335473ffffffffffffffffffffffffffffffffffffffff166103e1565b61031760a35481565b6104c76104a3366004614539565b609f6020526000908152604090208054600182015460029092015490919060ff1683565b6040805193845260208401929092521515908201526060016102d7565b61028f6104f2366004614264565b611f61565b610317609a5481565b61028f61050e366004614264565b611ff3565b61028f610521366004613f01565b6121b5565b610317610534366004614539565b60a06020526000908152604090205481565b61028f610554366004613f01565b612289565b61056c610567366004614539565b61235d565b604080518251815260208084015190820152918101511515908201526060016102d7565b61028f61059e366004614264565b6123c7565b6098546103e19073ffffffffffffffffffffffffffffffffffffffff1681565b610317609c5481565b61028f6105da366004613f01565b612459565b61028f6105ed366004614264565b6124f6565b610317610600366004614264565b60a46020526000908152604090205481565b60975473ffffffffffffffffffffffffffffffffffffffff1633146106a45760405162461bcd60e51b815260206004820152602a60248201527f4f6e6c79207468652073657175656e6365722063616e2075706461746520646c60448201527f736d20616464726573730000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b609880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60975473ffffffffffffffffffffffffffffffffffffffff1633146107785760405162461bcd60e51b815260206004820152603160248201527f4f6e6c79207468652073657175656e6365722063616e2072656d6f766520667260448201527f6175642070726f6f662061646472657373000000000000000000000000000000606482015260840161069b565b73ffffffffffffffffffffffffffffffffffffffff16600090815260a16020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b33600090815260a1602052604090205460ff16151560011461084e5760405162461bcd60e51b815260206004820152602e60248201527f4f6e6c792066726175642070726f6f66207768697465206c6973742063616e2060448201527f6368616c6c656e67652064617461000000000000000000000000000000000000606482015260840161069b565b6000848152609e602090815260408083208151608081018352815463ffffffff8082168352640100000000820481169583019590955268010000000000000000810490941692810192909252909160608301906c01000000000000000000000000900460ff1660028111156108c5576108c561427d565b60028111156108d6576108d661427d565b90525090506001816060015160028111156108f3576108f361427d565b148015610909575042816040015163ffffffff16115b61097b5760405162461bcd60e51b815260206004820152602d60248201527f526f6c6c757053746f7265206d75737420626520636f6d6d697474656420616e60448201527f6420756e636f6e6669726d656400000000000000000000000000000000000000606482015260840161069b565b8251610986906125e6565b6098546020850151604080870151606088015191517fed82c0ee00000000000000000000000000000000000000000000000000000000815260ff9093166004840152602483015263ffffffff16604482015273ffffffffffffffffffffffffffffffffffffffff9091169063ed82c0ee9060640160206040518083038186803b158015610a1257600080fd5b505afa158015610a26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4a9190614556565b14610a975760405162461bcd60e51b815260206004820152601e60248201527f6d6574616461746120707265696d61676520697320696e636f72726563740000604482015260640161069b565b806020015163ffffffff1683600001516040015163ffffffff1614610b4a5760405162461bcd60e51b815260206004820152604260248201527f7365616368446174612773206461746173746f7265206964206973206e6f742060448201527f636f6e73697374656e74207769746820676976656e20726f6c6c75702073746f60648201527f7265000000000000000000000000000000000000000000000000000000000000608482015260a40161069b565b610b54828061456f565b604051610b629291906145d4565b60405190819003902083515114610be15760405162461bcd60e51b815260206004820152603260248201527f646973636c6f737572652070726f6f667320686561646572686173682070726560448201527f696d61676520697320696e636f72726563740000000000000000000000000000606482015260840161069b565b610c21610bee838061456f565b610bfe6040860160208701614539565b610c0b60408701876145e4565b610c1860608901896145e4565b89608001612713565b610c6d5760405162461bcd60e51b815260206004820152601d60248201527f646973636c6f737572652070726f6f66732061726520696e76616c6964000000604482015260640161069b565b6000610c81610c7c848061456f565b612d0b565b905063ffffffff8116610c9760408501856145e4565b9050610ca96040860160208701614539565b63ffffffff16610cb9919061467b565b1115610d2d5760405162461bcd60e51b815260206004820152602e60248201527f43616e206f6e6c792070726f766520646174612066726f6d207468652073797360448201527f74656d61746963206368756e6b73000000000000000000000000000000000000606482015260840161069b565b6000610d5f610d3f60408601866145e4565b88604051806080016040528060608152602001614c576060913951610f4a565b9050604051806080016040528060608152602001614c576060913951815114610e165760405162461bcd60e51b815260206004820152604260248201527f50617273696e67206572726f722c2070726f76656e20737472696e672069732060448201527f646966666572656e74206c656e677468207468616e206672617564207374726960648201527f6e67000000000000000000000000000000000000000000000000000000000000608482015260a40161069b565b604051806080016040528060608152602001614c576060913980519060200120818051906020012014610e8b5760405162461bcd60e51b815260206004820152601d60248201527f70726f76656e20737472696e6720213d20667261756420737472696e67000000604482015260640161069b565b6000878152609e6020908152604080832080547fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff166c02000000000000000000000000179055875181015163ffffffff908116808552609f8452828520548a518401519092168552938290206001015482518c8152938401949094529082015260608101919091527fca227c67a02028763083580d42e8bdef4bb49c393068d05983421cd7a4a2a5be906080015b60405180910390a150505050505050565b6060610f576020846146c2565b610fc95760405162461bcd60e51b815260206004820152602760248201527f43616e6e6f742073746172742072656164696e672066726f6d2061207061646460448201527f6564206279746500000000000000000000000000000000000000000000000000606482015260840161069b565b6000835b83835110156110e557600061101882610fe76020826146d6565b610ff290600161467b565b610ffd9060206146ea565b6110079190614727565b85516110139088614727565b612d1a565b90508388888581811061102d5761102d61473e565b905060200281019061103f919061456f565b849061104b858361467b565b926110589392919061476d565b60405160200161106a93929190614797565b604051602081830303815290604052935087878481811061108d5761108d61473e565b905060200281019061109f919061456f565b90506110ab828461467b565b14156110c757826110bb816147bf565b935050600191506110df565b6110d281600161467b565b6110dc908361467b565b91505b50610fcd565b5050949350505050565b60975473ffffffffffffffffffffffffffffffffffffffff16331461117c5760405162461bcd60e51b815260206004820152603160248201527f4f6e6c79207468652073657175656e6365722063616e20736574206c6174657360448201527f74206c3220626c6f636b206e756d626572000000000000000000000000000000606482015260840161069b565b609b55565b611189613dfe565b6000828152609e60209081526040918290208251608081018452815463ffffffff80821683526401000000008204811694830194909452680100000000000000008104909316938101939093529060608301906c01000000000000000000000000900460ff1660028111156112005761120061427d565b60028111156112115761121161427d565b90525092915050565b60975473ffffffffffffffffffffffffffffffffffffffff1633146112a75760405162461bcd60e51b815260206004820152603a60248201527f4f6e6c79207468652073657175656e6365722063616e2073657420667261756460448201527f2070726f6f66206164647265737320756e617661696c61626c65000000000000606482015260840161069b565b73ffffffffffffffffffffffffffffffffffffffff16600090815260a16020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b60975473ffffffffffffffffffffffffffffffffffffffff1633146113835760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79207468652073657175656e6365722063616e2073746f72652064617460448201527f6100000000000000000000000000000000000000000000000000000000000000606482015260840161069b565b6113a4604051806060016040528060388152602001614cb760389139612d32565b855160409081015163ffffffff166000908152609f6020522054851480156113e75750855160409081015163ffffffff166000908152609f602052206001015484145b80156114155750855160409081015163ffffffff166000908152609f602052206002015460ff161515811515145b6114ad5760405162461bcd60e51b815260206004820152605560248201527f446174612073746f72652065697468657220776173206e6f7420696e6974696160448201527f6c697a65642062792074686520726f6c6c757020636f6e74726163742c206f7260648201527f20697320616c726561647920636f6e6669726d65640000000000000000000000608482015260a40161069b565b6114ce604051806060016040528060308152602001614cef60309139612d32565b855160409081015163ffffffff16600090815260a060205220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff146115a25760405162461bcd60e51b815260206004820152605560248201527f446174612073746f72652065697468657220776173206e6f7420696e6974696160448201527f6c697a65642062792074686520726f6c6c757020636f6e74726163742c206f7260648201527f20697320616c726561647920636f6e6669726d65640000000000000000000000608482015260a40161069b565b6115c36040518060600160405280602f8152602001614c28602f9139612d32565b6098546040517f58942e7300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116906358942e739061161d908b908b908b90600401614841565b600060405180830381600087803b15801561163757600080fd5b505af115801561164b573d6000803e3d6000fd5b50505050806118355760408051608081018252875182015163ffffffff90811682528851830151166020820152609c54909182019061168a904261467b565b63ffffffff16815260200160019052609d546000908152609e602090815260409182902083518154928501519385015163ffffffff90811668010000000000000000027fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff958216640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009095169190921617929092179283168217815560608401519092909183917fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff9091167fffffffffffffffffffffffffffffffffffffff0000000000ffffffffffffffff909116176c010000000000000000000000008360028111156117a0576117a061427d565b02179055505050609b849055609d8054875160409081015163ffffffff16600090815260a06020529081208290557fc7c0900be05d2a0ad0f77852eb975d9e862d1db0a2238617dd0f77854782f67292906117fa836147bf565b909155508751604090810151815163ffffffff93841681529216602083015281018790526060810186905260800160405180910390a16119f6565b60405180608001604052808463ffffffff16815260200187600001516040015163ffffffff168152602001609c544261186e919061467b565b63ffffffff168152602001600190526000838152609e602090815260409182902083518154928501519385015163ffffffff90811668010000000000000000027fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff958216640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009095169190921617929092179283168217815560608401519092909183917fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff9091167fffffffffffffffffffffffffffffffffffffff0000000000ffffffffffffffff909116176c010000000000000000000000008360028111156119815761198161427d565b021790555050865160409081015163ffffffff908116600090815260a060209081529083902086905589518301518351878152921690820152908101879052606081018690527fc7c0900be05d2a0ad0f77852eb975d9e862d1db0a2238617dd0f77854782f672915060800160405180910390a15b5050505050505050565b604051806080016040528060608152602001614c576060913981565b60975473ffffffffffffffffffffffffffffffffffffffff163314611aa95760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79207468652073657175656e6365722063616e2073746f72652064617460448201527f6100000000000000000000000000000000000000000000000000000000000000606482015260840161069b565b609954611abc63ffffffff871643614727565b10611b095760405162461bcd60e51b815260206004820152601e60248201527f7374616b65732074616b656e2066726f6d20746f6f206c6f6e672061676f0000604482015260640161069b565b609854604080517f72d18e8d000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff16916372d18e8d916004808301926020929190829003018186803b158015611b7457600080fd5b505afa158015611b88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bac919061491e565b9050609860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dcf49ea733308a8a888f8f6040518863ffffffff1660e01b8152600401611c15979695949392919061493b565b602060405180830381600087803b158015611c2f57600080fd5b505af1158015611c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c67919061491e565b5060408051606081018252868152602080820187815285151583850190815263ffffffff86166000908152609f8452858120945185559151600185015551600290930180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169315159390931790925560a09052207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff905581611d0b57609a8490555b6040805163ffffffff83168152602081018790529081018590527fa99ca06ac3461399088feac88ec48dc5a47d61c3b6839eab20146f2c4ee535849060600160405180910390a1505050505050505050565b611d65612dc1565b611d6f6000612e28565b565b600054610100900460ff1615808015611d915750600054600160ff909116105b80611dab5750303b158015611dab575060005460ff166001145b611e1d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161069b565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611e7b57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b611e83612e9f565b6097805473ffffffffffffffffffffffffffffffffffffffff808a167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556098805489841690831617905560a28054928816929091169190911790556099849055609c839055609a829055609b8290558015611f5857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610f39565b50505050505050565b60975473ffffffffffffffffffffffffffffffffffffffff163314611fee5760405162461bcd60e51b815260206004820152603160248201527f4f6e6c79207468652073657175656e6365722063616e20736574206c6174657360448201527f74206c3220626c6f636b206e756d626572000000000000000000000000000000606482015260840161069b565b609a55565b60a25473ffffffffffffffffffffffffffffffffffffffff1633146120805760405162461bcd60e51b815260206004820152602f60248201527f4f6e6c7920746865207265207375626d69747465722063616e207375626d697460448201527f20726520726f6c6c757020646174610000000000000000000000000000000000606482015260840161069b565b6000818152609e602090815260408083208151608081018352815463ffffffff8082168352640100000000820481169583019590955268010000000000000000810490941692810192909252909160608301906c01000000000000000000000000900460ff1660028111156120f7576120f761427d565b60028111156121085761210861427d565b905250602081015190915063ffffffff16156121b15760a38054600090815260a46020526040812084905581547fee84ab0752d66e31e484f6855689d7067ecd900a6c5a198a2908f74e583e7d57929091612162836147bf565b909155506020838101805163ffffffff9081166000908152609f845260408082205493519092168152819020600101548151948552928401879052830152606082015260800160405180910390a15b5050565b60975473ffffffffffffffffffffffffffffffffffffffff1633146122425760405162461bcd60e51b815260206004820152603260248201527f4f6e6c79207468652073657175656e6365722063616e2075706461746520726560448201527f207375626d697474657220616464726573730000000000000000000000000000606482015260840161069b565b60a280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60975473ffffffffffffffffffffffffffffffffffffffff1633146123165760405162461bcd60e51b815260206004820152602f60248201527f4f6e6c79207468652073657175656e6365722063616e2075706461746520736560448201527f7175656e63657220616464726573730000000000000000000000000000000000606482015260840161069b565b609780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b612383604051806060016040528060008152602001600081526020016000151581525090565b5063ffffffff166000908152609f60209081526040918290208251606081018452815481526001820154928101929092526002015460ff1615159181019190915290565b60975473ffffffffffffffffffffffffffffffffffffffff1633146124545760405162461bcd60e51b815260206004820152603060248201527f4f6e6c79207468652073657175656e6365722063616e2075706461746520667260448201527f6175642070726f6f6620706572696f6400000000000000000000000000000000606482015260840161069b565b609c55565b612461612dc1565b73ffffffffffffffffffffffffffffffffffffffff81166124ea5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161069b565b6124f381612e28565b50565b60975473ffffffffffffffffffffffffffffffffffffffff1633146125835760405162461bcd60e51b815260206004820152602f60248201527f4f6e6c79207468652073657175656e6365722063616e2075706461746520736560448201527f7175656e63657220616464726573730000000000000000000000000000000000606482015260840161069b565b60005b609d548110156125d6576000818152609e6020526040902080547fffffffffffffffffffffffffffffffffffffff00000000000000000000000000169055806125ce816147bf565b915050612586565b50609d556001609a819055609b55565b600080826000015183602001518460400151856060015186608001518760a001518860c001518960e001516040516020016126d698979695949392919097885260e096871b7fffffffff0000000000000000000000000000000000000000000000000000000090811660208a015295871b8616602489015293861b851660288801529190941b909216602c85015260a09290921b7fffffffffffffffffffffffff000000000000000000000000000000000000000016603084015260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016603c830152605082015260700190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209392505050565b6000808567ffffffffffffffff81111561272f5761272f613f1c565b604051908082528060200260200182016040528015612758578160200160208202803683370190505b50905060006127e58b8b6040805160c08101825260006080820181815260a0830182905282526020820181905291810182905260608101919091525050604080518082018252823581526020838101358183015282516080810184529182528383013560e090811c918301919091526044840135811c92820192909252604890920135901c606082015290565b90508460005b81811015612ac4576128bb83612801838e614992565b8a8a858181106128135761281361473e565b905060200281019061282591906149ba565b8b8b868181106128375761283761473e565b905060200281019061284991906149ba565b6040018c8c8781811061285e5761285e61473e565b905060200281019061287091906149ba565b6080018036038101906128839190614a48565b8d8d888181106128955761289561473e565b90506020028101906128a791906149ba565b6128b69061010081019061456f565b612f24565b61292d5760405162461bcd60e51b815260206004820152602260248201527f52657665616c206661696c65642064756520746f206e6f6e203120706169726960448201527f6e67000000000000000000000000000000000000000000000000000000000000606482015260840161069b565b89898281811061293f5761293f61473e565b9050602002810190612951919061456f565b9050836020015160206129649190614aa5565b65ffffffffffff16146129df5760405162461bcd60e51b815260206004820152603860248201527f506f6c796e6f6d69616c206d757374206861766520612032353620626974206360448201527f6f656666696369656e7420666f722065616368207465726d0000000000000000606482015260840161069b565b8989828181106129f1576129f161473e565b9050602002810190612a03919061456f565b604051612a119291906145d4565b6040518091039020888883818110612a2b57612a2b61473e565b9050602002810190612a3d91906149ba565b35898984818110612a5057612a5061473e565b9050602002810190612a6291906149ba565b604051612a8993929160209081013591019283526020830191909152604082015260600190565b60405160208183030381529060405280519060200120848281518110612ab157612ab161473e565b60209081029190910101526001016127eb565b5060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000184604051602001612af99190614b06565b6040516020818303038152906040528051906020012060001c612b1c91906146c2565b9050612b26613e26565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001856000604051602001612b5b929190614b12565b6040516020818303038152906040528051906020012060001c612b7e91906146c2565b8082526020820152600089898281612b9857612b9861473e565b9050602002810190612baa91906149ba565b612bba9036819003810190614b2b565b90506000612bec8d8d6000818110612bd457612bd461473e565b9050602002810190612be6919061456f565b8661307b565b905060015b8a811015612ce957612c4383612c3e8e8e85818110612c1257612c1261473e565b9050602002810190612c2491906149ba565b612c349036819003810190614b2b565b6020880151613121565b6131c8565b92506000612c748f8f84818110612c5c57612c5c61473e565b9050602002810190612c6e919061456f565b8861307b565b90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000180828760016020020151098408602086015186519194507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001910960208601525080612ce1816147bf565b915050612bf1565b50612cf6828a868461326c565b97505050505050505098975050505050505050565b604482013560e01c5b92915050565b6000818310612d295781612d2b565b825b9392505050565b6124f381604051602401612d469190614251565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f41304fac00000000000000000000000000000000000000000000000000000000179052613347565b60335473ffffffffffffffffffffffffffffffffffffffff163314611d6f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161069b565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16612f1c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161069b565b611d6f613368565b6000612fdf83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505060208b0151612f75915065ffffffffffff166133ee565b86516020808201519151818a01518083015190516040805194850195909552938301919091526060820152608081019190915260a00160405160208183030381529060405280519060200120612fd48b8d604001518e606001516136ee565b63ffffffff1661385a565b61302b5760405162461bcd60e51b815260206004820181905260248201527f496e636f7272656374207a65726f20706f6c79206d65726b6c652070726f6f66604482015260640161069b565b875160009061304b90612c3e613046368b90038b018b614b2b565b613872565b905061306e61305f36889003880188614b2b565b8683613069613931565b6139f1565b9998505050505050505050565b600080836001825b8281101561311557600088828961309b82602061467b565b926130a89392919061476d565b6130b191614b5d565b90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000180848309860894507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878409925061310d60208361467b565b915050613083565b50919695505050505050565b604080518082019091526000808252602082015261313d613e44565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561317057613172565bfe5b50806131c05760405162461bcd60e51b815260206004820152600d60248201527f65632d6d756c2d6661696c656400000000000000000000000000000000000000604482015260640161069b565b505092915050565b60408051808201909152600080825260208201526131e4613e62565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156131705750806131c05760405162461bcd60e51b815260206004820152600d60248201527f65632d6164642d6661696c656400000000000000000000000000000000000000604482015260640161069b565b6040805180820190915260018152600260208201526000908161329761329183613872565b86613121565b604080518082019091527f220ac48bb1f91fd93f502a3d0caa077ac70e0af8819b9d8fa26a168a2c558a5781527f08f54b82af08ceaf7cd5f180bac94870f6d8100a9c9afa9dd09a44916538911260208201529091506132f781836131c8565b9150600061330d61330785613872565b87613121565b9050600061331b8a836131c8565b905061333984613330368c90038c018c614a48565b83613069613931565b9a9950505050505050505050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b600054610100900460ff166133e55760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161069b565b611d6f33612e28565b6000806133fa83613c66565b90508061342957507fe82cea94884b1b895ea0742840a3b19249a723810fd1b04d8564d675b0a416f192915050565b806001141561345a57507f4843774a80fc8385b31024f5bd18b42e62de439206ab9468d42d826796d41f6792915050565b806002141561348b57507f092d3e5f87f5293e7ab0cc2ca6b0b5e4adb5e0011656544915f7cea34e69e5ab92915050565b80600314156134bc57507f494b208540ec8624fbbb3f2c64ffccdaf6253f8f4e50c0d93922d88195b0775592915050565b80600414156134ed57507ffdb44b84a82893cfa0e37a97f09ffc4298ad5e62be1bea1d03320ae836213d2292915050565b806005141561351e57507f3f50cb08231d2a76853ba9dbb20dad45a1b75c57cdaff6223bfe069752cff3d492915050565b806006141561354f57507fbb39eebd8138eefd5802a49d571e65b3e0d4e32277c28fbf5fbca66e7fb0431092915050565b806007141561358057507ff0a39b513e11fa80cbecbf352f69310eddd5cd03148768e0e9542bd600b133ec92915050565b80600814156135b157507f038cca2238865414efb752cc004fffec9e6069b709f495249cdf36efbd5952f692915050565b80600914156135e257507f2a26b054ed559dd255d8ac9060ebf6b95b768d87de767f8174ad2f9a4e48dd0192915050565b80600a141561361357507f1fe180d0bc4ff7c69fefa595b3b5f3c284535a280f6fdcf69b20770d1e20e1fc92915050565b80600b141561364457507f60e34ad57c61cd6fdd8177437c30e4a30334e63d7683989570cf27020efc820192915050565b80600c141561367557507feda2417e770ddbe88f083acf06b6794dfb76301314a32bd0697440d76f6cd9cc92915050565b80600d14156136a657507f8cbe9b8cf92ce70e3bec8e1e72a0f85569017a7e43c3db50e4a5badb8dea7ce892915050565b60405162461bcd60e51b815260206004820152601660248201527f4c6f67206e6f7420696e2076616c69642072616e676500000000000000000000604482015260640161069b565b6000806136fb8385614992565b9050600061370e8563ffffffff16613cd6565b9050600061371c8684614b99565b63ffffffff161561372e576001613731565b60005b60ff1661373e8785614bbc565b6137489190614992565b905060006137646137598385614bdf565b63ffffffff16613cd6565b90508663ffffffff168863ffffffff1610156137a65780613785828a613d03565b61379190610100614bdf565b61379b9190614bbc565b945050505050612d2b565b6137b08784614c02565b6137ba9082614c02565b63ffffffff168863ffffffff1610156137ec578061378581856137dd8b8d614c02565b6137e79190614992565b613d03565b60405162461bcd60e51b815260206004820152603260248201527f43616e6e6f7420637265617465206e756d626572206f66206672616d6520686960448201527f67686572207468616e20706f737369626c650000000000000000000000000000606482015260840161069b565b600083613868868585613d42565b1495945050505050565b6040805180820190915260008082526020820152815115801561389757506020820151155b156138b5575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4784602001516138fa91906146c2565b613924907f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47614727565b905292915050565b919050565b613939613e80565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082018252858152602080820185905282518084019093528583528201839052600091613a20613ea0565b60005b6002811015613be5576000613a398260066146ea565b9050848260028110613a4d57613a4d61473e565b60200201515183613a5f83600061467b565b600c8110613a6f57613a6f61473e565b6020020152848260028110613a8657613a8661473e565b60200201516020015183826001613a9d919061467b565b600c8110613aad57613aad61473e565b6020020152838260028110613ac457613ac461473e565b6020020151515183613ad783600261467b565b600c8110613ae757613ae761473e565b6020020152838260028110613afe57613afe61473e565b6020020151516001602002015183613b1783600361467b565b600c8110613b2757613b2761473e565b6020020152838260028110613b3e57613b3e61473e565b602002015160200151600060028110613b5957613b5961473e565b602002015183613b6a83600461467b565b600c8110613b7a57613b7a61473e565b6020020152838260028110613b9157613b9161473e565b602002015160200151600160028110613bac57613bac61473e565b602002015183613bbd83600561467b565b600c8110613bcd57613bcd61473e565b60200201525080613bdd816147bf565b915050613a23565b50613bee613ebf565b60006020826101808560086107d05a03fa9050808015613170575080613c565760405162461bcd60e51b815260206004820152601560248201527f70616972696e672d6f70636f64652d6661696c65640000000000000000000000604482015260640161069b565b5051151598975050505050505050565b6000808211613cb75760405162461bcd60e51b815260206004820152601360248201527f4c6f67206d75737420626520646566696e656400000000000000000000000000604482015260640161069b565b60005b600183821c14612d145780613cce816147bf565b915050613cba565b600060015b82816001901b1015613cf95780613cf1816147bf565b915050613cdb565b6001901b92915050565b600080613d158463ffffffff16613c66565b613d20906020614c02565b90508063ffffffff16613d3284613db0565b63ffffffff16901c949350505050565b60008260205b85518111613da757613d5b6002856146c2565b613d7c57816000528086015160205260406000209150600284049350613d95565b8086015160005281602052604060002091506002840493505b613da060208261467b565b9050613d48565b50949350505050565b600080805b6020811015613df7576001811b84811663ffffffff1615613de457613ddb82601f614727565b6001901b831792505b5080613def816147bf565b915050613db5565b5092915050565b604080516080810182526000808252602082018190529181018290529060608201905b905290565b60405180604001604052806002906020820280368337509192915050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280613e93613e26565b8152602001613e21613e26565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461392c57600080fd5b600060208284031215613f1357600080fd5b612d2b82613edd565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715613f6e57613f6e613f1c565b60405290565b604051610100810167ffffffffffffffff81118282101715613f6e57613f6e613f1c565b6040805190810167ffffffffffffffff81118282101715613f6e57613f6e613f1c565b63ffffffff811681146124f357600080fd5b803561392c81613fbb565b80356bffffffffffffffffffffffff8116811461392c57600080fd5b803560ff8116811461392c57600080fd5b600081830361016081121561401957600080fd5b614021613f4b565b91506101008082121561403357600080fd5b61403b613f74565b915083358252602084013561404f81613fbb565b602083015261406060408501613fcd565b604083015261407160608501613fcd565b606083015261408260808501613fcd565b608083015261409360a08501613fd8565b60a08301526140a460c08501613edd565b60c083015260e084013560e08301528183526140c1818501613ff4565b6020840152505061012082013560408201526140e06101408301613fcd565b606082015292915050565b6000806000806101c0858703121561410257600080fd5b843593506020850135925061411a8660408701614005565b91506101a085013567ffffffffffffffff81111561413757600080fd5b8501610100818803121561414a57600080fd5b939692955090935050565b6000806000806060858703121561416b57600080fd5b843567ffffffffffffffff8082111561418357600080fd5b818701915087601f83011261419757600080fd5b8135818111156141a657600080fd5b8860208260051b85010111156141bb57600080fd5b6020928301999098509187013596604001359550909350505050565b60005b838110156141f25781810151838201526020016141da565b83811115614201576000848401525b50505050565b6000815180845261421f8160208601602086016141d7565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000612d2b6020830184614207565b60006020828403121561427657600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600381106142e3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b600060808201905063ffffffff808451168352806020850151166020840152806040850151166040840152506060830151613df760608401826142ac565b60008083601f84011261433757600080fd5b50813567ffffffffffffffff81111561434f57600080fd5b60208301915083602082850101111561436757600080fd5b9250929050565b8035801515811461392c57600080fd5b600080600080600080600080610220898b03121561439b57600080fd5b883567ffffffffffffffff8111156143b257600080fd5b6143be8b828c01614325565b90995097506143d290508a60208b01614005565b955061018089013594506101a089013593506101c08901356143f381613fbb565b92506101e0890135915061440a6102008a0161436e565b90509295985092959890939650565b63ffffffff85811682528481166020830152831660408201526080810161444360608301846142ac565b95945050505050565b60008060008060008060008060e0898b03121561446857600080fd5b883567ffffffffffffffff81111561447f57600080fd5b61448b8b828c01614325565b909950975061449e905060208a01613ff4565b955060408901356144ae81613fbb565b9450606089013593506080890135925060a08901356144cc81613fbb565b915061440a60c08a0161436e565b60008060008060008060c087890312156144f357600080fd5b6144fc87613edd565b955061450a60208801613edd565b945061451860408801613edd565b9350606087013592506080870135915060a087013590509295509295509295565b60006020828403121561454b57600080fd5b8135612d2b81613fbb565b60006020828403121561456857600080fd5b5051919050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126145a457600080fd5b83018035915067ffffffffffffffff8211156145bf57600080fd5b60200191503681900382131561436757600080fd5b8183823760009101908152919050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261461957600080fd5b83018035915067ffffffffffffffff82111561463457600080fd5b6020019150600581901b360382131561436757600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561468e5761468e61464c565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826146d1576146d1614693565b500690565b6000826146e5576146e5614693565b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156147225761472261464c565b500290565b6000828210156147395761473961464c565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000808585111561477d57600080fd5b8386111561478a57600080fd5b5050820193919092039150565b600084516147a98184602089016141d7565b8201838582376000930192835250909392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156147f1576147f161464c565b5060010190565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600061018080835261485681840186886147f8565b915050825180516020840152602081015163ffffffff8082166040860152806040840151166060860152806060840151166080860152505060808101516148a560a085018263ffffffff169052565b5060a08101516bffffffffffffffffffffffff1660c08481019190915281015173ffffffffffffffffffffffffffffffffffffffff1660e0808501919091520151610100830152602083015160ff16610120830152604083015161014083015260609092015163ffffffff166101609091015292915050565b60006020828403121561493057600080fd5b8151612d2b81613fbb565b73ffffffffffffffffffffffffffffffffffffffff88811682528716602082015260ff8616604082015263ffffffff85811660608301528416608082015260c060a0820181905260009061306e90830184866147f8565b600063ffffffff8083168185168083038211156149b1576149b161464c565b01949350505050565b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee18336030181126149ee57600080fd5b9190910192915050565b600082601f830112614a0957600080fd5b614a11613f98565b806040840185811115614a2357600080fd5b845b81811015614a3d578035845260209384019301614a25565b509095945050505050565b600060808284031215614a5a57600080fd5b6040516040810181811067ffffffffffffffff82111715614a7d57614a7d613f1c565b604052614a8a84846149f8565b8152614a9984604085016149f8565b60208201529392505050565b600065ffffffffffff80831681851681830481118215151615614aca57614aca61464c565b02949350505050565b60008151602080840160005b83811015614afb57815187529582019590820190600101614adf565b509495945050505050565b6000612d2b8284614ad3565b6000614b1e8285614ad3565b9283525050602001919050565b600060408284031215614b3d57600080fd5b614b45613f98565b82358152602083013560208201528091505092915050565b80356020831015612d14577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602084900360031b1b1692915050565b600063ffffffff80841680614bb057614bb0614693565b92169190910692915050565b600063ffffffff80841680614bd357614bd3614693565b92169190910492915050565b600063ffffffff80831681851681830481118215151615614aca57614aca61464c565b600063ffffffff83811690831681811015614c1f57614c1f61464c565b03939250505056fe30303030303030303030303030313131313131313131313131313131313131313030303030303030303030303030302d5f2860204f2060295f2d202d5f2860206f2060295f2d202d5f286020512060295f2d2042495444414f204a5553542052454b5420594f55207c5f2860204f2060295f7c202d207c5f2860206f2060295f7c202d207c5f286020512060295f7c6161616161616161616161616161616161616161616161616161616161616161616161616161616161616161616161616161616161616161303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030a2646970667358221220e1b692fe4a926c17fee7b7c8ccc94ef518e80005f114fd4f0aa196e6966a2cfb64736f6c63430008090033",
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

// DataStoreIdToL2RollUpBlock is a free data retrieval call binding the contract method 0x92f30a45.
//
// Solidity: function dataStoreIdToL2RollUpBlock(uint32 ) view returns(uint256 startL2BlockNumber, uint256 endBL2BlockNumber, bool isReRollup)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) DataStoreIdToL2RollUpBlock(opts *bind.CallOpts, arg0 uint32) (struct {
	StartL2BlockNumber *big.Int
	EndBL2BlockNumber  *big.Int
	IsReRollup         bool
}, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "dataStoreIdToL2RollUpBlock", arg0)

	outstruct := new(struct {
		StartL2BlockNumber *big.Int
		EndBL2BlockNumber  *big.Int
		IsReRollup         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartL2BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndBL2BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsReRollup = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// DataStoreIdToL2RollUpBlock is a free data retrieval call binding the contract method 0x92f30a45.
//
// Solidity: function dataStoreIdToL2RollUpBlock(uint32 ) view returns(uint256 startL2BlockNumber, uint256 endBL2BlockNumber, bool isReRollup)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) DataStoreIdToL2RollUpBlock(arg0 uint32) (struct {
	StartL2BlockNumber *big.Int
	EndBL2BlockNumber  *big.Int
	IsReRollup         bool
}, error) {
	return _BVMEigenDataLayrChain.Contract.DataStoreIdToL2RollUpBlock(&_BVMEigenDataLayrChain.CallOpts, arg0)
}

// DataStoreIdToL2RollUpBlock is a free data retrieval call binding the contract method 0x92f30a45.
//
// Solidity: function dataStoreIdToL2RollUpBlock(uint32 ) view returns(uint256 startL2BlockNumber, uint256 endBL2BlockNumber, bool isReRollup)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) DataStoreIdToL2RollUpBlock(arg0 uint32) (struct {
	StartL2BlockNumber *big.Int
	EndBL2BlockNumber  *big.Int
	IsReRollup         bool
}, error) {
	return _BVMEigenDataLayrChain.Contract.DataStoreIdToL2RollUpBlock(&_BVMEigenDataLayrChain.CallOpts, arg0)
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

// GetL2ConfirmedBlockNumber is a free data retrieval call binding the contract method 0x8bea6cae.
//
// Solidity: function getL2ConfirmedBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) GetL2ConfirmedBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "getL2ConfirmedBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL2ConfirmedBlockNumber is a free data retrieval call binding the contract method 0x8bea6cae.
//
// Solidity: function getL2ConfirmedBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) GetL2ConfirmedBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.GetL2ConfirmedBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// GetL2ConfirmedBlockNumber is a free data retrieval call binding the contract method 0x8bea6cae.
//
// Solidity: function getL2ConfirmedBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) GetL2ConfirmedBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.GetL2ConfirmedBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// GetL2RollUpBlockByDataStoreId is a free data retrieval call binding the contract method 0xc96c0d38.
//
// Solidity: function getL2RollUpBlockByDataStoreId(uint32 _dataStoreId) view returns((uint256,uint256,bool))
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) GetL2RollUpBlockByDataStoreId(opts *bind.CallOpts, _dataStoreId uint32) (BVMEigenDataLayrChainBatchRollupBlock, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "getL2RollUpBlockByDataStoreId", _dataStoreId)

	if err != nil {
		return *new(BVMEigenDataLayrChainBatchRollupBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(BVMEigenDataLayrChainBatchRollupBlock)).(*BVMEigenDataLayrChainBatchRollupBlock)

	return out0, err

}

// GetL2RollUpBlockByDataStoreId is a free data retrieval call binding the contract method 0xc96c0d38.
//
// Solidity: function getL2RollUpBlockByDataStoreId(uint32 _dataStoreId) view returns((uint256,uint256,bool))
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) GetL2RollUpBlockByDataStoreId(_dataStoreId uint32) (BVMEigenDataLayrChainBatchRollupBlock, error) {
	return _BVMEigenDataLayrChain.Contract.GetL2RollUpBlockByDataStoreId(&_BVMEigenDataLayrChain.CallOpts, _dataStoreId)
}

// GetL2RollUpBlockByDataStoreId is a free data retrieval call binding the contract method 0xc96c0d38.
//
// Solidity: function getL2RollUpBlockByDataStoreId(uint32 _dataStoreId) view returns((uint256,uint256,bool))
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) GetL2RollUpBlockByDataStoreId(_dataStoreId uint32) (BVMEigenDataLayrChainBatchRollupBlock, error) {
	return _BVMEigenDataLayrChain.Contract.GetL2RollUpBlockByDataStoreId(&_BVMEigenDataLayrChain.CallOpts, _dataStoreId)
}

// GetL2StoredBlockNumber is a free data retrieval call binding the contract method 0x301b39ab.
//
// Solidity: function getL2StoredBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) GetL2StoredBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "getL2StoredBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL2StoredBlockNumber is a free data retrieval call binding the contract method 0x301b39ab.
//
// Solidity: function getL2StoredBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) GetL2StoredBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.GetL2StoredBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// GetL2StoredBlockNumber is a free data retrieval call binding the contract method 0x301b39ab.
//
// Solidity: function getL2StoredBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) GetL2StoredBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.GetL2StoredBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// GetRollupStoreByRollupBatchIndex is a free data retrieval call binding the contract method 0x2e72866b.
//
// Solidity: function getRollupStoreByRollupBatchIndex(uint256 _rollupBatchIndex) view returns((uint32,uint32,uint32,uint8))
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) GetRollupStoreByRollupBatchIndex(opts *bind.CallOpts, _rollupBatchIndex *big.Int) (BVMEigenDataLayrChainRollupStore, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "getRollupStoreByRollupBatchIndex", _rollupBatchIndex)

	if err != nil {
		return *new(BVMEigenDataLayrChainRollupStore), err
	}

	out0 := *abi.ConvertType(out[0], new(BVMEigenDataLayrChainRollupStore)).(*BVMEigenDataLayrChainRollupStore)

	return out0, err

}

// GetRollupStoreByRollupBatchIndex is a free data retrieval call binding the contract method 0x2e72866b.
//
// Solidity: function getRollupStoreByRollupBatchIndex(uint256 _rollupBatchIndex) view returns((uint32,uint32,uint32,uint8))
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) GetRollupStoreByRollupBatchIndex(_rollupBatchIndex *big.Int) (BVMEigenDataLayrChainRollupStore, error) {
	return _BVMEigenDataLayrChain.Contract.GetRollupStoreByRollupBatchIndex(&_BVMEigenDataLayrChain.CallOpts, _rollupBatchIndex)
}

// GetRollupStoreByRollupBatchIndex is a free data retrieval call binding the contract method 0x2e72866b.
//
// Solidity: function getRollupStoreByRollupBatchIndex(uint256 _rollupBatchIndex) view returns((uint32,uint32,uint32,uint8))
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) GetRollupStoreByRollupBatchIndex(_rollupBatchIndex *big.Int) (BVMEigenDataLayrChainRollupStore, error) {
	return _BVMEigenDataLayrChain.Contract.GetRollupStoreByRollupBatchIndex(&_BVMEigenDataLayrChain.CallOpts, _rollupBatchIndex)
}

// L2ConfirmedBlockNumber is a free data retrieval call binding the contract method 0x5d42ffb7.
//
// Solidity: function l2ConfirmedBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) L2ConfirmedBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "l2ConfirmedBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ConfirmedBlockNumber is a free data retrieval call binding the contract method 0x5d42ffb7.
//
// Solidity: function l2ConfirmedBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) L2ConfirmedBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.L2ConfirmedBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// L2ConfirmedBlockNumber is a free data retrieval call binding the contract method 0x5d42ffb7.
//
// Solidity: function l2ConfirmedBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) L2ConfirmedBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.L2ConfirmedBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// L2StoredBlockNumber is a free data retrieval call binding the contract method 0x990fca66.
//
// Solidity: function l2StoredBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) L2StoredBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "l2StoredBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2StoredBlockNumber is a free data retrieval call binding the contract method 0x990fca66.
//
// Solidity: function l2StoredBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) L2StoredBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.L2StoredBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
}

// L2StoredBlockNumber is a free data retrieval call binding the contract method 0x990fca66.
//
// Solidity: function l2StoredBlockNumber() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) L2StoredBlockNumber() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.L2StoredBlockNumber(&_BVMEigenDataLayrChain.CallOpts)
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

// Parse is a free data retrieval call binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) pure returns(bytes provenString)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) Parse(opts *bind.CallOpts, polys [][]byte, startIndex *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "parse", polys, startIndex, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Parse is a free data retrieval call binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) pure returns(bytes provenString)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) Parse(polys [][]byte, startIndex *big.Int, length *big.Int) ([]byte, error) {
	return _BVMEigenDataLayrChain.Contract.Parse(&_BVMEigenDataLayrChain.CallOpts, polys, startIndex, length)
}

// Parse is a free data retrieval call binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) pure returns(bytes provenString)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) Parse(polys [][]byte, startIndex *big.Int, length *big.Int) ([]byte, error) {
	return _BVMEigenDataLayrChain.Contract.Parse(&_BVMEigenDataLayrChain.CallOpts, polys, startIndex, length)
}

// ReRollupBatchIndex is a free data retrieval call binding the contract method 0xff2e0749.
//
// Solidity: function reRollupBatchIndex(uint256 ) view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) ReRollupBatchIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "reRollupBatchIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReRollupBatchIndex is a free data retrieval call binding the contract method 0xff2e0749.
//
// Solidity: function reRollupBatchIndex(uint256 ) view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) ReRollupBatchIndex(arg0 *big.Int) (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.ReRollupBatchIndex(&_BVMEigenDataLayrChain.CallOpts, arg0)
}

// ReRollupBatchIndex is a free data retrieval call binding the contract method 0xff2e0749.
//
// Solidity: function reRollupBatchIndex(uint256 ) view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) ReRollupBatchIndex(arg0 *big.Int) (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.ReRollupBatchIndex(&_BVMEigenDataLayrChain.CallOpts, arg0)
}

// ReRollupIndex is a free data retrieval call binding the contract method 0x927f2032.
//
// Solidity: function reRollupIndex() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) ReRollupIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "reRollupIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReRollupIndex is a free data retrieval call binding the contract method 0x927f2032.
//
// Solidity: function reRollupIndex() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) ReRollupIndex() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.ReRollupIndex(&_BVMEigenDataLayrChain.CallOpts)
}

// ReRollupIndex is a free data retrieval call binding the contract method 0x927f2032.
//
// Solidity: function reRollupIndex() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) ReRollupIndex() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.ReRollupIndex(&_BVMEigenDataLayrChain.CallOpts)
}

// ReSubmitterAddress is a free data retrieval call binding the contract method 0x758b8147.
//
// Solidity: function reSubmitterAddress() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) ReSubmitterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "reSubmitterAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReSubmitterAddress is a free data retrieval call binding the contract method 0x758b8147.
//
// Solidity: function reSubmitterAddress() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) ReSubmitterAddress() (common.Address, error) {
	return _BVMEigenDataLayrChain.Contract.ReSubmitterAddress(&_BVMEigenDataLayrChain.CallOpts)
}

// ReSubmitterAddress is a free data retrieval call binding the contract method 0x758b8147.
//
// Solidity: function reSubmitterAddress() view returns(address)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) ReSubmitterAddress() (common.Address, error) {
	return _BVMEigenDataLayrChain.Contract.ReSubmitterAddress(&_BVMEigenDataLayrChain.CallOpts)
}

// RollupBatchIndex is a free data retrieval call binding the contract method 0x3c762984.
//
// Solidity: function rollupBatchIndex() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) RollupBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "rollupBatchIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RollupBatchIndex is a free data retrieval call binding the contract method 0x3c762984.
//
// Solidity: function rollupBatchIndex() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) RollupBatchIndex() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.RollupBatchIndex(&_BVMEigenDataLayrChain.CallOpts)
}

// RollupBatchIndex is a free data retrieval call binding the contract method 0x3c762984.
//
// Solidity: function rollupBatchIndex() view returns(uint256)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) RollupBatchIndex() (*big.Int, error) {
	return _BVMEigenDataLayrChain.Contract.RollupBatchIndex(&_BVMEigenDataLayrChain.CallOpts)
}

// RollupBatchIndexRollupStores is a free data retrieval call binding the contract method 0x59cb6391.
//
// Solidity: function rollupBatchIndexRollupStores(uint256 ) view returns(uint32 originDataStoreId, uint32 dataStoreId, uint32 confirmAt, uint8 status)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCaller) RollupBatchIndexRollupStores(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OriginDataStoreId uint32
	DataStoreId       uint32
	ConfirmAt         uint32
	Status            uint8
}, error) {
	var out []interface{}
	err := _BVMEigenDataLayrChain.contract.Call(opts, &out, "rollupBatchIndexRollupStores", arg0)

	outstruct := new(struct {
		OriginDataStoreId uint32
		DataStoreId       uint32
		ConfirmAt         uint32
		Status            uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OriginDataStoreId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.DataStoreId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfirmAt = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// RollupBatchIndexRollupStores is a free data retrieval call binding the contract method 0x59cb6391.
//
// Solidity: function rollupBatchIndexRollupStores(uint256 ) view returns(uint32 originDataStoreId, uint32 dataStoreId, uint32 confirmAt, uint8 status)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) RollupBatchIndexRollupStores(arg0 *big.Int) (struct {
	OriginDataStoreId uint32
	DataStoreId       uint32
	ConfirmAt         uint32
	Status            uint8
}, error) {
	return _BVMEigenDataLayrChain.Contract.RollupBatchIndexRollupStores(&_BVMEigenDataLayrChain.CallOpts, arg0)
}

// RollupBatchIndexRollupStores is a free data retrieval call binding the contract method 0x59cb6391.
//
// Solidity: function rollupBatchIndexRollupStores(uint256 ) view returns(uint32 originDataStoreId, uint32 dataStoreId, uint32 confirmAt, uint8 status)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainCallerSession) RollupBatchIndexRollupStores(arg0 *big.Int) (struct {
	OriginDataStoreId uint32
	DataStoreId       uint32
	ConfirmAt         uint32
	Status            uint8
}, error) {
	return _BVMEigenDataLayrChain.Contract.RollupBatchIndexRollupStores(&_BVMEigenDataLayrChain.CallOpts, arg0)
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

// ConfirmData is a paid mutator transaction binding the contract method 0x4618ed87.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, uint256 startL2Block, uint256 endL2Block, uint32 originDataStoreId, uint256 reConfirmedBatchIndex, bool isReRollup) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) ConfirmData(opts *bind.TransactOpts, data []byte, searchData IDataLayrServiceManagerDataStoreSearchData, startL2Block *big.Int, endL2Block *big.Int, originDataStoreId uint32, reConfirmedBatchIndex *big.Int, isReRollup bool) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "confirmData", data, searchData, startL2Block, endL2Block, originDataStoreId, reConfirmedBatchIndex, isReRollup)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x4618ed87.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, uint256 startL2Block, uint256 endL2Block, uint32 originDataStoreId, uint256 reConfirmedBatchIndex, bool isReRollup) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) ConfirmData(data []byte, searchData IDataLayrServiceManagerDataStoreSearchData, startL2Block *big.Int, endL2Block *big.Int, originDataStoreId uint32, reConfirmedBatchIndex *big.Int, isReRollup bool) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.ConfirmData(&_BVMEigenDataLayrChain.TransactOpts, data, searchData, startL2Block, endL2Block, originDataStoreId, reConfirmedBatchIndex, isReRollup)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x4618ed87.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, uint256 startL2Block, uint256 endL2Block, uint32 originDataStoreId, uint256 reConfirmedBatchIndex, bool isReRollup) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) ConfirmData(data []byte, searchData IDataLayrServiceManagerDataStoreSearchData, startL2Block *big.Int, endL2Block *big.Int, originDataStoreId uint32, reConfirmedBatchIndex *big.Int, isReRollup bool) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.ConfirmData(&_BVMEigenDataLayrChain.TransactOpts, data, searchData, startL2Block, endL2Block, originDataStoreId, reConfirmedBatchIndex, isReRollup)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress, address _reSubmitterAddress, uint256 _block_stale_measure, uint256 _fraudProofPeriod, uint256 _l2SubmittedBlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) Initialize(opts *bind.TransactOpts, _sequencer common.Address, _dataManageAddress common.Address, _reSubmitterAddress common.Address, _block_stale_measure *big.Int, _fraudProofPeriod *big.Int, _l2SubmittedBlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "initialize", _sequencer, _dataManageAddress, _reSubmitterAddress, _block_stale_measure, _fraudProofPeriod, _l2SubmittedBlockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress, address _reSubmitterAddress, uint256 _block_stale_measure, uint256 _fraudProofPeriod, uint256 _l2SubmittedBlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) Initialize(_sequencer common.Address, _dataManageAddress common.Address, _reSubmitterAddress common.Address, _block_stale_measure *big.Int, _fraudProofPeriod *big.Int, _l2SubmittedBlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.Initialize(&_BVMEigenDataLayrChain.TransactOpts, _sequencer, _dataManageAddress, _reSubmitterAddress, _block_stale_measure, _fraudProofPeriod, _l2SubmittedBlockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress, address _reSubmitterAddress, uint256 _block_stale_measure, uint256 _fraudProofPeriod, uint256 _l2SubmittedBlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) Initialize(_sequencer common.Address, _dataManageAddress common.Address, _reSubmitterAddress common.Address, _block_stale_measure *big.Int, _fraudProofPeriod *big.Int, _l2SubmittedBlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.Initialize(&_BVMEigenDataLayrChain.TransactOpts, _sequencer, _dataManageAddress, _reSubmitterAddress, _block_stale_measure, _fraudProofPeriod, _l2SubmittedBlockNumber)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x15fda737.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) ProveFraud(opts *bind.TransactOpts, fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs BVMEigenDataLayrChainDisclosureProofs) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "proveFraud", fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x15fda737.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) ProveFraud(fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs BVMEigenDataLayrChainDisclosureProofs) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.ProveFraud(&_BVMEigenDataLayrChain.TransactOpts, fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x15fda737.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) ProveFraud(fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs BVMEigenDataLayrChainDisclosureProofs) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.ProveFraud(&_BVMEigenDataLayrChain.TransactOpts, fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
}

// RemoveFraudProofAddress is a paid mutator transaction binding the contract method 0x060ee9a4.
//
// Solidity: function removeFraudProofAddress(address _address) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) RemoveFraudProofAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "removeFraudProofAddress", _address)
}

// RemoveFraudProofAddress is a paid mutator transaction binding the contract method 0x060ee9a4.
//
// Solidity: function removeFraudProofAddress(address _address) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) RemoveFraudProofAddress(_address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.RemoveFraudProofAddress(&_BVMEigenDataLayrChain.TransactOpts, _address)
}

// RemoveFraudProofAddress is a paid mutator transaction binding the contract method 0x060ee9a4.
//
// Solidity: function removeFraudProofAddress(address _address) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) RemoveFraudProofAddress(_address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.RemoveFraudProofAddress(&_BVMEigenDataLayrChain.TransactOpts, _address)
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

// ResetRollupBatchData is a paid mutator transaction binding the contract method 0xf7db9795.
//
// Solidity: function resetRollupBatchData(uint256 _rollupBatchIndex) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) ResetRollupBatchData(opts *bind.TransactOpts, _rollupBatchIndex *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "resetRollupBatchData", _rollupBatchIndex)
}

// ResetRollupBatchData is a paid mutator transaction binding the contract method 0xf7db9795.
//
// Solidity: function resetRollupBatchData(uint256 _rollupBatchIndex) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) ResetRollupBatchData(_rollupBatchIndex *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.ResetRollupBatchData(&_BVMEigenDataLayrChain.TransactOpts, _rollupBatchIndex)
}

// ResetRollupBatchData is a paid mutator transaction binding the contract method 0xf7db9795.
//
// Solidity: function resetRollupBatchData(uint256 _rollupBatchIndex) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) ResetRollupBatchData(_rollupBatchIndex *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.ResetRollupBatchData(&_BVMEigenDataLayrChain.TransactOpts, _rollupBatchIndex)
}

// SetFraudProofAddress is a paid mutator transaction binding the contract method 0x32c58f7a.
//
// Solidity: function setFraudProofAddress(address _address) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) SetFraudProofAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "setFraudProofAddress", _address)
}

// SetFraudProofAddress is a paid mutator transaction binding the contract method 0x32c58f7a.
//
// Solidity: function setFraudProofAddress(address _address) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) SetFraudProofAddress(_address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.SetFraudProofAddress(&_BVMEigenDataLayrChain.TransactOpts, _address)
}

// SetFraudProofAddress is a paid mutator transaction binding the contract method 0x32c58f7a.
//
// Solidity: function setFraudProofAddress(address _address) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) SetFraudProofAddress(_address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.SetFraudProofAddress(&_BVMEigenDataLayrChain.TransactOpts, _address)
}

// StoreData is a paid mutator transaction binding the contract method 0x5e4a3056.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint256 startL2Block, uint256 endL2Block, uint32 totalOperatorsIndex, bool isReRollup) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) StoreData(opts *bind.TransactOpts, header []byte, duration uint8, blockNumber uint32, startL2Block *big.Int, endL2Block *big.Int, totalOperatorsIndex uint32, isReRollup bool) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "storeData", header, duration, blockNumber, startL2Block, endL2Block, totalOperatorsIndex, isReRollup)
}

// StoreData is a paid mutator transaction binding the contract method 0x5e4a3056.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint256 startL2Block, uint256 endL2Block, uint32 totalOperatorsIndex, bool isReRollup) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) StoreData(header []byte, duration uint8, blockNumber uint32, startL2Block *big.Int, endL2Block *big.Int, totalOperatorsIndex uint32, isReRollup bool) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.StoreData(&_BVMEigenDataLayrChain.TransactOpts, header, duration, blockNumber, startL2Block, endL2Block, totalOperatorsIndex, isReRollup)
}

// StoreData is a paid mutator transaction binding the contract method 0x5e4a3056.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint256 startL2Block, uint256 endL2Block, uint32 totalOperatorsIndex, bool isReRollup) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) StoreData(header []byte, duration uint8, blockNumber uint32, startL2Block *big.Int, endL2Block *big.Int, totalOperatorsIndex uint32, isReRollup bool) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.StoreData(&_BVMEigenDataLayrChain.TransactOpts, header, duration, blockNumber, startL2Block, endL2Block, totalOperatorsIndex, isReRollup)
}

// SubmitReRollUpInfo is a paid mutator transaction binding the contract method 0x9a71e29c.
//
// Solidity: function submitReRollUpInfo(uint256 batchIndex) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) SubmitReRollUpInfo(opts *bind.TransactOpts, batchIndex *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "submitReRollUpInfo", batchIndex)
}

// SubmitReRollUpInfo is a paid mutator transaction binding the contract method 0x9a71e29c.
//
// Solidity: function submitReRollUpInfo(uint256 batchIndex) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) SubmitReRollUpInfo(batchIndex *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.SubmitReRollUpInfo(&_BVMEigenDataLayrChain.TransactOpts, batchIndex)
}

// SubmitReRollUpInfo is a paid mutator transaction binding the contract method 0x9a71e29c.
//
// Solidity: function submitReRollUpInfo(uint256 batchIndex) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) SubmitReRollUpInfo(batchIndex *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.SubmitReRollUpInfo(&_BVMEigenDataLayrChain.TransactOpts, batchIndex)
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

// UnavailableFraudProofAddress is a paid mutator transaction binding the contract method 0x0a33202e.
//
// Solidity: function unavailableFraudProofAddress(address _address) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) UnavailableFraudProofAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "unavailableFraudProofAddress", _address)
}

// UnavailableFraudProofAddress is a paid mutator transaction binding the contract method 0x0a33202e.
//
// Solidity: function unavailableFraudProofAddress(address _address) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) UnavailableFraudProofAddress(_address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UnavailableFraudProofAddress(&_BVMEigenDataLayrChain.TransactOpts, _address)
}

// UnavailableFraudProofAddress is a paid mutator transaction binding the contract method 0x0a33202e.
//
// Solidity: function unavailableFraudProofAddress(address _address) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) UnavailableFraudProofAddress(_address common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UnavailableFraudProofAddress(&_BVMEigenDataLayrChain.TransactOpts, _address)
}

// UpdateDataLayrManagerAddress is a paid mutator transaction binding the contract method 0x02d777de.
//
// Solidity: function updateDataLayrManagerAddress(address _dataManageAddress) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) UpdateDataLayrManagerAddress(opts *bind.TransactOpts, _dataManageAddress common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "updateDataLayrManagerAddress", _dataManageAddress)
}

// UpdateDataLayrManagerAddress is a paid mutator transaction binding the contract method 0x02d777de.
//
// Solidity: function updateDataLayrManagerAddress(address _dataManageAddress) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) UpdateDataLayrManagerAddress(_dataManageAddress common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateDataLayrManagerAddress(&_BVMEigenDataLayrChain.TransactOpts, _dataManageAddress)
}

// UpdateDataLayrManagerAddress is a paid mutator transaction binding the contract method 0x02d777de.
//
// Solidity: function updateDataLayrManagerAddress(address _dataManageAddress) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) UpdateDataLayrManagerAddress(_dataManageAddress common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateDataLayrManagerAddress(&_BVMEigenDataLayrChain.TransactOpts, _dataManageAddress)
}

// UpdateFraudProofPeriod is a paid mutator transaction binding the contract method 0xd7fbc2e2.
//
// Solidity: function updateFraudProofPeriod(uint256 _fraudProofPeriod) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) UpdateFraudProofPeriod(opts *bind.TransactOpts, _fraudProofPeriod *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "updateFraudProofPeriod", _fraudProofPeriod)
}

// UpdateFraudProofPeriod is a paid mutator transaction binding the contract method 0xd7fbc2e2.
//
// Solidity: function updateFraudProofPeriod(uint256 _fraudProofPeriod) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) UpdateFraudProofPeriod(_fraudProofPeriod *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateFraudProofPeriod(&_BVMEigenDataLayrChain.TransactOpts, _fraudProofPeriod)
}

// UpdateFraudProofPeriod is a paid mutator transaction binding the contract method 0xd7fbc2e2.
//
// Solidity: function updateFraudProofPeriod(uint256 _fraudProofPeriod) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) UpdateFraudProofPeriod(_fraudProofPeriod *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateFraudProofPeriod(&_BVMEigenDataLayrChain.TransactOpts, _fraudProofPeriod)
}

// UpdateL2ConfirmedBlockNumber is a paid mutator transaction binding the contract method 0x2e64b4c0.
//
// Solidity: function updateL2ConfirmedBlockNumber(uint256 _l2ConfirmedBlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) UpdateL2ConfirmedBlockNumber(opts *bind.TransactOpts, _l2ConfirmedBlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "updateL2ConfirmedBlockNumber", _l2ConfirmedBlockNumber)
}

// UpdateL2ConfirmedBlockNumber is a paid mutator transaction binding the contract method 0x2e64b4c0.
//
// Solidity: function updateL2ConfirmedBlockNumber(uint256 _l2ConfirmedBlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) UpdateL2ConfirmedBlockNumber(_l2ConfirmedBlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateL2ConfirmedBlockNumber(&_BVMEigenDataLayrChain.TransactOpts, _l2ConfirmedBlockNumber)
}

// UpdateL2ConfirmedBlockNumber is a paid mutator transaction binding the contract method 0x2e64b4c0.
//
// Solidity: function updateL2ConfirmedBlockNumber(uint256 _l2ConfirmedBlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) UpdateL2ConfirmedBlockNumber(_l2ConfirmedBlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateL2ConfirmedBlockNumber(&_BVMEigenDataLayrChain.TransactOpts, _l2ConfirmedBlockNumber)
}

// UpdateL2StoredBlockNumber is a paid mutator transaction binding the contract method 0x9495de40.
//
// Solidity: function updateL2StoredBlockNumber(uint256 _l2StoredBlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) UpdateL2StoredBlockNumber(opts *bind.TransactOpts, _l2StoredBlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "updateL2StoredBlockNumber", _l2StoredBlockNumber)
}

// UpdateL2StoredBlockNumber is a paid mutator transaction binding the contract method 0x9495de40.
//
// Solidity: function updateL2StoredBlockNumber(uint256 _l2StoredBlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) UpdateL2StoredBlockNumber(_l2StoredBlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateL2StoredBlockNumber(&_BVMEigenDataLayrChain.TransactOpts, _l2StoredBlockNumber)
}

// UpdateL2StoredBlockNumber is a paid mutator transaction binding the contract method 0x9495de40.
//
// Solidity: function updateL2StoredBlockNumber(uint256 _l2StoredBlockNumber) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) UpdateL2StoredBlockNumber(_l2StoredBlockNumber *big.Int) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateL2StoredBlockNumber(&_BVMEigenDataLayrChain.TransactOpts, _l2StoredBlockNumber)
}

// UpdateReSubmitterAddress is a paid mutator transaction binding the contract method 0xafab4ac5.
//
// Solidity: function updateReSubmitterAddress(address _reSubmitterAddress) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) UpdateReSubmitterAddress(opts *bind.TransactOpts, _reSubmitterAddress common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "updateReSubmitterAddress", _reSubmitterAddress)
}

// UpdateReSubmitterAddress is a paid mutator transaction binding the contract method 0xafab4ac5.
//
// Solidity: function updateReSubmitterAddress(address _reSubmitterAddress) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) UpdateReSubmitterAddress(_reSubmitterAddress common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateReSubmitterAddress(&_BVMEigenDataLayrChain.TransactOpts, _reSubmitterAddress)
}

// UpdateReSubmitterAddress is a paid mutator transaction binding the contract method 0xafab4ac5.
//
// Solidity: function updateReSubmitterAddress(address _reSubmitterAddress) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) UpdateReSubmitterAddress(_reSubmitterAddress common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateReSubmitterAddress(&_BVMEigenDataLayrChain.TransactOpts, _reSubmitterAddress)
}

// UpdateSequencerAddress is a paid mutator transaction binding the contract method 0xc8fff01b.
//
// Solidity: function updateSequencerAddress(address _sequencer) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactor) UpdateSequencerAddress(opts *bind.TransactOpts, _sequencer common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.contract.Transact(opts, "updateSequencerAddress", _sequencer)
}

// UpdateSequencerAddress is a paid mutator transaction binding the contract method 0xc8fff01b.
//
// Solidity: function updateSequencerAddress(address _sequencer) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainSession) UpdateSequencerAddress(_sequencer common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateSequencerAddress(&_BVMEigenDataLayrChain.TransactOpts, _sequencer)
}

// UpdateSequencerAddress is a paid mutator transaction binding the contract method 0xc8fff01b.
//
// Solidity: function updateSequencerAddress(address _sequencer) returns()
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainTransactorSession) UpdateSequencerAddress(_sequencer common.Address) (*types.Transaction, error) {
	return _BVMEigenDataLayrChain.Contract.UpdateSequencerAddress(&_BVMEigenDataLayrChain.TransactOpts, _sequencer)
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

// BVMEigenDataLayrChainReRollupBatchDataIterator is returned from FilterReRollupBatchData and is used to iterate over the raw logs and unpacked data for ReRollupBatchData events raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainReRollupBatchDataIterator struct {
	Event *BVMEigenDataLayrChainReRollupBatchData // Event containing the contract specifics and raw log

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
func (it *BVMEigenDataLayrChainReRollupBatchDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BVMEigenDataLayrChainReRollupBatchData)
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
		it.Event = new(BVMEigenDataLayrChainReRollupBatchData)
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
func (it *BVMEigenDataLayrChainReRollupBatchDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BVMEigenDataLayrChainReRollupBatchDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BVMEigenDataLayrChainReRollupBatchData represents a ReRollupBatchData event raised by the BVMEigenDataLayrChain contract.
type BVMEigenDataLayrChainReRollupBatchData struct {
	ReRollupIndex      *big.Int
	RollupBatchIndex   *big.Int
	StratL2BlockNumber *big.Int
	EndL2BlockNumber   *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReRollupBatchData is a free log retrieval operation binding the contract event 0xee84ab0752d66e31e484f6855689d7067ecd900a6c5a198a2908f74e583e7d57.
//
// Solidity: event ReRollupBatchData(uint256 reRollupIndex, uint256 rollupBatchIndex, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) FilterReRollupBatchData(opts *bind.FilterOpts) (*BVMEigenDataLayrChainReRollupBatchDataIterator, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.FilterLogs(opts, "ReRollupBatchData")
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainReRollupBatchDataIterator{contract: _BVMEigenDataLayrChain.contract, event: "ReRollupBatchData", logs: logs, sub: sub}, nil
}

// WatchReRollupBatchData is a free log subscription operation binding the contract event 0xee84ab0752d66e31e484f6855689d7067ecd900a6c5a198a2908f74e583e7d57.
//
// Solidity: event ReRollupBatchData(uint256 reRollupIndex, uint256 rollupBatchIndex, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) WatchReRollupBatchData(opts *bind.WatchOpts, sink chan<- *BVMEigenDataLayrChainReRollupBatchData) (event.Subscription, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.WatchLogs(opts, "ReRollupBatchData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BVMEigenDataLayrChainReRollupBatchData)
				if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "ReRollupBatchData", log); err != nil {
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

// ParseReRollupBatchData is a log parse operation binding the contract event 0xee84ab0752d66e31e484f6855689d7067ecd900a6c5a198a2908f74e583e7d57.
//
// Solidity: event ReRollupBatchData(uint256 reRollupIndex, uint256 rollupBatchIndex, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) ParseReRollupBatchData(log types.Log) (*BVMEigenDataLayrChainReRollupBatchData, error) {
	event := new(BVMEigenDataLayrChainReRollupBatchData)
	if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "ReRollupBatchData", log); err != nil {
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
	RollupBatchIndex   *big.Int
	DataStoreId        uint32
	StratL2BlockNumber *big.Int
	EndL2BlockNumber   *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreConfirmed is a free log retrieval operation binding the contract event 0xc7c0900be05d2a0ad0f77852eb975d9e862d1db0a2238617dd0f77854782f672.
//
// Solidity: event RollupStoreConfirmed(uint256 rollupBatchIndex, uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) FilterRollupStoreConfirmed(opts *bind.FilterOpts) (*BVMEigenDataLayrChainRollupStoreConfirmedIterator, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.FilterLogs(opts, "RollupStoreConfirmed")
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainRollupStoreConfirmedIterator{contract: _BVMEigenDataLayrChain.contract, event: "RollupStoreConfirmed", logs: logs, sub: sub}, nil
}

// WatchRollupStoreConfirmed is a free log subscription operation binding the contract event 0xc7c0900be05d2a0ad0f77852eb975d9e862d1db0a2238617dd0f77854782f672.
//
// Solidity: event RollupStoreConfirmed(uint256 rollupBatchIndex, uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
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

// ParseRollupStoreConfirmed is a log parse operation binding the contract event 0xc7c0900be05d2a0ad0f77852eb975d9e862d1db0a2238617dd0f77854782f672.
//
// Solidity: event RollupStoreConfirmed(uint256 rollupBatchIndex, uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
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
	DataStoreId        uint32
	StratL2BlockNumber *big.Int
	EndL2BlockNumber   *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreInitialized is a free log retrieval operation binding the contract event 0xa99ca06ac3461399088feac88ec48dc5a47d61c3b6839eab20146f2c4ee53584.
//
// Solidity: event RollupStoreInitialized(uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) FilterRollupStoreInitialized(opts *bind.FilterOpts) (*BVMEigenDataLayrChainRollupStoreInitializedIterator, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.FilterLogs(opts, "RollupStoreInitialized")
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainRollupStoreInitializedIterator{contract: _BVMEigenDataLayrChain.contract, event: "RollupStoreInitialized", logs: logs, sub: sub}, nil
}

// WatchRollupStoreInitialized is a free log subscription operation binding the contract event 0xa99ca06ac3461399088feac88ec48dc5a47d61c3b6839eab20146f2c4ee53584.
//
// Solidity: event RollupStoreInitialized(uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
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

// ParseRollupStoreInitialized is a log parse operation binding the contract event 0xa99ca06ac3461399088feac88ec48dc5a47d61c3b6839eab20146f2c4ee53584.
//
// Solidity: event RollupStoreInitialized(uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
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
	RollupBatchIndex   *big.Int
	DataStoreId        uint32
	StratL2BlockNumber *big.Int
	EndL2BlockNumber   *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreReverted is a free log retrieval operation binding the contract event 0xca227c67a02028763083580d42e8bdef4bb49c393068d05983421cd7a4a2a5be.
//
// Solidity: event RollupStoreReverted(uint256 rollupBatchIndex, uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) FilterRollupStoreReverted(opts *bind.FilterOpts) (*BVMEigenDataLayrChainRollupStoreRevertedIterator, error) {

	logs, sub, err := _BVMEigenDataLayrChain.contract.FilterLogs(opts, "RollupStoreReverted")
	if err != nil {
		return nil, err
	}
	return &BVMEigenDataLayrChainRollupStoreRevertedIterator{contract: _BVMEigenDataLayrChain.contract, event: "RollupStoreReverted", logs: logs, sub: sub}, nil
}

// WatchRollupStoreReverted is a free log subscription operation binding the contract event 0xca227c67a02028763083580d42e8bdef4bb49c393068d05983421cd7a4a2a5be.
//
// Solidity: event RollupStoreReverted(uint256 rollupBatchIndex, uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
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

// ParseRollupStoreReverted is a log parse operation binding the contract event 0xca227c67a02028763083580d42e8bdef4bb49c393068d05983421cd7a4a2a5be.
//
// Solidity: event RollupStoreReverted(uint256 rollupBatchIndex, uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber)
func (_BVMEigenDataLayrChain *BVMEigenDataLayrChainFilterer) ParseRollupStoreReverted(log types.Log) (*BVMEigenDataLayrChainRollupStoreReverted, error) {
	event := new(BVMEigenDataLayrChainRollupStoreReverted)
	if err := _BVMEigenDataLayrChain.contract.UnpackLog(event, "RollupStoreReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

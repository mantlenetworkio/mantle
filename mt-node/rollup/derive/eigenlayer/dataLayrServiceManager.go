package eigenlayer

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

// BLSSignatureCheckerSignatoryTotals is an auto generated low-level Go binding around an user-defined struct.
type BLSSignatureCheckerSignatoryTotals struct {
	SignedStakeFirstQuorum  *big.Int
	SignedStakeSecondQuorum *big.Int
	TotalStakeFirstQuorum   *big.Int
	TotalStakeSecondQuorum  *big.Int
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

// DataLayrServiceManagerMetaData contains all meta data concerning the DataLayrServiceManager contract.
var DataLayrServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIInvestmentManager\",\"name\":\"_investmentManager\",\"type\":\"address\"},{\"internalType\":\"contractIEigenLayrDelegation\",\"name\":\"_eigenLayrDelegation\",\"type\":\"address\"},{\"internalType\":\"contractIRepository\",\"name\":\"_repository\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feePerBytePerTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"adversaryThresholdBasisPoints\",\"type\":\"uint16\"}],\"name\":\"AdversaryThresholdBasisPointsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"BombVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmDataStore\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"EphemeralKeyRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"FeePerBytePerTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"FirstQuorumThresholdPercentageSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feePayer\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIDataLayrServiceManager.DataStoreSearchData\",\"name\":\"searchData\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"InitDataStore\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PaymentManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"quorumThresholdBasisPoints\",\"type\":\"uint16\"}],\"name\":\"QuorumThresholdBasisPointsUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"SecondQuorumThresholdPercentageSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"taskNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"signedStakeFirstQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"signedStakeSecondQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"pubkeyHashes\",\"type\":\"bytes32[]\"}],\"name\":\"SignatoryRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"log_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"log_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"log_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"log_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"log_named_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"log_named_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"log_named_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"}],\"name\":\"log_named_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"log_named_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"log_named_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"log_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"logs\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIP_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION_SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IS_TEST\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DATASTORE_DURATION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DATASTORE_DURATION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUM_DS_PER_BLOCK_PER_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adversaryThresholdBasisPoints\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feePerBytePerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"storePeriodLength\",\"type\":\"uint32\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"taskNumberToConfirm\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stakesBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"signedStakeFirstQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signedStakeSecondQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStakeFirstQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStakeSecondQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structBLSSignatureChecker.SignatoryTotals\",\"name\":\"signedTotals\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"compressedSignatoryRecord\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreSearchData\",\"name\":\"searchData\",\"type\":\"tuple\"}],\"name\":\"confirmDataStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataLayrBombVerifier\",\"outputs\":[{\"internalType\":\"contractDataLayrBombVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataLayrLowDegreeChallenge\",\"outputs\":[{\"internalType\":\"contractDataLayrLowDegreeChallenge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataLayrPaymentManager\",\"outputs\":[{\"internalType\":\"contractIDataLayrPaymentManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dataStoreHashesForDurationAtTimestamp\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"dataStoreIdToSignatureHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataStoresForDuration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"one_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"two_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"three_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"four_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"five_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"six_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"seven_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"latestTime\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenLayrDelegation\",\"outputs\":[{\"internalType\":\"contractIEigenLayrDelegation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ephemeralKeyRegistry\",\"outputs\":[{\"internalType\":\"contractEphemeralKeyRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePerBytePerTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstQuorumThresholdPercentage\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"freezeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getDataStoreHashesForDurationAtTimestamp\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"}],\"name\":\"getNumDataStoresForDuration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feePayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalOperatorsIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"initDataStore\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"investmentManager\",\"outputs\":[{\"internalType\":\"contractIInvestmentManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTime\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"log2NumPowersOfTau\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numPowersOfTau\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumThresholdBasisPoints\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"repository\",\"outputs\":[{\"internalType\":\"contractIRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"unbondedAfter\",\"type\":\"uint32\"}],\"name\":\"revokeSlashingAbility\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondQuorumThresholdPercentage\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_adversaryThresholdBasisPoints\",\"type\":\"uint16\"}],\"name\":\"setAdversaryThresholdBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataLayrBombVerifier\",\"name\":\"_dataLayrBombVerifier\",\"type\":\"address\"}],\"name\":\"setBombVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEphemeralKeyRegistry\",\"name\":\"_ephemeralKeyRegistry\",\"type\":\"address\"}],\"name\":\"setEphemeralKeyRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePerBytePerTime\",\"type\":\"uint256\"}],\"name\":\"setFeePerBytePerTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_firstQuorumThresholdPercentage\",\"type\":\"uint128\"}],\"name\":\"setFirstQuorumThresholdPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataLayrLowDegreeChallenge\",\"name\":\"_dataLayrLowDegreeChallenge\",\"type\":\"address\"}],\"name\":\"setLowDegreeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataLayrPaymentManager\",\"name\":\"_dataLayrPaymentManager\",\"type\":\"address\"}],\"name\":\"setPaymentManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_quorumThresholdBasisPoints\",\"type\":\"uint16\"}],\"name\":\"setQuorumThresholdBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_secondQuorumThresholdPercentage\",\"type\":\"uint128\"}],\"name\":\"setSecondQuorumThresholdPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"packedDataStoreSearchData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"initTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"name\":\"stakeWithdrawalVerification\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"name\":\"verifyDataStoreMetadata\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"zeroPolynomialCommitmentMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DataLayrServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DataLayrServiceManagerMetaData.ABI instead.
var DataLayrServiceManagerABI = DataLayrServiceManagerMetaData.ABI

// DataLayrServiceManager is an auto generated Go binding around an Ethereum contract.
type DataLayrServiceManager struct {
	DataLayrServiceManagerCaller     // Read-only binding to the contract
	DataLayrServiceManagerTransactor // Write-only binding to the contract
	DataLayrServiceManagerFilterer   // Log filterer for contract events
}

// DataLayrServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataLayrServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLayrServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataLayrServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLayrServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataLayrServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLayrServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataLayrServiceManagerSession struct {
	Contract     *DataLayrServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DataLayrServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataLayrServiceManagerCallerSession struct {
	Contract *DataLayrServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DataLayrServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataLayrServiceManagerTransactorSession struct {
	Contract     *DataLayrServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DataLayrServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataLayrServiceManagerRaw struct {
	Contract *DataLayrServiceManager // Generic contract binding to access the raw methods on
}

// DataLayrServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataLayrServiceManagerCallerRaw struct {
	Contract *DataLayrServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DataLayrServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataLayrServiceManagerTransactorRaw struct {
	Contract *DataLayrServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataLayrServiceManager creates a new instance of DataLayrServiceManager, bound to a specific deployed contract.
func NewDataLayrServiceManager(address common.Address, backend bind.ContractBackend) (*DataLayrServiceManager, error) {
	contract, err := bindDataLayrServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManager{DataLayrServiceManagerCaller: DataLayrServiceManagerCaller{contract: contract}, DataLayrServiceManagerTransactor: DataLayrServiceManagerTransactor{contract: contract}, DataLayrServiceManagerFilterer: DataLayrServiceManagerFilterer{contract: contract}}, nil
}

// NewDataLayrServiceManagerCaller creates a new read-only instance of DataLayrServiceManager, bound to a specific deployed contract.
func NewDataLayrServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*DataLayrServiceManagerCaller, error) {
	contract, err := bindDataLayrServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerCaller{contract: contract}, nil
}

// NewDataLayrServiceManagerTransactor creates a new write-only instance of DataLayrServiceManager, bound to a specific deployed contract.
func NewDataLayrServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DataLayrServiceManagerTransactor, error) {
	contract, err := bindDataLayrServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerTransactor{contract: contract}, nil
}

// NewDataLayrServiceManagerFilterer creates a new log filterer instance of DataLayrServiceManager, bound to a specific deployed contract.
func NewDataLayrServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DataLayrServiceManagerFilterer, error) {
	contract, err := bindDataLayrServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerFilterer{contract: contract}, nil
}

// bindDataLayrServiceManager binds a generic wrapper to an already deployed contract.
func bindDataLayrServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DataLayrServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLayrServiceManager *DataLayrServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataLayrServiceManager.Contract.DataLayrServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLayrServiceManager *DataLayrServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.DataLayrServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLayrServiceManager *DataLayrServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.DataLayrServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLayrServiceManager *DataLayrServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataLayrServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.contract.Transact(opts, method, params...)
}

// BIPMULTIPLIER is a free data retrieval call binding the contract method 0xa3c7eaf0.
//
// Solidity: function BIP_MULTIPLIER() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) BIPMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "BIP_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BIPMULTIPLIER is a free data retrieval call binding the contract method 0xa3c7eaf0.
//
// Solidity: function BIP_MULTIPLIER() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) BIPMULTIPLIER() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.BIPMULTIPLIER(&_DataLayrServiceManager.CallOpts)
}

// BIPMULTIPLIER is a free data retrieval call binding the contract method 0xa3c7eaf0.
//
// Solidity: function BIP_MULTIPLIER() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) BIPMULTIPLIER() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.BIPMULTIPLIER(&_DataLayrServiceManager.CallOpts)
}

// DURATIONSCALE is a free data retrieval call binding the contract method 0x31a219c5.
//
// Solidity: function DURATION_SCALE() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) DURATIONSCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "DURATION_SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATIONSCALE is a free data retrieval call binding the contract method 0x31a219c5.
//
// Solidity: function DURATION_SCALE() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) DURATIONSCALE() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.DURATIONSCALE(&_DataLayrServiceManager.CallOpts)
}

// DURATIONSCALE is a free data retrieval call binding the contract method 0x31a219c5.
//
// Solidity: function DURATION_SCALE() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) DURATIONSCALE() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.DURATIONSCALE(&_DataLayrServiceManager.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) ISTEST() (bool, error) {
	return _DataLayrServiceManager.Contract.ISTEST(&_DataLayrServiceManager.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) ISTEST() (bool, error) {
	return _DataLayrServiceManager.Contract.ISTEST(&_DataLayrServiceManager.CallOpts)
}

// MAXDATASTOREDURATION is a free data retrieval call binding the contract method 0x578ae5a1.
//
// Solidity: function MAX_DATASTORE_DURATION() view returns(uint8)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) MAXDATASTOREDURATION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "MAX_DATASTORE_DURATION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXDATASTOREDURATION is a free data retrieval call binding the contract method 0x578ae5a1.
//
// Solidity: function MAX_DATASTORE_DURATION() view returns(uint8)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) MAXDATASTOREDURATION() (uint8, error) {
	return _DataLayrServiceManager.Contract.MAXDATASTOREDURATION(&_DataLayrServiceManager.CallOpts)
}

// MAXDATASTOREDURATION is a free data retrieval call binding the contract method 0x578ae5a1.
//
// Solidity: function MAX_DATASTORE_DURATION() view returns(uint8)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) MAXDATASTOREDURATION() (uint8, error) {
	return _DataLayrServiceManager.Contract.MAXDATASTOREDURATION(&_DataLayrServiceManager.CallOpts)
}

// MINDATASTOREDURATION is a free data retrieval call binding the contract method 0x1fdab6e4.
//
// Solidity: function MIN_DATASTORE_DURATION() view returns(uint8)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) MINDATASTOREDURATION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "MIN_DATASTORE_DURATION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MINDATASTOREDURATION is a free data retrieval call binding the contract method 0x1fdab6e4.
//
// Solidity: function MIN_DATASTORE_DURATION() view returns(uint8)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) MINDATASTOREDURATION() (uint8, error) {
	return _DataLayrServiceManager.Contract.MINDATASTOREDURATION(&_DataLayrServiceManager.CallOpts)
}

// MINDATASTOREDURATION is a free data retrieval call binding the contract method 0x1fdab6e4.
//
// Solidity: function MIN_DATASTORE_DURATION() view returns(uint8)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) MINDATASTOREDURATION() (uint8, error) {
	return _DataLayrServiceManager.Contract.MINDATASTOREDURATION(&_DataLayrServiceManager.CallOpts)
}

// NUMDSPERBLOCKPERDURATION is a free data retrieval call binding the contract method 0x5f87abbb.
//
// Solidity: function NUM_DS_PER_BLOCK_PER_DURATION() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) NUMDSPERBLOCKPERDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "NUM_DS_PER_BLOCK_PER_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NUMDSPERBLOCKPERDURATION is a free data retrieval call binding the contract method 0x5f87abbb.
//
// Solidity: function NUM_DS_PER_BLOCK_PER_DURATION() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) NUMDSPERBLOCKPERDURATION() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.NUMDSPERBLOCKPERDURATION(&_DataLayrServiceManager.CallOpts)
}

// NUMDSPERBLOCKPERDURATION is a free data retrieval call binding the contract method 0x5f87abbb.
//
// Solidity: function NUM_DS_PER_BLOCK_PER_DURATION() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) NUMDSPERBLOCKPERDURATION() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.NUMDSPERBLOCKPERDURATION(&_DataLayrServiceManager.CallOpts)
}

// AdversaryThresholdBasisPoints is a free data retrieval call binding the contract method 0x3594b60f.
//
// Solidity: function adversaryThresholdBasisPoints() view returns(uint16)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) AdversaryThresholdBasisPoints(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "adversaryThresholdBasisPoints")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// AdversaryThresholdBasisPoints is a free data retrieval call binding the contract method 0x3594b60f.
//
// Solidity: function adversaryThresholdBasisPoints() view returns(uint16)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) AdversaryThresholdBasisPoints() (uint16, error) {
	return _DataLayrServiceManager.Contract.AdversaryThresholdBasisPoints(&_DataLayrServiceManager.CallOpts)
}

// AdversaryThresholdBasisPoints is a free data retrieval call binding the contract method 0x3594b60f.
//
// Solidity: function adversaryThresholdBasisPoints() view returns(uint16)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) AdversaryThresholdBasisPoints() (uint16, error) {
	return _DataLayrServiceManager.Contract.AdversaryThresholdBasisPoints(&_DataLayrServiceManager.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0x39fe2e71.
//
// Solidity: function calculateFee(uint256 totalBytes, uint256 _feePerBytePerTime, uint32 storePeriodLength) pure returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) CalculateFee(opts *bind.CallOpts, totalBytes *big.Int, _feePerBytePerTime *big.Int, storePeriodLength uint32) (*big.Int, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "calculateFee", totalBytes, _feePerBytePerTime, storePeriodLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x39fe2e71.
//
// Solidity: function calculateFee(uint256 totalBytes, uint256 _feePerBytePerTime, uint32 storePeriodLength) pure returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) CalculateFee(totalBytes *big.Int, _feePerBytePerTime *big.Int, storePeriodLength uint32) (*big.Int, error) {
	return _DataLayrServiceManager.Contract.CalculateFee(&_DataLayrServiceManager.CallOpts, totalBytes, _feePerBytePerTime, storePeriodLength)
}

// CalculateFee is a free data retrieval call binding the contract method 0x39fe2e71.
//
// Solidity: function calculateFee(uint256 totalBytes, uint256 _feePerBytePerTime, uint32 storePeriodLength) pure returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) CalculateFee(totalBytes *big.Int, _feePerBytePerTime *big.Int, storePeriodLength uint32) (*big.Int, error) {
	return _DataLayrServiceManager.Contract.CalculateFee(&_DataLayrServiceManager.CallOpts, totalBytes, _feePerBytePerTime, storePeriodLength)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) CollateralToken() (common.Address, error) {
	return _DataLayrServiceManager.Contract.CollateralToken(&_DataLayrServiceManager.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) CollateralToken() (common.Address, error) {
	return _DataLayrServiceManager.Contract.CollateralToken(&_DataLayrServiceManager.CallOpts)
}

// DataLayrBombVerifier is a free data retrieval call binding the contract method 0x5e69019c.
//
// Solidity: function dataLayrBombVerifier() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) DataLayrBombVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "dataLayrBombVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataLayrBombVerifier is a free data retrieval call binding the contract method 0x5e69019c.
//
// Solidity: function dataLayrBombVerifier() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) DataLayrBombVerifier() (common.Address, error) {
	return _DataLayrServiceManager.Contract.DataLayrBombVerifier(&_DataLayrServiceManager.CallOpts)
}

// DataLayrBombVerifier is a free data retrieval call binding the contract method 0x5e69019c.
//
// Solidity: function dataLayrBombVerifier() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) DataLayrBombVerifier() (common.Address, error) {
	return _DataLayrServiceManager.Contract.DataLayrBombVerifier(&_DataLayrServiceManager.CallOpts)
}

// DataLayrLowDegreeChallenge is a free data retrieval call binding the contract method 0xdc145394.
//
// Solidity: function dataLayrLowDegreeChallenge() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) DataLayrLowDegreeChallenge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "dataLayrLowDegreeChallenge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataLayrLowDegreeChallenge is a free data retrieval call binding the contract method 0xdc145394.
//
// Solidity: function dataLayrLowDegreeChallenge() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) DataLayrLowDegreeChallenge() (common.Address, error) {
	return _DataLayrServiceManager.Contract.DataLayrLowDegreeChallenge(&_DataLayrServiceManager.CallOpts)
}

// DataLayrLowDegreeChallenge is a free data retrieval call binding the contract method 0xdc145394.
//
// Solidity: function dataLayrLowDegreeChallenge() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) DataLayrLowDegreeChallenge() (common.Address, error) {
	return _DataLayrServiceManager.Contract.DataLayrLowDegreeChallenge(&_DataLayrServiceManager.CallOpts)
}

// DataLayrPaymentManager is a free data retrieval call binding the contract method 0x0b6a8390.
//
// Solidity: function dataLayrPaymentManager() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) DataLayrPaymentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "dataLayrPaymentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataLayrPaymentManager is a free data retrieval call binding the contract method 0x0b6a8390.
//
// Solidity: function dataLayrPaymentManager() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) DataLayrPaymentManager() (common.Address, error) {
	return _DataLayrServiceManager.Contract.DataLayrPaymentManager(&_DataLayrServiceManager.CallOpts)
}

// DataLayrPaymentManager is a free data retrieval call binding the contract method 0x0b6a8390.
//
// Solidity: function dataLayrPaymentManager() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) DataLayrPaymentManager() (common.Address, error) {
	return _DataLayrServiceManager.Contract.DataLayrPaymentManager(&_DataLayrServiceManager.CallOpts)
}

// DataStoreHashesForDurationAtTimestamp is a free data retrieval call binding the contract method 0x1bd2b3cf.
//
// Solidity: function dataStoreHashesForDurationAtTimestamp(uint8 , uint256 , uint256 ) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) DataStoreHashesForDurationAtTimestamp(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int, arg2 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "dataStoreHashesForDurationAtTimestamp", arg0, arg1, arg2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataStoreHashesForDurationAtTimestamp is a free data retrieval call binding the contract method 0x1bd2b3cf.
//
// Solidity: function dataStoreHashesForDurationAtTimestamp(uint8 , uint256 , uint256 ) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) DataStoreHashesForDurationAtTimestamp(arg0 uint8, arg1 *big.Int, arg2 *big.Int) ([32]byte, error) {
	return _DataLayrServiceManager.Contract.DataStoreHashesForDurationAtTimestamp(&_DataLayrServiceManager.CallOpts, arg0, arg1, arg2)
}

// DataStoreHashesForDurationAtTimestamp is a free data retrieval call binding the contract method 0x1bd2b3cf.
//
// Solidity: function dataStoreHashesForDurationAtTimestamp(uint8 , uint256 , uint256 ) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) DataStoreHashesForDurationAtTimestamp(arg0 uint8, arg1 *big.Int, arg2 *big.Int) ([32]byte, error) {
	return _DataLayrServiceManager.Contract.DataStoreHashesForDurationAtTimestamp(&_DataLayrServiceManager.CallOpts, arg0, arg1, arg2)
}

// DataStoreIdToSignatureHash is a free data retrieval call binding the contract method 0xfc2c6058.
//
// Solidity: function dataStoreIdToSignatureHash(uint32 ) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) DataStoreIdToSignatureHash(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "dataStoreIdToSignatureHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataStoreIdToSignatureHash is a free data retrieval call binding the contract method 0xfc2c6058.
//
// Solidity: function dataStoreIdToSignatureHash(uint32 ) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) DataStoreIdToSignatureHash(arg0 uint32) ([32]byte, error) {
	return _DataLayrServiceManager.Contract.DataStoreIdToSignatureHash(&_DataLayrServiceManager.CallOpts, arg0)
}

// DataStoreIdToSignatureHash is a free data retrieval call binding the contract method 0xfc2c6058.
//
// Solidity: function dataStoreIdToSignatureHash(uint32 ) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) DataStoreIdToSignatureHash(arg0 uint32) ([32]byte, error) {
	return _DataLayrServiceManager.Contract.DataStoreIdToSignatureHash(&_DataLayrServiceManager.CallOpts, arg0)
}

// DataStoresForDuration is a free data retrieval call binding the contract method 0x33223aea.
//
// Solidity: function dataStoresForDuration() view returns(uint32 one_duration, uint32 two_duration, uint32 three_duration, uint32 four_duration, uint32 five_duration, uint32 six_duration, uint32 seven_duration, uint32 dataStoreId, uint32 latestTime)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) DataStoresForDuration(opts *bind.CallOpts) (struct {
	OneDuration   uint32
	TwoDuration   uint32
	ThreeDuration uint32
	FourDuration  uint32
	FiveDuration  uint32
	SixDuration   uint32
	SevenDuration uint32
	DataStoreId   uint32
	LatestTime    uint32
}, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "dataStoresForDuration")

	outstruct := new(struct {
		OneDuration   uint32
		TwoDuration   uint32
		ThreeDuration uint32
		FourDuration  uint32
		FiveDuration  uint32
		SixDuration   uint32
		SevenDuration uint32
		DataStoreId   uint32
		LatestTime    uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OneDuration = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.TwoDuration = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ThreeDuration = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.FourDuration = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.FiveDuration = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.SixDuration = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.SevenDuration = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.DataStoreId = *abi.ConvertType(out[7], new(uint32)).(*uint32)
	outstruct.LatestTime = *abi.ConvertType(out[8], new(uint32)).(*uint32)

	return *outstruct, err

}

// DataStoresForDuration is a free data retrieval call binding the contract method 0x33223aea.
//
// Solidity: function dataStoresForDuration() view returns(uint32 one_duration, uint32 two_duration, uint32 three_duration, uint32 four_duration, uint32 five_duration, uint32 six_duration, uint32 seven_duration, uint32 dataStoreId, uint32 latestTime)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) DataStoresForDuration() (struct {
	OneDuration   uint32
	TwoDuration   uint32
	ThreeDuration uint32
	FourDuration  uint32
	FiveDuration  uint32
	SixDuration   uint32
	SevenDuration uint32
	DataStoreId   uint32
	LatestTime    uint32
}, error) {
	return _DataLayrServiceManager.Contract.DataStoresForDuration(&_DataLayrServiceManager.CallOpts)
}

// DataStoresForDuration is a free data retrieval call binding the contract method 0x33223aea.
//
// Solidity: function dataStoresForDuration() view returns(uint32 one_duration, uint32 two_duration, uint32 three_duration, uint32 four_duration, uint32 five_duration, uint32 six_duration, uint32 seven_duration, uint32 dataStoreId, uint32 latestTime)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) DataStoresForDuration() (struct {
	OneDuration   uint32
	TwoDuration   uint32
	ThreeDuration uint32
	FourDuration  uint32
	FiveDuration  uint32
	SixDuration   uint32
	SevenDuration uint32
	DataStoreId   uint32
	LatestTime    uint32
}, error) {
	return _DataLayrServiceManager.Contract.DataStoresForDuration(&_DataLayrServiceManager.CallOpts)
}

// EigenLayrDelegation is a free data retrieval call binding the contract method 0x33d2433a.
//
// Solidity: function eigenLayrDelegation() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) EigenLayrDelegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "eigenLayrDelegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenLayrDelegation is a free data retrieval call binding the contract method 0x33d2433a.
//
// Solidity: function eigenLayrDelegation() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) EigenLayrDelegation() (common.Address, error) {
	return _DataLayrServiceManager.Contract.EigenLayrDelegation(&_DataLayrServiceManager.CallOpts)
}

// EigenLayrDelegation is a free data retrieval call binding the contract method 0x33d2433a.
//
// Solidity: function eigenLayrDelegation() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) EigenLayrDelegation() (common.Address, error) {
	return _DataLayrServiceManager.Contract.EigenLayrDelegation(&_DataLayrServiceManager.CallOpts)
}

// EphemeralKeyRegistry is a free data retrieval call binding the contract method 0xcce36eff.
//
// Solidity: function ephemeralKeyRegistry() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) EphemeralKeyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "ephemeralKeyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EphemeralKeyRegistry is a free data retrieval call binding the contract method 0xcce36eff.
//
// Solidity: function ephemeralKeyRegistry() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) EphemeralKeyRegistry() (common.Address, error) {
	return _DataLayrServiceManager.Contract.EphemeralKeyRegistry(&_DataLayrServiceManager.CallOpts)
}

// EphemeralKeyRegistry is a free data retrieval call binding the contract method 0xcce36eff.
//
// Solidity: function ephemeralKeyRegistry() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) EphemeralKeyRegistry() (common.Address, error) {
	return _DataLayrServiceManager.Contract.EphemeralKeyRegistry(&_DataLayrServiceManager.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) Failed() (bool, error) {
	return _DataLayrServiceManager.Contract.Failed(&_DataLayrServiceManager.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) Failed() (bool, error) {
	return _DataLayrServiceManager.Contract.Failed(&_DataLayrServiceManager.CallOpts)
}

// FeePerBytePerTime is a free data retrieval call binding the contract method 0xd21eed4f.
//
// Solidity: function feePerBytePerTime() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) FeePerBytePerTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "feePerBytePerTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePerBytePerTime is a free data retrieval call binding the contract method 0xd21eed4f.
//
// Solidity: function feePerBytePerTime() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) FeePerBytePerTime() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.FeePerBytePerTime(&_DataLayrServiceManager.CallOpts)
}

// FeePerBytePerTime is a free data retrieval call binding the contract method 0xd21eed4f.
//
// Solidity: function feePerBytePerTime() view returns(uint256)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) FeePerBytePerTime() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.FeePerBytePerTime(&_DataLayrServiceManager.CallOpts)
}

// FirstQuorumThresholdPercentage is a free data retrieval call binding the contract method 0x982a0792.
//
// Solidity: function firstQuorumThresholdPercentage() view returns(uint128)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) FirstQuorumThresholdPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "firstQuorumThresholdPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstQuorumThresholdPercentage is a free data retrieval call binding the contract method 0x982a0792.
//
// Solidity: function firstQuorumThresholdPercentage() view returns(uint128)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) FirstQuorumThresholdPercentage() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.FirstQuorumThresholdPercentage(&_DataLayrServiceManager.CallOpts)
}

// FirstQuorumThresholdPercentage is a free data retrieval call binding the contract method 0x982a0792.
//
// Solidity: function firstQuorumThresholdPercentage() view returns(uint128)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) FirstQuorumThresholdPercentage() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.FirstQuorumThresholdPercentage(&_DataLayrServiceManager.CallOpts)
}

// GetDataStoreHashesForDurationAtTimestamp is a free data retrieval call binding the contract method 0xed82c0ee.
//
// Solidity: function getDataStoreHashesForDurationAtTimestamp(uint8 duration, uint256 timestamp, uint32 index) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) GetDataStoreHashesForDurationAtTimestamp(opts *bind.CallOpts, duration uint8, timestamp *big.Int, index uint32) ([32]byte, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "getDataStoreHashesForDurationAtTimestamp", duration, timestamp, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDataStoreHashesForDurationAtTimestamp is a free data retrieval call binding the contract method 0xed82c0ee.
//
// Solidity: function getDataStoreHashesForDurationAtTimestamp(uint8 duration, uint256 timestamp, uint32 index) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) GetDataStoreHashesForDurationAtTimestamp(duration uint8, timestamp *big.Int, index uint32) ([32]byte, error) {
	return _DataLayrServiceManager.Contract.GetDataStoreHashesForDurationAtTimestamp(&_DataLayrServiceManager.CallOpts, duration, timestamp, index)
}

// GetDataStoreHashesForDurationAtTimestamp is a free data retrieval call binding the contract method 0xed82c0ee.
//
// Solidity: function getDataStoreHashesForDurationAtTimestamp(uint8 duration, uint256 timestamp, uint32 index) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) GetDataStoreHashesForDurationAtTimestamp(duration uint8, timestamp *big.Int, index uint32) ([32]byte, error) {
	return _DataLayrServiceManager.Contract.GetDataStoreHashesForDurationAtTimestamp(&_DataLayrServiceManager.CallOpts, duration, timestamp, index)
}

// GetNumDataStoresForDuration is a free data retrieval call binding the contract method 0x73441c4e.
//
// Solidity: function getNumDataStoresForDuration(uint8 duration) view returns(uint32)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) GetNumDataStoresForDuration(opts *bind.CallOpts, duration uint8) (uint32, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "getNumDataStoresForDuration", duration)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetNumDataStoresForDuration is a free data retrieval call binding the contract method 0x73441c4e.
//
// Solidity: function getNumDataStoresForDuration(uint8 duration) view returns(uint32)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) GetNumDataStoresForDuration(duration uint8) (uint32, error) {
	return _DataLayrServiceManager.Contract.GetNumDataStoresForDuration(&_DataLayrServiceManager.CallOpts, duration)
}

// GetNumDataStoresForDuration is a free data retrieval call binding the contract method 0x73441c4e.
//
// Solidity: function getNumDataStoresForDuration(uint8 duration) view returns(uint32)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) GetNumDataStoresForDuration(duration uint8) (uint32, error) {
	return _DataLayrServiceManager.Contract.GetNumDataStoresForDuration(&_DataLayrServiceManager.CallOpts, duration)
}

// InvestmentManager is a free data retrieval call binding the contract method 0x4b31bb10.
//
// Solidity: function investmentManager() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) InvestmentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "investmentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InvestmentManager is a free data retrieval call binding the contract method 0x4b31bb10.
//
// Solidity: function investmentManager() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) InvestmentManager() (common.Address, error) {
	return _DataLayrServiceManager.Contract.InvestmentManager(&_DataLayrServiceManager.CallOpts)
}

// InvestmentManager is a free data retrieval call binding the contract method 0x4b31bb10.
//
// Solidity: function investmentManager() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) InvestmentManager() (common.Address, error) {
	return _DataLayrServiceManager.Contract.InvestmentManager(&_DataLayrServiceManager.CallOpts)
}

// LatestTime is a free data retrieval call binding the contract method 0x7dfd16d7.
//
// Solidity: function latestTime() view returns(uint32)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) LatestTime(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "latestTime")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTime is a free data retrieval call binding the contract method 0x7dfd16d7.
//
// Solidity: function latestTime() view returns(uint32)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) LatestTime() (uint32, error) {
	return _DataLayrServiceManager.Contract.LatestTime(&_DataLayrServiceManager.CallOpts)
}

// LatestTime is a free data retrieval call binding the contract method 0x7dfd16d7.
//
// Solidity: function latestTime() view returns(uint32)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) LatestTime() (uint32, error) {
	return _DataLayrServiceManager.Contract.LatestTime(&_DataLayrServiceManager.CallOpts)
}

// Log2NumPowersOfTau is a free data retrieval call binding the contract method 0xa50017a1.
//
// Solidity: function log2NumPowersOfTau() view returns(uint48)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) Log2NumPowersOfTau(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "log2NumPowersOfTau")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Log2NumPowersOfTau is a free data retrieval call binding the contract method 0xa50017a1.
//
// Solidity: function log2NumPowersOfTau() view returns(uint48)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) Log2NumPowersOfTau() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.Log2NumPowersOfTau(&_DataLayrServiceManager.CallOpts)
}

// Log2NumPowersOfTau is a free data retrieval call binding the contract method 0xa50017a1.
//
// Solidity: function log2NumPowersOfTau() view returns(uint48)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) Log2NumPowersOfTau() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.Log2NumPowersOfTau(&_DataLayrServiceManager.CallOpts)
}

// NumPowersOfTau is a free data retrieval call binding the contract method 0x046bf4a6.
//
// Solidity: function numPowersOfTau() view returns(uint48)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) NumPowersOfTau(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "numPowersOfTau")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumPowersOfTau is a free data retrieval call binding the contract method 0x046bf4a6.
//
// Solidity: function numPowersOfTau() view returns(uint48)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) NumPowersOfTau() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.NumPowersOfTau(&_DataLayrServiceManager.CallOpts)
}

// NumPowersOfTau is a free data retrieval call binding the contract method 0x046bf4a6.
//
// Solidity: function numPowersOfTau() view returns(uint48)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) NumPowersOfTau() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.NumPowersOfTau(&_DataLayrServiceManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) Paused() (bool, error) {
	return _DataLayrServiceManager.Contract.Paused(&_DataLayrServiceManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) Paused() (bool, error) {
	return _DataLayrServiceManager.Contract.Paused(&_DataLayrServiceManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) PauserRegistry() (common.Address, error) {
	return _DataLayrServiceManager.Contract.PauserRegistry(&_DataLayrServiceManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _DataLayrServiceManager.Contract.PauserRegistry(&_DataLayrServiceManager.CallOpts)
}

// QuorumThresholdBasisPoints is a free data retrieval call binding the contract method 0xb569157b.
//
// Solidity: function quorumThresholdBasisPoints() view returns(uint16)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) QuorumThresholdBasisPoints(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "quorumThresholdBasisPoints")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// QuorumThresholdBasisPoints is a free data retrieval call binding the contract method 0xb569157b.
//
// Solidity: function quorumThresholdBasisPoints() view returns(uint16)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) QuorumThresholdBasisPoints() (uint16, error) {
	return _DataLayrServiceManager.Contract.QuorumThresholdBasisPoints(&_DataLayrServiceManager.CallOpts)
}

// QuorumThresholdBasisPoints is a free data retrieval call binding the contract method 0xb569157b.
//
// Solidity: function quorumThresholdBasisPoints() view returns(uint16)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) QuorumThresholdBasisPoints() (uint16, error) {
	return _DataLayrServiceManager.Contract.QuorumThresholdBasisPoints(&_DataLayrServiceManager.CallOpts)
}

// Repository is a free data retrieval call binding the contract method 0xe9176c60.
//
// Solidity: function repository() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) Repository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "repository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Repository is a free data retrieval call binding the contract method 0xe9176c60.
//
// Solidity: function repository() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) Repository() (common.Address, error) {
	return _DataLayrServiceManager.Contract.Repository(&_DataLayrServiceManager.CallOpts)
}

// Repository is a free data retrieval call binding the contract method 0xe9176c60.
//
// Solidity: function repository() view returns(address)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) Repository() (common.Address, error) {
	return _DataLayrServiceManager.Contract.Repository(&_DataLayrServiceManager.CallOpts)
}

// SecondQuorumThresholdPercentage is a free data retrieval call binding the contract method 0x41345163.
//
// Solidity: function secondQuorumThresholdPercentage() view returns(uint128)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) SecondQuorumThresholdPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "secondQuorumThresholdPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondQuorumThresholdPercentage is a free data retrieval call binding the contract method 0x41345163.
//
// Solidity: function secondQuorumThresholdPercentage() view returns(uint128)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) SecondQuorumThresholdPercentage() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.SecondQuorumThresholdPercentage(&_DataLayrServiceManager.CallOpts)
}

// SecondQuorumThresholdPercentage is a free data retrieval call binding the contract method 0x41345163.
//
// Solidity: function secondQuorumThresholdPercentage() view returns(uint128)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) SecondQuorumThresholdPercentage() (*big.Int, error) {
	return _DataLayrServiceManager.Contract.SecondQuorumThresholdPercentage(&_DataLayrServiceManager.CallOpts)
}

// StakeWithdrawalVerification is a free data retrieval call binding the contract method 0xa88d92aa.
//
// Solidity: function stakeWithdrawalVerification(bytes packedDataStoreSearchData, uint256 initTimestamp, uint256 unlockTime) view returns()
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) StakeWithdrawalVerification(opts *bind.CallOpts, packedDataStoreSearchData []byte, initTimestamp *big.Int, unlockTime *big.Int) error {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "stakeWithdrawalVerification", packedDataStoreSearchData, initTimestamp, unlockTime)

	if err != nil {
		return err
	}

	return err

}

// StakeWithdrawalVerification is a free data retrieval call binding the contract method 0xa88d92aa.
//
// Solidity: function stakeWithdrawalVerification(bytes packedDataStoreSearchData, uint256 initTimestamp, uint256 unlockTime) view returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) StakeWithdrawalVerification(packedDataStoreSearchData []byte, initTimestamp *big.Int, unlockTime *big.Int) error {
	return _DataLayrServiceManager.Contract.StakeWithdrawalVerification(&_DataLayrServiceManager.CallOpts, packedDataStoreSearchData, initTimestamp, unlockTime)
}

// StakeWithdrawalVerification is a free data retrieval call binding the contract method 0xa88d92aa.
//
// Solidity: function stakeWithdrawalVerification(bytes packedDataStoreSearchData, uint256 initTimestamp, uint256 unlockTime) view returns()
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) StakeWithdrawalVerification(packedDataStoreSearchData []byte, initTimestamp *big.Int, unlockTime *big.Int) error {
	return _DataLayrServiceManager.Contract.StakeWithdrawalVerification(&_DataLayrServiceManager.CallOpts, packedDataStoreSearchData, initTimestamp, unlockTime)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) TaskNumber() (uint32, error) {
	return _DataLayrServiceManager.Contract.TaskNumber(&_DataLayrServiceManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) TaskNumber() (uint32, error) {
	return _DataLayrServiceManager.Contract.TaskNumber(&_DataLayrServiceManager.CallOpts)
}

// VerifyDataStoreMetadata is a free data retrieval call binding the contract method 0xbda39e6f.
//
// Solidity: function verifyDataStoreMetadata(uint8 duration, uint256 timestamp, uint32 index, (bytes32,uint32,uint32,uint32,uint96,address,bytes32) metadata) view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) VerifyDataStoreMetadata(opts *bind.CallOpts, duration uint8, timestamp *big.Int, index uint32, metadata IDataLayrServiceManagerDataStoreMetadata) (bool, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "verifyDataStoreMetadata", duration, timestamp, index, metadata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDataStoreMetadata is a free data retrieval call binding the contract method 0xbda39e6f.
//
// Solidity: function verifyDataStoreMetadata(uint8 duration, uint256 timestamp, uint32 index, (bytes32,uint32,uint32,uint32,uint96,address,bytes32) metadata) view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) VerifyDataStoreMetadata(duration uint8, timestamp *big.Int, index uint32, metadata IDataLayrServiceManagerDataStoreMetadata) (bool, error) {
	return _DataLayrServiceManager.Contract.VerifyDataStoreMetadata(&_DataLayrServiceManager.CallOpts, duration, timestamp, index, metadata)
}

// VerifyDataStoreMetadata is a free data retrieval call binding the contract method 0xbda39e6f.
//
// Solidity: function verifyDataStoreMetadata(uint8 duration, uint256 timestamp, uint32 index, (bytes32,uint32,uint32,uint32,uint96,address,bytes32) metadata) view returns(bool)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) VerifyDataStoreMetadata(duration uint8, timestamp *big.Int, index uint32, metadata IDataLayrServiceManagerDataStoreMetadata) (bool, error) {
	return _DataLayrServiceManager.Contract.VerifyDataStoreMetadata(&_DataLayrServiceManager.CallOpts, duration, timestamp, index, metadata)
}

// ZeroPolynomialCommitmentMerkleRoots is a free data retrieval call binding the contract method 0x3367a3fb.
//
// Solidity: function zeroPolynomialCommitmentMerkleRoots(uint256 ) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerCaller) ZeroPolynomialCommitmentMerkleRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DataLayrServiceManager.contract.Call(opts, &out, "zeroPolynomialCommitmentMerkleRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZeroPolynomialCommitmentMerkleRoots is a free data retrieval call binding the contract method 0x3367a3fb.
//
// Solidity: function zeroPolynomialCommitmentMerkleRoots(uint256 ) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) ZeroPolynomialCommitmentMerkleRoots(arg0 *big.Int) ([32]byte, error) {
	return _DataLayrServiceManager.Contract.ZeroPolynomialCommitmentMerkleRoots(&_DataLayrServiceManager.CallOpts, arg0)
}

// ZeroPolynomialCommitmentMerkleRoots is a free data retrieval call binding the contract method 0x3367a3fb.
//
// Solidity: function zeroPolynomialCommitmentMerkleRoots(uint256 ) view returns(bytes32)
func (_DataLayrServiceManager *DataLayrServiceManagerCallerSession) ZeroPolynomialCommitmentMerkleRoots(arg0 *big.Int) ([32]byte, error) {
	return _DataLayrServiceManager.Contract.ZeroPolynomialCommitmentMerkleRoots(&_DataLayrServiceManager.CallOpts, arg0)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0xdeaf4498.
//
// Solidity: function checkSignatures(bytes data) returns(uint32 taskNumberToConfirm, uint32 stakesBlockNumber, bytes32 msgHash, (uint256,uint256,uint256,uint256) signedTotals, bytes32 compressedSignatoryRecord)
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) CheckSignatures(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "checkSignatures", data)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0xdeaf4498.
//
// Solidity: function checkSignatures(bytes data) returns(uint32 taskNumberToConfirm, uint32 stakesBlockNumber, bytes32 msgHash, (uint256,uint256,uint256,uint256) signedTotals, bytes32 compressedSignatoryRecord)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) CheckSignatures(data []byte) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.CheckSignatures(&_DataLayrServiceManager.TransactOpts, data)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0xdeaf4498.
//
// Solidity: function checkSignatures(bytes data) returns(uint32 taskNumberToConfirm, uint32 stakesBlockNumber, bytes32 msgHash, (uint256,uint256,uint256,uint256) signedTotals, bytes32 compressedSignatoryRecord)
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) CheckSignatures(data []byte) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.CheckSignatures(&_DataLayrServiceManager.TransactOpts, data)
}

// ConfirmDataStore is a paid mutator transaction binding the contract method 0x51899515.
//
// Solidity: function confirmDataStore(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) ConfirmDataStore(opts *bind.TransactOpts, data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "confirmDataStore", data, searchData)
}

// ConfirmDataStore is a paid mutator transaction binding the contract method 0x51899515.
//
// Solidity: function confirmDataStore(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) ConfirmDataStore(data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.ConfirmDataStore(&_DataLayrServiceManager.TransactOpts, data, searchData)
}

// ConfirmDataStore is a paid mutator transaction binding the contract method 0x51899515.
//
// Solidity: function confirmDataStore(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) ConfirmDataStore(data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.ConfirmDataStore(&_DataLayrServiceManager.TransactOpts, data, searchData)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operator) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "freezeOperator", operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operator) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) FreezeOperator(operator common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.FreezeOperator(&_DataLayrServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operator) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) FreezeOperator(operator common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.FreezeOperator(&_DataLayrServiceManager.TransactOpts, operator)
}

// InitDataStore is a paid mutator transaction binding the contract method 0xdcf49ea7.
//
// Solidity: function initDataStore(address feePayer, address confirmer, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex, bytes header) returns(uint32 index)
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) InitDataStore(opts *bind.TransactOpts, feePayer common.Address, confirmer common.Address, duration uint8, blockNumber uint32, totalOperatorsIndex uint32, header []byte) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "initDataStore", feePayer, confirmer, duration, blockNumber, totalOperatorsIndex, header)
}

// InitDataStore is a paid mutator transaction binding the contract method 0xdcf49ea7.
//
// Solidity: function initDataStore(address feePayer, address confirmer, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex, bytes header) returns(uint32 index)
func (_DataLayrServiceManager *DataLayrServiceManagerSession) InitDataStore(feePayer common.Address, confirmer common.Address, duration uint8, blockNumber uint32, totalOperatorsIndex uint32, header []byte) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.InitDataStore(&_DataLayrServiceManager.TransactOpts, feePayer, confirmer, duration, blockNumber, totalOperatorsIndex, header)
}

// InitDataStore is a paid mutator transaction binding the contract method 0xdcf49ea7.
//
// Solidity: function initDataStore(address feePayer, address confirmer, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex, bytes header) returns(uint32 index)
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) InitDataStore(feePayer common.Address, confirmer common.Address, duration uint8, blockNumber uint32, totalOperatorsIndex uint32, header []byte) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.InitDataStore(&_DataLayrServiceManager.TransactOpts, feePayer, confirmer, duration, blockNumber, totalOperatorsIndex, header)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) Pause() (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.Pause(&_DataLayrServiceManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.Pause(&_DataLayrServiceManager.TransactOpts)
}

// RevokeSlashingAbility is a paid mutator transaction binding the contract method 0xfb3f2922.
//
// Solidity: function revokeSlashingAbility(address operator, uint32 unbondedAfter) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) RevokeSlashingAbility(opts *bind.TransactOpts, operator common.Address, unbondedAfter uint32) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "revokeSlashingAbility", operator, unbondedAfter)
}

// RevokeSlashingAbility is a paid mutator transaction binding the contract method 0xfb3f2922.
//
// Solidity: function revokeSlashingAbility(address operator, uint32 unbondedAfter) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) RevokeSlashingAbility(operator common.Address, unbondedAfter uint32) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.RevokeSlashingAbility(&_DataLayrServiceManager.TransactOpts, operator, unbondedAfter)
}

// RevokeSlashingAbility is a paid mutator transaction binding the contract method 0xfb3f2922.
//
// Solidity: function revokeSlashingAbility(address operator, uint32 unbondedAfter) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) RevokeSlashingAbility(operator common.Address, unbondedAfter uint32) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.RevokeSlashingAbility(&_DataLayrServiceManager.TransactOpts, operator, unbondedAfter)
}

// SetAdversaryThresholdBasisPoints is a paid mutator transaction binding the contract method 0x516d8616.
//
// Solidity: function setAdversaryThresholdBasisPoints(uint16 _adversaryThresholdBasisPoints) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) SetAdversaryThresholdBasisPoints(opts *bind.TransactOpts, _adversaryThresholdBasisPoints uint16) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "setAdversaryThresholdBasisPoints", _adversaryThresholdBasisPoints)
}

// SetAdversaryThresholdBasisPoints is a paid mutator transaction binding the contract method 0x516d8616.
//
// Solidity: function setAdversaryThresholdBasisPoints(uint16 _adversaryThresholdBasisPoints) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) SetAdversaryThresholdBasisPoints(_adversaryThresholdBasisPoints uint16) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetAdversaryThresholdBasisPoints(&_DataLayrServiceManager.TransactOpts, _adversaryThresholdBasisPoints)
}

// SetAdversaryThresholdBasisPoints is a paid mutator transaction binding the contract method 0x516d8616.
//
// Solidity: function setAdversaryThresholdBasisPoints(uint16 _adversaryThresholdBasisPoints) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) SetAdversaryThresholdBasisPoints(_adversaryThresholdBasisPoints uint16) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetAdversaryThresholdBasisPoints(&_DataLayrServiceManager.TransactOpts, _adversaryThresholdBasisPoints)
}

// SetBombVerifier is a paid mutator transaction binding the contract method 0x227e375d.
//
// Solidity: function setBombVerifier(address _dataLayrBombVerifier) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) SetBombVerifier(opts *bind.TransactOpts, _dataLayrBombVerifier common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "setBombVerifier", _dataLayrBombVerifier)
}

// SetBombVerifier is a paid mutator transaction binding the contract method 0x227e375d.
//
// Solidity: function setBombVerifier(address _dataLayrBombVerifier) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) SetBombVerifier(_dataLayrBombVerifier common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetBombVerifier(&_DataLayrServiceManager.TransactOpts, _dataLayrBombVerifier)
}

// SetBombVerifier is a paid mutator transaction binding the contract method 0x227e375d.
//
// Solidity: function setBombVerifier(address _dataLayrBombVerifier) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) SetBombVerifier(_dataLayrBombVerifier common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetBombVerifier(&_DataLayrServiceManager.TransactOpts, _dataLayrBombVerifier)
}

// SetEphemeralKeyRegistry is a paid mutator transaction binding the contract method 0xf198b50e.
//
// Solidity: function setEphemeralKeyRegistry(address _ephemeralKeyRegistry) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) SetEphemeralKeyRegistry(opts *bind.TransactOpts, _ephemeralKeyRegistry common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "setEphemeralKeyRegistry", _ephemeralKeyRegistry)
}

// SetEphemeralKeyRegistry is a paid mutator transaction binding the contract method 0xf198b50e.
//
// Solidity: function setEphemeralKeyRegistry(address _ephemeralKeyRegistry) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) SetEphemeralKeyRegistry(_ephemeralKeyRegistry common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetEphemeralKeyRegistry(&_DataLayrServiceManager.TransactOpts, _ephemeralKeyRegistry)
}

// SetEphemeralKeyRegistry is a paid mutator transaction binding the contract method 0xf198b50e.
//
// Solidity: function setEphemeralKeyRegistry(address _ephemeralKeyRegistry) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) SetEphemeralKeyRegistry(_ephemeralKeyRegistry common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetEphemeralKeyRegistry(&_DataLayrServiceManager.TransactOpts, _ephemeralKeyRegistry)
}

// SetFeePerBytePerTime is a paid mutator transaction binding the contract method 0x772eefe3.
//
// Solidity: function setFeePerBytePerTime(uint256 _feePerBytePerTime) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) SetFeePerBytePerTime(opts *bind.TransactOpts, _feePerBytePerTime *big.Int) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "setFeePerBytePerTime", _feePerBytePerTime)
}

// SetFeePerBytePerTime is a paid mutator transaction binding the contract method 0x772eefe3.
//
// Solidity: function setFeePerBytePerTime(uint256 _feePerBytePerTime) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) SetFeePerBytePerTime(_feePerBytePerTime *big.Int) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetFeePerBytePerTime(&_DataLayrServiceManager.TransactOpts, _feePerBytePerTime)
}

// SetFeePerBytePerTime is a paid mutator transaction binding the contract method 0x772eefe3.
//
// Solidity: function setFeePerBytePerTime(uint256 _feePerBytePerTime) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) SetFeePerBytePerTime(_feePerBytePerTime *big.Int) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetFeePerBytePerTime(&_DataLayrServiceManager.TransactOpts, _feePerBytePerTime)
}

// SetFirstQuorumThresholdPercentage is a paid mutator transaction binding the contract method 0x2433fb6e.
//
// Solidity: function setFirstQuorumThresholdPercentage(uint128 _firstQuorumThresholdPercentage) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) SetFirstQuorumThresholdPercentage(opts *bind.TransactOpts, _firstQuorumThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "setFirstQuorumThresholdPercentage", _firstQuorumThresholdPercentage)
}

// SetFirstQuorumThresholdPercentage is a paid mutator transaction binding the contract method 0x2433fb6e.
//
// Solidity: function setFirstQuorumThresholdPercentage(uint128 _firstQuorumThresholdPercentage) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) SetFirstQuorumThresholdPercentage(_firstQuorumThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetFirstQuorumThresholdPercentage(&_DataLayrServiceManager.TransactOpts, _firstQuorumThresholdPercentage)
}

// SetFirstQuorumThresholdPercentage is a paid mutator transaction binding the contract method 0x2433fb6e.
//
// Solidity: function setFirstQuorumThresholdPercentage(uint128 _firstQuorumThresholdPercentage) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) SetFirstQuorumThresholdPercentage(_firstQuorumThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetFirstQuorumThresholdPercentage(&_DataLayrServiceManager.TransactOpts, _firstQuorumThresholdPercentage)
}

// SetLowDegreeChallenge is a paid mutator transaction binding the contract method 0x2417859b.
//
// Solidity: function setLowDegreeChallenge(address _dataLayrLowDegreeChallenge) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) SetLowDegreeChallenge(opts *bind.TransactOpts, _dataLayrLowDegreeChallenge common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "setLowDegreeChallenge", _dataLayrLowDegreeChallenge)
}

// SetLowDegreeChallenge is a paid mutator transaction binding the contract method 0x2417859b.
//
// Solidity: function setLowDegreeChallenge(address _dataLayrLowDegreeChallenge) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) SetLowDegreeChallenge(_dataLayrLowDegreeChallenge common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetLowDegreeChallenge(&_DataLayrServiceManager.TransactOpts, _dataLayrLowDegreeChallenge)
}

// SetLowDegreeChallenge is a paid mutator transaction binding the contract method 0x2417859b.
//
// Solidity: function setLowDegreeChallenge(address _dataLayrLowDegreeChallenge) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) SetLowDegreeChallenge(_dataLayrLowDegreeChallenge common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetLowDegreeChallenge(&_DataLayrServiceManager.TransactOpts, _dataLayrLowDegreeChallenge)
}

// SetPaymentManager is a paid mutator transaction binding the contract method 0x7e702dc8.
//
// Solidity: function setPaymentManager(address _dataLayrPaymentManager) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) SetPaymentManager(opts *bind.TransactOpts, _dataLayrPaymentManager common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "setPaymentManager", _dataLayrPaymentManager)
}

// SetPaymentManager is a paid mutator transaction binding the contract method 0x7e702dc8.
//
// Solidity: function setPaymentManager(address _dataLayrPaymentManager) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) SetPaymentManager(_dataLayrPaymentManager common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetPaymentManager(&_DataLayrServiceManager.TransactOpts, _dataLayrPaymentManager)
}

// SetPaymentManager is a paid mutator transaction binding the contract method 0x7e702dc8.
//
// Solidity: function setPaymentManager(address _dataLayrPaymentManager) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) SetPaymentManager(_dataLayrPaymentManager common.Address) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetPaymentManager(&_DataLayrServiceManager.TransactOpts, _dataLayrPaymentManager)
}

// SetQuorumThresholdBasisPoints is a paid mutator transaction binding the contract method 0x07dab8b3.
//
// Solidity: function setQuorumThresholdBasisPoints(uint16 _quorumThresholdBasisPoints) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) SetQuorumThresholdBasisPoints(opts *bind.TransactOpts, _quorumThresholdBasisPoints uint16) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "setQuorumThresholdBasisPoints", _quorumThresholdBasisPoints)
}

// SetQuorumThresholdBasisPoints is a paid mutator transaction binding the contract method 0x07dab8b3.
//
// Solidity: function setQuorumThresholdBasisPoints(uint16 _quorumThresholdBasisPoints) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) SetQuorumThresholdBasisPoints(_quorumThresholdBasisPoints uint16) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetQuorumThresholdBasisPoints(&_DataLayrServiceManager.TransactOpts, _quorumThresholdBasisPoints)
}

// SetQuorumThresholdBasisPoints is a paid mutator transaction binding the contract method 0x07dab8b3.
//
// Solidity: function setQuorumThresholdBasisPoints(uint16 _quorumThresholdBasisPoints) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) SetQuorumThresholdBasisPoints(_quorumThresholdBasisPoints uint16) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetQuorumThresholdBasisPoints(&_DataLayrServiceManager.TransactOpts, _quorumThresholdBasisPoints)
}

// SetSecondQuorumThresholdPercentage is a paid mutator transaction binding the contract method 0xe4038e18.
//
// Solidity: function setSecondQuorumThresholdPercentage(uint128 _secondQuorumThresholdPercentage) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) SetSecondQuorumThresholdPercentage(opts *bind.TransactOpts, _secondQuorumThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "setSecondQuorumThresholdPercentage", _secondQuorumThresholdPercentage)
}

// SetSecondQuorumThresholdPercentage is a paid mutator transaction binding the contract method 0xe4038e18.
//
// Solidity: function setSecondQuorumThresholdPercentage(uint128 _secondQuorumThresholdPercentage) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) SetSecondQuorumThresholdPercentage(_secondQuorumThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetSecondQuorumThresholdPercentage(&_DataLayrServiceManager.TransactOpts, _secondQuorumThresholdPercentage)
}

// SetSecondQuorumThresholdPercentage is a paid mutator transaction binding the contract method 0xe4038e18.
//
// Solidity: function setSecondQuorumThresholdPercentage(uint128 _secondQuorumThresholdPercentage) returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) SetSecondQuorumThresholdPercentage(_secondQuorumThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.SetSecondQuorumThresholdPercentage(&_DataLayrServiceManager.TransactOpts, _secondQuorumThresholdPercentage)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLayrServiceManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DataLayrServiceManager *DataLayrServiceManagerSession) Unpause() (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.Unpause(&_DataLayrServiceManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DataLayrServiceManager *DataLayrServiceManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _DataLayrServiceManager.Contract.Unpause(&_DataLayrServiceManager.TransactOpts)
}

// DataLayrServiceManagerAdversaryThresholdBasisPointsUpdatedIterator is returned from FilterAdversaryThresholdBasisPointsUpdated and is used to iterate over the raw logs and unpacked data for AdversaryThresholdBasisPointsUpdated events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerAdversaryThresholdBasisPointsUpdatedIterator struct {
	Event *DataLayrServiceManagerAdversaryThresholdBasisPointsUpdated // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerAdversaryThresholdBasisPointsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerAdversaryThresholdBasisPointsUpdated)
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
		it.Event = new(DataLayrServiceManagerAdversaryThresholdBasisPointsUpdated)
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
func (it *DataLayrServiceManagerAdversaryThresholdBasisPointsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerAdversaryThresholdBasisPointsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerAdversaryThresholdBasisPointsUpdated represents a AdversaryThresholdBasisPointsUpdated event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerAdversaryThresholdBasisPointsUpdated struct {
	AdversaryThresholdBasisPoints uint16
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterAdversaryThresholdBasisPointsUpdated is a free log retrieval operation binding the contract event 0x1bdc513ac13a36cd49087fef52b034cb5833bd75154db5239f27daa6bde17042.
//
// Solidity: event AdversaryThresholdBasisPointsUpdated(uint16 adversaryThresholdBasisPoints)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterAdversaryThresholdBasisPointsUpdated(opts *bind.FilterOpts) (*DataLayrServiceManagerAdversaryThresholdBasisPointsUpdatedIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "AdversaryThresholdBasisPointsUpdated")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerAdversaryThresholdBasisPointsUpdatedIterator{contract: _DataLayrServiceManager.contract, event: "AdversaryThresholdBasisPointsUpdated", logs: logs, sub: sub}, nil
}

// WatchAdversaryThresholdBasisPointsUpdated is a free log subscription operation binding the contract event 0x1bdc513ac13a36cd49087fef52b034cb5833bd75154db5239f27daa6bde17042.
//
// Solidity: event AdversaryThresholdBasisPointsUpdated(uint16 adversaryThresholdBasisPoints)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchAdversaryThresholdBasisPointsUpdated(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerAdversaryThresholdBasisPointsUpdated) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "AdversaryThresholdBasisPointsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerAdversaryThresholdBasisPointsUpdated)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "AdversaryThresholdBasisPointsUpdated", log); err != nil {
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

// ParseAdversaryThresholdBasisPointsUpdated is a log parse operation binding the contract event 0x1bdc513ac13a36cd49087fef52b034cb5833bd75154db5239f27daa6bde17042.
//
// Solidity: event AdversaryThresholdBasisPointsUpdated(uint16 adversaryThresholdBasisPoints)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseAdversaryThresholdBasisPointsUpdated(log types.Log) (*DataLayrServiceManagerAdversaryThresholdBasisPointsUpdated, error) {
	event := new(DataLayrServiceManagerAdversaryThresholdBasisPointsUpdated)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "AdversaryThresholdBasisPointsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerBombVerifierSetIterator is returned from FilterBombVerifierSet and is used to iterate over the raw logs and unpacked data for BombVerifierSet events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerBombVerifierSetIterator struct {
	Event *DataLayrServiceManagerBombVerifierSet // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerBombVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerBombVerifierSet)
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
		it.Event = new(DataLayrServiceManagerBombVerifierSet)
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
func (it *DataLayrServiceManagerBombVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerBombVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerBombVerifierSet represents a BombVerifierSet event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerBombVerifierSet struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBombVerifierSet is a free log retrieval operation binding the contract event 0x875303dc4b1493d311d0dc6908455605e5fd8deae1190665f8f5a4365c58fe58.
//
// Solidity: event BombVerifierSet(address previousAddress, address newAddress)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterBombVerifierSet(opts *bind.FilterOpts) (*DataLayrServiceManagerBombVerifierSetIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "BombVerifierSet")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerBombVerifierSetIterator{contract: _DataLayrServiceManager.contract, event: "BombVerifierSet", logs: logs, sub: sub}, nil
}

// WatchBombVerifierSet is a free log subscription operation binding the contract event 0x875303dc4b1493d311d0dc6908455605e5fd8deae1190665f8f5a4365c58fe58.
//
// Solidity: event BombVerifierSet(address previousAddress, address newAddress)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchBombVerifierSet(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerBombVerifierSet) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "BombVerifierSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerBombVerifierSet)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "BombVerifierSet", log); err != nil {
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

// ParseBombVerifierSet is a log parse operation binding the contract event 0x875303dc4b1493d311d0dc6908455605e5fd8deae1190665f8f5a4365c58fe58.
//
// Solidity: event BombVerifierSet(address previousAddress, address newAddress)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseBombVerifierSet(log types.Log) (*DataLayrServiceManagerBombVerifierSet, error) {
	event := new(DataLayrServiceManagerBombVerifierSet)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "BombVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerConfirmDataStoreIterator is returned from FilterConfirmDataStore and is used to iterate over the raw logs and unpacked data for ConfirmDataStore events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerConfirmDataStoreIterator struct {
	Event *DataLayrServiceManagerConfirmDataStore // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerConfirmDataStoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerConfirmDataStore)
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
		it.Event = new(DataLayrServiceManagerConfirmDataStore)
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
func (it *DataLayrServiceManagerConfirmDataStoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerConfirmDataStoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerConfirmDataStore represents a ConfirmDataStore event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerConfirmDataStore struct {
	DataStoreId uint32
	HeaderHash  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmDataStore is a free log retrieval operation binding the contract event 0xfbb7f4f1b0b9ad9e75d69d22c364e13089418d86fcb5106792a53046c0fb33aa.
//
// Solidity: event ConfirmDataStore(uint32 dataStoreId, bytes32 headerHash)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterConfirmDataStore(opts *bind.FilterOpts) (*DataLayrServiceManagerConfirmDataStoreIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "ConfirmDataStore")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerConfirmDataStoreIterator{contract: _DataLayrServiceManager.contract, event: "ConfirmDataStore", logs: logs, sub: sub}, nil
}

// WatchConfirmDataStore is a free log subscription operation binding the contract event 0xfbb7f4f1b0b9ad9e75d69d22c364e13089418d86fcb5106792a53046c0fb33aa.
//
// Solidity: event ConfirmDataStore(uint32 dataStoreId, bytes32 headerHash)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchConfirmDataStore(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerConfirmDataStore) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "ConfirmDataStore")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerConfirmDataStore)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "ConfirmDataStore", log); err != nil {
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

// ParseConfirmDataStore is a log parse operation binding the contract event 0xfbb7f4f1b0b9ad9e75d69d22c364e13089418d86fcb5106792a53046c0fb33aa.
//
// Solidity: event ConfirmDataStore(uint32 dataStoreId, bytes32 headerHash)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseConfirmDataStore(log types.Log) (*DataLayrServiceManagerConfirmDataStore, error) {
	event := new(DataLayrServiceManagerConfirmDataStore)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "ConfirmDataStore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerEphemeralKeyRegistrySetIterator is returned from FilterEphemeralKeyRegistrySet and is used to iterate over the raw logs and unpacked data for EphemeralKeyRegistrySet events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerEphemeralKeyRegistrySetIterator struct {
	Event *DataLayrServiceManagerEphemeralKeyRegistrySet // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerEphemeralKeyRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerEphemeralKeyRegistrySet)
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
		it.Event = new(DataLayrServiceManagerEphemeralKeyRegistrySet)
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
func (it *DataLayrServiceManagerEphemeralKeyRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerEphemeralKeyRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerEphemeralKeyRegistrySet represents a EphemeralKeyRegistrySet event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerEphemeralKeyRegistrySet struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEphemeralKeyRegistrySet is a free log retrieval operation binding the contract event 0x5db33fb0e580eb7ed3895fca3d40083f8a10ae79230d4ed70338725a3d6c8eb7.
//
// Solidity: event EphemeralKeyRegistrySet(address previousAddress, address newAddress)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterEphemeralKeyRegistrySet(opts *bind.FilterOpts) (*DataLayrServiceManagerEphemeralKeyRegistrySetIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "EphemeralKeyRegistrySet")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerEphemeralKeyRegistrySetIterator{contract: _DataLayrServiceManager.contract, event: "EphemeralKeyRegistrySet", logs: logs, sub: sub}, nil
}

// WatchEphemeralKeyRegistrySet is a free log subscription operation binding the contract event 0x5db33fb0e580eb7ed3895fca3d40083f8a10ae79230d4ed70338725a3d6c8eb7.
//
// Solidity: event EphemeralKeyRegistrySet(address previousAddress, address newAddress)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchEphemeralKeyRegistrySet(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerEphemeralKeyRegistrySet) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "EphemeralKeyRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerEphemeralKeyRegistrySet)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "EphemeralKeyRegistrySet", log); err != nil {
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

// ParseEphemeralKeyRegistrySet is a log parse operation binding the contract event 0x5db33fb0e580eb7ed3895fca3d40083f8a10ae79230d4ed70338725a3d6c8eb7.
//
// Solidity: event EphemeralKeyRegistrySet(address previousAddress, address newAddress)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseEphemeralKeyRegistrySet(log types.Log) (*DataLayrServiceManagerEphemeralKeyRegistrySet, error) {
	event := new(DataLayrServiceManagerEphemeralKeyRegistrySet)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "EphemeralKeyRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerFeePerBytePerTimeSetIterator is returned from FilterFeePerBytePerTimeSet and is used to iterate over the raw logs and unpacked data for FeePerBytePerTimeSet events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerFeePerBytePerTimeSetIterator struct {
	Event *DataLayrServiceManagerFeePerBytePerTimeSet // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerFeePerBytePerTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerFeePerBytePerTimeSet)
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
		it.Event = new(DataLayrServiceManagerFeePerBytePerTimeSet)
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
func (it *DataLayrServiceManagerFeePerBytePerTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerFeePerBytePerTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerFeePerBytePerTimeSet represents a FeePerBytePerTimeSet event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerFeePerBytePerTimeSet struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeePerBytePerTimeSet is a free log retrieval operation binding the contract event 0xcd1b2c2a220284accd1f9effd811cdecb6beaa4638618b48bbea07ce7ae16996.
//
// Solidity: event FeePerBytePerTimeSet(uint256 previousValue, uint256 newValue)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterFeePerBytePerTimeSet(opts *bind.FilterOpts) (*DataLayrServiceManagerFeePerBytePerTimeSetIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "FeePerBytePerTimeSet")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerFeePerBytePerTimeSetIterator{contract: _DataLayrServiceManager.contract, event: "FeePerBytePerTimeSet", logs: logs, sub: sub}, nil
}

// WatchFeePerBytePerTimeSet is a free log subscription operation binding the contract event 0xcd1b2c2a220284accd1f9effd811cdecb6beaa4638618b48bbea07ce7ae16996.
//
// Solidity: event FeePerBytePerTimeSet(uint256 previousValue, uint256 newValue)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchFeePerBytePerTimeSet(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerFeePerBytePerTimeSet) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "FeePerBytePerTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerFeePerBytePerTimeSet)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "FeePerBytePerTimeSet", log); err != nil {
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

// ParseFeePerBytePerTimeSet is a log parse operation binding the contract event 0xcd1b2c2a220284accd1f9effd811cdecb6beaa4638618b48bbea07ce7ae16996.
//
// Solidity: event FeePerBytePerTimeSet(uint256 previousValue, uint256 newValue)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseFeePerBytePerTimeSet(log types.Log) (*DataLayrServiceManagerFeePerBytePerTimeSet, error) {
	event := new(DataLayrServiceManagerFeePerBytePerTimeSet)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "FeePerBytePerTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerFirstQuorumThresholdPercentageSetIterator is returned from FilterFirstQuorumThresholdPercentageSet and is used to iterate over the raw logs and unpacked data for FirstQuorumThresholdPercentageSet events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerFirstQuorumThresholdPercentageSetIterator struct {
	Event *DataLayrServiceManagerFirstQuorumThresholdPercentageSet // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerFirstQuorumThresholdPercentageSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerFirstQuorumThresholdPercentageSet)
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
		it.Event = new(DataLayrServiceManagerFirstQuorumThresholdPercentageSet)
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
func (it *DataLayrServiceManagerFirstQuorumThresholdPercentageSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerFirstQuorumThresholdPercentageSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerFirstQuorumThresholdPercentageSet represents a FirstQuorumThresholdPercentageSet event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerFirstQuorumThresholdPercentageSet struct {
	PreviousThreshold *big.Int
	NewThreshold      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFirstQuorumThresholdPercentageSet is a free log retrieval operation binding the contract event 0x4458887b3b2c02a432e79083729e6cc9971e4d1d5f7186ce31a79b5716523472.
//
// Solidity: event FirstQuorumThresholdPercentageSet(uint256 previousThreshold, uint256 newThreshold)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterFirstQuorumThresholdPercentageSet(opts *bind.FilterOpts) (*DataLayrServiceManagerFirstQuorumThresholdPercentageSetIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "FirstQuorumThresholdPercentageSet")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerFirstQuorumThresholdPercentageSetIterator{contract: _DataLayrServiceManager.contract, event: "FirstQuorumThresholdPercentageSet", logs: logs, sub: sub}, nil
}

// WatchFirstQuorumThresholdPercentageSet is a free log subscription operation binding the contract event 0x4458887b3b2c02a432e79083729e6cc9971e4d1d5f7186ce31a79b5716523472.
//
// Solidity: event FirstQuorumThresholdPercentageSet(uint256 previousThreshold, uint256 newThreshold)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchFirstQuorumThresholdPercentageSet(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerFirstQuorumThresholdPercentageSet) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "FirstQuorumThresholdPercentageSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerFirstQuorumThresholdPercentageSet)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "FirstQuorumThresholdPercentageSet", log); err != nil {
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

// ParseFirstQuorumThresholdPercentageSet is a log parse operation binding the contract event 0x4458887b3b2c02a432e79083729e6cc9971e4d1d5f7186ce31a79b5716523472.
//
// Solidity: event FirstQuorumThresholdPercentageSet(uint256 previousThreshold, uint256 newThreshold)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseFirstQuorumThresholdPercentageSet(log types.Log) (*DataLayrServiceManagerFirstQuorumThresholdPercentageSet, error) {
	event := new(DataLayrServiceManagerFirstQuorumThresholdPercentageSet)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "FirstQuorumThresholdPercentageSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerInitDataStoreIterator is returned from FilterInitDataStore and is used to iterate over the raw logs and unpacked data for InitDataStore events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerInitDataStoreIterator struct {
	Event *DataLayrServiceManagerInitDataStore // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerInitDataStoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerInitDataStore)
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
		it.Event = new(DataLayrServiceManagerInitDataStore)
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
func (it *DataLayrServiceManagerInitDataStoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerInitDataStoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerInitDataStore represents a InitDataStore event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerInitDataStore struct {
	FeePayer   common.Address
	SearchData IDataLayrServiceManagerDataStoreSearchData
	Header     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitDataStore is a free log retrieval operation binding the contract event 0x24c9bb2620686dbe734e777d7b61fc8c088e12fb07e265a6bb90f9f9e0896012.
//
// Solidity: event InitDataStore(address feePayer, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, bytes header)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterInitDataStore(opts *bind.FilterOpts) (*DataLayrServiceManagerInitDataStoreIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "InitDataStore")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerInitDataStoreIterator{contract: _DataLayrServiceManager.contract, event: "InitDataStore", logs: logs, sub: sub}, nil
}

// WatchInitDataStore is a free log subscription operation binding the contract event 0x24c9bb2620686dbe734e777d7b61fc8c088e12fb07e265a6bb90f9f9e0896012.
//
// Solidity: event InitDataStore(address feePayer, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, bytes header)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchInitDataStore(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerInitDataStore) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "InitDataStore")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerInitDataStore)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "InitDataStore", log); err != nil {
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

// ParseInitDataStore is a log parse operation binding the contract event 0x24c9bb2620686dbe734e777d7b61fc8c088e12fb07e265a6bb90f9f9e0896012.
//
// Solidity: event InitDataStore(address feePayer, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, bytes header)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseInitDataStore(log types.Log) (*DataLayrServiceManagerInitDataStore, error) {
	event := new(DataLayrServiceManagerInitDataStore)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "InitDataStore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerPausedIterator struct {
	Event *DataLayrServiceManagerPaused // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerPaused)
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
		it.Event = new(DataLayrServiceManagerPaused)
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
func (it *DataLayrServiceManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerPaused represents a Paused event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*DataLayrServiceManagerPausedIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerPausedIterator{contract: _DataLayrServiceManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerPaused) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerPaused)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParsePaused(log types.Log) (*DataLayrServiceManagerPaused, error) {
	event := new(DataLayrServiceManagerPaused)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerPaymentManagerSetIterator is returned from FilterPaymentManagerSet and is used to iterate over the raw logs and unpacked data for PaymentManagerSet events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerPaymentManagerSetIterator struct {
	Event *DataLayrServiceManagerPaymentManagerSet // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerPaymentManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerPaymentManagerSet)
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
		it.Event = new(DataLayrServiceManagerPaymentManagerSet)
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
func (it *DataLayrServiceManagerPaymentManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerPaymentManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerPaymentManagerSet represents a PaymentManagerSet event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerPaymentManagerSet struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaymentManagerSet is a free log retrieval operation binding the contract event 0xa3044efb81dffce20bbf49cae117f167852a973364ae504dfade51a8d022c95a.
//
// Solidity: event PaymentManagerSet(address previousAddress, address newAddress)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterPaymentManagerSet(opts *bind.FilterOpts) (*DataLayrServiceManagerPaymentManagerSetIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "PaymentManagerSet")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerPaymentManagerSetIterator{contract: _DataLayrServiceManager.contract, event: "PaymentManagerSet", logs: logs, sub: sub}, nil
}

// WatchPaymentManagerSet is a free log subscription operation binding the contract event 0xa3044efb81dffce20bbf49cae117f167852a973364ae504dfade51a8d022c95a.
//
// Solidity: event PaymentManagerSet(address previousAddress, address newAddress)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchPaymentManagerSet(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerPaymentManagerSet) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "PaymentManagerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerPaymentManagerSet)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "PaymentManagerSet", log); err != nil {
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

// ParsePaymentManagerSet is a log parse operation binding the contract event 0xa3044efb81dffce20bbf49cae117f167852a973364ae504dfade51a8d022c95a.
//
// Solidity: event PaymentManagerSet(address previousAddress, address newAddress)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParsePaymentManagerSet(log types.Log) (*DataLayrServiceManagerPaymentManagerSet, error) {
	event := new(DataLayrServiceManagerPaymentManagerSet)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "PaymentManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerQuorumThresholdBasisPointsUpdateIterator is returned from FilterQuorumThresholdBasisPointsUpdate and is used to iterate over the raw logs and unpacked data for QuorumThresholdBasisPointsUpdate events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerQuorumThresholdBasisPointsUpdateIterator struct {
	Event *DataLayrServiceManagerQuorumThresholdBasisPointsUpdate // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerQuorumThresholdBasisPointsUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerQuorumThresholdBasisPointsUpdate)
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
		it.Event = new(DataLayrServiceManagerQuorumThresholdBasisPointsUpdate)
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
func (it *DataLayrServiceManagerQuorumThresholdBasisPointsUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerQuorumThresholdBasisPointsUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerQuorumThresholdBasisPointsUpdate represents a QuorumThresholdBasisPointsUpdate event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerQuorumThresholdBasisPointsUpdate struct {
	QuorumThresholdBasisPoints uint16
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterQuorumThresholdBasisPointsUpdate is a free log retrieval operation binding the contract event 0xf301989e9074791776ece0876bacf50f53acf6ef124c99a76284c7473b5d459b.
//
// Solidity: event QuorumThresholdBasisPointsUpdate(uint16 quorumThresholdBasisPoints)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterQuorumThresholdBasisPointsUpdate(opts *bind.FilterOpts) (*DataLayrServiceManagerQuorumThresholdBasisPointsUpdateIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "QuorumThresholdBasisPointsUpdate")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerQuorumThresholdBasisPointsUpdateIterator{contract: _DataLayrServiceManager.contract, event: "QuorumThresholdBasisPointsUpdate", logs: logs, sub: sub}, nil
}

// WatchQuorumThresholdBasisPointsUpdate is a free log subscription operation binding the contract event 0xf301989e9074791776ece0876bacf50f53acf6ef124c99a76284c7473b5d459b.
//
// Solidity: event QuorumThresholdBasisPointsUpdate(uint16 quorumThresholdBasisPoints)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchQuorumThresholdBasisPointsUpdate(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerQuorumThresholdBasisPointsUpdate) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "QuorumThresholdBasisPointsUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerQuorumThresholdBasisPointsUpdate)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "QuorumThresholdBasisPointsUpdate", log); err != nil {
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

// ParseQuorumThresholdBasisPointsUpdate is a log parse operation binding the contract event 0xf301989e9074791776ece0876bacf50f53acf6ef124c99a76284c7473b5d459b.
//
// Solidity: event QuorumThresholdBasisPointsUpdate(uint16 quorumThresholdBasisPoints)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseQuorumThresholdBasisPointsUpdate(log types.Log) (*DataLayrServiceManagerQuorumThresholdBasisPointsUpdate, error) {
	event := new(DataLayrServiceManagerQuorumThresholdBasisPointsUpdate)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "QuorumThresholdBasisPointsUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerSecondQuorumThresholdPercentageSetIterator is returned from FilterSecondQuorumThresholdPercentageSet and is used to iterate over the raw logs and unpacked data for SecondQuorumThresholdPercentageSet events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerSecondQuorumThresholdPercentageSetIterator struct {
	Event *DataLayrServiceManagerSecondQuorumThresholdPercentageSet // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerSecondQuorumThresholdPercentageSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerSecondQuorumThresholdPercentageSet)
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
		it.Event = new(DataLayrServiceManagerSecondQuorumThresholdPercentageSet)
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
func (it *DataLayrServiceManagerSecondQuorumThresholdPercentageSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerSecondQuorumThresholdPercentageSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerSecondQuorumThresholdPercentageSet represents a SecondQuorumThresholdPercentageSet event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerSecondQuorumThresholdPercentageSet struct {
	PreviousThreshold *big.Int
	NewThreshold      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSecondQuorumThresholdPercentageSet is a free log retrieval operation binding the contract event 0x74450771d766586a9e6988dd3926b008a37d3b3fd1f7298f44f72ba3efab1a3a.
//
// Solidity: event SecondQuorumThresholdPercentageSet(uint256 previousThreshold, uint256 newThreshold)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterSecondQuorumThresholdPercentageSet(opts *bind.FilterOpts) (*DataLayrServiceManagerSecondQuorumThresholdPercentageSetIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "SecondQuorumThresholdPercentageSet")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerSecondQuorumThresholdPercentageSetIterator{contract: _DataLayrServiceManager.contract, event: "SecondQuorumThresholdPercentageSet", logs: logs, sub: sub}, nil
}

// WatchSecondQuorumThresholdPercentageSet is a free log subscription operation binding the contract event 0x74450771d766586a9e6988dd3926b008a37d3b3fd1f7298f44f72ba3efab1a3a.
//
// Solidity: event SecondQuorumThresholdPercentageSet(uint256 previousThreshold, uint256 newThreshold)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchSecondQuorumThresholdPercentageSet(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerSecondQuorumThresholdPercentageSet) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "SecondQuorumThresholdPercentageSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerSecondQuorumThresholdPercentageSet)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "SecondQuorumThresholdPercentageSet", log); err != nil {
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

// ParseSecondQuorumThresholdPercentageSet is a log parse operation binding the contract event 0x74450771d766586a9e6988dd3926b008a37d3b3fd1f7298f44f72ba3efab1a3a.
//
// Solidity: event SecondQuorumThresholdPercentageSet(uint256 previousThreshold, uint256 newThreshold)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseSecondQuorumThresholdPercentageSet(log types.Log) (*DataLayrServiceManagerSecondQuorumThresholdPercentageSet, error) {
	event := new(DataLayrServiceManagerSecondQuorumThresholdPercentageSet)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "SecondQuorumThresholdPercentageSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerSignatoryRecordIterator is returned from FilterSignatoryRecord and is used to iterate over the raw logs and unpacked data for SignatoryRecord events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerSignatoryRecordIterator struct {
	Event *DataLayrServiceManagerSignatoryRecord // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerSignatoryRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerSignatoryRecord)
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
		it.Event = new(DataLayrServiceManagerSignatoryRecord)
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
func (it *DataLayrServiceManagerSignatoryRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerSignatoryRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerSignatoryRecord represents a SignatoryRecord event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerSignatoryRecord struct {
	MsgHash                 [32]byte
	TaskNumber              uint32
	SignedStakeFirstQuorum  *big.Int
	SignedStakeSecondQuorum *big.Int
	PubkeyHashes            [][32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSignatoryRecord is a free log retrieval operation binding the contract event 0x34d57e230be557a52d94166eb9035810e61ac973182a92b09e6b0e99110665a9.
//
// Solidity: event SignatoryRecord(bytes32 msgHash, uint32 taskNumber, uint256 signedStakeFirstQuorum, uint256 signedStakeSecondQuorum, bytes32[] pubkeyHashes)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterSignatoryRecord(opts *bind.FilterOpts) (*DataLayrServiceManagerSignatoryRecordIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "SignatoryRecord")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerSignatoryRecordIterator{contract: _DataLayrServiceManager.contract, event: "SignatoryRecord", logs: logs, sub: sub}, nil
}

// WatchSignatoryRecord is a free log subscription operation binding the contract event 0x34d57e230be557a52d94166eb9035810e61ac973182a92b09e6b0e99110665a9.
//
// Solidity: event SignatoryRecord(bytes32 msgHash, uint32 taskNumber, uint256 signedStakeFirstQuorum, uint256 signedStakeSecondQuorum, bytes32[] pubkeyHashes)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchSignatoryRecord(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerSignatoryRecord) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "SignatoryRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerSignatoryRecord)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "SignatoryRecord", log); err != nil {
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

// ParseSignatoryRecord is a log parse operation binding the contract event 0x34d57e230be557a52d94166eb9035810e61ac973182a92b09e6b0e99110665a9.
//
// Solidity: event SignatoryRecord(bytes32 msgHash, uint32 taskNumber, uint256 signedStakeFirstQuorum, uint256 signedStakeSecondQuorum, bytes32[] pubkeyHashes)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseSignatoryRecord(log types.Log) (*DataLayrServiceManagerSignatoryRecord, error) {
	event := new(DataLayrServiceManagerSignatoryRecord)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "SignatoryRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerUnpausedIterator struct {
	Event *DataLayrServiceManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerUnpaused)
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
		it.Event = new(DataLayrServiceManagerUnpaused)
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
func (it *DataLayrServiceManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerUnpaused represents a Unpaused event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*DataLayrServiceManagerUnpausedIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerUnpausedIterator{contract: _DataLayrServiceManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerUnpaused)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseUnpaused(log types.Log) (*DataLayrServiceManagerUnpaused, error) {
	event := new(DataLayrServiceManagerUnpaused)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogIterator struct {
	Event *DataLayrServiceManagerLog // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLog)
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
		it.Event = new(DataLayrServiceManagerLog)
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
func (it *DataLayrServiceManagerLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLog represents a Log event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLog(opts *bind.FilterOpts) (*DataLayrServiceManagerLogIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogIterator{contract: _DataLayrServiceManager.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLog) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLog)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLog(log types.Log) (*DataLayrServiceManagerLog, error) {
	event := new(DataLayrServiceManagerLog)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogAddressIterator struct {
	Event *DataLayrServiceManagerLogAddress // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogAddress)
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
		it.Event = new(DataLayrServiceManagerLogAddress)
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
func (it *DataLayrServiceManagerLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogAddress represents a LogAddress event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogAddress(opts *bind.FilterOpts) (*DataLayrServiceManagerLogAddressIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogAddressIterator{contract: _DataLayrServiceManager.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogAddress) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogAddress)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_address", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogAddress(log types.Log) (*DataLayrServiceManagerLogAddress, error) {
	event := new(DataLayrServiceManagerLogAddress)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogBytesIterator struct {
	Event *DataLayrServiceManagerLogBytes // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogBytes)
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
		it.Event = new(DataLayrServiceManagerLogBytes)
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
func (it *DataLayrServiceManagerLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogBytes represents a LogBytes event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogBytes(opts *bind.FilterOpts) (*DataLayrServiceManagerLogBytesIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogBytesIterator{contract: _DataLayrServiceManager.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogBytes) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogBytes)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_bytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogBytes(log types.Log) (*DataLayrServiceManagerLogBytes, error) {
	event := new(DataLayrServiceManagerLogBytes)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogBytes32Iterator struct {
	Event *DataLayrServiceManagerLogBytes32 // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogBytes32)
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
		it.Event = new(DataLayrServiceManagerLogBytes32)
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
func (it *DataLayrServiceManagerLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogBytes32 represents a LogBytes32 event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*DataLayrServiceManagerLogBytes32Iterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogBytes32Iterator{contract: _DataLayrServiceManager.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogBytes32) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogBytes32)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogBytes32(log types.Log) (*DataLayrServiceManagerLogBytes32, error) {
	event := new(DataLayrServiceManagerLogBytes32)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogIntIterator struct {
	Event *DataLayrServiceManagerLogInt // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogInt)
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
		it.Event = new(DataLayrServiceManagerLogInt)
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
func (it *DataLayrServiceManagerLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogInt represents a LogInt event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogInt(opts *bind.FilterOpts) (*DataLayrServiceManagerLogIntIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogIntIterator{contract: _DataLayrServiceManager.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogInt) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogInt)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_int", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogInt(log types.Log) (*DataLayrServiceManagerLogInt, error) {
	event := new(DataLayrServiceManagerLogInt)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedAddressIterator struct {
	Event *DataLayrServiceManagerLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogNamedAddress)
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
		it.Event = new(DataLayrServiceManagerLogNamedAddress)
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
func (it *DataLayrServiceManagerLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogNamedAddress represents a LogNamedAddress event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*DataLayrServiceManagerLogNamedAddressIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogNamedAddressIterator{contract: _DataLayrServiceManager.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogNamedAddress)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_address", log); err != nil {
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

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogNamedAddress(log types.Log) (*DataLayrServiceManagerLogNamedAddress, error) {
	event := new(DataLayrServiceManagerLogNamedAddress)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedBytesIterator struct {
	Event *DataLayrServiceManagerLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogNamedBytes)
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
		it.Event = new(DataLayrServiceManagerLogNamedBytes)
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
func (it *DataLayrServiceManagerLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogNamedBytes represents a LogNamedBytes event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*DataLayrServiceManagerLogNamedBytesIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogNamedBytesIterator{contract: _DataLayrServiceManager.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogNamedBytes)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogNamedBytes(log types.Log) (*DataLayrServiceManagerLogNamedBytes, error) {
	event := new(DataLayrServiceManagerLogNamedBytes)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedBytes32Iterator struct {
	Event *DataLayrServiceManagerLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogNamedBytes32)
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
		it.Event = new(DataLayrServiceManagerLogNamedBytes32)
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
func (it *DataLayrServiceManagerLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogNamedBytes32 represents a LogNamedBytes32 event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*DataLayrServiceManagerLogNamedBytes32Iterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogNamedBytes32Iterator{contract: _DataLayrServiceManager.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogNamedBytes32)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogNamedBytes32(log types.Log) (*DataLayrServiceManagerLogNamedBytes32, error) {
	event := new(DataLayrServiceManagerLogNamedBytes32)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedDecimalIntIterator struct {
	Event *DataLayrServiceManagerLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogNamedDecimalInt)
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
		it.Event = new(DataLayrServiceManagerLogNamedDecimalInt)
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
func (it *DataLayrServiceManagerLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*DataLayrServiceManagerLogNamedDecimalIntIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogNamedDecimalIntIterator{contract: _DataLayrServiceManager.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogNamedDecimalInt)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogNamedDecimalInt(log types.Log) (*DataLayrServiceManagerLogNamedDecimalInt, error) {
	event := new(DataLayrServiceManagerLogNamedDecimalInt)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedDecimalUintIterator struct {
	Event *DataLayrServiceManagerLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogNamedDecimalUint)
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
		it.Event = new(DataLayrServiceManagerLogNamedDecimalUint)
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
func (it *DataLayrServiceManagerLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*DataLayrServiceManagerLogNamedDecimalUintIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogNamedDecimalUintIterator{contract: _DataLayrServiceManager.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogNamedDecimalUint)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogNamedDecimalUint(log types.Log) (*DataLayrServiceManagerLogNamedDecimalUint, error) {
	event := new(DataLayrServiceManagerLogNamedDecimalUint)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedIntIterator struct {
	Event *DataLayrServiceManagerLogNamedInt // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogNamedInt)
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
		it.Event = new(DataLayrServiceManagerLogNamedInt)
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
func (it *DataLayrServiceManagerLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogNamedInt represents a LogNamedInt event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*DataLayrServiceManagerLogNamedIntIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogNamedIntIterator{contract: _DataLayrServiceManager.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogNamedInt)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_int", log); err != nil {
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

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogNamedInt(log types.Log) (*DataLayrServiceManagerLogNamedInt, error) {
	event := new(DataLayrServiceManagerLogNamedInt)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedStringIterator struct {
	Event *DataLayrServiceManagerLogNamedString // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogNamedString)
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
		it.Event = new(DataLayrServiceManagerLogNamedString)
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
func (it *DataLayrServiceManagerLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogNamedString represents a LogNamedString event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*DataLayrServiceManagerLogNamedStringIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogNamedStringIterator{contract: _DataLayrServiceManager.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogNamedString) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogNamedString)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_string", log); err != nil {
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

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogNamedString(log types.Log) (*DataLayrServiceManagerLogNamedString, error) {
	event := new(DataLayrServiceManagerLogNamedString)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedUintIterator struct {
	Event *DataLayrServiceManagerLogNamedUint // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogNamedUint)
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
		it.Event = new(DataLayrServiceManagerLogNamedUint)
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
func (it *DataLayrServiceManagerLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogNamedUint represents a LogNamedUint event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*DataLayrServiceManagerLogNamedUintIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogNamedUintIterator{contract: _DataLayrServiceManager.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogNamedUint)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogNamedUint(log types.Log) (*DataLayrServiceManagerLogNamedUint, error) {
	event := new(DataLayrServiceManagerLogNamedUint)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogStringIterator struct {
	Event *DataLayrServiceManagerLogString // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogString)
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
		it.Event = new(DataLayrServiceManagerLogString)
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
func (it *DataLayrServiceManagerLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogString represents a LogString event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogString(opts *bind.FilterOpts) (*DataLayrServiceManagerLogStringIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogStringIterator{contract: _DataLayrServiceManager.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogString) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogString)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_string", log); err != nil {
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

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogString(log types.Log) (*DataLayrServiceManagerLogString, error) {
	event := new(DataLayrServiceManagerLogString)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogUintIterator struct {
	Event *DataLayrServiceManagerLogUint // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogUint)
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
		it.Event = new(DataLayrServiceManagerLogUint)
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
func (it *DataLayrServiceManagerLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogUint represents a LogUint event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogUint(opts *bind.FilterOpts) (*DataLayrServiceManagerLogUintIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogUintIterator{contract: _DataLayrServiceManager.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogUint) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogUint)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_uint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogUint(log types.Log) (*DataLayrServiceManagerLogUint, error) {
	event := new(DataLayrServiceManagerLogUint)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataLayrServiceManagerLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogsIterator struct {
	Event *DataLayrServiceManagerLogs // Event containing the contract specifics and raw log

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
func (it *DataLayrServiceManagerLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLayrServiceManagerLogs)
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
		it.Event = new(DataLayrServiceManagerLogs)
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
func (it *DataLayrServiceManagerLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLayrServiceManagerLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLayrServiceManagerLogs represents a Logs event raised by the DataLayrServiceManager contract.
type DataLayrServiceManagerLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) FilterLogs(opts *bind.FilterOpts) (*DataLayrServiceManagerLogsIterator, error) {

	logs, sub, err := _DataLayrServiceManager.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &DataLayrServiceManagerLogsIterator{contract: _DataLayrServiceManager.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *DataLayrServiceManagerLogs) (event.Subscription, error) {

	logs, sub, err := _DataLayrServiceManager.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLayrServiceManagerLogs)
				if err := _DataLayrServiceManager.contract.UnpackLog(event, "logs", log); err != nil {
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

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_DataLayrServiceManager *DataLayrServiceManagerFilterer) ParseLogs(log types.Log) (*DataLayrServiceManagerLogs, error) {
	event := new(DataLayrServiceManagerLogs)
	if err := _DataLayrServiceManager.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

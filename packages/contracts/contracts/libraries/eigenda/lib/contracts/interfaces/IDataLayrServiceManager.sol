// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IServiceManager.sol";
import "./IEigenLayrDelegation.sol";
import "./IDataLayrPaymentManager.sol";

interface IDataLayrServiceManager is IServiceManager {
    //Relevant metadata for a given datastore
    struct DataStoreMetadata {
        bytes32 headerHash;
        uint32 durationDataStoreId;
        uint32 globalDataStoreId;
        uint32 blockNumber;
        uint96 fee;
        address confirmer;
        bytes32 signatoryRecordHash;
    }

    //Stores the data required to index a given datastore's metadata
    struct DataStoreSearchData {
        DataStoreMetadata metadata;
        uint8 duration;
        uint256 timestamp;
        uint32 index;
    }

    struct SignatoryRecordMinusDataStoreId {
        bytes32[] nonSignerPubkeyHashes;
        uint256 signedStakeFirstQuorum;
        uint256 signedStakeSecondQuorum;
    }

    struct DataStoresForDuration {
        uint32 one_duration;
        uint32 two_duration;
        uint32 three_duration;
        uint32 four_duration;
        uint32 five_duration;
        uint32 six_duration;
        uint32 seven_duration;
        uint32 dataStoreId;
        uint32 latestTime;
    }

    struct DataStoreHashInputs {
        bytes32 headerHash;
        uint32 dataStoreId;
        uint32 blockNumber;
        uint256 fee;
    }

    /**
     * @notice This function is used for
     * - notifying via Ethereum that the disperser has asserted the data blob
     * into DataLayr and is waiting to obtain quorum of DataLayr operators to sign,
     * - asserting the metadata corresponding to the data asserted into DataLayr
     * - escrow the service fees that DataLayr operators will receive from the disperser
     * on account of their service.
     *
     * This function returns the index of the data blob in dataStoreIdsForDuration[duration][block.timestamp]
     */
    /**
     * @param feePayer is the address that will be paying the fees for this datastore. check DataLayrPaymentManager for further details
     * @param confirmer is the address that must confirm the datastore
     * @param header is the summary of the data that is being asserted into DataLayr,
     *  type DataStoreHeader struct {
     *   KzgCommit      [64]byte
     *   Degree         uint32 
     *   NumSys         uint32
     *   NumPar         uint32
     *   OrigDataSize   uint32 
     *   Disperser      [20]byte
     *   LowDegreeProof [64]byte 
     *  }
     * @param duration for which the data has to be stored by the DataLayr operators.
     * This is a quantized parameter that describes how many factors of DURATION_SCALE
     * does this data blob needs to be stored. The quantization process comes from ease of
     * implementation in DataLayrBombVerifier.sol.
     * @param blockNumber is the block number in Ethereum for which the confirmation will
     * consult total + operator stake amounts.
     * -- must not be more than 'BLOCK_STALE_MEASURE' (defined in DataLayr) blocks in past
     * @return index The index in the array `dataStoreHashesForDurationAtTimestamp[duration][block.timestamp]` at which the DataStore's hash was stored.
     */
    function initDataStore(
        address feePayer,
        address confirmer,
        uint8 duration,
        uint32 blockNumber,
        uint32 totalOperatorsIndex,
        bytes calldata header
    )
        external
        returns (uint32);

    /**
     * @notice This function is used for
     * - disperser to notify that signatures on the message, comprising of hash( headerHash ),
     * from quorum of DataLayr nodes have been obtained,
     * - check that the aggregate signature is valid,
     * - and check whether quorum has been achieved or not.
     */
    /**
     * @param data Input to the `checkSignatures` function, which is of the format:
     * <
     * bytes32 msgHash,
     * uint48 index of the totalStake corresponding to the dataStoreId in the 'totalStakeHistory' array of the BLSRegistryWithBomb
     * uint32 numberOfNonSigners,
     * uint256[numberOfSigners][4] pubkeys of nonsigners,
     * uint32 apkIndex,
     * uint256[4] apk,
     * uint256[2] sigma
     * >
     */
    function confirmDataStore(bytes calldata data, DataStoreSearchData memory searchData) external;

    /// @notice number of leaves in the root tree
    function numPowersOfTau() external view returns (uint48);

    /// @notice number of layers in the root tree
    function log2NumPowersOfTau() external view returns (uint48);

    /// @notice Unit of measure (in time) for the duration of DataStores
    function DURATION_SCALE() external view returns (uint256);

    /// @notice The longest allowed duation of a DataStore, measured in `DURATION_SCALE`
    function MAX_DATASTORE_DURATION() external view returns (uint8);

    /// @notice Returns the hash of the `index`th DataStore with the specified `duration` at the specified UTC `timestamp`.
    function getDataStoreHashesForDurationAtTimestamp(uint8 duration, uint256 timestamp, uint32 index)
        external
        view
        returns (bytes32);

    /**
     * @notice returns the number of data stores for the @param duration
     */
    function getNumDataStoresForDuration(uint8 duration) external view returns (uint32);

    /// @notice Collateral token used for placing collateral on challenges & payment commits
    function collateralToken() external view returns (IERC20);

    /**
     * @notice contract used for handling payment challenges
     */
    function dataLayrPaymentManager() external view returns (IDataLayrPaymentManager);

    /**
     * @notice Checks that the hash of the `index`th DataStore with the specified `duration` at the specified UTC `timestamp` matches the supplied `metadata`.
     * Returns 'true' if the metadata matches the hash, and 'false' otherwise.
     */
   function verifyDataStoreMetadata(uint8 duration, uint256 timestamp, uint32 index, DataStoreMetadata memory metadata) external view returns (bool);
}

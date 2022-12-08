// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "../../interfaces/IRepository.sol";
import "../../interfaces/IEigenLayrDelegation.sol";
import "../../interfaces/IDelegationTerms.sol";

import "./DataLayrServiceManagerStorage.sol";
import "../BLSSignatureChecker.sol";

import "../../libraries/BytesLib.sol";
import "../../libraries/Merkle.sol";
import "../../libraries/DataStoreUtils.sol";
import "../../permissions/Pausable.sol";

import "../Repository.sol";
import "./DataLayrChallengeUtils.sol";
import "hardhat/console.sol";

/**
 * @title Primary entrypoint for procuring services from DataLayr.
 * @author Layr Labs, Inc.
 * @notice This contract is used for:
 * - initializing the data store by the disperser
 * - confirming the data store by the disperser with inferred aggregated signatures of the quorum
 * - freezing operators as the result of various "challenges"
 */
contract DataLayrServiceManager is DataLayrServiceManagerStorage, BLSSignatureChecker, Pausable, DSTest {
    using BytesLib for bytes;
    // sanity checks. should always require *some* signatures, but never *all* signatures
    uint128 internal constant MIN_THRESHOLD_PERCENTAGE = 1;
    uint128 internal constant MAX_THRESHOLD_PERCENTAGE = 99;

    //quorumThresholdBasisPoints is the minimum basis points of total registered operators that must sign the datastore
    uint16 public quorumThresholdBasisPoints = 9000;

    /**
    * adversaryThresholdBasisPoints is the maximum basis points of total registered
    * operators that witholds their chunks before the data can no longer be reconstructed
    * TODO: Change for prod!
    */
    uint16 public adversaryThresholdBasisPoints = 4000;

    uint128 public firstQuorumThresholdPercentage;
    uint128 public secondQuorumThresholdPercentage;

    /// @notice Keeps track of the number of DataStores for each duration, the total number of DataStores, and the `latestTime` until which operators must serve.
    DataStoresForDuration public dataStoresForDuration;

    // EVENTS
    event InitDataStore(
        address feePayer,
        IDataLayrServiceManager.DataStoreSearchData searchData,
        bytes header
    );

    /**
     * @notice Emitted when a DataStore is confirmed.
     * @param dataStoreId The ID for the DataStore inside of the specified duration (i.e. *not* the globalDataStoreId)
     * @param headerHash The headerHash of the DataStore.
     */
    event ConfirmDataStore(uint32 dataStoreId, bytes32 headerHash);

    event QuorumThresholdBasisPointsUpdate(uint16 quorumThresholdBasisPoints);
    event AdversaryThresholdBasisPointsUpdated(uint16 adversaryThresholdBasisPoints);

    modifier checkValidThresholds(uint16 _quorumThresholdBasisPoints, uint16 _adversaryThresholdBasisPoints) {
        require(_quorumThresholdBasisPoints > _adversaryThresholdBasisPoints,
            "DataLayrServiceManager.validThresholds: Quorum threshold must be strictly greater than adversary");
        _;
    }

    event FeePerBytePerTimeSet(uint256 previousValue, uint256 newValue);

    event BombVerifierSet(address previousAddress, address newAddress);

    event PaymentManagerSet(address previousAddress, address newAddress);

    event EphemeralKeyRegistrySet(address previousAddress, address newAddress);

    event FirstQuorumThresholdPercentageSet(uint256 previousThreshold, uint256 newThreshold);

    event SecondQuorumThresholdPercentageSet(uint256 previousThreshold, uint256 newThreshold);

    constructor(
        IInvestmentManager _investmentManager,
        IEigenLayrDelegation _eigenLayrDelegation,
        IRepository _repository,
        IERC20 _collateralToken,
        IPauserRegistry _pauserRegistry,
        uint256 _feePerBytePerTime
    )
        DataLayrServiceManagerStorage(_investmentManager, _eigenLayrDelegation, _collateralToken)
        BLSSignatureChecker(_repository)
        checkValidThresholds(quorumThresholdBasisPoints, adversaryThresholdBasisPoints)
    {
        _setFeePerBytePerTime(_feePerBytePerTime);
        dataStoresForDuration.dataStoreId = 1;
        dataStoresForDuration.latestTime = 1;
        _initializePauser(_pauserRegistry);
        // set default values
        _setFirstQuorumThresholdPercentage(90);
        _setSecondQuorumThresholdPercentage(90);
    }

    function setLowDegreeChallenge(DataLayrLowDegreeChallenge _dataLayrLowDegreeChallenge)
        external
        onlyRepositoryGovernance
    {
        dataLayrLowDegreeChallenge = _dataLayrLowDegreeChallenge;
    }

    function setBombVerifier(DataLayrBombVerifier _dataLayrBombVerifier) external onlyRepositoryGovernance {
        emit BombVerifierSet(address(dataLayrBombVerifier), address(_dataLayrBombVerifier));
        dataLayrBombVerifier = _dataLayrBombVerifier;
    }

    function setPaymentManager(DataLayrPaymentManager _dataLayrPaymentManager) external onlyRepositoryGovernance {
        emit PaymentManagerSet(address(dataLayrPaymentManager), address(_dataLayrPaymentManager));
        dataLayrPaymentManager = _dataLayrPaymentManager;
    }

    function setEphemeralKeyRegistry(EphemeralKeyRegistry _ephemeralKeyRegistry) external onlyRepositoryGovernance {
        emit EphemeralKeyRegistrySet(address(ephemeralKeyRegistry), address(_ephemeralKeyRegistry));
        ephemeralKeyRegistry = _ephemeralKeyRegistry;
    }

    function setQuorumThresholdBasisPoints(uint16 _quorumThresholdBasisPoints)
        external
        onlyRepositoryGovernance
        checkValidThresholds(_quorumThresholdBasisPoints, adversaryThresholdBasisPoints)
    {
        quorumThresholdBasisPoints = _quorumThresholdBasisPoints;
        emit QuorumThresholdBasisPointsUpdate(quorumThresholdBasisPoints);
    }

    function setAdversaryThresholdBasisPoints(uint16 _adversaryThresholdBasisPoints)
        external
        onlyRepositoryGovernance
        checkValidThresholds(quorumThresholdBasisPoints, _adversaryThresholdBasisPoints)
    {
        adversaryThresholdBasisPoints = _adversaryThresholdBasisPoints;
        emit AdversaryThresholdBasisPointsUpdated(adversaryThresholdBasisPoints);
    }


    function setFirstQuorumThresholdPercentage(uint128 _firstQuorumThresholdPercentage) external onlyRepositoryGovernance {
        _setFirstQuorumThresholdPercentage(_firstQuorumThresholdPercentage);
    }

    function setSecondQuorumThresholdPercentage(uint128 _secondQuorumThresholdPercentage) external onlyRepositoryGovernance {
        _setSecondQuorumThresholdPercentage(_secondQuorumThresholdPercentage);
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
        whenNotPaused
        returns (uint32 index)
    {
        console.log("dlsm:01010101010101010100101010");
        bytes32 headerHash = keccak256(header);
        uint32 storePeriodLength;
        IDataLayrServiceManager.DataStoreMetadata memory metadata;
        console.log("dlsm:0000000000000000000000000");
        {
            uint256 totalBytes;
            {
                //fetch the total number of operators for the blockNumber from which stakes are being read from
                uint32 totalOperators = IQuorumRegistry(address(_registry())).getTotalOperators(blockNumber, totalOperatorsIndex);

                totalBytes = DataStoreUtils.getTotalBytes(header, totalOperators);

                require(totalBytes >= MIN_STORE_SIZE, "DataLayrServiceManager.initDataStore: totalBytes < MIN_STORE_SIZE");
                require(totalBytes <= MAX_STORE_SIZE, "DataLayrServiceManager.initDataStore: totalBytes > MAX_STORE_SIZE");


                /**
                * @notice coding ratio is numSys/numOperators (where numOperators = numSys + numPar).  This is the minimum
                *   percentage of all chunks require to reconstruct the data.
                *
                * quorumThresholdBasisPoints is the minimum percentage of total registered operators that must sign the datastore
                * adversaryThresholdBasisPoints is the maximum percentage of total registered operators that witholds their chunks
                *    before the data can no longer be reconstructed.
                *
                * adversaryThresholdBasisPoints <  quorumThresholdBasisPoints, there cannot be more dishonest signers than actual signers
                *
                * quorumThresholdBasisPoints - adversaryThresholdBasisPoints represents the minimum percentage
                *   of operators that must be honest signers. This value must be greater than or equal to the coding ratio
                *   in order to ensure the data is available.
                */
                require(quorumThresholdBasisPoints - adversaryThresholdBasisPoints >= DataStoreUtils.getCodingRatio(header, totalOperators), "DataLayrServiceManager.initDataStore: Coding ratio is too high");

            }

            require(duration >= 1 && duration <= MAX_DATASTORE_DURATION, "DataLayrServiceManager.initDataStore: Invalid duration");

            // compute time and fees
            // computing the actual period for which data blob needs to be stored
            storePeriodLength = uint32(duration * DURATION_SCALE);

            // evaluate the total service fees that msg.sender has to put in escrow for paying out
            // the DataLayr nodes for their service

            uint256 fee = calculateFee(totalBytes, feePerBytePerTime, storePeriodLength);


            // require that disperser has sent enough fees to this contract to pay for this datastore.
            // This will revert if the deposits are not high enough due to undeflow.
            dataLayrPaymentManager.payFee(msg.sender, feePayer, fee);

            // Recording the initialization of datablob store along with auxiliary info
            //store metadata locally to be stored
            metadata = IDataLayrServiceManager.DataStoreMetadata({
                headerHash: headerHash,
                durationDataStoreId: getNumDataStoresForDuration(duration),
                globalDataStoreId: dataStoresForDuration.dataStoreId,
                blockNumber: blockNumber,
                fee: uint96(fee),
                confirmer: confirmer,
                signatoryRecordHash: bytes32(0)
            });
         }
         console.log("dlsm:111111111111111");


        /**
         * Stores the hash of the datastore's metadata into the `dataStoreHashesForDurationAtTimestamp` mapping.
         * We iterate through the mapping and store the hash in the first available empty storage slot.
         * This hash is stored to be checked during the quorum signature verification, ensuring that the correct dataStore is signed and confirmed.
         */
        {
            // uint g = gasleft();
            //iterate the index throughout the loop
            for (; index < NUM_DS_PER_BLOCK_PER_DURATION; index++) {
                if (dataStoreHashesForDurationAtTimestamp[duration][block.timestamp][index] == 0) {
                    dataStoreHashesForDurationAtTimestamp[duration][block.timestamp][index] =
                        DataStoreUtils.computeDataStoreHash(metadata);
                    // recording the empty slot
                    break;
                }
            }

            // reverting we looped through all of the indecies without finding an empty element
            require(
                index != NUM_DS_PER_BLOCK_PER_DURATION,
                "DataLayrServiceManager.initDataStore: number of initDatastores for this duration and block has reached its limit"
            );
        }
        console.log("dlsm:2222222222222222222");
        // sanity check on blockNumber
        {
            require(
                blockNumber <= block.number, "DataLayrServiceManager.initDataStore: specified blockNumber is in future"
            );

            require(
                (blockNumber + BLOCK_STALE_MEASURE) >= block.number,
                "DataLayrServiceManager.initDataStore: specified blockNumber is too far in past"
            );
        }
        console.log("dlsm:333333333333333333333333333333");
        IDataLayrServiceManager.DataStoreSearchData memory searchData = IDataLayrServiceManager.DataStoreSearchData({
            duration: duration,
            timestamp: block.timestamp,
            index: index,
            metadata: metadata
        });
        console.log("dlsm:444444444444444444444444444444");
        // emit event to represent initialization of data store
        emit InitDataStore(feePayer, searchData, header);

        // Updating dataStoresForDuration
        /**
         * @notice sets the latest time until which any of the active DataLayr operators that haven't committed
         * yet to deregistration are supposed to serve.
         */
        // recording the expiry time until which the DataLayr operators, who sign up to
        // part of the quorum, have to store the data
        uint32 _latestTime = uint32(block.timestamp) + storePeriodLength;
        if (_latestTime > dataStoresForDuration.latestTime) {
            dataStoresForDuration.latestTime = _latestTime;
        }
        console.log("dlsm:5555555555555555555");
        // increments the number of datastores for the specific duration of the asserted DataStore
        _incrementDataStoresForDuration(duration);

        // increment the counter
        ++dataStoresForDuration.dataStoreId;

        console.log("dlsm:66666666666666666666");
        return index;
    }

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
    function confirmDataStore(bytes calldata data, DataStoreSearchData memory searchData) external whenNotPaused {
        /**
         * Verify that the signatures provided by the disperser are indeed from DataLayr operators who have agreed to be in the quorum.
         * Additionally, pull relevant information from the provided `data` param, which we subsequently check the integrity of.
         */
        (
            uint32 dataStoreIdToConfirm,
            uint32 blockNumberFromTaskHash,
            bytes32 msgHash,
            SignatoryTotals memory signedTotals,
            bytes32 signatoryRecordHash
        ) = checkSignatures(data);

        /**
         * Make sure that the nodes signed the hash of dsid, headerHash, duration, timestamp, and index to avoid malleability in case of reorgs.
         * This keeps bomb and storage conditions fixed (to a single blockchain fork).
         */
        require(
            msgHash
                == keccak256(
                    abi.encodePacked(
                        dataStoreIdToConfirm,
                        searchData.metadata.headerHash,
                        searchData.duration,
                        searchData.timestamp,
                        searchData.index
                    )
                ),
            "DataLayrServiceManager.confirmDataStore: msgHash is not consistent with search data"
        );

        //make sure the address confirming is the prespecified `confirmer`
        require(
            msg.sender == searchData.metadata.confirmer,
            "DataLayrServiceManager.confirmDataStore: Sender is not authorized to confirm this datastore"
        );
        // check that the DataStore has not already been confirmed
        require(
            searchData.metadata.signatoryRecordHash == bytes32(0),
            "DataLayrServiceManager.confirmDataStore: SignatoryRecord must be bytes32(0)"
        );
        // verify integrity of `dataStoreIdToConfirm` provided as part of `data` input
        require(
            searchData.metadata.globalDataStoreId == dataStoreIdToConfirm,
            "DataLayrServiceManager.confirmDataStore: gloabldatastoreid is does not agree with data"
        );
        // verify integrity of `blockNumberFromTaskHash` provided as part of `data` input
        require(
            searchData.metadata.blockNumber == blockNumberFromTaskHash,
            "DataLayrServiceManager.confirmDataStore: blocknumber does not agree with data"
        );

        //Check if provided calldata matches the hash stored in dataStoreIDsForDuration in initDataStore
        //verify consistency of signed data with stored data
        bytes32 dsHash = DataStoreUtils.computeDataStoreHash(searchData.metadata);

        require(
            dataStoreHashesForDurationAtTimestamp[searchData.duration][searchData.timestamp][searchData.index] == dsHash,
            "DataLayrServiceManager.confirmDataStore: provided calldata does not match corresponding stored hash from initDataStore"
        );

        // add the computed signatoryRecordHash to the searchData
        searchData.metadata.signatoryRecordHash = signatoryRecordHash;

        // computing a new DataStoreIdsForDuration hash that includes the signatory record as well
        bytes32 newDsHash = DataStoreUtils.computeDataStoreHash(searchData.metadata);

        //storing new hash that now includes the signatory record
        dataStoreHashesForDurationAtTimestamp[searchData.duration][searchData.timestamp][searchData.index] = newDsHash;



        // check that signatories own at least a threshold percentage of the two stake sets (i.e. eth & eigen) implying quorum has been achieved
        require(
            (signedTotals.signedStakeFirstQuorum * BIP_MULTIPLIER) / signedTotals.totalStakeFirstQuorum
                >= quorumThresholdBasisPoints
                && (signedTotals.signedStakeSecondQuorum * BIP_MULTIPLIER) / signedTotals.totalStakeSecondQuorum
                    >= quorumThresholdBasisPoints,
            "DataLayrServiceManager.confirmDataStore: signatories do not own at least threshold percentage of both quorums"
        );



        emit ConfirmDataStore(dataStoresForDuration.dataStoreId, searchData.metadata.headerHash);

    }

    // called in the event of challenge resolution
    function freezeOperator(address operator) external {
        require(
            msg.sender == address(dataLayrLowDegreeChallenge)
                || msg.sender == address(dataLayrBombVerifier)
                || msg.sender == address(ephemeralKeyRegistry)
                || msg.sender == address(dataLayrPaymentManager),
            "DataLayrServiceManager.freezeOperator: Only challenge resolvers can slash operators"
        );
        ISlasher(investmentManager.slasher()).freezeOperator(operator);
    }

    // called in the event of deregistration
    function revokeSlashingAbility(address operator, uint32 unbondedAfter) external onlyRegistry {
        ISlasher(investmentManager.slasher()).revokeSlashingAbility(operator, unbondedAfter);
    }

    function setFeePerBytePerTime(uint256 _feePerBytePerTime) external onlyRepositoryGovernance {
        _setFeePerBytePerTime(_feePerBytePerTime);
    }

    // VIEW FUNCTIONS
    /**
     * @notice Checks that the hash of the `index`th DataStore with the specified `duration` at the specified UTC `timestamp` matches the supplied `metadata`.
     * Returns 'true' if the metadata matches the hash, and 'false' otherwise.
     */
    function verifyDataStoreMetadata(
        uint8 duration,
        uint256 timestamp,
        uint32 index,
        DataStoreMetadata memory metadata
    )
        external
        view
        returns (bool)
    {
        return(
            getDataStoreHashesForDurationAtTimestamp(
                duration,
                timestamp,
                index
            ) == DataStoreUtils.computeDataStoreHash(metadata)
        );
    }

    /// @notice Returns the hash of the `index`th DataStore with the specified `duration` at the specified UTC `timestamp`.
    function getDataStoreHashesForDurationAtTimestamp(uint8 duration, uint256 timestamp, uint32 index)
        public
        view
        returns (bytes32)
    {
        return dataStoreHashesForDurationAtTimestamp[duration][timestamp][index];
    }

    /**
     * @notice returns the number of data stores for the @param duration
     */
    function getNumDataStoresForDuration(uint8 duration) public view returns (uint32) {
        if (duration == 1) {
            return dataStoresForDuration.one_duration;
        }
        if (duration == 2) {
            return dataStoresForDuration.two_duration;
        }
        if (duration == 3) {
            return dataStoresForDuration.three_duration;
        }
        if (duration == 4) {
            return dataStoresForDuration.four_duration;
        }
        if (duration == 5) {
            return dataStoresForDuration.five_duration;
        }
        if (duration == 6) {
            return dataStoresForDuration.six_duration;
        }
        if (duration == 7) {
            return dataStoresForDuration.seven_duration;
        }
        revert("DataLayrServiceManager.getNumDataStoresForDuration: invalid duration");
    }

    function taskNumber() external view returns (uint32) {
        return dataStoresForDuration.dataStoreId;
    }

    /**
     * @notice Verifies that a DataStore exists which was created *at or before* `initTimestamp` *AND* that expires *strictly prior to* the
     * specified `unlockTime`.
     * @dev Function reverts if the verification fails.
     * @param packedDataStoreSearchData should be the same format as the output of `DataStoreUtils.packDataStoreSearchData(dataStoreSearchData)`
     */
    function stakeWithdrawalVerification(
        bytes calldata packedDataStoreSearchData,
        uint256 initTimestamp,
        uint256 unlockTime
    )
        external
        view
    {
        IDataLayrServiceManager.DataStoreSearchData memory searchData =
            DataStoreUtils.unpackDataStoreSearchData(packedDataStoreSearchData);
        bytes32 dsHash = DataStoreUtils.computeDataStoreHash(searchData.metadata);
        require(
            dataStoreHashesForDurationAtTimestamp[searchData.duration][searchData.timestamp][searchData.index] == dsHash,
            "DataLayrServiceManager.stakeWithdrawalVerification: provided calldata does not match corresponding stored hash from (initDataStore)"
        );

        /**
         * Now we check that the specified DataStore was created *at or before*  the `initTimestamp`, i.e. when the user undelegated, deregistered, etc. *AND*
         * that the user's funds are set to unlock *prior* to the expiration of the DataStore.
         * In other words, we are checking that a user was active when the specified DataStore was created, and is trying to unstake/undelegate/etc. funds prior
         * to them fully serving out their commitment to storing their share of the data.
         */
        require(
            (initTimestamp >= searchData.timestamp)
                && (unlockTime < searchData.timestamp + (searchData.duration * DURATION_SCALE)),
            "DataLayrServiceManager.stakeWithdrawalVerification: task does not meet requirements"
        );
    }

    /// @notice Returns the `latestTime` until which operators must serve.
    function latestTime() external view returns (uint32) {
        return dataStoresForDuration.latestTime;
    }

    // INTERNAL FUNTIONS

    /**
     * @notice increments the number of data stores for the @param duration
     */
    function _incrementDataStoresForDuration(uint8 duration) internal {
        if (duration == 1) {
            ++dataStoresForDuration.one_duration;
        }
        if (duration == 2) {
            ++dataStoresForDuration.two_duration;
        }
        if (duration == 3) {
            ++dataStoresForDuration.three_duration;
        }
        if (duration == 4) {
            ++dataStoresForDuration.four_duration;
        }
        if (duration == 5) {
            ++dataStoresForDuration.five_duration;
        }
        if (duration == 6) {
            ++dataStoresForDuration.six_duration;
        }
        if (duration == 7) {
            ++dataStoresForDuration.seven_duration;
        }
    }

    function calculateFee(uint256 totalBytes, uint256 _feePerBytePerTime, uint32 storePeriodLength)
        public
        pure
        returns (uint256)
    {
        return uint256(totalBytes * _feePerBytePerTime * storePeriodLength);
    }

    function _setFeePerBytePerTime(uint256 _feePerBytePerTime) internal {
        emit FeePerBytePerTimeSet(feePerBytePerTime, _feePerBytePerTime);
        feePerBytePerTime = _feePerBytePerTime;
    }

    function _setFirstQuorumThresholdPercentage(uint128 _firstQuorumThresholdPercentage) internal {
        require(
            _firstQuorumThresholdPercentage >= MIN_THRESHOLD_PERCENTAGE,
            "DataLayrServiceManager.setFirstQuorumThresholdPercentage: input too low"
        );
        require(
            _firstQuorumThresholdPercentage <= MAX_THRESHOLD_PERCENTAGE,
            "DataLayrServiceManager.setFirstQuorumThresholdPercentage: input too high"
        );
        emit FirstQuorumThresholdPercentageSet(firstQuorumThresholdPercentage, _firstQuorumThresholdPercentage);
        firstQuorumThresholdPercentage = _firstQuorumThresholdPercentage;
    }

    function _setSecondQuorumThresholdPercentage(uint128 _secondQuorumThresholdPercentage) internal {
        require(
            _secondQuorumThresholdPercentage >= MIN_THRESHOLD_PERCENTAGE,
            "DataLayrServiceManager.setSecondQuorumThresholdPercentage: input too low"
        );
        require(
            _secondQuorumThresholdPercentage <= MAX_THRESHOLD_PERCENTAGE,
            "DataLayrServiceManager.setSecondQuorumThresholdPercentage: input too high"
        );
        emit SecondQuorumThresholdPercentageSet(secondQuorumThresholdPercentage, _secondQuorumThresholdPercentage);
        secondQuorumThresholdPercentage = _secondQuorumThresholdPercentage;
    }

}

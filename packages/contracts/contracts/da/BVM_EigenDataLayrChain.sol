// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";
import { DataLayrDisclosureLogic } from "../libraries/eigenda/DataLayrDisclosureLogic.sol";
import { IDataLayrServiceManager } from "../libraries/eigenda/lib/contracts/interfaces/IDataLayrServiceManager.sol";
import { BN254 } from "../libraries/eigenda/BN254.sol";
import { DataStoreUtils } from "../libraries/eigenda/lib/contracts/libraries/DataStoreUtils.sol";
import { Parser } from "../libraries/eigenda/Parse.sol";


contract BVM_EigenDataLayrChain is OwnableUpgradeable, ReentrancyGuardUpgradeable, Parser {
    using SafeMathUpgradeable for uint256;
    using AddressUpgradeable for address;

    enum RollupStoreStatus {
        UNCOMMITTED,
        COMMITTED,
        REVERTED
    }

    struct DisclosureProofs {
        bytes header;
        uint32 firstChunkNumber;
        bytes[] polys;
        DataLayrDisclosureLogic.MultiRevealProof[] multiRevealProofs;
        BN254.G2Point polyEquivalenceProof;
    }

    uint256 public SUBMISSION_INTERVAL;
    address public sequencer;
    address public dataManageAddress;
    uint256 public BLOCK_STALE_MEASURE;
    uint256 public l2BlockNumber;
    uint256 public fraudProofPeriod = 1 days;

    bytes public constant FRAUD_STRING = '-_(` O `)_- -_(` o `)_- -_(` Q `)_- BITDAO JUST REKT YOU |_(` O `)_| - |_(` o `)_| - |_(` Q `)_|';

    uint256 internal constant DATA_STORE_INITIALIZED_BUT_NOT_CONFIRMED = type(uint256).max;

    struct RollupStore {
        uint32 dataStoreId;
        uint32 confirmAt;
        RollupStoreStatus status;
    }

    //mapping from the rollup's store id to datastore id
    mapping(uint256 => RollupStore) public rollupStores;
    /**
     * @notice mapping used to track whether or not this contract initiated specific dataStores, as well as
     * to track how the link between dataStoreId and rollupStoreNumber
     * @dev We use this so we don't create a subgraph temporarily
     */
    mapping(uint32 => uint256) public dataStoreIdToRollupStoreNumber;
    uint256 public rollupStoreNumber;

    event RollupStoreInitialized(uint32 dataStoreId);
    event RollupStoreConfirmed(uint32 rollupStoreNumber);
    event RollupStoreReverted(uint32 rollupStoreNumber);

    function initialize(address _sequencer, address _dataManageAddress, uint256 _submissionInterval, uint256 _block_stale_measure) public initializer {
        __Ownable_init();
        sequencer = _sequencer;
        dataManageAddress = _dataManageAddress;
        SUBMISSION_INTERVAL = _submissionInterval;
        BLOCK_STALE_MEASURE = _block_stale_measure;
        l2BlockNumber = 0;
    }

    /**
     * @notice Returns the block number of the latest submitted L2.
     * If no submitted yet then this function will return the starting block number.
     *
     * @return Latest submitted L2 block number.
     */
    function latestBlockNumber() public view returns (uint256) {
        return l2BlockNumber;
    }

    /**
     * @notice Computes the block number of the next L2 block that needs to be checkpointed.
     *
     * @return Next L2 block number.
     */
    function nextBlockNumber() public view returns (uint256) {
        return latestBlockNumber() + SUBMISSION_INTERVAL;
    }

    /**
     * @notice Called by the (staked) sequencer to pay for a datastore and post some metadata (in the `header` parameter) about it on chain.
     * Since the sequencer must encode the data before they post the header on chain, they must use a *snapshot* of the number and stakes of DataLayr operators
     * from a previous block number, specified by the `blockNumber` input.
     * @param header of data to be stored
     * @param duration is the duration to store the datastore for
     * @param blockNumber is the previous block number which was used to encode the data for storage
     * @param totalOperatorsIndex is index in the totalOperators array of DataLayr referring to what the total number of operators was at `blockNumber`
     * @dev The specified `blockNumber `must be less than `BLOCK_STALE_MEASURE` blocks in the past.
     */
    function storeData(
        bytes calldata header,
        uint8 duration,
        uint32 blockNumber,
        uint256 _l2BlockNumber,
        uint32 totalOperatorsIndex
    ) external {
        require(msg.sender == sequencer, "Only the sequencer can store data");

        require(block.number - blockNumber < BLOCK_STALE_MEASURE, "stakes taken from too long ago");
        uint32 dataStoreId = IDataLayrServiceManager(dataManageAddress).taskNumber();
        l2BlockNumber = _l2BlockNumber;
        //Initialize and pay for the datastore
        IDataLayrServiceManager(dataManageAddress).initDataStore(
            msg.sender,
            address(this),
            duration,
            blockNumber,
            totalOperatorsIndex,
            header
        );
        dataStoreIdToRollupStoreNumber[dataStoreId] = DATA_STORE_INITIALIZED_BUT_NOT_CONFIRMED;
        emit RollupStoreInitialized(dataStoreId);
    }

    /**
     * @notice After the `storeData `transaction is included in a block and doesn’t revert, the sequencer will disperse the data to the DataLayr nodes off chain
     * and get their signatures that they have stored the data. Now, the sequencer has to post the signature on chain and get it verified.
     * @param data Input of the header information for a dataStore and signatures for confirming the dataStore -- used as input to the `confirmDataStore` function
     * of the DataLayrServiceManager -- see the DataLayr docs for more info on this.
     * @param searchData Data used to specify the dataStore being confirmed. Must be provided so other contracts can properly look up the dataStore.
     * @dev Only dataStores created through this contract can be confirmed by calling this function.
     */
    function confirmData(
        bytes calldata data,
        IDataLayrServiceManager.DataStoreSearchData memory searchData
    ) external {
        require(msg.sender == sequencer, "Only the sequencer can store data");
        require(
            dataStoreIdToRollupStoreNumber[searchData.metadata.globalDataStoreId] ==
            DATA_STORE_INITIALIZED_BUT_NOT_CONFIRMED,
            "Data store either was not initialized by the rollup contract, or is already confirmed"
        );
        IDataLayrServiceManager(dataManageAddress).confirmDataStore(data, searchData);
        //store the rollups view of the datastore
        rollupStores[rollupStoreNumber] = RollupStore({
        dataStoreId: searchData.metadata.globalDataStoreId,
        confirmAt: uint32(block.timestamp + fraudProofPeriod),
        status: RollupStoreStatus.COMMITTED
        });
        //store link between dataStoreId and rollupStoreNumber
        dataStoreIdToRollupStoreNumber[searchData.metadata.globalDataStoreId] = rollupStoreNumber;
        emit RollupStoreConfirmed(uint32(rollupStoreNumber++));
    }

    /**
  * @notice Called by a challenger (this could be anyone -- "challenger" is not a permissioned role) to prove that fraud has occurred.
     * First, a subset of data included in a dataStore that was initiated by the sequencer is proven, and then the presence of fraud in the data is checked.
     * For the sake of this example, "fraud occurring" means that the sequencer included the forbidden `FRAUD_STRING` in a dataStore that they initiated.
     * In pratical use, "fraud occurring" might mean including data that specifies an invalid transaction or invalid state transition.
     * @param fraudulentStoreNumber The *rollupStoreNumber* to prove fraud on
     * @param startIndex The index to begin reading the proven data from
     * @param searchData Data used to specify the dataStore being fraud-proven. Must be provided so other contracts can properly look up the dataStore.
     * @param disclosureProofs Non-interactive polynomial proofs that prove that the specific data of interest was part of the dataStore in question.
     * @dev This function is only callable if:
     * -the sequencer is staked,
     * -the dataStore in question has been confirmed, and
     * -the fraudproof period for the dataStore has not yet passed.
     */
    function proveFraud(
        uint256 fraudulentStoreNumber,
        uint256 startIndex,
        IDataLayrServiceManager.DataStoreSearchData memory searchData,
        DisclosureProofs calldata disclosureProofs
    ) public {
        RollupStore memory rollupStore = rollupStores[fraudulentStoreNumber];
        require(rollupStore.status == RollupStoreStatus.COMMITTED && rollupStore.confirmAt > block.timestamp, "RollupStore must be committed and unconfirmed");
        //verify that the provided metadata is correct for the challenged data store
        require(
            IDataLayrServiceManager(dataManageAddress).getDataStoreHashesForDurationAtTimestamp(
                searchData.duration,
                searchData.timestamp,
                searchData.index
            ) == DataStoreUtils.computeDataStoreHash(searchData.metadata),
            "metadata preimage is incorrect"
        );
        //make sure search data, disclosure proof, and rollupstore are all consistent with each other
        require(searchData.metadata.globalDataStoreId == rollupStore.dataStoreId, "seachData's datastore id is not consistent with given rollup store");
        require(searchData.metadata.headerHash == keccak256(disclosureProofs.header), "disclosure proofs headerhash preimage is incorrect");
        //verify that all of the provided polynomials are in fact part of the data
        require(DataLayrDisclosureLogic.batchNonInteractivePolynomialProofs(
                disclosureProofs.header,
                disclosureProofs.firstChunkNumber,
                disclosureProofs.polys,
                disclosureProofs.multiRevealProofs,
                disclosureProofs.polyEquivalenceProof
            ), "disclosure proofs are invalid");


        // get the number of systematic symbols from the header
        uint32 numSys = DataLayrDisclosureLogic.getNumSysFromHeader(disclosureProofs.header);
        require(disclosureProofs.firstChunkNumber + disclosureProofs.polys.length <= numSys, "Can only prove data from the systematic chunks");
        //parse proven data
        bytes memory provenString = parse(disclosureProofs.polys, startIndex, FRAUD_STRING.length);
        //sanity check
        require(provenString.length == FRAUD_STRING.length, "Parsing error, proven string is different length than fraud string");
        //check whether provenString == FRAUD_STRING
        require(keccak256(provenString) == keccak256(FRAUD_STRING), "proven string != fraud string");
        //slash sequencer because fraud is proven
        rollupStores[fraudulentStoreNumber].status = RollupStoreStatus.REVERTED;
        emit RollupStoreReverted(uint32(fraudulentStoreNumber));
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { DataLayrDisclosureLogic } from "../libraries/eigenda/DataLayrDisclosureLogic.sol";
import { IDataLayrServiceManager } from "../libraries/eigenda/lib/contracts/interfaces/IDataLayrServiceManager.sol";
import { BN254 } from "../libraries/eigenda/BN254.sol";


// This contract is a rollup built on top of datalayr that allows an enshrined sequencer to store data
// on datalayr, but slashes them if the data they store ever contains a certain unallowed message
contract EigenDataLayrChain {
    enum RollupStoreStatus {
        UNCOMMITTED,
        COMMITTED,
        REVERTED
    }

    //the enshrined address that stores data on datalayr for the rollup
    address public sequencer;
    IERC20 public immutable stakeToken;
    uint256 public immutable neededStake;
    uint256 public immutable BLOCK_STALE_MEASURE;
    uint256 public immutable fraudProofPeriod = 1 days;

    uint256 internal constant DATA_STORE_INITIALIZED_BUT_NOT_CONFIRMED = type(uint256).max;

    bytes32 internal constant _DESIGNATED_IMPLEMENTATION_SLOT = 0xa1832b4681bf2b3269fd2d0163abf215d243dd098d1ece37dd0775a7f2a8c09d;

    struct RollupStore {
        //The corresponding datastore in DataLayr to the rollup block number
        uint32 dataStoreId;
        //The time at which the rollup block will be confirmed if not successfully challenged
        uint32 confirmAt;
        //the status of the rollup store
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
    // The datalayr service manager, the contract where datastores are initialized and confirmed
    IDataLayrServiceManager public dlsm;
    //TODO: Link this to DataLayr interface instead of redeclaration
    struct DisclosureProofs {
        bytes header;
        uint32 firstChunkNumber;
        bytes[] polys;
        DataLayrDisclosureLogic.MultiRevealProof[] multiRevealProofs;
        BN254.G2Point polyEquivalenceProof;
    }

    event RollupStoreInitialized(uint32 dataStoreId);
    event RollupStoreConfirmed(uint32 rollupStoreNumber);
    event RollupStoreReverted(uint32 rollupStoreNumber);

    //initialize the storage vars
    constructor(
        address _sequencer,
        IERC20 _stakeToken,
        uint256 _neededStake,
        IDataLayrServiceManager _dlsm
    ) {
        sequencer = _sequencer;
        stakeToken = _stakeToken;
        dlsm = _dlsm;
        neededStake = _neededStake;
        BLOCK_STALE_MEASURE = 100;
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
        uint32 totalOperatorsIndex
    ) external {
        require(msg.sender == sequencer, "Only the sequencer can store data");

        require(block.number - blockNumber < BLOCK_STALE_MEASURE, "stakes taken from too long ago");
        uint32 dataStoreId = dlsm.taskNumber();
        //Initialize and pay for the datastore
        dlsm.initDataStore(
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
        dlsm.confirmDataStore(data, searchData);
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
}

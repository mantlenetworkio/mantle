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
import "hardhat/console.sol";


contract BVM_EigenDataLayrChainUpgrade is OwnableUpgradeable, ReentrancyGuardUpgradeable, Parser {
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

    address public sequencer;
    address public dataManageAddress;
    uint256 public BLOCK_STALE_MEASURE;
    uint256 public l2StoredBlockNumber;
    uint256 public l2ConfirmedBlockNumber;
    uint256 public fraudProofPeriod;
    uint256 public rollupBatchIndex;

    bytes public constant FRAUD_STRING = '-_(` O `)_- -_(` o `)_- -_(` Q `)_- BITDAO JUST REKT YOU |_(` O `)_| - |_(` o `)_| - |_(` Q `)_|';
    uint256 internal constant DATA_STORE_INITIALIZED_BUT_NOT_CONFIRMED = type(uint256).max;

    struct RollupStore {
        uint32 originDataStoreId;
        uint32 dataStoreId;
        uint32 confirmAt;
        RollupStoreStatus status;
    }

    struct BatchRollupBlock {
        uint256 startL2BlockNumber;
        uint256 endBL2BlockNumber;
        bool    isReRollup;
    }

    mapping(uint256 => RollupStore) public rollupBatchIndexRollupStores;
    mapping(uint32 => BatchRollupBlock) public dataStoreIdToL2RollUpBlock;
    mapping(uint32 => uint256) public dataStoreIdToRollupStoreNumber;
    mapping(address => bool) private fraudProofWhitelist;


    event RollupStoreInitialized(uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber);
    event RollupStoreConfirmed(uint256 rollupBatchIndex, uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber);
    event RollupStoreReverted(uint256 rollupBatchIndex, uint32 dataStoreId, uint256 stratL2BlockNumber, uint256 endL2BlockNumber);

    function initialize(address _sequencer, address _dataManageAddress, uint256 _block_stale_measure, uint256 _fraudProofPeriod, uint256 _l2SubmittedBlockNumber) public initializer {
        __Ownable_init();
        sequencer = _sequencer;
        dataManageAddress = _dataManageAddress;
        BLOCK_STALE_MEASURE = _block_stale_measure;
        fraudProofPeriod = _fraudProofPeriod;
        l2StoredBlockNumber = _l2SubmittedBlockNumber;
        l2ConfirmedBlockNumber = _l2SubmittedBlockNumber;
    }

    /**
     * @notice Returns the block number of the latest stored L2.
     * @return Latest stored L2 block number.
     */
    function getL2StoredBlockNumber() public view returns (uint256) {
        return l2StoredBlockNumber;
    }

    /**
     * @notice Returns the block number of the latest stored L2.
     * @return Latest stored L2 block number.
     */
    function getL2ConfirmedBlockNumber() public view returns (uint256) {
        return l2ConfirmedBlockNumber;
    }

    /**
     * @notice Returns the rollup store by l2 block number
     * @return RollupStore.
     */
    function getRollupStoreByRollupBatchIndex(uint256 _rollupBatchIndex) public view returns (RollupStore memory) {
        return rollupBatchIndexRollupStores[_rollupBatchIndex];
    }

    /**
    * @notice Returns the l2 block number by store id
     * @return BatchRollupBlock.
     */
    function getL2RollUpBlockByDataStoreId(uint32 _dataStoreId) public view returns (BatchRollupBlock memory) {
        return dataStoreIdToL2RollUpBlock[_dataStoreId];
    }

    /**
    * @notice set fraud proof address
    * @param _address for fraud proof
    */
    function setFraudProofAddress(address _address) external {
        require(msg.sender == sequencer, "Only the sequencer can set fraud proof address unavailable");
        fraudProofWhitelist[_address] = true;
    }

    /**
    * @notice unavailable fraud proof address
    * @param _address for fraud proof
    */
    function unavailableFraudProofAddress(address _address) external {
        require(msg.sender == sequencer, "Only the sequencer can remove fraud proof address");
        fraudProofWhitelist[_address] = false;
    }

    /**
    * @notice remove fraud proof address
    * @param _address for fraud proof
    */
    function removeFraudProofAddress(address _address) external {
        require(msg.sender == sequencer, "Only the sequencer can remove fraud proof address");
        delete fraudProofWhitelist[_address];
    }

    /**
    * @notice update fraud proof period
    * @param _fraudProofPeriod fraud proof period
    */
    function updateFraudProofPeriod(uint256 _fraudProofPeriod) external {
        require(msg.sender == sequencer, "Only the sequencer can update fraud proof period");
        fraudProofPeriod = _fraudProofPeriod;
    }

    /**
    * @notice update dlsm address
    * @param _dataManageAddress dlsm address
    */
    function updateDataLayrManagerAddress(address _dataManageAddress) external {
        require(msg.sender == sequencer, "Only the sequencer can update dlsm address");
        dataManageAddress = _dataManageAddress;
    }

    /**
    * @notice update l2 latest store block number
    * @param _l2StoredBlockNumber l2 latest block number
    */
    function updateL2StoredBlockNumber(uint256 _l2StoredBlockNumber) external {
        require(msg.sender == sequencer, "Only the sequencer can set latest l2 block number");
        l2StoredBlockNumber = _l2StoredBlockNumber;
    }

    /**
    * @notice update l2 latest confirm block number
    * @param _l2ConfirmedBlockNumber l2 latest block number
    */
    function updateL2ConfirmedBlockNumber(uint256 _l2ConfirmedBlockNumber) external {
        require(msg.sender == sequencer, "Only the sequencer can set latest l2 block number");
        l2ConfirmedBlockNumber = _l2ConfirmedBlockNumber;
    }

    /**
    * @notice update sequencer address
    * @param _sequencer update sequencer address
    */
    function updateSequencerAddress(address _sequencer) external {
        require(msg.sender == sequencer, "Only the sequencer can update sequencer address");
        sequencer = _sequencer;
    }

    /**
    * @notice reset batch rollup batch data
    * @param _rollupBatchIndex update rollup index
    */
    function resetRollupBatchData(uint256 _rollupBatchIndex) external {
        require(msg.sender == sequencer, "Only the sequencer can update sequencer address");
        for (uint256 i = 0; i < rollupBatchIndex; i++) {
            delete rollupBatchIndexRollupStores[i];
        }
        rollupBatchIndex = _rollupBatchIndex;
        l2StoredBlockNumber = 1;
        l2ConfirmedBlockNumber = 1;
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
        uint256 startL2Block,
        uint256 endL2Block,
        uint32 totalOperatorsIndex,
        bool   isReRollup
    ) external {
        require(msg.sender == sequencer, "Only the sequencer can store data");
        require(block.number - blockNumber < BLOCK_STALE_MEASURE, "stakes taken from too long ago");
        uint32 dataStoreId = IDataLayrServiceManager(dataManageAddress).taskNumber();
        IDataLayrServiceManager(dataManageAddress).initDataStore(
            msg.sender,
            address(this),
            duration,
            blockNumber,
            totalOperatorsIndex,
            header
        );
        dataStoreIdToL2RollUpBlock[dataStoreId] = BatchRollupBlock({
        startL2BlockNumber: startL2Block,
        endBL2BlockNumber: endL2Block,
        isReRollup: isReRollup
        });
        dataStoreIdToRollupStoreNumber[dataStoreId] = DATA_STORE_INITIALIZED_BUT_NOT_CONFIRMED;
        if (!isReRollup) {
            l2StoredBlockNumber = endL2Block;
        }
        emit RollupStoreInitialized(dataStoreId, startL2Block, endL2Block);
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
        IDataLayrServiceManager.DataStoreSearchData memory searchData,
        uint256 startL2Block,
        uint256 endL2Block,
        uint32 originDataStoreId,
        uint256 reConfirmedBatchIndex,
        bool isReRollup
    ) external {
        require(msg.sender == sequencer, "Only the sequencer can store data");
        require(dataStoreIdToL2RollUpBlock[searchData.metadata.globalDataStoreId].startL2BlockNumber == startL2Block &&
        dataStoreIdToL2RollUpBlock[searchData.metadata.globalDataStoreId].endBL2BlockNumber == endL2Block &&
            dataStoreIdToL2RollUpBlock[searchData.metadata.globalDataStoreId].isReRollup == isReRollup,
            "Data store either was not initialized by the rollup contract, or is already confirmed"
        );
        require(
            dataStoreIdToRollupStoreNumber[searchData.metadata.globalDataStoreId] == DATA_STORE_INITIALIZED_BUT_NOT_CONFIRMED,
            "Data store either was not initialized by the rollup contract, or is already confirmed"
        );
        IDataLayrServiceManager(dataManageAddress).confirmDataStore(data, searchData);
        if (!isReRollup) {
            rollupBatchIndexRollupStores[rollupBatchIndex] = RollupStore({
            originDataStoreId: searchData.metadata.globalDataStoreId,
            dataStoreId: searchData.metadata.globalDataStoreId,
            confirmAt: uint32(block.timestamp + fraudProofPeriod),
            status: RollupStoreStatus.COMMITTED
            });
            l2ConfirmedBlockNumber = endL2Block;
            dataStoreIdToRollupStoreNumber[searchData.metadata.globalDataStoreId] = rollupBatchIndex;
            emit RollupStoreConfirmed(uint32(rollupBatchIndex++), searchData.metadata.globalDataStoreId, startL2Block, endL2Block);
        } else {
            rollupBatchIndexRollupStores[reConfirmedBatchIndex] = RollupStore({
            originDataStoreId: originDataStoreId,
            dataStoreId: searchData.metadata.globalDataStoreId,
            confirmAt: uint32(block.timestamp + fraudProofPeriod),
            status: RollupStoreStatus.COMMITTED
            });
            dataStoreIdToRollupStoreNumber[searchData.metadata.globalDataStoreId] = reConfirmedBatchIndex;
            emit RollupStoreConfirmed(reConfirmedBatchIndex, searchData.metadata.globalDataStoreId, startL2Block, endL2Block);
        }
    }

    /**
  * @notice Called by a challenger (this could be anyone -- "challenger" is not a permissioned role) to prove that fraud has occurred.
     * First, a subset of data included in a dataStore that was initiated by the sequencer is proven, and then the presence of fraud in the data is checked.
     * For the sake of this example, "fraud occurring" means that the sequencer included the forbidden `FRAUD_STRING` in a dataStore that they initiated.
     * In pratical use, "fraud occurring" might mean including data that specifies an invalid transaction or invalid state transition.
     * @param fraudulentStoreNumber The rollup l2Block to prove fraud on
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
    ) external {
        require(fraudProofWhitelist[msg.sender] == true, "Only fraud proof white list can challenge data");
        RollupStore memory rollupStore = rollupBatchIndexRollupStores[fraudulentStoreNumber];
        require(rollupStore.status == RollupStoreStatus.COMMITTED && rollupStore.confirmAt > block.timestamp, "RollupStore must be committed and unconfirmed");
        require(
            IDataLayrServiceManager(dataManageAddress).getDataStoreHashesForDurationAtTimestamp(
                searchData.duration,
                searchData.timestamp,
                searchData.index
            ) == DataStoreUtils.computeDataStoreHash(searchData.metadata),
            "metadata preimage is incorrect"
        );
        require(searchData.metadata.globalDataStoreId == rollupStore.dataStoreId, "seachData's datastore id is not consistent with given rollup store");
        require(searchData.metadata.headerHash == keccak256(disclosureProofs.header), "disclosure proofs headerhash preimage is incorrect");
        require(DataLayrDisclosureLogic.batchNonInteractivePolynomialProofs(
                disclosureProofs.header,
                disclosureProofs.firstChunkNumber,
                disclosureProofs.polys,
                disclosureProofs.multiRevealProofs,
                disclosureProofs.polyEquivalenceProof
            ), "disclosure proofs are invalid");
        uint32 numSys = DataLayrDisclosureLogic.getNumSysFromHeader(disclosureProofs.header);
        require(disclosureProofs.firstChunkNumber + disclosureProofs.polys.length <= numSys, "Can only prove data from the systematic chunks");
        bytes memory provenString = parse(disclosureProofs.polys, startIndex, FRAUD_STRING.length);
        require(provenString.length == FRAUD_STRING.length, "Parsing error, proven string is different length than fraud string");
        require(keccak256(provenString) == keccak256(FRAUD_STRING), "proven string != fraud string");
        rollupBatchIndexRollupStores[fraudulentStoreNumber].status = RollupStoreStatus.REVERTED;
        emit RollupStoreReverted(
            fraudulentStoreNumber,
            searchData.metadata.globalDataStoreId,
            dataStoreIdToL2RollUpBlock[searchData.metadata.globalDataStoreId].startL2BlockNumber,
            dataStoreIdToL2RollUpBlock[searchData.metadata.globalDataStoreId].endBL2BlockNumber
        );
    }
}

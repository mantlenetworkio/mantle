// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/* Library Imports */
import {Lib_BVMCodec} from "../../libraries/codec/Lib_BVMCodec.sol";
import {Lib_AddressResolver} from "../../libraries/resolver/Lib_AddressResolver.sol";
import {Lib_MerkleTree} from "../../libraries/utils/Lib_MerkleTree.sol";
import {CrossDomainEnabled} from "../../libraries/bridge/CrossDomainEnabled.sol";

/* Interface Imports */
import {IStateCommitmentChain} from "./IStateCommitmentChain.sol";
import {ICanonicalTransactionChain} from "./ICanonicalTransactionChain.sol";
import {IBondManager} from "../verification/IBondManager.sol";
import {IChainStorageContainer} from "./IChainStorageContainer.sol";
import {ITssGroupManager} from "../tss/ITssGroupManager.sol";
import {ITssRewardContract} from "../../L2/predeploys/iTssRewardContract.sol";

/**
 * @title StateCommitmentChain
 * @dev The State Commitment Chain (SCC) contract contains a list of proposed state roots which
 * Proposers assert to be a result of each transaction in the Canonical Transaction Chain (CTC).
 * Elements here have a 1:1 correspondence with transactions in the CTC, and should be the unique
 * state root calculated off-chain by applying the canonical transactions one by one.
 *
 */
contract StateCommitmentChain is IStateCommitmentChain, Lib_AddressResolver, CrossDomainEnabled {
    /*************
     * Constants *
     *************/

    uint256 public FRAUD_PROOF_WINDOW;
    uint256 public SEQUENCER_PUBLISH_WINDOW;

    /***************
     * Constructor *
     ***************/

    /**
     * @param _libAddressManager Address of the Address Manager.
     */
    constructor(
        address _libAddressManager,
        address _l1messenger,
        uint256 _fraudProofWindow,
        uint256 _sequencerPublishWindow
    ) Lib_AddressResolver(_libAddressManager) CrossDomainEnabled(address(0)) {
        messenger = _l1messenger;
        FRAUD_PROOF_WINDOW = _fraudProofWindow;
        SEQUENCER_PUBLISH_WINDOW = _sequencerPublishWindow;
    }

    /********************
     * Public Functions *
     ********************/

    /**
     * Accesses the batch storage container.
     * @return Reference to the batch storage container.
     */
    function batches() public view returns (IChainStorageContainer) {
        return IChainStorageContainer(resolve("ChainStorageContainer-SCC-batches"));
    }

    /**
     * @inheritdoc IStateCommitmentChain
     */
    function getTotalElements() public view returns (uint256 _totalElements) {
        (uint40 totalElements,) = _getBatchExtraData();
        return uint256(totalElements);
    }

    /**
     * @inheritdoc IStateCommitmentChain
     */
    function getTotalBatches() public view returns (uint256 _totalBatches) {
        return batches().length();
    }

    /**
     * @inheritdoc IStateCommitmentChain
     */
    // slither-disable-next-line external-function
    function getLastSequencerTimestamp() public view returns (uint256 _lastSequencerTimestamp) {
        (, uint40 lastSequencerTimestamp) = _getBatchExtraData();
        return uint256(lastSequencerTimestamp);
    }

    /**
     * @inheritdoc IStateCommitmentChain
     */
    // slither-disable-next-line external-function
    function appendStateBatch(bytes32[] memory _batch, uint256 _shouldStartAtElement, bytes memory _signature) public {
        // Fail fast in to make sure our batch roots aren't accidentally made fraudulent by the
        // publication of batches by some other user.
        require(
            _shouldStartAtElement == getTotalElements(),
            "Actual batch start index does not match expected start index."
        );

        // Proposers must have previously staked at the BondManager
        require(
            IBondManager(resolve("BondManager")).isCollateralized(msg.sender),
            "Proposer does not have enough collateral posted"
        );

        require(_batch.length > 0, "Cannot submit an empty state batch.");

        require(
            getTotalElements() + _batch.length <=
            ICanonicalTransactionChain(resolve("CanonicalTransactionChain")).getTotalElements(),
            "Number of state roots cannot exceed the number of canonical transactions."
        );

        // Call tss group register contract to verify the signature
        _checkClusterSignature(_batch, _shouldStartAtElement, _signature);

        // Pass the block's timestamp and the publisher of the data
        // to be used in the fraud proofs
        _appendBatch(_batch, _signature, abi.encode(block.timestamp, msg.sender));

        // Update distributed state batch, and emit message
        _distributeTssReward(_batch, _shouldStartAtElement);
    }

    /**
     * @inheritdoc IStateCommitmentChain
     */
    // slither-disable-next-line external-function
    function deleteStateBatch(Lib_BVMCodec.ChainBatchHeader memory _batchHeader) public {
        require(
            msg.sender == resolve("BVM_FraudVerifier"),
            "State batches can only be deleted by the BVM_FraudVerifier."
        );

        require(_isValidBatchHeader(_batchHeader), "Invalid batch header.");

        require(
            insideFraudProofWindow(_batchHeader),
            "State batches can only be deleted within the fraud proof window."
        );

        _deleteBatch(_batchHeader);
    }

    /**
     * @inheritdoc IStateCommitmentChain
     */
    // slither-disable-next-line external-function
    function verifyStateCommitment(
        bytes32 _element,
        Lib_BVMCodec.ChainBatchHeader memory _batchHeader,
        Lib_BVMCodec.ChainInclusionProof memory _proof
    ) public view returns (bool) {
        require(_isValidBatchHeader(_batchHeader), "Invalid batch header.");

        require(
            Lib_MerkleTree.verify(
                _batchHeader.batchRoot,
                _element,
                _proof.index,
                _proof.siblings,
                _batchHeader.batchSize
            ),
            "Invalid inclusion proof."
        );

        return true;
    }

    /**
     * @inheritdoc IStateCommitmentChain
     */
    function insideFraudProofWindow(Lib_BVMCodec.ChainBatchHeader memory _batchHeader)
    public
    view
    returns (bool _inside)
    {
        (uint256 timestamp,) = abi.decode(_batchHeader.extraData, (uint256, address));

        require(timestamp != 0, "Batch header timestamp cannot be zero");
        return (timestamp + FRAUD_PROOF_WINDOW) > block.timestamp;
    }

    /**
    * @inheritdoc IStateCommitmentChain
     */
    // slither-disable-next-line external-function
    function rollBackL2Chain(uint256 _shouldRollBack, uint256 _shouldStartAtElement, bytes memory _signature) public {
        // Fail fast in to make sure our batch roots aren't accidentally made fraudulent by the
        // publication of batches by some other user.
        require(
            _shouldStartAtElement == getTotalElements(),
            "Actual batch start index does not match expected start index."
        );

        // Proposers must have previously staked at the BondManager
        require(
            IBondManager(resolve("BondManager")).isCollateralized(msg.sender),
            "Proposer does not have enough collateral posted"
        );

        _checkRollBackSignature(_shouldRollBack,_signature);

    }


    /**
    * @inheritdoc IStateCommitmentChain
     */
    // slither-disable-next-line external-function
    function rollBackMessage(uint256 _shouldRollBack) public {}


    /**********************
     * Internal Functions *
     **********************/

    /**
     * Parses the batch context from the extra data.
     * @return Total number of elements submitted.
     * @return Timestamp of the last batch submitted by the sequencer.
     */
    function _getBatchExtraData() internal view returns (uint40, uint40) {
        bytes27 extraData = batches().getGlobalMetadata();

        // solhint-disable max-line-length
        uint40 totalElements;
        uint40 lastSequencerTimestamp;
        assembly {
            extraData := shr(40, extraData)
            totalElements := and(
            extraData,
            0x000000000000000000000000000000000000000000000000000000FFFFFFFFFF
            )
            lastSequencerTimestamp := shr(
            40,
            and(extraData, 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF0000000000)
            )
        }
        // solhint-enable max-line-length

        return (totalElements, lastSequencerTimestamp);
    }

    /**
     * Encodes the batch context for the extra data.
     * @param _totalElements Total number of elements submitted.
     * @param _lastSequencerTimestamp Timestamp of the last batch submitted by the sequencer.
     * @return Encoded batch context.
     */
    function _makeBatchExtraData(uint40 _totalElements, uint40 _lastSequencerTimestamp)
    internal
    pure
    returns (bytes27)
    {
        bytes27 extraData;
        assembly {
            extraData := _totalElements
            extraData := or(extraData, shl(40, _lastSequencerTimestamp))
            extraData := shl(40, extraData)
        }

        return extraData;
    }

    /**
     * Appends a batch to the chain.
     * @param _batch Elements within the batch.
     * @param _shouldStartAtElement Relative rollup block height.
     * @param _signature Signature of batch roots and rollup start height.
     */
    function _checkClusterSignature(bytes32[] memory _batch, uint256 _shouldStartAtElement, bytes memory _signature)
    internal
    {
        // abi hash encode to bytes
        require(
            ITssGroupManager(resolve("Proxy__TSS_GroupManager")).verifySign(
                keccak256(abi.encode(_batch, _shouldStartAtElement)), _signature),
            "verify signature failed"
        );
    }

    /**
     * Appends a batch to the chain.
     * @param _batch Elements within the batch.
     * @param _extraData Any extra data to append to the batch.
     */
    function _appendBatch(bytes32[] memory _batch, bytes memory _signature, bytes memory _extraData) internal {
        address sequencer = resolve("BVM_Proposer");
        (uint40 totalElements, uint40 lastSequencerTimestamp) = _getBatchExtraData();

        if (msg.sender == sequencer) {
            lastSequencerTimestamp = uint40(block.timestamp);
        } else {
            // We keep track of the last batch submitted by the sequencer so there's a window in
            // which only the sequencer can publish state roots. A window like this just reduces
            // the chance of "system breaking" state roots being published while we're still in
            // testing mode. This window should be removed or significantly reduced in the future.
            require(
                lastSequencerTimestamp + SEQUENCER_PUBLISH_WINDOW < block.timestamp,
                "Cannot publish state roots within the sequencer publication window."
            );
        }

        // For efficiency reasons getMerkleRoot modifies the `_batch` argument in place
        // while calculating the root hash therefore any arguments passed to it must not
        // be used again afterwards
        Lib_BVMCodec.ChainBatchHeader memory batchHeader = Lib_BVMCodec.ChainBatchHeader({
        batchIndex : getTotalBatches(),
        batchRoot : Lib_MerkleTree.getMerkleRoot(_batch),
        batchSize : _batch.length,
        prevTotalElements : totalElements,
        signature : _signature,
        extraData : _extraData
        });

        emit StateBatchAppended(
            batchHeader.batchIndex,
            batchHeader.batchRoot,
            batchHeader.batchSize,
            batchHeader.prevTotalElements,
            batchHeader.signature,
            batchHeader.extraData
        );

        batches().push(
            Lib_BVMCodec.hashBatchHeader(batchHeader),
            _makeBatchExtraData(
                uint40(batchHeader.prevTotalElements + batchHeader.batchSize),
                lastSequencerTimestamp
            )
        );
    }

    /**
     * Removes a batch and all subsequent batches from the chain.
     * @param _batchHeader Header of the batch to remove.
     */
    function _deleteBatch(Lib_BVMCodec.ChainBatchHeader memory _batchHeader) internal {
        require(_batchHeader.batchIndex < batches().length(), "Invalid batch index.");

        require(_isValidBatchHeader(_batchHeader), "Invalid batch header.");

        // slither-disable-next-line reentrancy-events
        batches().deleteElementsAfterInclusive(
            _batchHeader.batchIndex,
            _makeBatchExtraData(uint40(_batchHeader.prevTotalElements), 0)
        );

        // slither-disable-next-line reentrancy-events
        emit StateBatchDeleted(_batchHeader.batchIndex, _batchHeader.batchRoot);
    }

    /**
     * Distribute Reward to tss node.
     * @param _batch rollup batch.
     * @param  _shouldStartAtElement.
     */
    function _distributeTssReward(bytes32[] memory _batch, uint256 _shouldStartAtElement) internal {
        // get address of tss group member
        address[] memory tssMembers = ITssGroupManager(resolve("Proxy__TSS_GroupManager")).getTssGroupUnJailMembers();
        require(tssMembers.length > 0, "get tss members in error");

        // construct calldata for claimReward call
        bytes memory message = abi.encodeWithSelector(
            ITssRewardContract.claimReward.selector,
            _shouldStartAtElement,
            _batch.length,
            block.timestamp,
            tssMembers
        );

        // send call data into L2, hardcode address
        sendCrossDomainMessage(
            address(0x4200000000000000000000000000000000000020),
            2000000,
            message
        );

        // emit message
        emit DistributeTssReward(
            _shouldStartAtElement,
            _batch.length,
            block.timestamp,
            tssMembers
        );
    }

    /**
     * Checks that a batch header matches the stored hash for the given index.
     * @param _batchHeader Batch header to validate.
     * @return Whether or not the header matches the stored one.
     */
    function _isValidBatchHeader(Lib_BVMCodec.ChainBatchHeader memory _batchHeader)
    internal
    view
    returns (bool)
    {
        return Lib_BVMCodec.hashBatchHeader(_batchHeader) == batches().get(_batchHeader.batchIndex);
    }


    /**
     * roll back l2 chain to should start block number.
     * @param _shouldRollBack roll back number.
     * @param _signature Signature of roll back number.
     */
    function _checkRollBackSignature(uint256 _shouldRollBack, bytes memory _signature)
    internal
    {
        // abi hash encode to bytes
        require(
            ITssGroupManager(resolve("Proxy__TSS_GroupManager")).verifySign(
                keccak256(abi.encode(_shouldRollBack)), _signature),
            "verify signature failed"
        );


        // construct calldata for claimReward call
        bytes memory message = abi.encodeWithSelector(
            IStateCommitmentChain.rollBackMessage.selector,
                _shouldRollBack
        );

        // send call data into L2, hardcode address
        sendCrossDomainMessage(
            address(0xDeADdeaDdEaDdeADdEaDDeADDEaddEaDDEad2222),
            2000000,
            message
        );

        //emit roll back l2chain from block number
        emit RollBackL2Chain(_shouldRollBack);
    }
}

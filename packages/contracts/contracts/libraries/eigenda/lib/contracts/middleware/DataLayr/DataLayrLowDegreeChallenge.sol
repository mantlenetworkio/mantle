// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../../interfaces/IRepository.sol";
import "../../interfaces/IRepositoryAccess.sol";
import "../../interfaces/IQuorumRegistry.sol";
import "../../interfaces/IDataLayrServiceManager.sol";

import "../Repository.sol";

import "./DataLayrChallengeUtils.sol";

import "../../libraries/Merkle.sol";
import "../../libraries/BLS.sol";
import "../../libraries/BytesLib.sol";
import "../../libraries/DataStoreUtils.sol";

/**
 * @title Used to create and manage low degree challenges related to DataLayr.
 * @author Layr Labs, Inc.
 */
contract DataLayrLowDegreeChallenge {
    using SafeERC20 for IERC20;
    using BytesLib for bytes;

    IQuorumRegistry public immutable dlRegistry;
    DataLayrChallengeUtils public immutable challengeUtils;
    IDataLayrServiceManager public immutable dataLayrServiceManager;

    //Fixed gas limit to ensure pairing precompile doesn't use entire gas limit upon reversion
    uint256 public pairingGasLimit;

    enum ChallengeStatus {
        UNSUCCESSFUL,
        SUCCESSFUL
    }

    struct LowDegreeChallenge {
        // challenger's address
        address challenger;
        // challenge status
        ChallengeStatus status;
    }

    struct NonSignerExclusionProof {
        address signerAddress;
        uint32 operatorHistoryIndex;
    }

    event SuccessfulLowDegreeChallenge(bytes32 indexed headerHash, address challenger);

    mapping(bytes32 => LowDegreeChallenge) public lowDegreeChallenges;

    //POT refers to Powers of Tau
    uint256 internal constant MAX_POT_DEGREE = (2 ** 28);
    uint256 internal constant POT_TREE_HEIGHT = 28;

    modifier onlyRepositoryGovernance() {
        require(msg.sender == IRepositoryAccess(address(dataLayrServiceManager)).repository().owner(), "onlyRepositoryGovernance");
        _;
    }

    constructor(
        IDataLayrServiceManager _dataLayrServiceManager,
        IQuorumRegistry _dlRegistry,
        DataLayrChallengeUtils _challengeUtils,
        uint256 _gasLimit
    ) {
        dataLayrServiceManager = _dataLayrServiceManager;
        dlRegistry = _dlRegistry;
        challengeUtils = _challengeUtils;
        pairingGasLimit = _gasLimit;
    }

    /**
     *   @notice verifies all challenger inputs against stored hashes, computes low degreeness proof and
     *   freezes operator if verified as being excluded from nonsigner set.
     *   @param header is the header for the datastore in question.
     *   @param potElement is the G2 point of the POT element we are computing the pairing for (x^{n-m})
     *   @param potMerkleProof is the merkle proof for the POT element.
     *   @param dataStoreSearchData is the all relevant data about the datastore being challenged
     *   @param signatoryRecord is the record of signatures on said datastore
     */
    function challengeLowDegreenessProof(
        bytes calldata header,
        BN254.G2Point memory potElement,
        bytes memory potMerkleProof,
        IDataLayrServiceManager.DataStoreSearchData calldata dataStoreSearchData,
        IDataLayrServiceManager.SignatoryRecordMinusDataStoreId calldata signatoryRecord
    ) external {
        require(
            dataStoreSearchData.metadata.headerHash == keccak256(header),
            "provided datastore searchData does not match provided header"
        );

        /// @dev Refer to the datastore header spec for makeup of header
        BN254.G1Point memory lowDegreenessProof;

        //Slice the header to retrieve the lowdegreeness proof (a G1 point)
        assembly {
            mstore(lowDegreenessProof, calldataload(add(header.offset, 116)))
            mstore(add(lowDegreenessProof, 32), calldataload(add(header.offset, 148)))
        }

        //prove searchData, including nonSignerPubkeyHashes (in the form of signatory record in the metadata) matches stored searchData
        require(
            dataLayrServiceManager.verifyDataStoreMetadata(
                dataStoreSearchData.duration,
                dataStoreSearchData.timestamp,
                dataStoreSearchData.index,
                dataStoreSearchData.metadata
            ),
            "DataLayrLowDegreeChallenge.challengeLowDegreeHeader: Provided metadata does not match stored datastore metadata hash"
        );

        bytes32 signatoryRecordHash = DataStoreUtils.computeSignatoryRecordHash(
            dataStoreSearchData.metadata.globalDataStoreId,
            signatoryRecord.nonSignerPubkeyHashes,
            signatoryRecord.signedStakeFirstQuorum,
            signatoryRecord.signedStakeSecondQuorum
        );

        require(
            dataStoreSearchData.metadata.signatoryRecordHash == signatoryRecordHash,
            "DataLayrLowDegreeChallenge.lowDegreeChallenge: provided signatoryRecordHash does not match signatorRecordHash in provided searchData"
        );

        require(
            !verifyLowDegreenessProof(header, potElement, potMerkleProof, lowDegreenessProof),
            "DataLayrLowDegreeChallenge.lowDegreeChallenge: low degreeness proof verified successfully"
        );

        bytes32 dataStoreHash = keccak256(abi.encode(dataStoreSearchData));
        lowDegreeChallenges[dataStoreHash] = LowDegreeChallenge(msg.sender, ChallengeStatus.SUCCESSFUL);

        emit SuccessfulLowDegreeChallenge(dataStoreHash, msg.sender);
    }

    ///@notice slash an operator who signed a headerHash but failed a subsequent challenge
    function freezeOperatorsForLowDegreeChallenge(
        NonSignerExclusionProof[] memory nonSignerExclusionProofs,
        uint256 nonSignerIndex,
        IDataLayrServiceManager.DataStoreSearchData calldata searchData,
        IDataLayrServiceManager.SignatoryRecordMinusDataStoreId calldata signatoryRecord
    ) external {
        for (uint256 i; i < nonSignerExclusionProofs.length; i++) {
            address operator = nonSignerExclusionProofs[i].signerAddress;
            uint32 operatorHistoryIndex = nonSignerExclusionProofs[i].operatorHistoryIndex;

            // verify that operator was active *at the blockNumber*
            bytes32 operatorPubkeyHash = dlRegistry.getOperatorPubkeyHash(operator);
            IQuorumRegistry.OperatorStake memory operatorStake =
                dlRegistry.getStakeFromPubkeyHashAndIndex(operatorPubkeyHash, operatorHistoryIndex);
            require(
                // operator must have become active/registered before (or at) the block number
                (operatorStake.updateBlockNumber <= searchData.metadata.blockNumber)
                // operator must have still been active after (or until) the block number
                // either there is a later update, past the specified blockNumber, or they are still active
                && (
                    operatorStake.nextUpdateBlockNumber >= searchData.metadata.blockNumber
                        || operatorStake.nextUpdateBlockNumber == 0
                ),
                "DataLayrChallengeBase.slashOperator: operator was not active during blockNumber specified by dataStoreId / headerHash"
            );

            if (signatoryRecord.nonSignerPubkeyHashes.length != 0) {
                // check that operator was *not* in the non-signer set (i.e. they *did* sign)
                challengeUtils.checkExclusionFromNonSignerSet(operatorPubkeyHash, nonSignerIndex, signatoryRecord);
            }

            dataLayrServiceManager.freezeOperator(operator);
        }
    }

    /**
     * @notice This function verifies that a polynomial's degree is not greater than a provided degree and returns true if 
               the inputs to the pairing are valid and the pairing is successful.
     * @param header is the header information, which contains the kzg metadata (commitment and degree to check against)
     * @param potElement is the G2 point of the POT element we are computing the pairing for (x^{n-m})
     * @param potMerkleProof is the merkle proof for the POT element.
     * @param lowDegreenessProof is the provided G1 point which is the product of the POTElement and the polynomial, i.e., [(x^{n-m})*p(x)]_1
     *        This function computes the pairing e([p(x)]_1, [x^{n-m}]_2) = e([(x^{n-m})*p(x)]_1, [1]_2)
     */

    function verifyLowDegreenessProof(
        bytes calldata header,
        BN254.G2Point memory potElement,
        bytes memory potMerkleProof,
        BN254.G1Point memory lowDegreenessProof
    ) public view returns (bool) {
        require(potMerkleProof.length/32 ==  POT_TREE_HEIGHT, "DataLayrLowDegreeChallenge.verifyLowDegreenessProof: incorrect proof length");

        //retreiving the kzg commitment to the data in the form of a polynomial
        DataLayrChallengeUtils.DataStoreKZGMetadata memory dskzgMetadata =
            challengeUtils.getDataCommitmentAndMultirevealDegreeAndSymbolBreakdownFromHeader(header);

        //the index of the merkle tree containing the potElement
        uint256 potIndex = MAX_POT_DEGREE - dskzgMetadata.degree * challengeUtils.nextPowerOf2(dskzgMetadata.numSys);
        //computing hash of the powers of Tau element to verify merkle inclusion
        bytes32 hashOfPOTElement = keccak256(abi.encodePacked(potElement.X, potElement.Y));

        require(
            Merkle.checkMembership(hashOfPOTElement, potIndex, BLS.powersOfTauMerkleRoot, potMerkleProof),
            "DataLayrLowDegreeChallenge.proveLowDegreeness: PoT merkle proof failed"
        );

        BN254.G2Point memory negativeG2 = BN254.G2Point({X: [BLS.nG2x1, BLS.nG2x0], Y: [BLS.nG2y1, BLS.nG2y0]});

        (bool precompileWorks, bool pairingSuccessful) =
            BN254.safePairing(dskzgMetadata.c, potElement, lowDegreenessProof, negativeG2, pairingGasLimit);

        return (precompileWorks && pairingSuccessful);
    }

    //update pairing gas limit
    function setPairingGasLimit(uint256 newGasLimit) external onlyRepositoryGovernance {
        pairingGasLimit = newGasLimit;
    }
}

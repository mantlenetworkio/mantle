// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../../interfaces/IRepository.sol";
import "../../interfaces/IQuorumRegistry.sol";
import "../../interfaces/IDataLayrServiceManager.sol";
import "../../interfaces/IDataLayrPaymentManager.sol";
import "../Repository.sol";
import "../../libraries/DataStoreUtils.sol";
import "../../middleware/PaymentManager.sol";

/**
 * @title This contract is used for doing interactive payment challenges on DataLayr.
 * @author Layr Labs, Inc.
 * @notice Most functionality is inherited from the `PaymentManager` contract, with `respondToPaymentChallengeFinal` implementing DataLayr-specific functionality.
 */
contract DataLayrPaymentManager is PaymentManager, IDataLayrPaymentManager, Initializable {
    IDataLayrServiceManager public immutable dataLayrServiceManager;

    constructor(
        IERC20 _paymentToken,
        uint256 _paymentFraudproofCollateral,
        IRepository _repository,
        IDataLayrServiceManager _dataLayrServiceManager,
        IPauserRegistry _pauserReg
    )
        PaymentManager(_paymentToken, _paymentFraudproofCollateral, _repository, _pauserReg)
        initializer
    {
        dataLayrServiceManager = _dataLayrServiceManager;
    }

    /**
     * @notice Used to perform the final step in a payment challenge, in which the 'trueAmount' is determined and the winner of the challenge is decided.
     * This function is called by a party after the other party has bisected the challenged payments to a difference of one, i.e., further bisection
     * is not possible. Once the payments can no longer be bisected, the function resolves the challenge by determining who is wrong.
     */
    function respondToPaymentChallengeFinal(
        address operator,
        uint256 stakeIndex,
        uint48 nonSignerIndex,
        bytes32[] memory nonSignerPubkeyHashes,
        TotalStakes calldata totalStakesSigned,
        IDataLayrServiceManager.DataStoreSearchData calldata searchData
    )
        external
    {
        // copy challenge struct to memory
        PaymentChallenge memory challenge = operatorToPaymentChallenge[operator];

        require(
            block.timestamp < challenge.settleAt,
            "DataLayrPaymentManager.respondToPaymentChallengeFinal: challenge has already passed settlement time"
        );

        //checks that searchData is valid by checking against the hash stored in DLSM's dataStoreHashesForDurationAtTimestamp
        require(
            dataLayrServiceManager.verifyDataStoreMetadata(
                    searchData.duration,
                    searchData.timestamp,
                    searchData.index,
                    searchData.metadata
            ), "DataLayrPaymentManager.respondToPaymentChallengeFinal: search.metadata preimage is incorrect"
        );

        // recalculate the signatoryRecordHash, to verify integrity of `nonSignerPubkey` and `totalStakesSigned` inputs.
        bytes32 providedSignatoryRecordHash = DataStoreUtils.computeSignatoryRecordHash(
                searchData.metadata.globalDataStoreId,
                nonSignerPubkeyHashes,
                totalStakesSigned.signedStakeFirstQuorum,
                totalStakesSigned.signedStakeSecondQuorum
        );
        
        //checking that `nonSignerPubKeyHashes` and `totalStakesSigned` are correct, now that we know that searchData is valid
        require(
            providedSignatoryRecordHash == searchData.metadata.signatoryRecordHash,
            "DataLayrPaymentManager.respondToPaymentChallengeFinal: provided nonSignerPubKeyHashes or totalStakesSigned is incorrect"
        );

        // get the registry address
        IQuorumRegistry registry = IQuorumRegistry(address(repository.registry()));

        // pull the operator's pubkey hash
        bytes32 operatorPubkeyHash = registry.getOperatorPubkeyHash(operator);

        // calculate the true amount deserved
        uint96 trueAmount;

        //2^32 is an impossible index because it is more than the max number of registrants
        //the challenger marks 2^32 as the index to show that operator has not signed
        if (nonSignerIndex == 1 << 32) {
            for (uint256 i = 0; i < nonSignerPubkeyHashes.length;) {
                require(
                    nonSignerPubkeyHashes[i] != operatorPubkeyHash,
                    "DataLayrPaymentManager.respondToPaymentChallengeFinal: Operator was not a signatory"
                );

                unchecked {
                    ++i;
                }
            }
            IQuorumRegistry.OperatorStake memory operatorStake =
                registry.getStakeFromPubkeyHashAndIndex(operatorPubkeyHash, stakeIndex);

            // scoped block helps fix stack too deep
            {
                require(
                    operatorStake.updateBlockNumber <= searchData.metadata.blockNumber,
                    "DataLayrPaymentManager.respondToPaymentChallengeFinal: Operator stake index is too late"
                );

                require(
                    operatorStake.nextUpdateBlockNumber == 0
                        || operatorStake.nextUpdateBlockNumber > searchData.metadata.blockNumber,
                    "DataLayrPaymentManager.respondToPaymentChallengeFinal: Operator stake index is too early"
                );
            }

            require(
                searchData.metadata.globalDataStoreId == challenge.fromTaskNumber,
                "DataLayrPaymentManager.respondToPaymentChallengeFinal: Loaded DataStoreId does not match challenged"
            );

            // look up the voteWeigher address
            IVoteWeigher voteWeigher = repository.voteWeigher();

            /*
             *  searchData.metadata.fee is the total fee for a datastore
             *  This needs to be split up among operators, so we multiply 
             *  the fee by operatorStake/totalStakesSigned to get the operator's
             *  share of the stake.  Then we multiply by quorumBips/MAX_BIPS to get
             *  the percentage of the operator's fee for that quorum
            */
            trueAmount = uint96(
                (
                    (
                        (
                            uint256(searchData.metadata.fee) * uint256(voteWeigher.quorumBips(0))
                                * uint256(operatorStake.firstQuorumStake)
                        ) / totalStakesSigned.signedStakeFirstQuorum
                    )
                        + (
                            (
                                uint256(searchData.metadata.fee) * uint256(voteWeigher.quorumBips(1))
                                    * uint256(operatorStake.secondQuorumStake)
                            ) / totalStakesSigned.signedStakeSecondQuorum
                        )
                ) / MAX_BIPS
            );
        } else {
            //either the operator must have been a non signer or the task was based off of stakes before the operator registered
            require(
                nonSignerPubkeyHashes[nonSignerIndex] == operatorPubkeyHash
                    || searchData.metadata.blockNumber < registry.getFromBlockNumberForOperator(operator),
                "DataLayrPaymentManager.respondToPaymentChallengeFinal: Signer index is incorrect"
            );
        }

        {
            //final entity is the entity calling this function, i.e., it is their turn to make the final response
            bool finalEntityCorrect = trueAmount != challenge.amount1;
            /*
        * if status is OPERATOR_TURN_ONE_STEP, it is the operator's turn. This means the challenger was the one who set challenge.amount1 last.  
        * If trueAmount != challenge.amount1, then the challenger is wrong (doesn't mean operator is right).
        */
            if (challenge.status == ChallengeStatus.OPERATOR_TURN_ONE_STEP) {
                _resolve(challenge, finalEntityCorrect ? challenge.operator : challenge.challenger);
            }
            /*
        * if status is CHALLENGER_TURN_ONE_STEP, it is the challenger's turn. This means the operator was the one who set challenge.amount1 last.  
        * If trueAmount != challenge.amount1, then the operator is wrong and the challenger is correct
        */
            else if (challenge.status == ChallengeStatus.CHALLENGER_TURN_ONE_STEP) {
                _resolve(challenge, finalEntityCorrect ? challenge.challenger : challenge.operator);
            } else {
                revert("DataLayrPaymentManager.respondToPaymentChallengeFinal: Not in one step challenge phase");
            }

            challenge.status = ChallengeStatus.RESOLVED;
        }

        // update challenge struct in storage
        operatorToPaymentChallenge[operator] = challenge;
    }
}

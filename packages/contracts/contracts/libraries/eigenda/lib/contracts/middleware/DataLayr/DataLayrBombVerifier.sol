// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9.0;

import "../../interfaces/IDataLayrServiceManager.sol";
import "../../interfaces/IQuorumRegistry.sol";
import "../../interfaces/IEphemeralKeyRegistry.sol";
import "../../libraries/DataStoreUtils.sol";
import "./DataLayrChallengeUtils.sol";
import "../../libraries/DataStoreUtils.sol";
import "../../libraries/BN254.sol";

/**
 * @title Used to check Proofs of Custody in DataLayr.
 * @author Layr Labs, Inc.
 * @notice The core slashing module of DataLayr. Using Proofs of Custody, DataLayr is able to slash operators who are provably not storing their data.
 * @dev In order to prove that an operator wasn’t storing data at certain time, a challenger proves the following:
 * 1) The existence of a certain datastore referred to as the DETONATION datastore
 * 2) The existence of a certain datastore referred to as the BOMB datastore, which the operator has certified to storing,
 * that is chosen on-chain via the result of a function of the DETONATION datastore's header hash
 * 3) The data that the operator was storing for the BOMB datastore, when hashed with the operator's ephemeral key and the DETONATION datastore's
 * header hash, is below a certain threshold defined by the `DataLayrBombVerifier` contract
 * 4) The operator certified the storing of DETONATION datastore
 * If these 4 points are proved, the operator is slashed.
 * The operator should be checking the following above requirements against each new header hash it receives in order to not be slashed.
 */
contract DataLayrBombVerifier {
    /// @notice This struct is exactly IDataLayrServiceManager.DataStoreSearchData without the duration, used for identifying the correct bomb datastore.
    struct DataStoresForDuration {
        uint256 timestamp;
        uint32 index;
        IDataLayrServiceManager.DataStoreMetadata metadata;
    }

    /**
     * @notice This struct includes proofs of the metadata of the DETONATION datastore and as many potential BOMB datastores as it takes to find one
     * that the operator signed and the datastore from which the operator to be slashed started serving DataLayr.
     */
    struct DataStoreProofs {
        IDataLayrServiceManager.DataStoreSearchData operatorFromDataStore;
        IDataLayrServiceManager.DataStoreSearchData[] bombDataStores;
        IDataLayrServiceManager.DataStoreSearchData detonationDataStore;
    }

    /**
     * @notice `operatorIndex` is the index in the operator's list of indexes in the `dlRegistry`
     * that provides the operator's index at the time of the BOMB datastore
     * `totalOperatorsIndex` is the index in the list of total operators over time in the `dlRegistry`
     * that provides the total number of operators at the timeof the BOMB datastore
     * `detonationNonSignerIndex` is the index within the non-signer list of the DETONATION datastore that proves
     * that the operator signed on the availability of the DETONATION datastore
     */
    struct Indexes {
        uint32 operatorIndex;
        uint32 totalOperatorsIndex;
        uint256 detonationNonSignerIndex;
        uint256[] successiveSignerIndexes;
    }

    /// @notice Proof of the data itself -- from the BOMB datastore -- that the operator was responsible for storing
    struct DisclosureProof {
        bytes header;
        bytes poly;
        DataLayrChallengeUtils.MultiRevealProof multiRevealProof;
        BN254.G2Point polyEquivalenceProof;
    }

    /// @notice determines how often bombs will 'appear'
    // bomb will trigger every once every ~2^(256-249) = 2^7 = 128 chances
    // BOMB_THRESHOLD can be tuned up to increase the chance of bombs and therefore
    // reduce the expected value of not storing the data
    // BOMB_THRESHOLD can be tuned down to decrease the chance of bombs and therefore
    // increase the amount of nodes that will sign off on datastores
    uint256 public constant BOMB_THRESHOLD = uint256(2) ** uint256(249);

    // This is the interval in which the bomb must be proven
    uint256 public constant BOMB_FRAUDRPOOF_INTERVAL = 7 days;

    IDataLayrServiceManager public immutable dlsm;
    IQuorumRegistry public immutable dlRegistry;
    DataLayrChallengeUtils public immutable challengeUtils;
    IEphemeralKeyRegistry public immutable dlekRegistry;

    /**
     * @notice Initializes with the `DataLayrServiceManager`, the BLS registry used by DataLayr, the `DataLayrChallengeUtils` library to help
     * abstract away some common logic, and, the `IEphemeralKeyRegistry` used by DataLayr.
     */
    constructor(
        IDataLayrServiceManager _dlsm,
        IQuorumRegistry _dlRegistry,
        DataLayrChallengeUtils _challengeUtils,
        IEphemeralKeyRegistry _dlekRegistry
    ) {
        dlsm = _dlsm;
        dlRegistry = _dlRegistry;
        challengeUtils = _challengeUtils;
        dlekRegistry = _dlekRegistry;
    }

    /**
     * @notice Used to prove that `operator` improperly signed a datastore meeting the BOMB condition, making the operator subject to slashing.
     * The header hash of the DETONATION datastore id is mapped to one of the active datastores at the time, the BOMB datastore. If the operator
     * signed off on the availabilty of the BOMB datastore, then the function proceeds. If not, datastores are iterated through consecutively by
     * id until a datastore that the operator has signed is found. The first datastore that the operator hash signed is considered the BOMB datastore.
     * Ultimately, the BOMB condition is checked against a combined hash of the data stored in the BOMB datastore, the operator's ephemeral key, and
     * the header hash of the DETONATION datastore
     *
     * @dev Exceptionally verbose explanation:
     * The DETONATION datastore is the datastore whose header hash is mapped to one of the active datastores at its time of initialization.
     * The datastore that the DETONATION datastore is mapped to is called the BOMB datastore.
     * The BOMB datastore is the datastore whose data, when hashed with some auxiliary information, returned a hash value below BOMB_THRESHOLD
     * (the BOMB condition).
     * If such was the case, the operator should not have signed the DETONATION datastore, and thus 'detonated the bomb', making them subject
     * to slashing.
     *
     * In datalayr, every datastore is a potential DETONATION datastore, and it's corresponding potential BOMB datastore should
     * always be checked for the BOMB condition
     * The sender of this function is a party that is proving the existence of a certain operator that signed a DETONATION datastore whose corresponding
     * BOMB datastore met the BOMB condition
     *
     * tick, tick, tick, tick, ⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️
     *
     * @param operator is the address of the operator to slash
     * @param dataStoreProofs are the proofs of the datastores needed to calculate the bomb function
     * @param signatoryRecords are the signatory records needed to prove that the operator signed the DETONATION datastore and to prove the correct
     * BOMB datastore that the operator signed and all of the signatory records of potential BOMB datastores that the operator did not sign
     * @param sandwichProofs are the proofs of which datastores were active for each duration at the time the DETONATION datastore was initialized
     * and will be explained in more detail further in the document
     * @param disclosureProof is the proof of the data the operator was storing of the BOMB datastore
     *
     * @dev signatoryRecords input is formatted as following, with 'n' being its length:
     * signatoryRecords[0] is for the 'detonation' DataStore
     * signatoryRecords[1] through (inclusive) signatoryRecords[n-2] is for the DataStores starting at the 'bomb' DataStore returned by the
     * 'verifyBombDataStoreId' function and any immediately following series DataStores *that the operator did NOT sign*
     * signatoryRecords[n] is for the DataStore that is ultimately treated as the 'bomb' DataStore
     * this will be the first DataStore at or after the DataStore returned by the 'verifyBombDataStoreId' function *that the operator DID sign*
     */
    function verifyBomb(
        address operator,
        DataStoreProofs calldata dataStoreProofs,
        Indexes calldata indexes,
        IDataLayrServiceManager.SignatoryRecordMinusDataStoreId[] calldata signatoryRecords,
        DataStoresForDuration[2][2][] calldata sandwichProofs,
        DisclosureProof calldata disclosureProof
    ) external {
        // verify integrity of submitted metadata by checking against its stored hashes
        require(
            verifyMetadataPreImage(dataStoreProofs.operatorFromDataStore),
            "DataLayrBombVerifier.verifyBomb: operatorFrom metadata preimage incorrect"
        );
        require(
            verifyMetadataPreImage(dataStoreProofs.detonationDataStore),
            "DataLayrBombVerifier.verifyBomb: detonation metadata preimage incorrect"
        );

        {
            /**
             * require that either operator is still actively registered, OR
             * they were previously active and they deregistered within the last 'BOMB_FRAUDRPOOF_INTERVAL'
             */
            //get the id of the datastore the operator has been serving since
            uint32 fromDataStoreId = dlRegistry.getFromTaskNumberForOperator(operator);
            //deregisterTime is 0 if the operator is still registered and serving
            //otherwise it is the time at will/have stopped serving all of their existing datstores
            uint256 deregisterTime = dlRegistry.getOperatorDeregisterTime(operator);
            //Require that the operator is actively registered or, if they have deregistered, it is still before the 'BOMB_FRAUDRPOOF_INTERVAL' has passed
            require(
                fromDataStoreId != 0
                    && (deregisterTime == 0 || block.timestamp <= deregisterTime + BOMB_FRAUDRPOOF_INTERVAL),
                "DataLayrBombVerifier.verifyBomb: invalid operator or time"
            );
        }

        // get globalDataStoreId at bomb DataStore
        uint32 bombGlobalDataStoreId = verifyBombDataStoreId(operator, dataStoreProofs, sandwichProofs);

        /*
            this large block with for loop is used to iterate through DataStores
            although technically the pseudo-random DataStore containing the bomb is already determined, it is possible
            that the operator did not sign the 'bomb' DataStore (note that this is different than signing the 'detonator' DataStore!).
            In this specific case, the 'bomb' is actually contained in the next DataStore that the operator did indeed sign.
            The loop iterates through to find this next DataStore, thus determining the true 'bomb' DataStore.
        */
        /**
         * @notice Check that the DataLayr operator against whom bomb is being verified, was
         * actually part of the quorum for the detonation dataStoreId.
         *
         * The burden of responsibility lies with the challenger to show that the DataLayr operator
         * is not part of the non-signers for the dump. Towards that end, challenger provides
         * an index such that if the relationship among nonSignerPubkeyHashes (nspkh) is:
         * uint256(nspkh[0]) <uint256(nspkh[1]) < ...< uint256(nspkh[index])< uint256(nspkh[index+1]),...
         * then,
         * uint256(nspkh[index]) <  uint256(operatorPubkeyHash) < uint256(nspkh[index+1])
         */
        /**
         * @dev checkSignatures in DataLayrBLSSignatureChecker.sol enforces the invariant that the hash of all
         * non-signers' pubkeys are recorded in the compressed signatory record in a strictly ascending order.
         */
        // first we verify that the operator did indeed sign the 'detonation' DataStore
        {
            //the block number since the operator has been active
            uint32 operatorActiveFromBlockNumber = dlRegistry.getFromBlockNumberForOperator(operator);
            // fetch hash of operator's pubkey
            bytes32 operatorPubkeyHash = dlRegistry.getOperatorPubkeyHash(operator);

            // The BOMB datastore must be a datastore for which a signature from the operator has been submitted on chain
            // Then, we have an attestation that they have stored said data, so they can check it for the BOMB condition
            uint256 ultimateBombDataStoreIndex = dataStoreProofs.bombDataStores.length - 1;
            //verify all non signed DataStores from bomb till reaching the first signed DataStore to get the correct BOMB datastore
            for (uint256 i = 0; i < ultimateBombDataStoreIndex; ++i) {
                require(
                    dataStoreProofs.bombDataStores[i].metadata.globalDataStoreId == bombGlobalDataStoreId,
                    "DataLayrBombVerifier.verifyBomb: bombDataStore is not for correct id"
                );
                //verify the preimage of the i'th bombDataStore is consistent with storage
                require(
                    verifyMetadataPreImage(dataStoreProofs.bombDataStores[i]),
                    "DataLayrBombVerifier.verifyBomb: bombDataStores metadata preimage incorrect"
                );

                //There are 2 conditions under which the BOMB datastore id must increment
                //1. The BOMB datastore is based off of stakes before the operator joined
                //2. The BOMB datastore included the stake of the operator, but the operator did not sign it
                //This conditional statement checks (1)
                if (dataStoreProofs.bombDataStores[i].metadata.blockNumber >= operatorActiveFromBlockNumber) {
                    //If we make it inside of this loop, then the BOMB datastore included the operator's stake
                    //So we check the proof that the operator did not sign for this datastore
                    // Verify that the signatoryRecord supplied as input related to the i'th potential BOMB datastore is correct
                    require(
                        //will be bytes32(0) if this datastore was never confirmed
                        dataStoreProofs.bombDataStores[i].metadata.signatoryRecordHash == bytes32(0)
                            || dataStoreProofs.bombDataStores[i].metadata.signatoryRecordHash
                                == DataStoreUtils.computeSignatoryRecordHash(
                                    bombGlobalDataStoreId,
                                    signatoryRecords[i].nonSignerPubkeyHashes,
                                    signatoryRecords[i].signedStakeFirstQuorum,
                                    signatoryRecords[i].signedStakeSecondQuorum
                                ),
                        "DataLayrBombVerifier.verifyBomb: Bomb datastore signatory record does not match hash"
                    );

                    // verify that the operator was in the non-signer set (i.e did *NOT* sign) for this datastore
                    require(
                        signatoryRecords[i].nonSignerPubkeyHashes[indexes.successiveSignerIndexes[i]]
                            == operatorPubkeyHash,
                        "DataLayrBombVerifier.verifyBomb: Incorrect Bomb datastore nonsigner proof"
                    );
                }
                ++bombGlobalDataStoreId;
                //if we have incremented through all datastores after the initial BOMB datastore id, up until detonation, then there is
                //no proof of custody needed, as they have not signed/been paid for a while!
                if (bombGlobalDataStoreId == dataStoreProofs.detonationDataStore.metadata.globalDataStoreId) {
                    return;
                }
            }

            // verify that the correct BOMB dataStoreId (the first the operator signed at or above the pseudo-random dataStoreId) matches the provided data
            require(
                dataStoreProofs.bombDataStores[ultimateBombDataStoreIndex].metadata.globalDataStoreId
                    == bombGlobalDataStoreId,
                "DataLayrBombVerifier.verifyBomb: provided bomb datastore id must be as calculated"
            );

            //verify the preimage of the last provided BOMB datastore (the valid one) is consistent with storage
            require(
                verifyMetadataPreImage(dataStoreProofs.bombDataStores[ultimateBombDataStoreIndex]),
                "DataLayrBombVerifier.verifyBomb: BOMB datastore metadata preimage incorrect"
            );

            //Verify that the signatory record supplied as input related to the ultimate 'bomb' DataStore is correct
            require(
                dataStoreProofs.bombDataStores[ultimateBombDataStoreIndex].metadata.signatoryRecordHash
                    == DataStoreUtils.computeSignatoryRecordHash(
                        dataStoreProofs.bombDataStores[ultimateBombDataStoreIndex].metadata.globalDataStoreId,
                        signatoryRecords[ultimateBombDataStoreIndex].nonSignerPubkeyHashes,
                        signatoryRecords[ultimateBombDataStoreIndex].signedStakeFirstQuorum,
                        signatoryRecords[ultimateBombDataStoreIndex].signedStakeSecondQuorum
                    ),
                "DataLayrBombVerifier.verifyBomb: BOMB datastore sig record does not match hash"
            );

            //require that the detonation is happening for a datastore using the operators stake
            require(
                dataStoreProofs.bombDataStores[ultimateBombDataStoreIndex].metadata.blockNumber
                    >= operatorActiveFromBlockNumber,
                "DataLayrBombVerifier.verfiyBomb: BOMB datastore was not using the operator's stake"
            );

            // check that operator was *not* in the non-signer set (i.e. they *did* sign) for the ultimate 'bomb' DataStore
            if (signatoryRecords[ultimateBombDataStoreIndex].nonSignerPubkeyHashes.length != 0) {
                // check that operator was *not* in the non-signer set (i.e. they did sign)
                //not super critic: new call here, maybe change comment
                challengeUtils.checkExclusionFromNonSignerSet(
                    operatorPubkeyHash, indexes.operatorIndex, signatoryRecords[ultimateBombDataStoreIndex]
                );
            }

            //Verify that the operator *did* sign the DETONATION datastore
            uint256 lastSignatoryRecordIndex = signatoryRecords.length - 1;

            // Verify that the signatoryRecord supplied as input related to the 'detonation' DataStore is correct
            //NOTE that signatoryRecords[signatoryRecords.length - 1] is the signatory record for the DETONATION datastore
            require(
                dataStoreProofs.detonationDataStore.metadata.signatoryRecordHash
                    == DataStoreUtils.computeSignatoryRecordHash(
                        dataStoreProofs.detonationDataStore.metadata.globalDataStoreId,
                        signatoryRecords[lastSignatoryRecordIndex].nonSignerPubkeyHashes,
                        signatoryRecords[lastSignatoryRecordIndex].signedStakeFirstQuorum,
                        signatoryRecords[lastSignatoryRecordIndex].signedStakeSecondQuorum
                    ),
                "DataLayrBombVerifier.verifyBomb: Detonation singatory record does not match hash"
            );
            // require that the detonation is happening for a datastore using the operators stake
            require(
                dataStoreProofs.detonationDataStore.metadata.blockNumber >= operatorActiveFromBlockNumber,
                "DataLayrBombVerifier.verfiyBomb: Detonation datastore was not using the operator's stake"
            );

            // check that operator was *not* in the non-signer set (i.e. they did sign) for the 'detonation' DataStore
            if (signatoryRecords[lastSignatoryRecordIndex].nonSignerPubkeyHashes.length != 0) {
                // check that operator was *not* in the non-signer set (i.e. they did sign)
                //not super critic: new call here, maybe change comment
                challengeUtils.checkExclusionFromNonSignerSet(
                    operatorPubkeyHash, indexes.detonationNonSignerIndex, signatoryRecords[lastSignatoryRecordIndex]
                );
            }
        }

        // check the disclosure of the data chunk that the operator committed to storing
        require(
            nonInteractivePolynomialProof(
                // headerHashes.bombHeaderHash,
                operator,
                indexes.operatorIndex,
                indexes.totalOperatorsIndex,
                bombGlobalDataStoreId,
                disclosureProof,
                dataStoreProofs.operatorFromDataStore
            ),
            "DataLayrBombVerifier.verifyBomb: I from multireveal is not the commitment of poly"
        );

        // fetch the operator's ephemeral key for the DETONATION datastore
        bytes32 ek = dlekRegistry.getEphemeralKeyForTaskNumber(
            operator, dataStoreProofs.detonationDataStore.metadata.globalDataStoreId
        );

        // The bomb "condition" is that keccak(data, ek, headerHash) < BOMB_THRESHOLD
        // If it is met, there was a ........  ⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️⏲️
        // 💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣
        // 💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣💣
        // 💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥
        // 💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥💥
        require(
            uint256(
                keccak256(
                    abi.encodePacked(disclosureProof.poly, ek, dataStoreProofs.detonationDataStore.metadata.headerHash)
                )
            ) < BOMB_THRESHOLD,
            "DataLayrBombVerifier.verifyBomb: No bomb"
        );

        // trigger slashing
        dlsm.freezeOperator(operator);
    }

    /// @notice Returns globalDataStoreId at bomb DataStore
    function verifyBombDataStoreId(
        address operator,
        DataStoreProofs calldata dataStoreProofs,
        DataStoresForDuration[2][2][] calldata sandwichProofs
    ) internal view returns (uint32) {
        uint256 fromTime;
        {
            // get the dataStoreId at which the operator registered
            uint32 fromDataStoreId = dlRegistry.getFromTaskNumberForOperator(operator);

            // ensure that operatorFromHeaderHash corresponds to the correct dataStoreId (i.e. the one at which the operator registered)
            require(
                fromDataStoreId == dataStoreProofs.operatorFromDataStore.metadata.globalDataStoreId,
                "DataLayrBombVerifier.verifyBombDataStoreId: headerHash is not for correct operator from datastore"
            );
            // store the initTime of the dataStoreId at which the operator registered in memory
            fromTime = dataStoreProofs.operatorFromDataStore.timestamp;
        }

        // find the specific DataStore containing the bomb, specified by durationIndex and calculatedDataStoreId
        // 'verifySandwiches' gets a pseudo-randomized durationIndex and durationDataStoreId, as well as the nextGlobalDataStoreIdAfterBomb
        (uint8 durationIndex, uint32 calculatedDataStoreId, uint32 nextGlobalDataStoreIdAfterDetonationTimestamp) =
        verifySandwiches(
            uint256(dataStoreProofs.detonationDataStore.metadata.headerHash),
            fromTime,
            dataStoreProofs.detonationDataStore.timestamp,
            sandwichProofs
        );

        require(
            sandwichProofs.length == dlsm.MAX_DATASTORE_DURATION() + 1,
            "DataLayrBombVerifier.verifyBombDataStoreId: Incorrect sandwich proof length. *must account for last proof of bomb datastoremetdata"
        );

        // fetch the durationDataStoreId and globalDataStoreId for the specific 'detonation' DataStore specified by the parameters
        // check that the specified bombDataStore info matches the calculated info
        require(
            dataStoreProofs.bombDataStores[0].duration == (durationIndex + 1),
            "DataLayrBombVerifier.verifyBombDataStoreId: bomb datastore id's duration is the same as calculated"
        );
        require(
            dataStoreProofs.bombDataStores[0].metadata.durationDataStoreId == calculatedDataStoreId,
            "DataLayrBombVerifier.verifyBombDataStoreId: bomb datastore id provided is not the same as calculated"
        );

        // get the dataStoreId for 'detonationHeaderHash'
        // check that the dataStoreId for the provided detonationHeaderHash matches the calculated value
        require(
            dataStoreProofs.detonationDataStore.metadata.globalDataStoreId
                == nextGlobalDataStoreIdAfterDetonationTimestamp,
            "DataLayrBombVerifier.verifyBombDataStoreId: next datastore after bomb does not match provided detonation datastore"
        );
        // return globalDataStoreId at bomb DataStore
        return dataStoreProofs.bombDataStores[0].metadata.globalDataStoreId;
    }

    /**
     * @notice This function verifies the sandwich proof for each duration, then maps the `detonationHeaderHashValue` to one of the active datastores determined
     * by the sandwich proofs. This is the first potential BOMB datastore. The function returns the duration and id of the potential BOMB datastore and the
     * id of the earliest datastore after or at `detonationDataStoreInitTimestamp`
     *
     * @param detonationHeaderHashValue is the integer value of the DETONATION datastore's header hash
     * @param fromTime is the time from which the operator to slash hash been serving DataLayr
     * @param detonationDataStoreInitTimestamp is the time at which the DETONATION datastore was initialized
     * @param sandwichProofs are proofs for each duration of the different datastore ids that were active (not expired) at `detonationDataStoreInitTimestamp`
     *
     * @param sandwichProofs is a list of length equal to the number of durations that datastores can be stored. Each element is
     * 2 sandwich proofs of the datastores surrounding the boundaries of the duration. For example, if the first duration is 1 day,
     * then sandwichProofs[0][0] is a proof of the 2 datastores for duration 1 day surrounding @param detonationDataStoreInitTimestamp - 1 day or
     * @param fromTime. sandwichProofs[0][1] is a proof of the 2 datastores for duration 1 day surrounding @param detonationDataStoreInitTimestamp
     *
     * Then the BOMB datastore is picked from random by taking @param detonationHeaderHashValue and taking it modulo the number of active datastores
     * for example, if the current datastores for durations are
     *
     * Duration 1: ids 10-20
     * Duration 2: ids 1
     * Duration 3: ids 10-12
     * Duration 4: ids 10-22
     * Duration 5: ids 10-20
     * Duration 6: ids 10-20
     * Duration 7: ids 10-20
     * Duration 8: ids 10-20
     * Duration 9: ids 10-20
     * Duration 10: ids 10
     * Duration 11: ids 13-20
     * Duration 12: ids 10-20
     * Duration 13: none
     * Duration 14: ids 1-20
     *
     * and @param detonationHeaderHashValue is 100, the initially selected BOMB datastore will be
     * 100 - (11 datastores in Duration 1) - (1 datastores in Duration 2)  - (3 datastores in Duration 3)  - (13 datastores in Duration 4)
     * - (11 datastores in Duration 5) - (11 datastores in Duration 6) - (11 datastores in Duration 7) - (11 datastores in Duration 8)
     * - (11 datastores in Duration 9) - (1 datastores in Duration 10) - (8 datastores in Duration 11)
     * = the 8th datastore in Duration 12 which is: Duration 12, id 17
     *
     * If the challenged DLN did not sign off on the availability of the initial selected BOMB datastore, then the algorithm checks the
     * incremental datastores *by global datastore id*. This means that if Duration 12, id 17 was global datastore id 50, then it is checked whether
     * the DLN signed off on global datastore id 51, 52, ... until a signed datastore is found. There is a possible gas greiving attack that should
     * be thought about more here.
     *
     * @dev returns a pseudo-randomized durationIndex and durationDataStoreId, as well as the nextGlobalDataStoreIdAfterBomb
     */
    function verifySandwiches(
        uint256 detonationHeaderHashValue,
        uint256 fromTime,
        uint256 detonationDataStoreInitTimestamp,
        DataStoresForDuration[2][2][] calldata sandwichProofs
    )
        internal
        view
        returns (
            uint8 durationIndex,
            uint32 calculatedDataStoreId,
            uint32 nextGlobalDataStoreIdAfterDetonationTimestamp
        )
    {
        uint32 numberActiveDataStores;
        //This is a list of the number of active datastores for each duration
        //at the time of initialization of the DETONATION datastore
        uint32[] memory numberActiveDataStoresForDuration = new uint32[](
            dlsm.MAX_DATASTORE_DURATION()
        );
        //This is a list of the ids of the earliest active datastore for
        //each duration at the time of initialization of the DETONATION datastore
        uint32[] memory firstDataStoreForDuration = new uint32[](
            dlsm.MAX_DATASTORE_DURATION()
        );

        nextGlobalDataStoreIdAfterDetonationTimestamp = type(uint32).max;
        /**
         * For each duration,
         * If both timestamps in `sandwichProofs[i][0][0]` and `sandwichProofs[i][0][1]` are 0, then check that the number of datastore's for the
         * duration are 0 and go onto the next duration. This is a way of signifying that there have been no existing datastores for a duration.
         * If the sandwich proofs are nontrivial, set `sandwichTimestamp` equal to the later of 
             * when the operator started serving 
             * `detonationDataStoreInitTimestamp - DURATION_SCALE*duration`
         * since we only care about the datastores that were active at the time that were initialized after the operator registered.
         * Then, we verify the sandwich proof provided in `sandwichProof[i][0]` of the first datastore at or after `sandwichTimestamp`
         * and set `firstDataStoreForDuration[i]` equal to its id .
         * We also verify the sandwich proof provided in `sandwichProof[i][1]` of the first datastore at or after `detonationDataStoreInitTimestamp`
         * If the id of the datastore referred to by `sandwichProof[i][1][1]` is greater than the variable `nextGlobalDataStoreIdAfterDetonationTimestamp`,
         * then `nextGlobalDataStoreIdAfterDetonationTimestamp` is sent to `sandwichProof[i][1][1]`'s id.
         * Overall, `nextGlobalDataStoreIdAfterDetonationTimestamp` will be set to the id of datastore with the highest duration initialized
         * at `detonationDataStoreInitTimestamp`.
         * Then, `numberActiveDataStoresForDuration[i]` is set to `sandwichProof[i][1][1]`'s durationDataStoreId minus the `firstDataStoreForDuration[i]`
         * to get the number of active datastores for that duration at the `detonationDataStoreInitTimestamp`.
         * Finally, the sum `numberActiveDataStores` is increased by `numberActiveDataStoresForDuration[i]`.
         */
        for (uint8 i = 0; i < dlsm.MAX_DATASTORE_DURATION(); ++i) {
            // NOTE THAT i is loop index and (i + 1) is duration
            //If there are no datastores for certain duration, the prover should set the timestamps for the first sandwich proofs for that duration
            //equal to zero
            if (
                sandwichProofs[i][0][0].timestamp == sandwichProofs[i][0][1].timestamp
                    && sandwichProofs[i][0][0].timestamp == 0
            ) {
                //prover is claiming no datastores for given duration
                require(
                    dlsm.getNumDataStoresForDuration(i + 1) == 0,
                    "DataLayrBombVerifier.verifySandwiches: DataStores for duration are not 0"
                );
                //if storage agrees with provers claims, continue to next duration
                continue;
            }
            /**
             * Calculate the greater of ((init time of detonationDataStoreInitTimestamp) - duration) and fromTime.
             * Since 'fromTime' is the time at which the operator registered, if
             * fromTime is > (init time of detonationDataStoreInitTimestamp) - duration), then we only care about DataStores
             * starting from 'fromTime'
             */
            uint256 sandwichTimestamp =
                max(detonationDataStoreInitTimestamp - ((i + 1) * dlsm.DURATION_SCALE()), fromTime);
            /**
             * @dev Verify the sandwich proof for the given duratioxn. `verifyDataStoreIdSandwich` will return the the second datastore in the sandwich's metadata.
             * The second datastore is the first datastore after `sandwichTimestamp`. This is the first active datastore for the duration which was initialized at
             * or after the `sandwichTimestamp`. We store its durationDataStoreId in the `firstDataStoreForDuration` memory array.
             */
            firstDataStoreForDuration[i] =
                verifyDataStoreIdSandwich(sandwichTimestamp, i + 1, sandwichProofs[i][0]).durationDataStoreId;
            // verify the sandwich proof and store the metadata of the first datastore after `detonationDataStoreInitTimestamp` for the given duration
            IDataLayrServiceManager.DataStoreMetadata memory detonationDataStoreMetadata =
                verifyDataStoreIdSandwich(detonationDataStoreInitTimestamp, i + 1, sandwichProofs[i][1]);
            //The DETONATION datastore id is the nextGlobalDataStoreIdAfterDetonationTimestamp: the datastore with the lowest datastoreid after
            //the detonationDataStoreMetadata
            if (nextGlobalDataStoreIdAfterDetonationTimestamp > detonationDataStoreMetadata.globalDataStoreId) {
                nextGlobalDataStoreIdAfterDetonationTimestamp = detonationDataStoreMetadata.globalDataStoreId;
            }
            //record number of DataStores for duration
            numberActiveDataStoresForDuration[i] =
                detonationDataStoreMetadata.durationDataStoreId - firstDataStoreForDuration[i];
            // add number of DataStores (for this specific duration) to sum
            numberActiveDataStores += numberActiveDataStoresForDuration[i];
        }

        /**
         * Find the pseudo-randomly determined DataStore containing the bomb
         * by taking detonationHeaderHashValue modulo the *total* number of active datastores at the time
         */
        uint32 selectedDataStoreIndex = uint32(detonationHeaderHashValue % numberActiveDataStores);
        /**
         * Find the durationIndex and offset within the set of DataStores for that specific duration from the 'selectedDataStoreIndex'.
         * We can think of this as the DataStore 'location' specified by `selectedDataStoreIndex`, inside of a table of dataStores with one row per duration.
         */
        uint32 offset;
        (durationIndex, offset) = calculateCorrectIndexAndDurationOffsetFromNumberActiveDataStoresForDuration(
            selectedDataStoreIndex, numberActiveDataStoresForDuration
        );

        calculatedDataStoreId = firstDataStoreForDuration[durationIndex] + offset;
        /**
         * Return the pseudo-randomized `durationIndex` and `durationDataStoreId`, as specified by `selectedDataStoreIndex`,
         * as well as the `nextGlobalDataStoreIdAfterBomb`
         */
        return (durationIndex, calculatedDataStoreId, nextGlobalDataStoreIdAfterDetonationTimestamp);
    }

    /**
     * @notice Checks that the provided `sandwich` data accurately specifies the first dataStore, with the specified `duration`,
     * which was created at or after `sandwichTimestamp`
     * @notice For a certain @param duration, checks that the two datastores provided in @param sandwich
     * are the datastores just before and after (or equal) @param sandwichTimestamp, in that order
     * @notice For the given `duration` and `sandwichTimestamp`, this function returns the metadata of the earliest datastore that was stored
     * for `duration` and which was initialized after `sandwichTimestamp`.
     * @dev `sandwich[0]` is the search data for the latest datastore with `duration`, which was created *before* `sandwichTimestamp`. `sandwich[1]` is the
     * search data for the earliest datastore with the same `duration`, which was created *after* `sandwichTimestamp`.
     * @dev This function hashes the metadata in `sandwich` to verify its correctness and checks that the initialization times of the provided datastores are
     * before and after `sandwichTime`, respectively, and verifies that their ids are consecutive.
     * @return the metadata for the first dataStore, with the specified `duration`, which was created at or after `sandwichTimestamp`, i.e. the metadata
     * contained within `sandwich[1]`.
     */
    function verifyDataStoreIdSandwich(
        uint256 sandwichTimestamp,
        uint8 duration,
        DataStoresForDuration[2] calldata sandwich
    ) internal view returns (IDataLayrServiceManager.DataStoreMetadata memory) {
        // make sure that the first timestamp is strictly before the sandwichTimestamp
        require(
            sandwich[0].timestamp < sandwichTimestamp,
            "DataLayrBombVerifier.verifyDataStoreIdSandwich: sandwich[0].timestamp must be before sandwich time"
        );
        // make sure that the second timestamp is at or after the sandwichTimestamp
        require(
            sandwich[1].timestamp >= sandwichTimestamp || sandwich[1].timestamp == 0,
            "DataLayrBombVerifier.verifyDataStoreIdSandwich: sandwich[1].timestamp must be at or after sandwich time or 0"
        );

        /**
         * @dev If sandwichTimestamp is before the first datastore for the given duration, then sandwich[0].timestamp should be set equal to 0,
         * because there is no datastore before sandwichTimestamp for the duration.
         */
        if (sandwich[0].timestamp != 0) {
            /**
             * Since we've entered this code path, there *is* a datastore before sandwichTimestamp for the duration.
             * Verify that the provided metadata of the datastore before sandwichTimestamp (i.e. sandwich[0]) agrees with the stored hash.
             */
            require(
                dlsm.verifyDataStoreMetadata(
                    duration,
                    sandwich[0].timestamp,
                    sandwich[0].index,
                    sandwich[0].metadata
                ), "DataLayrBombVerifier.verifyDataStoreIdSandwich: sandwich[0].metadata preimage is incorrect"
            );

        } else {
            //if there is no data stores for the duration, then make sure metadata is consistent with that for future checks
            require(
                sandwich[0].metadata.durationDataStoreId == 0,
                "DataLayrBombVerifier.verifyDataStoreIdSandwich: sandwich[0].timstamp was 0 but duration datastore id was not"
            );
        }
        /**
         * @dev If sandwichTimestamp is after the last datastore for the given duration, then sandwich[1].timestamp should be set equal to 0,
         * because there is no datastore at or after sandwichTimestamp for the duration.
         */
        if (sandwich[1].timestamp != 0) {
            /**
             * Since we've entered this code path, there *is* a datastore at or after sandwichTimestamp for the duration.
             * Verify that the provided metadata of the datastore before sandwichTimestamp (i.e. sandwich[1]) agrees with the stored hash.
             */
            require(
                dlsm.verifyDataStoreMetadata(
                    duration,
                    sandwich[1].timestamp,
                    sandwich[1].index,
                    sandwich[1].metadata
                ),"DataLayrBombVerifier.verifyDataStoreIdSandwich: sandwich[1].metadata preimage is incorrect"
            );

            // make sure that sandwich[0] and sandwich[1] are consecutive datastores for the duration by checking that their
            // durationDataStoreIds are consecutive
            require(
                sandwich[0].metadata.durationDataStoreId + 1 == sandwich[1].metadata.durationDataStoreId,
                "DataLayrBombVerifier.verifyDataStoreIdSandwich: x and y datastore must be incremental or y datastore is not first in the duration"
            );
        } else {
            // if sandwich[1].timestamp is 0, the prover is claiming that there is no datastore at or after sandwichTimestamp for the duration
            require(
                dlsm.getNumDataStoresForDuration(duration) == sandwich[0].metadata.durationDataStoreId,
                "DataLayrBombVerifier.verifyDataStoreIdSandwich: x datastore is not the last datastore in the duration or no datastores for duration"
            );
        }
        return sandwich[1].metadata;
    }

    /**
     * @notice given an ordered list of groups and the number of elements in each group, as well as a total offset, calculates which group and index within the group
     * the offset points to.
     * @dev Inputs are a pseudo-random 'offset' value and an array of the number of active DataStores, ordered by duration.
     * Given the 'offset' value, this function moves through the 'duration' bins, and returns the bin and offset *within that bin* corresponding to 'offset'.
     * In other words, it finds the position for the 'offset'-th entry, specified by a duration 'bin'
     * and a value corresponding to the index of a DataStore within that bin.
     */
    function calculateCorrectIndexAndDurationOffsetFromNumberActiveDataStoresForDuration(
        uint32 offset,
        uint32[] memory numberActiveDataStoresForDuration
    ) internal pure returns (uint8 durationIndex, uint32 offsetRemaining) {
        offsetRemaining = offset;
        durationIndex = 0;
        for (; durationIndex < numberActiveDataStoresForDuration.length; ++durationIndex) {
            //we use > not >= because offsetRemaining should be the index within the correct duration
            if (numberActiveDataStoresForDuration[durationIndex] > offsetRemaining) {
                break;
            }
            offsetRemaining -= numberActiveDataStoresForDuration[durationIndex];
        }

        return (uint8(durationIndex), offsetRemaining);
    }

    /**
     * @notice Gets the specific coset/chunk number that `operator` was assigned for the datastore specified by `searchData`.
     * @param operatorIndex and @param totalOperatorsIndex are the indexes in their relative arrays in the `dlRegistry` that
     * are used to prove the total number of operators and the index of `operator` at `searchData.metadata.blockNumber`.
     * These are used to calculate which cosets were used for the datastore, based off of other information in the header.
     */
    function getChunkNumber(
        address operator,
        uint32 operatorIndex,
        uint32 totalOperatorsIndex,
        IDataLayrServiceManager.DataStoreSearchData calldata searchData
    ) internal view returns (uint32 chunkNumber) {
        // Verify that the provided `searchData` is correct
        require(
            dlsm.verifyDataStoreMetadata(
                searchData.duration,
                searchData.timestamp,
                searchData.index,
                searchData.metadata
            ), "DataLayrBombVerifier.getChunkNumber: search.metadataclear preimage is incorrect"
        );

        // Check that the specified dataStore has been confirmed
        require(searchData.metadata.signatoryRecordHash != bytes32(0), "Datastore is not committed yet");

        /**
         * Get the index of the given operator at `searchData.metadata.blockNumber` via the `operatorIndex` input,
         * and checks in the registry to make sure the index is accurate
         */
        operatorIndex = dlRegistry.getOperatorIndex(operator, searchData.metadata.blockNumber, operatorIndex);

        /**
         * Gets the totalNumber of operators at `searchData.metadata.blockNumber` via the `totalOperatorsIndex` input,
         * and checks in the registry to make sure the index is accurate
         */
        totalOperatorsIndex = dlRegistry.getTotalOperators(searchData.metadata.blockNumber, totalOperatorsIndex);

        // Calculate the coset given to the operator
        return (operatorIndex + searchData.metadata.globalDataStoreId) % totalOperatorsIndex;
    }

    /**
     * @notice This function verifies that `disclosureProof.poly` was the data that operator was storing for `dataStoreId` and it is the polynomial that is
     * commited to by the KZG commitment `disclosureProof.interpolationPoly`.
     * @dev The coset of the zero polynomial is determined from `operatorIndex` and `totalOperatorsIndex` and the metadata of the datastore provided by `searchData`.
     * @return 'true' if the proof succeeded. The function will revert otherwise.
     */
    function nonInteractivePolynomialProof(
        address operator,
        uint32 operatorIndex,
        uint32 totalOperatorsIndex,
        uint32 dataStoreId,
        DisclosureProof calldata disclosureProof,
        IDataLayrServiceManager.DataStoreSearchData calldata searchData
    ) internal view returns (bool) {
        // Fetch the chunkNumber that `operator` recieved for `dataStoreId`
        uint32 chunkNumber = getChunkNumber(operator, operatorIndex, totalOperatorsIndex, searchData);
        // Make sure that `dataStoreId` is consistent with the datastore in `searchData`
        require(
            searchData.metadata.globalDataStoreId == dataStoreId,
            "DataLayrBombVerifier.nonInteractivePolynomialProof: searchData does not match provided dataStoreId"
        );
        // Ensure that the headerhash in `searchData` is consistent with the header that is revealed against in `disclosureProof`
        require(
            searchData.metadata.headerHash == keccak256(disclosureProof.header),
            "DataLayrBombVerifier.nonInteractivePolynomialProof: hash of dislosure proof header does not match provided searchData"
        );

        /**
         * Verify that `disclosureProof.poly` is the data that is stored at the zero polynomial defined by `chunkNumber` against the overarching
         * polynomial commitment defined in `disclosureProof.header`
         */
        return challengeUtils.nonInteractivePolynomialProof(
            disclosureProof.header,
            chunkNumber,
            disclosureProof.poly,
            disclosureProof.multiRevealProof,
            disclosureProof.polyEquivalenceProof
        );
    }

    /// @notice This function verifies that `searchData` provided matches the data that was stored for the given datastore
    function verifyMetadataPreImage(IDataLayrServiceManager.DataStoreSearchData calldata searchData)
        internal
        view
        returns (bool)
    {
        return dlsm.getDataStoreHashesForDurationAtTimestamp(
            searchData.duration, searchData.timestamp, searchData.index
        ) == DataStoreUtils.computeDataStoreHash(searchData.metadata);
    }

    function max(uint256 x, uint256 y) internal pure returns (uint256) {
        return x > y ? x : y;
    }
}

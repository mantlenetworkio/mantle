// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "../interfaces/IBLSRegistry.sol";
import "../libraries/BytesLib.sol";
import "../libraries/DataStoreUtils.sol";
import "../libraries/BLS.sol";
import "../permissions/RepositoryAccess.sol";

/**
 * @title Used for checking BLS aggregate signatures from the operators of a `BLSRegistry`.
 * @author Layr Labs, Inc.
 * @notice This is the contract for checking the validity of aggregate operator signatures.
 */
abstract contract BLSSignatureChecker is RepositoryAccess {
    using BytesLib for bytes;
    // DATA STRUCTURES
    /**
     * @notice this data structure is used for recording the details on the total stake of the registered
     * operators and those operators who are part of the quorum for a particular taskNumber
     */

    struct SignatoryTotals {
        // total stake of the operators who are in the first quorum
        uint256 signedStakeFirstQuorum;
        // total stake of the operators who are in the second quorum
        uint256 signedStakeSecondQuorum;
        // total amount staked by all operators (irrespective of whether they are in the quorum or not)
        uint256 totalStakeFirstQuorum;
        // total amount staked by all operators (irrespective of whether they are in the quorum or not)
        uint256 totalStakeSecondQuorum;
    }

    // EVENTS
    /**
     * @notice used for recording the event that signature has been checked in checkSignatures function.
     */
    event SignatoryRecord(
        bytes32 msgHash,
        uint32 taskNumber,
        uint256 signedStakeFirstQuorum,
        uint256 signedStakeSecondQuorum,
        // uint256 totalStakeFirstQuorum,
        // uint256 totalStakeSecondQuorum,
        bytes32[] pubkeyHashes
    );

    // solhint-disable-next-line no-empty-blocks
    constructor(IRepository _repository) RepositoryAccess(_repository) {}

    // CONSTANTS -- commented out lines are due to inline assembly supporting *only* 'direct number constants' (for now, at least)
    uint256 internal constant BYTE_LENGTH_totalStakeIndex = 6;
    uint256 internal constant BYTE_LENGTH_stakesBlockNumber = 4;
    uint256 internal constant BYTE_LENGTH_taskNumberToConfirm = 4;
    uint256 internal constant BYTE_LENGTH_numberNonSigners = 4;
    // specifying a G2 public key requires 4 32-byte slots worth of data
    uint256 internal constant BYTE_LENGTH_PUBLIC_KEY = 128;
    uint256 internal constant BYTE_LENGTH_stakeIndex = 4;
    // uint256 internal constant BYTE_LENGTH_NON_SIGNER_INFO = BYTE_LENGTH_PUBLIC_KEY + BYTE_LENGTH_stakeIndex;
    uint256 internal constant BYTE_LENGTH_NON_SIGNER_INFO = 132;
    uint256 internal constant BYTE_LENGTH_apkIndex = 4;

    // uint256 internal constant BIT_SHIFT_totalStakeIndex = 256 - (BYTE_LENGTH_totalStakeIndex * 8);
    uint256 internal constant BIT_SHIFT_totalStakeIndex = 208;
    // uint256 internal constant BIT_SHIFT_stakesBlockNumber = 256 - (BYTE_LENGTH_stakesBlockNumber * 8);
    uint256 internal constant BIT_SHIFT_stakesBlockNumber = 224;
    // uint256 internal constant BIT_SHIFT_taskNumberToConfirm = 256 - (BYTE_LENGTH_taskNumberToConfirm * 8);
    uint256 internal constant BIT_SHIFT_taskNumberToConfirm = 224;
    // uint256 internal constant BIT_SHIFT_numberNonSigners = 256 - (BYTE_LENGTH_numberNonSigners * 8);
    uint256 internal constant BIT_SHIFT_numberNonSigners = 224;
    // uint256 internal constant BIT_SHIFT_stakeIndex = 256 - (BYTE_LENGTH_stakeIndex * 8);
    uint256 internal constant BIT_SHIFT_stakeIndex = 224;
    // uint256 internal constant BIT_SHIFT_apkIndex = 256 - (BYTE_LENGTH_apkIndex * 8);
    uint256 internal constant BIT_SHIFT_apkIndex = 224;

    uint256 internal constant CALLDATA_OFFSET_totalStakeIndex = 32;
    // uint256 internal constant CALLDATA_OFFSET_stakesBlockNumber = CALLDATA_OFFSET_totalStakeIndex + BYTE_LENGTH_totalStakeIndex;
    uint256 internal constant CALLDATA_OFFSET_stakesBlockNumber = 38;
    // uint256 internal constant CALLDATA_OFFSET_taskNumberToConfirm = CALLDATA_OFFSET_stakesBlockNumber + BYTE_LENGTH_stakesBlockNumber;
    uint256 internal constant CALLDATA_OFFSET_taskNumberToConfirm = 42;
    // uint256 internal constant CALLDATA_OFFSET_numberNonSigners = CALLDATA_OFFSET_taskNumberToConfirm + BYTE_LENGTH_taskNumberToConfirm;
    uint256 internal constant CALLDATA_OFFSET_numberNonSigners = 46;
    // uint256 internal constant CALLDATA_OFFSET_NonsignerPubkeys = CALLDATA_OFFSET_numberNonSigners + BYTE_LENGTH_numberNonSigners;
    uint256 internal constant CALLDATA_OFFSET_NonsignerPubkeys = 50;

    /**
     * @notice This function is called by disperser when it has aggregated all the signatures of the operators
     * that are part of the quorum for a particular taskNumber and is asserting them into on-chain. The function
     * checks that the claim for aggregated signatures are valid.
     *
     * The thesis of this procedure entails:
     * - computing the aggregated pubkey of all the operators that are not part of the quorum for
     * this specific taskNumber (represented by aggNonSignerPubkey)
     * - getting the aggregated pubkey of all registered nodes at the time of pre-commit by the
     * disperser (represented by pk),
     * - do subtraction of aggNonSignerPubkey from pk over Jacobian coordinate system to get aggregated pubkey
     * of all operators that are part of quorum.
     * - use this aggregated pubkey to verify the aggregated signature under BLS scheme.
     */
    /**
     * @dev This calldata is of the format:
     * <
     * bytes32 msgHash, the taskHash for which disperser is calling checkSignatures
     * uint48 index of the totalStake corresponding to the dataStoreId in the 'totalStakeHistory' array of the BLSRegistryWithBomb
     * uint32 blockNumber, the blockNumber at which the task was initated
     * uint32 taskNumberToConfirm
     * uint32 numberOfNonSigners,
     * {uint256[4], apkIndex}[numberOfNonSigners] the public key and the index to query of `pubkeyHashToStakeHistory` for each nonsigner,
     * in affine coordinates, arranges as (x_0, x_1), (y_0, y_1)
     * uint32 stakeIndex is the index in the stake history from which quorum stake info is retreived.
     * uint32 apkIndex, the index in the `apkUpdates` array at which we want to load the aggregate public key
     * uint256[4] apk (aggregate public key),
     * uint256[2] sigma, the aggregate signature itself
     * >
     * 
     * @dev Before signature verification, the function verifies operator stake information.  This includes ensuring that the provided `stakesBlockNumber`
     * is correct, i.e., ensure that the stake returned from the specified block number is recent enough and that the stake is either the most recent update
     * for the total stake (or the operator) or latest before the stakesBlockNumber.
     * The next step involves computing the aggregated pub key of all the operators that are not part of the quorum for this specific taskNumber.
     * We use a loop to iterate through the `nonSignerPK` array, loading each individual public key from calldata. Before the loop, we isolate the first public key
     * calldataload - this implementation saves us one `BLS.addJac` operation, which would be performed in the i=0 iteration otherwise.
     * Within the loop, each non-signer public key is loaded from the calldata into memory.  The most recent staking-related information is retrieved and is subtracted
     * from the total stake of validators in the quorum.  Then the aggregate public key and the aggregate non-signer public key is subtracted from it.
     * Finally  the siganture is verified by computing the elliptic curve pairing.
     */
    function checkSignatures(bytes calldata data)
        public
        returns (
            uint32 taskNumberToConfirm,
            uint32 stakesBlockNumber,
            bytes32 msgHash,
            SignatoryTotals memory signedTotals,
            bytes32 compressedSignatoryRecord
        )
    {
        // temporary variable used to hold various numbers
        uint256 placeholder;

        uint256 pointer;

        assembly {
            pointer := data.offset
            /**
             * Get the 32 bytes immediately after the function signature and length + offset encoding of 'bytes
             * calldata' input type, which represents the msgHash for which the disperser is calling `checkSignatures`
             */
            msgHash := calldataload(pointer)

            // Get the 6 bytes immediately after the above, which represent the index of the totalStake in the 'totalStakeHistory' array
            placeholder := shr(BIT_SHIFT_totalStakeIndex, calldataload(add(pointer, CALLDATA_OFFSET_totalStakeIndex)))
        }

        // fetch the 4 byte stakesBlockNumber, the block number from which stakes are going to be read from
        assembly {
            stakesBlockNumber :=
                shr(BIT_SHIFT_stakesBlockNumber, calldataload(add(pointer, CALLDATA_OFFSET_stakesBlockNumber)))
        }

        // obtain registry contract for querying information on stake later
        IBLSRegistry registry = IBLSRegistry(address(_registry()));

        /**
         * @dev Instantiate the memory object used for holding the aggregated public key of all operators that are *not* part of the quorum.
         * @dev Note that we are storing points in G2 using Jacobian coordinates - [x0, x1, y0, y1, z0, z1]
         */
        uint256[6] memory aggNonSignerPubkey;

        // get information on total stakes
        IQuorumRegistry.OperatorStake memory localStakeObject = registry.getTotalStakeFromIndex(placeholder);

        // check that the returned OperatorStake object is the most recent for the stakesBlockNumber
        _validateOperatorStake(localStakeObject, stakesBlockNumber);

        // copy total stakes amounts to `signedTotals` -- the 'signedStake' amounts are decreased later, to reflect non-signers
        signedTotals.totalStakeFirstQuorum = localStakeObject.firstQuorumStake;
        signedTotals.signedStakeFirstQuorum = localStakeObject.firstQuorumStake;
        signedTotals.totalStakeSecondQuorum = localStakeObject.secondQuorumStake;
        signedTotals.signedStakeSecondQuorum = localStakeObject.secondQuorumStake;

        assembly {
            //fetch the task number to avoid replay signing on same taskhash for different datastore
            taskNumberToConfirm :=
                shr(BIT_SHIFT_taskNumberToConfirm, calldataload(add(pointer, CALLDATA_OFFSET_taskNumberToConfirm)))
            // get the 4 bytes immediately after the above, which represent the
            // number of operators that aren't present in the quorum
            // slither-disable-next-line write-after-write
            placeholder := shr(BIT_SHIFT_numberNonSigners, calldataload(add(pointer, CALLDATA_OFFSET_numberNonSigners)))
        }

        // we have read (32 + 6 + 4 + 4 + 4) = 50 bytes of calldata so far
        pointer += CALLDATA_OFFSET_NonsignerPubkeys;

        // to be used for holding the pub key hashes of the operators that aren't part of the quorum
        bytes32[] memory pubkeyHashes = new bytes32[](placeholder);

        /**
         * @dev The next step involves computing the aggregated pub key of all the operators
         * that are not part of the quorum for this specific taskNumber.
         */

        /**
         * @dev loading pubkey for the first operator that is not part of the quorum as listed in the calldata;
         * Note that this need not be a special case and *could* be subsumed in the for loop below.
         * However, this implementation saves one 'addJac' operation, which would be performed in the i=0 iteration otherwise.
         * @dev Recall that `placeholder` here is the number of operators *not* included in the quorum
         */
        if (placeholder != 0) {
            //load compressed pubkey and the index in the stakes array into memory
            uint32 stakeIndex;
            assembly {
                /**
                 * @notice retrieving the pubkey of the node in Jacobian coordinates
                 */
                // sigma_x0
                mstore(aggNonSignerPubkey, calldataload(pointer))

                // sigma_x1
                mstore(add(aggNonSignerPubkey, 0x20), calldataload(add(pointer, 32)))

                // sigma_y0
                mstore(add(aggNonSignerPubkey, 0x40), calldataload(add(pointer, 64)))

                // sigma_y1
                mstore(add(aggNonSignerPubkey, 0x60), calldataload(add(pointer, 96)))

                // converting Affine coordinates to Jacobian coordinates
                // [(x_0, x_1), (y_0, y_1)] => [(x_0, x_1), (y_0, y_1), (1,0)]
                // source: https://crypto.stackexchange.com/questions/19598/how-can-convert-affine-to-jacobian-coordinates
                // sigma_z0
                mstore(add(aggNonSignerPubkey, 0x80), 1)
                // sigma_z1
                mstore(add(aggNonSignerPubkey, 0xA0), 0)

                /**
                 * @notice retrieving the index of the stake of the operator in pubkeyHashToStakeHistory in
                 * Registry.sol that was recorded at the time of pre-commit.
                 */
                stakeIndex := shr(BIT_SHIFT_stakeIndex, calldataload(add(pointer, BYTE_LENGTH_PUBLIC_KEY)))
            }
            // We have read (32 + 32 + 32 + 32 + 4) = 132 additional bytes of calldata in the above assembly block.
            // Update pointer accordingly.
            unchecked {
                pointer += BYTE_LENGTH_NON_SIGNER_INFO;
            }

            // get pubkeyHash and add it to pubkeyHashes of operators that aren't part of the quorum.
            bytes32 pubkeyHash = BLS.hashPubkey(aggNonSignerPubkey);

            pubkeyHashes[0] = pubkeyHash;

            // querying the VoteWeigher for getting information on the operator's stake
            // at the time of pre-commit
            localStakeObject = registry.getStakeFromPubkeyHashAndIndex(pubkeyHash, stakeIndex);

            // check that the returned OperatorStake object is the most recent for the stakesBlockNumber
            _validateOperatorStake(localStakeObject, stakesBlockNumber);

            // subtract operator stakes from totals
            signedTotals.signedStakeFirstQuorum -= localStakeObject.firstQuorumStake;
            signedTotals.signedStakeSecondQuorum -= localStakeObject.secondQuorumStake;
        }

        // temporary variable for storing the pubkey of operators in Jacobian coordinates
        uint256[6] memory pk;
        pk[4] = 1;

        for (uint256 i = 1; i < placeholder;) {
            //load compressed pubkey and the index in the stakes array into memory
            uint32 stakeIndex;
            assembly {
                /// @notice retrieving the pubkey of the operator that is not part of the quorum
                mstore(pk, calldataload(pointer))
                mstore(add(pk, 0x20), calldataload(add(pointer, 32)))
                mstore(add(pk, 0x40), calldataload(add(pointer, 64)))
                mstore(add(pk, 0x60), calldataload(add(pointer, 96)))

                /**
                 * @notice retrieving the index of the stake of the operator in pubkeyHashToStakeHistory in
                 * Registry.sol that was recorded at the time of pre-commit.
                 */
                // slither-disable-next-line variable-scope
                stakeIndex := shr(BIT_SHIFT_stakeIndex, calldataload(add(pointer, BYTE_LENGTH_PUBLIC_KEY)))
            }

            // We have read (32 + 32 + 32 + 32 + 4) = 132 additional bytes of calldata in the above assembly block.
            // Update pointer accordingly.
            unchecked {
                pointer += BYTE_LENGTH_NON_SIGNER_INFO;
            }

            // get pubkeyHash and add it to pubkeyHashes of operators that aren't part of the quorum.
            bytes32 pubkeyHash = BLS.hashPubkey(pk);

            //pubkeys should be ordered in ascending order of hash to make proofs of signing or
            // non signing constant time
            /**
             * @dev this invariant is used in forceOperatorToDisclose in ServiceManager.sol
             */
            require(uint256(pubkeyHash) > uint256(pubkeyHashes[i - 1]), "Pubkey hashes must be in ascending order");

            // recording the pubkey hash
            pubkeyHashes[i] = pubkeyHash;

            // querying the VoteWeigher for getting information on the operator's stake
            // at the time of pre-commit
            localStakeObject = registry.getStakeFromPubkeyHashAndIndex(pubkeyHash, stakeIndex);

            // check that the returned OperatorStake object is the most recent for the stakesBlockNumber
            _validateOperatorStake(localStakeObject, stakesBlockNumber);

            //subtract validator stakes from totals
            signedTotals.signedStakeFirstQuorum -= localStakeObject.firstQuorumStake;
            signedTotals.signedStakeSecondQuorum -= localStakeObject.secondQuorumStake;

            // add the pubkey of the operator to the aggregate pubkeys in Jacobian coordinate system.
            // slither-disable-next-line unused-return
            BLS.addJac(aggNonSignerPubkey, pk);

            unchecked {
                ++i;
            }
        }
        // usage of a scoped block here minorly decreases gas usage
        {
            uint32 apkIndex;
            assembly {
                //get next 32 bits which would be the apkIndex of apkUpdates in Registry.sol
                apkIndex := shr(BIT_SHIFT_apkIndex, calldataload(pointer))

                // Update pointer to account for the 4 bytes specifying the apkIndex
                pointer := add(pointer, BYTE_LENGTH_apkIndex)

                /**
                 * @notice Get the aggregated publickey at the moment when pre-commit happened
                 * @devAaggregated pubkey given as part of calldata instead of being retrieved from voteWeigher reduces number of SLOADs
                 */
                mstore(pk, calldataload(pointer))
                mstore(add(pk, 0x20), calldataload(add(pointer, 32)))
                mstore(add(pk, 0x40), calldataload(add(pointer, 64)))
                mstore(add(pk, 0x60), calldataload(add(pointer, 96)))
            }

            // We have read (32 + 32 + 32 + 32) = 128 additional bytes of calldata in the above assembly block.
            // Update pointer accordingly.
            unchecked {
                pointer += BYTE_LENGTH_PUBLIC_KEY;
            }

            // make sure the caller has provided the correct aggPubKey
            require(
                registry.getCorrectApkHash(apkIndex, stakesBlockNumber) == BLS.hashPubkey(pk),
                "BLSSignatureChecker.checkSignatures: Incorrect apk provided"
            );
        }

        // input for call to ecPairing precompile contract
        uint256[12] memory input = [
            uint256(0),
            uint256(0),
            uint256(0),
            uint256(0),
            uint256(0),
            uint256(0),
            uint256(0),
            uint256(0),
            uint256(0),
            uint256(0),
            uint256(0),
            uint256(0)
        ];

        // if at least 1 non-signer
        if (placeholder != 0) {
            /**
             * @notice need to subtract aggNonSignerPubkey from the apk to get aggregate signature of all
             * operators that are part of the quorum
             */
            // negate aggNonSignerPubkey
            aggNonSignerPubkey[2] = (BLS.MODULUS - aggNonSignerPubkey[2]) % BLS.MODULUS;
            aggNonSignerPubkey[3] = (BLS.MODULUS - aggNonSignerPubkey[3]) % BLS.MODULUS;

            // do the addition in Jacobian coordinates
            // slither-disable-next-line unused-return
            BLS.addJac(pk, aggNonSignerPubkey);

            // reorder for pairing
            (input[3], input[2], input[5], input[4]) = BLS.jacToAff(pk);
            // if zero non-signers
        } else {
            //else copy it to input
            //reorder for pairing
            (input[3], input[2], input[5], input[4]) = (pk[0], pk[1], pk[2], pk[3]);
        }

        /**
         * @notice now we verify that e(H(m), pk)e(sigma, -g2) == 1
         */

        // compute the point in G1
        (input[0], input[1]) = BLS.hashToG1(msgHash);

        // insert negated coordinates of the generator for G2
        input[8] = BLS.nG2x1;
        input[9] = BLS.nG2x0;
        input[10] = BLS.nG2y1;
        input[11] = BLS.nG2y0;

        assembly {
            // next in calldata are the signatures
            // sigma_x0
            mstore(add(input, 0xC0), calldataload(pointer))
            // sigma_x1
            mstore(add(input, 0xE0), calldataload(add(pointer, 0x20)))

            // check the pairing; if incorrect, revert
            if iszero(
                // staticcall address 8 (ecPairing precompile), forward all gas, send 384 bytes (0x180 in hex) = 12 (32-byte) inputs.
                // store the return data in input[11] (352 bytes / '0x160' in hex), and copy only 32 bytes of return data (since precompile returns boolean)
                staticcall(not(0), 0x08, input, 0x180, add(input, 0x160), 0x20)
            ) { revert(0, 0) }
        }

        // check that the provided signature is correct
        require(input[11] == 1, "BLSSignatureChecker.checkSignatures: Pairing unsuccessful");

        emit SignatoryRecord(
            msgHash,
            taskNumberToConfirm,
            signedTotals.signedStakeFirstQuorum,
            signedTotals.signedStakeSecondQuorum,
            // signedTotals.totalStakeFirstQuorum,
            // signedTotals.totalStakeSecondQuorum,
            pubkeyHashes
            );

        // set compressedSignatoryRecord variable used for fraudproofs
        compressedSignatoryRecord = DataStoreUtils.computeSignatoryRecordHash(
            taskNumberToConfirm,
            pubkeyHashes,
            signedTotals.signedStakeFirstQuorum,
            signedTotals.signedStakeSecondQuorum
        );

        // return taskNumber, stakesBlockNumber, msgHash, total stakes that signed, and a hash of the signatories
        return (taskNumberToConfirm, stakesBlockNumber, msgHash, signedTotals, compressedSignatoryRecord);
    }

    // simple internal function for validating that the OperatorStake returned from a specified index is the correct one
    function _validateOperatorStake(IQuorumRegistry.OperatorStake memory opStake, uint32 stakesBlockNumber)
        internal
        pure
    {
        // check that the stake returned from the specified index is recent enough
        require(opStake.updateBlockNumber <= stakesBlockNumber, "Provided stake index is too early");

        /**
         * check that stake is either the most recent update for the total stake (or the operator),
         * or latest before the stakesBlockNumber
         */
        require(
            opStake.nextUpdateBlockNumber == 0 || opStake.nextUpdateBlockNumber > stakesBlockNumber,
            "Provided stake index is not the most recent for blockNumber"
        );
    }
}

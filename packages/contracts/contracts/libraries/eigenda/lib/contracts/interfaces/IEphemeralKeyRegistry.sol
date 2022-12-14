// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

/**
 * @title Interface for an Ephemeral Key Registry, designed for use with Proofs of Custody.
 * @author Layr Labs, Inc.
 * @notice See the Dankrad's excellent article for an intro to Proofs of Custody:
 * https://dankradfeist.de/ethereum/2021/09/30/proofs-of-custody.html.
 */

interface IEphemeralKeyRegistry {
    /**
     * @notice Used by operator to post their first ephemeral key hash via BLSRegistry (on registration).
     * This effectively serves as a commitment to the ephemeral key - when it is revealed during the disclosure period, it can be verified against the hash.
     * @param EKHash is the hash of the Ephemeral key that is being currently used by the
     * @param operator for signing on bomb-based queries.
     */
    function postFirstEphemeralKeyHash(address operator, bytes32 EKHash) external;

    /**
     * @notice Used by the operator to post their ephemeral key preimage via BLSRegistry (on degregistration) after the expiry of its usage.
     * This function is called only when operator is going to de-register from the middleware.
     * Check its usage in the `deregisterOperator` function of the BLSRegistryWithBomb contract.
     * @param prevEK is the preimage.
     */
    function postLastEphemeralKeyPreImage(address operator, bytes32 prevEK) external;

    /**
     * @notice Used by the operator to update their ephemeral key hash and post their previous ephemeral key after the expiry of its usage.
     * Revealing of current ephemeral key and describing the hash of the new ephemeral key done together.
     * @param prevEK is the previous ephemeral key, checked against the `msg.sender`'s existing keyHash.
     * @param newEKHash is the hash of the new ephemeral key.
     * @dev The function must be called within the `REVEAL_PERIOD` which is the time window within which an operator can reveal their EK preimage
     * without the risk of being frontrun and slashed (outside of the reveal period, if an attacker sees the preimage in the mempool, they can frontrun
     * that reveal transaction and cause slashing of the operator).  
     */
    function revealAndUpdateEphemeralKey(bytes32 prevEK, bytes32 newEKHash) external;

    /// @notice retrieve the operator's current ephemeral key hash
    function getCurrEphemeralKeyHash(address operator) external returns (bytes32);

    /// @notice retrieve the operator's current ephemeral key itself
    function getLatestEphemeralKey(address operator) external returns (bytes32);

    /**
     * @notice Used for proving that an operator hasn't reveal their ephemeral key within the reveal window, triggering slashing of the operator.
     * @param operator The operator with a stale ephemeral key
     */
    function proveStaleUnrevealedEphemeralKey(address operator) external;

    /**
     * @notice Used for proving that an operator leaked an ephemeral key that was not supposed to be revealed, triggering slashing of the operator.
     * @param operator The operator who leaked their ephemeral key.
     * @param leakedEphemeralKey The ephemeral key for the operator, which they were not supposed to reveal.
     */
    function proveLeakedEphemeralKey(address operator, bytes32 leakedEphemeralKey) external;

    /// @notice Returns the UTC timestamp at which the operator last renewed their ephemeral key
    function getLastEKPostTimestamp(address operator) external returns (uint192);

    /**
     * @notice This function is used for getting the ephemeral key pertaining to a particular taskNumber, for an operator
     * @param operator The operator whose ephemeral key we are interested in.
     * @param taskNumber The taskNumber for which we want to retrieve the operator's ephemeral key
     */
    function getEphemeralKeyForTaskNumber(address operator, uint32 taskNumber) external view returns (bytes32);
}

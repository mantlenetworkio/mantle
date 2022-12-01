// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../interfaces/IEphemeralKeyRegistry.sol";
import "../interfaces/IQuorumRegistry.sol";
import "../permissions/RepositoryAccess.sol";

import "forge-std/Test.sol";

/**
 * @title Registry of Ephemeral Keys for operators, designed for use with Proofs of Custody.
 * @author Layr Labs, Inc.
 * @notice This contract has the functionality for ---
 * (1) storing revealed ephemeral keys for each operator from past,
 * (2) checking if ephemeral keys revealed too early and then slashing if needed,
 * (3) recording when a previous ephemeral key is made inactive
 * @notice See the Dankrad's excellent article for an intro to Proofs of Custody:
 * https://dankradfeist.de/ethereum/2021/09/30/proofs-of-custody.html.
 */
contract EphemeralKeyRegistry is IEphemeralKeyRegistry, RepositoryAccess, DSTest {
    // DATA STRUCTURES
    struct EKEntry {
        // hash of the ephemeral key, to be revealed after usage
        bytes32 keyHash;
        // the ephemeral key itself, i.e. the preimage of `keyHash`
        bytes32 ephemeralKey;
        // timestamp when the keyhash is first recorded
        uint192 timestamp;
        // task number of the middleware from which ephemeral key is being used
        uint32 startTaskNumber;
        // task number of the middleware  until which ephemeral key is being used
        uint32 endTaskNumber;
    }

    // max amount of time that an operator can take to update their ephemeral key
    uint256 public constant UPDATE_PERIOD = 18 days;

    // max amout of time operator has to submit and confirm the ephemeral key reveal transaction
    uint256 public constant REVEAL_PERIOD = 3 days;

    // operator => history of ephemeral keys, hashes of them, timestamp at which they were posted, and start/end taskNumbers
    mapping(address => EKEntry[]) public EKHistory;

    // solhint-disable-next-line no-empty-blocks
    constructor(IRepository _repository) RepositoryAccess(_repository) {}

    /**
     * @notice Used by operator to post their first ephemeral key hash via BLSRegistry (on registration).
     * This effectively serves as a commitment to the ephemeral key - when it is revealed during the disclosure period, it can be verified against the hash.
     * @param EKHash is the hash of the Ephemeral key that is being currently used by the
     * @param operator for signing on bomb-based queries.
     */
    function postFirstEphemeralKeyHash(address operator, bytes32 EKHash) external onlyRegistry {
        // record the new EK entry
        EKHistory[operator].push(
            EKEntry({
                keyHash: EKHash,
                ephemeralKey: bytes32(0),
                timestamp: uint192(block.timestamp),
                startTaskNumber: repository.serviceManager().taskNumber(),
                endTaskNumber: 0
            })
        );
    }

    /**
     * @notice Used by the operator to post their ephemeral key preimage via BLSRegistry (on degregistration) after the expiry of its usage.
     * This function is called only when operator is going to de-register from the middleware.
     * Check its usage in the `deregisterOperator` function of the BLSRegistryWithBomb contract.
     * @param prevEK is the preimage.
     */
    function postLastEphemeralKeyPreImage(address operator, bytes32 prevEK) external {
        IQuorumRegistry registry = IQuorumRegistry(address(_registry()));

        require(
            // check if call is coming from the 'Registry'
            msg.sender == address(registry)
            // otherwise, check if operator recently de-registered
            // specifically, check if the operator de-registered within (UPDATE_PERIOD + REVEAL_PERIOD) of the current time
            || (
                (operator == msg.sender)
                    && (block.timestamp <= registry.getOperatorDeregisterTime(operator) + UPDATE_PERIOD + REVEAL_PERIOD)
            ),
            "EphemeralKeyRegistry.postLastEphemeralKeyPreImage: onlyRegistry OR must have recently de-registered"
        );

        // retrieve the most recent EK entry for the operator
        uint256 historyLength = _getEKHistoryLength(operator);
        EKEntry memory existingEKEntry = EKHistory[operator][historyLength - 1];

        // check that the preimage matches with the hash
        require(
            existingEKEntry.keyHash == keccak256(abi.encode(prevEK)),
            "EphemeralKeyRegistry.postLastEphemeralKeyPreImage: Ephemeral key does not match previous ephemeral key commitment"
        );

        uint32 currentTaskNumber = repository.serviceManager().taskNumber();

        // update the EK entry
        existingEKEntry.ephemeralKey = prevEK;
        existingEKEntry.endTaskNumber = currentTaskNumber - 1;
        EKHistory[operator][historyLength - 1] = existingEKEntry;
    }

    /**
     * @notice Used by the operator to update their ephemeral key hash and post their previous ephemeral key after the expiry of its usage.
     * Revealing of current ephemeral key and describing the hash of the new ephemeral key done together.
     * @param prevEK is the previous ephemeral key, checked against the `msg.sender`'s existing keyHash.
     * @param newEKHash is the hash of the new ephemeral key.
     * @dev The function must be called within the `REVEAL_PERIOD` which is the time window within which an operator can reveal their EK preimage
     * without the risk of being frontrun and slashed (outside of the reveal period, if an attacker sees the preimage in the mempool, they can frontrun
     * that reveal transaction and cause slashing of the operator).  
     */
    function revealAndUpdateEphemeralKey(bytes32 prevEK, bytes32 newEKHash) external {
        // retrieve the most recent EK entry for the operator
        uint256 historyLength = _getEKHistoryLength(msg.sender);
        EKEntry memory existingEKEntry = EKHistory[msg.sender][historyLength - 1];

        // verify that the operator is active
        IQuorumRegistry registry = IQuorumRegistry(address(_registry()));
        require(
            registry.isActiveOperator(msg.sender),
            "EphemeralKeyRegistry.updateEphemeralKeyPreImage: operator is not active"
        );

        require(
            existingEKEntry.keyHash == keccak256(abi.encode(prevEK)),
            "EphemeralKeyRegistry.updateEphemeralKeyPreImage: Ephemeral key does not match previous ephemeral key commitment"
        );

        // checking the validity period of the ephemeral key update
        require(
            block.timestamp >= existingEKEntry.timestamp + UPDATE_PERIOD,
            "EphemeralKeyRegistry.updateEphemeralKeyPreImage: key update cannot be completed too early"
        );
        require(
            block.timestamp <= existingEKEntry.timestamp + UPDATE_PERIOD + REVEAL_PERIOD,
            "EphemeralKeyRegistry.updateEphemeralKeyPreImage: key update cannot be completed as update window has expired"
        );

        uint32 currentTaskNumber = repository.serviceManager().taskNumber();

        // updating the previous EK entry
        existingEKEntry.ephemeralKey = prevEK;
        existingEKEntry.endTaskNumber = currentTaskNumber - 1;
        EKHistory[msg.sender][historyLength - 1] = existingEKEntry;

        // new EK entry
        EKEntry memory newEKEntry;
        newEKEntry.keyHash = newEKHash;
        newEKEntry.timestamp = uint192(block.timestamp);
        newEKEntry.startTaskNumber = currentTaskNumber;
        EKHistory[msg.sender].push(newEKEntry);
    }

//TODO: `getLatestEphemeralKey` seems to be a better implementation than `getCurrEphemeralKey`.  Perhaps remove the latter?

    /// @notice retrieve the operator's current ephemeral key hash
    function getCurrEphemeralKeyHash(address operator) external view returns (bytes32) {
        uint256 historyLength = _getEKHistoryLength(operator);
        return EKHistory[operator][historyLength - 1].keyHash;
    }

    /// @notice retrieve the operator's current ephemeral key itself
    function getLatestEphemeralKey(address operator) external view returns (bytes32) {
        uint256 historyLength = _getEKHistoryLength(operator);
        if (EKHistory[operator][historyLength - 1].ephemeralKey != bytes32(0)) {
            return EKHistory[operator][historyLength - 1].ephemeralKey;
            // recent EKEntry is still within UPDATE_PERIOD
        } else {
            if (historyLength == 1) {
                revert("EphemeralKeyRegistry.getLatestEphemeralKey: no ephemeralKey posted yet");
            } else {
                return EKHistory[operator][historyLength - 2].ephemeralKey;
            }
        }
    }

    /**
     * @notice This function is used for getting the ephemeral key pertaining to a particular taskNumber, for an operator
     * @param operator The operator whose ephemeral key we are interested in.
     * @param taskNumber The taskNumber for which we want to retrieve the operator's ephemeral key
     */
    function getEphemeralKeyForTaskNumber(address operator, uint32 taskNumber) external view returns (bytes32) {
        uint256 historyLength = _getEKHistoryLength(operator);
        unchecked {
            historyLength -= 1;
        }
        EKEntry memory existingEKEntry = EKHistory[operator][historyLength];

        if (existingEKEntry.startTaskNumber >= taskNumber) {
            revert(
                "EphemeralKeyRegistry.getEphemeralKeyForTaskNumber: taskNumber corresponds to latest EK which is still unrevealed"
            );
        } else {
            for (; historyLength > 0; --historyLength) {
                if (
                    (taskNumber >= EKHistory[msg.sender][historyLength].startTaskNumber)
                        && (taskNumber <= EKHistory[msg.sender][historyLength].endTaskNumber)
                ) {
                    return EKHistory[msg.sender][historyLength].ephemeralKey;
                }
            }
        }
        revert("EphemeralKeyRegistry.getEphemeralKeyForTaskNumber: did not find EK");
    }

    /**
     * @notice Used for proving that an operator hasn't reveal their ephemeral key within the reveal window, triggering slashing of the operator.
     * @param operator The operator with a stale ephemeral key
     */
    function proveStaleUnrevealedEphemeralKey(address operator) external {
        uint256 historyLength = _getEKHistoryLength(operator);
        EKEntry memory existingEKEntry = EKHistory[operator][historyLength - 1];

        // check that the ephemeral key is not yet revealed
        require(
            existingEKEntry.ephemeralKey == bytes32(0),
            "EphemeralKeyRegistry.proveStaleEphemeralKey: ephemeralKey is already revealed"
        );

        // check that the ephemeral key is actually stale
        require(
            block.timestamp > existingEKEntry.timestamp + UPDATE_PERIOD + REVEAL_PERIOD,
            "EphemeralKeyRegistry.proveStaleEphemeralKey: ephemeralKey is not stale"
        );

        //trigger slashing of operator who hasn't updated their EK
        IServiceManager serviceManager = repository.serviceManager();
        serviceManager.freezeOperator(operator);
    }

    /**
     * @notice Used for proving that an operator leaked an ephemeral key that was not supposed to be revealed, triggering slashing of the operator.
     * @param operator The operator who leaked their ephemeral key.
     * @param leakedEphemeralKey The ephemeral key for the operator, which they were not supposed to reveal.
     */
    function proveLeakedEphemeralKey(address operator, bytes32 leakedEphemeralKey) external {
        uint256 historyLength = _getEKHistoryLength(operator);
        EKEntry memory existingEKEntry = EKHistory[operator][historyLength - 1];

        // First we check if we are still within the `UPDATE_PERIOD`, meaning that the operator still shouldn't have revealed the ephemeral key.
        if (block.timestamp < existingEKEntry.timestamp + UPDATE_PERIOD) {
            /**
             * Verify that the hash of the provided "leaked" preimage of the ephemeral key matches the stored hash. 
             * If it is a match, we call on the `serviceManager` contract to slash the operator.
             */
            if (existingEKEntry.keyHash == keccak256(abi.encode(leakedEphemeralKey))) {
                //trigger slashing function of the operator
                IServiceManager serviceManager = repository.serviceManager();
                serviceManager.freezeOperator(operator);
            }
        }
    }

    /// @notice Returns the UTC timestamp at which the operator last renewed their ephemeral key
    function getLastEKPostTimestamp(address operator) external view returns (uint192) {
        uint256 historyLength = _getEKHistoryLength(operator);
        EKEntry memory existingEKEntry = EKHistory[operator][historyLength - 1];
        return existingEKEntry.timestamp;
    }

    function _getEKHistoryLength(address operator) internal view returns (uint256) {
        uint256 historyLength = EKHistory[operator].length;
        if (historyLength == 0) {
            revert("EphemeralKeyRegistry._getEKHistoryLength: historyLength == 0");
        }
        return historyLength;
    }
}

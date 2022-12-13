// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./IRegistry.sol";

/**
 * @title Interface for a `Registry`-type contract that uses either 1 or 2 quorums.
 * @author Layr Labs, Inc.
 * @notice This contract does not currently support n-quorums where n >= 3.
 * Note in particular the presence of only `firstQuorumStake` and `secondQuorumStake` in the `OperatorStake` struct.
 */
interface IQuorumRegistry is IRegistry {
    // DATA STRUCTURES
    enum Status
    {
        // default is inactive
        INACTIVE,
        ACTIVE
    }

    /**
     * @notice  Data structure for storing info on operators to be used for:
     * - sending data by the sequencer
     * - payment and associated challenges
     */
    struct Operator {
        // hash of pubkey of the operator
        bytes32 pubkeyHash;
        // id is always unique
        uint32 id;
        // corresponds to position in operatorList
        uint32 index;
        // start taskNumber from which the  operator has been registered
        uint32 fromTaskNumber;
        // start block from which the  operator has been registered
        uint32 fromBlockNumber;
        // UTC time until which this operator is supposed to serve its obligations to this middleware
        // set only when committing to deregistration
        uint32 serveUntil;
        // UTC time at which the operator deregistered. If set to zero then the operator has not deregistered.
        uint32 deregisterTime;
        // indicates whether the operator is actively registered for serving the middleware or not
        Status status;
    }

    // struct used to give definitive ordering to operators at each blockNumber
    struct OperatorIndex {
        // blockNumber number at which operator index changed
        // note that the operator's index is different *for this block number*, i.e. the *new* index is *inclusive* of this value
        uint32 toBlockNumber;
        // index of the operator in array of operators, or the total number of operators if in the 'totalOperatorsHistory'
        uint32 index;
    }

    /// @notice struct used to store the stakes of an individual operator or the sum of all operators' stakes, for storage
    struct OperatorStake {
        // the block number at which the stake amounts were updated and stored
        uint32 updateBlockNumber;
        // the block number at which the *next update* occurred.
        /// @notice This entry has the value **0** until another update takes place.
        uint32 nextUpdateBlockNumber;
        // stake weight for the first quorum
        uint96 firstQuorumStake;
        // stake weight for the second quorum. Will always be zero in the event that only one quorum is used
        uint96 secondQuorumStake;
    }

    function getLengthOfTotalStakeHistory() external view returns (uint256);

    /**
     * @notice Returns the `index`-th entry in the dynamic array of total stake, `totalStakeHistory`.
     * @dev Function will revert in the event that `index` is out-of-bounds.
     */
    function getTotalStakeFromIndex(uint256 index) external view returns (OperatorStake memory);

    /// @notice Returns the unique ID of the specified `operator`.
    function getOperatorId(address operator) external returns (uint32);

    /// @notice Returns the stored pubkeyHash for the specified `operator`.
    function getOperatorPubkeyHash(address operator) external view returns (bytes32);

    /// @notice Returns task number from when `operator` has been registered.
    function getFromTaskNumberForOperator(address operator) external view returns (uint32);

    /// @notice Returns block number from when `operator` has been registered.
    function getFromBlockNumberForOperator(address operator) external view returns (uint32);

    /**
     * @notice Returns the stake weight corresponding to `pubkeyHash`, at the
     * `index`-th entry in the `pubkeyHashToStakeHistory[pubkeyHash]` array.
     * @param pubkeyHash Hash of the public key of the operator of interest.
     * @param index Array index for lookup, within the dynamic array `pubkeyHashToStakeHistory[pubkeyHash]`.
     * @dev Function will revert if `index` is out-of-bounds.
     */
    function getStakeFromPubkeyHashAndIndex(bytes32 pubkeyHash, uint256 index)
        external
        view
        returns (OperatorStake memory);

    /**
     * @notice Looks up the `operator`'s index in the dynamic array `operatorList` at the specified `blockNumber`.
     * @param index Used to specify the entry within the dynamic array `pubkeyHashToIndexHistory[pubkeyHash]` to 
     * read data from, where `pubkeyHash` is looked up from `operator`'s registration info
     * @param blockNumber Is the desired block number at which we wish to query the operator's position in the `operatorList` array
     * @dev Function will revert in the event that the specified `index` input does not identify the appropriate entry in the
     * array `pubkeyHashToIndexHistory[pubkeyHash]` to pull the info from.
    */
    function getOperatorIndex(address operator, uint32 blockNumber, uint32 index) external view returns (uint32);

    /**
     * @notice Looks up the number of total operators at the specified `blockNumber`.
     * @param index Input used to specify the entry within the dynamic array `totalOperatorsHistory` to read data from.
     * @dev This function will revert if the provided `index` is out of bounds.
    */
    function getTotalOperators(uint32 blockNumber, uint32 index) external view returns (uint32);

    /// @notice Returns the current number of operators of this service.
    function numOperators() external view returns (uint32);

    /**
     * @notice Returns the time at which the `operator` deregistered.
     * @dev Function will return **0** in the event that the operator is actively registered.
     */
    function getOperatorDeregisterTime(address operator) external view returns (uint256);

    /**
     * @notice Returns the most recent stake weights for the `operator`
     * @dev Function returns weights of **0** in the event that the operator has no stake history
     */
    function operatorStakes(address operator) external view returns (uint96, uint96);

    /// @notice Returns the stake amounts from the latest entry in `totalStakeHistory`.
    function totalStake() external view returns (uint96, uint96);
}

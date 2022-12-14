// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./RegistryBase.sol";
import "../interfaces/IBLSPublicKeyCompendium.sol";
import "../interfaces/IBLSRegistry.sol";
import "forge-std/Test.sol";

/**
 * @title A Registry-type contract using aggregate BLS signatures.
 * @author Layr Labs, Inc.
 * @notice This contract is used for
 * - registering new operators
 * - committing to and finalizing de-registration as an operator
 * - updating the stakes of the operator
 */
contract BLSRegistry is RegistryBase, IBLSRegistry {
    using BytesLib for bytes;

    // Hash of the zero public key
    bytes32 internal constant ZERO_PK_HASH = hex"012893657d8eb2efad4de0a91bcd0e39ad9837745dec3ea923737ea803fc8e3d";

    /// @notice contract used for looking up operators' BLS public keys
    IBLSPublicKeyCompendium public immutable pubkeyCompendium;

    /**
     * @notice list of keccak256(apk_x0, apk_x1, apk_y0, apk_y1) of operators, and the block numbers at which the aggregate
     * pubkeys were updated. This occurs whenever a new operator registers or deregisters.
     */
    ApkUpdate[] internal _apkUpdates;

    /**
     * @notice used for storing current aggregate public key
     * @dev Initialized value of APK is the generator of G2 group. It is necessary in order to do
     * addition in Jacobian coordinate system.
     */
    uint256[4] public apk;

    // EVENTS
    /**
     * @notice Emitted upon the registration of a new operator for the middleware
     * @param operator Address of the new operator
     * @param pkHash The keccak256 hash of the operator's public key
     * @param pk The operator's public key itself
     * @param apkHashIndex The index of the latest (i.e. the new) APK update
     * @param apkHash The keccak256 hash of the new Aggregate Public Key
     */
    event Registration(
        address indexed operator,
        bytes32 pkHash,
        uint256[4] pk,
        uint32 apkHashIndex,
        bytes32 apkHash,
        string socket
    );

    constructor(
        Repository _repository,
        IEigenLayrDelegation _delegation,
        IInvestmentManager _investmentManager,
        uint32 _unbondingPeriod,
        uint8 _NUMBER_OF_QUORUMS,
        uint256[] memory _quorumBips,
        StrategyAndWeightingMultiplier[] memory _firstQuorumStrategiesConsideredAndMultipliers,
        StrategyAndWeightingMultiplier[] memory _secondQuorumStrategiesConsideredAndMultipliers,
        IBLSPublicKeyCompendium _pubkeyCompendium
    )
        RegistryBase(
            _repository,
            _delegation,
            _investmentManager,
            _unbondingPeriod,
            _NUMBER_OF_QUORUMS,
            _quorumBips,
            _firstQuorumStrategiesConsideredAndMultipliers,
            _secondQuorumStrategiesConsideredAndMultipliers
        )
    {
        /**
         * @dev Initialized value of APK is the generator of G2 group. It is necessary in order to do
         * addition in Jacobian coordinate system.
         */
        uint256[4] memory initApk = [BLS.G2x0, BLS.G2x1, BLS.G2y0, BLS.G2y1];
        _processApkUpdate(initApk);
        //set compendium
        pubkeyCompendium = _pubkeyCompendium;
    }

    /**
     * @notice called for registering as an operator
     * @param operatorType specifies whether the operator want to register as staker for one or both quorums
     * @param pkBytes is the abi encoded bn254 G2 public key of the operator
     * @param socket is the socket address of the operator
     */
    function registerOperator(uint8 operatorType, bytes calldata pkBytes, string calldata socket) external virtual {
        _registerOperator(msg.sender, operatorType, pkBytes, socket);
    }

    /**
     * @param operator is the node who is registering to be a operator
     * @param operatorType specifies whether the operator want to register as staker for one or both quorums
     * @param pkBytes is the serialized bn254 G2 public key of the operator
     * @param socket is the socket address of the operator
     */
    function _registerOperator(address operator, uint8 operatorType, bytes calldata pkBytes, string calldata socket)
        internal
    {
        // validate the registration of `operator` and find their `OperatorStake`
        OperatorStake memory _operatorStake = _registrationStakeEvaluation(operator, operatorType);

        /// @notice evaluate the new aggregated pubkey
        uint256[4] memory newApk;
        uint256[4] memory pk = _parseSerializedPubkey(pkBytes);

        // getting pubkey hash
        bytes32 pubkeyHash = BLS.hashPubkey(pk);

        // our addition algorithm doesn't work in this case, since it won't properly handle `x + x`, per @gpsanant
        require(
            pubkeyHash != _apkUpdates[_apkUpdates.length - 1].apkHash,
            "BLSRegistry._registerOperator: Apk and pubkey cannot be the same"
        );

        require(pubkeyHash != ZERO_PK_HASH, "BLSRegistry._registerOperator: Cannot register with 0x0 public key");

        require(
            pubkeyCompendium.pubkeyHashToOperator(pubkeyHash) == operator,
            "BLSRegistry._registerOperator: operator does not own pubkey"
        );

        require(
            pubkeyHashToStakeHistory[pubkeyHash].length == 0, "BLSRegistry._registerOperator: pubkey already registered"
        );

        {
            // add pubkey to aggregated pukkey in Jacobian coordinates
            uint256[6] memory newApkJac =
                BLS.addJac([pk[0], pk[1], pk[2], pk[3], 1, 0], [apk[0], apk[1], apk[2], apk[3], 1, 0]);

            // convert back to Affine coordinates
            (newApk[0], newApk[1], newApk[2], newApk[3]) = BLS.jacToAff(newApkJac);
        }

        // record the APK update and get the hash of the new APK
        bytes32 newApkHash = _processApkUpdate(newApk);

        // add the operator to the list of registrants and do accounting
        _addRegistrant(operator, pubkeyHash, _operatorStake);

        emit Registration(operator, pubkeyHash, pk, uint32(_apkUpdates.length - 1), newApkHash, socket);
    }

    /**
     * @notice Used by an operator to de-register itself from providing service to the middleware.
     * @param pubkeyToRemoveAff is the sender's pubkey in affine coordinates
     * @param index is the sender's location in the dynamic array `operatorList`
     */
    function deregisterOperator(uint256[4] memory pubkeyToRemoveAff, uint32 index) external virtual returns (bool) {
        _deregisterOperator(msg.sender, pubkeyToRemoveAff, index);
        return true;
    }

    /**
     * @notice Used to process de-registering an operator from providing service to the middleware.
     * @param operator The operator to be deregistered
     * @param pubkeyToRemoveAff is the sender's pubkey in affine coordinates
     * @param index is the sender's location in the dynamic array `operatorList`
     */
    function _deregisterOperator(address operator, uint256[4] memory pubkeyToRemoveAff, uint32 index) internal {
        // verify that the `operator` is an active operator and that they've provided the correct `index`
        _deregistrationCheck(operator, index);

        /// @dev Fetch operator's stored pubkeyHash
        bytes32 pubkeyHash = registry[operator].pubkeyHash;
        /// @dev Verify that the stored pubkeyHash matches the 'pubkeyToRemoveAff' input
        require(
            pubkeyHash == BLS.hashPubkey(pubkeyToRemoveAff),
            "BLSRegistry._deregisterOperator: pubkey input does not match stored pubkeyHash"
        );

        // Perform necessary updates for removing operator, including updating operator list and index histories
        _removeOperator(pubkeyHash, index);

        // get existing aggregate public key
        uint256[4] memory pk = apk;
        // remove signer's pubkey from aggregate public key
        (pk[0], pk[1], pk[2], pk[3]) = BLS.removePubkeyFromAggregate(pubkeyToRemoveAff, pk);
        // update the aggregate public key of all registered operators and record this update in history
        _processApkUpdate(pk);
    }

    /**
     * @notice Used for updating information on deposits of nodes.
     * @param operators are the nodes whose deposit information is getting updated
     */
    function updateStakes(address[] calldata operators) external {
        // copy total stake to memory
        OperatorStake memory _totalStake = totalStakeHistory[totalStakeHistory.length - 1];

        // placeholders reused inside of loop
        OperatorStake memory currentStakes;
        bytes32 pubkeyHash;
        uint256 operatorsLength = operators.length;
        // iterating over all the tuples that are to be updated
        for (uint256 i = 0; i < operatorsLength;) {
            // get operator's pubkeyHash
            pubkeyHash = registry[operators[i]].pubkeyHash;
            // fetch operator's existing stakes
            currentStakes = pubkeyHashToStakeHistory[pubkeyHash][pubkeyHashToStakeHistory[pubkeyHash].length - 1];
            // decrease _totalStake by operator's existing stakes
            _totalStake.firstQuorumStake -= currentStakes.firstQuorumStake;
            _totalStake.secondQuorumStake -= currentStakes.secondQuorumStake;

            // update the stake for the i-th operator
            currentStakes = _updateOperatorStake(operators[i], pubkeyHash, currentStakes);

            // increase _totalStake by operator's updated stakes
            _totalStake.firstQuorumStake += currentStakes.firstQuorumStake;
            _totalStake.secondQuorumStake += currentStakes.secondQuorumStake;

            unchecked {
                ++i;
            }
        }

        // update storage of total stake
        _recordTotalStakeUpdate(_totalStake);
    }

    //TODO: The subgraph doesnt like uint256[4][] argument here. Figure this out laters
    // // TODO: de-dupe code copied from `updateStakes`, if reasonably possible
    // /**
    //  * @notice Used for removing operators that no longer meet the minimum requirements
    //  * @param operators are the nodes who will potentially be booted
    //  */
    // function bootOperators(
    //     address[] calldata operators,
    //     uint256[4][] memory pubkeysToRemoveAff,
    //     uint32[] memory indices
    // )
    //     external
    // {
    //     // copy total stake to memory
    //     OperatorStake memory _totalStake = totalStakeHistory[totalStakeHistory.length - 1];

    //     // placeholders reused inside of loop
    //     OperatorStake memory currentStakes;
    //     bytes32 pubkeyHash;
    //     uint256 operatorsLength = operators.length;
    //     // iterating over all the tuples that are to be updated
    //     for (uint256 i = 0; i < operatorsLength;) {
    //         // get operator's pubkeyHash
    //         pubkeyHash = registry[operators[i]].pubkeyHash;
    //         // fetch operator's existing stakes
    //         currentStakes = pubkeyHashToStakeHistory[pubkeyHash][pubkeyHashToStakeHistory[pubkeyHash].length - 1];
    //         // decrease _totalStake by operator's existing stakes
    //         _totalStake.firstQuorumStake -= currentStakes.firstQuorumStake;
    //         _totalStake.secondQuorumStake -= currentStakes.secondQuorumStake;

    //         // update the stake for the i-th operator
    //         currentStakes = _updateOperatorStake(operators[i], pubkeyHash, currentStakes);

    //         // remove the operator from the list of operators if they do *not* meet the minimum requirements
    //         if (
    //             (currentStakes.firstQuorumStake < minimumStakeFirstQuorum)
    //                 && (currentStakes.secondQuorumStake < minimumStakeSecondQuorum)
    //         ) {
    //             // TODO: optimize better if possible? right now this pushes an APK update for each operator removed.
    //             _deregisterOperator(operators[i], pubkeysToRemoveAff[i], indices[i]);
    //         }
    //         // in the case that the operator *does indeed* meet the minimum requirements
    //         else {
    //             // increase _totalStake by operator's updated stakes
    //             _totalStake.firstQuorumStake += currentStakes.firstQuorumStake;
    //             _totalStake.secondQuorumStake += currentStakes.secondQuorumStake;
    //         }

    //         unchecked {
    //             ++i;
    //         }
    //     }

    //     // update storage of total stake
    //     _recordTotalStakeUpdate(_totalStake);
    // }

    /**
     * @notice Updates the stored APK to `newApk`, calculates its hash, and pushes new entries to the `_apkUpdates` array
     * @param newApk The updated APK. This will be the `apk` after this function runs!
     */
    function _processApkUpdate(uint256[4] memory newApk) internal returns (bytes32) {
        // update stored aggregate public key
        // slither-disable-next-line costly-loop
        apk = newApk;

        // find the hash of aggregate pubkey
        bytes32 newApkHash = keccak256(abi.encodePacked(newApk[0], newApk[1], newApk[2], newApk[3]));

        // store the apk hash and the current block number in which the aggregated pubkey is being updated
        _apkUpdates.push(ApkUpdate({
            apkHash: newApkHash,
            blockNumber: uint32(block.number)
        }));

        return newApkHash;
    }

    // pkBytes = abi.encodePacked(pk.X.A1, pk.X.A0, pk.Y.A1, pk.Y.A0)
    function _parseSerializedPubkey(bytes calldata pkBytes) internal pure returns (uint256[4] memory) {
        uint256[4] memory pk;
        assembly {
            mstore(add(pk, 32), calldataload(pkBytes.offset))
            mstore(pk, calldataload(add(pkBytes.offset, 32)))
            mstore(add(pk, 96), calldataload(add(pkBytes.offset, 64)))
            mstore(add(pk, 64), calldataload(add(pkBytes.offset, 96)))
        }
        return pk;
    }

    /**
     * @notice get hash of a historical aggregated public key corresponding to a given index;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getCorrectApkHash(uint256 index, uint32 blockNumber) external view returns (bytes32) {
        // check that the `index`-th APK update occurred at or before `blockNumber`
        require(blockNumber >= _apkUpdates[index].blockNumber, "BLSRegistry.getCorrectApkHash: index too recent");

        // if not last update
        if (index != _apkUpdates.length - 1) {
            // check that there was not *another APK update* that occurred at or before `blockNumber`
            require(blockNumber < _apkUpdates[index + 1].blockNumber, "BLSRegistry.getCorrectApkHash: Not latest valid apk update");
        }

        return _apkUpdates[index].apkHash;
    }

    /// @notice returns the total number of APK updates that have ever occurred (including one for initializing the pubkey as the generator)
    function getApkUpdatesLength() external view returns (uint256) {
        return _apkUpdates.length;
    }

    /// @notice returns the `ApkUpdate` struct at `index` in the list of APK updates
    function apkUpdates(uint256 index) external view returns (ApkUpdate memory) {
        return _apkUpdates[index];
    }

    /// @notice returns the APK hash that resulted from the `index`th APK update
    function apkHashes(uint256 index) external view returns (bytes32) {
        return _apkUpdates[index].apkHash;
    }

    /// @notice returns the block number at which the `index`th APK update occurred
    function apkUpdateBlockNumbers(uint256 index) external view returns (uint32) {
        return _apkUpdates[index].blockNumber;
    }
}

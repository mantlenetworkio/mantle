// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../interfaces/IServiceManager.sol";
import "../interfaces/IRegistry.sol";
import "../interfaces/IEphemeralKeyRegistry.sol";
import "../interfaces/IBLSPublicKeyCompendium.sol";
import "../libraries/BytesLib.sol";
import "./BLSRegistry.sol";

// import "forge-std/Test.sol";

/**
 * @title Adds Proof of Custody functionality to the `BLSRegistry` contract.
 * @author Layr Labs, Inc.
 * @notice See the Dankrad's excellent article for an intro to Proofs of Custody:
 * https://dankradfeist.de/ethereum/2021/09/30/proofs-of-custody.html.
 * This contract relies on an `EphemeralKeyRegistry` to store operator's ephemeral keys.
 */
contract BLSRegistryWithBomb is BLSRegistry {
    using BytesLib for bytes;

    // TODO: either make this immutable *or* add a method to change it
    IEphemeralKeyRegistry public ephemeralKeyRegistry;

    constructor(
        Repository _repository,
        IEigenLayrDelegation _delegation,
        IInvestmentManager _investmentManager,
        IEphemeralKeyRegistry _ephemeralKeyRegistry,
        uint32 _unbondingPeriod,
        uint8 _NUMBER_OF_QUORUMS,
        uint256[] memory _quorumBips,
        StrategyAndWeightingMultiplier[] memory _firstQuorumStrategiesConsideredAndMultipliers,
        StrategyAndWeightingMultiplier[] memory _secondQuorumStrategiesConsideredAndMultipliers,
        IBLSPublicKeyCompendium _pubkeyCompendium
    )
        BLSRegistry(
            _repository,
            _delegation,
            _investmentManager,
            _unbondingPeriod,
            _NUMBER_OF_QUORUMS,
            _quorumBips,
            _firstQuorumStrategiesConsideredAndMultipliers,
            _secondQuorumStrategiesConsideredAndMultipliers,
            _pubkeyCompendium
        )
    {
        ephemeralKeyRegistry = _ephemeralKeyRegistry;
    }

    /**
     * @notice Used by an operator to de-register itself from providing service to the middleware.
     * For detailed comments, see the `deregisterOperator` function in BLSRegistry.sol.
     * Same as `BLSRegistry.deregisterOperator` except adds an external call to `ephemeralKeyRegistry.postLastEphemeralKeyPreImage(msg.sender, finalEphemeralKey)`,
     * passing along the additional argument `finalEphemeralKey`.
     */
    function deregisterOperator(uint256[4] memory pubkeyToRemoveAff, uint32 index, bytes32 finalEphemeralKey)
        external
        returns (bool)
    {
        _deregisterOperator(msg.sender, pubkeyToRemoveAff, index);

        //post last ephemeral key reveal on chain
        ephemeralKeyRegistry.postLastEphemeralKeyPreImage(msg.sender, finalEphemeralKey);

        return true;
    }

    /**
     * @notice called for registering as an operator.
     * For detailed comments, see the `registerOperator` function in BLSRegistry.sol.
     * same as `BLSRegistry.registerOperator` except adds an external call to `ephemeralKeyRegistry.postFirstEphemeralKeyHash(msg.sender, ephemeralKeyHash)`,
     * passing along the additional argument `ephemeralKeyHash`.
     */
    function registerOperator(
        uint8 operatorType,
        bytes32 ephemeralKeyHash,
        bytes calldata pkBytes,
        string calldata socket
    ) external {
        _registerOperator(msg.sender, operatorType, pkBytes, socket);

        //add ephemeral key to ephemeral key registry
        ephemeralKeyRegistry.postFirstEphemeralKeyHash(msg.sender, ephemeralKeyHash);
    }

    // the following function overrides the base function of BLSRegistry -- we want operators to provide additional arguments, so these versions (without those args) revert
    function registerOperator(uint8, bytes calldata, string calldata) external pure override {
        revert("BLSRegistryWithBomb.registerOperator: must register with ephemeral key");
    }

    // the following function overrides the base function of BLSRegistry -- we want operators to provide additional arguments, so these versions (without those args) revert
    function deregisterOperator(uint256[4] memory, uint32) external pure override returns (bool) {
        revert("BLSRegistryWithBomb.deregisterOperator: must deregister with ephemeral key");
    }
}

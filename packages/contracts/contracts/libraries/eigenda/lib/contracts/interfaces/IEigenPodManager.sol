// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./IInvestmentManager.sol";
import "./IEigenPod.sol";

/**
 * @title Interface for factory that creates and manages solo staking pods that have their withdrawal credentials pointed to EigenLayer.
 * @author Layr Labs, Inc.
 */

interface IEigenPodManager {
    //This struct helps manage the info about a certain pod owner's pod
    struct EigenPodInfo {
        uint128 balance; //total balance of all validators in the pod
        uint128 depositedBalance; //amount of balance deposited into EigenLayer
    }

    /**
     * @notice Creates an EigenPod for the sender.
     * @dev Function will revert if the `msg.sender` already has an EigenPod.
     */
    function createPod() external;

    /**
     * @notice Stakes for a new beacon chain validator on the sender's EigenPod. 
     * Also creates an EigenPod for the sender if they don't have one already.
     * @param pubkey The 48 bytes public key of the beacon chain validator.
     * @param signature The validator's signature of the deposit data.
     * @param depositDataRoot The root/hash of the deposit data for the validator's deposit.
     */
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable;

    /**
     * @notice Updates the beacon chain balance of the EigenPod, freezing the owner if they have overcommitted beacon chain ETH to EigenLayer.
     * @param podOwner The owner of the pod to udpate the balance of.
     * @param balanceToRemove The balance to remove before increasing, used when updating a validators balance.
     * @param balanceToAdd The balance to add after decreasing, used when updating a validators balance.
     * @dev Callable only by the `podOwner`'s EigenPod.
     */
    function updateBeaconChainBalance(address podOwner, uint64 balanceToRemove, uint64 balanceToAdd) external;

    /**
     * @notice Stakes beacon chain ETH into EigenLayer by adding BeaconChainETH shares to InvestmentManager.
     * @param podOwner The owner of the pod whose balance must be restaked.
     * @param amount The amount of beacon chain ETH to restake.
     * @dev Callable only by the `podOwner`'s EigenPod.
     */
    function depositBeaconChainETH(address podOwner, uint64 amount) external;

    /**
     * @notice Withdraws ETH that has been withdrawn from the beacon chain from the EigenPod.
     * @param podOwner The owner of the pod whose balance must be withdrawn.
     * @param recipient The recipient of withdrawn ETH.
     * @param amount The amount of ETH to withdraw.
     * @dev Callable only by the InvestmentManager contract.
     */
    function withdrawBeaconChainETH(address podOwner, address recipient, uint256 amount) external;

    /**
     * @notice Updates the oracle contract that provides the beacon chain state root
     * @param newBeaconChainOracle is the new oracle contract being pointed to
     * @dev Callable only by the owner of the InvestmentManager (i.e. governance).
     */
    function updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) external;

    /// @notice Returns the address of the `podOwner`'s EigenPod (whether it is deployed yet or not).
    function getPod(address podOwner) external view returns(IEigenPod);

    /// @notice returns the current EigenPodInfo for the `podOwner`'s EigenPod.
    function getPodInfo(address podOwner) external view returns(EigenPodInfo memory);

    /// @notice Oracle contract that provides updates to the beacon chain's state
    function beaconChainOracle() external view returns(IBeaconChainOracle);    

    /// @notice Returns the latest beacon chain state root posted to the beaconChainOracle.
    function getBeaconChainStateRoot() external view returns(bytes32);

    /// @notice EigenLayer's InvestmentManager contract
    function investmentManager() external view returns(IInvestmentManager);

    function getBalance(address podOwner) external view returns (uint128);

    function getDepositedBalance(address podOwner) external view returns (uint128);

    function hasPod(address podOwner) external view returns (bool);
}

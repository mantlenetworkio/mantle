// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../libraries/BeaconChainProofs.sol";
import "./IEigenPodManager.sol";
import "./IBeaconChainOracle.sol";

/**
 * @title Interface for solo staking pods that have their withdrawal credentials pointed to EigenLayer.
 * @author Layr Labs, Inc.
 */

interface IEigenPod {
    enum VALIDATOR_STATUS {
        INACTIVE, // doesnt exist
        ACTIVE, // staked on ethpos and withdrawal credentials are pointed to the EigenPod
        OVERCOMMITTED, // proven to be overcommitted to EigenLayer
        WITHDRAWN // withdrawn from the Beacon Chain
    }

    // this struct keeps track of PartialWithdrawalClaims
    struct PartialWithdrawalClaim {
        PARTIAL_WITHDRAWAL_CLAIM_STATUS status;
        // block at which the PartialWithdrawalClaim was created
        uint32 creationBlockNumber;
        // last block (inclusive) in which the PartialWithdrawalClaim can be fraudproofed
        uint32 fraudproofPeriodEndBlockNumber;
        // amount of ETH -- in Gwei -- to be withdrawn until completion of this claim
        uint64 partialWithdrawalAmountGwei;
    }

    enum PARTIAL_WITHDRAWAL_CLAIM_STATUS {
        REDEEMED,
        PENDING,
        FAILED
    }

    /// @notice The length, in blocks, if the fraud proof period following a claim on the amount of partial withdrawals in an EigenPod
    function PARTIAL_WITHDRAWAL_FRAUD_PROOF_PERIOD_BLOCKS() external view returns(uint32);

    /// @notice The amount of eth, in gwei, that is restaked per validator
    function REQUIRED_BALANCE_GWEI() external view returns(uint64);

    /// @notice The amount of eth, in wei, that is added to the penalty balance of the pod in case a validator's beacon chain balance ever falls
    ///         below REQUIRED_BALANCE_GWEI
    /// @dev currently this is set to REQUIRED_BALANCE_GWEI
    function OVERCOMMITMENT_PENALTY_AMOUNT_GWEI() external view returns(uint64);

    /// @notice The amount of eth, in wei, that is restaked per validator
    function REQUIRED_BALANCE_WEI() external view returns(uint256);

    /// @notice The amount of eth, in gwei, that can be part of a full withdrawal at the minimum
    function MIN_FULL_WITHDRAWAL_AMOUNT_GWEI() external view returns(uint64);

    /// @notice this is a mapping of validator indices to a Validator struct containing pertinent info about the validator
    function validatorStatus(uint40 validatorIndex) external view returns(VALIDATOR_STATUS);

    /// @return claim is the partial withdrawal claim at the provided index
    function getPartialWithdrawalClaim(uint256 index) external view returns(PartialWithdrawalClaim memory);
        
    /// @return length : the number of partial withdrawal claims ever made for this EigenPod
    function getPartialWithdrawalClaimsLength() external view returns(uint256);

    /// @notice the amount of execution layer ETH in this contract that is staked in EigenLayer (i.e. withdrawn from beaconchain but not EigenLayer), 
    function restakedExecutionLayerGwei() external view returns(uint64);

    /// @notice the excess balance from full withdrawals over RESTAKED_BALANCE_PER_VALIDATOR or partial withdrawals
    function instantlyWithdrawableBalanceGwei() external view returns(uint64);

    /// @notice the amount of penalties that have been paid from instantlyWithdrawableBalanceGwei or from partial withdrawals. These can be rolled
    ///         over from restakedExecutionLayerGwei into instantlyWithdrawableBalanceGwei when all existing penalties have been paid
    function rollableBalanceGwei() external view returns(uint64);

    /// @notice the total amount of gwei outstanding (i.e. to-be-paid) penalties due to over committing to EigenLayer on behalf of this pod
    function penaltiesDueToOvercommittingGwei() external view returns(uint64);

    /// @notice Used to initialize the pointers to contracts crucial to the pod's functionality, in beacon proxy construction from EigenPodManager
    function initialize(IEigenPodManager _eigenPodManager, address owner) external;

    /// @notice Called by EigenPodManager when the owner wants to create another ETH validator.
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable;

    /**
     * @notice Transfers `amountWei` in ether from this contract to the specified `recipient` address
     * @notice Called by EigenPodManager to withdrawBeaconChainETH that has been added to the EigenPod's balance due to a withdrawal from the beacon chain.
     * @dev Called during withdrawal or slashing.
     * @dev Note that this function is marked as non-reentrant to prevent the recipient calling back into it
     */
    function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) external;

    /// @notice The single EigenPodManager for EigenLayer
    function eigenPodManager() external view returns (IEigenPodManager);

    /// @notice The owner of this EigenPod
    function podOwner() external view returns (address);

    /**
     * @notice This function verifies that the withdrawal credentials of the podOwner are pointed to
     * this contract.  It verifies the provided proof of the ETH validator against the beacon chain state
     * root, marks the validator as 'active' in EigenLayer, and credits the restaked ETH in Eigenlayer.
     * @param proof is the bytes that prove the ETH validator's metadata against a beacon chain state root
     * @param validatorFields are the fields of the "Validator Container", refer to consensus specs 
     * for details: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
     */
    function verifyCorrectWithdrawalCredentials(
        uint40 validatorIndex,
        bytes calldata proof, 
        bytes32[] calldata validatorFields
    ) external;
    
    /**
     * @notice This function records an overcommitment of stake to EigenLayer on behalf of a certain ETH validator.
     *         If successful, the overcommitted balance is penalized (available for withdrawal whenever the pod's balance allows).
     *         The ETH validator's shares in the enshrined beaconChainETH strategy are also removed from the InvestmentManager and undelegated.
     * @param proof is the bytes that prove the ETH validator's metadata against a beacon state root
     * @param validatorFields are the fields of the "Validator Container", refer to consensus specs 
     * @param beaconChainETHStrategyIndex is the index of the beaconChainETHStrategy for the pod owner for the callback to 
     *                                    the InvestmentManger in case it must be removed from the list of the podOwners strategies
     * @dev For more details on the Beacon Chain spec, see: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
     */
    function verifyOvercommittedStake(
        uint40 validatorIndex,
        bytes calldata proof, 
        bytes32[] calldata validatorFields,
        uint256 beaconChainETHStrategyIndex
    ) external;

    /**
     * @notice This function records a full withdrawal on behalf of one of the Ethereum validators for this EigenPod
     * @param proof is the information needed to check the veracity of the block number and withdrawal being proven
     * @param blockNumberRoot is block number at which the withdrawal being proven is claimed to have happened
     * @param withdrawalFields are the fields of the withdrawal being proven
     * @param beaconChainETHStrategyIndex is the index of the beaconChainETHStrategy for the pod owner for the callback to 
     *                                    the EigenPodManager to the InvestmentManager in case it must be removed from the 
     *                                    podOwner's list of strategies
     */
    function verifyBeaconChainFullWithdrawal(
        BeaconChainProofs.WithdrawalAndBlockNumberProof calldata proof,
        bytes32 blockNumberRoot,
        bytes32[] calldata withdrawalFields,
        uint256 beaconChainETHStrategyIndex
    ) external;

    /**
     * @notice This function records a balance snapshot for the EigenPod. Its main functionality is to begin an optimistic
     *         claim process on the partial withdrawable balance for the EigenPod owner. The owner is claiming that they have 
     *         proven all full withdrawals until block.number, allowing their partial withdrawal balance to be easily calculated 
     *         via  
     *              address(this).balance / GWEI_TO_WEI = 
     *                  restakedExecutionLayerGwei + 
     *                  instantlyWithdrawableBalanceGwei + 
     *                  partialWithdrawalsGwei
     *         If any other full withdrawals are proven to have happened before block.number, the partial withdrawal is marked as failed
     * @param expireBlockNumber this is the block number before which the call to this function must be mined. To avoid race conditions with pending withdrawals,
     *                          if there are any pending full withrawals to this Eigenpod, this parameter should be set to the blockNumber at which the next full withdrawal
     *                          for a validator on this EigenPod is going to occur.
     * @dev The sender should be able to safely set the value of `expireBlockNumber` to type(uint32).max if there are no pending full withdrawals to this Eigenpod.
     */
    function recordPartialWithdrawalClaim(uint32 expireBlockNumber) external;

    /// @notice This function allows pod owners to redeem their partial withdrawals after the fraudproof period has elapsed
    function redeemLatestPartialWithdrawal(address recipient) external;

    /** 
     * @notice Withdraws instantlyWithdrawableBalanceGwei to the specified `recipient`
     * @dev Note that this function is marked as non-reentrant to prevent the recipient calling back into it
     */
    function withdrawInstantlyWithdrawableBalanceGwei(address recipient) external;

    /**
     * @notice Rebalances restakedExecutionLayerGwei in case penalties were previously paid from instantlyWithdrawableBalanceGwei or partial 
     *         withdrawal, so the EigenPod thinks podOwner has more restakedExecutionLayerGwei and staked balance than beaconChainETH on EigenLayer
     * @param amountGwei is the amount, in gwei, to roll over
     */
    function rollOverRollableBalance(uint64 amountGwei) external;

    /**
     * @notice Pays off existing penalties due to overcommitting to EigenLayer. Funds for paying penalties are deducted:
     *         1) first, from the execution layer ETH that is restaked in EigenLayer, because 
     *            it is the ETH that is actually supposed to be restaked
     *         2) second, from the instantlyWithdrawableBalanceGwei, to avoid allowing instant withdrawals
     *            from instantlyWithdrawableBalanceGwei, in case the balance of the contract is not enough 
     *            to cover the entire penalty
     */
    function payOffPenalties() external;
}
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "./interfaces/IDelegationManager.sol";
import "./interfaces/IDelegationShare.sol";
import "./interfaces/IDelegation.sol";
import "./interfaces/IDelegationSlasher.sol";

/**
 * @title Storage variables for the `InvestmentManager` contract.
 * @author Layr Labs, Inc.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract DelegationManagerStorage is IDelegationManager {
    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
    /// @notice The EIP-712 typehash for the deposit struct used by the contract
    bytes32 public constant DEPOSIT_TYPEHASH =
        keccak256("Deposit(address strategy,address token,uint256 amount,uint256 nonce,uint256 expiry)");
    /// @notice EIP-712 Domain separator
    bytes32 public DOMAIN_SEPARATOR;
    // staker => number of signed deposit nonce (used in depositIntoStrategyOnBehalfOf)
    mapping(address => uint256) public nonces;
    /**
     * @notice When a staker undelegates or an operator deregisters, their stake can still be slashed based on tasks/services created
     * within `REASONABLE_STAKES_UPDATE_PERIOD` of the present moment. In other words, this is the lag between undelegation/deregistration
     * and the staker's/operator's funds no longer being slashable due to misbehavior *on a new task*.
     */
    uint256 public constant REASONABLE_STAKES_UPDATE_PERIOD = 30 seconds;

    // fixed waiting period for withdrawals
    // TODO: set this to a proper interval for production
    uint32 public constant WITHDRAWAL_WAITING_PERIOD = 10 seconds;

    // maximum length of dynamic arrays in `investorStrats` mapping, for sanity's sake
    uint8 internal constant MAX_INVESTOR_DELEGATION_LENGTH = 32;

    // delegation system contracts
    IDelegation public immutable delegation;
    IDelegationSlasher public immutable delegationSlasher;

    // staker => IDelegationShare => number of shares which they currently hold
    mapping(address => mapping(IDelegationShare => uint256)) public investorDelegationShares;
    // staker => array of DelegationShare in which they have nonzero shares
    mapping(address => IDelegationShare[]) public investorDelegations;
    // hash of withdrawal inputs, aka 'withdrawalRoot' => timestamps & address related to the withdrawal
    mapping(bytes32 => WithdrawalStorage) public queuedWithdrawals;
    // staker => cumulative number of queued withdrawals they have ever initiated. only increments (doesn't decrement)
    mapping(address => uint256) public numWithdrawalsQueued;

    constructor(IDelegation _delegation, IDelegationSlasher _delegationSlasher) {
        delegation = _delegation;
        delegationSlasher = _delegationSlasher;
    }
}

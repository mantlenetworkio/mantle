// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "./interfaces/IDelegationCallback.sol";
import "./interfaces/IDelegation.sol";

/**
 * @title Base implementation of `IInvestmentStrategy` interface, designed to be inherited from by more complex strategies.
 * @notice Simple, basic, "do-nothing" InvestmentStrategy that holds a single underlying token and returns it on withdrawals.
 * Implements minimal versions of the IInvestmentStrategy functions, this contract is designed to be inherited by
 * more complex investment strategies, which can then override its functions as necessary.
 */
abstract contract DelegationCallbackBase is Initializable, PausableUpgradeable, IDelegationCallback {
    /// @notice DelegationManager contract
    IDelegation public delegation;

    /// @notice Simply checks that the `msg.sender` is the `DelegationManager`, which is an address stored immutably at construction.
    modifier onlyDelegation() {
        require(msg.sender == address(delegation), "DelegationShareBase.onlyDelegationManager");
        _;
    }

    function payForService(IERC20 token, uint256 amount) external payable {}

    function onDelegationWithdrawn(
        address delegator,
        IDelegationShare[] memory investorDelegationShares,
        uint256[] memory investorShares
    ) external {}

    function onDelegationReceived(
        address delegator,
        IDelegationShare[] memory investorDelegationShares,
        uint256[] memory investorShares
    ) external {}
}

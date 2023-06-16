// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title Minimal interface for an `IDelegationShares` contract.
 * @notice Custom `DelegationShares` implementations may expand extensively on this interface.
 */
interface IDelegationShare {
    /**
     * @notice Used to deposit tokens into this DelegationShares
     * @param token is the ERC20 token being deposited
     * @param amount is the amount of token being deposited
     * @dev This function is only callable by the delegationManager contract. It is invoked inside of the delegationManager's
     * `depositInto` function, and individual share balances are recorded in the delegationManager as well.
     * @return newShares is the number of new shares issued at the current exchange ratio.
     */
    function deposit(address depositor, IERC20 token, uint256 amount) external returns (uint256);

    /**
     * @notice Used to withdraw tokens from this DelegationLedger, to the `depositor`'s address
     * @param token is the ERC20 token being transferred out
     * @param amountShares is the amount of shares being withdrawn
     * @dev This function is only callable by the delegationManager contract. It is invoked inside of the delegationManager's
     * other functions, and individual share balances are recorded in the delegationManager as well.
     */
    function withdraw(address depositor, IERC20 token, uint256 amountShares) external;

    /**
     * @notice Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.
     * @notice In contrast to `sharesToUnderlyingView`, this function **may** make state modifications
     * @param amountShares is the amount of shares to calculate its conversion into the underlying token
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function sharesToUnderlying(uint256 amountShares) external returns (uint256);

    /**
     * @notice Used to convert an amount of underlying tokens to the equivalent amount of shares in this ledger.
     * @notice In contrast to `underlyingToSharesView`, this function **may** make state modifications
     * @param amountUnderlying is the amount of `underlyingToken` to calculate its conversion into ledger shares
     * @dev Implementation for these functions in particular may vary signifcantly for different ledgers
     */
    function underlyingToShares(uint256 amountUnderlying) external view returns (uint256);

    /**
     * @notice convenience function for fetching the current underlying value of all of the `user`'s shares in
     * this ledger. In contrast to `userUnderlyingView`, this function **may** make state modifications
     */
    function userUnderlying(address user) external returns (uint256);

     /**
     * @notice Used to convert a number of shares to the equivalent amount of underlying tokens for this ledger.
     * @notice In contrast to `sharesToUnderlying`, this function guarantees no state modifications
     * @param amountShares is the amount of shares to calculate its conversion into the underlying token
     * @dev Implementation for these functions in particular may vary signifcantly for different ledgers
     */
    function sharesToUnderlyingView(uint256 amountShares) external view returns (uint256);

    /**
     * @notice Used to convert an amount of underlying tokens to the equivalent amount of shares in this ledger.
     * @notice In contrast to `underlyingToShares`, this function guarantees no state modifications
     * @param amountUnderlying is the amount of `underlyingToken` to calculate its conversion into ledger shares
     * @dev Implementation for these functions in particular may vary signifcantly for different ledgers
     */
    function underlyingToSharesView(uint256 amountUnderlying) external view returns (uint256);

    /**
     * @notice convenience function for fetching the current underlying value of all of the `user`'s shares in
     * this ledger. In contrast to `userUnderlying`, this function guarantees no state modifications
     */
    function userUnderlyingView(address user) external view returns (uint256);

    /// @notice The underyling token for shares in this DelegationShares
    function underlyingToken() external view returns (IERC20);

    /// @notice The total number of extant shares in thie InvestmentStrategy
    function totalShares() external view returns (uint256);

    /// @notice Returns either a brief string explaining the strategy's goal & purpose, or a link to metadata that explains in more detail.
    function explanation() external view returns (string memory);
}

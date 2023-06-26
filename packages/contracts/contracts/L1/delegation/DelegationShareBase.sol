// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";

import "./interfaces/IDelegationManager.sol";

/**
 * @title Base implementation of `IDelegationShare` interface, designed to be inherited from by more complex strategies.
 * @author Layr Labs, Inc.
 * @notice Simple, basic, "do-nothing" DelegationShare that holds a single underlying token and returns it on withdrawals.
 * Implements minimal versions of the IDelegationShare functions, this contract is designed to be inherited by
 * more complex delegation contracts, which can then override its functions as necessary.
 */
abstract contract DelegationShareBase is Initializable, PausableUpgradeable, IDelegationShare {
    using SafeERC20 for IERC20;

    /// @notice DelegationManager contract
    IDelegationManager public delegationManager;

    /// @notice The underyling token for shares in this DelegationShare
    IERC20 public underlyingToken;

    /// @notice The total number of extant shares in the DelegationShare
    uint256 public totalShares;

    event Deposit(address depositor, address token, uint256 amount);

    event Withdraw(address depositor, address token, uint256 amount);

    /// @notice Simply checks that the `msg.sender` is the `DelegationManager`, which is an address stored immutably at construction.
    modifier onlyDelegationManager() {
        require(msg.sender == address(delegationManager), "DelegationShareBase.onlyDelegationManager");
        _;
    }

    /**
     * @notice Used to deposit tokens into this DelegationShare
     * @param token is the ERC20 token being deposited
     * @param amount is the amount of token being deposited
     * @dev This function is only callable by the DelegationManager contract. It is invoked inside of the delegationManager's
     * `depositIntoStrategy` function, and individual share balances are recorded in the delegationManager as well.
     * @return newShares is the number of new shares issued at the current exchange ratio.
     */
    function deposit(address depositor, IERC20 token, uint256 amount)
        external
        virtual
        override
        whenNotPaused
        onlyDelegationManager
        returns (uint256 newShares)
    {
        require(token == underlyingToken, "DelegationShareBase.deposit: Can only deposit underlyingToken");
        // be ware of lines below, if min amount is too small there will be a share calculation exploit problem
        (bool success, bytes memory data) = address(token).call(
            abi.encodeWithSignature("decimals()")
        );
        require(success, "underlyingToken have no method with decimals");
        uint256 decimals = uint256(bytes32(data));
        require(amount >= 1*10**decimals, "amount must gt 1 unit");

        /**
         * @notice calculation of newShares *mirrors* `underlyingToShares(amount)`, but is different since the balance of `underlyingToken`
         * has already been increased due to the `delegationManager` transferring tokens to this delegation contract prior to calling this function
         */
        uint256 priorTokenBalance = _tokenBalance() - amount;
        if (priorTokenBalance == 0 || totalShares == 0) {
            newShares = amount;
        } else {
            newShares = (amount * totalShares) / priorTokenBalance;
        }

        totalShares += newShares;
        emit Deposit(depositor, address(token), amount);
        return newShares;
    }

    /**
     * @notice Used to withdraw tokens from this DelegationShare, to the `depositor`'s address
     * @param token is the ERC20 token being transferred out
     * @param amountShares is the amount of shares being withdrawn
     * @dev This function is only callable by the delegationManager contract. It is invoked inside of the delegationManager's
     * other functions, and individual share balances are recorded in the delegationManager as well.
     */
    function withdraw(address depositor, IERC20 token, uint256 amountShares)
        external
        virtual
        override
        whenNotPaused
        onlyDelegationManager
    {
        require(token == underlyingToken, "DelegationShareBase.withdraw: Can only withdraw the strategy token");
        require(
            amountShares <= totalShares,
            "DelegationShareBase.withdraw: amountShares must be less than or equal to totalShares"
        );
        // copy `totalShares` value prior to decrease
        uint256 priorTotalShares = totalShares;
        // Decrease `totalShares` to reflect withdrawal. Unchecked arithmetic since we just checked this above.
        unchecked {
            totalShares -= amountShares;
        }
        /**
         * @notice calculation of amountToSend *mirrors* `sharesToUnderlying(amountShares)`, but is different since the `totalShares` has already
         * been decremented
         */
        uint256 amountToSend;
        if (priorTotalShares == amountShares) {
            amountToSend = _tokenBalance();
        } else {
            amountToSend = (_tokenBalance() * amountShares) / priorTotalShares;
        }
        underlyingToken.safeTransfer(depositor, amountToSend);
        emit Withdraw(depositor, address(token), amountToSend);
    }

    /**
     * @notice Currently returns a brief string explaining the strategy's goal & purpose, but for more complex
     * strategies, may be a link to metadata that explains in more detail.
     */
    function explanation() external pure virtual override returns (string memory) {
        // return "Base DelegationShare implementation to inherit from for more complex implementations";
        return "Mantle token DelegationShare implementation for submodules as an example";
    }

    /**
     * @notice Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.
     * @notice In contrast to `sharesToUnderlying`, this function guarantees no state modifications
     * @param amountShares is the amount of shares to calculate its conversion into the underlying token
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function sharesToUnderlyingView(uint256 amountShares) public view virtual override returns (uint256) {
        if (totalShares == 0) {
            return amountShares;
        } else {
            return (_tokenBalance() * amountShares) / totalShares;
        }
    }

    /**
     * @notice Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.
     * @notice In contrast to `sharesToUnderlyingView`, this function **may** make state modifications
     * @param amountShares is the amount of shares to calculate its conversion into the underlying token
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function sharesToUnderlying(uint256 amountShares) public view virtual override returns (uint256) {
        return sharesToUnderlyingView(amountShares);
    }

    /**
     * @notice Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.
     * @notice In contrast to `underlyingToShares`, this function guarantees no state modifications
     * @param amountUnderlying is the amount of `underlyingToken` to calculate its conversion into strategy shares
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function underlyingToSharesView(uint256 amountUnderlying) public view virtual returns (uint256) {
        uint256 tokenBalance = _tokenBalance();
        if (tokenBalance == 0 || totalShares == 0) {
            return amountUnderlying;
        } else {
            return (amountUnderlying * totalShares) / tokenBalance;
        }
    }

    /**
     * @notice Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.
     * @notice In contrast to `underlyingToSharesView`, this function **may** make state modifications
     * @param amountUnderlying is the amount of `underlyingToken` to calculate its conversion into strategy shares
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function underlyingToShares(uint256 amountUnderlying) external view virtual returns (uint256) {
        return underlyingToSharesView(amountUnderlying);
    }

    /**
     * @notice convenience function for fetching the current underlying value of all of the `user`'s shares in
     * this strategy. In contrast to `userUnderlying`, this function guarantees no state modifications
     */
    function userUnderlyingView(address user) external view virtual returns (uint256) {
        return sharesToUnderlyingView(shares(user));
    }

    /**
     * @notice convenience function for fetching the current underlying value of all of the `user`'s shares in
     * this strategy. In contrast to `userUnderlyingView`, this function **may** make state modifications
     */
    function userUnderlying(address user) external virtual returns (uint256) {
        return sharesToUnderlying(shares(user));
    }

    /**
     * @notice convenience function for fetching the current total shares of `user` in this strategy, by
     * querying the `delegationManager` contract
     */
    function shares(address user) public view virtual returns (uint256) {
        return IDelegationManager(delegationManager).investorDelegationShares(user, IDelegationShare(address(this)));
    }

    /// @notice Internal function used to fetch this contract's current balance of `underlyingToken`.
    // slither-disable-next-line dead-code
    function _tokenBalance() internal view virtual returns (uint256) {
        return underlyingToken.balanceOf(address(this));
    }
}

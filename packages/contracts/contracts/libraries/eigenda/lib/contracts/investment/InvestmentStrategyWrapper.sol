// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../interfaces/IInvestmentManager.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title Extremely simple implementation of `IInvestmentStrategy` interface.
 * @author Layr Labs, Inc.
 * @notice Simple, basic, "do-nothing" InvestmentStrategy that holds a single underlying token and returns it on withdrawals.
 * Assumes shares are always 1-to-1 with the underlyingToken.
 * @dev Unlike `InvestmentStrategyBase`, this contract is *not* designed to be inherited from.
 */
contract InvestmentStrategyWrapper is IInvestmentStrategy {
    using SafeERC20 for IERC20;

    /// @notice EigenLayer's InvestmentManager contract
    IInvestmentManager public immutable investmentManager;

    /// @notice The underyling token for shares in this InvestmentStrategy
    IERC20 public immutable underlyingToken;

    /// @notice The total number of extant shares in thie InvestmentStrategy
    uint256 public totalShares;

    modifier onlyInvestmentManager() {
        require(msg.sender == address(investmentManager), "InvestmentStrategyWrapper.onlyInvestmentManager");
        _;
    }

    constructor(IInvestmentManager _investmentManager, IERC20 _underlyingToken) {
        investmentManager = _investmentManager;
        underlyingToken = _underlyingToken;
    }

    /**
     * @notice Used to deposit tokens into this InvestmentStrategy
     * @param token is the ERC20 token being deposited
     * @param amount is the amount of token being deposited
     * @dev This function is only callable by the investmentManager contract. It is invoked inside of the investmentManager's
     * `depositIntoStrategy` function, and individual share balances are recorded in the investmentManager as well.
     * @return newShares is the number of new shares issued at the current exchange ratio.
     */
    function deposit(IERC20 token, uint256 amount) external virtual override onlyInvestmentManager returns (uint256) {
        require(token == underlyingToken, "InvestmentStrategyWrapper.deposit: Can only deposit underlyingToken");
        totalShares += amount;
        return amount;
    }

    /**
     * @notice Used to withdraw tokens from this InvestmentStrategy, to the `depositor`'s address
     * @param token is the ERC20 token being transferred out
     * @param amountShares is the amount of shares being withdrawn
     * @dev This function is only callable by the investmentManager contract. It is invoked inside of the investmentManager's
     * other functions, and individual share balances are recorded in the investmentManager as well.
     */
    function withdraw(address depositor, IERC20 token, uint256 amountShares)
        external
        virtual
        override
        onlyInvestmentManager
    {
        require(token == underlyingToken, "InvestmentStrategyWrapper.withdraw: Can only withdraw the strategy token");
        require(
            amountShares <= totalShares,
            "InvestmentStrategyWrapper.withdraw: amountShares must be less than or equal to totalShares"
        );
        // Decrease `totalShares` to reflect withdrawal. Unchecked arithmetic since we just checked this above.
        unchecked {
            totalShares -= amountShares;
        }
        underlyingToken.safeTransfer(depositor, amountShares);
    }

    /**
     * @notice Currently returns a brief string explaining the strategy's goal & purpose, but for more complex
     * strategies, may be a link to metadata that explains in more detail.
     */
    function explanation() external pure virtual override returns (string memory) {
        return "Wrapper InvestmentStrategy to simply store tokens. Assumes fixed 1-to-1 share-underlying exchange.";
    }

    /**
     * @notice Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.
     * @notice In contrast to `sharesToUnderlying`, this function guarantees no state modifications
     * @param amountShares is the amount of shares to calculate its conversion into the underlying token
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function sharesToUnderlyingView(uint256 amountShares) public view virtual override returns (uint256) {
        return amountShares;
    }

    /**
     * @notice Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.
     * @notice In contrast to `sharesToUnderlyingView`, this function **may** make state modifications
     * @param amountShares is the amount of shares to calculate its conversion into the underlying token
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function sharesToUnderlying(uint256 amountShares) public view virtual override returns (uint256) {
        return amountShares;
    }

    /**
     * @notice Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.
     * @notice In contrast to `underlyingToShares`, this function guarantees no state modifications
     * @param amountUnderlying is the amount of `underlyingToken` to calculate its conversion into strategy shares
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function underlyingToSharesView(uint256 amountUnderlying) external view virtual returns (uint256) {
        return amountUnderlying;
    }

    /**
     * @notice Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.
     * @notice In contrast to `underlyingToSharesView`, this function **may** make state modifications
     * @param amountUnderlying is the amount of `underlyingToken` to calculate its conversion into strategy shares
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function underlyingToShares(uint256 amountUnderlying) external view virtual returns (uint256) {
        return amountUnderlying;
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
     * querying the `investmentManager` contract
     */
    function shares(address user) public view virtual returns (uint256) {
        return IInvestmentManager(investmentManager).investorStratShares(user, IInvestmentStrategy(address(this)));
    }
}

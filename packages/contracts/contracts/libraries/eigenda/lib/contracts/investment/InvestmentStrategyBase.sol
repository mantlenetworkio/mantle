// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../interfaces/IInvestmentManager.sol";
import "../permissions/Pausable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

/**
 * @title Base implementation of `IInvestmentStrategy` interface, designed to be inherited from by more complex strategies.
 * @author Layr Labs, Inc.
 * @notice Simple, basic, "do-nothing" InvestmentStrategy that holds a single underlying token and returns it on withdrawals.
 * Implements minimal versions of the IInvestmentStrategy functions, this contract is designed to be inherited by
 * more complex investment strategies, which can then override its functions as necessary.
 */
contract InvestmentStrategyBase is Initializable, Pausable, IInvestmentStrategy {
    using SafeERC20 for IERC20;

    /// @notice EigenLayer's InvestmentManager contract
    IInvestmentManager public immutable investmentManager;

    /// @notice The underyling token for shares in this InvestmentStrategy
    IERC20 public underlyingToken;

    /// @notice The total number of extant shares in thie InvestmentStrategy
    uint256 public totalShares;

    /// @notice Simply checks that the `msg.sender` is the `investmentManager`, which is an address stored immutably at construction.
    modifier onlyInvestmentManager() {
        require(msg.sender == address(investmentManager), "InvestmentStrategyBase.onlyInvestmentManager");
        _;
    }

    /// @notice Since this contract is designed to be initializable, the constructor simply sets `investmentManager`, the only immutable variable.
    constructor(IInvestmentManager _investmentManager) {
        investmentManager = _investmentManager;
        _disableInitializers();
    }

    /// @notice Sets the `underlyingToken` and `pauserRegistry` for the strategy.
    function initialize(IERC20 _underlyingToken, IPauserRegistry _pauserRegistry) public initializer {
        underlyingToken = _underlyingToken;
        _initializePauser(_pauserRegistry);
    }

    /**
     * @notice Used to deposit tokens into this InvestmentStrategy
     * @param token is the ERC20 token being deposited
     * @param amount is the amount of token being deposited
     * @dev This function is only callable by the investmentManager contract. It is invoked inside of the investmentManager's
     * `depositIntoStrategy` function, and individual share balances are recorded in the investmentManager as well.
     * @return newShares is the number of new shares issued at the current exchange ratio.
     */
    function deposit(IERC20 token, uint256 amount)
        external
        virtual
        override
        whenNotPaused
        onlyInvestmentManager
        returns (uint256 newShares)
    {
        require(token == underlyingToken, "InvestmentStrategyBase.deposit: Can only deposit underlyingToken");

        /**
         * @notice calculation of newShares *mirrors* `underlyingToShares(amount)`, but is different since the balance of `underlyingToken`
         * has already been increased due to the `investmentManager` transferring tokens to this strategy prior to calling this function
         */
        uint256 priorTokenBalance = _tokenBalance() - amount;
        if (priorTokenBalance == 0 || totalShares == 0) {
            newShares = amount;
        } else {
            newShares = (amount * totalShares) / priorTokenBalance;
        }

        totalShares += newShares;
        return newShares;
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
        whenNotPaused
        onlyInvestmentManager
    {
        require(token == underlyingToken, "InvestmentStrategyBase.withdraw: Can only withdraw the strategy token");
        require(
            amountShares <= totalShares,
            "InvestmentStrategyBase.withdraw: amountShares must be less than or equal to totalShares"
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
    }

    /**
     * @notice Currently returns a brief string explaining the strategy's goal & purpose, but for more complex
     * strategies, may be a link to metadata that explains in more detail.
     */
    function explanation() external pure virtual override returns (string memory) {
        return "Base InvestmentStrategy implementation to inherit from for more complex implementations";
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
     * querying the `investmentManager` contract
     */
    function shares(address user) public view virtual returns (uint256) {
        return IInvestmentManager(investmentManager).investorStratShares(user, IInvestmentStrategy(address(this)));
    }

    /// @notice Internal function used to fetch this contract's current balance of `underlyingToken`.
    // slither-disable-next-line dead-code
    function _tokenBalance() internal view virtual returns (uint256) {
        return underlyingToken.balanceOf(address(this));
    }
}

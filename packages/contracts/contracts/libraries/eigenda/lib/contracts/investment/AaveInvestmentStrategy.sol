// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./aave/ILendingPool.sol";
import "./InvestmentStrategyBase.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

/**
 * @title InvestmentStrategy that lends tokens out on AAVE.
 * @author Layr Labs, Inc.
 * @notice Passively lends tokens on AAVE. Does not perform any borrowing.
 * @dev This contract is designed to accept deposits and process withdrawals in *either* the underlyingToken or aTokens
 */
abstract contract AaveInvestmentStrategy is Initializable, InvestmentStrategyBase {
    using SafeERC20 for IERC20;

    ILendingPool public lendingPool;
    IERC20 public aToken;

    // solhint-disable-next-line no-empty-blocks
    constructor(IInvestmentManager _investmentManager) InvestmentStrategyBase(_investmentManager) {}

    function initialize(
        IERC20 _underlyingToken,
        ILendingPool _lendingPool,
        IERC20 _aToken,
        IPauserRegistry _pauserRegistry
    ) public initializer {
        super.initialize(_underlyingToken, _pauserRegistry);
        lendingPool = _lendingPool;
        aToken = _aToken;
        underlyingToken.safeApprove(address(_lendingPool), type(uint256).max);
    }

    /**
     * @notice Used to deposit tokens into this InvestmentStrategy
     * @param token is the ERC20 token being deposited
     * @param amount is the amount of token being deposited
     * @dev This function is only callable by the investmentManager contract. It is invoked inside of the investmentManager's
     * `depositIntoStrategy` function, and individual share balances are recorded in the investmentManager as well
     * @return newShares is the number of new shares issued at the current exchange ratio.
     * For this strategy, the exchange ratio is fixed at (1 underlying token) / (1 share) due to the nature of AAVE's ATokens.
     * @notice This strategy accepts deposits either in the form of `underlyingToken` OR `aToken`
     */
    function deposit(IERC20 token, uint256 amount)
        external
        override
        whenNotPaused
        onlyInvestmentManager
        returns (uint256 newShares)
    {
        uint256 aTokenIncrease;
        uint256 aTokensBefore;
        if (token == underlyingToken) {
            //deposit and the "shares" are in proportion to the new aTokens minted
            aTokensBefore = aToken.balanceOf(address(this));

            //tokens have already been transferred to this contract
            //underlyingToken.transferFrom(depositor, address(this), amounts[0]);
            lendingPool.deposit(address(underlyingToken), amount, address(this), 0);

            // increment in the aToken balance of this contract due to the new investment
            aTokenIncrease = aToken.balanceOf(address(this)) - aTokensBefore;
        } else if (token == aToken) {
            aTokenIncrease = amount;

            // total aToken with this contract before the new investment,
            // this includes interest rates accrued on existing investment
            aTokensBefore = aToken.balanceOf(address(this)) - amount;
        } else {
            revert("AaveInvestmentStrategy.deposit: can only deposit underlyingToken or aToken");
        }
        if (totalShares == 0) {
            // no existing investment into this investment strategy
            newShares = aTokenIncrease;
        } else {
            /**
             * @dev Evaluating the number of new shares that would be issued for the increase
             * in aToken at the current price of each share in terms of aToken. This
             * price is given by aTokensBefore / totalShares.
             */
            newShares = (aTokenIncrease * totalShares) / aTokensBefore;
        }

        // incrementing the total number of shares
        totalShares += newShares;
    }

    /**
     * @notice Used to withdraw tokens from this InvestmentStrategy, to the `depositor`'s address
     * @param token is the ERC20 token being transferred out
     * @param shareAmount is the amount of shares being withdrawn
     * @dev This function is only callable by the investmentManager contract. It is invoked inside of the investmentManager's
     * other functions, and individual share balances are recorded in the investmentManager as well
     * @notice This strategy distributes withdrawals either in the form of `underlyingToken` OR `aToken`
     */
    function withdraw(address depositor, IERC20 token, uint256 shareAmount) external override onlyInvestmentManager {
        uint256 toWithdraw = sharesToUnderlying(shareAmount);

        if (token == underlyingToken) {
            //withdraw from lendingPool
            uint256 amountWithdrawn = lendingPool.withdraw(address(underlyingToken), toWithdraw, depositor);

            // transfer the underlyingToken to the depositor
            underlyingToken.safeTransfer(depositor, amountWithdrawn);
        } else if (token == aToken) {
            aToken.safeTransfer(depositor, toWithdraw);
        } else {
            revert("AaveInvestmentStrategy.withdraw: can only withdraw as underlyingToken or aToken");
        }

        // update the total shares for this investment strategy
        totalShares -= shareAmount;
    }

    function explanation() external pure override returns (string memory) {
        return "A simple investment strategy that allows a single asset to be deposited and loans it out on Aave";
    }

    /// @notice Internal function used to fetch this contract's current balance of `aToken`.
    function _tokenBalance() internal view override returns (uint256) {
        return aToken.balanceOf(address(this));
    }
}

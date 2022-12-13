// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./AaveInvestmentStrategy.sol";
import "./LIDO/IStableSwapStateOracle.sol";

/**
 * @title InvestmentStrategy that lends LIDO's stETH (0xae7ab96520de3a18e5e111b5eaab095312d7fe84) out on AAVE.
 * @author Layr Labs, Inc.
 * @notice Passively lends tokens on AAVE. Does not perform any borrowing.
 * @dev This contract is designed to accept deposits and process withdrawals in *either* the underlyingToken or aTokens
 * @dev This contract uses LIDO's StableSwapStateOracle to determine the current stETH/ETH ratio -- see https://docs.lido.fi/contracts/stable-swap-state-oracle
 * The StableSwapStateOracle on Mainnet is here https://etherscan.io/address/0x3a6bd15abf19581e411621d669b6a2bbe741ffd6#readContract
 */
contract AaveInvestmentStrategy_STETH is AaveInvestmentStrategy {
    IStableSwapStateOracle public stableSwapOracle;

    // solhint-disable-next-line no-empty-blocks
    constructor(IInvestmentManager _investmentManager) AaveInvestmentStrategy(_investmentManager) {}

    function initialize(
        IERC20 _underlyingToken,
        ILendingPool _lendingPool,
        IERC20 _aToken,
        IStableSwapStateOracle _stableSwapOracle,
        IPauserRegistry _pauserRegistry
    ) external initializer {
        super.initialize(_underlyingToken, _lendingPool, _aToken, _pauserRegistry);
        stableSwapOracle = _stableSwapOracle;
    }

    /**
     * @notice Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.
     * This strategy uses LIDO's `stableSwapOracle` to estimate the conversion from stETH to ETH.
     * @notice In contrast to `sharesToUnderlying`, this function guarantees no state modifications
     * @param amountShares is the amount of shares to calculate its conversion into the underlying token
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function sharesToUnderlyingView(uint256 amountShares) public view override returns (uint256) {
        (,,, uint256 exchangeRate) = stableSwapOracle.getState();
        return (super.sharesToUnderlyingView(amountShares) * exchangeRate) / 1e18;
    }

    /**
     * @notice Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.
     * This strategy uses LIDO's `stableSwapOracle` to estimate the conversion from ETH to stETH.
     * @notice In contrast to `underlyingToShares`, this function guarantees no state modifications
     * @param amountUnderlying is the amount of `underlyingToken` to calculate its conversion into strategy shares
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function underlyingToSharesView(uint256 amountUnderlying) public view override returns (uint256) {
        (,,, uint256 exchangeRate) = stableSwapOracle.getState();
        return (super.underlyingToSharesView(amountUnderlying) * 1e18) / exchangeRate;
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./TestHelper.t.sol";
import "../contracts/investment/InvestmentManagerStorage.sol";

contract InvestmentStrategyTests is TestHelper {
    /// @notice This function tests to ensure that a delegation contract
    ///         cannot be intitialized multiple times
    function testCannotInitMultipleTimesDelegation() public cannotReinit {
        wethStrat.initialize(weth, pauserReg);
    }

    ///@notice This function tests to ensure that only the investmentManager
    ///         can deposit into a strategy
    ///@param invalidDepositor is the non-registered depositor
    function testInvalidCalltoDeposit(address invalidDepositor) public fuzzedAddress(invalidDepositor) {
        IERC20 underlyingToken = wethStrat.underlyingToken();

        cheats.startPrank(invalidDepositor);
        cheats.expectRevert(bytes("InvestmentStrategyBase.onlyInvestmentManager"));
        wethStrat.deposit(underlyingToken, 1e18);
        cheats.stopPrank();
    }

    ///@notice This function tests to ensure that only the investmentManager
    ///         can deposit into a strategy
    ///@param invalidWithdrawer is the non-registered withdrawer
    ///@param depositor is the depositor for which the shares are being withdrawn
    function testInvalidCalltoWithdraw(address depositor, address invalidWithdrawer)
        public
        fuzzedAddress(invalidWithdrawer)
    {
        IERC20 underlyingToken = wethStrat.underlyingToken();

        cheats.startPrank(invalidWithdrawer);
        cheats.expectRevert(bytes("InvestmentStrategyBase.onlyInvestmentManager"));
        wethStrat.withdraw(depositor, underlyingToken, 1e18);
        cheats.stopPrank();
    }

    ///@notice This function tests ensures that withdrawing for a depositor that never
    ///         actually deposited fails.
    ///@param depositor is the depositor for which the shares are being withdrawn
    function testWithdrawalExceedsTotalShares(address depositor, uint256 shares) public fuzzedAddress(depositor) {
        cheats.assume(shares > investmentManager.investorStratShares(depositor, wethStrat));
        IERC20 underlyingToken = wethStrat.underlyingToken();

        cheats.startPrank(address(investmentManager));

        cheats.expectRevert(
            bytes("InvestmentStrategyBase.withdraw: amountShares must be less than or equal to totalShares")
        );
        wethStrat.withdraw(depositor, underlyingToken, shares);

        cheats.stopPrank();
    }
}

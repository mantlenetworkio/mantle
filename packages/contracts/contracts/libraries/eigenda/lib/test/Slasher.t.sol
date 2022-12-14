// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./Deployer.t.sol";
import "./TestHelper.t.sol";

contract SlasherTests is TestHelper {
    /**
     * @notice this function tests the slashing process by first freezing
     * the operator and then calling the investmentManager.slashShares()
     * to actually enforce the slashing conditions.
     */
    function testSlashing() public {
        IInvestmentStrategy[] memory strategyArray = new IInvestmentStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);

        // hardcoded inputs
        address[2] memory accounts = [acct_0, acct_1];
        uint256[2] memory depositAmounts;
        uint256 amountToDeposit = 1e7;
        address _operator = operator;
        strategyArray[0] = wethStrat;
        tokensArray[0] = weth;

        // have `_operator` make deposits in WETH strategy
        _testWethDeposit(_operator, amountToDeposit);
        // register `_operator` as an operator
        _testRegisterAsOperator(_operator, IDelegationTerms(_operator));

        // make deposit in WETH strategy from each of `accounts`, then delegate them to `_operator`
        for (uint256 i = 0; i < accounts.length; i++) {
            depositAmounts[i] = _testWethDeposit(accounts[i], amountToDeposit);
            _testDelegateToOperator(accounts[i], _operator);
        }

        uint256[] memory shareAmounts = new uint256[](1);
        shareAmounts[0] = depositAmounts[0];

        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        // investmentManager.queueWithdrawal(strategyIndexes, strategyArray, tokensArray, shareAmounts, nonce);
        cheats.startPrank(address(slasher.delegation()));
        slasher.freezeOperator(_operator);
        cheats.stopPrank();

        uint256 prev_shares = delegation.operatorShares(_operator, strategyArray[0]);

        investmentManager.slashShares(_operator, acct_0, strategyArray, tokensArray, strategyIndexes, shareAmounts);

        require(
            delegation.operatorShares(_operator, strategyArray[0]) + shareAmounts[0] == prev_shares,
            "Malicious Operator slashed by incorrect amount"
        );
    }

    /**
     * @notice testing ownable permissions for slashing functions
     * addPermissionedContracts(), removePermissionedContracts()
     * and resetFrozenStatus().
     */
    function testOnlyOwnerFunctions(address incorrectCaller, address inputAddr)
        public
        fuzzedAddress(incorrectCaller)
        fuzzedAddress(inputAddr)
    {
        cheats.assume(incorrectCaller != slasher.owner());
        cheats.startPrank(incorrectCaller);
        address[] memory addressArray = new address[](1);
        addressArray[0] = inputAddr;
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        slasher.addGloballyPermissionedContracts(addressArray);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        slasher.removeGloballyPermissionedContracts(addressArray);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        slasher.resetFrozenStatus(addressArray);
        cheats.stopPrank();
    }
}

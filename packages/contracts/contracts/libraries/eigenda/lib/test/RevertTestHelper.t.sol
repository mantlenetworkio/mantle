// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../contracts/libraries/BytesLib.sol";
import "../test/TestHelper.t.sol";

contract RevertTestHelper is TestHelper {
    using BytesLib for bytes;

    function _testShouldRevertRegisterOperatorWithDataLayr(
        uint8 operatorIndex,
        uint8 operatorType,
        string memory socket,
        uint256 amountEigenToDeposit,
        uint256 amountEthToDeposit,
        bytes memory revertMessage
    ) public {
        address operator = signers[operatorIndex];
        //setting up operator's delegation terms
        weth.transfer(operator, 1e18);
        weth.transfer(_challenger, 1e18);
        _testRegisterAsOperator(operator, IDelegationTerms(operator));

        for (uint256 i; i < delegates.length; i++) {
            //initialize weth, eigen and eth balances for delegator
            eigenToken.transfer(delegates[i], amountEigenToDeposit);
            weth.transfer(delegates[i], amountEthToDeposit);
            cheats.deal(delegates[i], amountEthToDeposit);

            cheats.startPrank(delegates[i]);

            //deposit delegator's eigen into investment manager
            eigenToken.approve(address(investmentManager), type(uint256).max);

            investmentManager.depositIntoStrategy(eigenStrat, eigenToken, amountEigenToDeposit);

            //deposit weth into investment manager
            weth.approve(address(investmentManager), type(uint256).max);
            investmentManager.depositIntoStrategy(wethStrat, weth, amountEthToDeposit);
            cheats.stopPrank();

            uint256 operatorEigenSharesBefore = delegation.operatorShares(operator, eigenStrat);
            uint256 operatorWETHSharesBefore = delegation.operatorShares(operator, wethStrat);

            //delegate delegator's deposits to operator
            _testDelegateToOperator(delegates[i], operator);
            //testing to see if increaseOperatorShares worked
            assertTrue(
                delegation.operatorShares(operator, eigenStrat) - operatorEigenSharesBefore == amountEigenToDeposit
            );
            assertTrue(delegation.operatorShares(operator, wethStrat) - operatorWETHSharesBefore == amountEthToDeposit);
        }

        cheats.startPrank(operator);
        //whitelist the dlsm to slash the operator
        slasher.allowToSlash(address(dlsm));
        pubkeyCompendium.registerBLSPublicKey(registrationData[operatorIndex]);
        cheats.expectRevert(revertMessage);
        dlReg.registerOperator(operatorType, testEphemeralKey, registrationData[operatorIndex].slice(0, 128), socket);
        cheats.stopPrank();
    }
}

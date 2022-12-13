// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./TestHelper.t.sol";
import "../contracts/investment/InvestmentManagerStorage.sol";
import "./utils/DataStoreUtilsWrapper.sol";
import "./mocks/ServiceManagerMock.sol";

contract InvestmentTests is TestHelper {
    /**
     * @notice Verifies that it is possible to deposit WETH
     * @param amountToDeposit Fuzzed input for amount of WETH to deposit
     */
    function testWethDeposit(uint256 amountToDeposit) public returns (uint256 amountDeposited) {
        return _testWethDeposit(signers[0], amountToDeposit);
    }

// TODO: reimplement with queued withdrawals
/*
    ///@notice This test verifies that it is possible to withdraw WETH after depositing it
    ///@param amountToDeposit The amount of WETH to try depositing
    ///@param amountToWithdraw The amount of shares to try withdrawing
    function testWethWithdrawal(uint96 amountToDeposit, uint96 amountToWithdraw) public {
        // want to deposit at least 1 wei
        cheats.assume(amountToDeposit > 0);
        // want to withdraw at least 1 wei
        cheats.assume(amountToWithdraw > 0);
        // cannot withdraw more than we deposit
        cheats.assume(amountToWithdraw <= amountToDeposit);
        // hard-coded inputs
        address sender = signers[0];
        uint256 strategyIndex = 0;
        _testDepositToStrategy(sender, amountToDeposit, weth, wethStrat);
        _testWithdrawFromStrategy(sender, strategyIndex, amountToWithdraw, weth, wethStrat);
    }
*/
// TODO: reimplement with queued withdrawals
/*
    /**
     * @notice Verifies that a strategy gets removed from the dynamic array 'investorStrats' when the user no longer has any shares in the strategy
     * @param amountToDeposit Fuzzed input for the amount deposited into the strategy, prior to withdrawing all shares
     */
/*
    function testRemovalOfStrategyOnWithdrawal(uint96 amountToDeposit) public {
        // hard-coded inputs
        IInvestmentStrategy _strat = wethStrat;
        IERC20 underlyingToken = weth;
        address sender = signers[0];

        _testDepositToStrategy(sender, amountToDeposit, underlyingToken, _strat);
        uint256 investorStratsLengthBefore = investmentManager.investorStratsLength(sender);
        uint256 investorSharesBefore = investmentManager.investorStratShares(sender, _strat);
        _testWithdrawFromStrategy(sender, 0, investorSharesBefore, underlyingToken, _strat);
        uint256 investorSharesAfter = investmentManager.investorStratShares(sender, _strat);
        uint256 investorStratsLengthAfter = investmentManager.investorStratsLength(sender);
        assertEq(investorSharesAfter, 0, "testRemovalOfStrategyOnWithdrawal: did not remove all shares!");
        assertEq(
            investorStratsLengthBefore - investorStratsLengthAfter,
            1,
            "testRemovalOfStrategyOnWithdrawal: strategy not removed from dynamic array when it should be"
        );
    }
*/
    /**
     * Testing queued withdrawals in the investment manager
     * @notice This test registers `staker` as a delegate if `registerAsDelegate` is set to 'true', deposits `amountToDeposit` into a simple WETH strategy,
     * and then starts a queued withdrawal for `amountToWithdraw` of shares in the same WETH strategy. It then tries to call `completeQueuedWithdrawal`
     * and verifies that it correctly (passes) reverts in the event that the `staker` is (not) delegated.
     * @notice In the event that the call to `completeQueuedWithdrawal` correctly reverted above, this function then fast-forwards to just past the `unlockTime`
     * for the queued withdrawal and verifies that a call to `completeQueuedWithdrawal` completes appropriately.
     * @param staker The caller who will create the queued withdrawal.
     * @param registerAsOperator When true, `staker` will register as a delegate inside of the call to `_createQueuedWithdrawal`. Otherwise they will not.
     * @param amountToDeposit Fuzzed input of amount of WETH deposited. Currently `_createQueuedWithdrawal` uses this as an input to `_testWethDeposit`.
     * @param amountToWithdraw Fuzzed input of the amount of shares to queue the withdrawal for.
     */
    function testQueuedWithdrawal(
        address staker,
        bool registerAsOperator,
        uint96 amountToDeposit,
        uint96 amountToWithdraw
    )
        public
        fuzzedAddress(staker)
    {
        // want to deposit at least 1 wei
        cheats.assume(amountToDeposit > 0);
        // want to withdraw at least 1 wei
        cheats.assume(amountToWithdraw > 0);
        // cannot withdraw more than we deposit
        cheats.assume(amountToWithdraw <= amountToDeposit);

        IInvestmentStrategy[] memory strategyArray = new IInvestmentStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        uint256[] memory strategyIndexes = new uint256[](1);

        // harcoded inputs, also somewhat shared with `testFraudproofQueuedWithdrawal`
        {
            strategyArray[0] = wethStrat;
            tokensArray[0] = weth;
            shareAmounts[0] = amountToWithdraw;
            strategyIndexes[0] = 0;
        }

        IInvestmentManager.WithdrawerAndNonce memory withdrawerAndNonce =
            IInvestmentManager.WithdrawerAndNonce({withdrawer: staker, nonce: 0});

        // create the queued withdrawal
       (bytes32 withdrawalRoot, IInvestmentManager.QueuedWithdrawal memory queuedWithdrawal) = _createQueuedWithdrawal(
            staker,
            registerAsOperator,
            amountToDeposit,
            strategyArray,
            tokensArray,
            shareAmounts,
            strategyIndexes,
            withdrawerAndNonce
        );

        cheats.startPrank(staker);
        // If `staker` was actively delegated when queued withdrawal was initiated, then verify that the next call -- to `completeQueuedWithdrawal` -- reverts appropriately
        if (queuedWithdrawal.delegatedAddress != address(0)) {
            cheats.expectRevert(
                "InvestmentManager.completeQueuedWithdrawal: withdrawal waiting period has not yet passed and depositor was delegated when withdrawal initiated"
            );
        }

        // try to complete the queued withdrawal
        investmentManager.completeQueuedWithdrawal(queuedWithdrawal, true);
        // TODO: add checks surrounding successful completion (e.g. funds being correctly transferred)

        if (queuedWithdrawal.delegatedAddress != address(0)) {
            // retrieve information about the queued withdrawal
            // bytes32 withdrawalRoot = investmentManager.calculateWithdrawalRoot(strategyArray, tokensArray, shareAmounts, withdrawerAndNonce);
            // (uint32 initTimestamp, uint32 unlockTimestamp, address withdrawer) = investmentManager.queuedWithdrawals(withdrawalRoot);
            uint32 unlockTimestamp;
            {
                (, unlockTimestamp,) = investmentManager.queuedWithdrawals(withdrawalRoot);
            }
            // warp to unlock time (i.e. past fraudproof period) and verify that queued withdrawal works at this time
            cheats.warp(unlockTimestamp);
            investmentManager.completeQueuedWithdrawal(queuedWithdrawal, true);
        }
        cheats.stopPrank();
    }

    /**
     * @notice This test checks that fraudproofing queued withdrawals through the InvestmentManager is possible.
     * @param amountToDeposit Fuzzed input of amount of WETH deposited. Currently `_createQueuedWithdrawal` uses this as an input to `_testWethDeposit`.
     * @param amountToWithdraw Fuzzed input of the amount of shares to queue the withdrawal for.
     */
    function testFraudproofQueuedWithdrawal(uint96 amountToDeposit, uint96 amountToWithdraw) public {
        // want to deposit at least 1 wei
        cheats.assume(amountToDeposit > 0);
        // want to withdraw at least 1 wei
        cheats.assume(amountToWithdraw > 0);
        // cannot withdraw more than we deposit
        cheats.assume(amountToWithdraw <= amountToDeposit);

        IInvestmentStrategy[] memory strategyArray = new IInvestmentStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        uint256[] memory strategyIndexes = new uint256[](1);
        // harcoded inputs, also somewhat shared with `testQueuedWithdrawal`
        {
            strategyArray[0] = wethStrat;
            tokensArray[0] = weth;
            shareAmounts[0] = amountToWithdraw;
            strategyIndexes[0] = 0;
        }
        
        // harcoded inputs
        address staker = acct_0;
        bool registerAsOperator = true;
        IInvestmentManager.WithdrawerAndNonce memory withdrawerAndNonce =
            IInvestmentManager.WithdrawerAndNonce({withdrawer: staker, nonce: 0});

        (bytes32 withdrawalRoot, IInvestmentManager.QueuedWithdrawal memory queuedWithdrawal) = _createQueuedWithdrawal(
            staker,
            registerAsOperator,
            amountToDeposit,
            strategyArray,
            tokensArray,
            shareAmounts,
            strategyIndexes,
            withdrawerAndNonce
        );

        // warp to a later time -- beyond the window for the `REASONABLE_STAKES_UPDATE_PERIOD` -- and then initiate the queued withdrawal waiting period
        cheats.warp(block.timestamp + 8 days);
        _testStartQueuedWithdrawalWaitingPeriod(
            staker,
            withdrawalRoot,
            (uint32(block.timestamp) + 9 days)
        );

        ServiceManagerMock mock = new ServiceManagerMock();
        bytes memory calldataForStakeWithdrawalVerification;

        // give slashing permission to the mock contract
        {
            cheats.startPrank(slasher.owner());
            address[] memory contractsToGiveSlashingPermission = new address[](1);
            contractsToGiveSlashingPermission[0] = address(mock);
            slasher.addGloballyPermissionedContracts(contractsToGiveSlashingPermission);
            cheats.stopPrank();
        }

        // fraudproof the queued withdrawal

        // function fraudproofQueuedWithdrawal(
        //     IInvestmentStrategy[] calldata strategies,
        //     IERC20[] calldata tokens,
        //     uint256[] calldata shareAmounts,
        //     address depositor,
        //     WithdrawerAndNonce calldata withdrawerAndNonce,
        //     bytes calldata data,
        //     IServiceManager slashingContract
        // ) external {
        investmentManager.challengeQueuedWithdrawal(queuedWithdrawal, calldataForStakeWithdrawalVerification, IServiceManager(address(mock)));
    }

    /// @notice deploys 'numStratsToAdd' strategies using '_testAddStrategy' and then deposits '1e18' to each of them from 'signers[0]'
    /// @param numStratsToAdd is the number of strategies being added and deposited into
    function testDepositStrategies(uint16 numStratsToAdd) public {
        _testDepositStrategies(signers[0], 1e18, numStratsToAdd);
    }

    /// @notice Verifies that it is possible to deposit eigen.
    /// @param eigenToDeposit is amount of eigen to deposit into the eigen strategy
    function testDepositEigen(uint96 eigenToDeposit) public {
        // sanity check for inputs; keeps fuzzed tests from failing
        cheats.assume(eigenToDeposit < eigenTotalSupply);
        _testDepositEigen(signers[0], eigenToDeposit);
    }

    /**
     * @notice Tries to deposit an unsupported token into an `InvestmentStrategyBase` contract by calling `investmentManager.depositIntoStrategy`.
     * Verifies that reversion occurs correctly.
     */
    function testDepositUnsupportedToken() public {
        IERC20 token = new ERC20PresetFixedSupply(
            "badToken",
            "BADTOKEN",
            100,
            address(this)
        );
        token.approve(address(investmentManager), type(uint256).max);
        cheats.expectRevert(bytes("InvestmentStrategyBase.deposit: Can only deposit underlyingToken"));
        investmentManager.depositIntoStrategy(wethStrat, token, 10);
    }

    /**
     * @notice Tries to deposit into an unsupported strategy by calling `investmentManager.depositIntoStrategy`.
     * Verifies that reversion occurs correctly.
     */
    function testDepositNonexistantStrategy(address nonexistentStrategy) public fuzzedAddress(nonexistentStrategy) {
        // assume that the fuzzed address is not already a contract!
        uint256 size;
        assembly {
            size := extcodesize(nonexistentStrategy)
        }
        cheats.assume(size == 0);
        // check against calls from precompile addresses -- was getting fuzzy failures from this
        cheats.assume(uint160(nonexistentStrategy) > 9);

        // harcoded input
        uint256 testDepositAmount = 10;

        IERC20 token = new ERC20PresetFixedSupply(
            "badToken",
            "BADTOKEN",
            100,
            address(this)
        );
        token.approve(address(investmentManager), type(uint256).max);
        cheats.expectRevert();
        investmentManager.depositIntoStrategy(IInvestmentStrategy(nonexistentStrategy), token, testDepositAmount);
    }

    // TODO: add test(s) that confirm deposits + withdrawals *of zero shares* fail correctly.
}

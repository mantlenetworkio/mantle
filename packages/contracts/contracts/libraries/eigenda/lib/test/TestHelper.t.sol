// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../contracts/libraries/BytesLib.sol";
import "../test/Deployer.t.sol";


contract TestHelper is EigenLayrDeployer {
    using BytesLib for bytes;

    uint8 durationToInit = 2;
    uint256 public SECP256K1N_MODULUS = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141;
    uint256 public SECP256K1N_MODULUS_HALF = 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0;

    function _testInitiateDelegation(
        uint8 operatorIndex,
        uint256 amountEigenToDeposit, 
        uint256 amountEthToDeposit        
    )
        public returns (uint256 amountEthStaked, uint256 amountEigenStaked)
    {
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
        amountEthStaked += delegation.operatorShares(operator, wethStrat);
        amountEigenStaked += delegation.operatorShares(operator, eigenStrat);

        return (amountEthStaked, amountEigenStaked);
    }

    // simply tries to register 'sender' as a delegate, setting their 'DelegationTerms' contract in EigenLayrDelegation to 'dt'
    // verifies that the storage of EigenLayrDelegation contract is updated appropriately
    function _testRegisterAsOperator(address sender, IDelegationTerms dt) internal {
        cheats.startPrank(sender);
        delegation.registerAsOperator(dt);
        assertTrue(delegation.isOperator(sender), "testRegisterAsOperator: sender is not a delegate");

        assertTrue(
            delegation.delegationTerms(sender) == dt, "_testRegisterAsOperator: delegationTerms not set appropriately"
        );

        assertTrue(delegation.isDelegated(sender), "_testRegisterAsOperator: sender not marked as actively delegated");
        cheats.stopPrank();
    }

    /**
     * @notice Deposits `amountToDeposit` of WETH from address `sender` into `wethStrat`.
     * @param sender The address to spoof calls from using `cheats.startPrank(sender)`
     * @param amountToDeposit Amount of WETH that is first *transferred from this contract to `sender`* and then deposited by `sender` into `stratToDepositTo`
     */
    function _testWethDeposit(address sender, uint256 amountToDeposit) internal returns (uint256 amountDeposited) {
        cheats.assume(amountToDeposit <= wethInitialSupply);
        // transfer WETH to `sender` and have them deposit it into `strat`
        amountDeposited = _testDepositToStrategy(sender, amountToDeposit, weth, wethStrat);
    }

    /**
     * @notice Deposits `amountToDeposit` of EIGEN from address `sender` into `eigenStrat`.
     * @param sender The address to spoof calls from using `cheats.startPrank(sender)`
     * @param amountToDeposit Amount of EIGEN that is first *transferred from this contract to `sender`* and then deposited by `sender` into `stratToDepositTo`
     */
    function _testDepositEigen(address sender, uint256 amountToDeposit) public {
        _testDepositToStrategy(sender, amountToDeposit, eigenToken, eigenStrat);
    }

    /**
     * @notice Deposits `amountToDeposit` of `underlyingToken` from address `sender` into `stratToDepositTo`.
     * *If*  `sender` has zero shares prior to deposit, *then* checks that `stratToDepositTo` is correctly added to their `investorStrats` array.
     *
     * @param sender The address to spoof calls from using `cheats.startPrank(sender)`
     * @param amountToDeposit Amount of WETH that is first *transferred from this contract to `sender`* and then deposited by `sender` into `stratToDepositTo`
     */
    function _testDepositToStrategy(
        address sender,
        uint256 amountToDeposit,
        IERC20 underlyingToken,
        IInvestmentStrategy stratToDepositTo
    )
        internal
        returns (uint256 amountDeposited)
    {
        // deposits will revert when amountToDeposit is 0
        cheats.assume(amountToDeposit > 0);

        uint256 operatorSharesBefore = investmentManager.investorStratShares(sender, stratToDepositTo);
        // assumes this contract already has the underlying token!
        uint256 contractBalance = underlyingToken.balanceOf(address(this));
        // logging and error for misusing this function (see assumption above)
        if (amountToDeposit > contractBalance) {
            emit log("amountToDeposit > contractBalance");
            emit log_named_uint("amountToDeposit is", amountToDeposit);
            emit log_named_uint("while contractBalance is", contractBalance);
            revert("_testDepositToStrategy failure");
        } else {

            
            underlyingToken.transfer(sender, amountToDeposit);
            cheats.startPrank(sender);
            underlyingToken.approve(address(investmentManager), type(uint256).max);

            investmentManager.depositIntoStrategy(stratToDepositTo, underlyingToken, amountToDeposit);
            amountDeposited = amountToDeposit;

            //check if depositor has never used this strat, that it is added correctly to investorStrats array.
            if (operatorSharesBefore == 0) {
                // check that strategy is appropriately added to dynamic array of all of sender's strategies
                assertTrue(
                    investmentManager.investorStrats(sender, investmentManager.investorStratsLength(sender) - 1)
                        == stratToDepositTo,
                    "_depositToStrategy: investorStrats array updated incorrectly"
                );
            }

            
            


            //in this case, since shares never grow, the shares should just match the deposited amount
            assertEq(
                investmentManager.investorStratShares(sender, stratToDepositTo) - operatorSharesBefore,
                amountDeposited,
                "_depositToStrategy: shares should match deposit"
            );
        }
        cheats.stopPrank();
    }

// TODO: reimplement with queued withdrawals
/*
    //checks that it is possible to withdraw from the given `stratToWithdrawFrom`
    function _testWithdrawFromStrategy(
        address sender,
        uint256 strategyIndex,
        uint256 amountSharesToWithdraw,
        IERC20 underlyingToken,
        IInvestmentStrategy stratToWithdrawFrom
    )
        internal
    {
        // fetch the length of `sender`'s dynamic `investorStrats` array
        uint256 investorStratsLengthBefore = investmentManager.investorStratsLength(sender);
        // fetch `sender`'s existing share amount
        uint256 existingShares = investmentManager.investorStratShares(sender, stratToWithdrawFrom);
        // fetch `sender`'s existing balance of `underlyingToken`
        uint256 senderUnderlyingBalanceBefore = underlyingToken.balanceOf(sender);

        // sanity checks on `strategyIndex` input
        if (strategyIndex >= investorStratsLengthBefore) {
            emit log("_testWithdrawFromStrategy: attempting to withdraw from out-of-bounds index");
            revert("_testWithdrawFromStrategy: attempting to withdraw from out-of-bounds index");
        }
        assertEq(address(stratToWithdrawFrom), address(investmentManager.investorStrats(sender, strategyIndex)));

        cheats.prank(sender);
        //trying to withdraw more than the amountDeposited will fail, so we expect a revert and *short-circuit* if it happens
        if (amountSharesToWithdraw > existingShares) {
            cheats.expectRevert(bytes("InvestmentManager._removeShares: shareAmount too high"));
            investmentManager.withdrawFromStrategy(
                strategyIndex, stratToWithdrawFrom, underlyingToken, amountSharesToWithdraw
            );
            return;
        } else {
            investmentManager.withdrawFromStrategy(
                strategyIndex, stratToWithdrawFrom, underlyingToken, amountSharesToWithdraw
            );
        }

        uint256 senderUnderlyingBalanceAfter = underlyingToken.balanceOf(sender);

        assertEq(
            amountSharesToWithdraw,
            senderUnderlyingBalanceAfter - senderUnderlyingBalanceBefore,
            "_testWithdrawFromStrategy: shares differ from 1-to-1 with underlyingToken?"
        );
        cheats.stopPrank();
    }
*/

    // tries to delegate from 'sender' to 'operator'
    // verifies that:
    //                  delegator has at least some shares
    //                  delegatedShares update correctly for 'operator'
    //                  delegated status is updated correctly for 'sender'
    function _testDelegateToOperator(address sender, address operator) internal {
        //delegator-specific information
        (IInvestmentStrategy[] memory delegateStrategies, uint256[] memory delegateShares) =
            investmentManager.getDeposits(sender);

        uint256 numStrats = delegateShares.length;
        assertTrue(numStrats > 0, "_testDelegateToOperator: delegating from address with no investments");
        uint256[] memory inititalSharesInStrats = new uint256[](numStrats);
        for (uint256 i = 0; i < numStrats; ++i) {
            inititalSharesInStrats[i] = delegation.operatorShares(operator, delegateStrategies[i]);
        }

        cheats.startPrank(sender);
        delegation.delegateTo(operator);
        cheats.stopPrank();

        assertTrue(
            delegation.delegatedTo(sender) == operator,
            "_testDelegateToOperator: delegated address not set appropriately"
        );
        assertTrue(
            delegation.delegationStatus(sender) == IEigenLayrDelegation.DelegationStatus.DELEGATED,
            "_testDelegateToOperator: delegated status not set appropriately"
        );

        for (uint256 i = 0; i < numStrats; ++i) {
            uint256 operatorSharesBefore = inititalSharesInStrats[i];
            uint256 operatorSharesAfter = delegation.operatorShares(operator, delegateStrategies[i]);
            assertTrue(
                operatorSharesAfter == (operatorSharesBefore + delegateShares[i]),
                "_testDelegateToOperator: delegatedShares not increased correctly"
            );
        }
    }

    // deploys a InvestmentStrategyBase contract and initializes it to treat `underlyingToken` as its underlying token
    function _testAddStrategyBase(IERC20 underlyingToken) internal returns (IInvestmentStrategy) {
        InvestmentStrategyBase strategy = InvestmentStrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(baseStrategyImplementation),
                    address(eigenLayrProxyAdmin),
                    abi.encodeWithSelector(InvestmentStrategyBase.initialize.selector, underlyingToken, pauserReg)
                )
            )
        );
        return strategy;
    }

    // deploys 'numStratsToAdd' strategies using '_testAddStrategyBase' and then deposits 'amountToDeposit' to each of them from 'sender'
    function _testDepositStrategies(address sender, uint256 amountToDeposit, uint16 numStratsToAdd) internal {
        // hard-coded inputs
        uint96 multiplier = 1e18;
        IERC20 underlyingToken = weth;

        cheats.assume(numStratsToAdd > 0 && numStratsToAdd <= 20);
        IInvestmentStrategy[] memory stratsToDepositTo = new IInvestmentStrategy[](
                numStratsToAdd
            );
        for (uint16 i = 0; i < numStratsToAdd; ++i) {
            stratsToDepositTo[i] = _testAddStrategyBase(underlyingToken);
            _testDepositToStrategy(sender, amountToDeposit, weth, InvestmentStrategyBase(address(stratsToDepositTo[i])));
        }
        for (uint16 i = 0; i < numStratsToAdd; ++i) {
            // check that strategy is appropriately added to dynamic array of all of sender's strategies
            assertTrue(
                investmentManager.investorStrats(sender, i) == stratsToDepositTo[i],
                "investorStrats array updated incorrectly"
            );

            // TODO: perhaps remove this is we can. seems brittle if we don't track the number of strategies somewhere
            //store strategy in mapping of strategies
            strategies[i] = IInvestmentStrategy(address(stratsToDepositTo[i]));
        }
        // add strategies to dlRegistry
        for (uint16 i = 0; i < numStratsToAdd; ++i) {
            VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] memory ethStratsAndMultipliers =
            new VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[](
                    1
                );
            ethStratsAndMultipliers[0].strategy = stratsToDepositTo[i];
            ethStratsAndMultipliers[0].multiplier = multiplier;
            dlReg.addStrategiesConsideredAndMultipliers(0, ethStratsAndMultipliers);
        }
    }


    /**
     * @notice Creates a queued withdrawal from `staker`. Begins by registering the staker as a delegate (if specified), then deposits `amountToDeposit`
     * into the WETH strategy, and then queues a withdrawal using
     * `investmentManager.queueWithdrawal(strategyIndexes, strategyArray, tokensArray, shareAmounts, withdrawerAndNonce)`
     * @notice After initiating a queued withdrawal, this test checks that `investmentManager.canCompleteQueuedWithdrawal` immediately returns the correct
     * response depending on whether `staker` is delegated or not.
     * @param staker The address to initiate the queued withdrawal
     * @param registerAsOperator If true, `staker` will also register as a delegate in the course of this function
     * @param amountToDeposit The amount of WETH to deposit
     */
    function _createQueuedWithdrawal(
        address staker,
        bool registerAsOperator,
        uint256 amountToDeposit,
        IInvestmentStrategy[] memory strategyArray,
        IERC20[] memory tokensArray,
        uint256[] memory shareAmounts,
        uint256[] memory strategyIndexes,
        IInvestmentManager.WithdrawerAndNonce memory withdrawerAndNonce
    )
        internal returns(bytes32 withdrawalRoot, IInvestmentManager.QueuedWithdrawal memory queuedWithdrawal)
    {
        require(amountToDeposit >= shareAmounts[0], "_createQueuedWithdrawal: sanity check failed");

        // we do this here to ensure that `staker` is delegated if `registerAsOperator` is true
        if (registerAsOperator) {
            assertTrue(!delegation.isDelegated(staker), "_createQueuedWithdrawal: staker is already delegated");
            _testRegisterAsOperator(staker, IDelegationTerms(staker));
            assertTrue(
                delegation.isDelegated(staker), "_createQueuedWithdrawal: staker isn't delegated when they should be"
            );
        }

        queuedWithdrawal = IInvestmentManager.QueuedWithdrawal({
            strategies: strategyArray,
            tokens: tokensArray,
            shares: shareAmounts,
            depositor: staker,
            withdrawerAndNonce: withdrawerAndNonce,
            delegatedAddress: delegation.delegatedTo(staker)
        });

        {
            //make deposit in WETH strategy
            uint256 amountDeposited = _testWethDeposit(staker, amountToDeposit);
            // We can't withdraw more than we deposit
            if (shareAmounts[0] > amountDeposited) {
                cheats.expectRevert("InvestmentManager._removeShares: shareAmount too high");
            }
        }

        //queue the withdrawal
        cheats.startPrank(staker);
        // TODO: check with 'undelegateIfPossible' = false, rather than just true
        withdrawalRoot = investmentManager.queueWithdrawal(strategyIndexes, strategyArray, tokensArray, shareAmounts, withdrawerAndNonce, true);
        // If `staker` was actively delegated at time of queuing the withdrawal, check that `canCompleteQueuedWithdrawal` correct returns 'false', and
        if (queuedWithdrawal.delegatedAddress != address(0)) {
            assertTrue(
                !investmentManager.canCompleteQueuedWithdrawal(queuedWithdrawal),
                "_createQueuedWithdrawal: user can immediately complete queued withdrawal (before waiting for fraudproof period), depsite being delegated"
            );
        }
        // If `staker` was *not* actively delegated at time of queuing the withdrawal, check that `canCompleteQueuedWithdrawal` correct returns 'true'
        else if (queuedWithdrawal.delegatedAddress == address(0)) {
            assertTrue(
                investmentManager.canCompleteQueuedWithdrawal(queuedWithdrawal),
                "_createQueuedWithdrawal: user *cannot* immediately complete queued withdrawal (before waiting for fraudproof period), despite *not* being delegated"
            );
        } else {
            revert("_createQueuedWithdrawal: staker was somehow neither delegated nor *not* delegated, simultaneously");
        }
        cheats.stopPrank();
        return (withdrawalRoot, queuedWithdrawal);
    }

    function _testStartQueuedWithdrawalWaitingPeriod(
        address withdrawer,
        bytes32 withdrawalRoot,
        uint32 stakeInactiveAfter
    ) internal {
        cheats.startPrank(withdrawer);
        // TODO: un-hardcode the '8 days' and '30 days' here
        // '8 days' accounts for the `REASONABLE_STAKES_UPDATE_PERIOD`
        cheats.warp(block.timestamp + 8 days);
        // '30 days' is used to prevent overflow in timestamps when stored as uint32 values (2^32 is in the year 2106 in UTC time)
        cheats.assume(stakeInactiveAfter < type(uint32).max - 30 days);
        cheats.assume(stakeInactiveAfter > block.timestamp);
        investmentManager.startQueuedWithdrawalWaitingPeriod(
                                        withdrawalRoot, 
                                        stakeInactiveAfter
                                    );
        cheats.stopPrank();
    }

    function getG2PublicKeyHash(bytes calldata data, address signer) public view returns(bytes32 pkHash){

        uint256[4] memory pk;
        // verify sig of public key and get pubkeyHash back, slice out compressed apk
        (pk[0], pk[1], pk[2], pk[3]) = BLS.verifyBLSSigOfPubKeyHash(data, signer);

        pkHash = keccak256(abi.encodePacked(pk[0], pk[1], pk[2], pk[3]));

        return pkHash;

    }

    function getG2PKOfRegistrationData(uint8 operatorIndex) internal view returns(uint256[4] memory){
        uint256[4] memory pubkey; 
        pubkey[0] = uint256(bytes32(registrationData[operatorIndex].slice(32,32)));
        pubkey[1] = uint256(bytes32(registrationData[operatorIndex].slice(0,32)));
        pubkey[2] = uint256(bytes32(registrationData[operatorIndex].slice(96,32)));
        pubkey[3] = uint256(bytes32(registrationData[operatorIndex].slice(64,32)));
        return pubkey;
    }


    function getVSfromVandS(uint8 v, bytes32 s) internal view returns(bytes32){
        if (uint256(s) > SECP256K1N_MODULUS_HALF) {
            s = bytes32(SECP256K1N_MODULUS - uint256(s));
        }

        bytes32 vs = s;
        if(v == 28){
            vs = bytes32(uint256(s) ^ (1 << 255));
        }

        return vs;
    }

}


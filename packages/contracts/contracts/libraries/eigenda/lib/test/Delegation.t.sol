// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/math/Math.sol";

import "../test/DataLayrTestHelper.t.sol";


import "../contracts/libraries/BytesLib.sol";


contract DelegationTests is DataLayrTestHelper {
    using BytesLib for bytes;
    using Math for uint256;

    uint256[] sharesBefore;
    uint256[] balanceBefore;
    uint256[] priorTotalShares;
    uint256[] strategyTokenBalance;

    uint256 public PRIVATE_KEY = 420;

    // packed info used to help handle stack-too-deep errors
    struct DataForTestWithdrawal {
        IInvestmentStrategy[] delegatorStrategies;
        uint256[] delegatorShares;
        IInvestmentManager.WithdrawerAndNonce withdrawerAndNonce;
    }

    /// @notice testing if an operator can register to themselves.
    function testSelfOperatorRegister() public {
        _testRegisterAdditionalSelfOperator(signers[0], registrationData[0], ephemeralKeyHashes[0]);
    }

    /// @notice testing if an operator can delegate to themselves.
    /// @param sender is the address of the operator.
    function testSelfOperatorDelegate(address sender) public {
        cheats.assume(sender != address(0));
        cheats.assume(sender != address(eigenLayrProxyAdmin));
        _testRegisterAsOperator(sender, IDelegationTerms(sender));
    }

    function testTwoSelfOperatorsRegister() public {
        _testRegisterAdditionalSelfOperator(signers[0], registrationData[0], ephemeralKeyHashes[0]);
        _testRegisterAdditionalSelfOperator(signers[1], registrationData[1], ephemeralKeyHashes[1]);
    }

    /// @notice registers a fixed address as a delegate, delegates to it from a second address,
    ///         and checks that the delegate's voteWeights increase properly
    /// @param operator is the operator being delegated to.
    /// @param staker is the staker delegating stake to the operator.
    function testDelegation(address operator, address staker, uint256 ethAmount, uint256 eigenAmount)
        public
        fuzzedAddress(operator)
        fuzzedAddress(staker)
    {
        cheats.assume(staker != operator);
        cheats.assume(ethAmount >= 0 && ethAmount <= 1e18);
        cheats.assume(eigenAmount >= 0 && eigenAmount <= 1e18);

        if (!delegation.isOperator(operator)) {
            _testRegisterAsOperator(operator, IDelegationTerms(operator));
        }

        uint256[3] memory amountsBefore;
        amountsBefore[0] = dlReg.weightOfOperator(operator, 0);
        amountsBefore[1] = dlReg.weightOfOperator(operator, 1);
        amountsBefore[2] = delegation.operatorShares(operator, wethStrat);

        //making additional deposits to the investment strategies
        assertTrue(delegation.isNotDelegated(staker) == true, "testDelegation: staker is not delegate");
        _testWethDeposit(staker, ethAmount);
        _testDepositEigen(staker, eigenAmount);
        _testDelegateToOperator(staker, operator);
        assertTrue(delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");

        (IInvestmentStrategy[] memory updatedStrategies, uint256[] memory updatedShares) =
            investmentManager.getDeposits(staker);

        {
            uint256 stakerEthWeight = investmentManager.investorStratShares(staker, updatedStrategies[0]);
            uint256 stakerEigenWeight = investmentManager.investorStratShares(staker, updatedStrategies[1]);

            uint256 operatorEthWeightAfter = dlReg.weightOfOperator(operator, 0);
            uint256 operatorEigenWeightAfter = dlReg.weightOfOperator(operator, 1);

            assertTrue(
                operatorEthWeightAfter - amountsBefore[0] == stakerEthWeight,
                "testDelegation: operatorEthWeight did not increment by the right amount"
            );
            assertTrue(
                operatorEigenWeightAfter - amountsBefore[1] == stakerEigenWeight,
                "Eigen weights did not increment by the right amount"
            );
        }
        {
            IInvestmentStrategy _strat = wethStrat;
            // IInvestmentStrategy _strat = investmentManager.investorStrats(staker, 0);
            assertTrue(address(_strat) != address(0), "investorStrats not updated correctly");

            assertTrue(
                delegation.operatorShares(operator, _strat) - updatedShares[0] == amountsBefore[2],
                "ETH operatorShares not updated correctly"
            );
        }
    }

    /// @notice tests delegation to EigenLayr via an ECDSA signatures - meta transactions are the future bby
    /// @param operator is the operator being delegated to.
    function testDelegateToBySignature(address operator, uint256 ethAmount, uint256 eigenAmount)
        public
        fuzzedAddress(operator)
    {
        cheats.assume(ethAmount >= 0 && ethAmount <= 1e18);
        cheats.assume(eigenAmount >= 0 && eigenAmount <= 1e18);
    

        if (!delegation.isOperator(operator)) {
            _testRegisterAsOperator(operator, IDelegationTerms(operator));
        }
        address staker = cheats.addr(PRIVATE_KEY);
        cheats.assume(staker != operator);

        //making additional deposits to the investment strategies
        assertTrue(delegation.isNotDelegated(staker) == true, "testDelegation: staker is not delegate");
        _testWethDeposit(staker, ethAmount);
        _testDepositEigen(staker, eigenAmount);

        uint256 nonceBefore = delegation.nonces(staker);

        bytes32 structHash = keccak256(abi.encode(delegation.DELEGATION_TYPEHASH(), staker, operator, nonceBefore, 0));
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", delegation.DOMAIN_SEPARATOR(), structHash));


        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(PRIVATE_KEY, digestHash);

        bytes32 vs = getVSfromVandS(v, s);
        
        delegation.delegateToBySignature(staker, operator, 0, r, vs);
        assertTrue(delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");
        assertTrue(nonceBefore + 1 == delegation.nonces(staker), "nonce not incremented correctly");
        assertTrue(delegation.delegatedTo(staker) == operator, "staker delegated to wrong operator");
    }

    /// @notice tests delegation to EigenLayr via an ECDSA signatures with invalid signature
    /// @param operator is the operator being delegated to.
    function testDelegateToByInvalidSignature(
        address operator, 
        uint256 ethAmount, 
        uint256 eigenAmount, 
        uint8 v,
        bytes32 r,
        bytes32 s
    )
        public
        fuzzedAddress(operator)
    {
        cheats.assume(ethAmount >= 0 && ethAmount <= 1e18);
        cheats.assume(eigenAmount >= 0 && eigenAmount <= 1e18);
    

        if (!delegation.isOperator(operator)) {
            _testRegisterAsOperator(operator, IDelegationTerms(operator));
        }
        address staker = cheats.addr(PRIVATE_KEY);
        cheats.assume(staker != operator);

        //making additional deposits to the investment strategies
        assertTrue(delegation.isNotDelegated(staker) == true, "testDelegation: staker is not delegate");
        _testWethDeposit(staker, ethAmount);
        _testDepositEigen(staker, eigenAmount);

        bytes32 vs = getVSfromVandS(v, s);
        
        cheats.expectRevert();
        delegation.delegateToBySignature(staker, operator, 0, r, vs);
        
    }

    /// @notice registers a fixed address as a delegate, delegates to it from a second address,
    /// and checks that the delegate's voteWeights increase properly
    /// @param operator is the operator being delegated to.
    /// @param staker is the staker delegating stake to the operator.
    function testDelegationMultipleStrategies(uint16 numStratsToAdd, address operator, address staker)
        public
        fuzzedAddress(operator)
        fuzzedAddress(staker)
    {
        cheats.assume(staker != operator);

        cheats.assume(numStratsToAdd > 0 && numStratsToAdd <= 20);
        uint96 operatorEthWeightBefore = dlReg.weightOfOperator(operator, 0);
        uint96 operatorEigenWeightBefore = dlReg.weightOfOperator(operator, 1);
        _testRegisterAsOperator(operator, IDelegationTerms(operator));
        _testDepositStrategies(staker, 1e18, numStratsToAdd);
        _testDepositEigen(staker, 1e18);
        _testDelegateToOperator(staker, operator);
        uint96 operatorEthWeightAfter = dlReg.weightOfOperator(operator, 0);
        uint96 operatorEigenWeightAfter = dlReg.weightOfOperator(operator, 1);
        assertTrue(
            operatorEthWeightAfter > operatorEthWeightBefore, "testDelegation: operatorEthWeight did not increase!"
        );
        assertTrue(
            operatorEigenWeightAfter > operatorEigenWeightBefore, "testDelegation: operatorEthWeight did not increase!"
        );
    }

    /// @notice test staker's ability to undelegate/withdraw from an operator.
    /// @param operator is the operator being delegated to.
    /// @param depositor is the staker delegating stake to the operator.
    function testWithdrawal(
            address operator, 
            address depositor,
            address withdrawer, 
            uint256 ethAmount,
            uint256 eigenAmount,
            uint32 stakeInactiveAfter,
            bool withdrawAsTokens
        ) 
            public 
            fuzzedAddress(operator) 
            fuzzedAddress(depositor) 
            fuzzedAddress(withdrawer) 
        {

        cheats.assume(depositor != operator);
        cheats.assume(ethAmount <= 1e18); 
        cheats.assume(eigenAmount <= 1e18); 
        cheats.assume(ethAmount > 0); 
        cheats.assume(eigenAmount > 0); 

        testDelegation(operator, depositor, ethAmount, eigenAmount);

        address delegatedTo = delegation.delegatedTo(depositor);

        // packed data structure to deal with stack-too-deep issues
        DataForTestWithdrawal memory dataForTestWithdrawal;

        // scoped block to deal with stack-too-deep issues
        {
            //delegator-specific information
            (IInvestmentStrategy[] memory delegatorStrategies, uint256[] memory delegatorShares) =
                investmentManager.getDeposits(depositor);
            dataForTestWithdrawal.delegatorStrategies = delegatorStrategies;
            dataForTestWithdrawal.delegatorShares = delegatorShares;

            IInvestmentManager.WithdrawerAndNonce memory withdrawerAndNonce = 
                IInvestmentManager.WithdrawerAndNonce({
                    withdrawer: withdrawer,
                    // harcoded nonce value
                    nonce: 0
                }
            );
            dataForTestWithdrawal.withdrawerAndNonce = withdrawerAndNonce;
        }

        uint256[] memory strategyIndexes = new uint256[](2);
        IERC20[] memory tokensArray = new IERC20[](2);
        {
            // hardcoded values
            strategyIndexes[0] = 0;
            strategyIndexes[1] = 0;
            tokensArray[0] = weth;
            tokensArray[1] = eigenToken;
        }

        //initiating queued withdrawal
// TODO: resolve the presence of duplicate-ish functions for creating queued withdrawals
        // (bytes32 withdrawalRoot, ) = _createQueuedWithdrawal(
        //     depositor,
        //     // hardcoded inputs for use with this function
        //     false,
        //     0,
        //     dataForTestWithdrawal.delegatorStrategies,
        //     tokensArray,
        //     dataForTestWithdrawal.delegatorShares,
        //     strategyIndexes,
        //     dataForTestWithdrawal.withdrawerAndNonce
        // );
        bytes32 withdrawalRoot = _testQueueWithdrawal(
            depositor,
            dataForTestWithdrawal.delegatorStrategies,
            tokensArray,
            dataForTestWithdrawal.delegatorShares,
            strategyIndexes,
            dataForTestWithdrawal.withdrawerAndNonce
        );

        _testStartQueuedWithdrawalWaitingPeriod(withdrawer, withdrawalRoot, stakeInactiveAfter);

        cheats.warp(stakeInactiveAfter + 1 days);

        if (withdrawAsTokens) {
            _testCompleteQueuedWithdrawalTokens(
                depositor,
                dataForTestWithdrawal.delegatorStrategies,
                tokensArray,
                dataForTestWithdrawal.delegatorShares,
                delegatedTo,
                dataForTestWithdrawal.withdrawerAndNonce
            );
        } else {
            _testCompleteQueuedWithdrawalShares(
                depositor,
                dataForTestWithdrawal.delegatorStrategies,
                tokensArray,
                dataForTestWithdrawal.delegatorShares,
                delegatedTo,
                dataForTestWithdrawal.withdrawerAndNonce
            );
        }
    }

    /// @notice test to see if an operator who is slashed/frozen
    ///         cannot be undelegated from by their stakers.
    /// @param operator is the operator being delegated to.
    /// @param staker is the staker delegating stake to the operator.
    function testSlashedOperatorWithdrawal(address operator, address staker, uint256 ethAmount, uint256 eigenAmount)
        public
        fuzzedAddress(operator)
        fuzzedAddress(staker)
    {
        cheats.assume(staker != operator);
        testDelegation(operator, staker, ethAmount, eigenAmount);

        address slashingContract = slasher.owner();

        address[] memory slashingContracts = new address[](1);
        slashingContracts[0] = slashingContract;

        cheats.startPrank(slashingContract);
        slasher.addGloballyPermissionedContracts(slashingContracts);
        slasher.freezeOperator(operator);
        cheats.stopPrank();

        (IInvestmentStrategy[] memory updatedStrategies, uint256[] memory updatedShares) =
            investmentManager.getDeposits(staker);

        IInvestmentManager.WithdrawerAndNonce memory withdrawerAndNonce =
            IInvestmentManager.WithdrawerAndNonce({withdrawer: staker, nonce: 0});

        uint256[] memory strategyIndexes = new uint256[](2);
        strategyIndexes[0] = 0;
        strategyIndexes[1] = 1;

        IERC20[] memory tokensArray = new IERC20[](2);
        tokensArray[0] = weth;
        tokensArray[0] = eigenToken;

        //initiating queued withdrawal
        cheats.expectRevert(
            bytes("InvestmentManager.onlyNotFrozen: staker has been frozen and may be subject to slashing")
        );
        _testQueueWithdrawal(staker, updatedStrategies, tokensArray, updatedShares, strategyIndexes, withdrawerAndNonce);
    }

    /// @notice This function tests to ensure that a delegation contract
    ///         cannot be intitialized multiple times
    function testCannotInitMultipleTimesDelegation() public cannotReinit {
        //delegation has already been initialized in the Deployer test contract
        delegation.initialize(pauserReg, address(this));
    }

    /// @notice This function tests to ensure that a you can't register as a delegate multiple times
    /// @param operator is the operator being delegated to.
    function testRegisterAsOperatorMultipleTimes(address operator) public fuzzedAddress(operator) {
        _testRegisterAsOperator(operator, IDelegationTerms(operator));
        cheats.expectRevert(bytes("EigenLayrDelegation.registerAsOperator: Delegate has already registered"));
        _testRegisterAsOperator(operator, IDelegationTerms(operator));
    }

    /// @notice This function tests to ensure that a staker cannot delegate to an unregistered operator
    /// @param delegate is the unregistered operator
    function testDelegationToUnregisteredDelegate(address delegate) public fuzzedAddress(delegate) {
        //deposit into 1 strategy for signers[1], who is delegating to the unregistered operator
        _testDepositStrategies(signers[1], 1e18, 1);
        _testDepositEigen(signers[1], 1e18);

        cheats.expectRevert(bytes("EigenLayrDelegation._delegate: operator has not yet registered as a delegate"));
        cheats.startPrank(signers[1]);
        delegation.delegateTo(delegate);
        cheats.stopPrank();
    }

    // @notice This function tests to ensure that a delegator can re-delegate to an operator after undelegating.
    // @param operator is the operator being delegated to.
    // @param staker is the staker delegating stake to the operator.
    function testRedelegateAfterWithdrawal(
            address operator, 
            address depositor, 
            address withdrawer, 
            uint256 ethAmount, 
            uint256 eigenAmount,
            uint32 stakeInactiveAfter,
            bool withdrawAsShares
        ) 
            public
            fuzzedAddress(operator) 
            fuzzedAddress(depositor)
            fuzzedAddress(withdrawer)
        {
        cheats.assume(depositor != operator);
        //this function performs delegation and subsequent withdrawal
        testWithdrawal(operator, depositor, withdrawer, ethAmount, eigenAmount, stakeInactiveAfter, withdrawAsShares);

        //warps past fraudproof time interval
        cheats.warp(block.timestamp + undelegationFraudproofInterval + 1);
        testDelegation(operator, depositor, ethAmount, eigenAmount);
    }

    //testing inclusion of nonsigners in DLN quorum, ensuring that nonsigner inclusion proof is working correctly.
    function testForNonSigners(uint256 ethAmount, uint256 eigenAmount) public {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e10);

        // address operator = signers[0];
        uint8 operatorType = 3;
        _testInitiateDelegation(0, eigenAmount, ethAmount);
        _testRegisterBLSPubKey(0);
        _testRegisterOperatorWithDataLayr(0, operatorType, testEphemeralKey, testSocket);

        NonSignerPK memory nonsignerPK;
        RegistrantAPK memory registrantAPK;
        SignerAggSig memory signerAggSig;

        nonsignerPK.xA0 = (uint256(10245738255635135293623161230197183222740738674756428343303263476182774511624));
        nonsignerPK.xA1 = (uint256(10281853605827367652226404263211738087634374304916354347419537904612128636245));
        nonsignerPK.yA0 = (uint256(3091447672609454381783218377241231503703729871039021245809464784750860882084));
        nonsignerPK.yA1 = (uint256(18210007982945446441276599406248966847525243540006051743069767984995839204266));

        registrantAPK.apk0 = uint256(20820493588973199354272631301248587752629863429201347184003644368113679196121);
        registrantAPK.apk1 = uint256(18507428821816114421698399069438744284866101909563082454551586195885282320634);
        registrantAPK.apk2 = uint256(1263326262781780932600377484793962587101562728383804037421955407439695092960);
        registrantAPK.apk3 = uint256(3512517006108887301063578607317108977425754510174956792003926207778790018672);
        
        signerAggSig.sigma0 = uint256(20617740300811009543012419127686924884246271121030353570695308863131407887373);
        signerAggSig.sigma1 = uint256(11071552465919207288683976891087172465162060876240494884992829947249670282179);


        uint32 numberOfSigners = 15;
        _testRegisterSigners(numberOfSigners, false);

        // scoped block helps fix 'stack too deep' errors
        {
            uint256 initTime = 1000000001;
            IDataLayrServiceManager.DataStoreSearchData memory searchData = _testInitDataStore(initTime, address(this), header);
            uint32 numberOfNonSigners = 1;
            uint32 dataStoreId = dlsm.taskNumber() - 1;

            bytes memory data = _getCallData(
                keccak256(
                    abi.encodePacked(
                        searchData.metadata.globalDataStoreId,
                        searchData.metadata.headerHash,
                        searchData.duration,
                        initTime,
                        uint32(0)
                    )
                ),
                numberOfNonSigners,
                registrantAPK,
                signerAggSig,
                nonsignerPK,
                searchData.metadata.blockNumber,
                dataStoreId
            );

            uint256 gasbefore = gasleft();

            dlsm.confirmDataStore(data, searchData);

            emit log_named_uint("gas cost", gasbefore - gasleft());
        }
    }

    //*******INTERNAL FUNCTIONS*********//
    function _testQueueWithdrawal(
        address depositor,
        IInvestmentStrategy[] memory strategyArray,
        IERC20[] memory tokensArray,
        uint256[] memory shareAmounts,
        uint256[] memory strategyIndexes,
        IInvestmentManager.WithdrawerAndNonce memory withdrawerAndNonce
    )
        internal
        returns (bytes32)
    {
        cheats.startPrank(depositor);

        bytes32 withdrawalRoot = investmentManager.queueWithdrawal(
            strategyIndexes,
            strategyArray,
            tokensArray,
            shareAmounts,
            withdrawerAndNonce,
            // TODO: make this an input
            true
        );
        cheats.stopPrank();
        return withdrawalRoot;
    }

    function _testCompleteQueuedWithdrawalShares(
        address depositor,
        IInvestmentStrategy[] memory strategyArray,
        IERC20[] memory tokensArray,
        uint256[] memory shareAmounts,
        address delegatedTo,
        IInvestmentManager.WithdrawerAndNonce memory withdrawerAndNonce
    )
        internal
    {
        cheats.startPrank(withdrawerAndNonce.withdrawer);

        for (uint256 i = 0; i < strategyArray.length; i++) {
            sharesBefore.push(investmentManager.investorStratShares(withdrawerAndNonce.withdrawer, strategyArray[i]));
        }

        IInvestmentManager.QueuedWithdrawal memory queuedWithdrawal = IInvestmentManager.QueuedWithdrawal({
            strategies: strategyArray,
            tokens: tokensArray,
            shares: shareAmounts,
            depositor: depositor,
            withdrawerAndNonce: withdrawerAndNonce,
            delegatedAddress: delegatedTo
        });

        // complete the queued withdrawal
        investmentManager.completeQueuedWithdrawal(queuedWithdrawal, false);

        for (uint256 i = 0; i < strategyArray.length; i++) {
            require(
                investmentManager.investorStratShares(withdrawerAndNonce.withdrawer, strategyArray[i])
                    == sharesBefore[i] + shareAmounts[i],
                "_testCompleteQueuedWithdrawalShares: withdrawer shares not incremented"
            );
        }
        cheats.stopPrank();
    }

    function _testCompleteQueuedWithdrawalTokens(
        address depositor,
        IInvestmentStrategy[] memory strategyArray,
        IERC20[] memory tokensArray,
        uint256[] memory shareAmounts,
        address delegatedTo,
        IInvestmentManager.WithdrawerAndNonce memory withdrawerAndNonce
    )
        internal
    {
        cheats.startPrank(withdrawerAndNonce.withdrawer);

        for (uint256 i = 0; i < strategyArray.length; i++) {
            balanceBefore.push(strategyArray[i].underlyingToken().balanceOf(withdrawerAndNonce.withdrawer));
            priorTotalShares.push(strategyArray[i].totalShares());
            strategyTokenBalance.push(strategyArray[i].underlyingToken().balanceOf(address(strategyArray[i])));
        }

        IInvestmentManager.QueuedWithdrawal memory queuedWithdrawal = IInvestmentManager.QueuedWithdrawal({
            strategies: strategyArray,
            tokens: tokensArray,
            shares: shareAmounts,
            depositor: depositor,
            withdrawerAndNonce: withdrawerAndNonce,
            delegatedAddress: delegatedTo
        });
        // complete the queued withdrawal
        investmentManager.completeQueuedWithdrawal(queuedWithdrawal, true);

        for (uint256 i = 0; i < strategyArray.length; i++) {
            //uint256 strategyTokenBalance = strategyArray[i].underlyingToken().balanceOf(address(strategyArray[i]));
            uint256 tokenBalanceDelta = strategyTokenBalance[i] * shareAmounts[i] / priorTotalShares[i];

            require(
                strategyArray[i].underlyingToken().balanceOf(withdrawerAndNonce.withdrawer)
                    == balanceBefore[i] + tokenBalanceDelta,
                "_testCompleteQueuedWithdrawalTokens: withdrawer balance not incremented"
            );
        }
        cheats.stopPrank();
    }
}

//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

 import "../contracts/interfaces/IEigenPod.sol";
 import "./utils/BeaconChainUtils.sol";
import "./Deployer.t.sol";


contract EigenPodTests is BeaconChainProofUtils, DSTest {
    using BytesLib for bytes;

    bytes pubkey = hex"88347ed1c492eedc97fc8c506a35d44d81f27a0c7a1c661b35913cfd15256c0cccbd34a83341f505c7de2983292f2cab";
    
    //hash tree root of list of validators
    bytes32 validatorTreeRoot;

    //hash tree root of individual validator container
    bytes32 validatorRoot;

    address podOwner = address(42000094993494);


    Vm cheats = Vm(HEVM_ADDRESS);
    EigenLayrDelegation public delegation;
    InvestmentManager public investmentManager;
    Slasher public slasher;
    PauserRegistry public pauserReg;

    ProxyAdmin public eigenLayrProxyAdmin;
    IBLSPublicKeyCompendium public blsPkCompendium;
    IEigenPodManager public eigenPodManager;
    IEigenPod public pod;
    IETHPOSDeposit public ethPOSDeposit;
    IBeacon public eigenPodBeacon;
    IBeaconChainOracle public beaconChainOracle;

    address[] public slashingContracts;
    address pauser = address(69);
    address unpauser = address(489);


    //performs basic deployment before each test
    function setUp() public {
        // deploy proxy admin for ability to upgrade proxy contracts
        eigenLayrProxyAdmin = new ProxyAdmin();

        // deploy pauser registry
        pauserReg = new PauserRegistry(pauser, unpauser);

        blsPkCompendium = new BLSPublicKeyCompendium();

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        EmptyContract emptyContract = new EmptyContract();
        delegation = EigenLayrDelegation(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayrProxyAdmin), ""))
        );
        investmentManager = InvestmentManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayrProxyAdmin), ""))
        );
        slasher = Slasher(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayrProxyAdmin), ""))
        );

        beaconChainOracle = new BeaconChainOracleMock();
        beaconChainOracle.setBeaconChainStateRoot(0xb08d5a1454de19ac44d523962096d73b85542f81822c5e25b8634e4e86235413);

        ethPOSDeposit = new ETHPOSDepositMock();
        pod = new EigenPod(ethPOSDeposit);

        eigenPodBeacon = new UpgradeableBeacon(address(pod));

        // this contract is deployed later to keep its address the same (for these tests)
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayrProxyAdmin), ""))
        );

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        EigenLayrDelegation delegationImplementation = new EigenLayrDelegation(investmentManager);
        InvestmentManager investmentManagerImplementation = new InvestmentManager(delegation, eigenPodManager, slasher);
        Slasher slasherImplementation = new Slasher(investmentManager, delegation);
        EigenPodManager eigenPodManagerImplementation = new EigenPodManager(ethPOSDeposit, eigenPodBeacon, investmentManager);


        address initialOwner = address(this);
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        eigenLayrProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(delegation))),
            address(delegationImplementation),
            abi.encodeWithSelector(EigenLayrDelegation.initialize.selector, pauserReg, initialOwner)
        );
        eigenLayrProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(investmentManager))),
            address(investmentManagerImplementation),
            abi.encodeWithSelector(InvestmentManager.initialize.selector, pauserReg, initialOwner)
        );
        eigenLayrProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(slasher))),
            address(slasherImplementation),
            abi.encodeWithSelector(Slasher.initialize.selector, pauserReg, initialOwner)
        );
        eigenLayrProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation),
            abi.encodeWithSelector(EigenPodManager.initialize.selector, beaconChainOracle)
        );

        slashingContracts.push(address(eigenPodManager));
        investmentManager.slasher().addGloballyPermissionedContracts(slashingContracts);
        
    }

    
    function testDeployAndVerifyNewEigenPod(bytes memory signature, bytes32 depositDataRoot) public {
        beaconChainOracle.setBeaconChainStateRoot(0xb08d5a1454de19ac44d523962096d73b85542f81822c5e25b8634e4e86235413);

        (beaconStateMerkleProof, validatorContainerFields, validatorMerkleProof, validatorTreeRoot, validatorRoot) = getInitialDepositProof();

        cheats.startPrank(podOwner);
        eigenPodManager.stake(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        IEigenPod newPod;

        newPod = eigenPodManager.getPod(podOwner);

        bytes32 validatorIndex = bytes32(uint256(0));
        bytes memory proofs = abi.encodePacked(validatorTreeRoot, beaconStateMerkleProof, validatorRoot, validatorIndex, validatorMerkleProof);
        newPod.verifyCorrectWithdrawalCredentials(pubkey, proofs, validatorContainerFields);

        uint64 validatorBalance = Endian.fromLittleEndianUint64(validatorContainerFields[2]);
        require(eigenPodManager.getBalance(podOwner) == validatorBalance, "Validator balance not updated correctly");

        IInvestmentStrategy beaconChainETHStrategy = investmentManager.beaconChainETHStrategy();
        uint256 beaconChainETHShares = investmentManager.investorStratShares(podOwner, beaconChainETHStrategy);

        require(beaconChainETHShares == validatorBalance, "investmentManager shares not updated correctly");
    }

    function testUpdateSlashedBeaconBalance(bytes memory signature, bytes32 depositDataRoot) public {
        //make initial deposit
        testDeployAndVerifyNewEigenPod(signature, depositDataRoot);

        beaconChainOracle.setBeaconChainStateRoot(0xc50c29e99864df7c8e181ef48bef732accf96a34ff685a9b6404077a893c03f9);
        //get updated proof, set beaconchain state root
        (beaconStateMerkleProof, validatorContainerFields, validatorMerkleProof, validatorTreeRoot, validatorRoot) = getSlashedDepositProof();
        

        IEigenPod eigenPod;
        eigenPod = eigenPodManager.getPod(podOwner);
        
        bytes32 validatorIndex = bytes32(uint256(0));
        bytes memory proofs = abi.encodePacked(validatorTreeRoot, beaconStateMerkleProof, validatorRoot, validatorIndex, validatorMerkleProof);
        eigenPod.verifyBalanceUpdate(pubkey, proofs, validatorContainerFields);
        
        uint64 validatorBalance = Endian.fromLittleEndianUint64(validatorContainerFields[2]);
        require(eigenPodManager.getBalance(podOwner) == validatorBalance, "Validator balance not updated correctly");
        require(investmentManager.slasher().isFrozen(podOwner), "podOwner not frozen successfully");

    }

    function testDeployNewEigenPodWithWrongPubkey(bytes memory wrongPubkey, bytes memory signature, bytes32 depositDataRoot) public {
        beaconChainOracle.setBeaconChainStateRoot(0xb08d5a1454de19ac44d523962096d73b85542f81822c5e25b8634e4e86235413);
        (beaconStateMerkleProof, validatorContainerFields, validatorMerkleProof, validatorTreeRoot, validatorRoot) = getInitialDepositProof();

        cheats.startPrank(podOwner);
        eigenPodManager.stake(wrongPubkey, signature, depositDataRoot);
        cheats.stopPrank();

        IEigenPod newPod;
        newPod = eigenPodManager.getPod(podOwner);

        bytes32 validatorIndex = bytes32(uint256(0));
        bytes memory proofs = abi.encodePacked(validatorTreeRoot, beaconStateMerkleProof, validatorRoot, validatorIndex, validatorMerkleProof);
        cheats.expectRevert(bytes("EigenPod.verifyCorrectWithdrawalCredentials: Proof is not for provided pubkey"));
        newPod.verifyCorrectWithdrawalCredentials(wrongPubkey, proofs, validatorContainerFields);
    }

    function testDeployNewEigenPodWithWrongWithdrawalCreds(address wrongWithdrawalAddress, bytes memory signature, bytes32 depositDataRoot) public {
        (beaconStateMerkleProof, validatorContainerFields, validatorMerkleProof, validatorTreeRoot, validatorRoot) = getInitialDepositProof();

        cheats.startPrank(podOwner);
        eigenPodManager.stake(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        IEigenPod newPod;
        newPod = eigenPodManager.getPod(podOwner);
        validatorContainerFields[1] = abi.encodePacked(bytes1(uint8(1)), bytes11(0), wrongWithdrawalAddress).toBytes32(0);

        bytes32 validatorIndex = bytes32(uint256(0));
        bytes memory proofs = abi.encodePacked(validatorTreeRoot, beaconStateMerkleProof, validatorRoot, validatorIndex, validatorMerkleProof);
        cheats.expectRevert(bytes("EigenPod.verifyValidatorFields: Invalid validator fields"));
        newPod.verifyCorrectWithdrawalCredentials(pubkey, proofs, validatorContainerFields);
    }

    function testDeployNewEigenPodWithActiveValidator(bytes memory signature, bytes32 depositDataRoot) public {
        beaconChainOracle.setBeaconChainStateRoot(0xb08d5a1454de19ac44d523962096d73b85542f81822c5e25b8634e4e86235413);
        (beaconStateMerkleProof, validatorContainerFields, validatorMerkleProof, validatorTreeRoot, validatorRoot) = getInitialDepositProof();
        

        cheats.startPrank(podOwner);
        eigenPodManager.stake(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        IEigenPod newPod;
        newPod = eigenPodManager.getPod(podOwner);

        bytes32 validatorIndex = bytes32(uint256(0));
        bytes memory proofs = abi.encodePacked(validatorTreeRoot, beaconStateMerkleProof, validatorRoot, validatorIndex, validatorMerkleProof);
        newPod.verifyCorrectWithdrawalCredentials(pubkey, proofs, validatorContainerFields);

        cheats.expectRevert(bytes("EigenPod.verifyCorrectWithdrawalCredentials: Validator not inactive"));
        newPod.verifyCorrectWithdrawalCredentials(pubkey, proofs, validatorContainerFields);
    }

    function testWithdrawalEigenPods(bytes memory signature, bytes32 depositDataRoot) public {
        //make initial deposit
        testDeployAndVerifyNewEigenPod(signature, depositDataRoot);

        uint128 balance = eigenPodManager.getBalance(podOwner);

         IEigenPod newPod;
        newPod = eigenPodManager.getPod(podOwner);
        newPod.topUpPodBalance{value : balance*(1**18)}();

        IInvestmentStrategy[] memory strategyArray = new IInvestmentStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        uint256[] memory strategyIndexes = new uint256[](1);
        IInvestmentManager.WithdrawerAndNonce memory withdrawerAndNonce =
            IInvestmentManager.WithdrawerAndNonce({withdrawer: podOwner, nonce: 0});
        bool undelegateIfPossible = false;
        {
            strategyArray[0] = investmentManager.beaconChainETHStrategy();
            shareAmounts[0] = balance;
            strategyIndexes[0] = 0;
        }

        uint256 podOwnerSharesBefore = investmentManager.investorStratShares(podOwner, investmentManager.beaconChainETHStrategy());
        
        cheats.startPrank(podOwner);
        investmentManager.queueWithdrawal(strategyIndexes, strategyArray, tokensArray, shareAmounts, withdrawerAndNonce, undelegateIfPossible);
        cheats.stopPrank();

        uint256 podOwnerSharesAfter = investmentManager.investorStratShares(podOwner, investmentManager.beaconChainETHStrategy());

        require(podOwnerSharesBefore - podOwnerSharesAfter == balance, "delegation shares not updated correctly");
    } 
}


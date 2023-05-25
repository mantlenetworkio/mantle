// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import '../../../../contracts/libraries/resolver/Lib_AddressManager.sol';
import '../../../../contracts/L1/rollup/StateCommitmentChain.sol';
import '../../../../contracts/L1/messaging/L1CrossDomainMessenger.sol';
import '../../../../contracts/L1/fraud-proof/verifier/subverifiers/StackOpVerifier.sol';
import '../../../../contracts/L1/fraud-proof/verifier/test-driver/VerifierTestDriver.sol';
import '../../../../contracts/L1/fraud-proof/verifier/VerifierEntry.sol';
import '../../../../contracts/L1/fraud-proof/verifier/libraries/VerificationContext.sol';
import '../../../../contracts/L1/fraud-proof/verifier/libraries/EVMTypesLib.sol';

import '../../../../contracts/L1/local/TestBitToken.sol';

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "forge-std/Test.sol";
import "../../mocks/EmptyContract.sol";
import {console} from "forge-std/console.sol";

contract RollUpTest is Test {

    ProxyAdmin public proxyAdmin;
    TransparentUpgradeableProxy public tu;

    StackOpVerifier public blockInitiationVerifier;
    StackOpVerifier public blockFinalizationVerifier;
    StackOpVerifier public interTxVerifier;
    StackOpVerifier public stackOpVerifier;
    StackOpVerifier public environmentalOpVerifier;
    StackOpVerifier public memoryOpVerifier;
    StackOpVerifier public storageOpVerifier;
    StackOpVerifier public callOpVerifier;
    StackOpVerifier public invalidOpVerifier;
    VerifierTestDriver public verifierTestDriver;
    VerifierEntry public verifierEntry;
    VerificationContext.Context public ctx;
    
    Lib_AddressManager public addressManager;
    StateCommitmentChain public scc;

    BitTokenERC20 public l1Bit;

    function setUp() public {

        addressManager = new Lib_AddressManager();
        scc = new StateCommitmentChain(address(addressManager), address(0), 0 ,0);
        addressManager.setAddress('BVM_Rolluper', address(this));
        addressManager.setAddress('StateCommitmentChain', address(scc));
        // vm.startPrank(deployer);
        proxyAdmin = new ProxyAdmin();

        blockInitiationVerifier = new StackOpVerifier();
        blockFinalizationVerifier = new StackOpVerifier();
        interTxVerifier = new StackOpVerifier();
        stackOpVerifier = new StackOpVerifier();
        environmentalOpVerifier = new StackOpVerifier();
        memoryOpVerifier = new StackOpVerifier();
        storageOpVerifier = new StackOpVerifier();
        callOpVerifier = new StackOpVerifier();
        invalidOpVerifier = new StackOpVerifier();
        verifierTestDriver = new VerifierTestDriver(
            address(blockInitiationVerifier),
            address(blockFinalizationVerifier),
            address(interTxVerifier),
            address(stackOpVerifier),
            address(environmentalOpVerifier),
            address(memoryOpVerifier),
            address(storageOpVerifier),
            address(callOpVerifier),
            address(invalidOpVerifier)
        );
        l1Bit = new BitTokenERC20("BitToken", "BIT");

        VerifierEntry verifierEntryImpl = new VerifierEntry();
        tu = new TransparentUpgradeableProxy(address(verifierEntryImpl), address(this), abi.encodeWithSelector(verifierEntry.initialize.selector));
        
        vm.prank(address(0));
        verifierEntryImpl.setVerifier(0, stackOpVerifier);
        tu.admin();
        tu.changeAdmin(address(this));

        vm.prank(address(0));
        verifierEntryImpl.setVerifier(0, stackOpVerifier);

    }

    function testInfo() public {

        // assertEq(verifierEntryImpl.owner(), address(proxyAdmin));
    }
}
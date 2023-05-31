// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import '../../../../contracts/libraries/resolver/Lib_AddressManager.sol';
import '../../../../contracts/L1/messaging/L1CrossDomainMessenger.sol';
import '../../../../contracts/L1/rollup/StateCommitmentChain.sol';

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

contract RollUPTest is Test {

    ProxyAdmin public proxyAdmin;
    TransparentUpgradeableProxy public proxy;
    EmptyContract public empt;

    Lib_AddressManager public addressManager;
    L1CrossDomainMessenger public l1cdm;
    StateCommitmentChain public scc;

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
    
    BitTokenERC20 public l1Bit;

    function setUp() public {
        
        // vm.startPrank(deployer);
        proxyAdmin = new ProxyAdmin();

        addressManager = new Lib_AddressManager();

        L1CrossDomainMessenger l1cdmImpl = new L1CrossDomainMessenger();
        proxy = new TransparentUpgradeableProxy(
            address(l1cdmImpl), 
            address(proxyAdmin), 
            abi.encodeWithSelector(L1CrossDomainMessenger.initialize.selector, address(addressManager)));
        l1cdm = L1CrossDomainMessenger(address(proxy));

        scc = new StateCommitmentChain(address(addressManager), address(l1cdm), 0, 0);

        l1Bit = new BitTokenERC20("BitToken", "BIT");

        VerifierEntry verifierEntryImpl = new VerifierEntry();
        proxy = new TransparentUpgradeableProxy(
            address(verifierEntryImpl), 
            address(proxyAdmin), 
            abi.encodeWithSelector(verifierEntry.initialize.selector));
        verifierEntry = VerifierEntry(address(proxy));

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

        // verifierEntry.setVerifier(0, address(blockInitiationVerifier));
    }

    function testInfo() public {

        assertEq(verifierEntry.owner(), address(this));

        console.log(l1cdm.paused());
    }

}
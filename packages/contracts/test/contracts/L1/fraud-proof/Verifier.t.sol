// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import '../../../../contracts/libraries/resolver/Lib_AddressManager.sol';
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

contract VerifierTest is Test {

    ProxyAdmin public proxyAdmin;

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

        EmptyContract emptyContract = new EmptyContract();

        verifierEntry = VerifierEntry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );

        VerifierEntry verifierEntryImpl = new VerifierEntry();
        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(verifierEntry))),
            address(verifierEntryImpl),
            abi.encodeWithSelector(verifierEntry.initialize.selector, address(this))
        );

    }

    function testInfo() public {

        assertEq(verifierEntry.owner(), address(this));
    }

    function testEntryVerifyProof() public {
        EVMTypesLib.Transaction memory tx;
        VerificationContext.Context memory ctx;

        tx.nonce= 0x0;
        tx.gasPrice= 0x3b9aca00;
        tx.gas= 0x1eaed;
        tx.to = address(0x0000000000000000000000000000000000000000);
        tx.value = 0x0;
        tx.data = hex"608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100d9565b60405180910390f35b610073600480360381019061006e919061009d565b61007e565b005b60008054905090565b8060008190555050565b60008135905061009781610103565b92915050565b6000602082840312156100b3576100b26100fe565b5b60006100c184828501610088565b91505092915050565b6100d3816100f4565b82525050565b60006020820190506100ee60008301846100ca565b92915050565b6000819050919050565b600080fd5b61010c816100f4565b811461011757600080fd5b5056fea2646970667358221220404e37f487a89a932dca5e77faaf6ca2de3b991f93d230604b1b8daaef64766264736f6c63430008070033";
        tx.v = 0x69d2;
        tx.r = 0xd4cb56830de320028fc55988eb0102257955519d9dc062d80a485abb122e9e67;
        tx.s = 0x30cd3a7bdef20deae43d760be21f524a8bbec4f97a59b07d6779da50a2a1af5f;

        ctx.coinbase = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
        ctx.timestamp = 0x63b7e9c1;
        ctx.number = 0x1;
        ctx.origin = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
        ctx.transaction = tx;
        ctx.inputRoot = 0x0000000000000000000000000000000000000000000000000000000000000000;
        ctx.txHash = 0x73ea1d2ee379cd800d4006953888d2e23a57675dfc44de09e1d0194fcd866b1f;

        verifierEntry.verifyOneStepProof(
            ctx,
            0x0,
            0xea576c527587e9781b29f73d5f0c3790bb6690bd1b3b080e4b1b213f53e9f872,
            hex"00000000000000010000000000000000000100000000000106f500000000000000004262f14003fcb1c7a13c958a3cd20818d6a5e3d7777745b13c05f14385c19e6d000000000000000060c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470000000000000000000000000000000000000000000000000744c19d2e8593c97867b3b6a3588f51cd9dbc5010a395cf199be4bbb353848b8c549a82442ec9d279699caa0abffb15160db5a18e264622fc440b66cff66bf3f0000000000000000000000000000000000000000000000000000000000000000cdc5a830e025de132066c7f43de48570407bfaebc30b96f499fc06d42f5602dfac5def5a7a8b39addbb94d47a566bd1900fb621a43073880e10d6c739ae503eb00000000000000000000000000000000000000000000000000000000000000000000000000000170608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100d9565b60405180910390f35b610073600480360381019061006e919061009d565b61007e565b005b60008054905090565b8060008190555050565b60008135905061009781610103565b92915050565b6000602082840312156100b3576100b26100fe565b5b60006100c184828501610088565b91505092915050565b6100d3816100f4565b82525050565b60006020820190506100ee60008301846100ca565b92915050565b6000819050919050565b600080fd5b61010c816100f4565b811461011757600080fd5b5056fea2646970667358221220404e37f487a89a932dca5e77faaf6ca2de3b991f93d230604b1b8daaef64766264736f6c63430008070033"
        );
    }

    function testTestDriverVerifyProof() public {
        EVMTypesLib.Transaction memory tx;
        VerificationContext.Context memory ctx;

        tx.nonce= 0x0;
        tx.gasPrice= 0x3b9aca00;
        tx.gas= 0x1eaed;
        tx.to = address(0x0000000000000000000000000000000000000000);
        tx.value = 0x0;
        tx.data = hex"608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100d9565b60405180910390f35b610073600480360381019061006e919061009d565b61007e565b005b60008054905090565b8060008190555050565b60008135905061009781610103565b92915050565b6000602082840312156100b3576100b26100fe565b5b60006100c184828501610088565b91505092915050565b6100d3816100f4565b82525050565b60006020820190506100ee60008301846100ca565b92915050565b6000819050919050565b600080fd5b61010c816100f4565b811461011757600080fd5b5056fea2646970667358221220404e37f487a89a932dca5e77faaf6ca2de3b991f93d230604b1b8daaef64766264736f6c63430008070033";
        tx.v = 0x69d2;
        tx.r = 0xd4cb56830de320028fc55988eb0102257955519d9dc062d80a485abb122e9e67;
        tx.s = 0x30cd3a7bdef20deae43d760be21f524a8bbec4f97a59b07d6779da50a2a1af5f;

        ctx.coinbase = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
        ctx.timestamp = 0x63b7e9c1;
        ctx.number = 0x1;
        ctx.origin = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
        ctx.transaction = tx;
        ctx.inputRoot = 0x0000000000000000000000000000000000000000000000000000000000000000;
        ctx.txHash = 0x73ea1d2ee379cd800d4006953888d2e23a57675dfc44de09e1d0194fcd866b1f;

        verifierTestDriver.verifyProof(
            ctx.coinbase,
            ctx.timestamp,
            ctx.number,
            ctx.origin,
            ctx.txHash,
            tx,
            0x0,
            0xea576c527587e9781b29f73d5f0c3790bb6690bd1b3b080e4b1b213f53e9f872,
            hex"00000000000000010000000000000000000100000000000106f500000000000000004262f14003fcb1c7a13c958a3cd20818d6a5e3d7777745b13c05f14385c19e6d000000000000000060c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470000000000000000000000000000000000000000000000000744c19d2e8593c97867b3b6a3588f51cd9dbc5010a395cf199be4bbb353848b8c549a82442ec9d279699caa0abffb15160db5a18e264622fc440b66cff66bf3f0000000000000000000000000000000000000000000000000000000000000000cdc5a830e025de132066c7f43de48570407bfaebc30b96f499fc06d42f5602dfac5def5a7a8b39addbb94d47a566bd1900fb621a43073880e10d6c739ae503eb00000000000000000000000000000000000000000000000000000000000000000000000000000170608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100d9565b60405180910390f35b610073600480360381019061006e919061009d565b61007e565b005b60008054905090565b8060008190555050565b60008135905061009781610103565b92915050565b6000602082840312156100b3576100b26100fe565b5b60006100c184828501610088565b91505092915050565b6100d3816100f4565b82525050565b60006020820190506100ee60008301846100ca565b92915050565b6000819050919050565b600080fd5b61010c816100f4565b811461011757600080fd5b5056fea2646970667358221220404e37f487a89a932dca5e77faaf6ca2de3b991f93d230604b1b8daaef64766264736f6c63430008070033"
        );
    }
}
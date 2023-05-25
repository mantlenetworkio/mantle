// SPDX-License-Identifier: Unlicense
pragma solidity >=0.8.0;

import {Vm} from "forge-std/Vm.sol";
import {DSTest} from "ds-test/test.sol";
import {console} from "forge-std/console.sol";
import {ProxyTester} from "./foundry-upgrades/ProxyTester.sol";

import '../../contracts/L1/tss/TssGroupManager.sol';

import '../../contracts/L1/tss/delegation/TssDelegation.sol';
import '../../contracts/L1/tss/delegation/TssDelegationManager.sol';
import '../../contracts/L1/tss/TssStakingSlashing.sol';

/*
TODO:
- Create a test to showcase the library
- Go over the cheatcodes + OZ docs and think of tests you can add into the code
- Add cheatcodes for added functionality (e.g storage slot) into the tester
*/

contract UpgradeTest is DSTest {
    ProxyTester proxy;

    address proxyAddress;

    address admin;

    TssGroupManager tssG;

    Vm constant vm =
        Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    function setUp() public {
        proxy = new ProxyTester();

        admin = vm.addr(69);

        tssG = new TssGroupManager();
    }

    function testDeployUUPS() public {
        proxy.setType("uups");
        proxyAddress = proxy.deploy(address(tssG), admin);
        assertEq(proxyAddress, proxy.proxyAddress());
        assertEq(proxyAddress, address(proxy.uups()));
        bytes32 tssGSlot = bytes32(
            uint256(keccak256("eip1967.proxy.implementation")) - 1
        );
        bytes32 proxySlot = vm.load(proxyAddress, tssGSlot);
        address addr;
        assembly {
            mstore(0, proxySlot)
            addr := mload(0)
        }
        assertEq(address(tssG), addr);

        console.log("proxy address: ", address(proxy));
        console.log("tssG address: ", address(tssG));
        console.log("tssG owner: ", tssG.owner());
    }

    function testUpgradeUUPS() public {
        testDeployUUPS();
        TssGroupManager newTssG = new TssGroupManager();
        /// Since the admin is an EOA, it doesn't have an owner
        proxy.upgrade(address(newTssG), admin, address(0));
        bytes32 tssGSlot = bytes32(
            uint256(keccak256("eip1967.proxy.implementation")) - 1
        );
        bytes32 proxySlot = vm.load(proxyAddress, tssGSlot);
        address addr;
        assembly {
            mstore(0, proxySlot)
            addr := mload(0)
        }
        assertEq(address(newTssG), addr);
    }

    function testDeployBeacon() public {
        proxy.setType("beaconProxy");
        // I will need an extra ProxyTester to become the beacon
        ProxyTester beaconTester = new ProxyTester();
        beaconTester.setType("beacon");
        beaconTester.deploy(address(tssG));
        proxy.deploy(address(beaconTester.beacon()));
        assertEq(address(tssG), beaconTester.beacon().implementation());
        bytes32 beaconSlot = bytes32(
            uint256(keccak256("eip1967.proxy.beacon")) - 1
        );
        bytes32 proxySlot = vm.load(proxy.proxyAddress(), beaconSlot);
        address addr;
        assembly {
            mstore(0, proxySlot)
            addr := mload(0)
        }
        assertEq(addr, beaconTester.beaconAddress());
    }
}

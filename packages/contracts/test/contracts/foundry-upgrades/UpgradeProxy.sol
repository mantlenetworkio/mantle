// SPDX-License-Identifier: Unlicense
pragma solidity >=0.8.0;

import {DeployProxy} from "./DeployProxy.sol";
import {ProxyAdminInterface} from "./IProxyAdmin.sol";

contract UpgradeProxy is DeployProxy {
    /// @notice Emitted when a proxy is upgraded. We define it here so that we can use
    /// the cheatcode and "expect" it.
    event Upgraded(address indexed implementation);

    /// @notice Upgrade a proxy smart contract without passing any calldata.
    function upgrade(
        address newImplementation,
        address admin,
        address owner
    ) public {
        // Load the expectEmit cheatcode for the Upgraded event. Since it's only a single argument, we only need
        // the first flag set to True.
        vm.expectEmit(true, false, false, false);
        emit Upgraded(newImplementation);
        // Check if the Admin address is a smart contract Admin or an EOA.
        // If it's an admin, the Upgrade will need to pass through the Admin, called by the owner EOA of the admin smart
        // contract.
        // Else, the upgrade function is called directly on the proxy, with the admin making the call.
        // vm.prank tells the Foundry VM to make the call from that particular address.
        if (isContract(admin)) {
            vm.prank(owner);
            ProxyAdminInterface(admin).upgrade(uups, newImplementation);
        } else {
            if (proxyType == ProxyType.UUPS) {
                vm.prank(admin);
                uups.upgradeTo(newImplementation);
            } else if (proxyType == ProxyType.Beacon) {
                vm.prank(admin);
                beacon.upgradeTo(newImplementation);
            } else if (proxyType == ProxyType.Transparent) {
                revert(
                    "Transparent ERC1967 proxies do not have upgradeable implementations"
                );
            }
        }
    }

    /// @notice Upgrade a proxy smart contract and also pass calldata to be called with the update.
    function upgrade(
        ProxyType proxy,
        address newImplementation,
        bytes memory data,
        address admin,
        address owner
    ) public {
        vm.expectEmit(true, false, false, false);
        emit Upgraded(newImplementation);
        if (isContract(admin)) {
            vm.prank(owner);
            ProxyAdminInterface(admin).upgradeAndCall(
                uups,
                newImplementation,
                data
            );
        } else {
            vm.prank(admin);
            uups.upgradeToAndCall(newImplementation, data);
        }
    }

    function upgradeBeacon(
        address newBeacon,
        bytes memory data,
        bool forceCall
    ) public {
        beaconProxy.upgradeBeaconToAndCall(newBeacon, data, forceCall);
    }

    function isContract(address addr) public view returns (bool) {
        uint256 size;
        assembly {
            size := extcodesize(addr)
        }
        return size > 0;
    }
}

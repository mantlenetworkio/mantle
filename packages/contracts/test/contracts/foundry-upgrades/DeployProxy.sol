// SPDX-License-Identifier: Unlicense
pragma solidity >=0.8.0;

// Import OZ Proxy contracts
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {UpgradeableBeaconProxy} from "./UpgradeableBeaconProxy.sol";
import {UpgradeableBeacon} from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import {Vm} from "forge-std/Vm.sol";
import {console} from "forge-std/console.sol";

contract DeployProxy {
    /// Cheatcodes address
    Vm constant vm =
        Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    ProxyType public proxyType;

    address public proxyAddress;

    address public beaconAddress;

    ERC1967Proxy public erc1967;

    TransparentUpgradeableProxy public uups;

    UpgradeableBeacon public beacon;

    UpgradeableBeaconProxy public beaconProxy;

    enum ProxyType {
        UUPS,
        BeaconProxy,
        Beacon,
        Transparent
    }

    function deploy(address implementation, bytes memory data)
        public
        returns (address)
    {
        if (proxyType == ProxyType.Transparent) {
            revert("Transparent proxy returns a single address");
        } else if (proxyType == ProxyType.UUPS) {
            revert("UUPS proxies require an admin address");
        } else if (proxyType == ProxyType.BeaconProxy) {
            return deployBeaconProxy(implementation, data);
        } else if (proxyType == ProxyType.Beacon) {
            return deployBeacon(implementation);
        }
    }

    function deploy(address implementation) public returns (address) {
        if (proxyType == ProxyType.Transparent) {
            revert("Transparent proxy returns a single address");
        } else if (proxyType == ProxyType.UUPS) {
            revert("UUPS proxies require an admin address");
        } else if (proxyType == ProxyType.BeaconProxy) {
            bytes memory data;
            return deployBeaconProxy(implementation, data);
        } else if (proxyType == ProxyType.Beacon) {
            return deployBeacon(implementation);
        }
    }

    function deploy(
        address implementation,
        address admin,
        bytes memory data
    ) public returns (address) {
        if (proxyType == ProxyType.Transparent) {
            revert("proxy implementation does't include admin address");
        } else if (proxyType == ProxyType.UUPS) {
            return deployUupsProxy(implementation, admin, data);
        } else if (proxyType == ProxyType.Beacon) {
            revert("proxy implementation does't include admin address");
        }
    }

    function deploy(address implementation, address admin)
        public
        returns (address)
    {
        if (proxyType == ProxyType.Transparent) {
            revert("proxy implementation does't include admin address");
        } else if (proxyType == ProxyType.UUPS) {
            bytes memory data;
            return deployUupsProxy(implementation, admin, data);
        } else if (proxyType == ProxyType.Beacon) {
            revert("proxy implementation does't include admin address");
        }
    }

    function deployBeacon(address impl) public returns (address) {
        beacon = new UpgradeableBeacon(impl);
        beaconAddress = address(beacon);
        vm.label(address(beaconAddress), "Upgradeable Beacon");
        return beaconAddress;
    }

    function deployBeaconProxy(address _beacon, bytes memory data)
        public
        returns (address)
    {
        beaconProxy = new UpgradeableBeaconProxy(_beacon, data);
        proxyAddress = address(beaconProxy);
        vm.label(proxyAddress, "Beacon Proxy");
        return proxyAddress;
    }

    function deployErc1967Proxy(address implementation, bytes memory data)
        public
        returns (address)
    {
        erc1967 = new ERC1967Proxy(implementation, data);
        proxyAddress = address(erc1967);
        vm.label(proxyAddress, "ERC1967 Proxy");
        return proxyAddress;
    }

    function deployUupsProxy(
        address implementation,
        address admin,
        bytes memory data
    ) public returns (address) {
        console.log("admin: ", admin);
        uups = new TransparentUpgradeableProxy(implementation, admin, data);
        proxyAddress = address(uups);
        vm.label(proxyAddress, "UUPS Proxy");
        return proxyAddress;
    }
}

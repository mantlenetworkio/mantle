// SPDX-License-Identifier: Unlicense
pragma solidity >=0.8.0;

import {UpgradeProxy} from "./UpgradeProxy.sol";

contract ProxyTester is UpgradeProxy {
    bool public typeSet;

    constructor() {}

    function setType(string memory _proxyType) public {
        require(typeSet == false, "ProxyTester has a proxy type already");
        if (keccak256(bytes(_proxyType)) == keccak256(bytes("uups"))) {
            proxyType = ProxyType.UUPS;
        } else if (keccak256(bytes(_proxyType)) == keccak256(bytes("beacon"))) {
            proxyType = ProxyType.Beacon;
        } else if (
            keccak256(bytes(_proxyType)) == keccak256(bytes("beaconProxy"))
        ) {
            proxyType = ProxyType.BeaconProxy;
        } else if (
            keccak256(bytes(_proxyType)) == keccak256(bytes("transparent"))
        ) {
            proxyType = ProxyType.Transparent;
        }
    }
}

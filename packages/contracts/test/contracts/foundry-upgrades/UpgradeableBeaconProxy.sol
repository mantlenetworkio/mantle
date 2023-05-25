// SPDX-License-Identifier: Unlicense
pragma solidity >=0.8.0;

import {BeaconProxy} from "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";

contract UpgradeableBeaconProxy is BeaconProxy {
    constructor(address beacon, bytes memory data) BeaconProxy(beacon, data) {}

    function upgradeBeaconToAndCall(
        address newBeacon,
        bytes memory data,
        bool forceCall
    ) public {
        _upgradeBeaconToAndCall(newBeacon, data, forceCall);
    }
}

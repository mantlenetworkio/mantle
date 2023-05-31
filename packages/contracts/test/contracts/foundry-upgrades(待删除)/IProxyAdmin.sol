// SPDX-License-Identifier: Unlicense
pragma solidity >=0.8.0;

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

interface ProxyAdminInterface {
    function upgradeAndCall(TransparentUpgradeableProxy proxy, address implementation, bytes memory data) external;
    function upgrade(TransparentUpgradeableProxy proxy, address implementation) external;
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

abstract contract WhiteList is OwnableUpgradeable {
    modifier whitelistOnly(address checkAddr) {
        if (!whitelist[checkAddr]) {
            revert("NOT_IN_WHITELIST");
        }
        _;
    }

    mapping(address => bool) public whitelist;

    /**
     * @notice Add to whitelist
     */
    function addToWhitelist(address[] calldata toAddAddresses) external onlyOwner {
        for (uint i = 0; i < toAddAddresses.length; i++) {
            whitelist[toAddAddresses[i]] = true;
        }
    }

    /**
     * @notice Remove from whitelist
     */
    function removeFromWhitelist(address[] calldata toRemoveAddresses) external onlyOwner {
        for (uint i = 0; i < toRemoveAddresses.length; i++) {
            delete whitelist[toRemoveAddresses[i]];
        }
    }
}

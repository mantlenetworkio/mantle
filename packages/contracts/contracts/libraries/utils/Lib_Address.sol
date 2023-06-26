// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/**
 * @title Lib_Address
 * @dev This library for convert bytes publicKey to address
 */
library Lib_Address {

    function publicKeyToAddress(bytes memory publicKey) internal pure returns (address) {
        require(publicKey.length == 64, "public key length must 64 bytes");
        return address(uint160(uint256(keccak256(publicKey))));
    }
}

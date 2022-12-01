// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/introspection/ERC165.sol";

/**
 * @title Technically EIP165-compliant, returns true for every interface.
 * @author Layr Labs, Inc.
 */
contract ERC165_Universal is IERC165 {
    function supportsInterface(bytes4) external pure returns (bool) {
        return true;
    }
}

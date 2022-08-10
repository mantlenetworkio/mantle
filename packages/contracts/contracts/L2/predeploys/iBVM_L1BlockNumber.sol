// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/**
 * @title iBVM_L1BlockNumber
 */
interface iBVM_L1BlockNumber {
    /********************
     * Public Functions *
     ********************/

    function getL1BlockNumber() external view returns (uint256);
}

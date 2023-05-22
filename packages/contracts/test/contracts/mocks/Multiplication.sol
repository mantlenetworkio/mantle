// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

contract Multiplication {
    constructor() {}

    function multiply() public view {
        uint256[3] memory coors;
        coors[0] = 1;
        coors[1] = 2;
        //this is the coefficient of the term with degree degree
        coors[2] = 255;
        uint256[2] memory product;
        assembly {
            if iszero(staticcall(not(0), 0x07, coors, 0x60, product, 0x40)) { revert(0, 0) }
        }
        // emit log_uint(product[0]);
        // emit log_uint(product[1]);
    }
}

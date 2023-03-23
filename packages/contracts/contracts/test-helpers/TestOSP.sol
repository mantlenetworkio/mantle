// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {TestERC20 } from "./TestERC20.sol";

contract TestOSP {
    TestERC20 public token;
    constructor(address _token) {
        token = TestERC20(_token);
    }

    function mint(uint256 amount) public {
        token.mint(msg.sender,amount);
    }

    function newToken() public{
        token = new TestERC20();
        token.mint(msg.sender,1000);
    }
}

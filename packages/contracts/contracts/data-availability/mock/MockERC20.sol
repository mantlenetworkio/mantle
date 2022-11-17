// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockERC20 is ERC20 {
    constructor(string memory _token, string memory _name, uint256 amount) ERC20(_token, _name) {
        _mint(msg.sender, amount);
    }
}
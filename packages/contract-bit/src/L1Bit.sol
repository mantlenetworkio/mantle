// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract L1Bit is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _mint(msg.sender, 1000000 * 10**18);
    }

    function mint(address to, uint256 amount) public virtual {
        /**
        ignore for testnet mint
        require(
            hasRole(MINTER_ROLE, _msgSender()),
            "TESTERC20MinterBurnerDecimals: must have minter role to mint"
        );
        */
        _mint(to, amount);
    }
}

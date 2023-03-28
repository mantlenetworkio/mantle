// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/* Library Imports */
import { Predeploys } from "../libraries/Predeploys.sol";

/* Contract Imports */
import { MantleMintableERC20 } from "../universal/MantleMintableERC20.sol";

/**
 * @title BVM_ETH
 * @dev The ETH predeploy provides an ERC20 interface for ETH deposited to Layer 2. Note that
 * unlike on Layer 1, Layer 2 accounts do not have a balance field.
 */
contract BVM_ETH is MantleMintableERC20 {
    /***************
     * Constructor *
     ***************/

    constructor()
        MantleMintableERC20(Predeploys.L2_STANDARD_BRIDGE, address(0), "Ether", "WETH")
    {}
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/* Library Imports */
import { Lib_PredeployAddresses } from "../../libraries/constants/Lib_PredeployAddresses.sol";

/* Contract Imports */
import { L2StandardERC20 } from "../../standards/L2StandardERC20.sol";

/**
 * @title BVM_ETH
 * @dev The ETH predeploy provides an ERC20 interface for ETH deposited to Layer 2. Note that
 * unlike on Layer 1, Layer 2 accounts do not have a balance field.
 */
contract BVM_ETH is L2StandardERC20 {
    /***************
     * Constructor *
     ***************/

    constructor()
        L2StandardERC20(Lib_PredeployAddresses.L2_STANDARD_BRIDGE, address(0), "Ether", "WETH", 18)
    {}
}

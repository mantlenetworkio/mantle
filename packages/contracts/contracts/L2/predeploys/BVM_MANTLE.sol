// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/* Library Imports */
import { Lib_PredeployAddresses } from "../../libraries/constants/Lib_PredeployAddresses.sol";

/* Contract Imports */
import { L2StandardERC20 } from "../../standards/L2StandardERC20.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";


/**
 * @title BVM_MANTLE
 * @dev The ETH predeploy provides an ERC20 interface for ETH deposited to Layer 2. Note that
 * unlike on Layer 1, Layer 2 accounts do not have a balance field.
 */
contract BVM_MANTLE is L2StandardERC20 {
    /***************
     * Constructor *
     ***************/
    // hardcode to mantle token mainnet address
    constructor()
        L2StandardERC20(Lib_PredeployAddresses.L2_STANDARD_BRIDGE, address(0x3c3a81e81dc49A522A592e7622A7E711c06bf354), "Mantle", "MNT", 18)
    {}


    function transfer(address recipient, uint256 amount) public virtual override(IERC20) returns (bool) {
        revert("BVM_MANTLE: transfer is disabled pending further community discussion.");
    }

    function approve(address spender, uint256 amount) public virtual override(IERC20) returns (bool) {
        revert("BVM_MANTLE: approve is disabled pending further community discussion.");
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public virtual override(IERC20) returns (bool) {
        revert("BVM_MANTLE: transferFrom is disabled pending further community discussion.");
    }

    function increaseAllowance(address spender, uint256 addedValue)
        public
        virtual
        override
        returns (bool)
    {
        revert("BVM_MANTLE: increaseAllowance is disabled pending further community discussion.");
    }

    function decreaseAllowance(address spender, uint256 subtractedValue)
        public
        virtual
        override
        returns (bool)
    {
        revert("BVM_MANTLE: decreaseAllowance is disabled pending further community discussion.");
    }
}

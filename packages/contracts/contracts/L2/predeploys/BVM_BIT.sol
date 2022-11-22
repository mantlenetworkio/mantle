// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/* Library Imports */
import { Lib_PredeployAddresses } from "../../libraries/constants/Lib_PredeployAddresses.sol";

/* Contract Imports */
import { L2StandardERC20 } from "../../standards/L2StandardERC20.sol";

/**
 * @title BVM_BIT
 * @dev The ETH predeploy provides an ERC20 interface for ETH deposited to Layer 2. Note that
 * unlike on Layer 1, Layer 2 accounts do not have a balance field.
 */
contract BVM_BIT is L2StandardERC20 {
    /***************
     * Constructor *
     ***************/
    // hardcode to bit token mainnet address
    constructor()
        L2StandardERC20(Lib_PredeployAddresses.L2_STANDARD_BRIDGE, address(0xdf3BD218A936A92be5e43592143ecc7a33cef514), "Bit Token", "BIT", 18)
    {}


    function transfer(address recipient, uint256 amount) public virtual override returns (bool) {
        revert("BVM_BIT: transfer is disabled pending further community discussion.");
    }

    function approve(address spender, uint256 amount) public virtual override returns (bool) {
        revert("BVM_BIT: approve is disabled pending further community discussion.");
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public virtual override returns (bool) {
        revert("BVM_BIT: transferFrom is disabled pending further community discussion.");
    }

    function increaseAllowance(address spender, uint256 addedValue)
        public
        virtual
        override
        returns (bool)
    {
        revert("BVM_BIT: increaseAllowance is disabled pending further community discussion.");
    }

    function decreaseAllowance(address spender, uint256 subtractedValue)
        public
        virtual
        override
        returns (bool)
    {
        revert("BVM_BIT: decreaseAllowance is disabled pending further community discussion.");
    }
}

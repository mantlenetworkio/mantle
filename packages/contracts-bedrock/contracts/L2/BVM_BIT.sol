// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/* Library Imports */
import { Predeploys } from "../libraries/Predeploys.sol";

/* Contract Imports */
import { MantleMintableERC20 } from "../universal/MantleMintableERC20.sol";

/**
 * @title BVM_BIT
 * @dev The ETH predeploy provides an ERC20 interface for ETH deposited to Layer 2. Note that
 * unlike on Layer 1, Layer 2 accounts do not have a balance field.
 */
contract BVM_BIT is MantleMintableERC20 {
    /***************
     * Constructor *
     ***************/
    // hardcode to bit token mainnet address
    constructor()
        MantleMintableERC20(Predeploys.L2_STANDARD_BRIDGE,address(0x1A4b46696b2bB4794Eb3D4c26f1c55F9170fa4C5),"Bit Token","BIT")
    {
    }


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

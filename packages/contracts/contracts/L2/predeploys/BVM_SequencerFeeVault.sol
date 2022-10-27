// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/* Library Imports */
import { Lib_PredeployAddresses } from "../../libraries/constants/Lib_PredeployAddresses.sol";

/* Contract Imports */
import { L2StandardBridge } from "../messaging/L2StandardBridge.sol";
import { IBVM_SequencerFeeVault } from "./iBVM_SequencerFeeVault.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title BVM_SequencerFeeVault
 * @dev Simple holding contract for fees paid to the Sequencer. Likely to be replaced in the future
 * but "good enough for now".
 */
contract BVM_SequencerFeeVault is IBVM_SequencerFeeVaultis ,Ownable{
    /*************
     * Constants *
     *************/

    // Minimum ETH balance that can be withdrawn in a single withdrawal.
    uint256 public constant MIN_WITHDRAWAL_AMOUNT = 15 ether;

    /*************
     * Variables *
     *************/

    // Address on L1 that will hold the fees once withdrawn. Dynamically initialized within l2geth.
    address public l1FeeWallet;

    uint256 public constant L1Gas = 200_000;

    /***************
     * Constructor *
     ***************/

    /**
     * @param _l1FeeWallet Initial address for the L1 wallet that will hold fees once withdrawn.
     * Currently HAS NO EFFECT in production because l2geth will mutate this storage slot during
     * the genesis block. This is ONLY for testing purposes.
     */
    constructor(address _l1FeeWallet, address _owner) {
        l1FeeWallet = _l1FeeWallet;
        transferOwnership(_owner);
    }

    /************
     * Fallback *
     ************/

    // slither-disable-next-line locked-ether
    receive() external payable {}

    /********************
     * Public Functions *
     ********************/
    function setL1FeeWallet(address _l1FeeWallet) public onlyOwner {
        l1FeeWallet = _l1FeeWallet;
    }
    function l1FeeWallet() external view  returns (address){
        return l1FeeWallet;
    }

// slither-disable-next-line external-function
    function withdraw() public {
        require(
            address(this).balance >= MIN_WITHDRAWAL_AMOUNT,
        // solhint-disable-next-line max-line-length
            "BVM_SequencerFeeVault: withdrawal amount must be greater than minimum withdrawal amount"
        );

        L2StandardBridge(Lib_PredeployAddresses.L2_STANDARD_BRIDGE).withdrawTo(
            Lib_PredeployAddresses.BVM_BIT,
            l1FeeWallet,
            address(this).balance,
            0,
            bytes("")
        );
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/* Library Imports */
import {Lib_PredeployAddresses} from "../../libraries/constants/Lib_PredeployAddresses.sol";

/* Contract Imports */
import {L2StandardBridge} from "../messaging/L2StandardBridge.sol";
import {IBVM_GasPriceOracle} from "./iBVM_GasPriceOracle.sol";

// TODO delete
import {ITssRewardContract} from "./iTssRewardContract.sol";

/**
 * @title BVM_SequencerFeeVault
 * @dev Simple holding contract for fees paid to the Sequencer. Likely to be replaced in the future
 * but "good enough for now".
 */
contract BVM_SequencerFeeVault {
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

    address public bvmGasPriceOracleAddress;

    uint256 public constant L1Gas = 200_000;

    /***************
     * Constructor *
     ***************/

    /**
     * @param _l1FeeWallet Initial address for the L1 wallet that will hold fees once withdrawn.
     * Currently HAS NO EFFECT in production because l2geth will mutate this storage slot during
     * the genesis block. This is ONLY for testing purposes.
     */
    constructor(address _l1FeeWallet, address _bvmGasPriceOracleAddress) {
        l1FeeWallet = _l1FeeWallet;
        bvmGasPriceOracleAddress = _bvmGasPriceOracleAddress;
    }

    /************
     * Fallback *
     ************/

    // slither-disable-next-line locked-ether
    receive() external payable {}

    /********************
     * Public Functions *
     ********************/

    // slither-disable-next-line external-function
    function withdraw() public {
        address to = l1FeeWallet;
        if (IBVM_GasPriceOracle(bvmGasPriceOracleAddress).IsBurning() == true) {
            to = address(0x000000000000000000000000000000000000dEaD);
            if (address(this).balance < MIN_WITHDRAWAL_AMOUNT) {
                return;
            }
        } else {
            require(
                address(this).balance >= MIN_WITHDRAWAL_AMOUNT,
            // solhint-disable-next-line max-line-length
                "BVM_SequencerFeeVault: withdrawal amount must be greater than minimum withdrawal amount"
            );
        }

        L2StandardBridge(Lib_PredeployAddresses.L2_STANDARD_BRIDGE).withdrawTo(
            Lib_PredeployAddresses.BVM_BIT,
            to,
            address(this).balance,
            0,
            bytes("")
        );
    }

    function tssClaim() public {
        address[] memory tssMembers = new address[](2);
        tssMembers[0] = address(0xCFc17379Ac80A9EF231772ACE60014fb84704cB4);
        tssMembers[1] = address(0x9D72b1e94C7075Be7E6da0E2104DB4302d02DB0E);

        bytes32[] memory _batch;
        uint256 _shouldStartAtElement;

        // construct calldata for claimReward call
        bytes memory message = abi.encodeWithSelector(
            ITssRewardContract.claimReward.selector,
            _shouldStartAtElement,
            250,
            block.timestamp,
            tssMembers
        );
        address target = address(0x4200000000000000000000000000000000000020);
        (bool success, ) = target.call(message);
        require(success == true,"call failed");
        // send call data into L2, hardcode address
//        sendCrossDomainMessage(
//            address(0x4200000000000000000000000000000000000020),
//            2000000,
//            message
//        );
    }
}

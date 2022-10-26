// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "./iTssRewardContract.sol";

/* Library Imports */
import {SafeMath} from "@openzeppelin/contracts/utils/math/SafeMath.sol";

/* Interface Imports */

/* External Imports */

/**
 * @title TssRewardContract
 * @dev Release to batch roll up tss members.
 */
contract TssRewardContract is ITssRewardContract {
    using SafeMath for uint256;

    mapping(uint256 => uint256) public ledger;
    address public deadAddress;
    address public owner;
    uint256 public dust;
    uint256 public totalAmount;
    uint256 public latsBatchTime;
    uint256 public sendAmountPerMinute;
    uint256 public sendAmountPerYear;

    // set call address
    constructor(address _deadAddress, address _owner) {
        deadAddress = _deadAddress;
        owner = _owner;
    }

    /**
     * Enforces that the modified function is only callable by a specific null address.
     *  authenticated to call this function.
     */
    modifier onlyFromDeadAddress() {
        require(
            msg.sender == deadAddress,
            "tss reward call message unauthenticated"
        );
        _;
    }

    modifier onlyOwner() {
        require(
            msg.sender == owner,
            "only be called by the owner of this contract"
        );
        _;
    }

    modifier checkBalance() {
        require(
            address(this).balance >= totalAmount,
            "balance record and contract balance are not equal"
        );
        _;
    }

    /**
     * @dev return the total undistributed amount
     */
    function queryReward() external view checkBalance returns (uint256) {
        return address(this).balance;
    }

    /**
     * @dev claimReward distribute reward to tss member.
     * @param _batchTime Batch corresponds to L1 Block Timestamp
     * @param _length The distribute batch block number
     * @param _tssMembers The address array of tss group members
     */
    function claimReward(uint256 _batchTime, address[] calldata _tssMembers)
    external
    virtual
    onlyFromDeadAddress
    checkBalance
    {
        //
        uint256 sendAmount = 0;
        uint256 batchAmount = 0;
        uint256 accu = 0;
        // sendAmount
        batchAmount = (_batchTime - latsBatchTime) * sendAmountPerMinute;
        sendAmount = batchAmount.div(_tssMembers.length);
        for (uint256 j = 0; j < _tssMembers.length; j++) {
            address payable addr = payable(_tssMembers[j]);
            accu = accu.add(sendAmount);
            totalAmount = totalAmount.sub(sendAmount);
            addr.transfer(sendAmount);
        }
        uint256 reserved = batchAmount.sub(accu);
        if (reserved > 0) {
            dust = dust.add(reserved);
        }
        emit DistributeTssReward(
            _batchTime,
            _tssMembers
        );
    }

    /**
     * @dev withdraw div dust
     */
    function withdrawDust() external onlyOwner checkBalance {
        uint256 amount = dust;
        totalAmount = totalAmount.sub(dust);
        dust = 0;
        if (amount > 0) {
            payable(owner).transfer(dust);
        }
    }

    /**
     * @dev clear balance
     */
    function withdraw() external onlyOwner checkBalance {
        totalAmount = 0;
        if (address(this).balance > 0) {
            payable(owner).transfer(address(this).balance);
        }
    }
}

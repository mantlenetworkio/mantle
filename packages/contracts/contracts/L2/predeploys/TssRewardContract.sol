// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "./iTssRewardContract.sol";

/* Library Imports */
import { SafeMath } from "@openzeppelin/contracts/utils/math/SafeMath.sol";

/* Interface Imports */

/* External Imports */

/**
 * @title TssRewardContract
 * @dev Collect L2 block gas reward per block and release to batch roll up tss members.
 */
contract TssRewardContract is ITssRewardContract {
    using SafeMath for uint256;

    mapping(uint256 => uint256) public ledger;
    address public deadAddress;
    address payable public owner;
    uint256 public bestBlockID;
    uint256 public dust;
    uint256 public totalAmount;

    // set call address
    constructor(address _deadAddress, address payable _owner) {
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
     * @param _blockStartHeight The block height at L2 which needs to distribute profits
     * @param _length The distribute batch block number
     * @param _tssMembers The address array of tss group members
     */
    function claimReward(uint256 _blockStartHeight, uint32 _length, address[] calldata _tssMembers)
        external
        virtual
        onlyFromDeadAddress
        checkBalance
    {
        uint256 blockAmount = ledger[_blockStartHeight];
        uint256 sendAmount = 0;
        uint256 accu = 0;
        // release reward from _blockStartHeight to _blockStartHeight + _length
        for (uint256 i=_blockStartHeight; i<_length; i++) {
            // calc send amount
            sendAmount = blockAmount.div(_tssMembers.length);
            // delete distributed height
            delete ledger[_blockStartHeight];
            for (uint256 j=0; j < _tssMembers.length; j++) {
                address payable addr = payable(_tssMembers[j]);
                accu = accu.add(sendAmount);
                totalAmount = totalAmount.sub(sendAmount);
                addr.transfer(sendAmount); // TODO use call to transfer
            }
            uint256 reserved = blockAmount.sub(accu);
            require(reserved >= 0, "release amount gt real balance");
            if (reserved > 0) {
                dust = dust.add(reserved);
            }
        }

        emit DistributeTssReward(
            _blockStartHeight,
            _length,
            _tssMembers
        );
    }

    /**
     * @dev update tss member gas reward by every block.
     * @param _blockID The block height at L2 which needs to distribute profits
     * @param _amount Distribute batch block number
     * @return _tssMembers Address array of tss group members
     */
    function updateReward(uint256 _blockID)
        external
        payable
        onlyFromDeadAddress
        checkBalance
        returns (bool)
    {
        // check update block ID
        require(_blockID == bestBlockID+1, "block id update illegal");
        // iter address to update balance
        bestBlockID = _blockID;
        totalAmount = totalAmount.add(msg.value);
        ledger[_blockID] = msg.value;
        return true;
    }

    function withdrawDust() external onlyOwner checkBalance {
        uint256 amount = dust;
        totalAmount = totalAmount.sub(dust);
        dust = 0;
        if (amount > 0) {
            owner.transfer(dust);
        }
    }

    function withdraw() external onlyOwner checkBalance {
        totalAmount = 0;
        if (address(this).balance > 0) {
            owner.transfer(address(this).balance);
        }
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/* Library Imports */

/* Interface Imports */
import { ITssRewardContract } from "./ITssRewardContract.sol";

/**
 * @title TssRewardContract
 * @dev Collect L2 block gas reward per block and release to batch roll up tss members.
 */
contract TssRewardContract is ITssRewardContract {
    mapping(uint256 => uint256) public ledger;
    address public deadAddress;
    uint256 public bestBlockID = 0;
    uint256 public dust = 0;

    // set call address
    constructor(uint256 _deadAddress) {
        deadAddress = _deadAddress;
    }

    /**
     * Enforces that the modified function is only callable by a specific cross-domain account.
     * @param _sourceDomainAccount The only account on the originating domain which is
     *  authenticated to call this function.
     */
    modifier onlyFromDeadAddress() {
        require(
            msg.sender == deadAddress,
            "tss reward call message unauthenticated"
        );
        _;
    }

    /**
     * @dev return the total undistributed amount
     */
    function queryReward() external view returns (uint256) {
        return address(this).balance;
    }

    /**
     * @dev claimReward distribute reward to tss member.
     * @param _blockStartHeight The block height at L2 which needs to distribute profits
     * @param _length The distribute batch block number
     * @param _tssMembers The address array of tss group members
     */
    function claimReward(uint256 _blockStartHeight, uint32 _length, address[] _tssMembers)
        external
        virtual
        onlyFromDeadAddress
    {
        uint256 totalAmount = ledger[blockStartHeight];
        uint256 sendAmount = 0;
        uint256 accu = 0;
        // release reward from _blockStartHeight to _blockStartHeight + _length
        for (i=_blockStartHeight; i<_length; i++) {
            // calc send amount
            sendAmount = amount / _tssMembers.length; // safemath
            // delete distributed height
            delete ledger[blockStartHeight];
            for (j=0; j < _tssMembers.length; j++) {
                accu += sendAmount;
                owner.transfer(sendAmount);
            }
            uint256 reserved = totalAmount - accu;
            require(reserved >= 0, "release amount gt real balance");
            if (reserved > 0) {
                dust += reserved;
            }
        }

        emit DistributeTssReward(
            _blockStartHeight,
            _length,
            tssMembers
        );
    }

    /**
     * @dev update tss member gas reward by every block.
     * @param _blockID The block height at L2 which needs to distribute profits
     * @param _amount Distribute batch block number
     * @return _tssMembers Address array of tss group members
     */
    function updateTssReward(uint256 _blockID, uint256 _amount)
        external
        onlyFromDeadAddress
        returns (bool)
    {
        // TODO check balance, require ledger sum balance match to contract global balance

        // check update block ID
        require(_blockID == bestBlockID+1, "block id update illegal");
        // iter address to update balance
        ledger[_blockID] = _amount;
        bestBlockID = _blockID;
        return true;
    }

    function withdrawDust() external {
        require(msg.sender == owner);
        owner.transfer(dust);
    }

    function withdraw() external {
        require(msg.sender == owner);
        owner.transfer(address(this).balance);
    }
}

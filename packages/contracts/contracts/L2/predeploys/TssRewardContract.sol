// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ITssRewardContract} from  "./iTssRewardContract.sol";
import {IBVM_GasPriceOracle} from "./iBVM_GasPriceOracle.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Lib_PredeployAddresses} from "../../libraries/constants/Lib_PredeployAddresses.sol";
import { CrossDomainEnabled } from "../../libraries/bridge/CrossDomainEnabled.sol";

/* Library Imports */
import {SafeMath} from "@openzeppelin/contracts/utils/math/SafeMath.sol";

/* Interface Imports */

/* External Imports */

/**
 * @title TssRewardContract
 * @dev Release to batch roll up tss members.
 */
contract TssRewardContract is Ownable,ITssRewardContract,CrossDomainEnabled {
    using SafeMath for uint256;

    mapping(uint256 => uint256) public ledger;
    address public bvmGasPriceOracleAddress;
    address public deadAddress;
    uint256 public dust;
    uint256 public bestBlockID;
    uint256 public totalAmount;
    uint256 public lastBatchTime;
    uint256 public sendAmountPerYear;
    address public sccAddress;

    uint256 public dustBlock;
    uint256 public waitingTime;
    // operator => tssreward
    mapping(address => uint256) public rewardDetails;
    //operator => claimer
    mapping(address => address) public operators;
    //claimer => operator
    mapping(address => address) public claimers;
    //operator => timestamp
    mapping(address => uint256) public claimTimes;
    //operator => claim number
    mapping(address => uint256) public claimAmout;
    address public stakeSlashAddress;


    // set call address
    constructor(address _deadAddress, address _owner, uint256 _sendAmountPerYear, address _bvmGasPriceOracleAddress,address _l2CrossDomainMessenger, address _sccAddress, uint256 _waitingTime, address _ssAddr)
    Ownable() CrossDomainEnabled(_l2CrossDomainMessenger)
    {
        transferOwnership(_owner);
        deadAddress = _deadAddress;
        sendAmountPerYear = _sendAmountPerYear;
        bvmGasPriceOracleAddress = _bvmGasPriceOracleAddress;
        sccAddress = _sccAddress;
        waitingTime = _waitingTime;
        stakeSlashAddress = _sccAddress;
    }

    // slither-disable-next-line locked-ether
    receive() external payable {}

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

    modifier checkBalance() {
        require(
            address(this).balance >= totalAmount,
            "balance record and contract balance are not equal"
        );
        _;
    }

    modifier onlyAuthorized() {
        address operator = claimers[msg.sender];
        require(operator != address(0),
        "The msg sender is not authorized by the validator"
        );
        _;
    }

    function setSendAmountPerYear(uint256 _sendAmountPerYear) public onlyOwner {
        sendAmountPerYear = _sendAmountPerYear;
    }

    function setWaitingTime(uint256 _waitingTime) public onlyOwner {
        waitingTime = _waitingTime;
    }

    function setSccAddr(address _sccAddress) public onlyOwner {
        sccAddress = _sccAddress;
    }

    function setStakeSlashAddr(address _ssAddre) public onlyOwner {
        stakeSlashAddress = _ssAddre;
    }

    function querySendAmountPerSecond() public view returns (uint256){
        return (sendAmountPerYear * 10 ** 18).div(365 * 24 * 60 * 60);
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
    function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] calldata _tssMembers)
    external
    virtual
    onlyFromCrossDomainAccount(sccAddress)
    {
        if (IBVM_GasPriceOracle(bvmGasPriceOracleAddress).IsBurning() != 1) {
            claimRewardByBlock(_blockStartHeight, _length, _tssMembers);
            return;
        }
        uint256 batchAmount = 0;
        // sendAmount
        if (lastBatchTime == 0) {
            lastBatchTime = _batchTime;
            return;
        }
        require(_batchTime > lastBatchTime,"args _batchTime must gther than last lastBatchTime");
        batchAmount = (_batchTime - lastBatchTime) * querySendAmountPerSecond() + dust;
        dust = 0;
        _distributeReward(batchAmount, _tssMembers, false);
        emit DistributeTssReward(
            lastBatchTime,
            _batchTime,
            batchAmount,
            _tssMembers
        );
        lastBatchTime = _batchTime;
    }

    /**
     * @dev claimReward distribute reward to tss member.
     * @param _blockStartHeight The block height at L2 which needs to distribute profits
     * @param _length The distribute batch block number
     * @param _tssMembers The address array of tss group members
     */
    function claimRewardByBlock(uint256 _blockStartHeight, uint32 _length, address[] calldata _tssMembers)
    internal
    {
        uint256 sendAmount = 0;
        uint256 batchAmount = 0;
        uint256 accu = 0;
        // release reward from _blockStartHeight to _blockStartHeight + _length - 1
        for (uint i = 0; i < _length; i++) {
            batchAmount = batchAmount.add(ledger[_blockStartHeight + i]);
            // delete distributed height
            delete ledger[_blockStartHeight + i];
        }
        if (batchAmount > 0) {
            batchAmount = batchAmount + dustBlock;
            dustBlock = 0;
            _distributeReward(batchAmount, _tssMembers, true);
        }
        emit DistributeTssRewardByBlock(
            _blockStartHeight,
            _length,
            sendAmount,
            _tssMembers
        );
    }

    /**
     * @dev update tss member gas reward by every block.
     * @param _blockID The block height at L2 which needs to distribute profits
     * @return _tssMembers Address array of tss group members
     */
    function updateReward(uint256 _blockID, uint256 _amount)
    external
    onlyFromDeadAddress
    checkBalance
    returns (bool)
    {
        // check update block ID
        require(_blockID == bestBlockID + 1, "block id update illegal");
        // iter address to update balance
        bestBlockID = _blockID;
        totalAmount = totalAmount.add(_amount);
        ledger[_blockID] = _amount;
        return true;
    }

    /**
     * @dev withdraw div dustBlock
     */
    function withdrawDust() external onlyOwner checkBalance {
        uint256 amount = dustBlock;
        totalAmount = totalAmount.sub(amount);
        dustBlock = 0;
        if (amount > 0) {
            payable(owner()).transfer(amount);
        }
    }

    /**
     * @dev clear balance
     */
    function withdraw() external onlyOwner checkBalance {
        totalAmount = 0;
        if (address(this).balance > 0) {
            payable(owner()).transfer(address(this).balance);
        }
    }

    function requestClaim() external onlyAuthorized returns (bool) {
        address operator = claimers[msg.sender];
        uint256 time = claimTimes[operator];
        require(time == 0,
        "You have already initiated a request to claim, please wait for the waiting period to pass"
        );
        claimTimes[operator] = block.timestamp;
        claimAmout[operator] = rewardDetails[operator];
        return true;
    }

    function queryClaimTime() external view onlyAuthorized returns (uint256) {
        address operator = claimers[msg.sender];
        uint256 remainTime = _remainTime(operator);
        return remainTime;
    }

    /**
     * @dev Claim reward
     */
    function claim() external onlyAuthorized {
        address operator = claimers[msg.sender];
        uint256 remainTime = _remainTime(operator);
        require(remainTime == 0,
        "please wait for the waiting period to pass"
        );
        _claim(operator);
        delete claimTimes[operator];
    }

    function setClaimer(address _operator, address _claimer)
    external
    virtual
    onlyFromCrossDomainAccount(stakeSlashAddress)
    {
        claimers[_claimer] = _operator;
        operators[_operator] = _claimer;
    }

    function _remainTime(address operator) internal view returns (uint256) {
        uint256 time = claimTimes[operator];
        require(time != 0,
        "please initiate a request to claim first"
        );
        uint256 last = time + waitingTime;
        uint256 remaining;
        if ( last > block.timestamp ) {
            remaining = last - block.timestamp;
        }else {
            remaining = 0;
        }
        return remaining;
    }


    function _claim(address _operator) internal {
        uint256 claimNumber = claimAmout[_operator];
        uint256 amount = rewardDetails[_operator];
        if (claimNumber > 0) {
            address claimer = operators[_operator];
            address payable addr = payable(claimer);
            addr.transfer(claimNumber);
            delete claimAmout[_operator];
            rewardDetails[_operator] = rewardDetails[_operator] - claimNumber;
        }
        emit Claim(_operator, amount);
    }

    function _distributeReward(uint256 amount, address[] calldata _tssMembers, bool isByBlock) internal {
        if (amount > 0) {
            uint256 sendAmount = 0;
            uint256 accu = 0;
            sendAmount = amount.div(_tssMembers.length);
            for (uint i = 0; i < _tssMembers.length; i++) {
                address operator = _tssMembers[i];
                rewardDetails[operator] += sendAmount;
                accu = accu.add(sendAmount);
                if (isByBlock) {
                    totalAmount = totalAmount.sub(sendAmount);
                }
            }
            uint256 reserved = amount.sub(accu);
            if (reserved > 0) {
                if (isByBlock) {
                    dustBlock = dustBlock.add(reserved);
                }else {
                    dust = dust.add(reserved);
                }

            }
        }
    }


}

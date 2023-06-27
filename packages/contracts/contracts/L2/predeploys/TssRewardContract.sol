// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

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
contract TssRewardContract is Ownable,ITssRewardContract,CrossDomainEnabled,ReentrancyGuardUpgradeable {
    using SafeMath for uint256;

    uint256 public lastBatchTime;
    uint256 public sendAmountPerYear;

    address public sccAddress;
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
    constructor(address _owner, uint256 _sendAmountPerYear,address _l2CrossDomainMessenger, address _sccAddress, uint256 _waitingTime, address _ssAddr)
    Ownable() CrossDomainEnabled(_l2CrossDomainMessenger)
    {
        transferOwnership(_owner);
        sendAmountPerYear = _sendAmountPerYear;
        sccAddress = _sccAddress;
        waitingTime = _waitingTime;
        stakeSlashAddress = _sccAddress;
    }

    // slither-disable-next-line locked-ether
    receive() external payable {}

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
    function queryReward() external view returns (uint256) {
        return address(this).balance;
    }

    /**
     * @dev claimReward distribute reward to tss member.
     * @param _blockStartHeight L2 rollup batch block start height.
     * @param _batchTime Batch corresponds to L1 Block Timestamp
     * @param _length The distribute batch block number
     * @param _tssMembers The address array of tss group members
     */
    function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] calldata _tssMembers)
    external
    virtual
    onlyFromCrossDomainAccount(sccAddress)
    {
        uint256 batchAmount = 0;
        // sendAmount
        if (lastBatchTime == 0) {
            lastBatchTime = _batchTime;
            return;
        }
        require(_batchTime > lastBatchTime,"args _batchTime must gther than last lastBatchTime");
        batchAmount = (_batchTime - lastBatchTime) * querySendAmountPerSecond();
        _distributeReward(batchAmount, _tssMembers);
        emit DistributeTssReward(
            lastBatchTime,
            _batchTime,
            batchAmount,
            _tssMembers
        );
        lastBatchTime = _batchTime;
    }

    /**
     * @dev clear balance
     */
    function withdraw() external onlyOwner {
        if (address(this).balance > 0) {
            (bool success, ) = owner().call{ value: address(this).balance }(new bytes(0));
            require(success, "TssReward withdraw: MNT transfer failed");
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

    function queryClaimTime() external view returns (uint256) {
        address operator = claimers[msg.sender];
        uint256 remainTime = _remainTime(operator);
        return remainTime;
    }

    /**
     * @dev Claim reward
     */
    function claim() external onlyAuthorized nonReentrant {
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
        if (operators[_operator] != address(0)) {
            delete claimers[operators[_operator]];
        }
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
        require(amount >= claimNumber,"The numerical value is incorrect");
        require(address(this).balance >= claimNumber,"The contract balance is insufficient to pay the reward value");
        if (claimNumber > 0) {
            address claimer = operators[_operator];
            delete claimAmout[_operator];
            rewardDetails[_operator] = rewardDetails[_operator] - claimNumber;
            (bool success, ) = claimer.call{ value: claimNumber }(new bytes(0));
            require(success, "TssReward claim: MNT transfer failed");
            emit Claim(_operator, claimNumber);
        }
    }


    function _distributeReward(uint256 amount, address[] calldata _tssMembers) internal {
        if (amount > 0) {
            for (uint i = 0; i < _tssMembers.length; i++) {
                address operator = _tssMembers[i];
                rewardDetails[operator] += amount;
            }
        }
    }

}

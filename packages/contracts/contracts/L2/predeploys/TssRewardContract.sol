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
    // staker => operator
    mapping(address => address) public delegatedTo;
    // operator => investment strategy => total number of shares delegated to them
    mapping(address => mapping(address => uint256)) public operatorStakerShares;
    //operator => stakers
    mapping(address => address[]) public operatorStakers;
    // operator => total number of shares
    mapping(address => uint256) public operatorShares;
    // operator or staker => tssreward
    mapping(address => uint256) public rewardDetails;
    //claimer => staker
    mapping(address => address) public claimers;



    // set call address
    constructor(address _deadAddress, address _owner, uint256 _sendAmountPerYear, address _bvmGasPriceOracleAddress,address _l2CrossDomainMessenger, address _sccAddress)
    Ownable() CrossDomainEnabled(_l2CrossDomainMessenger)
    {
        transferOwnership(_owner);
        deadAddress = _deadAddress;
        sendAmountPerYear = _sendAmountPerYear;
        bvmGasPriceOracleAddress = _bvmGasPriceOracleAddress;
        sccAddress = _sccAddress;
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

    function setSendAmountPerYear(uint256 _sendAmountPerYear) public onlyOwner {
        sendAmountPerYear = _sendAmountPerYear;
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
            batchAmount = batchAmount + dust;
            dust = 0;
            sendAmount = batchAmount.div(_tssMembers.length);
            for (uint j = 0; j < _tssMembers.length; j++) {
                address payable addr = payable(_tssMembers[j]);
                accu = accu.add(sendAmount);
                totalAmount = totalAmount.sub(sendAmount);
                addr.transfer(sendAmount);
            }
            uint256 reserved = batchAmount.sub(accu);
            if (reserved > 0) {
                dust = dust.add(reserved);
            }
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
     * @dev withdraw div dust
     */
    function withdrawDust() external onlyOwner checkBalance {
        uint256 amount = dust;
        totalAmount = totalAmount.sub(dust);
        dust = 0;
        if (amount > 0) {
        payable(owner()).transfer(dust);
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

    /**
     * @dev Increases the `staker`'s delegated shares
     * @param _operator the address of operator which staker chosed
     * @param _staker the address of staker
     * @param _shares the number of staker delegated for operator
     */
    function increaseDelegatedShares(address _operator, address _staker, uint256 _shares)
    external
    virtual
    onlyFromCrossDomainAccount(sccAddress) {
        _increaseDelegation(_operator, _staker, _shares);
    }

    /**
     * @dev Decreases the `staker`'s delegated shares
     * @param _operator the address of operator which staker chosed
     * @param _staker the address of staker
     * @param _shares the number of staker delegated for operator
     */
    function decreaseDelegatedShares(address _operator, address _staker, uint256 _shares)
    external
    virtual
    onlyFromCrossDomainAccount(sccAddress) {
        _decreasDelegation(_operator, _staker, _shares);
    }

    /**
     * @dev first stake and delegated shares
     * @param _operator the address of operator which staker chosed
     * @param _staker the address of staker
     * @param _shares the number of staker delegated for operator
     */
    function delegate(address _operator, address _staker, uint256 _shares)
    external
    virtual
    onlyFromCrossDomainAccount(sccAddress) {
        //store staker => operator
        delegatedTo[_staker] = _operator;
        operatorStakers[_operator].push(_staker);
        claimers[_staker] = _staker;
        _increaseDelegation(_operator, _staker, _shares);
    }

    /**
     * @dev Claim reward
     * @param _addr the address of reward owner
     */
    function claim(address _addr) external {
        _claim(_addr);
    }

    /**
     * @dev Claim reward and withdraw
     * @param _addr the address of reward owner
     */
    function claimWithdraw(address _addr) external {
        _claim(_addr);
        address staker = claimers[_addr];
        delete rewardDetails[staker];
        address operator = delegatedTo[staker];
        uint256 shares = operatorStakerShares[operator][staker];
        operatorShares[operator] -= shares;
        delete operatorStakerShares[operator][staker];
        delete delegatedTo[staker];
        delete claimers[_addr];
        _deleteStakers(operator, staker);
    }

    function setClaimer(address _staker, address _claimer)
    external
    virtual
    onlyFromCrossDomainAccount(sccAddress)
    {
        claimers[_claimer] = _staker;
    }



    function _claim(address _addr) internal {
        address staker = claimers[_addr];
        uint256 amount = rewardDetails[staker];
        if (amount > 0) {
            address payable addr = payable(_addr);
            addr.transfer(amount);
            rewardDetails[staker] = 0;
        }
        emit Claim(_addr, amount);
    }

    function _increaseDelegation(address _operator, address _staker, uint256 _shares) internal {
        //store operator => staker =>shares
        operatorStakerShares[_operator][_staker] += _shares;
        operatorShares[_operator] += _shares;
    }

    function _decreasDelegation(address _operator, address _staker, uint256 _shares) internal {
        operatorStakerShares[_operator][_staker] -= _shares;
        operatorShares[_operator] -= _shares;
    }

    function _deleteStakers(address _operator, address _staker) internal {
        uint arrayLength = operatorStakers[_operator].length;
        uint indexToBeDeleted;
        for (uint i=0; i<arrayLength; i++){
            if (operatorStakers[_operator][i] == _staker) {
                indexToBeDeleted = i;
                break;
            }
        }
        if (indexToBeDeleted < arrayLength-1) {
            operatorStakers[_operator][indexToBeDeleted] = operatorStakers[_operator][arrayLength-1];
        }
        operatorStakers[_operator].pop();
    }

    function _distributeReward(uint256 amount, address[] calldata _tssMembers,bool isByBlock) internal {
        if (amount > 0) {
            uint256 sendAmount = 0;
            uint256 totalShares = 0;
            uint256 accu = 0;
            for (uint i=0; i<_tssMembers.length; i++){
                totalShares = totalShares.add(operatorShares[_tssMembers[i]]);
            }
            for (uint j = 0; j < _tssMembers.length; j++) {
                address operator = _tssMembers[j];
                for (uint i = 0; i < operatorStakers[operator].length; i++) {
                    address staker = operatorStakers[operator][i];
                    sendAmount = (amount * operatorStakerShares[operator][staker]).div(totalShares);
                    rewardDetails[staker] += sendAmount;
                    accu = accu.add(sendAmount);
                    if (isByBlock) {
                        totalAmount = totalAmount.sub(sendAmount);
                    }
                }
            }
            uint256 reserved = amount.sub(accu);
            if (reserved > 0) {
                dust = dust.add(reserved);
            }
        }
    }


}

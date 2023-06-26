// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/**
 * @title ITssRewardContract
 */

interface ITssRewardContract {
    /**********
     * Events *
     **********/

    event DistributeTssReward(
        uint256 lastBatchTime,
        uint256 batchTime,
        uint256 amount,
        address[] tssMembers
    );

    event DistributeTssRewardByBlock(
        uint256   blockStartHeight,
        uint32     length,
        uint256    amount,
        address[] tssMembers
    );

    event Claim(
        address owner,
        uint256 amount
    );

    /********************
     * Public Functions *
     ********************/

    /**
     * @dev Query total undistributed balance.
     * @return Amount of undistributed rewards.
     */
    function queryReward() external view returns (uint256);

    /**
     * @dev Auto distribute reward to tss members.
     * @param _blockStartHeight L2 rollup batch block start height.
     * @param _length Rollup batch length.
     * @param _batchTime rollup batch time.
     * @param _tssMembers Tss member address array.
     */
    function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] calldata _tssMembers) external;

    /**
     * @dev clear contract(canonical).
     */
    function withdraw() external;

    /**
     * @dev Claim reward and withdraw
     */
    function claim() external;

    /**
     * @dev default claimer == staker, if staker is multi-signature address,must set claimer
     * @param _staker the address of staker
     * @param _claimer the address for staker to claim reward
     */
    function setClaimer(address _staker, address _claimer) external;

    /**
     * @dev Initiate a request to claim
     */
    function requestClaim() external returns (bool);

    /**
     * @dev Query the remaining time required to claim
     */
    function queryClaimTime() external returns (uint256);

    function setSccAddr(address sccAddr) external;

    function setStakeSlashAddr(address ssAddr) external;

    function setSendAmountPerYear(uint256) external;

    function setWaitingTime(uint256) external;

}

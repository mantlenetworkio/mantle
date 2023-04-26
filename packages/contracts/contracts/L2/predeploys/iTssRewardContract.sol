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
     * @param _tssMembers Tss member address array.
     */
    function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] calldata _tssMembers) external;

    /**
     * @dev Update deposit block gas into contract.
     * @param _blockID Update gas reward L2 block ID.
     * @return Update success.
     */
    function updateReward(uint256 _blockID, uint256 _amount) external returns (bool);
    /**
     * @dev withdraw dust.
     */
    function withdrawDust() external;

    /**
     * @dev clear contract(canonical).
     */
    function withdraw() external;

    /**
     * @dev Increases the `staker`'s delegated shares
     * @param _operator the address of operator which staker chosed
     * @param _staker the address of staker
     * @param _shares the number of staker delegated for operator
     */
    function increaseDelegatedShares(address _operator, address _staker, uint256 _shares) external;

    /**
     * @dev Decreases the `staker`'s delegated shares
     * @param _operator the address of operator which staker chosed
     * @param _staker the address of staker
     * @param _shares the number of staker delegated for operator
     */
    function decreaseDelegatedShares(address _operator, address _staker, uint256 _shares) external;

    /**
     * @dev first stake and delegated shares
     * @param _operator the address of operator which staker chosed
     * @param _staker the address of staker
     * @param _shares the number of staker delegated for operator
     */
    function delegate(address _operator, address _staker, uint256 _shares) external;

    /**
     * @dev Claim reward
     * @param _addr the address of reward owner
     */
    function claim(address _addr) external;

    /**
     * @dev Claim reward and withdraw
     * @param _addr the address of reward owner
     */
    function claimWithdraw(address _addr) external;

    /**
     * @dev default claimer == staker, if staker is multi-signature address,must set claimer
     * @param _staker the address of staker
     * @param _claimer the address for staker to claim reward
     */
    function setClaimer(address _staker, address _claimer) external;

}

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
        uint256 blockStartHeight,
        uint256 length,
        address[] tssMembers
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
    function claimReward(uint256 _blockStartHeight, uint32 _length, address[] calldata _tssMembers) external;

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
}

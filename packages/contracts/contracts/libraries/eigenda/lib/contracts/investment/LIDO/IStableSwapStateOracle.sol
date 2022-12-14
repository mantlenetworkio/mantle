// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface IStableSwapStateOracle {
    function getState()
        external
        view
        returns (uint256 _timestamp, uint256 _etherBalance, uint256 _stethBalance, uint256 _stethPrice);
}

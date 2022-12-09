// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

/// @notice this is an interface to allow a provided "receiver" address to receive ETH from the EigenPod contract
interface IBeaconChainETHReceiver {
    function receiveBeaconChainETH() external payable; 
}
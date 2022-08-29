// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

interface IStakingSlashing {
    struct DepositInfo {
        bytes pubKey;
        uint256 amount;
    }

    function staking(uint256, bytes memory) external;

    function withdrawToken() external;

    function quit() external;

    function clearQuitList() external;

    function slashing(bytes memory, bytes memory) external;

    function getDeposits(address) external returns (DepositInfo memory);

    function getSlashRecord(uint256, address) external returns (bool);
}

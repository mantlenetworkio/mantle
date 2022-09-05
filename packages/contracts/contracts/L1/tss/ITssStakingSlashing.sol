// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

interface IStakingSlashing {
    struct DepositInfo {
        address pledgor;
        bytes pubKey;
        uint256 amount;
    }

    // tx
    function setAddress(address , address ) external;
    function setSlashingParams(uint256[2] memory , uint256[2] memory) external;
    function staking(uint256 , bytes memory) external;
    function withdrawToken() external;
    function quitRequest() external;
    function clearQuitRequestList() external;
    function slashing(bytes memory, bytes memory) external;
    function unJail() external;

    // query
    function getSlashingParams() external view returns (uint256[2] memory, uint256[2] memory);
    function getQuitRequestList() external view returns (address[] memory);
    function getDeposits(address) external returns (DepositInfo memory);
    function batchGetDeposits(address[] calldata) external view returns (DepositInfo[] memory);
    function getSlashRecord(uint256, address) external view returns (bool);
    function isJailed(address user) external returns (bool);
}

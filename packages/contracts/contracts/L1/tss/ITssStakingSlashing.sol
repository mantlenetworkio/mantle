// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

interface IStakingSlashing {

    // tx
    function setAddress(address , address ) external;
    function setSlashingParams(uint256[2] calldata , uint256[2] calldata) external;
    function setPublicKey(bytes calldata) external;
    function quitRequest() external;
    function clearQuitRequestList() external;
    function slashing(bytes calldata, bytes calldata) external;
    function unJail() external;

    // query
    function getSlashingParams() external view returns (uint256[2] memory, uint256[2] memory);
    function getQuitRequestList() external view returns (address[] memory);
    function getSlashRecord(uint256, address) external view returns (bool);
    function isJailed(address) external returns (bool);
}

// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

interface IStakingSlashing {

    // tx
    function setTokenAddress(address) external;
    function setTssGroupAddress(address) external;
    function setRegulatoryAccount(address) external;
    function setClaimer(address, address) external;
    function setSlashingParams(uint256[2] calldata) external;
    function setTssManager(address) external;
    function quitRequest() external;
    function clearQuitRequestList() external;
    function slashing(bytes calldata, bytes calldata) external;
    function unJail() external;

    // query
    function getSlashingParams() external view returns (uint256[2] memory);
    function getQuitRequestList() external view returns (address[] memory);
    function getSlashRecord(uint256, address) external view returns (bool);
    function isJailed(address) external returns (bool);
    function isCanOperator(address) external returns (bool);

    //fund
    function deposit(uint256 amount) external returns (uint256);
    function withdraw() external;
    function completeWithdraw() external;
    function startWithdraw() external;
    function canCompleteQueuedWithdrawal() external returns (bool);

    //delegation
    function registerAsOperator(bytes calldata) external;
    function delegateTo(address) external;




}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

// import "forge-std/Test.sol";
// import "../../contracts/Interfaces/IServiceManager.sol";

contract ServiceManagerMock {
    // is IServiceManager, DSTest

    function taskNumber() external pure returns (uint32) {
        return 0;
    }

    function freezeOperator(address operator) external pure {}

    function revokeSlashingAbility(address operator, uint32 unbondedAfter) external pure {}

    // function collateralToken() external view returns (IERC20) {
    //     return IERC20(address(0));
    // }

    // function eigenLayrDelegation() external view returns (IEigenLayrDelegation) {
    //     return IEigenLayrDelegation(address(0));
    // }

    function stakeWithdrawalVerification(bytes calldata data, uint256 initTimestamp, uint256 unlockTime)
        external
        pure
    {}

    function latestTime() external pure returns (uint32) {
        return type(uint32).max;
    }
}

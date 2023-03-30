// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";

contract BVM_EigenDataLayrFee is OwnableUpgradeable, ReentrancyGuardUpgradeable {
    using SafeMathUpgradeable for uint256;
    using AddressUpgradeable for address;

    address public gasFeeAddress;
    uint256 userRollupFee;

    event RollupFeeHistory(uint256 l2Block, uint256 userRollupFee);

    function initialize(address _address) public initializer {
        __Ownable_init();
        userRollupFee = 0;
        gasFeeAddress = _address;
    }

    function setFeeAddress(address _address) public onlyOwner {
        gasFeeAddress = _address;
    }

    modifier onlyGasFee() {
        require(msg.sender == gasFeeAddress, "contract call is not gas fee address");
        _;
    }

    function setRollupFee(uint256 _l2Block, uint256 _userRollupFee) public onlyGasFee {
        userRollupFee = _userRollupFee;
        emit RollupFeeHistory(_l2Block, _userRollupFee);
    }

    function getRollupFee() public view returns (uint256) {
        return userRollupFee;
    }
}

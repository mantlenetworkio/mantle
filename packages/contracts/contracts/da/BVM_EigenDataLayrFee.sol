// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract BVM_EigenDataLayrFee is Initializable, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    using SafeMathUpgradeable for uint256;
    using AddressUpgradeable for address;

    address public gasFeeAddress;
    uint256 userRollupFee;

    event RollupFeeHistory(uint256 l2Block, uint256 userRollupFee);
    event FeeAddressUpdated(address oldFeeAddress, address newFeeAddress);

    constructor() {
        _disableInitializers();
    }

    function initialize(address _address) public initializer {
        require(_address != address(0), "initialize: can't set zero address to gasFeeAddress");
        __Ownable_init();
        userRollupFee = 0;
        gasFeeAddress = _address;
    }

    function setFeeAddress(address _address) public onlyOwner {
        require(_address != address(0), "setFeeAddress: address is the zero address");
        address oldGasFeeAddress = gasFeeAddress;
        gasFeeAddress = _address;
        emit FeeAddressUpdated(oldGasFeeAddress, gasFeeAddress);
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

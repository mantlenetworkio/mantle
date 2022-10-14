// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "./ISequencer.sol";

contract Sequencer is ISequencer, ReentrancyGuardUpgradeable, OwnableUpgradeable {
    using SafeERC20 for IERC20;

    mapping(address => Sequencer) public seqsDeposit;
    address[] public depositKeys;
    address bitToken;

    function initialize(address _bitToken) public initializer {
        __Ownable_init();
        __ReentrancyGuard_init();

        bitToken = _bitToken;
    }

    function changeBitAddress(address _bitToken) public onlyOwner {
        bitToken = _bitToken;
    }

    function createSequencer(
        uint256 _amount,
        address _mintAddress,
        bytes calldata _nodeID
    ) external nonReentrant {
        // check params todo: check nodeID
        require(_amount > 0, "invild amount");
        require(_mintAddress != address(0), "invild address, address can not be 0");

        // check already have sequencer
        require(seqsDeposit[msg.sender].mintAddress == address(0), "already have deposit");
        IERC20(bitToken).safeTransferFrom(msg.sender, address(this), _amount);

        uint256 index = depositKeys.length;
        seqsDeposit[msg.sender] = Sequencer({
            owner: msg.sender,
            mintAddress: _mintAddress,
            nodeID: _nodeID,
            amount: _amount,
            keyIndex: index
        });
        depositKeys[index] = msg.sender;
    }

    function deposit(uint256 _amount) external nonReentrant {
        // check params
        require(_amount > 0, "invild amount");
        // check already have sequencer
        require(seqsDeposit[msg.sender].mintAddress != address(0), "do not have create");

        // transfer
        IERC20(bitToken).safeTransferFrom(msg.sender, address(this), _amount);
        seqsDeposit[msg.sender].amount += _amount;
    }

    function withdraw(uint256 _amount) external nonReentrant {
        // check params
        require(_amount > 0, "invild amount");
        // check already have sequencer
        require(seqsDeposit[msg.sender].mintAddress != address(0), "do not have create");

        uint256 withdrawAmount = _amount;
        if (_amount > seqsDeposit[msg.sender].amount) {
            // when _amount > seqsDeposit.amount, withdraw all tokens
            withdrawAmount = seqsDeposit[msg.sender].amount;
        }

        // transfer
        IERC20(bitToken).safeTransfer(msg.sender, withdrawAmount);

        seqsDeposit[msg.sender].amount -= withdrawAmount;
        if (seqsDeposit[msg.sender].amount == 0) {
            deleteSequencer(msg.sender);
        }
    }

    function withdrawAll() external nonReentrant {
        // check already have sequencer
        require(seqsDeposit[msg.sender].mintAddress != address(0), "do not have create");

        uint256 withdrawAmount = seqsDeposit[msg.sender].amount;

        // transfer
        IERC20(bitToken).safeTransfer(msg.sender, withdrawAmount);
        deleteSequencer(msg.sender);
    }

    function getSequencers() external view returns (Sequencer[] memory) {
        Sequencer[] memory seqs = new Sequencer[](depositKeys.length);
        for (uint256 i = 0; i < depositKeys.length; i++) {
            address key = depositKeys[i];
            seqs[i] = seqsDeposit[key];
        }
        return seqs;
    }

    function getSequencer(address signer) external view returns (Sequencer memory) {
        return seqsDeposit[signer];
    }

    function isSequencer(address signer) external view returns (bool) {
        return seqsDeposit[signer].mintAddress != address(0);
    }

    function deleteSequencer(address signer) internal {
        uint256 index = seqsDeposit[signer].keyIndex;
        uint256 length = depositKeys.length;

        // delete
        depositKeys[index] = depositKeys[length - 1];
        delete depositKeys[length - 1];
        delete seqsDeposit[signer];
    }
}

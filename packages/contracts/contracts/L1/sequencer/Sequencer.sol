// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "./ISequencer.sol";

contract Sequencer is ISequencer, ReentrancyGuardUpgradeable, OwnableUpgradeable {
    using SafeERC20 for IERC20;

    mapping(address => Sequencer) public sequencers;
    address[] public owners;
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
        require(sequencers[msg.sender].mintAddress == address(0), "already have deposit");
        IERC20(bitToken).safeTransferFrom(msg.sender, address(this), _amount);

        uint256 index = owners.length;
        sequencers[msg.sender] = Sequencer({
            owner: msg.sender,
            mintAddress: _mintAddress,
            nodeID: _nodeID,
            amount: _amount,
            keyIndex: index
        });
        owners[index] = msg.sender;
    }

    function deposit(uint256 _amount) external nonReentrant {
        // check params
        require(_amount > 0, "invild amount");
        // check already have sequencer
        require(sequencers[msg.sender].mintAddress != address(0), "do not have create");

        // transfer
        IERC20(bitToken).safeTransferFrom(msg.sender, address(this), _amount);
        sequencers[msg.sender].amount += _amount;
    }

    function withdraw(uint256 _amount) external nonReentrant {
        // check params
        require(_amount > 0, "invild amount");
        // check already have sequencer
        require(sequencers[msg.sender].mintAddress != address(0), "do not have create");

        uint256 withdrawAmount = _amount;
        if (_amount > sequencers[msg.sender].amount) {
            // when _amount > sequencers.amount, withdraw all tokens
            withdrawAmount = sequencers[msg.sender].amount;
        }

        // transfer
        IERC20(bitToken).safeTransfer(msg.sender, withdrawAmount);

        sequencers[msg.sender].amount -= withdrawAmount;
        if (sequencers[msg.sender].amount == 0) {
            deleteSequencer(msg.sender);
        }
    }

    function withdrawAll() external nonReentrant {
        // check already have sequencer
        require(sequencers[msg.sender].mintAddress != address(0), "do not have create");

        uint256 withdrawAmount = sequencers[msg.sender].amount;

        // transfer
        IERC20(bitToken).safeTransfer(msg.sender, withdrawAmount);
        deleteSequencer(msg.sender);
    }

    function getSequencers() external view returns (Sequencer[] memory) {
        Sequencer[] memory seqs = new Sequencer[](owners.length);
        for (uint256 i = 0; i < owners.length; i++) {
            address key = owners[i];
            seqs[i] = sequencers[key];
        }
        return seqs;
    }

    function getSequencer(address signer) external view returns (Sequencer memory) {
        return sequencers[signer];
    }

    function isSequencer(address signer) external view returns (bool) {
        return sequencers[signer].mintAddress != address(0);
    }

    function deleteSequencer(address signer) internal {
        uint256 index = sequencers[signer].keyIndex;
        uint256 length = owners.length;

        // delete
        owners[index] = owners[length - 1];
        delete owners[length - 1];
        delete sequencers[signer];
    }
}

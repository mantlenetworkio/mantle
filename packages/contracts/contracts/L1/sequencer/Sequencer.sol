// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "./ISequencer.sol";

contract Sequencer is ReentrancyGuardUpgradeable, OwnableUpgradeable {
    using SafeERC20 for IERC20;

    struct SequencerInfo {
        address owner;
        address mintAddress;
        bytes nodeID;
        uint256 amount;
        uint256 keyIndex;
    }
    mapping(address => SequencerInfo) public sequencers;
    address[] public owners;
    address bitToken;

    event SequencerCreate(address, address, bytes);
    event SequencerUpdate(address, bytes, uint256);
    event SequencerDelete(address, bytes);

    function initialize(address _bitToken) public initializer {
        __Ownable_init();
        __ReentrancyGuard_init();

        bitToken = _bitToken;
    }

    function changeBitAddress(address _bitToken) public onlyOwner {
        bitToken = _bitToken;
    }

    // Create a new sequencer info and init amount
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
        sequencers[msg.sender] = SequencerInfo({
            owner: msg.sender,
            mintAddress: _mintAddress,
            nodeID: _nodeID,
            amount: _amount,
            keyIndex: index
        });
        owners[index] = msg.sender;
        emit SequencerCreate(msg.sender, _mintAddress, _nodeID);
    }

    // Check sequencer exist then add deposit amount
    function deposit(uint256 _amount) external nonReentrant {
        // check params
        require(_amount > 0, "invild amount");
        // check already have sequencer
        require(sequencers[msg.sender].mintAddress != address(0), "do not have create");

        // transfer
        IERC20(bitToken).safeTransferFrom(msg.sender, address(this), _amount);
        sequencers[msg.sender].amount += _amount;
        emit SequencerUpdate(
            sequencers[msg.sender].mintAddress,
            sequencers[msg.sender].nodeID,
            sequencers[msg.sender].amount
        );
    }

    // amount > deposit(signer).amount -> withdraw all
    // 0 < amount < deposit(signer).amount -> withdraw amount to signer
    // when deposit(signer).amount = 0, delete the sequencer
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
        emit SequencerUpdate(
            sequencers[msg.sender].mintAddress,
            sequencers[msg.sender].nodeID,
            sequencers[msg.sender].amount
        );

        if (sequencers[msg.sender].amount == 0) {
            deleteSequencer(msg.sender);
        }
    }

    // Check sequencer exist then withdraw all
    function withdrawAll() external nonReentrant {
        // check already have sequencer
        require(sequencers[msg.sender].mintAddress != address(0), "do not have create");

        uint256 withdrawAmount = sequencers[msg.sender].amount;

        // transfer
        IERC20(bitToken).safeTransfer(msg.sender, withdrawAmount);
        deleteSequencer(msg.sender);
    }

    // Return all sequencer infos
    function getSequencers() external view returns (SequencerInfo[] memory) {
        SequencerInfo[] memory seqs = new SequencerInfo[](owners.length);
        for (uint256 i = 0; i < owners.length; i++) {
            address key = owners[i];
            seqs[i] = sequencers[key];
        }
        return seqs;
    }

    // Return sequencer info by signer address
    function getSequencer(address signer) public view returns (SequencerInfo memory) {
        return sequencers[signer];
    }

    // Return if signer exist
    function isSequencer(address signer) public view returns (bool) {
        return sequencers[signer].mintAddress != address(0);
    }

    // Delete sequencer
    function deleteSequencer(address signer) internal {
        uint256 index = sequencers[signer].keyIndex;
        uint256 length = owners.length;

        emit SequencerDelete(sequencers[signer].mintAddress, sequencers[signer].nodeID);
        // delete
        owners[index] = owners[length - 1];
        delete owners[length - 1];
        delete sequencers[signer];
    }
}

// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

contract Sequencer is ReentrancyGuardUpgradeable, OwnableUpgradeable {
    using SafeERC20 for IERC20;

    struct SequencerInfo {
        address owner;
        address mintAddress;
        bytes nodeID;
        uint256 amount;
        uint256 keyIndex;
    }

    // Maps sequencers, set by signer address(signer -> sequencer)
    mapping(address => SequencerInfo) public sequencers;

    // Maps mint address to signer
    // to check and avoid mintAddress being bound by multiple singers
    // (Mint address -> signer)
    mapping(address => address) public rel;

    // Store the address of the signer
    address[] public owners;
    // Store bit token address
    address public bitToken;
    // Store the Epoch
    uint256 public epoch;
    // Store limit of sequencer
    uint256 public sequencerLimit;
    // Store scheduler
    bytes public scheduler;

    // SequencerCreate(signer, mintAddress, nodeID)
    event SequencerCreate(address, address, bytes);
    // SequencerUpdate(mintAddress, nodeID, amount)
    event SequencerUpdate(address, bytes, uint256);
    // SequencerDelete(mintAddress, nodeID)
    event SequencerDelete(address, bytes);

    function initialize(address _bitToken) public initializer {
        __Ownable_init();
        __ReentrancyGuard_init();

        bitToken = _bitToken;
        epoch = 0;
        sequencerLimit = 0;
    }

    /**
     * Update Epoch
     * @param nodeID new scheculer`s nodeID
     */
    function updateScheduler(bytes memory nodeID) public onlyOwner {
        scheduler = nodeID;
    }

    /**
     * Update Epoch
     * @param _limit new limit
     */
    function updateSequencerLimit(uint256 _limit) public onlyOwner {
        sequencerLimit = _limit;
    }

    /**
     * Update Epoch
     * @param _epoch new epoch
     */
    function updateEpoch(uint256 _epoch) public onlyOwner {
        epoch = _epoch;
    }

    /**
     * Update bit token address
     * @param _bitToken new ERC20 address of bit token
     */
    function updateBitAddress(address _bitToken) public onlyOwner {
        bitToken = _bitToken;
    }

    /**
     * Create a new sequencer info and init amount
     * @param _amount amount of bit token, will transfer to this contract when sequencer create
     * @param _mintAddress sequencer mint address
     * @param _nodeID sequencer node ID
     */
    function createSequencer(
        uint256 _amount,
        address _mintAddress,
        bytes calldata _nodeID
    ) external nonReentrant {
        // check params todo: check nodeID
        require(_amount > 0, "Invild amount");
        require(_mintAddress != address(0), "Invild address, address can not be 0");
        // check already have sequencer
        require(sequencers[msg.sender].mintAddress == address(0), "Already has been created");
        require(rel[_mintAddress] == address(0), "This mint address already has owner");
        IERC20(bitToken).safeTransferFrom(msg.sender, address(this), _amount);

        uint256 index = owners.length;
        sequencers[msg.sender] = SequencerInfo({
            owner: msg.sender,
            mintAddress: _mintAddress,
            nodeID: _nodeID,
            amount: _amount,
            keyIndex: index
        });
        owners.push(msg.sender);
        rel[_mintAddress] = msg.sender;
        emit SequencerCreate(msg.sender, _mintAddress, _nodeID);
    }

    /**
     * Check sequencer exist then add deposit amount
     * @param _amount amount of bit token
     */
    function deposit(uint256 _amount) external nonReentrant {
        // check params
        require(_amount > 0, "Invild amount");
        // check already have sequencer
        require(sequencers[msg.sender].mintAddress != address(0), "Sequencer not exist");

        // transfer
        IERC20(bitToken).safeTransferFrom(msg.sender, address(this), _amount);
        sequencers[msg.sender].amount += _amount;
        emit SequencerUpdate(
            sequencers[msg.sender].mintAddress,
            sequencers[msg.sender].nodeID,
            sequencers[msg.sender].amount
        );
    }

    /**
     * amount > deposit(signer).amount -> withdraw all
     * 0 < amount < deposit(signer).amount -> withdraw amount to signer
     * when deposit(signer).amount = 0, delete the sequencer
     * @param _amount amount of bit token
     */
    function withdraw(uint256 _amount) external nonReentrant {
        // check params
        require(_amount > 0, "Invild amount");
        // check already have sequencer
        require(sequencers[msg.sender].mintAddress != address(0), "Sequencer not exist");

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

    /**
     * Check sequencer exist then withdraw all.
     * This action will delete sequencer after withdraw
     */
    function withdrawAll() external nonReentrant {
        // check already have sequencer
        require(sequencers[msg.sender].mintAddress != address(0), "Do not have create");

        uint256 withdrawAmount = sequencers[msg.sender].amount;

        // transfer
        IERC20(bitToken).safeTransfer(msg.sender, withdrawAmount);

        emit SequencerUpdate(sequencers[msg.sender].mintAddress, sequencers[msg.sender].nodeID, 0);
        deleteSequencer(msg.sender);
    }

    /**
     * Return all sequencer infos
     * @return seqs all sequencers
     */
    function getSequencers() external view returns (SequencerInfo[] memory) {
        SequencerInfo[] memory seqs = new SequencerInfo[](owners.length);
        for (uint256 i = 0; i < owners.length; i++) {
            address key = owners[i];
            seqs[i] = sequencers[key];
        }
        return seqs;
    }

    /**
     * Return sequencer info by signer address
     * @param signer signer address, the key to find sequencer
     # @return seq sequencer info
     */
    function getSequencer(address signer) public view returns (SequencerInfo memory) {
        return sequencers[signer];
    }

    /**
     * Return owners
     # @return owners all owners
     */
    function getOwners() public view returns (address[] memory) {
        return owners;
    }

    /**
     * Return if signer exist
     */
    function isSequencer(address signer) public view returns (bool) {
        return sequencers[signer].mintAddress != address(0);
    }

    /**
     * Delete sequencer
     */
    function deleteSequencer(address signer) internal {
        uint256 index = sequencers[signer].keyIndex;
        uint256 length = owners.length;

        emit SequencerDelete(sequencers[signer].mintAddress, sequencers[signer].nodeID);
        // change index
        owners[index] = owners[length - 1];
        sequencers[owners[index]].keyIndex = index;

        // delete
        delete rel[sequencers[signer].mintAddress];
        owners.pop();
        delete sequencers[signer];
    }
}

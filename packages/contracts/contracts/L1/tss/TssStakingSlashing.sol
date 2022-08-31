// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { Lib_AddressResolver } from "../../libraries/resolver/Lib_AddressResolver.sol";
import { Lib_AddressManager } from "../../libraries/resolver/Lib_AddressManager.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "./ITssGroupManager.sol";
import "./ITssStakingSlashing.sol";

contract TssStakingSlashing is
    Initializable,
    Lib_AddressResolver,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    IStakingSlashing
{
    enum SlashType {
        nothing,
        uptime,
        animus
    }

    struct SlashMsg {
        uint256 batchIndex;
        address jailNode;
        address[] tssNodes;
        SlashType slashType;
    }
    // staking parameter settings
    // bit token contract address
    address public BitToken;
    // tss group contract address
    address public tssGroupContract;
    // storage staker infos (key:staker address)
    mapping(address => DepositInfo) public deposits;

    // slashing parameter settings
    // record the quit request
    address[] public quitList;
    // slashing amount of type uptime and animus (0:uptime , 1:animus)
    uint256[2] public slashAmount;
    // additional rewards for sender (0:uptime , 1:animus)
    uint256[2] public exIncome;
    // record the slash operate (map[batchIndex] -> (map[staker] -> slashed))
    mapping(uint256 => mapping(address => bool)) slashRecord;

    /**
     * @notice staking for himself
     * @param 0 staker address
     * @param 1 staker public key and deposit amount
     */
    event AddDeposit(address, DepositInfo);

    /**
     * @notice withdraw for himself
     * @param 0 staker address
     * @param 1 total amount withdraw
     */
    event Withdraw(address, uint256);

    /**
     * @notice slash tssnode
     * @param 0 slashed address
     * @param 1 slash type
     */
    event Slashing(address, SlashType);

    /***************
     * Constructor *
     ***************/

    /**
     * This contract is intended to be behind a delegate proxy.
     * We pass the zero address to the address resolver just to satisfy the constructor.
     * We still need to set this value in initialize().
     */
    constructor() Lib_AddressResolver(address(0)) {}

    /**
     * @notice initializes the contract setting and the deployer as the initial owner
     * @param _libAddressManager address manager contract address
     */
    function initialize(address _libAddressManager) public initializer {
        __Ownable_init();

        require(
            address(libAddressManager) == address(0),
            "L1CrossDomainMessenger already intialized."
        );
        libAddressManager = Lib_AddressManager(_libAddressManager);

        BitToken = resolve("L1_BitAddress");
        tssGroupContract = resolve("TssGroupManager");
    }

    /**
     * @notice change the bit token and tssGroup contract address
     * @param _token the erc20 bit token contract address
     * @param _tssGroup tssGroup contract address
     */
    function setAddress(address _token, address _tssGroup) public onlyOwner {
        BitToken = _token;
        tssGroupContract = _tssGroup;
    }

    /**
     * @notice set the slashing params (0 -> uptime , 1 -> animus)
     * @param _slashAmount the amount to be deducted for each type
     * @param _exIncome additional amount available to the originator of the report
     */
    function setSlashingParams(uint256[2] memory _slashAmount, uint256[2] memory _exIncome)
        public
        onlyOwner
    {
        for (uint256 i = 0; i < 2; i++) {
            require(_exIncome[i] > 0, "invalid amount");
            require(_slashAmount[i] > _exIncome[i], "slashAmount need bigger than exIncome");
            slashAmount[i] = _slashAmount[i];
            exIncome[i] = _exIncome[i];
        }
    }

    /**
     * @notice set the slashing params (0 -> uptime , 1 -> animus)
     * @return _slashAmount the amount to be deducted for each type
     */
    function getSlashingParams() public view returns (uint256[2] memory, uint256[2] memory) {
        return (slashAmount, exIncome);
    }

    /**
     * @notice staking entrance for user to deposit bit tokens
     * @param _amount deposit amount of bit token
     * @param _pubKey public key of sender
     */
    function staking(uint256 _amount, bytes memory _pubKey) public nonReentrant {
        // verify amount
        require(_amount > 0, "invalid amount");

        if (deposits[msg.sender].pubKey.length > 0) {
            // increase pledge amount
            require(isEqual(deposits[msg.sender].pubKey, _pubKey), "pubKey not equal");
        } else {
            // new to staking
            require(
                msg.sender == ITssGroupManager(tssGroupContract).publicKeyToAddress(_pubKey),
                "invalid pubKey"
            );
            deposits[msg.sender].pubKey = _pubKey;
        }

        // send bit token to staking contract , need user approve first
        require(
            IERC20(BitToken).transferFrom(msg.sender, address(this), _amount),
            "transfer erc20 token failed"
        );
        deposits[msg.sender].amount += _amount;

        emit AddDeposit(msg.sender, DepositInfo({ pubKey: _pubKey, amount: _amount }));
    }

    /**
     * @notice user who not elected to be validator to withdraw their bit token
     */
    function withdrawToken() public nonReentrant {
        uint256 amount = deposits[msg.sender].amount;
        require(amount > 0, "do not have deposit");
        bytes memory pubKey = deposits[msg.sender].pubKey;

        // when not in consensus period or be selected
        require(
            !ITssGroupManager(tssGroupContract).memberExistInActive(pubKey) &&
                !ITssGroupManager(tssGroupContract).memberExistActive(pubKey),
            "not at the right time"
        );

        delete deposits[msg.sender];

        require(IERC20(BitToken).transfer(msg.sender, amount), "erc20 transfer failed");

        emit Withdraw(msg.sender, amount);
    }

    /**
     * @notice send quit request for the next election
     */
    function quit() public nonReentrant {
        require(deposits[msg.sender].amount > 0, "do not have deposit");
        // when not in consensus period
        require(ITssGroupManager(tssGroupContract).inActiveIsEmpty(), "not at the right time");
        // is active member
        require(
            ITssGroupManager(tssGroupContract).memberExistActive(deposits[msg.sender].pubKey),
            "not at the active group"
        );
        for (uint256 i = 0; i < quitList.length; i++) {
            require(quitList[i] != msg.sender, "already in quitList");
        }
        quitList.push(msg.sender);
    }

    /**
     * @notice return the quit list
     */
    function getQuitList() public view returns (address[] memory) {
        return quitList;
    }

    /**
     * @notice clear the quit list
     */
    function clearQuitList() public onlyOwner {
        delete quitList;
    }

    /**
     * @notice verify the slash message then slash
     * @param _messageBytes the message that abi encode by type SlashMsg
     * @param _sig the signature of the hash keccak256(_messageBytes)
     */
    function slashing(bytes memory _messageBytes, bytes memory _sig) public nonReentrant {
        SlashMsg memory message = abi.decode(_messageBytes, (SlashMsg));
        // verify tss member state
        require(!isJailed(message.jailNode), "the node already jailed");
        require(
            ITssGroupManager(tssGroupContract).memberExistActive(deposits[msg.sender].pubKey),
            "sender not at the active group"
        );
        // have not slash before todo record by index and node address

        require(!slashRecord[message.batchIndex][message.jailNode], "already slashed");
        slashRecord[message.batchIndex][message.jailNode] = true;

        require(
            ITssGroupManager(tssGroupContract).verifySign(keccak256(_messageBytes), _sig),
            "singer not tss group pub key"
        );

        // slash tokens
        slash(message);
        emit Slashing(message.jailNode, message.slashType);
    }

    /**
     * @notice slash the staker and distribute rewards to voters
     * @param message the message about the slash infos
     */
    function slash(SlashMsg memory message) internal {
        // slashing params check
        for (uint256 i = 0; i < 2; i++) {
            require(slashAmount[i] > 0, "have not set the slash amount");
            require(exIncome[i] > 0, "have not set the extra income amount");
        }
        bytes memory jailNodePubKey = deposits[message.jailNode].pubKey;
        if (message.slashType == SlashType.uptime) {
            // jail and transfer deposits
            ITssGroupManager(tssGroupContract).memberJail(jailNodePubKey);
            transformDeposit(message.jailNode, 0, message.tssNodes);
        } else if (message.slashType == SlashType.animus) {
            // remove the member and transfer deposits
            ITssGroupManager(tssGroupContract).removeMember(jailNodePubKey);
            transformDeposit(message.jailNode, 1, message.tssNodes);
        } else {
            require(false, "err type for slashing");
        }
    }

    /**
     * @notice distribute rewards to voters
     * @param deduction address of the punished
     * @param slashType the type to punished
     * @param tssNodes participants other than the initiator
     */
    function transformDeposit(
        address deduction,
        uint256 slashType,
        address[] memory tssNodes
    ) internal {
        uint256 deductedAmount;
        uint256 totalTransfer;
        uint256 extraAmount;
        uint256 remainder;
        uint256 gain;

        if (deposits[deduction].amount > slashAmount[slashType]) {
            // deposit > slashAmount, deduct slashAmount then
            // distribute additional tokens for the sender
            deductedAmount = slashAmount[slashType];
            deposits[deduction].amount -= slashAmount[slashType];
            extraAmount = slashAmount[slashType] - exIncome[slashType];
        } else if (deposits[deduction].amount > exIncome[slashType]) {
            // exIncome < deposit <= slashAmount, deduct all token then
            // distribute additional tokens for the sender
            deductedAmount = deposits[deduction].amount;
            deposits[deduction].amount = 0;
            extraAmount = deductedAmount - exIncome[slashType];
        } else if (deposits[deduction].amount > 0) {
            // 0 < deposit <= exIncome, deduct all token
            deductedAmount = deposits[deduction].amount;
            deposits[deduction].amount = 0;
            extraAmount = deductedAmount;
        } else {
            require(false, "panic , invalid type");
        }

        // transfer sender the remainder as an additional bonus
        remainder = extraAmount % tssNodes.length;
        gain = (extraAmount - remainder) / tssNodes.length;

        deposits[msg.sender].amount += exIncome[slashType] + remainder;
        totalTransfer = exIncome[slashType] + remainder;
        for (uint256 i = 0; i < tssNodes.length; i++) {
            totalTransfer += gain;
            deposits[tssNodes[i]].amount += gain;
        }
        // The total transfer amount is the same as the deducted amount
        require(totalTransfer == deductedAmount, "panic, calculation error");
    }

    /**
     * @notice set tss node status unjail
     */
    function unJail() public {
        // slashing params check
        for (uint256 i = 0; i < 2; i++) {
            require(slashAmount[i] > 0, "have not set the slash amount");
            require(exIncome[i] > 0, "have not set the extra income amount");
        }
        require(deposits[msg.sender].amount >= slashAmount[1], "Insufficient balance");
        ITssGroupManager(tssGroupContract).memberUnJail(deposits[msg.sender].pubKey);
    }

    /**
     * @notice get the deposit info
     * @param user address of the staker
     */
    function getDeposits(address user) public view returns (DepositInfo memory) {
        return deposits[user];
    }

    /**
     * @notice get the slash record
     * @param batchIndex the index of batch
     * @param user address of the staker
     */
    function getSlashRecord(uint256 batchIndex, address user) public view returns (bool) {
        return slashRecord[batchIndex][user];
    }

    /**
     * @notice check the tssnode status
     * @param user address of the staker
     */
    function isJailed(address user) public returns (bool) {
        ITssGroupManager.TssMember memory tssMember = ITssGroupManager(tssGroupContract)
            .getTssMember(deposits[user].pubKey);

        return tssMember.status == ITssGroupManager.MemberStatus.jail;
    }

    /**
     * @notice check two bytes for equality
     */
    function isEqual(bytes memory byteListA, bytes memory byteListB) public pure returns (bool) {
        if (byteListA.length != byteListB.length) return false;
        for (uint256 i = 0; i < byteListA.length; i++) {
            if (byteListA[i] != byteListB[i]) return false;
        }
        return true;
    }
}

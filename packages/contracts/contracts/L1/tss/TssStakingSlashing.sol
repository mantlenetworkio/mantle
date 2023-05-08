// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import {DelegationShareBase} from "../delegation/DelegationShareBase.sol";
import {DelegationCallbackBase} from "../delegation/DelegationCallbackBase.sol";
import {IDelegationManager} from "../delegation/interfaces/IDelegationManager.sol";
import {IDelegationShare} from "../delegation/interfaces/IDelegation.sol";
import {IDelegation} from "../delegation/interfaces/IDelegation.sol";
import {CrossDomainEnabled} from "../../libraries/bridge/CrossDomainEnabled.sol";
import {ITssRewardContract} from "../../L2/predeploys/iTssRewardContract.sol";
import {TssDelegationManager} from "./delegation/TssDelegationManager.sol";

import "./ITssGroupManager.sol";
import "./ITssStakingSlashing.sol";
import "./WhiteList.sol";

contract TssStakingSlashing is
    Initializable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    IStakingSlashing,
    DelegationShareBase,
    DelegationCallbackBase,
    CrossDomainEnabled,
    WhiteList
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

    // tss group contract address
    address public tssGroupContract;
    //tss delegation manager address
    address public tssDelegationManagerContract;
    // storage operator infos (key:staker address)
    mapping(address => bytes) public operators;

    // slashing parameter settings
    // record the quit request
    address[] public quitRequestList;
    // slashing amount of type uptime and animus (0:uptime, 1:animus)
    uint256[2] public slashAmount;
    // additional rewards for sender (0:uptime, 1:animus)
    uint256[2] public exIncome;
    // record the slash operate (map[batchIndex] -> (map[staker] -> slashed))
    mapping(uint256 => mapping(address => bool)) slashRecord;




    /**
     * @notice slash tssnode
     * @param 0 slashed address
     * @param 1 slash type
     */
    event Slashing(address, SlashType);



    constructor()  CrossDomainEnabled(address(0)) {
        _disableInitializers();
    }

    /**
     * @notice initializes the contract setting and the deployer as the initial owner
     * @param _bitToken bit token contract address
     * @param _tssGroupContract address tss group manager contract address
     */
    function initialize(address _bitToken,
        address _tssGroupContract,
        address _delegationManager,
        address _delegation,
        address _l1messenger,
        address[] calldata stakerWhitelists,
        address[] calldata operatorWhitelists
        ) public initializer {
        __Ownable_init();
        __ReentrancyGuard_init();
        underlyingToken = IERC20(_bitToken);
        tssGroupContract = _tssGroupContract;
        tssDelegationManagerContract = _delegationManager;
        //initialize delegation
        delegationManager = IDelegationManager(_delegationManager);
        delegation = IDelegation(_delegation);
        messenger = _l1messenger;

        for (uint i = 0; i < stakerWhitelists.length; i++) {
            stakerWhitelist[stakerWhitelists[i]] = true;
        }
        for (uint i = 0; i < operatorWhitelists.length; i++) {
            operatorWhitelist[operatorWhitelists[i]] = true;
        }
    }

    /**
     * @notice change the bit token and tssGroup contract address
     * @param _token the erc20 bit token contract address
     * @param _tssGroup tssGroup contract address
     */
    function setAddress(address _token, address _tssGroup) public onlyOwner {
        underlyingToken = IERC20(_token);
        tssGroupContract = _tssGroup;
    }

    /**
     * @notice set the slashing params (0 -> uptime , 1 -> animus)
     * @param _slashAmount the amount to be deducted for each type
     * @param _exIncome additional amount available to the originator of the report
     */
    function setSlashingParams(uint256[2] calldata _slashAmount, uint256[2] calldata _exIncome)
        public
        onlyOwner
    {
        require(_slashAmount[1] > _slashAmount[0], "invalid param slashAmount, animus <= uptime");
        require(_exIncome[1] > _exIncome[0], "invalid param exIncome, animus <= uptime");

        for (uint256 i = 0; i < 2; i++) {
            require(_exIncome[i] > 0, "invalid amount");
            require(_slashAmount[i] > _exIncome[i], "slashAmount need bigger than exIncome");
            slashAmount[i] = _slashAmount[i];
            exIncome[i] = _exIncome[i];
        }
    }

    /**
     * @notice set the slashing params (0 -> uptime, 1 -> animus)
     * @return _slashAmount the amount to be deducted for each type
     */
    function getSlashingParams() public view returns (uint256[2] memory, uint256[2] memory) {
        return (slashAmount, exIncome);
    }



    function setPublicKey(bytes calldata _pubKey) public nonReentrant {
        if (delegation.isOperator(msg.sender)) {
            operators[msg.sender] = _pubKey;
        }
    }

    function onDelegationReceived(
        address delegator,
        address operator,
        IDelegationShare[] memory investorDelegationShares,
        uint256[] memory investorShares
    )external override onlyDelegation {
        uint256 delegationLength = investorDelegationShares.length;
        for (uint256 i = 0; i < delegationLength; i++) {
            if (address(investorDelegationShares[i]) == address(this)) {
                // construct calldata for increaseDelegatedShares call
                bytes memory message = abi.encodeWithSelector(
                    ITssRewardContract.increaseDelegatedShares.selector,
                    operator,
                    delegator,
                    investorShares[i]
                );

                // send call data into L2, hardcode address
                sendCrossDomainMessage(
                    address(0x4200000000000000000000000000000000000020),
                    2000000,
                    message
                );
            }
            break;
        }
    }

    function onDelegationWithdrawn(
        address delegator,
        address operator,
        IDelegationShare[] memory investorDelegationShares,
        uint256[] memory investorShares
    ) external override onlyDelegation {
        uint256 delegationLength = investorDelegationShares.length;
        for (uint256 i = 0; i < delegationLength; i++) {
            if (address(investorDelegationShares[i]) == address(this)) {
                // construct calldata for decreaseDelegatedShares call
                bytes memory message = abi.encodeWithSelector(
                    ITssRewardContract.decreaseDelegatedShares.selector,
                    operator,
                    delegator,
                    investorShares[i]
                );

                // send call data into L2, hardcode address
                sendCrossDomainMessage(
                    address(0x4200000000000000000000000000000000000020),
                    2000000,
                    message
                );
            }
            break;
        }
    }

    function setClaimer(
        address staker,
        address claimer
    ) external {
        require(msg.sender == staker, "msg sender is not the staker");
        bytes memory message = abi.encodeWithSelector(
            ITssRewardContract.setClaimer.selector,
            staker,
            claimer
        );
        // send call data into L2, hardcode address
        sendCrossDomainMessage(
            address(0x4200000000000000000000000000000000000020),
            2000000,
            message
        );
    }


    /**
     * @notice send quit request for the next election
     */
    function quitRequest() public nonReentrant {

        require(delegation.operatorShares(msg.sender, this) > 0, "do not have deposit");
        // when not in consensus period
        require(
            ITssGroupManager(tssGroupContract).memberExistInActive(operators[msg.sender]) ||
                ITssGroupManager(tssGroupContract).memberExistActive(operators[msg.sender]),
            "not at the inactive group or active group"
        );
        // is active member
        for (uint256 i = 0; i < quitRequestList.length; i++) {
            require(quitRequestList[i] != msg.sender, "already in quitRequestList");
        }
        quitRequestList.push(msg.sender);
    }

    /**
     * @notice return the quit list
     */
    function getQuitRequestList() public view returns (address[] memory) {
        return quitRequestList;
    }

    /**
     * @notice clear the quit list
     */
    function clearQuitRequestList() public onlyOwner {
        delete quitRequestList;
    }

    /**
     * @notice verify the slash message then slash
     * @param _messageBytes the message that abi encode by type SlashMsg
     * @param _sig the signature of the hash keccak256(_messageBytes)
     */
    function slashing(bytes calldata _messageBytes, bytes calldata _sig) public nonReentrant {
        SlashMsg memory message = abi.decode(_messageBytes, (SlashMsg));
        // verify tss member state not at jailed status
        require(!isJailed(message.jailNode), "the node already jailed");

        // have not slash before
        require(!slashRecord[message.batchIndex][message.jailNode], "already slashed");
        slashRecord[message.batchIndex][message.jailNode] = true;

        require(
            ITssGroupManager(tssGroupContract).verifySign(keccak256(_messageBytes), _sig),
            "signer not tss group pub key"
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
        bytes memory jailNodePubKey = operators[message.jailNode];
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
        uint256 deductedAmountShare;
        uint256 totalTransferShare;
        uint256 extraAmountShare;
        uint256 remainderShare;
        uint256 gainShare;
        uint256 _exIncomeShare = 0;
        // total slash slashAmount[slashType]
        // tssnodes get: gain = (slashAmount[slashType] - exIncome[slashType]) / tssnodes.length
        // sender get: remainder + _exIncome = (slashAmount[slashType] - exIncome[slashType]) % tssnodes.length + exIncome[slashType]
        // deductedAmount = tssnodes.length * gain + remainder + _exIncome = slashAmount[slashType]

        // check deposit > slashAmount, deduct slashAmount then
        // distribute additional tokens for the sender

        uint256 totalBalance = _tokenBalance();

        require(
            (delegation.operatorShares(deduction, this) * totalBalance) / totalShares >= slashAmount[slashType],
            "do not have enought shares"
        );
        // record total penalty
        deductedAmountShare = (slashAmount[slashType] * totalShares) / totalBalance;
        // record the sender's fixed additional income
        _exIncomeShare = (exIncome[slashType] * totalShares) / totalBalance;



        // record the deserving income for tss nodes
        extraAmountShare = deductedAmountShare - _exIncomeShare;
        // deserving income should subtract the remainder
        remainderShare = extraAmountShare % tssNodes.length;
        // record the gain for tss nodes
        gainShare = (extraAmountShare - remainderShare) / tssNodes.length;



        IDelegationShare[] memory delegationShares = new IDelegationShare[](1);
        delegationShares[0] = this;

        uint256[] memory delegationShareIndexes = new uint256[](1);
        delegationShareIndexes[0] = 0;

        uint256[] memory shareAmounts = new uint256[](1);
        shareAmounts[0] = _exIncomeShare + remainderShare;
        // sender get the fixed additional income and remainder
        TssDelegationManager(tssDelegationManagerContract).slashOperatorShares(deduction, msg.sender, delegationShares, delegationShareIndexes, shareAmounts);
        totalTransferShare = _exIncomeShare + remainderShare;

        // send gain to tss nodes
        for (uint256 i = 0; i < tssNodes.length; i++) {
            totalTransferShare += gainShare;
            shareAmounts[0] = gainShare;
            TssDelegationManager(tssDelegationManagerContract).slashOperatorShares(deduction, tssNodes[i], delegationShares, delegationShareIndexes, shareAmounts);
        }
        // The total transfer amount is the same as the deducted amount
        require(totalTransferShare == deductedAmountShare, "panic, calculation error");
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

        uint256 totalBalance = _tokenBalance();

        require((delegation.operatorShares(msg.sender, this) * totalBalance) / totalShares >= slashAmount[1], "Insufficient balance");
        ITssGroupManager(tssGroupContract).memberUnJail(operators[msg.sender]);
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
            .getTssMember(operators[user]);
        require(tssMember.publicKey.length == 64, "tss member not exist");
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

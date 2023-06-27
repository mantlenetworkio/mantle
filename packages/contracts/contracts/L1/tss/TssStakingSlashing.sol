// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/utils/math/SafeCast.sol";

import {DelegationShareBase} from "../delegation/DelegationShareBase.sol";
import {DelegationCallbackBase} from "../delegation/DelegationCallbackBase.sol";
import {IDelegationManager} from "../delegation/interfaces/IDelegationManager.sol";
import {IDelegationShare} from "../delegation/interfaces/IDelegation.sol";
import {IDelegation} from "../delegation/interfaces/IDelegation.sol";
import {CrossDomainEnabled} from "../../libraries/bridge/CrossDomainEnabled.sol";
import {ITssRewardContract} from "../../L2/predeploys/iTssRewardContract.sol";
import {TssDelegationManager} from "./delegation/TssDelegationManager.sol";
import {TssDelegation} from "./delegation/TssDelegation.sol";
import {WhiteList} from "../delegation/WhiteListBase.sol";
import {Lib_Address} from "../../libraries/utils/Lib_Address.sol";

import "./ITssGroupManager.sol";
import "./ITssStakingSlashing.sol";

contract TssStakingSlashing is
    Initializable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    IStakingSlashing,
    DelegationShareBase,
    DelegationCallbackBase,
    CrossDomainEnabled
{
    enum SlashType {
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
    //tss delegation address
    address public tssDelegationContract;
    // storage operator infos (key:staker address)
    mapping(address => bytes) public operators;

    // slashing parameter settings
    // record the quit request
    address[] public quitRequestList;
    // slashing amount of type uptime and animus (0:uptime, 1:animus)
    uint256[2] public slashAmount;
    // record the slash operate (map[batchIndex] -> (map[staker] -> slashed))
    mapping(uint256 => mapping(address => bool)) slashRecord;
    //EOA address
    address public regulatoryAccount;
    //msg sender => withdraw event
    mapping(address => bytes32) public withdrawalRoots;
    //msg sender => withdrawal
    mapping(address => IDelegationManager.QueuedWithdrawal) public withdrawals;
    //operator => stakers
    mapping(address => address[]) public stakers;
    //staker => operator
    mapping(address => address) public delegators;
    //operator => claimer
    mapping(address => address) public operatorClaimers;
    //claimer => operator
    mapping(address => address) public claimerOperators;
    bool public isSetParam;
    address public tssManager;


    /**
     * @notice slash tssnode
     * @param 0 slashed address
     * @param 1 slash type
     */
    event Slashing(address, SlashType);

    event WithdrawQueue(address,uint256);

    constructor()  CrossDomainEnabled(address(0)) {
        _disableInitializers();
    }

    /**
     * @notice initializes the contract setting and the deployer as the initial owner
     * @param _mantleToken mantle token contract address
     * @param _tssGroupContract address tss group manager contract address
     */
    function initialize(address _mantleToken,
        address _tssGroupContract,
        address _delegationManager,
        address _delegation,
        address _l1messenger,
        address _regulatoryAccount,
        address _tssManager
        ) public initializer {
        __Ownable_init();
        __ReentrancyGuard_init();
        underlyingToken = IERC20(_mantleToken);
        tssGroupContract = _tssGroupContract;
        tssDelegationManagerContract = _delegationManager;
        tssDelegationContract = _delegation;
        //initialize delegation
        delegationManager = IDelegationManager(_delegationManager);
        delegation = IDelegation(_delegation);
        messenger = _l1messenger;
        regulatoryAccount = _regulatoryAccount;
        tssManager = _tssManager;
    }

    /**
     * @notice change the mantle token and tssGroup contract address
     * @param _token the erc20 mantle token contract address
     */
    function setTokenAddress(address _token) public onlyOwner {
        require(_token != address(0),"Invalid address");
        underlyingToken = IERC20(_token);
    }

    function setTssGroupAddress(address _tssGroup) public onlyOwner{
        require(_tssGroup != address(0),"Invalid address");
        tssGroupContract = _tssGroup;
    }

    function setRegulatoryAccount(address _account) public onlyOwner {
        require(_account != address(0),"Invalid address");
        regulatoryAccount = _account;
    }

    function setTssManager(address _tssManager) public onlyOwner {
        require(_tssManager != address(0),"Invalid address");
        tssManager = _tssManager;
    }

    function setClaimer(
        address _operator,
        address _claimer
    ) external {
        require(msg.sender == _operator, "msg sender is diff with operator address");
        require(delegation.isOperator(msg.sender), "msg sender is not registered operator");
        require(claimerOperators[_claimer] == address(0), "the claimer has been used");
        if (operatorClaimers[_operator] != address(0)) {
            delete claimerOperators[operatorClaimers[_operator]];
        }
        operatorClaimers[_operator] = _claimer;
        claimerOperators[_claimer] = _operator;

        bytes memory message = abi.encodeWithSelector(
            ITssRewardContract.setClaimer.selector,
            _operator,
            _claimer
        );
        // send call data into L2, hardcode address
        sendCrossDomainMessage(
            address(0x4200000000000000000000000000000000000020),
            2000000,
            message
        );
    }

    /**
     * @notice set the slashing params (0 -> uptime , 1 -> animus)
     * @param _slashAmount the amount to be deducted for each type
     */
    function setSlashingParams(uint256[2] calldata _slashAmount)
        public
        onlyOwner
    {
        require(_slashAmount[1] > _slashAmount[0], "invalid param slashAmount, animus <= uptime");

        for (uint256 i = 0; i < 2; i++) {
            require(_slashAmount[i] > 0, "invalid amount");
            slashAmount[i] = _slashAmount[i];
        }
        isSetParam = true;
    }

    /**
     * @notice set the slashing params (0 -> uptime, 1 -> animus)
     */
    function getSlashingParams() public view returns (uint256[2] memory) {
        return slashAmount;
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
        require(tssManager == msg.sender,"TssStakingSlashing: msg.sender is not tssManager");
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
        require(isSetParam,"have not set the slash amount");
        bytes memory jailNodePubKey = operators[message.jailNode];
        if (message.slashType == SlashType.uptime) {
            // jail and transfer deposits
            ITssGroupManager(tssGroupContract).memberJail(jailNodePubKey);
            transformDeposit(message.jailNode, 0);
        } else if (message.slashType == SlashType.animus) {
            // remove the member and transfer deposits
            ITssGroupManager(tssGroupContract).memberJail(jailNodePubKey);
            transformDeposit(message.jailNode, 1);
        } else {
            revert("err type for slashing");
        }

    }

    /**
     * @notice distribute rewards to voters
     * @param deduction address of the punished
     * @param slashType the type to punished
     */
    function transformDeposit(
        address deduction,
        uint256 slashType
    ) internal {
        uint256 deductedAmountShare;

        uint256 totalBalance = _tokenBalance();

        require(
            (delegation.operatorShares(deduction, this) * totalBalance) / totalShares >= slashAmount[slashType],
            "do not have enought shares"
        );
        // record total penalty
        deductedAmountShare = (slashAmount[slashType] * totalShares) / totalBalance;

        uint256 operatorShare = delegation.operatorShares(deduction, this);

        IDelegationShare[] memory delegationShares = new IDelegationShare[](1);
        delegationShares[0] = this;

        uint256[] memory delegationShareIndexes = new uint256[](1);
        delegationShareIndexes[0] = 0;


        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = underlyingToken;

        address[] memory stakerS = stakers[deduction];
        for (uint256 i = 0; i < stakerS.length; i++){
            uint256 share = shares(stakerS[i]);
            uint256[] memory shareAmounts = new uint256[](1);
            shareAmounts[0] = deductedAmountShare * share / operatorShare;
            TssDelegationManager(tssDelegationManagerContract).slashShares(stakerS[i], regulatoryAccount, delegationShares,tokens, delegationShareIndexes, shareAmounts);
        }

    }

    /**
     * @notice set tss node status unjail
     */
    function unJail() public {
        // slashing params check
        require(isSetParam, "have not set the slash amount");
        require(isJailed(msg.sender), "An unjailed user doesn't need to call this method");

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

    function isCanOperator(address _addr) public returns (bool) {
        return TssDelegationManager(tssDelegationManagerContract).isCanOperator(_addr, this);
    }

    function deposit(uint256 amount) public returns (uint256) {
       uint256 shares = TssDelegationManager(tssDelegationManagerContract).depositInto(this, underlyingToken, amount, msg.sender);
       return shares;
    }

    function withdraw() external {
        require(delegation.isDelegated(msg.sender),"not delegator");

        require(
            withdrawalRoots[msg.sender] == bytes32(0),
            "msg sender already request withdraws"
        );

        uint256[] memory delegationIndexes = new uint256[](1);
        delegationIndexes[0] = 0;
        IDelegationShare[] memory delegationShares = new IDelegationShare[](1);
        delegationShares[0] = this;
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = underlyingToken;
        uint256[] memory sharesA = new uint256[](1);
        sharesA[0] = shares(msg.sender);
        uint256 nonce = TssDelegationManager(tssDelegationManagerContract).getWithdrawNonce(msg.sender);
        IDelegationManager.WithdrawerAndNonce memory withdrawerAndNonce = IDelegationManager.WithdrawerAndNonce({
            withdrawer: msg.sender,
            nonce: SafeCast.toUint96(nonce)
        });
        address operator = delegation.delegatedTo(msg.sender);

        IDelegationManager.QueuedWithdrawal memory queuedWithdrawal = IDelegationManager.QueuedWithdrawal({
            delegations: delegationShares,
            tokens: tokens,
            shares: sharesA,
            depositor: msg.sender,
            withdrawerAndNonce: withdrawerAndNonce,
            delegatedAddress: operator
        });
        withdrawals[msg.sender] = queuedWithdrawal;
        bytes32 withdrawRoot = TssDelegationManager(tssDelegationManagerContract).queueWithdrawal(msg.sender,delegationIndexes,delegationShares,tokens,sharesA,withdrawerAndNonce);
        withdrawalRoots[msg.sender] = withdrawRoot;
        emit WithdrawQueue(msg.sender, sharesA[0]);
    }

    function startWithdraw() external {
        require(
            withdrawalRoots[msg.sender] != bytes32(0),
            "msg sender must request withdraw first"
        );
        bytes32 withdrawRoot = withdrawalRoots[msg.sender];
        TssDelegationManager(tssDelegationManagerContract).startQueuedWithdrawalWaitingPeriod(withdrawRoot,msg.sender,0);
    }

    function canCompleteQueuedWithdrawal() external returns (bool) {

        require(
            withdrawalRoots[msg.sender] != bytes32(0),
            "msg sender did not request withdraws"
        );
        IDelegationManager.QueuedWithdrawal memory queuedWithdrawal = withdrawals[msg.sender];
        return delegationManager.canCompleteQueuedWithdrawal(queuedWithdrawal);
    }

    function completeWithdraw() external {

        require(
            withdrawalRoots[msg.sender] != bytes32(0),
            "msg sender did not request withdraws"
        );
        IDelegationManager.QueuedWithdrawal memory queuedWithdrawal = withdrawals[msg.sender];
        TssDelegationManager(tssDelegationManagerContract).completeQueuedWithdrawal(msg.sender, queuedWithdrawal, true);
        delete withdrawalRoots[msg.sender];
        delete withdrawals[msg.sender];
    }

    function registerAsOperator(bytes calldata _pubKey) external {
        require(msg.sender == Lib_Address.publicKeyToAddress(_pubKey), "public key not match");
        TssDelegation(tssDelegationContract).registerAsOperator(this, msg.sender);
        operators[msg.sender] = _pubKey;
    }

    function delegateTo(address _operator) external {
        TssDelegation(tssDelegationContract).delegateTo(_operator, msg.sender);
    }


    function onDelegationReceived(
        address delegator,
        address operator,
        IDelegationShare[] memory delegationShares,
        uint256[] memory investorShares
    )external override onlyDelegation {
        uint256 delegationLength = delegationShares.length;
        require(delegationLength == 1,"delegation only for tss");
        require(investorShares.length == 1,"delegation share only for tss");
        require(address(delegationShares[0]) == address(this),"must use current contract");
        if (delegators[delegator] == address(0)) {
            delegators[delegator] = operator;
            stakers[operator].push(delegator);
        }
    }

    function onDelegationWithdrawn(
        address delegator,
        address operator,
        IDelegationShare[] memory delegationShares,
        uint256[] memory investorShares
    ) external override onlyDelegation {
        uint256 delegationLength = delegationShares.length;
        require(delegationLength == 1,"delegation only for tss");
        require(investorShares.length == 1,"delegation share only for tss");
        require(address(delegationShares[0]) == address(this),"must use current contract");
        if (TssDelegationManager(tssDelegationManagerContract).getDelegationShares(delegator, delegationShares[0]) == investorShares[0]){
            address[] memory staker = stakers[operator];
            for (uint256 j = 0; j < staker.length; j++) {
                if (staker[j] == delegator) {
                    stakers[operator][j] = stakers[operator][staker.length -1];
                    stakers[operator].pop();
                    delete delegators[delegator];
                }
            }
        }
    }

}

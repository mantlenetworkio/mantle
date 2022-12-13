// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

import "../interfaces/IInvestmentManager.sol";
import "../interfaces/IEigenLayrDelegation.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IETHPOSDeposit.sol";
import "../interfaces/IEigenPod.sol";
import "../interfaces/IBeaconChainOracle.sol";

import "forge-std/Test.sol";



/**
 * @title The contract used for creating and managing EigenPods
 * @author Layr Labs, Inc.
 * @notice The main functionalities are:
 * - creating EigenPods
 * - staking for new validators on EigenPods
 * - keeping track of the balances of all validators of EigenPods, and their stake in EigenLayer
 * - withdrawing eth when withdrawals are initiated
 */
contract EigenPodManager is Initializable, IEigenPodManager 
{
    //TODO: change this to constant in prod
    IETHPOSDeposit immutable ethPOS;
    
    /// @notice Beacon proxy to which the EigenPods point
    IBeacon public immutable eigenPodBeacon;

    /// @notice EigenLayer's InvestmentManager contract
    IInvestmentManager public immutable investmentManager;

    /// @notice Oracle contract that provides updates to the beacon chain's state
    IBeaconChainOracle public beaconChainOracle;

    mapping(address => EigenPodInfo) public pods;

    event BeaconOracleUpdate(address newOracleAddress);

    modifier onlyEigenPod(address podOwner) {
        require(address(getPod(podOwner)) == msg.sender, "EigenPodManager.onlyEigenPod: not a pod");
        _;
    }

    modifier onlyInvestmentManager {
        require(msg.sender == address(investmentManager), "EigenPodManager.onlyEigenPod: not investmentManager");
        _;
    }

    modifier onlyInvestmentManagerOwner {
        require(msg.sender == Ownable(address(investmentManager)).owner(), "EigenPod.onlyInvestmentManagerOwner: not investment manager owner");
        _;
    }

    constructor(IETHPOSDeposit _ethPOS, IBeacon _eigenPodBeacon, IInvestmentManager _investmentManager) {
        ethPOS = _ethPOS;
        eigenPodBeacon = _eigenPodBeacon;
        investmentManager = _investmentManager;
        _disableInitializers();
    }

    function initialize(IBeaconChainOracle _beaconChainOracle) public initializer {
        beaconChainOracle = _beaconChainOracle;
        emit BeaconOracleUpdate(address(_beaconChainOracle));
    }

    /**
     * @notice Creates an EigenPod for the sender.
     * @dev Function will revert if the `msg.sender` already has an EigenPod.
     */
    function createPod() external {
        require(!hasPod(msg.sender), "EigenPodManager.createPod: Sender already has a pod");
        //deploy a pod if the sender doesn't have one already
        _deployPod();
    }

    /**
     * @notice Stakes for a new beacon chain validator on the sender's EigenPod. 
     * Also creates an EigenPod for the sender if they don't have one already.
     * @param pubkey The 48 bytes public key of the beacon chain validator.
     * @param signature The validator's signature of the deposit data.
     * @param depositDataRoot The root/hash of the deposit data for the validator's deposit.
     */
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable {
        IEigenPod pod = getPod(msg.sender);
        if(!hasPod(msg.sender)) {
            //deploy a pod if the sender doesn't have one already
            pod = _deployPod();
        }
        pod.stake{value: msg.value}(pubkey, signature, depositDataRoot);
    }

    /**
     * @notice Updates the beacon chain balance of the EigenPod, freezing the owner if they have overcommitted beacon chain ETH to EigenLayer.
     * @param podOwner The owner of the pod to udpate the balance of.
     * @param balanceToRemove The balance to remove before increasing, used when updating a validators balance.
     * @param balanceToAdd The balance to add after decreasing, used when updating a validators balance.
     * @dev Callable only by the `podOwner`'s EigenPod.
     */
    function updateBeaconChainBalance(address podOwner, uint64 balanceToRemove, uint64 balanceToAdd) external onlyEigenPod(podOwner) {
        uint128 newBalance = pods[podOwner].balance - balanceToRemove + balanceToAdd;
        pods[podOwner].balance = newBalance;
        /**
        * if the balance updates shows that the pod owner has more deposits into EigenLayer than beacon chain balance, freeze them
        * we also add the balance of the eigenPod in case withdrawals have occured so validator balances have been set to 0 
        * on the beacon chain the overall law is the amount InvestmentManager thinks is restaked <= balance of the withdrawal 
        * address + balance given from beacon chain state root if the investment manager ever thinks there is more 
        * restaked than there is, a freezing event is triggered
        */
        //TODO: add EigenPodManager as globally permissioned slashing contract
        if(pods[podOwner].depositedBalance > newBalance + msg.sender.balance) {
            investmentManager.slasher().freezeOperator(podOwner);
        }
    }

    /**
     * @notice Stakes beacon chain ETH into EigenLayer by adding BeaconChainETH shares to InvestmentManager.
     * @param podOwner The owner of the pod whose balance must be restaked.
     * @param amount The amount of beacon chain ETH to restake.
     * @dev Callable only by the `podOwner`'s EigenPod.
     */
    function depositBeaconChainETH(address podOwner, uint64 amount) external onlyEigenPod(podOwner) {
        //make sure that the podOwner hasn't over committed their stake, and deposit on their behalf
        require(pods[podOwner].depositedBalance + amount <= pods[podOwner].balance + address(getPod(podOwner)).balance, "EigenPodManager.depositBalanceIntoEigenLayer: cannot deposit more than balance");
        pods[podOwner].depositedBalance += amount;
        //deposit into InvestmentManager
        investmentManager.depositBeaconChainETH(podOwner, uint256(amount));
    }

    /**
     * @notice Withdraws ETH that has been withdrawn from the beacon chain from the EigenPod.
     * @param podOwner The owner of the pod whose balance must be withdrawn.
     * @param recipient The recipient of withdrawn ETH.
     * @param amount The amount of ETH to withdraw.
     * @dev Callable only by the InvestmentManager contract.
     */
    function withdrawBeaconChainETH(address podOwner, address recipient, uint256 amount) external onlyInvestmentManager {
        //subtract withdrawn amount from stake and balance
        pods[podOwner].depositedBalance = pods[podOwner].depositedBalance - uint128(amount);
        pods[podOwner].balance = pods[podOwner].balance - uint128(amount);
        getPod(podOwner).withdrawBeaconChainETH(recipient, amount);
    }

    /**
     * @notice Updates the oracle contract that provides the beacon chain state root
     * @param newBeaconChainOracle is the new oracle contract being pointed to
     * @dev Callable only by the owner of the InvestmentManager (i.e. governance).
     */
    function updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) external onlyInvestmentManagerOwner {
        beaconChainOracle = newBeaconChainOracle;
        emit BeaconOracleUpdate(address(newBeaconChainOracle));
    }


    // INTERNAL FUNCTIONS
    function _deployPod() internal returns (IEigenPod) {
        IEigenPod pod = 
            IEigenPod(
                Create2.deploy(
                    0, 
                    bytes32(uint256(uint160(msg.sender))), 
                    // set the beacon address to the eigenPodBeacon and initialize it
                    abi.encodePacked(
                        type(BeaconProxy).creationCode, 
                        abi.encode(eigenPodBeacon, abi.encodeWithSelector(IEigenPod.initialize.selector, IEigenPodManager(address(this)), msg.sender))
                    )
                )
            );
        return pod;
    }

    // VIEW FUNCTIONS
    /// @notice Returns the address of the `podOwner`'s EigenPod (whether it is deployed yet or not).
    function getPod(address podOwner) public view returns (IEigenPod) {
        return IEigenPod(
                Create2.computeAddress(
                    bytes32(uint256(uint160(podOwner))), //salt
                    keccak256(abi.encodePacked(
                        type(BeaconProxy).creationCode, 
                        abi.encode(eigenPodBeacon, abi.encodeWithSelector(IEigenPod.initialize.selector, IEigenPodManager(address(this)), podOwner))
                    )) //bytecode
                ));
    }

    /// @notice Returns 'true' if the `podOwner` has created an EigenPod, and 'false' otherwise.
    function hasPod(address podOwner) public view returns (bool) {
        return address(getPod(podOwner)).code.length > 0;
    }

    /// @notice returns the current EigenPodInfo for the `podOwner`'s EigenPod.
    function getPodInfo(address podOwner) external view returns (EigenPodInfo memory) {
        EigenPodInfo memory podInfo = pods[podOwner];
        return podInfo;
    }

    /// @notice Returns the latest beacon chain state root posted to the beaconChainOracle.
    function getBalance(address podOwner) external view returns (uint128) {
        return pods[podOwner].balance;
    }

    function getDepositedBalance(address podOwner) external view returns (uint128) {
        return pods[podOwner].depositedBalance;
    }

    function getBeaconChainStateRoot() external view returns(bytes32){
        return beaconChainOracle.getBeaconChainStateRoot();
    }
}
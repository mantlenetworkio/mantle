// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "../libraries/BeaconChainProofs.sol";
import "../libraries/BytesLib.sol";
import "../libraries/Endian.sol";
import "../interfaces/IETHPOSDeposit.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IEigenPod.sol";
import "../interfaces/IBeaconChainETHReceiver.sol";


/**
 * @title The implementation contract used for restaking beacon chain ETH on EigenLayer 
 * @author Layr Labs, Inc.
 * @notice The main functionalities are:
 * - creating new validators with their withdrawal credentials pointed to this contract
 * - proving from beacon chain state roots that withdrawal credentials are pointed to this contract
 * - proving from beacon chain state roots the balances of validators with their withdrawal credentials
 *   pointed to this contract
 * - updating aggregate balances in the EigenPodManager
 * - withdrawing eth when withdrawals are initiated
 */
contract EigenPod is IEigenPod, Initializable
{
    using BytesLib for bytes;

    //TODO: change this to constant in prod
    IETHPOSDeposit immutable ethPOS;

    /// @notice The single EigenPodManager for EigenLayer
    IEigenPodManager public eigenPodManager;

    /// @notice The owner of this EigenPod
    address public podOwner;

    /// @notice this is a mapping of validator keys to a Validator struct which holds info about the validator and their balances
    mapping(bytes32 => Validator) public validators;

    modifier onlyEigenPodManager {
        require(msg.sender == address(eigenPodManager), "EigenPod.InvestmentManager: not eigenPodManager");
        _;
    }


    constructor(IETHPOSDeposit _ethPOS) {
        ethPOS = _ethPOS;
        //TODO: uncomment for prod
        _disableInitializers();
    }

    /// @notice Used to initialize the pointers to contracts crucial to the pod's functionality, in beacon proxy construction from EigenPodManager
    function initialize(IEigenPodManager _eigenPodManager, address _podOwner) external initializer {
        eigenPodManager = _eigenPodManager;
        podOwner = _podOwner;
    }

    /// @notice Called by EigenPodManager when the owner wants to create another validator.
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable onlyEigenPodManager {
        // stake on ethpos
        ethPOS.deposit{value : msg.value}(pubkey, podWithdrawalCredentials(), signature, depositDataRoot);
    }

    /**
    * @notice This function verifies that the withdrawal credentials of the podOwner are pointed to
    * this contract.  It verifies the provided proof from the validator against the beacon chain state
    * root.
    * @param pubkey is the BLS public key for the validator.
    * @param proofs is
    * @param validatorFields are the fields of the "Validator Container", refer to consensus specs 
    * for details: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
     */
    function verifyCorrectWithdrawalCredentials(
        bytes calldata pubkey, 
        bytes calldata proofs, 
        bytes32[] calldata validatorFields
    ) external {
        //TODO: tailor this to production oracle
        bytes32 beaconStateRoot = eigenPodManager.getBeaconChainStateRoot();

        // get merklizedPubkey: https://github.com/prysmaticlabs/prysm/blob/de8e50d8b6bcca923c38418e80291ca4c329848b/beacon-chain/state/stateutil/sync_committee.root.go#L45
        bytes32 merklizedPubkey = sha256(abi.encodePacked(pubkey, bytes16(0)));

        require(validators[merklizedPubkey].status == VALIDATOR_STATUS.INACTIVE, "EigenPod.verifyCorrectWithdrawalCredentials: Validator not inactive");
        //verify validator proof
        BeaconChainProofs.verifyValidatorFields(
            beaconStateRoot,
            proofs,
            validatorFields
        );
        //require that the first field is the merkleized pubkey
        require(validatorFields[0] == merklizedPubkey, "EigenPod.verifyCorrectWithdrawalCredentials: Proof is not for provided pubkey");
        require(validatorFields[1] == podWithdrawalCredentials().toBytes32(0), "EigenPod.verifyCorrectWithdrawalCredentials: Proof is not for this EigenPod");
        //convert the balance field from 8 bytes of little endian to uint64 big endian 💪
        uint64 validatorBalance = Endian.fromLittleEndianUint64(validatorFields[2]);
        //update validator balance
        validators[merklizedPubkey].balance = validatorBalance;
        validators[merklizedPubkey].status = VALIDATOR_STATUS.ACTIVE;
        //update manager total balance for this pod
        //need to subtract zero and add the proven balance
        eigenPodManager.updateBeaconChainBalance(podOwner, 0, validatorBalance);
        eigenPodManager.depositBeaconChainETH(podOwner, validatorBalance);
    }

    function verifyBalanceUpdate(
        bytes calldata pubkey, 
        bytes calldata proofs, 
        bytes32[] calldata validatorFields
    ) external {
        //TODO: tailor this to production oracle
        bytes32 beaconStateRoot = eigenPodManager.getBeaconChainStateRoot();
        // get merklizedPubkey
        bytes32 merklizedPubkey = sha256(abi.encodePacked(pubkey, bytes16(0)));
        require(validators[merklizedPubkey].status == VALIDATOR_STATUS.ACTIVE, "EigenPod.verifyBalanceUpdate: Validator not active");
        //verify validator proof
        BeaconChainProofs.verifyValidatorFields(
            beaconStateRoot,
            proofs,
            validatorFields
        );
        //require that the first field is the merkleized pubkey
        require(validatorFields[0] == merklizedPubkey, "EigenPod.verifyBalanceUpdate: Proof is not for provided pubkey");
        //convert the balance field from 8 bytes of little endian to uint64 big endian 💪
        uint64 validatorBalance = Endian.fromLittleEndianUint64(validatorFields[2]);
        uint64 prevValidatorBalance = validators[merklizedPubkey].balance;
        //update validator balance
        validators[merklizedPubkey].balance = validatorBalance;
        //update manager total balance for this pod
        //need to subtract previous proven balance and add the current proven balance
        eigenPodManager.updateBeaconChainBalance(podOwner, prevValidatorBalance, validatorBalance);
        if(prevValidatorBalance < validatorBalance){
            eigenPodManager.depositBeaconChainETH(podOwner, validatorBalance - prevValidatorBalance);
        }
        
    }

    /**
     * @notice Transfers ether balance of this contract to the specified recipient address
     * @notice Called by EigenPodManager to withdrawBeaconChainETH that has been added to its balance due to a withdrawal from the beacon chain.
     * @dev Called during withdrawal or slashing.
     */
    function withdrawBeaconChainETH(
        address recipient,
        uint256 amount
    )
        external
        onlyEigenPodManager
    {
        //transfer ETH directly from pod to msg.sender 
        IBeaconChainETHReceiver(recipient).receiveBeaconChainETH{value: amount}();
    }
    //if you've been slashed on the Beacon chain, you can add balance to your pod to avoid getting slashed
    function topUpPodBalance() external payable {}

    // INTERNAL FUNCTIONS
    function podWithdrawalCredentials() internal view returns(bytes memory) {
        return abi.encodePacked(bytes1(uint8(1)), bytes11(0), address(this));
    }
}
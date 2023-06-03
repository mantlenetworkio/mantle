// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../delegation/Delegation.sol";
import "../../delegation/WhiteListBase.sol";


/**
 * @title The primary delegation contract.
 * @notice  This is the contract for delegation. The main functionalities of this contract are
 * - for enabling any staker to register as a delegate and specify the delegation terms it has agreed to
 * - for enabling anyone to register as an operator
 * - for a registered staker to delegate its stake to the operator of its agreed upon delegation terms contract
 * - for a staker to undelegate its assets
 * - for anyone to challenge a staker's claim to have fulfilled all its obligation before undelegation
 */
contract TssDelegation is Delegation {


    address public stakingSlash;




    // INITIALIZING FUNCTIONS
    constructor(IDelegationManager _delegationManager)
    Delegation(_delegationManager)
    {
        _disableInitializers();
    }


    function initializeT(
        address _stakingSlashing,
        address initialOwner
    ) external initializer {
        DOMAIN_SEPARATOR = keccak256(abi.encode(DOMAIN_TYPEHASH, bytes("Mantle"), block.chainid, address(this)));
        stakingSlash = _stakingSlashing;
         _transferOwnership(initialOwner);
    }

    modifier onlyStakingSlash() {
        require(msg.sender == stakingSlash, "contract call is not staking slashing");
        _;
    }

    function setStakingSlash(address _address) public onlyOwner {
        stakingSlash = _address;
    }

    function registerAsOperator(IDelegationCallback dt, address sender) external whitelistOnly(sender) onlyStakingSlash {

        require(
            address(delegationCallback[sender]) == address(0),
            "Delegation.registerAsOperator: Delegate has already registered"
        );
        // store the address of the delegation contract that the operator is providing.
        delegationCallback[sender] = dt;
        _delegate(sender, sender);
        emit RegisterOperator(address(dt), sender);
    }

    function delegateTo(address operator, address staker) external onlyStakingSlash whenNotPaused {
        _delegate(staker, operator);
    }


}

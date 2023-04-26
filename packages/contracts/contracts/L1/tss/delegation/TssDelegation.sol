// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../delegation/Delegation.sol";


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
    // INITIALIZING FUNCTIONS
    constructor(IDelegationManager _delegationManager)
    Delegation(_delegationManager)
    {
        _disableInitializers();
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../delegation/DelegationSlasher.sol";

/**
 * @title The primary 'slashing' contract.
 * @notice This contract specifies details on slashing. The functionalities are:
 * - adding contracts who have permission to perform slashing,
 * - revoking permission for slashing from specified contracts,
 * - calling investManager to do actual slashing.
 */
contract TssDelegationSlasher is DelegationSlasher {
    constructor(IDelegationManager _delegationManager, IDelegation _delegation)
    DelegationSlasher(_delegationManager, _delegation)
    {
        _disableInitializers();
    }
}

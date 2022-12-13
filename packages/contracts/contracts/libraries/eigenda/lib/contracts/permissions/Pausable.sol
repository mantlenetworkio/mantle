// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "../interfaces/IPauserRegistry.sol";

// import "forge-std/Test.sol";

/**
 * @title Adds pausability to a contract.
 * @author Layr Labs, Inc.
 * @notice Contracts that inherit from this contract define their own `pause` and `unpause` (and/or related) functions.
 * These functions should be permissioned as "onlyPauser" which defers to a `PauserRegistry` for determining access control.
 */
contract Pausable {
    /// @notice Address of the `PauserRegistry` contract that this contract defers to for determining access control (for pausing).
    IPauserRegistry public pauserRegistry;

    /// @dev whether or not the contract is currently paused
    bool private _paused;

    /**
     * @dev Emitted when the pause is triggered by `account`.
     */
    event Paused(address account);

    /**
     * @dev Emitted when the pause is lifted by `account`.
     */
    event Unpaused(address account);

    modifier onlyPauser() {
        require(msg.sender == pauserRegistry.pauser(), "msg.sender is not permissioned as pauser");
        _;
    }

    modifier onlyUnpauser() {
        require(msg.sender == pauserRegistry.unpauser(), "msg.sender is not permissioned as unpauser");
        _;
    }

    modifier whenNotPaused() {
        _requireNotPaused();
        _;
    }

    function _initializePauser(IPauserRegistry _pauserRegistry) internal {
        require(
            address(pauserRegistry) == address(0) && address(_pauserRegistry) != address(0),
            "Pausable._initializePauser: _initializePauser() can only be called once"
        );

        _paused = false;
        pauserRegistry = _pauserRegistry;
    }

    /**
     * @notice This function is used to pause an EigenLayer/DataLayer
     * contract functionality.  It is permissioned to the "PAUSER"
     * address, which is a low threshold multisig.
     */
    function pause() external onlyPauser {
        _paused = true;
        emit Paused(msg.sender);
    }

    /**
     * @notice This function is used to unpause an EigenLayer/DataLayer
     * contract functionality.  It is permissioned to the "UNPAUSER"
     * address, which is a reputed committee controlled, high threshold
     * multisig.
     */
    function unpause() external onlyUnpauser {
        _paused = false;
        emit Unpaused(msg.sender);
    }

    /**
     * @dev Returns true if the contract is paused, and false otherwise.
     */
    function paused() public view virtual returns (bool) {
        return _paused;
    }

    /**
     * @dev Throws if the contract is paused.
     */
    function _requireNotPaused() internal view virtual {
        require(!paused(), "Pausable: paused");
    }
}

// SPDX-License-Identifier: Apache-2.0

/*
 * Modifications Copyright 2022, Specular contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./libraries/Errors.sol";

// Exists only to reduce size of Rollup contract (maybe revert since Rollup fits under optimized compilation).
contract AssertionMap is Initializable {
    error ChildInboxSizeMismatch();

    error SiblingStateHashExists();

    struct Assertion {
        bytes32 stateHash; // Hash of execution state associated with assertion (see `RollupLib.stateHash`)
        uint256 inboxSize; // Inbox size this assertion advanced to
        uint256 parent; // Parent assertion ID
        uint256 deadline; // Confirmation deadline (L1 block timestamp)
        uint256 proposalTime; // L1 block number at which assertion was proposed
        // Staking state
        uint256 numStakers; // total number of stakers that have ever staked on this assertion. increasing only.
        // Child state
        uint256 childInboxSize; // child assertion inbox state
    }

    struct AssertionState {
        mapping(address => bool) stakers; // all stakers that have ever staked on this assertion.
        mapping(bytes32 => bool) childStateHashes; // child assertion vm hashes
    }

    mapping(uint256 => Assertion) public assertions;
    mapping(uint256 => AssertionState) private assertionStates; // mapping from assertionID to assertion state
    address public rollupAddress;

    modifier rollupOnly() {
        if (msg.sender != rollupAddress) {
            revert NotRollup(msg.sender, rollupAddress);
        }
        _;
    }

    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {}

    function setRollupAddress(address _rollupAddress) public {
        require(
            address(rollupAddress) == address(0),
            "rollupAddress already initialized."
        );
        require(_rollupAddress != address(0), "ZERO_ADDRESS");
        rollupAddress = _rollupAddress;
    }

    function getStateHash(uint256 assertionID) external view returns (bytes32) {
        return assertions[assertionID].stateHash;
    }

    function getInboxSize(uint256 assertionID) external view returns (uint256) {
        return assertions[assertionID].inboxSize;
    }

    function getParentID(uint256 assertionID) external view returns (uint256) {
        return assertions[assertionID].parent;
    }

    function getDeadline(uint256 assertionID) external view returns (uint256) {
        return assertions[assertionID].deadline;
    }

    function getProposalTime(uint256 assertionID) external view returns (uint256) {
        return assertions[assertionID].proposalTime;
    }

    function getNumStakers(uint256 assertionID) external view returns (uint256) {
        return assertions[assertionID].numStakers;
    }

    function isStaker(uint256 assertionID, address stakerAddress) external view returns (bool) {
        return assertionStates[assertionID].stakers[stakerAddress];
    }

    function createAssertion(
        uint256 assertionID,
        bytes32 stateHash,
        uint256 inboxSize,
        uint256 parentID,
        uint256 deadline
    ) external rollupOnly {
        Assertion storage parentAssertion = assertions[parentID];
        AssertionState storage parentAssertionState = assertionStates[parentID];
        // Child assertions must have same inbox size
        uint256 parentChildInboxSize = parentAssertion.childInboxSize;
        if (parentChildInboxSize == 0) {
            parentAssertion.childInboxSize = inboxSize;
        } else {
            if (inboxSize != parentChildInboxSize) {
                revert("ChildInboxSizeMismatch");
            }
        }
        if (parentAssertionState.childStateHashes[stateHash]) {
            revert("SiblingStateHashExists");
        }

        parentAssertionState.childStateHashes[stateHash] = true;

        assertions[assertionID] = Assertion(
            stateHash,
            inboxSize,
            parentID,
            deadline,
            block.number, // proposal time
            0, // numStakers
            0 // childInboxSize
        );
    }

    function stakeOnAssertion(uint256 assertionID, address stakerAddress) external rollupOnly {
        Assertion storage assertion = assertions[assertionID];
        assertionStates[assertionID].stakers[stakerAddress] = true;
        assertion.numStakers++;
    }

    function deleteAssertion(uint256 assertionID) external rollupOnly {
        delete assertions[assertionID];
    }

    function deleteAssertionForBatch(uint256 assertionID) external rollupOnly {
        bytes32 stateHash = assertions[assertionID].stateHash;
        uint256 parentID = assertions[assertionID].parent;
        delete assertions[assertionID];
        assertions[parentID].childInboxSize = 0;
        assertionStates[parentID].childStateHashes[stateHash] = false;
    }
}

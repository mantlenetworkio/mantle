// SPDX-License-Identifier: Apache-2.0

/*
 * Modifications Copyright 2022, Specular contributors
 *
 * This file was changed in accordance to Apache License, Version 2.0.
 *
 * Copyright 2021, Offchain Labs, Inc.
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

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "hardhat/console.sol";

import "./challenge/Challenge.sol";
import "./challenge/ChallengeLib.sol";
import "./AssertionMap.sol";
import "./IRollup.sol";
import "./RollupLib.sol";
import "./WhiteList.sol";
import "./verifier/IVerifier.sol";
import {Lib_AddressResolver} from "../../libraries/resolver/Lib_AddressResolver.sol";
import {Lib_AddressManager} from "../../libraries/resolver/Lib_AddressManager.sol";

abstract contract RollupBase is IRollup, Initializable {
    // Config parameters
    uint256 public confirmationPeriod; // number of L1 blocks
    uint256 public challengePeriod; // number of L1 blocks
    uint256 public minimumAssertionPeriod; // number of L1 blocks
    uint256 public baseStakeAmount; // number of stake tokens

    IERC20 public stakeToken;
    AssertionMap public override assertions;
    IVerifierEntry public verifier;

    struct Staker {
        bool isStaked;
        uint256 amountStaked;
        uint256 assertionID; // latest staked assertion ID
        address currentChallenge; // address(0) if none
    }

    struct Zombie {
        address stakerAddress;
        uint256 lastAssertionID;
    }

    struct ChallengeCtx {
        bool completed;
        address challengeAddress;
        address defenderAddress;
        address challengerAddress;
        uint256 defenderAssertionID;
        uint256 challengerAssertionID;
    }
}

contract Rollup is Lib_AddressResolver, RollupBase, Whitelist {
    modifier stakedOnly() {
        if (!isStaked(msg.sender)) {
            revert("NotStaked");
        }
        _;
    }

    // Assertion state
    uint256 public lastResolvedAssertionID;
    uint256 public lastConfirmedAssertionID;
    uint256 public lastCreatedAssertionID;

    // Staking state
    uint256 public numStakers; // current total number of stakers
    mapping(address => Staker) public stakers; // mapping from staker addresses to corresponding stakers
    mapping(address => uint256) public withdrawableFunds; // mapping from addresses to withdrawable funds (won in challenge)
    Zombie[] public zombies; // stores stakers that lost a challenge
    ChallengeCtx public challengeCtx;  // stores challenge context

    constructor() Lib_AddressResolver(address(0)) {
        _disableInitializers();
    }

    function initialize(
        address _owner,
        address _verifier,
        address _stakeToken,
        address _libAddressManager,
        address _assertionMap,
        uint256 _confirmationPeriod,
        uint256 _challengePeriod,
        uint256 _minimumAssertionPeriod,
        uint256 _baseStakeAmount,
        bytes32 _initialVMhash,
        address[] calldata whitelists
    ) public initializer {
        if (_owner == address(0) || _verifier == address(0)) {
            revert("ZeroAddress");
        }
        owner = _owner;
        stakeToken = IERC20(_stakeToken);
        verifier = IVerifierEntry(_verifier);

        if (address(libAddressManager) != address(0)) {
            revert("RedundantInitialized");
        }
        libAddressManager = Lib_AddressManager(_libAddressManager);

        if (address(assertions) != address(0)) {
            revert("RedundantInitialized");
        }
        assertions = AssertionMap(_assertionMap);

        confirmationPeriod = _confirmationPeriod;
        challengePeriod = _challengePeriod; // TODO: currently unused.
        minimumAssertionPeriod = _minimumAssertionPeriod;
        baseStakeAmount = _baseStakeAmount;

        assertions.setRollupAddress(address(this));
        lastResolvedAssertionID = 0;
        lastConfirmedAssertionID = 0;
        lastCreatedAssertionID = 0;

        assertions.createAssertion(
            lastResolvedAssertionID, // assertionID
            _initialVMhash,
            0, // inboxSize (genesis)
            0, // parentID
            block.number // deadline (unchallengeable)
        );

        for (uint i = 0; i < whitelists.length; i++) {
            whitelist[whitelists[i]] = true;
        }
    }

    /// @inheritdoc IRollup
    function isStaked(address addr) public view override returns (bool) {
        return stakers[addr].isStaked;
    }

    /// @inheritdoc IRollup
    function currentRequiredStake() public view override returns (uint256) {
        return baseStakeAmount;
    }

    /// @inheritdoc IRollup
    function confirmedInboxSize() public view override returns (uint256) {
        return assertions.getInboxSize(lastConfirmedAssertionID);
    }

    /// @inheritdoc IRollup
    function stake() external payable override whitelistOnly {
        if (isStaked(msg.sender)) {
            stakers[msg.sender].amountStaked += msg.value;
        } else {
            if (msg.value < baseStakeAmount) {
                revert("InsufficientStake");
            }
            stakers[msg.sender] = Staker(true, msg.value, 0, address(0));
            numStakers++;
            stakeOnAssertion(msg.sender, lastConfirmedAssertionID);
        }
    }

    /// @inheritdoc IRollup
    function unstake(uint256 stakeAmount) external override {
        requireStaked(msg.sender);
        // Require that staker is staked on a confirmed assertion.
        Staker storage staker = stakers[msg.sender];
        if (staker.assertionID > lastConfirmedAssertionID) {
            revert("StakedOnUnconfirmedAssertion");
        }
        if (stakeAmount > staker.amountStaked - currentRequiredStake()) {
            revert("InsufficientStake");
        }
        staker.amountStaked -= stakeAmount;
        // Note: we don't need to modify assertion state because you can only unstake from a confirmed assertion.
        (bool success,) = msg.sender.call{value: stakeAmount}("");
        if (!success) revert("TransferFailed");
    }

    /// @inheritdoc IRollup
    function removeStake(address stakerAddress) external override {
        requireStaked(stakerAddress);
        // Require that staker is staked on a confirmed assertion.
        Staker storage staker = stakers[stakerAddress];
        if (staker.assertionID > lastConfirmedAssertionID) {
            revert("StakedOnUnconfirmedAssertion");
        }
        deleteStaker(stakerAddress);
        // Note: we don't need to modify assertion state because you can only unstake from a confirmed assertion.
        (bool success,) = stakerAddress.call{value: staker.amountStaked}("");
        if (!success) revert("TransferFailed");
    }

    /// @inheritdoc IRollup
    function advanceStake(uint256 assertionID) external override stakedOnly {
        Staker storage staker = stakers[msg.sender];
        if (assertionID <= staker.assertionID && assertionID > lastCreatedAssertionID) {
            revert("AssertionOutOfRange");
        }
        // TODO: allow arbitrary descendant of current staked assertionID, not just child.
        if (staker.assertionID != assertions.getParentID(assertionID)) {
            revert("ParentAssertionUnstaked");
        }
        stakeOnAssertion(msg.sender, assertionID);
    }

    /// @inheritdoc IRollup
    function withdraw() external override {
        uint256 withdrawableFund = withdrawableFunds[msg.sender];
        withdrawableFunds[msg.sender] = 0;
        (bool success,) = msg.sender.call{value: withdrawableFund}("");
        if (!success) revert("TransferFailed");
    }

    /// @inheritdoc IRollup
    function createAssertion(
        bytes32 vmHash,
        uint256 inboxSize
    ) public override stakedOnly {
        uint256 parentID = stakers[msg.sender].assertionID;
        // Require that enough time has passed since the last assertion.
        if (block.number - assertions.getProposalTime(parentID) < minimumAssertionPeriod) {
            revert("MinimumAssertionPeriodNotPassed");
        }
        // Require that the assertion at least includes one transaction
        if (inboxSize <= assertions.getInboxSize(parentID)) {
            revert("EmptyAssertion");
        }

        // Initialize assertion.
        lastCreatedAssertionID++;
        emit AssertionCreated(lastCreatedAssertionID, msg.sender, vmHash, inboxSize);
        assertions.createAssertion(
            lastCreatedAssertionID, vmHash, inboxSize, parentID, newAssertionDeadline()
        );

        // Update stake.
        stakeOnAssertion(msg.sender, lastCreatedAssertionID);
    }

    /// @inheritdoc IRollup
    function createAssertionWithStateBatch(
        bytes32 vmHash,
        uint256 inboxSize,
        bytes32[] calldata _batch,
        uint256 _shouldStartAtElement,
        bytes calldata _signature
        ) external override stakedOnly { // todo batch submitter only
        // permissions only allow rollup proposer to submit assertion, only allow RollupContract to append new batch
        require(msg.sender == resolve("BVM_Rolluper"), "msg.sender is not rollup proposer, can't append batch");
        // create assertion
        createAssertion(vmHash, inboxSize);
        // append state batch
        address scc = resolve("StateCommitmentChain");
        (bool success, ) = scc.call(
            abi.encodeWithSignature("appendStateBatch(bytes32[],uint256,bytes)", _batch, _shouldStartAtElement, _signature)
        );
        require(success, "scc append state batch failed, revert all");
    }

    function challengeAssertion(address[2] calldata players, uint256[2] calldata assertionIDs)
        external
        override
        returns (address)
    {
        uint256 defenderAssertionID = assertionIDs[0];
        uint256 challengerAssertionID = assertionIDs[1];
        // Require IDs ordered and in-range.
        if (defenderAssertionID >= challengerAssertionID) {
            revert("WrongOrder");
        }
        if (challengerAssertionID > lastCreatedAssertionID) {
            revert("UnproposedAssertion");
        }
        if (lastConfirmedAssertionID >= defenderAssertionID) {
            revert("AssertionAlreadyResolved");
        }
        // Require that players have attested to sibling assertions.
        uint256 parentID = assertions.getParentID(defenderAssertionID);
        if (parentID != assertions.getParentID(challengerAssertionID)) {
            revert("DifferentParent");
        }

        // Require that neither player is currently engaged in a challenge.
        address defender = players[0];
        address challenger = players[1];
        requireUnchallengedStaker(defender);
        requireUnchallengedStaker(challenger);

        // TODO: Calculate upper limit for allowed node proposal time.

        // Initialize challenge.
        Challenge challenge = new Challenge();
        address challengeAddr = address(challenge);
        stakers[challenger].currentChallenge = challengeAddr;
        stakers[defender].currentChallenge = challengeAddr;

        challengeCtx = ChallengeCtx(false,challengeAddr,defender,challenger,defenderAssertionID,challengerAssertionID);
        emit AssertionChallenged(defenderAssertionID, challengeAddr);
        uint256 inboxSize = assertions.getInboxSize(parentID);
        bytes32 parentStateHash = assertions.getStateHash(parentID);
        bytes32 defenderStateHash = assertions.getStateHash(defenderAssertionID);
        challenge.initialize(
            defender,
            challenger,
            verifier,
            address(this),
            inboxSize,
            parentStateHash,
            defenderStateHash
        );
        return challengeAddr;
    }

    /// @inheritdoc IRollup
    function confirmFirstUnresolvedAssertion() external override {
        if (lastResolvedAssertionID >= lastCreatedAssertionID) {
            revert("NoUnresolvedAssertion");
        }

        // (1) there is at least one staker, and
        if (numStakers <= 0) revert("NoStaker");

        uint256 lastUnresolvedID = lastResolvedAssertionID + 1;

        // (2) challenge period has passed
        if (block.timestamp < assertions.getDeadline(lastUnresolvedID)) {
            revert("ChallengePeriodPending");
        }

        // (3) predecessor has been confirmed
        if (assertions.getParentID(lastUnresolvedID) != lastConfirmedAssertionID) {
            revert("InvalidParent");
        }

        // Remove old zombies
        // removeOldZombies();

        // (4) all stakers are staked on the block.
//        if (assertions.getNumStakers(lastUnresolvedID) != countStakedZombies(lastUnresolvedID) + numStakers) {
//            revert NotAllStaked();
//        }

        // Confirm assertion.
        // assertions.deleteAssertion(lastConfirmedAssertionID);
        lastResolvedAssertionID++;
        lastConfirmedAssertionID = lastResolvedAssertionID;
        emit AssertionConfirmed(lastResolvedAssertionID);
    }

    /// @inheritdoc IRollup
    function rejectFirstUnresolvedAssertion(address stakerAddress) external override {
        if (lastResolvedAssertionID >= lastCreatedAssertionID) {
            revert("NoUnresolvedAssertion");
        }

        uint256 firstUnresolvedAssertionID = lastResolvedAssertionID + 1;

        // First case - parent of first unresolved is last confirmed (`if` condition below). e.g.
        // [1] <- [3]           | valid chain ([1] is last confirmed, [3] is stakerAddress's unresolved assertion)
        //  ^---- [2]           | invalid chain ([2] is firstUnresolved)
        // Second case (trivial) - parent of first unresolved is not last confirmed. i.e.:
        //   parent is previous rejected, e.g.
        //   [1] <- [4]           | valid chain ([1] is last confirmed, [4] is stakerAddress's unresolved assertion)
        //   [2] <- [3]           | invalid chain ([3] is firstUnresolved)
        //   OR
        //   parent is previous confirmed, e.g.
        //   [1] <- [2] <- [4]    | valid chain ([2] is last confirmed, [4] is stakerAddress's unresolved assertion)
        //    ^---- [3]           | invalid chain ([3] is firstUnresolved)
        if (assertions.getParentID(firstUnresolvedAssertionID) == lastConfirmedAssertionID) {
            // 1a. challenge period has passed.
            if (block.timestamp < assertions.getDeadline(firstUnresolvedAssertionID)) {
                revert("ChallengePeriodPending");
            }
            // 1b. at least one staker exists (on a sibling)
            // - stakerAddress is indeed a staker
            requireStaked(stakerAddress);
            // - staker's assertion can't be a ancestor of firstUnresolved (because staker's assertion is also unresolved)
            if (stakers[stakerAddress].assertionID < firstUnresolvedAssertionID) {
                revert("AssertionAlreadyResolved");
            }
            // - staker's assertion can't be a descendant of firstUnresolved (because staker has never staked on firstUnresolved)
            if (assertions.isStaker(firstUnresolvedAssertionID, stakerAddress)) {
                revert("StakerStakedOnTarget");
            }
            // If a staker is staked on an assertion that is neither an ancestor nor a descendant of firstUnresolved, it must be a sibling, QED

            // 1c. no staker is staked on this assertion
            // removeOldZombies();
            if (assertions.getNumStakers(firstUnresolvedAssertionID) != countStakedZombies(firstUnresolvedAssertionID))
            {
                revert("StakersPresent");
            }
        }

        // Reject assertion.
        lastResolvedAssertionID++;
        emit AssertionRejected(lastResolvedAssertionID);
        assertions.deleteAssertion(lastResolvedAssertionID);
    }

    /// @inheritdoc IRollup
    function completeChallenge(address winner, address loser) external override {
        requireStaked(loser);

        address challenge = getChallenge(winner, loser);
        if (msg.sender != challenge) {
            revert("NotChallenge");
        }
        uint256 amountWon;
        uint256 loserStake = stakers[loser].amountStaked;
        uint256 winnerStake = stakers[winner].amountStaked;
        if (loserStake > baseStakeAmount) {
            // If loser has a higher stake than the winner, refund the difference.
            // Loser gets deleted anyways, so maybe unnecessary to set amountStaked.
            // stakers[loser].amountStaked = winnerStake;
            withdrawableFunds[loser] += (loserStake - baseStakeAmount);
            amountWon = baseStakeAmount;
        } else {
            amountWon = loserStake;
        }
        // Reward the winner with half the remaining stake
        stakers[winner].amountStaked += amountWon; // why +stake instead of +withdrawable?
        stakers[winner].currentChallenge = address(0);
        // Turning loser into zombie renders the loser's remaining stake inaccessible.
        uint256 assertionID = stakers[loser].assertionID;
        deleteStaker(loser);
        // Track as zombie so we can account for it during assertion resolution.
        zombies.push(Zombie(loser, assertionID));
        challengeCtx.completed = true;
    }

    /**
     * @notice Updates staker and assertion metadata.
     * @param stakerAddress Address of existing staker.
     * @param assertionID ID of existing assertion to stake on.
     */
    function stakeOnAssertion(address stakerAddress, uint256 assertionID) private {
        stakers[stakerAddress].assertionID = assertionID;
        assertions.stakeOnAssertion(assertionID, stakerAddress);
        emit StakerStaked(stakerAddress, assertionID);
    }

    /**
     * @notice Deletes the staker from global state. Does not touch assertion staker state.
     * @param stakerAddress Address of the staker to delete
     */
    function deleteStaker(address stakerAddress) private {
        numStakers--;
        delete stakers[stakerAddress];
    }

    /**
     * @notice Checks to see whether the two stakers are in the same challenge
     * @param staker1Address Address of the first staker
     * @param staker2Address Address of the second staker
     * @return Address of the challenge that the two stakers are in
     */
    function getChallenge(address staker1Address, address staker2Address) private view returns (address) {
        Staker storage staker1 = stakers[staker1Address];
        Staker storage staker2 = stakers[staker2Address];
        address challenge = staker1.currentChallenge;
        if (challenge == address(0)) {
            revert("NotInChallenge");
        }
        if (challenge != staker2.currentChallenge) {
            revert("InDifferentChallenge");
        }
        return challenge;
    }

    function newAssertionDeadline() private returns (uint256) {
        // TODO: account for prev assertion, gas
        // return block.number + confirmationPeriod;
        address scc = resolve("StateCommitmentChain");
        (bool success, bytes memory data) = scc.call(
            abi.encodeWithSignature("FRAUD_PROOF_WINDOW()")
        );
        uint256 confirmationWindow = uint256(bytes32(data));
        return block.timestamp + confirmationWindow;
    }

    // *****************
    // zombie processing
    // *****************

    /**
     * @notice Removes any zombies whose latest stake is earlier than the first unresolved assertion.
     * @dev Uses pop() instead of delete to prevent gaps, although order is not preserved
     */
    // function removeOldZombies() private {
    // }

    /**
     * @notice Counts the number of zombies staked on an assertion.
     * @dev O(n), where n is # of zombies (but is expected to be small).
     * This function could be uncallable if there are too many zombies. However,
     * removeOldZombies() can be used to remove any zombies that exist so that this
     * will then be callable.
     * @param assertionID The assertion on which to count staked zombies
     * @return The number of zombies staked on the assertion
     */
    function countStakedZombies(uint256 assertionID) private view returns (uint256) {
        uint256 numStakedZombies = 0;
        for (uint256 i = 0; i < zombies.length; i++) {
            if (assertions.isStaker(assertionID, zombies[i].stakerAddress)) {
                numStakedZombies++;
            }
        }
        return numStakedZombies;
    }

    // ************
    // requirements
    // ************

    function requireStaked(address stakerAddress) private view {
        if (!isStaked(stakerAddress)) {
            revert("NotStaked");
        }
    }

    function requireUnchallengedStaker(address stakerAddress) private view {
        requireStaked(stakerAddress);
        if (stakers[stakerAddress].currentChallenge != address(0)) {
            revert("ChallengedStaker");
        }
    }
}

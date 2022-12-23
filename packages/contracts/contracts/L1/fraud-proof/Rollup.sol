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

import "./challenge/Challenge.sol";
import "./challenge/ChallengeLib.sol";
import "./AssertionMap.sol";
import "./IRollup.sol";
import "./RollupLib.sol";
import "./verifier/IVerifier.sol";
import {Lib_AddressResolver} from "../../libraries/resolver/Lib_AddressResolver.sol";
import {IStateCommitmentChain} from "../rollup/IStateCommitmentChain.sol";

abstract contract RollupBase is IRollup, Initializable, Lib_AddressResolver {
    // Config parameters
    uint256 public confirmationPeriod; // number of L1 blocks
    uint256 public challengePeriod; // number of L1 blocks
    uint256 public minimumAssertionPeriod; // number of L1 blocks
    uint256 public maxGasPerAssertion; // L2 gas
    uint256 public baseStakeAmount; // number of stake tokens

    address public owner;
    IERC20 public stakeToken;
    AssertionMap public override assertions;
    IVerifier public verifier;

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
}

contract Rollup is RollupBase {
    modifier stakedOnly() {
        require(isStaked(msg.sender));
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

    function initialize(
        address _owner,
        address _verifier,
        address _stakeToken,
        uint256 _confirmationPeriod,
        uint256 _challengePeriod,
        uint256 _minimumAssertionPeriod,
        uint256 _maxGasPerAssertion,
        uint256 _baseStakeAmount,
        bytes32 _initialVMhash
    ) public initializer {
        require(_owner != address(0) && _verifier != address(0), "ZERO_ADDRESS");
        owner = _owner;
        stakeToken = IERC20(_stakeToken);
        verifier = IVerifier(_verifier);

        confirmationPeriod = _confirmationPeriod;
        challengePeriod = _challengePeriod; // TODO: currently unused.
        minimumAssertionPeriod = _minimumAssertionPeriod;
        maxGasPerAssertion = _maxGasPerAssertion;
        baseStakeAmount = _baseStakeAmount;

        assertions = new AssertionMap(address(this));
        assertions.createAssertion(
            0, // assertionID
            RollupLib.stateHash(RollupLib.ExecutionState(0, _initialVMhash)),
            0, // inboxSize (genesis)
            0, // total gasUsed
            0, // parentID
            block.number // deadline (unchallengeable)
        );
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
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
    function stake() external payable override {
        if (isStaked(msg.sender)) {
            stakers[msg.sender].amountStaked += msg.value;
        } else {
            require(msg.value >= baseStakeAmount, "INSUFFICIENT_STAKE");
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
        require(staker.assertionID <= lastConfirmedAssertionID, "STAKED_ON_UNCONFIRMED_ASSERTION");
        require(stakeAmount <= staker.amountStaked - currentRequiredStake(), "INSUFFICIENT_FUNDS");
        staker.amountStaked -= stakeAmount;
        // Note: we don't need to modify assertion state because you can only unstake from a confirmed assertion.
        (bool success,) = msg.sender.call{value: stakeAmount}("");
        require(success, "TRANSFER_FAILED");
    }

    // WARNING: this function is vulnerable to reentrancy attack!
    /// @inheritdoc IRollup
    function removeStake(address stakerAddress) external override {
        requireStaked(stakerAddress);
        // Require that staker is staked on a confirmed assertion.
        Staker storage staker = stakers[stakerAddress];
        require(staker.assertionID <= lastConfirmedAssertionID, "STAKED_ON_UNCONFIRMED_ASSERTION");
        deleteStaker(stakerAddress);
        // Note: we don't need to modify assertion state because you can only unstake from a confirmed assertion.
        (bool success,) = stakerAddress.call{value: staker.amountStaked}("");
        require(success, "TRANSFER_FAILED");
    }

    /// @inheritdoc IRollup
    function advanceStake(uint256 assertionID) external override stakedOnly {
        Staker storage staker = stakers[msg.sender];
        require(assertionID > staker.assertionID && assertionID <= lastCreatedAssertionID, "ASSERTION_OUT_OF_RANGE");
        // TODO: allow arbitrary descendant of current staked assertionID, not just child.
        require(staker.assertionID == assertions.getParentID(assertionID), "PARENT_ASSERTION_UNSTAKED");
        stakeOnAssertion(msg.sender, assertionID);
        emit AdvanceStake(assertionID);
    }

    /// @inheritdoc IRollup
    function withdraw() external override {
        uint256 withdrawableFund = withdrawableFunds[msg.sender];
        withdrawableFunds[msg.sender] = 0;
        (bool success,) = msg.sender.call{value: withdrawableFund}("");
        require(success, "TRANSFER_FAILED");
    }

    function createAssertionWithStateBatch(
        bytes32 vmHash,
        uint256 inboxSize,
        uint256 l2GasUsed,
        bytes32[] calldata _batch,
        uint256 _shouldStartAtElement,
        bytes calldata _signature
    ) external override stakedOnly {
        // permissions will be checked in appendStateBatch method, don't need to check again here
        require(msg.sender == resolve("BVM_Proposer"), "msg.sender is not proposer, can't append batch");

        // create assertion
        createAssertion(vmHash, inboxSize, l2GasUsed);

        // append state batch
        address scc = resolve("StateCommitmentChain");
        (bool success, bytes memory data) = scc.delegatecall(
            abi.encodeWithSignature("appendStateBatch(bytes32[],uint256,bytes)", _batch, _shouldStartAtElement, _signature)
        );
        require(success, "scc append state batch failed, revert all");
    }


    /// @inheritdoc IRollup
    function createAssertion(
        bytes32 vmHash,
        uint256 inboxSize,
        uint256 l2GasUsed
    ) external override stakedOnly {
        // TODO: determine if inboxSize needs to be included.
        RollupLib.ExecutionState memory endState = RollupLib.ExecutionState(l2GasUsed, vmHash);

        uint256 parentID = stakers[msg.sender].assertionID;
        // Require that enough time has passed since the last assertion.
        require(block.number - assertions.getProposalTime(parentID) >= minimumAssertionPeriod, "TIME_DELTA");
        // TODO: require(..., TOO_SMALL);
        uint256 prevL2GasUsed = assertions.getGasUsed(parentID);
        uint256 assertionGasUsed = l2GasUsed - prevL2GasUsed;
        // Require that the L2 gas used by the assertion is less than the limit.
        // TODO: arbitrum uses: timeSinceLastNode.mul(avmGasSpeedLimitPerBlock).mul(4) ?
        require(assertionGasUsed <= maxGasPerAssertion, "TOO_LARGE");
        // Require that the assertion at least includes one transaction
        require(inboxSize > assertions.getInboxSize(parentID), "NO_TXN");
        // Require that the assertion doesn't read past the end of the current inbox.
        require(inboxSize <= sequencerInbox.getInboxSize(), "INBOX_PAST_END");

        // Initialize assertion.
        lastCreatedAssertionID++;
        emit AssertionCreated(lastCreatedAssertionID, msg.sender, vmHash, inboxSize, l2GasUsed);
        assertions.createAssertion(
            lastCreatedAssertionID, RollupLib.stateHash(endState), inboxSize, l2GasUsed, parentID, newAssertionDeadline()
        );

        // Update stake.
        stakeOnAssertion(msg.sender, lastCreatedAssertionID);
    }

    function challengeAssertion(address[2] calldata players, uint256[2] calldata assertionIDs)
        external
        override
        returns (address)
    {
        uint256 defenderAssertionID = assertionIDs[0];
        uint256 challengerAssertionID = assertionIDs[1];
        // Require IDs ordered and in-range.
        require(defenderAssertionID < challengerAssertionID, "WRONG_ORDER");
        require(challengerAssertionID <= lastCreatedAssertionID, "NOT_PROPOSED");
        require(lastConfirmedAssertionID < defenderAssertionID, "ALREADY_RESOLVED");
        // Require that players have attested to sibling assertions.
        uint256 parentID = assertions.getParentID(defenderAssertionID);
        require(parentID == assertions.getParentID(challengerAssertionID), "DIFF_PARENT");

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
        emit AssertionChallenged(defenderAssertionID, challengeAddr);
        challenge.initialize(
            defender,
            challenger,
            verifier,
            address(this),
            assertions.getStateHash(parentID),
            assertions.getStateHash(defenderAssertionID)
        );

        return challengeAddr;
    }

    /// @inheritdoc IRollup
    function confirmFirstUnresolvedAssertion() external override {
        require(lastResolvedAssertionID < lastCreatedAssertionID, "NO_UNRESOLVED");

        // (1) there is at least one staker, and
        require(numStakers > 0, "NO_STAKERS");

        uint256 lastUnresolvedID = lastResolvedAssertionID + 1;
        // (2) challenge period has passed
        require(block.number >= assertions.getDeadline(lastUnresolvedID), "BEFORE_DEADLINE");
        // (3) predecessor has been confirmed
        require(assertions.getParentID(lastUnresolvedID) == lastConfirmedAssertionID, "INVALID_PARENT");

        // Remove old zombies
        // removeOldZombies();

        // (4) all stakers are staked on the block.
        require(
            assertions.getNumStakers(lastUnresolvedID) == countStakedZombies(lastUnresolvedID) + numStakers,
            "NOT_ALL_STAKED"
        );

        // Confirm assertion.
        // assertions.deleteAssertion(lastConfirmedAssertionID);
        lastResolvedAssertionID++;
        lastConfirmedAssertionID = lastResolvedAssertionID;
        emit AssertionConfirmed(lastResolvedAssertionID);
    }

    /// @inheritdoc IRollup
    function rejectFirstUnresolvedAssertion(address stakerAddress) external override {
        require(lastResolvedAssertionID < lastCreatedAssertionID, "NO_UNRESOLVED");

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
            require(block.number >= assertions.getDeadline(firstUnresolvedAssertionID), "BEFORE_DEADLINE");

            // 1b. at least one staker exists (on a sibling)
            // - stakerAddress is indeed a staker
            requireStaked(stakerAddress);
            // - staker's assertion can't be a ancestor of firstUnresolved (because staker's assertion is also unresolved)
            require(stakers[stakerAddress].assertionID >= firstUnresolvedAssertionID, "ASSERTION_ALREADY_RESOLVED");
            // - staker's assertion can't be a descendant of firstUnresolved (because staker has never staked on firstUnresolved)
            require(!assertions.isStaker(firstUnresolvedAssertionID, stakerAddress), "STAKER_STAKED_ON_TARGET");
            // If a staker is staked on an assertion that is neither an ancestor nor a descendant of firstUnresolved, it must be a sibling, QED

            // 1c. no staker is staked on this assertion
            // removeOldZombies();
            require(
                assertions.getNumStakers(firstUnresolvedAssertionID) == countStakedZombies(firstUnresolvedAssertionID),
                "HAS_STAKERS"
            );
        }

        // Reject assertion.
        lastResolvedAssertionID++;
        emit AssertionRejected(lastResolvedAssertionID);
        assertions.deleteAssertion(lastResolvedAssertionID);
    }

    /// @inheritdoc IRollup
    function completeChallenge(address winner, address loser) external override {
        require(msg.sender == getChallenge(winner, loser), "WRONG_SENDER");

        uint256 remainingLoserStake = stakers[loser].amountStaked;
        uint256 winnerStake = stakers[winner].amountStaked;
        if (remainingLoserStake > winnerStake) {
            // If loser has a higher stake than the winner, refund the difference.
            // Loser gets deleted anyways, so maybe unnecessary to set amountStaked.
            stakers[loser].amountStaked = winnerStake;
            withdrawableFunds[loser] += (winnerStake - remainingLoserStake);
            remainingLoserStake = winnerStake;
        }
        // Reward the winner with half the remaining stake
        uint256 amountWon = remainingLoserStake / 2;
        stakers[winner].amountStaked += amountWon; // why +stake instead of +withdrawable?
        stakers[winner].currentChallenge = address(0);
        // Credit the other half to the owner address
        withdrawableFunds[owner] += (remainingLoserStake - amountWon);
        // Turning loser into zombie renders the loser's remaining stake inaccessible.
        uint256 assertionID = stakers[loser].assertionID;
        deleteStaker(loser);
        // Track as zombie so we can account for it during assertion resolution.
        zombies.push(Zombie(loser, assertionID));
    }

    /**
     * @notice Updates staker and assertion metadata.
     * @param stakerAddress Address of existing staker.
     * @param assertionID ID of existing assertion to stake on.
     */
    function stakeOnAssertion(address stakerAddress, uint256 assertionID) private {
        stakers[stakerAddress].assertionID = assertionID;
        assertions.stakeOnAssertion(assertionID, stakerAddress);
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
        require(challenge != address(0), "NO_CHALLENGE");
        require(challenge == staker2.currentChallenge, "DIFF_IN_CHALLENGE");
        return challenge;
    }

    function newAssertionDeadline() private view returns (uint256) {
        // TODO: account for prev assertion, gas
        return block.number + confirmationPeriod;
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
        require(isStaked(stakerAddress), "NOT_STAKED");
    }

    function requireUnchallengedStaker(address stakerAddress) private view {
        requireStaked(stakerAddress);
        require(stakers[stakerAddress].currentChallenge == address(0), "IN_CHALLENGE");
    }
}

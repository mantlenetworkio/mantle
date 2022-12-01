// SPDX-License-Identifier: UNLICENSED

//modified version of Compound's 'GovernorAlpha' contract -- see https://github.com/compound-finance/compound-protocol/blob/master/contracts/Governance/GovernorAlpha.sol

/*
Copyright 2020 Compound Labs, Inc.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following 
	disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products
	derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING,
	BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL 
	THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
	(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
	HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
	ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

/**
 * @title Experimental governance contract.
 * @author Layr Labs, Inc.
 */
pragma solidity ^0.8.9;

import "../interfaces/IVoteWeigher.sol";
import "../interfaces/IRegistry.sol";
import "../interfaces/IQuorumRegistry.sol";
import "./Timelock.sol";
import "../permissions/RepositoryAccess.sol";

//TODO: better solutions for 'quorumVotes' and 'proposalThreshold'
contract Governor is RepositoryAccess {
    struct Proposal {
        /// @notice Unique id for looking up a proposal
        uint256 id;
        /// @notice Creator of the proposal
        address proposer;
        /// @notice The timestamp that the proposal will be available for execution, set once the vote succeeds
        uint256 eta;
        /// @notice the ordered list of target addresses for calls to be made
        address[] targets;
        /// @notice The ordered list of values (i.e. msg.value) to be passed to the calls to be made
        uint256[] values;
        /// @notice The ordered list of function signatures to be called
        string[] signatures;
        /// @notice The ordered list of calldata to be passed to each call
        bytes[] calldatas;
        /// @notice The UTC timestamp at which voting begins: holders must delegate their votes prior to this UTC timestamp
        uint256 startTime;
        /// @notice The UTC timestamp at which voting ends: votes must be cast prior to this UTC timestamp
        uint256 endTime;
        /// @notice Current number of votes from the first quorum in favor of this proposal
        uint256 forVotesFirstQuorum;
        /// @notice Current number of votes from the first quorum in opposition to this proposal
        uint256 againstVotesFirstQuorum;
        /// @notice Current number of votes from the second quorum in favor of this proposal
        uint256 forVotesSecondQuorum;
        /// @notice Current number of votes from the second quorum in opposition to this proposal
        uint256 againstVotesSecondQuorum;
        /// @notice Flag marking whether the proposal has been canceled
        bool canceled;
        /// @notice Flag marking whether the proposal has been executed
        bool executed;
    }

    /// @notice Ballot receipt record for a voter
    struct Receipt {
        /// @notice Whether or not a vote has been cast
        bool hasVoted;
        /// @notice Whether or not the voter supports the proposal
        bool support;
        /// @notice The number of votes the voter had, which were cast
        uint96 firstQuorumVotes;
        uint96 secondQuorumVotes;
    }

    /// @notice Possible states that a proposal may be in
    enum ProposalState {
        Pending,
        Active,
        Canceled,
        Defeated,
        Succeeded,
        Queued,
        Expired,
        Executed
    }
    /// @notice The address of the Protocol Timelock

    Timelock public timelock;
    /// @notice The maximum number of actions that can be included in a proposal
    uint16 public constant proposalMaxOperations = 10;
    /// @notice The delay before voting on a proposal may take place, once proposed. stored as uint256 in number of seconds
    uint256 public constant votingDelay = 2 days;
    /// @notice The duration of voting on a proposal, in seconds
    uint256 public constant votingPeriod = 7 days;
    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
    /// @notice The EIP-712 typehash for the ballot struct used by the contract
    bytes32 public constant BALLOT_TYPEHASH = keccak256("Ballot(uint256 proposalId,bool support)");

    IVoteWeigher public immutable VOTE_WEIGHTER;

    /// @notice The percentage of eth needed in support of a proposal required in order for a quorum
    /// to be reached for the eth and for a vote to succeed, if an eigen quorum is also reached
    uint16 public firstQuorumPercentage;

    /// @notice The percentage of eigen needed in support of a proposal required in order for a quorum
    /// to be reached for the eigen and for a vote to succeed, if an eth quorum is also reached
    uint16 public secondQuorumPercentage;

    /// @notice The percentage of eth required in order for a voter to become a proposer
    uint16 public proposalThresholdFirstQuorumPercentage;

    /// @notice The percentage of eigen required in order for a voter to become a proposer
    uint16 public proposalThresholdSecondQuorumPercentage;

    /// @notice The total number of proposals
    uint256 public proposalCount;

    /// @notice Address of multisig or other trusted party. Has zero extra votes, but can make proposals without meeting proposal
    ///         thresholds, and that pass by default
    ///         In other words, proposals created by the 'multsig' address require at least a quorum to reject them in order to
    ///         *not* pass, but otherwise operate like normal proposals, where the side with most votes wins
    address public multisig;

    /// @notice Receipts of ballots for the entire set of voters
    mapping(uint256 => mapping(address => Receipt)) public receipts;

    /// @notice The official record of all proposals ever proposed
    mapping(uint256 => Proposal) public proposals;

    /// @notice The latest proposal for each proposer
    mapping(address => uint256) public latestProposalIds;

    /// @notice An event emitted when a new proposal is created
    event ProposalCreated(
        uint256 id,
        address proposer,
        address[] targets,
        uint256[] values,
        string[] signatures,
        bytes[] calldatas,
        uint256 startTime,
        uint256 endTime,
        string description
    );

    /// @notice An event emitted when a vote has been cast on a proposal
    event VoteCast(
        address voter, uint256 proposalId, bool support, uint256 firstQuorumVotes, uint256 secondQuorumVotes
    );

    /// @notice An event emitted when a proposal has been canceled
    event ProposalCanceled(uint256 id);

    /// @notice An event emitted when a proposal has been queued in the Timelock
    event ProposalQueued(uint256 id, uint256 eta);

    /// @notice An event emitted when a proposal has been executed in the Timelock
    event ProposalExecuted(uint256 id);

    /// @notice Emitted when the 'multisig' address has been changed
    event MultisigTransferred(address indexed previousAddress, address indexed newAddress);

    /// @notice Emitted when the 'timelock' address has been changed
    event TimelockTransferred(address indexed previousAddress, address indexed newAddress);

    modifier onlyTimelock() {
        require(msg.sender == address(timelock), "onlyTimelock");
        _;
    }

    constructor(
        IRepository _repository,
        IVoteWeigher _VOTE_WEIGHTER,
        Timelock _timelock,
        address _multisig,
        uint16 _firstQuorumPercentage,
        uint16 _secondQuorumPercentage,
        uint16 _proposalThresholdFirstQuorumPercentage,
        uint16 _proposalThresholdSecondQuorumPercentage
    ) RepositoryAccess(_repository) {
        VOTE_WEIGHTER = _VOTE_WEIGHTER;
        _setTimelock(_timelock);
        _setMultisig(_multisig);
        _setQuorumsAndThresholds(
            _firstQuorumPercentage,
            _secondQuorumPercentage,
            _proposalThresholdFirstQuorumPercentage,
            _proposalThresholdSecondQuorumPercentage
        );
    }

    function propose(
        address[] memory targets,
        uint256[] memory values,
        string[] memory signatures,
        bytes[] memory calldatas,
        string memory description
    ) external returns (uint256) {
        (uint96 firstQuorumStake, uint96 secondQuorumStake) = _getVoterStakes(msg.sender);
        {
            // check percentage
            IQuorumRegistry registry = IQuorumRegistry(address(repository.registry()));
            (uint256 firstQuorumTotalStake, uint256 secondQuorumTotalStake) = registry.totalStake();
            require(
                (uint256(firstQuorumStake) * 100) / firstQuorumTotalStake
                    >= proposalThresholdFirstQuorumPercentage
                    || (uint256(secondQuorumStake) * 100) / secondQuorumTotalStake
                        >= proposalThresholdSecondQuorumPercentage || msg.sender == multisig,
                "RepositoryGovernance::propose: proposer votes below proposal threshold"
            );
        }
        require(
            targets.length == values.length && targets.length == signatures.length && targets.length == calldatas.length,
            "RepositoryGovernance::propose: proposal function information arity mismatch"
        );
        require(targets.length != 0, "RepositoryGovernance::propose: must provide actions");
        require(targets.length <= proposalMaxOperations, "RepositoryGovernance::propose: too many actions");

        uint256 latestProposalId = latestProposalIds[msg.sender];
        if (latestProposalId != 0) {
            ProposalState proposersLatestProposalState = state(latestProposalId);
            require(
                proposersLatestProposalState != ProposalState.Active,
                "RepositoryGovernance::propose: one live proposal per proposer, found an already active proposal"
            );
            require(
                proposersLatestProposalState != ProposalState.Pending,
                "RepositoryGovernance::propose: one live proposal per proposer, found an already pending proposal"
            );
        }

        ++proposalCount;
        Proposal memory newProposal = Proposal({
            id: proposalCount,
            proposer: msg.sender,
            eta: 0,
            targets: targets,
            values: values,
            signatures: signatures,
            calldatas: calldatas,
            startTime: block.timestamp + votingDelay,
            endTime: block.timestamp + votingDelay + votingPeriod,
            forVotesFirstQuorum: 0,
            againstVotesFirstQuorum: 0,
            forVotesSecondQuorum: 0,
            againstVotesSecondQuorum: 0,
            canceled: false,
            executed: false
        });

        proposals[newProposal.id] = newProposal;
        latestProposalIds[newProposal.proposer] = newProposal.id;

        emit ProposalCreated(
            newProposal.id,
            msg.sender,
            targets,
            values,
            signatures,
            calldatas,
            newProposal.startTime,
            newProposal.endTime,
            description
            );
        return newProposal.id;
    }

    function queue(uint256 proposalId) external {
        require(
            state(proposalId) == ProposalState.Succeeded,
            "RepositoryGovernance::queue: proposal can only be queued if it is succeeded"
        );
        Proposal storage proposal = proposals[proposalId];
        uint256 eta = block.timestamp + timelock.delay();
        for (uint256 i = 0; i < proposal.targets.length; ++i) {
            _queueOrRevert(proposal.targets[i], proposal.values[i], proposal.signatures[i], proposal.calldatas[i], eta);
        }
        proposal.eta = eta;
        emit ProposalQueued(proposalId, eta);
    }

    function _queueOrRevert(address target, uint256 value, string memory signature, bytes memory data, uint256 eta)
        internal
    {
        require(
            !timelock.queuedTransactions(keccak256(abi.encode(target, value, signature, data, eta))),
            "RepositoryGovernance::_queueOrRevert: proposal action already queued at eta"
        );
        // slither-disable-next-line unused-return
        timelock.queueTransaction(target, value, signature, data, eta);
    }

    function execute(uint256 proposalId) external payable {
        require(
            state(proposalId) == ProposalState.Queued,
            "RepositoryGovernance::execute: proposal can only be executed if it is queued"
        );
        Proposal storage proposal = proposals[proposalId];
        proposal.executed = true;
        for (uint256 i = 0; i < proposal.targets.length; ++i) {
            // slither-disable-next-line unused-return
            timelock.executeTransaction{value: proposal.values[i]}(
                proposal.targets[i], proposal.values[i], proposal.signatures[i], proposal.calldatas[i], proposal.eta
            );
        }
        emit ProposalExecuted(proposalId);
    }

    function cancel(uint256 proposalId) external {
        ProposalState stateOfProposal = state(proposalId);
        require(
            stateOfProposal != ProposalState.Executed, "RepositoryGovernance::cancel: cannot cancel executed proposal"
        );

        Proposal storage proposal = proposals[proposalId];
        require(proposal.proposer != multisig, "multisig does not have to meet threshold requirements");
        (uint96 firstQuorumStake, uint96 secondQuorumStake) = _getVoterStakes(proposal.proposer);
        {
            // check percentage
            IQuorumRegistry registry = IQuorumRegistry(address(_registry()));
            (uint256 firstQuorumTotalStake, uint256 secondQuorumTotalStake) = registry.totalStake();
            require(
                (uint256(firstQuorumStake) * 100) / firstQuorumTotalStake
                    >= proposalThresholdFirstQuorumPercentage
                    || (uint256(secondQuorumStake) * 100) / secondQuorumTotalStake
                        >= proposalThresholdSecondQuorumPercentage || msg.sender == multisig,
                "RepositoryGovernance::propose: proposer votes below proposal threshold"
            );
        }
        proposal.canceled = true;
        for (uint256 i = 0; i < proposal.targets.length; ++i) {
            timelock.cancelTransaction(
                proposal.targets[i], proposal.values[i], proposal.signatures[i], proposal.calldatas[i], proposal.eta
            );
        }

        emit ProposalCanceled(proposalId);
    }

    function getActions(uint256 proposalId)
        external
        view
        returns (
            address[] memory targets,
            uint256[] memory values,
            string[] memory signatures,
            bytes[] memory calldatas
        )
    {
        Proposal storage p = proposals[proposalId];
        return (p.targets, p.values, p.signatures, p.calldatas);
    }

    function getReceipt(uint256 proposalId, address voter) external view returns (Receipt memory) {
        return receipts[proposalId][voter];
    }

    function state(uint256 proposalId) public view returns (ProposalState) {
        require(proposalCount >= proposalId && proposalId > 0, "RepositoryGovernance::state: invalid proposal id");
        Proposal storage proposal = proposals[proposalId];
        IQuorumRegistry registry = IQuorumRegistry(address(repository.registry()));
        (uint256 firstQuorumTotalStake, uint256 secondQuorumTotalStake) = registry.totalStake();
        if (proposal.canceled) {
            return ProposalState.Canceled;
        } else if (block.timestamp <= proposal.startTime) {
            return ProposalState.Pending;
        } else if (block.timestamp <= proposal.endTime) {
            return ProposalState.Active;
        } else if (
            proposal.forVotesFirstQuorum <= proposal.againstVotesFirstQuorum
                || proposal.forVotesSecondQuorum <= proposal.againstVotesSecondQuorum
                || (
                    (
                        (proposal.forVotesFirstQuorum * 100) / firstQuorumTotalStake
                            < firstQuorumPercentage
                    ) && (proposal.proposer != multisig)
                )
                || (
                    (
                        (proposal.forVotesSecondQuorum * 100) / secondQuorumTotalStake
                            < secondQuorumPercentage
                    ) && (proposal.proposer != multisig)
                )
        ) {
            return ProposalState.Defeated;
        } else if (proposal.eta == 0) {
            return ProposalState.Succeeded;
        } else if (proposal.executed) {
            return ProposalState.Executed;
        } else if (block.timestamp >= proposal.eta + timelock.GRACE_PERIOD()) {
            return ProposalState.Expired;
        } else {
            return ProposalState.Queued;
        }
    }

    function castVote(uint256 proposalId, bool support) external {
        return _castVote(msg.sender, proposalId, support);
    }

    function castVoteBySig(uint256 proposalId, bool support, uint8 v, bytes32 r, bytes32 s) external {
        bytes32 domainSeparator = keccak256(abi.encode(DOMAIN_TYPEHASH, getChainId(), address(this)));
        bytes32 structHash = keccak256(abi.encode(BALLOT_TYPEHASH, proposalId, support));
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));
        address signatory = ecrecover(digest, v, r, s);
        require(signatory != address(0), "RepositoryGovernance::castVoteBySig: invalid signature");
        return _castVote(signatory, proposalId, support);
    }

    function _castVote(address voter, uint256 proposalId, bool support) internal {
        require(state(proposalId) == ProposalState.Active, "RepositoryGovernance::_castVote: voting is closed");
        Proposal storage proposal = proposals[proposalId];
        Receipt storage receipt = receipts[proposalId][voter];
        require(!receipt.hasVoted, "RepositoryGovernance::_castVote: voter already voted");
        (uint96 firstQuorumStake, uint96 secondQuorumStake) = _getVoterStakes(voter);

        if (support) {
            proposal.forVotesFirstQuorum = proposal.forVotesFirstQuorum + firstQuorumStake;
            proposal.forVotesSecondQuorum = proposal.forVotesSecondQuorum + secondQuorumStake;
        } else {
            proposal.againstVotesFirstQuorum = proposal.againstVotesFirstQuorum + firstQuorumStake;
            proposal.againstVotesSecondQuorum = proposal.againstVotesSecondQuorum + secondQuorumStake;
        }
        receipt.hasVoted = true;
        receipt.support = support;
        receipt.firstQuorumVotes = firstQuorumStake;
        receipt.secondQuorumVotes = secondQuorumStake;

        emit VoteCast(voter, proposalId, support, firstQuorumStake, secondQuorumStake);
    }

    function setQuorumsAndThresholds(
        uint16 _firstQuorumPercentage,
        uint16 _secondQuorumPercentage,
        uint16 _proposalThresholdFirstQuorumPercentage,
        uint16 _proposalThresholdSecondQuorumPercentage
    ) external onlyTimelock {
        _setQuorumsAndThresholds(
            _firstQuorumPercentage,
            _secondQuorumPercentage,
            _proposalThresholdFirstQuorumPercentage,
            _proposalThresholdSecondQuorumPercentage
        );
    }

    function _setQuorumsAndThresholds(
        uint16 _firstQuorumPercentage,
        uint16 _secondQuorumPercentage,
        uint16 _proposalThresholdFirstQuorumPercentage,
        uint16 _proposalThresholdSecondQuorumPercentage
    ) internal {
        require(_firstQuorumPercentage > 0 && _firstQuorumPercentage < 100, "bad _firstQuorumPercentage");
        require(_secondQuorumPercentage > 0 && _secondQuorumPercentage < 100, "bad _secondQuorumPercentage");
        require(
            _proposalThresholdFirstQuorumPercentage > 0 && _proposalThresholdFirstQuorumPercentage < 100,
            "bad _proposalThresholdFirstQuorumPercentage"
        );
        require(
            _proposalThresholdSecondQuorumPercentage > 0 && _proposalThresholdSecondQuorumPercentage < 100,
            "bad _proposalThresholdSecondQuorumPercentage"
        );

        firstQuorumPercentage = _firstQuorumPercentage;
        secondQuorumPercentage = _secondQuorumPercentage;
        proposalThresholdFirstQuorumPercentage = _proposalThresholdFirstQuorumPercentage;
        proposalThresholdSecondQuorumPercentage = _proposalThresholdSecondQuorumPercentage;
    }

    function setMultisig(address _multisig) external onlyTimelock {
        _setMultisig(_multisig);
    }

    function _setMultisig(address _multisig) internal {
        emit MultisigTransferred(multisig, _multisig);
        multisig = _multisig;
    }

    function getChainId() internal view returns (uint256) {
        return block.chainid;
    }

    function setTimelock(Timelock _timelock) external onlyTimelock {
        _setTimelock(_timelock);
    }

    function _setTimelock(Timelock _timelock) internal {
        emit TimelockTransferred(address(timelock), address(_timelock));
        timelock = _timelock;
    }

    // TODO: reintroduce a way to update stakes before simply fetching them?
    function _getVoterStakes(address user) internal view returns (uint96, uint96) {
        (uint96 firstQuorumStake, uint96 secondQuorumStake) = IQuorumRegistry(address(_registry())).operatorStakes(user);
        return (firstQuorumStake, secondQuorumStake);
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../interfaces/IEigenLayrDelegation.sol";
import "../interfaces/IInvestmentStrategy.sol";
import "../interfaces/IInvestmentManager.sol";
import "../permissions/RepositoryAccess.sol";

/**
 * @title Storage variables for the `VoteWeigherBase` contract.
 * @author Layr Labs, Inc.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract VoteWeigherBaseStorage is RepositoryAccess, IVoteWeigher {
    /**
     * @notice In weighing a particular investment strategy, the amount of underlying asset for that strategy is
     * multiplied by its multiplier, then divided by WEIGHTING_DIVISOR
     */
    struct StrategyAndWeightingMultiplier {
        IInvestmentStrategy strategy;
        uint96 multiplier;
    }

    /// @notice Constant used as a divisor in calculating weights.
    uint256 internal constant WEIGHTING_DIVISOR = 1e18;
    /// @notice Maximum length of dynamic arrays in the `strategiesConsideredAndMultipliers` mapping.
    uint8 internal constant MAX_WEIGHING_FUNCTION_LENGTH = 32;
    /// @notice Constant used as a divisor in dealing with BIPS amounts.
    uint256 internal constant MAX_BIPS = 10000;

    /// @notice The address of the Delegation contract for EigenLayr.
    IEigenLayrDelegation public immutable delegation;
    
    /// @notice The address of the InvestmentManager contract for EigenLayr.
    IInvestmentManager public immutable investmentManager;

    /// @notice Number of quorums that are being used by the middleware.
    uint256 public immutable NUMBER_OF_QUORUMS;

    /**
     * @notice mapping from quorum number to the list of strategies considered and their
     * corresponding multipliers for that specific quorum
     */
    mapping(uint256 => StrategyAndWeightingMultiplier[]) public strategiesConsideredAndMultipliers;

    /**
     * @notice This defines the earnings split between different quorums. Mapping is quorumNumber => BIPS which the quorum earns, out of the total earnings.
     * @dev The sum of all entries, i.e. sum(quorumBips[0] through quorumBips[NUMBER_OF_QUORUMS - 1]) should *always* be 10,000!
     */
    mapping(uint256 => uint256) public quorumBips;

    constructor(
        IRepository _repository,
        IEigenLayrDelegation _delegation,
        IInvestmentManager _investmentManager,
        uint8 _NUMBER_OF_QUORUMS,
        uint256[] memory _quorumBips
    ) RepositoryAccess(_repository) {
        // sanity check that the VoteWeigher is being initialized with at least 1 quorum
        require(_NUMBER_OF_QUORUMS != 0, "VoteWeigherBaseStorage.constructor: _NUMBER_OF_QUORUMS == 0");
        // verify that the provided `_quorumBips` is of the correct length
        require(
            _quorumBips.length == _NUMBER_OF_QUORUMS,
            "VoteWeigherBaseStorage.constructor: _quorumBips.length != _NUMBER_OF_QUORUMS"
        );
        uint256 totalQuorumBips;
        for (uint256 i; i < _NUMBER_OF_QUORUMS; ++i) {
            totalQuorumBips += _quorumBips[i];
            quorumBips[i] = _quorumBips[i];
        }
        // verify that the provided `_quorumBips` do indeed sum to 10,000!
        require(totalQuorumBips == MAX_BIPS, "VoteWeigherBaseStorage.constructor: totalQuorumBips != MAX_BIPS");
        delegation = _delegation;
        investmentManager = _investmentManager;
        NUMBER_OF_QUORUMS = _NUMBER_OF_QUORUMS;
    }
}

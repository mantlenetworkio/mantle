// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../interfaces/IInvestmentManager.sol";
import "../interfaces/IDelegationTerms.sol";
import "../interfaces/IEigenLayrDelegation.sol";

/**
 * @title Storage variables for the `EigenLayrDelegation` contract.
 * @author Layr Labs, Inc.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract EigenLayrDelegationStorage is IEigenLayrDelegation {
    /// @notice Gas budget provided in calls to DelegationTerms contracts
    uint256 internal constant LOW_LEVEL_GAS_BUDGET = 1e5;

    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /// @notice The EIP-712 typehash for the delegation struct used by the contract
    bytes32 public constant DELEGATION_TYPEHASH =
        keccak256("Delegation(address delegator,address operator,uint256 nonce,uint256 expiry)");

    /// @notice EIP-712 Domain separator
    bytes32 public DOMAIN_SEPARATOR;

    /// @notice The InvestmentManager contract for EigenLayr
    IInvestmentManager public immutable investmentManager;

    // operator => investment strategy => total number of shares delegated to them
    mapping(address => mapping(IInvestmentStrategy => uint256)) public operatorShares;

    // operator => delegation terms contract
    mapping(address => IDelegationTerms) public delegationTerms;

    // staker => operator
    mapping(address => address) public delegatedTo;

    // staker => whether they are delegated or not
    mapping(address => IEigenLayrDelegation.DelegationStatus) public delegationStatus;

    // delegator => number of signed delegation nonce (used in delegateToBySignature)
    mapping(address => uint256) public nonces;

    constructor(IInvestmentManager _investmentManager) {
        investmentManager = _investmentManager;
    }
}

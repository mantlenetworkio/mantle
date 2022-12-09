// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../interfaces/IRepository.sol";
import "../interfaces/IRepositoryAccess.sol";

/**
 * @title Defines access controls based around a single `Repository` contract.
 * @author Layr Labs, Inc.
 */
abstract contract RepositoryAccess is IRepositoryAccess {
    /// @notice The unique, immutable Repository contract associated with this contract.
    IRepository public immutable repository;

    /// @notice Irrevocably sets the (immutable) `repository` for this contract.
    constructor(IRepository _repository) {
        repository = _repository;
    }

    // MODIFIERS -- access controls based on stored addresses
    /// @notice when applied to a function, ensures that it is only callable by `repository` itself
    modifier onlyRepository() {
        require(msg.sender == address(repository), "onlyRepository");
        _;
    }

    /// @notice when applied to a function, ensures that it is only callable by the current `owner` of `repository`
    modifier onlyRepositoryGovernance() {
        require(msg.sender == address(_repositoryGovernance()), "onlyRepositoryGovernance");
        _;
    }

    /// @notice when applied to a function, ensures that it is only callable by the current `serviceManager` of `repository`
    modifier onlyServiceManager() {
        require(msg.sender == address(_serviceManager()), "onlyServiceManager");
        _;
    }

    /// @notice when applied to a function, ensures that it is only callable by the current `registry` of `repository`
    modifier onlyRegistry() {
        require(msg.sender == address(_registry()), "onlyRegistry");
        _;
    }

    // INTERNAL FUNCTIONS -- fetch info from repository
    function _repositoryGovernance() internal view returns (address) {
        return repository.owner();
    }

    function _serviceManager() internal view returns (IServiceManager) {
        return repository.serviceManager();
    }

    function _registry() internal view returns (IRegistry) {
        return repository.registry();
    }
}

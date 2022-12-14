// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../interfaces/IInvestmentManager.sol";
import "../interfaces/IEigenLayrDelegation.sol";
import "../interfaces/IRegistry.sol";
import "../interfaces/IRepository.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

/**
 * @title A central hub contract for a single middleware on EigenLayr.
 * @author Layr Labs, Inc.
 * @notice The Repository contract is intended to be the central source of "ground truth" for a single middleware on EigenLayr.
 * @notice The owner of the Repository for a middleware holds *tremendous power* – this role should only be given to a multisig or governance contract.
 * @dev Other contracts can refer to the Repository in order to determine the "official" contracts for the middleware, making it the central point
 * for upgrades-by-changing-contract-addresses.
 */
contract Repository is Ownable, Initializable, IRepository {
    /// @notice Address of the Delegation contract of EigenLayr.
    IEigenLayrDelegation public immutable delegation;
    /// @notice Address of the InvestmentManager contract of EigenLayr.
    IInvestmentManager public immutable investmentManager;
    /// @notice The VoteWeigher contract for this middleware, which determines quorum weighing functions.
    IVoteWeigher public voteWeigher;
    /// @notice The Registry contract for this middleware, where operators register and deregister.
    IRegistry public registry;
    /// @notice The ServiceManager contract for this middleware, where tasks are created / initiated.
    IServiceManager public serviceManager;

    event VoteWeigherSet(IVoteWeigher indexed previousAddress, IVoteWeigher indexed newAddress);
    event ServiceManagerSet(IServiceManager indexed previousAddress, IServiceManager indexed newAddress);
    event RegistrySet(IRegistry indexed previousAddress, IRegistry indexed newAddress);

    /// @notice Sets the (immutable) `delegation` and `investmentManager` addresses -- these are global to EigenLayr and should not change
    constructor(IEigenLayrDelegation _delegation, IInvestmentManager _investmentManager) {
        delegation = _delegation;
        investmentManager = _investmentManager;
    }

    /// @notice returns the owner of this repository contract
    function owner() public view override (Ownable, IRepository) returns (address) {
        return Ownable.owner();
    }

    /// @notice initializer (callable only once). Used for setting the associated contracts for the middleware and the intial owner of this contract.
    function initialize(
        IVoteWeigher _voteWeigher,
        IServiceManager _serviceManager,
        IRegistry _registry,
        address initialOwner
    ) external initializer {
        _setVoteWeigher(_voteWeigher);
        _setServiceManager(_serviceManager);
        _setRegistry(_registry);
        _transferOwnership(initialOwner);
    }

    /// @notice sets the vote weigher for the middleware
    function setVoteWeigher(IVoteWeigher _voteWeigher) external onlyOwner {
        _setVoteWeigher(_voteWeigher);
    }

    /// @notice sets the ServiceManager for the middleware
    function setServiceManager(IServiceManager _serviceManager) external onlyOwner {
        _setServiceManager(_serviceManager);
    }

    /// @notice sets the Registry for the middleware
    function setRegistry(IRegistry _registry) external onlyOwner {
        _setRegistry(_registry);
    }

    function _setVoteWeigher(IVoteWeigher _voteWeigher) internal {
        require(address(_voteWeigher) != address(0), "Repository._setVoteWeigher: zero address bad!");
        emit VoteWeigherSet(voteWeigher, _voteWeigher);
        voteWeigher = _voteWeigher;
    }

    function _setServiceManager(IServiceManager _serviceManager) internal {
        require(address(_serviceManager) != address(0), "Repository._setServiceManager: zero address bad!");
        emit ServiceManagerSet(serviceManager, _serviceManager);
        serviceManager = _serviceManager;
    }

    function _setRegistry(IRegistry _registry) internal {
        require(address(_registry) != address(0), "Repository._setRegistry: zero address bad!");
        emit RegistrySet(registry, _registry);
        registry = _registry;
    }
}

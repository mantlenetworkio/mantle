// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../../interfaces/IDataLayrServiceManager.sol";
import "../../interfaces/IEigenLayrDelegation.sol";
import "../../interfaces/IServiceManager.sol";
import "../../interfaces/IInvestmentManager.sol";
import "./DataLayrPaymentManager.sol";
import "./DataLayrLowDegreeChallenge.sol";
import "./DataLayrBombVerifier.sol";
import "../EphemeralKeyRegistry.sol";
import "../../permissions/RepositoryAccess.sol";

/**
 * @title Storage variables for the `DataLayrServiceManager` contract.
 * @author Layr Labs, Inc.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract DataLayrServiceManagerStorage is IDataLayrServiceManager, RepositoryAccess {
    /**
     *
     * CONSTANTS
     *
     */
    //TODO: mechanism to change any of these values?
    /// @notice Unit of measure (in time) for the duration of DataStores
    uint256 public constant DURATION_SCALE = 1 hours;
    uint256 public constant NUM_DS_PER_BLOCK_PER_DURATION = 20;
    // NOTE: these values are measured in *DURATION_SCALE*
    uint8 public constant MIN_DATASTORE_DURATION = 1;
    /// @notice The longest allowed duation of a DataStore, measured in `DURATION_SCALE`
    uint8 public constant MAX_DATASTORE_DURATION = 7;

    uint32 internal constant MIN_STORE_SIZE = 32;
    uint32 internal constant MAX_STORE_SIZE = 4e9;
    uint256 internal constant BLOCK_STALE_MEASURE = 100;
    uint256 public constant BIP_MULTIPLIER = 10000;

    /// @notice Collateral token used for placing collateral on challenges & payment commits
    IERC20 public immutable collateralToken;

    /**
     * @notice The EigenLayr delegation contract for this DataLayr which is primarily used by
     * delegators to delegate their stake to operators who would serve as DataLayr
     * nodes and so on.
     */
    /**
     * @dev For more details, see EigenLayrDelegation.sol.
     */
    IEigenLayrDelegation public immutable eigenLayrDelegation;

    IInvestmentManager public immutable investmentManager;

    /**
     * @notice service fee that will be paid out by the disperser to the DataLayr nodes
     * for storing per byte for per unit time.
     */
    uint256 public feePerBytePerTime;

    // TODO: set these values correctly
    /// @notice number of leaves in the root tree
    uint48 public constant numPowersOfTau = 0;
    /// @notice number of layers in the root tree
    uint48 public constant log2NumPowersOfTau = 0;

    //TODO: store these upon construction
    // Commitment(0), Commitment(x - w), Commitment((x-w)(x-w^2)), ...
    /**
     * @notice For a given l, zeroPolynomialCommitmentMerkleRoots[l] represents the root of merkle tree

                                    zeroPolynomialCommitmentMerkleRoots[l]
                                                        :    
                                                        :    
                         ____________ ....                             .... ____________              
                        |                                                               |
                        |                                                               |    
              _____h(h_1||h_2)______                                        ____h(h_{k-1}||h_{k}__________  
             |                      |                                      |                              |   
             |                      |                                      |                              |
            h_1                    h_2                                 h_{k-1}                         h_{k} 
             |                      |                                      |                              |  
             |                      |                                      |                              |  
     hash(x^l - w^l)       hash(x^l - (w^2)^l)                   hash(x^l - (w^{k-1})^l)        hash(x^l - (w^k)^l) 
     
     This tree is computed off-chain and only the Merkle roots are stored on-chain.
     */
    // CRITIC: does that mean there are only 32 possible 32 possible merkle trees?
    bytes32[32] public zeroPolynomialCommitmentMerkleRoots;

    /**
     * @notice mapping between the dataStoreId for a particular assertion of data into
     * DataLayr and a compressed information on the signatures of the DataLayr
     * nodes who signed up to be the part of the quorum.
     */
    mapping(uint32 => bytes32) public dataStoreIdToSignatureHash;

    //mapping from duration to timestamp to all of the ids of datastores that were initialized during that timestamp.
    //the third nested mapping just keeps track of a fixed number of datastores of a certain duration that can be
    //in that block
    mapping(uint8 => mapping(uint256 => bytes32[NUM_DS_PER_BLOCK_PER_DURATION])) public
        dataStoreHashesForDurationAtTimestamp;

    DataLayrLowDegreeChallenge public dataLayrLowDegreeChallenge;

    DataLayrBombVerifier public dataLayrBombVerifier;

    EphemeralKeyRegistry public ephemeralKeyRegistry;

    /**
     * @notice contract used for handling payment challenges
     */
    IDataLayrPaymentManager public dataLayrPaymentManager;

    constructor(
        IInvestmentManager _investmentManager,
        IEigenLayrDelegation _eigenLayrDelegation,
        IERC20 _collateralToken
    ) {
        investmentManager = _investmentManager;
        eigenLayrDelegation = _eigenLayrDelegation;
        collateralToken = _collateralToken;
    }
}

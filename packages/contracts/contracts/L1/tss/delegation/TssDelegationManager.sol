// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../delegation/DelegationManager.sol";


/**
 * @title The primary entry- and exit-point for funds into and out.
 * @notice This contract is for managing investments in different strategies. The main
 * functionalities are:
 * - adding and removing investment strategies that any delegator can invest into
 * - enabling deposit of assets into specified investment delegation(s)
 * - enabling removal of assets from specified investment delegation(s)
 * - recording deposit of ETH into settlement layer
 * - recording deposit for securing
 * - slashing of assets for permissioned strategies
 */
contract TssDelegationManager is DelegationManager {


    address public stakingSlash;




    /**
     * @param _delegation The delegation contract.
     * @param _delegationSlasher The primary slashing contract.
     */
    constructor(IDelegation _delegation, IDelegationSlasher _delegationSlasher)
    DelegationManager(_delegation, _delegationSlasher)
    {
        _disableInitializers();
    }


    modifier onlyStakingSlash() {
        require(msg.sender == stakingSlash, "contract call is not staking slashing");
        _;
    }

    function setStakingSlash(address _address) public onlyOwner {
        stakingSlash = _address;
    }


    /**
     * @notice Slashes the shares of a 'frozen' operator (or a staker delegated to one)
     * @param slashedAddress is the frozen address that is having its shares slashed
     * @param delegationIndexes is a list of the indices in `investorStrats[msg.sender]` that correspond to the strategies
     * for which `msg.sender` is withdrawing 100% of their shares
     * @param recipient The slashed funds are withdrawn as tokens to this address.
     * @dev delegationShares are removed from `investorStrats` by swapping the last entry with the entry to be removed, then
     * popping off the last entry in `investorStrats`. The simplest way to calculate the correct `delegationIndexes` to input
     * is to order the strategies *for which `msg.sender` is withdrawing 100% of their shares* from highest index in
     * `investorStrats` to lowest index
     */
    function slashOperatorShares(
        address slashedAddress,
        address recipient,
        IDelegationShare[] calldata delegationShares,
        uint256[] calldata delegationIndexes,
        uint256[] calldata shareAmounts
    )
        external
        whenNotPaused
        nonReentrant
        onlyStakingSlash
    {
        uint256 delegationIndex;
        uint256 strategiesLength = delegationShares.length;
        for (uint256 i = 0; i < strategiesLength;) {
            // the internal function will return 'true' in the event the delegation contract was
            // removed from the slashedAddress's array of strategies -- i.e. investorStrats[slashedAddress]
            if (_removeShares(slashedAddress, delegationIndexes[delegationIndex], delegationShares[i], shareAmounts[i])) {
                _addShares(recipient, delegationShares[i], shareAmounts[i]);
                delegation.increaseDelegatedShares(recipient, delegationShares[i], shareAmounts[i]);
                unchecked {
                    ++delegationIndex;
                }
            }

            // increment the loop
            unchecked {
                ++i;
            }
        }

        // modify delegated shares accordingly, if applicable
        delegation.decreaseDelegatedShares(slashedAddress, delegationShares, shareAmounts);


    }
    



}


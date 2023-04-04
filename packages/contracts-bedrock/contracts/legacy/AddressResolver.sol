// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/* Library Imports */
import { AddressManager } from "./AddressManager.sol";

/**
 * @title AddressResolver
 */
abstract contract AddressResolver {
    /*************
     * Variables *
     *************/

    AddressManager public addressManager;

    /***************
     * Constructor *
     ***************/

    /**
     * @param _addressManager Address of the Lib_AddressManager.
     */
    constructor(address _addressManager) {
        addressManager = AddressManager(_addressManager);
    }

    /********************
     * Public Functions *
     ********************/

    /**
     * Resolves the address associated with a given name.
     * @param _name Name to resolve an address for.
     * @return Address associated with the given name.
     */
    function resolve(string memory _name) public view returns (address) {
        return addressManager.getAddress(_name);
    }
}

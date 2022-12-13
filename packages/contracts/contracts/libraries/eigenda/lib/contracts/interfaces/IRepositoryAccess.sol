// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./IRepository.sol";

/**
 * @title This is the interface for the `RepositoryAccess` contract.
 * @author Layr Labs, Inc.
 */
interface IRepositoryAccess {
    function repository() external view returns (IRepository);
}

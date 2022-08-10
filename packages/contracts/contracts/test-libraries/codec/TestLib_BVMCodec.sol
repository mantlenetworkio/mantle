// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/* Library Imports */
import { Lib_BVMCodec } from "../../libraries/codec/Lib_BVMCodec.sol";

/**
 * @title TestLib_BVMCodec
 */
contract TestLib_BVMCodec {
    function encodeTransaction(Lib_BVMCodec.Transaction memory _transaction)
        public
        pure
        returns (bytes memory _encoded)
    {
        return Lib_BVMCodec.encodeTransaction(_transaction);
    }

    function hashTransaction(Lib_BVMCodec.Transaction memory _transaction)
        public
        pure
        returns (bytes32 _hash)
    {
        return Lib_BVMCodec.hashTransaction(_transaction);
    }
}

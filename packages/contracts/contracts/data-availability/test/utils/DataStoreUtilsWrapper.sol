// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../../contracts/libraries/DataStoreUtils.sol";

// wrapper around the DataStoreUtils library, so that we can call the library's functions that take inputs with 'calldata' location specified
contract DataStoreUtilsWrapper {
    function computeDataStoreHashExternal(IDataLayrServiceManager.DataStoreMetadata memory metadata)
        internal
        pure
        returns (bytes32)
    {
        return DataStoreUtils.computeDataStoreHash(metadata);
    }

    function packDataStoreMetadataExternal(IDataLayrServiceManager.DataStoreMetadata memory metadata)
        external
        pure
        returns (bytes memory)
    {
        return DataStoreUtils.packDataStoreMetadata(metadata);
    }

    function unpackDataStoreMetadataExternal(bytes calldata packedMetadata)
        external
        pure
        returns (IDataLayrServiceManager.DataStoreMetadata memory metadata)
    {
        return DataStoreUtils.unpackDataStoreMetadata(packedMetadata);
    }

    function packDataStoreSearchDataExternal(IDataLayrServiceManager.DataStoreSearchData memory searchData)
        external
        pure
        returns (bytes memory)
    {
        return DataStoreUtils.packDataStoreSearchData(searchData);
    }

    function unpackDataStoreSearchDataExternal(bytes calldata packedSearchData)
        external
        pure
        returns (IDataLayrServiceManager.DataStoreSearchData memory searchData)
    {
        return DataStoreUtils.unpackDataStoreSearchData(packedSearchData);
    }
}

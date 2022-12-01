// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "../interfaces/IDataLayrServiceManager.sol";

/**
 * @title Library of functions shared across DataLayr.
 * @author Layr Labs, Inc.
 */
library DataStoreUtils {
    uint16 public constant BIP_MULTIPLIER = 10000;

    uint256 public constant BYTES_PER_COEFFICIENT = 31;
    uint256 public constant BIT_SHIFT_degree = 224;
    uint256 public constant BIT_SHIFT_numSys = 224;
    uint256 public constant HEADER_OFFSET_degree = 64;
    uint256 public constant HEADER_OFFSET_numSys = 68;


    function getTotalBytes(bytes calldata header, uint32 totalChunks) internal pure returns(uint256) {
        uint256 numCoefficients;
        assembly {
            //numCoefficients = totalChunks * (degree + 1)
            //NOTE: degree + 1 is the number of coefficients
            numCoefficients := mul(totalChunks, add(shr(BIT_SHIFT_degree, calldataload(add(header.offset, HEADER_OFFSET_degree))), 1))
        }
        return numCoefficients * BYTES_PER_COEFFICIENT;
    }
    /// @param header of the datastore that the coding ratio is being retrieved for
    /// @param totalChunks the total number of chunks expected in this datastore
    /// @return codingRatio of the datastore in basis points
    function getCodingRatio(bytes calldata header, uint32 totalChunks) internal pure returns(uint16) {
        uint32 codingRatio;
        assembly {
            //codingRatio = numSys
            codingRatio := shr(BIT_SHIFT_numSys, calldataload(add(header.offset, HEADER_OFFSET_numSys)))
            //codingRatio = numSys * BIP_MULTIPLIER / totalChunks
            codingRatio := div(mul(codingRatio, BIP_MULTIPLIER), totalChunks)
        }
        return uint16(codingRatio);
    }

    function getDegree(bytes calldata header) internal pure returns (uint32) {
        uint32 degree;
        assembly {
            degree := shr(BIT_SHIFT_degree, calldataload(add(header.offset, HEADER_OFFSET_degree)))
        }
        return degree;
    }

    /// @notice Finds the `signatoryRecordHash`, used for fraudproofs.
    function computeSignatoryRecordHash(
        uint32 globalDataStoreId,
        bytes32[] memory nonSignerPubkeyHashes,
        uint256 signedStakeFirstQuorum,
        uint256 signedStakeSecondQuorum
    ) internal pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(globalDataStoreId, nonSignerPubkeyHashes, signedStakeFirstQuorum, signedStakeSecondQuorum)
        );
    }

    /// @notice Computes the hash of a single DataStore's metadata.
    function computeDataStoreHash(IDataLayrServiceManager.DataStoreMetadata memory metadata)
        internal
        pure
        returns (bytes32)
    {
        bytes32 dsHash = keccak256(
            abi.encodePacked(
                metadata.headerHash,
                metadata.durationDataStoreId,
                metadata.globalDataStoreId,
                metadata.blockNumber,
                metadata.fee,
                metadata.confirmer,
                metadata.signatoryRecordHash
            )
        );
        return dsHash;
    }

    /// @notice uses `abi.encodePacked` to encode a DataStore's metadata into a compressed format
    function packDataStoreMetadata(IDataLayrServiceManager.DataStoreMetadata memory metadata)
        internal
        pure
        returns (bytes memory)
    {
        return (
            abi.encodePacked(
                metadata.headerHash,
                metadata.durationDataStoreId,
                metadata.globalDataStoreId,
                metadata.blockNumber,
                metadata.fee,
                metadata.confirmer,
                metadata.signatoryRecordHash
            )
        );
    }

    /// @notice uses `abi.encodePacked` to encode a DataStore's searchData into a compressed format
    function packDataStoreSearchData(IDataLayrServiceManager.DataStoreSearchData memory searchData)
        internal
        pure
        returns (bytes memory)
    {
        return (
            abi.encodePacked(
                packDataStoreMetadata(searchData.metadata), searchData.duration, searchData.timestamp, searchData.index
            )
        );
    }

    // CONSTANTS -- commented out lines are due to inline assembly supporting *only* 'direct number constants' (for now, at least)
    // OBJECT BIT LENGTHS
    uint256 internal constant BIT_LENGTH_headerHash = 256;
    uint256 internal constant BIT_LENGTH_durationDataStoreId = 32;
    uint256 internal constant BIT_LENGTH_globalDataStoreId = 32;
    uint256 internal constant BIT_LENGTH_blockNumber = 32;
    uint256 internal constant BIT_LENGTH_fee = 96;
    uint256 internal constant BIT_LENGTH_confirmer = 160;
    uint256 internal constant BIT_LENGTH_signatoryRecordHash = 256;
    uint256 internal constant BIT_LENGTH_duration = 8;
    uint256 internal constant BIT_LENGTH_timestamp = 256;
    uint256 internal constant BIT_LENGTH_index = 32;

    // OBJECT BIT SHIFTS FOR READING FROM CALLDATA -- don't bother with using 'shr' if any of these is 0
    // uint256 internal constant BIT_SHIFT_headerHash = 256 - BIT_LENGTH_headerHash;
    // uint256 internal constant BIT_SHIFT_durationDataStoreId = 256 - BIT_LENGTH_durationDataStoreId;
    // uint256 internal constant BIT_SHIFT_globalDataStoreId = 256 - BIT_LENGTH_globalDataStoreId;
    // uint256 internal constant BIT_SHIFT_blockNumber = 256 - BIT_LENGTH_blockNumber;
    // uint256 internal constant BIT_SHIFT_fee = 256 - BIT_LENGTH_fee;
    // uint256 internal constant BIT_SHIFT_confirmer = 256 - BIT_LENGTH_confirmer;
    // uint256 internal constant BIT_SHIFT_signatoryRecordHash = 256 - BIT_LENGTH_signatoryRecordHash;
    // uint256 internal constant BIT_SHIFT_duration = 256 - BIT_LENGTH_duration;
    // uint256 internal constant BIT_SHIFT_timestamp = 256 - BIT_LENGTH_timestamp;
    // uint256 internal constant BIT_SHIFT_index = 256 - BIT_LENGTH_index;
    uint256 internal constant BIT_SHIFT_headerHash = 0;
    uint256 internal constant BIT_SHIFT_durationDataStoreId = 224;
    uint256 internal constant BIT_SHIFT_globalDataStoreId = 224;
    uint256 internal constant BIT_SHIFT_blockNumber = 224;
    uint256 internal constant BIT_SHIFT_fee = 160;
    uint256 internal constant BIT_SHIFT_confirmer = 96;
    uint256 internal constant BIT_SHIFT_signatoryRecordHash = 0;
    uint256 internal constant BIT_SHIFT_duration = 248;
    uint256 internal constant BIT_SHIFT_timestamp = 0;
    uint256 internal constant BIT_SHIFT_index = 224;

    // CALLDATA OFFSETS IN BYTES -- adding 7 and dividing by 8 here is for rounding *up* the bit amounts to bytes amounts
    // uint256 internal constant CALLDATA_OFFSET_headerHash = 0;
    // uint256 internal constant CALLDATA_OFFSET_durationDataStoreId = ((BIT_LENGTH_headerHash + 7) / 8);
    // uint256 internal constant CALLDATA_OFFSET_globalDataStoreId = CALLDATA_OFFSET_durationDataStoreId + ((BIT_LENGTH_durationDataStoreId + 7) / 8);
    // uint256 internal constant CALLDATA_OFFSET_blockNumber = CALLDATA_OFFSET_globalDataStoreId + ((BIT_LENGTH_globalDataStoreId + 7) / 8);
    // uint256 internal constant CALLDATA_OFFSET_fee = CALLDATA_OFFSET_blockNumber + ((BIT_LENGTH_blockNumber + 7) / 8);
    // uint256 internal constant CALLDATA_OFFSET_confirmer = CALLDATA_OFFSET_fee + ((BIT_LENGTH_fee + 7) / 8);
    // uint256 internal constant CALLDATA_OFFSET_signatoryRecordHash = CALLDATA_OFFSET_confirmer + ((BIT_LENGTH_confirmer + 7) / 8);
    // uint256 internal constant CALLDATA_OFFSET_duration = CALLDATA_OFFSET_signatoryRecordHash + ((BIT_LENGTH_signatoryRecordHash + 7) / 8);
    // uint256 internal constant CALLDATA_OFFSET_timestamp = CALLDATA_OFFSET_duration + ((BIT_LENGTH_duration + 7) / 8);
    // uint256 internal constant CALLDATA_OFFSET_index = CALLDATA_OFFSET_timestamp + ((BIT_LENGTH_timestamp + 7) / 8);
    uint256 internal constant CALLDATA_OFFSET_headerHash = 0;
    uint256 internal constant CALLDATA_OFFSET_durationDataStoreId = 32;
    uint256 internal constant CALLDATA_OFFSET_globalDataStoreId = 36;
    uint256 internal constant CALLDATA_OFFSET_blockNumber = 40;
    uint256 internal constant CALLDATA_OFFSET_fee = 44;
    uint256 internal constant CALLDATA_OFFSET_confirmer = 56;
    uint256 internal constant CALLDATA_OFFSET_signatoryRecordHash = 76;
    uint256 internal constant CALLDATA_OFFSET_duration = 108;
    uint256 internal constant CALLDATA_OFFSET_timestamp = 109;
    uint256 internal constant CALLDATA_OFFSET_index = 141;

    // MEMORY OFFSETS IN BYTES
    uint256 internal constant MEMORY_OFFSET_headerHash = 0;
    uint256 internal constant MEMORY_OFFSET_durationDataStoreId = 32;
    uint256 internal constant MEMORY_OFFSET_globalDataStoreId = 64;
    uint256 internal constant MEMORY_OFFSET_blockNumber = 96;
    uint256 internal constant MEMORY_OFFSET_fee = 128;
    uint256 internal constant MEMORY_OFFSET_confirmer = 160;
    uint256 internal constant MEMORY_OFFSET_signatoryRecordHash = 192;
    // I'm unsure why the memory-offsets work this way, but they do. See usage below.
    uint256 internal constant MEMORY_OFFSET_duration = 32;
    uint256 internal constant MEMORY_OFFSET_timestamp = 64;
    uint256 internal constant MEMORY_OFFSET_index = 96;

    /**
     * @notice Unpacks the packed metadata of a DataStore into a metadata struct.
     * @param packedMetadata should be in the same form as the output of `packDataStoreMetadata`
     */
    function unpackDataStoreMetadata(bytes calldata packedMetadata)
        internal
        pure
        returns (IDataLayrServiceManager.DataStoreMetadata memory metadata)
    {
        uint256 pointer;
        assembly {
            // fetch offset of `packedMetadata` input in calldata
            pointer := packedMetadata.offset
            mstore(
                // store in the headerHash memory location in `metadata`
                metadata,
                // read the headerHash from its calldata position in `packedMetadata`
                calldataload(pointer)
            )
            mstore(
                // store in the durationDataStoreId memory location in `metadata`
                add(metadata, MEMORY_OFFSET_durationDataStoreId),
                // read the durationDataStoreId from its calldata position in `packedMetadata`
                shr(BIT_SHIFT_durationDataStoreId, calldataload(add(pointer, CALLDATA_OFFSET_durationDataStoreId)))
            )
            mstore(
                // store in the globalDataStoreId memory location in `metadata`
                add(metadata, MEMORY_OFFSET_globalDataStoreId),
                // read the globalDataStoreId from its calldata position in `packedMetadata`
                shr(BIT_SHIFT_globalDataStoreId, calldataload(add(pointer, CALLDATA_OFFSET_globalDataStoreId)))
            )
            mstore(
                // store in the blockNumber memory location in `metadata`
                add(metadata, MEMORY_OFFSET_blockNumber),
                // read the blockNumber from its calldata position in `packedMetadata`
                shr(BIT_SHIFT_blockNumber, calldataload(add(pointer, CALLDATA_OFFSET_blockNumber)))
            )
            mstore(
                // store in the fee memory location in `metadata`
                add(metadata, MEMORY_OFFSET_fee),
                // read the fee from its calldata position in `packedMetadata`
                shr(BIT_SHIFT_fee, calldataload(add(pointer, CALLDATA_OFFSET_fee)))
            )
            mstore(
                // store in the confirmer memory location in `metadata`
                add(metadata, MEMORY_OFFSET_confirmer),
                // read the confirmer from its calldata position in `packedMetadata`
                shr(BIT_SHIFT_confirmer, calldataload(add(pointer, CALLDATA_OFFSET_confirmer)))
            )
            mstore(
                // store in the signatoryRecordHash memory location in `metadata`
                add(metadata, MEMORY_OFFSET_signatoryRecordHash),
                // read the signatoryRecordHash from its calldata position in `packedMetadata`
                calldataload(add(pointer, CALLDATA_OFFSET_signatoryRecordHash))
            )
        }
        return metadata;
    }

    /**
     * @notice Unpacks the packed searchData of a DataStore into a searchData struct.
     * @param packedSearchData should be in the same form as the output of `packDataStoreSearchData`
     */
    function unpackDataStoreSearchData(bytes calldata packedSearchData)
        internal
        pure
        returns (IDataLayrServiceManager.DataStoreSearchData memory searchData)
    {
        searchData.metadata = (unpackDataStoreMetadata(packedSearchData));
        uint256 pointer;
        assembly {
            // fetch offset of `packedSearchData` input in calldata
            pointer := packedSearchData.offset
            mstore(
                // store in the duration memory location of `searchData`
                add(searchData, MEMORY_OFFSET_duration),
                // read the duration from its calldata position in `packedSearchData`
                shr(BIT_SHIFT_duration, calldataload(add(pointer, CALLDATA_OFFSET_duration)))
            )
            mstore(
                // store in the timestamp memory location of `searchData`
                add(searchData, MEMORY_OFFSET_timestamp),
                // read the timestamp from its calldata position in `packedSearchData`
                calldataload(add(pointer, CALLDATA_OFFSET_timestamp))
            )
            mstore(
                // store in the index memory location of `searchData`
                add(searchData, MEMORY_OFFSET_index),
                // read the index from its calldata position in `packedSearchData`
                shr(BIT_SHIFT_index, calldataload(add(pointer, CALLDATA_OFFSET_index)))
            )
        }
        return searchData;
    }
}

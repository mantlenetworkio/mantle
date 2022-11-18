// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./utils/DataStoreUtilsWrapper.sol";

import "forge-std/Test.sol";

contract DataStoreUtilsTests is DSTest {
    DataStoreUtilsWrapper public dataStoreUtilsWrapper;

    function setUp() public {
        // deploy library wrapper contract so that we can call the library's functions that take inputs with 'calldata' location specified
        dataStoreUtilsWrapper = new DataStoreUtilsWrapper();
    }

    // unfuzzed version of testPackUnpackDataStoreMetadata
    function testPackUnpackDataStoreMetadataFixedInputs() public {
        testPackUnpackDataStoreMetadata(bytes32(uint256(1)), 2, 3, 4, 5, address(6), bytes32(uint256(7)));
    }

    // unfuzzed version of testPackUnpackDataStoreSearchData
    function testPackUnpackDataStoreSearchDataFixedInputs() public {
        testPackUnpackDataStoreSearchData(bytes32(uint256(1)), 2, 3, 4, 5, address(6), bytes32(uint256(7)), 8, 9, 10);
    }

    function testPackUnpackDataStoreMetadata(
        bytes32 headerHash,
        uint32 durationDataStoreId,
        uint32 globalDataStoreId,
        uint32 blockNumber,
        uint96 fee,
        address confirmer,
        bytes32 signatoryRecordHash
    ) public {
        // form struct from arguments
        IDataLayrServiceManager.DataStoreMetadata memory metadataStructBeforePacking = dataStoreMetadataFromArgs(
            headerHash, durationDataStoreId, globalDataStoreId, blockNumber, fee, confirmer, signatoryRecordHash
        );
        // pack the struct
        bytes memory packedMetadata = dataStoreUtilsWrapper.packDataStoreMetadataExternal(metadataStructBeforePacking);
        // unpack the struct
        IDataLayrServiceManager.DataStoreMetadata memory unpackedStruct =
            dataStoreUtilsWrapper.unpackDataStoreMetadataExternal(packedMetadata);
        // check the struct entries
        assertEq(
            headerHash,
            unpackedStruct.headerHash,
            "testPackUnpackDataStoreMetadata: unpacked headerHash does not match original one"
        );
        assertEq(
            durationDataStoreId,
            unpackedStruct.durationDataStoreId,
            "testPackUnpackDataStoreMetadata: unpacked durationDataStoreId does not match original one"
        );
        assertEq(
            globalDataStoreId,
            unpackedStruct.globalDataStoreId,
            "testPackUnpackDataStoreMetadata: unpacked globalDataStoreId does not match original one"
        );
        assertEq(
            blockNumber,
            unpackedStruct.blockNumber,
            "testPackUnpackDataStoreMetadata: unpacked blockNumber does not match original one"
        );
        assertEq(fee, unpackedStruct.fee, "testPackUnpackDataStoreMetadata: unpacked fee does not match original one");
        assertEq(
            confirmer,
            unpackedStruct.confirmer,
            "testPackUnpackDataStoreMetadata: unpacked confirmer does not match original one"
        );
        assertEq(
            signatoryRecordHash,
            unpackedStruct.signatoryRecordHash,
            "testPackUnpackDataStoreMetadata: unpacked signatoryRecordHash does not match original one"
        );

        // failsafe extra check, in case we modify the struct entries and forget a specific check above -- just check the full bytes against each other
        require(
            keccak256(abi.encode(metadataStructBeforePacking)) == keccak256(abi.encode(unpackedStruct)),
            "testPackUnpackDataStoreMetadata: keccak256(abi.encode(metadataStructBeforePacking)) != keccak256(abi.encode(unpackedStruct))"
        );
    }

    function testPackUnpackDataStoreSearchData(
        bytes32 headerHash,
        uint32 durationDataStoreId,
        uint32 globalDataStoreId,
        uint32 blockNumber,
        uint96 fee,
        address confirmer,
        bytes32 signatoryRecordHash,
        uint8 duration,
        uint256 timestamp,
        uint32 index
    ) public {
        // form struct from arguments
        IDataLayrServiceManager.DataStoreSearchData memory searchDataStructBeforePacking = dataStoreSearchDataFromArgs(
            headerHash,
            durationDataStoreId,
            globalDataStoreId,
            blockNumber,
            fee,
            confirmer,
            signatoryRecordHash,
            duration,
            timestamp,
            index
        );

        // pack the struct
        bytes memory packedSearchData =
            dataStoreUtilsWrapper.packDataStoreSearchDataExternal(searchDataStructBeforePacking);
        // unpack the struct
        IDataLayrServiceManager.DataStoreSearchData memory unpackedStruct =
            dataStoreUtilsWrapper.unpackDataStoreSearchDataExternal(packedSearchData);
        // check the struct entries
        assertEq(
            headerHash,
            unpackedStruct.metadata.headerHash,
            "testPackUnpackDataStoreSearchData: unpacked headerHash does not match original one"
        );
        assertEq(
            durationDataStoreId,
            unpackedStruct.metadata.durationDataStoreId,
            "testPackUnpackDataStoreSearchData: unpacked durationDataStoreId does not match original one"
        );
        assertEq(
            globalDataStoreId,
            unpackedStruct.metadata.globalDataStoreId,
            "testPackUnpackDataStoreSearchData: unpacked globalDataStoreId does not match original one"
        );
        assertEq(
            blockNumber,
            unpackedStruct.metadata.blockNumber,
            "testPackUnpackDataStoreSearchData: unpacked blockNumber does not match original one"
        );
        assertEq(
            fee,
            unpackedStruct.metadata.fee,
            "testPackUnpackDataStoreSearchData: unpacked fee does not match original one"
        );
        assertEq(
            confirmer,
            unpackedStruct.metadata.confirmer,
            "testPackUnpackDataStoreSearchData: unpacked confirmer does not match original one"
        );
        assertEq(
            signatoryRecordHash,
            unpackedStruct.metadata.signatoryRecordHash,
            "testPackUnpackDataStoreSearchData: unpacked signatoryRecordHash does not match original one"
        );
        assertEq(
            duration,
            unpackedStruct.duration,
            "testPackUnpackDataStoreSearchData: unpacked duration does not match original one"
        );
        assertEq(
            timestamp,
            unpackedStruct.timestamp,
            "testPackUnpackDataStoreSearchData: unpacked timestamp does not match original one"
        );
        assertEq(
            index, unpackedStruct.index, "testPackUnpackDataStoreSearchData: unpacked index does not match original one"
        );

        // failsafe extra check, in case we modify the struct entries and forget a specific check above -- just check the full bytes against each other
        require(
            keccak256(abi.encode(searchDataStructBeforePacking)) == keccak256(abi.encode(unpackedStruct)),
            "testPackUnpackDataStoreSearchData: keccak256(abi.encode(searchDataStructBeforePacking)) != keccak256(abi.encode(unpackedStruct))"
        );
    }

    function dataStoreMetadataFromArgs(
        bytes32 headerHash,
        uint32 durationDataStoreId,
        uint32 globalDataStoreId,
        uint32 blockNumber,
        uint96 fee,
        address confirmer,
        bytes32 signatoryRecordHash
    ) internal pure returns (IDataLayrServiceManager.DataStoreMetadata memory metadataStruct) {
        metadataStruct = IDataLayrServiceManager.DataStoreMetadata({
            headerHash: headerHash,
            durationDataStoreId: durationDataStoreId,
            globalDataStoreId: globalDataStoreId,
            blockNumber: blockNumber,
            fee: fee,
            confirmer: confirmer,
            signatoryRecordHash: signatoryRecordHash
        });
        return metadataStruct;
    }

    function dataStoreSearchDataFromArgs(
        bytes32 headerHash,
        uint32 durationDataStoreId,
        uint32 globalDataStoreId,
        uint32 blockNumber,
        uint96 fee,
        address confirmer,
        bytes32 signatoryRecordHash,
        uint8 duration,
        uint256 timestamp,
        uint32 index
    ) internal pure returns (IDataLayrServiceManager.DataStoreSearchData memory searchDataStruct) {
        searchDataStruct.metadata = dataStoreMetadataFromArgs(
            headerHash, durationDataStoreId, globalDataStoreId, blockNumber, fee, confirmer, signatoryRecordHash
        );
        searchDataStruct.duration = duration;
        searchDataStruct.timestamp = timestamp;
        searchDataStruct.index = index;
        return searchDataStruct;
    }
}

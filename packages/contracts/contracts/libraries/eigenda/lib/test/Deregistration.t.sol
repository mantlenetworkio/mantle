// SPDX-License-Verifier: UNLICENSED
pragma solidity ^0.8.9;

import "./DataLayrTestHelper.t.sol";
import "../contracts/libraries/BytesLib.sol";

contract DeregistrationTests is DataLayrTestHelper {
    using BytesLib for bytes;

    /**
    *   @notice Tests that optimistically deregistering works as intended, i.e., 
    *   the aggregate public key is updated correctly, and all storage is correctly updated.
    */
    function testBLSDeregistration(
        uint8 operatorIndex,
        uint256 ethAmount, 
        uint256 eigenAmount
    ) public fuzzedOperatorIndex(operatorIndex) {

        //TODO: probably a stronger test would be to register a few operators and then ensure that apk is updated correctly
        uint256[4] memory prevAPK;
        prevAPK[0] = dlReg.apk(0);
        prevAPK[1] = dlReg.apk(1);
        prevAPK[2] = dlReg.apk(2);
        prevAPK[3] = dlReg.apk(3);
        bytes32 prevAPKHash = BLS.hashPubkey(prevAPK);

        BLSRegistration(operatorIndex, ethAmount, eigenAmount);


        uint256[4] memory pubkeyToRemoveAff = getG2PKOfRegistrationData(operatorIndex);

        bytes32 pubkeyHash = BLS.hashPubkey(pubkeyToRemoveAff);                          

        _testDeregisterOperatorWithDataLayr(operatorIndex, pubkeyToRemoveAff, uint8(dlReg.numOperators()-1), testEphemeralKey);

        (,uint32 nextUpdateBlockNumber,uint96 firstQuorumStake, uint96 secondQuorumStake) = dlReg.pubkeyHashToStakeHistory(pubkeyHash, dlReg.getStakeHistoryLength(pubkeyHash)-1);
        require( nextUpdateBlockNumber == 0, "Stake history not updated correctly");
        require( firstQuorumStake == 0, "Stake history not updated correctly");
        require( secondQuorumStake == 0, "Stake history not updated correctly");

        bytes32 currAPKHash = dlReg.apkHashes(dlReg.getApkUpdatesLength()-1);
        require(currAPKHash == prevAPKHash, "aggregate public key has not been updated correctly following deregistration");

    }

    /**
    *   @notice Tests that deregistering with an incorrect public key
    *           reverts.
    */
    function testMismatchedPubkeyHashAndProvidedPubkeyHash(
        uint8 operatorIndex,
        uint256 ethAmount, 
        uint256 eigenAmount,
        uint256[4] memory pubkeyToRemoveAff
    ) public fuzzedOperatorIndex(operatorIndex) {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e18);
        cheats.assume(BLS.hashPubkey(pubkeyToRemoveAff) != BLS.hashPubkey(getG2PKOfRegistrationData(operatorIndex)));

    
        BLSRegistration(operatorIndex, ethAmount, eigenAmount);
        uint8 operatorListIndex = uint8(dlReg.numOperators()-1);
        cheats.expectRevert(bytes("BLSRegistry._deregisterOperator: pubkey input does not match stored pubkeyHash"));
        _testDeregisterOperatorWithDataLayr(operatorIndex, pubkeyToRemoveAff, operatorListIndex, testEphemeralKey);
    }

    /**
    *   @notice Tests that posting an ephemeral key that does
    *   not match the posted ek hash in the ekregistry
    *   reverts the deregistration.
    */
    function testEphemeralKeyDoesNotMatchPostedHash(
        uint8 operatorIndex,
        uint256 ethAmount, 
        uint256 eigenAmount,
        bytes32 badEphemeralKey
    ) public fuzzedOperatorIndex(operatorIndex) {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e18);
        cheats.assume(badEphemeralKey != testEphemeralKey);

        BLSRegistration(operatorIndex, ethAmount, eigenAmount);

        uint256[4] memory pubkeyToRemoveAff = getG2PKOfRegistrationData(operatorIndex);
        uint8 operatorListIndex = uint8(dlReg.numOperators()-1);
        cheats.expectRevert(bytes("EphemeralKeyRegistry.postLastEphemeralKeyPreImage: Ephemeral key does not match previous ephemeral key commitment"));
        _testDeregisterOperatorWithDataLayr(operatorIndex, pubkeyToRemoveAff, operatorListIndex, badEphemeralKey);
    }
    /**
    *   @notice Tests that deregistering an operator who has already 
    *           been deregistered/was never registered reverts
    */
    function testDeregisteringAlreadyDeregisteredOperator(
        uint8 operatorIndex,
        uint256 ethAmount, 
        uint256 eigenAmount
    ) public fuzzedOperatorIndex(operatorIndex) {

        BLSRegistration(operatorIndex, ethAmount, eigenAmount);

        uint256[4] memory pubkeyToRemoveAff = getG2PKOfRegistrationData(operatorIndex);      

        uint8 operatorListIndex = uint8(dlReg.numOperators());                  
        _testDeregisterOperatorWithDataLayr(operatorIndex, pubkeyToRemoveAff, operatorListIndex-1, testEphemeralKey);
        
        cheats.expectRevert(bytes("RegistryBase._deregistrationCheck: Operator is not registered"));
        _testDeregisterOperatorWithDataLayr(operatorIndex, pubkeyToRemoveAff, operatorListIndex-1, testEphemeralKey);

    }


    /// @notice Helper function that performs registration 
    function BLSRegistration(
        uint8 operatorIndex,
        uint256 ethAmount, 
        uint256 eigenAmount
    ) internal fuzzedOperatorIndex(operatorIndex) {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e18);
        
        uint8 operatorType = 3;
        _testInitiateDelegation(
            operatorIndex,
            eigenAmount,
            ethAmount
        );
        _testRegisterBLSPubKey(operatorIndex);
        _testRegisterOperatorWithDataLayr(
            operatorIndex,
            operatorType,
            testEphemeralKeyHash,
            testSocket
        );
    }

}
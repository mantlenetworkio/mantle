// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./DataLayrTestHelper.t.sol";
import "../contracts/libraries/BytesLib.sol";

contract RegistrationTests is DataLayrTestHelper {
    using BytesLib for bytes;

    /// @notice This test ensures that the optimistic flow for BLS registration 
    ///         works as intended by checking storage updates in the registration contracts.
    function testBLSRegistration(
        uint8 operatorIndex,
        uint256 ethAmount, 
        uint256 eigenAmount
    ) fuzzedOperatorIndex(operatorIndex) public {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e18);
        
        uint8 operatorType = 3;
        (
            uint256 amountEthStaked, 
            uint256 amountEigenStaked
        ) = _testInitiateDelegation(
                operatorIndex,
                eigenAmount,
                ethAmount
            );

        _testRegisterBLSPubKey(operatorIndex);

        uint256[4] memory pk;
        pk = getG2PKOfRegistrationData(operatorIndex);
        bytes32 hashofPk = BLS.hashPubkey(pk);
        require(pubkeyCompendium.operatorToPubkeyHash(signers[operatorIndex]) == hashofPk, "hash not stored correctly");
        require(pubkeyCompendium.pubkeyHashToOperator(hashofPk) == signers[operatorIndex], "hash not stored correctly");

        {

            uint96 ethStakedBefore = dlReg.getTotalStakeFromIndex(dlReg.getLengthOfTotalStakeHistory()-1).firstQuorumStake;
            uint96 eigenStakedBefore = dlReg.getTotalStakeFromIndex(dlReg.getLengthOfTotalStakeHistory()-1).secondQuorumStake;
            _testRegisterOperatorWithDataLayr(
                operatorIndex,
                operatorType,
                testEphemeralKeyHash,
                testSocket
            );

            uint256 numOperators = dlReg.numOperators();
            require(dlReg.operatorList(numOperators-1) == signers[operatorIndex], "operatorList not updated");

        
            uint96 ethStakedAfter = dlReg.getTotalStakeFromIndex(dlReg.getLengthOfTotalStakeHistory()-1).firstQuorumStake;
            uint96 eigenStakedAfter = dlReg.getTotalStakeFromIndex(dlReg.getLengthOfTotalStakeHistory()-1).secondQuorumStake;


            require(ethStakedAfter - ethStakedBefore == amountEthStaked, "eth quorum staked value not updated correctly");
            require(eigenStakedAfter - eigenStakedBefore == amountEigenStaked, "eigen quorum staked value not updated correctly");
        }
    }

    /// @notice Tests that registering the same public key twice reverts appropriately.
    function testRegisterPublicKeyTwice(uint8 operatorIndex) fuzzedOperatorIndex(operatorIndex) public {
        cheats.startPrank(signers[operatorIndex]);
        //try to register the same pubkey twice
        pubkeyCompendium.registerBLSPublicKey(registrationData[operatorIndex]);
        cheats.expectRevert(
            "BLSPublicKeyCompendium.registerBLSPublicKey: operator already registered pubkey"
        );
        pubkeyCompendium.registerBLSPublicKey(registrationData[operatorIndex]);
    }

    /// @notice Tests that re-registering while an msg.sender is actively registered reverts.
    function testRegisterWhileAlreadyActive(
        uint8 operatorIndex, 
        uint256 ethAmount, 
        uint256 eigenAmount
    ) fuzzedOperatorIndex(operatorIndex) public {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e18);
        
        uint8 operatorType = 3;
        _testInitiateDelegation(
            operatorIndex,
            eigenAmount,
            ethAmount
        );
        _testRegisterBLSPubKey(
            operatorIndex
        );
        _testRegisterOperatorWithDataLayr(
            operatorIndex,
            operatorType,
            testEphemeralKeyHash,
            testSocket
        );
        cheats.startPrank(signers[operatorIndex]);

        //try to register after already registered
        cheats.expectRevert(
            "RegistryBase._registrationStakeEvaluation: Operator is already registered"
        );
        dlReg.registerOperator(
            3,
            bytes32(0),
            registrationData[operatorIndex].slice(0, 128),
            testSocket
        );
        cheats.stopPrank();
    }

    /// @notice Test that when operator tries to register with DataLayr with a public key 
    ///         that they haven't registered in the BLSPublicKeyCompendium, it fails
    function testOperatorDoesNotOwnPublicKey(
        uint8 operatorIndex, 
        uint256 ethAmount, 
        uint256 eigenAmount
    ) fuzzedOperatorIndex(operatorIndex) public {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e18);

        uint8 operatorType = 3;
        _testInitiateDelegation(
            operatorIndex,
            eigenAmount,
            ethAmount
        );
        //registering the operator without having registered their BLS public key
        cheats.expectRevert(bytes("BLSRegistry._registerOperator: operator does not own pubkey"));

        _testRegisterOperatorWithDataLayr(
            operatorIndex,
            operatorType,
            testEphemeralKeyHash,
            testSocket
        );
    } 
    /// @notice Tests that registering without having delegated in any quorum reverts
    function testRegisterForDataLayrWithNeitherQuorum(
        uint8 operatorIndex,
        uint256 ethAmount,
        uint256 eigenAmount
    ) fuzzedOperatorIndex(operatorIndex) public {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e18);
        uint8 noQuorumOperatorType = 0;

        _testInitiateDelegation(
            operatorIndex,
            eigenAmount,
            ethAmount
        );
        _testRegisterBLSPubKey(
            operatorIndex
        );
        cheats.expectRevert(bytes("RegistryBase._registrationStakeEvaluation: Must register as at least one type of validator"));
        _testRegisterOperatorWithDataLayr(
            operatorIndex,
            noQuorumOperatorType,
            testEphemeralKeyHash,
            testSocket
        );
    }
    /// @notice Tests that registering without adequate quorum stake reverts
    function testRegisterWithoutEnoughQuorumStake(
        uint8 operatorIndex
    ) fuzzedOperatorIndex(operatorIndex) public {
        _testRegisterBLSPubKey(
            operatorIndex
        );

        uint8 operatorType = 1;
        cheats.expectRevert(bytes("RegistryBase._registrationStakeEvaluation: Must register as at least one type of validator"));
        _testRegisterOperatorWithDataLayr(operatorIndex, operatorType, testEphemeralKeyHash, testSocket);
        
        operatorType = 2;
        cheats.expectRevert(bytes("RegistryBase._registrationStakeEvaluation: Must register as at least one type of validator"));
        _testRegisterOperatorWithDataLayr(operatorIndex, operatorType, testEphemeralKeyHash, testSocket);

        operatorType = 3;
        cheats.expectRevert(bytes("RegistryBase._registrationStakeEvaluation: Must register as at least one type of validator"));
        _testRegisterOperatorWithDataLayr(operatorIndex, operatorType, testEphemeralKeyHash, testSocket);
    }


    /// @notice Test if registering with public key = 0 reverts.
    function testRegisterWithZeroPubKey(
        uint8 operatorIndex,
        uint256 ethAmount,
        uint256 eigenAmount
    ) fuzzedOperatorIndex(operatorIndex) public {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e18);

        bytes memory zeroData = hex"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";
        address operator = signers[operatorIndex];
        uint8 operatorType = 3;
        bytes32 apkHashBefore = dlReg.apkHashes(dlReg.getApkUpdatesLength()-1);
        emit log_named_bytes32("apkHashBefore", apkHashBefore);

        _testInitiateDelegation(
            operatorIndex,
            eigenAmount,
            ethAmount
        );
        cheats.startPrank(operator);
        //whitelist the dlsm to slash the operator
        slasher.allowToSlash(address(dlsm));

        cheats.expectRevert(bytes("BLSPublicKeyCompendium.registerBLSPublicKey: Cannot register with 0x0 public key"));
        pubkeyCompendium.registerBLSPublicKey(zeroData);

        cheats.expectRevert(bytes("BLSRegistry._registerOperator: Cannot register with 0x0 public key"));
        dlReg.registerOperator(operatorType, testEphemeralKeyHash, zeroData, testSocket);
        cheats.stopPrank(); 
    }

    /// @notice Tests for registering without having opted into slashing.
    function testRegisterWithoutSlashingOptIn(
        uint8 operatorIndex,
        uint256 ethAmount,
        uint256 eigenAmount
     ) fuzzedOperatorIndex(operatorIndex) public {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e18); 

        uint8 operatorType = 3;

        _testInitiateDelegation(
            operatorIndex,
            eigenAmount,
            ethAmount
        );

        cheats.startPrank(signers[operatorIndex]);
        pubkeyCompendium.registerBLSPublicKey(registrationData[operatorIndex]);
        cheats.stopPrank();

        cheats.expectRevert(bytes("RegistryBase._addRegistrant: operator must be opted into slashing by the serviceManager"));
        _testRegisterOperatorWithDataLayr(
            operatorIndex,
            operatorType,
            testEphemeralKeyHash,
            testSocket
        );
     }
    /// @notice Tests that when operator registers with the same public key 
    ///         as the current aggregate public key, it reverts.  Currently
    ///         signature aggregation doesn't support x + x.
    function testRegisteringWithSamePubKeyAsAggPubKey(
        uint8 operatorIndex,
        uint256 ethAmount,
        uint256 eigenAmount
     ) fuzzedOperatorIndex(operatorIndex) public {
        cheats.assume(ethAmount > 0 && ethAmount < 1e18);
        cheats.assume(eigenAmount > 0 && eigenAmount < 1e18);
        uint256[4] memory prevAPK;
        prevAPK[0] = dlReg.apk(0);
        prevAPK[1] = dlReg.apk(1);
        prevAPK[2] = dlReg.apk(2);
        prevAPK[3] = dlReg.apk(3);
        bytes memory packedAPK = abi.encodePacked(
                                    bytes32(prevAPK[1]),
                                    bytes32(prevAPK[0]),
                                    bytes32(prevAPK[3]),
                                    bytes32(prevAPK[2])
                                    );
        
        uint8 operatorType = 3;

        _testInitiateDelegation(
            operatorIndex,
            eigenAmount,
            ethAmount
        );
         cheats.startPrank(signers[operatorIndex]);
        // pubkeyCompendium.registerBLSPublicKey(packedAPK);

        cheats.expectRevert(bytes("BLSRegistry._registerOperator: Apk and pubkey cannot be the same"));
        dlReg.registerOperator(operatorType, testEphemeralKeyHash, packedAPK, testSocket);
        cheats.stopPrank();
     }
}

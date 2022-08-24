// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "./ITssGroupManager.sol";


contract TssGroupManager is OwnableUpgradeable, ReentrancyGuardUpgradeable, ITssGroupManager {
    using SafeMathUpgradeable for uint256;
    using ECDSAUpgradeable for bytes32;
    using AddressUpgradeable for address;
    bytes confirmGroupPublicKey;
    uint256 threshold;
    uint256 gRoundId;
    uint256 confirmNumber;
    address[] addresses;

    event TssMemberAppended(
        uint256       _roundId,
        bytes         _publicKey,
        address       _nodeAddress,
        MemberStatus  _status
    );

    TssMember[] ActiveTssMembers;
    bytes[] InActiveTssMember;
    mapping (bytes => bytes) public memberGroupKey;        // memberKey=>groupKey
    mapping (bytes => TssMember) public TssMemberInfo;     // address=>TssMember

    /**
    * @inheritdoc ITssGroupManager
    */
    // slither-disable-next-line external-function
    function setTssGroupMember(uint256 _threshold, bytes[] memory _batchPublicKey) public override onlyOwner {
        require(
            (InActiveTssMember.length > 0),
            "inactive tss member array is not empty"
        );
        for(uint i = 0; i < _batchPublicKey.length; i++) {
            InActiveTssMember.push(_batchPublicKey[i]);
        }
        threshold = _threshold;
        gRoundId = gRoundId + 1;
    }

    /**
    * @inheritdoc ITssGroupManager
    */
    // slither-disable-next-line external-function
    function setGroupPublicKey(bytes memory _publicKey, bytes memory _groupPublicKey) public override {
        for(uint i = 0; i < InActiveTssMember.length; i++) {
            require(
                isEqual(InActiveTssMember[i], _publicKey),
                "tss public key is not in InActiveTssMember"
            );
        }
        if (memberGroupKey[_publicKey].length == 0){
            confirmNumber = confirmNumber + 1;
        }
        memberGroupKey[_publicKey] = _groupPublicKey;
        if (confirmNumber == InActiveTssMember.length) {
            for(uint i = 1; i <= InActiveTssMember.length; i++) {
                bytes memory tempGroupKey = InActiveTssMember[i-1];
                require(
                    isEqual(InActiveTssMember[i-1], InActiveTssMember[i]),
                    "tss members groupPublicKey is not equal"
                );
            }
            updateTssMember();
        }
    }

    /**
    * @inheritdoc ITssGroupManager
    */
    // slither-disable-next-line external-function
    function getTssGroupInfo() public override returns (uint256, uint256, bytes memory, TssMember[] memory) {
        return (gRoundId, threshold, confirmGroupPublicKey, ActiveTssMembers);
    }

    /**
    * @inheritdoc ITssGroupManager
    */
    // slither-disable-next-line external-function
    function memberJail(bytes memory _publicKey) public override {
        for(uint i = 0; i < ActiveTssMembers.length; i++) {
            if (isEqual(ActiveTssMembers[i].publicKey, _publicKey)) {
                ActiveTssMembers[i].status = MemberStatus.jail;
                TssMemberInfo[_publicKey] = ActiveTssMembers[i];
            }
        }
    }

    /**
    * @inheritdoc ITssGroupManager
    */
    // slither-disable-next-line external-function
    function memberUnJail(bytes memory _publicKey) public override {
        for(uint i = 0; i < ActiveTssMembers.length; i++) {
            if (isEqual(ActiveTssMembers[i].publicKey, _publicKey)) {
                ActiveTssMembers[i].status = MemberStatus.unJail;
                TssMemberInfo[_publicKey] = ActiveTssMembers[i];
            }
        }
    }

    /**
    * @inheritdoc ITssGroupManager
    */
    // slither-disable-next-line external-function
    function removeMember(bytes memory _publicKey) public override {
        for(uint i = 0; i < ActiveTssMembers.length; i++) {
            if (isEqual(ActiveTssMembers[i].publicKey, _publicKey)) {
                delete ActiveTssMembers[i];
                delete TssMemberInfo[_publicKey];
            }
        }
    }

    /**
    * @inheritdoc ITssGroupManager
    */
    // slither-disable-next-line external-function
    function getTssGroupUnJailMembers() public override returns (address[] memory) {
        delete addresses;
        for(uint i = 0; i < ActiveTssMembers.length; i++) {
            if (ActiveTssMembers[i].status == MemberStatus.unJail) {
                addresses.push(
                    ActiveTssMembers[i].nodeAddress
                );
            }
        }
        return addresses;
    }

    /**
    * @inheritdoc ITssGroupManager
    */
    // slither-disable-next-line external-function
    function getTssGroupMembers() public override returns (TssMember[] memory) {
        return ActiveTssMembers;
    }

    /**
    * @inheritdoc ITssGroupManager
    */
    // slither-disable-next-line external-function
    function getTssMember(bytes memory _publicKey) public override returns (TssMember memory) {
        return TssMemberInfo[_publicKey];
    }

    /**
    * @inheritdoc ITssGroupManager
    */
    // slither-disable-next-line external-function
    function memberExistActive(bytes memory _publicKey) public override returns (bool) {
        for (uint i =0; i < ActiveTssMembers.length; i++) {
            if (isEqual(InActiveTssMember[i], _publicKey)) {
                return true;
            }
        }
        return false;
    }

    /**
    * @inheritdoc ITssGroupManager
    */
    // slither-disable-next-line external-function
    function memberExistInActive(bytes memory _publicKey) public override returns (bool) {
        for (uint i =0; i < InActiveTssMember.length; i++) {
            if (isEqual(InActiveTssMember[i], _publicKey)) {
                return true;
            }
        }
        return false;
    }

    /**
    * @inheritdoc ITssGroupManager
    */
    function verifySign(bytes32 _message, bytes memory _sig) public override returns (bool, address) {
        bytes32 ethSignedMessageHash = getEthSignedMessageHash(_message);
        return (recover(ethSignedMessageHash, _sig) == publicKeyToAddress(confirmGroupPublicKey), recover(ethSignedMessageHash, _sig));
    }

    function publicKeyToAddress (bytes memory publicKey) public returns (address) {
        require (publicKey.length == 64);
        return address (uint160 (uint256 (keccak256 (publicKey))));
    }

    function updateTssMember() public {
        for(uint i = 1; i < InActiveTssMember.length; i++) {
            ActiveTssMembers.push(
                TssMember({
                    publicKey: InActiveTssMember[i],
                    nodeAddress: publicKeyToAddress(InActiveTssMember[i]),
                    status: MemberStatus.unJail
                })
            );
            emit TssMemberAppended(
                gRoundId,
                InActiveTssMember[i],
                publicKeyToAddress(InActiveTssMember[i]),
                MemberStatus.unJail
            );
        }
        confirmNumber = 0;
        confirmGroupPublicKey = memberGroupKey[InActiveTssMember[0]];
    }

    function getEthSignedMessageHash(bytes32 _messageHash) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(
                "\x19Ethereum Signed Message:\n32",
                _messageHash
            ));
    }

    function recover(bytes32 _ethSignedMessageHash, bytes memory _sig) public pure returns (address){
        (bytes32 r, bytes32 s, uint8 v) = _split(_sig);
        return ecrecover(_ethSignedMessageHash, v, r, s);
    }

    function _split(bytes memory _sig) internal pure returns (bytes32 r, bytes32 s, uint8 v){
        require(_sig.length == 65, "invalid signature length");
        assembly {
            r :=mload(add(_sig, 32))
            s :=mload(add(_sig, 64))
            v :=byte(0, mload(add(_sig, 96)))
        }
    }

    function isEqual(bytes memory byteListA, bytes memory byteListB) public pure returns (bool) {
        if (byteListA.length != byteListB.length) return false;
        for(uint i = 0; i < byteListA.length; i ++) {
            if(byteListA[i] != byteListB[i]) return false;
        }
        return true;
    }
}

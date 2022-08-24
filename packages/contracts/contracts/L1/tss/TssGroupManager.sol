//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "./ITssGroupManager.sol";


contract TssGroupManager is OwnableUpgradeable, ReentrancyGuardUpgradeable, Lib_AddressResolver, ITssGroupManager {
    using SafeMathUpgradeable for uint256;
    using ECDSAUpgradeable for bytes32;
    using AddressUpgradeable for address;
    bytes confirmGroupPublicKey;
    uint256 threshold;
    uint256 gRoundId;
    uint256 confirmNumber;

    event TssMemberAppended(
        uint256       _roundId,
        bytes         _publicKey,
        address       _nodeAddress,
        MemberStatus  _status
    );

    TssMember[] ActiveTssMembers;
    bytes[] InActiveTssMember;
    mapping (bytes => bytes) public memberGroupKey;        // memberKey=>groupKey
    mapping (address => TssMember) public TssMemberInfo;   // address=>TssMember

    constructor() Lib_AddressResolver(address(0)) {}

    function setTssGroupMember(uint256 _threshold, bytes[] memory _batchPublicKey) public onlyOwner {
        for(uint i = 0; i < _batchPublicKey.length; i++) {
            InActiveTssMember.push(_batchPublicKey[i]);
        }
        threshold = _threshold;
        gRoundId = gRoundId + 1;
    }

    function setGroupPublicKey(bytes memory _publicKey, bytes memory _groupPublicKey) public {
        for(uint i = 0; i < InActiveTssMember.length; i++) {
            require(
                isEqual(InActiveTssMember[i], _publicKey),
                "tss public key is not in InActiveTssMember"
            );
        }
        if (memberGroupKey[_publicKey]){
            confirmNumber = confirmNumber + 1;
        }
        memberGroupKey[_publicKey] = _groupPublicKey;
        if (confirmNumber == InActiveTssMember.length) {
            for(uint i = 1; i <= InActiveTssMember.length; i++) {
                bytes tempGroupKey = InActiveTssMember[i-1];
                require(
                    isEqual(InActiveTssMember[i-1], InActiveTssMember[i]),
                    "tss members groupPublicKey is not equal"
                );
            }
            updateTssMember();
        }
    }

    function getTssGroupInfo() public returns (uint256, uint256, bytes, TssMember[])  {
        return (gRoundId, threshold, confirmGroupPublicKey, ActiveTssMembers);
    }

    function memberJail(bytes memory _publicKey) public {
        for(uint i = 0; i < ActiveTssMembers.length; i++) {
            if (isEqual(ActiveTssMembers[i].publicKey, _publicKey)) {
                ActiveTssMembers[i].status = MemberStatus.jail;
                TssMemberInfo[publicKeyToAddress(_publicKey)] = ActiveTssMembers[i];
            }
        }
    }

    function memberUnJail(bytes memory _publicKey) public {
        for(uint i = 0; i < ActiveTssMembers.length; i++) {
            if (isEqual(ActiveTssMembers[i].publicKey, _publicKey)) {
                ActiveTssMembers[i].status = MemberStatus.unJail;
                TssMemberInfo[publicKeyToAddress(_publicKey)] = ActiveTssMembers[i];
            }
        }
    }

    function removeMember(bytes memory _publicKey) public {
        for(uint i = 0; i < ActiveTssMembers.length; i++) {
            if (isEqual(ActiveTssMembers[i].publicKey, _publicKey)) {
                delete ActiveTssMembers[i];
                delete TssMemberInfo[publicKeyToAddress(_publicKey)];
            }
        }
    }

    function getTssGroupUnJailMembers() public returns (address[] memory) {
        address[] addresses;
        for(uint i = 0; i < ActiveTssMembers.length; i++) {
            if (ActiveTssMembers[i].status == MemberStatus.unJail) {
                addresses.push(ActiveTssMembers[i].nodeAddress);
            }
        }
        return addresses;
    }

    function getTssGroupMembers() public returns (TssMember[] memory) {
        return ActiveTssMembers;
    }

    function getTssMember(bytes _address) public returns (TssMember memory) {
        return TssMemberInfo[_address];
    }

    function memberExistActive(bytes _address) public returns (bool) {
        if (TssMemberInfo[_address]) {
            return true;
        }
        return false;
    }

    function memberExistInActive(bytes _address) public returns (bool) {
        for (uint i =0; i < InActiveTssMember.length; i++) {
            if (publicKeyToAddress(InActiveTssMember[i]) == _address) {
                return true;
            }
        }
        return false;
    }

    function verifySign(bytes32 memory _message, bytes memory _sig)  public returns (bool, address) {
        bytes32 ethSignedMessageHash = getEthSignedMessageHash(_message);
        return (recover(ethSignedMessageHash, _sig) == publicKeyToAddress(confirmGroupPublicKey), recover(ethSignedMessageHash, _sig));
    }

    function publicKeyToAddress (bytes memory publicKey) public returns (address) {
        require (publicKey.length == 64);
        return address (uint160 (uint256 (keccak256 (publicKey))));
    }

    function updateTssMember() public {
        for(uint i = 1; i < InActiveTssMember.length; i++) {
            TssMember memory tm = TssMember({
                publicKey: InActiveTssMember[i],
                nodeAddress: publicKeyToAddress(InActiveTssMember[i]),
                status: MemberStatus.unJail
            });
            ActiveTssMembers[i] = tm;
            emit TssMemberAppended(
                gRoundId,
                InActiveTssMember[i],
                publicKeyToAddress(InActiveTssMember[i]),
                MemberStatus.unJail
            );
        }
        confirmNumber = 0;
        confirmGroupPublicKey = memberGroupKey[ActiveTssMember[0]];
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

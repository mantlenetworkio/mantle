// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "./ITssGroupManager.sol";

contract TssGroupManager is
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    ITssGroupManager
{
    using SafeMathUpgradeable for uint256;
    using ECDSAUpgradeable for bytes32;
    using AddressUpgradeable for address;
    bytes confirmGroupPublicKey;
    address confirmGroupAddress;
    uint256 threshold;
    uint256 gRoundId;
    uint256 confirmNumber;
    address[] addresses;
    address public stakingSlash;

    bytes[] activeTssMembers; // active tss member group
    bytes[] inActiveTssMembers; // inactive tss member group
    mapping(bytes => TssMember) public tssActiveMemberInfo; // Tss member publicKey => tssMember
    mapping(bytes => bytes) private memberGroupKey; // user publicKey => Cpk
    mapping(bytes => uint256) private groupKeyCounter; // Cpk counter
    mapping(bytes => bool) private isSubmitGroupKey; // submit group key or not
    mapping(bytes => bool) public isInActiveMember; // tss member exist or not

    event tssGroupMemberAppend(uint256 _roundId, uint256 _threshold, bytes[] _inActiveTssMembers);

    event tssActiveMemberAppended(uint256 _roundId, bytes _groupKey, bytes[] activeTssMembers);

    function initialize() public initializer {
        __Ownable_init();
        gRoundId = 0;
        confirmNumber = 0;
        threshold = 0;
    }

    modifier onlyStakingSlash() {
        require(msg.sender == stakingSlash, "contract call is not staking slashing");
        _;
    }

    function setStakingSlash(address _address) public onlyOwner {
        stakingSlash = _address;
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function setTssGroupMember(uint256 _threshold, bytes[] memory _batchPublicKey)
        public
        override
        onlyOwner
    {
        require((_batchPublicKey.length > 0), "batch public key is empty");
        require(_threshold < _batchPublicKey.length, "threshold must less than tss member");
        // require((inActiveTssMembers.length == 0), "inactive tss member array is not empty");
        if(inActiveTssMembers.length > 0) {
            for (uint256 i = 0; i < inActiveTssMembers.length; i++) {
                // re-election clear data
                delete groupKeyCounter[memberGroupKey[inActiveTssMembers[i]]];
                delete memberGroupKey[inActiveTssMembers[i]];
                delete isSubmitGroupKey[inActiveTssMembers[i]];
                delete isInActiveMember[inActiveTssMembers[i]];
            }
            delete inActiveTssMembers;
        }
        for (uint256 i = 0; i < _batchPublicKey.length; i++) {
            inActiveTssMembers.push(_batchPublicKey[i]);
            isInActiveMember[_batchPublicKey[i]] = true;
            isSubmitGroupKey[_batchPublicKey[i]] = false;
        }
        threshold = _threshold;
        gRoundId = gRoundId + 1;
        confirmNumber = 0;
        emit tssGroupMemberAppend(gRoundId, _threshold, _batchPublicKey);
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function setGroupPublicKey(bytes memory _publicKey, bytes memory _groupPublicKey)
        public
        override
    {
        require(isInActiveMember[_publicKey] == true, "your public key is not in InActiveMember");
        require(msg.sender == publicKeyToAddress(_publicKey), "public key not match");

        if (isSubmitGroupKey[_publicKey] == false) {
            isSubmitGroupKey[_publicKey] = true;
            confirmNumber = confirmNumber + 1;
        }
        if (!isEqual(memberGroupKey[_publicKey], _groupPublicKey)) {
            groupKeyCounter[_groupPublicKey] += 1;
            if (memberGroupKey[_publicKey].length != 0) {
                groupKeyCounter[memberGroupKey[_publicKey]] -= 1;
            }
            memberGroupKey[_publicKey] = _groupPublicKey;
        }
        if (groupKeyCounter[_groupPublicKey] == inActiveTssMembers.length) {
            updateTssMember(_groupPublicKey);
        }
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function getTssGroupInfo()
        public
        view
        override
        returns (
            uint256,
            uint256,
            bytes memory,
            bytes[] memory
        )
    {
        if (inActiveTssMembers.length > 0) {
            return (gRoundId - 1, threshold, confirmGroupPublicKey, activeTssMembers);
        }
        return (gRoundId, threshold, confirmGroupPublicKey, activeTssMembers);
    }

    /**
    * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function getTssInactiveGroupInfo() public view override returns (uint256, uint256,  bytes[] memory){
        return (gRoundId, threshold, inActiveTssMembers);
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function memberJail(bytes memory _publicKey) public override onlyStakingSlash {
        tssActiveMemberInfo[_publicKey].status = MemberStatus.jail;
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function memberUnJail(bytes memory _publicKey) public override onlyStakingSlash {
        tssActiveMemberInfo[_publicKey].status = MemberStatus.unJail;
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function removeMember(bytes memory _publicKey) public override onlyStakingSlash {
        for (uint256 i = 0; i < activeTssMembers.length; i++) {
            if (isEqual(activeTssMembers[i], _publicKey)) {
                removeActiveTssMembers(i);
            }
        }
        delete tssActiveMemberInfo[_publicKey];
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function getTssGroupUnJailMembers() public override returns (address[] memory) {
        delete addresses;
        for (uint256 i = 0; i < activeTssMembers.length; i++) {
            if (tssActiveMemberInfo[activeTssMembers[i]].status == MemberStatus.unJail) {
                addresses.push(tssActiveMemberInfo[activeTssMembers[i]].nodeAddress);
            }
        }
        return addresses;
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function getTssGroupMembers() public view override returns (bytes[] memory) {
        return activeTssMembers;
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function getTssMember(bytes memory _publicKey) public view override returns (TssMember memory) {
        return tssActiveMemberInfo[_publicKey];
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function memberExistActive(bytes memory _publicKey) public view override returns (bool) {
        if (tssActiveMemberInfo[_publicKey].publicKey.length > 0) {
            return true;
        } else {
            return false;
        }
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function memberExistInActive(bytes memory _publicKey) public view override returns (bool) {
        return isInActiveMember[_publicKey];
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function inActiveIsEmpty() public view override returns (bool) {
        return inActiveTssMembers.length == 0;
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function verifySign(bytes32 _message, bytes memory _sig) public view override returns (bool) {
        return (recover(_message, _sig) == confirmGroupAddress);
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function publicKeyToAddress(bytes memory publicKey) public pure override returns (address) {
        require(publicKey.length == 64, "public key length must 64 bytes");
        return address(uint160(uint256(keccak256(publicKey))));
    }

    function updateTssMember(bytes memory _groupPublicKey) private {
        if (activeTssMembers.length > 0) {
            for (uint256 i = 0; i < activeTssMembers.length; i++) {
                delete tssActiveMemberInfo[activeTssMembers[i]];    // delete tss active member map
            }
            delete activeTssMembers;  // delete active members
        }
        for (uint256 i = 0; i < inActiveTssMembers.length; i++) {
            activeTssMembers.push(inActiveTssMembers[i]);
            tssActiveMemberInfo[inActiveTssMembers[i]] = TssMember({
                publicKey: inActiveTssMembers[i],
                nodeAddress: publicKeyToAddress(inActiveTssMembers[i]),
                status: MemberStatus.unJail
            });
            // election finish clear InActiveMember data
            delete groupKeyCounter[memberGroupKey[inActiveTssMembers[i]]];
            delete memberGroupKey[inActiveTssMembers[i]];
            delete isSubmitGroupKey[inActiveTssMembers[i]];
            delete isInActiveMember[inActiveTssMembers[i]];
        }
        delete inActiveTssMembers;
        confirmGroupPublicKey = _groupPublicKey;
        confirmGroupAddress = publicKeyToAddress(_groupPublicKey);
        emit tssActiveMemberAppended(gRoundId, _groupPublicKey, activeTssMembers);
    }

    function recover(bytes32 _ethSignedMessageHash, bytes memory _sig)
        public
        pure
        returns (address)
    {
        (bytes32 r, bytes32 s, uint8 v) = _split(_sig);
        return ecrecover(_ethSignedMessageHash, v, r, s);
    }

    function _split(bytes memory _sig)
        internal
        pure
        returns (
            bytes32 r,
            bytes32 s,
            uint8 v
        )
    {
        require(_sig.length == 65, "invalid signature length");
        assembly {
            r := mload(add(_sig, 32))
            s := mload(add(_sig, 64))
            v := byte(0, mload(add(_sig, 96)))
        }
    }

    function isEqual(bytes memory byteListA, bytes memory byteListB) public pure returns (bool) {
        if (byteListA.length != byteListB.length) return false;
        for (uint256 i = 0; i < byteListA.length; i++) {
            if (byteListA[i] != byteListB[i]) return false;
        }
        return true;
    }

    function removeActiveTssMembers(uint256 _index) private {
        require(_index < activeTssMembers.length, "index out of bound");
        for (uint256 i = _index; i < activeTssMembers.length - 1; i++) {
            activeTssMembers[i] = activeTssMembers[i + 1];
        }
        activeTssMembers.pop();
    }
}

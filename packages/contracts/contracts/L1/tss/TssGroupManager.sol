// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {Lib_Address} from "../../libraries/utils/Lib_Address.sol";
import "./ITssGroupManager.sol";
import "./ITssStakingSlashing.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract TssGroupManager is
    Initializable,
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
    uint256 tempThreshold;
    address public stakingSlash;

    bytes[] activeTssMembers; // active tss member group
    bytes[] inActiveTssMembers; // inactive tss member group
    mapping(bytes => TssMember) public tssActiveMemberInfo; // Tss member publicKey => tssMember
    mapping(bytes => bytes) private _memberGroupKey; // user publicKey => Cpk
    mapping(bytes => uint256) private _groupKeyCounter; // Cpk counter
    mapping(bytes => bool) private _isSubmitGroupKey; // submit group key or not
    mapping(bytes => bool) public isInActiveMember; // tss member exist or not

    event tssGroupMemberAppend(uint256 _roundId, uint256 _threshold, bytes[] _inActiveTssMembers);

    event tssActiveMemberAppended(uint256 _roundId, bytes _groupKey, bytes[] activeTssMembers);

    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init();
        gRoundId = 0;
        threshold = 0;
        tempThreshold = 0;
    }

    modifier onlyStakingSlash() {
        require(msg.sender == stakingSlash, "contract call is not staking slashing");
        _;
    }

    function setStakingSlash(address _address) public onlyOwner {
        require(_address != address(0), "param _address is the zero address");
        stakingSlash = _address;
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function setTssGroupMember(uint256 _threshold, bytes[] calldata _batchPublicKey)
        public
        override
        onlyOwner
    {
        require((_batchPublicKey.length > 0), "batch public key is empty");
        require(_threshold < _batchPublicKey.length, "threshold must less than tss member");
        for (uint256 i = 0; i < _batchPublicKey.length; i++) {
            address operator = Lib_Address.publicKeyToAddress(_batchPublicKey[i]);
            require(IStakingSlashing(stakingSlash).isCanOperator(operator),"batch public keys has a node ,can not be operator");
        }

        if(inActiveTssMembers.length > 0) {
            for (uint256 i = 0; i < inActiveTssMembers.length; i++) {
                // re-election clear data
                delete _groupKeyCounter[_memberGroupKey[inActiveTssMembers[i]]];
                delete _memberGroupKey[inActiveTssMembers[i]];
                delete _isSubmitGroupKey[inActiveTssMembers[i]];
                delete isInActiveMember[inActiveTssMembers[i]];
            }
            delete inActiveTssMembers;
        }
        for (uint256 i = 0; i < _batchPublicKey.length; i++) {
            inActiveTssMembers.push(_batchPublicKey[i]);
            isInActiveMember[_batchPublicKey[i]] = true;
            _isSubmitGroupKey[_batchPublicKey[i]] = false;
        }
        tempThreshold = _threshold;
        emit tssGroupMemberAppend(gRoundId + 1, _threshold, _batchPublicKey);
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function setGroupPublicKey(bytes calldata _publicKey, bytes calldata _groupPublicKey)
        public
        override
    {
        require(isInActiveMember[_publicKey], "your public key is not in InActiveMember");
        require(msg.sender == Lib_Address.publicKeyToAddress(_publicKey), "public key not match");
        require(_groupPublicKey.length == 64, "Invalid groupPublicKey length");

        if (!_isSubmitGroupKey[_publicKey]) {
            _isSubmitGroupKey[_publicKey] = true;
        }
        if (!_isEqual(_memberGroupKey[_publicKey], _groupPublicKey)) {
            _groupKeyCounter[_groupPublicKey] += 1;
            if (_memberGroupKey[_publicKey].length != 0) {
                _groupKeyCounter[_memberGroupKey[_publicKey]] -= 1;
            }
            _memberGroupKey[_publicKey] = _groupPublicKey;
        }
        if (_groupKeyCounter[_groupPublicKey] >= inActiveTssMembers.length) {
            _updateTssMember(_groupPublicKey);
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
        return (gRoundId, threshold, confirmGroupPublicKey, activeTssMembers);
    }

    /**
    * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function getTssInactiveGroupInfo() public view override returns (uint256, uint256,  bytes[] memory){
        return (gRoundId + 1, tempThreshold, inActiveTssMembers);
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function memberJail(bytes calldata _publicKey) public override onlyStakingSlash {
        tssActiveMemberInfo[_publicKey].status = MemberStatus.jail;
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function memberUnJail(bytes calldata _publicKey) public override onlyStakingSlash {
        tssActiveMemberInfo[_publicKey].status = MemberStatus.unJail;
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function removeMember(bytes calldata _publicKey) public override onlyOwner {
        require(
            activeTssMembers.length > threshold + 1,
            "TssGroupManager removeMember: active members must more than threshold plus one"
        );
        for (uint256 i = 0; i < activeTssMembers.length; i++) {
            if (_isEqual(activeTssMembers[i], _publicKey)) {
                _removeActiveTssMembers(i);
                break;
            }
        }
        delete tssActiveMemberInfo[_publicKey];
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function getTssGroupUnJailMembers() public view override returns (address[] memory) {
        uint256 expectedLength;
        for (uint256 i = 0; i < activeTssMembers.length; i++) {
            if (tssActiveMemberInfo[activeTssMembers[i]].status == MemberStatus.unJail) {
                expectedLength++;
            }
        }
        address[] memory _addresses = new address[](expectedLength);
        uint256 index;
        for (uint256 i = 0; i < activeTssMembers.length; i++) {
            if (tssActiveMemberInfo[activeTssMembers[i]].status == MemberStatus.unJail) {
                _addresses[index] = tssActiveMemberInfo[activeTssMembers[i]].nodeAddress;
                index++;
            }
        }
        return _addresses;
    }

    function isTssGroupUnJailMembers(address _addr) public view override returns (bool) {
        for (uint256 i = 0; i < activeTssMembers.length; i++) {
            if (tssActiveMemberInfo[activeTssMembers[i]].status == MemberStatus.unJail) {
                if ( _addr == tssActiveMemberInfo[activeTssMembers[i]].nodeAddress) {
                    return true;
                }
            }
        }
        return false;
    }

    function memberExistActive(address _addr) public view override returns (bool) {
        for (uint256 i = 0; i < activeTssMembers.length; i++) {
            if ( _addr == tssActiveMemberInfo[activeTssMembers[i]].nodeAddress) {
                return true;
            }
        }
        return false;
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
    function getTssMember(bytes calldata _publicKey) public view override returns (TssMember memory) {
        return tssActiveMemberInfo[_publicKey];
    }

    /**
     * @inheritdoc ITssGroupManager
     */
    // slither-disable-next-line external-function
    function memberExistActive(bytes calldata _publicKey) public view override returns (bool) {
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
    function memberExistInActive(bytes calldata _publicKey) public view override returns (bool) {
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
    function verifySign(bytes32 _message, bytes calldata _sig) public view override returns (bool) {
        return (recover(_message, _sig) == confirmGroupAddress);
    }

    function _updateTssMember(bytes calldata _groupPublicKey) private {
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
                nodeAddress: Lib_Address.publicKeyToAddress(inActiveTssMembers[i]),
                status: MemberStatus.unJail
            });
            // election finish clear InActiveMember data
            delete _groupKeyCounter[_memberGroupKey[inActiveTssMembers[i]]];
            delete _memberGroupKey[inActiveTssMembers[i]];
            delete _isSubmitGroupKey[inActiveTssMembers[i]];
            delete isInActiveMember[inActiveTssMembers[i]];
        }
        delete inActiveTssMembers;
        confirmGroupPublicKey = _groupPublicKey;
        confirmGroupAddress = Lib_Address.publicKeyToAddress(_groupPublicKey);
        threshold = tempThreshold;
        gRoundId = gRoundId + 1;
        emit tssActiveMemberAppended(gRoundId, _groupPublicKey, activeTssMembers);
    }

    function recover(bytes32 _ethSignedMessageHash, bytes calldata _sig)
        public
        pure
        returns (address)
    {
        (bytes32 r, bytes32 s, uint8 v) = _split(_sig);
        address signer = ecrecover(_ethSignedMessageHash, v, r, s);
        require(signer != address(0), "ecrecover failed");
        return signer;
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
        if (v < 27) v += 27;
    }

    function _isEqual(bytes memory byteListA, bytes memory byteListB) private pure returns (bool) {
        if (byteListA.length != byteListB.length) return false;
        for (uint256 i = 0; i < byteListA.length; i++) {
            if (byteListA[i] != byteListB[i]) return false;
        }
        return true;
    }

    function _removeActiveTssMembers(uint256 _index) private {
        activeTssMembers[_index] = activeTssMembers[activeTssMembers.length - 1];
        activeTssMembers.pop();
    }
}

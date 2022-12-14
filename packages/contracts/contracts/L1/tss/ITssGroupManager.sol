// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

interface ITssGroupManager {
    enum MemberStatus {
        unJail,
        jail
    }

    struct TssMember {
        bytes         publicKey;
        address       nodeAddress;
        MemberStatus  status;
    }

    function setTssGroupMember(uint256 _threshold, bytes[] memory _batchPublicKey) external;
    function setGroupPublicKey(bytes memory _publicKey, bytes memory _groupPublicKey) external;
    function getTssGroupInfo() external returns (uint256, uint256, bytes memory, bytes[] memory);
    function getTssInactiveGroupInfo() external returns (uint256, uint256, bytes[] memory);
    function memberJail(bytes memory _publicKey) external;
    function memberUnJail(bytes memory _publicKey) external;
    function removeMember(bytes memory _publicKey) external;
    function getTssGroupUnJailMembers() external returns (address[] memory);
    function getTssGroupMembers() external returns (bytes[] memory);
    function getTssMember(bytes memory _publicKey) external returns (TssMember memory);
    function memberExistActive(bytes memory _publicKey) external returns (bool);
    function memberExistInActive(bytes memory _publicKey) external returns (bool);
    function inActiveIsEmpty() external returns (bool);
    function verifySign(bytes32 _message, bytes memory _sig) external returns (bool);
    function publicKeyToAddress (bytes memory publicKey) external returns (address);
}

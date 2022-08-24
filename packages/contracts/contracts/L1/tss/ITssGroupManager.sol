pragma solidity ^0.8.0;

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

    function setTssGroupMember(uint256 _threshold, bytes[] memory _batchPublicKey) public onlyOwner;
    function setGroupPublicKey(bytes memory _publicKey, bytes memory _groupPublicKey) public;
    function getTssGroupInfo() public returns (uint256, uint256, bytes, TssMember[]);
    function memberJail(bytes memory _publicKey) public;
    function memberUnJail(bytes memory _publicKey) public;
    function removeMember(bytes memory _publicKey) public;
    function getTssGroupUnJailMembers() public returns (address[] memory);
    function getTssGroupMembers() public returns (TssMember[] memory);
    function getTssMember(bytes _address) public returns (TssMember memory);
    function memberExistActive(bytes _address) public returns (bool);
    function memberExistInActive(bytes _address) public returns (bool);
    function verifySign(bytes32 memory _message, bytes memory _sig)  public returns (bool, address);
}

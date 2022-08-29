# TssGroupManager









## Methods

### getEthSignedMessageHash

```solidity
function getEthSignedMessageHash(bytes32 _messageHash) external pure returns (bytes32)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _messageHash | bytes32 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bytes32 | undefined

### getTssGroupInfo

```solidity
function getTssGroupInfo() external view returns (uint256, uint256, bytes, bytes[])
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined
| _1 | uint256 | undefined
| _2 | bytes | undefined
| _3 | bytes[] | undefined

### getTssGroupMembers

```solidity
function getTssGroupMembers() external view returns (bytes[])
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bytes[] | undefined

### getTssGroupUnJailMembers

```solidity
function getTssGroupUnJailMembers() external nonpayable returns (address[])
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address[] | undefined

### getTssMember

```solidity
function getTssMember(bytes _publicKey) external view returns (struct ITssGroupManager.TssMember)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _publicKey | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | ITssGroupManager.TssMember | undefined

### inActiveIsEmpty

```solidity
function inActiveIsEmpty() external view returns (bool)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### initialize

```solidity
function initialize(address _libAddressManager) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _libAddressManager | address | undefined

### isEqual

```solidity
function isEqual(bytes byteListA, bytes byteListB) external pure returns (bool)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| byteListA | bytes | undefined
| byteListB | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### isInActiveMember

```solidity
function isInActiveMember(bytes) external view returns (bool)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### isSubmitGroupKey

```solidity
function isSubmitGroupKey(bytes) external view returns (bool)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### libAddressManager

```solidity
function libAddressManager() external view returns (contract Lib_AddressManager)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | contract Lib_AddressManager | undefined

### memberExistActive

```solidity
function memberExistActive(bytes _publicKey) external view returns (bool)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _publicKey | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### memberExistInActive

```solidity
function memberExistInActive(bytes _publicKey) external view returns (bool)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _publicKey | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### memberGroupKey

```solidity
function memberGroupKey(bytes) external view returns (bytes)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bytes | undefined

### memberJail

```solidity
function memberJail(bytes _publicKey) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _publicKey | bytes | undefined

### memberUnJail

```solidity
function memberUnJail(bytes _publicKey) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _publicKey | bytes | undefined

### owner

```solidity
function owner() external view returns (address)
```



*Returns the address of the current owner.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### publicKeyToAddress

```solidity
function publicKeyToAddress(bytes publicKey) external pure returns (address)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| publicKey | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### recover

```solidity
function recover(bytes32 _ethSignedMessageHash, bytes _sig) external pure returns (address)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _ethSignedMessageHash | bytes32 | undefined
| _sig | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### removeMember

```solidity
function removeMember(bytes _publicKey) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _publicKey | bytes | undefined

### renounceOwnership

```solidity
function renounceOwnership() external nonpayable
```



*Leaves the contract without owner. It will not be possible to call `onlyOwner` functions anymore. Can only be called by the current owner. NOTE: Renouncing ownership will leave the contract without an owner, thereby removing any functionality that is only available to the owner.*


### resolve

```solidity
function resolve(string _name) external view returns (address)
```

Resolves the address associated with a given name.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _name | string | Name to resolve an address for.

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | Address associated with the given name.

### setGroupPublicKey

```solidity
function setGroupPublicKey(bytes _publicKey, bytes _groupPublicKey) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _publicKey | bytes | undefined
| _groupPublicKey | bytes | undefined

### setStakingSlash

```solidity
function setStakingSlash(address _address) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _address | address | undefined

### setTssGroupMember

```solidity
function setTssGroupMember(uint256 _threshold, bytes[] _batchPublicKey) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _threshold | uint256 | undefined
| _batchPublicKey | bytes[] | undefined

### transferOwnership

```solidity
function transferOwnership(address newOwner) external nonpayable
```



*Transfers ownership of the contract to a new account (`newOwner`). Can only be called by the current owner.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| newOwner | address | undefined

### tssActiveMemberInfo

```solidity
function tssActiveMemberInfo(bytes) external view returns (bytes publicKey, address nodeAddress, enum ITssGroupManager.MemberStatus status)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| publicKey | bytes | undefined
| nodeAddress | address | undefined
| status | enum ITssGroupManager.MemberStatus | undefined

### updateTssMember

```solidity
function updateTssMember(bytes _groupPublicKey) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _groupPublicKey | bytes | undefined

### verifySign

```solidity
function verifySign(bytes32 _message, bytes _sig) external view returns (bool)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _message | bytes32 | undefined
| _sig | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined



## Events

### Initialized

```solidity
event Initialized(uint8 version)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| version  | uint8 | undefined |

### OwnershipTransferred

```solidity
event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| previousOwner `indexed` | address | undefined |
| newOwner `indexed` | address | undefined |

### tssActiveMemberAppended

```solidity
event tssActiveMemberAppended(uint256 _roundId, bytes _groupKey, bytes[] activeTssMembers)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _roundId  | uint256 | undefined |
| _groupKey  | bytes | undefined |
| activeTssMembers  | bytes[] | undefined |

### tssGroupMemberAppend

```solidity
event tssGroupMemberAppend(uint256 _roundId, uint256 _threshold, bytes[] _inActiveTssMembers)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _roundId  | uint256 | undefined |
| _threshold  | uint256 | undefined |
| _inActiveTssMembers  | bytes[] | undefined |




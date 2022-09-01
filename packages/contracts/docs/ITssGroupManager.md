# ITssGroupManager









## Methods

### getTssGroupInfo

```solidity
function getTssGroupInfo() external nonpayable returns (uint256, uint256, bytes, bytes[])
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
function getTssGroupMembers() external nonpayable returns (bytes[])
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

### getTssInactiveGroupInfo

```solidity
function getTssInactiveGroupInfo() external nonpayable returns (uint256, uint256, bytes[])
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined
| _1 | uint256 | undefined
| _2 | bytes[] | undefined

### getTssMember

```solidity
function getTssMember(bytes _publicKey) external nonpayable returns (struct ITssGroupManager.TssMember)
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
function inActiveIsEmpty() external nonpayable returns (bool)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### memberExistActive

```solidity
function memberExistActive(bytes _publicKey) external nonpayable returns (bool)
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
function memberExistInActive(bytes _publicKey) external nonpayable returns (bool)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _publicKey | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

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

### publicKeyToAddress

```solidity
function publicKeyToAddress(bytes publicKey) external nonpayable returns (address)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| publicKey | bytes | undefined

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

### setGroupPublicKey

```solidity
function setGroupPublicKey(bytes _publicKey, bytes _groupPublicKey) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _publicKey | bytes | undefined
| _groupPublicKey | bytes | undefined

### setTssGroupMember

```solidity
function setTssGroupMember(uint256 _threshold, bytes[] _batchPublicKey) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _threshold | uint256 | undefined
| _batchPublicKey | bytes[] | undefined

### verifySign

```solidity
function verifySign(bytes32 _message, bytes _sig) external nonpayable returns (bool)
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





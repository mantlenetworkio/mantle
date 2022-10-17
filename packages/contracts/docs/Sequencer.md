# Sequencer









## Methods

### bitToken

```solidity
function bitToken() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### changeBitAddress

```solidity
function changeBitAddress(address _bitToken) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _bitToken | address | undefined

### createSequencer

```solidity
function createSequencer(uint256 _amount, address _mintAddress, bytes _nodeID) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _amount | uint256 | undefined
| _mintAddress | address | undefined
| _nodeID | bytes | undefined

### deposit

```solidity
function deposit(uint256 _amount) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _amount | uint256 | undefined

### getOwners

```solidity
function getOwners() external view returns (address[])
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address[] | undefined

### getSequencer

```solidity
function getSequencer(address signer) external view returns (struct Sequencer.SequencerInfo)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| signer | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | Sequencer.SequencerInfo | undefined

### getSequencers

```solidity
function getSequencers() external view returns (struct Sequencer.SequencerInfo[])
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | Sequencer.SequencerInfo[] | undefined

### initialize

```solidity
function initialize(address _bitToken) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _bitToken | address | undefined

### isSequencer

```solidity
function isSequencer(address signer) external view returns (bool)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| signer | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### owner

```solidity
function owner() external view returns (address)
```



*Returns the address of the current owner.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### owners

```solidity
function owners(uint256) external view returns (address)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### rel

```solidity
function rel(address) external view returns (address)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### renounceOwnership

```solidity
function renounceOwnership() external nonpayable
```



*Leaves the contract without owner. It will not be possible to call `onlyOwner` functions anymore. Can only be called by the current owner. NOTE: Renouncing ownership will leave the contract without an owner, thereby removing any functionality that is only available to the owner.*


### sequencers

```solidity
function sequencers(address) external view returns (address owner, address mintAddress, bytes nodeID, uint256 amount, uint256 keyIndex)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| owner | address | undefined
| mintAddress | address | undefined
| nodeID | bytes | undefined
| amount | uint256 | undefined
| keyIndex | uint256 | undefined

### transferOwnership

```solidity
function transferOwnership(address newOwner) external nonpayable
```



*Transfers ownership of the contract to a new account (`newOwner`). Can only be called by the current owner.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| newOwner | address | undefined

### withdraw

```solidity
function withdraw(uint256 _amount) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _amount | uint256 | undefined

### withdrawAll

```solidity
function withdrawAll() external nonpayable
```








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

### SequencerCreate

```solidity
event SequencerCreate(address, address, bytes)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0  | address | undefined |
| _1  | address | undefined |
| _2  | bytes | undefined |

### SequencerDelete

```solidity
event SequencerDelete(address, bytes)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0  | address | undefined |
| _1  | bytes | undefined |

### SequencerUpdate

```solidity
event SequencerUpdate(address, bytes, uint256)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0  | address | undefined |
| _1  | bytes | undefined |
| _2  | uint256 | undefined |




# TssRewardContract



> TssRewardContract



*Release to batch roll up tss members.*

## Methods

### bestBlockID

```solidity
function bestBlockID() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### bvmGasPriceOracleAddress

```solidity
function bvmGasPriceOracleAddress() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### claimReward

```solidity
function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] _tssMembers) external nonpayable
```



*claimReward distribute reward to tss member.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _blockStartHeight | uint256 | undefined
| _length | uint32 | The distribute batch block number
| _batchTime | uint256 | Batch corresponds to L1 Block Timestamp
| _tssMembers | address[] | The address array of tss group members

### claimRewardPub

```solidity
function claimRewardPub(address[] _tssMembers) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _tssMembers | address[] | undefined

### deadAddress

```solidity
function deadAddress() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### dust

```solidity
function dust() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### lastBatchAmount

```solidity
function lastBatchAmount() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### lastSendAmount

```solidity
function lastSendAmount() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### lastSub

```solidity
function lastSub() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### latsBatchTime

```solidity
function latsBatchTime() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### ledger

```solidity
function ledger(uint256) external view returns (uint256)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### owner

```solidity
function owner() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### queryOwner

```solidity
function queryOwner() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### queryReward

```solidity
function queryReward() external view returns (uint256)
```



*return the total undistributed amount*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### querySendAmountPerSecond

```solidity
function querySendAmountPerSecond() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### sendAmountPerYear

```solidity
function sendAmountPerYear() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### step

```solidity
function step() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### totalAmount

```solidity
function totalAmount() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### updateReward

```solidity
function updateReward(uint256 _blockID, uint256 _amount) external nonpayable returns (bool)
```



*update tss member gas reward by every block.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _blockID | uint256 | The block height at L2 which needs to distribute profits
| _amount | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | _tssMembers Address array of tss group members

### withdraw

```solidity
function withdraw() external nonpayable
```



*clear balance*


### withdrawDust

```solidity
function withdrawDust() external nonpayable
```



*withdraw div dust*




## Events

### DistributeTssReward

```solidity
event DistributeTssReward(uint256 batchTime, address[] tssMembers)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| batchTime  | uint256 | undefined |
| tssMembers  | address[] | undefined |

### DistributeTssRewardByBlock

```solidity
event DistributeTssRewardByBlock(uint256 blockStartHeight, uint32 length, address[] tssMembers)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| blockStartHeight  | uint256 | undefined |
| length  | uint32 | undefined |
| tssMembers  | address[] | undefined |




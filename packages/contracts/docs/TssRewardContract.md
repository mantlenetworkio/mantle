# TssRewardContract



> TssRewardContract



*Collect L2 block gas reward per block and release to batch roll up tss members.*

## Methods

### bestBlockID

```solidity
function bestBlockID() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### claimReward

```solidity
function claimReward(uint256 _blockStartHeight, uint32 _length, address[] _tssMembers) external nonpayable
```



*claimReward distribute reward to tss member.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _blockStartHeight | uint256 | The block height at L2 which needs to distribute profits
| _length | uint32 | The distribute batch block number
| _tssMembers | address[] | The address array of tss group members

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
function owner() external view returns (address payable)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address payable | undefined

### queryReward

```solidity
function queryReward() external view returns (uint256)
```



*return the total undistributed amount*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### updateReward

```solidity
function updateReward(uint256 _blockID, uint256 _amount) external payable returns (bool)
```



*update tss member gas reward by every block.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _blockID | uint256 | The block height at L2 which needs to distribute profits
| _amount | uint256 | Distribute batch block number

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | _tssMembers Address array of tss group members

### withdraw

```solidity
function withdraw() external nonpayable
```



*clear contract(canonical).*


### withdrawDust

```solidity
function withdrawDust() external nonpayable
```



*withdraw dust.*




## Events

### DistributeTssReward

```solidity
event DistributeTssReward(uint256 blockStartHeight, uint256 length, address[] tssMembers)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| blockStartHeight  | uint256 | undefined |
| length  | uint256 | undefined |
| tssMembers  | address[] | undefined |




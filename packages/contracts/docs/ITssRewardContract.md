# ITssRewardContract



> ITssRewardContract





## Methods

### claimReward

```solidity
function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] _tssMembers) external nonpayable
```



*Auto distribute reward to tss members.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _blockStartHeight | uint256 | L2 rollup batch block start height.
| _length | uint32 | Rollup batch length.
| _batchTime | uint256 | undefined
| _tssMembers | address[] | Tss member address array.

### queryReward

```solidity
function queryReward() external view returns (uint256)
```



*Query total undistributed balance.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | Amount of undistributed rewards.

### updateReward

```solidity
function updateReward(uint256 _blockID, uint256 _amount) external nonpayable returns (bool)
```



*Update deposit block gas into contract.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _blockID | uint256 | Update gas reward L2 block ID.
| _amount | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | Update success.

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
event DistributeTssReward(uint256 lastBatchTime, uint256 batchTime, uint256 amount, address[] tssMembers)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| lastBatchTime  | uint256 | undefined |
| batchTime  | uint256 | undefined |
| amount  | uint256 | undefined |
| tssMembers  | address[] | undefined |

### DistributeTssRewardByBlock

```solidity
event DistributeTssRewardByBlock(uint256 blockStartHeight, uint32 length, uint256 amount, address[] tssMembers)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| blockStartHeight  | uint256 | undefined |
| length  | uint32 | undefined |
| amount  | uint256 | undefined |
| tssMembers  | address[] | undefined |




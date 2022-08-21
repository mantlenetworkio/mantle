# ITssRewardContract



> ITssRewardContract





## Methods

### claimReward

```solidity
function claimReward(uint256 _blockStartHeight, uint32 _length, address[] _tssMembers) external nonpayable
```



*Auto distribute reward to tss members.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _blockStartHeight | uint256 | L2 rollup batch block start height.
| _length | uint32 | Rollup batch length.
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
function updateReward(uint256 _blockID, uint256 _amount) external payable returns (bool)
```



*Update deposit block gas into contract.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _blockID | uint256 | Update gas reward L2 block ID.
| _amount | uint256 | Update gas reward amount.

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
event DistributeTssReward(uint256 blockStartHeight, uint256 length, address[] tssMembers)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| blockStartHeight  | uint256 | undefined |
| length  | uint256 | undefined |
| tssMembers  | address[] | undefined |




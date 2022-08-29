# IStakingSlashing









## Methods

### clearQuitList

```solidity
function clearQuitList() external nonpayable
```






### getDeposits

```solidity
function getDeposits(address) external nonpayable returns (struct IStakingSlashing.DepositInfo)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | IStakingSlashing.DepositInfo | undefined

### getSlashRecord

```solidity
function getSlashRecord(uint256, address) external nonpayable returns (bool)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined
| _1 | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### quit

```solidity
function quit() external nonpayable
```






### slashing

```solidity
function slashing(bytes, bytes) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | bytes | undefined
| _1 | bytes | undefined

### staking

```solidity
function staking(uint256, bytes) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined
| _1 | bytes | undefined

### withdrawToken

```solidity
function withdrawToken() external nonpayable
```










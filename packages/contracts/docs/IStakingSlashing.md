# IStakingSlashing









## Methods

### batchGetDeposits

```solidity
function batchGetDeposits(address[]) external view returns (struct IStakingSlashing.DepositInfo[])
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | address[] | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | IStakingSlashing.DepositInfo[] | undefined

### clearQuitRequestList

```solidity
function clearQuitRequestList() external nonpayable
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

### getQuitRequestList

```solidity
function getQuitRequestList() external view returns (address[])
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address[] | undefined

### getSlashRecord

```solidity
function getSlashRecord(uint256, address) external view returns (bool)
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

### getSlashingParams

```solidity
function getSlashingParams() external view returns (uint256[2], uint256[2])
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256[2] | undefined
| _1 | uint256[2] | undefined

### isJailed

```solidity
function isJailed(address) external nonpayable returns (bool)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### quitRequest

```solidity
function quitRequest() external nonpayable
```






### setAddress

```solidity
function setAddress(address, address) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined
| _1 | address | undefined

### setSlashingParams

```solidity
function setSlashingParams(uint256[2], uint256[2]) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | uint256[2] | undefined
| _1 | uint256[2] | undefined

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

### unJail

```solidity
function unJail() external nonpayable
```






### withdrawToken

```solidity
function withdrawToken() external nonpayable
```










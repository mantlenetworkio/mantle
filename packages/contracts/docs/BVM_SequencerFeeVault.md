# BVM_SequencerFeeVault



> BVM_SequencerFeeVault



*Simple holding contract for fees paid to the Sequencer. Likely to be replaced in the future but &quot;good enough for now&quot;.*

## Methods

### L1Gas

```solidity
function L1Gas() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### burner

```solidity
function burner() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### bvmGasPriceOracleAddress

```solidity
function bvmGasPriceOracleAddress() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### l1FeeWallet

```solidity
function l1FeeWallet() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### minWithdrawalAmount

```solidity
function minWithdrawalAmount() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### owner

```solidity
function owner() external view returns (address)
```



*Returns the address of the current owner.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### renounceOwnership

```solidity
function renounceOwnership() external nonpayable
```



*Leaves the contract without owner. It will not be possible to call `onlyOwner` functions anymore. Can only be called by the current owner. NOTE: Renouncing ownership will leave the contract without an owner, thereby removing any functionality that is only available to the owner.*


### setBurner

```solidity
function setBurner(address _burner) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _burner | address | undefined

### setL1FeeWallet

```solidity
function setL1FeeWallet(address _l1FeeWallet) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _l1FeeWallet | address | undefined

### setMinWithdrawalAmount

```solidity
function setMinWithdrawalAmount(uint256 _minWithdrawalAmount) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _minWithdrawalAmount | uint256 | undefined

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
function withdraw() external nonpayable
```








## Events

### OwnershipTransferred

```solidity
event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| previousOwner `indexed` | address | undefined |
| newOwner `indexed` | address | undefined |




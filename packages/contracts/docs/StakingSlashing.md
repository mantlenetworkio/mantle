# StakingSlashing









## Methods

### BitToken

```solidity
function BitToken() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### clearQuitList

```solidity
function clearQuitList() external nonpayable
```

clear the quit list




### deposits

```solidity
function deposits(address) external view returns (bytes pubKey, uint256 amount)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| pubKey | bytes | undefined
| amount | uint256 | undefined

### exIncome

```solidity
function exIncome(uint256) external view returns (uint256)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### getDeposits

```solidity
function getDeposits(address user) external view returns (struct IStakingSlashing.DepositInfo)
```

get the deposit info



#### Parameters

| Name | Type | Description |
|---|---|---|
| user | address | address of the staker

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | IStakingSlashing.DepositInfo | undefined

### getQuitList

```solidity
function getQuitList() external view returns (address[])
```

return the quit list




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address[] | undefined

### getSlashRecord

```solidity
function getSlashRecord(uint256 batchIndex, address user) external view returns (bool)
```

get the slash record



#### Parameters

| Name | Type | Description |
|---|---|---|
| batchIndex | uint256 | the index of batch
| user | address | address of the staker

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### getSlashingParams

```solidity
function getSlashingParams() external view returns (uint256[2], uint256[2])
```

set the slashing params (0 -&gt; uptime , 1 -&gt; animus)




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256[2] | _slashAmount the amount to be deducted for each type
| _1 | uint256[2] | undefined

### initialize

```solidity
function initialize(address _libAddressManager) external nonpayable
```

initializes the contract setting and the deployer as the initial owner



#### Parameters

| Name | Type | Description |
|---|---|---|
| _libAddressManager | address | address manager contract address

### isEqual

```solidity
function isEqual(bytes byteListA, bytes byteListB) external pure returns (bool)
```

check two bytes for equality



#### Parameters

| Name | Type | Description |
|---|---|---|
| byteListA | bytes | undefined
| byteListB | bytes | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### isJailed

```solidity
function isJailed(address user) external nonpayable returns (bool)
```

check the tssnode status



#### Parameters

| Name | Type | Description |
|---|---|---|
| user | address | address of the staker

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

### owner

```solidity
function owner() external view returns (address)
```



*Returns the address of the current owner.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### quit

```solidity
function quit() external nonpayable
```

send quit request for the next election




### quitList

```solidity
function quitList(uint256) external view returns (address)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

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

### setAddress

```solidity
function setAddress(address _token, address _tssGroup) external nonpayable
```

change the bit token and tssGroup contract address



#### Parameters

| Name | Type | Description |
|---|---|---|
| _token | address | the erc20 bit token contract address
| _tssGroup | address | tssGroup contract address

### setSlashingParams

```solidity
function setSlashingParams(uint256[2] _slashAmount, uint256[2] _exIncome) external nonpayable
```

set the slashing params (0 -&gt; uptime , 1 -&gt; animus)



#### Parameters

| Name | Type | Description |
|---|---|---|
| _slashAmount | uint256[2] | the amount to be deducted for each type
| _exIncome | uint256[2] | additional amount available to the originator of the report

### slashAmount

```solidity
function slashAmount(uint256) external view returns (uint256)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### slashing

```solidity
function slashing(bytes _messageBytes, bytes _sig) external nonpayable
```

verify the slash message then slash



#### Parameters

| Name | Type | Description |
|---|---|---|
| _messageBytes | bytes | the message that abi encode by type SlashMsg
| _sig | bytes | the signature of the hash keccak256(_messageBytes)

### staking

```solidity
function staking(uint256 _amount, bytes _pubKey) external nonpayable
```

staking entrance for user to deposit bit tokens



#### Parameters

| Name | Type | Description |
|---|---|---|
| _amount | uint256 | deposit amount of bit token
| _pubKey | bytes | public key of sender

### transferOwnership

```solidity
function transferOwnership(address newOwner) external nonpayable
```



*Transfers ownership of the contract to a new account (`newOwner`). Can only be called by the current owner.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| newOwner | address | undefined

### tssGroupContract

```solidity
function tssGroupContract() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### unJail

```solidity
function unJail() external nonpayable
```

set tss node status unjail




### withdrawToken

```solidity
function withdrawToken() external nonpayable
```

user who not elected to be validator to withdraw their bit token






## Events

### AddDeposit

```solidity
event AddDeposit(address, IStakingSlashing.DepositInfo)
```

staking for himself



#### Parameters

| Name | Type | Description |
|---|---|---|
| _0  | address | undefined |
| _1  | IStakingSlashing.DepositInfo | undefined |

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

### Slashing

```solidity
event Slashing(address, enum StakingSlashing.SlashType)
```

slash tssnode



#### Parameters

| Name | Type | Description |
|---|---|---|
| _0  | address | undefined |
| _1  | enum StakingSlashing.SlashType | undefined |

### Withdraw

```solidity
event Withdraw(address, uint256)
```

withdraw for himself



#### Parameters

| Name | Type | Description |
|---|---|---|
| _0  | address | undefined |
| _1  | uint256 | undefined |




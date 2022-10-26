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

### createSequencer

```solidity
function createSequencer(uint256 _amount, address _mintAddress, bytes _nodeID) external nonpayable
```

Create a new sequencer info and init amount



#### Parameters

| Name | Type | Description |
|---|---|---|
| _amount | uint256 | amount of bit token, will transfer to this contract when sequencer create
| _mintAddress | address | sequencer mint address
| _nodeID | bytes | sequencer node ID

### deposit

```solidity
function deposit(uint256 _amount) external nonpayable
```

Check sequencer exist then add deposit amount



#### Parameters

| Name | Type | Description |
|---|---|---|
| _amount | uint256 | amount of bit token

### epoch

```solidity
function epoch() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### getOwners

```solidity
function getOwners() external view returns (address[])
```

Return owners




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address[] | owners all owners

### getSequencer

```solidity
function getSequencer(address signer) external view returns (struct Sequencer.SequencerInfo)
```

Return sequencer info by signer address



#### Parameters

| Name | Type | Description |
|---|---|---|
| signer | address | signer address, the key to find sequencer

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | Sequencer.SequencerInfo | seq sequencer info

### getSequencers

```solidity
function getSequencers() external view returns (struct Sequencer.SequencerInfo[])
```

Return all sequencer infos




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | Sequencer.SequencerInfo[] | seqs all sequencers

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

Return if signer exist



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


### scheduler

```solidity
function scheduler() external view returns (bytes)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bytes | undefined

### sequencerLimit

```solidity
function sequencerLimit() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

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

### updateBitAddress

```solidity
function updateBitAddress(address _bitToken) external nonpayable
```

Update bit token address



#### Parameters

| Name | Type | Description |
|---|---|---|
| _bitToken | address | new ERC20 address of bit token

### updateEpoch

```solidity
function updateEpoch(uint256 _epoch) external nonpayable
```

Update Epoch



#### Parameters

| Name | Type | Description |
|---|---|---|
| _epoch | uint256 | new epoch

### updateScheduler

```solidity
function updateScheduler(bytes nodeID) external nonpayable
```

Update Epoch



#### Parameters

| Name | Type | Description |
|---|---|---|
| nodeID | bytes | new scheculer`s nodeID

### updateSequencerLimit

```solidity
function updateSequencerLimit(uint256 _limit) external nonpayable
```

Update Epoch



#### Parameters

| Name | Type | Description |
|---|---|---|
| _limit | uint256 | new limit

### withdraw

```solidity
function withdraw(uint256 _amount) external nonpayable
```

amount &gt; deposit(signer).amount -&gt; withdraw all 0 &lt; amount &lt; deposit(signer).amount -&gt; withdraw amount to signer when deposit(signer).amount = 0, delete the sequencer



#### Parameters

| Name | Type | Description |
|---|---|---|
| _amount | uint256 | amount of bit token

### withdrawAll

```solidity
function withdrawAll() external nonpayable
```

Check sequencer exist then withdraw all. This action will delete sequencer after withdraw






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




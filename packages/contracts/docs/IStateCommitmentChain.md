# IStateCommitmentChain



> IStateCommitmentChain





## Methods

### appendStateBatch

```solidity
function appendStateBatch(bytes32[] _batch, uint256 _shouldStartAtElement, bytes _signature) external nonpayable
```

Appends a batch of state roots to the chain.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _batch | bytes32[] | Batch of state roots.
| _shouldStartAtElement | uint256 | Index of the element at which this batch should start.
| _signature | bytes | undefined

### deleteStateBatch

```solidity
function deleteStateBatch(Lib_BVMCodec.ChainBatchHeader _batchHeader) external nonpayable
```

Deletes all state roots after (and including) a given batch.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _batchHeader | Lib_BVMCodec.ChainBatchHeader | Header of the batch to start deleting from.

### getLastSequencerTimestamp

```solidity
function getLastSequencerTimestamp() external view returns (uint256 _lastSequencerTimestamp)
```

Retrieves the timestamp of the last batch submitted by the sequencer.




#### Returns

| Name | Type | Description |
|---|---|---|
| _lastSequencerTimestamp | uint256 | Last sequencer batch timestamp.

### getTotalBatches

```solidity
function getTotalBatches() external view returns (uint256 _totalBatches)
```

Retrieves the total number of batches submitted.




#### Returns

| Name | Type | Description |
|---|---|---|
| _totalBatches | uint256 | Total submitted batches.

### getTotalElements

```solidity
function getTotalElements() external view returns (uint256 _totalElements)
```

Retrieves the total number of elements submitted.




#### Returns

| Name | Type | Description |
|---|---|---|
| _totalElements | uint256 | Total submitted elements.

### insideFraudProofWindow

```solidity
function insideFraudProofWindow(Lib_BVMCodec.ChainBatchHeader _batchHeader) external view returns (bool _inside)
```

Checks whether a given batch is still inside its fraud proof window.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _batchHeader | Lib_BVMCodec.ChainBatchHeader | Header of the batch to check.

#### Returns

| Name | Type | Description |
|---|---|---|
| _inside | bool | Whether or not the batch is inside the fraud proof window.

### rollBackL2Chain

```solidity
function rollBackL2Chain(uint256 _shouldRollBack, uint256 _shouldStartAtElement, bytes _signature) external nonpayable
```

Emit event to notify sequencers to roll back.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _shouldRollBack | uint256 | roll back to should start .
| _shouldStartAtElement | uint256 | Index of the element at which this batch should start
| _signature | bytes | signature of rollback message

### rollBackMessage

```solidity
function rollBackMessage(uint256 _shouldRollBack) external nonpayable
```

interface for send domain message



#### Parameters

| Name | Type | Description |
|---|---|---|
| _shouldRollBack | uint256 | roll back to should start .

### verifyStateCommitment

```solidity
function verifyStateCommitment(bytes32 _element, Lib_BVMCodec.ChainBatchHeader _batchHeader, Lib_BVMCodec.ChainInclusionProof _proof) external view returns (bool _verified)
```

Verifies a batch inclusion proof.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _element | bytes32 | Hash of the element to verify a proof for.
| _batchHeader | Lib_BVMCodec.ChainBatchHeader | Header of the batch in which the element was included.
| _proof | Lib_BVMCodec.ChainInclusionProof | Merkle inclusion proof for the element.

#### Returns

| Name | Type | Description |
|---|---|---|
| _verified | bool | undefined



## Events

### DistributeTssReward

```solidity
event DistributeTssReward(uint256 indexed _startBlockNumber, uint256 _length, uint256 indexed _batchTime, address[] _tssMembers)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _startBlockNumber `indexed` | uint256 | undefined |
| _length  | uint256 | undefined |
| _batchTime `indexed` | uint256 | undefined |
| _tssMembers  | address[] | undefined |

### RollBackL2Chain

```solidity
event RollBackL2Chain(uint256 indexed _startBlockNumber)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _startBlockNumber `indexed` | uint256 | undefined |

### StateBatchAppended

```solidity
event StateBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _signature, bytes _extraData)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _batchIndex `indexed` | uint256 | undefined |
| _batchRoot  | bytes32 | undefined |
| _batchSize  | uint256 | undefined |
| _prevTotalElements  | uint256 | undefined |
| _signature  | bytes | undefined |
| _extraData  | bytes | undefined |

### StateBatchDeleted

```solidity
event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _batchIndex `indexed` | uint256 | undefined |
| _batchRoot  | bytes32 | undefined |




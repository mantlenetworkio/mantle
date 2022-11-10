# CrossDomainMessenger





CrossDomainMessenger is a base contract that provides the core logic for the L1 and L2         cross-chain messenger contracts. It&#39;s designed to be a universal interface that only         needs to be extended slightly to provide low-level message passing functionality on each         chain it&#39;s deployed on. Currently only designed for message passing between two paired         chains and does not support one-to-many interactions.



## Methods

### MESSAGE_VERSION

```solidity
function MESSAGE_VERSION() external view returns (uint16)
```

Current message version identifier.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint16 | undefined

### MIN_GAS_CALLDATA_OVERHEAD

```solidity
function MIN_GAS_CALLDATA_OVERHEAD() external view returns (uint64)
```

Extra gas added to base gas for each byte of calldata in a message.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint64 | undefined

### MIN_GAS_CONSTANT_OVERHEAD

```solidity
function MIN_GAS_CONSTANT_OVERHEAD() external view returns (uint64)
```

Constant overhead added to the base gas for a message.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint64 | undefined

### MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR

```solidity
function MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR() external view returns (uint64)
```

Denominator for dynamic overhead added to the base gas for a message.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint64 | undefined

### MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR

```solidity
function MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR() external view returns (uint64)
```

Numerator for dynamic overhead added to the base gas for a message.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint64 | undefined

### OTHER_MESSENGER

```solidity
function OTHER_MESSENGER() external view returns (address)
```

Address of the paired CrossDomainMessenger contract on the other chain.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### baseGas

```solidity
function baseGas(bytes _message, uint32 _minGasLimit) external pure returns (uint64)
```

Computes the amount of gas required to guarantee that a given message will be         received on the other chain without running out of gas. Guaranteeing that a message         will not run out of gas is important because this ensures that a message can always         be replayed on the other chain if it fails to execute completely.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _message | bytes | Message to compute the amount of required gas for.
| _minGasLimit | uint32 | Minimum desired gas limit when message goes to target.

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint64 | Amount of gas required to guarantee message receipt.

### messageNonce

```solidity
function messageNonce() external view returns (uint256)
```

Retrieves the next message nonce. Message version will be added to the upper two         bytes of the message nonce. Message version allows us to treat messages as having         different structures.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | Nonce of the next message to be sent, with added message version.

### owner

```solidity
function owner() external view returns (address)
```



*Returns the address of the current owner.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### pause

```solidity
function pause() external nonpayable
```

Allows the owner of this contract to temporarily pause message relaying. Backup         security mechanism just in case. Owner should be the same as the upgrade wallet to         maintain the security model of the system as a whole.




### paused

```solidity
function paused() external view returns (bool)
```



*Returns true if the contract is paused, and false otherwise.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### receivedMessages

```solidity
function receivedMessages(bytes32) external view returns (bool)
```

Mapping of message hashes to boolean receipt values. Note that a message will only         be present in this mapping if it failed to be relayed on this chain at least once.         If a message is successfully relayed on the first attempt, then it will only be         present within the successfulMessages mapping.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | bytes32 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### relayMessage

```solidity
function relayMessage(uint256 _nonce, address _sender, address _target, uint256 _value, uint256 _minGasLimit, bytes _message) external payable
```

Relays a message that was sent by the other CrossDomainMessenger contract. Can only         be executed via cross-chain call from the other messenger OR if the message was         already received once and is currently being replayed.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _nonce | uint256 | Nonce of the message being relayed.
| _sender | address | Address of the user who sent the message.
| _target | address | Address that the message is targeted at.
| _value | uint256 | ETH value to send with the message.
| _minGasLimit | uint256 | Minimum amount of gas that the message can be executed with.
| _message | bytes | Message to send to the target.

### renounceOwnership

```solidity
function renounceOwnership() external nonpayable
```



*Leaves the contract without owner. It will not be possible to call `onlyOwner` functions anymore. Can only be called by the current owner. NOTE: Renouncing ownership will leave the contract without an owner, thereby removing any functionality that is only available to the owner.*


### sendMessage

```solidity
function sendMessage(address _target, bytes _message, uint32 _minGasLimit) external payable
```

Sends a message to some target address on the other chain. Note that if the call         always reverts, then the message will be unrelayable, and any ETH sent will be         permanently locked. The same will occur if the target on the other chain is         considered unsafe (see the _isUnsafeTarget() function).



#### Parameters

| Name | Type | Description |
|---|---|---|
| _target | address | Target contract or wallet address.
| _message | bytes | Message to trigger the target address with.
| _minGasLimit | uint32 | Minimum gas limit that the message can be executed with.

### successfulMessages

```solidity
function successfulMessages(bytes32) external view returns (bool)
```

Mapping of message hashes to boolean receipt values. Note that a message will only         be present in this mapping if it has successfully been relayed on this chain, and         can therefore not be relayed again.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | bytes32 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### transferOwnership

```solidity
function transferOwnership(address newOwner) external nonpayable
```



*Transfers ownership of the contract to a new account (`newOwner`). Can only be called by the current owner.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| newOwner | address | undefined

### unpause

```solidity
function unpause() external nonpayable
```

Allows the owner of this contract to resume message relaying once paused.




### xDomainMessageSender

```solidity
function xDomainMessageSender() external view returns (address)
```

Retrieves the address of the contract or wallet that initiated the currently         executing message on the other chain. Will throw an error if there is no message         currently being executed. Allows the recipient of a call to see who triggered it.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | Address of the sender of the currently executing message on the other chain.



## Events

### FailedRelayedMessage

```solidity
event FailedRelayedMessage(bytes32 indexed msgHash)
```

Emitted whenever a message fails to be relayed on this chain.



#### Parameters

| Name | Type | Description |
|---|---|---|
| msgHash `indexed` | bytes32 | Hash of the message that failed to be relayed. |

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

### Paused

```solidity
event Paused(address account)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| account  | address | undefined |

### RelayedMessage

```solidity
event RelayedMessage(bytes32 indexed msgHash)
```

Emitted whenever a message is successfully relayed on this chain.



#### Parameters

| Name | Type | Description |
|---|---|---|
| msgHash `indexed` | bytes32 | Hash of the message that was relayed. |

### SentMessage

```solidity
event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
```

Emitted whenever a message is sent to the other chain.



#### Parameters

| Name | Type | Description |
|---|---|---|
| target `indexed` | address | Address of the recipient of the message. |
| sender  | address | Address of the sender of the message. |
| message  | bytes | Message to trigger the recipient address with. |
| messageNonce  | uint256 | Unique nonce attached to the message. |
| gasLimit  | uint256 | Minimum gas limit that the message can be executed with. |

### SentMessageExtension1

```solidity
event SentMessageExtension1(address indexed sender, uint256 value)
```

Additional event data to emit, required as of Bedrock. Cannot be merged with the         SentMessage event without breaking the ABI of this contract, this is good enough.



#### Parameters

| Name | Type | Description |
|---|---|---|
| sender `indexed` | address | Address of the sender of the message. |
| value  | uint256 | ETH value sent along with the message to the recipient. |

### Unpaused

```solidity
event Unpaused(address account)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| account  | address | undefined |




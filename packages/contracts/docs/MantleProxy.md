# MantleProxy



> Proxy

Proxy is a transparent proxy that passes through the call if the caller is the owner or         if the caller is address(0), meaning that the call originated from an off-chain         simulation.



## Methods

### admin

```solidity
function admin() external nonpayable returns (address)
```

Gets the owner of the proxy contract.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | Owner address.

### changeAdmin

```solidity
function changeAdmin(address _admin) external nonpayable
```

Changes the owner of the proxy contract. Only callable by the owner.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _admin | address | New owner of the proxy contract.

### implementation

```solidity
function implementation() external nonpayable returns (address)
```

Queries the implementation address.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | Implementation address.

### upgradeTo

```solidity
function upgradeTo(address _implementation) external nonpayable
```

Set the implementation contract address. The code at the given address will execute         when this contract is called.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _implementation | address | Address of the implementation contract.

### upgradeToAndCall

```solidity
function upgradeToAndCall(address _implementation, bytes _data) external payable returns (bytes)
```

Set the implementation and call a function in a single transaction. Useful to ensure         atomic execution of initialization-based upgrades.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _implementation | address | Address of the implementation contract.
| _data | bytes | Calldata to delegatecall the new implementation with.

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bytes | undefined



## Events

### AdminChanged

```solidity
event AdminChanged(address previousAdmin, address newAdmin)
```

An event that is emitted each time the owner is upgraded. This event is part of the         EIP-1967 specification.



#### Parameters

| Name | Type | Description |
|---|---|---|
| previousAdmin  | address | The previous owner of the contract |
| newAdmin  | address | The new owner of the contract |

### Upgraded

```solidity
event Upgraded(address indexed implementation)
```

An event that is emitted each time the implementation is changed. This event is part         of the EIP-1967 specification.



#### Parameters

| Name | Type | Description |
|---|---|---|
| implementation `indexed` | address | The address of the implementation contract |




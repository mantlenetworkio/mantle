# MintableERC721Factory



> MintableERC721Factory

Factory contract for creating MintableERC721 contracts.



## Methods

### BRIDGE

```solidity
function BRIDGE() external view returns (address)
```

Address of the ERC721 bridge on this network.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### REMOTE_CHAIN_ID

```solidity
function REMOTE_CHAIN_ID() external view returns (uint256)
```

Chain ID for the remote network.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### bridge

```solidity
function bridge() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | Address of the ERC721 bridge on this network.

### createMintableERC721

```solidity
function createMintableERC721(address _remoteToken, string _name, string _symbol) external nonpayable returns (address)
```

Creates an instance of the standard ERC721.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _remoteToken | address | Address of the corresponding token on the other domain.
| _name | string | ERC721 name.
| _symbol | string | ERC721 symbol.

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### isMintableERC721

```solidity
function isMintableERC721(address) external view returns (bool)
```

Tracks addresses created by this factory.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### remoteChainId

```solidity
function remoteChainId() external view returns (uint256)
```

Chain ID for the remote network.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### version

```solidity
function version() external view returns (string)
```

Returns the full semver contract version.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | string | Semver contract version as a string.



## Events

### MintableERC721Created

```solidity
event MintableERC721Created(address indexed localToken, address indexed remoteToken, address deployer)
```

Emitted whenever a new MintableERC721 contract is created.



#### Parameters

| Name | Type | Description |
|---|---|---|
| localToken `indexed` | address | Address of the token on the this domain. |
| remoteToken `indexed` | address | Address of the token on the remote domain. |
| deployer  | address | Address of the initiator of the deployment |




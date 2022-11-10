# L1ERC721Bridge



> L1ERC721Bridge

The L1 ERC721 bridge is a contract which works together with the L2 ERC721 bridge to         make it possible to transfer ERC721 tokens from Ethereum to Mantle. This contract         acts as an escrow for ERC721 tokens deposited into L2.



## Methods

### MESSENGER

```solidity
function MESSENGER() external view returns (contract CrossDomainMessenger)
```

Messenger contract on this domain.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | contract CrossDomainMessenger | undefined

### OTHER_BRIDGE

```solidity
function OTHER_BRIDGE() external view returns (address)
```

Address of the bridge on the other network.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### bridgeERC721

```solidity
function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) external nonpayable
```

Initiates a bridge of an NFT to the caller&#39;s account on the other chain. Note that         this function can only be called by EOAs. Smart contract wallets should use the         `bridgeERC721To` function after ensuring that the recipient address on the remote         chain exists. Also note that the current owner of the token on this chain must         approve this contract to operate the NFT before it can be bridged.         **WARNING**: Do not bridge an ERC721 that was originally deployed on Mantle. This         bridge only supports ERC721s originally deployed on Ethereum. Users will need to         wait for the one-week challenge period to elapse before their Mantle-native NFT         can be refunded on L2.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _localToken | address | Address of the ERC721 on this domain.
| _remoteToken | address | Address of the ERC721 on the remote domain.
| _tokenId | uint256 | Token ID to bridge.
| _minGasLimit | uint32 | Minimum gas limit for the bridge message on the other domain.
| _extraData | bytes | Optional data to forward to the other chain. Data supplied here will not                     be used to execute any code on the other chain and is only emitted as                     extra data for the convenience of off-chain tooling.

### bridgeERC721To

```solidity
function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) external nonpayable
```

Initiates a bridge of an NFT to some recipient&#39;s account on the other chain. Note         that the current owner of the token on this chain must approve this contract to         operate the NFT before it can be bridged.         **WARNING**: Do not bridge an ERC721 that was originally deployed on Mantle. This         bridge only supports ERC721s originally deployed on Ethereum. Users will need to         wait for the one-week challenge period to elapse before their Mantle-native NFT         can be refunded on L2.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _localToken | address | Address of the ERC721 on this domain.
| _remoteToken | address | Address of the ERC721 on the remote domain.
| _to | address | Address to receive the token on the other domain.
| _tokenId | uint256 | Token ID to bridge.
| _minGasLimit | uint32 | Minimum gas limit for the bridge message on the other domain.
| _extraData | bytes | Optional data to forward to the other chain. Data supplied here will not                     be used to execute any code on the other chain and is only emitted as                     extra data for the convenience of off-chain tooling.

### deposits

```solidity
function deposits(address, address, uint256) external view returns (bool)
```

Mapping of L1 token to L2 token to ID to boolean, indicating if the given L1 token         by ID was deposited for a given L2 token.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined
| _1 | address | undefined
| _2 | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### finalizeBridgeERC721

```solidity
function finalizeBridgeERC721(address _localToken, address _remoteToken, address _from, address _to, uint256 _tokenId, bytes _extraData) external nonpayable
```

Completes an ERC721 bridge from the other domain and sends the ERC721 token to the         recipient on this domain.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _localToken | address | Address of the ERC721 token on this domain.
| _remoteToken | address | Address of the ERC721 token on the other domain.
| _from | address | Address that triggered the bridge on the other domain.
| _to | address | Address to receive the token on this domain.
| _tokenId | uint256 | ID of the token being deposited.
| _extraData | bytes | Optional data to forward to L2. Data supplied here will not be used to                     execute any code on L2 and is only emitted as extra data for the                     convenience of off-chain tooling.

### messenger

```solidity
function messenger() external view returns (contract CrossDomainMessenger)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | contract CrossDomainMessenger | Messenger contract on this domain.

### otherBridge

```solidity
function otherBridge() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | Address of the bridge on the other network.

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

### ERC721BridgeFinalized

```solidity
event ERC721BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
```

Emitted when an ERC721 bridge from the other network is finalized.



#### Parameters

| Name | Type | Description |
|---|---|---|
| localToken `indexed` | address | undefined |
| remoteToken `indexed` | address | undefined |
| from `indexed` | address | undefined |
| to  | address | undefined |
| tokenId  | uint256 | undefined |
| extraData  | bytes | undefined |

### ERC721BridgeInitiated

```solidity
event ERC721BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
```

Emitted when an ERC721 bridge to the other network is initiated.



#### Parameters

| Name | Type | Description |
|---|---|---|
| localToken `indexed` | address | undefined |
| remoteToken `indexed` | address | undefined |
| from `indexed` | address | undefined |
| to  | address | undefined |
| tokenId  | uint256 | undefined |
| extraData  | bytes | undefined |




# MintableERC721



> OptimismMintableERC721

This contract is the remote representation for some token that lives on another network,         typically an Mantle representation of an Ethereum-based token. Standard reference         implementation that can be extended or modified according to your needs.



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

Chain ID of the chain where the remote token is deployed.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### REMOTE_TOKEN

```solidity
function REMOTE_TOKEN() external view returns (address)
```

Address of the token on the remote domain.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### approve

```solidity
function approve(address to, uint256 tokenId) external nonpayable
```



*See {IERC721-approve}.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| to | address | undefined
| tokenId | uint256 | undefined

### balanceOf

```solidity
function balanceOf(address owner) external view returns (uint256)
```



*See {IERC721-balanceOf}.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| owner | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### baseTokenURI

```solidity
function baseTokenURI() external view returns (string)
```

Base token URI for this token.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | string | undefined

### bridge

```solidity
function bridge() external view returns (address)
```

Address of the ERC721 bridge on this network.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### burn

```solidity
function burn(address _from, uint256 _tokenId) external nonpayable
```

Burns a token ID from a user.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _from | address | Address of the user to burn the token from.
| _tokenId | uint256 | Token ID to burn.

### getApproved

```solidity
function getApproved(uint256 tokenId) external view returns (address)
```



*See {IERC721-getApproved}.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| tokenId | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### isApprovedForAll

```solidity
function isApprovedForAll(address owner, address operator) external view returns (bool)
```



*See {IERC721-isApprovedForAll}.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| owner | address | undefined
| operator | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### name

```solidity
function name() external view returns (string)
```



*See {IERC721Metadata-name}.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | string | undefined

### ownerOf

```solidity
function ownerOf(uint256 tokenId) external view returns (address)
```



*See {IERC721-ownerOf}.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| tokenId | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### remoteChainId

```solidity
function remoteChainId() external view returns (uint256)
```

Chain ID of the chain where the remote token is deployed.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### remoteToken

```solidity
function remoteToken() external view returns (address)
```

Address of the token on the remote domain.




#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### safeMint

```solidity
function safeMint(address _to, uint256 _tokenId) external nonpayable
```

Mints some token ID for a user, checking first that contract recipients         are aware of the ERC721 protocol to prevent tokens from being forever locked.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _to | address | Address of the user to mint the token for.
| _tokenId | uint256 | Token ID to mint.

### safeTransferFrom

```solidity
function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) external nonpayable
```



*See {IERC721-safeTransferFrom}.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| from | address | undefined
| to | address | undefined
| tokenId | uint256 | undefined
| _data | bytes | undefined

### setApprovalForAll

```solidity
function setApprovalForAll(address operator, bool approved) external nonpayable
```



*See {IERC721-setApprovalForAll}.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| operator | address | undefined
| approved | bool | undefined

### supportsInterface

```solidity
function supportsInterface(bytes4 _interfaceId) external view returns (bool)
```

Checks if a given interface ID is supported by this contract.



#### Parameters

| Name | Type | Description |
|---|---|---|
| _interfaceId | bytes4 | The interface ID to check.

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | True if the interface ID is supported, false otherwise.

### symbol

```solidity
function symbol() external view returns (string)
```



*See {IERC721Metadata-symbol}.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | string | undefined

### tokenByIndex

```solidity
function tokenByIndex(uint256 index) external view returns (uint256)
```



*See {IERC721Enumerable-tokenByIndex}.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| index | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### tokenOfOwnerByIndex

```solidity
function tokenOfOwnerByIndex(address owner, uint256 index) external view returns (uint256)
```



*See {IERC721Enumerable-tokenOfOwnerByIndex}.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| owner | address | undefined
| index | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### tokenURI

```solidity
function tokenURI(uint256 tokenId) external view returns (string)
```



*See {IERC721Metadata-tokenURI}.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| tokenId | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | string | undefined

### totalSupply

```solidity
function totalSupply() external view returns (uint256)
```



*See {IERC721Enumerable-totalSupply}.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### transferFrom

```solidity
function transferFrom(address from, address to, uint256 tokenId) external nonpayable
```



*See {IERC721-transferFrom}.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| from | address | undefined
| to | address | undefined
| tokenId | uint256 | undefined



## Events

### Approval

```solidity
event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| owner `indexed` | address | undefined |
| approved `indexed` | address | undefined |
| tokenId `indexed` | uint256 | undefined |

### ApprovalForAll

```solidity
event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| owner `indexed` | address | undefined |
| operator `indexed` | address | undefined |
| approved  | bool | undefined |

### Burn

```solidity
event Burn(address indexed account, uint256 tokenId)
```

Emitted when a token is burned.



#### Parameters

| Name | Type | Description |
|---|---|---|
| account `indexed` | address | undefined |
| tokenId  | uint256 | undefined |

### Mint

```solidity
event Mint(address indexed account, uint256 tokenId)
```

Emitted when a token is minted.



#### Parameters

| Name | Type | Description |
|---|---|---|
| account `indexed` | address | undefined |
| tokenId  | uint256 | undefined |

### Transfer

```solidity
event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| from `indexed` | address | undefined |
| to `indexed` | address | undefined |
| tokenId `indexed` | uint256 | undefined |




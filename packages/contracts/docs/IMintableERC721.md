# IMintableERC721



> IOptimismMintableERC721

Interface for contracts that are compatible with the OptimismMintableERC721 standard.         Tokens that follow this standard can be easily transferred across the ERC721 bridge.



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



*Gives permission to `to` to transfer `tokenId` token to another account. The approval is cleared when the token is transferred. Only a single account can be approved at a time, so approving the zero address clears previous approvals. Requirements: - The caller must own the token or be an approved operator. - `tokenId` must exist. Emits an {Approval} event.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| to | address | undefined
| tokenId | uint256 | undefined

### balanceOf

```solidity
function balanceOf(address owner) external view returns (uint256 balance)
```



*Returns the number of tokens in ``owner``&#39;s account.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| owner | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| balance | uint256 | undefined

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
function getApproved(uint256 tokenId) external view returns (address operator)
```



*Returns the account approved for `tokenId` token. Requirements: - `tokenId` must exist.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| tokenId | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| operator | address | undefined

### isApprovedForAll

```solidity
function isApprovedForAll(address owner, address operator) external view returns (bool)
```



*Returns if the `operator` is allowed to manage all of the assets of `owner`. See {setApprovalForAll}*

#### Parameters

| Name | Type | Description |
|---|---|---|
| owner | address | undefined
| operator | address | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### ownerOf

```solidity
function ownerOf(uint256 tokenId) external view returns (address owner)
```



*Returns the owner of the `tokenId` token. Requirements: - `tokenId` must exist.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| tokenId | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| owner | address | undefined

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
function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) external nonpayable
```



*Safely transfers `tokenId` token from `from` to `to`. Requirements: - `from` cannot be the zero address. - `to` cannot be the zero address. - `tokenId` token must exist and be owned by `from`. - If the caller is not `from`, it must be approved to move this token by either {approve} or {setApprovalForAll}. - If `to` refers to a smart contract, it must implement {IERC721Receiver-onERC721Received}, which is called upon a safe transfer. Emits a {Transfer} event.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| from | address | undefined
| to | address | undefined
| tokenId | uint256 | undefined
| data | bytes | undefined

### setApprovalForAll

```solidity
function setApprovalForAll(address operator, bool _approved) external nonpayable
```



*Approve or remove `operator` as an operator for the caller. Operators can call {transferFrom} or {safeTransferFrom} for any token owned by the caller. Requirements: - The `operator` cannot be the caller. Emits an {ApprovalForAll} event.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| operator | address | undefined
| _approved | bool | undefined

### supportsInterface

```solidity
function supportsInterface(bytes4 interfaceId) external view returns (bool)
```



*Returns true if this contract implements the interface defined by `interfaceId`. See the corresponding https://eips.ethereum.org/EIPS/eip-165#how-interfaces-are-identified[EIP section] to learn more about how these ids are created. This function call must use less than 30 000 gas.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| interfaceId | bytes4 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | undefined

### tokenByIndex

```solidity
function tokenByIndex(uint256 index) external view returns (uint256)
```



*Returns a token ID at a given `index` of all the tokens stored by the contract. Use along with {totalSupply} to enumerate all tokens.*

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
function tokenOfOwnerByIndex(address owner, uint256 index) external view returns (uint256 tokenId)
```



*Returns a token ID owned by `owner` at a given `index` of its token list. Use along with {balanceOf} to enumerate all of ``owner``&#39;s tokens.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| owner | address | undefined
| index | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| tokenId | uint256 | undefined

### totalSupply

```solidity
function totalSupply() external view returns (uint256)
```



*Returns the total amount of tokens stored by the contract.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### transferFrom

```solidity
function transferFrom(address from, address to, uint256 tokenId) external nonpayable
```



*Transfers `tokenId` token from `from` to `to`. WARNING: Usage of this method is discouraged, use {safeTransferFrom} whenever possible. Requirements: - `from` cannot be the zero address. - `to` cannot be the zero address. - `tokenId` token must be owned by `from`. - If the caller is not `from`, it must be approved to move this token by either {approve} or {setApprovalForAll}. Emits a {Transfer} event.*

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
| account `indexed` | address | Address of the account the token was burned from. |
| tokenId  | uint256 | Token ID of the burned token. |

### Mint

```solidity
event Mint(address indexed account, uint256 tokenId)
```

Emitted when a token is minted.



#### Parameters

| Name | Type | Description |
|---|---|---|
| account `indexed` | address | Address of the account the token was minted to. |
| tokenId  | uint256 | Token ID of the minted token. |

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




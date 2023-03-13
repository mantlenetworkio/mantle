
# @mantleio/sdk

[![codecov](https://codecov.io/gh/mantlenetworkio/mantle/branch/master/graph/badge.svg?token=0VTG7PG7YR&flag=sdk-tests)](https://codecov.io/gh/mantlenetworkio/mantle)

The `@mantleio/sdk` package provides a set of tools for interacting with Mantle.

## Installation

```
npm install @mantleio/sdk
```

## Docs

You can find auto-generated API documentation over at [sdk.mantle.io](https://sdk.mantle.io).

## Using the SDK

### CrossChainMessenger

The [`CrossChainMessenger`](https://github.com/mantlenetworkio/mantle/blob/develop/packages/sdk/src/cross-chain-messenger.ts) class simplifies the process of moving assets and data between Ethereum and Mantle.
You can use this class to, for example, initiate a withdrawal of ERC20 tokens from Mantle back to Ethereum, accurately track when the withdrawal is ready to be finalized on Ethereum, and execute the finalization transaction after the challenge period has elapsed.
The `CrossChainMessenger` can handle deposits and withdrawals of ETH and any ERC20-compatible token.
Detailed API descriptions can be found at [sdk.mantle.io](https://sdk.mantle.io/classes/crosschainmessenger).
The `CrossChainMessenger` automatically connects to all relevant contracts so complex configuration is not necessary.

### L2Provider and related utilities

The Mantle SDK includes [various utilities](https://github.com/mantlenetworkio/mantle/blob/develop/packages/sdk/src/l2-provider.ts) for handling Mantle's [transaction fee model](https://community.mantle.io/docs/developers/build/transaction-fees/).
For instance, [`estimateTotalGasCost`](https://sdk.mantle.io/modules.html#estimateTotalGasCost) will estimate the total cost (in wei) to send at transaction on Mantle including both the L2 execution cost and the L1 data cost.
You can also use the [`asL2Provider`](https://sdk.mantle.io/modules.html#asL2Provider) function to wrap an ethers Provider object into an `L2Provider` which will have all of these helper functions attached.

### Other utilities

The SDK contains other useful helper functions and constants.
For a complete list, refer to the auto-generated [SDK documentation](https://sdk.mantle.io/)

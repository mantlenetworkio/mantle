<div align="center">

<p><img src="./docs/assets/Mantle full logo_black_transparent.png" width="600"></p>

<p>
  <a href="https://github.com/mantlenetworkio/mantle/actions/workflows/ts-packages.yml?query=branch%3Amain"><img src="https://github.com/mantlenetworkio/mantle/workflows/typescript%20/%20contracts/badge.svg" /></a>
  <a href="https://github.com/mantlenetworkio/mantle/actions/workflows/integration.yml?query=branch%3Amain"><img src="https://github.com/mantlenetworkio/mantle/workflows/integration/badge.svg" /></a>
  <a href="https://github.com/mantlenetworkio/mantle/actions/workflows/geth.yml?query=branch%3Amain"><img src="https://github.com/mantlenetworkio/mantle/workflows/geth%20unit%20tests/badge.svg" /></a>
</p>

</div>

## Introduction
___

Mantle is a suite of Ethereum scaling solutions including an optimistic rollup and ZK rollup built using an iterative modular chain approach, and supported by BitDAO’s native token $BIT.

It is designed to bolster support for hyper-scaled throughput decentralized applications (dApps) — from derivatives decentralized exchanges (DEXs), to gaming, to operations of decentralized autonomous organizations (DAOs).

<br/>
<br/>

## Goals
---

Layer 2 rollups are built to address the scaling limitations of Ethereum by taking transaction processing to a separate execution layer, but this inevitably exposes users to the high gas fees and limited throughput on Layer 1.

Mantle's modular architecture helps achieve greater network efficiency for resources at hand, while maintaining the same level of security for all network actors. Increased network performance also enables better fraud proof and ZK proof technology, essentially unlocking the true potential of L2 rollups.

Different parts of the Mantle tech stack are specialized to tackle specific issues.

- [**Decentralized Sequencer**](https://docs.mantle.xyz/introducing-mantle/a-gentle-introduction/decentralized-sequencer): A permissionless set of L2 block producers
- [**Data Availability**](https://docs.mantle.xyz/introducing-mantle/a-gentle-introduction/solving-data-availability): Implementing EigenDA, an innovative re-staking solution that levarges Ethereum's validator network to bring the security of L1 to L2
- [**EVM-level Fraud Proofs**](https://docs.mantle.xyz/introducing-mantle/a-gentle-introduction/fraud-proofs): Improved fraud proofs that are evaluated using EVM-level instructions

> We encourage you to check out the [**Mantle tech docs**](https://docs.mantle.xyz) to learn more about the inner workings of Mantle.

</br>
</br>

## Quick Start
---

### Useful Addresses

|         Name          | Value                                |
| :-------------------: | ------------------------------------ |
| Testnet Token Faucet  | https://faucet.testnet.mantle.xyz/   |
| Mantle Testnet Bridge | https://bridge.testnet.mantle.xyz/   |
|    Mantle Explorer    | https://explorer.testnet.mantle.xyz/ |
|  Mantle Node RPC URL  | https://rpc.testnet.mantle.xyz/      |
|       Chain ID        | 5001                                 |

### Set up Local Environment

Setting up local L1 and L2 nodes may be particularly useful for testing out Mantle SDK methods.

1. Make sure your system has the following tools set up and running.
   - [Git](https://git-scm.com/downloads) - to fetch node software
   - [Node.js](https://nodejs.org/en/) - to run node instances
   - [Yarn](https://classic.yarnpkg.com/lang/en/docs/install/) - for dependancy management

2. Run L1 and L2 node instances

```bash
 git clone https://github.com/mantlenetworkio/mantle.git
 cd mantle/ops
 make up
 # check status
 make ps
```

dApps need to connect to nodes for fetching block data and sending transactions to the Mantle network. Our JSON-RPC API supports **HTTPS** and **WebSocket** connections.

|  Service  | URL                             |
| :-------: | ------------------------------- |
|    RPC    | https://rpc.testnet.mantle.xyz/ |
| WebSocket | wss://ws.testnet.mantle.xyz     |


### Using the Mantle SDK

You can use `npm` or `yarn` package managers to download and install the `@mantleio/sdk` package. We'll use `yarn` in this example.

1. Set up a project directory.

```bash
mkdir MantleSDK
cd MantleSDK
npm init --yes
```

2. Download and install the SDK package using this command.

```bash
yarn add -D @mantleio/sdk
```

3. Create a `.js` script and get started by making a request, for instance, to fetch the current L1 gas price.

```js
const ethers = require("ethers")
const mantle = require("@mantleio/sdk")

const l2RpcProvider = new ethers.providers.JsonRpcProvider("https://rpc.testnet.mantle.xyz")

async function main() {

    console.log(await mantle.getL1GasPrice(l2RpcProvider))
}

main();
```
4. Run your script using the `node <filename>.js` command to see the output.

> Feel free to browse through our [compilation of tutorials](https://mantlenetworkio.github.io/mantle-tutorial/) that use the Mantle SDK to demonstrate common functionality such as bridging assets between Mantle and Ethereum, and more.


### Using the Node RPC API

You can invoke the API endpoints by sending `curl` requests as well. Let's look at an example of a simple curl request being sent to invoke the `rollup_gasPrices` method that returns a JSON object containing the L1 and L2 gas prices used by a [Sequencer](https://docs.mantle.xyz/for-validators/network-roles#sequencers) to calculate the transaction gas fees.

> Want to get a better understanding of how gas fees are calculated on Mantle? Check out [the section on fee basics](https://docs.mantle.xyz/for-validators/transaction-fees-on-l2) in the tech docs.

```sh
# Request

curl -X POST --data '{"jsonrpc":"2.0","method":"rollup_gasPrices","params":[],"id":1}' <node url>
```

```json
// Response

{
  "jsonrpc":"2.0",
  "id":1,
  "result":{
    "l1GasPrice":"0x254aa66732",
    "l2GasPrice":"0xf3792"
  }
}
```

Check out [DEVELOP.md](./DEVELOP.md) for how we develop the Mantle.
<br/>
<br/>

## How to Contribute
___


Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of our contribution process.
Then check out our list of [good first issues](https://github.com/mantlenetworkio/mantle/contribute) to find something fun to work on!


<br/>
<br/>


## Directory Structure
___
<pre>
root
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/common-ts">common-ts</a>: Common tools for building apps in TypeScript
│   ├── <a href="./packages/contracts">contracts</a>: L1 and L2 smart contracts for Mantle
│   ├── <a href="./packages/core-utils">core-utils</a>: Low-level utilities that make building Mantle easier
│   ├── <a href="./packages/data-transport-layer">data-transport-layer</a>: Service for indexing Mantle-related L1 data
│   ├── <a href="./packages/fault-detector">fault-detector</a>: Service for detecting Sequencer faults
│   ├── <a href="./packages/message-relayer">message-relayer</a>: Tool for automatically relaying L1<>L2 messages in development
│   ├── <a href="./packages/replica-healthcheck">replica-healthcheck</a>: Service for monitoring the health of a replica node
│   └── <a href="./packages/sdk">sdk</a>: provides a set of tools for interacting with Mantle

~~ Production ~~
├── <a href="./batch-submitter">batch-submitter</a>: Service for submitting batches of transactions and results to L1
├── <a href="./bss-core">bss-core</a>: Core batch-submitter logic and utilities
├── <a href="./gas-oracle">gas-oracle</a>: Service for updating L1 gas prices on L2
├── <a href="./integration-tests">integration-tests</a>: Various integration tests for the Mantle network
├── <a href="./l2geth">l2geth</a>: Mantle client software, a fork of <a href="https://github.com/ethereum/go-ethereum/tree/v1.9.10">geth v1.9.10</a>  (deprecated for BEDROCK upgrade)
├── <a href="./l2geth-exporter">l2geth-exporter</a>: A prometheus exporter to collect/serve metrics from an L2 geth node
├── <a href="./op-exporter">op-exporter</a>: A prometheus exporter to collect/serve metrics from an Mantle node
├── <a href="./proxyd">proxyd</a>: Configurable RPC request router and proxy
├── <a href="./technical-documents">technical-documents</a>: audits and post-mortem documents
</pre>


## License
___

Code forked from [`optimism`](https://github.com/ethereum-optimism/optimism) under the name [`optimism`](https://github.com/mantlenetworkio/bitnetwork/tree/master/l2geth) is licensed under the [GNU GPLv3](https://gist.github.com/kn9ts/cbe95340d29fc1aaeaa5dd5c059d2e60) in accordance with the [original license](https://github.com/ethereum-optimism/optimism/blob/master/COPYING).

All other files within this repository are licensed under the [MIT License](https://github.com/mantlenetworkio/bitnetwork/blob/master/LICENSE) unless stated otherwise.

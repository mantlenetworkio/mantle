<div align="center">
  <h1> Mantle Wiki</h1>
</div>
<p align="center">
  <a href="https://github.com/bitnetworkio/mantle/actions/workflows/ts-packages.yml?query=branch%3Amaster"><img src="https://github.com/bitnetworkio/mantle/workflows/typescript%20/%20contracts/badge.svg" /></a>
  <a href="https://github.com/bitnetworkio/mantle/actions/workflows/integration.yml?query=branch%3Amaster"><img src="https://github.com/bitnetworkio/mantle/workflows/integration/badge.svg" /></a>
  <a href="https://github.com/bitnetworkio/mantle/actions/workflows/geth.yml?query=branch%3Amaster"><img src="https://github.com/bitnetworkio/mantle/workflows/geth%20unit%20tests/badge.svg" /></a>
</p>

## Documentation & Introduction
___

Mantle is a suite of Ethereum scaling solutions including an optimistic rollup and ZK rollup built using an iterative modular chain approach, and supported by BitDAO’s native token $BIT.

It is designed to bolster support for hyper-scaled throughput decentralized applications (dApps) — from derivatives decentralized exchanges (DEXs), to gaming, to operations of decentralized autonomous organizations (DAOs).

<br/>
<br/>

## Quick Start

___

Check out [DEVELOP.md](./DEVELOP.md) for how we develop the Mantle.
<br/>
<br/>

## How to contribute

___


Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of our contribution process.
Then check out our list of [good first issues](https://github.com/bitdao-io/mantle/contribute) to find something fun to work on!


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




<br/>
<br/>


## License
___

Code forked from [`optimism`](https://github.com/ethereum-optimism/optimism) under the name [`l2geth`](https://github.com/bitdao-io/bitnetwork/tree/master/l2geth) is licensed under the [GNU GPLv3](https://gist.github.com/kn9ts/cbe95340d29fc1aaeaa5dd5c059d2e60) in accordance with the [original license](https://github.com/ethereum-optimism/optimism/blob/master/COPYING).

All other files within this repository are licensed under the [MIT License](https://github.com/bitdao-io/bitnetwork/blob/master/LICENSE) unless stated otherwise.

# Changelog

## 0.2.0-alpha - 2022-12-23

### Features
- Implement the mechanism of triggering the Rollback instruction when the MPC signature fails ([#377](https://github.com/mantlenetworkio/mantle/pull/377)).
- Add data availability layer [eigenlayer](https://www.eigenlayer.com/) ([#395](https://github.com/mantlenetworkio/mantle/pull/395)) .
- Add the rollback mechanism ([#430](https://github.com/mantlenetworkio/mantle/pull/430)).

### Improvements
- Update hardhat version ([#382](https://github.com/mantlenetworkio/mantle/pull/382)).
- Update datalayer to support l1 hardhat and bit token ([#410](https://github.com/mantlenetworkio/mantle/pull/410)).
- Update golang version to 1.19 ([#424](https://github.com/mantlenetworkio/mantle/pull/424)).

### Bug Fixes
- Optimized tss roll back codes ([#454](https://github.com/mantlenetworkio/mantle/pull/454)).
- Complete data verification logic for eigenlayer ([#433](https://github.com/mantlenetworkio/mantle/pull/433)) .

## 0.1.1 - 2022-12-01

### Improvements
- Optimize code submission for deployed contracts ([#343](https://github.com/mantlenetworkio/mantle/pull/343)).
- Adjust l2 charging logic([#317](https://github.com/mantlenetworkio/mantle/pull/317)).

### Bug Fixes
- Delete automatic withdraw ([#323](https://github.com/mantlenetworkio/mantle/pull/323)).
- Fix tssReward contract.batchTime update exception ([#299](https://github.com/mantlenetworkio/mantle/pull/299)) .
- Fix gasPriceOracle contract modifier checkValue ([#320](https://github.com/mantlenetworkio/mantle/pull/320)).

### Deprecated
- Delete automatic burning ([#328](https://github.com/mantlenetworkio/mantle/pull/328)) .

## 0.1.0 - 2022-11-11

### Features
- MPC validators module in Layer 2 network, to minimize the trust risk of L2 execution results by threshold signature scheme(TSS).
- Native Token replacement in Layer 2 Network, to empower BitDAO's ecosystem by replacing the native token with BIT.
- Token Reward, to inspire community, organization and individual to run a l2geth-node and TSS-node.
- Gas fee adjustment, the Bit destruction mechanism will be triggered after collecting enough GasFee in the special contract.

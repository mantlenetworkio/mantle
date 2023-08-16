# Changelog

## [v0.4.3](https://github.com/mantlenetworkio/mantle/commits/v0.4.3) - 2023-08-16

### Features
- DA
  - Added DA license([#1263](https://github.com/mantlenetworkio/mantle/pull/1263))
- RollUp Service
  - Optimized the rollup service([#1301](https://github.com/mantlenetworkio/mantle/pull/1301))
  - Log optimization for mt-batcher and da-retriever([#1329](https://github.com/mantlenetworkio/mantle/pull/1329))
  - Modified for Rollup services optimization code([#1334](https://github.com/mantlenetworkio/mantle/pull/1334))
- SDK
  - Updated @mantleio/sdk to version 0.2.2([#1349](https://github.com/mantlenetworkio/mantle/pull/1349))

### Bug Fixes
- Fixed spelling errors in some documents([#1300](https://github.com/mantlenetworkio/mantle/pull/1300))
- Added Mantle audit report files([#1333](https://github.com/mantlenetworkio/mantle/pull/1333))
- Deleted error judge logic,which causing incorrect judgment logic and resulting in timeout configuration failure([#1337](https://github.com/mantlenetworkio/mantle/pull/1337))
- Handled testnet MNT token address compatibility([#1338](https://github.com/mantlenetworkio/mantle/pull/1338))
- Fixed gas oracle go mod and dockerfile([#1341](https://github.com/mantlenetworkio/mantle/pull/1341),[#1342](https://github.com/mantlenetworkio/mantle/pull/1342))
- Fixed gas oracle update frequency and block log query range([#1344](https://github.com/mantlenetworkio/mantle/pull/1344),[#1347](https://github.com/mantlenetworkio/mantle/pull/1347))
- Updated gas oracle metric, fixed bigint cmp formula and fee_Scalar false alarm([#1345](https://github.com/mantlenetworkio/mantle/pull/1345),[#1346](https://github.com/mantlenetworkio/mantle/pull/1346))


## [v0.4.3-alpha.0](https://github.com/mantlenetworkio/mantle/commits/v0.4.3-alpha.0) - 2023-08-03

### Features
- DTL
  - Optimized performance by separating L1 and L2 process execution time for synchronous data ([#1298](https://github.com/mantlenetworkio/mantle/pull/1298))
- SDK
  - Enhanced compatibility by aligning the standard-bridge with the selected chainID ([#1268](https://github.com/mantlenetworkio/mantle/pull/1268))
- GasOracle
  - Improved metric accuracy by modifying token_ratio ([#1310](https://github.com/mantlenetworkio/mantle/pull/1310))

### Bug Fixes
- Resolved a 'nonce too high' error caused by the DA verifier([#1313](https://github.com/mantlenetworkio/mantle/pull/1313))
- Fixed a panic issue during traceCall operation, eliminating unexpected program crashes ([#1305](https://github.com/mantlenetworkio/mantle/pull/1305))
- Handled unhandled errors in DA upgrade tool and NewStateTransition of l2geth, enhancing error handling and preventing potential issues ([#1288](https://github.com/mantlenetworkio/mantle/pull/1288), [#1295](https://github.com/mantlenetworkio/mantle/pull/1295))
- Strengthened seed generation process security by introducing passphrase protection ([#1297](https://github.com/mantlenetworkio/mantle/pull/1297))
- Addressed nonce overflow and other security-related concerns, bolstering system security ([#1293](https://github.com/mantlenetworkio/mantle/pull/1293), [#1292](https://github.com/mantlenetworkio/mantle/pull/1292))
- Improved JWT secret key configuration and handling for enhanced security ([#1226](https://github.com/mantlenetworkio/mantle/pull/1226))
- Renamed 'rollup' to 'Accepted on layer1', enhancing clarity and reflecting its role in the network's hierarchy ([#1314](https://github.com/mantlenetworkio/mantle/pull/1314))
- Resolved the issue of the 'Log/txstatus' element not found, addressing a gap in log information ([#1264](https://github.com/mantlenetworkio/mantle/pull/1264))
- Cleaned up unused DA contracts, optimizing the codebase by removing redundant contracts ([#1283](https://github.com/mantlenetworkio/mantle/pull/1283))
- Rectified Mantle token compilation errors on the Foundry platform, ensuring successful compilation ([#1282](https://github.com/mantlenetworkio/mantle/pull/1282))
- Fixed go mod and mt-challenger configuration issues, modified 'WMANTLE9' to 'WMANTLE' for compatibility and clarity ([#1302](https://github.com/mantlenetworkio/mantle/pull/1302), [#1281](https://github.com/mantlenetworkio/mantle/pull/1281))
- Addressed 'sigma MNT-35' issue, conducted a TSS security audit fix, ensuring accurate handling of 'sigma MNT-35' and enhancing TSS security ([#1257](https://github.com/mantlenetworkio/mantle/pull/1257))
- Fixed ConsenSys audit issue: cs-6.18, ensuring compliance with audit requirements ([#1243](https://github.com/mantlenetworkio/mantle/pull/1243))
- Addressed various ConsenSys audit issues (cs-6.30, cs-6.42, cs-6.40, cs-6.37, cs-6.14, cs-6.34, cs-6.22, cs-6.33), ensuring codebase compliance ([#1219](https://github.com/mantlenetworkio/mantle/pull/1219), [#1212](https://github.com/mantlenetworkio/mantle/pull/1212), [#1210](https://github.com/mantlenetworkio/mantle/pull/1210), [#1207](https://github.com/mantlenetworkio/mantle/pull/1207), [#1204](https://github.com/mantlenetworkio/mantle/pull/1204), [#1194](https://github.com/mantlenetworkio/mantle/pull/1194), [#1191](https://github.com/mantlenetworkio/mantle/pull/1191), [#1188](https://github.com/mantlenetworkio/mantle/pull/1188))
- Resolved lock threshold validation issue, enhancing security and preventing potential vulnerabilities ([#1227](https://github.com/mantlenetworkio/mantle/pull/1227))
- Fixed jwt secret key configuration issue, improving the handling of secret keys ([#1226](https://github.com/mantlenetworkio/mantle/pull/1226))
- Resolved ConsenSys minor audit issues (cs-6.35, cs-6.43, cs-6.44, cs-6.45), ensuring adherence to audit recommendations ([#1217](https://github.com/mantlenetworkio/mantle/pull/1217), [#1183](https://github.com/mantlenetworkio/mantle/pull/1183))
- Removed unused struct related to 'mnt-37', added a default message handler for function ProcessOneMessage:cs-6.26 ([#1202](https://github.com/mantlenetworkio/mantle/pull/1202), [#1198](https://github.com/mantlenetworkio/mantle/pull/1198))


## [v0.4.2](https://github.com/mantlenetworkio/mantle/commits/v0.4.2) - 2023-07-10

### Features
- DA
  - Adjust checking strategies for staker staking of MantleDa contracts([#1120](https://github.com/mantlenetworkio/mantle/pull/1120), [#1103](https://github.com/mantlenetworkio/mantle/pull/1103))
  - Code optimization for batch-submitter, dtl and mt-batcher([#1063](https://github.com/mantlenetworkio/mantle/pull/1063), [#1045](https://github.com/mantlenetworkio/mantle/pull/1045), [#1043](https://github.com/mantlenetworkio/mantle/pull/1043))
- TSS
  - Add permission verification to tss http api([#854](https://github.com/mantlenetworkio/mantle/pull/854))
- Packages
  - Updated the addresses for mainnet contracts in SDK, ensuring accurate and reliable contract interactions([#1247](https://github.com/mantlenetworkio/mantle/pull/1247))
- Documents
  - Removed outdated description in README, streamlining and improving clarity in the documentation([#1252](https://github.com/mantlenetworkio/mantle/pull/1252))
  - Updated the link for Optimism repository, providing an up-to-date and relevant resource([#1255](https://github.com/mantlenetworkio/mantle/pull/1255))

### Bug Fixes
- Fixed a bug in DTL data synchronization, ensuring accurate and reliable data transfer, resolving data inconsistency issues([#1258](https://github.com/mantlenetworkio/mantle/pull/1258), [#1250](https://github.com/mantlenetworkio/mantle/pull/1250))
- Fixed the contract logic in BVM_EigenDataLayrChain, updated the deployment process for EigenDataLayrChain([#1215](https://github.com/mantlenetworkio/mantle/pull/1215), [#1223](https://github.com/mantlenetworkio/mantle/pull/1223))
- Bugfix and upgrade for the DA tool, resolving issues and improving functionality([#1240](https://github.com/mantlenetworkio/mantle/pull/1240))
- Updated gasoracle metrics by adjusting the metrics port configuration([#1246](https://github.com/mantlenetworkio/mantle/pull/1246))
- Bugfix for transaction receipt unmarshaling, resolving issues related to data processing([#1236](https://github.com/mantlenetworkio/mantle/pull/1236))
- Hotfix for verifer sync range, addressing synchronization issues and improving data consistency([#1213](https://github.com/mantlenetworkio/mantle/pull/1213))
- Fixed the contract logic in BVM_EigenDataLayrChain, improving functionality and resolving issues([#1203](https://github.com/mantlenetworkio/mantle/pull/1203))


## [v0.4.2-beta.0](https://github.com/mantlenetworkio/mantle/commits/v0.4.2-beta.0) - 2023-06-30

### Bug Fixes
- Fixed the issue in the process of generating the aggregate public key([#1196](https://github.com/mantlenetworkio/mantle/pull/1196))
- Resolved a bug in the TSS node where block retrieval failed during verification([#1196](https://github.com/mantlenetworkio/mantle/pull/1196))
- Enhance the documentation for Mantle([#1190](https://github.com/mantlenetworkio/mantle/pull/1190))
- Optimize the configuration for deploying TSS staking([#1157](https://github.com/mantlenetworkio/mantle/pull/1157))
- Address and resolve the bug in the DTL Datastore([#1164](https://github.com/mantlenetworkio/mantle/pull/1164))
- Improve the resolution of DTL transaction status([#1163](https://github.com/mantlenetworkio/mantle/pull/1163))
- Enhance the measurement of parameterized gasUsed within GasOracle system([#1185](https://github.com/mantlenetworkio/mantle/pull/1185))
- Incorporate mainnet deployment into the implementation configration directory([#1166](https://github.com/mantlenetworkio/mantle/pull/1166))
- Upgrade the @mantleio/contracts package version to new release([#1155](https://github.com/mantlenetworkio/mantle/pull/1155))


## [v0.4.2-alpha.0](https://github.com/mantlenetworkio/mantle/commits/v0.4.2-alpha.0) - 2023-06-27

### Features
- Gas Oracle
  - Add additional sources for exchange rates and optimize the calculation method for token ratio([#1014](https://github.com/mantlenetworkio/mantle/pull/1014),[#1108](https://github.com/mantlenetworkio/mantle/pull/1108))
  - Improve metrics monitoring([#1102](https://github.com/mantlenetworkio/mantle/pull/1102))
  - Added a strategy of real-time adjustment of layer1 overhead based on rollup capacity, to obtain a lower tx fee experience([#926](https://github.com/mantlenetworkio/mantle/pull/926))
- L2 Fee Calculation
  - Support L2 fee collection, optimize the method for setting L2 gas price and allow for floating within a certain range([#1144](https://github.com/mantlenetworkio/mantle/pull/1144))
- Upgrade Framework
  - Enhance the upgrade framework of l2geth to support management of upgrade heights across different networks([#1007](https://github.com/mantlenetworkio/mantle/pull/1007))
- Batch Submitter
  - Expose Tss expected response with metric data, enrich handle logics for unexpected case, enrich control workflow([#1107](https://github.com/mantlenetworkio/mantle/pull/1107))
- L2geth
  - Support debug api debug_traceCall([#940](https://github.com/mantlenetworkio/mantle/pull/940))

### Bug Fixes
- Fix issues of missing permission verification in contract([#1118](https://github.com/mantlenetworkio/mantle/pull/1118))
- Fix issues of unreasonable contract naming convention([#1095](https://github.com/mantlenetworkio/mantle/pull/1095))
- Fix smart contract related bugs in audit reports([#1043](https://github.com/mantlenetworkio/mantle/pull/1043), [#1138](https://github.com/mantlenetworkio/mantle/pull/1138))
- Adjust checking strategies for staker staking of MantleDa contracts([#1120](https://github.com/mantlenetworkio/mantle/pull/1120), [#1103](https://github.com/mantlenetworkio/mantle/pull/1103))
- Code optimization for batch-submitter, dtl and mt-batcher([#1063](https://github.com/mantlenetworkio/mantle/pull/1063), [#1045](https://github.com/mantlenetworkio/mantle/pull/1045), [#1043](https://github.com/mantlenetworkio/mantle/pull/1043))
- Add permission verification to tss http api([#854](https://github.com/mantlenetworkio/mantle/pull/854))

## [v0.4.1](https://github.com/mantlenetworkio/mantle/commits/v0.4.1) - 2023-06-25

### Bug Fixes
- Remove hsm credential in log and fixed tss docker file compile bug([#1109](https://github.com/mantlenetworkio/mantle/pull/1109))

## [v0.4.1-beta.0](https://github.com/mantlenetworkio/mantle/commits/v0.4.1-beta.0) - 2023-06-19

### Bug Fixes
- Fixed mt-batcher-fee creating signature null pointer from hsm([#971](https://github.com/mantlenetworkio/mantle/pull/971))
- Fixed replica returning http error from dt query txstatus([#979](https://github.com/mantlenetworkio/mantle/pull/979))
- Fixed the issue with improper use of go.work([#980](https://github.com/mantlenetworkio/mantle/pull/980)、[#982](https://github.com/mantlenetworkio/mantle/pull/982)、[#983](https://github.com/mantlenetworkio/mantle/pull/983)、[#996](https://github.com/mantlenetworkio/mantle/pull/996))
- Fixed an issue where a problem node could not be slashed by placing the logic of whether to penalize tssnode in tssmanager and having tssnode verify whether it can execute([#998](https://github.com/mantlenetworkio/mantle/pull/998))
- Fixed a data synchronization bug by optimizing the data synchronization logic of DTL([#1001](https://github.com/mantlenetworkio/mantle/pull/1001))

## [v0.4.0](https://github.com/mantlenetworkio/mantle/commits/v0.4.0) - 2023-06-08

### Features
- Integrate fraud proof feature into Mantle network([#814](https://github.com/mantlenetworkio/mantle/pull/814))
- Adding delegation functionality for fraud proof and MPC modules, supporting external nodes to participate in MPC.([#826](https://github.com/mantlenetworkio/mantle/pull/826))
- $BIT -> $MNT([#942](https://github.com/mantlenetworkio/mantle/pull/942))
  - Use Mantle token instead of BitDAO token as the native token of the Mantle network.
  - Use Mantle token for various purposes such as gas fee payments and other applications.
- Integrated MantleDA (MantleDA built on EigenDA technology),  add the joining and exiting mechanism function of da node. （[#896](https://github.com/mantlenetworkio/mantle/pull/896)、[#947](https://github.com/mantlenetworkio/mantle/pull/947)）
- Define fraud proof metrics to serve data visualization and alert monitoring([#892](https://github.com/mantlenetworkio/mantle/pull/892))
- Define tssmanager metrics to serve data visualization and alert monitoring([#951](https://github.com/mantlenetworkio/mantle/pull/951))

### Bug Fixes
- Fixed an issue where l1 tip didn’t take effect when calculating l1 price([#869](https://github.com/mantlenetworkio/mantle/pull/869))

## [0.3.2](https://github.com/mantlenetworkio/mantle/commits/v0.3.2) - 2023-05-10

### Features
- Reduce the amount of data sent to Layer1 rollup using EigenLayer, significantly lowering gas fees for Layer2([#811](https://github.com/mantlenetworkio/mantle/pull/811)、[#825](https://github.com/mantlenetworkio/mantle/pull/825))

## [0.3.1](https://github.com/mantlenetworkio/mantle/commits/v0.3.1) - 2023-04-13

### Features
- Changed the rollup method from Layer1 CTC contract to EigenDA([#779](https://github.com/mantlenetworkio/mantle/pull/779))
- Added support for data verification and recovery based on DA‘s api service in the Verifier node([#803](https://github.com/mantlenetworkio/mantle/pull/803))

### Improvements
- Optimized re-rollup process([#789](https://github.com/mantlenetworkio/mantle/pull/789))
- Optimized code for contract upgrades([#798](https://github.com/mantlenetworkio/mantle/pull/798))

## [0.3.0](https://github.com/mantlenetworkio/mantle/commits/v0.3.0) - 2023-03-16

### Features
- Verifier syncs data from the DA network, verifies and generates new blocks([#638](https://github.com/mantlenetworkio/mantle/pull/638))
- Make a verification mechanism for the data stored in the DA layer, and rollup again if the storage fails([#670](https://github.com/mantlenetworkio/mantle/pull/670))

### Improvements
- Improve rollback logic to account for Layer-1 reorgs([#635](https://github.com/mantlenetworkio/mantle/pull/635))
- dockerfile upgrade go1.19 ([#656](https://github.com/mantlenetworkio/mantle/pull/656))

### Bug Fixes
- Fix the array out-of-bounds bug of batchsubmitter([#622](https://github.com/mantlenetworkio/mantle/pull/622))
- Optimized performance when tss processes large batch of stateroots verification([#627](https://github.com/mantlenetworkio/mantle/pull/627))
- Solve the problem of inconsistent blockhash generated by the sequencer and verifier nodes due to the gas limit([#630](https://github.com/mantlenetworkio/mantle/pull/630))
- Verify stateroot when verifier syncs data from Layer1 and DA([#648](https://github.com/mantlenetworkio/mantle/pull/648))
- Fixed the state root bug for verifier nodes([#658](https://github.com/mantlenetworkio/mantle/pull/658))

## [0.2.0](https://github.com/mantlenetworkio/mantle/commits/v0.2.0) - 2023-01-09

### Features
- add white list for da fraud proof([#496](https://github.com/mantlenetworkio/mantle/issues/496))
- support mt-batcher and mt-challenger send transaction by EIP1559([#524](https://github.com/mantlenetworkio/mantle/issues/524))
- da challenger integrate to mantle([#527](https://github.com/mantlenetworkio/mantle/issues/527))
- add da rollup min and max size config([#528](https://github.com/mantlenetworkio/mantle/issues/528))
- Add WBIT Token on mantle network([#540](https://github.com/mantlenetworkio/mantle/issues/540))

### Improvements
- Make challenge period time for fraud proofs configurable([#461](https://github.com/mantlenetworkio/mantle/pull/461))
- Updates to main README.md([#543](https://github.com/mantlenetworkio/mantle/pull/543))
- sdk bump to v0.1.4([#546](https://github.com/mantlenetworkio/mantle/pull/546))

### Bug Fixes
- fix bug when tss manager get signature from store ([#538](https://github.com/mantlenetworkio/mantle/pull/538))

## [0.2.0-alpha](https://github.com/mantlenetworkio/mantle/commits/v0.2.0-alpha) - 2022-12-23

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

## [0.1.1](https://github.com/mantlenetworkio/mantle/commits/v0.1.1) - 2022-12-01


### Improvements
- Optimize code submission for deployed contracts ([#343](https://github.com/mantlenetworkio/mantle/pull/343)).
- Adjust l2 charging logic([#317](https://github.com/mantlenetworkio/mantle/pull/317)).

### Bug Fixes
- Delete automatic withdraw ([#323](https://github.com/mantlenetworkio/mantle/pull/323)).
- Fix tssReward contract.batchTime update exception ([#299](https://github.com/mantlenetworkio/mantle/pull/299)) .
- Fix gasPriceOracle contract modifier checkValue ([#320](https://github.com/mantlenetworkio/mantle/pull/320)).

### Deprecated
- Delete automatic burning ([#328](https://github.com/mantlenetworkio/mantle/pull/328)) .

## [0.1.0](https://github.com/mantlenetworkio/mantle/commits/v0.1.0) - 2022-11-11


### Features
- MPC validators module in Layer 2 network, to minimize the trust risk of L2 execution results by threshold signature scheme(TSS).
- Native Token replacement in Layer 2 Network, to empower BitDAO's ecosystem by replacing the native token with BIT.
- Token Reward, to inspire community, organization and individual to run a l2geth-node and TSS-node.
- Gas fee adjustment, the Bit destruction mechanism will be triggered after collecting enough GasFee in the special contract.

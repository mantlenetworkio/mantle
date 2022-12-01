# EigenLayr Smart Contract Tests
This document provides some context for EigenLayr's automated tests, design philsophy for them, and broad tracking of goals.
Existing tests are broken into a few suites.
Several suites -- namely 'Investment', 'Delegation', and 'DataLayr' -- depend on a shared setup and use quite a few shared/modular functions; this shared setup/functionality is contained in the 'Deployer.t.sol' file.
Other existing suites have more limited dependencies and currently cover:
-'Disclosure' -- covers forced disclosure of data in DataLayr
-'NFGT' -- covers basic testing of experimental NFGT token

## High Level Goals
1. Better Comments -- comments should insightful/explanatory, perhaps in NatSpec or other more proper format.
2. Simpler Tests -- ideally a test breaking should make it easy to identify which **specific** functionality has been broken.
3. Tests Separated into Separate Files -- tests should be organized into 'suites' which focus on a subset of functionality (e.g. auth/permissions, investment, delegation, etc).
4. Modular Functions -- we should avoid copy-pasting large code blocks. Complex tests should consist mostly of composing smaller modular blocks, which may also be used elsewhere.
5. Reduce Hard-Coding -- while hard-coding may be useful in initially writing a test, we should avoid hard-coding values when possible. This also helps with creating fuzzed tests, as well as making tests less 'brittle', so they only break when they changes are harmful. In the interim, where possible please declare variables with fixed values rather than directly hard-coding values.

## TODOs
-forced disclosure for DL
-robustness of sig checker
-fault/revert testing (ensuring that improper inputs cause failures/reversions)
-shares in delegation
-more involved/complex investment tests (esp. verifying updates to dynamic arrays)
-permission/auth suite
-DelegationTerms suite
-Governance suite
-Slashing suite
-proof that libraries are good (esp. BLS -- perhaps an opportunity for formal verification methods?)

## Aspirations
-more integration with off-chain code
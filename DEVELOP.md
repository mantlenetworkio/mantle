
## Development Quick Start

### Dependencies

You'll need the following:

* [Git](https://git-scm.com/downloads)
* [NodeJS](https://nodejs.org/en/download/)
* [Node Version Manager](https://github.com/nvm-sh/nvm)
* [Yarn](https://classic.yarnpkg.com/en/docs/install)
* [Docker](https://docs.docker.com/get-docker/)
* [Docker Compose](https://docs.docker.com/compose/install/)
* [Foundry](https://getfoundry.sh)

### Setup

Clone the repository, open it, and install nodejs packages with `yarn`:

```bash
git clone git@github.com:bitnetworkio/mantle.git
cd mantle
yarn install
```

### Install the Correct Version of NodeJS

Using `nvm`, install the correct version of NodeJS.

```
nvm use
```

### Building the TypeScript packages

[foundry](https://github.com/foundry-rs/foundry) is used for some smart contract
development in the monorepo. It is required to build the TypeScript packages
and compile the smart contracts. Install foundry [here](https://getfoundry.sh/).

To build all of the [TypeScript packages](./packages), run:

```bash
yarn clean
yarn build
```

Packages compiled when on one branch may not be compatible with packages on a different branch.
**You should recompile all packages whenever you move from one branch to another.**
Use the above commands to recompile the packages.

### Building the rest of the system

If you want to run an mantle node OR **if you want to run the integration tests**, you'll need to build the rest of the system.

```bash
cd ops
export COMPOSE_DOCKER_CLI_BUILD=1 # these environment variables significantly speed up build time
export DOCKER_BUILDKIT=1
docker-compose build
```

This will build the following containers:

* [`l1_chain`](https://hub.docker.com/r/ethereumbitnetwork/hardhat): simulated L1 chain using hardhat-evm as a backend
* [`deployer`](https://hub.docker.com/r/ethereumbitnetwork/deployer): process that deploys L1 smart contracts to the L1 chain
* [`dtl`](https://hub.docker.com/r/ethereumbitnetwork/data-transport-layer): service that indexes transaction data from the L1 chain
* [`l2geth`](https://hub.docker.com/r/ethereumbitnetwork/l2geth): L2 geth node running in Sequencer mode
* [`verifier`](https://hub.docker.com/r/ethereumbitnetwork/go-ethereum): L2 geth node running in Verifier mode
* [`relayer`](https://hub.docker.com/r/ethereumbitnetwork/message-relayer): helper process that relays messages between L1 and L2
* [`batch_submitter`](https://hub.docker.com/r/ethereumbitnetwork/batch-submitter): service that submits batches of Sequencer transactions to the L1 chain
* [`integration_tests`](https://hub.docker.com/r/ethereumbitnetwork/integration-tests): integration tests in a box

If you want to make a change to a container, you'll need to take it down and rebuild it.
For example, if you make a change in l2geth:

```bash
cd ops
docker-compose stop -- l2geth
docker-compose build -- l2geth
docker-compose start l2geth
```

Source code changes can have an impact on more than one container.
**If you're unsure about which containers to rebuild, just rebuild them all**:

```bash
cd ops
docker-compose down
docker-compose build
docker-compose up
```

**If a node process exits with exit code: 137** you may need to increase the default memory limit of docker containers

Finally, **if you're running into weird problems and nothing seems to be working**, run:

```bash
cd mantle
yarn clean
yarn build
cd ops
docker-compose down -v
docker-compose build
docker-compose up
```

#### Viewing docker container logs

By default, the `docker-compose up` command will show logs from all services, and that
can be hard to filter through. In order to view the logs from a specific service, you can run:

```bash
docker-compose logs --follow <service name>
```

### Running tests

Before running tests: **follow the above instructions to get everything built.**

#### Running unit tests

Run unit tests for all packages in parallel via:

```bash
yarn test
```

To run unit tests for a specific package:

```bash
cd packages/package-to-test
yarn test
```

#### Running integration tests

Follow above instructions for building the whole stack.
Build and run the integration tests:

```bash
cd integration-tests
yarn build
yarn test:integration
```
#### Running contract static analysis

We perform static analysis with [`slither`](https://github.com/crytic/slither).
You must have Python 3.x installed to run `slither`.
To run `slither` locally, do:

```bash
cd packages/contracts
pip3 install slither-analyzer
yarn test:slither
```
## Set Up Local Environment

Check out [ops/README.local.md](./ops/README.local.md).

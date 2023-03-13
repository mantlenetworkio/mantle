# @mantleio/chain-mon

[![codecov](https://codecov.io/gh/mantlenetworkio/mantle/branch/develop/graph/badge.svg?token=0VTG7PG7YR&flag=chain-mon-tests)](https://codecov.io/gh/mantlenetworkio/mantle)

`chain-mon` is a collection of chain monitoring services.

## Installation

Clone, install, and build the Mantle monorepo:

```
git clone https://github.com/mantlenetworkio/mantle.git
yarn install
yarn build
```

## Running a service

Copy `.env.example` into a new file named `.env`, then set the environment variables listed there depending on the service you want to run.
Once your environment variables have been set, run via:

```
yarn start:<service name>
```

For example, to run `drippie-mon`, execute:

```
yarn start:drippie-mon
```

# @bidaoio/fault-detector

The `fault-detector` is a simple service for detecting discrepancies between your local view of the Mantle network and the L2 output proposals published to Ethereum.

## Installation

Clone, install, and build the Mantle monorepo:

```
git clone https://github.com/mantlenetworkio/mantle.git
yarn install
yarn build
```

## Running the service

Copy `.env.example` into a new file named `.env`, then set the environment variables listed there.
Once your environment variables have been set, run the service via:

```
yarn start
```

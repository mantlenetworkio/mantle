#!/bin/sh

set -e

RETRIES=${RETRIES:-40}

if [ $USE_SECRET_MANAGER ] ;then
  echo "use secret-manager"
else
  echo "doesn't use secret-manager"
fi


if [[ ! -z "$URL" ]]; then
    # get the addrs from the URL provided
    ADDRESSES=$(curl --fail --show-error --silent --retry-connrefused --retry $RETRIES --retry-delay 5 $URL)
    # set the env
    export CTC_ADDRESS=$(echo $ADDRESSES | jq -r '.CanonicalTransactionChain')
    export SCC_ADDRESS=$(echo $ADDRESSES | jq -r '.StateCommitmentChain')
    export FP_ROLLUP_ADDRESS=$(echo $ADDRESSES | jq -r '.Proxy__Rollup')
fi


# waits for l2geth to be up
curl -XPOST \
    -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
    --fail \
    --show-error \
    --silent \
    --retry-connrefused \
    --retry $RETRIES \
    --retry-delay 1 \
    --output /dev/null \
    $L2_ETH_RPC

# go
exec batch-submitter "$@"

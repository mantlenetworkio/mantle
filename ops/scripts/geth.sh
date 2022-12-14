#!/bin/sh

# FIXME: Cannot use set -e since bash is not installed in Dockerfile
# set -e

RETRIES=${RETRIES:-40}
VERBOSITY=${VERBOSITY:-6}

# get the genesis file from the deployer
curl \
    --fail \
    --show-error \
    --silent \
    --retry-connrefused \
    --retry-all-errors \
    --retry $RETRIES \
    --retry-delay 5 \
    $ROLLUP_STATE_DUMP_PATH \
    -o genesis.json

# wait for the dtl to be up, else geth will crash if it cannot connect
curl \
    --fail \
    --show-error \
    --silent \
    --output /dev/null \
    --retry-connrefused \
    --retry $RETRIES \
    --retry-delay 1 \
    $ROLLUP_CLIENT_HTTP

# import the key that will be used to locally sign blocks
# this key does not have to be kept secret in order to be secure
# we use an insecure password ("pwd") to lock/unlock the password
echo "Importing private key"
echo $BLOCK_SIGNER_KEY > key.prv
echo "pwd" > password
geth account import --password ./password ./key.prv

# initialize the geth node with the genesis file
echo "Initializing Geth node"
geth --verbosity="$VERBOSITY" "$@" init genesis.json
echo $BLOCK_SCHEDULER_ADDRESS
echo $SCHEDULER_P2P_ENODE

# start the geth node
echo "Starting Geth node"
if [ $IS_SEQUENCER == "true" ] ;then
  echo "we are sequencer node!!"
  exec geth \
    --verbosity="$VERBOSITY" \
    --password ./password \
    --allow-insecure-unlock \
    --unlock $BLOCK_SIGNER_ADDRESS \
    --bootnodes $SCHEDULER_P2P_ENODE \
    --nat $NAT \
    --port $P2P_PORT \
    --mine \
    --scheduler.address $BLOCK_SCHEDULER_ADDRESS \
#    --miner.etherbase $BLOCK_SIGNER_ADDRESS \
    --sequencer.mode="true" \
    "$@"
else [ $IS_SCHEDULER == "true" ]
  exec geth \
    --verbosity="$VERBOSITY" \
    --password ./password \
    --allow-insecure-unlock \
    --unlock $BLOCK_SIGNER_ADDRESS \
    --nat $NAT \
    --mine \
    --scheduler.address $BLOCK_SCHEDULER_ADDRESS \
#    --miner.etherbase $BLOCK_SIGNER_ADDRESS \
    "$@"
fi

  exec geth \
    --verbosity="$VERBOSITY" \
    --password ./password \
    --allow-insecure-unlock \
    --unlock $BLOCK_SIGNER_ADDRESS \
    --mine \
    --scheduler.address $BLOCK_SCHEDULER_ADDRESS \
#    --miner.etherbase $BLOCK_SIGNER_ADDRESS \
    "$@"
#
#./geth --datadir data --scheduler.address 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 --password password.txt --unlock 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 --nodiscover --verbosity 5
#./geth --datadir data1 --rpcport 8085 --port 30306 --password password.txt --unlock 0x70997970c51812dc3a010c7d01b50e0d17dc79c8 --nodiscover --scheduler.address 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 --sequencer.mode=true --verbosity 5
#./geth --datadir data2 --rpcport 8086 --port 30307 --password password.txt --unlock 0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc --nodiscover --scheduler.address 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 --sequencer.mode=true --verbosity 5

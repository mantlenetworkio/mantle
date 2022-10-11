## scripts usage

### init-key.sh

platform: aws ubuntu
user: root


```shell
export PASSWORD=${replace_string}, i.e: export PASSWORD=mantle@123
```

### init-l1-contract

#### build the images
```shell
docker build -t mantle/initl1 -f ops/docker/initl1/Dockerfile .
```

#### start deploying contracts(i.e: goerli-qa)
```shell


docker run --net bridge -itd  --restart on-failure  \
 -e "AUTOMATICALLY_TRANSFER_OWNERSHIP=true" \
-e "ETHERSCAN_API_KEY=B1XAN986315AME96W9QK7X1RGQ6WJMWEPW" \
-e "CONTRACTS_RPC_URL=https://eth-goerli.g.alchemy.com/v2/821_LFssCCQnEG3mHnP7tSrc87IQKsUp" \
-e "CONTRACTS_DEPLOYER_KEY=114983389b91256e6e26cb0e30371ceb5385e3fb309cd0e7823b86fff27c9ba4" \
-e "CONTRACTS_TARGET_NETWORK=goerli-qa"  \
--name=init-contract mantle/initl1

```

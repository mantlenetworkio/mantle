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
-e "CONTRACTS_DEPLOYER_KEY=e4bf8c09fc7bb5c3eb932260b9fcf0f2a3fecb61512b0e979afb4ce1187bfe70" \
-e "CONTRACTS_TARGET_NETWORK=goerli-qa"  \
--name=init-contract mantle/initl1

```

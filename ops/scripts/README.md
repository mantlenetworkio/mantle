## scripts usage

### init-key.sh

platform: aws ubuntu
user: root


```shell
export PASSWORD=${replace_string}, i.e: export PASSWORD=bitnetwork@123
```

### init-l1-contract

#### build the images
```shell
docker build -t bitnetwork/initl1 -f ops/docker/initl1/Dockerfile .
```

#### start deploying contracts(i.e: goerli-qa)
```shell


docker run --net bridge -itd  --restart on-failure  \
 -e "AUTOMATICALLY_TRANSFER_OWNERSHIP=true" \
-e "ETHERSCAN_API_KEY=B1XAN986315AME96W9QK7X1RGQ6WJMWEPW" \
-e "CONTRACTS_RPC_URL=https://eth-goerli.g.alchemy.com/v2/821_LFssCCQnEG3mHnP7tSrc87IQKsUp" \
-e "CONTRACTS_DEPLOYER_KEY=d04820edb3ce80d3b24595b14e32555280599ea4349a5601788154adbe19d6d4" \
-e "CONTRACTS_TARGET_NETWORK=goerli-qa"  \
--name=init-contract bitnetwork/initl1

```

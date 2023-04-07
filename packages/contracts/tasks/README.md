### task

##### urgent-upgrade-task

###### local:setL1BridgeChugCode

tranfer to addressmanager owner
```shell
cast send -f 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 \
--private-key ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
 --legacy --rpc-url http://localhost:9545 \
 --value 10ether 0xd5add52d36399570e56c183d949da83ac29aa7d6
```

query balances
```shell
cast balance --rpc-url http://localhost:9545 \
0xd5add52d36399570e56c183d949da83ac29aa7d6
```

convert unit
```shell
cast --to-unit 1000000000000000000wei ether
```

upgradecode
```shell
yarn hardhat setL1BridgeChugCode --contract 0x52753615226F8aC8a464bfecb11Ef798CFF3793f --network local

```

###### local:setTssGroupManagerCode

upgradecode
```shell
yarn hardhat udpateTssGroupManagerCode --contract 0xa83239cf2b900682001f9144144B5E5e5788A631 --network local

```

```shell
cast call --rpc-url  http://localhost:9545 \
--private-key 26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e \
--from 0xd5add52d36399570e56c183d949da83ac29aa7d6 0xa83239cf2b900682001f9144144B5E5e5788A631 "implementation()"
```


###### local:setTssStakingSlashCode

upgradecode
```shell
yarn hardhat updateTssStakingSlashCode --contract 0xe6cd9e7b620964bECd42c7Ad41e56724f515E284 --network local

```


```shell
cast call --rpc-url  http://localhost:9545 \
--private-key 26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e \
--from 0xd5add52d36399570e56c183d949da83ac29aa7d6 0xe6cd9e7b620964bECd42c7Ad41e56724f515E284 "implementation()"
```


###### local:updateEigenDataLayrChainCode

upgradecode
```shell
yarn hardhat updateEigenDataLayrChainCode --contract 0x5a0069E211A28cBD1a7dbD585877596FeD07805b --network local

```

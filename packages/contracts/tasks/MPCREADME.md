## upgrade-task

### branch
- rde：feature/ctc-rollup-switch-to-eigenlayer
- mantle:archive/urgent-contract-upgrade
- eigenlayr-contracts-mantle:archive/urgent_upgrade


### urgent-upgrade-task-l1


#### catuion：
for every time we upgrade the contract on mainnet，we need to change the
deploy address.

#### local:updateL1BridgeChug        mpc verified

tranfer eth balance to addressmanager owner
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

mpcUpdateL1BridgeChug
```shell
yarn hardhat mpcUpdateL1BridgeChug --contract 0x802e4857306AF17f654452112a50E58f4a58e28D --network goerli
```

```shell
cast call --rpc-url  https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28 0x802e4857306AF17f654452112a50E58f4a58e28D "getTestUpdate()"
cast send --rpc-url  https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28  \
--private-key 574108d5a6bfa179e9727887e315c1cd08ec2f8ca09dd4f00b1abd591011375f \
--from 0xFca9E706e5b0AE97B8F98F747F57Fa64f75EC48D 0x802e4857306AF17f654452112a50E58f4a58e28D "setTestUpdate(uint256)"  121
cast call --rpc-url  https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28 0x1eFfED12A9408c830fD19535F89131f1829debc4 "testUpdate()"
```

#### local: mpcUpdateL1CrossDomainMessenger   mpc verified

```shell
yarn hardhat mpcUpdateL1CrossDomainMessenger --contract 0xfDe0ef603c3DCbcbB94F451238caC39Adad41918 --network goerli
```

query owner:
```
cast call --rpc-url  http://localhost:9545 \
--private-key 26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e \
--from 0xd5add52d36399570e56c183d949da83ac29aa7d6 0x19C22f181280dF6Ad1d97285cdD430173Df91C12 "owner()"
```

```shell
cast call --rpc-url  https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28 0x1eFfED12A9408c830fD19535F89131f1829debc4 "getTestUpdate()"
cast send --rpc-url  https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28 --private-key 574108d5a6bfa179e9727887e315c1cd08ec2f8ca09dd4f00b1abd591011375f \
--from 0xFca9E706e5b0AE97B8F98F747F57Fa64f75EC48D 0xfDe0ef603c3DCbcbB94F451238caC39Adad41918 "setTestUpdate(uint256)"  12
cast call --rpc-url  https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28 0x1eFfED12A9408c830fD19535F89131f1829debc4 "testUpdate()"
```


#### local:mpcUpdateTssGroupManager
if we want to query the implementation, we must use the admin address and privatekey.

upgrade
```shell
yarn hardhat mpcUpdateTssGroupManager --contract 0xa83239cf2b900682001f9144144B5E5e5788A631 --network goerli
```

query implementation
```shell

cast call --rpc-url  https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28  0x4A1fC57Be953Fa5C24868457C025F116863Be7ba "owner()"

cast call --rpc-url  http://localhost:9545 \
--private-key 26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e \
--from 0xd5add52d36399570e56c183d949da83ac29aa7d6 0xa83239cf2b900682001f9144144B5E5e5788A631 "implementation()"
```

```shell
cast call --rpc-url  https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28 0x1eFfED12A9408c830fD19535F89131f1829debc4 "getTestUpdate()"
cast call --rpc-url  https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28 0x1eFfED12A9408c830fD19535F89131f1829debc4 "setTestUpdate(uint256)"  12
cast call --rpc-url  https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28 0x1eFfED12A9408c830fD19535F89131f1829debc4 "testUpdate()"
```



#### local:mpcUpdateTssStakingSlash

update
```shell
yarn hardhat mpcUpdateTssStakingSlash --contract 0xe6cd9e7b620964bECd42c7Ad41e56724f515E284 --network goerli
```

query implementation:
```shell
cast call --rpc-url  http://localhost:9545 \
--private-key 26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e \
--from 0xd5add52d36399570e56c183d949da83ac29aa7d6 0xe6cd9e7b620964bECd42c7Ad41e56724f515E284 "implementation()"
```


### urgent-upgrade-task-eigenda

eigenda有相同的proxyadmin地址，可以通过proxyadmin来查询某个proxy的implementation地址。
```shell
/**
 * @dev This is an auxiliary contract meant to be assigned as the admin of a {TransparentUpgradeableProxy}. For an
 * explanation of why you would want to use this see the documentation for {TransparentUpgradeableProxy}.
 */
contract ProxyAdmin is Ownable {
    /**
     * @dev Returns the current implementation of `proxy`.
     *
     * Requirements:
     *
     * - This contract must be the admin of `proxy`.
     */
    function getProxyImplementation(TransparentUpgradeableProxy proxy) public view virtual returns (address) {
        // We need to manually run the static call since the getter cannot be flagged as view
        // bytes4(keccak256("implementation()")) == 0x5c60da1b
        (bool success, bytes memory returndata) = address(proxy).staticcall(hex"5c60da1b");
        require(success);
        return abi.decode(returndata, (address));
    }

```

#### local:mpcUpdateEigenDataLayrChain

query implementation demo:
```shell
cast call --rpc-url  http://localhost:9545 \
--private-key 74e58c0127a59c8745568e7b4b6f41a4ad27875d2678358e0a0431f8385e5e9d \
--from 0xd5da011954f654e8192ffc3bd8469fd997c360fc 0x216C8c9815fe6e43222D94De5598e17C014A1b99 "getProxyImplementation(address)" 0x535211625Ec42aAB8BF018A33afa9729e6AaB634


0x000000000000000000000000839f3616eeac84496e1b8e519fe824247f5a8250
```
- private-key da-01 owner
- from owner
- proxyadmin contract address: 0x216C8c9815fe6e43222D94De5598e17C014A1b99
- transparentproxy contract address: 0x535211625Ec42aAB8BF018A33afa9729e6AaB634

##### EigenLayrDelegation
部分部署合约地址路径：mantle/datalayr-mantle/contracts/eignlayr-contracts/data

```shell
cast call --rpc-url  http://localhost:9545 \
--private-key 74e58c0127a59c8745568e7b4b6f41a4ad27875d2678358e0a0431f8385e5e9d \
--from 0xd5da011954f654e8192ffc3bd8469fd997c360fc 0x216C8c9815fe6e43222D94De5598e17C014A1b99 "getProxyImplementation(address)" 0x535211625Ec42aAB8BF018A33afa9729e6AaB634

0x000000000000000000000000839f3616eeac84496e1b8e519fe824247f5a8250
```

##### InvestmentManager
```shell
cast call --rpc-url  http://localhost:9545 \
--private-key 74e58c0127a59c8745568e7b4b6f41a4ad27875d2678358e0a0431f8385e5e9d \
--from 0xd5da011954f654e8192ffc3bd8469fd997c360fc 0x216C8c9815fe6e43222D94De5598e17C014A1b99 "getProxyImplementation(address)" 0xEeD0Ca239a6664b905fE49d992965584FFA8fBe0

0x0000000000000000000000009c938c57460b0b6c77fdf8b33dd70e0992d2a736
```

##### Slasher
```shell
cast call --rpc-url  http://localhost:9545 \
--private-key 74e58c0127a59c8745568e7b4b6f41a4ad27875d2678358e0a0431f8385e5e9d \
--from 0xd5da011954f654e8192ffc3bd8469fd997c360fc 0x216C8c9815fe6e43222D94De5598e17C014A1b99 "getProxyImplementation(address)" 0xdF9DAc4dA10C04eB92F1c208877EB5A94681dF3e

0x000000000000000000000000c623aca11542b6fe4636cca97a23420b701aa2d4
```

##### EigenPodManager
```shell
cast call --rpc-url  http://localhost:9545 \
--private-key 74e58c0127a59c8745568e7b4b6f41a4ad27875d2678358e0a0431f8385e5e9d \
--from 0xd5da011954f654e8192ffc3bd8469fd997c360fc 0x216C8c9815fe6e43222D94De5598e17C014A1b99 "getProxyImplementation(address)" 0xcaeBE568937802510001DFf643aAD22C6F5Fb4e2

0x00000000000000000000000064061d4abed2203376a5c3b6a5c4783eaeb81e54
```

##### InvestmentStrategyBase WETH
wethstart proxy: 0xc8eef68bec1DdE402cc2075A3AcD5E478A20d8A7

```shell
cast call --rpc-url  http://localhost:9545 \
--private-key 74e58c0127a59c8745568e7b4b6f41a4ad27875d2678358e0a0431f8385e5e9d \
--from 0xd5da011954f654e8192ffc3bd8469fd997c360fc 0x216C8c9815fe6e43222D94De5598e17C014A1b99 "getProxyImplementation(address)" 0xc8eef68bec1DdE402cc2075A3AcD5E478A20d8A7

0x0000000000000000000000002fbaa9d0e2a80577a1342447e9ef8ed18c86800e
```

##### InvestmentStrategyBase BIT
BIT proxy: 0xe0D430af034477F8Dc7977E35177c4216a42BDa3


```shell
cast call --rpc-url  http://localhost:9545 \
--private-key 74e58c0127a59c8745568e7b4b6f41a4ad27875d2678358e0a0431f8385e5e9d \
--from 0xd5da011954f654e8192ffc3bd8469fd997c360fc 0x216C8c9815fe6e43222D94De5598e17C014A1b99 "getProxyImplementation(address)" 0xe0D430af034477F8Dc7977E35177c4216a42BDa3

0x0000000000000000000000002fbaa9d0e2a80577a1342447e9ef8ed18c86800e
```


#### upgrade contract test

```shell
forge script script/Upgrade.s.sol:EigenLayrUpgrade \
--rpc-url http://localhost:9545  --private-key \
74e58c0127a59c8745568e7b4b6f41a4ad27875d2678358e0a0431f8385e5e9d \
--broadcast --slow -vvvv
```

### urgent-upgrade-task-l2
for l2 upgrade, we use **hard-code model.**

#### sequence 1
change the code by state_transition.go
https://github.com/mantlenetworkio/mantle/blob/ac8ed3fa93ea26029a4c086ff07459314de70650/l2geth/core/state_transition.go#L39

#### sequence 2
change the height by state_transition.go

#### create a new release, upgrade the binary for sequencer and verifier
https://github.com/mantlenetworkio/mantle/releases

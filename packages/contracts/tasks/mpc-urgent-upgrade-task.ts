import { task } from 'hardhat/config'
import { ethers } from 'ethers'

// @ts-ignore
import { getContractDefinition, getContractFactory } from '../src'

task('mpcUpdateL1BridgeChugCode')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const provider = new ethers.providers.JsonRpcProvider(
      'https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28'
    )
    //0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e is the key of addressmananger
    //address: 0xd5add52d36399570e56c183d949da83ac29aa7d6
    //this is a account of goerli-network, we will use the gnosis-safe on mainnet
    const ownerWallet = new ethers.Wallet(
      '0x574108d5a6bfa179e9727887e315c1cd08ec2f8ca09dd4f00b1abd591011375f',
      provider
    )
    const l1ChugSplashProxy = getContractFactory('L1ChugSplashProxy').attach(
      taskArgs.contract
    )
    console.log('Query before setCode')
    console.log(
      'Implementation :',
      await l1ChugSplashProxy
        .connect(accounts[0].provider)
        .callStatic.getImplementation({
          from: ethers.constants.AddressZero,
        })
    )

    const upgrade = getContractDefinition('L1StandardBridgeUpdate')

    console.log('set Code:')
    console.log(upgrade.deployedBytecode)


    // const upgradeContract = getContractFactory(
    //   'L1StandardBridgeUpdate'
    // ).attach(taskArgs.contract)
    // await l1ChugSplashProxy
    //   .connect(ownerWallet)
    //   .setCode(upgrade.deployedBytecode)
    //
    // console.log('Query after setCode')
    // console.log(
    //   'Implementation :',
    //   await l1ChugSplashProxy
    //     .connect(accounts[0].provider)
    //     .callStatic.getImplementation({
    //       from: ethers.constants.AddressZero,
    //     })
    // )

    // console.log(
    //   'version: ',
    //   await upgradeContract.connect(accounts[0]).getVersion()
    // )
    // await upgradeContract.connect(ownerWallet).updateL1TokenBridge(taskArgs.contract)
    // console.log(await upgradeContract.connect(accounts[0]).l1TokenBridge())
  })


task('mpcUpdateL1CrossDomainMessenger')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const provider = new ethers.providers.JsonRpcProvider(
      'https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28'
    )

    //0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e is the key of addressmananger
    //address: 0xd5add52d36399570e56c183d949da83ac29aa7d6
    //this is a account of goerli-network, we will use the gnosis-safe on mainnet
    const ownerWallet = new ethers.Wallet(
      '0x574108d5a6bfa179e9727887e315c1cd08ec2f8ca09dd4f00b1abd591011375f',
      provider
    )


    const l1CrossDomainMessenger = getContractFactory('L1CrossDomainMessengerUpgrade')
    const L1CrossDomainMessenger = await l1CrossDomainMessenger
      .connect(ownerWallet)
      .deploy()
    await L1CrossDomainMessenger.deployed()
    console.log('L1CrossDomainMessenger : ', L1CrossDomainMessenger.address)

    // const libAddressManager = getContractFactory('Lib_AddressManager').attach(taskArgs.contract);
    //
    // await libAddressManager.connect(ownerWallet)
    //   .setAddress('BVM_L1CrossDomainMessenger',L1CrossDomainMessenger.address)
  })

task('mpcUpdateTssGroupManager')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const provider = new ethers.providers.JsonRpcProvider(
      'https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28'
    )

    //0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e is the key of addressmananger
    //address: 0xd5add52d36399570e56c183d949da83ac29aa7d6
    //this is a account of goerli-network, we will use the gnosis-safe on mainnet
    const ownerWallet = new ethers.Wallet(
      '0x574108d5a6bfa179e9727887e315c1cd08ec2f8ca09dd4f00b1abd591011375f',
      provider
    )

    // const l1TransparentProxy = getContractFactory('TransparentUpgradeableProxy').attach(taskArgs.contract)
    //
    // console.log('Query before setImplementation')
    // console.log(
    //   'Implementation :',
    //   await l1TransparentProxy.connect(ownerWallet).callStatic.implementation()
    // )

    const tssGroupManagerUpgrade = getContractFactory('TssGroupManagerUpgrade')
    const TssGroupManagerUpgrade = await tssGroupManagerUpgrade
      .connect(ownerWallet)
      .deploy()
    await TssGroupManagerUpgrade.deployed()
    console.log('TssGroupManagerUpgrade : ', TssGroupManagerUpgrade.address)

    // await l1TransparentProxy.connect(ownerWallet).upgradeTo(TssGroupManagerUpgrade.address)
    // console.log('Upgrade success，new impl : ', await l1TransparentProxy
    //   .connect(ownerWallet)
    //   .callStatic.implementation())
  })



task('mpcUpdateTssStakingSlashCode')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const provider = new ethers.providers.JsonRpcProvider(
      'https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28'
    )

    //0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e is the key of addressmananger
    //address: 0xd5add52d36399570e56c183d949da83ac29aa7d6
    //this is a account of goerli-network, we will use the gnosis-safe on mainnet
    const ownerWallet = new ethers.Wallet(
      '0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e',
      provider
    )

    const l1TransparentProxy = getContractFactory('TransparentUpgradeableProxy').attach(taskArgs.contract)

    console.log('Query before setImplementation')
    console.log(
      'Implementation :',
      await l1TransparentProxy.connect(ownerWallet).callStatic.implementation()
    )

    const tssStakingSlashingUpgrade = getContractFactory('TssStakingSlashingUpgrade')
    const TssStakingSlashingUpgrade = await tssStakingSlashingUpgrade
      .connect(ownerWallet)
      .deploy()
    await TssStakingSlashingUpgrade.deployed()
    console.log('TssStakingSlashingUpgrade : ', TssStakingSlashingUpgrade.address)

    // await l1TransparentProxy.connect(ownerWallet).upgradeTo(TssStakingSlashingUpgrade.address)
    // console.log('Upgrade success，new impl : ', await l1TransparentProxy
    //   .connect(ownerWallet)
    //   .callStatic.implementation())
  })



task('mpcUpdateEigenDataLayrChainCode')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const provider = new ethers.providers.JsonRpcProvider(
      'https://goerli.infura.io/v3/d6167662f2104fbc8d5a947e59dbaa28'
    )

    //0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e is the key of addressmananger
    //address: 0xd5add52d36399570e56c183d949da83ac29aa7d6
    //this is a account of goerli-network, we will use the gnosis-safe on mainnet
    const ownerWallet = new ethers.Wallet(
      '0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e',
      provider
    )

    const l1TransparentProxy = getContractFactory('TransparentUpgradeableProxy').attach(taskArgs.contract)

    console.log('Query before setImplementation')
    console.log(
      'Implementation :',
      await l1TransparentProxy.connect(ownerWallet).callStatic.implementation()
    )

    const bvm_EigenDataLayrChainUpgrade = getContractFactory('BVM_EigenDataLayrChainUpgrade')
    const BVM_EigenDataLayrChainUpgrade = await bvm_EigenDataLayrChainUpgrade
      .connect(ownerWallet)
      .deploy()
    await BVM_EigenDataLayrChainUpgrade.deployed()
    console.log('BVM_EigenDataLayrChainUpgrade : ', BVM_EigenDataLayrChainUpgrade.address)

    // await l1TransparentProxy.connect(ownerWallet).upgradeTo(BVM_EigenDataLayrChainUpgrade.address)
    // console.log('Upgrade success，new impl : ', await l1TransparentProxy
    //   .connect(ownerWallet)
    //   .callStatic.implementation())
  })




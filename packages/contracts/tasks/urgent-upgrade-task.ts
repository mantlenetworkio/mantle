import { task } from 'hardhat/config'
import { ethers } from 'ethers'

// @ts-ignore
import { getContractDefinition, getContractFactory } from '../src'

task('setL1BridgeChugCode')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const provider = new ethers.providers.JsonRpcProvider(
      'http://localhost:9545'
    )
    //0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e is the key of addressmananger
    //address: 0xd5add52d36399570e56c183d949da83ac29aa7d6
    const ownerWallet = new ethers.Wallet(
      '0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e',
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

    const upgrade = getContractDefinition('L1StandardBridgeUpgrade')
    const upgradeContract = getContractFactory(
      'L1StandardBridgeUpgrade'
    ).attach(taskArgs.contract)
    console.log('set Code')

    await l1ChugSplashProxy
      .connect(ownerWallet)
      .setCode(upgrade.deployedBytecode)

    console.log('Query after setCode')
    console.log(
      'Implementation :',
      await l1ChugSplashProxy
        .connect(accounts[0].provider)
        .callStatic.getImplementation({
          from: ethers.constants.AddressZero,
        })
    )

    console.log(
      'version: ',
      await upgradeContract.connect(accounts[0]).getVersion()
    )
  })

task('updateTssGroupManagerCode')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const provider = new ethers.providers.JsonRpcProvider(
      'http://localhost:9545'
    )

    //0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e is the key of addressmananger
    //address: 0xd5add52d36399570e56c183d949da83ac29aa7d6
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

    const tssGroupManagerUpgrade = getContractFactory('TssGroupManagerUpgrade')
    const TssGroupManagerUpgrade = await tssGroupManagerUpgrade
      .connect(ownerWallet)
      .deploy()
    await TssGroupManagerUpgrade.deployed()
    console.log('TssGroupManagerUpgrade : ', TssGroupManagerUpgrade.address)

    await l1TransparentProxy.connect(ownerWallet).upgradeTo(TssGroupManagerUpgrade.address)
    console.log('Upgrade success，new impl : ', await l1TransparentProxy
      .connect(ownerWallet)
      .callStatic.implementation())
  })



task('updateTssStakingSlashCode')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const provider = new ethers.providers.JsonRpcProvider(
      'http://localhost:9545'
    )

    //0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e is the key of addressmananger
    //address: 0xd5add52d36399570e56c183d949da83ac29aa7d6
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

    await l1TransparentProxy.connect(ownerWallet).upgradeTo(TssStakingSlashingUpgrade.address)
    console.log('Upgrade success，new impl : ', await l1TransparentProxy
      .connect(ownerWallet)
      .callStatic.implementation())
  })



task('updateEigenDataLayrChainCode')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const provider = new ethers.providers.JsonRpcProvider(
      'http://localhost:9545'
    )

    //0x26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e is the key of addressmananger
    //address: 0xd5add52d36399570e56c183d949da83ac29aa7d6
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

    await l1TransparentProxy.connect(ownerWallet).upgradeTo(BVM_EigenDataLayrChainUpgrade.address)
    console.log('Upgrade success，new impl : ', await l1TransparentProxy
      .connect(ownerWallet)
      .callStatic.implementation())
  })


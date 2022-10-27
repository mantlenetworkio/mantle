import { task } from 'hardhat/config'
import { ethers } from 'ethers'

import { getContractFactory } from '../src/contract-defs'
import { getContractArtifact } from '../src/contract-artifacts'

task('getOwner')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const l1ChugSplashProxy = getContractFactory('L1ChugSplashProxy').attach(
      taskArgs.contract
    )
    const currentOwner = await l1ChugSplashProxy
      .connect(accounts[0].provider)
      .callStatic.getOwner({
        from: ethers.constants.AddressZero,
      })
    console.log('currentOwner: ', currentOwner)
  })

task('getImp')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const l1ChugSplashProxy = getContractFactory('L1ChugSplashProxy').attach(
      taskArgs.contract
    )
    console.log(
      'Implementation :',
      await l1ChugSplashProxy
        .connect(accounts[0].provider)
        .callStatic.getImplementation({
          from: ethers.constants.AddressZero,
        })
    )
  })

task('setCode')
  .addParam('contract', 'proxy address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const ownerWallet = new ethers.Wallet(
      '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
      accounts[0].provider
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

    const upgrade = getContractArtifact('L1StandardBridgeUpgrade')
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

task('upgradeTest').setAction(async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners()

  const ownerWallet = new ethers.Wallet(
    '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
    accounts[0].provider
  )
  const l1ChugSplashProxy = getContractFactory('L1ChugSplashProxy')
  const proxy = await l1ChugSplashProxy
    .connect(ownerWallet)
    .deploy(ownerWallet.address)

  const testContract = getContractArtifact('Test')
  await proxy.connect(ownerWallet).setCode(testContract.deployedBytecode)

  console.log(proxy.address)
  const implementation = await proxy
    .connect(accounts[0].provider)
    .callStatic.getImplementation({
      from: ethers.constants.AddressZero,
    })
  console.log(implementation)

  const test = getContractFactory('Test').attach(proxy.address)
  await test.connect(accounts[0]).setVersion()
  console.log(await test.connect(accounts[0]).version())

  const testUpgradeContract = getContractArtifact('TestUpgrade')
  await proxy.connect(ownerWallet).setCode(testUpgradeContract.deployedBytecode)
  const testUpgrade = getContractFactory('Test').attach(proxy.address)
  await testUpgrade.connect(accounts[0]).setVersion()
  console.log(await testUpgrade.connect(accounts[0]).version())
})

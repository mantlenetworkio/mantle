import { task } from 'hardhat/config'
import { ethers } from 'ethers'

// @ts-ignore
import { getContractDefinition, getContractFactory } from '../src'

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

task('splashProxyUpgradeTest').setAction(async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners()
  const ownerWallet = new ethers.Wallet(
    '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
    accounts[0].provider
  )

  const l1ChugSplashProxy = getContractFactory('L1ChugSplashProxy')
  const proxy = await l1ChugSplashProxy
    .connect(ownerWallet)
    .deploy(ownerWallet.address)

  const testContract = getContractDefinition('Test')
  await proxy.connect(ownerWallet).setCode(testContract.deployedBytecode)

  const implementation = await proxy
    .connect(accounts[0].provider)
    .callStatic.getImplementation({
      from: ethers.constants.AddressZero,
    })
  console.log(implementation)

  const test = getContractFactory('Test').attach(proxy.address)
  await test.connect(accounts[0]).setVersion()
  console.log('Test version', await test.connect(accounts[0]).version())

  const testUpgradeContract = getContractDefinition('TestUpgrade')
  await proxy.connect(ownerWallet).setCode(testUpgradeContract.deployedBytecode)

  const testUpgrade = getContractFactory('TestUpgrade').attach(proxy.address)
  await testUpgrade.connect(accounts[0]).setVersion()
  await testUpgrade.connect(accounts[0]).setCheckNum(11)
  console.log(
    'Check number:',
    await testUpgrade.connect(accounts[0]).checkNum()
  )
  console.log('Test version:', await testUpgrade.connect(accounts[0]).version())
})

task('transparentUpgradea').setAction(async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners()
  const ownerWallet = new ethers.Wallet(
    '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
    accounts[0].provider
  )
  const testContract = getContractFactory('Test')
  const TestContract = await testContract.connect(ownerWallet).deploy()
  await TestContract.deployed()
  console.log('TestContract : ', TestContract.address)

  const callData = testContract.interface.encodeFunctionData('initialize', [])

  // TransparentUpgradeableProxy
  const l1ChugSplashProxy = getContractFactory('TransparentUpgradeableProxy')
  const proxy = await l1ChugSplashProxy
    .connect(ownerWallet)
    .deploy(TestContract.address, ownerWallet.address, callData)

  let implementation = await proxy
    .connect(accounts[0].provider)
    .callStatic.implementation();
  console.log(implementation)

  const test = getContractFactory('Test').attach(proxy.address)
  await test.connect(accounts[1]).setVersion()
  console.log('Test version', await test.connect(accounts[1]).version())

  const testUpgradeContract = getContractFactory('TestUpgrade')
  const TestUpgradeContract = await testUpgradeContract
    .connect(ownerWallet)
    .deploy()
  await TestUpgradeContract.deployed()
  console.log('TestUpgradeContract : ', TestUpgradeContract.address)

  await proxy.connect(ownerWallet).upgradeTo(TestUpgradeContract.address)
  implementation = await proxy
    .connect(accounts[0].provider)
    .callStatic.implementation()
  console.log('Upgrade success，new impl : ', implementation)

  const testUpgrade = getContractFactory('TestUpgrade').attach(proxy.address)
  await testUpgrade.connect(accounts[1]).setVersion()
  await testUpgrade.connect(accounts[1]).setCheckNum(11)
  console.log(
    'Check number:',
    await testUpgrade.connect(accounts[1]).checkNum()
  )
  console.log('Test version:', await testUpgrade.connect(accounts[1]).version())
})

import {task} from "hardhat/config";


task('deploySequencer')
  .addParam('l1bitaddress')
  .setAction(async (taskArgs, hre) => {
    const sequencerFactory = await hre.ethers.getContractFactory('Sequencer')
    const sequencerProxy = await hre.upgrades.deployProxy(
      sequencerFactory,
      [
        taskArgs.l1bitaddress,
      ]
    )
    await sequencerProxy.deployed()
    const proxyAdmin = await hre.upgrades.admin.getInstance()
    const impAddr = await proxyAdmin.getProxyImplementation(sequencerProxy.address)
    console.log('sequencer proxy address :', sequencerProxy.address.toLocaleLowerCase())
    console.log('sequencer address :', impAddr)
    console.log('sequencer contract owner :', await sequencerProxy.owner())
    console.log('sequencer bit address :', await sequencerProxy.bitToken())
  })

task('upgradeSequencer')
  .addParam('sequencerproxy')
  .setAction(async (taskArgs, hre) => {
    const sequencerFactory = await hre.ethers.getContractFactory('Sequencer')
    const sequencerProxy = await hre.upgrades.upgradeProxy(taskArgs.sequencerproxy,sequencerFactory)
    await sequencerProxy.deployed()
    const proxyAdmin = await hre.upgrades.admin.getInstance()
    const impAddr = await proxyAdmin.getProxyImplementation(sequencerProxy.address)
    console.log('sequencer proxy address:', sequencerProxy.address.toLocaleLowerCase())
    console.log('sequencer proxy owner :', await sequencerProxy.owner())
    console.log('sequencer address:', impAddr)
    console.log('sequencer bit address :', await sequencerProxy.bitToken())
  })

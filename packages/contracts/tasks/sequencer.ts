import {task} from "hardhat/config";


task('deploySequencer')
  .addParam('l1bitaddress')
  .setAction(async (taskArgs, hre) => {
    const sequencerFactory = await hre.ethers.getContractFactory('Sequencer')
    const sequencer = await hre.upgrades.deployProxy(
      sequencerFactory,
      [
        taskArgs.l1bitaddress,
      ]
    )
    await sequencer.deployed()
    console.log('sequencer_proxy_address:', sequencer.address.toLocaleLowerCase())
    console.log('sequencer contract owner :', await sequencer.owner())
    console.log('sequencer bit address :', await sequencer.bitToken())
  })

task("updateScheduler")
  .addParam("sequencer")
  .addParam("scheduler")
  .setAction(async (taskArgs, hre) => {
    const sequencerFactory = await hre.ethers.getContractFactory('Sequencer')
    const sequencer = sequencerFactory.attach(taskArgs.sequencer)
    await sequencer.updateScheduler(taskArgs.scheduler)
    console.log(await sequencer.scheduler())
  })

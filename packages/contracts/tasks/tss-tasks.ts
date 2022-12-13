import {task} from 'hardhat/config'
import {getContractFactory} from '../src/contract-defs'
import {HexToBytes} from "../src/deploy-utils";
import {utils} from "ethers";


task("setTssGroupMember")
  .addParam('contract', "tss group contract address")
  .addParam('threshold', "tss threshold")
  .addParam('batchpublickey', "tss group batch publicKey")
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const tssGroupManager = getContractFactory('TssGroupManager').attach(taskArgs.contract)
    const thresholdP = taskArgs.threshold

    const batchPublicKey = []
    //support parse special array data for batchpublickey
    let re = /\[/gi
    taskArgs.batchpublickey = taskArgs.batchpublickey.replace(re, '')
    re = /\]/gi
    taskArgs.batchpublickey = taskArgs.batchpublickey.replace(re, '')
    taskArgs.batchpublickey = taskArgs.batchpublickey.split(',')
    for (const pk of taskArgs.batchpublickey) {
      const ret = await HexToBytes(pk)
      batchPublicKey.push(ret)
    }
    await tssGroupManager.connect(accounts[0]).setTssGroupMember(thresholdP, batchPublicKey)
  })

task("setSlashingParams")
  .addParam('contract', "tss staking slashing contract address")
  .addParam('slashamount0', "slash uptime amount")
  .addParam('slashamount1', "slash animus amount")
  .addParam('exincome0', "slash uptime extra income")
  .addParam('exincome1', "slash animus extra income")
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const tssStakingSlashing = getContractFactory('TssStakingSlashing').attach(taskArgs.contract)
    let slashParams = [taskArgs.slashamount0, taskArgs.slashamount1]
    let exIncomes = [taskArgs.exincome0, taskArgs.exincome1]
    await tssStakingSlashing.connect(accounts[0]).setSlashingParams(slashParams, exIncomes)
  })

task('getSlashingParams')
  .addParam('contract', 'tss staking slashing contract address')
  .setAction(async (taskArgs, hre) => {
    const tssStakingSlashing = getContractFactory('TssStakingSlashing').attach(taskArgs.contract)
    const params = await tssStakingSlashing.connect(hre.ethers.provider).getSlashingParams()
    if (params.length === 0) {
      console.log('slashing params has not been set')
      return
    }
    console.log('minimum deposit amount required: ', params[0][1].toString())
    console.log('liveness slashing amount: ', params[0][0].toString())
    console.log('animus slashing amount: ', params[0][1].toString())
    console.log(
      'additional rewards for sender by liveness case: ',
      params[1][0].toString()
    )
    console.log(
      'additional rewards for sender by animus case: ',
      params[1][1].toString()
    )
  })

task("clearQuitRequestList")
  .addParam('contract', "tss staking slashing contract address")
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const tssStakingSlashing = getContractFactory('TssStakingSlashing').attach(taskArgs.contract)
    await tssStakingSlashing.connect(accounts[0]).clearQuitRequestList()
  })

task("staking")
  .addParam('contract', "tss staking slashing contract address")
  .addParam('amount')
  .addParam('pubkey')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const tssStakingSlashing = getContractFactory('TssStakingSlashing').attach(taskArgs.contract)
    await tssStakingSlashing.connect(accounts[0]).staking(taskArgs.amount, taskArgs.pubkey)
  })

task('withdrawToken')
  .addParam('contract', 'tss staking slashing contract address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const tssStakingSlashing = getContractFactory('TssStakingSlashing').attach(
      taskArgs.contract
    )
    await tssStakingSlashing.connect(accounts[0]).withdrawToken()
  })

task('unjail')
  .addParam('contract', 'tss staking slashing contract address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const tssStakingSlashing = getContractFactory('TssStakingSlashing').attach(
      taskArgs.contract
    )
    await tssStakingSlashing.connect(accounts[0]).unJail()
  })

task(`checkIsTssMember`)
  .addParam('contract', 'tss group manager contract address')
  .addParam('pubkey')
  .setAction(async (taskArgs, hre) => {
    const tssGroupManager = getContractFactory('TssGroupManager')
      .attach(taskArgs.contract)
    const inactiveInfo = await tssGroupManager
      .connect(hre.ethers.provider)
      .getTssInactiveGroupInfo()
    const inactiveMembers = inactiveInfo[3]
    let found = false
    for (const i in inactiveMembers) {
      if (inactiveMembers[i] === taskArgs.pubkey) {
        found = true
        break
      }
    }
    if (found) {
      console.log('true')
      return
    }

    const activeInfo = await tssGroupManager
      .connect(hre.ethers.provider)
      .getTssGroupInfo()
    const activeMembers = activeInfo[3]
    for (const i in activeMembers) {
      if (activeMembers[i] === taskArgs.pubkey) {
        found = true
        break
      }
    }
    console.log(found)
  })

task(`getTssActiveGroupInfo`)
  .addParam('contract', 'tss group manager contract address')
  .setAction(async (taskArgs, hre) => {
    const tssGroupManager = getContractFactory('TssGroupManager')
      .attach(taskArgs.contract)
    const info = await tssGroupManager
      .connect(hre.ethers.provider)
      .getTssGroupInfo()
    console.log(info)
  })

task(`getTssInactiveGroupInfo`)
  .addParam('contract', 'tss group manager contract address')
  .setAction(async (taskArgs, hre) => {
    const tssGroupManager = getContractFactory('TssGroupManager')
      .attach(taskArgs.contract)
    const info = await tssGroupManager
      .connect(hre.ethers.provider)
      .getTssInactiveGroupInfo()
    console.log(info)
  })

task(`getTssGroupUnJailMembers`)
  .addParam(`contract`)
  .setAction(async (taskArgs, hre) => {
    const tssGroupManager = getContractFactory('TssGroupManager')
      .attach(taskArgs.contract)
    const addresses = await tssGroupManager
      .connect(hre.ethers.provider)
      .getTssGroupUnJailMembers()
    console.log(addresses)
  })

task(`quitRequest`)
  .addParam('contract', 'tss staking slashing contract address')
  .setAction(async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()
    const tssStakingSlashing = getContractFactory('TssStakingSlashing')
      .attach(taskArgs.contract)
    const response = await tssStakingSlashing
      .connect(accounts[0])
      .quitRequest()
    console.log(response)
  })

task(`getQuitRequestList`)
  .addParam('contract', 'tss staking slashing contract address')
  .setAction(async (taskArgs, hre) => {
    const tssStakingSlashing = getContractFactory('TssStakingSlashing')
      .attach(taskArgs.contract)
    const response = await tssStakingSlashing
      .connect(hre.ethers.provider)
      .getQuitRequestList()
    console.log(response)
  })

task(`upgradeTssManager`)
  .addParam('proxy', '')
  .setAction(async (taskArgs, hre) => {
    console.log('Upgrading TssGroupManager')
    const accounts = await hre.ethers.getSigners()
    const NewTssManager = await hre.ethers.getContractFactory('TssGroupManager')
    const newTssManager = await NewTssManager.deploy()
    await newTssManager.deployed()
    console.log('newTssManager deployed to: ', newTssManager.address)

    const proxy = await getContractFactory(
      'TransparentUpgradeableProxy'
    ).attach(taskArgs.proxy)
    await proxy.connect(accounts[1]).upgradeTo(newTssManager.address)
    console.log('TssGroupManager Upgraded')

  })

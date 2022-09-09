import {task} from 'hardhat/config'
import {getContractFactory} from '../src/contract-defs'
import {HexToBytes} from "../src/deploy-utils";
import {ethers} from "ethers";


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
  .addParam('prikey')
  .setAction(async (taskArgs, hre) => {
    const signer = new hre.ethers.Wallet(
      taskArgs.prikey,
      hre.ethers.provider
    )
    const tssStakingSlashing = getContractFactory('TssStakingSlashing').attach(taskArgs.contract)
    await tssStakingSlashing.connect(signer).staking(taskArgs.amount, taskArgs.pubkey)
  })

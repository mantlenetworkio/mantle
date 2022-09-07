import { task } from 'hardhat/config'
import { getContractFactory } from '../src/contract-defs'

task("setSlashingParams")
    .addParam('contract', "tss staking slashing contract address")
    .addParam('slashamount0', "slash uptime amount")
    .addParam('slashamount1', "slash animus amount")
    .addParam('exincome0', "slash uptime extra income")
    .addParam('exincome1', "slash animus extra income")
    .setAction(async (taskArgs, hre) => {
        const tssStakingSlashing = getContractFactory('TssStakingSlashing').attach(taskArgs.contract)
        let slashParams = [taskArgs.slashamount0, taskArgs.slashamount1]
        let exIncomes = [taskArgs.exincome0, taskArgs.exincome1]
        await tssStakingSlashing.setSlashingParams(slashParams, exIncomes)
    })

task("clearQuitRequestList")
    .addParam('contract', "tss staking slashing contract address")
    .setAction(async (taskArgs, hre) => {
        const tssStakingSlashing = getContractFactory('TssStakingSlashing').attach(taskArgs.contract)
        await tssStakingSlashing.clearQuitRequestList()
    })

task("staking")
    .addParam('contract', "tss staking slashing contract address")
    .addParam('amount')
    .addParam('pubkey')
    .setAction(async (taskArgs, hre) => {
        const tssStakingSlashing = getContractFactory('TssStakingSlashing').attach(taskArgs.contract)
        await tssStakingSlashing.staking(taskArgs.amount, taskArgs.pubkey)
    })

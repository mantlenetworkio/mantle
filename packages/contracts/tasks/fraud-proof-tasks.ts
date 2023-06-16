import { task } from 'hardhat/config'
import { ethers } from 'ethers'
import { hexStringEquals } from '@mantleio/core-utils'

// @ts-ignore
import { getContractFactory } from '../src'
// @ts-ignore
import { names } from '../src/address-names'

// eslint-disable-next-line @typescript-eslint/no-var-requires
const fs = require('fs')

task('setAddress')
  .addParam('addressmanager', 'Lib_AddressManager contract address')
  .addParam('name', 'contract name')
  .addParam('address', 'contract address')
  .setAction(async (taskArgs) => {
    const provider = new ethers.providers.JsonRpcProvider(
      process.env.CONTRACTS_RPC_URL
    )
    const addressManagerKey = process.env.BVM_ADDRESS_MANAGER_KEY
    const managerWallet = new ethers.Wallet(addressManagerKey, provider)
    const lib_AddressManager = await getContractFactory(
      'Lib_AddressManager'
    ).attach(taskArgs.addressmanager)

    console.log(
      'Lib_AddressManager owner:',
      await lib_AddressManager.connect(managerWallet).owner()
    )
    console.log(
      'The name has address before set',
      await lib_AddressManager.connect(managerWallet).getAddress(taskArgs.name)
    )
    await lib_AddressManager
      .connect(managerWallet)
      .setAddress(taskArgs.name, taskArgs.address)
    console.log(
      'The name has address after set',
      await lib_AddressManager.connect(managerWallet).getAddress(taskArgs.name)
    )
  })

task('whiteListInit')
  .addParam('rollup', 'Rollup contract address')
  .addParam('stakers', 'Rollup contract stakers address')
  .setAction(async (taskArgs) => {
    const provider = new ethers.providers.JsonRpcProvider(
      process.env.CONTRACTS_RPC_URL
    )
    const deployerKey = process.env.CONTRACTS_DEPLOYER_KEY
    const proposerAddr = process.env.BVM_ROLLUPER_ADDRESS
    const entryOwner = new ethers.Wallet(deployerKey, provider)

    console.log(
      'balance: ',
      entryOwner.address,
      (await entryOwner.getBalance()).toString()
    )
    const SequencerENV = process.env.BVM_ROLLUPER_ADDRESS
    const Validator1ENV = process.env.BVM_VERIFIER1_ADDRESS
    const operatorWhitelist = [SequencerENV, Validator1ENV]
    console.log('operatorWhitelist:', operatorWhitelist)

    const StakerWhitelist = taskArgs.stakers.split(',')
    console.log('StakerWhitelist:', StakerWhitelist)

    const rollup = await getContractFactory('Rollup').attach(taskArgs.rollup)
    await rollup.connect(entryOwner).addToOperatorWhitelist(operatorWhitelist)
    await rollup.connect(entryOwner).addToStakerWhitelist(StakerWhitelist)

    console.log('transferOwnerShip')
    await rollup.connect(entryOwner).transferOwnership(proposerAddr)
  })

task('rollupStake')
  .addParam('rollup', 'Rollup contract address')
  .addParam('stakerkeys', 'Rollup contract stakers address key')
  .addParam('amount', 'amount to stake', '0.1')
  .setAction(async (taskArgs) => {
    const provider = new ethers.providers.JsonRpcProvider(
      process.env.CONTRACTS_RPC_URL
    )
    const mantleToken = process.env.L1_MANTLE_ADDRESS
    const mantle = await getContractFactory('L1MantleToken').attach(mantleToken)
    const rollup = await getContractFactory('Rollup').attach(taskArgs.rollup)

    const SequencerAddress = process.env.BVM_ROLLUPER_ADDRESS
    const Validator1Address = process.env.BVM_VERIFIER1_ADDRESS
    const stakerKeys = taskArgs.stakerkeys.split(',')
    const operators = [SequencerAddress, Validator1Address]

    const deployerKey = process.env.CONTRACTS_DEPLOYER_KEY
    const deployer = new ethers.Wallet(deployerKey, provider)
    const amount = ethers.utils.parseEther('1.0') // 替换为转账金额
    const mantleAmount = ethers.utils.parseEther('100.0') // 替换为转账金额

    // 使用索引的 for 循环迭代数组
    for (let i = 0; i < stakerKeys.length; i++) {
      const stakerWallet = new ethers.Wallet(stakerKeys[i], provider)
      // 构造交易对象
      const transaction = {
        to: stakerWallet.address,
        value: amount,
      }

      try {
        const response = await deployer.sendTransaction(transaction)
        console.log('Transaction hash:', response.hash)
      } catch (error) {
        console.error('Failed to send transaction:', error)
      }

      await mantle.connect(deployer).transfer(stakerWallet.address, mantleAmount)
      console.log(
        'balance: ',
        stakerWallet.address,
        (await mantle.connect(deployer).balanceOf(stakerWallet.address)).toString()
      )

      console.log(
        'ETH Balance:',
        stakerWallet.address,
        ' ',
        await stakerWallet.getBalance()
      )
      await mantle
        .connect(stakerWallet)
        .approve(taskArgs.rollup, ethers.utils.parseEther(taskArgs.amount))
      console.log(
        'ETH Balance:',
        stakerWallet.address,
        ' ',
        await stakerWallet.getBalance()
      )

      console.log('stake', stakerWallet.address, operators[i])
      await rollup
        .connect(stakerWallet)
        .stake(ethers.utils.parseEther(taskArgs.amount), operators[i])
    }
  })

task(`deployVerifier`)
  .addParam('verifier', 'verifier entry address')
  .setAction(async (taskArgs) => {
    const provider = new ethers.providers.JsonRpcProvider(
      process.env.CONTRACTS_RPC_URL
    )
    const deployerKey = process.env.CONTRACTS_DEPLOYER_KEY
    const entryOwner = new ethers.Wallet(deployerKey, provider)
    const BlockInitiationVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.BlockInitiationVerifier
    )
    const blockInitiationVerifier = await BlockInitiationVerifier.connect(
      entryOwner
    ).deploy()
    await blockInitiationVerifier.deployed()
    console.log('blockInitiationVerifier : ', blockInitiationVerifier.address)

    const BlockFinalizationVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.BlockFinalizationVerifier
    )
    const blockFinalizationVerifier = await BlockFinalizationVerifier.connect(
      entryOwner
    ).deploy()
    await blockFinalizationVerifier.deployed()
    console.log(
      'blockFinalizationVerifier : ',
      blockFinalizationVerifier.address
    )

    const InterTxVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.BlockInitiationVerifier
    )
    const interTxVerifier = await InterTxVerifier.connect(entryOwner).deploy()
    await interTxVerifier.deployed()
    console.log('interTxVerifier : ', interTxVerifier.address)

    const StackOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.StackOpVerifier
    )
    const stackOpVerifier = await StackOpVerifier.connect(entryOwner).deploy()
    await stackOpVerifier.deployed()
    console.log('stackOpVerifier : ', stackOpVerifier.address)

    const EnvironmentalOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.EnvironmentalOpVerifier
    )
    const environmentalOpVerifier = await EnvironmentalOpVerifier.connect(
      entryOwner
    ).deploy()
    await environmentalOpVerifier.deployed()
    console.log('environmentalOpVerifier : ', environmentalOpVerifier.address)

    const MemoryOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.MemoryOpVerifier
    )
    const memoryOpVerifier = await MemoryOpVerifier.connect(entryOwner).deploy()
    await memoryOpVerifier.deployed()
    console.log('memoryOpVerifier : ', memoryOpVerifier.address)

    const StorageOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.StorageOpVerifier
    )
    const storageOpVerifier = await StorageOpVerifier.connect(
      entryOwner
    ).deploy()
    await storageOpVerifier.deployed()
    console.log('storageOpVerifier : ', storageOpVerifier.address)

    const CallOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.CallOpVerifier
    )
    const callOpVerifier = await CallOpVerifier.connect(entryOwner).deploy()
    await callOpVerifier.deployed()
    console.log('callOpVerifier : ', callOpVerifier.address)

    const InvalidOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.InvalidOpVerifier
    )
    const invalidOpVerifier = await InvalidOpVerifier.connect(
      entryOwner
    ).deploy()
    await invalidOpVerifier.deployed()
    console.log('invalidOpVerifier : ', invalidOpVerifier.address)

    const Proxy__Verifier = getContractFactory(
      names.managed.fraud_proof.VerifierEntry
    ).attach(taskArgs.verifier)
    await Proxy__Verifier.connect(entryOwner).setVerifier(
      0,
      await stackOpVerifier.address
    )
    await Proxy__Verifier.connect(entryOwner).setVerifier(
      1,
      await environmentalOpVerifier.address
    )
    await Proxy__Verifier.connect(entryOwner).setVerifier(
      2,
      await memoryOpVerifier.address
    )
    await Proxy__Verifier.connect(entryOwner).setVerifier(
      3,
      await storageOpVerifier.address
    )
    await Proxy__Verifier.connect(entryOwner).setVerifier(
      4,
      await callOpVerifier.address
    )
    await Proxy__Verifier.connect(entryOwner).setVerifier(
      5,
      await invalidOpVerifier.address
    )
    await Proxy__Verifier.connect(entryOwner).setVerifier(
      6,
      await interTxVerifier.address
    )
    await Proxy__Verifier.connect(entryOwner).setVerifier(
      7,
      await blockInitiationVerifier.address
    )
    await Proxy__Verifier.connect(entryOwner).setVerifier(
      8,
      await blockFinalizationVerifier.address
    )

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(entryOwner).blockInitiationVerifier(),
        blockInitiationVerifier.address
      )
    ) {
      console.log('blockInitiationVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(entryOwner).blockFinalizationVerifier(),
        blockFinalizationVerifier.address
      )
    ) {
      console.log('blockFinalizationVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(entryOwner).interTxVerifier(),
        interTxVerifier.address
      )
    ) {
      console.log('interTxVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(entryOwner).stackOpVerifier(),
        stackOpVerifier.address
      )
    ) {
      console.log('stackOpVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(entryOwner).environmentalOpVerifier(),
        environmentalOpVerifier.address
      )
    ) {
      console.log('environmentalOpVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(entryOwner).memoryOpVerifier(),
        memoryOpVerifier.address
      )
    ) {
      console.log('memoryOpVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(entryOwner).storageOpVerifier(),
        storageOpVerifier.address
      )
    ) {
      console.log('storageOpVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(entryOwner).callOpVerifier(),
        callOpVerifier.address
      )
    ) {
      console.log('callOpVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(entryOwner).invalidOpVerifier(),
        invalidOpVerifier.address
      )
    ) {
      console.log('invalidOpVerifier address not equal')
    }

    const VerifierTestDriver = await getContractFactory('VerifierTestDriver')
    const verifierTestDriver = await VerifierTestDriver.connect(
      entryOwner
    ).deploy(
      blockInitiationVerifier.address,
      blockFinalizationVerifier.address,
      interTxVerifier.address,
      stackOpVerifier.address,
      environmentalOpVerifier.address,
      memoryOpVerifier.address,
      storageOpVerifier.address,
      callOpVerifier.address,
      invalidOpVerifier.address
    )
    await verifierTestDriver.deployed()

    console.log('verifierTestDriver deployed to:', verifierTestDriver.address)
  })

task(`genOsp`)
  .addParam('hash', 'the transaction hash to prove')
  .addParam('step', 'the step to prove')
  .setAction(async (taskArgs) => {
    // TODO change hardcode url to env variable
    const provider = new ethers.providers.JsonRpcProvider(
      'http://localhost:8545'
    )
    const res = await provider.send('debug_generateProofForTest', [
      taskArgs.hash,
      0,
      0,
      parseInt(taskArgs.step, 10),
    ])
    fs.writeFileSync(
      './test/data/json/fraud-proof/osp.json',
      JSON.stringify(res)
    )
    console.log('wrote proof to ./test/data/json/fraud-proof/osp.json')
  })

task(`verifyOsp`)
  // .addParam('addr', 'VerifierTestDriver contract address')
  .setAction(async () => {
    const provider = new ethers.providers.JsonRpcProvider(
      process.env.CONTRACTS_RPC_URL
    )
    const ownerWallet = new ethers.Wallet(
      '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
      provider
    )
    const verifierTestDriver = await getContractFactory(
      'VerifierTestDriver'
    ).attach('0xc6e7DF5E7b4f2A278906862b61205850344D4e7d')
    const { ctx, proof } = JSON.parse(
      fs.readFileSync('./test/data/json/fraud-proof/osp.json')
    )
    console.log(`processing tx ${ctx.txnHash}`)

    const transaction = [
      ctx.txNonce,
      ctx.gasPrice,
      ctx.gas,
      ctx.recipient,
      ctx.value,
      ctx.input,
      ctx.txV,
      ctx.txR,
      ctx.txS,
    ]

    const res = await verifierTestDriver
      .connect(ownerWallet)
      .verifyProof(
        ctx.coinbase,
        ctx.timestamp,
        ctx.blockNumber,
        ctx.origin,
        ctx.txnHash,
        transaction,
        proof.verifier,
        proof.currHash,
        proof.proof
      )

    if (!hexStringEquals(proof.nextHash, res)) {
      console.log(
        'next hash not equal with proof: ',
        proof.nextHash,
        ' res: ',
        res
      )
    } else {
      console.log('verify success')
    }
  })

module.exports = {}

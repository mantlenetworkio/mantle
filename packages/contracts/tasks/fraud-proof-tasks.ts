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
      'http://localhost:9545'
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
  .setAction(async (taskArgs) => {
    const provider = new ethers.providers.JsonRpcProvider(
      'http://localhost:9545'
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
    // const Validator2ENV = process.env.BVM_VERIFIER2_ADDRESS
    const whiteListToAdd = [SequencerENV, Validator1ENV]
    console.log('whiteList:', whiteListToAdd)
    const rollup = await getContractFactory('Rollup').attach(taskArgs.rollup)
    await rollup.connect(entryOwner).addToOperatorWhitelist(whiteListToAdd)
    await rollup.connect(entryOwner).addToStakerWhitelist(whiteListToAdd)
    console.log('transferOwnerShip')
    await rollup.connect(entryOwner).transferOwnership(proposerAddr)
  })

task('rollupStake')
  .addParam('rollup', 'Rollup contract address')
  .addParam('amount', 'amount to stake', '0.1')
  .setAction(async (taskArgs) => {
    const provider = new ethers.providers.JsonRpcProvider(
      'http://localhost:9545'
    )
    const mantleToken = process.env.L1_MANTLE_ADDRESS
    const verifier1Key = process.env.BVM_VERIFIER1_KEY
    const proposerKey = process.env.BVM_PROPOSER_KEY

    const proposerWallet = new ethers.Wallet(proposerKey, provider)
    const verifier1Wallet = new ethers.Wallet(verifier1Key, provider)

    const wallets = [proposerWallet, verifier1Wallet]
    const rollup = await getContractFactory('Rollup').attach(taskArgs.rollup)
    const mantle = await getContractFactory('L1MantleToken').attach(mantleToken)
    for (const w of wallets) {
      console.log("ETH Balance:",w.address," ",await w.getBalance())
      // await mantle.connect(w).mint(ethers.utils.parseEther(taskArgs.amount))
      await mantle
        .connect(w)
        .approve(taskArgs.rollup, ethers.utils.parseEther(taskArgs.amount))
      console.log(
        'balance: ',
        w.address,
        (await mantle.connect(w).balanceOf(w.address)).toString()
      )
      await rollup.connect(w).stake(ethers.utils.parseEther(taskArgs.amount), w.address)
    }
  })

task(`deployVerifier`)
  .addParam('verifier', 'verifier entry address')
  .setAction(async (taskArgs) => {
    const provider = new ethers.providers.JsonRpcProvider(
      'http://localhost:9545'
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
    const memoryOpVerifier = await MemoryOpVerifier.connect(
      entryOwner
    ).deploy()
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
      'http://localhost:9545'
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

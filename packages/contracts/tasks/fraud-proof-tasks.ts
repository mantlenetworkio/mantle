import { task } from 'hardhat/config'
import { ethers } from 'ethers'
import { hexStringEquals } from '@mantleio/core-utils'
import { boolean, int } from 'hardhat/internal/core/params/argumentTypes'

// @ts-ignore
import { getContractFactory } from '../src'
// @ts-ignore
import { names } from '../src/address-names'

// eslint-disable-next-line @typescript-eslint/no-var-requires
const fs = require('fs')

task(`deployVerifier`)
  .addParam('address', 'verifier entry address')
  .setAction(async (taskArgs) => {
    const provider = new ethers.providers.JsonRpcProvider(
      'http://localhost:9545'
    )
    const ownerWallet = new ethers.Wallet(
      '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
      provider
    )

    const BlockInitiationVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.BlockInitiationVerifier
    )
    const blockInitiationVerifier = await BlockInitiationVerifier.connect(
      ownerWallet
    ).deploy()
    await blockInitiationVerifier.deployed()
    console.log('blockInitiationVerifier : ', blockInitiationVerifier.address)

    const BlockFinalizationVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.BlockFinalizationVerifier
    )
    const blockFinalizationVerifier = await BlockFinalizationVerifier.connect(
      ownerWallet
    ).deploy()
    await blockFinalizationVerifier.deployed()
    console.log(
      'blockFinalizationVerifier : ',
      blockFinalizationVerifier.address
    )

    const InterTxVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.BlockInitiationVerifier
    )
    const interTxVerifier = await InterTxVerifier.connect(ownerWallet).deploy()
    await interTxVerifier.deployed()
    console.log('interTxVerifier : ', interTxVerifier.address)

    const StackOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.StackOpVerifier
    )
    const stackOpVerifier = await StackOpVerifier.connect(ownerWallet).deploy()
    await stackOpVerifier.deployed()
    console.log('stackOpVerifier : ', stackOpVerifier.address)

    const EnvironmentalOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.EnvironmentalOpVerifier
    )
    const environmentalOpVerifier = await EnvironmentalOpVerifier.connect(
      ownerWallet
    ).deploy()
    await environmentalOpVerifier.deployed()
    console.log('environmentalOpVerifier : ', environmentalOpVerifier.address)

    const MemoryOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.MemoryOpVerifier
    )
    const memoryOpVerifier = await MemoryOpVerifier.connect(
      ownerWallet
    ).deploy()
    await memoryOpVerifier.deployed()
    console.log('memoryOpVerifier : ', memoryOpVerifier.address)

    const StorageOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.StorageOpVerifier
    )
    const storageOpVerifier = await StorageOpVerifier.connect(
      ownerWallet
    ).deploy()
    await storageOpVerifier.deployed()
    console.log('storageOpVerifier : ', storageOpVerifier.address)

    const CallOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.CallOpVerifier
    )
    const callOpVerifier = await CallOpVerifier.connect(ownerWallet).deploy()
    await callOpVerifier.deployed()
    console.log('callOpVerifier : ', callOpVerifier.address)

    const InvalidOpVerifier = getContractFactory(
      names.managed.fraud_proof.SubVerifiers.InvalidOpVerifier
    )
    const invalidOpVerifier = await InvalidOpVerifier.connect(
      ownerWallet
    ).deploy()
    await invalidOpVerifier.deployed()
    console.log('invalidOpVerifier : ', invalidOpVerifier.address)

    const Proxy__Verifier = getContractFactory(
      names.managed.fraud_proof.VerifierEntry
    ).attach(taskArgs.address)
    await Proxy__Verifier.connect(ownerWallet).setVerifier(
      0,
      await stackOpVerifier.address
    )
    await Proxy__Verifier.connect(ownerWallet).setVerifier(
      1,
      await environmentalOpVerifier.address
    )
    await Proxy__Verifier.connect(ownerWallet).setVerifier(
      2,
      await memoryOpVerifier.address
    )
    await Proxy__Verifier.connect(ownerWallet).setVerifier(
      3,
      await storageOpVerifier.address
    )
    await Proxy__Verifier.connect(ownerWallet).setVerifier(
      4,
      await callOpVerifier.address
    )
    await Proxy__Verifier.connect(ownerWallet).setVerifier(
      5,
      await invalidOpVerifier.address
    )
    await Proxy__Verifier.connect(ownerWallet).setVerifier(
      6,
      await interTxVerifier.address
    )
    await Proxy__Verifier.connect(ownerWallet).setVerifier(
      7,
      await blockInitiationVerifier.address
    )
    await Proxy__Verifier.connect(ownerWallet).setVerifier(
      8,
      await blockFinalizationVerifier.address
    )

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(ownerWallet).blockInitiationVerifier(),
        blockInitiationVerifier.address
      )
    ) {
      console.log('blockInitiationVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(ownerWallet).blockFinalizationVerifier(),
        blockFinalizationVerifier.address
      )
    ) {
      console.log('blockFinalizationVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(ownerWallet).interTxVerifier(),
        interTxVerifier.address
      )
    ) {
      console.log('interTxVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(ownerWallet).stackOpVerifier(),
        stackOpVerifier.address
      )
    ) {
      console.log('stackOpVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(ownerWallet).environmentalOpVerifier(),
        environmentalOpVerifier.address
      )
    ) {
      console.log('environmentalOpVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(ownerWallet).memoryOpVerifier(),
        memoryOpVerifier.address
      )
    ) {
      console.log('memoryOpVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(ownerWallet).storageOpVerifier(),
        storageOpVerifier.address
      )
    ) {
      console.log('storageOpVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(ownerWallet).callOpVerifier(),
        callOpVerifier.address
      )
    ) {
      console.log('callOpVerifier address not equal')
    }

    if (
      !hexStringEquals(
        await Proxy__Verifier.connect(ownerWallet).invalidOpVerifier(),
        invalidOpVerifier.address
      )
    ) {
      console.log('invalidOpVerifier address not equal')
    }

    const VerifierTestDriver = await getContractFactory('VerifierTestDriver')
    const verifierTestDriver = await VerifierTestDriver.connect(
      ownerWallet
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

task(`genOspForStep`)
  .addParam('fraud', '', false, boolean)
  .addParam('hash', 'the transaction hash to prove')
  .addParam('step', 'the step to prove', 1, int)
  .setAction(async (taskArgs) => {
    const provider = new ethers.providers.JsonRpcProvider(
      'http://localhost:8545'
    )
    const res = await provider.send('debug_generateProofForStep', [
      taskArgs.fraud,
      taskArgs.hash,
      taskArgs.step,
    ])
    fs.writeFileSync(
      './test/data/json/fraud-proof/ospStep.json',
      JSON.stringify(res)
    )
    console.log('wrote proof to ./test/data/json/fraud-proof/ospStep.json')
  })

task(`genOspForOpcode`)
  .addParam('fraud', '', false, boolean)
  .addParam('hash', 'the transaction hash to prove')
  .addParam('opcode', 'the op code to prove', 22, int)
  .setAction(async (taskArgs) => {
    const provider = new ethers.providers.JsonRpcProvider(
      'http://localhost:8545'
    )
    const res = await provider.send('debug_generateProofForOpcode', [
      taskArgs.fraud,
      taskArgs.hash,
      taskArgs.opcode,
    ])
    fs.writeFileSync(
      './test/data/json/fraud-proof/ospOp.json',
      JSON.stringify(res)
    )
    console.log('wrote proof to ./test/data/json/fraud-proof/ospOp.json')
  })

task(`verifyOsp`)
  // .addParam('addr', 'VerifierTestDriver contract address')
  .setAction(async (taskArgs) => {
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

    const path = []
    if (fs.existsSync('./test/data/json/fraud-proof/ospStep.json')) {
      path.push('./test/data/json/fraud-proof/ospStep.json')
    }
    if (fs.existsSync('./test/data/json/fraud-proof/ospOp.json')) {
      path.push('./test/data/json/fraud-proof/ospOp.json')
    }

    for (const item of path) {
      const { ctx, proof } = JSON.parse(fs.readFileSync(item))
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
    }
  })

task(`testOpcode`)
  .addParam('fraud', '', false, boolean)
  .setAction(async (taskArgs) => {
    const providerL1 = new ethers.providers.JsonRpcProvider(
      'http://localhost:9545'
    )
    const providerL2 = new ethers.providers.JsonRpcProvider(
      'http://localhost:8545'
    )
    const ownerWalletL1 = new ethers.Wallet(
      '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
      providerL1
    )
    const ownerWalletL2 = new ethers.Wallet(
      '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
      providerL2
    )
    const tokenFactory = getContractFactory('TestERC20')
    const token = await tokenFactory
      .connect(ownerWalletL2)
      .deploy()
    await token.deployed()
    console.log('token : ', token.address)
    console.log('deploy tx : ', token.deployTransaction.hash)
    const TestFactory = getContractFactory('TestOSP')
    const test = await TestFactory.connect(ownerWalletL2).deploy(token.address)
    await test.deployed()
    const mintRes = await test.newToken({ gasLimit: 1000000 })
    console.log('mint tx', mintRes.hash)

    const opcode = [
      1, // opcode: ADD stakeOp verifier 0
      52, // opcode: CALLVALUE EnvironmentalOp verifier 1
      // todo : cannot verify
      // 82, // opcode: MSTORE MemoryOp verifier 2
      // 84, // opcode: SLOAD StorageOp verifier 3
      // todo : cannot generate proof
      // 241, // opcode: CALL CallOp verifier 4
      // 240,
      256, // opcode: Invalid InvalidOp verifier 5
    ]
    const osp = []
    for (const item of opcode) {
      console.log('\n Start opcode: ', item)
      const res = await providerL2.send('debug_generateProofForOpcode', [
        taskArgs.fraud,
        mintRes.hash,
        item,
      ])
      const json = JSON.stringify(res)
      const { proof } = JSON.parse(json)
      // eslint-disable-next-line eqeqeq
      if (proof.opcode != '') {
        console.log('proof for opcode', item, 'success')
        osp.push(json)
        continue
      }
      console.log('tx do not have action about opcode:', item)
    }
    const verifierTestDriver = await getContractFactory(
      'VerifierTestDriver'
    ).attach('0xc6e7DF5E7b4f2A278906862b61205850344D4e7d')

    for (const item of osp) {
      const { ctx, proof } = JSON.parse(item)
      console.log(`processing verifier ${proof.verifier}`)

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
        .connect(ownerWalletL1)
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
        console.log('next hash not equal with proof: ', proof.nextHash)
      } else {
        console.log('verify success')
      }
    }
  })
module.exports = {}

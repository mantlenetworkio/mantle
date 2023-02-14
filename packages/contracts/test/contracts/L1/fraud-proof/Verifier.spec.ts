// eslint-disable-next-line import/order
import { Contract, Signer } from 'ethers'

// @ts-ignore
import { ethers } from 'hardhat'

// @ts-ignore
import { deploy } from '../../../helpers'
import { ctx, proof } from '../../../data/json/fraud-proof/fp_challenge.json'

describe('RollUp', () => {
  let accounts: Signer[]
  let verifierTestDriver: Contract
  let stackOpVerifier: Contract
  let verifier: Contract

  const tx = [
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

  before('setup', async () => {
    accounts = await ethers.getSigners()
    await deployVerifier()
    await deployVerifierEntry()
  })

  it('verifierEntry verifyProof', async () => {
    const ctxCall = {
      coinbase: ctx.coinbase,
      timestamp: ctx.timestamp,
      number: ctx.blockNumber,
      origin: ctx.origin,
      transaction: tx,
      inputRoot:
        '0x0000000000000000000000000000000000000000000000000000000000000000',
      txHash: ctx.txnHash,
    }

    await verifier.verifyOneStepProof(
      ctxCall,
      proof.verifier,
      proof.currHash,
      proof.proof
    )
  })

  it('VerifierTestDriver verifyProof', async () => {
    await verifierTestDriver.verifyProof(
      ctx.coinbase,
      ctx.timestamp,
      ctx.blockNumber,
      ctx.origin,
      ctx.txnHash,
      tx,
      proof.verifier,
      proof.currHash,
      proof.proof
    )
  })

  const deployVerifier = async () => {
    stackOpVerifier = await deploy('StackOpVerifier')
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    verifierTestDriver = await deploy('VerifierTestDriver', {
      args: [
        stackOpVerifier.address,
        stackOpVerifier.address,
        stackOpVerifier.address,
        stackOpVerifier.address,
        stackOpVerifier.address,
        stackOpVerifier.address,
        stackOpVerifier.address,
        stackOpVerifier.address,
        stackOpVerifier.address,
      ],
    })
  }

  const deployVerifierEntry = async () => {
    const verifierEntry = await deploy('VerifierEntry')
    const callData = verifierEntry.interface.encodeFunctionData('initialize')
    const verifierProxy = await deploy('TransparentUpgradeableProxy', {
      args: [
        verifierEntry.address, // logic
        await accounts[2].getAddress(), // admin
        callData, // call data
      ],
    })

    verifier = new Contract(
      verifierProxy.address,
      verifierEntry.interface,
      accounts[0]
    )
    await verifier.setVerifier(0, stackOpVerifier.address)
  }
})

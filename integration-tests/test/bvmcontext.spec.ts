/* Imports: External */
import { ethers } from 'hardhat'
import { expectApprox } from '@mantleio/core-utils'
import { predeploys } from '@mantleio/contracts'
import { Contract, BigNumber } from 'ethers'

/* Imports: Internal */
import { expect } from './shared/setup'
import { envConfig, DEFAULT_TEST_GAS_L1 } from './shared/utils'
import { MantleEnv } from './shared/env'

/**
 * These tests cover the BVM execution contexts. In the BVM execution
 * of a L1 to L2 transaction, both `block.number` and `block.timestamp`
 * must be equal to the blocknumber/timestamp of the L1 transaction.
 */
describe('BVM Context: Layer 2 EVM Context', () => {
  let env: MantleEnv
  before(async () => {
    env = await MantleEnv.new()
  })

  let Multicall: Contract
  let BVMContextStorage: Contract
  beforeEach(async () => {
    const BVMContextStorageFactory = await ethers.getContractFactory(
      'BVMContextStorage',
      env.l2Wallet
    )
    const MulticallFactory = await ethers.getContractFactory(
      'Multicall',
      env.l2Wallet
    )

    BVMContextStorage = await BVMContextStorageFactory.deploy()
    await BVMContextStorage.deployTransaction.wait()
    Multicall = await MulticallFactory.deploy()
    await Multicall.deployTransaction.wait()
  })

  const numTxs = envConfig.BVMCONTEXT_SPEC_NUM_TXS

  it('enqueue: L1 contextual values are correctly set in L2', async () => {
    for (let i = 0; i < numTxs; i++) {
      // Send a transaction from L1 to L2. This will automatically update the L1 contextual
      // information like the L1 block number and L1 timestamp.
      const tx =
        await env.messenger.contracts.l1.L1CrossDomainMessenger.sendMessage(
          BVMContextStorage.address,
          '0x',
          2_000_000,
          {
            gasLimit: DEFAULT_TEST_GAS_L1,
          }
        )

      // Wait for the transaction to be sent over to L2.
      await tx.wait()
      const pair = await env.waitForXDomainTransaction(tx)

      // Get the L1 block that the enqueue transaction was in so that
      // the timestamp can be compared against the layer two contract
      const l1Block = await env.l1Provider.getBlock(pair.receipt.blockNumber)
      const l2Block = await env.l2Provider.getBlock(
        pair.remoteReceipt.blockNumber
      )

      // block.number should return the value of the L2 block number.
      const l2BlockNumber = await BVMContextStorage.blockNumbers(i)
      expect(l2BlockNumber.toNumber()).to.deep.equal(l2Block.number)

      // L1BLOCKNUMBER opcode should return the value of the L1 block number.
      const l1BlockNumber = await BVMContextStorage.l1BlockNumbers(i)
      expect(l1BlockNumber.toNumber()).to.deep.equal(l1Block.number)

      // L1 and L2 blocks will have approximately the same timestamp.
      const timestamp = await BVMContextStorage.timestamps(i)
      expectApprox(timestamp.toNumber(), l1Block.timestamp, {
        percentUpperDeviation: 5,
      })
      expect(timestamp.toNumber()).to.deep.equal(l2Block.timestamp)

      // Difficulty should always be zero.
      const difficulty = await BVMContextStorage.difficulty(i)
      expect(difficulty.toNumber()).to.equal(0)

      // Coinbase should always be sequencer fee vault.
      const coinbase = await BVMContextStorage.coinbases(i)
      expect(coinbase).to.equal(predeploys.BVM_SequencerFeeVault)
    }
  })

  it('should set correct BVM Context for `eth_call`', async () => {
    for (let i = 0; i < numTxs; i++) {
      // Make an empty transaction to bump the latest block number.
      const dummyTx = await env.l2Wallet.sendTransaction({
        to: `0x${'11'.repeat(20)}`,
        data: '0x',
      })
      await dummyTx.wait()

      const block = await env.l2Provider.getBlockWithTransactions('latest')
      const [, returnData] = await Multicall.callStatic.aggregate(
        [
          [
            BVMContextStorage.address,
            BVMContextStorage.interface.encodeFunctionData(
              'getCurrentBlockTimestamp'
            ),
          ],
          [
            BVMContextStorage.address,
            BVMContextStorage.interface.encodeFunctionData(
              'getCurrentBlockNumber'
            ),
          ],
          [
            BVMContextStorage.address,
            BVMContextStorage.interface.encodeFunctionData(
              'getCurrentL1BlockNumber'
            ),
          ],
        ],
        { blockTag: block.number }
      )

      const timestamp = BigNumber.from(returnData[0])
      const blockNumber = BigNumber.from(returnData[1])
      const l1BlockNumber = BigNumber.from(returnData[2])
      const tx = block.transactions[0] as any

      expect(tx.l1BlockNumber).to.deep.equal(l1BlockNumber.toNumber())
      expect(block.timestamp).to.deep.equal(timestamp.toNumber())
      expect(block.number).to.deep.equal(blockNumber.toNumber())
    }
  })

  /**
   * `rollup_getInfo` is a new RPC endpoint that is used to return the BVM
   * context. The data returned should match what is actually being used as the
   * BVM context.
   */
  // TODO: This test is not reliable. If we really care about this then we need to figure out a
  // more reliable way to test this behavior.
  it.skip('should return same timestamp and blocknumbers between `eth_call` and `rollup_getInfo`', async () => {
    // // As atomically as possible, call `rollup_getInfo` and Multicall for the
    // // blocknumber and timestamp. If this is not atomic, then the sequencer can
    // // happend to update the timestamp between the `eth_call` and the `rollup_getInfo`
    // const [info, [, returnData]] = await Promise.all([
    //   L2Provider.send('rollup_getInfo', []),
    //   Multicall.callStatic.aggregate([
    //     [
    //       BVMContextStorage.address,
    //       BVMContextStorage.interface.encodeFunctionData(
    //         'getCurrentBlockTimestamp'
    //       ),
    //     ],
    //     [
    //       BVMContextStorage.address,
    //       BVMContextStorage.interface.encodeFunctionData(
    //         'getCurrentL1BlockNumber'
    //       ),
    //     ],
    //   ]),
    // ])
    // const timestamp = BigNumber.from(returnData[0])
    // const blockNumber = BigNumber.from(returnData[1])
    // expect(info.ethContext.blockNumber).to.deep.equal(blockNumber.toNumber())
    // expect(info.ethContext.timestamp).to.deep.equal(timestamp.toNumber())
  })
})

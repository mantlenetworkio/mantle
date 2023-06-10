/* Imports: External */
import { StateBatchAppendedEvent } from '@mantleio/contracts/dist/types/contracts/L1/rollup/StateCommitmentChain'
import { getContractFactory } from '@mantleio/contracts'
import { BigNumber } from 'ethers'

/* Imports: Internal */
import {
  StateRootBatchEntry,
  StateBatchAppendedExtraData,
  StateBatchAppendedParsedEvent,
  StateRootEntry,
  EventHandlerSet,
} from '../../../types'

export const handleEventsStateCachedBatchAppended: EventHandlerSet<
  StateBatchAppendedEvent,
  StateBatchAppendedExtraData,
  StateBatchAppendedParsedEvent
> = {
  getExtraData: async (event) => {
    const eventBlock = await event.getBlock()
    const l1Transaction = await event.getTransaction()

    return {
      timestamp: eventBlock.timestamp,
      blockNumber: eventBlock.number,
      submitter: l1Transaction.from,
      l1TransactionHash: l1Transaction.hash,
      l1TransactionData: l1Transaction.data,
    }
  },
  parseEvent: (event, extraData) => {
    let stateRoots: any
    try {
      stateRoots = getContractFactory(
        'StateCommitmentChain'
      ).interface.decodeFunctionData(
        'appendStateBatch',
        extraData.l1TransactionData
      )[0]
    } catch (e) {
      stateRoots = getContractFactory('Rollup').interface.decodeFunctionData(
        'createAssertionWithStateBatch',
        extraData.l1TransactionData
      )[2]
    }
    const stateRootEntries: StateRootEntry[] = []
    for (let i = 0; i < stateRoots.length; i++) {
      stateRootEntries.push({
        index: event.args._prevTotalElements.add(BigNumber.from(i)).toNumber(),
        batchIndex: event.args._batchIndex.toNumber(),
        value: stateRoots[i],
        confirmed: true,
      })
    }

    // Using .toNumber() here and in other places because I want to move everything to use
    // BigNumber + hex, but that'll take a lot of work. This makes it easier in the future.
    const stateRootBatchEntry: StateRootBatchEntry = {
      index: event.args._batchIndex.toNumber(),
      blockNumber: BigNumber.from(extraData.blockNumber).toNumber(),
      timestamp: BigNumber.from(extraData.timestamp).toNumber(),
      submitter: extraData.submitter,
      size: event.args._batchSize.toNumber(),
      root: event.args._batchRoot,
      prevTotalElements: event.args._prevTotalElements.toNumber(),
      extraData: event.args._extraData,
      l1TransactionHash: extraData.l1TransactionHash,
      type: 'LEGACY', // There is currently only 1 state root batch type
    }

    return {
      stateRootBatchEntry,
      stateRootEntries,
    }
  },
  storeEvent: async (entry, db) => {
    // Defend against situations where we missed an event because the RPC provider
    // (infura/alchemy/whatever) is missing an event.
    await db.putStateRootCachedBatchEntries([entry.stateRootBatchEntry])
    await db.putStateRootCachedEntries(entry.stateRootEntries)
  },
}

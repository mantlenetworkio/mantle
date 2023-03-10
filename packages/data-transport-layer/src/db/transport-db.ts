/* Imports: External */
import { LevelUp } from 'levelup'
import { BigNumber } from 'ethers'
import { BatchType } from '@mantleio/core-utils'

import { SimpleDB } from './simple-db'
import { PATCH_CONTEXTS, BSS_HF1_INDEX } from '../config'
import {
  BatchTransactionEntry,
  DataStoreEntry,
  EnqueueEntry,
  RollupStoreEntry,
  StateRootBatchEntry,
  StateRootEntry,
  TransactionBatchEntry,
  TransactionEntry,
  TransactionListEntry
} from "../types/database-types";

const TRANSPORT_DA_DB_KEYS = {
  UPDATED_BATCH_INDEX: `da:updatedbatchindex`,
  BATCH_INDEX: `da:batchindex`,
  TXS_BY_DS_ID: `da:txsbydsid`,
  DATA_STORE_BY_ID: `da:datastorebyid`,
  TX_LIST_DS_ID: `da:txlistdsid`,
  DS_LIST_BY_BTACH_ID: `da:dssbybatchid`,
  ROLLUP_STORE_BY_BATCH_INDEX: `da:rollupstorebybatchindex`,
  UPDATED_ROLLUP_STORE_BY_BATCH_INDEX: `da:updatedrollupstorebybatchindex`,

  UPDATED_DS_ID: `da:updateddsid`,
}

const TRANSPORT_DB_KEYS = {
  ENQUEUE: `enqueue`,
  ENQUEUE_CTC_INDEX: `ctc:enqueue`,
  TRANSACTION: `transaction`,
  UNCONFIRMED_TRANSACTION: `unconfirmed:transaction`,
  UNCONFIRMED_HIGHEST: `unconfirmed:highest`,
  TRANSACTION_BATCH: `batch:transaction`,
  STATE_ROOT: `stateroot`,
  UNCONFIRMED_STATE_ROOT: `unconfirmed:stateroot`,
  STATE_ROOT_BATCH: `batch:stateroot`,
  STARTING_L1_BLOCK: `l1:starting`,
  HIGHEST_L2_BLOCK: `l2:highest`,
  HIGHEST_SYNCED_BLOCK: `synced:highest`,
  CONSISTENCY_CHECK: `consistency:checked`,
}

interface Indexed {
  index: number
}

interface ExtraTransportDBOptions {
  l2ChainId?: number
}

export class TransportDB {
  public db: SimpleDB
  public opts: ExtraTransportDBOptions

  constructor(leveldb: LevelUp, opts?: ExtraTransportDBOptions) {
    this.db = new SimpleDB(leveldb)
    this.opts = opts || {}
  }

  //======================================================================================
  //====================================DA==PART==START===================================
  //======================================================================================
  public async putUpdatedBatchIndex(index: number): Promise<void> {
    await this.db.put([
      {
        key: TRANSPORT_DA_DB_KEYS.UPDATED_BATCH_INDEX,
        index: 0,
        value: index,
      },
    ])
  }

  public async getUpdatedBatchIndex(): Promise<number> {
    return this.db.get(TRANSPORT_DA_DB_KEYS.UPDATED_BATCH_INDEX, 0)
  }

  public async putUpdatedRollupBatchIndex(index: number): Promise<void> {
    await this.db.put([
      {
        key: TRANSPORT_DA_DB_KEYS.UPDATED_ROLLUP_STORE_BY_BATCH_INDEX,
        index: 0,
        value: index,
      },
    ])
  }

  public async getUpdatedRollupBatchIndex(): Promise<number> {
    return this.db.get(TRANSPORT_DA_DB_KEYS.UPDATED_ROLLUP_STORE_BY_BATCH_INDEX, 0)
  }

  public async putUpdatedDsId(index: number): Promise<void> {
    await this.db.put([
      {
        key: TRANSPORT_DA_DB_KEYS.UPDATED_DS_ID,
        index: 0,
        value: index,
      },
    ])
  }

  public async getUpdatedDsId(): Promise<number> {
    return this.db.get(TRANSPORT_DA_DB_KEYS.UPDATED_DS_ID, 0)
  }

  public async putRollupStoreByBatchIndex<RollupStoreEntry extends Indexed>(
    entry: RollupStoreEntry,
    batchIndex: number
  ): Promise<void> {
    await this.db.put([
      {
        key: TRANSPORT_DA_DB_KEYS.ROLLUP_STORE_BY_BATCH_INDEX + batchIndex,
        index: 0,
        value: entry,
      },
    ])
  }

  public async getRollupStoreByBatchIndex(
    batchIndex: number
  ): Promise<RollupStoreEntry> {
    return this.db.get(
      TRANSPORT_DA_DB_KEYS.ROLLUP_STORE_BY_BATCH_INDEX + batchIndex,
      0
    )
  }

  public async putBatchTxByDataStoreId(
    entries: BatchTransactionEntry[],
    dsId: number
  ): Promise<void> {
    await this._putFullEntries(
      TRANSPORT_DA_DB_KEYS.TXS_BY_DS_ID + dsId,
      entries
    )
  }

  public async getBatchTxByDataStoreId(
    dsId: number
  ): Promise<BatchTransactionEntry[]> {
    return this._getFullEnties(TRANSPORT_DA_DB_KEYS.TXS_BY_DS_ID + dsId)
  }

  // public async getTxListByDSId(dsId: number): Promise<TransactionListEntry[]> {
  //   const lens = await this._getLatestEntryIndex(
  //     TRANSPORT_DA_DB_KEYS.TX_LIST_DS_ID + dsId
  //   )
  //   return this._getEntries(TRANSPORT_DA_DB_KEYS.TX_LIST_DS_ID + dsId, 0, lens+1)
  // }

  public async getTxListByDSId(dsId: number): Promise<TransactionListEntry[]> {

    return this._getFullEnties(TRANSPORT_DA_DB_KEYS.TX_LIST_DS_ID + dsId)
  }



  public async putTxListByDSId(
    entries: TransactionListEntry[],
    dsId: number
  ): Promise<void> {
    await this._putFullEntries(TRANSPORT_DA_DB_KEYS.TX_LIST_DS_ID + dsId, entries)
  }

  public async putDsById(entry: DataStoreEntry, dsId: number): Promise<void> {
    await this.db.put([
      {
        key: TRANSPORT_DA_DB_KEYS.DATA_STORE_BY_ID,
        index: dsId,
        value: entry,
      },
    ])
  }

  public async getDsById<DataStoreEntry extends Indexed>(
    dsId: number
  ): Promise<DataStoreEntry> {
    if (dsId === null) {
      return null
    }

    return this._getEntryByIndex(TRANSPORT_DA_DB_KEYS.DATA_STORE_BY_ID, dsId)
  }

  public async getLatestBatchIndex(): Promise<number> {
    return this.db.get(TRANSPORT_DA_DB_KEYS.BATCH_INDEX, 0)
  }

  public async putLatestBatchIndex(newBatchIndex: number): Promise<void> {
    return this.db.put([
      {
        key: TRANSPORT_DA_DB_KEYS.BATCH_INDEX,
        index: 0,
        value: newBatchIndex,
      },
    ])
  }

  //=====================================================================
  //=============================DA===PART==END==========================
  //=====================================================================

  public async putEnqueueEntries(entries: EnqueueEntry[]): Promise<void> {
    await this._putEntries(TRANSPORT_DB_KEYS.ENQUEUE, entries)
  }

  public async putTransactionEntries(
    entries: TransactionEntry[]
  ): Promise<void> {
    await this._putEntries(TRANSPORT_DB_KEYS.TRANSACTION, entries)
  }

  public async putUnconfirmedTransactionEntries(
    entries: TransactionEntry[]
  ): Promise<void> {
    await this._putEntries(TRANSPORT_DB_KEYS.UNCONFIRMED_TRANSACTION, entries)
  }

  public async putTransactionBatchEntries(
    entries: TransactionBatchEntry[]
  ): Promise<void> {
    await this._putEntries(TRANSPORT_DB_KEYS.TRANSACTION_BATCH, entries)
  }

  public async putStateRootEntries(entries: StateRootEntry[]): Promise<void> {
    await this._putEntries(TRANSPORT_DB_KEYS.STATE_ROOT, entries)
  }

  public async putUnconfirmedStateRootEntries(
    entries: StateRootEntry[]
  ): Promise<void> {
    await this._putEntries(TRANSPORT_DB_KEYS.UNCONFIRMED_STATE_ROOT, entries)
  }

  public async putStateRootBatchEntries(
    entries: StateRootBatchEntry[]
  ): Promise<void> {
    await this._putEntries(TRANSPORT_DB_KEYS.STATE_ROOT_BATCH, entries)
  }

  public async putTransactionIndexByQueueIndex(
    queueIndex: number,
    index: number
  ): Promise<void> {
    await this.db.put([
      {
        key: TRANSPORT_DB_KEYS.ENQUEUE_CTC_INDEX,
        index: queueIndex,
        value: index,
      },
    ])
  }

  public async getTransactionIndexByQueueIndex(index: number): Promise<number> {
    return this.db.get(TRANSPORT_DB_KEYS.ENQUEUE_CTC_INDEX, index)
  }

  public async getEnqueueByIndex(index: number): Promise<EnqueueEntry> {
    return this._getEntryByIndex(TRANSPORT_DB_KEYS.ENQUEUE, index)
  }

  public async getTransactionByIndex(index: number): Promise<TransactionEntry> {
    return this._getEntryByIndex(TRANSPORT_DB_KEYS.TRANSACTION, index)
  }

  public async getUnconfirmedTransactionByIndex(
    index: number
  ): Promise<TransactionEntry> {
    return this._getEntryByIndex(
      TRANSPORT_DB_KEYS.UNCONFIRMED_TRANSACTION,
      index
    )
  }

  public async getTransactionsByIndexRange(
    start: number,
    end: number
  ): Promise<TransactionEntry[]> {
    return this._getEntries(TRANSPORT_DB_KEYS.TRANSACTION, start, end)
  }

  public async getTransactionBatchByIndex(
    index: number
  ): Promise<TransactionBatchEntry> {
    const entry = (await this._getEntryByIndex(
      TRANSPORT_DB_KEYS.TRANSACTION_BATCH,
      index
    )) as TransactionBatchEntry
    if (entry && typeof entry.type === 'undefined') {
      entry.type = BatchType[BatchType.LEGACY]
    }
    return entry
  }

  public async getStateRootByIndex(index: number): Promise<StateRootEntry> {
    return this._getEntryByIndex(TRANSPORT_DB_KEYS.STATE_ROOT, index)
  }

  public async getUnconfirmedStateRootByIndex(
    index: number
  ): Promise<StateRootEntry> {
    return this._getEntryByIndex(
      TRANSPORT_DB_KEYS.UNCONFIRMED_STATE_ROOT,
      index
    )
  }

  public async getStateRootsByIndexRange(
    start: number,
    end: number
  ): Promise<StateRootEntry[]> {
    return this._getEntries(TRANSPORT_DB_KEYS.STATE_ROOT, start, end)
  }

  public async getStateRootBatchByIndex(
    index: number
  ): Promise<StateRootBatchEntry> {
    return this._getEntryByIndex(TRANSPORT_DB_KEYS.STATE_ROOT_BATCH, index)
  }

  public async getLatestEnqueue(): Promise<EnqueueEntry> {
    return this._getLatestEntry(TRANSPORT_DB_KEYS.ENQUEUE)
  }

  public async getLatestTransaction(): Promise<TransactionEntry> {
    return this._getLatestEntry(TRANSPORT_DB_KEYS.TRANSACTION)
  }

  public async getLatestUnconfirmedTransaction(): Promise<TransactionEntry> {
    return this._getLatestEntry(TRANSPORT_DB_KEYS.UNCONFIRMED_TRANSACTION)
  }

  public async getLatestTransactionBatch(): Promise<TransactionBatchEntry> {
    const entry = (await this._getLatestEntry(
      TRANSPORT_DB_KEYS.TRANSACTION_BATCH
    )) as TransactionBatchEntry
    if (entry && typeof entry.type === 'undefined') {
      entry.type = BatchType[BatchType.LEGACY]
    }
    return entry
  }

  public async getLatestStateRoot(): Promise<StateRootEntry> {
    return this._getLatestEntry(TRANSPORT_DB_KEYS.STATE_ROOT)
  }

  public async getLatestUnconfirmedStateRoot(): Promise<StateRootEntry> {
    return this._getLatestEntry(TRANSPORT_DB_KEYS.UNCONFIRMED_STATE_ROOT)
  }

  public async getLatestStateRootBatch(): Promise<StateRootBatchEntry> {
    return this._getLatestEntry(TRANSPORT_DB_KEYS.STATE_ROOT_BATCH)
  }

  public async getHighestL2BlockNumber(): Promise<number> {
    return this.db.get<number>(TRANSPORT_DB_KEYS.HIGHEST_L2_BLOCK, 0)
  }

  public async getConsistencyCheckFlag(): Promise<boolean> {
    return this.db.get<boolean>(TRANSPORT_DB_KEYS.CONSISTENCY_CHECK, 0)
  }

  public async putConsistencyCheckFlag(flag: boolean): Promise<void> {
    return this.db.put<boolean>([
      {
        key: TRANSPORT_DB_KEYS.CONSISTENCY_CHECK,
        index: 0,
        value: flag,
      },
    ])
  }

  public async putHighestL2BlockNumber(
    block: number | BigNumber
  ): Promise<void> {
    if (block <= (await this.getHighestL2BlockNumber())) {
      return
    }

    return this.db.put<number>([
      {
        key: TRANSPORT_DB_KEYS.HIGHEST_L2_BLOCK,
        index: 0,
        value: BigNumber.from(block).toNumber(),
      },
    ])
  }

  public async getHighestSyncedUnconfirmedBlock(): Promise<number> {
    return (
      (await this.db.get<number>(TRANSPORT_DB_KEYS.UNCONFIRMED_HIGHEST, 0)) || 0
    )
  }

  public async setHighestSyncedUnconfirmedBlock(block: number): Promise<void> {
    return this.db.put<number>([
      {
        key: TRANSPORT_DB_KEYS.UNCONFIRMED_HIGHEST,
        index: 0,
        value: block,
      },
    ])
  }

  public async getHighestSyncedL1Block(): Promise<number> {
    return (
      (await this.db.get<number>(TRANSPORT_DB_KEYS.HIGHEST_SYNCED_BLOCK, 0)) ||
      0
    )
  }

  public async setHighestSyncedL1Block(block: number): Promise<void> {
    return this.db.put<number>([
      {
        key: TRANSPORT_DB_KEYS.HIGHEST_SYNCED_BLOCK,
        index: 0,
        value: block,
      },
    ])
  }

  public async getStartingL1Block(): Promise<number> {
    return this.db.get<number>(TRANSPORT_DB_KEYS.STARTING_L1_BLOCK, 0)
  }

  public async setStartingL1Block(block: number): Promise<void> {
    return this.db.put<number>([
      {
        key: TRANSPORT_DB_KEYS.STARTING_L1_BLOCK,
        index: 0,
        value: block,
      },
    ])
  }

  // Not sure if this next section belongs in this class.

  public async getFullTransactionByIndex(
    index: number
  ): Promise<TransactionEntry> {
    const transaction = await this.getTransactionByIndex(index)
    if (transaction === null) {
      return null
    }

    return this._makeFullTransaction(transaction)
  }

  public async getLatestFullTransaction(): Promise<TransactionEntry> {
    return this.getFullTransactionByIndex(
      await this._getLatestEntryIndex(TRANSPORT_DB_KEYS.TRANSACTION)
    )
  }

  public async getFullTransactionsByIndexRange(
    start: number,
    end: number
  ): Promise<TransactionEntry[]> {
    const transactions = await this.getTransactionsByIndexRange(start, end)
    if (transactions === null) {
      return null
    }

    const fullTransactions = []
    for (const transaction of transactions) {
      fullTransactions.push(await this._makeFullTransaction(transaction))
    }

    return fullTransactions
  }

  private async _makeFullTransaction(
    transaction: TransactionEntry
  ): Promise<TransactionEntry> {
    // We only need to do extra work for L1 to L2 transactions.
    if (transaction.queueOrigin !== 'l1') {
      return transaction
    }

    const enqueue = await this.getEnqueueByIndex(transaction.queueIndex)
    if (enqueue === null) {
      return null
    }

    let timestamp = enqueue.timestamp

    // BSS HF1 activates at block 0 if not specified.
    const bssHf1Index = BSS_HF1_INDEX[this.opts.l2ChainId] || 0
    if (transaction.index >= bssHf1Index) {
      timestamp = transaction.timestamp
    }

    // Override with patch contexts if necessary
    const contexts = PATCH_CONTEXTS[this.opts.l2ChainId]
    if (contexts && contexts[transaction.index + 1]) {
      timestamp = contexts[transaction.index + 1]
    }

    return {
      ...transaction,
      ...{
        blockNumber: enqueue.blockNumber,
        timestamp,
        gasLimit: enqueue.gasLimit,
        target: enqueue.target,
        origin: enqueue.origin,
        data: enqueue.data,
      },
    }
  }

  public async _getLatestEntryIndex(key: string): Promise<number> {
    return this.db.get<number>(`${key}:latest`, 0) || 0
  }

  private async _putLatestEntryIndex(
    key: string,
    index: number
  ): Promise<void> {
    return this.db.put<number>([
      {
        key: `${key}:latest`,
        index: 0,
        value: index,
      },
    ])
  }

  private async _getLatestEntry<TEntry extends Indexed>(
    key: string
  ): Promise<TEntry | null> {
    return this._getEntryByIndex(key, await this._getLatestEntryIndex(key))
  }

  private async _putLatestEntry<TEntry extends Indexed>(
    key: string,
    entry: TEntry
  ): Promise<void> {
    const latest = await this._getLatestEntryIndex(key)
    if (entry.index >= latest) {
      await this._putLatestEntryIndex(key, entry.index)
    }
  }

  private async _putEntries<TEntry extends Indexed>(
    key: string,
    entries: TEntry[]
  ): Promise<void> {
    if (entries.length === 0) {
      return
    }

    await this.db.put<TEntry>(
      entries.map((entry) => {
        return {
          key: `${key}:index`,
          index: entry.index,
          value: entry,
        }
      })
    )

    await this._putLatestEntry(key, entries[entries.length - 1])
  }

  private async _getEntryByIndex<TEntry extends Indexed>(
    key: string,
    index: number
  ): Promise<TEntry | null> {
    if (index === null) {
      return null
    }
    return this.db.get<TEntry>(`${key}:index`, index)
  }

  private async _getEntries<TEntry extends Indexed>(
    key: string,
    startIndex: number,
    endIndex: number
  ): Promise<TEntry[] | []> {
    return this.db.range<TEntry>(`${key}:index`, startIndex, endIndex)
  }

  private async _getFullEnties<TEntry extends Indexed>(
    key: string
  ): Promise<TEntry[] | []> {
    if (key === null) {
      return []
    }
    const lastIndex = await this._getLatestEntryIndex(key)
    if (lastIndex === null) {
      return []
    }
    return this._getEntries(key, 0, lastIndex+1)
  }

  private async _putFullEntries<TEntry extends Indexed>(
    key: string,
    entries: TEntry[]
  ): Promise<void> {
    await this._putLatestEntryIndex(key, entries.length -1)
    await this._putEntries(key, entries)
  }
}

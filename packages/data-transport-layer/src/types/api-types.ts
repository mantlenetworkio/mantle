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
} from "./database-types";

export type EnqueueResponse = EnqueueEntry & {
  ctcIndex: number | null
}

export interface TransactionResponse {
  batch: TransactionBatchEntry
  transaction: TransactionEntry
}

export interface TransactionBatchResponse {
  batch: TransactionBatchEntry
  transactions: TransactionEntry[]
}

export interface StateRootResponse {
  batch: StateRootBatchEntry
  stateRoot: StateRootEntry
}

export interface StateRootBatchResponse {
  batch: StateRootBatchEntry
  stateRoots: StateRootEntry[]
}

export interface ContextResponse {
  blockNumber: number
  timestamp: number
  blockHash: string
}

export interface GasPriceResponse {
  gasPrice: string
}

export interface LatestTxBatchIndexResponse {
  batchIndex: number
}

export interface DataStoreListByBatchIndexResponse {
  dataStore: RollupStoreEntry
  batchIndex: number
}

export interface BatchTxByDataStoreIdResponse {
  dsId: number
  batchTx: BatchTransactionEntry[]
}

export interface DataStoreByIdResponse {
  dataStore: string
}

export interface TestResponse {
  len: number
  putdata: string
  data: string
}

export interface TxListByStoreIdResponse {
  txList: TransactionListEntry[]
  storeId: number
}

export type SyncingResponse =
  | {
      syncing: true
      highestKnownTransactionIndex: number
      currentTransactionIndex: number
    }
  | {
      syncing: false
      currentTransactionIndex: number
    }

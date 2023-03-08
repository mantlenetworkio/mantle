export interface DecodedSequencerBatchTransaction {
  sig: {
    r: string
    s: string
    v: number
  }
  value: string
  gasLimit: string
  gasPrice: string
  nonce: string
  target: string
  data: string
}

export interface EnqueueEntry {
  index: number
  target: string
  data: string
  gasLimit: string
  origin: string
  blockNumber: number
  timestamp: number
}

export interface DataStoreEntry {
  index: number
  dataStoreId: string
  storeNumber: string
  durationDataStoreId: string
  dataCommitment: string
  msgHash: string
  stakesFromBlockNumber: string
  initTime: string
  expireTime: number
  duration: number
  numSys: string
  numPar: string
  degree: string
  storePeriodLength: string
  fee: string
  confirmer: string
  header: string
  initTxHash: string
  initGasUsed: string
  initBlockNumber: string
  confirmed: string
  ethSigned: string
  eigenSigned: string
  nonSignerPubKeyHashes: string
  signatoryRecord: string
  confirmTxHash: string
  confirmGasUsed: string
}
export interface TransactionListEntry {
  index: number
  blockNumber: string
  txHash: string
}

export interface TransactionEntry {
  index: number
  batchIndex: number
  data: string
  blockNumber: number
  timestamp: number
  gasLimit: string
  target: string
  origin: string
  value: string
  queueOrigin: 'sequencer' | 'l1'
  queueIndex: number | null
  decoded: DecodedSequencerBatchTransaction | null
  confirmed: boolean
}

export interface RollupStoreEntry {
  index: number
  data_store_id: number
  status: string
  confirm_at: number
}

interface BatchEntry {
  index: number
  blockNumber: number
  timestamp: number
  submitter: string
  size: number
  root: string
  prevTotalElements: number
  extraData: string
  l1TransactionHash: string
  type: string
}

export type TransactionBatchEntry = BatchEntry
export type StateRootBatchEntry = BatchEntry

export interface StateRootEntry {
  index: number
  batchIndex: number
  value: string
  confirmed: boolean
}

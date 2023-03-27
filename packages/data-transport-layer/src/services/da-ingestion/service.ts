/* Imports: External */
import { sleep } from '@mantleio/core-utils'
import { BaseService, Metrics } from '@mantleio/common-ts'
import { BaseProvider } from '@ethersproject/providers'
import { LevelUp } from 'levelup'
import { constants } from 'ethers'
// eslint-disable-next-line import/order
import { Gauge, Counter } from 'prom-client'

/* Imports: Internal */
// import { serialize } from '@ethersproject/transactions'
import fetch from 'node-fetch'
import { MissingElementError } from './handlers/errors'
import { TransportDB } from '../../db/transport-db'
import { validators } from '../../utils'
import { L1DataTransportServiceOptions } from '../main/service'
import {
  TransactionEntry,
  DataStoreEntry,
  TransactionListEntry, RollupStoreEntry,
} from '../../types'

interface DaIngestionMetrics {
  highestSyncedL1Block: Gauge<string>
  missingElementCount: Counter<string>
  unhandledErrorCount: Counter<string>
}

const registerMetrics = ({
  client,
  registry,
}: Metrics): DaIngestionMetrics => ({
  highestSyncedL1Block: new client.Gauge({
    name: 'data_transport_layer_synced_da_data',
    help: 'Synced DA  Data',
    registers: [registry],
  }),
  missingElementCount: new client.Counter({
    name: 'data_transport_layer_da_missing_element_count',
    help: 'Number of times recovery from missing elements happens',
    registers: [registry],
  }),
  unhandledErrorCount: new client.Counter({
    name: 'data_transport_layer_da_unhandled_error_count',
    help: 'Number of times recovered from unhandled errors',
    registers: [registry],
  }),
})

export interface DaIngestionServiceOptions
  extends L1DataTransportServiceOptions {
  db: LevelUp
  metrics: Metrics
}

const optionSettings = {
  db: {
    validate: validators.isLevelUP,
  },
  pollingInterval: {
    default: 5000,
    validate: validators.isInteger,
  },

  dangerouslyCatchAllErrors: {
    default: false,
    validate: validators.isBoolean,
  },
  l1RpcProvider: {
    validate: (val: any) => {
      return validators.isString(val) || validators.isJsonRpcProvider(val)
    },
  },
  l2ChainId: {
    validate: validators.isInteger,
  },
}

export interface Range {
  start: number
  end: number
}

export class DaIngestionService extends BaseService<DaIngestionServiceOptions> {
  constructor(options: DaIngestionServiceOptions) {
    super('Da_Ingestion_Service', options, optionSettings)
  }

  private daIngestionMetrics: DaIngestionMetrics

  private state: {
    db: TransportDB
    l1RpcProvider: BaseProvider
    startingL1BlockNumber: number
    mtBatcherFetchUrl: string
  } = {} as any

  protected async _init(): Promise<void> {
    this.state.db = new TransportDB(this.options.db, {
      l2ChainId: this.options.l2ChainId,
    })

    const lastBatchIndex =  await this.state.db.getLastBatchIndex()
    if (lastBatchIndex <= 0 || lastBatchIndex === null) {
      await this.state.db.putLastBatchIndex(this.options.daInitBatch)
    }

    this.daIngestionMetrics = registerMetrics(this.metrics)

    this.state.mtBatcherFetchUrl =
      this.options.mtBatcherHost +
      ':' +
      this.options.mtBatcherFetchPort.toString()
  }

  protected async _start(): Promise<void> {
    while (this.running) {
      try {
        const batchIndexRange = await this.getBatchIndexRange()
        if (batchIndexRange.start >= batchIndexRange.end) {
          continue
        }
        await this.pareTransaction(batchIndexRange)
        await sleep(this.options.pollingInterval)
      } catch (err) {
        if (err instanceof MissingElementError) {
          this.logger.warn('recovering from a missing event', {
            message: err.toString(),
          })
        } else if (!this.running || this.options.dangerouslyCatchAllErrors) {
          this.daIngestionMetrics.unhandledErrorCount.inc()
          this.logger.error('Caught an unhandled error', {
            message: err.toString(),
            stack: err.stack,
            code: err.code,
          })
          await sleep(this.options.daPollingInterval)
        } else {
          throw err
        }
      }
    }
  }

  private async pareTransaction(batchIndexRange: Range) {
    const dataStore: DataStoreEntry[] = []
    const transactionEntries: TransactionEntry[] = []
    const exploreTransactionEntries: TransactionEntry[] = []
    for ( let index = batchIndexRange.start; index <= batchIndexRange.end;  index++) {
      const dataStoreRollupId = await this.GetRollupStoreByRollupBatchIndex(index)
      if (dataStoreRollupId['data_store_id'] === 0) {
        continue
      }
      const dataStore = await this.GetDataStoreById(dataStoreRollupId['data_store_id'].toString())
      if (dataStore === null) {
        continue
      }
      if (dataStore['Confirmed']) {
        // explore transaction list
        await this._storeTransactionListByDSId(dataStoreRollupId['data_store_id'])

        // batch transaction list
        await this._storeBatchTransactionsByDSId(dataStoreRollupId['data_store_id'])

        // put rollup store info to db
        await this.state.db.putRollupStoreByBatchIndex(
          {
            index: 0,
            data_store_id: dataStoreRollupId['data_store_id'],
            status: dataStoreRollupId['status'],
            confirm_at: dataStoreRollupId['confirm_at'],
          },
          index
        )
        // put data store to db
        const dataStoreEntry: DataStoreEntry = {
          dataStoreId: dataStore['Id'],
          storeNumber: dataStore['StoreNumber'],
          durationDataStoreId: dataStore['DurationDataStoreId'],
          index: dataStore['Index'],
          dataCommitment: dataStore['DataCommitment'],
          msgHash: dataStore['MsgHash'],
          stakesFromBlockNumber: dataStore['StakesFromBlockNumber'],
          initTime: dataStore['InitTime'],
          expireTime: dataStore['ExpireTime'],
          duration: dataStore['Duration'],
          numSys: dataStore['NumSys'],
          numPar: dataStore['NumPar'],
          degree: dataStore['Degree'],
          storePeriodLength: dataStore['StorePeriodLength'],
          fee: dataStore['Fee'],
          confirmer: dataStore['Confirmer'],
          header: dataStore['Header'],
          initTxHash: dataStore['InitTxHash'],
          initGasUsed: dataStore['InitGasUsed'],
          initBlockNumber: dataStore['InitBlockNumber'],
          confirmed: dataStore['Confirmed'],
          ethSigned: dataStore['EthSigned'],
          eigenSigned: dataStore['EigenSigned'],
          nonSignerPubKeyHashes: dataStore['NonSignerPubKeyHashes'],
          signatoryRecord: dataStore['SignatoryRecord'],
          confirmTxHash: dataStore['ConfirmTxHash'],
          confirmGasUsed: dataStore['ConfirmGasUsed'],
        }
        await this.state.db.putDsById(dataStoreEntry, dataStoreRollupId['data_store_id'])
      }
    }
  }

  private async getBatchIndexRange(): Promise<Range> {
    const latestBatchIndex = await this.state.db.getLastBatchIndex()
    const newTxBatchIndex: number = await this.GetLatestTransactionBatchIndex()
    const end = latestBatchIndex + this.options.daSyncStep
    if (newTxBatchIndex > latestBatchIndex) {
      await this.state.db.putLastBatchIndex(end)
      return {
        start: latestBatchIndex,
        end: end,
      }
    } else {
      return {
        start: latestBatchIndex,
        end: newTxBatchIndex,
      }
    }
  }

  private async _storeBatchTransactionsByDSId(storeId: number) {
    const transactionEntries: TransactionEntry[] = []
    if (storeId <= 0) {
      return []
    }
    const batchTxs = await this.GetBatchTransactionByDataStoreId(storeId)
      .then((rst) => {
        return rst
      })
      .catch((error) => {
        console.log('GetBatchTransactionByDataStoreId error ', error)
        return []
      })
    try {
      if (batchTxs.length === 0) {
        return
      }
      for (const batchTx of batchTxs) {
        const queueOrigin =
          batchTx['TxMeta']['queueOrigin'] === 1 ? 'l1' : 'sequencer'
        const txData =
          batchTx['TxMeta']['queueOrigin'] === 1
            ? null
            : batchTx['TxMeta']['rawTransaction']
        const sigData =
          batchTx['TxMeta']['queueOrigin'] === 1
            ? null
            : {
              v: batchTx['TxDetail']['v'],
              r: batchTx['TxDetail']['r'],
              s: batchTx['TxDetail']['s'],
            }
        transactionEntries.push({
          index: batchTx['TxMeta']['index'],
          batchIndex: 0,
          blockNumber: batchTx['TxMeta']['l1BlockNumber'],
          timestamp: batchTx['TxMeta']['l1Timestamp'],
          gasLimit: '0',
          target: constants.AddressZero,
          origin: batchTx['TxMeta']['l1MessageSender'],
          data: txData,
          queueOrigin,
          value: batchTx['TxDetail']['value'],
          queueIndex: batchTx['TxMeta']['queueIndex'],
          decoded: {
            sig: sigData,
            value: batchTx['TxDetail']['value'],
            gasLimit: '0x0',
            gasPrice: batchTx['TxDetail']['gasPrice'],
            nonce: batchTx['TxDetail']['nonce'],
            target: constants.AddressZero,
            data: batchTx['TxDetail']['input'],
          },
          confirmed: true,
        })
      }
      await this.state.db.putTransactions(transactionEntries)
      await this.state.db.putBatchTransactionByDsId(transactionEntries, storeId);
    }catch (error) {
      throw new Error(
        `eigen layer sync finish, error is: ${error}`
      )
    }
  }

  private async _storeTransactionListByDSId(storeId: number): Promise<void> {
    const txList = await this.GetTransactionListByStoreNumber(storeId)
    if (
      txList === null ||
      txList.length === 0 ||
      JSON.stringify(txList) === '{}'
    ) {
      return
    }

    const transactionEntries: TransactionListEntry[] = []
    for (const tx of txList) {
      const index_ = transactionEntries.length
      transactionEntries.push({
        index: index_,
        txIndex: tx['index'],
        blockNumber: tx['BlockNumber'],
        txHash: tx['TxHash'],
      })
    }
    await this.state.db.putTxListByDSId(transactionEntries, storeId)
  }

  private async GetLatestTransactionBatchIndex(): Promise<number> {
    const data = await fetch(
      this.state.mtBatcherFetchUrl + '/eigen/getLatestTransactionBatchIndex',
      {
        method: 'GET',
        headers: { Accept: 'application/json' },
      }
    )
      .then((res) => res.json())
      .catch((error) => {
        console.log(
          'GetLatestTransactionBatchIndex HTTP  error : status!=200 error info = ',
          error
        )
        return 1
      })
    let newTxBatchIndex: number = 1
    if (typeof data === 'number') {
      newTxBatchIndex = data
    }
    return newTxBatchIndex
  }

  private async GetRollupStoreByRollupBatchIndex(
    batchIndex: number
  ): Promise<any> {
    const requestData = JSON.stringify({
      batch_index: batchIndex,
    })
    // 👇️ const response: Response
    return fetch(
      this.state.mtBatcherFetchUrl + '/eigen/getRollupStoreByRollupBatchIndex',
      {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: requestData,
      }
    )
      .then((res) => res.json())
      .catch((error) => {
        return error
      })
  }

  private async GetBatchTransactionByDataStoreId(
    storeNumber: number
  ): Promise<any> {
    const requestData = JSON.stringify({
      store_number: storeNumber,
    })
    // 👇️ const response: Response
    return fetch(
      this.state.mtBatcherFetchUrl + '/dtl/getBatchTransactionByDataStoreId',
      {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: requestData,
      }
    ).then((res) => res.json()).catch((error) => {
        console.log(
          'GetBatchTransactionByDataStoreId  HTTP error status != 200 ',
          error
        )
        return error
      })
  }

  private async GetDataStoreById(storeNumber: string): Promise<any> {
    // 👇️ const response: Response
    return (
      fetch(this.state.mtBatcherFetchUrl + '/browser/getDataStoreById', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          store_id: storeNumber,
        }),
      })
        .then((res) => res.json())
        .catch((error) => {
          console.log('GetDataStoreById HTTP error status != 200 ', error)
          return error
        })
    )
  }

  private async GetTransactionListByStoreNumber(
    storeNumber: number
  ): Promise<any> {
    const requestData = JSON.stringify({
      store_number: storeNumber,
    })
    // 👇️ const response: Response
    return fetch(
      this.state.mtBatcherFetchUrl + '/browser/GetTransactionListByStoreNumber',
      {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: requestData,
      }
    )
      .then((res) => res.json())
      .catch((error) => {
        console.log('GetTransactionListByStoreNumber HTTP error status != 200 ')
        return error
      })
  }
}

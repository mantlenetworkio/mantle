/* Imports: External */
import { sleep } from '@mantleio/core-utils'
import { BaseService, Metrics } from '@mantleio/common-ts'
import { BaseProvider } from '@ethersproject/providers'
import { LevelUp } from 'levelup'
import { constants } from 'ethers'
// eslint-disable-next-line import/order
import { Gauge, Counter } from 'prom-client'

/* Imports: Internal */
import { serialize } from '@ethersproject/transactions'
import fetch from 'node-fetch'

import { MissingElementError } from './handlers/errors'
import { TransportDB } from '../../db/transport-db'
import { validators } from '../../utils'
import { L1DataTransportServiceOptions } from '../main/service'
import {
  DataStoreEntry,
  TransactionEntry,
  TransactionListEntry,
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
    name: 'data_transport_layer_highest_synced_l1_block',
    help: 'Highest Synced L1 Block Number',
    registers: [registry],
  }),
  missingElementCount: new client.Counter({
    name: 'data_transport_layer_missing_element_count',
    help: 'Number of times recovery from missing elements happens',
    registers: [registry],
  }),
  unhandledErrorCount: new client.Counter({
    name: 'data_transport_layer_l1_unhandled_error_count',
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

    this.daIngestionMetrics = registerMetrics(this.metrics)

    this.state.mtBatcherFetchUrl =
      this.options.mtBatcherHost +
      ':' +
      this.options.mtBatcherFetchPort.toString()
  }

  protected async _start(): Promise<void> {
    // This is our main function. It's basically just an infinite loop that attempts to stay in
    // sync with events coming from Ethereum. Loops as quickly as it can until it approaches the
    // tip of the chain, after which it starts waiting for a few seconds between each loop to avoid
    // unnecessary spam.
    while (this.running) {
      try {
        const lastBatchIndex = await this.state.db.getLatestBatchIndex()

        let newTxBatchIndex: number = -1
        const getNewBatchIndex = async () => {
          const rst = await this.GetLatestTransactionBatchIndex()
          if (typeof rst === 'number') {
            newTxBatchIndex = rst
          }
        }
        await getNewBatchIndex()
        if (newTxBatchIndex !== -1) {
          await this.state.db.putUpdatedBatchIndex(newTxBatchIndex)
        }
        let now_batch_Index = -1

        for (let i = lastBatchIndex; i <= newTxBatchIndex; i++) {
          now_batch_Index = i

          const dataStore = await this.GetRollupStoreByRollupBatchIndex(i)
            .then((rst) => {
              return rst
            })
            .catch((error) => {
              console.log('getRollupStoreByRollupBatchIndex error : ', error)
            })
          if (dataStore === null) {
            console.log('HTTP getRollup and get null data')
            now_batch_Index--
            break
          }
          if (dataStore['status'] === 0) {
            now_batch_Index--
            break
          }
          await this.state.db.putRollupStoreByBatchIndex(
            {
              index: 0,
              data_store_id: dataStore['data_store_id'],
              status: dataStore['status'],
              confirm_at: dataStore['confirm_at'],
            },
            i
          )

          await this._updateBatchTxByDSId(
            dataStore['data_store_id'],
            now_batch_Index
          )
          await this._storeDataStoreById(dataStore['data_store_id'])

        }
        await this.state.db.putUpdatedBatchIndex(now_batch_Index)


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
          await sleep(this.options.pollingInterval)
        } else {
          throw err
        }
      }
    }
  }
  private async _updateBatchTxByDSId(
    storeId: number,
    index: number
  ): Promise<void> {
    const transactionEntries: TransactionEntry[] = []

    const batchTxs = await this.GetBatchTransactionByDataStoreId(storeId)
      .then((rst) => {
        return rst
      })
      .catch((error) => {
        console.log('GetBatchTransactionByDataStoreId error ', error)
        return null
      })
    console.log('batch tx :', batchTxs)
    for (const batchTx of batchTxs) {
      const queueOrigin =
        batchTx['txMeta']['queueOrigin'] === 1 ? 'l1' : 'sequencer'

      const txData =
        batchTx['txMeta']['queueOrigin'] === 1
          ? null
          : serialize(
              {
                nonce: batchTx['txDetail']['nonce'],
                gasPrice: batchTx['txDetail']['gasPrice'],
                gasLimit: '0',
                to: batchTx['txDetail']['to'],
                value: batchTx['txDetail']['value'],
                data: batchTx['txDetail']['input'],
              },
              {
                v: batchTx['txDetail']['v'],
                r: batchTx['txDetail']['r'],
                s: batchTx['txDetail']['s'],
              }
            )
      const sigData =
        batchTx['txMeta']['queueOrigin'] === 1
          ? null
          : {
              v: batchTx['txDetail']['v'],
              r: batchTx['txDetail']['r'],
              s: batchTx['txDetail']['s'],
            }
      //TODO:
      transactionEntries.push({
        index: batchTx['txMeta']['index'],
        batchIndex: index,
        blockNumber: batchTx['txMeta']['l1BlockNumber'],
        timestamp: batchTx['txMeta'],
        gasLimit: '0',
        target: constants.AddressZero,
        origin: null,
        data: txData,
        queueOrigin,
        value: batchTx['txDetail']['value'],
        queueIndex: batchTx['txMeta']['queueIndex'],
        decoded: {
          sig: sigData,
          value: batchTx['txDetail']['value'],
          gasLimit: '0x0',
          gasPrice: batchTx['txDetail']['gasPrice'],
          nonce: batchTx['txDetail']['nonce'],
          target: constants.AddressZero,
          data: batchTx['txDetail']['input'],
        },
        confirmed: true,
      })
    }

    await this.state.db.putBatchTxByDataStoreId(transactionEntries, storeId)
  }

  private async _storeDataStoreById(storeId: number): Promise<void> {
    const dataStore = await this.GetDataStoreById(storeId)
    if (dataStore === null) {
      return
    }
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
    await this.state.db.putDsById(dataStoreEntry, storeId)
  }

  private async _storeTransactionListByDSId(storeId: number): Promise<void> {
    const txList = await this.GetTransactionListByStoreNumber(storeId)
    if (txList === null || txList.length === 0) {
      return
    }
    const transactionEntries: TransactionListEntry[] = []
    const i = 0
    for (const tx of txList) {
      transactionEntries.push({
        index: i,
        blockNumber: tx['BlockNumber'],
        txHash: tx['TxHash'],
      })
    }
    await this.state.db.putTxListByDSId(transactionEntries, storeId)
  }

  private async GetLatestTransactionBatchIndex(): Promise<any> {
    return fetch(
      this.state.mtBatcherFetchUrl + '/eigen/getLatestTransactionBatchIndex',
      {
        method: 'GET',
        headers: { Accept: 'application/json' },
      }
    )
      .then((res) => {
        if (res.status === 200) {
          return res.json()
        }
        console.log('HTTP status error : status!=200')
        return -1
      })
      .catch((error) => {
        console.log('HTTP  error : status!=200 error info = ', error)
        return -1
      })
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
      .then((res) => {
        return res
      })
      .catch((error) => {
        return error
      })
  }

  private async GetBatchTransactionByDataStoreId(
    storeNumber: number
  ): Promise<any> {
    console.log('storeNumber', storeNumber)
    const requestData = JSON.stringify({
      store_number: storeNumber,
    })
    // 👇️ const response: Response
    return fetch(
      this.state.mtBatcherFetchUrl + '/eigen/getBatchTransactionByDataStoreId',
      {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: requestData,
      }
    )
      .then((res) => {
        if (res.status === 200) {
          return res
        }
        console.log('HTTP status != 200 ', res.json())
        return res
      })
      .then((res) => res.json())
      .catch((error) => {
        console.log('HTTP error status != 200 ')
        return error
      })
  }

  private async GetDataStoreById(storeNumber: number): Promise<any> {
    console.log('storeNumber', storeNumber)
    const requestData = JSON.stringify({
      store_id: storeNumber,
    })
    // 👇️ const response: Response
    return fetch(this.state.mtBatcherFetchUrl + '/browser/getDataStoreById', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: requestData,
    })
      .then((res) => {
        if (res.status === 200) {
          return res
        }
        console.log('HTTP status != 200 ', res.json())
        return res
      })
      .then((res) => res.json())
      .catch((error) => {
        console.log('HTTP error status != 200 ')
        return error
      })
  }

  private async GetTransactionListByStoreNumber(
    storeNumber: number
  ): Promise<any> {
    console.log('storeNumber', storeNumber)
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
      .then((res) => {
        if (res.status === 200) {
          return res
        }
        console.log('HTTP status != 200 ', res.json())
        return res
      })
      .then((res) => res.json())
      .catch((error) => {
        console.log('HTTP error status != 200 ')
        return error
      })
  }
}

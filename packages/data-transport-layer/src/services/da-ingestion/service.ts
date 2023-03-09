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
    while (this.running) {
      try {
        const maxDsId = await this.updateBatchIndexLoop()
        await this.updateRollupDataStoreLoop()
        await this.updateTransactionListByDsIDLoop(maxDsId)

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
          await sleep(this.options.pollingInterval)
        } else {
          throw err
        }
      }
    }
  }
  private async updateRollupDataStoreLoop(): Promise<void> {
    const latestBatchIndex = await this.GetLatestTransactionBatchIndex()
    let updatedBatchIndex = await this.state.db.getUpdatedRollupBatchIndex()
    console.log(latestBatchIndex,updatedBatchIndex)
    if (updatedBatchIndex === null) {
      updatedBatchIndex = 0
    }
    if (latestBatchIndex <= updatedBatchIndex) {
      return
    }
    const loopTime =
      updatedBatchIndex + 10 < latestBatchIndex
        ? updatedBatchIndex + 10
        : latestBatchIndex
    for (let i = updatedBatchIndex; i <= loopTime; i++) {
      const dataStore = await this.GetRollupStoreByRollupBatchIndex(i)
        .then((rst) => {
          return rst
        })
        .catch((error) => {
          console.log('getRollupStoreByRollupBatchIndex error : ', error)
        })
      if (dataStore === null) {
        console.log('HTTP getRollup and get null data')
        break
      }
      if (dataStore['status'] === 0) {
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
      await this.state.db.putUpdatedRollupBatchIndex(i)
    }
  }
  private async updateBatchIndexLoop(): Promise<number> {
    const lastBatchIndex = await this.state.db.getUpdatedBatchIndex()

    let newTxBatchIndex: number = -1
    const getNewBatchIndex = async () => {
      const rst = await this.GetLatestTransactionBatchIndex()
      if (typeof rst === 'number' && rst !== 0) {
        newTxBatchIndex = rst
      }
    }
    await getNewBatchIndex()
    if (newTxBatchIndex !== -1) {
      await this.state.db.putLatestBatchIndex(newTxBatchIndex)
    }
    let maxDsId: number = 0
    for (let i = lastBatchIndex; i <= newTxBatchIndex; i++) {
      const now_batchIndex = i
      const dataStore = await this.GetRollupStoreByRollupBatchIndex(i)
      if (dataStore === null) {
        console.log('HTTP getRollup and get null data')
        break
      }
      if (dataStore['status'] === 0) {
        break
      }
      maxDsId =
        maxDsId < dataStore['data_store_id']
          ? dataStore['data_store_id']
          : maxDsId
      console.log("maxDsId:",maxDsId,dataStore['data_store_id'])


      await this._storeDataStoreById(dataStore['data_store_id'])
      await this._updateBatchTxByDSId(
        dataStore['data_store_id'],
        now_batchIndex
      )
      await this.state.db.putUpdatedBatchIndex(i)
    }

    return maxDsId
  }

  private async updateTransactionListByDsIDLoop(
    endDsId: number
  ): Promise<void> {
    console.log('endDsId : ', endDsId)
    let startDsId = await this.state.db.getUpdatedDsId()
    if (startDsId === null) {
      startDsId = 1
    }
    if (startDsId >= endDsId) {
      return
    }
    const loopTime = endDsId - startDsId > 10 ? startDsId + 10 : endDsId
    for (let dsId = startDsId; dsId <= loopTime; dsId++) {
      await this._storeDataStoreById(dsId)
      console.log('updated data store entry ,data store id = ', dsId)
      await this._storeTransactionListByDSId(dsId)
      console.log('updated tx list by dsId ,data store id = ', dsId)
      await this.state.db.putUpdatedDsId(dsId)
    }
  }
  private async _updateBatchTxByDSId(
    storeId: number,
    index: number
  ): Promise<void> {
    const transactionEntries: TransactionEntry[] = []
    console.log("start update batch txs by dsid")
    const batchTxs = await this.GetBatchTransactionByDataStoreId(storeId)
      .then((rst) => {
        return rst
      })
      .catch((error) => {
        console.log('GetBatchTransactionByDataStoreId error ', error)
        return null
      })
    for (const batchTx of batchTxs) {
      const queueOrigin =
        batchTx['TxMeta']['queueOrigin'] === 1 ? 'l1' : 'sequencer'

      // const txData =
      //   batchTx['TxMeta']['queueOrigin'] === 1
      //     ? null
      //     : serialize(
      //         {
      //           nonce: batchTx['TxDetail']['nonce'],
      //           gasPrice: batchTx['TxDetail']['gasPrice'],
      //           gasLimit: 0,
      //           to: batchTx['TxDetail']['to'],
      //           value: batchTx['TxDetail']['value'],
      //           data: batchTx['TxDetail']['input'],
      //         },
      //         {
      //           v: batchTx['TxDetail']['v'],
      //           r: batchTx['TxDetail']['r'],
      //           s: batchTx['TxDetail']['s'],
      //         }
      //       )
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
      //TODO:
      transactionEntries.push({
        index: batchTx['TxMeta']['index'],
        batchIndex: index,
        blockNumber: batchTx['TxMeta']['l1BlockNumber'],
        timestamp: batchTx['TxMeta'],
        gasLimit: '0',
        target: constants.AddressZero,
        origin: null,
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
    console.log("txs store = ",storeId)
    await this.state.db.putBatchTxByDataStoreId(transactionEntries, storeId)
  }

  private async _storeDataStoreById(storeId: number): Promise<void> {
    if (storeId === null) {
      return
    }
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
    console.log('txList :', txList)
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
      .then((res) => res.json())
      .catch((error) => {
        console.log(
          'GetLatestTransactionBatchIndex HTTP  error : status!=200 error info = ',
          error
        )
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
      this.state.mtBatcherFetchUrl + '/eigen/getBatchTransactionByDataStoreId',
      {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: requestData,
      }
    )
      .then((res) => res.json())
      .catch((error) => {
        console.log(
          'GetBatchTransactionByDataStoreId  HTTP error status != 200 ',
          error
        )
        return error
      })
  }

  private async GetDataStoreById(storeNumber: number): Promise<any> {
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
        // .then((res) => res.json())
        .catch((error) => {
          console.log('GetDataStoreById HTTP error status != 200 ', error)
          return error
        })
    )
  }

  private async GetTransactionListByStoreNumber(
    storeNumber: number
  ): Promise<any> {
    console.log('store_number', storeNumber)
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

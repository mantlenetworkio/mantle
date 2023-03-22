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
        const dataStoreIdRange = await this.getDataStoreIdRange(batchIndexRange)
        if (dataStoreIdRange === null) {
          await sleep(this.options.pollingInterval)
          continue
        }
        await this.updateBatchTransactionsByDataStoreIdRange(dataStoreIdRange)

        await this.updateTransactionListAndDataStoreByDsIDLoop(dataStoreIdRange)
        await this.updateRollupDataStoreLoop(batchIndexRange)

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
  private async updateRollupDataStoreLoop(
    batchIndexRange: Range
  ): Promise<void> {
    if (batchIndexRange === null) {
      return
    }
    const latestBatchIndex = batchIndexRange.end
    let updatedBatchIndex = await this.state.db.getUpdatedRollupBatchIndex()
    if (updatedBatchIndex === null) {
      updatedBatchIndex = 0
    }
    if (
      latestBatchIndex <= updatedBatchIndex ||
      latestBatchIndex - updatedBatchIndex > 10 * 2
    ) {
      return
    }

    for (let i = updatedBatchIndex; i <= latestBatchIndex; i++) {
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
  private async getBatchIndexRange(): Promise<Range> {
    const lastBatchIndex = (await this.state.db.getUpdatedBatchIndex()) || 1
    const newTxBatchIndex: number = await this.GetLatestTransactionBatchIndex()
    if (newTxBatchIndex <= lastBatchIndex) {
      return null
    }
    await this.state.db.putLatestBatchIndex(newTxBatchIndex)
    const loopTime =
      newTxBatchIndex > lastBatchIndex + 10
        ? lastBatchIndex + 10
        : newTxBatchIndex
    return {
      start: lastBatchIndex,
      end: loopTime,
    }
  }

  private async getDataStoreIdRange(batchIndexRange: Range): Promise<Range> {
    if (batchIndexRange === null) {
      return null
    }
    const dataStore = await this.GetRollupStoreByRollupBatchIndex(
      batchIndexRange.start
    )
    if (dataStore === null || dataStore['status'] === 0) {
      console.log('HTTP getRollup and get null data')
      return null
    }
    const startDsId = dataStore['data_store_id']
    let dataStore_ = await this.GetRollupStoreByRollupBatchIndex(
      batchIndexRange.end
    )
    if (dataStore_ === null) {
      console.log('HTTP getRollup and get null data')
      return null
    } else if (dataStore_['status'] === 0) {
      dataStore_ = await this.GetRollupStoreByRollupBatchIndex(
        batchIndexRange.end - 1
      )
      if (dataStore_ === null) {
        console.log('HTTP getRollup and get null data')
        return null
      }
    }
    const endDsId = dataStore_['data_store_id']
    return {
      start: startDsId,
      end: endDsId,
    }
  }

  private async updateTransactionListAndDataStoreByDsIDLoop(
    dataStoreIdRange: Range
  ): Promise<void> {
    let startDsId = await this.state.db.getUpdatedDsId()
    if (startDsId === null) {
      startDsId = 1
    }
    if (startDsId >= dataStoreIdRange.end) {
      return
    }
    startDsId =
      startDsId > dataStoreIdRange.start ? dataStoreIdRange.start : startDsId
    for (let dsId = startDsId; dsId <= dataStoreIdRange.end; dsId++) {
      await this._storeDataStoreById(dsId)
      await this._storeTransactionListByDSId(dsId)
      await this.state.db.putUpdatedDsId(dsId)
    }
  }
  private async updateBatchTransactionsByDataStoreIdRange(
    dataStoreIdRange: Range
  ): Promise<void> {
    if (dataStoreIdRange == null) {
      return
    }

    for (
      let dataStoreId = dataStoreIdRange.start;
      dataStoreId <= dataStoreIdRange.end;
      dataStoreId++
    ) {
      this._storeBatchTransactionsByDSId(dataStoreId)
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
      await this.state.db.putBatchTxByDsId(batchTxs[0]['TxMeta']['index'], batchTxs[batchTxs.length -1]['TxMeta']['index'], storeId);
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
      await this.state.db.putTransactions(transactionEntries)
    }catch (error) {
      console.log("eigen layer sync finish")
    }
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

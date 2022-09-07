package index

import (
	"context"
	"time"

	"github.com/bitdao-io/mantle/l2geth/common"
	"github.com/bitdao-io/mantle/l2geth/common/hexutil"
	"github.com/bitdao-io/mantle/l2geth/ethclient"
	"github.com/bitdao-io/mantle/l2geth/log"
)

const scanRange = 10

type Indexer struct {
	store           IndexerStore
	l1Cli           *ethclient.Client
	l1ConfirmBlocks int
	sccContractAddr common.Address
	hook            Hook
	taskInterval    time.Duration
	stopChan        chan struct{}
}

type Hook interface {
	AfterStateBatchIndexed([32]byte) error
}

func NewIndexer(store IndexerStore, l1url string, l1ConfirmBlocks int, sccContractAddr string, taskInterval string) (Indexer, error) {
	taskIntervalDur, err := time.ParseDuration(taskInterval)
	if err != nil {
		return Indexer{}, nil
	}
	l1Cli, err := ethclient.Dial(l1url)
	if err != nil {
		return Indexer{}, err
	}
	address := common.HexToAddress(sccContractAddr)
	return Indexer{
		store:           store,
		l1Cli:           l1Cli,
		l1ConfirmBlocks: l1ConfirmBlocks,
		sccContractAddr: address,
		taskInterval:    taskIntervalDur,
		stopChan:        make(chan struct{}),
	}, nil
}

func (o Indexer) SetHook(hook Hook) Indexer {
	o.hook = hook
	return o
}

func (o Indexer) Start() {
	scannedHeight, err := o.store.GetScannedHeight()
	if err != nil {
		panic(err)
	}
	go o.ObserveStateBatchAppended(scannedHeight)
}

func (o Indexer) Stop() {
	close(o.stopChan)
}

func (o Indexer) ObserveStateBatchAppended(scannedHeight uint64) {
	queryTicker := time.NewTicker(o.taskInterval)
	for {
		func() {
			currentHeader, err := o.l1Cli.HeaderByNumber(context.Background(), nil)
			if err != nil {
				log.Error("failed to call layer1 HeaderByNumber", err)
				return
			}
			latestConfirmedBlockHeight := currentHeader.Number.Uint64() - uint64(o.l1ConfirmBlocks)

			startHeight := scannedHeight + 1
			endHeight := startHeight + scanRange
			if latestConfirmedBlockHeight < endHeight {
				endHeight = latestConfirmedBlockHeight
			}
			events, err := FilterStateBatchAppendedEvent(o.l1Cli, int64(startHeight), int64(endHeight), o.sccContractAddr)
			if err != nil {
				log.Error("failed to scan stateBatchAppended event", err)
				return
			}

			if len(events) != 0 {
				for _, event := range events {
					for stateBatchRoot, batchIndex := range event {
						var retry bool
						for !retry {
							retry = indexBatch(o.store, stateBatchRoot, batchIndex)
						}
						if err := o.hook.AfterStateBatchIndexed(stateBatchRoot); err != nil {
							log.Error("errors occur when executed hook AfterStateBatchIndexed", "err", err)
						}
					}
				}
			}

			scannedHeight = endHeight
			for err != nil { // retry until update successfully
				if err = o.store.UpdateHeight(scannedHeight); err != nil {
					log.Error("failed to update scannedHeight, retry", err)
					time.Sleep(2 * time.Second)
				}
			}
		}()

		select {
		case <-o.stopChan:
			return
		case <-queryTicker.C:
		}

	}
}

func indexBatch(store StateBatchStore, stateBatchRoot [32]byte, batchIndex uint64) (retry bool) {
	found, stateBatch := store.GetStateBatch(stateBatchRoot)
	if !found {
		log.Error("can not find the state batch with root, skip this batch", "root", hexutil.Encode(stateBatchRoot[:]))
		return false
	}
	stateBatch.BatchIndex = batchIndex
	if err := store.SetStateBatch(stateBatch); err != nil { // update stateBatch with index
		log.Error("failed to SetStateBatch with index", err)
		time.Sleep(2 * time.Second)
		return true
	}

	if err := store.IndexStateBatch(batchIndex, stateBatchRoot); err != nil {
		log.Error("failed to IndexStateBatch", err)
		time.Sleep(2 * time.Second)
		return true
	}
	return false
}

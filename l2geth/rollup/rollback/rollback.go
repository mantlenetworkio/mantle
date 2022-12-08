package rollback

import "github.com/mantlenetworkio/mantle/l2geth/eth/gasprice"

type EthAPIBackend struct {
	extRPCEnabled   bool
	eth             *Ethereum
	gpo             *gasprice.Oracle
	rollupGpo       *gasprice.RollupOracle
	verifier        bool
	gasLimit        uint64
	UsingBVM        bool
	MaxCallDataSize int
}

func (b *EthAPIBackend) SetHead(number uint64) {
	if number == 0 {
		log.Info("Cannot reset to genesis")
		return
	}
	if !b.UsingBVM {
		b.eth.protocolManager.downloader.Cancel()
	}
	b.eth.blockchain.SetHead(number)

	// Make sure to reset the LatestL1{Timestamp,BlockNumber}
	block := b.eth.blockchain.CurrentBlock()
	txs := block.Transactions()
	if len(txs) == 0 {
		log.Error("No transactions found in block", "number", number)
		return
	}

	tx := txs[0]
	blockNumber := tx.L1BlockNumber()
	if blockNumber == nil {
		log.Error("No L1BlockNumber found in transaction", "number", number)
		return
	}

	b.eth.syncService.SetLatestL1Timestamp(tx.L1Timestamp())
	b.eth.syncService.SetLatestL1BlockNumber(blockNumber.Uint64())
}

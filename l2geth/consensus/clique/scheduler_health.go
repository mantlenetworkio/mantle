package clique

import (
	"fmt"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/consensus/clique/synchronizer"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

type sequencerHealther struct {
	SeqSet           synchronizer.SequencerSequencerInfos
	SequencersPoints map[common.Address]uint64
}

var (
	// TODO set by contract
	initPoints uint64 = 6
)

func NewSequencerHealther() *sequencerHealther {
	return &sequencerHealther{}
}

// SetSequencerHealthChecker when update sequencerSet sequencerHealthChecker will be reset
func (schedulerInst *Scheduler) SetSequencerHealthChecker(seqSets synchronizer.SequencerSequencerInfos) {
	schedulerInst.sequencerHealther.SequencersPoints = make(map[common.Address]uint64)
	schedulerInst.sequencerHealther.SeqSet = seqSets
	for _, seqSet := range seqSets {
		schedulerInst.sequencerHealther.SequencersPoints[common.Address(seqSet.MintAddress)] = initPoints
	}
	for key := range schedulerInst.sequencerHealther.SequencersPoints {
		schedulerInst.sequencerHealther.SequencersPoints[key] = initPoints
	}
}

func (schedulerInst *Scheduler) checkSequencer() {
	blockNumber := schedulerInst.blockchain.CurrentHeader().Number.Uint64()
	sequencer := schedulerInst.currentStartMsg.Sequencer
	if (blockNumber - schedulerInst.currentStartMsg.StartHeight) >= schedulerInst.expectMinTxsCount {
		return
	}
	// deduct points
	schedulerInst.deductPoints(sequencer)
	return
}

// punishSequencer Check the health score of the current Sequencer. If
// the health score reaches the lower limit, the Sequencer is removed
// from the collection of block producers on the day
func (schedulerInst *Scheduler) punishSequencer(sequencer common.Address) {
	var newSeqSet synchronizer.SequencerSequencerInfos
	for _, seqSet := range schedulerInst.sequencerHealther.SeqSet {
		if seqSet.MintAddress.String() == sequencer.String() {
			continue
		}
		newSeqSet = append(newSeqSet, seqSet)
	}
	delete(schedulerInst.sequencerHealther.SequencersPoints, sequencer)
	// get changes
	changes := compareSequencerSet(schedulerInst.sequencerSet.Sequencers, newSeqSet)
	log.Debug(fmt.Sprintf("Get sequencer set success, have changes: %d", len(changes)))

	// update sequencer set and consensus_engine
	schedulerInst.l.Lock()
	defer schedulerInst.l.Unlock()
	err := schedulerInst.sequencerSet.UpdateWithChangeSet(changes)
	if err != nil {
		log.Error("sequencer set update failed", "err", err)
		return
	}
}

func (schedulerInst *Scheduler) deductPoints(sequencer common.Address) {
	if schedulerInst.zeroPoints(sequencer) {
		schedulerInst.punishSequencer(sequencer)
		return
	}
	schedulerInst.sequencerHealther.SequencersPoints[sequencer] = schedulerInst.sequencerHealther.SequencersPoints[sequencer] - 1
	return
}

func (schedulerInst *Scheduler) zeroPoints(sequencer common.Address) bool {
	return schedulerInst.sequencerHealther.SequencersPoints[sequencer] == 0
}

func (schedulerInst *Scheduler) GetExpectMinTxsCount(batchSize uint64) (uint64, error) {
	var pendingTxCount uint64
	pendingTxs, err := schedulerInst.txpool.Pending()
	if err != nil {
		return 0, err
	}
	for _, txs := range pendingTxs {
		pendingTxCount += uint64(len(txs))
	}
	if pendingTxCount > batchSize {
		pendingTxCount = batchSize
	}
	return pendingTxCount, nil
}

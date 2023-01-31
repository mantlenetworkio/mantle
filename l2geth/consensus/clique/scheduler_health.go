package clique

import (
	"math/big"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/consensus/clique/synchronizer"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

type healthAssessor struct {
	SeqSet           synchronizer.SequencerSequencerInfos
	SequencersPoints map[common.Address]uint64
}

var (
	// TODO set by config
	initPoints uint64 = 6
)

func NewHealthAssessor() *healthAssessor {
	return &healthAssessor{}
}

// setSequencerHealthPoints when update sequencerSet sequencerHealthChecker will be reset
func (schedulerInst *Scheduler) setSequencerHealthPoints(seqSets synchronizer.SequencerSequencerInfos) {
	schedulerInst.sequencerAssessor.SequencersPoints = make(map[common.Address]uint64)
	schedulerInst.sequencerAssessor.SeqSet = seqSets
	for _, seqSet := range seqSets {
		schedulerInst.sequencerAssessor.SequencersPoints[common.Address(seqSet.MintAddress)] = initPoints
	}
}

// checkSequencer checks the working status of the previous producer at the beginning
// of the new batch and deduct points of the sequencer that did not achieve the expected goal
func (schedulerInst *Scheduler) checkSequencer() {
	sequencer := schedulerInst.currentStartMsg.Sequencer
	if sequencer.String() == (common.Address{}).String() {
		// just return for first running
		return
	}
	blockNumber := schedulerInst.blockchain.CurrentHeader().Number.Uint64()
	log.Info("check sequencer", "sequencer", sequencer.String(),
		"current_height", blockNumber,
		"last_batch_start_height", schedulerInst.currentStartMsg.StartHeight,
		"expect_min_txs_count", schedulerInst.expectMinTxsCount,
		"current_point", schedulerInst.sequencerAssessor.SequencersPoints[sequencer])

	if (blockNumber - schedulerInst.currentStartMsg.StartHeight + 1) >= schedulerInst.expectMinTxsCount {
		return
	}
	schedulerInst.deductPoints(sequencer)
	return
}

// punishSequencer Check the health score of the current Sequencer. If
// the health score reaches the lower limit, the Sequencer is removed
// from the collection of block producers on the day
func (schedulerInst *Scheduler) punishSequencer(sequencer common.Address) {
	if len(schedulerInst.sequencerSet.Sequencers) == 1 {
		log.Info("Only have one sequencer, don't punish it")
		return
	}
	log.Info("Punish sequencer", "sequencer", sequencer.String())
	var newSeqSet synchronizer.SequencerSequencerInfos
	for _, seqSet := range schedulerInst.sequencerAssessor.SeqSet {
		if seqSet.MintAddress.String() == sequencer.String() {
			seqSet.Amount = big.NewInt(0)
		}
		newSeqSet = append(newSeqSet, seqSet)
	}
	delete(schedulerInst.sequencerAssessor.SequencersPoints, sequencer)
	changes := compareSequencerSet(schedulerInst.sequencerSet.Sequencers, newSeqSet)

	// update sequencer set and consensus_engine
	schedulerInst.sequencerSetMtx.Lock()
	defer schedulerInst.sequencerSetMtx.Unlock()
	err := schedulerInst.sequencerSet.UpdateWithChangeSet(changes)
	if err != nil {
		log.Error("sequencer set update failed", "err", err)
		return
	}
}

// deductPoints deduct points for the specified sequencer
func (schedulerInst *Scheduler) deductPoints(sequencer common.Address) {
	if schedulerInst.sequencerAssessor.SequencersPoints[sequencer] > 0 {
		schedulerInst.sequencerAssessor.SequencersPoints[sequencer] = schedulerInst.sequencerAssessor.SequencersPoints[sequencer] - 1
	} else {
		log.Error("try to deduct point from sequencer with zero point", "sequencer", sequencer.String())
		return
	}
	if schedulerInst.zeroPoints(sequencer) {
		schedulerInst.punishSequencer(sequencer)
	}
	return
}

func (schedulerInst *Scheduler) zeroPoints(sequencer common.Address) bool {
	return schedulerInst.sequencerAssessor.SequencersPoints[sequencer] == 0
}

// getExpectMinTxsCount returns the minimum amount of transactions that the producer should
// complete according to the number of transactions contained in the current txpool
func (schedulerInst *Scheduler) getExpectMinTxsCount(batchSize uint64) (uint64, error) {
	pendingTxCount, err := schedulerInst.verifiedTxCount()
	if err != nil {
		return 0, err
	}
	if pendingTxCount > batchSize {
		pendingTxCount = batchSize
	}
	return pendingTxCount, nil
}

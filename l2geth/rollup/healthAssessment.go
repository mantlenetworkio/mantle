package rollup

import (
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"sync"
)

type SequencerHealthChecker struct {
	BlockNumber      uint64
	TxsMinimumLength uint64
	SequencerHealth
}

type SequencerHealth struct {
	SequencersPoints map[common.Address]uint64
	pointsLock       sync.Mutex
}

var sequencersHealth *SequencerHealth

const InitPoints = 6

var SequencersPoints map[common.Address]uint64

var sequencerHealthChecker = &SequencerHealthChecker{}

func (s *SyncService) InitSequencerHealthChecker(sequencers []common.Address) {
	sequencerHealthChecker.SequencerHealth.SequencersPoints = make(map[common.Address]uint64)
	for _, sequencer := range sequencers {
		sequencerHealthChecker.SequencerHealth.SequencersPoints[sequencer] = InitPoints
	}
	s.resetPoints()
}

//func (s *SyncService) checker() {
//	ticker := time.NewTicker(time.Minute * 1)
//	for range ticker.C {
//		s.CheckBlockState()
//	}
//}

func (s *SyncService) PreCheck(sequencer common.Address) (bool, error) {
	return s.CheckBlockState(sequencer)

}

func (s *SyncService) CheckBlockState(sequencer common.Address) (bool, error) {
	blockNumber := s.bc.CurrentHeader().Number.Uint64()
	// TODO
	var maxTxsLength uint64
	if (blockNumber - sequencerHealthChecker.BlockNumber) >= sequencerHealthChecker.TxsMinimumLength {
		s.Reset(maxTxsLength)
		return true, nil
	} else {
		if health := s.SelfHealthCheck(); !health {
			// TODO Alert or whatever
			return true, nil
		} else {
			s.deductPoints(sequencer)
			// deduct points
			s.deductPoints(sequencer)
			s.Reset(maxTxsLength)
		}
	}
	return true, nil
}

func checkoutTxsLength() {

}

func (s *SyncService) Reset(maxTxsLength uint64) error {
	// TODO reset
	s.resetBlockState(maxTxsLength)
	return nil
}

// Reset block update time, PendingTx status
func (s *SyncService) resetBlockState(maxTxsLenght uint64) {
	var pendingTxCount uint64
	pendingTxs, err := s.txpool.Pending()
	if err != nil {

	}
	for _, txs := range pendingTxs {
		pendingTxCount += uint64(len(txs))
	}
	var blockState SequencerHealthChecker
	if pendingTxCount >= maxTxsLenght {
		pendingTxCount = maxTxsLenght
	}
	blockState.TxsMinimumLength = pendingTxCount
	blockState.BlockNumber = s.bc.CurrentHeader().Number.Uint64()
}

// SelfHealthCheck Self health check
func (s *SyncService) SelfHealthCheck() bool {
	// TODO
	return true
}

// punishSequencer Check the health score of the current Sequencer. If
// the health score reaches the lower limit, the Sequencer is removed
// from the collection of block producers on the day
func (s *SyncService) punishSequencer() {
	// TODO handle L1 sequencer contract
	//
}

func (s *SyncService) deductPoints(sequencer common.Address) {
	if s.checkPoints(sequencer) {
		// TODO
		s.punishSequencer()
		return
	}
	SequencersPoints[sequencer] = SequencersPoints[sequencer] - 1
	return
}

func (s *SyncService) checkPoints(sequencer common.Address) bool {
	return SequencersPoints[sequencer] == 0
}

func (s *SyncService) resetPoints() error {
	for key := range sequencersHealth.SequencersPoints {
		sequencersHealth.pointsLock.Lock()
		SequencersPoints[key] = InitPoints
		sequencersHealth.pointsLock.Unlock()
	}
}

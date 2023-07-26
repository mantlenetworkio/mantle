package txmgr

import (
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
)

type SendState struct {
	minedTxs         map[common.Hash]struct{}
	nonceTooLowCount uint64
	mu               sync.RWMutex

	safeAbortNonceTooLowCount uint64
}

func NewSendState(safeAbortNonceTooLowCount uint64) *SendState {
	if safeAbortNonceTooLowCount == 0 {
		panic("txmgr: safeAbortNonceTooLowCount cannot be zero")
	}

	return &SendState{
		minedTxs:                  make(map[common.Hash]struct{}),
		nonceTooLowCount:          0,
		safeAbortNonceTooLowCount: safeAbortNonceTooLowCount,
	}
}

func (s *SendState) ProcessSendError(err error) {
	if err == nil {
		return
	}

	if !strings.Contains(err.Error(), core.ErrNonceTooLow.Error()) {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.nonceTooLowCount++
}

func (s *SendState) TxMined(txHash common.Hash) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.minedTxs[txHash] = struct{}{}
}

func (s *SendState) TxNotMined(txHash common.Hash) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, wasMined := s.minedTxs[txHash]
	delete(s.minedTxs, txHash)

	if len(s.minedTxs) == 0 && wasMined {
		s.nonceTooLowCount = 0
	}
}

func (s *SendState) ShouldAbortImmediately() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.minedTxs) > 0 {
		return false
	}

	return s.nonceTooLowCount >= s.safeAbortNonceTooLowCount
}

func (s *SendState) IsWaitingForConfirmation() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.minedTxs) > 0
}

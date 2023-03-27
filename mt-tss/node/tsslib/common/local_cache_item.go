package common

import (
	"sync"

	"github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/messages"
)

type LocalCacheItem struct {
	Msg           *messages.WireMessage
	Hash          string
	lock          *sync.Mutex
	ConfirmedList map[string]string
}

func NewLocalCacheItem(msg *messages.WireMessage, hash string) *LocalCacheItem {
	return &LocalCacheItem{
		Msg:           msg,
		Hash:          hash,
		lock:          &sync.Mutex{},
		ConfirmedList: make(map[string]string),
	}
}

// UpdateConfirmList add the given party's hash into the confirm list
func (l *LocalCacheItem) UpdateConfirmList(P2PID, hash string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.ConfirmedList[P2PID] = hash
}

// TotalConfirmParty number of parties that already confirmed their hash
func (l *LocalCacheItem) TotalConfirmParty() int {
	l.lock.Lock()
	defer l.lock.Unlock()
	return len(l.ConfirmedList)
}

func (l *LocalCacheItem) GetPeers() []string {
	peers := make([]string, 0, len(l.ConfirmedList))
	l.lock.Lock()
	defer l.lock.Unlock()
	for peer := range l.ConfirmedList {
		peers = append(peers, peer)
	}
	return peers
}

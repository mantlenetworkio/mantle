package abnormal

import "sync"

type ShareMgr struct {
	requested map[string]bool
	reqLocker *sync.Mutex
}

func NewTssShareMgr() *ShareMgr {
	return &ShareMgr{
		reqLocker: &sync.Mutex{},
		requested: make(map[string]bool),
	}
}

func (sm *ShareMgr) Set(key string) {
	sm.reqLocker.Lock()
	defer sm.reqLocker.Unlock()
	sm.requested[key] = true
}

func (sm *ShareMgr) QueryAndDelete(key string) bool {
	sm.reqLocker.Lock()
	defer sm.reqLocker.Unlock()
	ret := sm.requested[key]
	delete(sm.requested, key)
	return ret
}

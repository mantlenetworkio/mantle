package abnormal

import (
	"errors"
	"sync"

	"github.com/bnb-chain/tss-lib/tss"
)

const (
	HashCheckFail       = "hash check failed"
	TssTimeout          = "Tss timeout"
	TssSyncFail         = "signers fail to sync before keygen/keysign"
	TssBrokenMsg        = "tss share verification failed"
	InternalError       = "fail to start the join party "
	GenerateNewKeyError = "fail to generate new key"
	SignatureError      = "fail to signature message"
)

var (
	ErrHashFromOwner     = errors.New(" hash sent from data owner")
	ErrNotEnoughPeer     = errors.New("not enough nodes to evaluate hash")
	ErrNotMajority       = errors.New("message we received does not match the majority")
	ErrTssTimeOut        = errors.New("error Tss Timeout")
	ErrHashCheck         = errors.New("error in processing hash check")
	ErrHashInconsistency = errors.New("fail to agree on the hash value")
)

type PartyInfo struct {
	Party      tss.Party
	PartyIDMap map[string]*tss.PartyID
}

type Node struct {
	Pubkey    string `json:"pubkey"`
	Data      []byte `json:"data"`
	Signature []byte `json:"signature,omitempty"`
}

type Abnormal struct {
	FailReason   string `json:"fail_reason"`
	IsUnicast    bool   `json:"is_broadcast"`
	Nodes        []Node `json:"abnormal_peers,omitempty"`
	AbnormalLock *sync.RWMutex
}

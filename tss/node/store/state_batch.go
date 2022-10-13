package store

import (
	"encoding/json"
	"github.com/mantlenetworkio/mantle/tss/index"
)

func (s *Storage) SetStateBatch(info index.StateBatchInfo) error {
	bz, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return s.db.Put(getStateBatchKey(info.BatchRoot), bz, nil)
}

func (s *Storage) GetStateBatch(root [32]byte) (bool, index.StateBatchInfo) {
	bz, err := s.db.Get(getStateBatchKey(root), nil)
	if err != nil {
		return handleError2(index.StateBatchInfo{}, err)
	}
	var sbi index.StateBatchInfo
	if err = json.Unmarshal(bz, &sbi); err != nil {
		return true, index.StateBatchInfo{}
	}
	return true, sbi
}

func (s *Storage) IndexStateBatch(index uint64, root [32]byte) error {
	return s.db.Put(getIndexStateBatchKey(index), root[:], nil)
}

func (s *Storage) GetIndexStateBatch(index uint64) (bool, [32]byte) {
	bz, err := s.db.Get(getIndexStateBatchKey(index), nil)
	if err != nil {
		return handleError2([32]byte{}, err)
	}
	var stateRoot [32]byte
	copy(stateRoot[:], bz)
	return true, stateRoot
}

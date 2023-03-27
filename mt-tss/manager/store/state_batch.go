package store

import (
	"encoding/json"

	"github.com/mantlenetworkio/mantle/mt-tss/index"
)

func (s *Storage) SetOutput(info index.OutputInfo) error {
	bz, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return s.db.Put(getOutputKey(info.OutputRoot), bz, nil)
}

func (s *Storage) GetOutput(root [32]byte) (bool, index.OutputInfo) {
	bz, err := s.db.Get(getOutputKey(root), nil)
	if err != nil {
		return handleError2(index.OutputInfo{}, err)
	}
	var sbi index.OutputInfo
	if err = json.Unmarshal(bz, &sbi); err != nil {
		return true, index.OutputInfo{}
	}
	return true, sbi
}

func (s *Storage) IndexOutput(index uint64, root [32]byte) error {
	return s.db.Put(getIndexOutputKey(index), root[:], nil)
}

func (s *Storage) GetIndexStateBatch(index uint64) (bool, [32]byte) {
	bz, err := s.db.Get(getIndexOutputKey(index), nil)
	if err != nil {
		return handleError2([32]byte{}, err)
	}
	var stateRoot [32]byte
	copy(stateRoot[:], bz)
	return true, stateRoot
}

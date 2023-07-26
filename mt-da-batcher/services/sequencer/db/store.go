package db

import (
	"github.com/ethereum/go-ethereum/log"
)

type Store struct {
	db *LevelDBStore
}

func NewStore(path string) (*Store, error) {
	db, err := NewLevelDBStore(path)
	if err != nil {
		log.Info("Could not create leveldb database.")
		return nil, err
	}
	return &Store{
		db: db,
	}, nil
}

func (s *Store) GetReRollupBatchIndex() (uint64, bool) {
	key := []byte("ReRollupBatchIndex")
	data, err := s.db.Get(key)
	if err != nil {
		return 0, false
	}
	bn := toUint64(data)
	return bn, true
}

func (s *Store) SetReRollupBatchIndex(bn uint64) bool {
	key := []byte("ReRollupBatchIndex")
	data := toByteArray(bn)
	err := s.db.Put(key, data)
	return err == nil
}

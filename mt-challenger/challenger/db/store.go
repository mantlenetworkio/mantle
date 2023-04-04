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

func (s *Store) GetLatestBatchIndex() (uint64, bool) {
	key := []byte("BatchIndex")
	data, err := s.db.Get(key)
	if err != nil {
		return 0, false
	}
	bn := toUint64(data)
	return bn, true
}

func (s *Store) SetLatestBatchIndex(bn uint64) bool {
	key := []byte("BatchIndex")
	data := toByteArray(bn)
	err := s.db.Put(key, data)
	return err == nil
}

func (s *Store) GetReRollupIndex() (uint64, bool) {
	key := []byte("ReRollupIndex")
	data, err := s.db.Get(key)
	if err != nil {
		return 0, false
	}
	bn := toUint64(data)
	return bn, true
}

func (s *Store) SetLatestReRollupIndex(bn uint64) bool {
	key := []byte("ReRollupIndex")
	data := toByteArray(bn)
	err := s.db.Put(key, data)
	return err == nil
}

func (s *Store) GetToolReRollupIndex() (uint64, bool) {
	key := []byte("ToolReRollupIndex")
	data, err := s.db.Get(key)
	if err != nil {
		return 0, false
	}
	bn := toUint64(data)
	return bn, true
}

func (s *Store) SetToolReRollupIndex(bn uint64) bool {
	key := []byte("ToolReRollupIndex")
	data := toByteArray(bn)
	err := s.db.Put(key, data)
	return err == nil
}

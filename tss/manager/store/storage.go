package store

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

type Storage struct {
	db *leveldb.DB
}

func NewStorage(levelDbFolder string) (*Storage, error) {
	var db *leveldb.DB
	var err error
	if len(levelDbFolder) == 0 {
		// no directory given, use in memory store
		mStore := storage.NewMemStorage()
		db, err = leveldb.Open(mStore, nil)
		if err != nil {
			return nil, fmt.Errorf("fail to in memory open level db: %w", err)
		}
	} else {
		db, err = leveldb.OpenFile(levelDbFolder, nil)
		if err != nil {
			return nil, fmt.Errorf("fail to open level db %s: %w", levelDbFolder, err)
		}
	}
	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

package db

import (
	"encoding/binary"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBStore struct {
	*leveldb.DB
}

func (d *LevelDBStore) Put(key []byte, value []byte) error {
	return d.DB.Put(key, value, nil)
}

func (d *LevelDBStore) Get(key []byte) ([]byte, error) {
	return d.DB.Get(key, nil)
}

func (d *LevelDBStore) Delete(key []byte) error {
	return d.DB.Delete(key, nil)
}

func NewLevelDBStore(path string) (*LevelDBStore, error) {
	handle, err := leveldb.OpenFile(path, nil)
	return &LevelDBStore{handle}, err
}

func toByteArray(i uint64) []byte {
	arr := make([]byte, 8)
	binary.BigEndian.PutUint64(arr[0:8], uint64(i))
	return arr
}

func toUint64(arr []byte) uint64 {
	i := binary.BigEndian.Uint64(arr)
	return i
}

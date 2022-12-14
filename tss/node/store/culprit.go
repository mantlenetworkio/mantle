package store

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
)

func (s *Storage) AddCulprits(culprits []string) {
	bz, err := s.db.Get(getCulpritsKey(), nil)
	if err != nil && err != leveldb.ErrNotFound {
		panic(err)
	}
	if len(bz) > 0 {
		var data []string
		err = json.Unmarshal(bz, &data)
		if err != nil {
			panic(err)
		}
		culprits = append(culprits, data...)
	}
	bz, err = json.Marshal(culprits)
	if err != nil {
		panic(err)
	}
	if err = s.db.Put(getCulpritsKey(), bz, nil); err != nil {
		panic(err)
	}
}

func (s *Storage) GetCulprits() []string {
	bz, err := s.db.Get(getCulpritsKey(), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil
		}
		panic(err)
	}
	var ret []string
	if err = json.Unmarshal(bz, &ret); err != nil {
		panic(err)
	}
	return ret
}

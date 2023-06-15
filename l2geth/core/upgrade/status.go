package upgrade

import (
	"github.com/mantlenetworkio/mantle/l2geth/ethdb"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

/*
	For upgrade based on L1 block number,
*/

var (
	mockUpgradeFlag = []byte("mockUpgradeFlag")
)

func ReadUpgradeFlag(db ethdb.Reader, key []byte) []byte {
	exist, err := db.Has(key)
	if err != nil {
		log.Error("Failed to get upgrade status", "err", err)
	}
	if !exist {
		return nil
	}
	data, err := db.Get(key)
	if err != nil {
		log.Error("Failed to get upgrade status", "err", err)
	}
	return data
}

func ExistUpgradeFlag(db ethdb.Reader, key []byte) bool {
	exist, err := db.Has(key)
	if err != nil {
		log.Error("Failed to check upgrade status", "err", err)
	}

	return exist
}

func WriteUpgradeFlag(db ethdb.Writer, key, data []byte) {
	if err := db.Put(key, data); err != nil {
		log.Error("Failed to write upgrade status", "err", err)
	}
}

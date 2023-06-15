package upgrade

import (
	"github.com/mantlenetworkio/mantle/l2geth/ethdb"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

/*
	For upgrade based on L1 block number, record if the upgrade took place.
*/

var (
	mockUpgradeFlag = []byte("mockUpgradeFlag")
)

func readUpgradeFlag(db ethdb.Reader, key []byte) []byte {
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

func existUpgradeFlag(db ethdb.Reader, key []byte) bool {
	exist, err := db.Has(key)
	if err != nil {
		log.Error("Failed to check upgrade status", "err", err)
	}

	return exist
}

func writeUpgradeFlag(db ethdb.Writer, key, data []byte) {
	if err := db.Put(key, data); err != nil {
		log.Error("Failed to write upgrade status", "err", err)
	}
}

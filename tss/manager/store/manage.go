package store

import (
	"encoding/json"
	"github.com/bitdao-io/bitnetwork/l2geth/common"
	"github.com/bitdao-io/bitnetwork/tss/manager/types"
	"github.com/bitdao-io/bitnetwork/tss/slash"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func (s *Storage) SetSlashingTx(info slash.SlashingInfo, txBytes []byte) {
	tx := types.SlashingTx{
		Info:    info,
		TxBytes: txBytes,
	}
	bz, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	if err = s.db.Put(getSlashingTxKey(info.Address, info.BatchIndex), bz, nil); err != nil {
		panic(err)
	}
}

func (s *Storage) ListSlashingTx() (slashingTxs []types.SlashingTx) {
	iterator := s.db.NewIterator(util.BytesPrefix(SlashingTxKeyPrefix), nil)
	defer iterator.Release()
	for iterator.Next() {
		buf := iterator.Value()
		if len(buf) == 0 {
			continue
		}
		var slashingTx types.SlashingTx
		if err := json.Unmarshal(iterator.Value(), &slashingTx); err != nil {
			panic(err)
		}
		slashingTxs = append(slashingTxs, slashingTx)
	}
	return
}

func (s *Storage) RemoveSlashingTx(address common.Address, batchIndex uint64) {
	if err := s.db.Delete(getSlashingTxKey(address, batchIndex), nil); err != nil {
		panic(err)
	}
}

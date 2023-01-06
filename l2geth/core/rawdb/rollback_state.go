package rawdb

import (
	"bytes"

	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/ethdb"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/rlp"
)

// WriteRollbackStates stores a RollbackStates into the database
func WriteRollbackStates(db ethdb.KeyValueWriter, rollbackStates types.RollbackStates) {
	//// Write the encoded rollbackStates
	data, err := rlp.EncodeToBytes(rollbackStates)
	if err != nil {
		log.Crit("Failed to RLP encode rollbackStates", "err", err)
	}
	//key := headerKey(number, hash)
	if err := db.Put(rollbackStatesKey, data); err != nil {
		log.Crit("Failed to store rollbackStates", "err", err)
	}
}

// ReadRollbackStates will read rollback states.
func ReadRollbackStates(db ethdb.KeyValueReader) types.RollbackStates {
	data, _ := db.Get(rollbackStatesKey)
	if len(data) == 0 {
		return types.RollbackStates{}
	}
	var rollbackStates types.RollbackStates
	if err := rlp.Decode(bytes.NewReader(data), rollbackStates); err != nil {
		log.Error("Invalid rollbackStates RLP", "err", err)
		return nil
	}
	return rollbackStates
}

// ReadRollbackStates will read rollback states.
func ReadLatestRollbackState(db ethdb.KeyValueReader) *types.RollbackState {
	data, _ := db.Get(rollbackStatesKey)
	if len(data) == 0 {
		return &types.RollbackState{}
	}
	var rollbackStates types.RollbackStates
	if err := rlp.Decode(bytes.NewReader(data), rollbackStates); err != nil {
		log.Error("Invalid rollbackStates RLP", "err", err)
		return nil
	}
	return rollbackStates[len(rollbackStates)-1]
}

package ether

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/state"
)

// getBVMETHTotalSupply returns BVM ETH's total supply by reading
// the appropriate storage slot.
func getBVMETHTotalSupply(db *state.StateDB) *big.Int {
	key := getBVMETHTotalSupplySlot()
	return db.GetState(BVMETHAddress, key).Big()
}

func getBVMETHTotalSupplySlot() common.Hash {
	position := common.Big2
	key := common.BytesToHash(common.LeftPadBytes(position.Bytes(), 32))
	return key
}

func GetBVMETHTotalSupplySlot() common.Hash {
	return getBVMETHTotalSupplySlot()
}

// getBVMETHBalance gets a user's BVM ETH balance from state by querying the
// appropriate storage slot directly.
func getBVMETHBalance(db *state.StateDB, addr common.Address) *big.Int {
	return db.GetState(BVMETHAddress, CalcBVMETHStorageKey(addr)).Big()
}

package tssreward

import (
	"github.com/bitdao-io/bitnetwork/l2geth/common"
	"github.com/bitdao-io/bitnetwork/l2geth/core/rawdb"
	"github.com/bitdao-io/bitnetwork/l2geth/core/state"
	"github.com/bitdao-io/bitnetwork/l2geth/core/vm"
	"github.com/bitdao-io/bitnetwork/l2geth/params"
	"github.com/bitdao-io/bitnetwork/l2geth/rollup/dump"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestUpdateTssRewardData(t *testing.T) {
	blockNumber := big.NewInt(100)
	fee := big.NewInt(10000)
	data, err := PacketData(blockNumber, fee)
	require.NoError(t, err)
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()))
	vmctx := vm.Context{
		CanTransfer: func(vm.StateDB, common.Address, *big.Int) bool { return true },
		Transfer:    func(vm.StateDB, common.Address, common.Address, *big.Int) {},
	}
	vmenv := vm.NewEVM(vmctx, statedb, params.AllEthashProtocolChanges, vm.Config{ExtraEips: []int{2200}})
	zeroAddress := vm.AccountRef(common.Address{})
	_, _, err = vmenv.Call(zeroAddress, dump.L2ExcuteFeeWallet, data, 0, big.NewInt(0))
}

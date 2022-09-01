package types

import (
	"github.com/bitdao-io/bitnetwork/tss/index"
	"github.com/bitdao-io/bitnetwork/tss/slash"
)

type NodeStore interface {
	index.StateBatchStore
	index.ScanHeightStore
	slash.SlashingStore
}

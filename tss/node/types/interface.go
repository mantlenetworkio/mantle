package types

import (
	"github.com/bitdao-io/mantle/tss/index"
	"github.com/bitdao-io/mantle/tss/slash"
)

type NodeStore interface {
	index.StateBatchStore
	index.ScanHeightStore
	slash.SlashingStore
}

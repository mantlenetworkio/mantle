package types

import (
	"github.com/mantlenetworkio/mantle/tss/index"
	"github.com/mantlenetworkio/mantle/tss/slash"
)

type NodeStore interface {
	index.StateBatchStore
	index.ScanHeightStore
	slash.SlashingStore
}

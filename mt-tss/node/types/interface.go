package types

import (
	"github.com/mantlenetworkio/mantle/mt-tss/index"
	"github.com/mantlenetworkio/mantle/mt-tss/slash"
)

type NodeStore interface {
	index.OutputStore
	index.ScanHeightStore
	slash.SlashingStore
}

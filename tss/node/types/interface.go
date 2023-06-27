package types

import (
	"github.com/mantlenetworkio/mantle/tss/index"
	"github.com/mantlenetworkio/mantle/tss/slash"
)

type TssMemberStore interface {
	SetInactiveMembers(TssMembers) error
	GetInactiveMembers() (bool, TssMembers)
	SetActiveMembers(TssMembers) error
	GetActiveMembers() (bool, TssMembers)
}

type NodeStore interface {
	index.StateBatchStore
	index.ScanHeightStore
	slash.SlashingStore
	TssMemberStore
}

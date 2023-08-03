package types

import (
	"github.com/mantlenetworkio/mantle/tss/index"
	"github.com/mantlenetworkio/mantle/tss/slash"
)

type TssMemberStore interface {
	SetInactiveMembers(TssMembers) error
	GetInactiveMembers() (TssMembers, error)
	SetActiveMembers(TssMembers) error
	GetActiveMembers() (TssMembers, error)
}

type NodeStore interface {
	index.StateBatchStore
	index.ScanHeightStore
	slash.SlashingStore
	TssMemberStore
}

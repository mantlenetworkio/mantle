package types

import (
	"github.com/tenderly/optimism/tss/index"
	"github.com/tenderly/optimism/tss/slash"
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

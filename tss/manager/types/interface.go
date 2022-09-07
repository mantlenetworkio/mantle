package types

import (
	tss "github.com/bitdao-io/mantle/tss/common"
	"github.com/bitdao-io/mantle/tss/index"
	"github.com/bitdao-io/mantle/tss/slash"
)

type SignService interface {
	SignStateBatch(request tss.SignStateRequest) ([]byte, error)
	SignTxBatch() error
}

type TssQueryService interface {
	QueryActiveInfo() (TssCommitteeInfo, error)
	QueryInactiveInfo() (TssCommitteeInfo, error)
}

type CPKStore interface {
	Insert(CpkData) error
	GetByElectionId(uint64) (CpkData, error)
}

type ManagerStore interface {
	CPKStore
	index.StateBatchStore
	index.ScanHeightStore
	slash.SlashingStore
}

package types

import (
	tss "github.com/bitdao-io/bitnetwork/tss/common"
	"github.com/bitdao-io/bitnetwork/tss/index"
	"github.com/bitdao-io/bitnetwork/tss/slash"
)

type SignService interface {
	SignStateBatch(request tss.SignStateRequest) ([]byte, error)
	SignTxBatch() error
}

type TssQueryService interface {
	QueryActiveInfo() TssCommitteeInfo
	QueryInactiveInfo() TssCommitteeInfo
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

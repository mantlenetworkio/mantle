package types

import (
	tss "github.com/bitdao-io/bitnetwork/tss/common"
	"github.com/bitdao-io/bitnetwork/tss/index"
)

type SignService interface {
	SignStateBatch(request tss.SignStateRequest) ([]byte, error)
	SignTxBatch() error
}

type TssQueryService interface {
	QueryInfo() TssCommitteeInfo
}

type CPKStore interface {
	Insert(CpkData) error
	GetByElectionId(uint64) (CpkData, error)
}

type ManagerStore interface {
	CPKStore
	index.StateBatchStore
	index.ScanHeightStore
}

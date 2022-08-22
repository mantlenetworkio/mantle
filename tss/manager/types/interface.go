package types

import (
	tss "github.com/bitdao-io/bitnetwork/tss/types"
)

type SignService interface {
	SignStateBatch(request tss.SignStateRequest) ([]byte, error)
	SignTxBatch() error
}

type TssQueryService interface {
	QueryInfo() TssInfos
}

type CPKStore interface {
	Insert(CpkData) error
	GetByElectionId(uint64) (CpkData, error)
}

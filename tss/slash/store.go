package slash

import (
	"github.com/ethereum/go-ethereum/common"
)

type SigningInfo struct {
	Address             common.Address `json:"address"`
	StartBatchIndex     uint64         `json:"start_batch_index"`
	IndexOffset         uint64         `json:"index_offset"`
	MissedBlocksCounter uint64         `json:"missed_blocks_counter"`
}

type SlashingInfo struct {
	Address    common.Address `json:"address"`
	BatchIndex uint64         `json:"batch_index"`
	ElectionId uint64         `json:"election_id"`
	SlashType  byte           `json:"slash_type"`
}

type SlashingStore interface {
	SetSigningInfo(SigningInfo)
	GetSigningInfo(common.Address) (bool, SigningInfo)

	GetNodeMissedBatchBitArray(common.Address, uint64) bool
	SetNodeMissedBatchBitArray(common.Address, uint64, bool)
	ClearNodeMissedBatchBitArray(common.Address)

	SetSlashingInfo(SlashingInfo)
	ListSlashingInfo() []SlashingInfo

	GetSlashingInfo(common.Address, uint64) (bool, SlashingInfo)
	IsInSlashing(common.Address) bool
	RemoveSlashingInfo(common.Address, uint64)

	AddCulprits([]string)
	GetCulprits() []string
}

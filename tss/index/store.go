package index

type StateBatchInfo struct {
	BatchRoot    [32]byte `json:"batch_root"`
	ElectionId   uint64   `json:"election_id"`
	AbsentNodes  []string `json:"absent_nodes"`
	WorkingNodes []string `json:"working_nodes"`
	BatchIndex   uint64   `json:"batch_index"`
}

type StateBatchStore interface {
	SetStateBatch(StateBatchInfo) error
	GetStateBatch([32]byte) (bool, StateBatchInfo)
	IndexStateBatch(uint64, [32]byte) error
	GetIndexStateBatch(index uint64) (bool, [32]byte)
}

type ScanHeightStore interface {
	UpdateHeight(uint64) error
	GetScannedHeight() (uint64, error)
}

type IndexerStore interface {
	StateBatchStore
	ScanHeightStore
}

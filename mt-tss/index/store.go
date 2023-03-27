package index

type OutputInfo struct {
	OutputRoot   [32]byte `json:"output_root"`
	ElectionId   uint64   `json:"election_id"`
	AbsentNodes  []string `json:"absent_nodes"`
	WorkingNodes []string `json:"working_nodes"`
	OutputIndex  uint64   `json:"output_index"`
}

type OutputStore interface {
	SetOutput(OutputInfo) error
	GetOutput([32]byte) (bool, OutputInfo)
	IndexOutput(uint64, [32]byte) error
	GetIndexOutput(index uint64) (bool, [32]byte)
}

type ScanHeightStore interface {
	UpdateHeight(uint64) error
	GetScannedHeight() (uint64, error)
}

type IndexerStore interface {
	OutputStore
	ScanHeightStore
}

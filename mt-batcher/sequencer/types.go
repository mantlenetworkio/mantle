package sequencer

import "math/big"

type AggregateSignature struct {
	AggSig           []byte
	AggPubKey        []byte
	NonSignerPubKeys [][]byte
}

type DisperseMeta struct {
	Sigs            AggregateSignature
	ApkIndex        uint32
	TotalStakeIndex uint64
	StoreNumber     uint32
}

type StoreParams struct {
	BlockNumber         uint32
	TotalOperatorsIndex uint32
	OrigDataSize        uint32 // unique nonce for each data store
	NumTotal            uint32 // total number data node active on chain
	Quorum              uint32 // minimal amount of signatures from data node
	NumSys              uint32 // number of data node which contains the systematic chunk
	NumPar              uint32 // number of data node which contains the parity chunk
	Duration            uint32 // duration which data is stored

	// Data and Encoding
	KzgCommit      []byte // elliptic curve kzg commitmetn
	LowDegreeProof []byte
	Degree         uint32   // degree of the polynomial
	TotalSize      uint64   // total size of the data
	Order          []uint32 // mapping for deciding the storer of each coded data chunk

	// Chain
	Fee        *big.Int
	HeaderHash []byte
	Disperser  []byte
}

package types

import "math/big"

type SignStateRequest struct {
	StartBlock          big.Int    `json:"start_block"`
	OffsetStartsAtIndex big.Int    `json:"offset_starts_at_index"`
	StateRoots          [][32]byte `json:"state_roots"`
}

type AskResponse struct {
	Result bool `json:"result"`
}

type NodeSignStateRequest struct {
	ClusterPublicKey string           `json:"cluster_public_key"`
	Timestamp        int64            `json:"timestamp"`
	Nodes            []string         `json:"nodes"`
	StateBatch       SignStateRequest `json:"state_batch"`
}

type SignResponse struct {
	Signature SignatureData `json:"signature"`
}

type KeygenRequest struct {
	Nodes     []string `json:"nodes"`
	Timestamp int64    `json:"timestamp"`
}

type KeygenResponse struct {
	ClusterPublicKey string `json:"cluster_public_key"`
}

type SignatureData struct {
	// Ethereum-style recovery byte; only the first byte is relevant
	SignatureRecovery []byte `json:"signature_recovery,omitempty"`
	// Signature components R, S
	R []byte `json:"r,omitempty"`
	S []byte `json:"s,omitempty"`
	// M represents the original message digest that was signed M
	M []byte `json:"m,omitempty"`
}

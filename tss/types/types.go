package types

import "math/big"

type AskStateRequest struct {
	StartBlock          big.Int    `json:"start_block"`
	OffsetStartsAtIndex big.Int    `json:"offset_starts_at_index"`
	StateRoots          [][32]byte `json:"state_roots"`
}

type AskResponse struct {
	Result bool `json:"result"`
}

type BatchSignRequest struct {
	Timestamp  int64         `json:"timestamp"`
	Signs      []SignRequest `json:"signs"`
	PoolPubKey string        `json:"pool_pub_key"`
}

type SignRequest struct {
	UniqueId string   `json:"unique_id"`
	Nodes    []string `json:"nodes"`
}

type SignResponse struct {
	UniqueId  string        `json:"unique_id"`
	Signature SignatureData `json:"signature"`
}

type KeygenRequest struct {
	Nodes     []string `json:"nodes"`
	Threshold int      `json:"threshold"`
}

type KeygenResponse struct {
	PoolPubKey string `json:"pool_pub_key"`
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

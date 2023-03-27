package tsslib

import (
	keygen2 "github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/keygen"
	keysign2 "github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/keysign"
)

type Server interface {
	Start() error
	Stop()
	GetLocalPeerID() string
	Keygen(req keygen2.Request) (keygen2.Response, error)
	KeySign(req keysign2.Request) (keysign2.Response, error)
	ExportPeerAddress() map[string]string
	GetParticipants(poolPubkey string) ([]string, error)
}

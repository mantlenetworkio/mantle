package keygen

import (
	"github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/common"
)

type Response struct {
	PubKey          string        `json:"pubKey"`
	Address         []byte        `json:"address"`
	PubKeyByte      []byte        `json:"pubKey_byte"`
	Status          common.Status `json:"status"`
	FailReason      string        `json:"fail_reason"`
	AbnormalPubKeys []string      `json:"abnormal_pub_keys"`
}

func NewResponse(pubkey string, address, pubkeyByte []byte, status common.Status, failReason string, abnormalPubkeys []string) Response {
	return Response{
		PubKey:          pubkey,
		Address:         address,
		PubKeyByte:      pubkeyByte,
		Status:          status,
		FailReason:      failReason,
		AbnormalPubKeys: abnormalPubkeys,
	}
}

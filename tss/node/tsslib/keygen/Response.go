package keygen

import (
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/common"
)

type Response struct {
	PubKey          string        `json:"pubKey"`
	Address         []byte        `json:"address"`
	Status          common.Status `json:"status"`
	FailReason      string        `json:"fail_reason"`
	AbnormalPubKeys []string      `json:"abnormal_pub_keys"`
}

func NewResponse(pubkey string, address []byte, status common.Status, failReason string, abnormalPubkeys []string) Response {
	return Response{
		PubKey:          pubkey,
		Address:         address,
		Status:          status,
		FailReason:      failReason,
		AbnormalPubKeys: abnormalPubkeys,
	}
}

package keysign

import (
	common2 "github.com/bnb-chain/tss-lib/common"
	"github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/common"
)

type Response struct {
	SignatureData *common2.SignatureData `json:"signature_data"`
	Status        common.Status          `json:"status"`
	FailReason    string                 `json:"failReason"`
	Culprits      []string               `json:"culprits"`
}

func NewResponse(signature *common2.SignatureData, status common.Status, failReason string, culprits []string) Response {
	return Response{
		SignatureData: signature,
		Status:        status,
		FailReason:    failReason,
		Culprits:      culprits,
	}
}

package clique

import (
	"encoding/hex"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"testing"
)

func TestProducerSerializeB(t *testing.T) {
	address := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	bps := &types.BatchPeriodStartMsg{
		ReorgIndex:   0,
		BatchIndex:   0,
		StartHeight:  1,
		MaxHeight:    100,
		ExpireTime:   1669787879,
		MinerAddress: address,
		SequencerSet: []common.Address{
			address,
		},
		Signature: common.Hex2Bytes("2020a0bbf67b08b1df594333c1ead3a771d9742d2f33798e050da744b1255bb67860d672a5055429cc53d17e6c57550989b39cf997e2fb58d1ec6aae198a471501"),
	}
	buf := bps.SerializeBatchPeriodStartMsg()
	t.Log(hex.EncodeToString(buf))
	str := make([]byte, extraVanity)
	end := make([]byte, extraSeal)
	buf = append(str, buf...)
	buf = append(buf, end...)
	t.Log(hex.EncodeToString(buf))
}

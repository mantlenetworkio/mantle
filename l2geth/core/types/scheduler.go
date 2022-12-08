package types

import (
	"encoding/binary"
	"errors"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
)

const (
	uint64Length = 8
)

type MsgVerify interface {
	GetData() []byte
	GetSignature() []byte
	GetSigner() (common.Address, error)
}

var _ MsgVerify = &BatchPeriodStartMsg{}

type BatchPeriodStartMsg struct {
	ReorgIndex   uint64
	BatchIndex   uint64
	StartHeight  uint64
	MaxHeight    uint64
	ExpireTime   uint64
	MinerAddress common.Address
	SequencerSet []common.Address
	Signature    []byte
}

func (bps *BatchPeriodStartMsg) GetSigner() (common.Address, error) {
	if bps.Signature == nil {
		return common.Address{}, errors.New("msg do not have signature")
	}
	pubEcr, err := crypto.SigToPub(crypto.Keccak256(bps.GetData()), bps.GetSignature())
	if err != nil {
		return common.Address{}, errors.New("signature ecrecover failed")
	}
	addressEcr := crypto.PubkeyToAddress(*pubEcr)
	return addressEcr, nil
}

func (bps *BatchPeriodStartMsg) GetData() []byte {
	if bps == nil || len(bps.SequencerSet) == 0 {
		return nil
	}
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, bps.ReorgIndex)
	buf = binary.BigEndian.AppendUint64(buf, bps.BatchIndex)
	buf = binary.BigEndian.AppendUint64(buf, bps.StartHeight)
	buf = binary.BigEndian.AppendUint64(buf, bps.MaxHeight)
	buf = binary.BigEndian.AppendUint64(buf, bps.ExpireTime)

	buf = append(buf, bps.MinerAddress.Bytes()...)

	for _, sequencer := range bps.SequencerSet {
		buf = append(buf, sequencer.Bytes()...)
	}
	return buf
}

func (bps *BatchPeriodStartMsg) GetSignature() []byte {
	return bps.Signature
}

func (bps *BatchPeriodStartMsg) SerializeBatchPeriodStartMsg() []byte {
	if bps == nil || len(bps.SequencerSet) == 0 || len(bps.Signature) != crypto.SignatureLength {
		return nil
	}
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, bps.ReorgIndex)
	buf = binary.BigEndian.AppendUint64(buf, bps.BatchIndex)
	buf = binary.BigEndian.AppendUint64(buf, bps.StartHeight)
	buf = binary.BigEndian.AppendUint64(buf, bps.MaxHeight)
	buf = binary.BigEndian.AppendUint64(buf, bps.ExpireTime)

	buf = append(buf, bps.MinerAddress.Bytes()...)

	for _, sequencer := range bps.SequencerSet {
		buf = append(buf, sequencer.Bytes()...)
	}

	buf = append(buf, bps.Signature...)

	return buf
}

func (bps *BatchPeriodStartMsg) Hash() common.Hash {
	if bps == nil {
		return common.Hash{}
	}
	return rlpHash(bps)
}

func IsValidBatchPeriodStartMsgBuf(buf []byte) bool {
	sequencerSetLength := len(buf) - (uint64Length*5 + common.AddressLength + crypto.SignatureLength)
	if sequencerSetLength <= 0 || sequencerSetLength%common.AddressLength != 0 {
		return false
	}
	return true
}

func DeserializeBatchPeriodStartMsg(buf []byte) BatchPeriodStartMsg {
	sequencerSetLength := len(buf) - (uint64Length*5 + common.AddressLength + crypto.SignatureLength)
	if sequencerSetLength <= 0 || sequencerSetLength%common.AddressLength != 0 {
		return BatchPeriodStartMsg{}
	}

	reorgIndex := binary.BigEndian.Uint64(buf[:uint64Length])
	batchIndex := binary.BigEndian.Uint64(buf[uint64Length : uint64Length*2])
	startHeight := binary.BigEndian.Uint64(buf[uint64Length*2 : uint64Length*3])
	maxHeight := binary.BigEndian.Uint64(buf[uint64Length*3 : uint64Length*4])
	expireTime := binary.BigEndian.Uint64(buf[uint64Length*4 : uint64Length*5])

	var sequencerSet []common.Address
	for idx := 0; idx < sequencerSetLength/common.AddressLength; idx++ {
		sequencerSet = append(sequencerSet, common.BytesToAddress(buf[uint64Length*5+common.AddressLength*(idx+1):uint64Length*5+common.AddressLength*(idx+2)]))
	}

	return BatchPeriodStartMsg{
		ReorgIndex:   reorgIndex,
		BatchIndex:   batchIndex,
		StartHeight:  startHeight,
		MaxHeight:    maxHeight,
		ExpireTime:   expireTime,
		MinerAddress: common.BytesToAddress(buf[uint64Length*5 : uint64Length*5+common.AddressLength]),
		SequencerSet: sequencerSet,
		Signature:    buf[len(buf)-crypto.SignatureLength:],
	}
}

type BatchPeriodEndMsg struct {
	ReorgIndex   uint64
	BatchIndex   uint64
	StartHeight  uint64
	EndHeight    uint64
	MinerAddress common.Address
	Signatures   [][]byte
	Signature    []byte
}

func (bps *BatchPeriodEndMsg) Hash() common.Hash {
	if bps == nil {
		return common.Hash{}
	}
	return rlpHash(bps)
}

type FraudProofReorgMsg struct {
	ReorgIndex    uint64
	ReorgToHeight uint64
	TssSignature  []byte
}

func (bps *FraudProofReorgMsg) Hash() common.Hash {
	if bps == nil {
		return common.Hash{}
	}
	return rlpHash(bps)
}

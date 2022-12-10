package types

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/rlp"
)

const (
	uint64Length = 8
)

type MsgVerify interface {
	Serialize() []byte
	GetSignData() []byte
	GetSignature() []byte
	GetSigner() (common.Address, error)
	Hash() common.Hash
}

var _ MsgVerify = &BatchPeriodStartMsg{}

type BatchPeriodStartMsg struct {
	ReorgIndex  uint64
	BatchIndex  uint64
	StartHeight uint64
	MaxHeight   uint64
	ExpireTime  uint64
	Sequencer   common.Address
	Signature   []byte
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

	return BatchPeriodStartMsg{
		ReorgIndex:  reorgIndex,
		BatchIndex:  batchIndex,
		StartHeight: startHeight,
		MaxHeight:   maxHeight,
		ExpireTime:  expireTime,
		Sequencer:   common.BytesToAddress(buf[uint64Length*5 : uint64Length*5+common.AddressLength]),
		Signature:   buf[len(buf)-crypto.SignatureLength:],
	}
}

func (bps *BatchPeriodStartMsg) Serialize() []byte {
	if bps == nil || len(bps.Signature) != crypto.SignatureLength {
		return nil
	}
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, bps.ReorgIndex)
	buf = binary.BigEndian.AppendUint64(buf, bps.BatchIndex)
	buf = binary.BigEndian.AppendUint64(buf, bps.StartHeight)
	buf = binary.BigEndian.AppendUint64(buf, bps.MaxHeight)
	buf = binary.BigEndian.AppendUint64(buf, bps.ExpireTime)

	buf = append(buf, bps.Sequencer.Bytes()...)

	buf = append(buf, bps.Signature...)

	return buf
}

func (bps *BatchPeriodStartMsg) GetSigner() (common.Address, error) {
	if bps.Signature == nil {
		return common.Address{}, errors.New("msg do not have signature")
	}
	pubEcr, err := crypto.SigToPub(crypto.Keccak256(bps.GetSignData()), bps.GetSignature())
	if err != nil {
		return common.Address{}, errors.New("signature ecrecover failed")
	}
	addressEcr := crypto.PubkeyToAddress(*pubEcr)
	return addressEcr, nil
}

func (bps *BatchPeriodStartMsg) GetSignData() []byte {
	if bps == nil {
		return nil
	}
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, bps.ReorgIndex)
	buf = binary.BigEndian.AppendUint64(buf, bps.BatchIndex)
	buf = binary.BigEndian.AppendUint64(buf, bps.StartHeight)
	buf = binary.BigEndian.AppendUint64(buf, bps.MaxHeight)
	buf = binary.BigEndian.AppendUint64(buf, bps.ExpireTime)

	buf = append(buf, bps.Sequencer.Bytes()...)
	return buf
}

func (bps *BatchPeriodStartMsg) GetSignature() []byte {
	return bps.Signature
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

type BatchPeriodAnswerMsg struct {
	Sequencer  common.Address
	StartIndex uint64
	Txs        Transactions
	Signature  []byte
}

func DeserializeBatchPeriodAnswerMsg(buf []byte) (BatchPeriodAnswerMsg, error) {
	txsBytes := len(buf) - (common.AddressLength + uint64Length + crypto.SignatureLength)
	if txsBytes <= 0 {
		return BatchPeriodAnswerMsg{}, fmt.Errorf("invalid BatchPeriodAnswerMsg bytes")
	}

	sequencer := common.BytesToAddress(buf[:common.AddressLength])
	startIndex := binary.BigEndian.Uint64(buf[common.AddressLength : common.AddressLength+uint64Length])

	var txs Transactions

	startPos := common.AddressLength + uint64Length
	for {
		if startPos >= len(buf)-crypto.SignatureLength {
			break
		}
		txBytesLength := int(binary.BigEndian.Uint64(buf[startPos : startPos+uint64Length]))
		if startPos+uint64Length+txBytesLength > len(buf)-crypto.SignatureLength {
			return BatchPeriodAnswerMsg{}, fmt.Errorf("invalid tx bytes")
		}
		txBytes := buf[startPos+uint64Length : startPos+uint64Length+txBytesLength]
		var tx *Transaction
		err := rlp.DecodeBytes(txBytes, &tx)
		if err != nil {
			return BatchPeriodAnswerMsg{}, err
		}
		txs = append(txs, tx)

		startPos = startPos + uint64Length + txBytesLength
	}

	signature := buf[len(buf)-crypto.SignatureLength:]

	return BatchPeriodAnswerMsg{
		Sequencer:  sequencer,
		StartIndex: startIndex,
		Txs:        txs,
		Signature:  signature,
	}, nil
}

func (bpa *BatchPeriodAnswerMsg) Serialize() []byte {
	if bpa == nil || len(bpa.Signature) != crypto.SignatureLength {
		return nil
	}
	buf := bpa.Sequencer.Bytes()
	binary.BigEndian.PutUint64(buf, bpa.StartIndex)
	for i, _ := range bpa.Txs {
		txBytes := bpa.Txs.GetRlp(i)

		var txBytesLengthBytes = make([]byte, 8)
		binary.BigEndian.PutUint64(txBytesLengthBytes, uint64(len(txBytes)))
		buf = append(buf, txBytesLengthBytes...)
		buf = append(buf, txBytes...)
	}

	buf = append(buf, bpa.Signature...)

	return buf
}

func (bpa *BatchPeriodAnswerMsg) GetSigner() (common.Address, error) {
	if bpa.Signature == nil {
		return common.Address{}, errors.New("msg do not have signature")
	}
	pubEcr, err := crypto.SigToPub(crypto.Keccak256(bpa.GetSignData()), bpa.GetSignature())
	if err != nil {
		return common.Address{}, errors.New("signature ecrecover failed")
	}
	addressEcr := crypto.PubkeyToAddress(*pubEcr)
	return addressEcr, nil
}

func (bpa *BatchPeriodAnswerMsg) GetSignData() []byte {
	if bpa == nil {
		return nil
	}
	buf := bpa.Sequencer.Bytes()
	binary.BigEndian.PutUint64(buf, bpa.StartIndex)
	for _, tx := range bpa.Txs {
		txHash := tx.Hash()
		buf = append(buf, txHash[:]...)
	}

	return buf
}

func (bpa *BatchPeriodAnswerMsg) GetSignature() []byte {
	return bpa.Signature
}

func (bpa *BatchPeriodAnswerMsg) Hash() common.Hash {
	if bpa == nil {
		return common.Hash{}
	}
	return rlpHash(bpa)
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

package types

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

const (
	uint64Length = 8
)

type MsgVerify interface {
	GetSignData() []byte
	GetSignature() []byte
	GetSigner() (common.Address, error)
	Hash() common.Hash
}

var _ MsgVerify = &BatchPeriodStartMsg{}
var _ MsgVerify = &BatchPeriodAnswerMsg{}
var _ MsgVerify = &BatchTxSetProof{}

func VerifySigner(msg MsgVerify, addr common.Address) bool {
	signer, err := msg.GetSigner()
	if err != nil {
		log.Error("Verify signature err", "err", err)
	}
	if !bytes.Equal(signer.Bytes(), addr.Bytes()) {
		log.Error("Verify signature failed", "addr", addr, "signer", signer.String())
		return false
	}
	return true
}

type BatchPeriodStartMsg struct {
	RollbackStates RollbackStates
	BatchIndex     uint64
	StartHeight    uint64
	MaxHeight      uint64
	ExpireTime     uint64
	Sequencer      common.Address
	Signature      []byte
}

func (bps *BatchPeriodStartMsg) Serialize() []byte {
	if bps == nil || len(bps.Signature) != crypto.SignatureLength {
		return nil
	}
	var buf = make([]byte, 8)
	for _, rollbackState := range bps.RollbackStates {
		binary.BigEndian.AppendUint64(buf, rollbackState.BlockNumber)
		buf = append(buf, rollbackState.BlockHash.Bytes()...)
	}
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
	for _, rollbackState := range bps.RollbackStates {
		binary.BigEndian.AppendUint64(buf, rollbackState.BlockNumber)
		buf = append(buf, rollbackState.BlockHash.Bytes()...)
	}
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

type BatchPeriodAnswerMsg struct {
	Sequencer  common.Address
	BatchIndex uint64
	StartIndex uint64
	Txs        Transactions
	Signature  []byte
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
	buf = binary.BigEndian.AppendUint64(buf, bpa.BatchIndex)
	buf = binary.BigEndian.AppendUint64(buf, bpa.StartIndex)
	for _, tx := range bpa.Txs {
		tempTxs := make(Transactions, 1, 1)
		tempTxs[0] = tx
		txHashRoot := DeriveSha(tempTxs)
		buf = append(buf, txHashRoot[:]...)
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

func (bpa *BatchPeriodAnswerMsg) ToBatchTxSetProof() *BatchTxSetProof {
	result := BatchTxSetProof{
		Sequencer:  bpa.Sequencer,
		BatchIndex: bpa.BatchIndex,
		StartIndex: bpa.StartIndex,
		Signature:  bpa.Signature,
	}
	for _, tx := range bpa.Txs {
		tempTxs := make(Transactions, 1, 1)
		tempTxs[0] = tx
		txHashRoot := DeriveSha(tempTxs)
		result.TxHashSet = append(result.TxHashSet, txHashRoot)
	}

	return &result
}

type BatchTxSetProof struct {
	Sequencer  common.Address
	BatchIndex uint64
	StartIndex uint64
	TxHashSet  []common.Hash
	Signature  []byte
}

func DecodeBatchTxSetProof(buf []byte) (BatchTxSetProof, error) {
	if len(buf) <= common.AddressLength+uint64Length+uint64Length+crypto.SignatureLength {
		return BatchTxSetProof{}, fmt.Errorf("BatchPeriodAnswerMsg SignData length is too short")
	}
	txHashSetLength := len(buf) - common.AddressLength - uint64Length - uint64Length - crypto.SignatureLength
	if txHashSetLength%common.HashLength != 0 {
		return BatchTxSetProof{}, fmt.Errorf("BatchPeriodAnswerMsg SignData contains invalid tx hash set")
	}

	sequencer := common.BytesToAddress(buf[:common.AddressLength])
	batchIndex := binary.BigEndian.Uint64(buf[common.AddressLength : common.AddressLength+uint64Length])
	startIndex := binary.BigEndian.Uint64(buf[common.AddressLength+uint64Length : common.AddressLength+uint64Length*2])

	var txs []common.Hash

	startPos := common.AddressLength + uint64Length*2
	for {
		if startPos >= len(buf)-crypto.SignatureLength {
			break
		}
		txHashBytes := buf[startPos : startPos+common.HashLength]
		var txHash common.Hash
		copy(txHash[:], txHashBytes)
		txs = append(txs, txHash)

		startPos = startPos + common.HashLength
	}

	signature := buf[len(buf)-crypto.SignatureLength:]

	return BatchTxSetProof{
		Sequencer:  sequencer,
		BatchIndex: batchIndex,
		StartIndex: startIndex,
		TxHashSet:  txs,
		Signature:  signature,
	}, nil
}

func (btsp *BatchTxSetProof) Serialize() []byte {
	if btsp == nil || len(btsp.TxHashSet) == 0 || len(btsp.Signature) != crypto.SignatureLength {
		return nil
	}
	buf := btsp.Sequencer[:]
	buf = binary.BigEndian.AppendUint64(buf, btsp.BatchIndex)
	buf = binary.BigEndian.AppendUint64(buf, btsp.StartIndex)
	for _, txHash := range btsp.TxHashSet {
		buf = append(buf, txHash[:]...)
	}

	buf = append(buf, btsp.Signature...)
	return buf
}

func (btsp *BatchTxSetProof) GetSignData() []byte {
	if btsp == nil {
		return nil
	}
	buf := btsp.Sequencer[:]
	buf = binary.BigEndian.AppendUint64(buf, btsp.BatchIndex)
	buf = binary.BigEndian.AppendUint64(buf, btsp.StartIndex)
	for _, txHash := range btsp.TxHashSet {
		buf = append(buf, txHash[:]...)
	}
	return buf
}

func (btsp *BatchTxSetProof) GetSignature() []byte {
	return btsp.Signature
}

func (btsp *BatchTxSetProof) GetSigner() (common.Address, error) {
	pubEcr, err := crypto.SigToPub(crypto.Keccak256(btsp.GetSignData()), btsp.GetSignature())
	if err != nil {
		return common.Address{}, errors.New("signature ecrecover failed")
	}

	return crypto.PubkeyToAddress(*pubEcr), nil
}

func (btsp *BatchTxSetProof) Hash() common.Hash {
	if btsp == nil {
		return common.Hash{}
	}
	return rlpHash(btsp)
}

func (btsp *BatchTxSetProof) ContainTxHashOrNot(txHash common.Hash, height uint64) bool {
	if btsp == nil {
		return false
	}
	if height < btsp.StartIndex || height-btsp.StartIndex >= uint64(len(btsp.TxHashSet)) {
		return false
	}
	if bytes.Equal(txHash[:], btsp.TxHashSet[height-btsp.StartIndex][:]) {
		return true
	}
	return false
}

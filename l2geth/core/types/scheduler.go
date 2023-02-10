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
	Sequencer   common.Address
	BatchIndex  uint64
	StartHeight uint64
	RootHash    common.Hash
	Txs         Transactions
	Signature   []byte
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
	buf = binary.BigEndian.AppendUint64(buf, bpa.StartHeight)
	buf = append(buf, bpa.RootHash[:]...)

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

func (bpa *BatchPeriodAnswerMsg) ToBatchTxSetProof(tx *Transaction) (*BatchTxSetProof, error) {
	var txHashSet []common.Hash
	for _, tx := range bpa.Txs {
		tempTxs := make(Transactions, 1, 1)
		tempTxs[0] = tx
		txHashRoot := DeriveSha(tempTxs)
		txHashSet = append(txHashSet, txHashRoot)
	}
	proof, rootHash, err := generateMerkleProof(txHashSet, DeriveSha(Transactions{tx}))
	if err != nil {
		return nil, err
	}
	result := BatchTxSetProof{
		Sequencer:   bpa.Sequencer,
		BatchIndex:  bpa.BatchIndex,
		StartHeight: bpa.StartHeight,
		RootHash:    rootHash,
		Proof:       proof,
		Signature:   bpa.Signature,
	}
	return &result, nil
}

type BatchTxSetProof struct {
	Sequencer   common.Address
	BatchIndex  uint64
	StartHeight uint64
	RootHash    common.Hash
	Proof       []common.Hash
	Signature   []byte
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
	startHeight := binary.BigEndian.Uint64(buf[common.AddressLength+uint64Length : common.AddressLength+uint64Length*2])
	rootHash := common.BytesToHash(buf[common.AddressLength+uint64Length*2 : common.AddressLength+uint64Length*2+common.HashLength])

	var proof []common.Hash

	startPos := common.AddressLength + uint64Length*2
	for {
		if startPos >= len(buf)-crypto.SignatureLength {
			break
		}
		txHashBytes := buf[startPos : startPos+common.HashLength]
		var proofItem common.Hash
		copy(proofItem[:], txHashBytes)
		proof = append(proof, proofItem)

		startPos = startPos + common.HashLength
	}

	signature := buf[len(buf)-crypto.SignatureLength:]

	return BatchTxSetProof{
		Sequencer:   sequencer,
		BatchIndex:  batchIndex,
		StartHeight: startHeight,
		RootHash:    rootHash,
		Proof:       proof,
		Signature:   signature,
	}, nil
}

func (btsp *BatchTxSetProof) Serialize() []byte {
	if btsp == nil || len(btsp.Proof) == 0 || len(btsp.Signature) != crypto.SignatureLength {
		return nil
	}
	buf := btsp.Sequencer[:]
	buf = append(buf, btsp.RootHash[:]...)
	for _, txHash := range btsp.Proof {
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
	buf = binary.BigEndian.AppendUint64(buf, btsp.StartHeight)
	buf = append(buf, btsp.RootHash[:]...)
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

func (btsp *BatchTxSetProof) ValidProof(txHash common.Hash) bool {
	if btsp == nil {
		return false
	}
	return validateMerkleProof(btsp.Proof, txHash, btsp.RootHash)
}

func calculateProofLength(totalLeaf int) int {
	merkleProofLength := 0
	divResult := totalLeaf
	hasMod := false
	for divResult != 0 {
		tempDivResult := divResult / 2
		if tempDivResult*2 < divResult {
			hasMod = true
		}
		divResult = tempDivResult
		merkleProofLength++
	}
	if !hasMod {
		merkleProofLength--
	}
	return merkleProofLength
}

func calculateProof(leafA, leafB common.Hash) common.Hash {
	if leafA.Big().Cmp(leafB.Big()) <= 0 {
		return crypto.Keccak256Hash(leafA[:], leafB[:])
	} else {
		return crypto.Keccak256Hash(leafB[:], leafA[:])
	}
}

func generateMerkleProof(hashSet []common.Hash, txHash common.Hash) ([]common.Hash, common.Hash, error) {
	merkleProofLength := calculateProofLength(len(hashSet))
	merkleProof := make([]common.Hash, merkleProofLength, merkleProofLength)

	merkleProofInputs := make([]common.Hash, 0, len(hashSet))
	var leaf common.Hash
	found := false
	for idx, hash := range hashSet {
		idxBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(idxBytes, uint64(idx))
		merkleProofInputs[idx] = crypto.Keccak256Hash(idxBytes, hash[:])
		if bytes.Equal(txHash[:], hash[:]) {
			leaf = merkleProofInputs[idx]
			found = true
		}
	}
	if !found {
		return nil, common.Hash{}, fmt.Errorf("hashSet doesn't contain txHash")
	}

	tempProofLength := (len(hashSet) + 1) / 2
	tempProof := make([]common.Hash, tempProofLength, tempProofLength)
	var merkleRoot common.Hash

	proofIdx := 0
	for len(tempProof) >= 1 {
		for idx := 0; idx < len(tempProof); idx++ {
			matched := false
			if bytes.Equal(leaf[:], merkleProofInputs[2*idx][:]) {
				merkleProof[proofIdx] = merkleProofInputs[2*idx+1]
				proofIdx++
				matched = true
			} else if bytes.Equal(leaf[:], merkleProofInputs[2*idx+1][:]) {
				merkleProof[proofIdx] = merkleProofInputs[2*idx]
				proofIdx++
				matched = true
			}
			tempProof[idx] = calculateProof(merkleProofInputs[2*idx], merkleProofInputs[2*idx+1])
			if matched {
				leaf = tempProof[idx]
			}
		}
		merkleProofInputs = tempProof

		if len(tempProof) == 1 {
			merkleRoot = tempProof[0]
			break
		}
		tempProofLength = (tempProofLength + 1) / 2
		tempProof = make([]common.Hash, tempProofLength, tempProofLength)
	}
	return merkleProof, merkleRoot, nil
}

func validateMerkleProof(merkleProof []common.Hash, txHash common.Hash, rootHash common.Hash) bool {
	hash := txHash
	for _, proof := range merkleProof {
		hash = calculateProof(hash, proof)
	}
	return bytes.Equal(hash[:], rootHash[:])
}

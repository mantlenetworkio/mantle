// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"bytes"
	"math"
	"math/big"
	"reflect"
	"testing"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/params"
	"github.com/mantlenetworkio/mantle/l2geth/rlp"
)

func TestLegacyReceiptDecoding(t *testing.T) {
	tests := []struct {
		name   string
		encode func(*Receipt) ([]byte, error)
	}{
		{
			"StoredReceiptRLP",
			encodeAsStoredReceiptRLP,
		},
		{
			"V4StoredReceiptRLP",
			encodeAsV4StoredReceiptRLP,
		},
		{
			"V3StoredReceiptRLP",
			encodeAsV3StoredReceiptRLP,
		},
	}

	tx := NewTransaction(1, common.HexToAddress("0x1"), big.NewInt(1), 1, big.NewInt(1), nil)
	receipt := &Receipt{
		Status:            ReceiptStatusFailed,
		CumulativeGasUsed: 1,
		Logs: []*Log{
			{
				Address: common.BytesToAddress([]byte{0x11}),
				Topics:  []common.Hash{common.HexToHash("dead"), common.HexToHash("beef")},
				Data:    []byte{0x01, 0x00, 0xff},
			},
			{
				Address: common.BytesToAddress([]byte{0x01, 0x11}),
				Topics:  []common.Hash{common.HexToHash("dead"), common.HexToHash("beef")},
				Data:    []byte{0x01, 0x00, 0xff},
			},
		},
		TxHash:          tx.Hash(),
		ContractAddress: common.BytesToAddress([]byte{0x01, 0x11, 0x11}),
		GasUsed:         111111,
	}
	receipt.Bloom = CreateBloom(Receipts{receipt})

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			enc, err := tc.encode(receipt)
			if err != nil {
				t.Fatalf("Error encoding receipt: %v", err)
			}
			var dec ReceiptForStorage
			if err := rlp.DecodeBytes(enc, &dec); err != nil {
				t.Fatalf("Error decoding RLP receipt: %v", err)
			}
			// Check whether all consensus fields are correct.
			if dec.Status != receipt.Status {
				t.Fatalf("Receipt status mismatch, want %v, have %v", receipt.Status, dec.Status)
			}
			if dec.CumulativeGasUsed != receipt.CumulativeGasUsed {
				t.Fatalf("Receipt CumulativeGasUsed mismatch, want %v, have %v", receipt.CumulativeGasUsed, dec.CumulativeGasUsed)
			}
			if dec.Bloom != receipt.Bloom {
				t.Fatalf("Bloom data mismatch, want %v, have %v", receipt.Bloom, dec.Bloom)
			}
			if len(dec.Logs) != len(receipt.Logs) {
				t.Fatalf("Receipt log number mismatch, want %v, have %v", len(receipt.Logs), len(dec.Logs))
			}
			for i := 0; i < len(dec.Logs); i++ {
				if dec.Logs[i].Address != receipt.Logs[i].Address {
					t.Fatalf("Receipt log %d address mismatch, want %v, have %v", i, receipt.Logs[i].Address, dec.Logs[i].Address)
				}
				if !reflect.DeepEqual(dec.Logs[i].Topics, receipt.Logs[i].Topics) {
					t.Fatalf("Receipt log %d topics mismatch, want %v, have %v", i, receipt.Logs[i].Topics, dec.Logs[i].Topics)
				}
				if !bytes.Equal(dec.Logs[i].Data, receipt.Logs[i].Data) {
					t.Fatalf("Receipt log %d data mismatch, want %v, have %v", i, receipt.Logs[i].Data, dec.Logs[i].Data)
				}
			}
		})
	}
}

func encodeAsStoredReceiptRLP(want *Receipt) ([]byte, error) {
	stored := &storedReceiptRLP{
		PostStateOrStatus: want.statusEncoding(),
		CumulativeGasUsed: want.CumulativeGasUsed,
		Logs:              make([]*LogForStorage, len(want.Logs)),
	}
	for i, log := range want.Logs {
		stored.Logs[i] = (*LogForStorage)(log)
	}
	return rlp.EncodeToBytes(stored)
}

func encodeAsV4StoredReceiptRLP(want *Receipt) ([]byte, error) {
	stored := &v4StoredReceiptRLP{
		PostStateOrStatus: want.statusEncoding(),
		CumulativeGasUsed: want.CumulativeGasUsed,
		TxHash:            want.TxHash,
		ContractAddress:   want.ContractAddress,
		Logs:              make([]*LogForStorage, len(want.Logs)),
		GasUsed:           want.GasUsed,
	}
	for i, log := range want.Logs {
		stored.Logs[i] = (*LogForStorage)(log)
	}
	return rlp.EncodeToBytes(stored)
}

func encodeAsV3StoredReceiptRLP(want *Receipt) ([]byte, error) {
	stored := &v3StoredReceiptRLP{
		PostStateOrStatus: want.statusEncoding(),
		CumulativeGasUsed: want.CumulativeGasUsed,
		Bloom:             want.Bloom,
		TxHash:            want.TxHash,
		ContractAddress:   want.ContractAddress,
		Logs:              make([]*LogForStorage, len(want.Logs)),
		GasUsed:           want.GasUsed,
	}
	for i, log := range want.Logs {
		stored.Logs[i] = (*LogForStorage)(log)
	}
	return rlp.EncodeToBytes(stored)
}

// Tests that receipt data can be correctly derived from the contextual infos
func TestDeriveFields(t *testing.T) {
	// Create a few transactions to have receipts for
	txs := Transactions{
		NewContractCreation(1, big.NewInt(1), 1, big.NewInt(1), nil),
		NewTransaction(2, common.HexToAddress("0x2"), big.NewInt(2), 2, big.NewInt(2), nil),
	}
	// Create the corresponding receipts
	receipts := Receipts{
		&Receipt{
			Status:            ReceiptStatusFailed,
			CumulativeGasUsed: 1,
			Logs: []*Log{
				{Address: common.BytesToAddress([]byte{0x11})},
				{Address: common.BytesToAddress([]byte{0x01, 0x11})},
			},
			TxHash:          txs[0].Hash(),
			ContractAddress: common.BytesToAddress([]byte{0x01, 0x11, 0x11}),
			GasUsed:         1,
		},
		&Receipt{
			PostState:         common.Hash{2}.Bytes(),
			CumulativeGasUsed: 3,
			Logs: []*Log{
				{Address: common.BytesToAddress([]byte{0x22})},
				{Address: common.BytesToAddress([]byte{0x02, 0x22})},
			},
			TxHash:          txs[1].Hash(),
			ContractAddress: common.BytesToAddress([]byte{0x02, 0x22, 0x22}),
			GasUsed:         2,
		},
	}
	// Clear all the computed fields and re-derive them
	number := big.NewInt(1)
	hash := common.BytesToHash([]byte{0x03, 0x14})

	clearComputedFieldsOnReceipts(t, receipts)
	if err := receipts.DeriveFields(params.TestChainConfig, hash, number.Uint64(), txs); err != nil {
		t.Fatalf("DeriveFields(...) = %v, want <nil>", err)
	}
	// Iterate over all the computed fields and check that they're correct
	signer := MakeSigner(params.TestChainConfig, number)

	logIndex := uint(0)
	for i := range receipts {
		if receipts[i].TxHash != txs[i].Hash() {
			t.Errorf("receipts[%d].TxHash = %s, want %s", i, receipts[i].TxHash.String(), txs[i].Hash().String())
		}
		if receipts[i].BlockHash != hash {
			t.Errorf("receipts[%d].BlockHash = %s, want %s", i, receipts[i].BlockHash.String(), hash.String())
		}
		if receipts[i].BlockNumber.Cmp(number) != 0 {
			t.Errorf("receipts[%c].BlockNumber = %s, want %s", i, receipts[i].BlockNumber.String(), number.String())
		}
		if receipts[i].TransactionIndex != uint(i) {
			t.Errorf("receipts[%d].TransactionIndex = %d, want %d", i, receipts[i].TransactionIndex, i)
		}
		if receipts[i].GasUsed != txs[i].Gas() {
			t.Errorf("receipts[%d].GasUsed = %d, want %d", i, receipts[i].GasUsed, txs[i].Gas())
		}
		if txs[i].To() != nil && receipts[i].ContractAddress != (common.Address{}) {
			t.Errorf("receipts[%d].ContractAddress = %s, want %s", i, receipts[i].ContractAddress.String(), (common.Address{}).String())
		}
		from, _ := Sender(signer, txs[i])
		contractAddress := crypto.CreateAddress(from, txs[i].Nonce())
		if txs[i].To() == nil && receipts[i].ContractAddress != contractAddress {
			t.Errorf("receipts[%d].ContractAddress = %s, want %s", i, receipts[i].ContractAddress.String(), contractAddress.String())
		}
		for j := range receipts[i].Logs {
			if receipts[i].Logs[j].BlockNumber != number.Uint64() {
				t.Errorf("receipts[%d].Logs[%d].BlockNumber = %d, want %d", i, j, receipts[i].Logs[j].BlockNumber, number.Uint64())
			}
			if receipts[i].Logs[j].BlockHash != hash {
				t.Errorf("receipts[%d].Logs[%d].BlockHash = %s, want %s", i, j, receipts[i].Logs[j].BlockHash.String(), hash.String())
			}
			if receipts[i].Logs[j].TxHash != txs[i].Hash() {
				t.Errorf("receipts[%d].Logs[%d].TxHash = %s, want %s", i, j, receipts[i].Logs[j].TxHash.String(), txs[i].Hash().String())
			}
			if receipts[i].Logs[j].TxHash != txs[i].Hash() {
				t.Errorf("receipts[%d].Logs[%d].TxHash = %s, want %s", i, j, receipts[i].Logs[j].TxHash.String(), txs[i].Hash().String())
			}
			if receipts[i].Logs[j].TxIndex != uint(i) {
				t.Errorf("receipts[%d].Logs[%d].TransactionIndex = %d, want %d", i, j, receipts[i].Logs[j].TxIndex, i)
			}
			if receipts[i].Logs[j].Index != logIndex {
				t.Errorf("receipts[%d].Logs[%d].Index = %d, want %d", i, j, receipts[i].Logs[j].Index, logIndex)
			}
			logIndex++
		}
	}
}

func clearComputedFieldsOnReceipts(t *testing.T, receipts Receipts) {
	t.Helper()

	for _, receipt := range receipts {
		clearComputedFieldsOnReceipt(t, receipt)
	}
}

func clearComputedFieldsOnReceipt(t *testing.T, receipt *Receipt) {
	t.Helper()

	receipt.TxHash = common.Hash{}
	receipt.BlockHash = common.Hash{}
	receipt.BlockNumber = big.NewInt(math.MaxUint32)
	receipt.TransactionIndex = math.MaxUint32
	receipt.ContractAddress = common.Address{}
	receipt.GasUsed = 0

	clearComputedFieldsOnLogs(t, receipt.Logs)
}

func clearComputedFieldsOnLogs(t *testing.T, logs []*Log) {
	t.Helper()

	for _, log := range logs {
		clearComputedFieldsOnLog(t, log)
	}
}

func clearComputedFieldsOnLog(t *testing.T, log *Log) {
	t.Helper()

	log.BlockNumber = math.MaxUint32
	log.BlockHash = common.Hash{}
	log.TxHash = common.Hash{}
	log.TxIndex = math.MaxUint32
	log.Index = math.MaxUint32
}

func TestDecodeReceiptsBytes(t *testing.T) {

	tx := NewTransaction(1, common.HexToAddress("0x4932"), big.NewInt(math.MaxUint32), math.MaxUint64, big.NewInt(math.MaxUint32), []byte{0x11})
	receipt := &Receipt{
		PostState:         common.Hash{1}.Bytes(),
		Status:            ReceiptStatusSuccessful,
		CumulativeGasUsed: 1800,
		Logs: []*Log{
			{Address: common.BytesToAddress([]byte("0x112"))},
			{Address: common.BytesToAddress([]byte{0x01, 0x11})},
		},
		TxHash:           tx.Hash(),
		ContractAddress:  common.BytesToAddress([]byte{0x02, 0x22, 0x22}),
		GasUsed:          10000,
		BlockHash:        common.BytesToHash([]byte{0x03, 0x33, 0x33}),
		BlockNumber:      big.NewInt(math.MaxUint32),
		TransactionIndex: math.MaxUint32,
		L1GasPrice:       big.NewInt(int64(100)),
		L1GasUsed:        big.NewInt(int64(100)),
		L1Fee:            big.NewInt(int64(100)),
		FeeScalar:        big.NewFloat(10.12),
		DAGasPrice:       big.NewInt(int64(10)),
		DAGasUsed:        big.NewInt(int64(10)),
		DAFee:            big.NewInt(int64(10)),
	}
	receipt.Bloom = CreateBloom(Receipts{receipt})

	receiptByte, err := encodeAsStoredReceiptRLPV2(receipt)
	if err != nil {
		t.Fatalf("Error encoding receipt: %v", err)
		return
	}
	var s ReceiptForStorage
	err = rlp.DecodeReceiptsBytes(receiptByte, &s)
	if err != nil {
		t.Fatalf("Error decoding RLP receipt: %v", err)
	}

	t.Log(receipt)
	t.Log(s)

	if !bytes.Equal(receipt.PostState, s.PostState) {
		t.Errorf("receipt.PostState = %v, want %v", receipt.PostState, s.PostState)
	}

	for i := 0; i < len(s.Logs); i++ {
		if s.Logs[i].Address != receipt.Logs[i].Address {
			t.Fatalf("receipt log %d address mismatch, want %v, have %v", i, receipt.Logs[i].Address, s.Logs[i].Address)
		}
	}
	if receipt.CumulativeGasUsed != s.CumulativeGasUsed {
		t.Errorf("receipt.CumulativeGasUsed = %v, want %v", receipt.CumulativeGasUsed, s.CumulativeGasUsed)
	}
	if receipt.Bloom != s.Bloom {
		t.Errorf("receipt.Bloom = %v, want %v", receipt.Bloom, &s.Bloom)
	}
	if receipt.L1GasPrice.String() != s.L1GasPrice.String() {
		t.Errorf("receipt.L1GasPrice = %v, want %v", receipt.L1GasPrice, s.L1GasPrice)
	}
	if receipt.L1Fee.String() != s.L1Fee.String() {
		t.Errorf("receipt.L1Fee = %v, want %v", receipt.L1Fee, s.L1Fee)
	}
	if receipt.L1GasUsed.String() != s.L1GasUsed.String() {
		t.Errorf("receipt.L1GasUsed = %v, want %v", receipt.L1GasUsed, s.L1GasUsed)
	}
	if receipt.FeeScalar.String() != s.FeeScalar.String() {
		t.Errorf("receipt.FeeScalar = %v, want %v", receipt.FeeScalar, s.FeeScalar)
	}
	if receipt.DAFee.String() != s.DAFee.String() {
		t.Errorf("receipt.DAFee = %v, want %v", receipt.DAFee, s.DAFee)
	}
	if receipt.DAGasUsed.String() != s.DAGasUsed.String() {
		t.Errorf("receipt.DAGasUsed = %v, want %v", receipt.DAGasUsed, s.DAGasUsed)
	}
	if receipt.DAGasPrice.String() != s.DAGasPrice.String() {
		t.Errorf("receipt.L1GasPrice = %v, want %v", receipt.DAGasPrice, s.DAGasPrice)
	}

}

func encodeAsStoredReceiptRLPV2(want *Receipt) ([]byte, error) {
	stored := storedReceiptRLP{
		PostStateOrStatus: want.statusEncoding(),
		CumulativeGasUsed: want.CumulativeGasUsed,
		Logs:              make([]*LogForStorage, len(want.Logs)),
		L1GasUsed:         want.L1GasUsed,
		L1GasPrice:        want.L1GasPrice,
		L1Fee:             want.L1Fee,
		FeeScalar:         want.FeeScalar.String(),
		DAGasUsed:         want.DAGasUsed,
		DAGasPrice:        want.DAGasPrice,
		DAFee:             want.DAFee,
	}
	for i, log := range want.Logs {
		stored.Logs[i] = (*LogForStorage)(log)
	}
	return rlp.EncodeToBytes(stored)
}

// testEncodeRLP
var receiptTests = []Receipt{
	//error rlp: cannot encode negative *big.Int
	{BlockNumber: big.NewInt(-1)},
	{DAFee: big.NewInt(-1)},
	{DAGasPrice: big.NewInt(-1)},
	{DAGasUsed: big.NewInt(-1)},
	{L1Fee: big.NewInt(-1)},
	{L1GasPrice: big.NewInt(-1)},
	{L1GasUsed: big.NewInt(-1)},

	//info Missing longitude,Ten in total
	{FeeScalar: big.NewFloat(10.123456789)},
}

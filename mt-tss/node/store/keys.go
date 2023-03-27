package store

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
)

var (
	CPKDataKeyPrefix                 = []byte{0x01}
	OutputKeyPrefix                  = []byte{0x02}
	IndexOutputKeyPrefix             = []byte{0x03}
	SigningInfoKeyPrefix             = []byte{0x04}
	NodeMissedBatchBitArrayKeyPrefix = []byte{0x05}
	SlashingInfoKeyPrefix            = []byte{0x06}
	ScannedHeightKeyPrefix           = []byte{0x07}
	CulpritsKeyPrefix                = []byte{0x08}
)

func getCPKDataKey(electionId uint64) []byte {
	electionIdBz := make([]byte, 8)
	binary.BigEndian.PutUint64(electionIdBz, electionId)
	return append(CPKDataKeyPrefix, electionIdBz...)
}

func getOutputKey(batchRoot [32]byte) []byte {
	return append(OutputKeyPrefix, batchRoot[:]...)
}

func getIndexOutputKey(index uint64) []byte {
	indexBz := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBz, index)
	return append(IndexOutputKeyPrefix, indexBz...)
}

func getSigningInfoKey(address common.Address) []byte {
	return append(SigningInfoKeyPrefix, address.Bytes()...)
}

// key: prefix + address + index
func getNodeMissedBatchBitArrayKey(address common.Address, index uint64) []byte {
	indexBz := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBz, index)
	return append(getNodeMissedBatchBitArrayAddressPrefixKey(address), indexBz...)
}

func getNodeMissedBatchBitArrayAddressPrefixKey(address common.Address) []byte {
	return append(NodeMissedBatchBitArrayKeyPrefix, address.Bytes()...)
}

// key: prefix + address + batchIndex
func getSlashingInfoKey(address common.Address, batchIndex uint64) []byte {
	indexBz := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBz, batchIndex)
	return append(getSlashingInfoAddressKey(address), indexBz...)
}

func getSlashingInfoAddressKey(address common.Address) []byte {
	return append(SlashingInfoKeyPrefix, address.Bytes()...)
}

func getScannedHeightKey() []byte {
	return ScannedHeightKeyPrefix
}

func getCulpritsKey() []byte {
	return CulpritsKeyPrefix
}

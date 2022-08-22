package store

import "encoding/binary"

var (
	CPKDataKeyPrefix = []byte{0x01}
)

func getCPKDataKey(electionId uint64) []byte {
	electionIdBz := make([]byte, 8)
	binary.BigEndian.PutUint64(electionIdBz, electionId)
	return append(CPKDataKeyPrefix, electionIdBz...)
}

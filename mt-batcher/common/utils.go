package common

import (
	"fmt"
	"github.com/Layr-Labs/datalayr/common/contracts"
	"github.com/Layr-Labs/datalayr/common/graphView"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"os"
)

func CreateUploadHeader(params StoreParams) ([]byte, error) {
	var kzgCommitArray [64]byte
	copy(kzgCommitArray[:], params.KzgCommit)
	var lowDegreeProof [64]byte
	copy(lowDegreeProof[:], params.LowDegreeProof)
	var disperserArray [20]byte
	copy(disperserArray[:], params.Disperser)

	h := contracts.DataStoreHeader{
		KzgCommit:      kzgCommitArray,
		LowDegreeProof: lowDegreeProof,
		Degree:         params.Degree,
		NumSys:         params.NumSys,
		NumPar:         params.NumPar,
		OrigDataSize:   params.OrigDataSize,
		Disperser:      disperserArray,
	}
	uploadHeader, _, err := contracts.CreateUploadHeader(h)
	if err != nil {
		return nil, err
	}
	return uploadHeader, nil
}

func GetMessageHash(event graphView.DataStoreInit) []byte {
	msg := make([]byte, 0)
	msg = append(msg, uint32ToByteSlice(event.StoreNumber)...)
	msg = append(msg, event.DataCommitment[:]...)
	msg = append(msg, byte(event.Duration))
	msg = append(msg, packTo(uint32ToByteSlice(event.InitTime), 32)...)
	msg = append(msg, uint32ToByteSlice(event.Index)...)
	msgHash := crypto.Keccak256(msg)
	return msgHash
}

func uint32ToByteSlice(x uint32) []byte {
	res := make([]byte, 4)
	res[0] = byte(x >> 24)
	res[1] = byte((x >> 16) & 255)
	res[2] = byte((x >> 8) & 255)
	res[3] = byte(x & 255)
	return res
}

func packTo(x []byte, n int) []byte {
	for i := len(x); i < n; i++ {
		x = append([]byte{byte(0)}, x...)
	}
	return x
}

func MakeCalldata(
	params StoreParams,
	meta DisperseMeta,
	storeNumber uint32,
	msgHash [32]byte,
) []byte {

	totalStakeIndexBytes := bigIntToBytes(
		new(big.Int).SetUint64(meta.TotalStakeIndex),
		6,
	)

	storeNumberBytes := bigIntToBytes(
		new(big.Int).SetUint64(uint64(storeNumber)),
		4,
	)

	stakesFromBlockNumberBytes := bigIntToBytes(
		new(big.Int).SetUint64(uint64(params.BlockNumber)),
		4,
	)

	numNonPubKeysBytes := bigIntToBytes(
		new(big.Int).SetUint64(uint64(len(meta.Sigs.NonSignerPubKeys))),
		4,
	)

	flattenedNonPubKeysBytes := make([]byte, 0)
	for i := 0; i < len(meta.Sigs.NonSignerPubKeys); i++ {
		flattenedNonPubKeysBytes = append(
			flattenedNonPubKeysBytes,
			meta.Sigs.NonSignerPubKeys[i]...,
		)
	}

	apkIndexBytes := bigIntToBytes(
		new(big.Int).SetUint64(uint64(meta.ApkIndex)),
		4,
	)

	var calldata []byte
	calldata = append(calldata, msgHash[:]...)
	calldata = append(calldata, totalStakeIndexBytes...)
	calldata = append(calldata, stakesFromBlockNumberBytes...)
	calldata = append(calldata, storeNumberBytes...)
	calldata = append(calldata, numNonPubKeysBytes...)
	calldata = append(calldata, flattenedNonPubKeysBytes...)
	calldata = append(calldata, apkIndexBytes...)
	calldata = append(calldata, meta.Sigs.AggPubKey...)
	calldata = append(calldata, meta.Sigs.AggSig...)
	return calldata

}

func bigIntToBytes(n *big.Int, packTo int) []byte {
	bigIntBytes := n.Bytes()
	bigIntLen := len(bigIntBytes)
	intBytes := make([]byte, packTo)

	if bigIntLen > packTo {
		fmt.Println("Cannot pad bytes: Desired length is less than existing length")
		os.Exit(1)
	}

	for i := 0; i < bigIntLen; i++ {
		intBytes[packTo-1-i] = bigIntBytes[bigIntLen-1-i]
	}
	return intBytes
}

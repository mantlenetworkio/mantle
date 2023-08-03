package common

import (
	"fmt"
	"math/big"
	"os"

	"github.com/Layr-Labs/datalayr/common/header"
)

func CreateUploadHeader(params StoreParams) ([]byte, error) {
	var kzgCommitArray [64]byte
	copy(kzgCommitArray[:], params.KzgCommit)
	var lowDegreeProof [64]byte
	copy(lowDegreeProof[:], params.LowDegreeProof)
	var disperserArray [20]byte
	copy(disperserArray[:], params.Disperser)

	h := header.DataStoreHeader{
		KzgCommit:      kzgCommitArray,
		LowDegreeProof: lowDegreeProof,
		Degree:         params.Degree,
		NumSys:         params.NumSys,
		NumPar:         params.NumPar,
		OrigDataSize:   params.OrigDataSize,
		Disperser:      disperserArray,
	}
	uploadHeader, _, err := header.CreateUploadHeader(h)
	if err != nil {
		return nil, err
	}
	return uploadHeader, nil
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

	referenceBlockNumberBytes := bigIntToBytes(
		new(big.Int).SetUint64(uint64(params.ReferenceBlockNumber)),
		4,
	)

	numNonPubKeysBytes := bigIntToBytes(
		new(big.Int).SetUint64(uint64(len(meta.Sigs.NonSignerPubkeys))),
		4,
	)

	flattenedNonPubKeysBytes := make([]byte, 0)
	for i := 0; i < len(meta.Sigs.NonSignerPubkeys); i++ {
		flattenedNonPubKeysBytes = append(
			flattenedNonPubKeysBytes,
			meta.Sigs.NonSignerPubkeys[i]...,
		)
	}

	apkIndexBytes := bigIntToBytes(
		new(big.Int).SetUint64(uint64(meta.ApkIndex)),
		4,
	)

	var calldata []byte
	calldata = append(calldata, msgHash[:]...)
	calldata = append(calldata, totalStakeIndexBytes...)
	calldata = append(calldata, referenceBlockNumberBytes...)
	calldata = append(calldata, storeNumberBytes...)
	calldata = append(calldata, numNonPubKeysBytes...)
	calldata = append(calldata, flattenedNonPubKeysBytes...)
	calldata = append(calldata, apkIndexBytes...)
	calldata = append(calldata, meta.Sigs.StoredAggPubkeyG1...)
	calldata = append(calldata, meta.Sigs.UsedAggPubkeyG2...)
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

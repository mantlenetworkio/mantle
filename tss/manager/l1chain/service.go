package l1chain

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/influxdata/influxdb/pkg/slices"
	"github.com/mantlenetworkio/mantle/tss/slash"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/tss/bindings/tgm"
	"github.com/mantlenetworkio/mantle/tss/manager/types"
)

type QueryService struct {
	ethClient             *ethclient.Client
	tssGroupManagerCaller *tgm.TssGroupManagerCaller
	confirmBlocks         uint64
	slashingStore         slash.SlashingStore
}

func NewQueryService(url, tssGroupContractAddress string, confirmBlocks int, store slash.SlashingStore) QueryService {
	cli, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	tssGroupManagerCaller, err := tgm.NewTssGroupManagerCaller(common.HexToAddress(tssGroupContractAddress), cli)
	if err != nil {
		panic(err)
	}
	return QueryService{
		ethClient:             cli,
		tssGroupManagerCaller: tssGroupManagerCaller,
		confirmBlocks:         uint64(confirmBlocks),
		slashingStore:         store,
	}
}

func (q QueryService) QueryActiveInfo() (types.TssCommitteeInfo, error) {
	currentBlockNumber, err := q.ethClient.BlockNumber(context.Background())
	if err != nil {
		return types.TssCommitteeInfo{}, err
	}
	electionId, threshold, cpk, activeTssMembers, err := q.tssGroupManagerCaller.GetTssGroupInfo(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(currentBlockNumber - q.confirmBlocks)})
	if err != nil {
		return types.TssCommitteeInfo{}, err
	}
	if len(cpk) == 0 {
		return types.TssCommitteeInfo{}, errors.New("cpk is not confirmed")
	}
	unmarshalledCPK, err := crypto.UnmarshalPubkey(append([]byte{0x04}, cpk...))
	if err != nil {
		log.Error("fail to unmarshal cpk", "err", err)
		return types.TssCommitteeInfo{}, nil
	}
	compressCPK := crypto.CompressPubkey(unmarshalledCPK)

	unjailMembers, err := q.tssGroupManagerCaller.GetTssGroupUnJailMembers(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(currentBlockNumber - q.confirmBlocks)})
	if err != nil {
		log.Error("fail to GetTssGroupUnJailMembers", "err", err)
		return types.TssCommitteeInfo{}, nil
	}

	var hasJailMembers bool
	if len(unjailMembers) < len(activeTssMembers) {
		log.Info("found jailed members from L1", "jailed number", len(activeTssMembers)-len(unjailMembers))
		hasJailMembers = true
	}
	// need to exclude the culprits
	culprits := q.slashingStore.GetCulprits()
	tssMembers := make([]string, 0)

	for _, m := range activeTssMembers {
		unmarshalled, err := crypto.UnmarshalPubkey(append([]byte{0x04}, m...))
		if err != nil {
			log.Error("fail to unmarshal tss member", "err", err)
			return types.TssCommitteeInfo{}, nil
		}
		compressed := crypto.CompressPubkey(unmarshalled)
		hexEncoded := hex.EncodeToString(compressed)
		if slices.Exists(culprits, hexEncoded) { // exclude culprits
			continue
		}

		if hasJailMembers {
			addr := crypto.PubkeyToAddress(*unmarshalled)
			if !isAddrExist(unjailMembers, addr) { // exclude jailed address
				continue
			}
		}
		tssMembers = append(tssMembers, hexEncoded)
	}

	return types.TssCommitteeInfo{
		ElectionId:    electionId.Uint64(),
		Threshold:     int(threshold.Int64()),
		ClusterPubKey: hex.EncodeToString(compressCPK),
		TssMembers:    tssMembers,
	}, nil
}

func isAddrExist(set []common.Address, find common.Address) bool {
	for _, s := range set {
		if bytes.Compare(s.Bytes(), find.Bytes()) == 0 {
			return true
		}
	}
	return false
}

func (q QueryService) QueryInactiveInfo() (types.TssCommitteeInfo, error) {
	currentBlockNumber, err := q.ethClient.BlockNumber(context.Background())
	if err != nil {
		return types.TssCommitteeInfo{}, err
	}
	electionId, threshold, inactiveTssMembers, err := q.tssGroupManagerCaller.GetTssInactiveGroupInfo(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(currentBlockNumber - q.confirmBlocks)})
	if len(inactiveTssMembers) == 0 {
		return types.TssCommitteeInfo{}, nil
	}
	tssMembers := make([]string, len(inactiveTssMembers), len(inactiveTssMembers))
	for i, m := range inactiveTssMembers {
		// raw public key(64bytes) ==> compress public key(33bytes)
		unmarshalled, err := crypto.UnmarshalPubkey(append([]byte{0x04}, m...))
		if err != nil {
			log.Error("fail to unmarshal tss member", "err", err)
			return types.TssCommitteeInfo{}, nil
		}
		compressed := crypto.CompressPubkey(unmarshalled)
		hexEncoded := hex.EncodeToString(compressed)

		// raw public key(64bytes) ==> uncompressed format: 0x04||rawPK (65bytes)
		// uncompressed := append([]byte{0x04}, m...)
		// hexEncoded := hex.EncodeToString(uncompressed)
		tssMembers[i] = hexEncoded
	}
	return types.TssCommitteeInfo{
		ElectionId: electionId.Uint64(),
		Threshold:  int(threshold.Int64()),
		TssMembers: tssMembers,
	}, nil
}

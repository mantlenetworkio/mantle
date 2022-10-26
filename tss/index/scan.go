package index

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/tss/bindings/scc"
)

var stateBatchAppendedTopicHash = crypto.Keccak256Hash([]byte("StateBatchAppended(uint256,bytes32,uint256,uint256,bytes,bytes)"))

func FilterStateBatchAppendedEvent(cli *ethclient.Client, startHeight, endHeight int64, contract common.Address) ([]map[[32]byte]uint64, error) {
	filter := ethereum.FilterQuery{
		FromBlock: big.NewInt(startHeight),
		ToBlock:   big.NewInt(endHeight),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{stateBatchAppendedTopicHash}},
	}
	logs, err := cli.FilterLogs(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	sccFilter, err := scc.NewStateCommitmentChainFilterer(contract, cli)
	if err != nil {
		return nil, err
	}

	result := make([]map[[32]byte]uint64, len(logs))
	for i, lg := range logs {
		event, err := sccFilter.ParseStateBatchAppended(lg)
		if err != nil {
			return nil, err
		}
		m := make(map[[32]byte]uint64)
		m[event.BatchRoot] = event.BatchIndex.Uint64()
		result[i] = m
		log.Info("got StateBatchAppended event", "batchRoot", hex.EncodeToString(event.BatchRoot[:]), "batchIndex", event.BatchIndex.Uint64())
	}
	return result, nil
}

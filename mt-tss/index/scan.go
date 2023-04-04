package index

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/log"
	"github.com/mantlenetworkio/mantle/mt-bindings/bindings"
)

var stateBatchAppendedTopicHash = crypto.Keccak256Hash([]byte("OutputProposed(bytes32,uint256,uint256,uint256)"))

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

	sccFilter, err := bindings.NewL2OutputOracleFilterer(contract, cli)
	if err != nil {
		return nil, err
	}

	result := make([]map[[32]byte]uint64, len(logs))
	for i, lg := range logs {
		event, err := sccFilter.ParseOutputProposed(lg)
		if err != nil {
			return nil, err
		}
		m := make(map[[32]byte]uint64)
		m[event.OutputRoot] = event.L2OutputIndex.Uint64()
		result[i] = m
		log.Info("got StateBatchAppended event", "batchRoot", hex.EncodeToString(event.OutputRoot[:]), "batchIndex", event.L2OutputIndex.Uint64())
	}
	return result, nil
}

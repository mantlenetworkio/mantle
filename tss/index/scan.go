package index

import (
	"context"
	"math/big"

	ethereum "github.com/mantlenetworkio/mantle/l2geth"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	ethcrypto "github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/ethclient"
)

var stateBatchAppendedTopicHash = ethcrypto.Keccak256Hash([]byte("PacketSent(bytes)"))

func FilterStateBatchAppendedEvent(cli *ethclient.Client, startHeight, endHeight int64, contract common.Address) ([]map[[32]byte]uint64, error) {
	filter := ethereum.FilterQuery{
		FromBlock: big.NewInt(startHeight),
		ToBlock:   big.NewInt(endHeight),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{stateBatchAppendedTopicHash}},
	}
	_, err := cli.FilterLogs(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

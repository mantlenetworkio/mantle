package clique

import (
	"github.com/bitdao-io/mantle/l2geth/consensus"
)

type API struct {
	chain  consensus.ChainReader
	clique *Clique
}

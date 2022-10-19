package clique

import (
	"github.com/mantlenetworkio/mantle/l2geth/consensus"
)

type API struct {
	chain  consensus.ChainReader
	clique *Clique
}

// Producers returns the current producer the node tries to produce block.
func (api *API) Producers() Producers {
	api.clique.lock.RLock()
	defer api.clique.lock.RUnlock()

	return api.clique.producers
}

package coterie

import (
	"github.com/mantlenetworkio/mantle/l2geth/consensus"
)

type API struct {
	chain   consensus.ChainReader
	coterie *Coterie
}

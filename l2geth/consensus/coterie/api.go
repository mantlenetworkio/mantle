package coterie

import (
	"github.com/bitdao-io/mantle/l2geth/consensus"
)

type API struct {
	chain   consensus.ChainReader
	coterie *Coterie
}

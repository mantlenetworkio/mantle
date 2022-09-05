package coterie

import (
	"github.com/bitdao-io/bitnetwork/l2geth/consensus"
)

type API struct {
	chain   consensus.ChainReader
	coterie *Coterie
}

package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Staker represents fraud-proof assertion stake status
type Staker struct {
	IsStaked         bool
	AmountStaked     *big.Int
	AssertionID      *big.Int
	CurrentChallenge common.Address
}

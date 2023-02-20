package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Staker represents fraud-proof assertion stake status
type Staker struct {
	IsStaked         bool
	AmountStaked     *big.Int
	AssertionID      *big.Int
	CurrentChallenge common.Address
}

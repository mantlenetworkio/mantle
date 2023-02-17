package types

import (
	"github.com/mantlenetworkio/mantle/fraud-proof/bindings"
	"math/big"

	"github.com/mantlenetworkio/mantle/l2geth/common"
)

// Assertion represents disputable assertion for L1 rollup contract
type Assertion struct {
	ID           *big.Int
	VmHash       common.Hash
	InboxSize    *big.Int
	GasUsed      *big.Int
	Parent       *big.Int
	Deadline     *big.Int
	ProposalTime *big.Int
}

func (a *Assertion) Copy() *Assertion {
	return &Assertion{
		ID:        new(big.Int).Set(a.ID),
		VmHash:    a.VmHash,
		GasUsed:   new(big.Int).Set(a.GasUsed),
		InboxSize: new(big.Int).Set(a.InboxSize),
		Deadline:  new(big.Int).Set(a.Deadline),
	}
}

func (a *Assertion) GetParentAssertion(assertionMap *bindings.AssertionMapCallerSession) (*Assertion, error) {
	ret, err := assertionMap.Assertions(a.Parent)
	if err != nil {
		return nil, err
	}
	return &Assertion{
		ID:           a.Parent,
		VmHash:       ret.StateHash,
		InboxSize:    ret.InboxSize,
		GasUsed:      ret.GasUsed,
		Parent:       ret.Parent,
		Deadline:     ret.Deadline,
		ProposalTime: ret.ProposalTime,
	}, nil
}
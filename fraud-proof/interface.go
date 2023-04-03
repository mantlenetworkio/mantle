package rollup

import "math/big"

type FraudProver interface {
	Start() error

	Stop() error

	CreateAssertionWithStateBatch([][32]byte, *big.Int, []byte, interface{}) error

	CreateAssertion(interface{}) error

	GetLatestAssertion() (interface{}, error)

	InChallenge() bool

	RespondChallenge() error

	GenerateState() (interface{}, error)
}

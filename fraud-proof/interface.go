package rollup

import "math/big"

type FraudProover interface {
	Start() error
	Stop() error

	CreateAssertion(interface{}) error

	CreateAssertionWithStateBatch([][32]byte, *big.Int, []byte, interface{}) error

	GetLatestAssertion() (interface{}, error)

	InChallenge() bool

	RespondChallenge() error

	GenerateState() (interface{}, error)
}

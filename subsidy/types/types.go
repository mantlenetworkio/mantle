package types

import "time"

type PayerState struct {
	LastPayTime time.Time
	EndBlock    uint64
	PayTxHash   string
}

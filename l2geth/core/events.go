// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
)

type PeerAddEvent struct {
	PeerId []byte
	Has    chan bool
}

type BatchPeriodStartEvent struct {
	Msg   *types.BatchPeriodStartMsg
	ErrCh chan error
}

type L1ToL2TxStartEvent struct {
	ErrCh       chan error
	SchedulerCh chan struct{}
}

type BatchPeriodAnswerEvent struct {
	Msg   *types.BatchPeriodAnswerMsg
	ErrCh chan error
}

type BatchEndEvent struct{}

type FraudProofReorgEvent struct {
	Msg   *types.FraudProofReorgMsg
	ErrCh chan error
}

// NewTxsEvent is posted when a batch of transactions enter the transaction pool.
type NewTxsEvent struct {
	Txs        []*types.Transaction
	TxSetProof *types.BatchTxSetProof
	ErrCh      chan error
}

// NewMinedBlockEvent is posted when a block has been imported.
type NewMinedBlockEvent struct{ Block *types.Block }

// RemovedLogsEvent is posted when a reorg happens
type RemovedLogsEvent struct{ Logs []*types.Log }

type ChainEvent struct {
	Block *types.Block
	Hash  common.Hash
	Logs  []*types.Log
}

type ChainSideEvent struct {
	Block *types.Block
}

type ChainHeadEvent struct{ Block *types.Block }

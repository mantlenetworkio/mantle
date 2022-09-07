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

package eth

import (
	"github.com/bitdao-io/mantle/l2geth/consensus/coterie"
)

// consensusProtocolLength is the number of implemented message.
var consensusProtocolLength = 2

const consensusMaxMsgSize = 10 * 1024 * 1024 // Maximum cap on the size of a protocol message

// eth protocol message codes
const (
	GetProducersMsg = 0x01
	ProducersMsg    = 0x02
)

// getProducersData represents a producers query.
type getProducersData struct {
	Number uint64 // Block hash from which to retrieve producers (excludes Hash)
}

// producersDate represents a producer set.
type producersData struct {
	Number       uint64 // Number represents Block hash from which to retrieve producers (excludes Hash)
	Epoch        uint64 // Epoch represents the block number for each producer
	schedulerId  string // schedulerId represents scheduler's peer.id
	SequencerSet coterie.SequencerSet
}

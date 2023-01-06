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

// consensusProtocolLength is the number of implemented message.
var consensusProtocolLength = uint64(3)

const consensusMaxMsgSize = 10 * 1024 * 1024 // Maximum cap on the size of a protocol message

// protocolName is the official short name of the protocol used during capability negotiation.
const consensusProtocolName = "clique"

// eth protocol message codes
const (
	BatchPeriodStartMsg  = 0x00
	BatchPeriodAnswerMsg = 0x01
	FraudProofReorgMsg   = 0x02
)

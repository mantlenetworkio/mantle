// Copyright 2022, Specular contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package state

import (
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/rlp"
)

type LogSeries struct {
	Logs           []*types.Log
	AccumulateHash common.Hash
	Bloom          types.Bloom
}

func LogSeriesFromLogs(logs []*types.Log) *LogSeries {
	ls := make([]*types.Log, len(logs))
	copy(ls, logs)
	h := common.Hash{}
	for _, l := range ls {
		// TODO: should we check rlp encode error here?
		logBytes, _ := rlp.EncodeToBytes(l)
		h = crypto.Keccak256Hash(h.Bytes(), logBytes)
	}
	bin := types.LogsBloom(ls)
	return &LogSeries{
		Logs:           ls,
		AccumulateHash: h,
		Bloom:          types.BytesToBloom(bin.Bytes()),
	}
}

func (l *LogSeries) Hash() common.Hash {
	return common.BytesToHash(l.EncodeState())
}

func (l *LogSeries) EncodeState() []byte {
	return crypto.Keccak256(l.AccumulateHash.Bytes(), l.Bloom.Bytes())
}

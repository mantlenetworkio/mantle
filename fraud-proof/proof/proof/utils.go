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

package proof

import (
	"strings"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
	"github.com/mantlenetworkio/mantle/l2geth/params"
)

func GetAccountProof(
	db vm.StateDB,
	address common.Address,
) (*MPTProof, error) {
	accountProof, err := db.GetProof(address)
	if err != nil {
		return nil, err
	}
	return &MPTProof{accountProof}, nil
}

func GetStorageProof(
	db vm.StateDB,
	address common.Address,
	key common.Hash,
) (*MPTProof, *MPTProof, error) {
	accountProof, err := db.GetProof(address)
	if err != nil {
		return nil, nil, err
	}
	storageProof, err := db.GetStorageProof(address, key)
	if err != nil {
		return nil, nil, err
	}
	return &MPTProof{accountProof}, &MPTProof{storageProof}, nil
}

func calcCellNum(offset, size uint64) uint64 {
	return (offset+size+31)/32 - offset/32
}

func IsStackError(err error) bool {
	if err == nil {
		return false
	}
	return strings.HasPrefix(err.Error(), "stack")
}

func IsStopTokenError(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == "stop token"
}

func precompile(rules params.Rules, addr common.Address) (vm.PrecompiledContract, bool) {
	var precompiles map[common.Address]vm.PrecompiledContract
	switch {
	case rules.IsBerlin:
		precompiles = vm.PrecompiledContractsBerlin
	case rules.IsIstanbul:
		precompiles = vm.PrecompiledContractsIstanbul
	case rules.IsByzantium:
		precompiles = vm.PrecompiledContractsByzantium
	default:
		precompiles = vm.PrecompiledContractsHomestead
	}
	p, ok := precompiles[addr]
	return p, ok
}

func isEnvironmentalOp(opcode vm.OpCode) bool {
	return opcode == vm.ADDRESS ||
		opcode == vm.ORIGIN ||
		opcode == vm.CALLER ||
		opcode == vm.CALLVALUE ||
		opcode == vm.CODESIZE ||
		opcode == vm.CALLDATASIZE ||
		opcode == vm.GASPRICE ||
		opcode == vm.BLOCKHASH ||
		opcode == vm.COINBASE ||
		opcode == vm.TIMESTAMP ||
		opcode == vm.NUMBER ||
		opcode == vm.DIFFICULTY ||
		opcode == vm.GASLIMIT ||
		opcode == vm.CHAINID
}

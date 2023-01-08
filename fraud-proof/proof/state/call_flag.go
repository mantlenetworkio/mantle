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
	"fmt"

	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

type CallFlag byte

const (
	CALLFLAG_CALL         = 0
	CALLFLAG_CALLCODE     = 1
	CALLFLAG_DELEGATECALL = 2
	CALLFLAG_STATICCALL   = 3
	CALLFLAG_CREATE       = 4
	CALLFLAG_CREATE2      = 5
)

func (f CallFlag) IsCreate() bool {
	return f == CALLFLAG_CREATE || f == CALLFLAG_CREATE2
}

func OpCodeToCallFlag(op vm.OpCode) CallFlag {
	if op == vm.CALL {
		return CALLFLAG_CALL
	} else if op == vm.CALLCODE {
		return CALLFLAG_CALLCODE
	} else if op == vm.DELEGATECALL {
		return CALLFLAG_DELEGATECALL
	} else if op == vm.STATICCALL {
		return CALLFLAG_STATICCALL
	} else if op == vm.CREATE {
		return CALLFLAG_CREATE
	} else if op == vm.CREATE2 {
		return CALLFLAG_CREATE2
	}
	log.Error(fmt.Sprintf("Opcode %s is not call", op))
	panic(fmt.Sprintf("Opcode %s is not call", op))
}

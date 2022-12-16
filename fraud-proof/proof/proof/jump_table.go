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
	"github.com/mantlenetworkio/mantle/fraud-proof/proof/state"
	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
)

type genProofFunc func(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error)

type proofGen struct {
	genProof genProofFunc
}

type jumpTable [256]*proofGen

var proofJumpTable = newProofJumpTable()

func newProofJumpTable() jumpTable {
	tbl := jumpTable{
		vm.STOP: {
			genProof: opStopProof,
		},
		vm.ADD: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.MUL: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.SUB: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.DIV: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.SDIV: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.MOD: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.SMOD: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.ADDMOD: {
			genProof: makeStackOnlyInstructionProof(3),
		},
		vm.MULMOD: {
			genProof: makeStackOnlyInstructionProof(3),
		},
		vm.EXP: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.SIGNEXTEND: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.LT: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.GT: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.SLT: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.SGT: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.EQ: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.ISZERO: {
			genProof: makeStackOnlyInstructionProof(1),
		},
		vm.AND: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.OR: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.XOR: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.NOT: {
			genProof: makeStackOnlyInstructionProof(1),
		},
		vm.BYTE: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.SHL: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.SHR: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.SAR: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.SHA3: {
			genProof: opKeccak256Proof,
		},
		vm.ADDRESS: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.BALANCE: {
			genProof: opBalanceProof,
		},
		vm.ORIGIN: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.CALLER: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.CALLVALUE: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.CALLDATALOAD: {
			genProof: opCallDataLoadProof,
		},
		vm.CALLDATASIZE: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.CALLDATACOPY: {
			genProof: opCallDataCopyProof,
		},
		vm.CODESIZE: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.CODECOPY: {
			genProof: opCodeCopyProof,
		},
		vm.GASPRICE: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.EXTCODESIZE: {
			genProof: opExtCodeSizeProof,
		},
		vm.EXTCODECOPY: {
			genProof: opExtCodeCopyProof,
		},
		vm.RETURNDATASIZE: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.RETURNDATACOPY: {
			genProof: opReturnDataCopyProof,
		},
		vm.EXTCODEHASH: {
			genProof: opExtCodeHashProof,
		},
		vm.BLOCKHASH: {
			genProof: opBlockHashProof,
		},
		vm.COINBASE: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.TIMESTAMP: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.NUMBER: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.DIFFICULTY: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.GASLIMIT: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.CHAINID: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.SELFBALANCE: {
			genProof: opSelfBalanceProof,
		},
		// vm.BASEFEE: {
		// 	genProof: makeStackOnlyInstructionProof(0),
		// },
		vm.POP: {
			genProof: makeStackOnlyInstructionProof(1),
		},
		vm.MLOAD: {
			genProof: opMLoadProof,
		},
		vm.MSTORE: {
			genProof: opMStoreProof,
		},
		vm.MSTORE8: {
			genProof: opMStore8Proof,
		},
		vm.SLOAD: {
			genProof: opSLoadProof,
		},
		vm.SSTORE: {
			genProof: opSStoreProof,
		},
		vm.JUMP: {
			genProof: makeStackOnlyInstructionProof(1),
		},
		vm.JUMPI: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.PC: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.MSIZE: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.GAS: {
			genProof: makeStackOnlyInstructionProof(0),
		},
		vm.JUMPDEST: {
			genProof: opJumpDestProof,
		},
		vm.PUSH1: {
			genProof: opPushProof,
		},
		vm.PUSH2: {
			genProof: opPushProof,
		},
		vm.PUSH3: {
			genProof: opPushProof,
		},
		vm.PUSH4: {
			genProof: opPushProof,
		},
		vm.PUSH5: {
			genProof: opPushProof,
		},
		vm.PUSH6: {
			genProof: opPushProof,
		},
		vm.PUSH7: {
			genProof: opPushProof,
		},
		vm.PUSH8: {
			genProof: opPushProof,
		},
		vm.PUSH9: {
			genProof: opPushProof,
		},
		vm.PUSH10: {
			genProof: opPushProof,
		},
		vm.PUSH11: {
			genProof: opPushProof,
		},
		vm.PUSH12: {
			genProof: opPushProof,
		},
		vm.PUSH13: {
			genProof: opPushProof,
		},
		vm.PUSH14: {
			genProof: opPushProof,
		},
		vm.PUSH15: {
			genProof: opPushProof,
		},
		vm.PUSH16: {
			genProof: opPushProof,
		},
		vm.PUSH17: {
			genProof: opPushProof,
		},
		vm.PUSH18: {
			genProof: opPushProof,
		},
		vm.PUSH19: {
			genProof: opPushProof,
		},
		vm.PUSH20: {
			genProof: opPushProof,
		},
		vm.PUSH21: {
			genProof: opPushProof,
		},
		vm.PUSH22: {
			genProof: opPushProof,
		},
		vm.PUSH23: {
			genProof: opPushProof,
		},
		vm.PUSH24: {
			genProof: opPushProof,
		},
		vm.PUSH25: {
			genProof: opPushProof,
		},
		vm.PUSH26: {
			genProof: opPushProof,
		},
		vm.PUSH27: {
			genProof: opPushProof,
		},
		vm.PUSH28: {
			genProof: opPushProof,
		},
		vm.PUSH29: {
			genProof: opPushProof,
		},
		vm.PUSH30: {
			genProof: opPushProof,
		},
		vm.PUSH31: {
			genProof: opPushProof,
		},
		vm.PUSH32: {
			genProof: opPushProof,
		},
		vm.DUP1: {
			genProof: makeStackOnlyInstructionProof(1),
		},
		vm.DUP2: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.DUP3: {
			genProof: makeStackOnlyInstructionProof(3),
		},
		vm.DUP4: {
			genProof: makeStackOnlyInstructionProof(4),
		},
		vm.DUP5: {
			genProof: makeStackOnlyInstructionProof(5),
		},
		vm.DUP6: {
			genProof: makeStackOnlyInstructionProof(6),
		},
		vm.DUP7: {
			genProof: makeStackOnlyInstructionProof(7),
		},
		vm.DUP8: {
			genProof: makeStackOnlyInstructionProof(8),
		},
		vm.DUP9: {
			genProof: makeStackOnlyInstructionProof(9),
		},
		vm.DUP10: {
			genProof: makeStackOnlyInstructionProof(10),
		},
		vm.DUP11: {
			genProof: makeStackOnlyInstructionProof(11),
		},
		vm.DUP12: {
			genProof: makeStackOnlyInstructionProof(12),
		},
		vm.DUP13: {
			genProof: makeStackOnlyInstructionProof(13),
		},
		vm.DUP14: {
			genProof: makeStackOnlyInstructionProof(14),
		},
		vm.DUP15: {
			genProof: makeStackOnlyInstructionProof(15),
		},
		vm.DUP16: {
			genProof: makeStackOnlyInstructionProof(16),
		},
		vm.SWAP1: {
			genProof: makeStackOnlyInstructionProof(2),
		},
		vm.SWAP2: {
			genProof: makeStackOnlyInstructionProof(3),
		},
		vm.SWAP3: {
			genProof: makeStackOnlyInstructionProof(4),
		},
		vm.SWAP4: {
			genProof: makeStackOnlyInstructionProof(5),
		},
		vm.SWAP5: {
			genProof: makeStackOnlyInstructionProof(6),
		},
		vm.SWAP6: {
			genProof: makeStackOnlyInstructionProof(7),
		},
		vm.SWAP7: {
			genProof: makeStackOnlyInstructionProof(8),
		},
		vm.SWAP8: {
			genProof: makeStackOnlyInstructionProof(9),
		},
		vm.SWAP9: {
			genProof: makeStackOnlyInstructionProof(10),
		},
		vm.SWAP10: {
			genProof: makeStackOnlyInstructionProof(11),
		},
		vm.SWAP11: {
			genProof: makeStackOnlyInstructionProof(12),
		},
		vm.SWAP12: {
			genProof: makeStackOnlyInstructionProof(13),
		},
		vm.SWAP13: {
			genProof: makeStackOnlyInstructionProof(14),
		},
		vm.SWAP14: {
			genProof: makeStackOnlyInstructionProof(15),
		},
		vm.SWAP15: {
			genProof: makeStackOnlyInstructionProof(16),
		},
		vm.SWAP16: {
			genProof: makeStackOnlyInstructionProof(17),
		},
		vm.LOG0: {
			genProof: makeLogProof(0),
		},
		vm.LOG1: {
			genProof: makeLogProof(1),
		},
		vm.LOG2: {
			genProof: makeLogProof(2),
		},
		vm.LOG3: {
			genProof: makeLogProof(3),
		},
		vm.LOG4: {
			genProof: makeLogProof(4),
		},
		vm.CREATE: {
			genProof: opCreateProof,
		},
		vm.CALL: {
			genProof: opCallProof,
		},
		vm.CALLCODE: {
			genProof: opCallCodeProof,
		},
		vm.RETURN: {
			genProof: opReturnProof,
		},
		vm.DELEGATECALL: {
			genProof: opDelegateCallProof,
		},
		vm.CREATE2: {
			genProof: opCreate2Proof,
		},
		vm.STATICCALL: {
			genProof: opStaticCallProof,
		},
		vm.REVERT: {
			genProof: opRevertProof,
		},
		vm.SELFDESTRUCT: {
			genProof: opSelfDestructProof,
		},
	}

	for i, entry := range tbl {
		if entry == nil {
			tbl[i] = &proofGen{genProof: opInvalidProof}
		}
	}

	return tbl
}

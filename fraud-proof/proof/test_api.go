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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/mantlenetworkio/mantle/fraud-proof/proof/prover"
	oss "github.com/mantlenetworkio/mantle/fraud-proof/proof/state"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core"
	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
	"github.com/mantlenetworkio/mantle/l2geth/rpc"
)

func (api *ProverAPI) GenerateProofForTest(ctx context.Context, hash common.Hash, cumulativeGasUsed, blockGasUsed *big.Int, step uint64, config *ProverConfig) (json.RawMessage, error) {
	transaction, blockHash, blockNumber, index, err := api.backend.GetTransaction(ctx, hash)
	if err != nil {
		return nil, err
	}
	block, err := api.backend.BlockByNumber(ctx, rpc.BlockNumber(blockNumber))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("block #%d not found", blockNumber)
	}
	parent, err := api.backend.BlockByNumber(ctx, rpc.BlockNumber(blockNumber-1))
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, fmt.Errorf("parent block #%d not found", blockNumber-1)
	}
	// It shouldn't happen in practice.
	if blockNumber == 0 {
		return nil, errors.New("genesis is not traceable")
	}
	reexec := defaultProveReexec
	if config != nil && config.Reexec != nil {
		reexec = *config.Reexec
	}
	// get tx ctx
	msg, txContext, statedb, err := api.backend.StateAtTransaction(ctx, block, int(index), reexec)
	if err != nil {
		return nil, err
	}
	receipts, err := api.backend.GetReceipts(ctx, blockHash)
	if err != nil {
		return nil, err
	}
	// get block ctx
	chainCtx := createChainContext(api.backend, ctx)
	vmctx := core.NewEVMBlockContext(parent.Header(), chainCtx, nil)
	// calc block state hash
	blockHashTree, err := oss.BlockHashTreeFromBlockContext(&vmctx)
	if err != nil {
		return nil, err
	}
	// calc its
	its := oss.InterStateFromCaptured(
		blockNumber,
		index,
		statedb,
		blockGasUsed,
		block.Transactions(),
		receipts,
		blockHashTree,
	)
	// new test prover
	testProver := prover.NewTestProver(
		step,
		transaction,
		&txContext,
		receipts[index],
		api.backend.ChainConfig().Rules(vmctx.BlockNumber),
		blockNumber,
		index,
		statedb,
		*its,
		blockHashTree,
	)
	// new evm
	vmenv := vm.NewEVM(txContext, statedb, api.backend.ChainConfig(), vm.Config{Debug: true, Tracer: testProver})
	statedb.Prepare(hash, block.Hash(), int(index))
	_, _, _, err = core.ApplyMessage(vmenv, msg, new(core.GasPool).AddGas(msg.Gas()))
	if err != nil {
		return nil, fmt.Errorf("tracing failed: %w", err)
	}
	return testProver.GetResult()
}

package signer

import (
	"encoding/json"
	"github.com/bitdao-io/bitnetwork/l2geth/common/hexutil"
	tsstypes "github.com/bitdao-io/bitnetwork/tss/types"
	"github.com/rs/zerolog"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"math/big"
	"sync"
)

func (p *Processor) Verify() {
	defer p.wg.Done()
	logger := p.logger.With().Str("step", "verify event").Logger()
	logger.Info().Msg("start to verify events ")

	go func() {
		defer func() {
			logger.Info().Msg("exit verify event process")
		}()
		for {
			select {
			case <-p.stopChan:
				return
			case req := <-p.askRequestChan:
				var askRequest tsstypes.AskStateRequest
				if err := json.Unmarshal(req.Params, &askRequest); err != nil {
					logger.Error().Msg("failed to unmarshal ask request")
					continue
				}
				var resId = req.ID

				// if offset value does not equal  length of state roots array, return false
				if len(askRequest.StateRoots) != int(askRequest.OffsetStartsAtIndex.Int64()) {
					askResponse := tsstypes.AskResponse{
						Result: false,
					}

					RpcResponse := tdtypes.NewRPCSuccessResponse(resId, askResponse)
					p.wsClient.SendMsg(RpcResponse)
				} else {
					wg := &sync.WaitGroup{}
					wg.Add(int(askRequest.OffsetStartsAtIndex.Int64()))
					var result = true
					var rLock = &sync.Mutex{}
					quit := make(chan struct{})
					for index, stateRoot := range askRequest.StateRoots {
						go p.verify(askRequest.StartBlock, index, stateRoot, logger, wg, result, rLock, quit)
					}
					wg.Wait()
					if result {
						p.UpdateWaitSignEvents(resId.(tdtypes.JSONRPCStringID).String(), askRequest)
					}
					askResponse := tsstypes.AskResponse{
						Result: result,
					}
					RpcResponse := tdtypes.NewRPCSuccessResponse(resId, askResponse)
					p.wsClient.SendMsg(RpcResponse)

				}

			}

		}
	}()

}

func (p *Processor) verify(start big.Int, index int, stateRoot [32]byte, logger zerolog.Logger, wg *sync.WaitGroup, result bool, lock *sync.Mutex, quit chan struct{}) {
	defer wg.Done()
	defer logger.Info().Msgf("start block number:(%s),index (%s), verify done", start.Int64(), index)

	select {
	case <-quit:
		return
	default:
		offset := new(big.Int).SetInt64(int64(index))
		blockNumber := offset.Add(offset, &start)
		logger.Info().Msgf("verify block number %s", blockNumber)
		lock.Lock()
		defer lock.Unlock()
		block, err := p.l2Client.BlockByNumber(p.ctx, blockNumber)
		if err != nil {
			logger.Err(err).Msgf("failed to get block by (%s) ", blockNumber)
			result = false
			quit <- struct{}{}
		} else {
			if hexutil.Encode(stateRoot[:]) != block.Root().String() {
				logger.Info().Msgf("block number (%s) stateroot doesn't same", blockNumber)
				result = false
				quit <- struct{}{}
			}
		}
	}
}

func (p *Processor) UpdateWaitSignEvents(uniqueId string, msg tsstypes.AskStateRequest) {
	p.waitSignLock.Lock()
	defer p.waitSignLock.Unlock()
	p.waitSignMsgs[uniqueId] = msg
}

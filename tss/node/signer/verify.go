package signer

import (
	"encoding/json"
	"math/big"
	"sync"

	"github.com/mantlenetworkio/mantle/l2geth/common/hexutil"
	"github.com/mantlenetworkio/mantle/tss/common"
	"github.com/rs/zerolog"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
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
				var askRequest common.SignStateRequest
				var RpcResponse tdtypes.RPCResponse
				if err := json.Unmarshal(req.Params, &askRequest); err != nil {
					logger.Error().Msg("failed to unmarshal ask request")
					RpcResponse = tdtypes.NewRPCErrorResponse(req.ID, 201, "failed to unmarshal ", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}
				var resId = req.ID
				offset := len(askRequest.StateRoots)

				wg := &sync.WaitGroup{}
				wg.Add(offset)
				var result = true
				var rLock = &sync.Mutex{}
				quit := make(chan struct{})
				for index, stateRoot := range askRequest.StateRoots {
					go p.verify(askRequest.StartBlock, index, stateRoot, logger, wg, result, rLock, quit)
				}
				wg.Wait()
				if result {
					hash, err := signMsgToHash(askRequest)
					if err != nil {
						logger.Err(err).Msg("failed to conv msg to hash")
						RpcResponse = tdtypes.NewRPCErrorResponse(req.ID, 201, "failed to conv msg to hash", err.Error())
						p.wsClient.SendMsg(RpcResponse)
						continue
					} else {
						hashStr := hexutil.Encode(hash)
						p.UpdateWaitSignEvents(hashStr, askRequest)
					}
				}
				askResponse := common.AskResponse{
					Result: result,
				}
				RpcResponse = tdtypes.NewRPCSuccessResponse(resId, askResponse)
				p.wsClient.SendMsg(RpcResponse)

			}

		}
	}()

}

func (p *Processor) verify(start string, index int, stateRoot [32]byte, logger zerolog.Logger, wg *sync.WaitGroup, result bool, lock *sync.Mutex, quit chan struct{}) {
	defer wg.Done()
	defer logger.Info().Msgf("start block number:(%s),index (%s), verify done", start, index)

	select {
	case <-quit:
		return
	default:
		offset := new(big.Int).SetInt64(int64(index))
		startBig, _ := new(big.Int).SetString(start, 10)
		blockNumber := offset.Add(offset, startBig)
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

func (p *Processor) UpdateWaitSignEvents(uniqueId string, msg common.SignStateRequest) {
	p.waitSignLock.Lock()
	defer p.waitSignLock.Unlock()
	p.waitSignMsgs[uniqueId] = msg
}

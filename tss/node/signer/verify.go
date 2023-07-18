package signer

import (
	"encoding/json"
	"math/big"
	"sync"

	"github.com/rs/zerolog"

	"github.com/mantlenetworkio/mantle/l2geth/common/hexutil"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/tss/common"

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
					if err := p.wsClient.SendMsg(RpcResponse); err != nil {
						logger.Error().Err(err).Msg("failed to send msg to manager")
					}
					continue
				}
				if askRequest.StartBlock == nil ||
					askRequest.OffsetStartsAtIndex == nil ||
					askRequest.StartBlock.Cmp(big.NewInt(0)) < 0 ||
					askRequest.OffsetStartsAtIndex.Cmp(big.NewInt(0)) < 0 {
					logger.Error().Msg("StartBlock and OffsetStartsAtIndex must not be nil or negative")
					RpcResponse = tdtypes.NewRPCErrorResponse(req.ID, 201, "invalid askRequest", "StartBlock and OffsetStartsAtIndex must not be nil or negative")
					if err := p.wsClient.SendMsg(RpcResponse); err != nil {
						logger.Error().Err(err).Msg("failed to send msg to manager")
					}
					return
				}
				var resId = req.ID
				var size = len(askRequest.StateRoots)
				logger.Info().Msgf("stateroots size %d ", size)
				if len(askRequest.StateRoots) == 0 {
					logger.Error().Msg("stateroots size is empty")
					RpcResponse = tdtypes.NewRPCErrorResponse(req.ID, 201, "stateroots size is empty ", "do not need to sign")
					if err := p.wsClient.SendMsg(RpcResponse); err != nil {
						logger.Error().Err(err).Msg("failed to send msg to manager")
					}
					continue
				} else {
					wg := &sync.WaitGroup{}
					wg.Add(1)
					result, err := p.verify(askRequest.StartBlock, size-1, askRequest.StateRoots[size-1], logger, wg)
					if !result {
						if err != nil {
							logger.Error().Msgf("failed to verify block %s", err.Error())
							RpcResponse = tdtypes.NewRPCErrorResponse(req.ID, 201, "get error when verify ", err.Error())
							if err := p.wsClient.SendMsg(RpcResponse); err != nil {
								logger.Error().Err(err).Msg("failed to send msg to manager")
							}
							continue
						}
					} else {
						hash, err := signMsgToHash(askRequest)
						if err != nil {
							logger.Err(err).Msg("failed to conv msg to hash")
							RpcResponse = tdtypes.NewRPCErrorResponse(req.ID, 201, "failed to conv msg to hash", err.Error())
							if err := p.wsClient.SendMsg(RpcResponse); err != nil {
								logger.Error().Err(err).Msg("failed to send msg to manager")
							}
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
					if err := p.wsClient.SendMsg(RpcResponse); err != nil {
						logger.Error().Err(err).Msg("failed to send msg to manager")
					}

				}
			}

		}
	}()
}

func (p *Processor) verify(start *big.Int, index int, stateRoot [32]byte, logger zerolog.Logger, wg *sync.WaitGroup) (bool, error) {
	defer wg.Done()

	offset := new(big.Int).SetInt64(int64(index))
	blockNumber := offset.Add(offset, start)
	logger.Info().Msgf("start to query block by number %d", blockNumber)

	value, ok := p.GetVerify(blockNumber.String())
	if ok {
		if value {
			return value, nil
		}
	}

	var block *types.Block
	var err error
	for i := 0; i < 3; i++ {
		block, err = p.l2Client.BlockByNumber(p.ctx, blockNumber)
		if err == nil {
			break
		} else {
			logger.Info().Msgf("retry to query block by number %d, times %d", blockNumber, i)
		}
	}
	if err != nil {
		logger.Err(err).Msgf("failed to get block by (%d) ", blockNumber)
		return false, err
	} else {
		if hexutil.Encode(stateRoot[:]) != block.Root().String() {
			logger.Info().Msgf("block number (%d) state root doesn't same, state root (%s) , block root (%s)", blockNumber, hexutil.Encode(stateRoot[:]), block.Root().String())
			p.CacheVerify(blockNumber.String(), false)
			return false, nil
		} else {
			logger.Info().Msgf("block number (%d) verify success", blockNumber)
			p.CacheVerify(blockNumber.String(), true)
			return true, nil
		}
	}
}

func (p *Processor) UpdateWaitSignEvents(uniqueId string, msg common.SignStateRequest) {
	p.waitSignLock.Lock()
	defer p.waitSignLock.Unlock()
	p.waitSignMsgs[uniqueId] = msg
}

func (p *Processor) CacheVerify(key string, value bool) bool {
	p.cacheVerifyLock.Lock()
	defer p.cacheVerifyLock.Unlock()
	return p.cacheVerify.Set(key, value)
}

func (p *Processor) GetVerify(key string) (bool, bool) {
	p.cacheVerifyLock.RLock()
	defer p.cacheVerifyLock.RUnlock()
	return p.cacheVerify.Get(key)
}

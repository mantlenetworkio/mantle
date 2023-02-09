package signer

import (
	"encoding/json"
	"math/big"
	"sync"

	"github.com/mantlenetworkio/mantle/l2geth/common/hexutil"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
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
				var getBlockErr = false
				var resultErr error
				for index, stateRoot := range askRequest.StateRoots {
					go func(fIndex int, fStateRoot [32]byte) {
						resultTmp, err := p.verify(askRequest.StartBlock, fIndex, fStateRoot, logger, wg)
						if !resultTmp {
							result = resultTmp
							if err != nil {
								logger.Error().Msgf("failed to verify block %s", err.Error())
								getBlockErr = true
								resultErr = err
							}
						}
					}(index, stateRoot)
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

				if getBlockErr {
					RpcResponse = tdtypes.NewRPCErrorResponse(req.ID, 201, "get error when verify ", resultErr.Error())
					p.wsClient.SendMsg(RpcResponse)
				} else {
					askResponse := common.AskResponse{
						Result: result,
					}
					RpcResponse = tdtypes.NewRPCSuccessResponse(resId, askResponse)
					p.wsClient.SendMsg(RpcResponse)
				}

			}

		}
	}()

}

func (p *Processor) verify(start string, index int, stateRoot [32]byte, logger zerolog.Logger, wg *sync.WaitGroup) (bool, error) {
	defer wg.Done()
	defer logger.Info().Msgf("start block number:(%s),index (%d), verify done", start, index)

	offset := new(big.Int).SetInt64(int64(index))
	startBig, _ := new(big.Int).SetString(start, 10)
	blockNumber := offset.Add(offset, startBig)
	logger.Info().Msgf("verify block number %d", blockNumber)

	value, ok := p.cacheVerify.Get(blockNumber.String())
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
			bol := p.CacheVerify(blockNumber.String(), false)
			logger.Info().Msgf("cache verify behavior %s ", bol)
			return false, nil
		} else {
			logger.Info().Msgf("block number (%d) verify success", blockNumber)
			bol := p.CacheVerify(blockNumber.String(), true)
			logger.Info().Msgf("cache verify behavior %s ", bol)
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

//func group(slice [][32]byte, segment int) [][][32]byte {
//	size := len(slice)
//	var segments = make([][][32]byte, 0)
//	quantity := size / segment
//	remainder := size % segment
//	i := 0
//	for ; i < quantity; i++ {
//		segments = append(segments, slice[i*segment:(i+1)*segment])
//	}
//	if remainder != 0 {
//		segments = append(segments, slice[i*segment:i*segment+remainder])
//	}
//	return segments
//}

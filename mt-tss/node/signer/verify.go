package signer

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/mantlenetworkio/mantle/mt-tss/common"
	"github.com/rs/zerolog"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"math/big"
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
				var askRequest common.SignOutputRequest
				var RpcResponse tdtypes.RPCResponse
				if err := json.Unmarshal(req.Params, &askRequest); err != nil {
					logger.Error().Msg("failed to unmarshal ask request")
					RpcResponse = tdtypes.NewRPCErrorResponse(req.ID, 201, "failed to unmarshal ", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}
				var resId = req.ID

				result, err := p.verify(askRequest, logger)
				if err != nil {
					RpcResponse = tdtypes.NewRPCErrorResponse(req.ID, 201, "get error when verify ", resultErr.Error())
					p.wsClient.SendMsg(RpcResponse)
				} else {
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

		}
	}()

}

func (p *Processor) verify(outputRequest common.SignOutputRequest, logger zerolog.Logger) (bool, error) {

	l2BlockNumber := new(big.Int).SetUint64(outputRequest.L2BlockNumber)
	value, ok := p.GetVerify(l2BlockNumber.String())
	if ok {
		if value {
			return value, nil
		}
	}

	output, err := p.rollupClient.OutputAtBlock(context.Background(), l2BlockNumber.Uint64())

	if err != nil {
		logger.Err(err).Msgf("failed to fetch output at block (%d) ", l2BlockNumber)
		return false, err
	} else {
		if hexutil.Encode(outputRequest.OutputRoot[:]) != output.OutputRoot.String() {
			logger.Info().Msgf("block number (%d) output doesn't same, request root (%s) , block root (%s)", l2BlockNumber, hexutil.Encode(outputRequest.OutputRoot[:]), output.OutputRoot.String())
			p.CacheVerify(l2BlockNumber.String(), false)
			return false, nil
		} else {
			logger.Info().Msgf("block number (%d) verify success", l2BlockNumber)
			p.CacheVerify(l2BlockNumber.String(), true)
			return true, nil
		}
	}
}

func (p *Processor) UpdateWaitSignEvents(uniqueId string, msg common.SignOutputRequest) {
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

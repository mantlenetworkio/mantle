package signer

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/mantlenetworkio/mantle/l2geth/common/hexutil"
	tsscommon "github.com/mantlenetworkio/mantle/tss/common"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func (p *Processor) SignRollBack() {
	defer p.wg.Done()
	logger := p.logger.With().Str("step", "sign Roll Back Message").Logger()

	logger.Info().Msg("start to sign roll back message ")

	go func() {
		defer func() {
			logger.Info().Msg("exit sign roll back process")
		}()
		for {
			select {
			case <-p.stopChan:
				return
			case req := <-p.signRollBachChan:
				var resId = req.ID.(tdtypes.JSONRPCStringID).String()
				logger.Info().Msgf("dealing resId (%s) ", resId)

				var nodeSignRequest tsscommon.NodeSignRequest
				rawMsg := json.RawMessage{}
				nodeSignRequest.RequestBody = &rawMsg

				if err := json.Unmarshal(req.Params, &nodeSignRequest); err != nil {
					logger.Error().Msg("failed to unmarshal roll back request")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}
				var requestBody tsscommon.RollBackRequest
				if err := json.Unmarshal(rawMsg, &requestBody); err != nil {
					logger.Error().Msg("failed to umarshal roll back params request body")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}
				nodeSignRequest.RequestBody = requestBody
				startBlock, _ := new(big.Int).SetString(requestBody.StartBlock, 10)
				hashTx, err := tsscommon.RollBackHash(startBlock)
				if err != nil {
					logger.Err(err).Msg("failed to encode roll back msg")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}

				var signResponse tsscommon.SignResponse

				hashStr := hexutil.Encode(hashTx)
				signByte, ok := p.GetSign(hashStr)
				if ok {
					logger.Info().Msg("singer get roll back signature from cache")
					signResponse = tsscommon.SignResponse{
						Signature: signByte,
					}
				} else {
					data, culprits, err := p.handleSign(nodeSignRequest, hashTx, logger)

					if err != nil {
						logger.Error().Msgf("roll back %s sign failed ", requestBody.StartBlock)
						var errorRes tdtypes.RPCResponse
						if len(culprits) > 0 {
							respData := strings.Join(culprits, ",")
							errorRes = tdtypes.NewRPCErrorResponse(req.ID, 100, err.Error(), respData)
							p.nodeStore.AddCulprits(culprits)
						} else {
							errorRes = tdtypes.NewRPCErrorResponse(req.ID, 201, "sign failed", err.Error())
						}
						er := p.wsClient.SendMsg(errorRes)
						if er != nil {
							logger.Err(er).Msg("failed to send msg to tss manager")
						}
						continue
					}
					signResponse = tsscommon.SignResponse{
						Signature: data,
					}
					bol := p.CacheSign(hashStr, data)
					logger.Info().Msgf("cache roll back sign byte behavior %s ", bol)
				}

				RpcResponse := tdtypes.NewRPCSuccessResponse(req.ID, signResponse)
				err = p.wsClient.SendMsg(RpcResponse)
				if err != nil {
					logger.Err(err).Msg("failed to sendMsg to bridge ")
				} else {
					logger.Info().Msg("send roll back sign response successfully")
				}
			}
		}
	}()
}

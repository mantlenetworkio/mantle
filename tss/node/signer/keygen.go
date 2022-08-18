package signer

import (
	"encoding/json"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/common"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/keygen"
	tsstypes "github.com/bitdao-io/bitnetwork/tss/types"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func (p *Processor) Keygen() {
	defer p.wg.Done()
	logger := p.logger.With().Str("step", "keygen").Logger()

	logger.Info().Msg("start to keygen ")

	go func() {
		defer func() {
			logger.Info().Msg("exit keygen process")
		}()
		for {
			select {
			case <-p.stopChan:
				return
			case req := <-p.keygenRequestChan:
				var resId = req.ID.(tdtypes.JSONRPCStringID).String()
				logger.Info().Msgf("dealing resId (%s) ", resId)

				var keyR tsstypes.KeygenRequest
				if err := json.Unmarshal(req.Params, &keyR); err != nil {
					logger.Error().Msg("failed to unmarshal ask request")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}

				var keygenReq = keygen.Request{
					Keys:      keyR.Nodes,
					ThresHold: keyR.Threshold,
				}
				resp, err := p.tssServer.Keygen(keygenReq)

				if err != nil {
					logger.Err(err).Msg("failed to keygen !")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 202, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
				} else {
					if resp.Status == common.Success {
						keygenResponse := tsstypes.KeygenResponse{
							PoolPubKey: resp.PubKey,
						}
						RpcResponse := tdtypes.NewRPCSuccessResponse(tdtypes.JSONRPCStringID(resId), keygenResponse)
						p.wsClient.SendMsg(RpcResponse)
					} else {
						RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 202, "failed", resp.FailReason)
						p.wsClient.SendMsg(RpcResponse)
					}
				}

			}
		}
	}()
}

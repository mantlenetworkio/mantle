package signer

import (
	"encoding/json"
	"github.com/mantlenetworkio/mantle/tss/common"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func (p *Processor) VerifySlash() {
	defer p.wg.Done()
	logger := p.logger.With().Str("step", "verify slash event").Logger()
	logger.Info().Msg("start to verify slash events ")

	go func() {
		defer func() {
			logger.Info().Msg("exit verify slash event process")
		}()
		for {
			select {
			case <-p.stopChan:
				return
			case req := <-p.askSlashChan:
				var askRequest common.SlashRequest
				var RpcResponse tdtypes.RPCResponse
				var resId = req.ID
				if err := json.Unmarshal(req.Params, &askRequest); err != nil {
					logger.Error().Msg("failed to unmarshal ask request")
					RpcResponse = tdtypes.NewRPCErrorResponse(resId, 201, "failed to unmarshal", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}

				var re bool
				if askRequest.SignType == common.SlashTypeCulprit {
					culprits := p.nodeStore.GetCulprits()
					if len(culprits) > 0 {
						for _, v := range culprits {
							if v == askRequest.Address.String() {
								re = true
								break
							}
						}
					}
				} else if askRequest.SignType == common.SlashTypeLiveness {
					result := p.nodeStore.GetNodeMissedBatchBitArray(askRequest.Address, askRequest.BatchIndex)
					re = result
				}
				if re {
					p.UpdateWaitSignSlashMsgs(askRequest)
				}
				askResponse := common.AskResponse{
					Result: re,
				}
				RpcResponse = tdtypes.NewRPCSuccessResponse(resId, askResponse)
				p.wsClient.SendMsg(RpcResponse)
			}
		}
	}()

}

func (p *Processor) UpdateWaitSignSlashMsgs(msg common.SlashRequest) {
	p.waitSignSlashLock.Lock()
	defer p.waitSignSlashLock.Unlock()
	p.waitSignSlashMsgs[msg.Address.String()][msg.BatchIndex] = msg
}

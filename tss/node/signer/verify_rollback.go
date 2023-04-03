package signer

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	bindings "github.com/mantlenetworkio/mantle/tss/bindings/fp"
	"github.com/mantlenetworkio/mantle/tss/common"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func (p *Processor) VerifyRollBack() {
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
			case req := <-p.askRollBackChan:
				var RpcResponse tdtypes.RPCResponse

				var resId = req.ID

				ret, err := p.doVerify(req)

				if err != nil {
					logger.Error().Msgf("failed to do verify for rollback, %s", err.Error())
					RpcResponse = tdtypes.NewRPCErrorResponse(req.ID, 201, "failed to do verify for rollback ", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}
				askResponse := common.AskResponse{
					Result: ret,
				}
				RpcResponse = tdtypes.NewRPCSuccessResponse(resId, askResponse)
				p.wsClient.SendMsg(RpcResponse)
			}
		}
	}()

}

func (p *Processor) doVerify(req tdtypes.RPCRequest) (bool, error) {
	var askRequest common.SignStateRequest
	if err := json.Unmarshal(req.Params, &askRequest); err != nil {
		return false, err
	}

	challengeAddr := ethcommon.HexToAddress(askRequest.Challenge)
	challengeContract, err := bindings.NewChallenge(challengeAddr, p.l1Client)
	if err != nil {
		return false, err
	}
	winner, err := challengeContract.Winner(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		return false, err
	}
	defender, err := challengeContract.Defender(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		return false, err
	}
	rollback, err := challengeContract.Rollback(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		return false, err
	}

	if winner.String() == defender.String() && !rollback {
		return true, nil
	} else {
		return false, nil
	}

}

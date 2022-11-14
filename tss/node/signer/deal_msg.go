package signer

import (
	"errors"
	"github.com/mantlenetworkio/mantle/tss/common"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func (p *Processor) ProcessMessage() {
	logger := p.logger.With().Str("step", "process websocket message").Logger()
	defer p.wg.Done()
	reqChan := make(chan tmtypes.RPCRequest)
	stopChan := make(chan struct{})
	if err := p.wsClient.RegisterResChannel(reqChan, stopChan); err != nil {
		logger.Err(err).Msg("failed to register request channel with websocket client")
		return
	}

	go func() {
		defer func() {
			close(stopChan)
		}()
		for {
			select {
			case rpcReq := <-reqChan:
				reqId := rpcReq.ID.(tdtypes.JSONRPCStringID).String()
				logger.Info().Str("reqId", reqId).Msgf("receive request method : %s", rpcReq.Method)
				if rpcReq.Method == common.AskStateBatch.String() {
					if err := p.writeChan(p.askRequestChan, rpcReq); err != nil {
						logger.Err(err).Msg("failed to write msg to ask channel,channel blocked ")
					}
				} else if rpcReq.Method == common.SignStateBatch.String() {
					if err := p.writeChan(p.signRequestChan, rpcReq); err != nil {
						logger.Err(err).Msg("failed to write msg to sign channel,channel blocked ")
					}
				} else if rpcReq.Method == "keygen" {
					if err := p.writeChan(p.keygenRequestChan, rpcReq); err != nil {
						logger.Err(err).Msg("failed to write msg to keygen channel,channel blocked")
					}
				} else if rpcReq.Method == common.AskSlash.String() {
					if err := p.writeChan(p.askSlashChan, rpcReq); err != nil {
						logger.Err(err).Msg("failed to write msg to ask slash channel,channel blocked")
					}
				} else if rpcReq.Method == common.SignSlash.String() {
					if err := p.writeChan(p.signSlashChan, rpcReq); err != nil {
						logger.Err(err).Msg("failed to write msg to sign slash channel,channel blocked")
					}
				} else {
					logger.Error().Msgf("unknown rpc request method : %s ", rpcReq.Method)
				}
			}

		}
	}()
}

func (p *Processor) writeChan(cache chan tdtypes.RPCRequest, msg tdtypes.RPCRequest) error {
	select {
	case cache <- msg:
		if msg.Method == common.AskStateBatch.String() {
			p.metrics.AskChannelCount.Set(float64(len(cache)))
		} else if msg.Method == common.SignStateBatch.String() {
			p.metrics.SignChannelCount.Set(float64(len(cache)))
		}
		return nil
	default:
		return errors.New(msg.Method + " channel blocked,can not write!")
	}
}

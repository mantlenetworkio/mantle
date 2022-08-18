package signer

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/bitdao-io/bitnetwork/l2geth/crypto"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/common"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/keysign"
	tsstypes "github.com/bitdao-io/bitnetwork/tss/types"
	"github.com/rs/zerolog"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"sync"
)

func (p *Processor) Sign() {
	defer p.wg.Done()
	logger := p.logger.With().Str("step", "sign Message").Logger()

	logger.Info().Msg("start to sign message ")

	go func() {
		defer func() {
			logger.Info().Msg("exit sign process")
		}()
		for {
			select {
			case <-p.stopChan:
				return
			case req := <-p.signRequestChan:
				var resId = req.ID.(tdtypes.JSONRPCStringID).String()
				logger.Info().Msgf("dealing resId (%s) ", resId)

				var batchSignRequest tsstypes.BatchSignRequest
				if err := json.Unmarshal(req.Params, &batchSignRequest); err != nil {
					logger.Error().Msg("failed to unmarshal ask request")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}
				v, ok := p.signRequests[resId]
				if !ok {
					p.signRequests[resId] = batchSignRequest.Timestamp
				} else {
					if v < batchSignRequest.Timestamp {
						p.signRequests[resId] = batchSignRequest.Timestamp
						ch, ok := p.signMsgQuitChan[resId]
						if ok {
							ch <- struct{}{}
						}
					}
				}
				err := checkMessages(batchSignRequest.Signs, p.waitSignMsgs)
				if err != nil {
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())

					p.wsClient.SendMsg(RpcResponse)
					logger.Err(err).Msg("check event failed")
					continue
				}

				wg := &sync.WaitGroup{}
				wg.Add(len(batchSignRequest.Signs))

				quit := make(chan struct{})
				p.signMsgQuitChan[resId] = quit
				for _, sign := range batchSignRequest.Signs {
					go p.distribute(resId, sign, quit, wg, batchSignRequest.PoolPubKey, logger)
				}
				wg.Wait()
				close(p.signMsgQuitChan[resId])
				delete(p.signMsgQuitChan, resId)
			}
		}
	}()
}

func (p *Processor) distribute(reqId string, sign tsstypes.SignRequest, quit <-chan struct{}, wg *sync.WaitGroup, poolPubKey string, logger zerolog.Logger) {

	defer wg.Done()
	select {
	case <-quit:
		return
	default:
		data, err := p.handleSign(sign, poolPubKey, logger)

		if err != nil {
			logger.Error().Msgf(" %s sign failed ", sign.UniqueId)
			er := p.wsClient.SendMsg(tdtypes.NewRPCErrorResponse(tdtypes.JSONRPCStringID(reqId), 201, "failed", err.Error()))
			if er != nil {
				logger.Err(er).Msg("failed to send msg to tss manager")
			} else {
				p.removeWaitEvent(sign.UniqueId)
			}
			return
		}

		signResponse := tsstypes.SignResponse{
			UniqueId:  sign.UniqueId,
			Signature: data,
		}
		RpcResponse := tdtypes.NewRPCSuccessResponse(tdtypes.JSONRPCStringID(reqId), signResponse)
		err = p.wsClient.SendMsg(RpcResponse)
		if err != nil {
			logger.Err(err).Msg("failed to sendMsg to bridge ")
		} else {
			p.removeWaitEvent(sign.UniqueId)
		}
	}

}

func (p *Processor) handleSign(sign tsstypes.SignRequest, poolPubKey string, logger zerolog.Logger) (tsstypes.SignatureData, error) {

	logger.Info().Msgf(" dealing sign (%s)", sign.UniqueId)
	msg, ok := p.waitSignMsgs[sign.UniqueId]

	if !ok {
		logger.Error().Msgf("msg (%s) doesn't verify ", sign.UniqueId)
		return tsstypes.SignatureData{}, errors.New("msg doesn't verify " + sign.UniqueId)
	}

	rawBytes := make([]byte, 0)
	for _, sr := range msg.StateRoots {
		rawBytes = append(rawBytes, sr[:]...)
	}
	rawBytes = append(rawBytes, msg.OffsetStartsAtIndex.Bytes()...)
	digestBz := crypto.Keccak256Hash(rawBytes).Bytes()

	signedData, culpritNodes, err := p.sign(digestBz, sign.Nodes, poolPubKey, logger)
	if err != nil {
		if len(culpritNodes) > 0 {
			logger.Err(err).Msgf(" sign failed with culpritNodes %s ", culpritNodes)
		}
		return tsstypes.SignatureData{}, err
	}
	return signedData, nil
}

func (p *Processor) sign(digestBz []byte, signerPubKeys []string, poolPubKey string, logger zerolog.Logger) (signatureData tsstypes.SignatureData, culpritNodes []string, err error) {

	logger.Info().Str("message", hex.EncodeToString(digestBz)).Msg("got message to be signed")
	keysignReq := keysign.NewRequest(poolPubKey, digestBz, signerPubKeys)
	keysignRes, err := p.tssServer.KeySign(keysignReq)
	if err != nil {
		logger.Err(err).Msg("fail to generate signature ")
		return signatureData, nil, err
	}
	if keysignRes.Status == common.Success {
		signatureData = tsstypes.SignatureData{
			SignatureRecovery: keysignRes.SignatureData.SignatureRecovery,
			R:                 keysignRes.SignatureData.R,
			S:                 keysignRes.SignatureData.S,
			M:                 keysignRes.SignatureData.M,
		}

		return signatureData, nil, nil
	} else {
		return signatureData, keysignRes.Culprits, errors.New(keysignRes.FailReason)
	}
}

func checkMessages(signs []tsstypes.SignRequest, waitSignMsgs map[string]tsstypes.AskStateRequest) error {
	for _, sign := range signs {
		_, ok := waitSignMsgs[sign.UniqueId]
		if !ok {
			return errors.New("event sign request has the event which unverified")
		}
	}
	return nil
}

func (p *Processor) removeWaitEvent(key string) {
	p.waitSignLock.Lock()
	defer p.waitSignLock.Unlock()
	delete(p.waitSignMsgs, key)
}

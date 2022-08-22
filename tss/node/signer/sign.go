package signer

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	l2common "github.com/bitdao-io/bitnetwork/l2geth/common"
	"github.com/bitdao-io/bitnetwork/l2geth/crypto"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/common"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/keysign"
	tsstypes "github.com/bitdao-io/bitnetwork/tss/types"
	"github.com/rs/zerolog"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
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

				var nodeSignRequest tsstypes.NodeSignStateRequest
				if err := json.Unmarshal(req.Params, &nodeSignRequest); err != nil {
					logger.Error().Msg("failed to unmarshal ask request")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}

				err, hash := checkMessages(nodeSignRequest.StateBatch, p.waitSignMsgs)
				if err != nil {
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())

					p.wsClient.SendMsg(RpcResponse)
					logger.Err(err).Msg("check event failed")
					continue
				}

				data, err := p.handleSign(nodeSignRequest, hash, logger)

				if err != nil {
					logger.Error().Msgf(" %s sign failed ", hash.String())
					er := p.wsClient.SendMsg(tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error()))
					if er != nil {
						logger.Err(er).Msg("failed to send msg to tss manager")
					} else {
						p.removeWaitEvent(hash.String())
					}
					return
				}

				signResponse := tsstypes.SignResponse{
					Signature: data,
				}
				RpcResponse := tdtypes.NewRPCSuccessResponse(req.ID, signResponse)
				err = p.wsClient.SendMsg(RpcResponse)
				if err != nil {
					logger.Err(err).Msg("failed to sendMsg to bridge ")
				} else {
					p.removeWaitEvent(hash.String())
				}
			}
		}
	}()
}

func (p *Processor) handleSign(sign tsstypes.NodeSignStateRequest, hash l2common.Hash, logger zerolog.Logger) (tsstypes.SignatureData, error) {

	logger.Info().Msgf(" timestamp (%s) ,dealing sign hex (%s)", sign.Timestamp, hash.String())

	signedData, culpritNodes, err := p.sign(hash.Bytes(), sign.Nodes, sign.ClusterPublicKey, logger)
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

func checkMessages(sign tsstypes.SignStateRequest, waitSignMsgs map[string]tsstypes.SignStateRequest) (error, l2common.Hash) {
	hash := signMsgToHash(sign)
	_, ok := waitSignMsgs[hash.String()]
	if !ok {
		return errors.New("event sign request has the event which unverified"), hash
	}

	return nil, hash
}

func signMsgToHash(msg tsstypes.SignStateRequest) l2common.Hash {
	rawBytes := make([]byte, 0)
	for _, sr := range msg.StateRoots {
		rawBytes = append(rawBytes, sr[:]...)
	}
	rawBytes = append(rawBytes, msg.OffsetStartsAtIndex.Bytes()...)
	return crypto.Keccak256Hash(rawBytes)
}

func (p *Processor) removeWaitEvent(key string) {
	p.waitSignLock.Lock()
	defer p.waitSignLock.Unlock()
	delete(p.waitSignMsgs, key)
}

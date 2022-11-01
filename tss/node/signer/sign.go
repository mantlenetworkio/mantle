package signer

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/influxdata/influxdb/pkg/slices"
	"github.com/mantlenetworkio/mantle/l2geth/common/hexutil"
	tsscommon "github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/index"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/common"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/keysign"
	"github.com/mantlenetworkio/mantle/tss/slash"
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

				var nodeSignRequest tsscommon.NodeSignRequest
				rawMsg := json.RawMessage{}
				nodeSignRequest.RequestBody = &rawMsg

				if err := json.Unmarshal(req.Params, &nodeSignRequest); err != nil {
					logger.Error().Msg("failed to unmarshal ask request")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}
				var requestBody tsscommon.SignStateRequest
				if err := json.Unmarshal(rawMsg, &requestBody); err != nil {
					logger.Error().Msg("failed to umarshal ask's params request body")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}
				nodeSignRequest.RequestBody = requestBody

				var data []byte
				err, hash, signByte := p.checkMessages(requestBody)
				hashStr := hexutil.Encode(hash)

				if err != nil {
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())

					p.wsClient.SendMsg(RpcResponse)
					logger.Err(err).Msg("check event failed")
					continue
				}

				//cache can not find the sign result by hashStr,we need to handle sign request.
				if signByte == nil {
					signData, culprits, err := p.handleSign(nodeSignRequest, hash, logger)
					if err != nil {
						logger.Error().Msgf(" %s sign failed ", hashStr)
						var errorRes tdtypes.RPCResponse
						if len(culprits) > 0 {
							respData := strings.Join(culprits, ",")
							errorRes = tdtypes.NewRPCErrorResponse(req.ID, tsscommon.CulpritErrorCode, err.Error(), respData)
							p.nodeStore.AddCulprits(culprits)

							//store slash info
							for _, culprit := range culprits {
								addr, err := tsscommon.NodeToAddress(culprit)
								if err != nil {
									logger.Error().Msgf("failed to convert node to address %s", culprit)
								}
								p.nodeStore.SetSlashingInfo(slash.SlashingInfo{
									Address:    addr,
									ElectionId: requestBody.ElectionId,
									BatchIndex: math.MaxUint64, // not real, just for identifying the specific slashing info.
									SlashType:  tsscommon.SlashTypeCulprit,
								})
							}

						} else {
							errorRes = tdtypes.NewRPCErrorResponse(req.ID, 201, "sign failed", err.Error())
						}
						er := p.wsClient.SendMsg(errorRes)
						if er != nil {
							logger.Err(er).Msg("failed to send msg to tss manager")
						} else {
							p.removeWaitEvent(hashStr)
						}

						continue
					}
					bol := p.CacheSign(hashStr, signData)
					logger.Info().Msgf("cache sign byte behavior %s ", bol)
					data = signData
				} else {
					data = signByte
				}

				signResponse := tsscommon.SignResponse{
					Signature: data,
				}
				RpcResponse := tdtypes.NewRPCSuccessResponse(req.ID, signResponse)
				logger.Info().Msg("start to send response to manager ")

				err = p.wsClient.SendMsg(RpcResponse)
				if err != nil {
					logger.Err(err).Msg("failed to sendMsg to tss manager ")
				} else {
					logger.Info().Msg("send sign response to manager successfully")
					err := p.storeStateBatch(requestBody.ElectionId, requestBody.StateRoots, nodeSignRequest.Nodes, nodeSignRequest.ClusterPublicKey)
					if err != nil {
						logger.Err(err).Msg("failed to store StateBatch to level db")
					}
					p.removeWaitEvent(hashStr)
				}
			}
		}
	}()
}

func (p *Processor) handleSign(sign tsscommon.NodeSignRequest, hashTx []byte, logger zerolog.Logger) ([]byte, []string, error) {

	logger.Info().Msgf(" timestamp (%d) ,dealing sign hex (%s)", sign.Timestamp, hexutil.Encode(hashTx))

	signedData, culpritNodes, err := p.sign(hashTx, sign.Nodes, sign.ClusterPublicKey, logger)
	if err != nil {
		if len(culpritNodes) > 0 {
			logger.Err(err).Msgf(" sign failed with culpritNodes %s ", culpritNodes)
		}
		return nil, culpritNodes, err
	}
	signatureBytes := getSignatureBytes(&signedData)
	return signatureBytes, nil, nil
}

func (p *Processor) sign(digestBz []byte, signerPubKeys []string, poolPubKey string, logger zerolog.Logger) (signatureData tsscommon.SignatureData, culpritNodes []string, err error) {

	logger.Info().Str("message", hex.EncodeToString(digestBz)).Msg("got message to be signed")
	keysignReq := keysign.NewRequest(poolPubKey, digestBz, signerPubKeys)
	keysignRes, err := p.tssServer.KeySign(keysignReq)
	if err != nil {
		logger.Err(err).Msg("fail to generate signature ")
		return signatureData, nil, err
	}
	if keysignRes.Status == common.Success {
		signatureData = tsscommon.SignatureData{
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

func (p *Processor) checkMessages(sign tsscommon.SignStateRequest) (err error, hashByte, signByte []byte) {
	hashByte, err = signMsgToHash(sign)
	if err != nil {
		return err, hashByte, nil
	}
	hashStr := hexutil.Encode(hashByte)

	signByte, ok := p.cacheSign.Get(hashStr)
	if ok {
		return nil, hashByte, signByte
	}
	_, ok = p.waitSignMsgs[hashStr]
	if !ok {
		return errors.New("sign request has the unverified state batch"), nil, nil
	}
	return nil, hashByte, nil
}

func signMsgToHash(msg tsscommon.SignStateRequest) ([]byte, error) {
	offsetStartsAtIndex, _ := new(big.Int).SetString(msg.OffsetStartsAtIndex, 10)
	return tsscommon.StateBatchHash(msg.StateRoots, offsetStartsAtIndex)
}

func (p *Processor) removeWaitEvent(key string) {
	p.waitSignLock.Lock()
	defer p.waitSignLock.Unlock()
	delete(p.waitSignMsgs, key)
}

func getSignatureBytes(sig *tsscommon.SignatureData) []byte {
	R := new(big.Int).SetBytes(sig.R)
	S := new(big.Int).SetBytes(sig.S)
	N := btcec.S256().N
	halfOrder := new(big.Int).Rsh(N, 1)
	if S.Cmp(halfOrder) == 1 {
		S.Sub(N, S)
	}
	rBytes := R.Bytes()
	sBytes := S.Bytes()
	cBytes := sig.SignatureRecovery

	sigBytes := make([]byte, 65)
	copy(sigBytes[32-len(rBytes):32], rBytes)
	copy(sigBytes[64-len(sBytes):64], sBytes)
	copy(sigBytes[64:65], cBytes)
	return sigBytes
}

func (p *Processor) storeStateBatch(electionId uint64, stateBatch [][32]byte, workingNodes []string, poolPubkey string) error {
	batchRoot, err := tsscommon.GetMerkleRoot(stateBatch)
	if err != nil {
		return err
	}

	paricipants, err := p.tssServer.GetParticipants(poolPubkey)
	if err != nil {
		return err
	}
	absentNodes := make([]string, 0)
	for _, n := range paricipants {
		if !slices.ExistsIgnoreCase(workingNodes, n) {
			absentNodes = append(absentNodes, n)
		}

	}

	sbi := index.StateBatchInfo{
		BatchRoot:    batchRoot,
		ElectionId:   electionId,
		AbsentNodes:  absentNodes,
		WorkingNodes: workingNodes,
	}
	if err = p.nodeStore.SetStateBatch(sbi); err != nil {
		return err
	}
	return nil
}

func (p *Processor) CacheSign(key string, value []byte) bool {
	p.cacheSignLock.Lock()
	defer p.cacheSignLock.Unlock()
	return p.cacheSign.Set(key, value)
}

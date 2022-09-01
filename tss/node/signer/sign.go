package signer

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	l2common "github.com/bitdao-io/bitnetwork/l2geth/common"
	"github.com/bitdao-io/bitnetwork/l2geth/common/hexutil"
	"github.com/bitdao-io/bitnetwork/l2geth/crypto"
	tsscommon "github.com/bitdao-io/bitnetwork/tss/common"
	"github.com/bitdao-io/bitnetwork/tss/index"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/common"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/keysign"
	"github.com/btcsuite/btcd/btcec"
	"github.com/influxdata/influxdb/pkg/slices"
	"github.com/rs/zerolog"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"math/big"
	"strings"
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

				err, hash := checkMessages(requestBody, p.waitSignMsgs)
				if err != nil {
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())

					p.wsClient.SendMsg(RpcResponse)
					logger.Err(err).Msg("check event failed")
					continue
				}

				data, culprits, err := p.handleSign(nodeSignRequest, hash.Bytes(), logger)

				if err != nil {
					logger.Error().Msgf(" %s sign failed ", hash.String())
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
					} else {
						p.removeWaitEvent(hash.String())
					}

					continue
				}

				signResponse := tsscommon.SignResponse{
					Signature: data,
				}
				RpcResponse := tdtypes.NewRPCSuccessResponse(req.ID, signResponse)
				err = p.wsClient.SendMsg(RpcResponse)
				if err != nil {
					logger.Err(err).Msg("failed to sendMsg to bridge ")
				} else {
					err := p.storeStateBatch(requestBody.ElectionId, requestBody.StateRoots, nodeSignRequest.Nodes, nodeSignRequest.ClusterPublicKey)
					if err != nil {
						logger.Err(err).Msg("failed to store StateBatch to level db")
					}
					p.removeWaitEvent(hash.String())
				}
			}
		}
	}()
}

func (p *Processor) handleSign(sign tsscommon.NodeSignRequest, hashTx []byte, logger zerolog.Logger) ([]byte, []string, error) {

	logger.Info().Msgf(" timestamp (%s) ,dealing sign hex (%s)", sign.Timestamp, hexutil.Encode(hashTx))

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

func checkMessages(sign tsscommon.SignStateRequest, waitSignMsgs map[string]tsscommon.SignStateRequest) (error, l2common.Hash) {
	hash := signMsgToHash(sign)
	_, ok := waitSignMsgs[hash.String()]
	if !ok {
		return errors.New("event sign request has the event which unverified"), hash
	}

	return nil, hash
}

func signMsgToHash(msg tsscommon.SignStateRequest) l2common.Hash {
	rawBytes := make([]byte, 0)
	for _, sr := range msg.StateRoots {
		rawBytes = append(rawBytes, sr[:]...)
	}
	offsetStartsAtIndex, _ := new(big.Int).SetString(msg.OffsetStartsAtIndex, 10)
	rawBytes = append(rawBytes, offsetStartsAtIndex.Bytes()...)
	return crypto.Keccak256Hash(rawBytes)
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

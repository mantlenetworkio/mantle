package manager

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/bitdao-io/bitnetwork/l2geth/crypto"
	"github.com/bitdao-io/bitnetwork/l2geth/log"
	"github.com/bitdao-io/bitnetwork/tss/manager/types"
	tss "github.com/bitdao-io/bitnetwork/tss/types"
	"github.com/bitdao-io/bitnetwork/tss/ws/server"
	"github.com/btcsuite/btcd/btcec"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func (m Manager) sign(ctx types.Context, request tss.SignStateRequest) ([]byte, error) {
	respChan := make(chan server.ResponseMsg)
	stopChan := make(chan struct{})

	if err := m.wsServer.RegisterResChannel(ctx.RequestId(), respChan, stopChan); err != nil {
		log.Error("failed to register response channel at signing step", err)
		return nil, err
	}

	errSendChan := make(chan struct{})
	var validSignatureBz []byte
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		cctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer func() {
			log.Info("exit signing process")
			cancel()
			close(stopChan)
			wg.Done()
		}()
		for {
			select {
			case <-errSendChan:
				return
			case resp := <-respChan:
				log.Info(fmt.Sprintf("signed response: %s", resp.RpcResponse.String()), "node", resp.SourceNode)
				if resp.RpcResponse.Error == nil {
					var signResponse tss.SignResponse
					if err := tmjson.Unmarshal(resp.RpcResponse.Result, &signResponse); err != nil {
						log.Error("failed to unmarshal sign response", err)
						continue
					}

					rawBytes := make([]byte, 0)
					for _, sr := range request.StateRoots {
						rawBytes = append(rawBytes, sr[:]...)
					}
					rawBytes = append(rawBytes, request.OffsetStartsAtIndex.Bytes()...)
					digestBz := crypto.Keccak256Hash(rawBytes).Bytes()
					if bytes.Compare(digestBz, signResponse.Signature.M) != 0 {
						log.Error("mismatched hash", "expected hash", hex.EncodeToString(digestBz), "actual", signResponse.Signature.M)
						continue
					}
					poolPubKeyBz, _ := hex.DecodeString(ctx.TssInfos().ClusterPubKey)
					signatureBz, err := getSignature(&signResponse.Signature)
					if err != nil {
						log.Error("failed to parse signature", err)
						continue
					}
					if !crypto.VerifySignature(poolPubKeyBz, digestBz, signatureBz[:64]) {
						log.Error("illegal signature")
						continue
					}
					validSignatureBz = signatureBz
					return
				}
			case <-cctx.Done():
				log.Warn("wait for signature timeout")
				return
			default:
			}
		}
	}()

	m.sendToNodes(ctx, request, errSendChan)
	wg.Wait()

	var err error
	if validSignatureBz == nil {
		err = errors.New("failed to generate signature")
	}
	return validSignatureBz, err
}

func (m Manager) sendToNodes(ctx types.Context, request tss.SignStateRequest, errSendChan chan struct{}) {
	nodes := ctx.Approvers()
	nodeRequest := tss.NodeSignStateRequest{
		ClusterPublicKey: ctx.TssInfos().ClusterPubKey,
		Timestamp:        time.Now().UnixMilli(),
		Nodes:            ctx.Approvers(),
		StateBatch:       request,
	}
	requestBz, err := json.Marshal(nodeRequest)
	if err != nil {
		log.Error("failed to json marshal node request", err)
		errSendChan <- struct{}{}
		return
	}

	rpcRequest := tmtypes.NewRPCRequest(tmtypes.JSONRPCStringID(ctx.RequestId()), "signState", requestBz)
	for _, node := range nodes {
		go func(node string, request tmtypes.RPCRequest) {
			if err := m.wsServer.SendMsg(
				server.RequestMsg{
					RpcRequest: request,
					TargetNode: node,
				}); err != nil {
				log.Error("failed to send sign request to nodes", err)
				errSendChan <- struct{}{}
				return
			}
		}(node, rpcRequest)
	}
}

func getSignature(sig *tss.SignatureData) ([]byte, error) {
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
	return sigBytes, nil
}

package manager

import (
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
	tss "github.com/bitdao-io/bitnetwork/tss/common"
	"github.com/bitdao-io/bitnetwork/tss/manager/types"
	"github.com/bitdao-io/bitnetwork/tss/ws/server"
	"github.com/btcsuite/btcd/btcec"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func (m Manager) sign(ctx types.Context, request interface{}, digestBz []byte, method tss.Method) (tss.SignResponse, error) {
	respChan := make(chan server.ResponseMsg)
	stopChan := make(chan struct{})

	if err := m.wsServer.RegisterResChannel(ctx.RequestId(), respChan, stopChan); err != nil {
		log.Error("failed to register response channel at signing step", err)
		return tss.SignResponse{}, err
	}

	errSendChan := make(chan struct{})
	var validSignatureResponse tss.SignResponse
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

					poolPubKeyBz, _ := hex.DecodeString(ctx.TssInfos().ClusterPubKey)
					if !crypto.VerifySignature(poolPubKeyBz, digestBz, signResponse.Signature[:64]) {
						log.Error("illegal signature")
						continue
					}
					validSignatureResponse = signResponse
					return
				}
			case <-cctx.Done():
				log.Warn("wait for signature timeout")
				return
			default:
			}
		}
	}()

	m.sendToNodes(ctx, request, method, errSendChan)
	wg.Wait()

	var err error
	if validSignatureResponse.Signature == nil {
		err = errors.New("failed to generate signature")
	}
	return validSignatureResponse, err
}

func (m Manager) sendToNodes(ctx types.Context, request interface{}, method tss.Method, errSendChan chan struct{}) {
	nodes := ctx.Approvers()
	nodeRequest := tss.NodeSignRequest{
		ClusterPublicKey: ctx.TssInfos().ClusterPubKey,
		Timestamp:        time.Now().UnixMilli(),
		Nodes:            ctx.Approvers(),
		RequestBody:      request,
	}
	requestBz, err := json.Marshal(nodeRequest)
	if err != nil {
		log.Error("failed to json marshal node request", err)
		errSendChan <- struct{}{}
		return
	}

	rpcRequest := tmtypes.NewRPCRequest(tmtypes.JSONRPCStringID(ctx.RequestId()), method.String(), requestBz)
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

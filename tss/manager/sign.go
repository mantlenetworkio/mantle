package manager

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/influxdata/influxdb/pkg/slices"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/mantlenetworkio/mantle/l2geth/log"
	tss "github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/manager/types"
	"github.com/mantlenetworkio/mantle/tss/ws/server"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

type Counter struct {
	count map[string]int
}

func (c *Counter) increment(node string) {
	if c.count == nil {
		c.count = make(map[string]int, 0)
	}
	num := c.count[node]
	c.count[node] = num + 1
}

func (c *Counter) satisfied(minNumber int) []string {
	var ret []string
	for n, ct := range c.count {
		if ct >= minNumber {
			ret = append(ret, n)
		}
	}
	return ret
}

func (m Manager) sign(ctx types.Context, request interface{}, digestBz []byte, method tss.Method) (tss.SignResponse, []string, error) {
	respChan := make(chan server.ResponseMsg)
	stopChan := make(chan struct{})

	if err := m.wsServer.RegisterResChannel(ctx.RequestId(), respChan, stopChan); err != nil {
		log.Error("failed to register response channel at signing step", err)
		return tss.SignResponse{}, nil, err
	}
	log.Info("Registered ResChannel with requestID", "requestID", ctx.RequestId())

	errSendChan := make(chan struct{})
	responseNodes := make(map[string]struct{})
	counter := &Counter{}
	var validSignResponse *tss.SignResponse
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		cctx, cancel := context.WithTimeout(context.Background(), m.signTimeout)
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
				if !slices.ExistsIgnoreCase(ctx.Approvers(), resp.SourceNode) { // ignore the message which the sender should not be involved in approver set
					continue
				}
				func() {
					defer func() {
						responseNodes[resp.SourceNode] = struct{}{}
					}()
					if resp.RpcResponse.Error == nil {
						var signResponse tss.SignResponse
						if err := tmjson.Unmarshal(resp.RpcResponse.Result, &signResponse); err != nil {
							log.Error("failed to unmarshal sign response", err)
							return
						}

						poolPubKeyBz, _ := hex.DecodeString(ctx.TssInfos().ClusterPubKey)
						if !crypto.VerifySignature(poolPubKeyBz, digestBz, signResponse.Signature[:64]) {
							log.Error("illegal signature")
							return
						}

						if method != tss.SignSlash { // if it is not signSlash, then exit when receiving the first valid response
							validSignResponse = &signResponse
							return
						}

						// if signing slashing, we chose a better gas price as the valid one
						if validSignResponse == nil {
							slashTxGasPrice, succ := new(big.Int).SetString(signResponse.SlashTxGasPrice, 10)
							if !succ {
								log.Error("wrong format of slashTxGasPrice")
								return
							}
							signResponse.SlashTxGasPriceBigInt = slashTxGasPrice
							validSignResponse = &signResponse
						} else {
							// if current gas price > last node gas price, replace it
							slashTxGasPrice, succ := new(big.Int).SetString(signResponse.SlashTxGasPrice, 10)
							if !succ {
								log.Error("wrong format of slashTxGasPrice")
								return
							}
							if slashTxGasPrice.Cmp(validSignResponse.SlashTxGasPriceBigInt) > 0 {
								signResponse.SlashTxGasPriceBigInt = slashTxGasPrice
								validSignResponse = &signResponse
							}
						}
					} else if resp.RpcResponse.Error.Code == tss.CulpritErrorCode {
						_, ok := responseNodes[resp.SourceNode]
						if ok { // ignore if handled
							return
						}

						culpritData := resp.RpcResponse.Error.Data
						culprits := strings.Split(culpritData, ",")
						for _, culprit := range culprits {
							if slices.Exists(ctx.Approvers(), culprit) {
								counter.increment(culprit)
							}
						}
					}
				}()

			case <-cctx.Done():
				log.Warn("wait for signature timeout")
				return
			default:
				if len(responseNodes) == len(ctx.Approvers()) {
					log.Info("received all signing responses")
					return
				}
			}
		}
	}()

	m.sendToNodes(ctx, request, method, errSendChan)
	wg.Wait()

	var culprits []string
	if validSignResponse == nil {
		culprits = counter.satisfied(ctx.TssInfos().Threshold + 1)
		return tss.SignResponse{}, culprits, errors.New("failed to generate signature")
	}
	return *validSignResponse, culprits, nil
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

package manager

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/influxdata/influxdb/pkg/slices"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	tmjson "github.com/tendermint/tendermint/libs/json"

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

func (m *Manager) sign(ctx types.Context, request interface{}, digestBz []byte, method tss.Method) (tss.SignResponse, []string, error) {
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
					if resp.RpcResponse.Error != nil {
						if resp.RpcResponse.Error.Code == tss.CulpritErrorCode {
							_, ok := responseNodes[resp.SourceNode]
							if ok { // ignore if handled
								return
							}

							culpritData := resp.RpcResponse.Error.Data
							culprits := strings.Split(culpritData, ",")
							culprits = deDuplication(culprits)
							for _, culprit := range culprits {
								if slices.ExistsIgnoreCase(ctx.Approvers(), culprit) {
									counter.increment(culprit)
								}
							}
						} else {
							log.Error("Unrecognized error code",
								"err_code", resp.RpcResponse.Error.Code,
								"err_data", resp.RpcResponse.Error.Data,
								"err_message", resp.RpcResponse.Error.Message)
						}
						return
					} else {
						var signResponse tss.SignResponse
						if err := tmjson.Unmarshal(resp.RpcResponse.Result, &signResponse); err != nil {
							log.Error("failed to unmarshal sign response", err)
							return
						}

						poolPubKeyBz, _ := hex.DecodeString(ctx.TssInfos().ClusterPubKey)
						if len(signResponse.Signature) < 64 {
							log.Error(fmt.Sprintf("invalid signature, expected length is no less than 64, actual length is %d", len(signResponse.Signature)))
							return
						}
						if !crypto.VerifySignature(poolPubKeyBz, digestBz, signResponse.Signature[:64]) {
							log.Error("illegal signature")
							return
						}

						validSignResponse = &signResponse
						return
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

func (m *Manager) sendToNodes(ctx types.Context, request interface{}, method tss.Method, errSendChan chan struct{}) {
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

func deDuplication(nodes []string) []string {
	var r []string
	h := make(map[string]bool)
	for _, node := range nodes {
		if _, ok := h[node]; !ok {
			r = append(r, node)
			h[node] = true
		}
	}
	return r
}

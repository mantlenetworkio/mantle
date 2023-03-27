package manager

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/influxdata/influxdb/pkg/slices"
	tss "github.com/mantlenetworkio/mantle/mt-tss/common"
	"github.com/mantlenetworkio/mantle/mt-tss/manager/types"
	"github.com/mantlenetworkio/mantle/mt-tss/ws/server"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func (m Manager) observeElection() {

	queryTicker := time.NewTicker(m.taskInterval + 30*time.Second)
	for {
		log.Info("trying to handle new election...", "stopGenKey", m.stopGenKey)
		if !m.stopGenKey {
			func() {
				// check if new round election is held(inactive tss members)
				tssInfo, err := m.tssQueryService.QueryInactiveInfo()
				if err != nil {
					log.Error("failed to query inactive info", "err", err)
					return
				}
				log.Info("query inactive members", "numbers", len(tssInfo.TssMembers))

				//tssMembers, threshold, electionId := getInactiveMembers()
				if len(tssInfo.TssMembers) > 0 {
					culprits := m.store.GetCulprits()
					if len(culprits) > 0 {
						for _, tm := range tssInfo.TssMembers {
							if slices.Exists(culprits, tm) {
								log.Error("Reject to keygen, culprit is involved in the elected members", "culprit", tm)
								return
							}
						}
					}

					// the CPK has not been confirmed in the latest election
					// start to generate CPK
					cpkData, err := m.store.GetByElectionId(tssInfo.ElectionId)
					if err != nil {
						log.Error("failed to get cpk from storage", "err", err)
						return
					}

					if len(cpkData.Cpk) != 0 && time.Now().Sub(cpkData.CreationTime).Hours() < m.cpkConfirmTimeout.Hours() { // cpk is generated, but has not been confirmed yet
						return
					}
					cpk, err := m.generateKey(tssInfo.TssMembers, tssInfo.Threshold)
					if err != nil {
						log.Error("failed to generate key", "err", err)
						return
					}

					if err = m.store.Insert(types.CpkData{
						Cpk:          cpk,
						ElectionId:   tssInfo.ElectionId,
						CreationTime: time.Now(),
					}); err != nil {
						log.Error("failed to get cpk from storage", "err", err)
					}
				}
			}()
		}

		select {
		case <-m.stopChan:
			return
		case <-queryTicker.C:
		}
	}
}

func (m Manager) generateKey(tssMembers []string, threshold int) (string, error) {
	availableNodes := m.availableNodes(tssMembers)
	if len(availableNodes) < len(tssMembers) {
		return "", errors.New("not enough available nodes to generate CPK")
	}
	requestId := randomRequestId()
	respChan := make(chan server.ResponseMsg)
	stopChan := make(chan struct{})
	if err := m.wsServer.RegisterResChannel(requestId, respChan, stopChan); err != nil {
		log.Error("failed to register response channel", "err", err)
		return "", err
	}

	sendError := make(chan struct{})
	clusterPublicKeys := make(map[string]string, 0)
	var anyError error
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		cctx, cancel := context.WithTimeout(context.Background(), m.keygenTimeout)
		defer func() {
			log.Info("exit accept keygen response goroutine")
			cancel()
			close(stopChan)
			wg.Done()
		}()
		for {
			select {
			case <-sendError:
				anyError = errors.New("failed to send request to node")
				log.Error("failed to send request to node")
				return
			case <-cctx.Done():
				anyError = errors.New("wait nodes for keygen response timeout")
				log.Error("wait nodes for keygen response timeout")
				return
			case resp := <-respChan:
				log.Info("received keygen response", "response", resp.RpcResponse.String(), "node", resp.SourceNode)
				if resp.RpcResponse.Error != nil {
					anyError = errors.New(resp.RpcResponse.Error.Error())
					log.Error("returns error", "node", resp.SourceNode)
					return
				}
				var keygenResp tss.KeygenResponse
				if err := tmjson.Unmarshal(resp.RpcResponse.Result, &keygenResp); err != nil {
					anyError = err
					log.Error("failed to Unmarshal KeygenResponse", "err", err)
					return
				}
				clusterPublicKeys[resp.SourceNode] = keygenResp.ClusterPublicKey
			default:
				if len(clusterPublicKeys) == len(availableNodes) {
					return
				}
			}
		}
	}()

	m.callKeygen(availableNodes, threshold, requestId, sendError)
	wg.Wait()

	if anyError != nil {
		return "", anyError
	}

	// check if existing different CPKs
	var base string
	for _, cpk := range clusterPublicKeys {
		if len(base) == 0 {
			base = cpk
			continue
		}
		if cpk != base {
			return "", errors.New("found different CPKs generated from tss members")
		}
	}

	if len(clusterPublicKeys) != len(availableNodes) {
		return "", errors.New("timeout")
	}
	return base, nil
}

func (m Manager) callKeygen(availableNodes []string, threshold int, requestId string, sendError chan struct{}) {
	for _, node := range availableNodes {
		nodeRequest := tss.KeygenRequest{
			Nodes:     availableNodes,
			Threshold: threshold,
			Timestamp: time.Now().UnixMilli(),
		}
		requestBz, _ := json.Marshal(nodeRequest)
		go func(node string, requestBz []byte) {
			requestMsg := server.RequestMsg{
				TargetNode: node,
				RpcRequest: tmtypes.NewRPCRequest(tmtypes.JSONRPCStringID(requestId), "keygen", requestBz),
			}
			if err := m.wsServer.SendMsg(requestMsg); err != nil {
				sendError <- struct{}{}
			}
		}(node, requestBz)
	}
}

package manager

import (
	"errors"
	"testing"
	"time"

	tss "github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/ws/server"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func TestKeygen(t *testing.T) {
	var afterMsgSent afterMsgSendFunc = func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		askResp := tss.KeygenResponse{
			ClusterPublicKey: "abcd",
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, askResp)
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}
	var queryAliveNodes queryAliveNodesFunc = func() []string {
		return []string{"a", "b", "c", "d"}
	}
	manager, _ := setup(afterMsgSent, queryAliveNodes)
	cpk, err := manager.generateKey([]string{"a", "b", "c", "d"}, 3)
	require.NoError(t, err)
	require.EqualValues(t, "abcd", cpk)
}

func TestInConsistCPKGen(t *testing.T) {
	var afterMsgSent afterMsgSendFunc = func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		askResp := tss.KeygenResponse{
			ClusterPublicKey: "abcd",
		}
		if request.TargetNode == "c" {
			askResp = tss.KeygenResponse{
				ClusterPublicKey: "abc",
			}
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, askResp)
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}
	var queryAliveNodes queryAliveNodesFunc = func() []string {
		return []string{"a", "b", "c", "d"}
	}
	manager, _ := setup(afterMsgSent, queryAliveNodes)
	cpk, err := manager.generateKey([]string{"a", "b", "c", "d"}, 3)
	require.Error(t, err)
	require.ErrorContains(t, err, "found different CPKs")
	require.EqualValues(t, 0, len(cpk))
}

func TestSendErrorKeyGen(t *testing.T) {
	var afterMsgSent afterMsgSendFunc = func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		return errors.New("mock send error")
	}
	var queryAliveNodes queryAliveNodesFunc = func() []string {
		return []string{"a", "b", "c", "d"}
	}
	manager, _ := setup(afterMsgSent, queryAliveNodes)
	cpk, err := manager.generateKey([]string{"a", "b", "c", "d"}, 3)
	require.Error(t, err)
	require.ErrorContains(t, err, "failed to send request to node")
	require.EqualValues(t, 0, len(cpk))
}

func TestTimeoutKeygen(t *testing.T) {
	var afterMsgSent afterMsgSendFunc = func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		if request.TargetNode == "c" {
			return nil
		}
		askResp := tss.KeygenResponse{
			ClusterPublicKey: "abcd",
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, askResp)
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}
	var queryAliveNodes queryAliveNodesFunc = func() []string {
		return []string{"a", "b", "c", "d"}
	}
	manager, _ := setup(afterMsgSent, queryAliveNodes)
	before := time.Now()
	cpk, err := manager.generateKey([]string{"a", "b", "c", "d"}, 3)
	cost := time.Now().Sub(before)
	require.Error(t, err)
	require.ErrorContains(t, err, "timeout")
	require.EqualValues(t, 0, len(cpk))
	require.True(t, cost.Seconds() >= manager.keygenTimeout.Seconds())
}

func TestNotEnoughAliveNodesKeygen(t *testing.T) {
	var queryAliveNodes queryAliveNodesFunc = func() []string {
		return []string{"a", "b", "c"}
	}
	manager, _ := setup(nil, queryAliveNodes)
	cpk, err := manager.generateKey([]string{"a", "b", "c", "d"}, 3)
	require.Error(t, err)
	require.ErrorContains(t, err, "not enough available nodes")
	require.EqualValues(t, 0, len(cpk))
}

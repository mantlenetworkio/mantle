package manager

import (
	"errors"
	"testing"
	"time"

	tss "github.com/bitdao-io/bitnetwork/tss/common"
	"github.com/bitdao-io/bitnetwork/tss/manager/types"
	"github.com/bitdao-io/bitnetwork/tss/ws/server"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

type mockWsManager struct {
	responseCh   chan server.ResponseMsg
	afterMsgSent func(server.RequestMsg, chan server.ResponseMsg) error
}

func (mock *mockWsManager) AliveNodes() []string {
	return nil
}

func (mock *mockWsManager) RegisterResChannel(id string, responseMsg chan server.ResponseMsg, stopChan chan struct{}) error {
	mock.responseCh = responseMsg
	return nil
}

func (mock *mockWsManager) SendMsg(request server.RequestMsg) error {
	return mock.afterMsgSent(request, mock.responseCh)
}

func TestAgreement(t *testing.T) {
	// all return true
	afterMsgSent := func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		askResp := tss.AskResponse{
			Result: true,
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, askResp)
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}
	manager, request := prepareParams(afterMsgSent)
	ctx := types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold: 3,
		})
	ctx, err := manager.agreement(ctx, request, "ask")
	require.NoError(t, err)
	approvers := ctx.Approvers()
	require.EqualValues(t, 4, len(approvers))
}

func TestOneRefuseAgreement(t *testing.T) {
	// one returns false, others return true
	afterMsgSent := func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		askResp := tss.AskResponse{
			Result: true,
		}
		if request.TargetNode == "a" {
			askResp = tss.AskResponse{
				Result: false,
			}
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, askResp)
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}
	manager, request := prepareParams(afterMsgSent)
	ctx := types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold: 3,
		})
	ctx, err := manager.agreement(ctx, request, "ask")
	require.NoError(t, err)
	approvers := ctx.Approvers()
	require.EqualValues(t, 3, len(approvers))
}

func TestSentErrorAgreement(t *testing.T) {
	afterMsgSent := func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		askResp := tss.AskResponse{
			Result: true,
		}
		if request.TargetNode == "a" {
			return errors.New("failed to sent to a ")
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, askResp)
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}

	manager, request := prepareParams(afterMsgSent)
	ctx := types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold: 3,
		})
	ctx, err := manager.agreement(ctx, request, "ask")
	require.Error(t, err)
	require.Contains(t, err.Error(), "not enough response")

	ctx = types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold: 2,
		})
	ctx, err = manager.agreement(ctx, request, "ask")
	require.NoError(t, err)
	require.EqualValues(t, 3, len(ctx.Approvers()))
}

func TestErrorRespAgreement(t *testing.T) {
	afterMsgSent := func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		askResp := tss.AskResponse{
			Result: true,
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, askResp)
		if request.TargetNode == "a" {
			rpcResp = tmtypes.NewRPCErrorResponse(request.RpcRequest.ID, -1, "error response", "")
		}
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}
	manager, request := prepareParams(afterMsgSent)
	ctx := types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold: 3,
		})
	ctx, err := manager.agreement(ctx, request, "ask")
	require.Error(t, err)
	require.Contains(t, err.Error(), "not enough response")

	ctx = types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold: 2,
		})
	ctx, err = manager.agreement(ctx, request, "ask")
	require.NoError(t, err)
	require.EqualValues(t, 3, len(ctx.Approvers()))
}

func TestTimeoutAgreement(t *testing.T) {
	afterMsgSent := func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		askResp := tss.AskResponse{
			Result: true,
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, askResp)
		if request.TargetNode == "a" {
			return nil
		}
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}
	manager, request := prepareParams(afterMsgSent)
	ctx := types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold: 2,
		})
	before := time.Now()
	ctx, err := manager.agreement(ctx, request, "ask")
	require.NoError(t, err)
	costTime := time.Now().Sub(before)
	require.True(t, costTime.Seconds() >= askTimeOutSeconds)
	require.EqualValues(t, 3, len(ctx.Approvers()))
}

func prepareParams(afterMsgSent func(request server.RequestMsg, respCh chan server.ResponseMsg) error) (Manager, tss.SignStateRequest) {
	mock := mockWsManager{
		afterMsgSent: afterMsgSent,
	}
	manager := Manager{
		wsServer: &mock,
	}
	request := tss.SignStateRequest{
		StartBlock:          "1",
		OffsetStartsAtIndex: "1",
		StateRoots:          [][32]byte{},
	}
	return manager, request
}

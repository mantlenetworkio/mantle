package manager

import (
	"time"

	tss "github.com/mantlenetworkio/mantle/mt-tss/common"
	"github.com/mantlenetworkio/mantle/mt-tss/manager/store"
	"github.com/mantlenetworkio/mantle/mt-tss/ws/server"
)

type afterMsgSendFunc func(server.RequestMsg, chan server.ResponseMsg) error
type queryAliveNodesFunc func() []string

type mockWsManager struct {
	responseCh      chan server.ResponseMsg
	afterMsgSent    afterMsgSendFunc
	queryAliveNodes queryAliveNodesFunc
}

func (mock *mockWsManager) AliveNodes() []string {
	if mock.queryAliveNodes != nil {
		return mock.queryAliveNodes()
	}
	return nil
}

func (mock *mockWsManager) RegisterResChannel(id string, responseMsg chan server.ResponseMsg, stopChan chan struct{}) error {
	mock.responseCh = responseMsg
	return nil
}

func (mock *mockWsManager) SendMsg(request server.RequestMsg) error {
	return mock.afterMsgSent(request, mock.responseCh)
}

func setup(afterMsgSent afterMsgSendFunc, queryAliveNodes queryAliveNodesFunc) (Manager, tss.SignStateRequest) {
	mock := mockWsManager{
		afterMsgSent:    afterMsgSent,
		queryAliveNodes: queryAliveNodes,
	}
	storage, err := store.NewStorage("")
	if err != nil {
		panic(err)
	}
	manager := Manager{
		wsServer: &mock,
		store:    storage,

		askTimeout:        5 * time.Second,
		signTimeout:       5 * time.Second,
		keygenTimeout:     5 * time.Second,
		cpkConfirmTimeout: 5 * time.Second,
	}
	request := tss.SignStateRequest{
		StartBlock:          "1",
		OffsetStartsAtIndex: "1",
		StateRoots:          [][32]byte{},
	}
	return manager, request
}

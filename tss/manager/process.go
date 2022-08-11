package manager

import (
	"errors"

	"github.com/ethereum-optimism/optimism/tss/manager/types"
	"github.com/ethereum-optimism/optimism/tss/ws/server"
)

const threshold = 3

type Manager struct {
	wsServer server.IWebsocketManager
}

func NewManager(wsServer server.IWebsocketManager) Manager {
	return Manager{
		wsServer: wsServer,
	}
}

func (m Manager) SignStateBatch(request types.SignStateRequest) ([]byte, error) {
	activeNodes, pass := m.checkActiveNodes()
	if !pass {
		return nil, errors.New("not enough active nodes to sign state")
	}
	approvers := m.getApprovers(activeNodes, request)
	return m.sign(approvers)
}

func (m Manager) SignTxBatch() error {
	return errors.New("not support for now")
}

func (m Manager) checkActiveNodes() ([]string, bool) {
	activeNodes := m.wsServer.AliveNodes()
	if len(activeNodes) < threshold {
		return nil, false
	}
	return activeNodes, true
}

func (m Manager) getApprovers(nodes []string, request types.SignStateRequest) []string {
	return nodes
}

func (m Manager) sign(nodes []string) ([]byte, error) {
	return nil, nil
}

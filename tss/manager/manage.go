package manager

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/bitdao-io/bitnetwork/l2geth/ethclient"
	"github.com/bitdao-io/bitnetwork/tss/manager/types"
	tss "github.com/bitdao-io/bitnetwork/tss/types"
	"github.com/bitdao-io/bitnetwork/tss/ws/server"
	"github.com/influxdata/influxdb/pkg/slices"
)

type Manager struct {
	wsServer        server.IWebsocketManager
	tssQueryService types.TssQueryService
	cpkStore        types.CPKStore

	l1Cli      *ethclient.Client
	stopGenKey bool
}

func NewManager(wsServer server.IWebsocketManager, tssQueryService types.TssQueryService, cpkStore types.CPKStore, l1Url string) (Manager, error) {
	l1Cli, err := ethclient.Dial(l1Url)
	if err != nil {
		return Manager{}, err
	}
	return Manager{
		wsServer:        wsServer,
		tssQueryService: tssQueryService,
		cpkStore:        cpkStore,
		l1Cli:           l1Cli,
	}, nil
}

func (m Manager) Start() {
	go m.observeElection()
}

func (m Manager) stopGenerateKey() {
	m.stopGenKey = true
}

func (m Manager) recoverGenerateKey() {
	m.stopGenKey = false
}

func (m Manager) SignStateBatch(request tss.SignStateRequest) ([]byte, error) {
	tssInfo := m.tssQueryService.QueryInfo()
	availableNodes := m.availableNodes(tssInfo.PartyPubKeys)
	if len(availableNodes) < tssInfo.Threshold+1 {
		return nil, errors.New("not enough available nodes to sign state")
	}
	ctx := types.NewContext().
		WithAvailableNodes(availableNodes).
		WithTssInfo(tssInfo).
		WithRequestId(randomRequestId())
	ctx, err := m.agreement(ctx, request)
	if err != nil {
		return nil, err
	}
	if len(ctx.Approvers()) < ctx.TssInfos().Threshold+1 {
		return nil, errors.New("failed to sign, not enough approvals from tss nodes")
	}
	return m.sign(ctx, request)
}

func (m Manager) SignTxBatch() error {
	return errors.New("not support for now")
}

func (m Manager) availableNodes(tssMembers []string) []string {
	aliveNodes := m.wsServer.AliveNodes()
	availableNodes := make([]string, 0)
	for _, n := range aliveNodes {
		if slices.Exists(tssMembers, n) {
			availableNodes = append(availableNodes, n)
		}
	}
	return availableNodes
}

// current time + 4 digit random number
func randomRequestId() string {
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	return time.Now().Format("20060102150405") + code
}

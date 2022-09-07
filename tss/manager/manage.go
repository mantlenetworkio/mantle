package manager

import (
	"errors"
	"fmt"
	"github.com/bitdao-io/mantle/tss/bindings/tsh"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/bitdao-io/mantle/l2geth/log"
	tss "github.com/bitdao-io/mantle/tss/common"
	"github.com/bitdao-io/mantle/tss/index"
	"github.com/bitdao-io/mantle/tss/manager/types"
	"github.com/bitdao-io/mantle/tss/slash"
	"github.com/bitdao-io/mantle/tss/ws/server"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/influxdata/influxdb/pkg/slices"
)

type Manager struct {
	wsServer                 server.IWebsocketManager
	tssQueryService          types.TssQueryService
	store                    types.ManagerStore
	l1Cli                    *ethclient.Client
	tssStakingSlashingCaller *tsh.TssStakingSlashingCaller
	l1ConfirmBlocks          int

	taskInterval          time.Duration
	confirmReceiptTimeout time.Duration
	keygenTimeout         time.Duration
	cpkConfirmTimeout     time.Duration
	askTimeout            time.Duration
	signTimeout           time.Duration

	stateSignatureCache map[[32]byte][]byte
	sigCacheLock        *sync.RWMutex
	stopGenKey          bool
	stopChan            chan struct{}
}

func NewManager(wsServer server.IWebsocketManager,
	tssQueryService types.TssQueryService,
	store types.ManagerStore,
	config tss.Configuration) (Manager, error) {
	taskIntervalDur, err := time.ParseDuration(config.TimedTaskInterval)
	if err != nil {
		return Manager{}, err
	}
	receiptConfirmTimeoutDur, err := time.ParseDuration(config.L1ReceiptConfirmTimeout)
	if err != nil {
		return Manager{}, err
	}
	keygenTimeoutDur, err := time.ParseDuration(config.Manager.KeygenTimeout)
	if err != nil {
		return Manager{}, err
	}
	cpkConfirmTimeoutDur, err := time.ParseDuration(config.Manager.CPKConfirmTimeout)
	if err != nil {
		return Manager{}, err
	}
	askTimeoutDur, err := time.ParseDuration(config.Manager.AskTimeout)
	if err != nil {
		return Manager{}, err
	}
	signTimeoutDur, err := time.ParseDuration(config.Manager.SignTimeout)
	if err != nil {
		return Manager{}, err
	}

	l1Cli, err := ethclient.Dial(config.L1Url)
	if err != nil {
		return Manager{}, err
	}
	tssStakingSlashingCaller, err := tsh.NewTssStakingSlashingCaller(common.HexToAddress(config.TssStakingSlashContractAddress), l1Cli)

	if err != nil {
		return Manager{}, err
	}
	return Manager{
		wsServer:                 wsServer,
		tssQueryService:          tssQueryService,
		store:                    store,
		l1Cli:                    l1Cli,
		l1ConfirmBlocks:          config.L1ConfirmBlocks,
		tssStakingSlashingCaller: tssStakingSlashingCaller,

		taskInterval:          taskIntervalDur,
		confirmReceiptTimeout: receiptConfirmTimeoutDur,
		keygenTimeout:         keygenTimeoutDur,
		cpkConfirmTimeout:     cpkConfirmTimeoutDur,
		askTimeout:            askTimeoutDur,
		signTimeout:           signTimeoutDur,

		stateSignatureCache: make(map[[32]byte][]byte),
		sigCacheLock:        &sync.RWMutex{},
		stopChan:            make(chan struct{}),
	}, nil
}

// Start launch a manager
func (m Manager) Start() {
	go m.observeElection()
	go m.slashing()
}

func (m Manager) Stop() {
	close(m.stopChan)
}

func (m Manager) stopGenerateKey() {
	m.stopGenKey = true
}

func (m Manager) recoverGenerateKey() {
	m.stopGenKey = false
}

func (m Manager) SignStateBatch(request tss.SignStateRequest) ([]byte, error) {
	offsetStartsAtIndex, _ := new(big.Int).SetString(request.OffsetStartsAtIndex, 10)
	digestBz, err := tss.StateBatchHash(request.StateRoots, offsetStartsAtIndex)
	if err != nil {
		return nil, err
	}
	if sig := m.getStateSignature(digestBz); len(sig) > 0 {
		return sig, nil
	}

	tssInfo, err := m.tssQueryService.QueryActiveInfo()
	if err != nil {
		return nil, err
	}
	availableNodes := m.availableNodes(tssInfo.TssMembers)
	if len(availableNodes) < tssInfo.Threshold+1 {
		return nil, errors.New("not enough available nodes to sign state")
	}
	stateBatchRoot, _ := tss.GetMerkleRoot(request.StateRoots)
	found, stateBatch := m.store.GetStateBatch(stateBatchRoot)
	if found && stateBatch.BatchIndex != 0 {
		return nil, errors.New("the state batch is already indexed on layer1")
	}
	ctx := types.NewContext().
		WithAvailableNodes(availableNodes).
		WithTssInfo(tssInfo).
		WithRequestId(randomRequestId()).
		WithElectionId(tssInfo.ElectionId)

	// ask tss nodes for the agreement
	ctx, err = m.agreement(ctx, request, tss.AskStateBatch)
	if err != nil {
		return nil, err
	}
	if len(ctx.Approvers()) < ctx.TssInfos().Threshold+1 {
		return nil, errors.New("failed to sign, not enough approvals from tss nodes")
	}

	request.ElectionId = tssInfo.ElectionId
	resp, culprits, err := m.sign(ctx, request, digestBz, tss.SignStateBatch)
	if err != nil {
		for _, culprit := range culprits {
			addr, err := tss.NodeToAddress(culprit)
			if err != nil {
				log.Error("failed to convert node to address", "public key", culprit, "err", err)
				continue
			}
			m.store.SetSlashingInfo(slash.SlashingInfo{
				Address:    addr,
				ElectionId: tssInfo.ElectionId,
				BatchIndex: math.MaxUint64, // not real, just for identifying the specific slashing info.
				SlashType:  tss.SlashTypeCulprit,
			})
		}
		m.store.AddCulprits(culprits)
		return nil, err
	}
	absents := make([]string, 0)
	for _, node := range tssInfo.TssMembers {
		if !slices.ExistsIgnoreCase(ctx.Approvers(), node) {
			// 并且node不在slashing中
			addr, _ := tss.NodeToAddress(node)
			if !m.store.IsInSlashing(addr) {
				absents = append(absents, node)
			}
		}
	}
	if err = m.afterSignStateBatch(ctx, request.StateRoots, absents); err != nil {
		log.Error("failed to execute afterSign", "err", err)
	}

	m.setStateSignature(digestBz, resp.Signature)
	return resp.Signature, nil
}

func (m Manager) SignTxBatch() error {
	return errors.New("not support for now")
}

func (m Manager) availableNodes(tssMembers []string) []string {
	aliveNodes := m.wsServer.AliveNodes()
	availableNodes := make([]string, 0)
	for _, n := range aliveNodes {
		if slices.ExistsIgnoreCase(tssMembers, n) {
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

func (m Manager) afterSignStateBatch(ctx types.Context, stateBatch [][32]byte, absentNodes []string) error {
	batchRoot, err := tss.GetMerkleRoot(stateBatch)
	if err != nil {
		return err
	}
	sbi := index.StateBatchInfo{
		BatchRoot:    batchRoot,
		ElectionId:   ctx.ElectionId(),
		AbsentNodes:  absentNodes,
		WorkingNodes: ctx.AvailableNodes(),
	}
	if err = m.store.SetStateBatch(sbi); err != nil {
		return err
	}
	return nil
}

func (m Manager) getStateSignature(digestBz []byte) []byte {
	m.sigCacheLock.RLock()
	defer m.sigCacheLock.RUnlock()
	var key [32]byte
	copy(key[:], digestBz)
	return m.stateSignatureCache[key]
}

func (m Manager) setStateSignature(digestBz []byte, sig []byte) {
	m.sigCacheLock.Lock()
	defer m.sigCacheLock.Unlock()
	var key [32]byte
	copy(key[:], digestBz)
	m.stateSignatureCache[key] = sig
}

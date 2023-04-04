package manager

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mantlenetworkio/mantle/mt-bindings/bindings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/influxdata/influxdb/pkg/slices"
	tss "github.com/mantlenetworkio/mantle/mt-tss/common"
	"github.com/mantlenetworkio/mantle/mt-tss/index"
	"github.com/mantlenetworkio/mantle/mt-tss/manager/types"
	"github.com/mantlenetworkio/mantle/mt-tss/slash"
	"github.com/mantlenetworkio/mantle/mt-tss/ws/server"
)

type Manager struct {
	wsServer                 server.IWebsocketManager
	tssQueryService          types.TssQueryService
	store                    types.ManagerStore
	l1Cli                    *ethclient.Client
	tssStakingSlashingCaller *bindings.TssStakingSlashingCaller
	tssGroupManagerCaller    *bindings.TssGroupManagerCaller
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
	tssStakingSlashingCaller, err := bindings.NewTssStakingSlashingCaller(common.HexToAddress(config.TssStakingSlashContractAddress), l1Cli)
	if err != nil {
		return Manager{}, err
	}
	tssGroupManagerCaller, err := bindings.NewTssGroupManagerCaller(common.HexToAddress(config.TssGroupContractAddress), l1Cli)
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
		tssGroupManagerCaller:    tssGroupManagerCaller,

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
	log.Info("manager is starting......")
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

func (m Manager) SignStateBatch(request tss.SignOutputRequest) ([]byte, error) {
	log.Info("received sign output request", "output", request.String())
	l1_block_number := new(big.Int).SetUint64(request.L1BlockNumber)
	l2_block_number := new(big.Int).SetUint64(request.L2BlockNumber)
	l1_block_hash := common.HexToHash(request.L1BlockHash)
	digestBz, err := tss.StateBatchHash(request.OutputRoot, l2_block_number, l1_block_hash, l1_block_number)
	if err != nil {
		return nil, err
	}
	if sig := m.getStateSignature(digestBz); len(sig) > 0 {
		log.Info("get stored signature ", "digest", hex.EncodeToString(digestBz), "sig", hex.EncodeToString(sig))
		response := tss.BatchSubmitterResponse{
			Signature: sig,
			RollBack:  false,
		}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Error("batch submitter response failed to marshal !")
			return nil, err
		}
		return responseBytes, nil
	}

	tssInfo, err := m.tssQueryService.QueryActiveInfo()
	if err != nil {
		return nil, err
	}
	availableNodes := m.availableNodes(tssInfo.TssMembers)
	if len(availableNodes) < tssInfo.Threshold+1 {
		return nil, errors.New("not enough available nodes to sign state")
	}

	found, outputInfo := m.store.GetOutput(request.OutputRoot)
	if found && outputInfo.OutputIndex != 0 {
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
	var resp tss.SignResponse
	var culprits []string
	var rollback bool
	var signErr error

	if len(ctx.Approvers()) < ctx.TssInfos().Threshold+1 {
		if len(ctx.UnApprovers()) < ctx.TssInfos().Threshold+1 {
			return nil, errors.New("failed to sign, approvals " + strings.Join(ctx.Approvers(), ",") + " ,unApprovals " + strings.Join(ctx.UnApprovers(), ","))
		}
		log.Warn("failed to approval from tss nodes , there is wrong state root in batch state roots.need to roll back l2chain to batch index !")
		//TODO bedrock rollback
		//change unApprovals to approvals to do sign
		ctx = ctx.WithApprovers(ctx.UnApprovers())
		rollback = true
		startBlock := new(big.Int).SetUint64(request.L2BlockNumber)
		rollBackRequest := tss.RollBackRequest{StartBlock: startBlock.String()}
		rollBackBz, err := tss.RollBackHash(startBlock)
		if err != nil {
			return nil, err
		}
		resp, culprits, signErr = m.sign(ctx, rollBackRequest, rollBackBz, tss.SignRollBack)
	} else {
		request.ElectionId = tssInfo.ElectionId
		resp, culprits, signErr = m.sign(ctx, request, digestBz, tss.SignStateBatch)
	}

	if signErr != nil {
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
		return nil, signErr
	}

	if !rollback {
		absents := make([]string, 0)
		for _, node := range tssInfo.TssMembers {
			if !slices.ExistsIgnoreCase(ctx.Approvers(), node) {
				addr, _ := tss.NodeToAddress(node)
				if !m.store.IsInSlashing(addr) {
					absents = append(absents, node)
				}
			}
		}
		if err = m.afterSignStateBatch(ctx, request.OutputRoot, absents); err != nil {
			log.Error("failed to execute afterSign", "err", err)
		}
		m.setStateSignature(digestBz, resp.Signature)
	}

	response := tss.BatchSubmitterResponse{
		Signature: resp.Signature,
		RollBack:  rollback,
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Error("batch submitter response failed to marshal !")
		return nil, err
	}
	return responseBytes, nil
}

func (m Manager) SignTxBatch() error {
	return errors.New("not support for now")
}

func (m Manager) availableNodes(tssMembers []string) []string {
	aliveNodes := m.wsServer.AliveNodes()
	log.Info("check available nodes", "expected", fmt.Sprintf("%v", tssMembers), "alive nodes", fmt.Sprintf("%v", aliveNodes))
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

func (m Manager) afterSignStateBatch(ctx types.Context, outputRoot [32]byte, absentNodes []string) error {
	sbi := index.OutputInfo{
		OutputRoot:   outputRoot,
		ElectionId:   ctx.ElectionId(),
		AbsentNodes:  absentNodes,
		WorkingNodes: ctx.AvailableNodes(),
	}
	log.Info("Store the signed state batch", "batchRoot", hex.EncodeToString(outputRoot[:]))
	if err := m.store.SetOutput(sbi); err != nil {
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

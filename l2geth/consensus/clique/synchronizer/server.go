package synchronizer

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	binding "github.com/mantlenetworkio/mantle/l2geth/consensus/clique/synchronizer/bindings"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

const (
	ctxTimeout = 100 * time.Second

	ENV_SEQUENCER_CONTRACT_ADDRESS = "SEQUENCER_CONTRACT_ADDRESS"
	ENV_SEQUENCER_L1_RPC           = "SEQUENCER_L1_RPC"
	GET_SEQUENCER_RETRY_TIMES      = 1000
	GET_SEQUENCER_TIMEOUT          = 10
)

// SequencerSequencerInfos implementation the sort interface
var _ sort.Interface = SequencerSequencerInfos{}

// SequencerSequencerInfos set infos for sort
type SequencerSequencerInfos []binding.SequencerSequencerInfo

// Len for sort
func (s SequencerSequencerInfos) Len() int {
	return len(s)
}

// Less for sort
func (s SequencerSequencerInfos) Less(i, j int) bool {
	return s[i].Amount.Int64() < s[j].Amount.Int64()
}

// Swap for sort
func (s SequencerSequencerInfos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Synchronizer set for connect with l1 sequencer contract
type Synchronizer struct {
	sequencerContractAddr common.Address
	rpcUrl                string
}

// NewSynchronizer get sequencer address and l1 rpc address from os env then create Synchronizer
func NewSynchronizer() *Synchronizer {
	seqAddr := os.Getenv(ENV_SEQUENCER_CONTRACT_ADDRESS)
	rpcUrl := os.Getenv(ENV_SEQUENCER_L1_RPC)
	return &Synchronizer{
		sequencerContractAddr: common.HexToAddress(seqAddr),
		rpcUrl:                rpcUrl,
	}
}

// GetSchedulerAddr query scheduler address from l1 sequencer contract
func (sync *Synchronizer) GetSchedulerAddr() (common.Address, error) {
	// Create eth client
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	rpcClient, err := rpc.DialContext(ctx, sync.rpcUrl)
	if err != nil {
		return common.Address{}, fmt.Errorf("create rpcClient failed: %s", err)
	}
	ethClient := ethclient.NewClient(rpcClient)

	seqContractInst, err := binding.NewSequencer(sync.sequencerContractAddr, ethClient)
	if err != nil {
		return common.Address{}, fmt.Errorf("create rpcClient failed: %s", err)
	}
	schedulerAddr, err := seqContractInst.Scheduler(nil)
	for schedulerAddr.String() == (common.Address{}.String()) && err == nil {
		log.Info("retry get scheduler", "addr", schedulerAddr)
		time.Sleep(GET_SEQUENCER_TIMEOUT * time.Second)
		schedulerAddr, err = seqContractInst.Scheduler(nil)
	}

	return schedulerAddr, err
}

// GetSequencerSet will query the sequencer set from l1 sequencer contract
func (sync *Synchronizer) GetSequencerSet() (SequencerSequencerInfos, error) {
	// Create eth client
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	rpcClient, err := rpc.DialContext(ctx, sync.rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("create rpcClient failed: %s", err)
	}
	ethClient := ethclient.NewClient(rpcClient)

	seqContractInst, err := binding.NewSequencer(sync.sequencerContractAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("create sequencer contract instance failed: %s", err)
	}
	num, err := seqContractInst.SequencerLimit(nil)
	if err != nil {
		return nil, err
	}
	// get sequencers from sequencer contract
	var seqInfos SequencerSequencerInfos
	for i := 0; i < GET_SEQUENCER_RETRY_TIMES; i++ {
		seqInfos, err = seqContractInst.GetSequencers(nil)
		if err != nil {
			continue
		}
		time.Sleep(GET_SEQUENCER_TIMEOUT * time.Second)
		if len(seqInfos) == 0 {
			return nil, fmt.Errorf("empty sequencer set")
		}
		if len(seqInfos) != 0 {
			break
		}
		log.Info("retry get sequencer", "time", time.Now(),
			"retry_time", i)
	}
	if len(seqInfos) == 0 {
		return nil, fmt.Errorf("empty sequencer set")
	}
	// sort by deposit amount ascending
	sort.Stable(seqInfos)
	// find users deposit and infos from the limit
	if seqInfos.Len() > int(num) {
		seqInfos = seqInfos[seqInfos.Len()-int(num):]
	}
	return seqInfos, nil
}

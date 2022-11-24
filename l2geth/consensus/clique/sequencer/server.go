package sequencer

import (
	"context"
	"errors"
	"os"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	binding "github.com/mantlenetworkio/mantle/l2geth/consensus/clique/sequencer/bindings"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

const ENV_SEQUENCER_CONTRACT_ADDRESS = "SEQUENCER_CONTRACT_ADDRESS"
const ENV_SEQUENCER_L1_RPC = "SEQUENCER_L1_RPC"
const ctxTimeout = 100 * time.Second

var rpcUrl string
var seqAddr string
var seqContract *binding.Sequencer

// Initialize will get Sequencer contract address and rpc url from os env, then init seqContract.
func Initialize() {
	// get params from os env
	seqAddr = os.Getenv(ENV_SEQUENCER_CONTRACT_ADDRESS)
	rpcUrl = os.Getenv(ENV_SEQUENCER_L1_RPC)
	addr := common.HexToAddress(seqAddr)

	// Create eth client
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	rpcClient, err := rpc.DialContext(ctx, rpcUrl)
	if err != nil {
		log.Error("create rpcClient failed: ", err)
		return
	}
	ethClient := ethclient.NewClient(rpcClient)

	// init sequencer contract
	seqContract, err = binding.NewSequencer(
		addr, ethClient,
	)
	if err != nil {
		log.Error("create seqContract failed: ", err)
		return
	}
}

// set infos for sort
type SequencerSequencerInfos []binding.SequencerSequencerInfo

func (seqs SequencerSequencerInfos) Len() int {
	return len(seqs)
}

func (s SequencerSequencerInfos) Less(i, j int) bool {
	return s[i].Amount.Int64() < s[j].Amount.Int64()
}

func (s SequencerSequencerInfos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func GetScheduler() (common.Address, error) {
	return seqContract.Scheduler(nil)
}

// GetSequencerSet will return the validator set
func GetSequencerSet() (SequencerSequencerInfos, error) {
	// get sequencer limit from sequencer contract
	num, err := seqContract.SequencerLimit(nil)
	if err != nil {
		return nil, err
	}
	// get sequencers from sequencer contract
	var seqInfos SequencerSequencerInfos
	seqInfos, err = seqContract.GetSequencers(nil)
	if err != nil {
		return nil, err
	}
	if len(seqInfos) == 0 {
		return nil, errors.New("Do not have sequencers")
	}
	// sort by deposit amount ascending
	sort.Stable(seqInfos)
	// find users deposit and infos from the limit
	if seqInfos.Len() > int(num) {
		seqInfos = seqInfos[seqInfos.Len()-int(num):]
	}
	return seqInfos, nil
}

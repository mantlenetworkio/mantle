package scheduler

import (
	"bytes"
	"context"
	"errors"
	"math"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/mantlenetworkio/mantle/bss-core/metrics"
	seq "github.com/mantlenetworkio/mantle/scheduler/bindings/sequencer"
	"github.com/mantlenetworkio/mantle/scheduler/service"
)

// stateRootSize is the size in bytes of a state root.
const stateRootSize = 32

var bigOne = new(big.Int).SetUint64(1) //nolint:unused

type Config struct {
	Name          string
	L1Client      *ethclient.Client
	SequencerAddr common.Address // sequencer contract address
	SequencerNum  int
}

// Driver implements the service.Driver interface.
var _ service.Driver = &Driver{}

type Driver struct {
	cfg               Config
	sequencerContract *seq.Sequencer
	seqz              *SequencerSet
	block             uint64     // current block height
	producer          *Sequencer // currenct producer
	metrics           *metrics.Base
}

func NewDriver(cfg Config) (*Driver, error) {
	seqContract, err := seq.NewSequencer(
		cfg.SequencerAddr, cfg.L1Client,
	)
	if err != nil {
		return nil, err
	}

	// get sequencers from Sequencer contract
	seqz, err := GetSequencerSet(seqContract, cfg.SequencerNum)
	if err != nil {
		return nil, err
	}

	return &Driver{
		cfg:               cfg,
		sequencerContract: seqContract,
		seqz:              NewSequencerSet(seqz),
		metrics:           metrics.NewBase("scheduler", cfg.Name),
	}, nil
}

// Name is an identifier used to prefix logs for a particular service.
func (d *Driver) Name() string {
	return d.cfg.Name
}

// Metrics returns the subservice telemetry object.
func (d *Driver) Metrics() metrics.Metrics {
	return d.metrics
}

// ExecuteProcess executes a process on this driver.
func (d *Driver) ExecuteProcess(ctx context.Context) error {
	seqz, err := GetSequencerSet(d.sequencerContract, d.cfg.SequencerNum)
	if err != nil {
		return err
	}
	changes := d.CompareValidatorSet(seqz)

	d.seqz.updateWithChangeSet(changes, true)
	return nil
}

// SwithProducer switches the current producer.
// This function will be called by liveless issue.
func (d *Driver) SwitchProducer(ctx context.Context) error {
	d.seqz.IncrementProducerPriority(0)
	return nil
}

type SequencerSequencerInfos []seq.SequencerSequencerInfo

func (seqs SequencerSequencerInfos) Len() int {
	return len(seqs)
}

func (s SequencerSequencerInfos) Less(i, j int) bool {
	return s[i].Amount.Int64() < s[j].Amount.Int64()
}

func (s SequencerSequencerInfos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// GetSequencerSet will return the validator set
func GetSequencerSet(seqContract *seq.Sequencer, num int) ([]*Sequencer, error) {
	// get sequencers from sequencer contract
	var seqInfos SequencerSequencerInfos
	seqInfos, err := seqContract.GetSequencers(nil)
	if err != nil {
		return nil, err
	}
	if len(seqInfos) == 0 {
		return nil, errors.New("Do not have sequencers")
	}
	sort.Stable(seqInfos)
	// find users deposit and infos
	if seqInfos.Len() > num {
		seqInfos = seqInfos[seqInfos.Len()-num:]
	}
	// set sequencer, voting power = deposit / 10^18
	scale := int64(math.Pow10(18))
	var seqz []*Sequencer
	for _, v := range seqInfos {
		seq := &Sequencer{
			Address:     v.MintAddress,
			NodeID:      v.NodeID,
			VotingPower: v.Amount.Div(v.Amount, big.NewInt(scale)).Int64(),
		}
		if err = seq.SequencerBasic(); err != nil {
			return nil, err
		}
		seqz = append(seqz, seq)
	}

	return seqz, nil
}

// CompareValidatorSet will return the update with Driver.seqz
func (d *Driver) CompareValidatorSet(new []*Sequencer) []*Sequencer {
	var changes []*Sequencer
	for i, v := range new {
		changed := true
		for _, seq := range d.seqz.Sequencers {
			if bytes.Equal(seq.Address.Bytes(), v.Address.Bytes()) && v.VotingPower == seq.VotingPower {
				changed = false
				break
			}
		}
		if changed {
			changes = append(changes, new[i])
		}
	}
	return changes
}

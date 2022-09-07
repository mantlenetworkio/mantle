package scheduler

import (
	"context"
	"math/big"

	"github.com/bitdao-io/mantle/batch-submitter/bindings/scc"
	"github.com/bitdao-io/mantle/bss-core/metrics"
	common2 "github.com/bitdao-io/mantle/l2geth/common"
	"github.com/bitdao-io/mantle/scheduler/service"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// stateRootSize is the size in bytes of a state root.
const stateRootSize = 32

var bigOne = new(big.Int).SetUint64(1) //nolint:unused

type Config struct {
	Name          string
	L1Client      *ethclient.Client
	SequencerAddr common.Address // sequencer contract address
}

// Driver implements the service.Driver interface.
var _ service.Driver = &Driver{}

type Driver struct {
	cfg               Config
	sequencerContract *scc.StateCommitmentChain
	seqz              *SequencerSet
	block             uint64     // current block height
	producer          *Sequencer // currenct producer
	metrics           *metrics.Base
}

func NewDriver(cfg Config) (*Driver, error) {
	sequencerContract, err := scc.NewStateCommitmentChain(
		cfg.SequencerAddr, cfg.L1Client,
	)
	if err != nil {
		return nil, err
	}

	// TODO: get sequencers from Sequencer contract
	// d.sequencerContract.getSequencers()
	sequencers := []common.Address{}
	var seqz []*Sequencer
	for _, sequencer := range sequencers {
		seq := &Sequencer{
			Address: common2.Address(sequencer),
		}
		seqz = append(seqz, seq)
	}

	return &Driver{
		cfg:               cfg,
		sequencerContract: sequencerContract,
		seqz:              NewSequencerSet(seqz),
		metrics:           metrics.NewBase("batch_submitter", cfg.Name),
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
	// TODO: get update sequencers from Sequencer contract
	// d.sequencerContract.getSequencers()
	updateSequencers := []common.Address{}
	var seqz []*Sequencer
	for _, sequencer := range updateSequencers {
		seq := &Sequencer{
			Address: common2.Address(sequencer),
		}
		seqz = append(seqz, seq)
	}
	d.seqz.updateWithChangeSet(seqz, true)
	return nil
}

// SwithProducer switches the current producer.
// This function will be called by liveless issue.
func (d *Driver) SwitchProducer(ctx context.Context) error {
	d.seqz.IncrementProducerPriority(0)
	return nil
}

package scheduler

import (
	"bytes"
	"context"
	"math/big"

	"github.com/bitdao-io/bitnetwork/bss-core/metrics"
	"github.com/bitdao-io/bitnetwork/l2geth/crypto"
	"github.com/bitdao-io/bitnetwork/scheduler/bindings/tgm"
	"github.com/bitdao-io/bitnetwork/scheduler/bindings/tsh"
	"github.com/bitdao-io/bitnetwork/scheduler/service"
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
	TSHAddr       common.Address // sequencer staking address
}

// Driver implements the service.Driver interface.
var _ service.Driver = &Driver{}

type Driver struct {
	cfg         Config
	tgmContract *tgm.TssGroupManager
	tshContract *tsh.TssStakingSlashing
	seqz        *SequencerSet
	block       uint64     // current block height
	producer    *Sequencer // currenct producer
	metrics     *metrics.Base
}

func NewDriver(cfg Config) (*Driver, error) {
	tgmContract, err := tgm.NewTssGroupManager(
		cfg.SequencerAddr, cfg.L1Client,
	)
	if err != nil {
		return nil, err
	}

	tshContract, err := tsh.NewTssStakingSlashing(
		cfg.TSHAddr, cfg.L1Client,
	)
	if err != nil {
		return nil, err
	}

	// get sequencers from Sequencer contract
	pubKeys, err := tgmContract.GetTssGroupMembers(nil)
	if err != nil {
		return nil, err
	}
	var seqz []*Sequencer
	for _, v := range pubKeys {
		ecdsaPubKey, err := crypto.UnmarshalPubkey(v)
		if err != nil {
			return nil, err
		}
		addr := crypto.PubkeyToAddress(*ecdsaPubKey)
		deposit, err := tshContract.GetDeposits(nil, common.BytesToAddress(addr.Bytes()))

		seq := &Sequencer{
			Address: addr,
			PubKey:  *ecdsaPubKey,
			// todo : caculate voting power
			VotingPower: deposit.Amount.Int64(),
		}
		seqz = append(seqz, seq)
	}

	return &Driver{
		cfg:         cfg,
		tgmContract: tgmContract,
		tshContract: tshContract,
		seqz:        NewSequencerSet(seqz),
		metrics:     metrics.NewBase("batch_submitter", cfg.Name),
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
	// get update sequencers from Sequencer contract
	pubKeys, err := d.tgmContract.GetTssGroupMembers(nil)
	if err != nil {
		return err
	}
	var seqz []*Sequencer
	for _, v := range pubKeys {
		ecdsaPubKey, err := crypto.UnmarshalPubkey(v)
		if err != nil {
			return err
		}
		addr := crypto.PubkeyToAddress(*ecdsaPubKey)
		deposit, err := d.tshContract.GetDeposits(nil, common.BytesToAddress(addr.Bytes()))

		seq := &Sequencer{
			Address: addr,
			PubKey:  *ecdsaPubKey,
			// todo : caculate voting power
			VotingPower: deposit.Amount.Int64(),
		}
		seqz = append(seqz, seq)
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

// CompareValidatorSet will return the update with Driver.seqz
func (d *Driver) CompareValidatorSet(new []*Sequencer) []*Sequencer {
	var changes []*Sequencer
	unchangeIndex := []int{}
	for i, v := range new {
		for _, seq := range d.seqz.Sequencers {
			if bytes.Equal(seq.Address.Bytes(), v.Address.Bytes()) && v.VotingPower == seq.VotingPower {
				unchangeIndex = append(unchangeIndex, i)
				break
			}
		}
	}
	for i := 0; len(unchangeIndex) > 0; i++ {
		if i == unchangeIndex[0] {
			unchangeIndex = unchangeIndex[1:]
			continue
		}
		changes = append(changes, new[i])
	}
	return changes
}

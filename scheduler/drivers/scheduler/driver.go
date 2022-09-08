package scheduler

import (
	"bytes"
	"context"
	"math"
	"math/big"

	"github.com/bitdao-io/mantle/batch-submitter/bindings/scc"
	"github.com/bitdao-io/mantle/bss-core/metrics"
	common2 "github.com/bitdao-io/mantle/l2geth/common"
	"github.com/bitdao-io/mantle/scheduler/service"
	"github.com/bitdao-io/bitnetwork/bss-core/metrics"
	common2 "github.com/bitdao-io/bitnetwork/l2geth/common"
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
	seqz, err := GetSequencerSet(tgmContract, tshContract)
	if err != nil {
		return nil, err
	}

	return &Driver{
		cfg:         cfg,
		tgmContract: tgmContract,
		tshContract: tshContract,
		seqz:        NewSequencerSet(seqz),
		metrics:     metrics.NewBase("scheduler", cfg.Name),
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
	seqz, err := GetSequencerSet(d.tgmContract, d.tshContract)
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

// GetSequencerSet will return the validator set
func GetSequencerSet(tgmContract *tgm.TssGroupManager, tshContract *tsh.TssStakingSlashing) ([]*Sequencer, error) {
	// get activeTssMembers public keys from tss group contract
	pubKeys, err := tgmContract.GetTssGroupMembers(nil)
	if err != nil {
		return nil, err
	}

	// find users deposit and infos
	var users []common.Address
	for _, v := range pubKeys {
		ecdsaPubKey, err := crypto.UnmarshalPubkey(v)
		if err != nil {
			return nil, err
		}
		addr := crypto.PubkeyToAddress(*ecdsaPubKey)
		users = append(users, common.BytesToAddress(addr.Bytes()))
	}
	valSetInfo, err := tshContract.BatchGetDeposits(nil, users)
	if err != nil {
		return nil, err
	}

	// set sequencer, voting power = deposit / 10^18
	scale := int64(math.Pow10(18))
	var seqz []*Sequencer
	for _, v := range valSetInfo {
		ecdsaPubKey, err := crypto.UnmarshalPubkey(v.PubKey)
		if err != nil {
			return nil, err
		}
		seq := &Sequencer{
			Address:     common2.BytesToAddress(v.Pledgor.Bytes()),
			PubKey:      *ecdsaPubKey,
			VotingPower: v.Amount.Div(v.Amount, big.NewInt(scale)).Int64(),
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

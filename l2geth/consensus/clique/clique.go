package clique

import (
	"bytes"
	"errors"
	"io"
	"math/big"
	"sync"
	"time"

	"golang.org/x/crypto/sha3"

	lru "github.com/hashicorp/golang-lru"
	"github.com/mantlenetworkio/mantle/l2geth/accounts"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/consensus"
	"github.com/mantlenetworkio/mantle/l2geth/consensus/misc"
	"github.com/mantlenetworkio/mantle/l2geth/core/state"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/ethdb"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/params"
	"github.com/mantlenetworkio/mantle/l2geth/rlp"
	"github.com/mantlenetworkio/mantle/l2geth/rollup/rcfg"
	"github.com/mantlenetworkio/mantle/l2geth/rpc"
)

var SequencerSetKey = []byte{0x00}

const (
	checkpointInterval = 1024 // Number of blocks after which to save the vote snapshot to the database
	inmemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	inmemorySignatures = 4096 // Number of recent block signatures to keep in memory

	wiggleTime = 500 * time.Millisecond // Random delay (per signer) to allow concurrent signers
)

var (
	extraVanity = 32                       // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = crypto.SignatureLength   // Fixed number of extra-data suffix bytes reserved for signer seal
	uncleHash   = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.
)

var (
	errUnknownBlock       = errors.New("unknown block")
	errMissingSignature   = errors.New("extra-data 65 byte signature suffix missing")
	errMissingVanity      = errors.New("extra-data 32 byte vanity prefix missing")
	errInvalidMixDigest   = errors.New("non-zero mix digest")
	errInvalidUncleHash   = errors.New("non empty uncle hash")
	errInvalidDifficulty  = errors.New("invalid difficulty")
	errUnauthorizedSigner = errors.New("unauthorized signer")
	errInvalidTimestamp   = errors.New("invalid timestamp")
)

type SignerFn func(accounts.Account, string, []byte) ([]byte, error)

type Clique struct {
	config *params.CliqueConfig // Consensus engine configuration parameters
	db     ethdb.Database

	signatures  *lru.ARCCache // Signatures of recent blocks to speed up mining
	producers   Producers
	signer      common.Address
	signFn      SignerFn
	schedulerID []byte
	lock        sync.RWMutex
}

func New(config *params.CliqueConfig, db ethdb.Database) *Clique {
	return &Clique{
		config: config,
		db:     db,
	}
}

func (c *Clique) SetProducers(data Producers) {
	c.producers = data
	c.schedulerID = data.SchedulerID
	data.store(c.db)
}

func (c *Clique) GetProducers(data GetProducers) Producers {
	return c.producers
}

// Authorize injects a private key into the consensus engine to mint new blocks with.
func (c *Clique) Authorize(signer common.Address, signFn SignerFn) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.signer = signer
	c.signFn = signFn
	// c.schedulerID = schedulerID
}

func (c *Clique) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (c *Clique) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return c.verifyHeader(chain, header, nil)
}

func (c *Clique) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			err := c.verifyHeader(chain, header, headers[:i])

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

func (c *Clique) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

func (c *Clique) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return c.verifySeal(chain, header, nil)
}

func (c *Clique) Prepare(chain consensus.ChainReader, header *types.Header) error {
	header.Difficulty = nil

	number := header.Number.Uint64()

	snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
	if number == 1 {
		c.producers = snap.Producers
	}
	if err != nil {
		return err
	}

	// if number > c.producers.Number+c.producers.Epoch {
	// 	return errUnknownBlock
	// }

	// if number < c.producers.Number {
	// 	return errUnknownBlock
	// }

	// header.Coinbase = c.producers.SequencerSet.GetProducer().Address

	// Ensure the extra data has all its components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	if number%c.config.Epoch == 0 {
		header.Extra = append(header.Extra, snap.Producers.serialize()...)
	}

	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	// Do not manipulate the timestamps when running with the bvm
	if !rcfg.UsingBVM {
		header.Time = parent.Time + c.config.Period
		if header.Time < uint64(time.Now().Unix()) {
			header.Time = uint64(time.Now().Unix())
		}
	}

	return nil
}

func (c *Clique) Finalize(
	chain consensus.ChainReader,
	header *types.Header,
	state *state.StateDB,
	txs []*types.Transaction,
	uncles []*types.Header,
) {
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

func (c *Clique) FinalizeAndAssemble(
	chain consensus.ChainReader,
	header *types.Header,
	state *state.StateDB,
	txs []*types.Transaction,
	uncles []*types.Header,
	receipts []*types.Receipt,
) (*types.Block, error) {
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

func (c *Clique) Seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()

	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	c.lock.RLock()
	signer, signFn := c.signer, c.signFn
	c.lock.RUnlock()

	if header.Number.Uint64() >= c.producers.Number+c.producers.Epoch {
		return errUnknownBlock
	}

	if header.Number.Uint64() < c.producers.Number {
		return errUnknownBlock
	}

	if c.signer != c.producers.SequencerSet.GetProducer().Address {
		return errUnauthorizedSigner
	}

	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	if rcfg.UsingBVM {
		delay = 0
	}
	// Sign all the things!
	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypeClique, CliqueRLP(header))
	if err != nil {
		return err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)
	// Wait until sealing is terminated or delay timeout.
	log.Trace("Waiting for slot to sign and propagate", "delay", common.PrettyDuration(delay))
	go func() {
		select {
		case <-stop:
			return
		case <-time.After(delay):
		}

		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header))
		}
	}()

	if number%c.config.Epoch == 0 {
		producers := deserialize(header.Extra[extraVanity : len(header.Extra)-extraSeal])
		c.producers = *producers
	}

	c.producers.SequencerSet.IncrementProducerPriority(1)

	return nil
}

func (c *Clique) SealHash(header *types.Header) common.Hash {
	return SealHash(header)
}

func (c *Clique) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	return nil
}

func (c *Clique) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{{
		Namespace: "clique",
		Version:   "1.0",
		Service:   &API{chain: chain, clique: c},
		Public:    false,
	}}
}

func (c *Clique) Close() error {
	return nil
}

func (c *Clique) verifyHeader(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	if header.Number == nil {
		return errUnknownBlock
	}

	if !rcfg.UsingBVM {
		// Don't waste time checking blocks from the future
		if header.Time > uint64(time.Now().Unix()) {
			return consensus.ErrFutureBlock
		}
	}

	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}

	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}
	// Ensure that the block's difficulty is empty
	if header.Difficulty != nil {
		return errInvalidDifficulty
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return c.verifyCascadingFields(chain, header, parents)
}

func (c *Clique) verifyCascadingFields(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to its parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}

	// Do not account for timestamps in consensus when running the bvm
	// changes. The timestamp must be montonic, meaning that it can be the same
	// or increase. L1 dictates the timestamp.
	if !rcfg.UsingBVM {
		if parent.Time+c.config.Period > header.Time {
			return errInvalidTimestamp
		}
	}

	return c.verifySeal(chain, header, parents)
}

// snapshot retrieves the authorization snapshot at a given point in time.
func (c *Clique) snapshot(chain consensus.ChainReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		headers []*types.Header
		snap    *Snapshot
	)
	for snap == nil {
		// If we're at the genesis, snapshot the initial state. Alternatively if we're
		// at a checkpoint block without a parent (light client CHT), or we have piled
		// up more headers than allowed to be reorged (chain reinit from a freezer),
		// consider the checkpoint trusted and snapshot it.
		if number == 0 || (number%c.config.Epoch == 0 && (len(headers) > params.ImmutabilityThreshold || chain.GetHeaderByNumber(number-1) == nil)) {
			checkpoint := chain.GetHeaderByNumber(number)
			if checkpoint != nil {
				hash := checkpoint.Hash()

				producers := deserialize(checkpoint.Extra[extraVanity : len(checkpoint.Extra)-extraSeal])

				snap = newSnapshot(c.config, c.signatures, number, hash, *producers)
				if err := snap.store(c.db); err != nil {
					return nil, err
				}
				log.Info("Stored checkpoint snapshot to disk", "number", number, "hash", hash)
				break
			}
		}
		// No snapshot for this header, gather the header and move backward
		// var header *types.Header
		// if len(parents) > 0 {
		// 	// If we have explicit parents, pick from there (enforced)
		// 	header = parents[len(parents)-1]
		// 	if header.Hash() != hash || header.Number.Uint64() != number {
		// 		return nil, consensus.ErrUnknownAncestor
		// 	}
		// 	parents = parents[:len(parents)-1]
		// } else {
		// 	// No explicit parents (or no more left), reach out to the database
		// 	header = chain.GetHeader(hash, number)
		// 	if header == nil {
		// 		return nil, consensus.ErrUnknownAncestor
		// 	}
		// }
		// headers = append(headers, header)
		// number, hash = number-1, header.ParentHash
	}

	// Previous snapshot found, apply any pending headers on top of it
	// for i := 0; i < len(headers)/2; i++ {
	// 	headers[i], headers[len(headers)-1-i] = headers[len(headers)-1-i], headers[i]
	// }
	// snap, err := snap.apply(headers)
	// if err != nil {
	// 	return nil, err
	// }

	// // If we've generated a new checkpoint snapshot, save to disk
	// if snap.Number%checkpointInterval == 0 && len(headers) > 0 {
	// 	if err = snap.store(c.db); err != nil {
	// 		return nil, err
	// 	}
	// 	log.Trace("Stored voting snapshot to disk", "number", snap.Number, "hash", snap.Hash)
	// }

	return snap, nil
}

func (c *Clique) verifySeal(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	// height < currentNumber: 							don't check
	// currentNumber <= height < currentNumber + epoch: check
	// height >= currentNumber + epoch: 				update current

	if header.Number.Uint64() >= c.producers.Number+c.producers.Epoch {
		return errUnknownBlock
	}

	if header.Number.Uint64() < c.producers.Number {
		return nil
	}

	signer, err := c.Author(header)
	if err != nil {
		return err
	}

	if signer != c.producers.SequencerSet.GetProducer().Address {
		return errUnauthorizedSigner
	}

	return nil
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header)
	hasher.Sum(hash[:0])
	return hash
}

func CliqueRLP(header *types.Header) []byte {
	b := new(bytes.Buffer)
	encodeSigHeader(b, header)
	return b.Bytes()
}

func encodeSigHeader(w io.Writer, header *types.Header) {
	err := rlp.Encode(w, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-crypto.SignatureLength], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	})
	if err != nil {
		panic("can't encode: " + err.Error())
	}
}

package coterie

import (
	"bytes"
	"errors"
	"io"
	"math/big"
	"sync"
	"time"

	"golang.org/x/crypto/sha3"

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

type Coterie struct {
	config      *params.CoterieConfig // Consensus engine configuration parameters
	db          ethdb.Database
	producers   Producers
	signer      common.Address
	signFn      SignerFn
	schedulerID string
	lock        sync.RWMutex
}

func New(config *params.CoterieConfig, db ethdb.Database) *Coterie {
	return &Coterie{
		config: config,
		db:     db,
	}
}

func (c *Coterie) SetProducers(data Producers) {
	c.producers = data
	c.schedulerID = data.SchedulerID
	data.store(c.db)
}

func (c *Coterie) GetProducers(data GetProducers) Producers {
	return c.producers
}

// Authorize injects a private key into the consensus engine to mint new blocks with.
func (c *Coterie) Authorize(signer common.Address, signFn SignerFn, schedulerID string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.signer = signer
	c.signFn = signFn
	c.schedulerID = schedulerID
}

func (c *Coterie) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (c *Coterie) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return c.verifyHeader(chain, header, nil)
}

func (c *Coterie) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
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

func (c *Coterie) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

func (c *Coterie) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return c.verifySeal(chain, header, nil)
}

func (c *Coterie) Prepare(chain consensus.ChainReader, header *types.Header) error {
	header.Difficulty = nil

	number := header.Number.Uint64()

	if header.Number.Uint64() >= c.producers.Number+c.producers.Epoch {
		return errUnknownBlock
	}

	if header.Number.Uint64() < c.producers.Number {
		return errUnknownBlock
	}

	header.Coinbase = c.producers.SequencerSet.GetProducer().Address

	// Ensure the extra data has all its components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

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

func (c *Coterie) Finalize(
	chain consensus.ChainReader,
	header *types.Header,
	state *state.StateDB,
	txs []*types.Transaction,
	uncles []*types.Header,
) {
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

func (c *Coterie) FinalizeAndAssemble(
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

func (c *Coterie) Seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
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
	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypeCoterie, CoterieRLP(header))
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

	c.producers.SequencerSet.IncrementProducerPriority(1)

	return nil

}

func (c *Coterie) SealHash(header *types.Header) common.Hash {
	return SealHash(header)
}

func (c *Coterie) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	return nil
}

func (c *Coterie) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{{
		Namespace: "coterie",
		Version:   "1.0",
		Service:   &API{chain: chain, coterie: c},
		Public:    false,
	}}
}

func (c *Coterie) Close() error {
	return nil
}

func (c *Coterie) verifyHeader(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
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

func (c *Coterie) verifyCascadingFields(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
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

func (c *Coterie) verifySeal(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
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

func CoterieRLP(header *types.Header) []byte {
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

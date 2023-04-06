package txmgr_test

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mantlenetworkio/mantle/mt-batcher/txmgr"
)

type testHarness struct {
	cfg       txmgr.Config
	mgr       txmgr.TxManager
	backend   *mockBackend
	gasPricer *gasPricer
}

func newTestHarnessWithConfig(cfg txmgr.Config) *testHarness {
	backend := newMockBackend()
	mgr := txmgr.NewSimpleTxManager(cfg, backend)

	return &testHarness{
		cfg:       cfg,
		mgr:       mgr,
		backend:   backend,
		gasPricer: newGasPricer(3),
	}
}

func newTestHarness() *testHarness {
	return newTestHarnessWithConfig(configWithNumConfs(1))
}

func configWithNumConfs(numConfirmations uint64) txmgr.Config {
	return txmgr.Config{
		ResubmissionTimeout:       time.Second,
		ReceiptQueryInterval:      50 * time.Millisecond,
		NumConfirmations:          numConfirmations,
		SafeAbortNonceTooLowCount: 3,
	}
}

type gasPricer struct {
	epoch         int64
	mineAtEpoch   int64
	baseGasTipFee *big.Int
	baseBaseFee   *big.Int
	mu            sync.Mutex
}

func newGasPricer(mineAtEpoch int64) *gasPricer {
	return &gasPricer{
		mineAtEpoch:   mineAtEpoch,
		baseGasTipFee: big.NewInt(5),
		baseBaseFee:   big.NewInt(7),
	}
}

func (g *gasPricer) expGasFeeCap() *big.Int {
	_, gasFeeCap := g.feesForEpoch(g.mineAtEpoch)
	return gasFeeCap
}

func (g *gasPricer) shouldMine(gasFeeCap *big.Int) bool {
	return g.expGasFeeCap().Cmp(gasFeeCap) == 0
}

func (g *gasPricer) feesForEpoch(epoch int64) (*big.Int, *big.Int) {
	epochBaseFee := new(big.Int).Mul(g.baseBaseFee, big.NewInt(epoch))
	epochGasTipCap := new(big.Int).Mul(g.baseGasTipFee, big.NewInt(epoch))
	epochGasFeeCap := txmgr.CalcGasFeeCap(epochBaseFee, epochGasTipCap)

	return epochGasTipCap, epochGasFeeCap
}

func (g *gasPricer) sample() (*big.Int, *big.Int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.epoch++
	epochGasTipCap, epochGasFeeCap := g.feesForEpoch(g.epoch)

	return epochGasTipCap, epochGasFeeCap
}

type minedTxInfo struct {
	gasFeeCap   *big.Int
	blockNumber uint64
}

type mockBackend struct {
	mu sync.RWMutex

	blockHeight uint64

	minedTxs map[common.Hash]minedTxInfo
}

func newMockBackend() *mockBackend {
	return &mockBackend{
		minedTxs: make(map[common.Hash]minedTxInfo),
	}
}

func (b *mockBackend) mine(txHash *common.Hash, gasFeeCap *big.Int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.blockHeight++
	if txHash != nil {
		b.minedTxs[*txHash] = minedTxInfo{
			gasFeeCap:   gasFeeCap,
			blockNumber: b.blockHeight,
		}
	}
}

func (b *mockBackend) BlockNumber(ctx context.Context) (uint64, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.blockHeight, nil
}

func (b *mockBackend) TransactionReceipt(
	ctx context.Context,
	txHash common.Hash,
) (*types.Receipt, error) {

	b.mu.RLock()
	defer b.mu.RUnlock()

	txInfo, ok := b.minedTxs[txHash]
	if !ok {
		return nil, nil
	}

	return &types.Receipt{
		TxHash:      txHash,
		GasUsed:     txInfo.gasFeeCap.Uint64(),
		BlockNumber: big.NewInt(int64(txInfo.blockNumber)),
	}, nil
}

func TestTxMgrConfirmAtMinGasPrice(t *testing.T) {
	t.Parallel()

	h := newTestHarness()

	gasPricer := newGasPricer(1)

	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		gasTipCap, gasFeeCap := gasPricer.sample()
		return types.NewTx(&types.DynamicFeeTx{
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
		}), nil
	}

	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		if gasPricer.shouldMine(tx.GasFeeCap()) {
			txHash := tx.Hash()
			h.backend.mine(&txHash, tx.GasFeeCap())
		}
		return nil
	}

	ctx := context.Background()
	receipt, err := h.mgr.Send(ctx, updateGasPrice, sendTx)
	require.Nil(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, gasPricer.expGasFeeCap().Uint64(), receipt.GasUsed)
}

func TestTxMgrNeverConfirmCancel(t *testing.T) {
	t.Parallel()

	h := newTestHarness()

	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		gasTipCap, gasFeeCap := h.gasPricer.sample()
		return types.NewTx(&types.DynamicFeeTx{
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
		}), nil
	}

	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		// Don't publish tx to backend, simulating never being mined.
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	receipt, err := h.mgr.Send(ctx, updateGasPrice, sendTx)
	require.Equal(t, err, context.DeadlineExceeded)
	require.Nil(t, receipt)
}

func TestTxMgrConfirmsAtHigherGasPrice(t *testing.T) {
	t.Parallel()

	h := newTestHarness()

	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		gasTipCap, gasFeeCap := h.gasPricer.sample()
		return types.NewTx(&types.DynamicFeeTx{
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
		}), nil
	}

	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		if h.gasPricer.shouldMine(tx.GasFeeCap()) {
			txHash := tx.Hash()
			h.backend.mine(&txHash, tx.GasFeeCap())
		}
		return nil
	}

	ctx := context.Background()
	receipt, err := h.mgr.Send(ctx, updateGasPrice, sendTx)
	require.Nil(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, h.gasPricer.expGasFeeCap().Uint64(), receipt.GasUsed)
}

var errRpcFailure = errors.New("rpc failure")

func TestTxMgrBlocksOnFailingRpcCalls(t *testing.T) {
	t.Parallel()

	h := newTestHarness()

	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		gasTipCap, gasFeeCap := h.gasPricer.sample()
		return types.NewTx(&types.DynamicFeeTx{
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
		}), nil
	}

	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		return errRpcFailure
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	receipt, err := h.mgr.Send(ctx, updateGasPrice, sendTx)
	require.Equal(t, err, context.DeadlineExceeded)
	require.Nil(t, receipt)
}

func TestTxMgrOnlyOnePublicationSucceeds(t *testing.T) {
	t.Parallel()

	h := newTestHarness()

	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		gasTipCap, gasFeeCap := h.gasPricer.sample()
		return types.NewTx(&types.DynamicFeeTx{
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
		}), nil
	}

	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		// Fail all but the final attempt.
		if !h.gasPricer.shouldMine(tx.GasFeeCap()) {
			return errRpcFailure
		}

		txHash := tx.Hash()
		h.backend.mine(&txHash, tx.GasFeeCap())
		return nil
	}

	ctx := context.Background()
	receipt, err := h.mgr.Send(ctx, updateGasPrice, sendTx)
	require.Nil(t, err)

	require.NotNil(t, receipt)
	require.Equal(t, h.gasPricer.expGasFeeCap().Uint64(), receipt.GasUsed)
}

func TestTxMgrConfirmsMinGasPriceAfterBumping(t *testing.T) {
	t.Parallel()

	h := newTestHarness()

	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		gasTipCap, gasFeeCap := h.gasPricer.sample()
		return types.NewTx(&types.DynamicFeeTx{
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
		}), nil
	}

	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		// Delay mining the tx with the min gas price.
		if h.gasPricer.shouldMine(tx.GasFeeCap()) {
			time.AfterFunc(5*time.Second, func() {
				txHash := tx.Hash()
				h.backend.mine(&txHash, tx.GasFeeCap())
			})
		}
		return nil
	}

	ctx := context.Background()
	receipt, err := h.mgr.Send(ctx, updateGasPrice, sendTx)
	require.Nil(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, h.gasPricer.expGasFeeCap().Uint64(), receipt.GasUsed)
}

func TestTxMgrDoesntAbortNonceTooLowAfterMiningTx(t *testing.T) {
	t.Parallel()

	h := newTestHarnessWithConfig(configWithNumConfs(2))

	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		gasTipCap, gasFeeCap := h.gasPricer.sample()
		return types.NewTx(&types.DynamicFeeTx{
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
		}), nil
	}

	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		switch {

		case tx.GasFeeCap().Cmp(h.gasPricer.expGasFeeCap()) < 0:
			return nil

		case h.gasPricer.shouldMine(tx.GasFeeCap()):
			txHash := tx.Hash()
			h.backend.mine(&txHash, tx.GasFeeCap())
			time.AfterFunc(5*time.Second, func() {
				h.backend.mine(nil, nil)
			})
			return nil

		default:
			return core.ErrNonceTooLow
		}
	}

	ctx := context.Background()
	receipt, err := h.mgr.Send(ctx, updateGasPrice, sendTx)
	require.Nil(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, h.gasPricer.expGasFeeCap().Uint64(), receipt.GasUsed)
}

func TestWaitMinedReturnsReceiptOnFirstSuccess(t *testing.T) {
	t.Parallel()

	h := newTestHarness()

	// Create a tx and mine it immediately using the default backend.
	tx := types.NewTx(&types.LegacyTx{})
	txHash := tx.Hash()
	h.backend.mine(&txHash, new(big.Int))

	ctx := context.Background()
	receipt, err := txmgr.WaitMined(ctx, h.backend, tx, 50*time.Millisecond, 1)
	require.Nil(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, receipt.TxHash, txHash)
}

func TestWaitMinedCanBeCanceled(t *testing.T) {
	t.Parallel()

	h := newTestHarness()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create an unimined tx.
	tx := types.NewTx(&types.LegacyTx{})

	receipt, err := txmgr.WaitMined(ctx, h.backend, tx, 50*time.Millisecond, 1)
	require.Equal(t, err, context.DeadlineExceeded)
	require.Nil(t, receipt)
}

func TestWaitMinedMultipleConfs(t *testing.T) {
	t.Parallel()

	const numConfs = 2

	h := newTestHarnessWithConfig(configWithNumConfs(numConfs))
	ctxt, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create an unimined tx.
	tx := types.NewTx(&types.LegacyTx{})
	txHash := tx.Hash()
	h.backend.mine(&txHash, new(big.Int))

	receipt, err := txmgr.WaitMined(ctxt, h.backend, tx, 50*time.Millisecond, numConfs)
	require.Equal(t, err, context.DeadlineExceeded)
	require.Nil(t, receipt)

	ctxt, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Mine an empty block, tx should now be confirmed.
	h.backend.mine(nil, nil)
	receipt, err = txmgr.WaitMined(ctxt, h.backend, tx, 50*time.Millisecond, numConfs)
	require.Nil(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, txHash, receipt.TxHash)
}

func TestManagerPanicOnZeroConfs(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("NewSimpleTxManager should panic when using zero conf")
		}
	}()

	_ = newTestHarnessWithConfig(configWithNumConfs(0))
}

type failingBackend struct {
	returnSuccessBlockNumber bool
	returnSuccessReceipt     bool
}

func (b *failingBackend) BlockNumber(ctx context.Context) (uint64, error) {
	if !b.returnSuccessBlockNumber {
		b.returnSuccessBlockNumber = true
		return 0, errRpcFailure
	}

	return 1, nil
}

func (b *failingBackend) TransactionReceipt(
	ctx context.Context, txHash common.Hash) (*types.Receipt, error) {

	if !b.returnSuccessReceipt {
		b.returnSuccessReceipt = true
		return nil, errRpcFailure
	}

	return &types.Receipt{
		TxHash:      txHash,
		BlockNumber: big.NewInt(1),
	}, nil
}

func TestWaitMinedReturnsReceiptAfterFailure(t *testing.T) {
	t.Parallel()

	var borkedBackend failingBackend

	tx := types.NewTx(&types.LegacyTx{})
	txHash := tx.Hash()

	ctx := context.Background()
	receipt, err := txmgr.WaitMined(ctx, &borkedBackend, tx, 50*time.Millisecond, 1)
	require.Nil(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, receipt.TxHash, txHash)
}

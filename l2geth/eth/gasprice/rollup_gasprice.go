package gasprice

import (
	"context"
	"math/big"
	"sync"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/rollup/fees"
)

// RollupOracle holds the L1 and L2 gas prices for fee calculation
type RollupOracle struct {
	l1GasPrice     *big.Int
	l2GasPrice     *big.Int
	daGasPrice     *big.Int
	overhead       *big.Int
	scalar         *big.Float
	isBurning      *big.Int
	charge         *big.Int
	sccAddress     common.Address
	l1GasPriceLock sync.RWMutex
	l2GasPriceLock sync.RWMutex
	daGasPriceLock sync.RWMutex
	overheadLock   sync.RWMutex
	scalarLock     sync.RWMutex
	isBurningLock  sync.RWMutex
	chargeLock     sync.RWMutex
	sccAddressLock sync.RWMutex
}

// NewRollupOracle returns an initialized RollupOracle
func NewRollupOracle() *RollupOracle {
	return &RollupOracle{
		l1GasPrice: new(big.Int),
		l2GasPrice: new(big.Int),
		overhead:   new(big.Int),
		scalar:     new(big.Float),
		isBurning:  new(big.Int),
		charge:     new(big.Int),
	}
}

// SuggestL1GasPrice returns the gas price which should be charged per byte of published
// data by the sequencer.
func (gpo *RollupOracle) SuggestL1GasPrice(ctx context.Context) (*big.Int, error) {
	gpo.l1GasPriceLock.RLock()
	defer gpo.l1GasPriceLock.RUnlock()
	return gpo.l1GasPrice, nil
}

// SetL1GasPrice returns the current L1 gas price
func (gpo *RollupOracle) SetL1GasPrice(gasPrice *big.Int) error {
	gpo.l1GasPriceLock.Lock()
	defer gpo.l1GasPriceLock.Unlock()
	gpo.l1GasPrice = gasPrice
	log.Info("Set L1 Gas Price", "gasprice", gpo.l1GasPrice)
	return nil
}

// SuggestL2GasPrice returns the gas price which should be charged per unit of gas
// set manually by the sequencer depending on congestion
func (gpo *RollupOracle) SuggestL2GasPrice(ctx context.Context) (*big.Int, error) {
	gpo.l2GasPriceLock.RLock()
	defer gpo.l2GasPriceLock.RUnlock()
	return gpo.l2GasPrice, nil
}

// SetL2GasPrice returns the current L2 gas price
func (gpo *RollupOracle) SetL2GasPrice(gasPrice *big.Int) error {
	gpo.l2GasPriceLock.Lock()
	defer gpo.l2GasPriceLock.Unlock()
	gpo.l2GasPrice = gasPrice
	log.Info("Set L2 Gas Price", "gasprice", gpo.l2GasPrice)
	return nil
}

// SuggestDAGasPrice returns the gas price which should be charged per byte of published
// data by the sequencer.
func (gpo *RollupOracle) SuggestDAGasPrice(ctx context.Context) (*big.Int, error) {
	gpo.daGasPriceLock.RLock()
	defer gpo.daGasPriceLock.RUnlock()
	return gpo.daGasPrice, nil
}

// SetDAGasPrice returns the current DA gas price
func (gpo *RollupOracle) SetDAGasPrice(daGasPrice *big.Int) error {
	gpo.daGasPriceLock.Lock()
	defer gpo.daGasPriceLock.Unlock()
	gpo.daGasPrice = daGasPrice
	log.Info("Set DA Gas Price", "daGasprice", gpo.daGasPrice)
	return nil
}

// SuggestOverhead returns the cached overhead value from the
// BVM_GasPriceOracle
func (gpo *RollupOracle) SuggestOverhead(ctx context.Context) (*big.Int, error) {
	gpo.overheadLock.RLock()
	defer gpo.overheadLock.RUnlock()
	return gpo.overhead, nil
}

// SetOverhead caches the overhead value that is set in the
// BVM_GasPriceOracle
func (gpo *RollupOracle) SetOverhead(overhead *big.Int) error {
	gpo.overheadLock.Lock()
	defer gpo.overheadLock.Unlock()
	gpo.overhead = overhead
	log.Info("Set batch overhead", "overhead", overhead)
	return nil
}

// SuggestScalar returns the cached scalar value
func (gpo *RollupOracle) SuggestScalar(ctx context.Context) (*big.Float, error) {
	gpo.scalarLock.RLock()
	defer gpo.scalarLock.RUnlock()
	return gpo.scalar, nil
}

// SetScalar sets the scalar value held in the BVM_GasPriceOracle
func (gpo *RollupOracle) SetScalar(scalar *big.Int, decimals *big.Int) error {
	gpo.scalarLock.Lock()
	defer gpo.scalarLock.Unlock()
	value := fees.ScaleDecimals(scalar, decimals)
	gpo.scalar = value
	log.Info("Set scalar", "scalar", gpo.scalar)
	return nil
}

// SetIsBurning sets the isBurning value held in the BVM_GasPriceOracle
func (gpo *RollupOracle) SetIsBurning(isBurning *big.Int) error {
	gpo.isBurningLock.Lock()
	defer gpo.isBurningLock.Unlock()
	gpo.isBurning = isBurning
	log.Info("Set isBurning", "isBurning", isBurning)
	return nil
}

// SetCharge sets the charge value held in the BVM_GasPriceOracle
func (gpo *RollupOracle) SetCharge(charge *big.Int) error {
	gpo.chargeLock.Lock()
	defer gpo.chargeLock.Unlock()
	gpo.charge = charge
	log.Info("Set charge", "charge", charge)
	return nil
}

// SetCharge sets the charge value held in the BVM_GasPriceOracle
func (gpo *RollupOracle) SetSCCAddress(sccAddress common.Address) error {
	gpo.sccAddressLock.Lock()
	defer gpo.sccAddressLock.Unlock()
	gpo.sccAddress = sccAddress
	log.Info("Set sccAddress", "sccAddress", sccAddress.Hex())
	return nil
}

// SCCAddress returns the cached SCCAddress value
func (gpo *RollupOracle) SCCAddress() common.Address {
	gpo.sccAddressLock.RLock()
	defer gpo.sccAddressLock.RUnlock()
	return gpo.sccAddress
}

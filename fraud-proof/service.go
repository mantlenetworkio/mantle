package rollup

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/mantlenetworkio/mantle/fraud-proof/rollup/services"
	"github.com/mantlenetworkio/mantle/fraud-proof/rollup/services/sequencer"
	"github.com/mantlenetworkio/mantle/fraud-proof/rollup/services/validator"
	"github.com/mantlenetworkio/mantle/l2geth/accounts"
	"github.com/mantlenetworkio/mantle/l2geth/accounts/keystore"
	"github.com/mantlenetworkio/mantle/l2geth/eth"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/node"
)

// RegisterFraudProofService registers rollup service configured by ctx
// Either a sequncer service or a validator service will be registered
func RegisterFraudProofService(stack *node.Node, cfg *services.Config) {
	// Unlock account for L1 transaction signer
	var ks *keystore.KeyStore
	if keystores := stack.AccountManager().Backends(keystore.KeyStoreType); len(keystores) > 0 {
		ks = keystores[0].(*keystore.KeyStore)
	}
	if ks == nil {
		log.Crit("Failed to register the Rollup service: keystore not found")
	}
	chainID := big.NewInt(int64(cfg.L1ChainID))
	json, err := ks.Export(accounts.Account{Address: cfg.StakeAddr}, cfg.Passphrase, cfg.Passphrase)
	if err != nil {
		log.Crit("Failed to register the Rollup service", "err", err)
	}
	auth, err := bind.NewTransactorWithChainID(bytes.NewReader(json), cfg.Passphrase, chainID)
	if err != nil {
		log.Crit("Failed to register the Rollup service", "err", err)
	}

	var ethService *eth.Ethereum
	if err := stack.Service(&ethService); err != nil {
		log.Crit("Failed to retrieve eth service backend", "err", err)
	}

	// Register services
	log.Info("Print log type", "cfg.Node", cfg.Node)
	if cfg.Node == services.NODE_SCHEDULER {
		log.Info("Start NODE_SCHEDULER...")
		sequencer.RegisterService(ethService, ethService.APIBackend, cfg, auth)
		log.Info("Start NODE_SCHEDULER end...")
	} else if cfg.Node == services.NODE_VERIFIER {
		log.Info("Start NODE_VERIFIER...")
		validator.RegisterService(ethService, ethService.APIBackend, cfg, auth)
		log.Info("Start NODE_VERIFIER end...")
	} else {
		log.Crit("Failed to register the Rollup service: Node type unkown", "type", cfg.Node)
	}
}

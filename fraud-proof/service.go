package rollup

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/mantlenetworkio/mantle/fraud-proof/proof"
	"github.com/mantlenetworkio/mantle/fraud-proof/rollup/services"
	"github.com/mantlenetworkio/mantle/l2geth/accounts"
	"github.com/mantlenetworkio/mantle/l2geth/accounts/keystore"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/node"
)

// RegisterRollupService registers rollup service configured by ctx
// Either a sequncer service or a validator service will be registered
func RegisterRollupService(stack *node.Node, eth services.Backend, proofBackend proof.Backend, cfg *services.Config) {
	// Unlock account for L1 transaction signer
	var ks *keystore.KeyStore
	if keystores := stack.AccountManager().Backends(keystore.KeyStoreType); len(keystores) > 0 {
		ks = keystores[0].(*keystore.KeyStore)
	}
	if ks == nil {
		log.Crit("Failed to register the Rollup service: keystore not found")
	}
	chainID := big.NewInt(int64(cfg.L1ChainID))
	json, err := ks.Export(accounts.Account{Address: cfg.Coinbase}, cfg.Passphrase, "")
	if err != nil {
		log.Crit("Failed to register the Rollup service", "err", err)
	}
	auth, err := bind.NewTransactorWithChainID(bytes.NewReader(json), cfg.Passphrase, chainID)
	if err != nil {
		log.Crit("Failed to register the Rollup service", "err", err)
	}

	// Register services
	if cfg.Node == services.NODE_SEQUENCER {
		fmt.Println(auth)
		// TODO-FIXME new service and start
		//sequencer.RegisterService(stack, eth, proofBackend, cfg, auth)
	} else if cfg.Node == services.NODE_VALIDATOR {
		fmt.Println(auth)
		// TODO-FIXME new service and start
		//validator.RegisterService(stack, eth, proofBackend, cfg, auth)
	} else {
		log.Crit("Failed to register the Rollup service: Node type unkown", "type", cfg.Node)
	}
}

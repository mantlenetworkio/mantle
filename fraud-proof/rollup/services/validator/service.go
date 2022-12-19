package validator

import (
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/specularl2/specular/clients/geth/specular/proof"
	"github.com/specularl2/specular/clients/geth/specular/rollup/services"
)

func RegisterService(stack *node.Node, eth services.Backend, proofBackend proof.Backend, cfg *services.Config, auth *bind.TransactOpts) {
	var validator node.Lifecycle
	var err error
	if cfg.Node == services.NODE_VALIDATOR {
		validator, err = New(eth, proofBackend, cfg, auth)
	}
	if err != nil {
		log.Crit("Failed to register the Rollup service", "err", err)
	}
	stack.RegisterLifecycle(validator)
	// stack.RegisterAPIs(seq.APIs())
	log.Info("Validator registered")
}

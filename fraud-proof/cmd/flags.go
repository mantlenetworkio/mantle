package cmd

import (
	"github.com/mantlenetworkio/mantle/fraud-proof/rollup/services"
	"github.com/mantlenetworkio/mantle/l2geth/cmd/utils"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"gopkg.in/urfave/cli.v1"
)

var (
	// Fraud Proofs Flags
	RollupNodeFlag = &cli.StringFlag{
		Name:  "rollup.node",
		Usage: "Start node as rollup sequencer or validator",
		Value: "",
	}
	RollupCoinBaseFlag = &cli.StringFlag{
		Name:  "rollup.coinbase",
		Usage: "The sequencer/validator address to be unlocked (pass passphrash via --password)",
		Value: "",
	}
	RollupL1EndpointFlag = &cli.StringFlag{
		Name:  "rollup.l1endpoint",
		Usage: "The api endpoint of L1 client",
		Value: "",
	}
	RollupL1ChainIDFlag = &cli.Uint64Flag{
		Name:  "rollup.l1chainid",
		Usage: "The chain ID of L1 client",
		Value: 31337,
	}
	RollupSequencerAddrFlag = &cli.StringFlag{
		Name:  "rollup.sequencer-addr",
		Usage: "The account address of sequencer",
		Value: "",
	}
	RollupRollupAddrFlag = &cli.StringFlag{
		Name:  "rollup.rollup-addr",
		Usage: "The contract address of L1 rollup",
		Value: "",
	}
	RollupRollupStakeAmount = &cli.Uint64Flag{
		Name:   "rollup.stake-amount",
		Usage:  "Required staking amount",
		EnvVar: "ROLLUP_STAKE_AMOUNT",
		Value:  1000000000000000000,
	}
)

//// RegisterEthService adds an Ethereum client to the stack.
//// The second return value is the full node instance, which may be nil if the
//// node is running as a light client.
//func RegisterEthService(stack *node.Node, cfg *ethconfig.Config) (ethapi.Backend, *eth.Ethereum) {
//	backend, err := eth.New(stack, cfg)
//	if err != nil {
//		utils.Fatalf("Failed to register the Ethereum service: %v", err)
//	}
//	//if cfg.LightServ > 0 {
//	//	_, err := les.NewLesServer(stack, backend, cfg)
//	//	if err != nil {
//	//		utils.Fatalf("Failed to create the LES server: %v", err)
//	//	}
//	//}
//	if err := ethcatalyst.Register(stack, backend); err != nil {
//		utils.Fatalf("Failed to register the Engine API service: %v", err)
//	}
//	stack.RegisterAPIs(tracers.APIs(backend.APIBackend))
//	// <specular modification>
//	stack.RegisterAPIs(proof.APIs(backend.APIBackend))
//	// <specular modification/>
//	return backend.APIBackend, backend
//}

func MakeFraudProofConfig(ctx *cli.Context) *services.Config {
	utils.CheckExclusive(ctx, RollupNodeFlag, utils.MiningEnabledFlag)
	utils.CheckExclusive(ctx, RollupNodeFlag, utils.DeveloperFlag)
	var passphrase string
	if list := utils.MakePasswordList(ctx); len(list) > 0 {
		passphrase = list[0]
	} else {
		utils.Fatalf("Failed to register the Rollup service: coinbase account locked")
	}
	cfg := &services.Config{
		Node:          ctx.String(RollupNodeFlag.Name),
		Coinbase:      common.HexToAddress(ctx.String(RollupCoinBaseFlag.Name)),
		Passphrase:    passphrase,
		L1Endpoint:    ctx.String(RollupL1EndpointFlag.Name),
		L1ChainID:     ctx.Uint64(RollupL1ChainIDFlag.Name),
		SequencerAddr: common.HexToAddress(ctx.String(RollupSequencerAddrFlag.Name)),
		//SequencerInboxAddr: common.HexToAddress(ctx.String(RollupSequencerInboxAddrFlag.Name)),
		RollupAddr:        common.HexToAddress(ctx.String(RollupRollupAddrFlag.Name)),
		RollupStakeAmount: ctx.Uint64(RollupRollupStakeAmount.Name),
	}
	return cfg
}

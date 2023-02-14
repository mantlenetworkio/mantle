package cmd

import (
	"github.com/mantlenetworkio/mantle/fraud-proof/rollup/services"
	"github.com/mantlenetworkio/mantle/l2geth/cmd/utils"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"gopkg.in/urfave/cli.v1"
)

var (
	// Fraud Proofs Flags
	FraudProofNodeFlag = &cli.StringFlag{
		Name:  "fp.node",
		Usage: "Start node as rollup sequencer or validator",
		Value: "",
	}
	FraudProofCoinBaseFlag = &cli.StringFlag{
		Name:  "fp.coinbase",
		Usage: "The sequencer/validator address to be unlocked (pass passphrash via --password)",
		Value: "",
	}
	FraudProofL1EndpointFlag = &cli.StringFlag{
		Name:  "fp.l1endpoint",
		Usage: "The api endpoint of L1 client",
		Value: "",
	}
	FraudProofL1ChainIDFlag = &cli.Uint64Flag{
		Name:  "fp.l1chainid",
		Usage: "The chain ID of L1 client",
		Value: 31337,
	}
	FraudProofSequencerAddrFlag = &cli.StringFlag{
		Name:  "fp.sequencer-addr",
		Usage: "The account address of sequencer",
		Value: "",
	}
	FraudProofRollupAddrFlag = &cli.StringFlag{
		Name:  "fp.rollup-addr",
		Usage: "The contract address of L1 rollup",
		Value: "",
	}
	FraudProofRollupStakeAmount = &cli.Uint64Flag{
		Name:   "fp.stake-amount",
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
	//utils.CheckExclusive(ctx, FraudProofNodeFlag, utils.MiningEnabledFlag)
	//utils.CheckExclusive(ctx, FraudProofNodeFlag, utils.DeveloperFlag)
	var passphrase string
	if list := utils.MakePasswordList(ctx); len(list) > 0 {
		passphrase = list[0]
	} else {
		utils.Fatalf("Failed to register the Rollup service: coinbase account locked")
	}
	cfg := &services.Config{
		Node:              ctx.String(utils.RollupRoleFlag.Name),
		Coinbase:          common.HexToAddress(ctx.String(FraudProofCoinBaseFlag.Name)),
		Passphrase:        passphrase,
		L1Endpoint:        ctx.String(FraudProofL1EndpointFlag.Name),
		L1ChainID:         ctx.Uint64(FraudProofL1ChainIDFlag.Name),
		SequencerAddr:     common.HexToAddress(ctx.String(FraudProofSequencerAddrFlag.Name)),
		RollupAddr:        common.HexToAddress(ctx.String(FraudProofRollupAddrFlag.Name)),
		RollupStakeAmount: ctx.Uint64(FraudProofRollupStakeAmount.Name),
	}
	return cfg
}

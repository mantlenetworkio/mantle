package challenger

import (
	"context"
	"github.com/Layr-Labs/datalayr/common/logging"
	ethc "github.com/ethereum/go-ethereum/common"
	bsscore "github.com/mantlenetworkio/mantle/bss-core"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/l1l2client"
	"github.com/mantlenetworkio/mantle/mt-challenger/challenger"
	"github.com/urfave/cli"
	"time"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(cliCtx *cli.Context) error {
		cfg, err := NewConfig(cliCtx)
		if err != nil {
			return err
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		logger, err := logging.GetLogger(cfg.LoggingConfig)
		if err != nil {
			return err
		}
		sequencerPrivKey, _, err := bsscore.ParseWalletPrivKeyAndContractAddr(
			"MtChallenger", cfg.Mnemonic, cfg.SequencerHDPath,
			cfg.PrivateKey, cfg.EigenContractAddress,
		)
		if err != nil {
			return err
		}
		l1Client, err := l1l2client.L1EthClientWithTimeout(ctx, cfg.L1EthRpc, cfg.DisableHTTP2)
		if err != nil {
			return err
		}
		logger.Info().Msg("MtChallenger l1Client init success")
		l2Client, err := l1l2client.DialL2EthClientWithTimeout(ctx, cfg.L2MtlRpc, cfg.DisableHTTP2)
		if err != nil {
			return err
		}
		logger.Info().Msg("MtChallenger l2Client init success")
		timeout, err := time.ParseDuration("12s")
		if err != nil {
			logger.Fatal().Err(err).Msg("MtChallenger improper timeout from config")
		}
		challengerConfig := &challenger.ChallengerConfig{
			L1Client:          l1Client,
			L2Client:          l2Client,
			EigenContractAddr: ethc.Address(common.HexToAddress(cfg.EigenContractAddress)),
			Logger:            logger,
			PrivKey:           sequencerPrivKey,
			GraphProvider:     cfg.GraphProvider,
			RetrieverSocket:   cfg.RetrieverSocket,
			KzgConfig:         cfg.KzgConfig,
			LastStoreNumber:   cfg.FromStoreNumber,
			Timeout:           timeout,
		}
		cLager, err := challenger.NewChallenger(ctx, challengerConfig)
		if err != nil {
			return err
		}
		if err := cLager.Start(); err != nil {
			return err
		}
		logger.Info().Msg("MtChallenger da challenger service start")
		defer cLager.Stop()
		return nil
	}
}

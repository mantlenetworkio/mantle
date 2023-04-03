package challenger

import (
	"context"
	"github.com/Layr-Labs/datalayr/common/logging"
	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/l1l2client"
	common2 "github.com/mantlenetworkio/mantle/mt-batcher/services/common"
	"github.com/mantlenetworkio/mantle/mt-challenger/challenger"
	"github.com/urfave/cli"
	"math/big"
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
		challengerPrivKey, _, err := common2.ParseWalletPrivKeyAndContractAddr(
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
		log.Info("MtChallenger l1Client init success")
		l2Client, err := l1l2client.DialL2EthClientWithTimeout(ctx, cfg.L2MtlRpc, cfg.DisableHTTP2)
		if err != nil {
			return err
		}
		log.Info("MtChallenger l2Client init success")
		timeout, err := time.ParseDuration("12s")
		if err != nil {
			log.Error("MtChallenger improper timeout from config", "err", err)
		}
		chainID, err := l1Client.ChainID(ctx)
		if err != nil {
			return err
		}

		signer := func(chainID *big.Int) challenger.SignerFn {
			s := common2.PrivateKeySignerFn(challengerPrivKey, chainID)
			return func(_ context.Context, addr ethc.Address, tx *types.Transaction) (*types.Transaction, error) {
				return s(addr, tx)
			}
		}
		challengerConfig := &challenger.ChallengerConfig{
			L1Client:                  l1Client,
			L2Client:                  l2Client,
			EigenContractAddr:         ethc.Address(common.HexToAddress(cfg.EigenContractAddress)),
			Logger:                    logger,
			PrivKey:                   challengerPrivKey,
			GraphProvider:             cfg.GraphProvider,
			RetrieverSocket:           cfg.RetrieverSocket,
			KzgConfig:                 cfg.KzgConfig,
			LastStoreNumber:           cfg.FromStoreNumber,
			Timeout:                   timeout,
			PollInterval:              cfg.PollInterval,
			DbPath:                    cfg.DbPath,
			CheckerBatchIndex:         cfg.CheckerBatchIndex,
			NeedReRollupBatch:         cfg.NeedReRollupBatch,
			ReRollupToolEnable:        cfg.ReRollupToolEnable,
			SignerFn:                  signer(chainID),
			ResubmissionTimeout:       cfg.ResubmissionTimeout,
			NumConfirmations:          cfg.NumConfirmations,
			SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
		}
		cLager, err := challenger.NewChallenger(ctx, challengerConfig)
		if err != nil {
			return err
		}
		if err := cLager.Start(); err != nil {
			return err
		}
		log.Info("MtChallenger da challenger service start")
		defer cLager.Stop()
		return nil
	}
}

package mt_batcher

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	sentry "github.com/getsentry/sentry-go"
	bsscore "github.com/mantlenetworkio/mantle/bss-core"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/l1l2client"
	"github.com/mantlenetworkio/mantle/mt-batcher/sequencer"
	"github.com/urfave/cli"
	"math/big"
	"strings"
	"time"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(cliCtx *cli.Context) error {
		cfg, err := NewConfig(cliCtx)
		if err != nil {
			return err
		}
		log.Info("Config parsed",
			"disperser", cfg.DisperserEndpoint,
			"mtlrpc", cfg.L2MtlRpc, "gitVersion", gitVersion)

		if cfg.SentryEnable {
			defer sentry.Flush(2 * time.Second)
		}
		log.Info("Initializing mantel da batch submitter")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		log.Info("Config LogLevel", "LogLevel", cfg.LogLevel)

		sequencerPrivKey, _, err := bsscore.ParseWalletPrivKeyAndContractAddr(
			"DaSequencer", cfg.Mnemonic, cfg.SequencerHDPath,
			cfg.PrivateKey, cfg.EigenContractAddress,
		)
		if err != nil {
			return err
		}
		l1Client, err := l1l2client.L1EthClientWithTimeout(ctx, cfg.L1EthRpc, cfg.DisableHTTP2)
		if err != nil {
			return err
		}
		log.Info("l1Client init success")
		chainID, err := l1Client.ChainID(ctx)
		if err != nil {
			return err
		}

		l2Client, err := l1l2client.DialL2EthClientWithTimeout(ctx, cfg.L2MtlRpc, cfg.DisableHTTP2)
		if err != nil {
			return err
		}
		log.Info("l2Client init success")

		MtBatherPrivateKey, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.PrivateKey, "0x"))
		if err != nil {
			return err
		}

		signer := func(chainID *big.Int) sequencer.SignerFn {
			s := PrivateKeySignerFn(MtBatherPrivateKey, chainID)
			return func(_ context.Context, addr ethc.Address, tx *types.Transaction) (*types.Transaction, error) {
				return s(addr, tx)
			}
		}

		driverConfig := &sequencer.DriverConfig{
			L1Client:                  l1Client,
			L2Client:                  l2Client,
			EigenContractAddr:         ethc.Address(common.HexToAddress(cfg.EigenContractAddress)),
			EigenFeeContractAddress:   ethc.Address(common.HexToAddress(cfg.EigenFeeContractAddress)),
			PrivKey:                   sequencerPrivKey,
			BlockOffset:               cfg.BlockOffset,
			RollUpMinSize:             cfg.RollUpMinSize,
			RollUpMaxSize:             cfg.RollUpMaxSize,
			EigenLayerNode:            cfg.EigenLayerNode,
			ChainID:                   chainID,
			DataStoreDuration:         uint64(cfg.DataStoreDuration),
			DataStoreTimeout:          cfg.DataStoreTimeout,
			DisperserSocket:           cfg.DisperserEndpoint,
			FeeSizeSec:                cfg.FeeSizeSec,
			PollInterval:              cfg.PollInterval,
			GraphProvider:             cfg.GraphProvider,
			ResubmissionTimeout:       cfg.ResubmissionTimeout,
			NumConfirmations:          cfg.NumConfirmations,
			SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
			FeeModelEnable:            cfg.FeeModelEnable,
			SignerFn:                  signer(chainID),
		}
		driver, err := sequencer.NewDriver(ctx, driverConfig)
		if err != nil {
			log.Error("new driver fail", "err", err)
			return err
		}
		if err := driver.Start(); err != nil {
			return err
		}
		log.Info("driver init success")

		defer driver.Stop()
		log.Info("mt batcher started")
		return nil
	}
}

func PrivateKeySignerFn(key *ecdsa.PrivateKey, chainID *big.Int) bind.SignerFn {
	from := crypto.PubkeyToAddress(key.PublicKey)
	signer := types.LatestSignerForChainID(chainID)
	return func(address ethc.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != from {
			return nil, bind.ErrNotAuthorized
		}
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}
}

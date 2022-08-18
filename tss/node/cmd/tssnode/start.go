package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/bitdao-io/bitnetwork/l2geth/crypto"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bitdao-io/bitnetwork/tss/node/config"
	"github.com/bitdao-io/bitnetwork/tss/node/server"
	sign "github.com/bitdao-io/bitnetwork/tss/node/signer"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/common"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a tss node",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNode(cmd)
		},
	}
	cmd.Flags().StringP("port", "p", "8080", "health server port")
	cmd.Flags().Bool("full", false, "default false, if true, the tss server will not start util all configured bootstrap peers are connected")
	cmd.Flags().Bool("non-prod", false, "default false, only for non-production environment")
	return cmd
}

func runNode(cmd *cobra.Command) error {
	cfgFile, _ := cmd.Flags().GetString("config")
	serverPort, _ := cmd.Flags().GetString("port")
	nonProd, _ := cmd.Flags().GetBool("non-prod")
	waitPeersFullConnected, _ := cmd.Flags().GetBool("full")
	cfg := config.DefaultConfiguration()
	err := config.LoadConfig(cfgFile, &cfg)
	if err != nil {
		log.Error().Err(err).Msg("fail to load config ")
		return err
	}

	if len(cfg.Key.PrivateKey) == 0 {
		return errors.New("need to config private key")
	}

	privKey, err := crypto.HexToECDSA(cfg.Key.PrivateKey)
	if err != nil {
		return err
	}

	//new tss server instance
	p2pPort, err := strconv.Atoi(cfg.P2PPort)
	if err != nil {
		log.Error().Err(err).Msg("p2p port value in config file, can not convert to int type")
	}

	cfgBz, _ := json.Marshal(cfg)
	log.Info().Str("config: ", string(cfgBz)).Msg("configuration file context")
	tssInstance, err := tsslib.NewTss(
		cfg.BootstrapPeers,
		waitPeersFullConnected,
		p2pPort,
		privKey,
		cfg.BaseDir,
		common.TssConfig{
			PreParamTimeout: cfg.PreParamTimeout,
			KeyGenTimeout:   cfg.KeyGenTimeout,
			KeySignTimeout:  cfg.KeySignTimeout,
			EnableMonitor:   false,
		},
		cfg.PreParamFile,
		cfg.ExternalIP,
		cfg.Secrets.Enable,
		cfg.Secrets.SecretId,
		cfg.Shamir,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to create tss server instance")
	}
	if err := tssInstance.Start(); err != nil {
		log.Error().Err(err).Msg("fail to start tss server")
	}
	pubkey := hex.EncodeToString(crypto.FromECDSAPub(&privKey.PublicKey))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	signer, err := sign.NewProcessor(cfg, ctx, tssInstance, privKey, pubkey)
	if err != nil {
		log.Error().Err(err).Msg("fail to new signer ")
		return err
	}
	if cfg.PauseSigning {
		signer.Pause()
	}
	signer.Start()

	hs := server.NewHttpServer(":"+serverPort, tssInstance, signer, nonProd)

	if err := hs.Start(); err != nil {
		log.Error().Err(err).Msg("fail to start http server")
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("stop signal received ")

	tssInstance.Stop()
	signer.Stop()
	hs.Stop()
	log.Info().Msg("server stopped")

	return nil
}

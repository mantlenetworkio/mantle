package tssnode

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/tss/node/store"

	tss "github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/index"
	"github.com/mantlenetworkio/mantle/tss/node/server"
	sign "github.com/mantlenetworkio/mantle/tss/node/signer"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/common"
	"github.com/mantlenetworkio/mantle/tss/slash"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "node",
		Short: "launch a tss node process",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runNode(cmd)
		},
	}
	return cmd
}

func runNode(cmd *cobra.Command) error {
	serverPort, _ := cmd.Flags().GetString("port")
	nonProd, _ := cmd.Flags().GetBool("non-prod")
	waitPeersFullConnected, _ := cmd.Flags().GetBool("full")
	cfg := tss.GetConfigFromCmd(cmd)

	if len(cfg.Node.PrivateKey) == 0 {
		return errors.New("need to config private key")
	}

	privKey, err := crypto.HexToECDSA(cfg.Node.PrivateKey)
	if err != nil {
		return err
	}

	//new level db storage
	store, err := store.NewStorage(cfg.Node.DBDir)
	if err != nil {
		return err
	}

	observer, err := index.NewIndexer(store, cfg.L1Url, cfg.L1ConfirmBlocks, cfg.SccContractAddress, cfg.TimedTaskInterval)
	if err != nil {
		return err
	}
	observer.SetHook(slash.NewSlashing(store, store, cfg.SignedBatchesWindow, cfg.MinSignedInWindow))
	observer.Start()

	//new tss server instance
	p2pPort, err := strconv.Atoi(cfg.Node.P2PPort)
	if err != nil {
		log.Error().Err(err).Msg("p2p port value in config file, can not convert to int type")
	}

	cfgBz, _ := json.Marshal(cfg)
	log.Info().Str("config: ", string(cfgBz)).Msg("configuration file context")
	tssInstance, err := tsslib.NewTss(
		cfg.Node.BootstrapPeers,
		waitPeersFullConnected,
		p2pPort,
		privKey,
		cfg.Node.BaseDir,
		common.TssConfig{
			PreParamTimeout: cfg.Node.PreParamTimeout,
			KeyGenTimeout:   cfg.Node.KeyGenTimeout,
			KeySignTimeout:  cfg.Node.KeySignTimeout,
			EnableMonitor:   false,
		},
		cfg.Node.PreParamFile,
		cfg.Node.ExternalIP,
		cfg.Node.Secrets.Enable,
		cfg.Node.Secrets.SecretId,
		cfg.Node.Shamir,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to create tss server instance")
	}
	if err := tssInstance.Start(); err != nil {
		log.Error().Err(err).Msg("fail to start tss server")
	}

	pubkey := crypto.CompressPubkey(&privKey.PublicKey)
	pubkeyHex := hex.EncodeToString(pubkey)

	localPubkeyBytes := crypto.FromECDSAPub(&privKey.PublicKey)
	// bytes len is 64
	localPubkeyBytes = localPubkeyBytes[1:]

	address := ethcrypto.PubkeyToAddress(privKey.PublicKey)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	signer, err := sign.NewProcessor(cfg, ctx, tssInstance, privKey, localPubkeyBytes, pubkeyHex, store, address)
	if err != nil {
		log.Error().Err(err).Msg("fail to new signer ")
		return err
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

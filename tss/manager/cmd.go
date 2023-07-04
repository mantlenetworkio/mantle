package manager

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/index"
	"github.com/mantlenetworkio/mantle/tss/manager/l1chain"
	"github.com/mantlenetworkio/mantle/tss/manager/router"
	"github.com/mantlenetworkio/mantle/tss/manager/store"
	"github.com/mantlenetworkio/mantle/tss/slash"
	"github.com/mantlenetworkio/mantle/tss/ws/server"
)

const (
	jwtExpiryTimeout = 60 * time.Second
	jwtKeyLength     = 32
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "manager",
		Short: "launch a tss manager process",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return run(cmd)
		},
	}
	return cmd
}

func run(cmd *cobra.Command) error {
	config := common.GetConfigFromCmd(cmd)
	log.Info("config info print", "MissSignedNumber", config.MissSignedNumber)
	log.Info("l1 start block number", "block number", config.L1StartBlockNumber)
	l1StartBlockNumber, err := strconv.ParseUint(
		config.L1StartBlockNumber, 10, 32,
	)
	if err != nil {
		return err
	}

	managerStore, err := store.NewStorage(config.Manager.DBDir)
	if err != nil {
		return err
	}
	queryService, err := l1chain.NewQueryService(config.L1Url, config.TssGroupContractAddress, config.L1ConfirmBlocks, managerStore)
	if err != nil {
		return err
	}
	if len(config.Manager.PrivateKey) == 0 {
		return errors.New("need to config private key")
	}
	wsServer, err := server.NewWSServer(config.Manager.WsAddr, queryService)
	if err != nil {
		return err
	}

	observer, err := index.NewIndexer(managerStore, config.L1Url, config.L1ConfirmBlocks, config.SccContractAddress, config.TimedTaskInterval, l1StartBlockNumber)
	if err != nil {
		return err
	}
	observer = observer.SetHook(slash.NewSlashing(managerStore, managerStore, config.MissSignedNumber))
	observer.Start()

	manager, err := NewManager(wsServer, queryService, managerStore, config)
	if err != nil {
		return err
	}
	manager.Start()

	registry := router.NewRegistry(manager, managerStore)
	r := gin.Default()
	registry.Register(r)

	var s *http.Server
	if config.Manager.JwtSecret != "" {
		jwtHandler, err := common.NewJwtHandler(r, config.Manager.JwtSecret)
		if err != nil {
			return err
		}

		// custom http configuration
		s = &http.Server{
			Addr:    config.Manager.HttpAddr,
			Handler: jwtHandler,
		}
	} else {
		s = &http.Server{
			Addr:    config.Manager.HttpAddr,
			Handler: r,
		}
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Error("api server starts failed", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	manager.Stop()
	observer.Stop()

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown", err)
		return err
	}

	log.Info("Server exiting")

	return nil
}

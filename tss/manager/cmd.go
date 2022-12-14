package manager

import (
	"context"
	"errors"
	"github.com/mantlenetworkio/mantle/tss/index"
	"github.com/mantlenetworkio/mantle/tss/manager/l1chain"
	"github.com/mantlenetworkio/mantle/tss/manager/store"
	"github.com/mantlenetworkio/mantle/tss/slash"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/manager/router"
	"github.com/mantlenetworkio/mantle/tss/ws/server"
	"github.com/spf13/cobra"
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
	log.Info("config info print", "SignedBatchesWindow", config.SignedBatchesWindow, "MinSignedInWindow", config.MinSignedInWindow)
	wsServer, err := server.NewWSServer(config.Manager.WsAddr)
	if err != nil {
		return err
	}
	managerStore, err := store.NewStorage(config.Manager.DBDir)
	if err != nil {
		return err
	}
	observer, err := index.NewIndexer(managerStore, config.L1Url, config.L1ConfirmBlocks, config.SccContractAddress, config.TimedTaskInterval)
	if err != nil {
		return err
	}
	observer = observer.SetHook(slash.NewSlashing(managerStore, managerStore, config.SignedBatchesWindow, config.MinSignedInWindow))
	observer.Start()

	queryService := l1chain.NewQueryService(config.L1Url, config.TssGroupContractAddress, config.L1ConfirmBlocks, managerStore)
	manager, err := NewManager(wsServer, queryService, managerStore, config)
	if err != nil {
		return err
	}
	manager.Start()

	registry := router.NewRegistry(manager, managerStore)
	r := gin.Default()
	registry.Register(r)

	// custom http configuration
	s := &http.Server{
		Addr:    config.Manager.HttpAddr,
		Handler: r,
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

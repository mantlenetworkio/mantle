package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/bitdao-io/bitnetwork/tss/index"
	"github.com/bitdao-io/bitnetwork/tss/manager/store"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bitdao-io/bitnetwork/l2geth/log"
	"github.com/bitdao-io/bitnetwork/tss/common"
	"github.com/bitdao-io/bitnetwork/tss/manager/router"
	"github.com/bitdao-io/bitnetwork/tss/ws/server"
	"github.com/gin-gonic/gin"
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
	fmt.Println(config)

	wsServer, err := server.NewWSServer("")
	if err != nil {
		return err
	}
	managerStore, err := store.NewStorage("")
	if err != nil {
		return err
	}
	observer, err := index.NewObserver(managerStore, "", "")
	if err != nil {
		return err
	}
	manager, err := NewManager(wsServer, nil, managerStore, observer, "")
	if err != nil {
		return err
	}
	manager.Start()

	registry := router.NewRegistry(manager)
	r := gin.Default()
	registry.Register(r)

	// custom http configuration
	s := &http.Server{
		Addr:         "",
		Handler:      r,
		ReadTimeout:  0,
		WriteTimeout: 0,
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

package manager

import (
	"context"
	"errors"
	"github.com/bitdao-io/bitnetwork/l2geth/log"
	"github.com/ethereum-optimism/optimism/tss"
	"github.com/ethereum-optimism/optimism/tss/manager/router"
	"github.com/ethereum-optimism/optimism/tss/ws/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Command(cfg tss.Configuration) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "manager",
		Short: "launch a tss manager process",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return run(cfg)
		},
	}
	return cmd
}

func run(cfg tss.Configuration) error {
	wsServer, err := server.NewWSServer("")
	if err != nil {
		return err
	}
	manager := NewManager(wsServer)

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

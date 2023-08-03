package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	cmm "github.com/mantlenetworkio/mantle/tss/common"
	sign "github.com/mantlenetworkio/mantle/tss/node/signer"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib"
)

type Server struct {
	logger    zerolog.Logger
	tssServer tsslib.Server
	signer    *sign.Processor
	s         *http.Server
	nonProd   bool
}

type GenRequest struct {
	PartyPubKeys []string `json:"party_pub_keys"`
	Threshold    int      `json:"threshold" mapstructure:"threshold"`
}

type MockEventRequest struct {
}

func NewHttpServer(addr string, t tsslib.Server, signer *sign.Processor, nonProd bool, jwtSecretStr string) (*Server, error) {
	hs := &Server{
		logger:    log.With().Str("module", "http").Logger(),
		tssServer: t,
		nonProd:   nonProd,
		signer:    signer,
	}
	handler, err := hs.newHandler(jwtSecretStr)
	if err != nil {
		return nil, err
	}
	s := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	hs.s = s
	return hs, nil
}

func (hs *Server) Start() error {
	if hs.s == nil {
		return errors.New("invalid http server instance")
	}
	if err := hs.s.ListenAndServe(); err != nil {
		hs.logger.Error().Err(err).Msg("api server starts failed")
		return err
	}
	return nil
}

func (hs *Server) Stop() {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := hs.s.Shutdown(c)
	if err != nil {
		hs.logger.Error().Err(err).Msg("Failed to shutdown the server gracefully")
	}
}

func (hs *Server) newHandler(jwtSecretStr string) (http.Handler, error) {
	router := mux.NewRouter()
	router.Handle("/metrics", promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer, promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{MaxRequestsInFlight: 3},
		),
	))

	router.Use(logMiddleware())
	if jwtSecretStr != "" {
		return cmm.NewJwtHandler(router, jwtSecretStr)
	} else {
		return router, nil
	}
}

func logMiddleware() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Debug().
				Str("route", r.URL.Path).
				Str("port", r.URL.Port()).
				Str("method", r.Method).
				Msg("HTTP request received")

			handler.ServeHTTP(w, r)
		})
	}
}

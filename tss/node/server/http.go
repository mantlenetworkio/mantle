package server

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	cmm "github.com/mantlenetworkio/mantle/tss/common"
	sign "github.com/mantlenetworkio/mantle/tss/node/signer"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/common"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/keygen"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/keysign"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	jwtExpiryTimeout = 60 * time.Second
	jwtKeyLength     = 32
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

type jwtHandler struct {
	keyFunc func(token *jwt.Token) (interface{}, error)
	handler http.Handler
}

// ServeHTTP implements http.Handler
func (handler *jwtHandler) ServeHTTP(out http.ResponseWriter, r *http.Request) {
	var (
		strToken string
		claims   jwt.RegisteredClaims
	)
	if auth := r.Header.Get("Authorization"); strings.HasPrefix(auth, "Bearer ") {
		strToken = strings.TrimPrefix(auth, "Bearer ")
	}
	if len(strToken) == 0 {
		http.Error(out, "missing token", http.StatusUnauthorized)
		return
	}
	// We explicitly set only HS256 allowed, and also disables the
	// claim-check: the RegisteredClaims internally requires 'iat' to
	// be no later than 'now', but we allow for a bit of drift.
	token, err := jwt.ParseWithClaims(strToken, &claims, handler.keyFunc,
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithoutClaimsValidation())

	switch {
	case err != nil:
		http.Error(out, err.Error(), http.StatusUnauthorized)
	case !token.Valid:
		http.Error(out, "invalid token", http.StatusUnauthorized)
	case !claims.VerifyExpiresAt(time.Now(), false): // optional
		http.Error(out, "token is expired", http.StatusUnauthorized)
	case claims.IssuedAt == nil:
		http.Error(out, "missing issued-at", http.StatusUnauthorized)
	case time.Since(claims.IssuedAt.Time) > jwtExpiryTimeout:
		http.Error(out, "stale token", http.StatusUnauthorized)
	case time.Until(claims.IssuedAt.Time) > jwtExpiryTimeout:
		http.Error(out, "future token", http.StatusUnauthorized)
	default:
		handler.handler.ServeHTTP(out, r)
	}
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
	router.Handle("/ping", http.HandlerFunc(hs.pingHandler)).Methods(http.MethodGet)
	router.Handle("/gen-key", http.HandlerFunc(hs.keyGenHandler)).Methods(http.MethodPost)
	router.Handle("/key-sign", http.HandlerFunc(hs.keySignHandler)).Methods(http.MethodPost)
	router.Handle("/metrics", promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer, promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{MaxRequestsInFlight: 3},
		),
	))

	router.Use(logMiddleware())

	return cmm.NewJwtHandler(router, jwtSecretStr)
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

func (hs *Server) pingHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (hs *Server) keyGenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer func() {
		if err := r.Body.Close(); nil != err {
			hs.logger.Error().Err(err).Msg("fail to close request body")
		}
	}()
	hs.logger.Info().Msg("receive key gen request")
	decoder := json.NewDecoder(r.Body)
	var genRequest GenRequest
	if err := decoder.Decode(&genRequest); nil != err {
		hs.logger.Error().Err(err).Msg("fail to decode keygen request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var keygenReq = keygen.Request{
		Keys:      genRequest.PartyPubKeys,
		ThresHold: genRequest.Threshold,
	}

	resp, err := hs.tssServer.Keygen(keygenReq)
	if err != nil {
		hs.logger.Error().Err(err).Msg("fail to key gen")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	hs.logger.Debug().Msgf("resp:%+v", resp)
	if resp.Status == common.Success {
		hs.logger.Info().Msgf("generated compressed public key: %s \n", resp.PubKey)
		addressBytes := resp.Address
		address := hex.EncodeToString(addressBytes)
		hs.logger.Info().Msgf("generated address: %s \n", address)

		buf, _ := json.Marshal(map[string]string{"pubKey": resp.PubKey, "address": address})
		if _, err = w.Write(buf); err != nil {
			hs.logger.Error().Err(err).Msg("fail to write to response")
		}
	} else {
		buf, err := json.Marshal(resp)
		if err != nil {
			hs.logger.Error().Err(err).Msg("fail to marshal response to json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if _, err = w.Write(buf); err != nil {
			hs.logger.Error().Err(err).Msg("fail to write to response")
		}
	}

}

func (hs *Server) keySignHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer func() {
		if err := r.Body.Close(); nil != err {
			hs.logger.Error().Err(err).Msg("fail to close request body")
		}
	}()
	hs.logger.Info().Msg("receive key sign request")

	var keySignReq keysign.Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&keySignReq); nil != err {
		hs.logger.Error().Err(err).Msg("fail to decode key sign request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hs.logger.Info().Msgf("request:%+v", keySignReq)
	signResp, err := hs.tssServer.KeySign(keySignReq)
	if err != nil {
		hs.logger.Error().Err(err).Msg("fail to key sign")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResult, err := json.MarshalIndent(signResp, "", "	")
	if err != nil {
		hs.logger.Error().Err(err).Msg("fail to marshal response to json message")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(jsonResult)
	if err != nil {
		hs.logger.Error().Err(err).Msg("fail to write response")
	}
}

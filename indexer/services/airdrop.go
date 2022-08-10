package services

import (
	"net/http"

	"github.com/bitdao-io/bitnetwork/indexer/db"
	"github.com/bitdao-io/bitnetwork/indexer/metrics"
	"github.com/bitdao-io/bitnetwork/indexer/server"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/mux"
)

var airdropLogger = log.New("service", "airdrop")

type Airdrop struct {
	db      *db.Database
	metrics *metrics.Metrics
}

func NewAirdrop(db *db.Database, metrics *metrics.Metrics) *Airdrop {
	return &Airdrop{
		db:      db,
		metrics: metrics,
	}
}

func (a *Airdrop) GetAirdrop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	airdrop, err := a.db.GetAirdrop(common.HexToAddress(address))
	if err != nil {
		airdropLogger.Error("db error getting airdrop", "err", err)
		server.RespondWithError(w, http.StatusInternalServerError, "database error")
		return
	}

	if airdrop == nil {
		server.RespondWithError(w, http.StatusNotFound, "airdrop not found")
		return
	}

	server.RespondWithJSON(w, http.StatusOK, airdrop)
}

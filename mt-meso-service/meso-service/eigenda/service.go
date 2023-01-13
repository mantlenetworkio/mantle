package eigenda

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	middleware2 "github.com/mantlenetworkio/mantle/mt-meso-service/middleware"
	"strconv"
)

type DaServiceConfig struct {
	EigenMiddleWareConfig *middleware2.EigenDaMiddleWareConfig
	DaServicePort         int
	Debug                 bool
}

type DaService struct {
	Cfg               *DaServiceConfig
	EigenDaMiddleWare *middleware2.EigenDaMiddleWare
	echo              *echo.Echo
}

func NewDaService(ctx context.Context, cfg *DaServiceConfig) (*DaService, error) {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Debug = cfg.Debug
	e.Use(middleware.Recover())
	eigenDaMiddleWare, err := middleware2.NewEigenDaMiddleWare(ctx, cfg.EigenMiddleWareConfig)
	if err != nil {
		log.Error("new eigen da middle ware fail", "err", err)
		return nil, err
	}
	server := &DaService{
		Cfg:               cfg,
		EigenDaMiddleWare: eigenDaMiddleWare,
		echo:              e,
	}
	server.routes()
	return server, nil
}

func (s *DaService) routes() {
	s.echo.GET("eigen/getBatchTransactionByDataStoreId", s.GetBatchTransactionByDataStoreId)
	s.echo.POST("eigen/getRollupStoreByRollupBatchIndex", s.GetRollupStoreByRollupBatchIndex)
	s.echo.POST("eigen/getLatestTransactionBatchIndex", s.GetLatestTransactionBatchIndex)
}

func (s *DaService) Run() {
	err := s.echo.Start(":" + strconv.Itoa(s.Cfg.DaServicePort))
	if err != nil {
		log.Error("eigen da sever start fail")
	}
	log.Info("eigen da sever start success", "port", s.Cfg.DaServicePort)
}

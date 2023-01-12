package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

type EigenDaServerConfig struct {
	Port  int
	Debug bool
}

type EigenDaServer struct {
	Cfg  *EigenDaServerConfig
	echo *echo.Echo
}

func NewServer(cfg *EigenDaServerConfig) *EigenDaServer {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Debug = cfg.Debug
	e.Use(middleware.Recover())
	server := &EigenDaServer{
		Cfg:  cfg,
		echo: e,
	}
	server.routes()
	return server
}

func (s *EigenDaServer) routes() {
	s.echo.GET("eigen/:get", s.GetBatchTransactionByDataStoreId)
	s.echo.GET("eigen/:get", s.GetLatestTransactionBatchIndex)
}

func (s *EigenDaServer) Run() {
	s.echo.Logger.Fatal(s.echo.Start(":" + strconv.Itoa(s.Cfg.Port)))
}

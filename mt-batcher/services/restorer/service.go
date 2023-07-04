package restorer

import (
	"context"
	"strconv"
	"sync"

	"github.com/Layr-Labs/datalayr/common/graphView"
	"github.com/ethereum/go-ethereum/log"
	gecho "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shurcooL/graphql"

	"github.com/mantlenetworkio/mantle/mt-batcher/bindings"
	"github.com/mantlenetworkio/mantle/mt-batcher/services/common"
)

type DaServiceConfig struct {
	EigenContract   *bindings.BVMEigenDataLayrChain
	RetrieverSocket string
	GraphProvider   string
	DaServicePort   int
	EigenLayerNode  int
}

type DaService struct {
	Ctx           context.Context
	Cfg           *DaServiceConfig
	GraphClient   *graphView.GraphClient
	GraphqlClient *graphql.Client
	echo          *gecho.Echo
	cancel        func()
	wg            sync.WaitGroup
}

func NewDaService(ctx context.Context, cfg *DaServiceConfig) (*DaService, error) {
	_, cancel := context.WithTimeout(ctx, common.DefaultTimeout)
	defer cancel()
	e := gecho.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Debug = true
	e.Use(middleware.Recover())
	graphClient := graphView.NewGraphClient(cfg.GraphProvider, nil)
	graphqlClient := graphql.NewClient(graphClient.GetEndpoint(), nil)
	server := &DaService{
		Ctx:           ctx,
		Cfg:           cfg,
		GraphClient:   graphClient,
		GraphqlClient: graphqlClient,
		echo:          e,
		cancel:        cancel,
	}
	server.routes()
	return server, nil
}

func (s *DaService) routes() {
	s.echo.GET("eigen/getLatestTransactionBatchIndex", s.GetLatestTransactionBatchIndex)
	s.echo.POST("eigen/getRollupStoreByRollupBatchIndex", s.GetRollupStoreByRollupBatchIndex)
	s.echo.POST("eigen/getBatchTransactionByDataStoreId", s.GetBatchTransactionByDataStoreId)
	s.echo.POST("dtl/getBatchTransactionByDataStoreId", s.GetDtlBatchTransactionByDataStoreId)
	s.echo.POST("browser/getDataStoreList", s.GetDataStoreList)
	s.echo.POST("browser/getDataStoreById", s.getDataStoreById)
	s.echo.POST("browser/GetTransactionListByStoreNumber", s.GetTransactionListByStoreNumber)
}

func (s *DaService) Start() error {
	defer s.wg.Done()
	err := s.echo.Start(":" + strconv.Itoa(s.Cfg.DaServicePort))
	if err != nil {
		log.Error("eigen da sever start fail")
		return err
	}
	log.Info("eigen da sever start success", "port", s.Cfg.DaServicePort)
	return nil
}

func (s *DaService) Stop() {
	s.cancel()
	s.wg.Wait()
}

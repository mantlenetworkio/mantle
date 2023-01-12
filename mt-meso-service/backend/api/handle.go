package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *EigenDaServer) GetLatestTransactionBatchIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func (s *EigenDaServer) GetBatchTransactionByDataStoreId(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

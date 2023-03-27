package router

import (
	"errors"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	tss "github.com/mantlenetworkio/mantle/mt-tss/common"
	"github.com/mantlenetworkio/mantle/mt-tss/manager/types"
)

type Registry struct {
	signService  types.SignService
	adminService types.AdminService
}

func NewRegistry(signService types.SignService, adminService types.AdminService) *Registry {
	return &Registry{
		signService:  signService,
		adminService: adminService,
	}
}

func (registry *Registry) SignStateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request tss.SignStateRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid request body"))
			return
		}

		signature, err := registry.signService.SignStateBatch(request)
		if err != nil {
			c.String(http.StatusInternalServerError, "failed to sign state")
			log.Error("failed to sign state", "error", err)
			return
		}

		if _, err = c.Writer.Write(signature); err != nil {
			log.Error("failed to write signature to response writer", "error", err)
		}
	}
}

func (registry *Registry) ResetHeightHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		heightStr := c.Param("height")
		height, err := strconv.Atoi(heightStr)
		if err != nil {
			c.String(http.StatusInternalServerError, "wrong height format")
			log.Error("failed to reset height", "error", err)
			return
		}
		err = registry.adminService.ResetScanHeight(uint64(height))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			log.Error("failed to reset height", "error", err)
			return
		}
	}
}

func (registry *Registry) GetHeightHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		height, err := registry.adminService.GetScannedHeight()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			log.Error("failed to get height", "error", err)
			return
		}
		c.String(http.StatusOK, strconv.FormatUint(height, 10))
	}
}

func (registry *Registry) DeleteSlashHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		addressStr := c.Query("address")
		if len(addressStr) == 0 {
			c.String(http.StatusBadRequest, "empty address")
			return
		}
		batchIndexStr := c.Query("index")
		if len(batchIndexStr) == 0 {
			c.String(http.StatusBadRequest, "empty index")
			return
		}
		address := common.HexToAddress(addressStr)
		index, err := strconv.Atoi(batchIndexStr)
		if err != nil {
			c.String(http.StatusBadRequest, "wrong format index")
			return
		}
		registry.adminService.RemoveSlashingInfo(address, uint64(index))
		c.String(http.StatusOK, "success")
	}
}

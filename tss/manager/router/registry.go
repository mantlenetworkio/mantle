package router

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	tss "github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/manager/types"
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

		_, succ := new(big.Int).SetString(request.StartBlock, 10)
		if !succ {
			c.JSON(http.StatusBadRequest, errors.New("wrong StartBlock, can not be converted to number"))
			return
		}
		_, succ = new(big.Int).SetString(request.OffsetStartsAtIndex, 10)
		if !succ {
			c.JSON(http.StatusBadRequest, errors.New("wrong OffsetStartsAtIndex, can not be converted to number"))
			return
		}
		var signature []byte
		var err error
		if request.Type == 0 {
			signature, err = registry.signService.SignStateBatch(request)
		} else if request.Type == 1 {
			signature, err = registry.signService.SignRollBack(request)
		}
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

package router

import (
	"errors"
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

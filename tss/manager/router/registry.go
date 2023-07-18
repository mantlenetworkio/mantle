package router

import (
	"errors"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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
		if request.StartBlock == nil ||
			request.OffsetStartsAtIndex == nil ||
			request.StartBlock.Cmp(big.NewInt(0)) < 0 ||
			request.OffsetStartsAtIndex.Cmp(big.NewInt(0)) < 0 {
			c.JSON(http.StatusBadRequest, errors.New("StartBlock and OffsetStartsAtIndex must not be nil or negative"))
			return
		}
		var signature []byte
		var err error
		if request.Type == 0 {
			signature, err = registry.signService.SignStateBatch(request)
		} else if request.Type == 1 {
			if !common.IsHexAddress(request.Challenge) {
				c.JSON(http.StatusBadRequest, errors.New("wrong challenge address, can not be converted to hex address"))
				return
			}
			signature, err = registry.signService.SignRollBack(request)
		} else {
			c.String(http.StatusBadRequest, "invalid request type %d, expected request type: 0 and 1", request.Type)
			return
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
		heightStr := c.PostForm("height")
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

func (registry *Registry) PrometheusHandler() gin.HandlerFunc {
	h := promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer, promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{MaxRequestsInFlight: 3},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

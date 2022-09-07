package router

import (
	"errors"
	"math/big"
	"net/http"

	"github.com/bitdao-io/mantle/l2geth/log"
	tss "github.com/bitdao-io/mantle/tss/common"
	"github.com/bitdao-io/mantle/tss/manager/types"
	"github.com/gin-gonic/gin"
)

type Registry struct {
	signService types.SignService
}

func NewRegistry(signService types.SignService) *Registry {
	return &Registry{
		signService: signService,
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

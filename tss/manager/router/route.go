package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (registry *Registry) Register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	v1Router := r.Group("/api/v1")
	v1Router.POST("/sign/state", registry.SignStateHandler())

	v1Router.POST("/admin/reset/height", registry.ResetHeightHandler())
}

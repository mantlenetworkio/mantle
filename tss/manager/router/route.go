package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (registry *Registry) Register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	v1Router := r.Group("/api/v1")
	v1Router.POST("/sign/state", registry.SignStateHandler())
	v1Router.GET("/metrics", registry.PrometheusHandler())

	v1Router.GET("/admin/height", registry.GetHeightHandler())
	v1Router.POST("/admin/reset/height", registry.ResetHeightHandler())
	v1Router.DELETE("/admin/delete/slash", registry.DeleteSlashHandler())

}

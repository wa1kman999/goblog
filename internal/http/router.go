package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	v1prefix = "/go_blog/v1"
)

// initRouter 初始化路由
func initRouter(router *gin.Engine) error {

	router.GET("/ready", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	router.GET("/healthy", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	return nil
}

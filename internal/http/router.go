package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userController "github.com/wa1kman999/goblog/internal/controller/user"
)

const (
	v1prefix = "/goblog/v1"
)

// initRouter 初始化路由
func initRouter(router *gin.Engine) error {

	router.GET("/ready", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	router.GET("/healthy", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	user := router.Group(v1prefix + "/")
	// 使用中间件
	{
		// 新建用户
		user.POST("/user", userController.CreateUser)
	}

	return nil
}

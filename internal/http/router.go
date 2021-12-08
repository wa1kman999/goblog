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
		// 通过id查询用户
		user.GET("/user/:id", userController.GetUserInfo)
		// 查询用户列表
		user.POST("/user/list", userController.GetUserList)
		// 编辑用户
		//user.PUT("/user", userController.EditUser)
	}

	return nil
}

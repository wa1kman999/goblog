package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	systemController "github.com/wa1kman999/goblog/internal/controller/system"
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

	/*
		前端展示页面接口
	*/
	router.POST("/login", userController.Login)
	user := router.Group(v1prefix + "/user")
	{
		// 新建用户
		user.POST("", userController.CreateUser)
		// 通过id查询用户
		user.GET("/:id", userController.GetUserInfo)
		// 查询用户列表
		user.POST("/list", userController.GetUserList)
		// 编辑用户
		//user.PUT("/user", userController.EditUser)
	}
	system := router.Group(v1prefix + "/system")
	{
		// 查询服务器状态
		system.GET("/state", systemController.GetServerInfo)
	}

	return nil
}

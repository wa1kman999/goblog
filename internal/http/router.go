package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	articleController "github.com/wa1kman999/goblog/internal/controller/article"
	categoryController "github.com/wa1kman999/goblog/internal/controller/category"
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
	// 用户相关的接口
	user := router.Group(v1prefix + "/user")
	{
		// 新建用户
		user.POST("", userController.CreateUser)
		// 通过id查询用户
		user.GET("/:id", userController.GetUserInfo)
		// 查询用户列表
		user.POST("/list", userController.GetUserList)
		// 编辑用户
		user.PUT("", userController.EditUser)
		// 删除用户
		user.DELETE("/:id", userController.DeleteUser)
	}
	// 文章相关的接口
	article := router.Group(v1prefix + "/article")
	{
		// 新建文章
		article.POST("", articleController.CreateArticle)
		// 上传图片到本地
		article.POST("/img", articleController.Upload)
		// 获取图片
		article.GET("/img/:path", articleController.GetImg)
		// 文章列表
		article.POST("/list", articleController.GetArticleList)
	}
	// 分类相关接口
	category := router.Group(v1prefix + "/category")
	{
		// 新建分类
		category.POST("", categoryController.CreateCategory)
		// 获取分类列表
		category.POST("/list", categoryController.GetCategoryList)
	}

	system := router.Group(v1prefix + "/system")
	{
		// 查询服务器状态
		system.GET("/state", systemController.GetServerInfo)
	}

	return nil
}

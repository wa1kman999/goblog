package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginSwaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wa1kman999/goblog/global"
	"github.com/wa1kman999/goblog/internal/http/middleware"
	"github.com/wa1kman999/goblog/pkg/common/constants"
)

var httpServer *http.Server

func router() http.Handler {
	r := gin.New()
	// 跨域中间件
	r.Use(middleware.Cors())

	// 请求日志
	//r.Use(middleware.Logger())
	if gin.Mode() == gin.DebugMode {
		r.Use(gin.Logger())
	}

	// 错误恢复
	r.Use(gin.Recovery())

	err := initRouter(r)
	if err != nil {
		log.Fatal("initialize router", err)
		return r
	}

	// release 模式下不提供接口文档访问
	if string(constants.ReleaseMode) != os.Getenv("GIN_MODE") {
		r.GET(v1prefix+"/swagger/*any", ginSwagger.WrapHandler(ginSwaggerFiles.Handler))
	}

	return r
}

// Serve 启动服务
func Serve() error {

	httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", global.GBConfig.System.Host, global.GBConfig.System.Port),
		Handler:      router(),
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	return httpServer.ListenAndServe()
}

// Shutdown 关闭服务
func Shutdown() error {
	log.Infof("正在关闭http服务")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		return err
	}
	log.Info("http服务成功关闭")
	return nil
}

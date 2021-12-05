package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wa1kman999/goblog/global"
	"github.com/wa1kman999/goblog/internal/http/middleware"
)

var httpServer *http.Server

func router() http.Handler {
	r := gin.New()
	// 跨域中间件
	r.Use(middleware.Cors())
	// 请求日志
	r.Use(middleware.DefaultLogger())

	err := initRouter(r)
	if err != nil {
		log.Fatal("initialize router", err)
		return r
	}

	// 错误恢复
	r.Use(gin.Recovery())

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
	log.Println("正在关闭http服务")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		return err
	}
	log.Println("http服务成功关闭")
	return nil
}

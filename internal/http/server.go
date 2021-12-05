package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var httpServer *http.Server

func router() http.Handler {
	r := gin.New()
	err := initRouter(r)
	if err != nil {
		log.Fatal("init router", err)
		return r
	}

	// 错误恢复
	r.Use(gin.Recovery())

	return r
}

// Serve 启动服务
func Serve() error {
	// TODO 读取配置文件

	//log.Println("正在启动http服务，监听端口", conf.Server.Port)

	httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", "localhost", "8090"),
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

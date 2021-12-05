package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/wa1kman999/goblog/config"
	"github.com/wa1kman999/goblog/initialize"
	httpServer "github.com/wa1kman999/goblog/internal/http"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	g.Go(func() error {
		if err := httpServer.Serve(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
		return nil
	})

	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		for {
			si := <-c
			switch si {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				return shutdown()
			case syscall.SIGHUP:
			default:
				return nil
			}
		}
	})

	if err := g.Wait(); err != nil {
		log.Printf("服务运行失败: %s", err)
		panic(err)
	}

}

// shutdown 关闭服务
func shutdown() error {
	// 关闭http服务
	if err := httpServer.Shutdown(); err != nil {
		return err
	}
	return nil
}

func init() {
	_ = config.Init()
	conf := config.Get()

	log.Println(fmt.Sprintf("redis %s %s %d ", conf.Redis.Addr, conf.Redis.Password, conf.Redis.DB))

	// 初始化mysql连接

	// 初始化redis
	initialize.RedisClient()
}

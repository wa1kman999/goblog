package initialize

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/wa1kman999/goblog/global"
	"go.uber.org/zap"
)

func RedisClient() error {
	redisCfg := global.GBConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GBLog.Error("redis connect ping failed, err:", zap.Error(err))
		return err
	}
	global.GBLog.Info("redis connect ping response:", zap.String("pong", pong))
	global.GBRedis = client
	return nil
}

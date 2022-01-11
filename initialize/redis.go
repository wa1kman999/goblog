package initialize

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/wa1kman999/goblog/global"
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
		logrus.Errorf("redis connect ping failed, err: %s", err.Error())
		return err
	}
	logrus.Infof("redis connect ping response: %s", pong)
	global.GBRedis = client
	return nil
}

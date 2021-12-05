package initialize

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/wa1kman999/goblog/config"
)

var globalRedis *redis.Client

func RedisClient() {
	redisCfg := config.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("redis connect ping failed, err: %s", err.Error())
	} else {
		log.Printf("redis connect ping response: %s", pong)
		globalRedis = client
	}
}

package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/wa1kman999/goblog/config"
	"gorm.io/gorm"
)

var (
	GBConfig config.Config
	GBRedis  *redis.Client
	GBMysql  *gorm.DB
)

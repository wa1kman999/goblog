package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/wa1kman999/goblog/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GBConfig config.Config
	GBRedis  *redis.Client
	GBLog    *zap.Logger
	GBMysql  *gorm.DB
)

package cache

import (
	"fmt"
	"time"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"github.com/go-redis/redis/v7"
)

type single struct{}

var cfg = config.GetConfig()

var zap = logging.NewLogger(cfg)

var singleton *single

var redisClient *redis.Client

func InitRedis(cfg *config.Config) *single {
	if singleton == nil {
		fmt.Println("creating redis connection...")
		redisClient = redis.NewClient(&redis.Options{
			Addr:               fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
			Password:           cfg.Redis.Password,
			DB:                 0,
			DialTimeout:        cfg.Redis.DialTimeout * time.Second,
			ReadTimeout:        cfg.Redis.ReadTimeout * time.Second,
			WriteTimeout:       cfg.Redis.WriteTimeout * time.Second,
			PoolSize:           cfg.Redis.PoolSize,
			PoolTimeout:        cfg.Redis.PoolTimeout,
			IdleTimeout:        500 * time.Millisecond,
			IdleCheckFrequency: cfg.Redis.IdleCheckFrequency * time.Millisecond,
		})
		singleton = &single{}
		zap.Info(logging.Postgres, logging.Migration, "Redis connected successfully", nil)
		return singleton
	} else {
		zap.Info(logging.Postgres, logging.Migration, "Redis is already configured", nil)
		return singleton
	}
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}

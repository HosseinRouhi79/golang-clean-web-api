package cache

import (
	"fmt"
	"time"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	logger "github.com/HosseinRouhi79/golang-clean-web-api/src/logs"
	"github.com/go-redis/redis/v7"
)

var zLog = logger.Logger()

type single struct{}

var singleton *single

var redisClient *redis.Client

func InitRedis(cfg *config.Config) (*single) {
	if singleton == nil {
		zLog.Info().Msg("creating redis connection...")
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
		return singleton
	} else {
		zLog.Info().Msg("redis is already created")
		return singleton
	}
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}

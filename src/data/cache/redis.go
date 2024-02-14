package cache

import (
	"fmt"
	"log"
	"time"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/go-redis/redis/v7"
)

type single struct{}

var singleton *single

var redisClient *redis.Client

func InitRedis(cfg *config.Config) error {
	if singleton != nil {
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

		_, err := redisClient.Ping().Result()
		if err != nil {
			return err
		}
		singleton = &single{}
		return nil
	} else {
		log.Println("redis is already created")
		return nil
	}
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}

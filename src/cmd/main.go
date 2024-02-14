package main

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	// cache.InitRedis(cfg)
	api.InitServer(cfg)
	defer cache.CloseRedis()
}

package main

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/cache"
)

func main() {
	cfg := config.GetConfig()
	api.InitServer(cfg)
	cache.InitRedis(cfg)
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
}

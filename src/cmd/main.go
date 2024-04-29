package main

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/cache"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/db"
	// "github.com/HosseinRouhi79/golang-clean-web-api/src/data/db/migrations"
)

func main() {

	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	cache.InitRedis(cfg)
	err := db.InitDB(cfg)
	if err != nil {
		panic("cannot connect to database(main error)")
	}
	// migrations.Up_1()
	api.InitServer(cfg)
	defer cache.CloseRedis()
	defer db.CloseDB()
}

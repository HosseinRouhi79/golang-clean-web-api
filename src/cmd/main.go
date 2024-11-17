package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/api"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/cache"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/db"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	go func() {
		log.Println("Starting pprof on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()

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

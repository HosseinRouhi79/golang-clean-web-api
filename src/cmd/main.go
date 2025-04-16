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
// @title           My API
// @version         1.0
// @description     This is my API server.
// @contact.name    API Support
// @contact.url     http://www.example.com/support
// @contact.email   support@example.com
// @host            localhost:5005
// @BasePath        /api/v1
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
	err := db.InitDB(cfg)
	if err != nil {
		panic("cannot connect to database(main error)")
	}
	// migrations.Up_1()
	api.InitServer(cfg)
	defer cache.CloseRedis()
	defer db.CloseDB()
}

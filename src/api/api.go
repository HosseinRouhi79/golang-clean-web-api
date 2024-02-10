package api

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/routers"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/gin-gonic/gin"
	"fmt"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery()) // => r1 := gin.Default()

	v1 := r.Group("/api/v1/")
	{
		healthGroup := v1.Group("health")
		routers.Health(healthGroup)
	} 

	v2 := r.Group("/api/v1/")
	{
		testGroup := v2.Group("test")
		routers.Test(testGroup)
	}

	v3 := r.Group("/api/v3/")
	{
		formGroup := v3.Group("form")
		routers.BodyBinder(formGroup)
	}

	if err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort)); err != nil {
		panic(err)
	}

}

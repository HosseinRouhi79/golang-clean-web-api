package api

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/routers"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery()) // => r1 := gin.Default()

	v1 := r.Group("/api/v1/")
	{
		healthGroup := v1.Group("health")
		routers.Health(healthGroup)
	} 
	r.Run(":8081")

}

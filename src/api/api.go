package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery()) // => r1 := gin.Default()

	v1 := r.Group("/api/v1/")
	{
		v1.GET("health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "OK!")
		})
	}
	r.Run(":8081")

}

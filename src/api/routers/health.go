package routers

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealth()
	r.GET("/", handler.Health)
	r.POST("/", handler.HealthPost)
	r.GET("/:id", handler.HealthPostByID)
}

func Test(r *gin.RouterGroup){
	handler := handlers.NewTest()
	r.GET("/", handler.HeaderBind)
	r.GET("/query", handler.QueryBind)
	r.GET("/query2/:id/:name", handler.UriBind)
}

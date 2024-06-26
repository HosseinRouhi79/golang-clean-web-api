package routers

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/handlers"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/middlewares"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/gin-gonic/gin"
)

func Country(r *gin.RouterGroup) {
	handler := handlers.Country{}
	handlerDelete := handlers.CountryDelete{}
	cfg := config.GetConfig()

	r.POST("/c/create", middlewares.Authentication(cfg), handler.Create)
	r.DELETE("/c/delete", middlewares.Authentication(cfg), handlerDelete.Delete)
}

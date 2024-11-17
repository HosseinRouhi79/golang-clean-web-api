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
	handlerUpdate := handlers.CountryUpdate{}
	handlerAssign := handlers.CitiesAssined{}
	cfg := config.GetConfig()

	r.POST("/c/create", middlewares.Authentication(cfg), handler.Create)
	r.PUT("/c/update", middlewares.Authentication(cfg), handlerUpdate.Update)
	r.PUT("/c/delete", middlewares.Authentication(cfg), handlerDelete.Delete)
	r.GET("/c/get/:id", middlewares.Authentication(cfg), handlers.GetByID)
	r.GET("/c/get/all", middlewares.Authentication(cfg), handlers.GetAllCountries)
	r.POST("/c/assign/cities", handlerAssign.AssignCityToCountry)
}

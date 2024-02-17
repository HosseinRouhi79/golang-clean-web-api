package api

import (
	"fmt"
	"log"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/middlewares"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/routers"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/validation"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.MobileValidator, true)
		err2 := val.RegisterValidation("password", validation.PasswordValidator, true)
		if err != nil {
			log.Print(err.Error())
		}
		if err2 != nil {
			log.Print(err.Error())
		}
	}
	r.Use(gin.Logger(), gin.Recovery(), middlewares.Limitter()) // => r1 := gin.Default()

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

	RegisterSwagger(r, cfg)
	if err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort)); err != nil {
		panic(err)
	}

}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("192.168.59.133:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

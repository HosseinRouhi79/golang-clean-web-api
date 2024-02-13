package api

import (
	"fmt"
	"log"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/routers"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/validation"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()
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

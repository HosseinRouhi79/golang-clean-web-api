package routers

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/handlers"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/middlewares"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.RouterGroup) {
	cfg := config.GetConfig()
	handler := handlers.AuthMobile{}
	handler2 := handlers.AuthRegister{}
	handler3 := handlers.UserHandler{}
	tokenHandler := handlers.TokenHandler{}
	r.POST("/register-login-mobile", handler.RLMobile)
	r.POST("/claim",middlewares.Authentication(cfg), tokenHandler.GetClaims)
	r.POST("/register", handler2.Register)
	r.POST("/send-otp", handler3.SendOtp)
	// r.POST("/refresh", handler.Refresh)
}

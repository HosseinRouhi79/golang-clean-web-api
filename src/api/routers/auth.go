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
	tokenHandler := handlers.TokenHandler{}
	r.POST("/register-login-mobile", handler.RLMobile)
	r.POST("/claim",middlewares.Authentication(cfg), tokenHandler.GetClaims)
	// r.POST("/register", handler.Register)
	// r.POST("/logout", handler.Logout)
	// r.POST("/refresh", handler.Refresh)
}

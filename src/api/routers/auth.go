package routers

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.RouterGroup) {
	handler := handlers.AuthMobile{}
	tokenHandler := handlers.TokenHandler{}
	r.POST("/register-login-mobile", handler.RLMobile)
	r.POST("/claim", tokenHandler.GetClaims)
	// r.POST("/register", handler.Register)
	// r.POST("/logout", handler.Logout)
	// r.POST("/refresh", handler.Refresh)
}

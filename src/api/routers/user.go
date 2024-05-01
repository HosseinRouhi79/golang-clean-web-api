package routers

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetOtp(r *gin.RouterGroup) {
	handler := handlers.UserHandler{}
	r.POST("/send-otp", handler.SendOtp)
}
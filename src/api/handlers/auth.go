package handlers

import (
	"net/http"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/dto"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/services"
	"github.com/gin-gonic/gin"
)

var cfg = config.GetConfig()

type AuthMobile struct {
	Mobile string
	Otp string
}

func (auth AuthMobile) RLMobile(c *gin.Context) {

	am := AuthMobile{}
	err := c.ShouldBind(&am)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	dto := dto.RegisterLoginByMobileDto{
		Mobile: am.Mobile,
		Otp: am.Otp,
	}
	
	userService := services.NewUserService(cfg)
	tokenDetail, err := userService.RegisterLoginByMobile(dto)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"token":  tokenDetail,
	})
}
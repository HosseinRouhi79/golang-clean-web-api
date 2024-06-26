package handlers

import (
	"net/http"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/dto"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/services"
	"github.com/gin-gonic/gin"
)

var cfg = config.GetConfig()

type TokenHandler struct {
	Token string `form:"token"`
}
type AuthMobile struct {
	Mobile string `form:"mobile"`
	Otp    string `form:"otp"`
}

// User_Auth godoc
// @Summary User Auth
// @Description Register Login
// @Tags auth
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param mobile formData string true "Mobile number"
// @Param otp formData string false "OTP"
// @Success 200 {object} helper.HTTPResponse "Success"
// @Failure 400 {object} helper.HTTPResponse "Failed"
// @Router /register-login-mobile/ [post]
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
		Otp:    am.Otp,
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

func (t TokenHandler) GetClaims(c *gin.Context) {
	err := c.ShouldBind(&t)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	tokenService := services.NewTokenService(cfg)
	mpClaims, _ := tokenService.GetClaims(t.Token)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"claims": mpClaims,
	})
}

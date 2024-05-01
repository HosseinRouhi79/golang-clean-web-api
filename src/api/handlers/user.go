package handlers

import (
	"strconv"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/helper"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/services"
	"github.com/gin-gonic/gin"
)

type OtpRequest struct {
	Mobile string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
}

type UserHandler struct {
	UserService services.UserService
	OtpService  services.OtpService
}

func (h UserHandler) SendOtp(c *gin.Context) {
	cfg := config.GetConfig()
	// logger := logging.NewLogger(cfg)
	h = UserHandler{
		OtpService:  *services.NewOtpService(cfg),
		UserService: *services.NewUserService(cfg),
	}
	req := OtpRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.UserService.Logger.Infof("error has occured: %s", err.Error())
		return
	}
	otpCode := helper.GenerateOtp()
	err = h.OtpService.SetOtp(req.Mobile, strconv.Itoa(otpCode))

	if err != nil {
		h.UserService.Logger.Infof("can not set otp code: %s", err.Error())
		return
	}

	// send OTP SMS
	c.JSON(200, gin.H{
		"status": true,
	})
}

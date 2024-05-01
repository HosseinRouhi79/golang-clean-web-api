package services

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/db"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct{
	Logger logging.Logger
	Cfg *config.Config
	Otp OtpService
	Db *gorm.DB
}


func NewUserService(cfg *config.Config) (*UserService) {
	logger := logging.NewLogger(cfg)
	otp := NewOtpService(cfg)
	db := db.GetDB()
	u := &UserService{
		Logger: logger,
		Otp: *otp,
		Db: db,
		Cfg: cfg,
	}
	return u
}
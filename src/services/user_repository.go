package services

import (
	"errors"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/models"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"gorm.io/gorm"
)

var cfg = config.GetConfig()
var logger = logging.NewLogger(cfg)

func (userService *UserService) ExistMobile(mobile string) (error, bool) {
	model := models.User{}
	err := userService.Db.Where("mobile_number = ?", mobile).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Info(logging.Postgres, logging.Api, "Record not found", nil)
			return nil, false
		}
		logger.Info(logging.Postgres, logging.Api, err.Error(), nil)
		return err, false
		
	}

	return nil, true
}

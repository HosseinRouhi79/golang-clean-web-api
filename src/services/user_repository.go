package services

import (
	"errors"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/constants"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/models"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"gorm.io/gorm"
)

var cfg = config.GetConfig()
var logger = logging.NewLogger(cfg)

func (userService *UserService) ExistBytMobile(mobile string) (error, bool) {
	model := models.User{}
	err := userService.Db.Table("users").Where("mobile_number = ?", mobile).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Info(logging.Postgres, logging.Api, "Record(mobile_number) not found", nil)
			return nil, false
		}
		logger.Info(logging.Postgres, logging.Api, err.Error(), nil)
		return err, false
		
	}

	return nil, true
}

func (userService *UserService) ExistByEmail(email string) (error, bool) {
	model := models.User{}
	err := userService.Db.Table("users").Where("email = ?", email).First(&model).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
            logger.Info(logging.Postgres, logging.Api, "Record(email) not found", nil)
            return nil, false
        }
        logger.Info(logging.Postgres, logging.Api, err.Error(), nil)
        return err, false
	}
	return nil, true
} 

func (userService *UserService) ExistByUsername(username string) (error, bool) {
	model := models.User{}
	err := userService.Db.Table("users").Where("username =?", username).First(&model).Error
	if err!= nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
            logger.Info(logging.Postgres, logging.Api, "Record(username) not found", nil)
            return nil, false
        }
        logger.Info(logging.Postgres, logging.Api, err.Error(), nil)
        return err, false
	}
	return nil, true
} 

func (userService *UserService) GetDefaultRole(username string) (err error, roleID int) {

	role := models.Role{}
	err = userService.Db.Table("roles").Where("name =?", constants.DefaultRole).First(&role).Error
	if err!= nil{
		userService.Logger.Info(logging.Postgres, logging.DefaultRoleNotFound, "Role not found", nil)
		return err, 0
	}
	return nil, role.Id

}


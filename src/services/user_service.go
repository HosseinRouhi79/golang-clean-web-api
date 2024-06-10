package services

import (
	"errors"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/db"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/models"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/dto"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Logger logging.Logger
	Cfg    *config.Config
	Otp    OtpService
	Db     *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	otp := NewOtpService(cfg)
	db := db.GetDB()
	u := &UserService{
		Logger: logger,
		Otp:    *otp,
		Db:     db,
		Cfg:    cfg,
	}
	return u
}

//Register
//Register/Login by mobile & otp
//Login by username & password

func (userService *UserService) Register(dto dto.RegisterationDto) error {

	model := models.User{}
	model.FirstName = dto.FirstName
	model.LastName = dto.LastName
	model.Username = dto.Username
	model.Email = dto.Email

	err, val := userService.ExistByEmail(model.Email)
	if err != nil {
		return err
	}
	if val {
		return errors.New("email already exist")
	}

	err, val = userService.ExistByUsername(model.Username)
	if err!= nil {
        return err
    }
	if val {
        return errors.New("username already exist")
    }

	err, roleID := userService.GetDefaultRole(model.Username)
	if err!= nil {
		userService.Logger.Info(logging.Postgres, logging.DefaultRoleNotFound, "Role not found",nil)
        return err
    }

	bp := []byte(model.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err!= nil {
        userService.Logger.Info(logging.Postgres, logging.HashPassword, "error in creating model(pass)", nil)
        return err
    }

	model.Password = string(hp)

	tx := userService.Db.Begin()
	err = tx.Create(&model).Error
	if err!= nil {
        tx.Rollback()
		userService.Logger.Info(logging.Postgres, logging.Select, "error in creating model", nil)
        return err
    }
	err = tx.Create(&models.UserRole{RoleId: roleID, UserId: model.Id}).Error
	if err!= nil {
		tx.Rollback()
        userService.Logger.Info(logging.Postgres, logging.Select, "error in creating user-role", nil)
        return err
	}
	tx.Commit()
	return nil

}

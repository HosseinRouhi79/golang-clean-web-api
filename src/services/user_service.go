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
	Token  TokenService
	Db     *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	otp := NewOtpService(cfg)
	token := NewTokenService(cfg)
	db := db.GetDB()
	u := &UserService{
		Logger: logger,
		Otp:    *otp,
		Db:     db,
		Cfg:    cfg,
		Token:  *token,
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
	if err != nil {
		return err
	}
	if val {
		return errors.New("username already exist")
	}

	roleID, err := userService.GetDefaultRole(model.Username)
	if err != nil {
		userService.Logger.Info(logging.Postgres, logging.DefaultRoleNotFound, "Role not found", nil)
		return err
	}

	bp := []byte(model.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		userService.Logger.Info(logging.Postgres, logging.HashPassword, "error in creating model(pass)", nil)
		return err
	}

	model.Password = string(hp)

	tx := userService.Db.Begin()
	err = tx.Create(&model).Error
	if err != nil {
		tx.Rollback()
		userService.Logger.Info(logging.Postgres, logging.Select, "error in creating model", nil)
		return err
	}
	err = tx.Create(&models.UserRole{RoleId: roleID, UserId: model.Id}).Error
	if err != nil {
		tx.Rollback()
		userService.Logger.Info(logging.Postgres, logging.Select, "error in creating user-role", nil)
		return err
	}
	tx.Commit()
	return nil

}

func (userService *UserService) RegisterLoginByMobile(dto dto.RegisterLoginByMobileDto) (tokenDetail *TokenDetail, err error) {
	err = userService.Otp.ValidateOtp(dto.Mobile, dto.Otp)
	if err != nil {
		userService.Logger.Info(logging.OTP, logging.Select, err.Error(), nil)
		return nil, err
	}
	err, exists := userService.ExistBytMobile(dto.Mobile)

	user := models.User{Mobile: dto.Mobile, Username: dto.Mobile}

	if !exists {

	} else {
		err := userService.Db.Table("users").
			Where("mobile_number = ?", dto.Mobile).
			Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
				return tx.Preload("Role")
			}).
			First(&user).Error

		if err != nil {
			return nil, err
		}

		tdto := TokenDto{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email,
			UserName: user.Username, UserID: user.Id}
		roles := []string{}
		if len(*user.UserRoles) > 0 {
			for _, ur := range *user.UserRoles {
				roles = append(roles, ur.Role.Name)
			}
		}
		tdto.Roles = roles
		tokenDetail, err := userService.Token.GenerateToken(&tdto)
		if err!= nil {
            userService.Logger.Info(logging.Internal, logging.Select, err.Error(), nil)
            return nil, err
        }

		return tokenDetail, nil
	}
	if err != nil {
		userService.Logger.Info(logging.Internal, logging.Select, err.Error(), nil)
		return nil, err
	}
	return nil, nil
}

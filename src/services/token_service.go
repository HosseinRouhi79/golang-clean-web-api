package services

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
)

type TokenService struct {
	Logger logging.Logger
	Cfg    *config.Config
}

type TokenDto struct {
	UserID       int
	FirstName    string
	LastName     string
	UserName     string
	Email        string
	MobileNumber string
	Roles        []string
}

func NewTokenService(cfg *config.Config) (*TokenService) {
	logger := logging.NewLogger(cfg)
	u := &TokenService{
		Logger: logger,
		Cfg: cfg,
	}
	return u
}
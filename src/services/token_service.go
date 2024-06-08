package services

import (
	"time"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"github.com/golang-jwt/jwt/v5"
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

type TokenDetail struct {
	AccessToken                string `json:"accessToken"`
	RefreshToken               string `json:"refreshToken"`
	AccessTokenExpirationTime  int64  `json:"accessTokenExpirationTime"`
	RefreshTokenExpirationTime int64  `json:"refreshTokenExpirationTime"`
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)
	u := &TokenService{
		Logger: logger,
		Cfg:    cfg,
	}
	return u
}

func (tokenService *TokenService) GenerateToken(tokenDto *TokenDto) (*TokenDetail, error) {
	token := TokenDetail{}
	token.AccessTokenExpirationTime = time.Now().Add(tokenService.Cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	token.RefreshTokenExpirationTime = time.Now().Add(tokenService.Cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	atc := jwt.MapClaims{} //access token claims


	return nil, nil
}

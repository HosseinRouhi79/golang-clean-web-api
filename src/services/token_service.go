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

	dto := jwt.MapClaims{} //access token claims

	dto["id"] = tokenDto.UserID
	dto["firstName"] = tokenDto.FirstName
	dto["lastName"] = tokenDto.LastName
	dto["userName"] = tokenDto.UserName
	dto["email"] = tokenDto.Email
	dto["mobileNumber"] = tokenDto.MobileNumber
	dto["roles"] = tokenDto.Roles
	dto["exp"] = token.AccessTokenExpirationTime

	atc := jwt.NewWithClaims(jwt.SigningMethodHS256, dto)
	var err error
	token.AccessToken, err = atc.SignedString([]byte(tokenService.Cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}

	dtor := jwt.MapClaims{} //refresh token claims
	dtor["id"] = tokenDto.UserID
	dtor["exp"] = token.RefreshTokenExpirationTime

	rtc := jwt.NewWithClaims(jwt.SigningMethodHS256, dtor)

	token.RefreshToken, err = rtc.SignedString([]byte(tokenService.Cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}
	return &token, nil
}

package middlewares

import (
	"strings"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/constants"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var claimsMap map[string]interface{}
		tokenService := services.NewTokenService(cfg)
		auth := c.GetHeader(constants.Authorization)

		token := strings.Split(auth, " ")
		accessToken := token[1]

		claimsMap, err = tokenService.GetClaims(accessToken)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				tokenService.Logger.Infof("token is expired: %s", err.Error())
			default:
				tokenService.Logger.Infof("error has occurred: %s", err.Error())
			}
		}

	}
}

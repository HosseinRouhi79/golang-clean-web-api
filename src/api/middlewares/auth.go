package middlewares

import (
	"fmt"
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
		if auth == "" {
			err = fmt.Errorf("auth header doesnt exist")
		} else {
			token := strings.Split(auth, " ")
			accessToken := token[1]

			claimsMap, err = tokenService.GetClaims(accessToken)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = fmt.Errorf("token is expired")
				default:
					err = fmt.Errorf("error has occured")
				}
			}
		}
		if err == nil {
			c.Set("id", claimsMap["id"])
			c.Set("userName", claimsMap["userName"])
			c.Set("firstName", claimsMap["firstName"])
			c.Set("lastName", claimsMap["lastName"])
			c.Set("mobileNumber", claimsMap["mobileNumber"])
			c.Set("email", claimsMap["email"])
			c.Set("roles", claimsMap["roles"])
			c.Set("exp", claimsMap["exp"])

			c.Next()
		}
	}
}

func Authorization(validRoles []string) gin.HandlerFunc{
	return func(c *gin.Context) {
		var err error
		var roleList []string
		roleval, ok := c.Keys["roles"]
		if !ok{
			err = fmt.Errorf("error has occured")
		} else{
			roleList = append(roleList, roleval.([]string)...)
		}
		if err == nil{
			for _, role := range roleList{
				for _, vRole := range validRoles{
					if role == vRole{
						c.Next()
					}
				}
			}
		}
	}
}
package middlewares

import (
	"time"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"github.com/gin-gonic/gin"
)

func StructuredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetConfig()
		logger := logging.NewLogger(cfg)
		start := time.Now()
		url := c.Request.URL

		c.Next()
		
		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now() // stop
		param.Latency = param.TimeStamp.Sub(start)

		keyMap := map[logging.ExtraKey]interface{}{
			logging.Path:    url,
			logging.Latency: param.Latency,
		}

		logger.Info(logging.RequestResponse, logging.Api, "", keyMap)
	}
}

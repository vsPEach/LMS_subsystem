package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Logger interface {
	Error(...any)
	Info(...any)
	Warn(...any)
	Infof(template string, args ...any)
	Errorf(template string, args ...any)
}

func Logging(logger Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()

		logger.Infof("IP: %s, Latency: %s, Path: %s, Raw: %s", param.ClientIP, param.Latency, path, raw)
	}
}

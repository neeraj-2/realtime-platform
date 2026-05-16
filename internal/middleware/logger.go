package middleware


import (
	"time"
	"realtime-platform/internal/config"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)

		config.Log.Info(
			"Incoming request",
		)

		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		config.Log.Sugar().Infof("%s %s - %d (%s)", method, path, statusCode, duration)
	}
}
package middlewares

import (
	"time"

	"github.com/TimX-21/auth-service-go/internal/config"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)

		fields := map[string]any{
			"method": c.Request.Method,
			"uri":    c.Request.RequestURI,
		}

		if lastErr := c.Errors.Last(); lastErr != nil {
			config.Log.Errorw("Error",
				"method", fields["method"],
				"uri", fields["uri"],
				"error", lastErr,
			)
			c.Abort()
			return
		}

		config.Log.Infow("Info",
			"method", c.Request.Method,
			"uri", c.Request.RequestURI,
			"status", c.Writer.Status(),
			"duration", duration,
		)

	}
}

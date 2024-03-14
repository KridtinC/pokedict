package handler

import (
	"time"

	"github.com/KridtinC/pokedict/internal/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		logger.Get().Info("Request received", zapcore.Field{
			Key:    "method",
			Type:   zapcore.StringType,
			String: c.Request.Method,
		},
			zapcore.Field{
				Key:    "path",
				Type:   zapcore.StringType,
				String: c.Request.URL.Path,
			},
		)
		c.Next()
		logger.Get().Info("Request completed", zapcore.Field{
			Key:    "duration",
			Type:   zapcore.StringType,
			String: time.Since(start).String(),
		})
	}
}

package middleware

import (
	"fmt"
	"time"

	authMiddle "github.com/3-shake/reckoner-dmh-server/auth/handler/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type middle struct {
	authMiddle.MiddlewareAuthHandler
}

type MiddlewareImpl interface {
	Logging(conf *zap.Config) gin.HandlerFunc
	Authenticate() gin.HandlerFunc
}

func NewMiddleware(authHandler authMiddle.MiddlewareAuthHandler) MiddlewareImpl {
	return &middle{authHandler}
}

func (this *middle) Logging(conf *zap.Config) gin.HandlerFunc {
	logger, _ := conf.Build()
	return func(c *gin.Context) {
		start := time.Now().UTC()
		c.Set("logger", logger)
		c.Next()

		latency := time.Since(start)

		path := c.Request.URL.Path
		method := c.Request.Method
		logger.Info(
			fmt.Sprintf("%s %s", c.Request.Method, path),
			zap.String("prefix", "ADN-API"),
			zap.String("path", path),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Int("status", c.Writer.Status()),
			zap.Time("time", start),
			zap.String("method", method),
			zap.Duration("latency", latency),
			zapFieldStringsByStringMap("header", c.Request.Header),
			zapFieldStringsByStringMap("query", c.Request.URL.Query()),
		)

		if len(c.Errors) > 0 {
			for _, errMsg := range c.Errors.Errors() {
				logger.Error(errMsg)
			}
		}
	}
}

func zapFieldStringsByStringMap(key string, m map[string][]string) zapcore.Field {
	return zap.Object(key, stringsByStringMarshaler(m))
}

func stringsByStringMarshaler(m map[string][]string) zapcore.ObjectMarshalerFunc {
	return func(inner zapcore.ObjectEncoder) error {
		for k, values := range m {
			err := inner.AddArray(k, zapcore.ArrayMarshalerFunc(func(inner zapcore.ArrayEncoder) error {
				for _, v := range values {
					inner.AppendString(v)
				}
				return nil
			}))
			if err != nil {
				return err
			}
		}
		return nil
	}
}

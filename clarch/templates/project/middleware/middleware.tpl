package middleware

import (
	"fmt"
	"time"

	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/project/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type middle struct {
	// authMiddle.MiddlewareAuthHandler
}

type MiddlewareImpl interface {
	Cors() gin.HandlerFunc
	Logging() gin.HandlerFunc
	// Authenticate() gin.HandlerFunc
}

func NewMiddleware() MiddlewareImpl {
	return &middle{}
}

func (this *middle) Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AddAllowMethods("DELETE")
	config.AddAllowHeaders("Access-Control-Allow-Origin", "Authorization")
	config.AllowAllOrigins = true
	return cors.New(config)
}

func (this *middle) Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		c.Next()

		latency := time.Since(start)

		path := c.Request.URL.Path
		method := c.Request.Method
		logger.Logger.Info(
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

// pkg/middleware/logger.go

package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	RequestIDKey = "RequestID"
)

// Logger represents the logger interface used by the middleware
type Logger interface {
	Named(name string) Logger
	With(args ...interface{}) Logger
	WithContext(ctx context.Context) Logger
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

// RequestLogger middleware for Gin that logs incoming HTTP requests
func RequestLogger(log Logger) gin.HandlerFunc {
	logger := log.Named("gin-middleware")

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		requestID := uuid.New().String()
		c.Set(RequestIDKey, requestID)

		ctx := context.WithValue(c.Request.Context(), RequestIDKey, requestID)
		c.Request = c.Request.WithContext(ctx)

		reqLogger := logger.WithContext(ctx).With(
			"request_id", requestID,
			"method", c.Request.Method,
			"path", path,
			"client_ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
		)
		reqLogger.Info("Incoming request")

		c.Next()
		duration := time.Since(start)

		logFields := []interface{}{
			"status", c.Writer.Status(),
			"duration", duration,
			"bytes", c.Writer.Size(),
		}
		if raw != "" {
			logFields = append(logFields, "query", raw)
		}
		if len(c.Errors) > 0 {
			logFields = append(logFields, "errors", c.Errors.String())
		}

		switch {
		case c.Writer.Status() >= 500:
			reqLogger.Error("Server error", logFields...)
		case c.Writer.Status() >= 400:
			reqLogger.Warn("Client error", logFields...)
		case c.Writer.Status() >= 300:
			reqLogger.Info("Redirect", logFields...)
		default:
			reqLogger.Info("Request completed", logFields...)
		}
	}
}

// Recovery middleware that logs panics
func Recovery(log Logger) gin.HandlerFunc {
	logger := log.Named("gin-recovery")

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				requestID, _ := c.Get(RequestIDKey)

				logger.Error("Panic recovered",
					"error", err,
					"request_id", requestID,
					"method", c.Request.Method,
					"path", c.Request.URL.Path,
					"client_ip", c.ClientIP(),
				)
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}

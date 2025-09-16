package middleware

import (
	"context"
	"go-commerce/infrastucture/log"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log request details here (method, path, headers, etc.)
		requestID := uuid.New().String()
		timeoutCtx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
		defer cancel()

		ctx := context.WithValue(timeoutCtx, "request_id", requestID)
		c.Request = c.Request.WithContext(ctx)

		startTime := time.Now()
		c.Next()
		latency := time.Since(startTime)

		requestLog := logrus.Fields{
			"request_id": requestID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     c.Writer.Status(),
			"latency":    latency,
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}

		if c.Writer.Status() == 200 || c.Writer.Status() == 201 {
			log.Logger.WithFields(requestLog).Info("request completed")
		} else { 
			log.Logger.WithFields(requestLog).Error("request error!")
		}
	}
}

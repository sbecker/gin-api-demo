package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
)

type RequestIDOptions struct {
	AllowSetting bool
}

func RequestID(options RequestIDOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string

		if options.AllowSetting {
			// If Set-Request-Id header is set on request, use that for
			// Request-Id response header. Otherwise, generate a new one.
			requestID = c.Request.Header.Get("Set-Request-Id")
		}

		if requestID == "" {
			requestID = uuid.New()
		}

		c.Writer.Header().Set("Request-Id", requestID)
		c.Next()
	}
}

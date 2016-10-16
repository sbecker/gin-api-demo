package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
)

func RequestID(c *gin.Context) {
	// If Set-Request-Id header is set on request, use that for
	// Request-Id response header. Otherwise, generate a new one.
	requestID := c.Request.Header.Get("Set-Request-Id")
	if requestID == "" {
		requestID = uuid.New()
	}
	c.Writer.Header().Set("Request-Id", requestID)
	c.Next()
}

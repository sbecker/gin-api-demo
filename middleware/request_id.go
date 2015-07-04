package middleware

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/gin-gonic/gin"
)

func RequestId(c *gin.Context) {
	// If Set-Request-Id header is set on request, use that for
	// Request-Id response header. Otherwise, generate a new one.
	requestId := c.Request.Header.Get("Set-Request-Id")
	if requestId == "" {
		requestId = uuid.New()
	}
	c.Writer.Header().Set("Request-Id", requestId)
	c.Next()
}

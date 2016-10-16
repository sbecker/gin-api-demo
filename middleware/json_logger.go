package middleware

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/sbecker/gin-api-demo/util"
)

// JSONLogMiddleware logs a gin HTTP request in JSON format, with some additional custom key/values
func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := util.GetDurationInMillseconds(start)

		entry := log.WithFields(log.Fields{
			"client_ip":  util.GetClientIP(c),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"user_id":    util.GetUserID(c),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
			// "api_version": util.ApiVersion,
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}

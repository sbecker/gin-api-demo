package middleware

import "github.com/gin-gonic/gin"

// from https://github.com/gin-gonic/gin/issues/29#issuecomment-89132826
func CORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // allow any origin domain
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://domain.com") // uncomment to restrict to single domain
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	} else {
		c.Next()
	}
}

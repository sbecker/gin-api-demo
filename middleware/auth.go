package middleware

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/sbecker/gin-api-demo/dao"
)

func Auth(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	r, _ := regexp.Compile("^Bearer (.+)$")

	match := r.FindStringSubmatch(authHeader)

	if len(match) == 0 {
		c.AbortWithStatus(401)
		return
	}
	tokenString := match[1]

	if len(tokenString) == 0 {
		c.AbortWithStatus(401)
		return
	}

	user, err := dao.GetUserByAuthToken(tokenString)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	c.Set("currentUser", user)
	c.Next()
}

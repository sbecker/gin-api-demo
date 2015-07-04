package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/sbecker/gin-api-demo/dao"
	"github.com/sbecker/gin-api-demo/serializers"
)

func GetAllUsers(c *gin.Context) {
	currentUser := getCurrentUser(c)
	users := dao.GetAllUsers(currentUser)
	c.JSON(200, serializers.SerializeUsers(users, currentUser))
}

func GetUserById(c *gin.Context) {
	currentUser := getCurrentUser(c)

	id, err := getStringParam(c, "id")
	if err != nil {
		c.JSON(400, "could not parse id")
		return
	}

	user, err := dao.GetUserById(id, currentUser)
	if err != nil {
		c.JSON(404, "Not Found")
		return
	}

	c.JSON(200, serializers.SerializeUser(user, currentUser))
}

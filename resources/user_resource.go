package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/sbecker/gin-api-demo/dao"
	"github.com/sbecker/gin-api-demo/serializers"
)

type UserResource struct {
}

func NewUserResource(e *gin.Engine) {
	u := UserResource{}

	// Setup Routes
	e.GET("/users", u.getAllUsers)
	e.GET("/users/:id", u.getUserByID)
}

func (r *UserResource) getAllUsers(c *gin.Context) {
	currentUser := getCurrentUser(c)
	users := dao.GetAllUsers(currentUser)
	c.JSON(200, serializers.SerializeUsers(users, currentUser, "/users"))
}

func (r *UserResource) getUserByID(c *gin.Context) {
	currentUser := getCurrentUser(c)

	id, err := getStringParam(c, "id")
	if err != nil {
		c.JSON(400, "could not parse id")
		return
	}

	user, err := dao.GetUserByID(id, currentUser)
	if err != nil {
		c.JSON(404, "Not Found")
		return
	}

	c.JSON(200, serializers.SerializeUser(user, currentUser))
}

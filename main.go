package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sbecker/gin-api-demo/dao"
	"github.com/sbecker/gin-api-demo/middleware"
	"github.com/sbecker/gin-api-demo/resources"
)

func main() {
	dao.InitDb()

	r := gin.Default()
	r.Use(middleware.RequestId)
	r.Use(middleware.Auth)
	r.Use(middleware.CORS)

	r.GET("/users", resources.GetAllUsers)
	r.GET("/users/:id", resources.GetUserById)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

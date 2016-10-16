package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sbecker/gin-api-demo/dao"
	"github.com/sbecker/gin-api-demo/middleware"
	"github.com/sbecker/gin-api-demo/resources"
	"github.com/sbecker/gin-api-demo/util"
)

func main() {
	dao.InitDb()

	util.UseJSONLogFormat()
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(middleware.JSONLogMiddleware())
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID(middleware.RequestIDOptions{AllowSetting: false}))
	r.Use(middleware.Auth())
	r.Use(middleware.CORS(middleware.CORSOptions{}))

	resources.NewUserResource(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

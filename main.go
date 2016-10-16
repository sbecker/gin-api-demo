package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/sbecker/gin-api-demo/dao"
	"github.com/sbecker/gin-api-demo/middleware"
	"github.com/sbecker/gin-api-demo/resources"
	"github.com/sbecker/gin-api-demo/util"
)

func main() {
	util.LoadEnvVars()
	util.UseJSONLogFormat()

	dao.InitDb()

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(middleware.JSONLogMiddleware())
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID(middleware.RequestIDOptions{AllowSetting: false}))
	r.Use(middleware.Auth())
	r.Use(middleware.CORS(middleware.CORSOptions{}))

	resources.NewUserResource(r)

	port := util.GetEnv("PORT", "8080")
	log.Info("Service starting on port " + port)

	r.Run(":" + port) // listen and serve
}

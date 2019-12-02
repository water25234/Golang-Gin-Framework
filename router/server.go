package router

import (
	"github.com/gin-gonic/gin"

	coreserver "github.com/water25234/Golang-Gin-Framework/core/server"
	"github.com/water25234/Golang-Gin-Framework/server"
)

var Router *gin.Engine

func init() {
	coreserver.SetServerGonfig()
	coreserver.SetAppConfig()
	server.InitRedis()
}

func StartServer() {
	Router = SetupRouter()
	Router.Run()
}

package router

import (
	"github.com/gin-gonic/gin"
	coreserver "github.com/water25234/Golang-Gin-Framework/core/server"
)

var Router *gin.Engine

func init() {
	coreserver.SetServerGonfig()
}

func StartServer() {
	Router = SetupRouter()
	Router.Run()
}

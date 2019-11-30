package router

import (
	apiv1 "../api/v1"
	apiv1user "../api/v1/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	helloRouting := router.Group("/hello")
	{
		helloRouting.GET("", apiv1.GetHello)

		helloRouting.DELETE("/:id", apiv1.DeleteHello)
	}

	userRouting := router.Group("/user")
	{
		userRouting.GET("/:uid", apiv1user.GetUser)
	}

	return router
}

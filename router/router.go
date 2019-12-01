package router

import (
	apiv1auth "../api/v1/auth"
	apiv1user "../api/v1/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	authRouting := router.Group("/auth")
	{
		authRouting.GET("", apiv1auth.GetAuth)

		authRouting.DELETE("/:id", apiv1auth.DeleteAuth)

		authRouting.POST("/:uid", apiv1auth.PostAuth)
	}

	userRouting := router.Group("/user")
	{
		userRouting.GET("/:uid", apiv1user.GetUser)
	}

	return router
}

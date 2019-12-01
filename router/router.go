package router

import (
	"github.com/gin-gonic/gin"
	apiv1auth "github.com/water25234/Golang-Gin-Framework/api/v1/auth"
	apiv1user "github.com/water25234/Golang-Gin-Framework/api/v1/user"
	"github.com/water25234/Golang-Gin-Framework/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	gin.New().Use(gin.Logger())

	gin.New().Use(gin.Recovery())

	gin.New().Group("api").Use(middleware.ExecuteThrottle())
	{
		authRouting := router.Group("api/v1/auth")
		{
			authRouting.GET("", apiv1auth.GetAuth)

			authRouting.DELETE("/:id", apiv1auth.DeleteAuth)

			authRouting.POST("/:uid", apiv1auth.PostAuth)
		}

		userRouting := router.Group("api/v1/user")
		{
			userRouting.GET("/:uid", apiv1user.GetUser)
		}
	}

	return router
}

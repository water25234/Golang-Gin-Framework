package router

import (
	"github.com/gin-gonic/gin"
	apiv1auth "github.com/water25234/Golang-Gin-Framework/api/v1/auth"
	apiv1user "github.com/water25234/Golang-Gin-Framework/api/v1/user"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//middlewares.executeThrottle()

	//middleware.executeThrottle(recoveryHandler)
	//r := gin.New()

	//router = router.Group("api")
	//r.Use(middleware.executeThrottle(recoveryHandler))

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

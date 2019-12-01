package apiv1user

import (
	"net/http"

	api "github.com/water25234/Golang-Gin-Framework/api/v1"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	uid := ctx.Param("uid")

	ctx.JSON(http.StatusOK, api.GetSuccessResponse(gin.H{"userId": uid}))
}

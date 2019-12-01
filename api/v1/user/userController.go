package apiv1user

import (
	"net/http"

	api "github.com/water25234/Golang-Gin-Framework/api/v1"
	"github.com/water25234/Golang-Gin-Framework/core/log"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	log.Info("test get user log")
	ctx.JSON(http.StatusOK, api.GetSuccessResponse(gin.H{"userId": uid}))
}

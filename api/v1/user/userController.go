package apiv1user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	ctx.JSON(http.StatusOK, gin.H{"userId": uid})
}

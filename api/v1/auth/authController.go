package apiv1auth

import (
	"net/http"

	api ".."

	"github.com/gin-gonic/gin"
)

func GetAuth(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/plain", []byte("Hello, Justin Home!"))
}

func DeleteAuth(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "Hello World DELETE Justin %s", id)
}

func PostAuth(ctx *gin.Context) {
	uid := ctx.Param("uid")

	ctx.JSON(http.StatusOK, api.GetSuccessResponse(gin.H{"userId": uid}))
}

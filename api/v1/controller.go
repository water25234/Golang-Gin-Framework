package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHello(ctx *gin.Context) {
	ctx.Data(200, "text/plain", []byte("Hello, Justin Home!"))
}

func DeleteHello(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "Hello World DELETE Justin %s", id)
}

package api

import "github.com/gin-gonic/gin"

func GetErrorResponse(data interface{}, desc string) map[string]interface{} {
	return gin.H{
		"metadata": gin.H{
			"status": "9999",
			"desc":   desc,
		},
		"data": data,
	}
}

func GetSuccessResponse(data interface{}) map[string]interface{} {
	return gin.H{
		"metadata": gin.H{
			"status": "0000",
			"desc":   "success",
		},
		"data": data,
	}
}

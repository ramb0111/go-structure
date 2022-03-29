package utils

import "github.com/gin-gonic/gin"

// ResponseHandler for response
func ResponseHandler(ctx *gin.Context, status int, response gin.H) {
	ctx.JSON(status, response)
}

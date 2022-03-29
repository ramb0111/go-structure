package utils

import "github.com/gin-gonic/gin"

// ErrorHandler for error handling
func ErrorHandler(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError struct
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ValidationError(ctx *gin.Context, status int, response gin.H) {
	ctx.JSON(status, response)
}

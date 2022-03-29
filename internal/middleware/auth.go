package middleware

import (
	"fmt"
	"net/http"

	"github.com/TransportMall/payment-systems/internal/utils"

	"github.com/gin-gonic/gin"
)

//TokenAuthMiddleware middleware
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, ok := utils.TokenValid(c.Request)
		if ok == false {
			utils.ErrorHandler(c, http.StatusUnauthorized, fmt.Errorf("Invalid Authorization"))
			return
		}
		c.Set("auth", claim)
		c.Next()
	}
}

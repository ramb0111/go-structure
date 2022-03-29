package routes

import (
	"github.com/TransportMall/payment-systems/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetUpAdminRoutes(r *gin.Engine) {

	r.GET("/users", controller.GetUsers)
}

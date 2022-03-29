package routes

import (
	"github.com/TransportMall/payment-systems/internal/controller"
	"github.com/TransportMall/payment-systems/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetUpUserRoutes function for user related routes
func SetUpUserRoutes(r *gin.Engine) {

	api := r.Group("/api/user")
	{
		api.POST("/signup", controller.UserSignup)
		api.POST("/login", controller.UserLogin)

		api.Use(middleware.TokenAuthMiddleware())
		api.GET("/me", controller.MyProfile)
	}
}

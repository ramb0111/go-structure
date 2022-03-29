package api

import (
	"fmt"
	"os"

	"github.com/TransportMall/payment-systems/internal/api/routes"
	"github.com/TransportMall/payment-systems/internal/repository"
	"github.com/TransportMall/payment-systems/internal/utils"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//InitializeServer starts server
func InitializeServer() {

	db := repository.InitDB()
	utils.InitRedis()

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.GET("/health-checkup", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running...",
		})
	})

	routes.SetUpAdminRoutes(router)
	routes.SetUpUserRoutes(router)

	appPort := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", appPort))
}

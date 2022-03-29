package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TransportMall/payment-systems/internal/api"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error occured while loading .env file")
	}

	appName := os.Getenv("APP_NAME")
	appPort := os.Getenv("APP_PORT")
	fmt.Println(fmt.Sprintf("%s starting on http://localhost:%s", appName, appPort))

	api.InitializeServer()
}

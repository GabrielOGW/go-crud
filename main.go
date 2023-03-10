package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GabrielOGW/go-crud/src/configuration/logger"
	"github.com/GabrielOGW/go-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user aplication")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	fmt.Println(os.Getenv(("TEST")))

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
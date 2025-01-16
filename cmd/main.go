package main

import (
	"BookShopGoProject/config"
	"BookShopGoProject/internal/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	fmt.Println("Current Working Directory:", currentPath)
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error load .env file")
	}

	config.InitDB()
	if config.DB == nil {
		log.Fatal("Database connection was not initialized properly")
	}
	router := gin.Default()

	routes.RegisterRoutes(router, config.DB)

	if err := router.Run(":6969"); err != nil {
		log.Fatal(err)
	}
}

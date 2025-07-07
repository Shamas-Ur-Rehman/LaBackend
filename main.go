package main

import (
	"Laorgaincs/config"
	"Laorgaincs/middleware"
	"Laorgaincs/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	config.ConnectCloudinary()

	r := gin.Default()

	// Apply middleware
	r.Use(middleware.Logger())
	r.Use(middleware.CORSMiddleware())

	// Register routes
	routes.ProductRoutes(r)

	r.Run(":8080")
}

package main

import (
	"backend/config"
	"backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDB()

	// Seed the database with sample data
	config.SeedDatabase(db)

	router := gin.Default()
	routes.SetupHotelRoutes(router, db)
	routes.SetupUserRoutes(router, db)

	router.Run(":8080")
}

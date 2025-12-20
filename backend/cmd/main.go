package main

import (
	"backend/config"
	"backend/internal/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDB()

	config.SeedDatabase(db)

	router := gin.Default()
	routes.SetupHotelRoutes(router, db)
	routes.SetupUserRoutes(router, db)
	routes.SetupBookingRoutes(router, db)
	routes.SetupAuthRoutes(router, db)

	router.Run(":" + os.Getenv("SERVER_PORT"))
}

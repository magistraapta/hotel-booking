package main

import (
	"backend/config"
	_ "backend/docs" // Swagger documentation
	"backend/internal/routes"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Hotel Booking API
// @version         1.0
// @description     A RESTful API for hotel booking management system
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@hotelbooking.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	config.LoadEnv()
	db := config.ConnectDB()

	config.SeedDatabase(db)

	router := gin.Default()
	config.SetupCORS(router)

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupHotelRoutes(router, db)
	routes.SetupUserRoutes(router, db)
	routes.SetupBookingRoutes(router, db)
	routes.SetupAuthRoutes(router, db)

	router.Run(":" + os.Getenv("SERVER_PORT"))
}

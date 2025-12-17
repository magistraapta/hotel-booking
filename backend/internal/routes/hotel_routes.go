package routes

import (
	"backend/internal/controller"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupHotelRoutes(router *gin.Engine, db *gorm.DB) {
	hotelRepository := repository.NewHotelRepository(db)
	hotelService := service.NewHotelService(hotelRepository)
	hotelController := controller.NewHotelController(hotelService)

	hotelRouter := router.Group("/hotels")
	{
		hotelRouter.POST("/", hotelController.CreateHotel)
		hotelRouter.GET("/", hotelController.GetAllHotels)
		hotelRouter.GET("/:id", hotelController.GetHotelById)
	}
}

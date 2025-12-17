package routes

import (
	"backend/internal/controller"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBookingRoutes(router *gin.Engine, db *gorm.DB) {
	bookingRepository := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepository)
	bookingController := controller.NewBookingController(bookingService)

	bookingRouter := router.Group("/bookings")
	{
		bookingRouter.POST("/", bookingController.CreateBooking)
		bookingRouter.GET("/", bookingController.GetAllBookings)
		bookingRouter.GET("/:id", bookingController.GetBookingById)
	}
}

package routes

import (
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBookingRoutes(router *gin.Engine, db *gorm.DB) {

	bookingRepository := repository.NewBookingRepository(db)
	hotelRepository := repository.NewHotelRepository(db)
	bookingService := service.NewBookingService(hotelRepository, bookingRepository)
	bookingController := controller.NewBookingController(bookingService)

	bookingRouter := router.Group("/bookings")
	{
		bookingRouter.POST("/", middleware.RequireLogin(), bookingController.CreateBooking)
		bookingRouter.GET("/", bookingController.GetAllBookings)
		bookingRouter.GET("/:id", bookingController.GetBookingById)
		bookingRouter.GET("/user/:user_id", bookingController.GetBookingsByUserId)
	}
}

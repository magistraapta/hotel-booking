package controller

import (
	"backend/internal/domain"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bookingService service.BookingService
}

func NewBookingController(bookingService service.BookingService) *BookingController {
	return &BookingController{bookingService: bookingService}
}

func (c *BookingController) CreateBooking(ctx *gin.Context) {
	var booking domain.Booking
	ctx.ShouldBindJSON(&booking)
	err := c.bookingService.CreateBooking(&booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Booking created successfully"})
}

func (c *BookingController) GetAllBookings(ctx *gin.Context) {
	bookings, err := c.bookingService.GetAllBookings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Bookings fetched successfully", "bookings": bookings})
}

func (c *BookingController) GetBookingById(ctx *gin.Context) {
	id := ctx.Param("id")
	booking, err := c.bookingService.GetBookingById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Booking fetched successfully", "booking": booking})
}

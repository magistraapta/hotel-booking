package controller

import (
	"backend/internal/domain"
	"backend/internal/service"
	"backend/internal/shared"
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
		ctx.JSON(http.StatusInternalServerError, shared.ApiResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, shared.ApiResponse{Message: "Booking created successfully", Data: booking})
}

func (c *BookingController) GetAllBookings(ctx *gin.Context) {
	bookings, err := c.bookingService.GetAllBookings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.ApiResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Bookings fetched successfully", Data: bookings})
}

func (c *BookingController) GetBookingById(ctx *gin.Context) {
	id := ctx.Param("id")
	booking, err := c.bookingService.GetBookingById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.ApiResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Booking fetched successfully", Data: booking})
}

package controller

import (
	"backend/internal/domain"
	"backend/internal/service"
	"backend/internal/shared"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bookingService service.BookingService
}

func NewBookingController(bookingService service.BookingService) *BookingController {
	return &BookingController{bookingService: bookingService}
}

// CreateBooking godoc
// @Summary      Create a new booking
// @Description  Create a new hotel booking (Requires authentication)
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        booking  body      domain.CreateBookingRequest  true  "Booking information"
// @Success      201      {object}  shared.ApiResponse{data=domain.Booking}
// @Failure      400      {object}  shared.ErrorResponse
// @Failure      401      {object}  shared.ErrorResponse
// @Failure      500      {object}  shared.ErrorResponse
// @Router       /bookings [post]
func (c *BookingController) CreateBooking(ctx *gin.Context) {
	var booking domain.CreateBookingRequest
	ctx.ShouldBindJSON(&booking)
	err := c.bookingService.CreateBooking(&booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{Message: err.Error(), Path: ctx.Request.URL.Path, Status: http.StatusInternalServerError, Timestamp: time.Now().Format(time.RFC3339)})
		return
	}
	ctx.JSON(http.StatusCreated, shared.ApiResponse{Message: "Booking created successfully", Data: booking})
}

// GetAllBookings godoc
// @Summary      Get all bookings
// @Description  Retrieve a list of all bookings
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Success      200  {object}  shared.ApiResponse{data=[]domain.Booking}
// @Failure      500  {object}  shared.ErrorResponse
// @Router       /bookings [get]
func (c *BookingController) GetAllBookings(ctx *gin.Context) {
	bookings, err := c.bookingService.GetAllBookings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{Message: err.Error(), Path: ctx.Request.URL.Path, Status: http.StatusInternalServerError, Timestamp: time.Now().Format(time.RFC3339)})
		return
	}
	ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Bookings fetched successfully", Data: bookings})
}

// GetBookingById godoc
// @Summary      Get booking by ID
// @Description  Retrieve a specific booking by its ID
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Booking ID"
// @Success      200  {object}  shared.ApiResponse{data=domain.Booking}
// @Failure      400  {object}  shared.ErrorResponse
// @Failure      500  {object}  shared.ErrorResponse
// @Router       /bookings/{id} [get]
func (c *BookingController) GetBookingById(ctx *gin.Context) {
	id := ctx.Param("id")
	booking, err := c.bookingService.GetBookingById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{Message: err.Error(), Path: ctx.Request.URL.Path, Status: http.StatusInternalServerError, Timestamp: time.Now().Format(time.RFC3339)})
		return
	}
	ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Booking fetched successfully", Data: booking})
}

// GetBookingsByUserId godoc
// @Summary      Get bookings by user ID
// @Description  Retrieve all bookings for a specific user
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Param        user_id  path      string  true  "User ID"
// @Success      200      {object}  shared.ApiResponse{data=[]domain.Booking}
// @Failure      400      {object}  shared.ErrorResponse
// @Failure      500      {object}  shared.ErrorResponse
// @Router       /bookings/user/{user_id} [get]
func (c *BookingController) GetBookingsByUserId(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	bookings, err := c.bookingService.GetBookingsByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{Message: err.Error(), Path: ctx.Request.URL.Path, Status: http.StatusInternalServerError, Timestamp: time.Now().Format(time.RFC3339)})
		return
	}
	ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Bookings fetched successfully", Data: bookings})
}

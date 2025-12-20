package controller

import (
	"backend/internal/domain"
	"backend/internal/service"
	"backend/internal/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HotelController struct {
	hotelService service.HotelService
}

func NewHotelController(hotelService service.HotelService) *HotelController {
	return &HotelController{hotelService: hotelService}
}

// CreateHotel godoc
// @Summary      Create a new hotel
// @Description  Create a new hotel (Admin only)
// @Tags         Hotels
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        hotel  body      domain.Hotel  true  "Hotel information"
// @Success      201    {object}  shared.ApiResponse{data=domain.Hotel}
// @Failure      400    {object}  shared.ErrorResponse
// @Failure      401    {object}  shared.ErrorResponse
// @Failure      403    {object}  shared.ErrorResponse
// @Failure      500    {object}  shared.ErrorResponse
// @Router       /hotels [post]
func (c *HotelController) CreateHotel(ctx *gin.Context) {
	var hotel domain.Hotel
	ctx.ShouldBindJSON(&hotel)
	err := c.hotelService.CreateHotel(&hotel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.NewInternalServerErrorResponse(err.Error(), ctx.Request.URL.Path))
		return
	}
	ctx.JSON(http.StatusCreated, shared.NewCreatedResponse("Hotel created successfully", hotel, ctx.Request.URL.Path))
}

// GetAllHotels godoc
// @Summary      Get all hotels
// @Description  Retrieve a list of all hotels
// @Tags         Hotels
// @Accept       json
// @Produce      json
// @Success      200  {object}  shared.ApiResponse{data=[]domain.Hotel}
// @Failure      500  {object}  shared.ErrorResponse
// @Router       /hotels [get]
func (c *HotelController) GetAllHotels(ctx *gin.Context) {
	hotels, err := c.hotelService.GetAllHotels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.NewInternalServerErrorResponse(err.Error(), ctx.Request.URL.Path))
		return
	}
	ctx.JSON(http.StatusOK, shared.NewSuccessResponse("Hotels fetched successfully", hotels, http.StatusOK, ctx.Request.URL.Path))
}

// GetHotelById godoc
// @Summary      Get hotel by ID
// @Description  Retrieve a specific hotel by its ID
// @Tags         Hotels
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Hotel ID"
// @Success      200  {object}  shared.ApiResponse{data=domain.Hotel}
// @Failure      400  {object}  shared.ErrorResponse
// @Failure      500  {object}  shared.ErrorResponse
// @Router       /hotels/{id} [get]
func (c *HotelController) GetHotelById(ctx *gin.Context) {
	id := ctx.Param("id")
	hotel, err := c.hotelService.GetHotelById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.NewInternalServerErrorResponse(err.Error(), ctx.Request.URL.Path))
		return
	}
	ctx.JSON(http.StatusOK, shared.NewSuccessResponse("Hotel fetched successfully", hotel, http.StatusOK, ctx.Request.URL.Path))
}

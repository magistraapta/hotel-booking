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

func (c *HotelController) GetAllHotels(ctx *gin.Context) {
	hotels, err := c.hotelService.GetAllHotels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.NewInternalServerErrorResponse(err.Error(), ctx.Request.URL.Path))
		return
	}
	ctx.JSON(http.StatusOK, shared.NewSuccessResponse("Hotels fetched successfully", hotels, http.StatusOK, ctx.Request.URL.Path))
}

func (c *HotelController) GetHotelById(ctx *gin.Context) {
	id := ctx.Param("id")
	hotel, err := c.hotelService.GetHotelById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.NewInternalServerErrorResponse(err.Error(), ctx.Request.URL.Path))
		return
	}
	ctx.JSON(http.StatusOK, shared.NewSuccessResponse("Hotel fetched successfully", hotel, http.StatusOK, ctx.Request.URL.Path))
}

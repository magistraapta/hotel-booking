package controller

import (
	"backend/internal/domain"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HotelController struct {
	hotelService *service.HotelService
}

func NewHotelController(hotelService *service.HotelService) *HotelController {
	return &HotelController{hotelService: hotelService}
}

func (c *HotelController) CreateHotel(ctx *gin.Context) {
	var hotel domain.Hotel
	ctx.ShouldBindJSON(&hotel)
	err := c.hotelService.CreateHotel(&hotel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Hotel created successfully"})
}

func (c *HotelController) GetAllHotels(ctx *gin.Context) {
	hotels, err := c.hotelService.GetAllHotels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hotels fetched successfully", "hotels": hotels})
}

func (c *HotelController) GetHotelById(ctx *gin.Context) {
	id := ctx.Param("id")
	hotel, err := c.hotelService.GetHotelById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hotel fetched successfully", "hotel": hotel})
}

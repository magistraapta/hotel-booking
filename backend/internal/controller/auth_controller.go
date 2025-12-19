package controller

import (
	"backend/internal/domain"
	"backend/internal/service"
	"backend/internal/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginRequest domain.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, shared.ApiResponse{Error: err.Error()})
		return
	}
	loginResponse, err := c.authService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.ApiResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Login successful", Data: loginResponse})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var registerRequest domain.RegisterRequest
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, shared.ApiResponse{Error: err.Error()})
		return
	}
	err = c.authService.Register(&registerRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.ApiResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Registration successful"})
}

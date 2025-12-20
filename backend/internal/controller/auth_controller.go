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

// Login godoc
// @Summary      User login
// @Description  Authenticate user and return JWT tokens
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        loginRequest  body      domain.LoginRequest  true  "Login credentials"
// @Success      200           {object}  shared.ApiResponse{data=domain.LoginResponse}
// @Failure      400           {object}  shared.ErrorResponse
// @Failure      500           {object}  shared.ErrorResponse
// @Router       /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var loginRequest domain.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, shared.NewBadRequestResponse(err.Error(), ctx.Request.URL.Path))
		return
	}
	loginResponse, err := c.authService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.NewInternalServerErrorResponse(err.Error(), ctx.Request.URL.Path))
		return
	}
	ctx.JSON(http.StatusOK, shared.NewSuccessResponse("Login successful", loginResponse, http.StatusOK, ctx.Request.URL.Path))
}

// Register godoc
// @Summary      User registration
// @Description  Register a new user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        registerRequest  body      domain.RegisterRequest  true  "User registration data"
// @Success      200              {object}  shared.ApiResponse
// @Failure      400              {object}  shared.ErrorResponse
// @Failure      500              {object}  shared.ErrorResponse
// @Router       /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var registerRequest domain.RegisterRequest
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, shared.NewBadRequestResponse(err.Error(), ctx.Request.URL.Path))
		return
	}
	err = c.authService.Register(&registerRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.NewInternalServerErrorResponse(err.Error(), ctx.Request.URL.Path))
		return
	}

	ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Registration successful"})
}

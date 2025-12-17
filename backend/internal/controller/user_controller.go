package controller

import (
	"backend/internal/domain"
	"backend/internal/service"
	"backend/internal/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user domain.User
	ctx.ShouldBindJSON(&user)
	err := c.userService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.ApiResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, shared.ApiResponse{Message: "User created successfully", Data: user})
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.ApiResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Users fetched successfully", Data: users})
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.ApiResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "User fetched successfully", Data: user})
}

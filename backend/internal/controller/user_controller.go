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

// CreateUser godoc
// @Summary      Create a new user
// @Description  Register a new user account
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      domain.User  true  "User information"
// @Success      201   {object}  shared.ApiResponse{data=domain.User}
// @Failure      400   {object}  shared.ErrorResponse
// @Failure      500   {object}  shared.ErrorResponse
// @Router       /users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user domain.User
	ctx.ShouldBindJSON(&user)
	err := c.userService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.NewInternalServerErrorResponse(err.Error(), ctx.Request.URL.Path))
		return
	}
	ctx.JSON(http.StatusCreated, shared.NewCreatedResponse("User created successfully", user, ctx.Request.URL.Path))
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieve a list of all users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  shared.ApiResponse{data=[]domain.User}
// @Failure      500  {object}  shared.ErrorResponse
// @Router       /users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.NewInternalServerErrorResponse(err.Error(), ctx.Request.URL.Path))
		return
	}
	ctx.JSON(http.StatusOK, shared.NewSuccessResponse("Users fetched successfully", users, http.StatusOK, ctx.Request.URL.Path))
}

// GetUserById godoc
// @Summary      Get user by ID
// @Description  Retrieve a specific user by their ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  shared.ApiResponse{data=domain.User}
// @Failure      400  {object}  shared.ErrorResponse
// @Failure      500  {object}  shared.ErrorResponse
// @Router       /users/{id} [get]
func (c *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shared.NewInternalServerErrorResponse(err.Error(), ctx.Request.URL.Path))
		return
	}
	ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "User fetched successfully", Data: user})
}

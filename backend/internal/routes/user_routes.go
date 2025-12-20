package routes

import (
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/internal/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/", userController.CreateUser)
		// for testing
		// TestAdminRoute godoc
		// @Summary      Test admin route
		// @Description  Test endpoint to verify admin authentication (for testing only)
		// @Tags         Users
		// @Accept       json
		// @Produce      json
		// @Security     BearerAuth
		// @Success      200  {object}  shared.ApiResponse
		// @Failure      401  {object}  shared.ErrorResponse
		// @Failure      403  {object}  shared.ErrorResponse
		// @Router       /users/admin [get]
		userRouter.GET("/admin", middleware.RequireAdmin(), func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Admin route accessed successfully"})
		})
		// TestAuthRoute godoc
		// @Summary      Test authentication route
		// @Description  Test endpoint to verify user authentication (for testing only)
		// @Tags         Users
		// @Accept       json
		// @Produce      json
		// @Security     BearerAuth
		// @Success      200  {object}  shared.ApiResponse
		// @Failure      401  {object}  shared.ErrorResponse
		// @Router       /users/test [get]
		userRouter.GET("/test", middleware.RequireLogin(), func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Test route accessed successfully"})
		})
	}
}

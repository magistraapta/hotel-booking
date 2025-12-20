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
		userRouter.GET("/admin", middleware.RequireAdmin(), func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Admin route accessed successfully"})
		})
		userRouter.GET("/test", middleware.RequireLogin(), func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, shared.ApiResponse{Message: "Test route accessed successfully"})
		})
	}
}

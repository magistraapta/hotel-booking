package routes

import (
	"backend/internal/controller"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/register", authController.Register)
		authRouter.POST("/refresh", authController.RefreshToken)
	}
}

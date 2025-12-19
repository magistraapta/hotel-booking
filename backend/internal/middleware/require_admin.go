package middleware

import (
	"backend/internal/auth"
	"backend/internal/shared"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, shared.ApiResponse{Error: "Unauthorized"})
			ctx.Abort()
			return
		}
		const bearerPrefix = "Bearer "
		if len(authHeader) < len(bearerPrefix) || !strings.HasPrefix(authHeader, bearerPrefix) {
			ctx.JSON(http.StatusUnauthorized, shared.ApiResponse{Error: "Unauthorized"})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, bearerPrefix)

		user, err := auth.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, shared.ApiResponse{Error: "Unauthorized"})
			ctx.Abort()
			return
		}

		if !user.IsAdmin {
			ctx.JSON(http.StatusForbidden, shared.ApiResponse{Error: "Forbidden: only admins can access this resource"})
			ctx.Abort()
			return
		}

		ctx.Set("user", &user)
		ctx.Next()
	}
}

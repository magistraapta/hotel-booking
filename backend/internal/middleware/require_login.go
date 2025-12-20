package middleware

import (
	"backend/internal/auth"
	"backend/internal/shared"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RequireLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, shared.ErrorResponse{Message: "Unauthorized", Path: ctx.Request.URL.Path, Status: http.StatusUnauthorized, Timestamp: time.Now().Format(time.RFC3339)})
			ctx.Abort()
			return
		}
		const bearerPrefix = "Bearer "
		if len(authHeader) < len(bearerPrefix) || !strings.HasPrefix(authHeader, bearerPrefix) {
			ctx.JSON(http.StatusUnauthorized, shared.ErrorResponse{Message: "Unauthorized", Path: ctx.Request.URL.Path, Status: http.StatusUnauthorized, Timestamp: time.Now().Format(time.RFC3339)})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, bearerPrefix)

		user, err := auth.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, shared.ErrorResponse{Message: "Unauthorized", Path: ctx.Request.URL.Path, Status: http.StatusUnauthorized, Timestamp: time.Now().Format(time.RFC3339)})
			ctx.Abort()
			return
		}

		ctx.Set("user", &user)
		ctx.Next()
	}
}

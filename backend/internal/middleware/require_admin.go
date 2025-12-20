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
			ctx.JSON(http.StatusUnauthorized, shared.NewUnauthorizedResponse("Unauthorized", ctx.Request.URL.Path))
			ctx.Abort()
			return
		}
		const bearerPrefix = "Bearer "
		if len(authHeader) < len(bearerPrefix) || !strings.HasPrefix(authHeader, bearerPrefix) {
			ctx.JSON(http.StatusUnauthorized, shared.NewUnauthorizedResponse("Unauthorized", ctx.Request.URL.Path))
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, bearerPrefix)

		user, err := auth.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, shared.NewUnauthorizedResponse("Unauthorized", ctx.Request.URL.Path))
			ctx.Abort()
			return
		}

		if !user.IsAdmin {
			ctx.JSON(http.StatusForbidden, shared.NewForbiddenResponse("Forbidden: only admins can access this resource", ctx.Request.URL.Path))
			ctx.Abort()
			return
		}

		ctx.Set("user", &user)
		ctx.Next()
	}
}

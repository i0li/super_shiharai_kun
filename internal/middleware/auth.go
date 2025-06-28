package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/auth/jwt"
	"github.com/i0li/super_shiharai_kun/internal/infra/logger"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, apperr.ErrResponse{Error: apperr.ErrMsgNoAuthorizationHeader})
			c.Abort()
			return
		}

		var tokenStr string
		if authParts := strings.Split(authHeader, " "); len(authParts) == 2 && authParts[0] == "Bearer" {
			tokenStr = authParts[1]
		}

		userID, err := jwt.VerifyToken(tokenStr)
		if err != nil {
			logger.L.Error(apperr.Wrap(err).DetailMessage())
			c.JSON(http.StatusUnauthorized, apperr.ErrResponse{Error: apperr.ErrMsgInvalidToken})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/i0li/super_shiharai_kun/internal/apperr"

	jwtUtil "github.com/i0li/super_shiharai_kun/internal/util"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(apperr.ErrResRequireAuthHeader.Code, apperr.ErrResRequireAuthHeader)
			c.Abort()
			return
		}

		var tokenStr string
		if authParts := strings.Split(authHeader, " "); len(authParts) == 2 && authParts[0] == "Bearer" {
			tokenStr = authParts[1]
		}

		userID, err := jwtUtil.VerifyToken(tokenStr)
		if err != nil {
			c.JSON(apperr.ErrResInvalidToken.Code, apperr.ErrResInvalidToken)
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

package util

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
)

func GetUserID(c *gin.Context) (int64, error) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		return 0, apperr.Wrap(errors.New("userID not found in context"))
	}
	userID, ok := userIDVal.(int64)
	if !ok {
		return 0, apperr.Wrap(errors.New("failed to convert userID"))
	}

	return userID, nil
}

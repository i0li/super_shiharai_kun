package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/infra/logger"
	"github.com/i0li/super_shiharai_kun/internal/usecase"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(uc usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: uc}
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type loginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(apperr.ErrResBadRequest.Code, apperr.ErrResBadRequest)
		return
	}

	tokenStr, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, apperr.ErrUnauthorized) {
			c.JSON(apperr.ErrResUnauthorized.Code, apperr.ErrResUnauthorized)
			return
		}

		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(apperr.ErrResInternalServerError.Code, apperr.ErrResInternalServerError)
		return
	}

	c.JSON(http.StatusOK, loginResponse{Token: tokenStr})
}

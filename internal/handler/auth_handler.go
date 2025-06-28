package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/logger"
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
		c.JSON(http.StatusBadRequest, apperr.ErrResponse{Error: apperr.ErrMsgBadRequest})
		return
	}

	userID, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		if err == apperr.ErrUnauthorized {
			c.JSON(http.StatusUnauthorized, apperr.ErrResponse{Error: apperr.ErrMsgUnauthorized})
			return
		}

		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(http.StatusInternalServerError, apperr.ErrResponse{Error: apperr.ErrMsgInternalServerError})
		return
	}

	token, err := h.usecase.GenerateToken(userID)
	if err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(http.StatusInternalServerError, apperr.ErrResponse{Error: apperr.ErrMsgInternalServerError})
		return
	}

	c.JSON(http.StatusOK, loginResponse{Token: token})
}

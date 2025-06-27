package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appErr "github.com/i0li/super_shiharai_kun/internal/errors"
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
		c.JSON(http.StatusBadRequest, appErr.ErrResponse{Error: appErr.ErrMsgBadRequest})
		return
	}

	userID, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		if err == appErr.ErrUnauthorized {
			c.JSON(http.StatusUnauthorized, appErr.ErrResponse{Error: appErr.ErrMsgUnauthorized})
			return
		}

		c.JSON(http.StatusInternalServerError, appErr.ErrResponse{Error: appErr.ErrMsgInternalServerError})
		return
	}

	token, err := h.usecase.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appErr.ErrResponse{Error: appErr.ErrMsgInternalServerError})
		return
	}

	c.JSON(http.StatusOK, loginResponse{Token: token})
}

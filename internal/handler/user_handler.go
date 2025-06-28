package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/infra/logger"
	"github.com/i0li/super_shiharai_kun/internal/usecase"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

type registerRequest struct {
	CompanyName string `json:"company_name" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(http.StatusBadRequest, apperr.ErrResponse{Error: apperr.ErrMsgBadRequest})
		return
	}

	if err := h.usecase.RegisterUser(
		req.CompanyName,
		req.Name,
		req.Email,
		req.Password,
	); err != nil {
		if errors.Is(err, apperr.ErrEmailAlreadyExists) {
			logger.L.Error(apperr.Wrap(err).DetailMessage())
			c.JSON(http.StatusBadRequest, apperr.ErrResponse{Error: apperr.ErrMsgBadRequest})
			return
		}

		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(http.StatusInternalServerError, apperr.ErrResponse{Error: apperr.ErrMsgInternalServerError})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

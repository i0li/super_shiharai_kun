package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appErr "github.com/i0li/super_shiharai_kun/internal/errors"
	"github.com/i0li/super_shiharai_kun/internal/model"
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
		c.JSON(http.StatusBadRequest, appErr.ErrResponse{Error: appErr.ErrMsgBadRequest})
		return
	}

	user := &model.User{
		CompanyName: req.CompanyName,
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
	}

	if err := h.usecase.RegisterUser(user); err != nil {
		if err == appErr.ErrEmailAlreadyExists {
			c.JSON(http.StatusBadRequest, appErr.ErrResponse{Error: appErr.ErrMsgBadRequest})
			return
		}

		c.JSON(http.StatusInternalServerError, appErr.ErrResponse{Error: appErr.ErrMsgInternalServerError})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

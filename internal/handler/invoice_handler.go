package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/infra/logger"
	"github.com/i0li/super_shiharai_kun/internal/types"
	"github.com/i0li/super_shiharai_kun/internal/usecase"
	"github.com/shopspring/decimal"
)

type InvoiceHandler struct {
	usecase usecase.InvoiceUsecase
}

func NewInvoiceHandler(uc usecase.InvoiceUsecase) *InvoiceHandler {
	return &InvoiceHandler{usecase: uc}
}

type createRequest struct {
	PaymentAmount  decimal.Decimal `json:"payment_amount" binding:"required"`
	PaymentDueDate types.Date      `json:"payment_due_date" binding:"required"`
}
type createResponse struct {
	InvoiceID int64 `json:"invoice_id"`
}

func (h *InvoiceHandler) CreateInvoice(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		logger.L.Error(apperr.Wrap(errors.New("userID not found from context")).DetailMessage())
		c.JSON(http.StatusInternalServerError, apperr.ErrResponse{Error: apperr.ErrMsgInternalServerError})
		return
	}
	userID, ok := userIDVal.(int64)
	if !ok {
		logger.L.Error(apperr.Wrap(errors.New("failed to convert context's userID")).DetailMessage())
		c.JSON(http.StatusInternalServerError, apperr.ErrResponse{Error: apperr.ErrMsgInternalServerError})
		return
	}

	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(http.StatusBadRequest, apperr.ErrResponse{Error: apperr.ErrMsgBadRequest})
		return
	}

	invoiceID, err := h.usecase.Create(userID, req.PaymentAmount, req.PaymentDueDate.Time)
	if err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(http.StatusInternalServerError, apperr.ErrResponse{Error: apperr.ErrMsgInternalServerError})
		return
	}

	c.JSON(http.StatusOK, createResponse{InvoiceID: invoiceID})
}

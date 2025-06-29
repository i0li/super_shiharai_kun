package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/infra/logger"
	"github.com/i0li/super_shiharai_kun/internal/usecase"
	"github.com/i0li/super_shiharai_kun/internal/util"
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
	PaymentDueDate string          `json:"payment_due_date" binding:"required"`
}
type createResponse struct {
	InvoiceID int64 `json:"invoice_id"`
}

func (h *InvoiceHandler) CreateInvoice(c *gin.Context) {
	userID, err := util.GetUserID(c)
	if err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(apperr.ErrResInternalServerError.Code, apperr.ErrResInternalServerError)
		return
	}

	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(apperr.ErrResBadRequest.Code, apperr.ErrResBadRequest)
		return
	}

	paymentDueDate, err := time.Parse("2006-01-02", req.PaymentDueDate)
	if err != nil {
		c.JSON(apperr.ErrResBadRequest.Code, apperr.ErrResBadRequest)
		return
	}

	invoiceID, err := h.usecase.Create(userID, req.PaymentAmount, paymentDueDate)
	if err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(apperr.ErrResInternalServerError.Code, apperr.ErrResInternalServerError)
		return
	}

	c.JSON(http.StatusOK, createResponse{InvoiceID: invoiceID})
}

type FindPayableInvoicesInPeriodRequest struct {
	StartDate string `form:"start_date" binding:"required"`
	EndDate   string `form:"end_date" binding:"required"`
}
type FindPayableInvoicesInPeriodResponse struct {
	Invoices []InvoiceItem `json:"invoices"`
}
type InvoiceItem struct {
	ID             int64           `json:"id"`
	IssueDate      string          `json:"issue_date"`
	PaymentAmount  decimal.Decimal `json:"payment_amount"`
	Fee            decimal.Decimal `json:"fee"`
	FeeRate        decimal.Decimal `json:"fee_rate"`
	TaxAmount      decimal.Decimal `json:"tax_amount"`
	TaxRate        decimal.Decimal `json:"tax_rate"`
	TotalAmount    decimal.Decimal `json:"total_amount"`
	PaymentDueDate string          `json:"payment_due_date"`
}

func (h *InvoiceHandler) FindPayableInvoicesInPeriod(c *gin.Context) {
	userID, err := util.GetUserID(c)
	if err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(apperr.ErrResInternalServerError.Code, apperr.ErrResInternalServerError)
		return
	}

	var req FindPayableInvoicesInPeriodRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(apperr.ErrResBadRequest.Code, apperr.ErrResBadRequest)
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(apperr.ErrResBadRequest.Code, apperr.ErrResBadRequest)
		return
	}
	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(apperr.ErrResBadRequest.Code, apperr.ErrResBadRequest)
		return
	}

	invoices, err := h.usecase.FindPayableInvoicesInPeriod(userID, startDate, endDate)
	if err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
		c.JSON(apperr.ErrResInternalServerError.Code, apperr.ErrResInternalServerError)
		return
	}

	var invoiceItems []InvoiceItem
	for _, invoice := range invoices {
		invoiceItems = append(invoiceItems, InvoiceItem{
			ID:             invoice.ID,
			IssueDate:      invoice.IssueDate.Format("2006-01-02"),
			PaymentAmount:  invoice.PaymentAmount,
			Fee:            invoice.Fee,
			FeeRate:        invoice.FeeRate,
			TaxAmount:      invoice.TaxAmount,
			TaxRate:        invoice.TaxRate,
			TotalAmount:    invoice.TotalAmount,
			PaymentDueDate: invoice.PaymentDueDate.Format("2006-01-02"),
		})
	}

	c.JSON(http.StatusOK, FindPayableInvoicesInPeriodResponse{
		Invoices: invoiceItems,
	})
}

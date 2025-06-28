package usecase

import (
	"time"

	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/model"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"github.com/shopspring/decimal"
)

var (
	InvoiceFeeRate, _ = decimal.NewFromString("0.04")
	InvoiceTaxRate, _ = decimal.NewFromString("0.10")
)

type InvoiceUsecase interface {
	Create(userID int64, paymentAmount decimal.Decimal, paymentDueDate time.Time) (int64, error)
}

type invoiceUsecase struct {
	repo repository.InvoiceRepository
}

func NewInvoiceUsecase(invoiceRepo repository.InvoiceRepository) InvoiceUsecase {
	return &invoiceUsecase{repo: invoiceRepo}
}

func (uc *invoiceUsecase) Create(userID int64, paymentAmount decimal.Decimal, paymentDueDate time.Time) (int64, error) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	invoice := &model.Invoice{
		UserID:         userID,
		IssueDate:      time.Now().In(jst),
		PaymentAmount:  paymentAmount,
		Fee:            decimal.Zero,
		FeeRate:        InvoiceFeeRate,
		TaxAmount:      decimal.Zero,
		TaxRate:        InvoiceTaxRate,
		TotalAmount:    decimal.Zero,
		PaymentDueDate: paymentDueDate,
	}
	invoice.CalculateTotalAmount()

	if err := uc.repo.Create(invoice); err != nil {
		return 0, apperr.Wrap(err)
	}

	return invoice.ID, nil
}

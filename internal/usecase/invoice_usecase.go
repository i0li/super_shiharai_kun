package usecase

import (
	"time"

	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/domain"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"github.com/shopspring/decimal"
)

var (
	InvoiceFeeRate, _ = decimal.NewFromString("0.04")
	InvoiceTaxRate, _ = decimal.NewFromString("0.10")
)

type InvoiceUsecase interface {
	Create(
		userID int64,
		paymentAmount decimal.Decimal,
		paymentDueDate time.Time,
	) (int64, error)

	FindPayableInvoicesInPeriod(
		userID int64,
		startDate, endDate time.Time,
	) ([]*domain.Invoice, error)
}

type invoiceUsecase struct {
	repo repository.InvoiceRepository
}

func NewInvoiceUsecase(invoiceRepo repository.InvoiceRepository) InvoiceUsecase {
	return &invoiceUsecase{repo: invoiceRepo}
}

func (uc *invoiceUsecase) Create(userID int64, paymentAmount decimal.Decimal, paymentDueDate time.Time) (int64, error) {
	invoice := &domain.Invoice{
		UserID:         userID,
		IssueDate:      time.Now(),
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

func (uc *invoiceUsecase) FindPayableInvoicesInPeriod(
	userID int64,
	startDate, endDate time.Time,
) ([]*domain.Invoice, error) {
	invoices, err := uc.repo.FindByPaymentDueDatePeriod(userID, startDate, endDate)
	if err != nil {
		return nil, apperr.Wrap(err)
	}

	return invoices, nil
}

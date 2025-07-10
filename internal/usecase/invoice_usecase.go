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

type Pagination struct {
	Limit  int
	Offset int
	Total  int64
}

type InvoiceOutput struct {
	ID             int64
	IssueDate      time.Time
	PaymentAmount  decimal.Decimal
	Fee            decimal.Decimal
	FeeRate        decimal.Decimal
	TaxAmount      decimal.Decimal
	TaxRate        decimal.Decimal
	TotalAmount    decimal.Decimal
	PaymentDueDate time.Time
}

type InvoiceUsecase interface {
	Create(
		userID int64,
		paymentAmount decimal.Decimal,
		paymentDueDate time.Time,
	) (int64, error)

	FindPayableInvoicesInPeriod(
		userID int64,
		startDate, endDate time.Time,
		limit int,
		offset int,
	) ([]*InvoiceOutput, Pagination, error)
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
	limit int,
	offset int,
) ([]*InvoiceOutput, Pagination, error) {
	invoices, err := uc.repo.FindByPaymentDueDatePeriod(userID, startDate, endDate, limit, offset)
	if err != nil {
		return nil, Pagination{}, apperr.Wrap(err)
	}

	total, err := uc.repo.CountByPaymentDueDatePeriod(userID, startDate, endDate)
	if err != nil {
		return nil, Pagination{}, apperr.Wrap(err)
	}

	var invoiceOutputs []*InvoiceOutput
	for _, invoice := range invoices {
		invoiceOutputs = append(invoiceOutputs, &InvoiceOutput{
			ID:             invoice.ID,
			IssueDate:      invoice.IssueDate,
			PaymentAmount:  invoice.PaymentAmount,
			Fee:            invoice.Fee,
			FeeRate:        invoice.FeeRate,
			TaxAmount:      invoice.TaxAmount,
			TaxRate:        invoice.TaxRate,
			TotalAmount:    invoice.TotalAmount,
			PaymentDueDate: invoice.PaymentDueDate,
		})
	}

	pagination := Pagination{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}

	return invoiceOutputs, pagination, nil
}
